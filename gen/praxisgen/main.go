package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/rightscale/rsc/gen"
	"github.com/rightscale/rsc/gen/writers"
)

// Data structure used to load content of index.json files
type Index map[string]map[string]map[string]interface{}

func main() {
	// 1. Parse command line arguments
	curDir, err := os.Getwd()
	check(err)
	var metadataDirVal = flag.String("metadata", curDir,
		"Path to directory(ies) containig metadata files (index.json, etc.). Multiple directories can be specified using a coma separated list.")
	var destDirVal = flag.String("output", curDir,
		"Path to output file")
	var pkgName = flag.String("pkg", "", "Name of generated package, e.g. \"rsapi16\"")
	var targetVersion = flag.String("target", "",
		"Target version, only generate code for this version.\nIf this option is specified then the generated code lives directly under package <pkg>, otherwise it lives under <pkg>.<version>.")
	var clientName = flag.String("client", "", "Name of API client go struct, e.g. \"Api16\".")
	flag.Parse()

	var metadataDirs = strings.Split(*metadataDirVal, ",")
	for _, metadataDir := range metadataDirs {
		if stat, err := os.Stat(metadataDir); err != nil || !stat.IsDir() {
			check(fmt.Errorf("%s is not a valid directory.", metadataDir))
		}
	}

	destDir := *destDirVal
	if stat, _ := os.Stat(destDir); stat != nil && !stat.IsDir() {
		check(fmt.Errorf("%s is not a valid directory.", destDir))
	}

	if *pkgName == "" {
		check(fmt.Errorf("-pkg option is required."))
	}

	if *clientName == "" {
		check(fmt.Errorf("-client option is required."))
	}

	var indeces = make(map[string]Index, len(metadataDirs)) // Index content mapped by directory path
	for _, metadataDir := range metadataDirs {
		var indexFile = path.Join(metadataDir, "index.json")
		indexContent, err := loadFile(indexFile)
		check(err)
		var index Index
		err = json.Unmarshal(indexContent, &index)
		if err != nil {
			check(fmt.Errorf("Cannot unmarshal JSON read from '%s': %s", indexFile, err.Error()))
		}
		indeces[metadataDir] = index
	}

	// 2. Analyze
	var descriptors = make(map[string]*gen.ApiDescriptor) // descriptors indexed by version
	for dirPath, index := range indeces {
		for version, resources := range index {
			if len(*targetVersion) > 0 {
				if version != *targetVersion {
					fmt.Fprintf(os.Stderr, "Skipping version \"%s\"\n", version)
					continue
				}
			}
			var apiResources = make(map[string]map[string]interface{}) // Resource properties indexed by name indexed by resource name
			for name, resource := range resources {
				// Skip built-in resources (?)
				if strings.HasSuffix(name, " (*)") {
					continue
				}
				var fileName = fmt.Sprintf("%s.json", resource["controller"])
				var resourcePath = path.Join(dirPath, version, "resources", fileName)
				var resourceData map[string]interface{}
				if err := unmarshal(resourcePath, &resourceData); err != nil {
					check(fmt.Errorf("Failed to unmarshal content of file %s: %s", resourcePath, err.Error()))
				}
				apiResources[name] = resourceData
			}

			var apiTypes = make(map[string]map[string]interface{}) // Type properties indexed by name indexed by type name
			var typesDir = path.Join(dirPath, version, "types")
			files, err := ioutil.ReadDir(typesDir)
			if err != nil {
				check(fmt.Errorf("Failed to load types: %s", err.Error()))
			}
			for _, file := range files {
				var typeData map[string]interface{}
				if err := unmarshal(path.Join(typesDir, file.Name()), &typeData); err != nil {
					check(fmt.Errorf("Failed to unmarshal content of file %s: %s", path.Join(typesDir, file.Name()), err.Error()))
				}
				var typeName, ok = typeData["name"]
				if !ok {
					check(fmt.Errorf("Missing \"name\" key for type defined in %s", path.Join(typesDir, file.Name())))
				}
				apiTypes[typeName.(string)] = typeData
			}
			var analyzer = NewApiAnalyzer(version, *clientName, apiResources, apiTypes)
			d, err := analyzer.Analyze()
			check(err)
			if existing, ok := descriptors[version]; ok {
				err := existing.Merge(d)
				if err != nil {
					check(fmt.Errorf("Conflict between multiple metadata directory: %s", err.Error()))
				}
			} else {
				descriptors[version] = d
			}
		}
	}

	// 3. Write code
	var genClients []string
	var genMetadata []string
	for version, descriptor := range descriptors {
		var pkg string
		if len(*targetVersion) == 0 {
			pkg = toPackageName(version)
		}
		os.MkdirAll(path.Join(destDir, pkg), 0755)
		var clientPath = path.Join(destDir, pkg, "codegen_client.go")
		var metadataPath = path.Join(destDir, pkg, "codegen_metadata.go")
		check(generateClient(descriptor, clientPath, *pkgName))
		check(generateMetadata(descriptor, metadataPath, *pkgName))
		genClients = append(genClients, clientPath)
		genMetadata = append(genMetadata, metadataPath)
	}

	// 4. Say something...
	for i := 0; i < len(genClients); i++ {
		fmt.Printf("%s\n%s\n", genClients[i], genMetadata[i])
	}
}

// Helper method that loads file file content and (JSON) unmarshals it into target
func unmarshal(path string, target *map[string]interface{}) error {
	content, err := loadFile(path)
	if err != nil {
		return fmt.Errorf("Failed to load media type JSON from '%s': %s", path, err.Error())
	}
	err = json.Unmarshal(content, target)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal JSON read from '%s': %s", path, err.Error())
	}
	return nil
}

// Convert version number in index.json to go package name
// "1.6" => "v1_6"
func toPackageName(version string) string {
	if version == "unversioned" {
		return "v0"
	}
	var parts = strings.Split(version, ".")
	var i = 1
	var p = parts[len(parts)-i]
	for p == "0" && i <= len(parts) {
		i += 1
		p = parts[len(parts)-i]
	}
	version = strings.Join(parts, "_")
	return fmt.Sprintf("v%s", version)
}

// Generate API client code, drives the code writer.
func generateClient(descriptor *gen.ApiDescriptor, codegen, pkg string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewClientWriter()
	if err != nil {
		return err
	}
	check(c.WriteHeader(pkg, f))
	for _, name := range descriptor.ResourceNames {
		resource := descriptor.Resources[name]
		c.WriteResourceHeader(name, f)
		check(c.WriteResource(resource, f))
	}
	c.WriteTypeSectionHeader(f)
	for _, name := range descriptor.TypeNames {
		t := descriptor.Types[name]
		c.WriteType(t, f)
	}
	f.Close()
	o, err := exec.Command("go", "fmt", codegen).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to format generated client code:\n%s", o)
	}
	return nil
}

// Generate API metadata, drives the metadata writer.
func generateMetadata(descriptor *gen.ApiDescriptor, codegen, pkg string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewMetadataWriter()
	if err != nil {
		return err
	}
	check(c.WriteHeader(pkg, f))
	check(c.WriteMetadata(descriptor, f))
	f.Close()
	o, err := exec.Command("go", "fmt", codegen).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to format generated metadata code:\n%s", o)
	}
	return nil
}

// Helper function that reads content from given file
func loadFile(file string) ([]byte, error) {
	if _, err := os.Stat(file); err != nil {
		return nil, fmt.Errorf("Cannot find '%s'", file)
	}
	js, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Cannot read '%s': %s", file, err.Error())
	}
	return js, nil
}

// Panic if error is not nil
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "praxisgen: %s\n", err.Error())
		os.Exit(1)
	}
}
