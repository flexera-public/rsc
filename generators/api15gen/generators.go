package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

// Generators use the result of the analyzer to produce artefacts. At the moment there are two
// generators: the ClientGenerator produces the actual API client methods while the CmdGenerator
// produces kingpin commands that can be loaded into a command line tool.
// In the future it probably makes sense to add a TestGenerator that produces tests for all the
// client methods.

// Common interface to all generators
type Generator interface {
	Generate(descriptor *ApiDescriptor, destDir string) error
}

// The ClientGenerator produces a single file that contains the actual API client methods.
type ClientGenerator struct {
	Filename string
}

// Factory method
func NewClientGenerator(filename string) *ClientGenerator {
	return &ClientGenerator{filename}
}

// Go through the descriptor results and generate the code using the code writer.
func (cg *ClientGenerator) Generate(descriptor *ApiDescriptor, destDir string) error {
	codegen := path.Join(destDir, cg.Filename)
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := NewCodeWriter()
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
		return fmt.Errorf("Failed to format generated code:\n%s", o)
	}
	return nil
}

// The CmdGenerator produces kingpin sub-commands that can be integrated into a command line tool.
type CmdGenerator struct {
	Filename string
}

// Factory method
func NewCmdGenerator(filename string) *CmdGenerator {
	return &CmdGenerator{filename}
}

// Go through the descriptor results and generate the sub-commands
func (cg *CmdGenerator) Generate(descriptor *ApiDescriptor, destDir string) error {
	codegen := path.Join(destDir, cg.Filename)
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := NewCmdsWriter(descriptor)
	if err != nil {
		return err
	}
	check(c.WriteHeader(descriptor, f))
	for _, name := range descriptor.ResourceNames {
		resource := descriptor.Resources[name]
		c.WriteResourceHeader(name, f)
		check(c.WriteResource(resource, f))
	}
	f.Close()
	o, err := exec.Command("go", "fmt", codegen).CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to format generated code:\n%s", o)
	}
	return nil
}
