//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated
// Mar 2, 2015 at 9:07pm (PST)
// Command:
// $ praxisgen -metadata=ssd/restful_doc -output=ssd -pkg=ssd -target=1.0 -client=Api
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
		Name:        "Schedule",
		Description: `A Schedule represents a recurring period during which a CloudApp should be running. It must have a unique name and an optional description. The recurrence rules follow the [Recurrence Rule format](https://tools.ietf.org/html/rfc5545#section-3.8.5.3).`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the schedules available in Designer.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/collections/%s/schedules",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/schedules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show detailed information about a given Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/collections/%s/schedules/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/schedules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Schedule`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/collections/%s/schedules",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/schedules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update one or more attributes of an existing Schedule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PATCH",
						Pattern:    "/collections/%s/schedules/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/schedules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Schedule to update`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete a Schedule from the system.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/collections/%s/schedules/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/schedules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Schedule to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete multiple Schedules from the system in bulk.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/collections/%s/schedules",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/schedules`),
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
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
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
		Name:        "Template",
		Description: `A Template represent a CloudApplication Template (CAT) that has been uploaded to this design collection.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the templates available in Designer along with some general details.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/collections/%s/templates",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of template IDs to retrieve`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
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
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show detailed information about a given Template. Use the views specified below for more information.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/collections/%s/templates/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates/([^/]+)`),
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
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Template`,
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
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new Template by uploading its content to Designer.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/collections/%s/templates",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update the content of an existing Template (a Template with the same "name" value in the CAT).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/collections/%s/templates/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Template to update`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete a Template from the system. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/collections/%s/templates/%s",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Template to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete multiple Templates from the system in bulk. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/collections/%s/templates",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates`),
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
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
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
						HttpMethod: "GET",
						Pattern:    "/collections/%s/templates/%s/download",
						Variables:  []string{"collection_id", "id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates/([^/]+)/download`),
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
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the Template to download`,
						Type:        "string",
						Location:    metadata.PathParam,
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
						HttpMethod: "POST",
						Pattern:    "/collections/%s/templates/actions/compile",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates/actions/compile`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "publish",
				Description: `Publish the given Template to the Catalog so that users can launch it.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/collections/%s/templates/actions/publish",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates/actions/publish`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "unpublish",
				Description: `Remove a publication from the Catalog by specifying its associated Template.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/collections/%s/templates/actions/unpublish",
						Variables:  []string{"collection_id"},
						Regexp:     regexp.MustCompile(`/collections/([^/]+)/templates/actions/unpublish`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "collection_id",
						Description: `The collection ID (currently the account ID)`,
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
