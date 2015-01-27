package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"bitbucket.org/pkg/inflect"
)

var (
	// Path to generated file
	targetFile string

	// Attribute types mapping
	attributeTypes map[string]string

	// Regexp used to match RightScale media type identifiers
	rightScaleType = regexp.MustCompile("^vnd\\.rightscale\\.")

	// Resources that don't have a media type
	NoMediaTypeResources = map[string]bool{
		"HealthCheck":   true,
		"Oauth2":        true,
		"Tag":           true,
		"UserDatas":     true,
		"ChildAccounts": true,
	}

	// Datetime attribute name regexp
	dateAttributeRegex = regexp.MustCompile("_at$")
)

func main() {
	destDefault, _ := filepath.Abs("codegen.go")
	metadata := flag.String("metadata", "./api_data.json",
		"Path to API 1.5 metadata file, defaults to './api_data.json'")
	attributes := flag.String("attributes", "./attributes.json",
		"Path to API 1.5 attribute types file, defaults to './attributes.json'")
	targetFile = *flag.String("output", destDefault,
		"Path to output file, defaults to './codegen.go'")
	flag.Parse()
	if len(*metadata) == 0 {
		check(fmt.Errorf("Specify path to metadata json with '-metadata /path/to/api_data.json'"))
	}
	at, err := ioutil.ReadFile(*attributes)
	check(err)
	check(json.Unmarshal(at, &attributeTypes))
	js, err := ioutil.ReadFile(*metadata)
	check(err)
	var content map[string]interface{}
	check(json.Unmarshal(js, &content))
	f, err := os.Create(targetFile)
	check(err)
	c, err := NewCodeWriter()
	check(err)
	check(c.WriteHeader(f))
	names := sortedKeys(content)
	for _, name := range names {
		data, err := analyzeResource(name, content[name])
		check(err)
		check(c.WriteResource(data, f))
	}
}

func analyzeResource(name string, resource interface{}) (*ResourceData, error) {
	res := resource.(map[string]interface{})

	// Compute actions
	methods := res["methods"].(map[string]interface{})
	actionNames := sortedKeys(methods)
	actions := make([]*ResourceAction, len(methods))
	idx := 0
	for _, actionName := range actionNames {
		m := methods[actionName]
		meth := m.(map[string]interface{})
		var params map[string]interface{}
		if p, ok := meth["parameters"]; ok {
			params = p.(map[string]interface{})
		}
		description := "<no description>"
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		httpMethod, path := parseRoute(meth["route"].(string))
		var contentType string
		if c, ok := meth["content_type"].(string); ok {
			contentType = c
		}
		actions[idx] = &ResourceAction{
			Name:        methodName(actionName, name),
			Description: description,
			HttpMethod:  httpMethod,
			Path:        path,
			Params:      parseParams(params),
			Return:      parseReturn(actionName, name, contentType),
		}
		idx += 1
	}

	// Compute description
	var description string
	if d, ok := res["description"].(string); ok {
		description = d
	}

	// Compute attributes
	var attributes []*ResourceAttribute
	var atts map[string]interface{}
	if a, ok := res["media_type"].(map[string]interface{}); ok {
		atts = a["attributes"].(map[string]interface{})
		attributes = make([]*ResourceAttribute, len(atts))
		for idx, n := range sortedKeys(atts) {
			attributes[idx] = &ResourceAttribute{n, attributeTypes[n]}
		}
	} else {
		attributes = []*ResourceAttribute{}
	}

	// We're done!
	return &ResourceData{
		Name:        name,
		Description: description,
		Actions:     actions,
		Attributes:  attributes,
	}, nil
}

func parseRoute(route string) (string, string) {
	if len(route) < 7 {
		return "", "<unknown route>" // :(((( some routes are empty
	}
	method := strings.TrimRight(route[0:7], " ")
	path := strings.TrimRight(strings.Split(route[8:], "{")[0], " ")

	return method, path
}

func parseParams(params map[string]interface{}) []*ActionParam {
	paramNames := sortedKeys(params)
	res := make([]*ActionParam, len(paramNames))
	j := 0
	for _, name := range paramNames {
		p := params[name]
		param := p.(map[string]interface{})
		n := parseParamName(name)
		class, ok := param["class"]
		if !ok {
			res[j] = &ActionParam{n, goType("String")} // Assume string, e.g. 'limit' of audit entries index
			j += 1
		} else if class != "Hash" {
			res[j] = &ActionParam{n, goType(class.(string))}
			j += 1
		}
	}
	return res[0:j]
}

func parseParamName(name string) string {
	r := regexp.MustCompile("[^[:alnum:]]+")
	p := r.ReplaceAllString(name, "_")
	return inflect.CamelizeDownFirst(p)
}

func parseReturn(kind, resName, contentType string) string {
	switch kind {
	case "show":
		return fmt.Sprintf("*%s", resourceType(resName))
	case "index":
		return fmt.Sprintf("[]%s", resourceType(resName))
	case "create":
		return "href"
	case "update", "destroy":
		return ""
	default:
		switch {
		case len(contentType) == 0:
			return ""
		case strings.Index(contentType, "application/vnd.rightscale.") == 0:
			if contentType == "application/vnd.rightscale.text" {
				return "string"
			}
			elems := strings.SplitN(contentType[27:], ";", 2)
			name := inflect.Camelize(elems[0])
			if len(elems) > 1 && elems[1] == "type=collection" {
				name = "[]" + name
			}
			return name
		default: // Shouldn't be here
			panic("api15gen: Unknown content type " + contentType)
		}
	}

}

// Name of go type corresponding to metadata type
func goType(apiType string) string {
	switch apiType {
	case "Integer":
		return "int"
	case "String":
		return "string"
	case "Array":
		return "[]string"
	case "Enumerable":
		return "map[string]string" // e.g. inputs
	default:
		panic("Unknown API type " + apiType)
	}
}

// Name of go type for resource with given name
// It should always be the same (camelized) but there are some resources that don't have a media
// type so for these we use a map.
func resourceType(resName string) string {
	if _, ok := NoMediaTypeResources[resName]; ok {
		return "map[string]string"
	} else {
		return inflect.Singularize(resName)
	}
}

// Heuristic to create method name from action and resource names
func methodName(action, resource string) string {
	if action != "index" && strings.Index(action, "multi") != 0 {
		resource = inflect.Singularize(resource)
	}
	return inflect.Camelize(action) + inflect.Camelize(resource)
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

// Panic if error is not nil
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "api15gen: %s\n", err.Error())
		os.Remove(targetFile)
		os.Exit(1)
	}
}
