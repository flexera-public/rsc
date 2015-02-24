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

func main() {
	// 1. Parse command line arguments
	curDir, err := os.Getwd()
	check(err)
	metadataDirVal := flag.String("metadata", curDir,
		"Path to directory containig metadata files (index.json, etc.)")
	destDirVal := flag.String("output", curDir,
		"Path to output file")
	pkgName := flag.String("pkg", "", "Name of generated package, e.g. \"rsapi16\"")
	flag.Parse()

	metadataDir := *metadataDirVal
	if stat, _ := os.Stat(metadataDir); !stat.IsDir() {
		check(fmt.Errorf("%s is not a valid directory.", metadataDir))
	}

	destDir := *destDirVal
	if stat, _ := os.Stat(destDir); stat != nil && !stat.IsDir() {
		check(fmt.Errorf("%s is not a valid directory.", destDir))
	}

	if *pkgName == "" {
		check(fmt.Errorf("-pkg option is required."))
	}

	indexFile := path.Join(metadataDir, "index.json")
	var index map[string]map[string]map[string]string
	indexContent, err := loadFile(indexFile)
	check(err)
	err = json.Unmarshal(indexContent, &index)
	if err != nil {
		check(fmt.Errorf("Cannot unmarshal JSON read from '%s': %s", indexFile, err.Error()))
	}

	// 2. Analyze
	var descriptors = make(map[string]*gen.ApiDescriptor) // descriptors indexed by version
	for version, resources := range index {
		var apiResources = make(map[string]map[string]interface{}) // Resource properies indexed by name indexed by resource name
		var apiTypes = make(map[string]map[string]interface{})     // Resource properies indexed by name indexed by resource name
		for name, resource := range resources {
			// Skip built-in resources (?)
			if strings.HasSuffix(name, " (*)") {
				continue
			}

			// Parse resource
			var fileName = fmt.Sprintf("%s::%s.json", toModuleName(version), name)
			var resourcePath = path.Join(metadataDir, version, "resources", fileName)
			var resourceData map[string]interface{}
			check(unmarshal(resourcePath, &resourceData))
			apiResources[name] = resourceData

			// Parse media type
			var media_type, ok = resource["media_type"]
			if !ok {
				check(fmt.Errorf("No media type defined for resource %s", name))
			}
			fileName = fmt.Sprintf("%s.json", resource["media_type"])
			var typePath = path.Join(metadataDir, version, "types", fileName)
			var typeData map[string]interface{}
			check(unmarshal(typePath, &typeData))
			apiTypes[resource["media_type"]] = typeData
		}
		var analyzer = NewApiAnalyzer(apiResources, apiTypes)
		descriptors[version] = analyzer.Analyze()
	}

	// 3. Write code.go
	for version, descriptor := range descriptors {
		os.MkdirAll(path.Join(destDir, version), 0755)
		check(generateClient(descriptor, path.Join(destDir, version, "codegen_client.go")))
		check(generateMetadata(descriptor, path.Join(destDir, version, "codegen_metadata.go")))
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

// Convert version number in index.json to ruby module name used by resource filenames
func toModuleName(version string) string {
	if version == "unversioned" {
		return "V0"
	}
	return fmt.Sprintf("V%s", strings.Replace(version, ".", "_", -1))
}

// Generate API client code, drives the code writer.
func generateClient(descriptor *gen.ApiDescriptor, codegen string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewClientWriter()
	if err != nil {
		return err
	}
	check(c.WriteHeader(f))
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
func generateMetadata(descriptor *gen.ApiDescriptor, codegen string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewMetadataWriter()
	if err != nil {
		return err
	}
	check(c.WriteHeader("rsapi15", f))
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
		fmt.Fprintf(os.Stderr, "api15gen: %s\n", err.Error())
		os.Exit(1)
	}
}
