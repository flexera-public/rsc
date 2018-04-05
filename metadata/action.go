package metadata

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Action represents a resource action.
type Action struct {
	Name         string
	Description  string
	PathPatterns []*PathPattern // Action path patterns and
	APIParams    []*ActionParam // Actual API request parameters
	CommandFlags []*ActionParam // Parameters initialized via command lines, these correspond to the leaves of all APIParams.
	Payload      string         // Name of payload type, only set for basic types
}

// PathParamNames returns the names of the action path parameters sorted alphabetically.
func (a *Action) PathParamNames() []string {
	return a.paramsByLocation(PathParam)
}

// QueryParamNames returns the names of the action query parameters sorted alphabetically.
func (a *Action) QueryParamNames() []string {
	return a.paramsByLocation(QueryParam)
}

// PayloadParamNames returns the names of the action payload parameters sorted alphabetically.
func (a *Action) PayloadParamNames() []string {
	return a.paramsByLocation(PayloadParam)
}

// MatchHref returns true if the given href matches one of the action's href patterns exactly
func (a *Action) MatchHref(href string) bool {
	hrefs := []string{href, href + "/"}
	for _, pattern := range a.PathPatterns {
		for _, href := range hrefs {
			indices := pattern.Regexp.FindStringIndex(href)
			// it is only an exact match if the pattern matches from the beginning to the length of the string
			if indices != nil && indices[0] == 0 && indices[1] == len(href) {
				return true
			}
		}
	}
	return false
}

// paramsByLocation is a helper method that returns the names of the parameters at the given
// location (path, query string or payload).
func (a *Action) paramsByLocation(loc Location) []string {
	var res []string
	for _, p := range a.APIParams {
		if p.Location == loc {
			res = append(res, p.Name)
		}
	}
	sort.Strings(res)
	return res
}

// Location indicates a parameter location (i.e. URL path, query string or request body).
type Location int

// Possible param location
const (
	PathParam Location = iota
	QueryParam
	PayloadParam
)

// ActionParam represents a resource action parameters.
type ActionParam struct {
	Name        string         // Param name
	Description string         // Param description
	Type        string         // Param type, one of "string", "[]string", "int", "[]int" or "map[string]string"
	Location    Location       // Param location, i.e. path, query or payload
	Mandatory   bool           // Whether parameter is mandatory
	NonBlank    bool           // Whether parameter value can be blank
	Regexp      *regexp.Regexp // Regular expression used to validate parameter values
	ValidValues []string       // List of valid values for parameter
}

// PathPattern represents a possible path for a given action.
type PathPattern struct {
	HTTPMethod string         // "GET", "POST", "PUT", "DELETE", ...
	Pattern    string         // Actual pattern, e.g. "/clouds/%s/instances/%s"
	Variables  []string       // Pattern variable names in order of appearance in pattern, e.g. "cloud_id", "id"
	Regexp     *regexp.Regexp // Regexp used to match href and capture variable values, e.g. "/clouds/([^/]+)/instances/([^/])+"
}

// URL returns a URL to the action given a set of values that can be used to substitute the action
// paths pattern variables. This method tries to use a many variables as possible so that
// "the longest" path gets used. So if for example an action has the patterns "/instances/:id" and
// "/clouds/:cloud_id/instances/:id" and both the :cloud_id and :id variable values are given as
// parameter, the method returns a URL built from substituting the values of the later (longer) path.
// The method returns an error in case no path pattern can have all its variables subsituted.
func (a *Action) URL(vars []*PathVariable) (*ActionPath, error) {
	candidates := make([]*ActionPath, len(a.PathPatterns))
	allMissing := []string{}
	j := 0
	for _, p := range a.PathPatterns {
		path, names := p.Substitute(vars)
		if path == "" {
			allMissing = append(allMissing, names...)
		} else {
			candidates[j] = &ActionPath{path, p.HTTPMethod, len(names)}
			j++
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

// ActionPath is a match built from a path pattern and given variable values.
type ActionPath struct {
	Path       string // Actual path, e.g. "/clouds/1/instances/42"
	HTTPMethod string // HTTP method
	weight     int    // Match relevance, i.e. number of variables consumed to produce match value
}

// Substitute attemps to substitute the path pattern variables with the given values.
// - If the substitution succeeds, it returns the resulting path and the list of variable names
//   that were used to build it.
// - If the substitution fails, it returns an empty string and the list of variable names that are
//   missing from the list of given values.
func (p *PathPattern) Substitute(vars []*PathVariable) (string, []string) {
	values := make([]interface{}, len(p.Variables))
	var missing []string
	var used []string
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

// PathVariable consists of a name and value.
type PathVariable struct {
	Name  string // e.g. "cloud_id"
	Value string // e.g. "1"
}

// ByLen makes it possible to sort path patterns by length.
type ByLen []*PathPattern

func (b ByLen) Len() int           { return len(b) }
func (b ByLen) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByLen) Less(i, j int) bool { return len(b[i].Pattern) > len(b[j].Pattern) }

// ByWeight makes it possible to sort path match by weight, from heaviest to lightest.
type ByWeight []*ActionPath

func (b ByWeight) Len() int           { return len(b) }
func (b ByWeight) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByWeight) Less(i, j int) bool { return b[i].weight > b[j].weight }
