package metadata

import (
	"fmt"
	"sort"
)

// Resource command
type Resource struct {
	Name        string
	Description string
	Actions     []*Action
}

// ExtractVariables takes a resource href and extracts the variables (cloud_id, id etc.) that
// make it up. It does that by matching the href against all the resource action path patterns
// and finding the longest one that matches.
func (r *Resource) ExtractVariables(href string) ([]*PathVariable, error) {
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
	var submatches = pattern.Regexp.FindAllStringSubmatch(href, -1)
	if len(submatches) == 0 {
		return []*PathVariable{}, nil
	}
	if len(submatches) > 1 {
		// Not right... there should only be one
		return nil, fmt.Errorf("Ambiguous href '%s', matches multiple path patterns of %s actions",
			href, r.Name)
	}
	var submatch = submatches[0]
	var variables = make([]*PathVariable, len(submatch)-1)
	for i, v := range submatch[1:] {
		variables[i] = &PathVariable{Name: pattern.Variables[i], Value: v}

	}
	return variables, nil
}
