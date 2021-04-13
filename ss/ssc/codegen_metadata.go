//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ praxisgen -metadata=ss/ssc/restful_doc -output=ss/ssc -pkg=ssc -target=1.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssc // import "gopkg.in/rightscale/rsc.v9/ss/ssc"

import (
	"regexp"

	"gopkg.in/rightscale/rsc.v9/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"AccountPreference": &metadata.Resource{
		Name: "AccountPreference",
		Description: `The AccountPreference resource stores preferences that apply account-wide, such as UI customization settings and other settings.
The Self-Service portal uses some of these preferences in the portal itself, and this resource allows you to extend the settings
to use in your own integration.`,
		Identifier: "application/vnd.rightscale.self_service.account_preference",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "created_by",
				FieldName: "CreatedBy",
				FieldType: "*User",
			},

			&metadata.Attribute{
				Name:      "group_name",
				FieldName: "GroupName",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "name",
				FieldName: "Name",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},

			&metadata.Attribute{
				Name:      "value",
				FieldName: "Value",
				FieldType: "string",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the AccountPreferences for this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/account_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/account_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by group, so that only AccountPreferences belonging to that group are returned`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by group, so that only AccountPreferences belonging to that group are returned`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Get details for a particular AccountPreference`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/account_preferences/%s",
						Variables:  []string{"account_id", "name"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/account_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new AccountPreference or update an existing AccountPreference with the new value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/catalog/accounts/%s/account_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/account_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "group_name",
						Description: `The group to place this AccountPreference in. Any string value is accepted - the group does not need to exist`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name for the new AccountPreference or AccountPreference to update (note this is the key for this resource)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "value",
						Description: `The value to set for this AccountPreference`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "group_name",
						Description: `The group to place this AccountPreference in. Any string value is accepted - the group does not need to exist`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name for the new AccountPreference or AccountPreference to update (note this is the key for this resource)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "value",
						Description: `The value to set for this AccountPreference`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete an AccountPreference`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/catalog/accounts/%s/account_preferences/%s",
						Variables:  []string{"account_id", "name"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/account_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Application": &metadata.Resource{
		Name: "Application",
		Description: `An Application is an element in the Catalog that can be launched by users. Applications are generally created by uploading CAT
files to the Designer and publishing them to the Catalog, though they can also be created via API calls to the Catalog directly without
going through Designer. If an Application was created from Designer through the publish action, it contains a link back to the Template
resource in Designer.
In the Self-Service portal, an Application is equivalent to an item in the Catalog. Most users have access to these Application resources
and can launch them to create Executions in the Manager application.`,
		Identifier: "application/vnd.rightscale.self_service.application",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "compilation_href",
				FieldName: "CompilationHref",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "compiled_cat",
				FieldName: "CompiledCat",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "compiler_ver",
				FieldName: "CompilerVer",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "created_by",
				FieldName: "CreatedBy",
				FieldType: "*User",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "long_description",
				FieldName: "LongDescription",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "name",
				FieldName: "Name",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "parameters",
				FieldName: "Parameters",
				FieldType: "[]*Parameter",
			},

			&metadata.Attribute{
				Name:      "required_parameters",
				FieldName: "RequiredParameters",
				FieldType: "[]string",
			},

			&metadata.Attribute{
				Name:      "schedule_required",
				FieldName: "ScheduleRequired",
				FieldType: "bool",
			},

			&metadata.Attribute{
				Name:      "schedules",
				FieldName: "Schedules",
				FieldType: "[]*Schedule",
			},

			&metadata.Attribute{
				Name:      "short_description",
				FieldName: "ShortDescription",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "template_info",
				FieldName: "TemplateInfo",
				FieldType: "*TemplateInfo",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the Applications available in the specified Catalog.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of Application IDs to retrieve. If not specified, all are returned.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of Application IDs to retrieve. If not specified, all are returned.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show detailed information about a given Application.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/catalogs/%s/applications/%s",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications/([^/]+)`),
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
				Description: `Create a new Application in the Catalog.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/catalog/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compiled_cat[cat_parser_gem_version]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[compiler_ver]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[conditions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[definitions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[imports]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[long_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[mappings]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[operations]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[outputs]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[package]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[parameters]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[permissions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[resources]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[rs_ca_ver]",
						Description: ``,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[short_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[source]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Long description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
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
						Name:        "schedules[][created_from]",
						Description: `optional HREF of the Schedule resource used to create this schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][description]",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][name]",
						Description: `The name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Short description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `If created from a Template, the template href can be provided to maintain the relationship between the resources.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compiled_cat",
						Description: `The compiled source of the CAT file. This can be obtained by calling Template.compile or Template.show in the Designer application.`,
						Type:        "*CompiledCAT",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Long description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
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
						Type:        "[]*Schedule",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Short description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `If created from a Template, the template href can be provided to maintain the relationship between the resources.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update the content of an existing Application.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/catalog/catalogs/%s/applications/%s",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compiled_cat[cat_parser_gem_version]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[compiler_ver]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[conditions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[definitions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[imports]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[long_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[mappings]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[operations]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[outputs]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[package]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[parameters]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[permissions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[resources]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[rs_ca_ver]",
						Description: ``,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[short_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[source]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Long description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of application`,
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
						Name:        "schedules[][created_from]",
						Description: `optional HREF of the Schedule resource used to create this schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][description]",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][name]",
						Description: `The name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Short description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `A template href can be provided to maintain the relationship between the resources.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compiled_cat",
						Description: `The compiled source of the CAT file. This can be obtained by calling Template.compile or Template.show in the Designer application.`,
						Type:        "*CompiledCAT",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Long description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of application`,
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
						Type:        "[]*Schedule",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Short description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `A template href can be provided to maintain the relationship between the resources.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_update",
				Description: `Update the content of multiple Applications.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/catalog/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compiled_cat[cat_parser_gem_version]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[compiler_ver]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[conditions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[definitions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[imports]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[long_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[mappings]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[operations]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[outputs]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[package]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[parameters]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[permissions]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[dependency_hashes][]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[resources]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[rs_ca_ver]",
						Description: ``,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[short_description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "compiled_cat[source]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID to update`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Long description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of application`,
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
						Name:        "schedules[][created_from]",
						Description: `optional HREF of the Schedule resource used to create this schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][description]",
						Description: `An optional description that will help users understand the purpose of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][name]",
						Description: `The name of the Schedule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][hour]",
						Description: `The hour of day from 0 to 23.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][minute]",
						Description: `The minute from 0 to 59.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedules[][start_recurrence][rule]",
						Description: `A RRULE string describing the recurrence rule.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Short description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `A template href can be provided to maintain the relationship between the resources.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "compiled_cat",
						Description: `The compiled source of the CAT file. This can be obtained by calling Template.compile or Template.show in the Designer application.`,
						Type:        "*CompiledCAT",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID to update`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "long_description",
						Description: `Long description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of application`,
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
						Type:        "[]*Schedule",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "short_description",
						Description: `Short description of application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `A template href can be provided to maintain the relationship between the resources.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete an Application from the Catalog`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/catalog/catalogs/%s/applications/%s",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete multiple Applications from the Catalog`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/catalog/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The Application IDs to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The Application IDs to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "download",
				Description: `Download the underlying CAT source of an Application.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/catalogs/%s/applications/%s/download",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications/([^/]+)/download`),
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
				Name:        "launch",
				Description: `Launches an Application by creating an Execution with ScheduledActions as needed to match the optional Schedule provided.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/catalog/catalogs/%s/applications/%s/actions/launch",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/catalogs/([^/]+)/applications/([^/]+)/actions/launch`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "defer_launch",
						Description: `Whether or not to defer launching the execution. Setting this value to true will keep the execution in not_started state until it is explicitly launched or the first scheduled start operation occurs.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `The description for the execution. The description of the Application will be used if none is provided.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "end_date",
						Description: `When the CloudApp should be automatically terminated.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name for the Execution. The Application name will be used if none is provided. This will be used as the name of the deployment (appended with a unique ID).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][name]",
						Description: `Name of configuration option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][type]",
						Description: `Type of configuration option.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"string", "number", "list"},
					},
					&metadata.ActionParam{
						Name:        "options[][value]",
						Description: `Configuration option value, a string, integer or array of strings depending on type`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedule_name",
						Description: `Name of the Schedule to use when launching. It must match one of the schedules attached to the Application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "defer_launch",
						Description: `Whether or not to defer launching the execution. Setting this value to true will keep the execution in not_started state until it is explicitly launched or the first scheduled start operation occurs.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `The description for the execution. The description of the Application will be used if none is provided.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "end_date",
						Description: `When the CloudApp should be automatically terminated.`,
						Type:        "*time.Time",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name for the Execution. The Application name will be used if none is provided. This will be used as the name of the deployment (appended with a unique ID).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options",
						Description: `The configuration options of the Execution. These are the values provided for the CloudApp parameters.`,
						Type:        "[]*ConfigurationOption",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "schedule_name",
						Description: `Name of the Schedule to use when launching. It must match one of the schedules attached to the Application`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"EndUser": &metadata.Resource{
		Name:        "EndUser",
		Description: ``,
		Identifier:  "application/vnd.rightscale.self_service.end_user",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "company",
				FieldName: "Company",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "first_name",
				FieldName: "FirstName",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "last_name",
				FieldName: "LastName",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "phone",
				FieldName: "Phone",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "timezone_name",
				FieldName: "TimezoneName",
				FieldType: "string",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Show all Self-Service Only End Users that belong to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/end_users",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/end_users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by user ID`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by user ID`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Grant a user Self-Service Only End User access to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/catalog/accounts/%s/end_users",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/end_users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user_ids[]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user_ids",
						Description: `User IDs to add as SS End Users to this account`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/catalog/accounts/%s/end_users",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/end_users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user_ids[]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user_ids",
						Description: `User IDs to remove as SS End Users to this account`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "non_ss_users",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/end_users/available",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/end_users/available`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"NotificationRule": &metadata.Resource{
		Name: "NotificationRule",
		Description: `A notification rule describes which notification should be created
        when events occur in the system. Events may be generated when an
        execution status changes or when an operation fails for example.
        A rule has a source which can be a specific resource or a group of
        resources (described via a link-like syntax), a target which
        corresponds to a user (for now) and a minimum severity used to filter
        out events with lower severities.`,
		Identifier: "application/vnd.rightscale.self_service.notification_rule",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "account_id",
				FieldName: "AccountId",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "category",
				FieldName: "Category",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "min_severity",
				FieldName: "MinSeverity",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "priority",
				FieldName: "Priority",
				FieldType: "int",
			},

			&metadata.Attribute{
				Name:      "source",
				FieldName: "Source",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "target",
				FieldName: "Target",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List all notification rules, potentially filtering by a collection of resources.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/notification_rules",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/notification_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by category.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "source",
						Description: `List all notification rules where the target is the current user.
                          The list can be further filtered by notification source: either by
                          source type or by specific source.
                          * To retrieve all notification rules that apply to all executions use:
                            GET nofication_rules?source==/api/projects/1234/executions
                          * To retrieve all notification rules that apply to a specific execution use:
                            GET notification_rules?source==/api/projects/1234/executions/5678`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: true,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "targets",
						Description: `Comma separated list of target ids. Note, currently only "me" is allowed.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by category.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "source",
						Description: `List all notification rules where the target is the current user.
                          The list can be further filtered by notification source: either by
                          source type or by specific source.
                          * To retrieve all notification rules that apply to all executions use:
                            GET nofication_rules?source==/api/projects/1234/executions
                          * To retrieve all notification rules that apply to a specific execution use:
                            GET notification_rules?source==/api/projects/1234/executions/5678`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: true,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "targets",
						Description: `Comma separated list of target ids. Note, currently only "me" is allowed.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Create one notification rule for a specific target and source.
          The source must be unique in the scope of target and account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/catalog/accounts/%s/notification_rules",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/notification_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "category",
						Description: `The type of notification for the resource.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"lifecycle", "scheduled"},
					},
					&metadata.ActionParam{
						Name: "min_severity",
						Description: `The lowest level of notifications for the target to receive.
                    Setting this to "error" will result in only receiving error notifications,
                    whereas setting it to "info" will result in receiving both info and error notifications,
                    and setting it to "none" will result in not receiving any notifications.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						ValidValues: []string{"error", "info", "none"},
					},
					&metadata.ActionParam{
						Name: "source",
						Description: `The resource (or resource collection) that would trigger the notification.
                    "/api/manager/projects/1234/executions" refers to ALL executions in the project,
                    "/api/manager/projects/1234/executions/5678" refers to just one execution, and
                    "/api/manager/projects/1234/executions?filter[]=created_by==me" refers to executions
                    created by the submitting user. The source must be unique in the scope of target and account.
                    Note that at this time, "me" is the only supported target filter.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name: "target",
						Description: `The notification target (user) that the rule applies to.
                    Note that at this time, "me" is the only supported target.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "category",
						Description: `The type of notification for the resource.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"lifecycle", "scheduled"},
					},
					&metadata.ActionParam{
						Name: "min_severity",
						Description: `The lowest level of notifications for the target to receive.
                    Setting this to "error" will result in only receiving error notifications,
                    whereas setting it to "info" will result in receiving both info and error notifications,
                    and setting it to "none" will result in not receiving any notifications.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						ValidValues: []string{"error", "info", "none"},
					},
					&metadata.ActionParam{
						Name: "source",
						Description: `The resource (or resource collection) that would trigger the notification.
                    "/api/manager/projects/1234/executions" refers to ALL executions in the project,
                    "/api/manager/projects/1234/executions/5678" refers to just one execution, and
                    "/api/manager/projects/1234/executions?filter[]=created_by==me" refers to executions
                    created by the submitting user. The source must be unique in the scope of target and account.
                    Note that at this time, "me" is the only supported target filter.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name: "target",
						Description: `The notification target (user) that the rule applies to.
                    Note that at this time, "me" is the only supported target.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
					},
				},
			},

			&metadata.Action{
				Name:        "patch",
				Description: `Change min severity of existing rule`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/api/catalog/accounts/%s/notification_rules/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/notification_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "min_severity",
						Description: `The lowest level of notifications for the target to receive.
                    Setting this to "error" will result in only receiving error notifications,
                    whereas setting it to "info" will result in receiving both info and error notifications,
                    and setting it to "none" will result in not receiving any notifications.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						ValidValues: []string{"error", "info", "none"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "min_severity",
						Description: `The lowest level of notifications for the target to receive.
                    Setting this to "error" will result in only receiving error notifications,
                    whereas setting it to "info" will result in receiving both info and error notifications,
                    and setting it to "none" will result in not receiving any notifications.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						ValidValues: []string{"error", "info", "none"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show one notification rule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/notification_rules/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/notification_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete one notification rule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/catalog/accounts/%s/notification_rules/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/notification_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete one or more notification rules by id or source and target.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/catalog/accounts/%s/notification_rules",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/notification_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `The exact source of the rule to be deleted`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "target",
						Description: `The notification target (user) that the rule applies to.
                    Note that at this time, "me" is the only supported target.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `The exact source of the rule to be deleted`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "target",
						Description: `The notification target (user) that the rule applies to.
                    Note that at this time, "me" is the only supported target.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  false,
					},
				},
			},
		},
	},
	"UserPreference": &metadata.Resource{
		Name: "UserPreference",
		Description: `The UserPreference resource stores preferences on a per user basis, such as default notification preference.
The Self-Service portal uses these preferences in the portal.`,
		Identifier: "application/vnd.rightscale.self_service.user_preference",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "created_by",
				FieldName: "CreatedBy",
				FieldType: "*User",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "timestamps",
				FieldName: "Timestamps",
				FieldType: "*TimestampsStruct",
			},

			&metadata.Attribute{
				Name:      "user_id",
				FieldName: "UserId",
				FieldType: "int",
			},

			&metadata.Attribute{
				Name:      "user_preference_info",
				FieldName: "UserPreferenceInfo",
				FieldType: "*UserPreferenceInfo",
			},

			&metadata.Attribute{
				Name:      "value",
				FieldName: "Value",
				FieldType: "string",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `List the UserPreference for users in this account.
Only administrators and infrastructure users may request the preferences of other users.
Users who are not members of the admin role need to specify a filter with their ID in order to retrieve their preferences.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/user_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/user_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by user, so that only UserPreference belonging to that user are returned. Use "me" as a shortcut for the current user ID.`,
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
						Description: `Filter by user, so that only UserPreference belonging to that user are returned. Use "me" as a shortcut for the current user ID.`,
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
				Description: `Get details for a particular UserPreference`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/user_preferences/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/user_preferences/([^/]+)`),
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
				Name: "create",
				Description: `Create a new UserPreference.
Multiple resources can be created at once with a multipart request.
Values are validated with the corresponding UserPreferenceInfo.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/catalog/accounts/%s/user_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/user_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user_id",
						Description: `Administrators can create preferences for other users by providing this value`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "user_preference_info_id",
						Description: `The ID for the UserPreferenceInfo defining this UserPreference`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "value",
						Description: `The value to set for this UserPreference`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user_id",
						Description: `Administrators can create preferences for other users by providing this value`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "user_preference_info_id",
						Description: `The ID for the UserPreferenceInfo defining this UserPreference`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "value",
						Description: `The value to set for this UserPreference`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update the value of a UserPreference.
Multiple values may be updated using a multipart request.
Values are validated with the corresponding UserPreferenceInfo.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/api/catalog/accounts/%s/user_preferences/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/user_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `In a multipart request, the ID of the UserPreference to update`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "value",
						Description: `The value to set for this UserPreference`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `In a multipart request, the ID of the UserPreference to update`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "value",
						Description: `The value to set for this UserPreference`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete a UserPreference`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/catalog/accounts/%s/user_preferences/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/user_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"UserPreferenceInfo": &metadata.Resource{
		Name: "UserPreferenceInfo",
		Description: `The UserPreferenceInfo resource defines the available user preferences supported by the system.
It is also used to validate values saved in UserPreference.`,
		Identifier: "application/vnd.rightscale.self_service.user_preference_info",
		Attributes: []*metadata.Attribute{
			&metadata.Attribute{
				Name:      "category",
				FieldName: "Category",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "default_value",
				FieldName: "DefaultValue",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "display_name",
				FieldName: "DisplayName",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "help_text",
				FieldName: "HelpText",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "href",
				FieldName: "Href",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "id",
				FieldName: "Id",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "kind",
				FieldName: "Kind",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "name",
				FieldName: "Name",
				FieldType: "string",
			},

			&metadata.Attribute{
				Name:      "value_constraint",
				FieldName: "ValueConstraint",
				FieldType: "[]string",
			},

			&metadata.Attribute{
				Name:      "value_range",
				FieldName: "ValueRange",
				FieldType: "*ValueRangeStruct",
			},

			&metadata.Attribute{
				Name:      "value_type",
				FieldName: "ValueType",
				FieldType: "string",
			},
		},
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the UserPreferenceInfo.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/user_preference_infos",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/user_preference_infos`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by category and/or name`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by category and/or name`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Get details for a particular UserPreferenceInfo`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/catalog/accounts/%s/user_preference_infos/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/api/catalog/accounts/([^/]+)/user_preference_infos/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
}
