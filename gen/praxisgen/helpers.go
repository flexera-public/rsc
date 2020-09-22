package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/go-openapi/inflect"
)

var (
	// Capture all alphanumerical parts to build go identifier from raw param name
	partsRegexp = regexp.MustCompile("[^[:alnum:]]+")

	// Check whether string only contains blank characters
	blankRegexp = regexp.MustCompile(`^\s*$`)

	// BuiltInTypes is the list of praxis built-in types for which there is no "type" member
	// in the corresponding JSON.
	BuiltInTypes = []string{"RsId", "IP", "QueryFilter", "Href", "Tag", "CSV"}
)

// Produce action return type name
func toGoReturnTypeName(name string, slice bool) string {
	slicePrefix := ""
	if slice {
		slicePrefix = "[]"
	}
	return fmt.Sprintf("%s*%s", slicePrefix, toGoTypeName(name))
}

// Produce go type name from given ruby type name
func toGoTypeName(name string) string {
	switch name {
	case "String", "Symbol":
		return "string"
	case "Integer":
		return "int"
	case "Boolean":
		return "bool"
	case "Struct", "Collection":
		panic("Uh oh, trying to infer a go type name for a unnamed struct or collection (" + name + ")")
	default:
		if strings.Contains(name, "::") {
			elems := strings.Split(name, "::")
			return strings.Join(elems[2:len(elems)], "")
		}
		return name
	}

}

// Parse native names into go parameter names
func toVarName(name string) string {
	if name == "options" {
		return "options_"
	}
	if name == "type" {
		return "type_"
	}
	p := partsRegexp.ReplaceAllString(name, "_")
	return inflect.CamelizeDownFirst(p)
}

// Return dumpable representation of given object
func prettify(o interface{}) string {
	s, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", o)
	}
	return string(s)
}

// Helper method that removes blank lines from strings
func removeBlankLines(doc string) string {
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

// Returns true if given name is the name of a built-in type
func isBuiltInType(name string) bool {
	for _, n := range BuiltInTypes {
		if name == n {
			return true
		}
	}
	return false
}

// Return keys of given maps sorted
func sortedKeys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	return keys
}
