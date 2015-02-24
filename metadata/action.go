package metadata

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Resource action
type Action struct {
	Name              string
	Description       string
	PathPatterns      []*PathPattern // Action path patterns and variables
	Params            []*ActionParam // All parameter details
	QueryParamNames   []string       // Query string parameter names, e.g. "filter" in /clouds?filter[]=name==foo
	PayloadParamNames []string       // Payload parameter names (top level keys of payload structure)
}

// Resource action parametersn
type ActionParam struct {
	Name        string         // Param name
	Description string         // Param description
	Type        string         // Param type, one of "string", "[]string", "int", "[]int" or "map[string]string"
	Mandatory   bool           // Whether parameter is mandatory
	NonBlank    bool           // Whether parameter value can be blank
	Regexp      *regexp.Regexp // Regular expression used to validate parameter values
	ValidValues []string       // List of valid values for parameter
}

// A path pattern represents a possible path for a given action.
type PathPattern struct {
	HttpMethod string         // "GET", "POST", "PUT", "DELETE", ...
	Pattern    string         // Actual pattern, e.g. "/clouds/%s/instances/%s"
	Variables  []string       // Pattern variable names in order of appearance in pattern, e.g. "cloud_id", "id"
	Regexp     *regexp.Regexp // Regexp used to match href and capture variable values, e.g. "/clouds/([^/]+)/instances/([^/])+"
}

// Url returns a URL to the action given a set of values that can be used to substitute the action
// paths pattern variables. This method tries to use a many variables as possible so that
// "the longest" path gets used. So if for example an action has the patterns "/instances/:id" and
// "/clouds/:cloud_id/instances/:id" and both the :cloud_id and :id variable values are given as
// parameter, the method returns a URL built from substituting the values of the later (longer) path.
// The method returns an error in case no path pattern can have all its variables subsituted.
func (a *Action) Url(vars []*PathVariable) (*ActionPath, error) {
	var candidates = make([]*ActionPath, len(a.PathPatterns))
	var allMissing = []string{}
	var j = 0
	for _, p := range a.PathPatterns {
		var path, names = p.Substitute(vars)
		if path == "" {
			allMissing = append(allMissing, names...)
		} else {
			candidates[j] = &ActionPath{path, p.HttpMethod, len(names)}
			j += 1
		}
	}
	if j == 0 {
		return nil, fmt.Errorf("Missing variables to instantiate action URL, one or more of the following variables are needed: %s",
			strings.Join(allMissing, ", "))
	}
	candidates = candidates[:j]
	sort.Sort(ByWeight(candidates))
	return candidates[0], nil
}

// A match built from a path pattern and given variable values
type ActionPath struct {
	Path       string // Actual path, e.g. "/clouds/1/instances/42"
	HttpMethod string // HTTP method
	weight     int    // Match relevance, i.e. number of variables consumed to produce match value
}

// Substitute attemps to substitute the path pattern variables with the given values.
// - If the substitution succeeds, it returns the resulting path and the list of variable names
//   that were used to build it.
// - If the substitution fails, it returns an empty string and the list of variable names that are
//   missing from the list of given values.
func (p *PathPattern) Substitute(vars []*PathVariable) (string, []string) {
	var values = make([]interface{}, len(p.Variables))
	var missing = []string{}
	var used = []string{}
	for i, n := range p.Variables {
		for _, v := range vars {
			if v.Name == n {
				values[i] = v.Value
				used = append(used, n)
				break
			}
		}
		if values[i] == nil {
			missing = append(missing, n)
		}
	}
	if len(missing) > 0 {
		return "", missing
	}
	return fmt.Sprintf(p.Pattern, values...), used
}

// A path variable consists of a name and value
type PathVariable struct {
	Name  string // e.g. "cloud_id"
	Value string // e.g. "1"
}

// Make it possible to sort path patterns by length
type ByLen []*PathPattern

func (b ByLen) Len() int           { return len(b) }
func (b ByLen) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByLen) Less(i, j int) bool { return len(b[i].Pattern) > len(b[j].Pattern) }

// Make it possible to sort path match by weight, from heaviest to lightest
type ByWeight []*ActionPath

func (b ByWeight) Len() int           { return len(b) }
func (b ByWeight) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByWeight) Less(i, j int) bool { return b[i].weight > b[j].weight }
