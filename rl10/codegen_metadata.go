//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ praxisgen -metadata=rl10/docs/api -output=rl10 -pkg=rl10 -target=unversioned -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rl10

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"DebugCookbookPath": &metadata.Resource{
		Name:        "DebugCookbookPath",
		Description: `Manipulate debug cookbook directory location`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Retrieve debug cookbook directory location`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/debug/cookbook",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/debug/cookbook`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Set debug cookbook directory location`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
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
				APIParams: []*metadata.ActionParam{
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
						HTTPMethod: "DELETE",
						Pattern:    "/rll/debug/cookbook",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/debug/cookbook`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"DockerControl": &metadata.Resource{
		Name:        "DockerControl",
		Description: `Manipulate the Docker integration in RightLink 10`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Show Docker integration features`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/docker/control",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/docker/control`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Enable/disable Docker integration features`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/rll/docker/control",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/docker/control`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "docker_host",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "enable_docker",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"none", "tags", "all"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "docker_host",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "enable_docker",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"none", "tags", "all"},
					},
				},
			},
		},
	},
	"Env": &metadata.Resource{
		Name:        "Env",
		Description: `Manipulate global script environment variables`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Retrieve all environment variables`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/env",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/env`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Retrieve environment variable value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/env/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/env/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Set environment variable value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
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
				APIParams: []*metadata.ActionParam{
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
						HTTPMethod: "DELETE",
						Pattern:    "/rll/env/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/env/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"LoginControl": &metadata.Resource{
		Name:        "LoginControl",
		Description: `Manipulate login policy settings`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Show login policy features`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/login/control",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/login/control`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Enable/disable login policy features`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/rll/login/control",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/login/control`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "enable_login",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"on", "off", "compat"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "enable_login",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"on", "off", "compat"},
					},
				},
			},
		},
	},
	"Proc": &metadata.Resource{
		Name:        "Proc",
		Description: `List of process variables, such as version, identity, and protocol_version`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List all process variables`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/proc",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/proc`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Retrieve process variable value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/proc/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/proc/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Set process variable value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/rll/proc/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/proc/([^/]+)`),
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
				APIParams: []*metadata.ActionParam{
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
		},
	},
	"Rl10": &metadata.Resource{
		Name:        "Rl10",
		Description: `Miscellaneous RightLink 10 local requests`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "upgrade",
				Description: `Relaunch the RightLink process using a specified binary`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
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
				APIParams: []*metadata.ActionParam{
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
						HTTPMethod: "POST",
						Pattern:    "/rll/run/recipe",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/run/recipe`),
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
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments",
						Description: `Script argument values`,
						Type:        "map[string]interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
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
						HTTPMethod: "POST",
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
				APIParams: []*metadata.ActionParam{
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
		Description: `Manipulate the TSS proxy (this is deprecated, please use the /rll/tss/control resource)`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "get_hostname",
				Description: `Get the TSS hostname to proxy (deprecated, RL10 knows the hostname)`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/tss/hostname",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/tss/hostname`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "put_hostname",
				Description: `Set the TSS hostname to proxy (deprecated, RL10 knows the hostname)`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
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
				APIParams: []*metadata.ActionParam{
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
	"TSSControl": &metadata.Resource{
		Name:        "TSSControl",
		Description: `Manipulate monitoring (TSS) settings`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Show monitoring features`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/tss/control",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/tss/control`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Enable/disable monitoring features`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/rll/tss/control",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/tss/control`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "enable_monitoring",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"false", "true", "none", "util", "extra", "all"},
					},
					&metadata.ActionParam{
						Name:        "tss_id",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "enable_monitoring",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"false", "true", "none", "util", "extra", "all"},
					},
					&metadata.ActionParam{
						Name:        "tss_id",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "put_control",
				Description: `Control the TSS monitoring (deprecated, use the update action)`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/rll/tss/control",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/tss/control`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "enable_monitoring",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "tss_id",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "enable_monitoring",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "tss_id",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"TSSPlugin": &metadata.Resource{
		Name:        "TSSPlugin",
		Description: `TSS Custom Plugins`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Get TSS plugins list`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/tss/exec",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/rll/tss/exec`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Add new TSS custom plugin`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/rll/tss/exec/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/tss/exec/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "executable",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "executable",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Get TSS plugin info`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/rll/tss/exec/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/tss/exec/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update TSS custom plugin`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/rll/tss/exec/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/tss/exec/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "executable",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "executable",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete TSS plugin info`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/rll/tss/exec/%s",
						Variables:  []string{"name"},
						Regexp:     regexp.MustCompile(`/rll/tss/exec/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
}
