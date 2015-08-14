package main // import "gopkg.in/rightscale/rsc.v3/gen/praxisgen"

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"bitbucket.org/pkg/inflect"
)

var (
	// Capture all alphanumerical parts to build go identifier from raw param name
	partsRegexp = regexp.MustCompile("[^[:alnum:]]+")

	// Check whether string only contains blank characters
	blankRegexp = regexp.MustCompile(`^\s*$`)

	// List of praxis built-in types for which there is no "type" member in the corresponding JSON
	BuiltInTypes = []string{"RsId", "IP", "QueryFilter", "Href", "Tag", "CSV"}
)

// Produce go type name from given ruby type name
func toGoTypeName(name string, usePointer bool) string {
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
		var typeName string
		if strings.Contains(name, "::") {
			elems := strings.Split(name, "::")
			typeName = strings.Join(elems[2:len(elems)], "")
		} else {
			typeName = name
		}
		if usePointer {
			typeName = "*" + typeName
		}
		return typeName
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
			i += 1
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
	for k, _ := range m {
		keys[idx] = k
		idx += 1
	}
	sort.Strings(keys)
	return keys
}
