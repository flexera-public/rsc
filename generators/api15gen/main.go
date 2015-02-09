package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	// 1. Parse command line arguments
	curDir, err := os.Getwd()
	check(err)
	metadataDirVal := flag.String("metadata", curDir,
		"Path to directory containig metadata files (api_data.json and attributes.json)")
	destDirVal := flag.String("output", curDir,
		"Path to output file")
	flag.Parse()

	metadataDir := *metadataDirVal
	if stat, _ := os.Stat(metadataDir); !stat.IsDir() {
		check(fmt.Errorf("%s is not a valid directory", metadataDir))
	}

	destDir := *destDirVal
	if stat, _ := os.Stat(destDir); !stat.IsDir() {
		check(fmt.Errorf("%s is not a valid directory", destDir))
	}

	apiDataFile := path.Join(metadataDir, "api_data.json")
	var apiData map[string]interface{}
	i := interface{}(apiData)
	check(loadJson(apiDataFile, &i))

	attributesFile := path.Join(metadataDir, "attributes.json")
	var attributes map[string]string
	i = interface{}(attributes)
	check(loadJson(attributesFile, &i))

	// 2. Analyze
	analyzer := NewApiAnalyzer(apiData, attributes)
	descriptor := analyzer.Analyze()

	// 3. Write codegen.go
	check(NewClientGenerator("codegen.go").Generate(descriptor, destDir))

	// 4. Write codegen_cmds.go
	check(NewCmdGenerator("codegen_cmds.go").Generate(descriptor, destDir))
}

// Helper function that reads and unmashals json from given file
func loadJson(file string, val *interface{}) error {
	if _, err := os.Stat(file); err != nil {
		return fmt.Errorf("Cannot find '%s'", file)
	}
	js, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Cannot read '%s': %s", file, err.Error())
	}
	err = json.Unmarshal(js, val)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal JSON read from '%s': %s", file, err.Error())
	}
	return nil
}

// Panic if error is not nil
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "api15gen: %s\n", err.Error())
		os.Exit(1)
	}
}
