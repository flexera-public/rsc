package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/rightscale/rsc/gen"
	"github.com/rightscale/rsc/gen/writers"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// 1. Parse command line arguments
	curDir, err := os.Getwd()
	kingpin.FatalIfError(err, "")
	metadataDirVal := flag.String("metadata", curDir,
		"Path to directory containig metadata files (api_data.json and attributes.json)")
	destDirVal := flag.String("output", curDir,
		"Path to output file")
	flag.Parse()

	metadataDir := *metadataDirVal
	if stat, _ := os.Stat(metadataDir); !stat.IsDir() {
		kingpin.Fatalf("%s is not a valid directory", metadataDir)
	}

	destDir := *destDirVal
	if stat, _ := os.Stat(destDir); stat != nil && !stat.IsDir() {
		kingpin.Fatalf("%s is not a valid directory", destDir)
	}

	apiDataFile := path.Join(metadataDir, "api_data.json")
	var apiData map[string]interface{}
	apiDataText, err := loadFile(apiDataFile)
	kingpin.FatalIfError(err, "")
	err = json.Unmarshal(apiDataText, &apiData)
	if err != nil {
		kingpin.Fatalf("Cannot unmarshal JSON read from '%s': %s", apiDataFile, err)
	}

	attributesFile := path.Join(metadataDir, "attributes.json")
	var attributes map[string]string
	attributesText, err := loadFile(attributesFile)
	kingpin.FatalIfError(err, "")
	err = json.Unmarshal(attributesText, &attributes)
	if err != nil {
		kingpin.Fatalf("Cannot unmarshal JSON read from '%s': %s", attributesFile, err)
	}

	// 2. Analyze
	analyzer := NewAPIAnalyzer(apiData, attributes)
	descriptor := analyzer.Analyze()

	// 3. Write codegen_client.go
	var clientPath = path.Join(destDir, "codegen_client.go")
	kingpin.FatalIfError(generateClient(descriptor, clientPath), "")

	// 4. Write codegen_metadata.go
	var metadataPath = path.Join(destDir, "codegen_metadata.go")
	kingpin.FatalIfError(generateMetadata(descriptor, metadataPath), "")

	// 5. Say something...
	fmt.Printf("%s\n%s\n", clientPath, metadataPath)
}

// Generate API client code, drives the code writer.
func generateClient(descriptor *gen.APIDescriptor, codegen string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewClientWriter()
	if err != nil {
		return err
	}
	kingpin.FatalIfError(c.WriteHeader("cm15", "1.5", false /*needTime*/, true /*needJSON*/, f), "")
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
func generateMetadata(descriptor *gen.APIDescriptor, codegen string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := writers.NewMetadataWriter()
	if err != nil {
		return err
	}
	kingpin.FatalIfError(c.WriteHeader("cm15", f), "")
	kingpin.FatalIfError(c.WriteMetadata(descriptor, f), "")
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
		return nil, fmt.Errorf("Cannot read '%s': %s", file, err)
	}
	return js, nil
}
