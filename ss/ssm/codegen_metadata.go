//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated
// Feb 27, 2015 at 5:28pm (PST)
// Command:
// $ praxisgen -metadata=ssm/restful_doc -output=ssm -pkg=ssm -target=1.0 -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssm

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"Execution": &metadata.Resource{
		Name:        "Execution",
		Description: `An Execution is a launched instance of a CloudApp`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List information about the Executions, or use a filter to only return certain Executions. A view can be used for various levels of detail.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/executions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by status, syntax is ["status==running"]`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of execution IDs to retrieve`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by status, syntax is ["status==running"]`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of execution IDs to retrieve`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "*string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show details for a given Execution. A view can be used for various levels of detail.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/executions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded", "source"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "*string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded", "source"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new execution from a CAT, a compiled CAT, an Application in the Catalog, or a Template in Designer`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `No description provided for delete.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/projects/%s/executions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "force",
						Description: `Force delete execution, bypass state validation so that non term can deleted.`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `ID of execution to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "force",
						Description: `Force delete execution, bypass state validation so that non term can deleted.`,
						Type:        "*bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `ID of execution to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete several executions from the database. Note: if an execution has not successfully been terminated, there may still be associated cloud resources running.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/projects/%s/executions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "force",
						Description: `Force delete execution, bypass state validation so that non term can deleted.`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "force",
						Description: `Force delete execution, bypass state validation so that non term can deleted.`,
						Type:        "*bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "download",
				Description: `Download the CAT source for the execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/executions/%s/download",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/([^/]+)/download`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_version",
						Description: `The API version (only valid value is currently "1.0")`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_version",
						Description: `The API version (only valid value is currently "1.0")`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "launch",
				Description: `Launch an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/%s/actions/launch",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/([^/]+)/actions/launch`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "start",
				Description: `Start an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/%s/actions/start",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/([^/]+)/actions/start`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "stop",
				Description: `Stop an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/%s/actions/stop",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/([^/]+)/actions/stop`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "terminate",
				Description: `Terminate an Execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/%s/actions/terminate",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/([^/]+)/actions/terminate`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the execution`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_launch",
				Description: `Launch several executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/actions/launch",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/actions/launch`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to launch`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to launch`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_start",
				Description: `Start several executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/actions/start",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/actions/start`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to start`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to start`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_stop",
				Description: `Stop several executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/actions/stop",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/actions/stop`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to stop`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to stop`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_terminate",
				Description: `Terminate several executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/executions/actions/terminate",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/executions/actions/terminate`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to terminate`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `List of execution IDs to terminate`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"Notification": &metadata.Resource{
		Name:        "Notification",
		Description: `The Notification resource represents a system notification that an action has occurred`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the most recent 50 Notifications. Use the filter parameter to specify specify Executions.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/notifications",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/notifications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The Notification IDs to return`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The Notification IDs to return`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Get details for a specific Notification`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/notifications/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/notifications/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Notification ID to return`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Notification ID to return`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"Operation": &metadata.Resource{
		Name:        "Operation",
		Description: `Operations represent actions that can be taken on an Execution.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Get the list of 50 most recent Operations (usually filtered by Execution).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/operations",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/operations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution ID or status`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `IDs of operations to filter on`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `The maximum number of operations to retrieve. The maximum (and default) limit is 50.If a limit of more than 50 is specified, only 50 operations will be returned`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by Execution ID or status`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `IDs of operations to filter on`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `The maximum number of operations to retrieve. The maximum (and default) limit is 50.If a limit of more than 50 is specified, only 50 operations will be returned`,
						Type:        "*int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "*string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Get the details for a specific Operation`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/operations/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/operations/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Operation to get details for`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Operation to get details for`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Optional view to return`,
						Type:        "*string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "expanded"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Trigger an Operation to run by specifying the Execution ID and the name of the Operation.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/operations",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/operations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"ScheduledAction": &metadata.Resource{
		Name:        "ScheduledAction",
		Description: `ScheduledActions describe a set of timed occurrences for an action to be run (at most once per day).`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List ScheduledAction resources in the project. The list can be filtered to a given execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/scheduled_actions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_actions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by execution id or execution creator (user) id.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by execution id or execution creator (user) id.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Retrieve given ScheduledAction resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/scheduled_actions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_actions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new ScheduledAction resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/scheduled_actions",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_actions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "patch",
				Description: `Update one or more ScheduledAction properties. If the ScheduledAction has the mandatory attribute set to true, the 'force' flag must be set in order to modify it. All ScheduledActions created through the UI are set to 'mandatory' by default. When the 'recurrence' is updated, the 'next_occurrence' will be modified accordingly unless it's also specified.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PATCH",
						Pattern:    "/projects/%s/scheduled_actions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_actions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction to update.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction to update.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete a ScheduledAction. If the ScheduledAction has the mandatory attribute set to true, the 'force' flag must be set in order to delete it.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/projects/%s/scheduled_actions/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_actions/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction to delete.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction to delete.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "skip",
				Description: `Skips the requested number of ScheduledAction occurrences. If no count is provided, one occurrence is skipped. On success, the next_occurrence view of the updated ScheduledAction is returned.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/scheduled_actions/%s/actions/skip",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_actions/([^/]+)/actions/skip`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledAction.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"ScheduledOperation": &metadata.Resource{
		Name:        "ScheduledOperation",
		Description: `ScheduledOperations describe a set of timed occurrences for an operation to be run (at most once per day).`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List ScheduledOperation resources in the project. The list can be filtered to a given execution.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/scheduled_operations",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_operations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by execution id or execution creator (user) id.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by execution id or execution creator (user) id.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Retrieve given ScheduledOperation resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/projects/%s/scheduled_operations/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_operations/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new ScheduledOperation resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/scheduled_operations",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_operations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "patch",
				Description: `Update one or more ScheduledOperation properties. If the ScheduledOperation has the mandatory attribute set to true, the 'force' flag must be set in order to modify it. All ScheduledOperations created through the UI are set to 'mandatory' by default. When the 'recurrence' is updated, the 'next_occurrence' will be modified accordingly unless it's also specified.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PATCH",
						Pattern:    "/projects/%s/scheduled_operations/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_operations/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation to update.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation to update.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete a ScheduledOperation. If the ScheduledOperation has the mandatory attribute set to true, the 'force' flag must be set in order to delete it.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/projects/%s/scheduled_operations/%s",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_operations/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation to delete.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation to delete.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "skip",
				Description: `Skips the requested number of ScheduledOperation occurrences. If no count is provided, one occurrence is skipped. On success, the next_occurrence view of the updated ScheduledOperation is returned.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/projects/%s/scheduled_operations/%s/actions/skip",
						Variables:  []string{"project_id", "id"},
						Regexp:     regexp.MustCompile(`/projects/([^/]+)/scheduled_operations/([^/]+)/actions/skip`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The id of the ScheduledOperation.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_id",
						Description: `The project ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
}
