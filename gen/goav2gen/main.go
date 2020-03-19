package main // import "gopkg.in/rightscale/rsc.v8/gen/goav2gen"

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/rightscale/rsc.v8/gen/writers"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// 1. Parse command line arguments
	curDir, err := os.Getwd()
	kingpin.FatalIfError(err, "")
	metadata := flag.String("metadata", curDir,
		"Path to directories containing openapi.json.")
	destDirVal := flag.String("output", curDir,
		"Path to output file")
	pkgName := flag.String("pkg", "", "Name of generated package, e.g. \"policy\"")
	version := flag.String("version", "", "Version of API to generate code for")
	clientName := flag.String("client", "", "Name of API client go struct, e.g. \"Policy\".")
	tool := flag.String("tool", "rsc", "Tool or library for which to generate code, supported values are 'rsc' or 'angular'")
	debug := flag.Bool("debug", false, "debug output")

	flag.Parse()

	destDir := *destDirVal
	if stat, _ := os.Stat(destDir); stat != nil && !stat.IsDir() {
		kingpin.Fatalf("%s is not a valid directory.", destDir)
	}

	if *version == "" {
		kingpin.Fatalf("-version option is required.")
	}

	if *tool == "rsc" {
		if *pkgName == "" {
			kingpin.Fatalf("-pkg option is required.")
		}

		if *clientName == "" {
			kingpin.Fatalf("-client option is required.")
		}
	}

	if *debug {
		dbg = fmt.Printf
	}

	if stat, err := os.Stat(*metadata); err != nil || !stat.IsDir() {
		kingpin.Fatalf("%s is not a valid directory.", *metadata)
	}

	docFile := path.Join(*metadata, "openapi.json")
	data, err := loadFile(docFile)
	kingpin.FatalIfError(err, "")
	var doc Doc
	err = json.Unmarshal(data, &doc)
	if err != nil {
		kingpin.Fatalf("Cannot unmarshal JSON read from '%s': %s", docFile, err)
	}

	attrFile := path.Join(*metadata, "attributes.json")
	var attrOverrides = make(map[string]string)
	if fileExists(attrFile) {
		attrData, err := loadFile(attrFile)
		kingpin.FatalIfError(err, "")
		err = json.Unmarshal(attrData, &attrOverrides)
		if err != nil {
			kingpin.Fatalf("Cannot unmarshal JSON read from '%s': %s", attrFile, err)
		}
	}
	typeFile := path.Join(*metadata, "types.json")
	var typeOverrides = make(map[string]string)
	if fileExists(typeFile) {
		typeData, err := loadFile(typeFile)
		kingpin.FatalIfError(err, "")
		err = json.Unmarshal(typeData, &typeOverrides)
		if err != nil {
			kingpin.Fatalf("Cannot unmarshal JSON read from '%s': %s", typeFile, err)
		}
	}

	// 2. Analyze
	a := NewAPIAnalyzer(*version, *clientName, &doc, attrOverrides, typeOverrides)
	api, err := a.Analyze()
	kingpin.FatalIfError(err, "")

	// 3. Write code
	var generated []string
	os.MkdirAll(path.Join(destDir), 0755)
	switch *tool {
	case "rsc":
		clientPath := path.Join(destDir, "codegen_client.go")
		metadataPath := path.Join(destDir, "codegen_metadata.go")
		kingpin.FatalIfError(writers.GenerateClient(*version, api, clientPath, *pkgName), "")
		kingpin.FatalIfError(writers.GenerateMetadata(api, metadataPath, *pkgName), "")
		generated = append(generated, clientPath)
		generated = append(generated, metadataPath)
	case "angular":
		files, err := writers.GenerateAngular(api, destDir)
		kingpin.FatalIfError(err, "")
		generated = append(generated, files...)
	default:
		kingpin.Fatalf("Invalid tool '%s', supported clients are 'rsc' and 'angular'", *tool)
	}

	// 4. Say something...
	for i := 0; i < len(generated); i++ {
		fmt.Printf("%s\n", generated[i])
	}
}

// loadFile reads content from existing file and returns a byte array
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

// fileExists reads content from existing file and returns a byte array
func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
