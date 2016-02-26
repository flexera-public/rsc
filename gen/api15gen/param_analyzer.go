package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/rightscale/rsc/gen"

	"bitbucket.org/pkg/inflect"
)

var (
	// Capture root of path
	rootRegexp = regexp.MustCompile("([^\\[]+)\\[")

	// Parent path regular expression
	parentPathRegexp = regexp.MustCompile(`^(.*)\[[^\]]+\]$`)

	// Child path regular expression
	childPathRegexp = regexp.MustCompile(`^.*\[([^\]]+)\](\[\])?$`)

	// Capture all alphanumerical parts to build go identifier from raw param name
	partsRegexp = regexp.MustCompile("[^[:alnum:]]+")
)

// ParamAnalyzer exposes the "Analyze" method which initializes all the fields but 'rawParams'
// which is initialized by the factory method.
// The analyzer takes a map describing the parameters of a method as found in the API JSON and
// produces the corresponding set of ActionParam structs.
type ParamAnalyzer struct {
	// Raw parameter hashes as found in JSON
	rawParams map[string]interface{}
	// Temporary data structure used by analysis
	parsed map[string]*gen.ActionParam
	// Leaf parameter names sorted alphabetically
	leafParamNames []string

	/* Fields below are computed by 'analyze' */

	// Parameter types indexed by name
	ParamTypes map[string]*gen.ObjectDataType
	// Parameters sorted alphabetically by name
	Params []*gen.ActionParam
	// Leaf parameters used for command line
	LeafParams []*gen.ActionParam
}

// NewAnalyzer is the factory method, it initializes the 'path' and 'rawParams' fields.
func NewAnalyzer(params map[string]interface{}) *ParamAnalyzer {
	return &ParamAnalyzer{rawParams: params}
}

// Analyze all parameters and categorize them
// Initialize all fields of ParamAnalyzer struct
func (p *ParamAnalyzer) Analyze() {
	// Order params using their length so "foo[bar]" is analyzed before "foo"
	params := p.rawParams
	paths := make([]string, len(params))
	i := 0
	for n := range params {
		paths[i] = n
		i++
	}
	sort.Strings(paths)
	sort.Sort(ByReverseLength(paths))
	rawLeafParams := []string{}
	for _, p := range paths {
		hasLeaf := false
		for _, r := range rawLeafParams {
			if strings.HasSuffix(r, "]") && strings.HasPrefix(r, p) {
				hasLeaf = true
				break
			}
		}
		if hasLeaf {
			continue
		}
		rawLeafParams = append(rawLeafParams, p)
	}
	sort.Strings(rawLeafParams)
	p.leafParamNames = rawLeafParams

	// Iterate through all params and build corresponding ActionParam structs
	p.parsed = map[string]*gen.ActionParam{}
	top := map[string]*gen.ActionParam{}
	for _, path := range paths {
		if strings.HasSuffix(path, "[*]") {
			// Cheat a little bit - there a couple of cases where parent type is
			// Hash instead of Enumerable, make that enumerable everywhere
			// There are also cases where there's no parent path, fix that up also
			matches := parentPathRegexp.FindStringSubmatch(path)
			if hashParam, ok := params[matches[1]].(map[string]interface{}); ok {
				hashParam["class"] = "Enumerable"
			} else {
				// Create parent
				rawParams := map[string]interface{}{}
				parentPath := matches[1]
				p.parsed[parentPath] = p.newParam(parentPath, rawParams,
					new(gen.EnumerableDataType))
				if parentPathRegexp.FindStringSubmatch(parentPath) == nil {
					top[parentPath] = p.parsed[parentPath]
				}
			}
			continue
		}
		var child *gen.ActionParam
		origPath := path
		origParam := params[path].(map[string]interface{})
		matches := parentPathRegexp.FindStringSubmatch(path)
		isTop := (matches == nil)
		if prev, ok := p.parsed[path]; ok {
			if isTop {
				top[path] = prev
			}
			continue
		}
		var branch []*gen.ActionParam
		for matches != nil {
			param := params[path].(map[string]interface{})
			parentPath := matches[1]
			var isArrayChild bool
			if strings.HasSuffix(parentPath, "[]") {
				isArrayChild = true
			}
			if parent, ok := p.parsed[parentPath]; ok {
				a, ok := parent.Type.(*gen.ArrayDataType)
				if ok {
					parent = a.ElemType
				}
				child = p.parseParam(path, param, child)
				if !parent.Mandatory {
					// Make required fields of optional hashes optional.
					child.Mandatory = false
				}
				branch = append(branch, child)
				if _, ok = parent.Type.(*gen.EnumerableDataType); !ok {
					o := parent.Type.(*gen.ObjectDataType)
					o.Fields = appendSorted(o.Fields, child)
					p.parsed[path] = child
				}
				break // No need to keep going back, we already have a parent
			} else {
				child = p.parseParam(path, param, child)
				branch = append(branch, child)
				p.parsed[path] = child
				if isArrayChild {
					// Generate array item as it's not listed explicitly in JSON
					itemPath := matches[1] + "[item]"
					typeName := p.typeName(matches[1])
					parent = p.newParam(itemPath, map[string]interface{}{},
						&gen.ObjectDataType{typeName, []*gen.ActionParam{child}})
					p.parsed[parentPath] = parent
					child = parent
					branch = append(branch, child)
					parentPath = parentPath[:len(parentPath)-2]
				}
			}
			path = parentPath
			matches = parentPathRegexp.FindStringSubmatch(path)
		}
		if isTop {
			if _, ok := p.parsed[path]; !ok {
				actionParam := p.parseParam(path, origParam, nil)
				p.parsed[path] = actionParam
			}
			top[path] = p.parsed[path]
		} else {
			matches := rootRegexp.FindStringSubmatch(origPath)
			rootPath := matches[1]
			if _, ok := p.parsed[rootPath]; !ok {
				p.parsed[rootPath] = p.parseParam(rootPath,
					params[rootPath].(map[string]interface{}), child)
			}
			actionParam, _ := p.parsed[rootPath]
			mandatory := actionParam.Mandatory
			if len(branch) > 0 {
				for i := len(branch) - 1; i >= 0; i-- {
					p := branch[i]
					if mandatory {
						if !p.Mandatory {
							mandatory = false
						}
					} else {
						p.Mandatory = false
					}
				}
			}
		}
	}
	// Now do a second pass on parsed params to generate their declarations
	p.ParamTypes = make(map[string]*gen.ObjectDataType)
	for _, param := range top {
		p.recordTypes(param.Type)
	}

	i = 0
	res := make([]*gen.ActionParam, len(top))
	for _, param := range top {
		res[i] = param
		i++
	}
	sort.Sort(gen.ByName(res))
	p.Params = res
}

// ByReverseLength makes it possible to sort an array of strings by length.
type ByReverseLength []string

func (s ByReverseLength) Len() int {
	return len(s)
}
func (s ByReverseLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByReverseLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

// Recursively record all type declarations
func (p *ParamAnalyzer) recordTypes(root gen.DataType) {
	if o, ok := root.(*gen.ObjectDataType); ok {
		if _, found := p.ParamTypes[o.TypeName]; !found {
			p.ParamTypes[o.TypeName] = o
			for _, f := range o.Fields {
				p.recordTypes(f.Type)
			}
		}
	} else if a, ok := root.(*gen.ArrayDataType); ok {
		p.recordTypes(a.ElemType.Type)
	}
}

// Sort action params by name
func appendSorted(params []*gen.ActionParam, param *gen.ActionParam) []*gen.ActionParam {
	params = append(params, param)
	sort.Sort(gen.ByName(params))
	return params
}

// Parse data type in context
func (p *ParamAnalyzer) parseDataType(path string, child *gen.ActionParam) gen.DataType {
	param := p.rawParams[path].(map[string]interface{})
	class := "String"
	if c, ok := param["class"].(string); ok {
		class = c
	}
	var res gen.DataType
	switch class {
	case "Integer":
		i := gen.BasicDataType("int")
		res = &i
	case "String":
		s := gen.BasicDataType("string")
		res = &s
	case "Array":
		if child != nil {
			res = &gen.ArrayDataType{child}
		} else {
			s := gen.BasicDataType("string")
			p := p.newParam(fmt.Sprintf("%s[item]", path),
				map[string]interface{}{}, &s)
			res = &gen.ArrayDataType{p}
		}
	case "SourceUpload":
		res = new(gen.SourceUploadDataType)
	case "FileUpload", "Tempfile":
		res = &gen.UploadDataType{TypeName: "FileUpload"}
	case "Enumerable":
		res = new(gen.EnumerableDataType)
	case "Hash":
		if current, ok := p.parsed[path]; ok {
			res = current.Type
			o := res.(*gen.ObjectDataType)
			o.Fields = appendSorted(o.Fields, child)
		} else {
			oname := p.typeName(path)
			res = &gen.ObjectDataType{oname, []*gen.ActionParam{child}}
		}
	}
	return res
}

func (p *ParamAnalyzer) typeName(path string) string {
	matches := childPathRegexp.FindStringSubmatch(path)
	res := path
	if matches != nil {
		res = matches[1]
	}
	return strings.Title(parseParamName(res))
}

// Build action param struct from json data
func (p *ParamAnalyzer) parseParam(path string, param map[string]interface{}, child *gen.ActionParam) *gen.ActionParam {
	dType := p.parseDataType(path, child)
	return p.newParam(path, param, dType)
}

// New parameter from raw values
func (p *ParamAnalyzer) newParam(path string, param map[string]interface{}, dType gen.DataType) *gen.ActionParam {
	var description, regexp string
	var mandatory, nonBlank bool
	var validValues []interface{}
	if d, ok := param["description"]; ok {
		description = d.(string)
	}
	if m, ok := param["mandatory"]; ok {
		mandatory = m.(bool)
	}
	if n, ok := param["non_blank"]; ok {
		nonBlank = n.(bool)
	}
	if r, ok := param["regexp"]; ok {
		regexp = r.(string)
	}
	if v, ok := param["valid_values"]; ok {
		validValues = v.([]interface{})
	}
	native := nativeNameFromPath(path)
	isLeaf := false
	if _, ok := dType.(*gen.EnumerableDataType); ok {
		isLeaf = true
	} else {
		for _, l := range p.leafParamNames {
			if path == l {
				isLeaf = true
				break
			}
		}
	}
	queryName := path
	if _, ok := dType.(*gen.ArrayDataType); ok {
		queryName += "[]"
	}
	actionParam := &gen.ActionParam{
		Name:        native,
		QueryName:   queryName,
		Description: removeBlankLines(description),
		VarName:     parseParamName(native),
		Type:        dType,
		Mandatory:   mandatory,
		NonBlank:    nonBlank,
		Regexp:      regexp,
		ValidValues: validValues,
	}
	if isLeaf {
		p.LeafParams = append(p.LeafParams, actionParam)
	}
	return actionParam
}

// Check whether string only contains blank characters
var blankRegexp = regexp.MustCompile(`^\s*$`)

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

// Extract name (leaf) from path
func nativeNameFromPath(path string) string {
	native := path
	matches := childPathRegexp.FindStringSubmatch(path)
	if matches != nil {
		native = matches[1]
	}

	return native
}

// Parse native names into go parameter names
func parseParamName(name string) string {
	if name == "r_s_version" {
		return "rsVersion"
	}
	if name == "type" {
		return "type_"
	}
	if name == "options" {
		return "options_"
	}
	p := partsRegexp.ReplaceAllString(name, "_")
	return inflect.CamelizeDownFirst(p)
}
