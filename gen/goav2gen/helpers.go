package main // import "gopkg.in/rightscale/rsc.v8/gen/goav2gen"

import (
	"encoding/json"
	"fmt"
	"mime"
	"regexp"
	"strings"

	"bitbucket.org/pkg/inflect"
	"gopkg.in/rightscale/rsc.v8/gen"
)

var (
	// Capture all alphanumerical parts to build go identifier from raw param name
	partsRegexp = regexp.MustCompile("[^[:alnum:]]+")

	// Check whether string only contains blank characters
	blankRegexp = regexp.MustCompile(`^\s*$`)
)

// Produce action return type name
func toGoReturnTypeName(name string, slice bool) string {
	slicePrefix := ""
	if slice {
		slicePrefix = "[]"
	}
	return fmt.Sprintf("%s*%s", slicePrefix, toGoTypeName(name))
}

// toGoTypeName converts from swagger/goa type names to go type names
func toGoTypeName(name string) string {
	switch name {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	}
	return name

}

// Parse native names into go parameter names
func toVarName(name string) string {
	p := partsRegexp.ReplaceAllString(name, "_")
	return inflect.CamelizeDownFirst(p)
}

var toMethodName = inflect.Camelize
var toTypeName = inflect.Camelize

// Return dumpable representation of given object
func prettify(o interface{}) string {
	s, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", o)
	}
	return string(s)
}

// cleans up description, removing blank lines and extraneous info
func cleanDescription(doc string) string {
	docBits := strings.Split(doc, "Required security scope")
	doc = docBits[0]

	lines := strings.Split(doc, "\n")
	fullLines := make([]string, len(lines))
	i := 0
	for _, line := range lines {
		if len(line) > 0 && !blankRegexp.MatchString(line) {
			fullLines[i] = line
			i++
		}
	}
	return strings.Join(fullLines[:i], "\n")
}

func mediaType(title string) string {
	if strings.Contains(title, "Mediatype") {
		ident, _, _ := mime.ParseMediaType(strings.TrimPrefix(title, "Mediatype identifier: "))
		return ident
	}
	return ""
}

func mediaTypeAttrs(title string) map[string]string {
	if strings.Contains(title, "Mediatype") {
		_, params, _ := mime.ParseMediaType(strings.TrimPrefix(title, "Mediatype identifier: "))
		return params
	}
	return map[string]string{}
}

func signature(dt gen.DataType) (sig string) {
	switch t := dt.(type) {
	case *gen.BasicDataType:
		sig = string(*t)
	case *gen.ArrayDataType:
		cs := t.ElemType.Signature()
		sig = fmt.Sprintf("[]%s", cs)
	case *gen.ObjectDataType:
		sig = fmt.Sprintf("*%s", t.TypeName)
	case *gen.EnumerableDataType:
		sig = "map[string]interface{}"
	case *gen.UploadDataType:
		sig = "*rsapi.FileUpload"
	case *gen.SourceUploadDataType:
		sig = "*rsapi.SourceUpload"
	}
	return
}

var dbg = func(msg string, args ...interface{}) (int, error) { return 0, nil }
var warn = fmt.Printf

func fail(msg string, args ...interface{}) {
	panic(fmt.Errorf(msg, args...))
}
