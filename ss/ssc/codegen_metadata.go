//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated
// Feb 27, 2015 at 5:28pm (PST)
// Command:
// $ praxisgen -metadata=ssc/restful_doc -output=ssc -pkg=ssc -target=1.0 -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssc

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"AccountPreference": &metadata.Resource{
		Name:        "AccountPreference",
		Description: `The AccountPreference resource stores preferences that apply account-wide, such as UI customization settings and other settings.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the AccountPreferences for this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/account_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/account_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by group, so that only AccountPreferences belonging to that group are returned`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "group",
						Description: `The group that this setting belongs to, simply for organizational purposes`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by group, so that only AccountPreferences belonging to that group are returned`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "group",
						Description: `The group that this setting belongs to, simply for organizational purposes`,
						Type:        "*string",
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
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/account_preferences/%s",
						Variables:  []string{"account_id", "name"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/account_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the AccountPreference to retrieve`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the AccountPreference to retrieve`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create a new AccountPreference or update an existing AccountPreference with the new value`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/accounts/%s/account_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/account_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
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
						HttpMethod: "DELETE",
						Pattern:    "/accounts/%s/account_preferences/%s",
						Variables:  []string{"account_id", "name"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/account_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the AccountPreference to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `The name of the AccountPreference to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"Application": &metadata.Resource{
		Name:        "Application",
		Description: `An Application is an element in the Catalog that can be launched by users`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the Applications available in the specified Catalog.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `An optional list of Application IDs to retrieve. If not specified, all are returned.`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
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
						HttpMethod: "GET",
						Pattern:    "/catalogs/%s/applications/%s",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID`,
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
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID`,
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
				Description: `Create a new Application in the Catalog.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update the content of an existing Application.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/catalogs/%s/applications/%s",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_update",
				Description: `Update the content of multiple Applications.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete an Application from the Catalog`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/catalogs/%s/applications/%s",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete multiple Applications from the Catalog`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/catalogs/%s/applications",
						Variables:  []string{"catalog_id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids[]",
						Description: `The Application IDs to delete`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
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
						HttpMethod: "GET",
						Pattern:    "/catalogs/%s/applications/%s/download",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications/([^/]+)/download`),
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
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID to delete`,
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
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
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
						HttpMethod: "POST",
						Pattern:    "/catalogs/%s/applications/%s/actions/launch",
						Variables:  []string{"catalog_id", "id"},
						Regexp:     regexp.MustCompile(`/catalogs/([^/]+)/applications/([^/]+)/actions/launch`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "catalog_id",
						Description: `The catalog ID (currently the account ID)`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The Application ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"NotificationRule": &metadata.Resource{
		Name:        "NotificationRule",
		Description: `A notification rule describes which notification should be created         when events occur in the system...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List all notification rules, potentially filtering by a collection of resources.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/notification_rules",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/notification_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `List all notification rules where the target is the current user.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
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
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `List all notification rules where the target is the current user.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "targets",
						Description: `Comma separated list of target ids. Note, currently only "me" is allowed.`,
						Type:        "*string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Create one notification rule for a specific target and source.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/accounts/%s/notification_rules",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/notification_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "patch",
				Description: `Change min severity of existing rule`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PATCH",
						Pattern:    "/accounts/%s/notification_rules/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/notification_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show one notification rule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/notification_rules/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/notification_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Delete one notification rule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/accounts/%s/notification_rules/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/notification_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Notification rule id`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete one or more notification rules by id or source and target.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/accounts/%s/notification_rules",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/notification_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `Account ID of the target and resource.`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"UserPreference": &metadata.Resource{
		Name:        "UserPreference",
		Description: `The UserPreference resource stores preferences on a per user basis, such as default notification preference.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the UserPreference for users in this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/user_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/user_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
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
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
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
				Description: `Get details for a particular UserPreference`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/user_preferences/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/user_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreference to retrieve`,
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
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreference to retrieve`,
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
				Description: `Create a new UserPreference.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/accounts/%s/user_preferences",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/user_preferences`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update the value of a UserPreference.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PATCH",
						Pattern:    "/accounts/%s/user_preferences/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/user_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreference to update, or "*" for a multipart request`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreference to update, or "*" for a multipart request`,
						Type:        "string",
						Location:    metadata.PathParam,
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
						HttpMethod: "DELETE",
						Pattern:    "/accounts/%s/user_preferences/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/user_preferences/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreference to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreference to delete`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"UserPreferenceInfo": &metadata.Resource{
		Name:        "UserPreferenceInfo",
		Description: `The UserPreferenceInfo resource defines the available user preferences supported by the system.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List the UserPreferenceInfo.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/user_preference_infos",
						Variables:  []string{"account_id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/user_preference_infos`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: `Filter by category and/or name`,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
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
						HttpMethod: "GET",
						Pattern:    "/accounts/%s/user_preference_infos/%s",
						Variables:  []string{"account_id", "id"},
						Regexp:     regexp.MustCompile(`/accounts/([^/]+)/user_preference_infos/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreferenceInfo to retrieve`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The account ID`,
						Type:        "string",
						Location:    metadata.PathParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `The ID of the UserPreferenceInfo to retrieve`,
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
