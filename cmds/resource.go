package cmds

import (
	"fmt"
	"sort"
)

// Resource command
type ResourceCmd struct {
	Name        string
	Description string
	Actions     []*ActionCmd
}

// ExtractVariables takes a resource href and extracts the variables (cloud_id, id etc.) that
// make it up. It does that by matching the href against all the resource action path patterns
// and finding the longest one that matches.
func (r *ResourceCmd) ExtractVariables(href string) ([]*PathVariable, error) {
	var matches = []*PathPattern{}
	for _, action := range r.Actions {
		for _, pattern := range action.PathPatterns {
			if pattern.Regexp.MatchString(href) {
				matches = append(matches, pattern)
			}
		}
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("Href does not match any action path of resource %s", r.Name)
	}
	sort.Sort(ByLen(matches))
	var pattern = matches[0]
	var values = pattern.Regexp.FindAllString(href, -1)
	var variables = make([]*PathVariable, len(values))
	for i, v := range values {
		variables[i] = &PathVariable{Name: pattern.Variables[i], Value: v}
	}
	return variables, nil
}
