//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ praxisgen -metadata=rl10/docs/api -output=rl10 -pkg=rl10 -target=unversioned -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rl10 // import "gopkg.in/rightscale/rsc.v1/rl10"

import (
	"regexp"

	"gopkg.in/rightscale/rsc.v1/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"DebugCookbookPath": &metadata.Resource{
		Name:        "DebugCookbookPath",
		Description: `Manipulate debug cookbook directory location`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Retrieve debug cookbook directory location`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/rll/debug/cookbook",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/debug/cookbook`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Set debug cookbook directory location`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/rll/debug/cookbook",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/debug/cookbook`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "path",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "path",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Remove debug cookbook directory location`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/rll/debug/cookbook",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/debug/cookbook`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Env": &metadata.Resource{
		Name:        "Env",
		Description: `Manipulate global script environment variables`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Retrieve all environment variables`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/rll/env",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/env`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Retrieve environment variable value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/rll/env/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/env/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Set environment variable value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/rll/env/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/env/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string",
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete environment variable`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/rll/env/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/env/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Proc": &metadata.Resource{
		Name:        "Proc",
		Description: `List of process variables, such as version, identity, and protocol_version`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List all process variables`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/rll/proc",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/proc`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Retrieve process variable value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/rll/proc/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/proc/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Rl10": &metadata.Resource{
		Name:        "Rl10",
		Description: `Miscellaneous RightLink 10 local requests`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "upgrade",
				Description: `Relaunch the RightLink process using a specified binary`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/rll/upgrade",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/upgrade`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "exec",
						Description: `Absolute path to binary`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "exec",
						Description: `Absolute path to binary`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "run_recipe",
				Description: `Run git-based scripts (as recipes) synchronously`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/rll/run/recipe",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/run/recipe`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "json",
						Description: `JSON hash of "name": "value" pairs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "recipe",
						Description: `Name of recipe`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "json",
						Description: `JSON hash of "name": "value" pairs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "recipe",
						Description: `Name of recipe`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "run_right_script",
				Description: `Run RightScripts synchronously`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/rll/run/right_script",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/run/right_script`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments",
						Description: `Script argument values`,
						Type:        "map",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "right_script",
						Description: `Name of script`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "right_script_id",
						Description: `Id of script`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments",
						Description: `Script argument values`,
						Type:        "map[string]interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "right_script",
						Description: `Name of script`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "right_script_id",
						Description: `Id of script`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"TSS": &metadata.Resource{
		Name:        "TSS",
		Description: `Manipulate the TSS proxy`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "get_hostname",
				Description: `Get the TSS hostname to proxy`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/rll/tss/hostname",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/tss/hostname`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "put_hostname",
				Description: `Set the TSS hostname to proxy`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/rll/tss/hostname",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/tss/hostname`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "hostname",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "hostname",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
}
