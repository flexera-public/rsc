package metadata

import (
	"fmt"
	"sort"
	"strings"
)

// Resource command
type Resource struct {
	Name        string
	Description string
	Actions     []*Action
	Identifier  string
	Links       map[string]string
}

// ExtractVariables takes a resource href and extracts the variables (cloud_id, id etc.) that
// make it up. It does that by matching the href against all the resource action path patterns
// and finding the longest one that matches.
func (r *Resource) ExtractVariables(href string) ([]*PathVariable, error) {
	matches := r.findMatches(href)
	if len(matches) == 0 && !strings.HasPrefix(href, "/") {
		href = "/" + href
		matches = r.findMatches(href)
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("Href '%s' does not match any action path of resource %s", href, r.Name)
	}
	sort.Sort(ByLen(matches))
	pattern := matches[0]
	submatches := pattern.Regexp.FindAllStringSubmatch(href, -1)
	if len(submatches) == 0 {
		return []*PathVariable{}, nil
	}
	if len(submatches) > 1 {
		// Not right... there should only be one
		return nil, fmt.Errorf("Ambiguous href '%s', matches multiple path patterns of %s actions",
			href, r.Name)
	}
	submatch := submatches[0]
	variables := make([]*PathVariable, len(submatch)-1)
	for i, v := range submatch[1:] {
		variables[i] = &PathVariable{Name: pattern.Variables[i], Value: v}

	}
	return variables, nil
}

// GetAction returns the action with the given name, returns nil if none is found.
func (r *Resource) GetAction(name string) *Action {
	for _, a := range r.Actions {
		if a.Name == name {
			return a
		}
	}
	return nil
}

// HasLink returns whether the resource has a link with the given name.
func (r *Resource) HasLink(name string) bool {
	for n, _ := range r.Links {
		if n == name {
			return true
		}
	}
	return false
}

// Find paths that match given href
func (r *Resource) findMatches(href string) []*PathPattern {
	var matches []*PathPattern
	for _, action := range r.Actions {
		for _, pattern := range action.PathPatterns {
			if pattern.Regexp.MatchString(href) || pattern.Regexp.MatchString(href+"/") {
				matches = append(matches, pattern)
			}
		}
	}
	return matches
}
