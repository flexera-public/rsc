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

	"bitbucket.org/pkg/inflect"

	"github.com/rightscale/rsc/gen"
	"github.com/rightscale/rsc/gen/writers"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Index is the data structure used to load content of index.json files.
type Index map[string]map[string]map[string]interface{}

func main() {
	// 1. Parse command line arguments
	curDir, err := os.Getwd()
	kingpin.FatalIfError(err, "")
	metadataDirVal := flag.String("metadata", curDir,
		"Path to directory(ies) containig metadata files (index.json, etc.). Multiple directories can be specified using a coma separated list.")
	destDirVal := flag.String("output", curDir,
		"Path to output file")
	pkgName := flag.String("pkg", "", "Name of generated package, e.g. \"rsapi16\"")
	targetVersion := flag.String("target", "", "Version of API to generate code for")
	clientName := flag.String("client", "", "Name of API client go struct, e.g. \"API16\".")
	tool := flag.String("tool", "rsc", "Tool or library for which to generate code, supported values are 'rsc' or 'angular'")
	flag.Parse()

	metadataDirs := strings.Split(*metadataDirVal, ",")
	for _, metadataDir := range metadataDirs {
		if stat, err := os.Stat(metadataDir); err != nil || !stat.IsDir() {
			kingpin.Fatalf("%s is not a valid directory.", metadataDir)
		}
	}

	destDir := *destDirVal
	if stat, _ := os.Stat(destDir); stat != nil && !stat.IsDir() {
		kingpin.Fatalf("%s is not a valid directory.", destDir)
	}

	if *targetVersion == "" {
		kingpin.Fatalf("-target option is required.")
	}

	if *tool == "rsc" {
		if *pkgName == "" {
			kingpin.Fatalf("-pkg option is required.")
		}

		if *clientName == "" {
			kingpin.Fatalf("-client option is required.")
		}
	}

	indeces := make(map[string]Index, len(metadataDirs)) // Index content mapped by directory path
	for _, metadataDir := range metadataDirs {
		indexFile := path.Join(metadataDir, "index.json")
		indexContent, err := loadFile(indexFile)
		kingpin.FatalIfError(err, "")
		var index Index
		err = json.Unmarshal(indexContent, &index)
		if err != nil {
			kingpin.Fatalf("Cannot unmarshal JSON read from '%s': %s", indexFile, err)
		}
		indeces[metadataDir] = index
	}

	// 2. Analyze
	descriptors := make(map[string]*gen.APIDescriptor) // descriptors indexed by version
	for dirPath, index := range indeces {
		for version, resources := range index {
			if len(*targetVersion) > 0 {
				if version != *targetVersion {
					fmt.Fprintf(os.Stderr, "Skipping version \"%s\"\n", version)
					continue
				}
			}
			apiResources := make(map[string]map[string]interface{}) // Resource properties indexed by name indexed by resource name
			for name, resource := range resources {
				// Skip built-in resources (?)
				if strings.HasSuffix(name, " (*)") {
					continue
				}
				fileName := strings.Replace(fmt.Sprintf("%s.json", resource["controller"]), "::", "-", -1)
				resourcePath := path.Join(dirPath, version, "resources", fileName)
				var resourceData map[string]interface{}
				if err := unmarshal(resourcePath, &resourceData); err != nil {
					kingpin.Fatalf("Failed to unmarshal content of file %s: %s", resourcePath, err)
				}
				apiResources[name] = resourceData
			}

			apiTypes := make(map[string]map[string]interface{}) // Type properties indexed by name indexed by type name
			typesDir := path.Join(dirPath, version, "types")
			files, _ := ioutil.ReadDir(typesDir)
			for _, file := range files {
				var typeData map[string]interface{}
				if err := unmarshal(path.Join(typesDir, file.Name()), &typeData); err != nil {
					kingpin.Fatalf("Failed to unmarshal content of file %s: %s", path.Join(typesDir, file.Name()), err)
				}
				typeName, ok := typeData["name"]
				if !ok {
					kingpin.Fatalf("Missing \"name\" key for type defined in %s", path.Join(typesDir, file.Name()))
				}
				apiTypes[typeName.(string)] = typeData
			}
			analyzer := NewAPIAnalyzer(version, *clientName, apiResources, apiTypes)
			d, err := analyzer.Analyze()
			kingpin.FatalIfError(err, "")
			if existing, ok := descriptors[version]; ok {
				err := existing.Merge(d)
				if err != nil {
					kingpin.Fatalf("Conflict between multiple metadata directory: %s", err)
				}
			} else {
				descriptors[version] = d
			}
		}
	}

	// 3. Write code
	var generated []string
	for version, descriptor := range descriptors {
		var pkg string
		if len(*targetVersion) == 0 {
			pkg = toPackageName(version)
		}
		os.MkdirAll(path.Join(destDir, pkg), 0755)
		switch *tool {
		case "rsc":
			clientPath := path.Join(destDir, pkg, "codegen_client.go")
			metadataPath := path.Join(destDir, pkg, "codegen_metadata.go")
			kingpin.FatalIfError(generateClient(*targetVersion, descriptor, clientPath, *pkgName), "")
			kingpin.FatalIfError(generateMetadata(descriptor, metadataPath, *pkgName), "")
			generated = append(generated, clientPath)
			generated = append(generated, metadataPath)
		case "angular":
			pkgPath := path.Join(destDir, pkg)
			files, err := generateAngular(descriptor, pkgPath)
			kingpin.FatalIfError(err, "")
			generated = append(generated, files...)
		default:
			kingpin.Fatalf("Invalid tool '%s', supported clients are 'rsc' and 'angular'", *tool)
		}
	}

	// 4. Say something...
	for i := 0; i < len(generated); i++ {
		fmt.Printf("%s\n", generated[i])
	}
}

// Helper method that loads file file content and (JSON) unmarshals it into target
func unmarshal(path string, target *map[string]interface{}) error {
	content, err := loadFile(path)
	if err != nil {
		return fmt.Errorf("Failed to load media type JSON from '%s': %s", path, err)
	}
	err = json.Unmarshal(content, target)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal JSON read from '%s': %s", path, err)
	}
	return nil
}

// Convert version number in index.json to go package name
// "1.6" => "v1_6"
func toPackageName(version string) string {
	if version == "unversioned" {
		return "v0"
	}
	parts := strings.Split(version, ".")
	i := 1
	p := parts[len(parts)-i]
	for p == "0" && i <= len(parts) {
		i++
		p = parts[len(parts)-i]
	}
	version = strings.Join(parts, "_")
	return fmt.Sprintf("v%s", version)
}

// Generate API client code, drives the code writer.
func generateClient(version string, descriptor *gen.APIDescriptor, codegen, pkg string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewClientWriter()
	if err != nil {
		return err
	}
	kingpin.FatalIfError(c.WriteHeader(pkg, version, descriptor.NeedTime, descriptor.NeedJSON, f), "")
	for _, name := range descriptor.ResourceNames {
		resource := descriptor.Resources[name]
		c.WriteResourceHeader(name, f)
		kingpin.FatalIfError(c.WriteResource(resource, f), "")
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
func generateMetadata(descriptor *gen.APIDescriptor, codegen, pkg string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewMetadataWriter()
	if err != nil {
		return err
	}
	kingpin.FatalIfError(c.WriteHeader(pkg, f), "")
	kingpin.FatalIfError(c.WriteMetadata(descriptor, f), "")
	f.Close()
	o, err := exec.Command("go", "fmt", codegen).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to format generated metadata code:\n%s", o)
	}
	return nil
}

// Generate API metadata, drives the metadata writer.
func generateAngular(descriptor *gen.APIDescriptor, pkgDir string) ([]string, error) {
	var files []string
	for _, name := range descriptor.ResourceNames {
		res := descriptor.Resources[name]
		codegen := path.Join(pkgDir, inflect.Underscore(name)+".js")
		f, err := os.Create(codegen)
		if err != nil {
			return files, err
		}
		c, err := writers.NewAngularWriter()
		if err != nil {
			return files, err
		}
		kingpin.FatalIfError(c.WriteResource(res, f), "")
		f.Close()
		files = append(files, codegen)
	}
	return files, nil
}

// Helper function that reads content from given file
func loadFile(file string) ([]byte, error) {
	if _, err := os.Stat(file); err != nil {
		return nil, fmt.Errorf("Cannot find '%s'", file)
	}
	js, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Cannot read '%s': %s", file, err)
	}
	return js, nil
}
