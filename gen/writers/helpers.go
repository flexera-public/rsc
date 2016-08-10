package writers

import (
	"fmt"
	"os"
	"strings"

	"github.com/rightscale/rsc/gen"
	"github.com/rightscale/rsc/gen/writers/text"
)

// reserved golang keywords and package names
var reserved = map[string]bool{
	"byte":       true,
	"complex128": true,
	"complex64":  true,
	"float32":    true,
	"float64":    true,
	"int":        true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"int8":       true,
	"rune":       true,
	"string":     true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uint8":      true,

	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"defer":       true,
	"else":        true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,

	// stdlib and rsc packages used by generated code
	"fmt":      true,
	"ioutil":   true,
	"json":     true,
	"time":     true,
	"metadata": true,
	"rsapi":    true,
}

// fixReserved appends an underscore on to Go reserved keywords.
func fixReserved(w string) string {
	if reserved[w] {
		w += "_"
	}
	return w
}

// Produce line comments by concatenating given strings and producing 80 characters long lines
// starting with "//"
func comment(elems ...string) string {
	var lines []string
	for _, e := range elems {
		lines = append(lines, strings.Split(e, "\n")...)
	}
	var trimmed = make([]string, len(lines))
	for i, l := range lines {
		trimmed[i] = strings.TrimLeft(l, " \t")
	}
	t := strings.Join(trimmed, "\n")
	return text.Indent(t, "// ")
}

// Serialize action parameters
func parameters(a *gen.Action) string {
	var m = a.MandatoryParams()
	var hasOptional = a.HasOptionalParams()
	var countParams = len(m)
	if hasOptional {
		countParams++
	}
	var params = make([]string, countParams)
	for i, param := range m {
		params[i] = fmt.Sprintf("%s %s", fixReserved(param.VarName), param.Signature())
	}
	if hasOptional {
		params[countParams-1] = "options rsapi.APIParams"
	}

	return strings.Join(params, ", ")
}

// Produces code that initializes a APIParams struct with the values of parameters for the given
// action and location.
func paramsInitializer(action *gen.Action, location int, varName string) string {
	var fields []string
	var optionals []*gen.ActionParam
	varName = fixReserved(varName)
	for _, param := range action.Params {
		if param.Location != location {
			continue
		}
		if param.Mandatory {
			name := param.Name
			if location == 1 { // QueryParam
				name = param.QueryName
			}
			fields = append(fields, fmt.Sprintf("\"%s\": %s,", name, fixReserved(param.VarName)))
		} else {
			optionals = append(optionals, param)
		}
	}
	if len(fields) == 0 && len(optionals) == 0 {
		return ""
	}
	var paramsDecl = fmt.Sprintf("rsapi.APIParams{\n%s\n}", strings.Join(fields, "\n\t"))
	if len(optionals) == 0 {
		return fmt.Sprintf("\n%s = %s", varName, paramsDecl)
	}
	var inits = make([]string, len(optionals))
	for i, opt := range optionals {
		name := opt.Name
		if location == 1 { // QueryParam
			name = opt.QueryName
		}
		inits[i] = fmt.Sprintf("\tvar %sOpt = options[\"%s\"]\n\tif %sOpt != nil {\n\t\t%s[\"%s\"] = %sOpt\n\t}",
			opt.VarName, opt.Name, opt.VarName, varName, name, opt.VarName)
	}
	var paramsInits = strings.Join(inits, "\n\t")
	return fmt.Sprintf("\n%s = %s\n%s", varName, paramsDecl, paramsInits)
}

// Command line used to run tool
func commandLine() string {
	return fmt.Sprintf("$ %s %s", os.Args[0], strings.Join(os.Args[1:], " "))
}

// Code that checks whether variable with given name and type contains a blank value (empty string,
// empty array or empy map).
// Return empty string if type of variable cannot produce blank values
func blankCondition(name string, t gen.DataType) (blank string) {
	name = fixReserved(name)
	switch actual := t.(type) {
	case *gen.BasicDataType:
		if *actual == "string" {
			blank = fmt.Sprintf("if %s == \"\" {", name)
		}
	case *gen.ArrayDataType:
		blank = fmt.Sprintf("if len(%s) == 0 {", name)
	case *gen.ObjectDataType:
		blank = fmt.Sprintf("if %s == nil {", name)
	case *gen.EnumerableDataType:
		blank = fmt.Sprintf("if len(%s) == 0 {", name)
	}
	return
}

// GET => Get
func toVerb(text string) (res string) {
	res = strings.ToUpper(string(text[0])) + strings.ToLower(text[1:])
	if text == "GET" || text == "POST" {
		res += "Raw"
	}
	return
}

// *ServerArrayLocator => ServerArrayLocator
func stripStar(text string) string {
	if text[0] == '*' {
		return text[1:]
	}
	return text
}

// Convert []interface{} into []string
func toStringArray(a []interface{}) []string {
	res := make([]string, len(a))
	for i, v := range a {
		res[i] = fmt.Sprintf("%v", v)
	}
	return res
}

// Escape ` in string to be wrapped in them
func escapeBackticks(d string) string {
	elems := strings.Split(d, "`")
	return strings.Join(elems, "` + `")
}

// Type of flag, one of "string", "[]string", "int", "map" or "file"
func flagType(param *gen.ActionParam) string {
	path := param.QueryName
	if _, ok := param.Type.(*gen.ArrayDataType); ok {
		return "[]string"
	} else if _, ok := param.Type.(*gen.EnumerableDataType); ok {
		return "map"
	} else if _, ok := param.Type.(*gen.SourceUploadDataType); ok {
		return "sourcefile"
	} else if _, ok := param.Type.(*gen.UploadDataType); ok {
		return "file"
	}
	b, ok := param.Type.(*gen.BasicDataType)
	if !ok {
		panic("Wooaat? a object leaf??? - " + path)
	}
	return string(*b)
}
