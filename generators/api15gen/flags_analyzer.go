package rsapi15

import (
	"fmt"
	"regexp"
	"strings"
)

// Flag definitions consist of a map of definition to unique name.
// Names can be used to prefix command line args.
type FlagDefs []*FlagDef

// A flag definition contains either a literal flag name or a flag pattern. It also defines the
// type of the value returned by the command line parser for that flag.
type FlagDef struct {
	Param       *ActionParam // Param or param field being defined by this flag
	Path        string       // Path to param field being defined by this flag, e.g. "server_array.datacenter_policy.weight"
	Value       string       // Actual flag name, e.g. "server_array.datacenter_policy.(\d+).weight"
	Type        string       // Flag type, "int" or "string" at the moment
	Description string       // Flag description if any
	IsArray     bool         // Whether param or param field is an array or child of an array element type
	IsEnum      bool         // Whether param or param field is an enum
	IsRequired  bool         // Whether value is required
}

// Generate flag definitions needed to create a command line parser that will gather all the
// data needed to build an instance of the parameter value.
func Flags(p *ActionParam) FlagDefs {
	children := childFlags(p)
	var res FlagDefs
	if len(children) > 0 {
		res = make(FlagDefs, len(children))
		for i, c := range children {
			value := p.Name
			path := p.Name
			if len(c.Value) > 0 {
				value += "." + c.Value
			}
			if c.IsArray || c.IsEnum {
				value = escapeValue(value)
			}
			if len(c.Path) > 0 {
				path += "." + c.Path
			}
			res[i] = &FlagDef{c.Param, path, value, c.Type, c.Description, c.IsArray,
				c.IsEnum, c.IsRequired}
		}
	} else {
		res = FlagDefs{&FlagDef{
			Param:       p,
			Path:        p.Name,
			Value:       p.Name,
			Type:        string(*p.Type.(*BasicDataType)),
			Description: p.Description,
			IsArray:     false,
			IsEnum:      false,
			IsRequired:  p.Mandatory,
		}}
	}
	return res
}

// Helper routine that uses a heuristic to escape a flag value to be used as a regexp.
// The difficulty is that some parts of the flag may be regular expressions while others are literal
// strings that need to be escaped. This algorithm relies on the fact that parts are separated by a
// dot and that flag value regexp parts start with "(".
// A more generic way of doing this is to keep all value parts in an array instead of a string and
// keeping track of which one is a regexp and which one isn't. Seems overkill here though.
func escapeValue(value string) string {
	parts := strings.Split(value, ".")
	escaped := make([]string, len(parts))
	for i, p := range parts {
		if p[0] == '(' {
			escaped[i] = p
		} else {
			escaped[i] = regexp.QuoteMeta(p)
		}
	}
	return strings.Join(escaped, `\.`)
}

// Child flags, used for recursion
func childFlags(p *ActionParam) []*FlagDef {
	var res []*FlagDef
	switch t := p.Type.(type) {
	case *BasicDataType:
		res = FlagDefs{}
	case *ArrayDataType:
		flags := childFlags(t.ElemType)
		if len(flags) > 0 {
			res = make([]*FlagDef, len(flags))
			for i, f := range flags {
				res[i] = &FlagDef{
					Param:       f.Param,
					Path:        f.Path,
					Value:       fmt.Sprintf(`(\d+).%s`, f.Value),
					Type:        f.Type,
					Description: f.Description,
					IsArray:     true,
					IsEnum:      false,
					IsRequired:  f.IsRequired,
				}
			}
		} else {
			res = FlagDefs{&FlagDef{
				Param:       p,
				Path:        "",
				Value:       `(\d+)`,
				Type:        "string",
				Description: p.Description,
				IsArray:     true,
				IsEnum:      false,
				IsRequired:  p.Mandatory,
			}}
		}
	case *ObjectDataType:
		res = FlagDefs{}
		for _, field := range t.Fields {
			flags := childFlags(field)
			temp := make([]*FlagDef, len(flags))
			if len(flags) > 0 {
				for i, f := range flags {
					fName := field.Name
					path := fName
					if len(f.Path) > 0 {
						path += "." + f.Path
					}
					temp[i] = &FlagDef{
						Param:       f.Param,
						Path:        path,
						Value:       fName + "." + f.Value,
						Type:        f.Type,
						Description: f.Description,
						IsArray:     f.IsArray,
						IsEnum:      f.IsEnum,
						IsRequired:  f.IsRequired,
					}
				}
				res = append(res, temp...)
			} else {
				// No child means field's type is a basic type
				fType := string(*field.Type.(*BasicDataType))
				res = append(res, &FlagDef{
					Param:       field,
					Path:        field.Name,
					Value:       field.Name,
					Type:        fType,
					Description: field.Description,
					IsArray:     false,
					IsEnum:      false,
					IsRequired:  field.Mandatory,
				})
			}
		}
	case *EnumerableDataType:
		res = FlagDefs{&FlagDef{
			Param:       p,
			Path:        "",
			Value:       `([a-z0-9_]+)`,
			Type:        "string",
			Description: p.Description,
			IsArray:     false,
			IsEnum:      true,
			IsRequired:  p.Mandatory,
		}}
	}
	return res
}
