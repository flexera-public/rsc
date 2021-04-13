package writers // import "gopkg.in/rightscale/rsc.v9/gen/writers"

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-openapi/inflect"
	"gopkg.in/rightscale/rsc.v9/gen"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Generate API client code, drives the code writer.
func GenerateClient(version string, descriptor *gen.APIDescriptor, codegen, pkg string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := NewClientWriter()
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
func GenerateMetadata(descriptor *gen.APIDescriptor, codegen, pkg string) error {
	f, err := os.Create(codegen)
	if err != nil {
		return err
	}
	c, err := NewMetadataWriter()
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
func GenerateAngular(descriptor *gen.APIDescriptor, pkgDir string) ([]string, error) {
	var files []string
	for _, name := range descriptor.ResourceNames {
		res := descriptor.Resources[name]
		codegen := filepath.Join(pkgDir, inflect.Underscore(name)+".js")
		f, err := os.Create(codegen)
		if err != nil {
			return files, err
		}
		c, err := NewAngularWriter()
		if err != nil {
			return files, err
		}
		kingpin.FatalIfError(c.WriteResource(res, f), "")
		f.Close()
		files = append(files, codegen)
	}
	return files, nil
}
