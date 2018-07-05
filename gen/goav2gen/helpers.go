package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"bitbucket.org/pkg/inflect"
	"github.com/rightscale/rsc/gen"
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
	// if name == "options" {
	// 	return "options_"
	// }
	// if name == "type" {
	// 	return "type_"
	// }
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

var mediateTypeRe = regexp.MustCompile(`(?i)Mediatype identifier: ([^;]+)(;.*$)?`)
var splitAttrs = regexp.MustCompile(`\s*;\s*`)

func mediaType(title string) string {
	if matches := mediateTypeRe.FindStringSubmatch(title); matches != nil {
		return matches[1]
	}
	return ""
}

func mediaTypeAttrs(title string) map[string]string {
	attrs := map[string]string{}
	if matches := mediateTypeRe.FindStringSubmatch(title); matches != nil {
		allAttrs := matches[2]

		attrBits := splitAttrs.Split(allAttrs, -1)

		for _, attr := range attrBits {
			if attr == "" {
				continue
			}
			kv := strings.Split(attr, "=")
			attrs[kv[0]] = kv[1]
		}
	}
	return attrs
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

func fail(msg string, args ...interface{}) {
	panic(fmt.Errorf(msg, args...))
}
