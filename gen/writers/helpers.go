package writers

import (
	"fmt"
	"os"
	"strings"

	"github.com/kr/text"
	"github.com/rightscale/rsc/gen"
)

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
		countParams += 1
	}
	var params = make([]string, countParams)
	for i, param := range m {
		params[i] = fmt.Sprintf("%s %s", param.VarName, param.Signature())
	}
	if hasOptional {
		params[countParams-1] = "options rsapi.ApiParams"
	}

	return strings.Join(params, ", ")
}

// Produces code that initializes a ApiParams struct with the values of parameters for the given
// action and location.
func paramsInitializer(action *gen.Action, location int, varName string) string {
	var names []string
	switch location {
	case gen.PathParam:
		names = action.PathParamNames
	case gen.QueryParam:
		names = action.QueryParamNames
	case gen.PayloadParam:
		names = action.PayloadParamNames
	}
	if len(names) == 0 {
		return ""
	}
	fields := []string{}
	var optionals []*gen.ActionParam
	for _, param := range action.Params {
		if param.Location != location {
			continue
		}
		if param.Mandatory {
			fields = append(fields, fmt.Sprintf("\"%s\": %s,", param.Name, param.VarName))
		} else {
			optionals = append(optionals, param)
		}
	}
	var paramsDecl = fmt.Sprintf("rsapi.ApiParams{\n%s\n}", strings.Join(fields, "\n\t"))
	if len(optionals) == 0 {
		return fmt.Sprintf("\n%s = %s", varName, paramsDecl)
	}
	var inits = make([]string, len(optionals))
	for i, opt := range optionals {
		inits[i] = fmt.Sprintf("\tvar %sOpt = options[\"%s\"]\n\tif %sOpt != nil {\n\t\t%s[\"%s\"] = %sOpt\n\t}",
			opt.VarName, opt.Name, opt.VarName, varName, opt.Name, opt.VarName)
	}
	var paramsInits = strings.Join(inits, "\n\t")
	return fmt.Sprintf("\n%s = %s\n%s", varName, paramsDecl, paramsInits)
}

// Return true if signature contains pointer, false otherwise
func isPointer(sig string) bool {
	return strings.HasPrefix(sig, "*")
}

// Command line used to run tool
func commandLine() string {
	return fmt.Sprintf("$ %s %s", os.Args[0], strings.Join(os.Args[1:], " "))
}

// Code that checks whether variable with given name and type contains a blank value (empty string,
// empty array or empy map).
// Return empty string if type of variable cannot produce blank values
func blankCondition(name string, t gen.DataType) (blank string) {
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

// Convert description into flag help message
func toHelp(long string) string {
	lines := strings.Split(long, "\n")
	if len(lines) < 2 {
		return long
	}
	if lines[0][len(lines[0])-1] == '.' {
		return lines[0]
	}
	sentences := strings.Split(long, ".")
	if strings.Count(sentences[0], "\n") > 0 {
		sentence := strings.Split(sentences[0], "\n")
		return strings.Join(sentence, " ") + "..."
	}
	return sentences[0]
}

// Type of flag, one of "string", "[]string", "int" or "map"
func flagType(param *gen.ActionParam) string {
	path := param.QueryName
	if strings.HasSuffix(path, "[]") {
		return "[]string"
	}
	if strings.Contains(path, "[]") {
		_, ok := param.Type.(*gen.BasicDataType)
		if !ok {
			if _, ok := param.Type.(*gen.EnumerableDataType); ok {
				return "[]string"
			}
			panic(fmt.Sprintf("Wooaat? an array with a non basic leaf??? - %#v", param.Type))
		}
		return "[]string"
	}
	if _, ok := param.Type.(*gen.ArrayDataType); ok {
		return "[]string"
	} else if _, ok := param.Type.(*gen.EnumerableDataType); ok {
		return "map"
	}
	b, ok := param.Type.(*gen.BasicDataType)
	if !ok {
		panic("Wooaat? a object leaf???")
	}
	return string(*b)
}
