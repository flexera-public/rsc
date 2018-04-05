//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ praxisgen -metadata=ss/ssd/restful_doc -output=ss/ssd -pkg=ssd -target=1.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssd

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"Schedule": &metadata.Resource{
		Name: "Schedule",
		Description: `A Schedule represents a recurring period during which a CloudApp should be running. It must have a unique name and an optional description. The recurrence rules follow the [Recurrence Rule format](https://tools.ietf.org/html/rfc5545#section-3.8.5.3).
Multiple Schedules can be associated with a Template when published to the Catalog. Users will be able to launch the resulting CloudApp with one of the associated schedule. Updating or deleting a Schedule will not affect CloudApps that were published with that Schedule.`,
		Identifier: "application/vnd.rightscale.self_service.schedule",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the schedules available in Designer.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/designer/collections/%s/schedules",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/schedules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show detailed information about a given Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/designer/collections/%s/schedules/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/schedules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/schedules",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/schedules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The unique name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence[hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence[minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence[rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence[hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence[minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence[rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The unique name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence",
						Description: `When to start a CloudApp`,
						Type:        "*Recurrence",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence",
						Description: `When to stop a CloudApp`,
						Type:        "*Recurrence",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update one or more attributes of an existing Schedule.
Note: updating a Schedule in Designer doesn't update it in the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/api/designer/collections/%s/schedules/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/schedules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The unique name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence[hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence[minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence[rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence[hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence[minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence[rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The unique name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_recurrence",
						Description: `When to start a CloudApp`,
						Type:        "*Recurrence",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "stop_recurrence",
						Description: `When to stop a CloudApp`,
						Type:        "*Recurrence",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Delete a Schedule from the system.
Note: deleting a Schedule from Designer doesn't remove it from the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/designer/collections/%s/schedules/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/schedules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "multi_delete",
				Description: `Delete multiple Schedules from the system in bulk.
Note: deleting a Schedule from Designer doesn't remove it from the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/designer/collections/%s/schedules",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/schedules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The IDs of the Schedules to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The IDs of the Schedules to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"Template": &metadata.Resource{
		Name: "Template",
		Description: `A Template represent a CloudApplication Template (CAT) that has been uploaded to this design collection.
For information on the syntax of a CAT file, please see the [CAT File Language Reference](http://docs.rightscale.com/ss/reference/ss_CAT_file_language.html) on the RightScale Docs
site.
A CAT file is compiled by Self-Service to make it ready for publication and subsequent launch by users. To
test your CAT file syntax, you can call the compile action with the source content. In order to
Publish your CAT to the Catalog where users can launch it, it must be uploaded to Designer first, and then
published to the Catalog.
CAT files are uniquely identified by the name of the CloudApplication, which is specified as the "name"
attribute inside of a CAT file.`,
		Identifier: "application/vnd.rightscale.self_service.template",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the templates available in Designer along with some general details.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/designer/collections/%s/templates",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by name, syntax is ["name==foo"]`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of template IDs to retrieve`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
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
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by name, syntax is ["name==foo"]`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of template IDs to retrieve`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
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
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show detailed information about a given Template. Use the views specified below for more information.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/designer/collections/%s/templates/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
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
				APIParams: []*metadata.ActionParam{
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
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new Template by uploading its content to Designer.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/templates",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `Multipart File Upload`,
						Type:        "file",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `Multipart File Upload`,
						Type:        "*rsapi.FileUpload",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "create_from_compilation",
				Description: `Create a new Template from a previously compiled CAT.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/templates/actions/create_from_compilation",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/actions/create_from_compilation`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compilation_href",
						Description: `The href of the compilation`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filename",
						Description: `The filename of the template`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compilation_href",
						Description: `The href of the compilation`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filename",
						Description: `The filename of the template`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update the content of an existing Template (a Template with the same "name" value in the CAT).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/designer/collections/%s/templates/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `Multipart File Upload`,
						Type:        "file",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `Multipart File Upload`,
						Type:        "*rsapi.FileUpload",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update_from_compilation",
				Description: `Update a Template from a previously compiled CAT.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/templates/%s/actions/update_from_compilation",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/([^/]+)/actions/update_from_compilation`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compilation_href",
						Description: `The href of the compilation`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filename",
						Description: `The filename of the template`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compilation_href",
						Description: `The href of the compilation`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filename",
						Description: `The filename of the template`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete a Template from the system. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/designer/collections/%s/templates/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete multiple Templates from the system in bulk. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/designer/collections/%s/templates",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The IDs of the Template to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The IDs of the Template to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "download",
				Description: `Download the source of a Template.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/designer/collections/%s/templates/%s/download",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/([^/]+)/download`),
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
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_version",
						Description: `The API version (only valid value is currently "1.0")`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "compile",
				Description: `Compile the Template, but don't save it to Designer. Useful for debugging a CAT file while you are still authoring it.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/templates/actions/compile",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/actions/compile`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `The source of the CAT as a string`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `The source of the CAT as a string`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "dependencies",
				Description: `Lists the Templates which the provided CAT source or Template directly or indirectly depend upon`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/templates/actions/dependencies",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/actions/dependencies`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `The source of the CAT as a string, mutually exclusive with template_id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_id",
						Description: `The id of the template, mutually exclusive with source, have predecedence over "source" if both parameters are given`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source",
						Description: `The source of the CAT as a string, mutually exclusive with template_id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_id",
						Description: `The id of the template, mutually exclusive with source, have predecedence over "source" if both parameters are given`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "dependents",
				Description: `List the Dependents templates available in Designer for the given package, even if no template actually define the package.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/designer/collections/%s/templates/actions/dependents",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/actions/dependents`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "package",
						Description: `The path of the Package to which lists the dependents`,
						Type:        "string",
						Location:    metadata.QueryParam,
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
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "package",
						Description: `The path of the Package to which lists the dependents`,
						Type:        "string",
						Location:    metadata.QueryParam,
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
			},

			&metadata.Action{
				Name:        "publish",
				Description: `Publish the given Template to the Catalog so that users can launch it.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/templates/actions/publish",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/actions/publish`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of a Template to publish`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Optionally override the Template long description used mostly for designers.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Optionally override the Template name for display in the Catalog`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "overridden_application_href",
						Description: `If re-publishing, you must specify the href of the Application in the Catalog that is being overridden`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedule_required",
						Description: `Whether the CloudApp requires a schedule to be provided at launch time. If set to false, allows user to pick from '24/7' schedule when launching in the UI`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Optionally override the Template short description for display in the Catalog`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of a Template to publish`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Optionally override the Template long description used mostly for designers.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Optionally override the Template name for display in the Catalog`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "overridden_application_href",
						Description: `If re-publishing, you must specify the href of the Application in the Catalog that is being overridden`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedule_required",
						Description: `Whether the CloudApp requires a schedule to be provided at launch time. If set to false, allows user to pick from '24/7' schedule when launching in the UI`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules",
						Description: `Schedules available to users when launching the application`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Optionally override the Template short description for display in the Catalog`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "unpublish",
				Description: `Remove a publication from the Catalog by specifying its associated Template.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/designer/collections/%s/templates/actions/unpublish",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/api/designer/collections/([^/]+)/templates/actions/unpublish`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Template to unpublish`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Template to unpublish`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
}
