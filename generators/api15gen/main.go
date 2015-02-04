package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
)

var (
	// Path to generated file
	targetFile string

	// Whether to keep generated file even if there were errors during generation
	keep bool
)

func main() {
	// 1. Parse command line arguments
	destDefault, _ := filepath.Abs("codegen.go")
	metadata := flag.String("metadata", "./api_data.json",
		"Path to API 1.5 metadata file")
	attributes := flag.String("attributes", "./attributes.json",
		"Path to API 1.5 attribute types file")
	targetFilePtr := flag.String("output", destDefault,
		"Path to output file")
	keepPtr := flag.Bool("keep", false, "Keep generated code even if generation produced errors")
	flag.Parse()
	targetFile = *targetFilePtr
	keep = *keepPtr
	if len(*metadata) == 0 {
		check(fmt.Errorf("Specify path to metadata json with '-metadata /path/to/api_data.json'"))
	}
	at, err := ioutil.ReadFile(*attributes)
	check(err)
	var attributeTypes map[string]string
	check(json.Unmarshal(at, &attributeTypes))
	js, err := ioutil.ReadFile(*metadata)
	check(err)
	var content map[string]interface{}
	check(json.Unmarshal(js, &content))

	// 2. Analyze
	analyzer := NewApiAnalyzer(content, attributeTypes)
	analyzer.Analyze()

	// 3. Write codegen.go
	f, err := os.Create(targetFile)
	check(err)
	c, err := NewCodeWriter()
	check(err)
	check(c.WriteHeader(f))
	names := sortedKeys(content)
	for _, name := range names {
		resource := analyzer.Resources[name]
		c.WriteResourceHeader(name, f)
		check(c.WriteResource(resource, f))
	}
	typeNames := make([]string, len(analyzer.Types))
	for n, _ := range analyzer.Types {
		typeNames = append(typeNames, n)
	}
	sort.Strings(typeNames)
	for _, name := range typeNames {
		t := analyzer.Types[name]
		c.WriteType(t, f)
	}
	f.Close()

	// 4. "go fmt" codegen.go
	o, err := exec.Command("go", "fmt", targetFile).CombinedOutput()
	if err != nil {
		check(fmt.Errorf("Failed to format generated code:\n%s", o))
	}
}

// Panic if error is not nil
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "api15gen: %s\n", err.Error())
		if !keep {
			os.Remove(targetFile)
		}
		os.Exit(1)
	}
}
