//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ api15gen -metadata=cm15 -output=cm15
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package cm15

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"Account": &metadata.Resource{
		Name:        "Account",
		Description: ``,
		Identifier:  "application/vnd.rightscale.account",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/accounts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/accounts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
		Links: map[string]string{
			"cluster": "Href of the cluster the account belongs to",
			"owner":   "Href of the owner",
			"self":    "Href of itself",
		},
	},
	"AccountGroup": &metadata.Resource{
		Name:        "AccountGroup",
		Description: ` An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.`,
		Identifier:  "application/vnd.rightscale.account_group",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Lists the AccountGroups owned by this Account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/account_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/account_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single AccountGroup.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/account_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/account_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"account": "Associated account",
			"self":    "Href of itself",
		},
	},
	"Alert": &metadata.Resource{
		Name:        "Alert",
		Description: `An Alert represents an AlertSpec bound to a running Instance.`,
		Identifier:  "application/vnd.rightscale.alert",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "destroy",
				Description: `Destroys the Alert.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/instances/%s/alerts/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/servers/%s/alerts/%s",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_arrays/%s/alerts/%s",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/deployments/%s/alerts/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/alerts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "disable",
				Description: `Disables the Alert indefinitely. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/alerts/%s/disable",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/alerts/%s/disable",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/alerts/%s/disable",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/alerts/%s/disable",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/alerts/%s/disable",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)/disable$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "enable",
				Description: `Enables the Alert indefinitely. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/alerts/%s/enable",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/alerts/%s/enable",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/alerts/%s/enable",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/alerts/%s/enable",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/alerts/%s/enable",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)/enable$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all Alerts.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/alerts",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/servers/%s/alerts",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s/alerts",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/alerts",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/alerts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/alerts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"alert_spec_href", "status"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"alert_spec_href", "status"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "quench",
				Description: `Suppresses the Alert from being triggered for a given time period. Idempotent.
Required parameters:
	duration: The time period in seconds to suppress Alert from being triggered.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/alerts/%s/quench",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/alerts/%s/quench",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/alerts/%s/quench",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/alerts/%s/quench",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/alerts/%s/quench",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)/quench$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "duration",
						Description: `The time period in seconds to suppress Alert from being triggered.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "duration",
						Description: `The time period in seconds to suppress Alert from being triggered.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows the attributes of a specified Alert.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/alerts/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/servers/%s/alerts/%s",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s/alerts/%s",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/alerts/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/alerts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"alert_spec": "Associated AlertSpec",
			"instance":   "Associated running Instance",
			"self":       "Href of itself",
		},
	},
	"AlertSpec": &metadata.Resource{
		Name: "AlertSpec",
		Description: `An AlertSpec defines the conditions under which an Alert is triggered and escalated.
Condition sentence: if &lt;file&gt;.&lt;variable&gt; &lt;condition&gt; '&lt;threshold&gt;' for &lt;duration&gt; min then escalate to '&lt;escalation_name&gt;'.`,
		Identifier: "application/vnd.rightscale.alert_spec",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new AlertSpec with the given parameters.
Required parameters:
	alert_spec`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/alert_specs",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/alert_specs",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/alert_specs",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/alert_specs",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/alert_specs",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/alert_specs$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "alert_spec[escalation_name]",
						Description: `Escalate to the named alert escalation when the alert is triggered. Must either escalate or vote.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[subject_href]",
						Description: `The href of the resource that this AlertSpec should be associated with. The subject can be a ServerTemplate, Server, ServerArray, or Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[description]",
						Description: `The description of the AlertSpec.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_type]",
						Description: `Vote to grow or shrink a ServerArray when the alert is triggered. Must either escalate or vote.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"grow", "shrink"},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[condition]",
						Description: `The condition (operator) in the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{">", ">=", "<", "<=", "==", "!="},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[threshold]",
						Description: `The threshold of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[duration]",
						Description: `The duration in minutes of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[variable]",
						Description: `The RRD variable of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_tag]",
						Description: `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[name]",
						Description: `The name of the AlertSpec.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[file]",
						Description: `The RRD path/file_name of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "alert_spec",
						Description: ``,
						Type:        "*AlertSpecParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given AlertSpec.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/instances/%s/alert_specs/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/servers/%s/alert_specs/%s",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_arrays/%s/alert_specs/%s",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_templates/%s/alert_specs/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/alert_specs/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alert_specs/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `No description provided for index.
Optional parameters:
	filter
	view
	with_inherited: Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/alert_specs",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/servers/%s/alert_specs",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s/alert_specs",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/alert_specs",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/alert_specs",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/alert_specs$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "with_inherited",
						Description: `Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "escalation_name", "name", "subject_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "escalation_name", "name", "subject_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
					&metadata.ActionParam{
						Name:        "with_inherited",
						Description: `Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `No description provided for show.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/alert_specs/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/servers/%s/alert_specs/%s",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s/alert_specs/%s",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/alert_specs/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/alert_specs/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alert_specs/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates an AlertSpec with the given parameters.
Required parameters:
	alert_spec`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/instances/%s/alert_specs/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/servers/%s/alert_specs/%s",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/server_arrays/%s/alert_specs/%s",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/server_templates/%s/alert_specs/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/alert_specs/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alert_specs/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "alert_spec[escalation_name]",
						Description: `Escalate to the named alert escalation when the alert is triggered.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[description]",
						Description: `The description of the AlertSpec.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[condition]",
						Description: `The condition (operator) in the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{">", ">=", "<", "<=", "==", "!="},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[threshold]",
						Description: `The threshold of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_type]",
						Description: `Vote to grow or shrink a ServerArray when the alert is triggered.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"grow", "shrink"},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[duration]",
						Description: `The duration in minutes of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[variable]",
						Description: `The RRD variable of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_tag]",
						Description: `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[name]",
						Description: `The name of the AlertSpec.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[file]",
						Description: `The RRD path/file_name of the condition sentence.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "alert_spec",
						Description: ``,
						Type:        "*AlertSpecParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"self":    "Href of itself",
			"subject": "Associated subject. The subject can be a ServerTemplate, Server, Server Array, or a running Instance.",
		},
	},
	"AuditEntry": &metadata.Resource{
		Name:        "AuditEntry",
		Description: `An Audit Entry can be used to track various activities of a resource.`,
		Identifier:  "application/vnd.rightscale.audit_entry",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "append",
				Description: `Updates the summary and appends more details to a given AuditEntry. Each audit entry detail is stored
as one chunk, the offset determines the location of that chunk within the overall audit entry details section.
For example, if you create an AuditEntry and append "DEF" at offset 10, and later append
"ABC" at offset 9, the overall audit entry details will be "ABCDEF". Use the \n character to
separate details by new lines.
Optional parameters:
	detail: The details to be appended to the audit entry record.
	notify: The event notification category. Defaults to 'None'.
	offset: The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.
	summary: The updated summary for the audit entry, maximum length is 255 characters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/audit_entries/%s/append",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)/append$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "summary",
						Description: `The updated summary for the audit entry, maximum length is 255 characters.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "detail",
						Description: `The details to be appended to the audit entry record.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "offset",
						Description: `The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "detail",
						Description: `The details to be appended to the audit entry record.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "offset",
						Description: `The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "summary",
						Description: `The updated summary for the audit entry, maximum length is 255 characters.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a new AuditEntry with the given parameters.
Required parameters:
	audit_entry
Optional parameters:
	notify: The event notification category. Defaults to 'None'.
	user_email: The email of the user (who created/triggered the audit entry). Only usable with instance role.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/audit_entries",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/audit_entries$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "audit_entry[auditee_href]",
						Description: `The href of the resource that this audit entry should be associated with (e.g. an instance's href).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_entry[summary]",
						Description: `The summary of the audit entry to be created, maximum length is 255 characters.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_entry[detail]",
						Description: `The initial details of the audit entry to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "user_email",
						Description: `The email of the user (who created/triggered the audit entry). Only usable with instance role.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "audit_entry",
						Description: ``,
						Type:        "*AuditEntryParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user_email",
						Description: `The email of the user (who created/triggered the audit entry). Only usable with instance role.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "detail",
				Description: `shows the details of a given AuditEntry.
Note that the media type of the response is simply text.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/audit_entries/%s/detail",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)/detail$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists AuditEntries of the account. Due to the potentially large number of audit entries, a start and end date must
be provided during an index call to limit the search. The format of the dates must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
2011/07/11 00:00:00 +0000.
A maximum of 1000 records will be returned by each index call.
Using the available filters, one can select or group which audit entries to retrieve.
Required parameters:
	end_date: The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).
	limit: Limit the audit entries to this number. The limit should >= 1 and <= 1000
	start_date: The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/audit_entries",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/audit_entries$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "start_date",
						Description: `The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "end_date",
						Description: `The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"auditee_href", "user_email"},
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `Limit the audit entries to this number. The limit should >= 1 and <= 1000`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "end_date",
						Description: `The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"auditee_href", "user_email"},
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `Limit the audit entries to this number. The limit should >= 1 and <= 1000`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "start_date",
						Description: `The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Lists the attributes of a given audit entry.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/audit_entries/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates the summary of a given AuditEntry.
Required parameters:
	audit_entry
Optional parameters:
	notify: The event notification category. Defaults to 'None'.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/audit_entries/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "audit_entry[summary]",
						Description: `The updated summary for the audit entry, maximum length is 255 characters.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_entry[offset]",
						Description: `The offset where the next details will be appended. Used in ordering of summary updates.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "audit_entry",
						Description: ``,
						Type:        "*AuditEntryParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"auditee": "Href of the resource that this Audit Entry relates to.",
			"detail":  "Href where the Audit Entry detail is available from.",
			"self":    "Href of itself",
		},
	},
	"Backup": &metadata.Resource{
		Name:        "Backup",
		Description: ``,
		Identifier:  "application/vnd.rightscale.backup",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "cleanup",
				Description: `Deletes old backups that meet the given criteria. For example, if a user calls cleanup with keep monthlies set to 12,
then the latest backup for each month, for 12 months, will be kept.
All backups belong to a particular 'lineage'. Backups are not constrained to a specific cloud or a specific deployment.
A lineage is account-specific. Hence, backups having the same lineage but belonging to different clouds are still considered
for cleanup.
If backups specific to a single cloud should be cleaned up, see the cloud_href parameter.
Definitions:
A snapshot is completed if its status is "available"
A snapshot is committed if it has a tag "rs_backup:committed=true"
A snapshot belongs to a backup "X" if it has a tag "rs_backup:backup_id=X"
A snapshot is part of a backup with size "Y" if it has a tag "rs_backup:count=Y"
A snapshot's position in a backup is "Z" if it has a tag "rs_backup:position=Z"
Backups are of 3 types:
Perfect backup: A backup which is completed (all the snapshots are completed) and committed (all the snapshots are committed) and the number of snapshots it found is equal to the number
in the "rs_backup:count=" tag on each of the Snapshots.
Imperfect backup: A backup which is not committed or if the number of snapshots it found is not equal to the number in the "rs_backup:count=" tag on each of the snapshots.
Partial Perfect backup: A snapshot which is neither perfect nor imperfect
An imperfect backup is picked up for cleanup only if there exists a perfect backup with a newer created_at timestamp.
No constraints will be applied on such imperfect backups and all of them will be destroyed.
For all the perfect backups, the constraints of keep_last and dailies etc. will be applied.
The algorithm for choosing the perfect backups to keep is simple. It is the union of those set of backups if each of those conditions are applied
independently. i.e backups_to_keep = backups_to_keep(keep_last) U backups_to_keep(dailies) U backups_to_keep(weeklies) U backups_to_keep(monthlies) U backups_to_keep(yearlies)
Hence, it is important to "commit" a backup to make it eligible for cleanup.
Required parameters:
	keep_last: The number of backups that should be kept.
	lineage: The lineage of the backups that are to be cleaned-up.
Optional parameters:
	cloud_href: Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.
	dailies: The number of daily backups(the latest one in each day) that should be kept.
	monthlies: The number of monthly backups(the latest one in each month) that should be kept.
	weeklies: The number of weekly backups(the latest one in each week) that should be kept.
	yearlies: The number of yearly backups(the latest one in each year) that should be kept.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/backups/cleanup",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/backups/cleanup$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cloud_href",
						Description: `Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "keep_last",
						Description: `The number of backups that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "monthlies",
						Description: `The number of monthly backups(the latest one in each month) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "weeklies",
						Description: `The number of weekly backups(the latest one in each week) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "yearlies",
						Description: `The number of yearly backups(the latest one in each year) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "dailies",
						Description: `The number of daily backups(the latest one in each day) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "lineage",
						Description: `The lineage of the backups that are to be cleaned-up.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cloud_href",
						Description: `Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "dailies",
						Description: `The number of daily backups(the latest one in each day) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "keep_last",
						Description: `The number of backups that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "lineage",
						Description: `The lineage of the backups that are to be cleaned-up.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "monthlies",
						Description: `The number of monthly backups(the latest one in each month) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "weeklies",
						Description: `The number of weekly backups(the latest one in each week) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "yearlies",
						Description: `The number of yearly backups(the latest one in each year) that should be kept.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Takes in an array of volume_attachment_hrefs and takes a snapshot of each.
The volume_attachment_hrefs must belong to the same instance.
Required parameters:
	backup`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/backups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/backups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup[volume_attachment_hrefs][]",
						Description: `List of volume attachment hrefs that are to be backed-up.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[description]",
						Description: `The description to be set on each of the volume snapshots`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[from_master]",
						Description: `Setting this to 'true' will create a tag 'rs_backup:from_master=true' on the snapshots so that one can filter them later.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name: "backup[lineage]",
						Description: `A unique value to create backups belonging to a particular system.
                                       This will be used to set the tag  e.g. 'rs_backup:lineage=prod_mysqldb'. `,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  true,
					},
					&metadata.ActionParam{
						Name:        "backup[name]",
						Description: `The name to be set on each of the volume snapshots.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup",
						Description: ``,
						Type:        "*BackupParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/backups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all of the backups with the given lineage tag. Filters can be used to search for a particular backup. If the
'latest_before' filter is set, only one backup is returned (the latest backup before the given timestamp).
To get the latest completed backup, the 'completed' filter should be set to 'true' and the 'latest_before' filter
should be set to the current timestamp. The format of the timestamp must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
2011/07/11 00:00:00 +0000.
To get the latest completed backup just before, say 25 June 2009, then the 'completed' filter
should be set to 'true' and the 'latest_before' filter should be set to 2009/06/25 00:00:00 +0000.
Required parameters:
	lineage: Backups belonging to this lineage.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/backups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/backups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "lineage",
						Description: `Backups belonging to this lineage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "committed", "completed", "from_master", "latest_before"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "committed", "completed", "from_master", "latest_before"},
					},
					&metadata.ActionParam{
						Name:        "lineage",
						Description: `Backups belonging to this lineage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "restore",
				Description: `Restores the given Backup.
This call will:
create the required number of Volumes from the volume_snapshots_hrefs in the given Backup,
attach them to the given Instance at the device specified in the Snapshot. If the devices are already being used
   on the Instance, the Task will denote that the restore has failed.
Required parameters:
	instance_href: The instance href that the backup will be restored to.
Optional parameters:
	backup`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/backups/%s/restore",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)/restore$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup[volume_type_href]",
						Description: `The href of the volume type. Each volume is created with this volume type instead of the default volume type for the cloud. A Name, Resource UID and optional Size is associated with a volume type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[description]",
						Description: `Each volume is created with this description instead of the volume snapshot's description`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance_href",
						Description: `The instance href that the backup will be restored to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[iops]",
						Description: `The number of IOPS (I/O Operations Per Second) each volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[name]",
						Description: `Each volume is created with this name instead of the volume snapshot's name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[size]",
						Description: `Each volume is created with this size in gigabytes (GB) instead of the volume snapshot's size (must be equal or larger). Some volume types have predefined sizes and do not allow selecting a custom size on volume creation.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup",
						Description: ``,
						Type:        "*BackupParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance_href",
						Description: `The instance href that the backup will be restored to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Lists the attributes of a given backup`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/backups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.
Required parameters:
	backup`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/backups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup[committed]",
						Description: `Setting this to 'true' will update the 'rs_backup:committed=false' tag to 'rs_backup:committed=true' on all the snapshots.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup",
						Description: ``,
						Type:        "*BackupParam3",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"self": "Href of itself",
		},
	},
	"ChildAccount": &metadata.Resource{
		Name:        "ChildAccount",
		Description: ``,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create an enterprise ChildAccount for this Account. The User will by default get an 'admin' role
on the ChildAccount to enable him/her to add, delete Users and Permissions.
For more information on the valid values for 'cluster_href', refer to the following:
http://support.rightscale.com/12-Guides/RightScale_API_1.5/Examples/ChildAccounts/Create
Required parameters:
	child_account`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/child_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/child_accounts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "child_account[cluster_href]",
						Description: `The href of the cluster in which to create the account. If not specified, will default to the cluster of the parent account.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "child_account[name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "child_account",
						Description: ``,
						Type:        "*ChildAccountParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the enterprise ChildAccounts available for this Account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/child_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/child_accounts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update an enterprise ChildAccount for this Account.
Required parameters:
	child_account`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/accounts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/accounts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/child_accounts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/child_accounts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "child_account[name]",
						Description: `The updated name for the account.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "child_account",
						Description: ``,
						Type:        "*ChildAccountParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Cloud": &metadata.Resource{
		Name:        "Cloud",
		Description: `Represents a Cloud (within the context of the account in the session).`,
		Identifier:  "application/vnd.rightscale.cloud",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Lists the clouds available to this account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/clouds$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_type", "description", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_type", "description", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single cloud.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},
		},
		Links: map[string]string{
			"datacenters":                  "Associated datacenters",
			"images":                       "Associated images",
			"instance_types":               "Associated instance types",
			"instances":                    "Associated instances",
			"ip_address_bindings":          "Associated IP Address Bindings",
			"ip_addresses":                 "Associated IP Addresses",
			"recurring_volume_attachments": "Associated recurring volume attachments",
			"security_groups":              "Associated security groups",
			"self":                         "Href of itself",
			"ssh_keys":                     "Associated ssh keys",
			"subnets":                      "Associated subnets",
			"volume_attachments":           "Associated volume attachments",
			"volume_snapshots":             "Associated volume snapshots",
			"volume_types":                 "Associated volume types",
			"volumes":                      "Associated volumes",
		},
	},
	"CloudAccount": &metadata.Resource{
		Name:        "CloudAccount",
		Description: `Represents a Cloud Account (An association between the account and a cloud).`,
		Identifier:  "application/vnd.rightscale.cloud_account",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a CloudAccount by passing in the respective credentials for each cloud.
For more information on the specific parameters for each cloud, refer to the following:
http://docs.rightscale.com/api/api_1.5_examples/cloudaccounts.html
Required parameters:
	cloud_account`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cloud_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cloud_account[cloud_href]",
						Description: `The href of the cloud if it is known. For valid values see support portal link above.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cloud_account[creds]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cloud_account[token]",
						Description: `The cloud token to identify a private cloud`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cloud_account",
						Description: ``,
						Type:        "*CloudAccountParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete a CloudAccount.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/cloud_accounts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the CloudAccounts (non-aws) available to this Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cloud_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "show",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cloud_accounts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
		Links: map[string]string{
			"account": "Associated account",
			"cloud":   "Associated cloud",
			"self":    "Href of itself",
		},
	},
	"Cookbook": &metadata.Resource{
		Name:        "Cookbook",
		Description: `Represents a given instance of a single cookbook.`,
		Identifier:  "application/vnd.rightscale.cookbook",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "destroy",
				Description: `Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/cookbooks/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "follow",
				Description: `Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.
Required parameters:
	value: Indicates if this action should follow (true) or unfollow (false) a Cookbook.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cookbooks/%s/follow",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/follow$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should follow (true) or unfollow (false) a Cookbook.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should follow (true) or unfollow (false) a Cookbook.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "freeze",
				Description: `Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.
Required parameters:
	value: Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cookbooks/%s/freeze",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/freeze$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the Cookbooks available to this account.
The extended_designer view is only available to accounts with the designer permission.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cookbooks",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbooks$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "namespace", "repository_href", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "namespace", "repository_href", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
			},

			&metadata.Action{
				Name: "obsolete",
				Description: `Marks a Cookbook as obsolete (or un-obsolete).
Required parameters:
	value: Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cookbooks/%s/obsolete",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/obsolete$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single Cookbook.
The extended_designer view is only available to accounts with the designer permission.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cookbooks/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
			},
		},
		Links: map[string]string{
			"cookbook_attachments": "Attachments for this Cookbook",
			"repository":           "Href of associated Repository",
			"self":                 "Href of itself",
		},
	},
	"CookbookAttachment": &metadata.Resource{
		Name:        "CookbookAttachment",
		Description: `Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.`,
		Identifier:  "application/vnd.rightscale.cookbook_attachment",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Attach a cookbook to a given resource.
Optional parameters:
	cookbook_attachment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cookbooks/%s/cookbook_attachments",
						Variables:  []string{"cookbook_id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/cookbook_attachments",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cookbook_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachment[server_template_href]",
						Description: `The href of the server template to attach the cookbook to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cookbook_attachment[cookbook_href]",
						Description: `The href of the cookbook to attach.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachment",
						Description: ``,
						Type:        "*CookbookAttachmentParam",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Detach a cookbook from a given resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/cookbooks/%s/cookbook_attachments/%s",
						Variables:  []string{"cookbook_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_templates/%s/cookbook_attachments/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/cookbook_attachments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists Cookbook Attachments.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cookbooks/%s/cookbook_attachments",
						Variables:  []string{"cookbook_id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/cookbook_attachments",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cookbook_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "multi_attach",
				Description: `Attach multiple cookbooks to a given resource.
Required parameters:
	cookbook_attachments`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/cookbook_attachments/multi_attach",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/multi_attach$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cookbook_attachments/multi_attach",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/multi_attach$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachments[server_template_href]",
						Description: `The href of the server template to attach the cookbooks to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cookbook_attachments[cookbook_hrefs][]",
						Description: `The hrefs of the cookbooks to be attached`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachments",
						Description: ``,
						Type:        "*CookbookAttachments",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "multi_detach",
				Description: `Detach multiple cookbooks from a given resource.
Required parameters:
	cookbook_attachments`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/cookbook_attachments/multi_detach",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/multi_detach$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/cookbook_attachments/multi_detach",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/multi_detach$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachments[cookbook_attachment_hrefs][]",
						Description: `The hrefs of the cookbook attachments to be detached`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachments",
						Description: ``,
						Type:        "*CookbookAttachments2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single cookbook attachment to a ServerTemplate.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cookbooks/%s/cookbook_attachments/%s",
						Variables:  []string{"cookbook_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/cookbook_attachments/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/cookbook_attachments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cookbook":        "Href of the attached Cookbook",
			"self":            "Href of itself",
			"server_template": "Href of ServerTemplate",
		},
	},
	"Credential": &metadata.Resource{
		Name: "Credential",
		Description: `A Credential provides an abstraction for sensitive textual information,
such as passphrases or cloud credentials, whose value should be encrypted
when stored in the database and not generally shown in the UI or through the
API. Credentials may then be used as inputs of type "Cred" in RightScripts
or Chef recipes. NOTE: Credential values may be updated through the API, but
values cannot be retrieved via the API.`,
		Identifier: "application/vnd.rightscale.credential",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new Credential with the given parameters.
Required parameters:
	credential`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/credentials",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/credentials$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credential[description]",
						Description: `The description of the Credential to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "credential[value]",
						Description: `The value of the Credential to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "credential[name]",
						Description: `The name of the Credential to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credential",
						Description: ``,
						Type:        "*CredentialParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a Credential.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/credentials/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/credentials/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the Credentials available to this account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/credentials",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/credentials$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single Credential. Credential values may be retrieved using the "sensitive" view by users with "admin" role only.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/credentials/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/credentials/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a Credential.
Required parameters:
	credential`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/credentials/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/credentials/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credential[description]",
						Description: `The updated description of the Credential.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "credential[value]",
						Description: `The updated value of the Credential.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "credential[name]",
						Description: `The updated name of the Credential.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credential",
						Description: ``,
						Type:        "*CredentialParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"self": "Href of itself",
		},
	},
	"Datacenter": &metadata.Resource{
		Name: "Datacenter",
		Description: `Datacenters represent isolated facilities within a cloud. The level and type of isolation is cloud dependent.
While Datacenters in large public clouds might correspond to different physical buildings, with different power,
internet links...etc., Datacenters within the context of a private cloud might simply correspond to having different network providers.
Spreading servers across distinct Datacenters helps minimize outages.`,
		Identifier: "application/vnd.rightscale.datacenter",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Lists all Datacenters for a particular cloud.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/datacenters",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/datacenters$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single Datacenter.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/datacenters/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/datacenters/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud": "Associated cloud",
			"self":  "Href of itself",
		},
	},
	"Deployment": &metadata.Resource{
		Name:        "Deployment",
		Description: `Deployments represent logical groupings of related assets such as servers, server arrays, default configuration settings...etc.`,
		Identifier:  "application/vnd.rightscale.deployment",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "clone",
				Description: `Clones a given deployment.
Optional parameters:
	deployment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/clone$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the cloned deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&metadata.ActionParam{
						Name:        "deployment[description]",
						Description: `The description for the cloned deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "deployment[name]",
						Description: `The name for the cloned deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment",
						Description: ``,
						Type:        "*DeploymentParam",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a new deployment with the given parameters.
Required parameters:
	deployment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/deployments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment[resource_group_href]",
						Description: `The href of the Windows Azure Resource Group attached to the deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&metadata.ActionParam{
						Name:        "deployment[description]",
						Description: `The description of the deployment to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "deployment[name]",
						Description: `The name of the deployment to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment",
						Description: ``,
						Type:        "*DeploymentParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given deployment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/deployments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists deployments of the account.
Using the available filters, one can select or group which deployments to retrieve.
The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
details please see Inputs#index.)
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/deployments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name", "resource_group_href", "server_tag_scope"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name", "resource_group_href", "server_tag_scope"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&metadata.Action{
				Name: "lock",
				Description: `Locks a given deployment. Idempotent.
Locking prevents servers from being deleted or moved from the deployment.
Other actions such as adding servers or renaming the deployment are still allowed.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/lock",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/lock$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "servers",
				Description: `Lists the servers belonging to this deployment. This call is equivalent to servers#index call, where the servers returned will
automatically be filtered by this deployment. See servers#index for details on other options and parameters.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/servers",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Lists the attributes of a given deployment.
The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
details please see Inputs#index.)
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&metadata.Action{
				Name:        "unlock",
				Description: `Unlocks a given deployment. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/unlock",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/unlock$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a given deployment.
Required parameters:
	deployment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/deployments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment[resource_group_href]",
						Description: `The href of the Windows Azure Resource Group attached to the deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&metadata.ActionParam{
						Name:        "deployment[description]",
						Description: `The updated description for the deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "deployment[name]",
						Description: `The updated name for the deployment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment",
						Description: ``,
						Type:        "*DeploymentParam3",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"alerts":         "Associated alerts",
			"inputs":         "List of configuration inputs",
			"lock_user":      "Href of the user who has locked this deployment",
			"resource_group": "Href of the ResourceGroup that this Deployment belongs to",
			"self":           "Href of itself",
			"server_arrays":  "Associated server arrays",
			"servers":        "Associated servers",
		},
	},
	"HealthCheck": &metadata.Resource{
		Name:        "HealthCheck",
		Description: ``,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Check health of RightApi controllers`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/health-check/",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/health-check/$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"IdentityProvider": &metadata.Resource{
		Name: "IdentityProvider",
		Description: `An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account,
and is trusted by the RightScale dashboard to authenticate your enterprise's end users.
To register an Identity Provider, contact your account manager.`,
		Identifier: "application/vnd.rightscale.identity_provider",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Lists the identity providers associated with this enterprise account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/identity_providers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/identity_providers$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show the specified identity provider, if associated with this enterprise account.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/identity_providers/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/identity_providers/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"self": "Href of itself",
		},
	},
	"Image": &metadata.Resource{
		Name: "Image",
		Description: `Images represent base VM image existing in a cloud. An image will define the initial Operating System and root disk contents
for a new Instance to have, and therefore it represents the basic starting point for creating a new one.`,
		Identifier: "application/vnd.rightscale.image",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Lists all Images for the given Cloud.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/images",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/images$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cpu_architecture", "description", "image_type", "name", "os_platform", "resource_uid", "visibility"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cpu_architecture", "description", "image_type", "name", "os_platform", "resource_uid", "visibility"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows information about a single Image.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/images/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/images/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud": "Associated cloud",
			"self":  "Href of itself",
		},
	},
	"Input": &metadata.Resource{
		Name: "Input",
		Description: `Inputs help extract dynamic information, usually specified at runtime, from repeatable configuration operations that can be codified.
Inputs are variables defined in and used by RightScripts/Recipes. The two main attributes of an input are 'name' and 'value'. The 'name'
identifies the input and the 'value', although a string encodes what type it is. It could be a text encoded as 'text:myvalue' or a credential
encoded as 'cred:MY_CRED' or a key etc. Please see support.rightscale.com for more info on input hierarchies and their different types.`,
		Identifier: "application/vnd.rightscale.input",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Retrieves the full list of existing inputs of the specified resource.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/inputs",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/inputs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/inputs",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/inputs$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/inputs",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/inputs$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs_2_0"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs_2_0"},
					},
				},
			},

			&metadata.Action{
				Name: "multi_update",
				Description: `Performs a bulk update of inputs on the specified resource.
If an input exists with the same name, its value will be updated. If an input does not
exist with a specified name, it will be ignored.
Input values are represented as strings.
There are two notations for inputs:
1.0 notation - The deprecated notation used in API 1.0 and in 1.5
    2.0 notation - The new notation that is partially supported in API 1.5, and will
        be the only notation supported in 2.0
Although the two notations are similar, they have a few important differences, in particular:
  With 2.0 notation, values MUST begin with a prefix identifying their type, followed
  by a colon (example: "text:foo"). With 1.0 notation, unprefixed values are generally
  taken to be text-type.
  With 2.0 notation, a sentinel value "inherit" is used to express that an input
  should use an inherited value. With 1.0 notation the empty string was used to express
  the same thing. (Due to requirement 1, empty string is no longer a valid input.)
  With 2.0 notation, each element of an array is an entire input value; arrays can
  contain cred, env, or even other arrays. With 1.0 notation, array elements are
  implicitly text values and there is no way to specify anything else.Note that the UI
  does not support complex-valued arrays; please use this feature with caution!
The following types of inputs are supported:
Type
Format
1.0 Example(s)
2.0 Example(s)
Text string
&lt;value&gt; (1.0 only)text:&lt;value&gt;
footext:footext:multi word value
text:footext:multi word value
Blank string(input is present but its value is empty-string)
text:blank (2.0 only)
text:
blank
Ignore (input is not present, input will inherit)
ignore$ignore (1.0 only)ignore:$ignore (1.0 only)
ignore$ignoreignore:$ignore
ignore
Dynamically-substituted environment value
env:&lt;value&gt;env:&lt;component&gt;:&lt;value&gt;
env:MY_ENV_VARenv:my_server:MY_ENV_VAR
env:MY_ENV_VARenv:my_server:MY_ENV_VAR
Credential value
cred:&lt;value&gt;
cred:abcd1234wxyz
cred:abcd1234wxyz
Private SSH key
key:&lt;value&gt;key:&lt;value&gt;:&lt;cloud_id&gt;
key:1234abcd5678key:1234abcd5678:1
key:1234abcd5678key:1234abcd5678:1
Array of values
array:&lt;value&gt;,... (1.0 only)array:["&lt;type&gt;:&lt;value&gt;",...] (2.0 only)
array:x,y(NOTE: 1.0 only supports text inputs for arrays)
array:["text:v1","text:v2"]array:["text:x","env:server_x:MY_VAR"]
Note that in the case of array inputs, the portion after the colon must be
valid JSON. In particular, when enclosing the input within double-quotes
(e.g. for use in cURL or Ruby), the double-quotes must be escaped.
Single-quotes may not be used within the array input, since they are not
valid for JSON strings.
The legacy format for providing inputs is as an array of name-value pairs
(ex: -d inputs[][name]="MY_INPUT" -d inputs[][value]="text:foobar"), however
the new format is supported for inputs provided as a hash
(ex: -d inputs[MY_INPUT]="text:foobar").
If the old format is used, the input is parsed using 1.0 semantics.
If the new format is used, the input is parsed using the new 2.0 semantics.
Required parameters:
	inputs`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/instances/%s/inputs/multi_update",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/inputs/multi_update$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/deployments/%s/inputs/multi_update",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/inputs/multi_update$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/server_templates/%s/inputs/multi_update",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/inputs/multi_update$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value to be updated with. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The name of the input to be updated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Instance": &metadata.Resource{
		Name: "Instance",
		Description: `Instances represent an entity that is runnable in the cloud.
An instance of type "next" is a container of information that expresses how to configure a future instance when we decide
to launch or start it.
A "next" instance generally only exists in the RightScale realm, and usually doesn't have any corresponding representation
existing in the cloud. However, if an instance is not of type "next", it will generally represent an existing running
(or provisioned) virtual machine existing in the cloud.`,
		Identifier: "application/vnd.rightscale.instance",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates and launches a raw instance using the provided parameters.
Required parameters:
	instance
Optional parameters:
	api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][create_default_port_forwarding_rules]",
						Description: `Automatically create default port forwarding rules (enabled by default). Supported by Azure cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][local_ssd_interface]",
						Description: `The type of SSD(s) to be created. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][create_boot_volume]",
						Description: `If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][delete_boot_volume]",
						Description: `If enabled, the associated volume will be deleted when the instance is terminated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][placement_tenancy]",
						Description: `The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "dedicated"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][availability_set]",
						Description: `Availability set for raw instance. Supported by Azure v2 cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][service_account]",
						Description: `Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][local_ssd_count]",
						Description: `Additional local SSDs. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][keep_alive_url]",
						Description: `The ulr of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][admin_username]",
						Description: `The user that will be granted administrative privileges. Supported by AzureRM cloud only. For more information, <a href="http://docs.rightscale.com/clouds/azure_resource_manager/reference/limitations.html">review the documentation</a>.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][max_spot_price]",
						Description: `Specify the max spot price you will pay for. Required when 'pricing_type' is 'spot'. Only applies to clouds which support spot-pricing and when 'spot' is chosen as the 'pricing_type'. Should be a Float value >= 0.001, eg: 0.095, 0.123, 1.23, etc...`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][keep_alive_id]",
						Description: `The id of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][ebs_optimized]",
						Description: `Whether the instance is able to connect to IOPS-enabled volumes.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][pricing_type]",
						Description: `Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance.WARNING:  By using spot pricing, you acknowledge that your instance/server/array may not be able to be launched (and arrays may be unable to grow) as newly launched instances might be stuck in bidding, and/or existing instances may be terminated at any time, due to the cloud's spot pricing changes and availability.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"fixed", "spot"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][preemptible]",
						Description: `Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][memory_mb]",
						Description: `The size of instance memory. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][num_cores]",
						Description: `The number of instance cores. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][metadata]",
						Description: `Extra data used for configuration, in query string format.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][disk_gb]",
						Description: `The size of root disk. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[security_group_hrefs][]",
						Description: `The hrefs of the security groups.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[placement_group_href]",
						Description: `The placement group to launch the instance in. Not supported by all clouds & instance types.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ramdisk_image_href]",
						Description: `The href of the ramdisk image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[instance_type_href]",
						Description: `The href of the instance type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[kernel_image_href]",
						Description: `The href of the kernel image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[deployment_href]",
						Description: `The href of the deployment to which the Instance will be added.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ssh_key_href]",
						Description: `The href of the SSH key to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[name]",
						Description: `The name of the instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
					&metadata.ActionParam{
						Name:        "instance",
						Description: ``,
						Type:        "*InstanceParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists instances of a given cloud, server array.
Using the available filters, it is possible to craft powerful queries about which instances to retrieve.
For example, one can easily list:
* instances that have names that contain "app"
* all instances of a given deployment
* instances belonging to a given server array (i.e., have the same parent_url)
To see the instances of a server array including the next_instance, use the URL "/api/clouds/:cloud_id/instances" with the filter "parent_href==/api/server_arrays/XX". To list only the running
instances of a server array, use the URL "/api/server_arrays/:server_array_id/current_instances"
The 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
details please see Inputs#index.)
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s/current_instances",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/current_instances$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"datacenter_href", "deployment_href", "name", "os_platform", "parent_href", "placement_group_href", "private_dns_name", "private_ip_address", "public_dns_name", "public_ip_address", "resource_uid", "server_template_href", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny", "sensitive"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"datacenter_href", "deployment_href", "name", "os_platform", "parent_href", "placement_group_href", "private_dns_name", "private_ip_address", "public_dns_name", "public_ip_address", "resource_uid", "server_template_href", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny", "sensitive"},
					},
				},
			},

			&metadata.Action{
				Name: "launch",
				Description: `Launches an instance using the parameters that this instance has been configured with.
Note that this action can only be performed in "next" instances, and not on instances that are already running.
Optional parameters:
	api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
	count: For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error
	inputs`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/launch",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/launch$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/launch",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/launch$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/launch",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/launch$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of that input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
					&metadata.ActionParam{
						Name:        "count",
						Description: `For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
					&metadata.ActionParam{
						Name:        "count",
						Description: `For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "lock",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/lock",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/lock$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "multi_run_executable",
				Description: `Runs a script or a recipe in the running instances.
This is an asynchronous function, which returns immediately after queuing the executable for execution.
Status of the execution can be tracked at the URL returned in the "Location" header.
Optional parameters:
	filter
	ignore_lock: Specifies the ability to ignore the lock(s) on the Instance(s).
	inputs
	recipe_name: The name of the recipe to be run.
	right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/multi_run_executable",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/multi_run_executable$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/multi_run_executable",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_run_executable$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of these inputs. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The name of inputs needed. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to be run.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock(s) on the Instance(s).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock(s) on the Instance(s).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to be run.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "multi_terminate",
				Description: `Terminates running instances.
Either a filter or the parameter 'terminate_all' must be provided.
Optional parameters:
	filter
	terminate_all: Specifies the ability to terminate all instances.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/multi_terminate",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/multi_terminate$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/multi_terminate",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_terminate$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "terminate_all",
						Description: `Specifies the ability to terminate all instances.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "terminate_all",
						Description: `Specifies the ability to terminate all instances.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "reboot",
				Description: `Reboot a running instance.
Note that this action can only succeed if the instance is running. One cannot reboot instances of type "next".`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/reboot",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/reboot$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/reboot",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/reboot$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "run_executable",
				Description: `Runs a script or a recipe in the running instance.
This is an asynchronous function, which returns immediately after queuing the executable for execution.
Status of the execution can be tracked at the URL returned in the "Location" header.
Note that this can only be performed on running instances.
Optional parameters:
	ignore_lock: Specifies the ability to ignore the lock on the Instance.
	inputs
	recipe_name: The name of the recipe to run.
	right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/run_executable",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/run_executable$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of these inputs. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The name of inputs needed. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to run.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock on the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock on the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to run.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows attributes of a single instance.
The 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
details please see Inputs#index.)
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny", "sensitive"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny", "sensitive"},
					},
				},
			},

			&metadata.Action{
				Name: "start",
				Description: `Starts an instance that has been stopped, resuming it to its previously saved volume state.
After an instance is started, the reference to your instance will have a different id.
The new id can be found by performing an index query with the appropriate filters on the
Instances resource, performing a show action on the Server resource for Server Instances, or
performing a current_instances action on the ServerArray resource for ServerArray Instances.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/start",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/start$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "stop",
				Description: `Stores the instance's current volume state to resume later using the 'start' action.
After an instance is stopped, the reference to your instance will have a different id.
The new id can be found by performing an index query with the appropriate filters on the
Instances resource, performing a show action on the Server resource for Server Instances, or performing a current_instances action on the ServerArray resource for ServerArray Instances.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/stop",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/stop$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "terminate",
				Description: `Terminates a running instance.
Note that this action can succeed only if the instance is running. One cannot terminate instances of type "next".`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/terminate",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/terminate$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/terminate",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/terminate$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "unlock",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/unlock",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/unlock$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a single instance.
Required parameters:
	instance`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/instances/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][create_default_port_forwarding_rules]",
						Description: `Automatically create default port forwarding rules (enabled by default). Supported by Azure cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][local_ssd_interface]",
						Description: `The type of SSD(s) to be created. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][create_boot_volume]",
						Description: `If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][delete_boot_volume]",
						Description: `If enabled, the associated volume will be deleted when the instance is terminated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][placement_tenancy]",
						Description: `The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "dedicated"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][availability_set]",
						Description: `Availability set for raw instance. Supported by Azure v2 cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][local_ssd_count]",
						Description: `Additional local SSDs. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][service_account]",
						Description: `Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][keep_alive_url]",
						Description: `The ulr of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][admin_username]",
						Description: `The user that will be granted administrative privileges. Supported by AzureRM cloud only. For more information, <a href="http://docs.rightscale.com/clouds/azure_resource_manager/reference/limitations.html">review the documentation</a>.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][max_spot_price]",
						Description: `Specify the max spot price you will pay for. Required when 'pricing_type' is 'spot'. Only applies to clouds which support spot-pricing and when 'spot' is chosen as the 'pricing_type'. Should be a Float value >= 0.001, eg: 0.095, 0.123, 1.23, etc...`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][keep_alive_id]",
						Description: `The id of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][pricing_type]",
						Description: `Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance.WARNING:  By using spot pricing, you acknowledge that your instance/server/array may not be able to be launched (and arrays may be unable to grow) as newly launched instances might be stuck in bidding, and/or existing instances may be terminated at any time, due to the cloud's spot pricing changes and availability.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"fixed", "spot"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][preemptible]",
						Description: `Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][num_cores]",
						Description: `The number of instance cores. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][memory_mb]",
						Description: `The size of instance memory. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][metadata]",
						Description: `Extra data used for configuration, in query string format.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][disk_gb]",
						Description: `The size of root disk. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[multi_cloud_image_href]",
						Description: `The href of the updated MultiCloudImage for the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[security_group_hrefs][]",
						Description: `The hrefs of the updated security groups.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[server_template_href]",
						Description: `The href of the updated ServerTemplate for the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[instance_type_href]",
						Description: `The href of the updated Instance Type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[private_ip_address]",
						Description: `The private ip address for the instance`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ramdisk_image_href]",
						Description: `The href of the updated ramdisk image for the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[kernel_image_href]",
						Description: `The href of the updated kernel image for the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[deployment_href]",
						Description: `The href of the updated Deployment for the Instance. This is only supported for Instances that are not associated with a Server or ServerArray.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[datacenter_href]",
						Description: `The href of the updated Datacenter / Zone for the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ssh_key_href]",
						Description: `The href of the updated SSH key for the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[image_href]",
						Description: `The href of the updated Image for the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[name]",
						Description: `The updated name to give the Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "instance",
						Description: ``,
						Type:        "*InstanceParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"alert_specs":        "Associated alert specs",
			"alerts":             "Associated alerts",
			"cloud":              "Associated Cloud",
			"datacenter":         "Associated datacenter",
			"deployment":         "Associated Deployment",
			"image":              "Associated image",
			"inputs":             "List of configuration inputs",
			"instance_type":      "Associated instance type",
			"kernel_image":       "Associated kernel image",
			"lock_user":          "Href of the user who has locked this instance",
			"monitoring_metrics": "Associated monitoring metrics",
			"multi_cloud_image":  "Associated multi cloud image",
			"parent":             "Parent Object (Server/ServerArray)",
			"placement_group":    "Associated placement group",
			"ramdisk_image":      "Associated ramdisk image",
			"self":               "Href of itself",
			"server_template":    "Associated ServerTemplate",
			"ssh_key":            "Associated ssh key",
			"volume_attachments": "Associated volume attachments",
		},
	},
	"InstanceType": &metadata.Resource{
		Name:        "InstanceType",
		Description: ``,
		Identifier:  "application/vnd.rightscale.instance_type",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Lists instance types.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instance_types",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instance_types$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cpu_architecture", "description", "name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cpu_architecture", "description", "name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single Instance type.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instance_types/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instance_types/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud": "Associated cloud",
			"self":  "Href of itself",
		},
	},
	"IpAddress": &metadata.Resource{
		Name:        "IpAddress",
		Description: `An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.`,
		Identifier:  "application/vnd.rightscale.ip_address",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new IpAddress with the given parameters.
Required parameters:
	ip_address`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/ip_addresses",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address[deployment_href]",
						Description: `The href of the Deployment that owns this IpAddress.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address[network_href]",
						Description: `(OpenStack Only) The href of the Network that the IpAddress will be associated to. This parameter is required for OpenStack with Neutron clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address[domain]",
						Description: `(Amazon Only) Pass vpc to create this IP for EC2-VPC only environments. Pass ec2_classic to create this IP for EC2-Classic environments. Defaults to ec2_classic.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"ec2_classic", "vpc"},
					},
					&metadata.ActionParam{
						Name:        "ip_address[name]",
						Description: `The name of the IpAddress to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address",
						Description: ``,
						Type:        "*IpAddressParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given IpAddress.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/ip_addresses/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the IpAddresses available to this account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_addresses",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"deployment_href", "name"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"deployment_href", "name"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single IpAddress.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_addresses/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a given IpAddress.
Required parameters:
	ip_address`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/ip_addresses/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address[deployment_href]",
						Description: `The href of the Deployment that owns this IpAddress.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address[name]",
						Description: `The updated name of the IpAddress.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address",
						Description: ``,
						Type:        "*IpAddressParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"deployment":          "Containing Deployment",
			"ip_address_bindings": "Associated ip address bindings",
			"network":             "Associated network",
			"self":                "Href of itself",
		},
	},
	"IpAddressBinding": &metadata.Resource{
		Name: "IpAddressBinding",
		Description: `An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.
The IpAddress is bound immediately for a current instance, or on launch for a next instance.
It also allows specifying port forwarding rules for that particular IpAddress and Instance pair.`,
		Identifier: "application/vnd.rightscale.ip_address_binding",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates an ip address binding which attaches a specified IpAddress resource
to a specified instance, and also allows for configuration of port forwarding
rules. If the instance specified is a current (running) instance, a one-time
IpAddressBinding will be created. If the instance is a next instance, then
a recurring IpAddressBinding is created, which will cause the IpAddress to
be bound each time the incarnator boots.
Required parameters:
	ip_address_binding`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/ip_addresses/%s/ip_address_bindings",
						Variables:  []string{"cloud_id", "ip_address_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/ip_address_bindings",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address_binding[public_ip_address_href]",
						Description: `The IpAddress to bind to the specified instance. Required unless port forwarding rule params are passed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[instance_href]",
						Description: `The Instance to which this IpAddress should be bound. Mutually exclusive with server_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[private_port]",
						Description: `Incoming network traffic will get forwarded to this port number on the specified Instance. If not specified, will use public port. Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[server_href]",
						Description: `The Server to which this IpAddress should be bound. Mutually exclusive with instance_href.Note: the Server must have a current_instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[public_port]",
						Description: `The incoming port for port forwarding. Incoming network traffic on this port will get forwarded (to the IP:Private Port of the specified Instance). Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[protocol]",
						Description: `Transport layer protocol of traffic that may be forwarded from public port to private port on the Instance. Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"UDP", "TCP"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address_binding",
						Description: ``,
						Type:        "*IpAddressBindingParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `No description provided for destroy.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/ip_addresses/%s/ip_address_bindings/%s",
						Variables:  []string{"cloud_id", "ip_address_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/ip_address_bindings/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the ip address bindings available to this account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_addresses/%s/ip_address_bindings",
						Variables:  []string{"cloud_id", "ip_address_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_address_bindings",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"instance_href", "ip_address_href"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"instance_href", "ip_address_href"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single ip address binding.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_addresses/%s/ip_address_bindings/%s",
						Variables:  []string{"cloud_id", "ip_address_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_address_bindings/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
		Links: map[string]string{
			"instance":   "Href of the Instance to which the IpAddress is bound",
			"ip_address": "Href of the IpAddress bound to the Instance",
			"self":       "Href of itself",
		},
	},
	"MonitoringMetric": &metadata.Resource{
		Name:        "MonitoringMetric",
		Description: `A monitoring metric is a stream of data that is captured in an instance. Metrics can be monitored, graphed and can be used as the basis for triggering alerts.`,
		Identifier:  "application/vnd.rightscale.monitoring_metric",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "data",
				Description: `Gives the raw monitoring data for a particular metric. The response will include different variables
associated with that metric and the data points for each of those variables.
To get the data for a certain duration, for e.g. for the last 10 minutes(600 secs), provide the variables
start="-600" and end="0".
Required parameters:
	end: An integer number of seconds from current time. e.g. -150 or 0
	start: An integer number of seconds from current time. e.g. -300`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/monitoring_metrics/%s/data",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/monitoring_metrics/([^/]+)/data$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "start",
						Description: `An integer number of seconds from current time. e.g. -300`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "end",
						Description: `An integer number of seconds from current time. e.g. -150 or 0 `,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "end",
						Description: `An integer number of seconds from current time. e.g. -150 or 0 `,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "start",
						Description: `An integer number of seconds from current time. e.g. -300`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the monitoring metrics available for the instance and their corresponding graph hrefs.
Making a request to the graph_href will return a png image corresponding to that monitoring metric.
Optional parameters:
	filter
	period: The time scale for which the graph is generated. Default is 'day'
	size: The size of the graph to be generated. Default is 'small'.
	title: The title of the graph.
	tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/monitoring_metrics",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/monitoring_metrics$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"plugin", "view"},
					},
					&metadata.ActionParam{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&metadata.ActionParam{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&metadata.ActionParam{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"plugin", "view"},
					},
					&metadata.ActionParam{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&metadata.ActionParam{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&metadata.ActionParam{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows attributes of a single monitoring metric.
Making a request to the graph_href will return a png image corresponding to that monitoring metric.
Optional parameters:
	period: The time scale for which the graph is generated. Default is 'day'.
	size: The size of the graph to be generated. Default is 'small'.
	title: The title of the graph.
	tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/monitoring_metrics/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/monitoring_metrics/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&metadata.ActionParam{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&metadata.ActionParam{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&metadata.ActionParam{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&metadata.ActionParam{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"data": "Href for getting raw monitoring data for the metric",
			"self": "Href of itself",
		},
	},
	"MultiCloudImage": &metadata.Resource{
		Name: "MultiCloudImage",
		Description: `A MultiCloudImage is a RightScale component that functions as a pointer to machine images in specific clouds
(e.g. AWS US-East, Rackspace). Each ServerTemplate can reference many MultiCloudImages that defines which
image should be used when a server is launched in a particular cloud.`,
		Identifier: "application/vnd.rightscale.multi_cloud_image",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "clone",
				Description: `Clones a given MultiCloudImage.
Required parameters:
	multi_cloud_image`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/multi_cloud_images/%s/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/clone$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image[description]",
						Description: `The description for the cloned MultiCloudImage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image[name]",
						Description: `The name for the cloned MultiCloudImage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image",
						Description: ``,
						Type:        "*MultiCloudImageParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "commit",
				Description: `Commits a given MultiCloudImage. Only HEAD revisions can be committed.
Required parameters:
	commit_message: The message associated with the commit.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/multi_cloud_images/%s/commit",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/commit$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a new MultiCloudImage with the given parameters.
Required parameters:
	multi_cloud_image`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/multi_cloud_images",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image[description]",
						Description: `The description of the MultiCloudImage to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image[name]",
						Description: `The name of the MultiCloudImage to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image",
						Description: ``,
						Type:        "*MultiCloudImageParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given MultiCloudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_templates/%s/multi_cloud_images/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/multi_cloud_images/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/multi_cloud_images",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name", "revision"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name", "revision"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/multi_cloud_images/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/multi_cloud_images/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).
Currently, the attributes you can update are only the 'direct' attributes of a server template.
Required parameters:
	multi_cloud_image`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/server_templates/%s/multi_cloud_images/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/multi_cloud_images/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image[description]",
						Description: `The updated description for the MultiCloudImage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image[name]",
						Description: `The updated name for the MultiCloudImage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image",
						Description: ``,
						Type:        "*MultiCloudImageParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"matchers": "Associated multi cloud image setting matchers",
			"self":     "Href of itself",
			"settings": "Associated multi cloud image settings",
		},
	},
	"MultiCloudImageMatcher": &metadata.Resource{
		Name: "MultiCloudImageMatcher",
		Description: `A MultiCloudImageMatcher generates MultiCloudImageSettings for all clouds of a
given cloud type. For now, only one type of matcher is supported
(fingerprint). Fingerprint will match images based upon a checksum as returned
by the cloud and is supported CloudStack, OpenStack, and vSphere clouds. Pass
in an example image with an image_href from which to generate the fingerprint.`,
		Identifier: "application/vnd.rightscale.multi_cloud_image_matcher",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new setting matcher for an existing MultiCloudImage.
Required parameters:
	multi_cloud_image_matcher`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/multi_cloud_images/%s/matchers",
						Variables:  []string{"multi_cloud_image_id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/matchers$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_matcher[image_href]",
						Description: `The href of the example Image to use. Mandatory if specifying fingerprint type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_matcher[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_matcher",
						Description: ``,
						Type:        "*MultiCloudImageMatcherParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a MultiCloudImage setting matcher.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/multi_cloud_images/%s/matchers/%s",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/matchers/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the MultiCloudImage setting matchers.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/multi_cloud_images/%s/matchers",
						Variables:  []string{"multi_cloud_image_id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/matchers$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "rematch",
				Description: `Generates new MultiCloudImageSettings based upon match_criteria. Returns hash of created/updated/destroyed settings.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/multi_cloud_images/%s/matchers/%s/rematch",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/matchers/([^/]+)/rematch$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single MultiCloudImage setting matcher.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/multi_cloud_images/%s/matchers/%s",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/matchers/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
		Links: map[string]string{
			"cloud":             "Associated cloud",
			"multi_cloud_image": "Associated multi cloud image",
			"self":              "Href of itself",
		},
	},
	"MultiCloudImageSetting": &metadata.Resource{
		Name: "MultiCloudImageSetting",
		Description: `A MultiCloudImageSetting defines which
settings should be used when a server is launched in a cloud.`,
		Identifier: "application/vnd.rightscale.multi_cloud_image_setting",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new setting for an existing MultiCloudImage.
Required parameters:
	multi_cloud_image_setting`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/multi_cloud_images/%s/settings",
						Variables:  []string{"multi_cloud_image_id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[instance_type_href]",
						Description: `The href of the instance type. Mandatory if specifying cloud_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[ramdisk_image_href]",
						Description: `The href of the ramdisk image. Optional if specifying cloud_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[kernel_image_href]",
						Description: `The href of the kernel image. Optional if specifying cloud_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[cloud_href]",
						Description: `The href of the Cloud to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[image_href]",
						Description: `The href of the Image to use. Mandatory if specifying cloud_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting",
						Description: ``,
						Type:        "*MultiCloudImageSettingParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a MultiCloudImage setting.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/multi_cloud_images/%s/settings/%s",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the MultiCloudImage settings.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/multi_cloud_images/%s/settings",
						Variables:  []string{"multi_cloud_image_id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single MultiCloudImage setting.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/multi_cloud_images/%s/settings/%s",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates a settings for a MultiCloudImage.
Required parameters:
	multi_cloud_image_setting`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/multi_cloud_images/%s/settings/%s",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[instance_type_href]",
						Description: `The href of the instance type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[ramdisk_image_href]",
						Description: `The href of the ramdisk image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[kernel_image_href]",
						Description: `The href of the kernel image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[cloud_href]",
						Description: `The href of the Cloud to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting",
						Description: ``,
						Type:        "*MultiCloudImageSettingParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":             "Associated cloud",
			"image":             "Associated image",
			"instance_type":     "Associated instance type",
			"multi_cloud_image": "Associated multi cloud image",
			"self":              "Href of itself",
		},
	},
	"Network": &metadata.Resource{
		Name:        "Network",
		Description: `A Network is a logical grouping of network devices.`,
		Identifier:  "application/vnd.rightscale.network",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new network.
Required parameters:
	network`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/networks",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/networks$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network[instance_tenancy]",
						Description: `The launch policy for AWS instances in the Network. Specify 'default' to allow instances to decide their own launch policy. Specify 'dedicated' to force all instances to be launched as 'dedicated'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[deployment_href]",
						Description: `The href of the Deployment that owns this Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[description]",
						Description: `The description for the Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[cidr_block]",
						Description: `The range of IP addresses for the Network. This parameter is required for Amazon clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[cloud_href]",
						Description: `The Cloud to create the Network in`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[name]",
						Description: `The name for the Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network",
						Description: ``,
						Type:        "*NetworkParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the given network(s).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/networks/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/networks/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists networks in this account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/networks",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/networks$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cidr_block", "cloud_href", "deployment_href", "name", "resource_uid"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cidr_block", "cloud_href", "deployment_href", "name", "resource_uid"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows attributes of a single network.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/networks/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/networks/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates the given network.
Required parameters:
	network`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/networks/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/networks/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network[route_table_href]",
						Description: `Sets the default RouteTable for this Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[deployment_href]",
						Description: `The href of the Deployment that owns this Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[description]",
						Description: `The updated description for the Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[name]",
						Description: `The updated name for the Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network",
						Description: ``,
						Type:        "*NetworkParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":               "Href of the Cloud the network is in",
			"default_route_table": "The href of the RouteTable that is currently active",
			"deployment":          "Containing Deployment",
			"self":                "Href of itself",
		},
	},
	"NetworkGateway": &metadata.Resource{
		Name:        "NetworkGateway",
		Description: `A NetworkGateway is an interface that allows traffic to be routed between networks.`,
		Identifier:  "application/vnd.rightscale.network_gateway",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a new NetworkGateway.
Required parameters:
	network_gateway`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/network_gateways",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_gateways$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_gateway[description]",
						Description: `The description to be set on the NetworkGateway.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[cloud_href]",
						Description: `The cloud to create the NetworkGateway in.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[name]",
						Description: `The name to be set on the NetworkGateway.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[type]",
						Description: `The type of the NetworkGateway.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"vpn", "internet"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_gateway",
						Description: ``,
						Type:        "*NetworkGatewayParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing NetworkGateway.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/network_gateways/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_gateways/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the NetworkGateways available to this account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/network_gateways",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_gateways$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "name", "network_href"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "name", "network_href"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single NetworkGateway.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/network_gateways/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_gateways/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update an existing NetworkGateway.
Required parameters:
	network_gateway`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/network_gateways/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_gateways/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_gateway[network_href]",
						Description: `Pass a blank string to detach from the specified Network, or pass a valid Network href to attach to the specified network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[description]",
						Description: `The description to be set on the NetworkGateway.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[name]",
						Description: `The name to be set on the NetworkGateway.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_gateway",
						Description: ``,
						Type:        "*NetworkGatewayParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":   "The cloud the NetworkGateway is in",
			"network": "The href of the Network this NetworkGateway is attached to, if any",
			"self":    "Href of itself",
		},
	},
	"NetworkOptionGroup": &metadata.Resource{
		Name: "NetworkOptionGroup",
		Description: `A key/value pair hash containing options for configuring a Network.
The key/value pairs are stored in the "options" parameter.
Keys correspond to the type of option to set, and values correspond
to the value of the particular option being set.
Option keys that are supported vary depending on cloud -- please consult
your particular cloud's documentation for available option keys.`,
		Identifier: "application/vnd.rightscale.network_option_group",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a new NetworkOptionGroup.
Required parameters:
	network_option_group`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/network_option_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group[description]",
						Description: `Description of this NetworkOptionGroup`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[cloud_href]",
						Description: `The Cloud to create this NetworkOptionGroup in`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[options]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[type]",
						Description: `Type of this NetworkOptionGroup`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[name]",
						Description: `Name of this NetworkOptionGroup`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group",
						Description: ``,
						Type:        "*NetworkOptionGroupParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing NetworkOptionGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/network_option_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `List NetworkOptionGroups available in this account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/network_option_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "description", "name", "type"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "description", "name", "type"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single NetworkOptionGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/network_option_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update an existing NetworkOptionGroup.
Required parameters:
	network_option_group`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/network_option_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group[description]",
						Description: `Update the description`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[options]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[name]",
						Description: `Update the name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group",
						Description: ``,
						Type:        "*NetworkOptionGroupParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud": "Associated cloud.",
			"self":  "This NetworkOptionGroup's href",
		},
	},
	"NetworkOptionGroupAttachment": &metadata.Resource{
		Name: "NetworkOptionGroupAttachment",
		Description: `Resource for attaching NetworkOptionGroups to Networks.
A single NetworkOptionGroup can be attached to many Networks.
A Network/Subnet can have many NetworkOptionGroups attached, as long as the
NetworkOptionGroups each have different types.
This resource describes the attachment details between a particular
NetworkOptionGroup and Network.
Amazon currently only supports attaching NetworkOptionGroups to Networks.
Other clouds in the future may support attaching to Subnets.`,
		Identifier: "application/vnd.rightscale.network_option_group_attachment",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a new NetworkOptionGroupAttachment.
Required parameters:
	network_option_group_attachment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/network_option_group_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group_attachment[network_option_group_href]",
						Description: `The NetworkOptionGroup to attach to the specified resource.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_option_group_attachment[network_href]",
						Description: `The Network to attach the specified NetworkOptionGroup to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group_attachment",
						Description: ``,
						Type:        "*NetworkOptionGroupAttachmentParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing NetworkOptionGroupAttachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/network_option_group_attachments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `List NetworkOptionGroupAttachments in this account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/network_option_group_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "network_href", "network_option_group_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "network_href", "network_option_group_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single NetworkOptionGroupAttachment.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/network_option_group_attachments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update an existing NetworkOptionGroupAttachment.
Required parameters:
	network_option_group_attachment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/network_option_group_attachments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group_attachment[network_option_group_href]",
						Description: `The NetworkOptionGroup to attach to the specified resource.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group_attachment",
						Description: ``,
						Type:        "*NetworkOptionGroupAttachmentParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":                "Associated Cloud",
			"network":              "Network being attached with the NetworkOptionGroup",
			"network_option_group": "NetworkOptionGroup being attached to a networking resource",
			"self":                 "Href of itself",
		},
	},
	"Oauth2": &metadata.Resource{
		Name:        "Oauth2",
		Description: ``,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Perform an OAuth 2.0 token_refresh operation to obtain an access token that
can be used in lieu of an API session cookie. (In other words, creates a
session using OAuth 2.0).
Note that an API-Version header is required with your request, and that the server
may respond with a 301 Moved Permanently if you include an account_id parameter and
your account is hosted in another RightScale cluster.
The request parameters and response format are all as per the OAuth 2.0
Internet Draft standard v23. In brief:
 - Successful responses include an access token, an expires-in timestamp, and a token type
 - The token type is always "bearer"
 - To use a bearer token, include header "Authorization: Bearer " with your API requests
 - The client must refresh the access token before it expires
# Example Request using Curl (with prettified response):
curl -i -H X-API-Version:1.5 -x POST https://my.rightscale.com/api/oauth2 -d "grant_type=refresh_token" -d "refresh_token=abcd1234deadbeef"
{
  "access_token": "xyzzy",
  "expires_in":   3600,
  "token_type":   "bearer"
}
Required parameters:
	grant_type: Type of grant.
Optional parameters:
	account_id: The client's account ID (only needed for instance agent clients).
	client_id: The client ID (only needed for confidential clients).
	client_secret: The client secret (only needed for confidential clients).
	r_s_version: The RightAgent protocol version the client conforms to (only needed for instance agent clients).
	refresh_token: The refresh token obtained from OAuth grant.
	right_link_version: The RightLink gem version the client conforms to (only needed for instance agent clients).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/oauth2/",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/oauth2/$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_link_version",
						Description: `The RightLink gem version the client conforms to (only needed for instance agent clients).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "client_secret",
						Description: `The client secret (only needed for confidential clients).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "refresh_token",
						Description: `The refresh token obtained from OAuth grant.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "r_s_version",
						Description: `The RightAgent protocol version the client conforms to (only needed for instance agent clients).`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "grant_type",
						Description: `Type of grant.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"refresh_token"},
					},
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The client's account ID (only needed for instance agent clients).`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "client_id",
						Description: `The client ID (only needed for confidential clients).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The client's account ID (only needed for instance agent clients).`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "client_id",
						Description: `The client ID (only needed for confidential clients).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "client_secret",
						Description: `The client secret (only needed for confidential clients).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "grant_type",
						Description: `Type of grant.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"refresh_token"},
					},
					&metadata.ActionParam{
						Name:        "r_s_version",
						Description: `The RightAgent protocol version the client conforms to (only needed for instance agent clients).`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "refresh_token",
						Description: `The refresh token obtained from OAuth grant.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_link_version",
						Description: `The RightLink gem version the client conforms to (only needed for instance agent clients).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Permission": &metadata.Resource{
		Name: "Permission",
		Description: `  Please note that API 1.5 does not support operations on Governance Groups
  or Orgs and only allows management of the following CM Roles:
    admin, actor, observer,
    aws_architect, publisher,
    designer, signup_wiz,
    enterprise_manager, server_login,
    library, security_manager,
    instance, server_superuser,
    infrastructure, ss_end_user,
    ss_designer, ss_observer
  Moreover, this API allows management of only roles granted directly
  on an account, to an individual user.`,
		Identifier: "application/vnd.rightscale.permission",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a permission, thereby granting some user a particular role
with respect to the current account.
The 'observer' role has a special status; it must be granted before
a user is eligible for any other permission in a given account.
When provisioning users, always create the observer permission FIRST;
creating any other permission before it will result in an error.
For more information about the roles available and the privileges
they confer, please refer to the following page of the RightScale
support portal:
  http://support.rightscale.com/15-References/Lists/List_of_User_Roles
Required parameters:
	permission`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/permissions",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/permissions$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "permission[role_title]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "permission[user_href]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "permission",
						Description: ``,
						Type:        "*PermissionParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "destroy",
				Description: `Destroy a permission, thereby revoking a user's role with respect
to the current account.
The 'observer' role has a special status; it cannot be revoked if
a user has any other roles, because other roles become useless without
being able to read data pertaining to the account.
When deprovisioning user, always destroy the observer permission LAST;
destroying it while the user has other permissions will result in an error.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/permissions/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/permissions/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `List all permissions for all users of the current account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/permissions",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/permissions$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"user_href"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"user_href"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single permission.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/permissions/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/permissions/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
		Links: map[string]string{
			"account": "Href of the account to which this permission belongs",
			"self":    "Href of itself",
			"user":    "Href of the user to which this permission belongs",
		},
	},
	"PlacementGroup": &metadata.Resource{
		Name:        "PlacementGroup",
		Description: ``,
		Identifier:  "application/vnd.rightscale.placement_group",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a PlacementGroup.
Required parameters:
	placement_group`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/placement_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/placement_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "placement_group[cloud_specific_attributes][account_type]",
						Description: `AzureRM: The type of Storage Account.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "placement_group[deployment_href]",
						Description: `The href of the Deployment that owns this PlacementGroup.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "placement_group[description]",
						Description: `The description of the Placement Group to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "placement_group[cloud_href]",
						Description: `The Href of the Cloud in which the PlacementGroup should be created. Note: This feature is not supported for all clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "placement_group[name]",
						Description: `The name of the Placement Group to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "placement_group",
						Description: ``,
						Type:        "*PlacementGroupParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Destroys a PlacementGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/placement_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/placement_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all PlacementGroups in an account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/placement_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/placement_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "deployment_href", "name", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "deployment_href", "name", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows information about a single PlacementGroup.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/placement_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/placement_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":      "Href of the Cloud that this PlacementGroup belongs to",
			"deployment": "Containing Deployment",
			"self":       "Href of itself",
		},
	},
	"Preference": &metadata.Resource{
		Name:        "Preference",
		Description: `A Preference is a user and account-specific setting. Preferences are used in many part of the RightScale platform and can be used for custom purposes if desired.`,
		Identifier:  "application/vnd.rightscale.preference",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the given preference.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/preferences/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/preferences/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all preferences.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/preferences",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/preferences$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single preference.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/preferences/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/preferences/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `If 'id' is known, updates preference with given contents.
Otherwise, creates new preference.
Note: If create, will return '201 Created' and the location of the new preference.
Required parameters:
	preference`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/preferences/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/preferences/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "preference[contents]",
						Description: `The updated contents for the Preference.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "preference",
						Description: ``,
						Type:        "*PreferenceParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"self": "Href of itself",
		},
	},
	"Publication": &metadata.Resource{
		Name: "Publication",
		Description: `A Publication is a revisioned component shared with a set of Account Groups.
If shared with your account, it can be imported in to your account.`,
		Identifier: "application/vnd.rightscale.publication",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "import",
				Description: `Imports the given publication and its subordinates to this account.
Only non-HEAD revisions that are shared with the account can be imported.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/publications/%s/import",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/publications/([^/]+)/import$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the publications available to this account. Only non-HEAD revisions are possible.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/publications",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/publications$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name", "publisher", "revision"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name", "publisher", "revision"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single publication. Only non-HEAD revisions are possible.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/publications/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/publications/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"lineage": "Href of Publication Lineage",
			"self":    "Href of itself",
		},
	},
	"PublicationLineage": &metadata.Resource{
		Name: "PublicationLineage",
		Description: `A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.
It is shared among all revisions of a Publication within the marketplace.
Publication Lineages are different than lineages that exist within an account.`,
		Identifier: "application/vnd.rightscale.publication_lineage",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "show",
				Description: `Show information about a single publication lineage. Only non-HEAD revisions are possible.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/publication_lineages/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/publication_lineages/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"self": "Href of itself",
		},
	},
	"RecurringVolumeAttachment": &metadata.Resource{
		Name:        "RecurringVolumeAttachment",
		Description: `A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a Server/ServerArray the next time an instance is launched.`,
		Identifier:  "application/vnd.rightscale.recurring_volume_attachment",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new recurring volume attachment.
Required parameters:
	recurring_volume_attachment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/recurring_volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volumes/%s/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volume_snapshots/%s/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_snapshot_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[volume_type_href]",
						Description: `The href of the volume type. Can be required by some clouds in case if you attaching volume snapshot.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[runnable_href]",
						Description: `The href of the server or server array to attach to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[storage_href]",
						Description: `The href of the volume or volume snapshot to be attached on launch of a next instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[settings]",
						Description: `Additional parameters concerning created attachment. For example, ':delete_on_termination => true' will schedule volume deletion if instance was terminated.`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[device]",
						Description: `The device location where the volume or volume snapshot will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment",
						Description: ``,
						Type:        "*RecurringVolumeAttachmentParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given recurring volume attachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/recurring_volume_attachments/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volumes/%s/recurring_volume_attachments/%s",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volume_snapshots/%s/recurring_volume_attachments/%s",
						Variables:  []string{"cloud_id", "volume_snapshot_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all recurring volume attachments.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/recurring_volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes/%s/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_snapshots/%s/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_snapshot_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"runnable_href", "storage_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"runnable_href", "storage_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single recurring volume attachment.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/recurring_volume_attachments/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes/%s/recurring_volume_attachments/%s",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_snapshots/%s/recurring_volume_attachments/%s",
						Variables:  []string{"cloud_id", "volume_snapshot_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":    "Associated cloud",
			"runnable": "Associated Server/ServerArray",
			"self":     "Href of itself",
			"storage":  "Associated Volume/VolumeSnapshot",
		},
	},
	"Repository": &metadata.Resource{
		Name: "Repository",
		Description: `A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.
RightScale currently supports the following types of repositores: git, svn, and URLs of compressed files (tar, tgz, gzip).`,
		Identifier: "application/vnd.rightscale.repository",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "cookbook_import",
				Description: `Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.
Required parameters:
	asset_hrefs: Hrefs of the assets that should be imported.
Optional parameters:
	follow: A flag indicating whether imported cookbooks should be followed.
	namespace: The namespace to import into.
	repository_commit_reference: Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.
	with_dependencies: A flag indicating whether dependencies should automatically be imported.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/repositories/%s/cookbook_import",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/cookbook_import$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository_commit_reference",
						Description: `Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "with_dependencies",
						Description: `A flag indicating whether dependencies should automatically be imported.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "asset_hrefs[]",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
					&metadata.ActionParam{
						Name:        "follow",
						Description: `A flag indicating whether imported cookbooks should be followed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "asset_hrefs",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "follow",
						Description: `A flag indicating whether imported cookbooks should be followed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
					&metadata.ActionParam{
						Name:        "repository_commit_reference",
						Description: `Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "with_dependencies",
						Description: `A flag indicating whether dependencies should automatically be imported.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "cookbook_import_preview",
				Description: `Retrieves a preview of the effects of a Cookbook import.
NOTE: This action is for RightScale internal use only. The response is
free-form JSON with no associated mediatype.
DO NOT USE, THIS ACTION IS SUBJECT TO CHANGE AT ANYTIME.
Required parameters:
	asset_hrefs: Hrefs of the assets that should be imported.
	namespace: The namespace to import into.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/repositories/%s/cookbook_import_preview",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/cookbook_import_preview$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "asset_hrefs[]",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "asset_hrefs",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a Repository.
The following types of inputs are supported for the credential fields:
Type
Format
Example(s)
Text string
text:&lt;value&gt;
text:-----BEGIN RSA PRIVATE KEY-----text:secret
Credential value
cred:&lt;value&gt;
cred:my ssh keycred:svn_1_password
Required parameters:
	repository`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/repositories",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/repositories$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository[asset_paths][cookbooks][]",
						Description: `The cookbook paths for the repository`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][password]",
						Description: `The password, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][username]",
						Description: `The user name, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][ssh_key]",
						Description: `The SSH key, or credential, for the repository (only valid for git repositories).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[commit_reference]",
						Description: `The revision for the repository`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[source_type]",
						Description: `The source type for the repository.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"git", "svn", "download"},
					},
					&metadata.ActionParam{
						Name:        "repository[auto_import]",
						Description: `Whether cookbooks should automatically be imported upon repository creation.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "repository[description]",
						Description: `The description for the repository.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[source]",
						Description: `The URL for the repository.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[name]",
						Description: `The repository name.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository",
						Description: ``,
						Type:        "*RepositoryParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the specified Repositories.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/repositories/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all Repositories for this Account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/repositories",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/repositories$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "refetch",
				Description: `Refetches all RepositoryAssets associated with the Repository.
Note that a refetch simply updates RightScale's view of the contents of the repository.
You must perform an import to use the assets in your design objects (or use the auto import parameter).
Optional parameters:
	auto_import: Whether cookbooks should automatically be imported after repositories are fetched.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/repositories/%s/refetch",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/refetch$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "auto_import",
						Description: `Whether cookbooks should automatically be imported after repositories are fetched.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "auto_import",
						Description: `Whether cookbooks should automatically be imported after repositories are fetched.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "resolve",
				Description: `Show a list of repositories that have imported cookbooks with the given names.
This operation returns a list of repositories that would later satisfy a call
to the swap_repository
action on a ServerTemplate.
Optional parameters:
	imported_cookbook_name: A list of cookbook names that were imported by the repository.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/repositories/resolve",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/repositories/resolve$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "imported_cookbook_name[]",
						Description: `A list of cookbook names that were imported by the repository.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "imported_cookbook_name",
						Description: `A list of cookbook names that were imported by the repository.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows a specified Repository.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/repositories/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates a specified Repository.
The following types of inputs are supported for the credential fields:
Type
Format
Example(s)
Text string
text:&lt;value&gt;
text:-----BEGIN RSA PRIVATE KEY-----text:secret
Credential value
cred:&lt;value&gt;
cred:my ssh keycred:svn_1_password
Required parameters:
	repository`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/repositories/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository[asset_paths][cookbooks][]",
						Description: `The updated cookbook paths for the repository`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][username]",
						Description: `The updated user name, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][password]",
						Description: `The updated password, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][ssh_key]",
						Description: `The updated SSH key for the repository (only valid for git repositories).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[commit_reference]",
						Description: `The updated commit reference (tag, branch, revision...) for the repository`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[description]",
						Description: `The updated description for the repository.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[source_type]",
						Description: `The updated source type for the repository.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"git", "svn", "download"},
					},
					&metadata.ActionParam{
						Name:        "repository[source]",
						Description: `The updated URL for the repository.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[name]",
						Description: `The updated repository name.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository",
						Description: ``,
						Type:        "*RepositoryParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"repository_assets": "Assets fetched from the repository",
			"self":              "Href of itself",
		},
	},
	"RepositoryAsset": &metadata.Resource{
		Name: "RepositoryAsset",
		Description: `A RepositoryAsset represents an item discovered in a Repository. These assets represent only a view of the Repository
the last time it was scraped. In order to use these assets, you must import them into your account.`,
		Identifier: "application/vnd.rightscale.repository_asset",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `List a repository's current assets.
Repository assests are the cookbook details that were scraped from a
given repository.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/repositories/%s/repository_assets",
						Variables:  []string{"repository_id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/repository_assets$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single asset.
A repository assest are the cookbook details that were scraped from a
repository.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/repositories/%s/repository_assets/%s",
						Variables:  []string{"repository_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/repository_assets/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"self": "Href of itself",
		},
	},
	"ResourceGroup": &metadata.Resource{
		Name:        "ResourceGroup",
		Description: ``,
		Identifier:  "application/vnd.rightscale.resource_group",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a ResourceGroup.
Required parameters:
	resource_group`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/resource_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/resource_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_group[deployment_href]",
						Description: `The Href of the Deployment that owns this Resource Group.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "resource_group[description]",
						Description: `The description of the Resource Group to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "resource_group[cloud_href]",
						Description: `The Href of the Cloud in which the ResourceGroup should be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "resource_group[name]",
						Description: `The name of the Resource Group to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_group",
						Description: ``,
						Type:        "*ResourceGroupParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Destroys a ResourceGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/resource_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/resource_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all ResourceGroups in an account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/resource_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/resource_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "name", "state"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "name", "state"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows information about a single ResourceGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/resource_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/resource_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a given ResourceGroup.
Required parameters:
	resource_group`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/resource_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/resource_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_group[deployment_href]",
						Description: `The Href of the Deployment that owns this Resource Group.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "resource_group[description]",
						Description: `The description of the Resource Group to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_group",
						Description: ``,
						Type:        "*ResourceGroupParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":      "Href of the Cloud that this ResourceGroup belongs to",
			"deployment": "Associated Deployment",
			"self":       "Href of itself",
		},
	},
	"RightScript": &metadata.Resource{
		Name: "RightScript",
		Description: `A RightScript is an executable piece of code that can be run on a server
during the boot, operational, or decommission phases.
All revisions of
a RightScript belong to a RightScript lineage that is exposed by the
"lineage" attribute (NOTE: This attribute is merely a string to locate
all revisions of a RightScript and NOT a working URL).`,
		Identifier: "application/vnd.rightscale.right_script",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "commit",
				Description: `Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.
Required parameters:
	right_script`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/right_scripts/%s/commit",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/commit$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script[commit_message]",
						Description: `The message to be included with the requested commit`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script",
						Description: ``,
						Type:        "*RightScriptParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `No description provided for create.
Required parameters:
	right_script`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/right_scripts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/right_scripts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script[description]",
						Description: `The description of the RightScript to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script[packages]",
						Description: `Space-separated list of package names needed in order to successfully run the script.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script[source]",
						Description: `The script source to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script[name]",
						Description: `The name of the RightScript to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script",
						Description: ``,
						Type:        "*RightScriptParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `No description provided for destroy.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/right_scripts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists RightScripts.
Optional parameters:
	filter
	latest_only: Whether or not to return only the latest version for each lineage.
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/right_scripts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/right_scripts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "latest_only",
						Description: `Whether or not to return only the latest version for each lineage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"lineage", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"lineage", "name"},
					},
					&metadata.ActionParam{
						Name:        "latest_only",
						Description: `Whether or not to return only the latest version for each lineage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single RightScript.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/right_scripts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs_2_0"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs_2_0"},
					},
				},
			},

			&metadata.Action{
				Name:        "show_source",
				Description: `Returns the script source for a RightScript`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/right_scripts/%s/source",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/source$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates RightScript name/description
Required parameters:
	right_script`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/right_scripts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script[description]",
						Description: `The new description for the RightScript`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script[packages]",
						Description: `The new list of packages for the RightScript`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script[source]",
						Description: `The script source to be updated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script[name]",
						Description: `The new name for the RightScript`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script",
						Description: ``,
						Type:        "*RightScriptParam3",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "update_source",
				Description: `Updates the source of the given RightScript
Required parameters:
	filename: The file name to update the RightScript source with.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/right_scripts/%s/source",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/source$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `The file name to update the RightScript source with.`,
						Type:        "sourcefile",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `The file name to update the RightScript source with.`,
						Type:        "*rsapi.SourceUpload",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"self":   "Href of itself",
			"source": "Returns the RightScript source",
		},
	},
	"RightScriptAttachment": &metadata.Resource{
		Name:        "RightScriptAttachment",
		Description: ``,
		Identifier:  "application/vnd.rightscale.right_script_attachment",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Uploads the RightScript attachment links it to the RightScript. Create expects HTTP request to be formatted as multipart mime.
Required parameters:
	right_script_attachment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/right_scripts/%s/attachments",
						Variables:  []string{"right_script_id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_attachment[filename]",
						Description: `The file name of the RightScript attachment to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script_attachment[content]",
						Description: `The content of the RightScript attachment to be created.`,
						Type:        "file",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_attachment",
						Description: ``,
						Type:        "*RightScriptAttachmentParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/right_scripts/%s/attachments/%s",
						Variables:  []string{"right_script_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists RightScript attachments.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/right_scripts/%s/attachments",
						Variables:  []string{"right_script_id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"filename"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"filename"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single RightScript attachment.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/right_scripts/%s/attachments/%s",
						Variables:  []string{"right_script_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Uploads and updates existing attachment in a RightScript. Update expects HTTP request to formatted as multipart mime.
Required parameters:
	right_script_attachment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/right_scripts/%s/attachments/%s",
						Variables:  []string{"right_script_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/attachments/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_attachment[filename]",
						Description: `The new file name for the RightScript attachment.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script_attachment[content]",
						Description: `The new content for the RightScript attachment.`,
						Type:        "file",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_attachment",
						Description: ``,
						Type:        "*RightScriptAttachmentParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"right_script": "The RightScript to which the attachment is attached",
			"self":         "Href of itself",
		},
	},
	"Route": &metadata.Resource{
		Name: "Route",
		Description: `A Route defines how networking traffic should be routed from one
destination to another. See next_hop_type for available endpoint targets.`,
		Identifier: "application/vnd.rightscale.route",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a new Route.
Required parameters:
	route`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/routes",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/routes$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/route_tables/%s/routes",
						Variables:  []string{"route_table_id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route[cloud_specific_attributes][instance_tags][]",
						Description: `A list of instance tags to which this route applies. Omitting this value will result in creation of global route.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[cloud_specific_attributes][priority]",
						Description: `Priority is used to break ties in the case where there is more than one matching route of maximum length. A lower value is higher priority.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route[destination_cidr_block]",
						Description: `The destination (CIDR IP address) for the Route.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[route_table_href]",
						Description: `The RouteTable to create the Route in.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name: "route[next_hop_href]",
						Description: `The href of the Route's next hop.
                         Required if route[next_hop_type] is 'instance', 'network_interface', or 'network_gateway'.
                         Not allowed otherwise.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_type]",
						Description: `The Route's next hop type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"instance", "network_interface", "network_gateway", "ip_string", "url"},
					},
					&metadata.ActionParam{
						Name: "route[next_hop_url]",
						Description: `The URL of the Route's next hop.
                         Required if route[next_hop_type] is 'url'.
                         Not allowed otherwise.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_ip]",
						Description: `The IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[description]",
						Description: `The description to be set on the Route.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route",
						Description: ``,
						Type:        "*RouteParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing Route.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/routes/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/routes/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/route_tables/%s/routes/%s",
						Variables:  []string{"route_table_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `List Routes available in this account.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/routes",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/routes$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/route_tables/%s/routes",
						Variables:  []string{"route_table_id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "description", "network_href", "next_hop_href", "next_hop_ip", "next_hop_type", "next_hop_url", "route_table_href", "state"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "description", "network_href", "next_hop_href", "next_hop_ip", "next_hop_type", "next_hop_url", "route_table_href", "state"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single Route.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/routes/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/routes/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/route_tables/%s/routes/%s",
						Variables:  []string{"route_table_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update an existing Route.
Required parameters:
	route`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/routes/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/routes/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/route_tables/%s/routes/%s",
						Variables:  []string{"route_table_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route[destination_cidr_block]",
						Description: `The updated destination (CIDR IP address) for the Route.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_href]",
						Description: `The updated href of the Route's next hop. Required if route[next_hop_type] is 'instance', 'network_interface', or 'network_gateway'. Not allowed otherwise.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_type]",
						Description: `The updated Route's next hop type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"instance", "network_interface", "network_gateway", "ip_string"},
					},
					&metadata.ActionParam{
						Name:        "route[description]",
						Description: `The updated description of the Route.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_ip]",
						Description: `The updated IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route",
						Description: ``,
						Type:        "*RouteParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":       "The cloud that this Route is in.",
			"next_hop":    "Href of the next_hop (only if next_hop_type isn't ip_string)",
			"route_table": "The RouteTable that this Route belongs to.",
			"self":        "Link to this Route.",
		},
	},
	"RouteTable": &metadata.Resource{
		Name:        "RouteTable",
		Description: `Grouped listing of Routes`,
		Identifier:  "application/vnd.rightscale.route_table",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a new RouteTable.
Required parameters:
	route_table`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/route_tables",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/route_tables$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route_table[network_href]",
						Description: `The Network to create the RouteTable in.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route_table[description]",
						Description: `The description to be set on the RouteTable.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route_table[cloud_href]",
						Description: `The cloud to create the RouteTable in.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route_table[name]",
						Description: `The name to be set on the RouteTable.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route_table",
						Description: ``,
						Type:        "*RouteTableParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing RouteTable.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/route_tables/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `List RouteTables available in this account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/route_tables",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/route_tables$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "name", "network_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "name", "network_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single RouteTable.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/route_tables/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update an existing RouteTable.
Required parameters:
	route_table`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/route_tables/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route_table[description]",
						Description: `The description to be set on the RouteTable.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route_table[name]",
						Description: `The name to be set on the RouteTable.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route_table",
						Description: ``,
						Type:        "*RouteTableParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":   "Associated Cloud",
			"network": "Associated Network",
			"routes":  "Associated Routes",
			"self":    "Href of itself",
		},
	},
	"RunnableBinding": &metadata.Resource{
		Name: "RunnableBinding",
		Description: `A RunnableBinding represents an item in a runlist of a ServerTemplate. These items could be
RightScript or Chef recipes, and could be associated with any one of the three runlists of a
ServerTemplate (boot, operational, decommission).`,
		Identifier: "application/vnd.rightscale.runnable_binding",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Bind an executable to the given ServerTemplate.
An executable may be either a RightScript or Chef Recipe.
The resource must be editable.
Required parameters:
	runnable_binding`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/runnable_bindings",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding[right_script_href]",
						Description: `The RightScript href. Note: recipe cannot be specified when this param is given.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_binding[position]",
						Description: `The position of the executable in the execution order. If not specified, will be added to the end. If specified, will be inserted in that location and cause all others to move down.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_binding[sequence]",
						Description: `The sequence at which this executable should be run. Default is 'operational'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"boot", "decommission", "operational"},
					},
					&metadata.ActionParam{
						Name:        "runnable_binding[recipe]",
						Description: `The Chef recipe name. Note: right_script_href cannot be specified when this param is given.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding",
						Description: ``,
						Type:        "*RunnableBindingParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "destroy",
				Description: `Unbind an executable from the given resource.
The resource must be editable.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_templates/%s/runnable_bindings/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the executables bound to the ServerTemplate.
An excutable may be either a RightScript or Chef Recipe.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/runnable_bindings",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "multi_update",
				Description: `Update attributes for multiple bound executables.
The resource must be editable.
Required parameters:
	runnable_bindings`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/server_templates/%s/runnable_bindings/multi_update",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings/multi_update$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_bindings[][right_script_href]",
						Description: `The updated RightScript href. Note: recipe cannot be specified when this param is given.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][position]",
						Description: `The updated position of the RunnableBinding in the execution order. If specified, will be inserted in that location and cause all others to move down.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][sequence]",
						Description: `The sequence at which this executable should be run.  Default is 'operational'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"boot", "decommission", "operational"},
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][recipe]",
						Description: `The updated Chef recipe name. Note: right_script_href cannot be specified when this param is given.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][id]",
						Description: `The ID of the RunnableBinding to update.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_bindings",
						Description: ``,
						Type:        "[]*RunnableBindings",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single executable binding.
An excutable may be either a RightScript or Chef Recipe.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s/runnable_bindings/%s",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"right_script":    "Href of associated RightScript (only returned if RightScript executable)",
			"self":            "Href of itself",
			"server_template": "Href of associated ServerTemplate",
		},
	},
	"Scheduler": &metadata.Resource{
		Name:        "Scheduler",
		Description: `Provide RightLink with the ability to schedule script executions on instances`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "schedule_recipe",
				Description: `Schedules a chef recipe for execution on the current instance
Optional parameters:
	arguments: Serialized recipe execution arguments values keyed by name
	audit_id: Optional, reuse audit if specified
	audit_period: RunlistPolicy audit period
	formal_values: Formal input parameter values
	policy: RunlistPolicy policy name
	recipe: Chef recipe name, overridden by recipe_id
	recipe_id: ServerTemplateChefRecipe ID
	thread: RunlistPolicy thread name`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/right_net/scheduler/schedule_recipe",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/right_net/scheduler/schedule_recipe$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "formal_values",
						Description: `Formal input parameter values`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "audit_period",
						Description: `RunlistPolicy audit period`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "arguments",
						Description: `Serialized recipe execution arguments values keyed by name`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "recipe_id",
						Description: `ServerTemplateChefRecipe ID`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_id",
						Description: `Optional, reuse audit if specified`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "policy",
						Description: `RunlistPolicy policy name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "recipe",
						Description: `Chef recipe name, overridden by recipe_id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "thread",
						Description: `RunlistPolicy thread name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments",
						Description: `Serialized recipe execution arguments values keyed by name`,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "audit_id",
						Description: `Optional, reuse audit if specified`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_period",
						Description: `RunlistPolicy audit period`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "formal_values",
						Description: `Formal input parameter values`,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "policy",
						Description: `RunlistPolicy policy name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "recipe",
						Description: `Chef recipe name, overridden by recipe_id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recipe_id",
						Description: `ServerTemplateChefRecipe ID`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "thread",
						Description: `RunlistPolicy thread name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "schedule_right_script",
				Description: `Schedules a RightScript for execution on the current instance
Optional parameters:
	arguments: Serialized script execution arguments values keyed by name
	audit_id: Optional, reuse audit if specified
	audit_period: RunlistPolicy audit period
	formal_values: Formal input parameter values
	policy: RunlistPolicy policy name
	right_script: RightScript name, overridden by right_script_id
	right_script_id: RightScript ID
	thread: RunlistPolicy thread name`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/right_net/scheduler/schedule_right_script",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/right_net/scheduler/schedule_right_script$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_id",
						Description: `RightScript ID`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "formal_values",
						Description: `Formal input parameter values`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "audit_period",
						Description: `RunlistPolicy audit period`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "right_script",
						Description: `RightScript name, overridden by right_script_id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "arguments",
						Description: `Serialized script execution arguments values keyed by name`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "audit_id",
						Description: `Optional, reuse audit if specified`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "policy",
						Description: `RunlistPolicy policy name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "thread",
						Description: `RunlistPolicy thread name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "arguments",
						Description: `Serialized script execution arguments values keyed by name`,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "audit_id",
						Description: `Optional, reuse audit if specified`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_period",
						Description: `RunlistPolicy audit period`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "formal_values",
						Description: `Formal input parameter values`,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "policy",
						Description: `RunlistPolicy policy name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "right_script",
						Description: `RightScript name, overridden by right_script_id`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script_id",
						Description: `RightScript ID`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "thread",
						Description: `RunlistPolicy thread name`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"SecurityGroup": &metadata.Resource{
		Name: "SecurityGroup",
		Description: `Security Groups represent network security profiles that contain lists of firewall rules for different ports and source IP addresses, as well as
trust relationships amongst different security groups.`,
		Identifier: "application/vnd.rightscale.security_group",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a security group.
Required parameters:
	security_group`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/security_groups",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group[deployment_href]",
						Description: `The href of the Deployment that owns this SecurityGroup.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group[network_href]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group[description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group[name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group",
						Description: ``,
						Type:        "*SecurityGroupParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete security group(s)`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/security_groups/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists Security Groups.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/security_groups",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"deployment_href", "name", "network_href", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"deployment_href", "name", "network_href", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single Security Group.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/security_groups/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":                "Associated cloud",
			"deployment":           "Containing Deployment",
			"network":              "Associated network",
			"security_group_rules": "Associated security_group_rules",
			"self":                 "Href of itself",
		},
	},
	"SecurityGroupRule": &metadata.Resource{
		Name:        "SecurityGroupRule",
		Description: ``,
		Identifier:  "application/vnd.rightscale.security_group_rule",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a security group rule for a security group.
The following flavors are supported:
1. group-based TCP/UDP
2. group-based ICMP
3. CIDR-based TCP/UDP
4. CIDR-based ICMP
Required parameters:
	security_group_rule`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/security_group_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/security_groups/%s/security_group_rules",
						Variables:  []string{"cloud_id", "security_group_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][start_port]",
						Description: `Start of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][icmp_code]",
						Description: `ICMP code. Required if protocol is 'icmp'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][icmp_type]",
						Description: `ICMP type. Required if protocol is 'icmp'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][end_port]",
						Description: `End of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[security_group_href]",
						Description: `Security Group to add rule to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[group_owner]",
						Description: `Owner of source Security Group. Required if source_type is 'group'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[source_type]",
						Description: `Source type. May be a CIDR block or another Security Group.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"cidr_ips", "group"},
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[group_name]",
						Description: `Name of source Security Group. Required if source_type is 'group'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[direction]",
						Description: `Direction of traffic.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"ingress", "egress"},
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol]",
						Description: `Protocol to filter on.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"tcp", "udp", "icmp", "all"},
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[priority]",
						Description: `Lower takes precedence. Supported by AzureRM cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[cidr_ips]",
						Description: `An IP address range in CIDR notation. Required if source_type is 'cidr_ips'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[action]",
						Description: `Allow or deny rule. Defaults to allow. Supported by AzureRM cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"allow", "deny"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group_rule",
						Description: ``,
						Type:        "*SecurityGroupRuleParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete security group rule(s)`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/security_group_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/security_groups/%s/security_group_rules/%s",
						Variables:  []string{"cloud_id", "security_group_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists SecurityGroupRules.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/security_group_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/security_groups/%s/security_group_rules",
						Variables:  []string{"cloud_id", "security_group_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single SecurityGroupRule.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/security_group_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/security_groups/%s/security_group_rules/%s",
						Variables:  []string{"cloud_id", "security_group_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Required parameters:
	security_group_rule`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/security_group_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/security_groups/%s/security_group_rules/%s",
						Variables:  []string{"cloud_id", "security_group_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group_rule[description]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group_rule",
						Description: ``,
						Type:        "*SecurityGroupRuleParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"security_group": "Associated security group",
			"self":           "Href of itself",
		},
	},
	"Server": &metadata.Resource{
		Name: "Server",
		Description: `Servers represent the notion of a server/machine from the RightScale's perspective. A Server, does not always
have a corresponding VM running or provisioned in a cloud. Some clouds use the word "servers" to refer to created VM's. These allocated VM's
are not called Servers in the RightScale API, they are called Instances.
A Server always has a next_instance association, which will define the configuration to apply to a new instance when
the server is launched or started (starting servers is not yet supported through this API).
Once a Server is launched/started a current_instance relationship will exist.
Accessing the current_instance of a server results in immediate runtime modification of this running server.
Changes to the next_instance association prepares the
configuration for the next instance launch/start (therefore they have no effect until such operation is performed).`,
		Identifier: "application/vnd.rightscale.server",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "clone",
				Description: `Clones a given server.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/clone$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a new server, and configures its corresponding "next" instance with the received parameters.
Required parameters:
	server`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/servers$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/servers",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][create_default_port_forwarding_rules]",
						Description: `Automatically create default port forwarding rules (enabled by default). Supported by Azure cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][local_ssd_interface]",
						Description: `The type of SSD(s) to be created. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][delete_boot_volume]",
						Description: `If enabled, the associated volume will be deleted when the instance is terminated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][create_boot_volume]",
						Description: `If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][placement_tenancy]",
						Description: `The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "dedicated"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][availability_set]",
						Description: `Availability set for raw instance. Supported by Azure v2 cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][local_ssd_count]",
						Description: `Additional local SSDs. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][service_account]",
						Description: `Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][admin_username]",
						Description: `The user that will be granted administrative privileges. Supported by AzureRM cloud only. For more information, <a href="http://docs.rightscale.com/clouds/azure_resource_manager/reference/limitations.html">review the documentation</a>.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][max_spot_price]",
						Description: `Specify the max spot price you will pay for. Required when 'pricing_type' is 'spot'. Only applies to clouds which support spot-pricing and when 'spot' is chosen as the 'pricing_type'. Should be a Float value >= 0.001, eg: 0.095, 0.123, 1.23, etc...`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][keep_alive_url]",
						Description: `The ulr of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][keep_alive_id]",
						Description: `The id of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][pricing_type]",
						Description: `Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance.WARNING:  By using spot pricing, you acknowledge that your instance/server/array may not be able to be launched (and arrays may be unable to grow) as newly launched instances might be stuck in bidding, and/or existing instances may be terminated at any time, due to the cloud's spot pricing changes and availability.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"fixed", "spot"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][preemptible]",
						Description: `Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][memory_mb]",
						Description: `The size of instance memory. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][num_cores]",
						Description: `The number of instance cores. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][metadata]",
						Description: `Extra data used for configuration, in query string format.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][disk_gb]",
						Description: `The size of root disk. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server[instance][associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][multi_cloud_image_href]",
						Description: `The href of the Multi Cloud Image to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][security_group_hrefs][]",
						Description: `The hrefs of the security groups.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][server_template_href]",
						Description: `The href of the Server Template.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][ramdisk_image_href]",
						Description: `The href of the Ramdisk Image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][private_ip_address]",
						Description: `The private ip address for the instance`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][instance_type_href]",
						Description: `The href of the Instance Type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][kernel_image_href]",
						Description: `The href of the Kernel Image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs][][value]",
						Description: `The value of that Input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs][][name]",
						Description: `The Input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][ssh_key_href]",
						Description: `The href of the SSH key to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_href]",
						Description: `The href of the cloud that the Server should be added to.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[deployment_href]",
						Description: `The href of the deployment to which the Server will be added.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[description]",
						Description: `The Server description.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[optimized]",
						Description: `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[name]",
						Description: `The name of the Server.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server",
						Description: ``,
						Type:        "*ServerParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given server.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/servers/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/deployments/%s/servers/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "disable_runnable_bindings",
				Description: `Disables a list of runnable bindings associated with a given server.
Optional parameters:
	runnable_binding_hrefs: List of Runnable Bindings.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/disable_runnable_bindings",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/disable_runnable_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs[]",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "enable_runnable_bindings",
				Description: `Enables a list of runnable bindings associated with a given server.
Optional parameters:
	runnable_binding_hrefs: List of Runnable Bindings.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/enable_runnable_bindings",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/enable_runnable_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs[]",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists servers.
By using the available filters, it is possible to retrieve servers that have common characteristics.
For example, one can list:
* servers that have names that contain "app_server"
* all servers of a given deployment
For more filters, please see the 'index' action on 'Instances' resource as most of the attributes belong to
a 'current_instance' than to a server.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/servers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/servers$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/servers",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "deployment_href", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "deployment_href", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&metadata.Action{
				Name: "launch",
				Description: `Launches the "next" instance of this server. This function is equivalent to invoking the launch action on the
URL of this servers next_instance. See Instances#launch for details.
Optional parameters:
	api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
	count: For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error
	inputs`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/launch",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/launch$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of that input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
					&metadata.ActionParam{
						Name:        "count",
						Description: `For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
					&metadata.ActionParam{
						Name:        "count",
						Description: `For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows the information of a single server.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/servers/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/servers/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&metadata.Action{
				Name: "terminate",
				Description: `Terminates the current instance of this server. This function is equivalent to invoking the terminate action on the
URL of this servers current_instance. See Instances#terminate for details.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/terminate",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/terminate$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "unwrap",
				Description: `No description provided for unwrap.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/%s/unwrap",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/unwrap$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/servers/%s/unwrap",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/([^/]+)/unwrap$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a single server.
Required parameters:
	server`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/servers/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/deployments/%s/servers/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][delete_boot_volume]",
						Description: `If enabled, the associated volume will be deleted when the instance is terminated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][create_boot_volume]",
						Description: `If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[description]",
						Description: `The updated description for the server.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[optimized]",
						Description: `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[name]",
						Description: `The updated server name.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server",
						Description: ``,
						Type:        "*ServerParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "wrap_instance",
				Description: `Wrap an existing instance and set current instance for new server
Required parameters:
	server`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/servers/wrap_instance",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/servers/wrap_instance$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/servers/wrap_instance",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/wrap_instance$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server[instance][multi_cloud_image_href]",
						Description: `The href of the Multi Cloud Image to use.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][server_template_href]",
						Description: `The href of the Server Template.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[deployment_href]",
						Description: `The href of the deployment to which the Server will be added.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][href]",
						Description: `The href of the Instance around which the server should be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[description]",
						Description: `The Server description.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[name]",
						Description: `The name of the Server.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server",
						Description: ``,
						Type:        "*ServerParam3",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"alert_specs":      "Associated AlertSpecs",
			"alerts":           "Associated Alerts",
			"current_instance": "Associated current instance",
			"deployment":       "Associated deployment",
			"next_instance":    "Associated next instance",
			"self":             "Href of itself",
		},
	},
	"ServerArray": &metadata.Resource{
		Name: "ServerArray",
		Description: `A server array represents a logical group of instances and allows to resize(grow/shrink) that group based on certain elasticity parameters.
A server array just like a server always has a next_instance association, which will define the configuration to apply when a new instance is launched.
But unlike a server which has a current_instance relationship, the server array has a
current_instances relationship that gives the information about
all the running instances in the array. Changes to the next_instance association prepares the configuration for the next instance that is to be launched
in the array and will therefore not affect any of the currently running instances.`,
		Identifier: "application/vnd.rightscale.server_array",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "clone",
				Description: `Clones a given server array.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/clone$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a new server array, and configures its corresponding "next" instance with the received parameters.
Required parameters:
	server_array`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_arrays$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/deployments/%s/server_arrays",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][create_default_port_forwarding_rules]",
						Description: `Automatically create default port forwarding rules (enabled by default). Supported by Azure cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]",
						Description: `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][collect_audit_entries]",
						Description: `The audit SQS queue that will store audit entries.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][voters_tag_predicate]",
						Description: `The Voters Tag that RightScale will use in order to determine when to scale up/down.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][algorithm]",
						Description: `The algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"max_10", "avg_10"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][decision_threshold]",
						Description: `The percentage of servers that must agree in order to trigger an alert before an action is taken.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][max_age]",
						Description: `The threshold (in seconds) before a resize action occurs on the server array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][regexp]",
						Description: `The regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][local_ssd_interface]",
						Description: `The type of SSD(s) to be created. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][delete_boot_volume]",
						Description: `If enabled, the associated volume will be deleted when the instance is terminated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][create_boot_volume]",
						Description: `If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][placement_tenancy]",
						Description: `The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "dedicated"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][availability_set]",
						Description: `Availability set for raw instance. Supported by Azure v2 cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][service_account]",
						Description: `Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][local_ssd_count]",
						Description: `Additional local SSDs. Supported by GCE cloud only`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][admin_username]",
						Description: `The user that will be granted administrative privileges. Supported by AzureRM cloud only. For more information, <a href="http://docs.rightscale.com/clouds/azure_resource_manager/reference/limitations.html">review the documentation</a>.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][max_spot_price]",
						Description: `Specify the max spot price you will pay for. Required when 'pricing_type' is 'spot'. Only applies to clouds which support spot-pricing and when 'spot' is chosen as the 'pricing_type'. Should be a Float value >= 0.001, eg: 0.095, 0.123, 1.23, etc...`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][keep_alive_url]",
						Description: `The ulr of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][keep_alive_id]",
						Description: `The id of keep alive. Supported by UCA cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][pricing_type]",
						Description: `Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance.WARNING:  By using spot pricing, you acknowledge that your instance/server/array may not be able to be launched (and arrays may be unable to grow) as newly launched instances might be stuck in bidding, and/or existing instances may be terminated at any time, due to the cloud's spot pricing changes and availability.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"fixed", "spot"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][preemptible]",
						Description: `Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][memory_mb]",
						Description: `The size of instance memory. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][num_cores]",
						Description: `The number of instance cores. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][metadata]",
						Description: `Extra data used for configuration, in query string format.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][disk_gb]",
						Description: `The size of root disk. Supported by UCA cloud only.`,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_calm_time]",
						Description: `The time (in minutes) on how long you want to wait before you repeat another action.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_down_by]",
						Description: `The number of servers to scale down by.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][min_count]",
						Description: `The minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][max_count]",
						Description: `The maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_up_by]",
						Description: `The number of servers to scale up by.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][min_count]",
						Description: `The minimum number of servers that must be operational at all times in the server array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][max_count]",
						Description: `The maximum number of servers that can be operational at the same time in the server array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][time]",
						Description: `Specifies the time when an alert-based array resizes.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][day]",
						Description: `Specifies the day when an alert-based array resizes.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][multi_cloud_image_href]",
						Description: `The href of the MultiCloudImage to be used.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][security_group_hrefs][]",
						Description: `The hrefs of the Security Groups.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][server_template_href]",
						Description: `The ServerTemplate that will be used to create the worker instances in the server array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][instance_type_href]",
						Description: `The href of the Instance Type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][ramdisk_image_href]",
						Description: `The href of the Ramdisk Image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][weight]",
						Description: `Instance allocation (should total 100%).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][kernel_image_href]",
						Description: `The href of the Kernel Image.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][inputs][][value]",
						Description: `The value of that Input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][inputs]",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][datacenter_href]",
						Description: `The href of the Datacenter / Zone. For multiple Datacenters, use 'datacenter_policy' instead.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][max]",
						Description: `Max instances (0 for unlimited).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][inputs][][name]",
						Description: `The Input name.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][ssh_key_href]",
						Description: `The href of the SSH Key to be used.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][subnet_hrefs][]",
						Description: `The hrefs of the updated Subnets.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_href]",
						Description: `The href of the Cloud that the array will be associated with.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][image_href]",
						Description: `The href of the Image to be used.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[deployment_href]",
						Description: `The href of the deployment for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[description]",
						Description: `The description for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[array_type]",
						Description: `The array type for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"alert", "queue"},
					},
					&metadata.ActionParam{
						Name:        "server_array[optimized]",
						Description: `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[state]",
						Description: `The status of the server array. If active, the server array is enabled for scaling actions.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"enabled", "disabled"},
					},
					&metadata.ActionParam{
						Name:        "server_array[name]",
						Description: `The name for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_array",
						Description: ``,
						Type:        "*ServerArrayParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "current_instances",
				Description: `List the running instances belonging to the server array. See Instances#index for details.
This action is slightly different from invoking the index action on the Instances resource with the filter "parent_href == /api/server_arrays/XX" because the
latter will include 'next_instance' as well.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s/current_instances",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/current_instances$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny", "sensitive"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny", "sensitive"},
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given server array.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_arrays/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/deployments/%s/server_arrays/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "disable_runnable_bindings",
				Description: `Disables a list of runnable bindings associated with a given server.
Optional parameters:
	runnable_binding_hrefs: List of Runnable Bindings.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/disable_runnable_bindings",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/disable_runnable_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs[]",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "enable_runnable_bindings",
				Description: `Enables a list of runnable bindings associated with a given server.
Optional parameters:
	runnable_binding_hrefs: List of Runnable Bindings.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/enable_runnable_bindings",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/enable_runnable_bindings$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs[]",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding_hrefs",
						Description: `List of Runnable Bindings.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists server arrays.
By using the available filters, it is possible to retrieve server arrays that have common characteristics.
For example, one can list:
* arrays that have names that contain "my_server_array"
* all arrays of a given deployment
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_arrays$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/server_arrays",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "deployment_href", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"cloud_href", "deployment_href", "name"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&metadata.Action{
				Name: "launch",
				Description: `Launches a new instance in the server array with the configuration defined in the 'next_instance'. This function is equivalent to invoking the launch action on the
URL of this server_array's next_instance. See Instances#launch for details.
Optional parameters:
	api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
	count: For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error
	inputs`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/launch",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/launch$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of that input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
					&metadata.ActionParam{
						Name:        "count",
						Description: `For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
					&metadata.ActionParam{
						Name:        "count",
						Description: `For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "multi_run_executable",
				Description: `Run an executable on all instances of this array. This function is equivalent to invoking the "multi_run_executable" action on the instances resource
(Instances#multi_run_executable with the filter "parent_href == /api/server_arrays/XX"). To run an executable on a subset of the instances of the array, provide additional filters. To run an executable
a single instance, invoke the action "run_executable" directly on the instance (see Instances#run_executable)
Optional parameters:
	filter
	ignore_lock: Specifies the ability to ignore the lock(s) on the Instance(s).
	inputs
	recipe_name: The name of the recipe to be run.
	right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/multi_run_executable",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_run_executable$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of these inputs. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The name of inputs needed. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to be run.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock(s) on the Instance(s).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock(s) on the Instance(s).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to be run.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "multi_terminate",
				Description: `Terminate all instances of this array. This function is equivalent to invoking the "multi_terminate" action on the instances resource ( Instances#multi_terminate with
the filter "parent_href == /api/server_arrays/XX"). To terminate a subset of the instances of the array, provide additional filters. To terminate a single instance,
invoke the action "terminate" directly on the instance (see Instances#terminate)
Optional parameters:
	filter
	terminate_all: Specifies the ability to terminate all instances.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_arrays/%s/multi_terminate",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_terminate$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "terminate_all",
						Description: `Specifies the ability to terminate all instances.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "terminate_all",
						Description: `Specifies the ability to terminate all instances.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Shows the information of a single server array.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/deployments/%s/server_arrays/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a single server array.
Required parameters:
	server_array`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/server_arrays/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/deployments/%s/server_arrays/%s",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]",
						Description: `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][collect_audit_entries]",
						Description: `The updated audit SQS queue that will store audit entries.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][voters_tag_predicate]",
						Description: `The updated Voters Tag that RightScale will use in order to determine when to scale up/down.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][algorithm]",
						Description: `The updated algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"max_10", "avg_10"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][decision_threshold]",
						Description: `The updated percentage of servers that must agree in order to trigger an alert before an action is taken.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][max_age]",
						Description: `The updated threshold (in seconds) before a resize action occurs on the server array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][regexp]",
						Description: `The updated regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_calm_time]",
						Description: `The updated time (in minutes) on how long you want to wait before you repeat another action.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_down_by]",
						Description: `The updated number of servers to scale down by.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][max_count]",
						Description: `The updated maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][min_count]",
						Description: `The updated minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_up_by]",
						Description: `The updated number of servers to scale up by.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][min_count]",
						Description: `The updated minimum number of servers that must be operational at all times in the server array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][max_count]",
						Description: `The updated maximum number of servers that can be operational at the same time in the server array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][time]",
						Description: `The updated time when an alert-based array resizes.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][day]",
						Description: `The updated day when an alert-based array resizes.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][weight]",
						Description: `Instance allocation (should total 100%).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][max]",
						Description: `Max instances (0 for unlimited).`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[deployment_href]",
						Description: `The updated href of the deployment for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[description]",
						Description: `The updated description for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[array_type]",
						Description: `The updated array type for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"alert", "queue"},
					},
					&metadata.ActionParam{
						Name:        "server_array[optimized]",
						Description: `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[state]",
						Description: `The updated status of the server array. If active, the server array is enabled for scaling actions.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"enabled", "disabled"},
					},
					&metadata.ActionParam{
						Name:        "server_array[name]",
						Description: `The updated name for the Server Array.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_array",
						Description: ``,
						Type:        "*ServerArrayParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"alert_specs":       "Associated AlertSpecs",
			"alerts":            "Associated Alerts",
			"current_instances": "Associated current instances",
			"deployment":        "Associated deployment",
			"next_instance":     "Associated next instance",
			"self":              "Href of itself",
		},
	},
	"ServerTemplate": &metadata.Resource{
		Name: "ServerTemplate",
		Description: `ServerTemplates allow you to pre-configure servers by starting from a base image and adding scripts that run during the boot,
operational, and shutdown phases. A ServerTemplate is a description of how a new instance will be configured when it is
provisioned by your cloud provider.
All revisions of a ServerTemplate belong to a ServerTemplate lineage that is exposed by the "lineage" attribute.
(NOTE: This attribute is merely a string to locate all revisions of a ServerTemplate and NOT a working URL)`,
		Identifier: "application/vnd.rightscale.server_template",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "clone",
				Description: `Clones a given ServerTemplate.
Required parameters:
	server_template`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/clone$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template[description]",
						Description: `The description for the cloned ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_template[name]",
						Description: `The name for the cloned ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template",
						Description: ``,
						Type:        "*ServerTemplateParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "commit",
				Description: `Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.
Required parameters:
	commit_head_dependencies: Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.
	commit_message: The message associated with the commit.
	freeze_repositories: Freeze the repositories.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/commit",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/commit$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "commit_head_dependencies",
						Description: `Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "freeze_repositories",
						Description: `Freeze the repositories.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "commit_head_dependencies",
						Description: `Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "freeze_repositories",
						Description: `Freeze the repositories.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a new ServerTemplate with the given parameters.
Required parameters:
	server_template`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_templates$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template[description]",
						Description: `The description of the ServerTemplate to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_template[name]",
						Description: `The name of the ServerTemplate to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template",
						Description: ``,
						Type:        "*ServerTemplateParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_templates/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "detect_changes_in_head",
				Description: `Identifies RightScripts attached to the resource that differ from their HEAD.
If the attached revision of the RightScript is the HEAD, then this will indicate
a difference between it and the latest committed revision in the same lineage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/detect_changes_in_head",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/detect_changes_in_head$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.
The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
details please see Inputs#index.)
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_templates$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "lineage", "multi_cloud_image_href", "name", "revision"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"description", "lineage", "multi_cloud_image_href", "name", "revision"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&metadata.Action{
				Name: "publish",
				Description: `Publishes a given ServerTemplate and its subordinates.
Only non-HEAD revisions that are owned by the account can be published.
Required parameters:
	account_group_hrefs: List of hrefs of account groups to publish to.
	descriptions
Optional parameters:
	allow_comments: Allow users to leave comments on this ServerTemplate.
	categories: List of Categories.
	email_comments: Email me when a user comments on this ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/publish",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/publish$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_group_hrefs[]",
						Description: `List of hrefs of account groups to publish to.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "descriptions[notes]",
						Description: `New Revision Notes.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "descriptions[short]",
						Description: `Short Description.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "descriptions[long]",
						Description: `Long Description.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "allow_comments",
						Description: `Allow users to leave comments on this ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "email_comments",
						Description: `Email me when a user comments on this ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "categories[]",
						Description: `List of Categories.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_group_hrefs",
						Description: `List of hrefs of account groups to publish to.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "allow_comments",
						Description: `Allow users to leave comments on this ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "categories",
						Description: `List of Categories.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "descriptions",
						Description: ``,
						Type:        "*Descriptions",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "email_comments",
						Description: `Email me when a user comments on this ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "resolve",
				Description: `Enumerates all attached cookbooks, missing dependencies and bound executables.
Version constraints on missing dependencies and the state of the Chef Recipes;
whether or not the cookbook or recipe itself could be found among the
attachments, will also be reported.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/resolve",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/resolve$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single ServerTemplate. HEAD revisions have a revision of 0.
The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
details please see Inputs#index.)
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_templates/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
			},

			&metadata.Action{
				Name: "swap_repository",
				Description: `In-place replacement of attached cookbooks from a given repository.
For each attached cookbook coming from the source repository, replace it by
attaching a cookbook of identical name coming from the target repository.
In order for the operation to be successful, all attachments that came from the
source repository must exist in the target repository.
If multiple cookbooks of a given name exist in the target repository, preference
is given by the following order (top most being the highest preference):
  Name & Version Match / Primary Namespace
  Name & Version Match / Alternate Namespace
  Name Match / Primary Namespace
  Name Match / Alternate Namespace
If multiple cookbooks still have the same preference for the replacement, the operation is
indeterministic.
Required parameters:
	source_repository_href: The repository whose cookbook attachments are to be replaced.
	target_repository_href: The repository whose cookbook attachments are to be utilized.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_templates/%s/swap_repository",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/swap_repository$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source_repository_href",
						Description: `The repository whose cookbook attachments are to be replaced.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "target_repository_href",
						Description: `The repository whose cookbook attachments are to be utilized.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source_repository_href",
						Description: `The repository whose cookbook attachments are to be replaced.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "target_repository_href",
						Description: `The repository whose cookbook attachments are to be utilized.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).
Currently, the attributes you can update are only the 'direct' attributes of a server template. To
manage multi cloud images of a ServerTemplate, please see the resource 'ServerTemplateMultiCloudImages'.
Required parameters:
	server_template`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/server_templates/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template[description]",
						Description: `The updated description for the ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_template[name]",
						Description: `The updated name for the ServerTemplate.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template",
						Description: ``,
						Type:        "*ServerTemplateParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"alert_specs":               "Associated AlertSpecs",
			"cookbook_attachments":      "Attached Chef Cookbooks",
			"default_multi_cloud_image": "The default MultiCloud Image",
			"inputs":                    "List of configuration inputs",
			"multi_cloud_images":        "Associated MultiCloud Images",
			"publication":               "Associated privately shared Publication",
			"runnable_bindings":         "Attached RightScripts and Chef Recipes",
			"self":                      "Href of itself",
		},
	},
	"ServerTemplateMultiCloudImage": &metadata.Resource{
		Name: "ServerTemplateMultiCloudImage",
		Description: `This resource represents links between ServerTemplates and MultiCloud Images and enables you to effectively
add/delete MultiCloudImages to ServerTemplates and make them the default one.`,
		Identifier: "application/vnd.rightscale.server_template_multi_cloud_image",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new ServerTemplateMultiCloudImage with the given parameters.
Required parameters:
	server_template_multi_cloud_image`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_template_multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template_multi_cloud_image[multi_cloud_image_href]",
						Description: `The href of the MultiCloud Image to be used.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_template_multi_cloud_image[server_template_href]",
						Description: `The href of the ServerTemplate to be used.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template_multi_cloud_image",
						Description: ``,
						Type:        "*ServerTemplateMultiCloudImageParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given ServerTemplateMultiCloudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/server_template_multi_cloud_images/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists the ServerTemplateMultiCloudImages available to this account.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_template_multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"is_default", "multi_cloud_image_href", "server_template_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"is_default", "multi_cloud_image_href", "server_template_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "make_default",
				Description: `Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/server_template_multi_cloud_images/%s/make_default",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images/([^/]+)/make_default$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_template_multi_cloud_images/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"multi_cloud_image": "Associated MultiCloud Image",
			"self":              "Href of itself",
			"server_template":   "Associated ServerTemplate",
		},
	},
	"Session": &metadata.Resource{
		Name: "Session",
		Description: `The sessions resource is in charge of creating API sessions that are bound to a given account. The sequence for login into the API is the following:
* Perform a POST request to /api/sessions ('create' action) to my.rightscale.com or to any more specific hosts saved from previous sessions.
* If the targeted host is not appropriate for the specific account being accessed it will return a 302 http code with a URL with which the client must retry the same POST request.
* If the targeted host is the right one and the login is successful, it will return a 204 http code, along with two cookies that will need to be saved and passed in any subsequent API request.
* If there is an authentication or authorization problem with the POST request an error (typically 401 or 422 ) may be returned at any point in the above sequence.
* If the session expires, it will return a 403 http code with a "Session cookie is expired or invalid" message.
Note that all API calls irrespective of the resource it is acting on, should pass a header "X_API_VERSION" with the value "1.5".`,
		Identifier: "application/vnd.rightscale.session",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "accounts",
				Description: `List all the accounts that a user has access to.
This call may be executed outside of an existing session. Doing so requires passing a username and password in the
request body. The idea is that it should be possible to list accounts that can be used to create a session.
Upon successfully authenticating the credentials, the system will return a 200 OK code and return the list of accounts.
If an 302 redirect code is returned, the client is responsible of re-issuing the GET request against the content of the received Location header, passing the exact same parameters again.
Example Request using Curl (not using an existing session):
curl -i -H X_API_VERSION:1.5 -X GET -d email='email@me.com' -d password='mypassword' https://my.rightscale.com/api/sessions/accounts
Example Request using Curl (using an existing session):
curl -i -H X_API_VERSION:1.5 -X GET -b mycookies https://my.rightscale.com/api/sessions/accounts
Optional parameters:
	email: The email to login with if not using existing session.
	password: The corresponding password.
	view: Extended view shows account permissions and products`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/sessions/accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/sessions/accounts$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "password",
						Description: `The corresponding password.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "email",
						Description: `The email to login with if not using existing session.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Extended view shows account permissions and products`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "email",
						Description: `The email to login with if not using existing session.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "password",
						Description: `The corresponding password.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Extended view shows account permissions and products`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what
features are available within its privileges.
Example Request using Curl:
curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions
Optional parameters:
	view: Whoami view provides links to the logged-in principal and the account being accessed`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/sessions",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/sessions$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `Whoami view provides links to the logged-in principal and the account being accessed`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "whoami"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `Whoami view provides links to the logged-in principal and the account being accessed`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "whoami"},
					},
				},
			},

			&metadata.Action{
				Name: "index_instance_session",
				Description: `Shows the full attributes of the instance (that has the token used to log-in).
This call can be used by an instance to get it's own details.
Example Request using Curl:
curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions/instance`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/sessions/instance",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/sessions/instance$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
		Links: map[string]string{
			"account":                            "The account currently being accessed",
			"account_groups":                     "Available AccountGroups",
			"accounts":                           "Available Accounts",
			"alert_specs":                        "Available AlertSpecs",
			"alerts":                             "Available Alerts",
			"audit_entries":                      "Available AuditEntries",
			"backups":                            "Available Backups",
			"child_accounts":                     "Available ChildAccounts",
			"cloud_accounts":                     "Available CloudAccounts",
			"cloud_flow_processes":               "Available Cloud Workflow processes",
			"clouds":                             "Available Clouds",
			"cookbooks":                          "Available Cookbooks",
			"credentials":                        "Available Credentials",
			"deployments":                        "Available Deployments",
			"identity_providers":                 "Available IdentityProviders",
			"instance":                           "The currently logged-in instance",
			"multi_cloud_images":                 "Available MultiCloudImages",
			"network_gateways":                   "Available NetworkGateways",
			"network_option_group_attachments":   "Available NetworkOptionGroupAttachments",
			"network_option_groups":              "Available NetworkOptionGroups",
			"networks":                           "Available Networks",
			"permissions":                        "Available Permissions",
			"placement_groups":                   "Available PlacementGroups",
			"preferences":                        "Available Preferences",
			"publication_lineages":               "Available PublicationLineages",
			"publications":                       "Available Publications",
			"repositories":                       "Available Repositories",
			"right_scripts":                      "Available RightScripts",
			"route_tables":                       "Available RouteTables",
			"routes":                             "Available Routes",
			"security_group_rules":               "Available SecurityGroupRules",
			"self":                               "Href of itself",
			"server_arrays":                      "Available ServerArrays",
			"server_template_multi_cloud_images": "Available ServerTemplateMultiCloudImages",
			"server_templates":                   "Available ServerTemplates",
			"servers":                            "Available Servers",
			"tags":                               "Search for Tags",
			"user":                               "The currently logged-in user",
			"users":                              "Available Users",
		},
	},
	"SshKey": &metadata.Resource{
		Name: "SshKey",
		Description: `Ssh Keys represent a created SSH Key that exists in the cloud.
An ssh key might also contain the private part of the key, and can be used to login to instances launched with it.`,
		Identifier: "application/vnd.rightscale.ssh_key",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new ssh key.
Required parameters:
	ssh_key`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/ssh_keys",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ssh_key[name]",
						Description: `The name for the SSH key to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ssh_key",
						Description: ``,
						Type:        "*SshKeyParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given ssh key.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/ssh_keys/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists ssh keys.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ssh_keys",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single ssh key.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/ssh_keys/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud": "Associated cloud",
			"self":  "Href of itself",
		},
	},
	"Subnet": &metadata.Resource{
		Name: "Subnet",
		Description: `A Subnet is a logical grouping of network devices. An Instance can have many
Subnets.`,
		Identifier: "application/vnd.rightscale.subnet",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new subnet.
Required parameters:
	subnet`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/subnets",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/subnets",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "subnet[datacenter_href]",
						Description: `The associated Datacenter.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[network_href]",
						Description: `The associated Network.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[description]",
						Description: `The description for the Subnet.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[cidr_block]",
						Description: `The range of IP addresses for the Subnet.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[name]",
						Description: `The name for the Subnet.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "subnet",
						Description: ``,
						Type:        "*SubnetParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the given subnet(s).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/instances/%s/subnets/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/subnets/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists subnets of a given cloud.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/subnets",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/subnets",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"datacenter_href", "instance_href", "name", "network_href", "resource_uid", "visibility"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"datacenter_href", "instance_href", "name", "network_href", "resource_uid", "visibility"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows attributes of a single subnet.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/subnets/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/subnets/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Updates name and description for the current subnet.
Required parameters:
	subnet`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/instances/%s/subnets/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/subnets/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "subnet[route_table_href]",
						Description: `The RouteTable to associate/dissociate with this Subnet. Pass this parameter with an empty string to reset this Subnet to use the default RouteTable.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subnet[description]",
						Description: `The updated description for the Subnet.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subnet[cidr_block]",
						Description: `The range of IP addresses for the Subnet.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[name]",
						Description: `The updated name for the Subnet.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "subnet",
						Description: ``,
						Type:        "*SubnetParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"datacenter":            "Href of the Datacenter the Subnet is in",
			"effective_route_table": "Href of the RouteTable associated with the Subnet (may be inherited).",
			"network":               "Href of the Network the Subnet is in",
			"route_table":           "Href of the RouteTable explicitly associated with the Subnet, if any.",
			"self":                  "Href of itself",
		},
	},
	"Tag": &metadata.Resource{
		Name: "Tag",
		Description: `A tag or machine tag is a useful way of attaching useful metadata to an object/resource.
Tags are commonly used as an extra label or identifier.
For example, you might want to add a tag to an EBS Snapshot or AMI so that you can find it more quickly.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "by_resource",
				Description: `Get tags for a list of resource hrefs.
The hrefs can belong to various resource types and the tags for a non-existent href will be empty.
Required parameters:
	resource_hrefs: Hrefs of the resources for which tags are to be returned.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/tags/by_resource",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/by_resource$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which tags are to be returned.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs",
						Description: `Hrefs of the resources for which tags are to be returned.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "by_tag",
				Description: `Search for resources having a list of tags in a specific resource_type.
The search criteria can contain plain tags ("my_db_server"), machine tags ("server:db=true"), or
namespace &amp; predicate wildcards ("server:db=*"). The result set includes links to the resources.
If match_all is "true", then the search is an "AND" operation -- only resources having all the
tags are returned. If match_all has any other value or is missing, the search is performed
as an "OR" operation.
PLEASE NOTE the behavior of the include_tags_with_prefix parameter: if it is absent,
or blank, then you will find resources that match your query but receive no information about
the tags that apply to those resources. (Your search will also complete much more quickly in
this case.)
For example, a search with tag[]="server:db=true" and include_tags_with_prefix="backup:"
will return resources that are tagged as a DB server, and also return all "backup" related tags
for every matching resource.
Required parameters:
	resource_type: Search among a single resource type.
	tags: The tags which must be present on the resource.
Optional parameters:
	include_tags_with_prefix: If included, all tags matching this prefix will be returned. If not included, no tags will be returned.
	match_all: If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.
	with_deleted: If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/tags/by_tag",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/by_tag$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "include_tags_with_prefix",
						Description: `If included, all tags matching this prefix will be returned. If not included, no tags will be returned.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "resource_type",
						Description: `Search among a single resource type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"servers", "instances", "volumes", "volume_snapshots", "deployments", "server_templates", "multi_cloud_images", "images", "server_arrays", "accounts", "placement_groups"},
					},
					&metadata.ActionParam{
						Name:        "with_deleted",
						Description: `If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "match_all",
						Description: `If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "tags[]",
						Description: `The tags which must be present on the resource.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "include_tags_with_prefix",
						Description: `If included, all tags matching this prefix will be returned. If not included, no tags will be returned.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "match_all",
						Description: `If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "resource_type",
						Description: `Search among a single resource type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"servers", "instances", "volumes", "volume_snapshots", "deployments", "server_templates", "multi_cloud_images", "images", "server_arrays", "accounts", "placement_groups"},
					},
					&metadata.ActionParam{
						Name:        "tags",
						Description: `The tags which must be present on the resource.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "with_deleted",
						Description: `If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "multi_add",
				Description: `Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.
The hrefs can belong to various resource types. If a resource for a href could not be found, an
error is returned and no tags are added for any resource.
No error will be raised if the resource already has the tag(s) you are trying to add.
Note: At this point, tags on 'next_instance' are not supported and one has to add tags to the 'server'.
Required parameters:
	resource_hrefs: Hrefs of the resources for which the tags are to be added.
	tags: Tags to be added.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/tags/multi_add",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/multi_add$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which the tags are to be added.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tags[]",
						Description: `Tags to be added.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs",
						Description: `Hrefs of the resources for which the tags are to be added.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tags",
						Description: `Tags to be added.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "multi_delete",
				Description: `Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.
The hrefs can belong to various resource types. If a resource for a href could not be found, an
error is returned and no tags are deleted for any resource.
Note that no error will be raised if the resource does not have the tag(s) you are trying to delete.
Required parameters:
	resource_hrefs: Hrefs of the resources for which tags are to be deleted.
	tags: Tags to be deleted.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/tags/multi_delete",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/multi_delete$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which tags are to be deleted.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tags[]",
						Description: `Tags to be deleted.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs",
						Description: `Hrefs of the resources for which tags are to be deleted.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tags",
						Description: `Tags to be deleted.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
	},
	"Task": &metadata.Resource{
		Name: "Task",
		Description: `Tasks represent processes that happen (or have happened) asynchronously within the context of an Instance.
An example of a type of task is an operational script that runs in an instance.
Task resources can be returned by certain API calls, such as Instances.run_executable, Backups.restore, and others.`,
		Identifier: "application/vnd.rightscale.task",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "show",
				Description: `Displays information of a given task within the context of an instance.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/live/tasks/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/live/tasks/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/server_arrays/%s/live/tasks/%s",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/live/tasks/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},
		},
		Links: map[string]string{
			"self": "Href of itself",
		},
	},
	"User": &metadata.Resource{
		Name:        "User",
		Description: ``,
		Identifier:  "application/vnd.rightscale.user",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Create a user. If a user already exists with the same email, that user will be returned.
Creating a user alone will not enable the user to access this account. You have to create
'permissions' for that user before it can be used. Performing a 'show' on a new user
will fail unless you immediately create an 'observer' permission on the current account.
Note that information about users and their permissions must be propagated globally across all
RightScale clusters, and this can take some time (less than 60 seconds under normal circumstances)
so the users you create may not be able to login for a minute or two after you create them. However,
you may create or destroy permissions for newly-created users with no delay.
To create a user that will login using password authentication, include the 'password' parameter
with your request.
To create an SSO-enabled user, you must specify the identity_provider that will be vouching for
this user's identity, as well as the principal_uid (SAML NameID or OpenID identity URL) that
the identity provider will assert for this user. Identity providers should be specified by
their API href; you can obtain a list of the identity providers available to your account by
invoking the 'index' action of the identity_providers API resource.
Required parameters:
	user`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/users",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/users$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user[identity_provider_href]",
						Description: `The RightScale API href of the Identity Provider through which this user will login to RightScale. Required to create an SSO-authenticated user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[principal_uid]",
						Description: `The principal identifier (SAML NameID or OpenID identity URL) of this user. Required to create an SSO-authenticated user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[timezone_name]",
						Description: `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[first_name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[last_name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[password]",
						Description: `The password of this user. Required to create a password-authenticated user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[company]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[phone]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[email]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user",
						Description: ``,
						Type:        "*UserParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `List the users available to the account the user is logged in to. Therefore, to list the users of
a child account, the user has to login to the child account first.
Optional parameters:
	filter`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/users",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/users$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"email", "first_name", "last_name"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"email", "first_name", "last_name"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single user.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/users/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/users/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update a user's contact information, change their password, or update their SSO settings.
In order to update a user record, one of the following criteria must be met:
1. You've authenticated and are the user being modified, and you provide a valid current_password.
2. You're an admin and the user is linked to your enterprise SSO provider.
3. You're an admin and the user's email matches the email_domain of your enterprise SSO provider.
In other words: you can update yourself if you know your own password, you can update
yourself or others if you're an admin and they're linked to your SSO provider, and you can update any user
if you're an admin and their email address is known to belong to your organization.
For information about enabling canonical email domain ownership for your enterprise, please
talk to your RightScale account manager or contact our support team.
To update a user's contact information, simply pass the desired values for email, first_name,
and so forth.
To update a user's password, provide the desired new_password.
To set or update a user's SSO information, you may provide a just a principal_uid (to maintain the
user's existing identity provider) or you may provide an identity_provider_href and a
principal_uid (to switch identity providers as well as specify a new user identity).
In the context of SAML, principal_uid is equivalent to the SAML NameID or Subject claim.
RightScale cannot predict or influence the NameID value that your SAML IdP will send to us for
Required parameters:
	user`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/users/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/users/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user[identity_provider_href]",
						Description: `The updated RightScale API href of the associated Identity Provider.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[current_password]",
						Description: `The current password for the user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[timezone_name]",
						Description: `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[principal_uid]",
						Description: `The updated principal identifier (SAML NameID or OpenID identity URL) of this user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[current_email]",
						Description: `The existing email of this user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[new_password]",
						Description: `The new password for this user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[first_name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[login_name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "user[last_name]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[new_email]",
						Description: `The updated email of this user.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[company]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[phone]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user",
						Description: ``,
						Type:        "*UserParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"identity_provider": "Associated identity provider",
			"self":              "Href of itself",
		},
	},
	"UserData": &metadata.Resource{
		Name:        "UserData",
		Description: ``,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `No description provided for show.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/user_data/",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/user_data/$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Volume": &metadata.Resource{
		Name:        "Volume",
		Description: `A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).`,
		Identifier:  "application/vnd.rightscale.volume",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new volume.
Required parameters:
	volume`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volumes",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume[parent_volume_snapshot_href]",
						Description: `The href of the snapshot from which the volume will be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[placement_group_href]",
						Description: `The href of the Placement Group. This option can not be used in combination with parent_volume_snapshot_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[volume_type_href]",
						Description: `The href of the volume type. A Name, Resource UID and optional Size is associated with a Volume Type.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[datacenter_href]",
						Description: `The href of the Datacenter / Zone that the Volume will be in. This parameter is required for non-OpenStack clouds.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[deployment_href]",
						Description: `The href of the Deployment that owns this Volume.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[description]",
						Description: `The description of the Volume to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[image_href]",
						Description: `The href of the Image that should be used as a source`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[encrypted]",
						Description: `A flag indicating whether Volume should be encrypted. Only available on clouds supporting volume encryption.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "volume[iops]",
						Description: `The number of IOPS (I/O Operations Per Second) this Volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[name]",
						Description: `The name for the Volume to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[size]",
						Description: `The size of a Volume to be created in gigabytes (GB). Some Volume Types have predefined sizes and do not allow selecting a custom size on Volume creation.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume",
						Description: ``,
						Type:        "*VolumeParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given volume.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volumes/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists volumes.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"datacenter_href", "deployment_href", "description", "name", "parent_volume_snapshot_href", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"datacenter_href", "deployment_href", "description", "name", "parent_volume_snapshot_href", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single volume.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `No description provided for update.
Required parameters:
	volume`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/clouds/%s/volumes/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume[allowed_instance_hrefs][remove][]",
						Description: `Hrefs for instances to remove from allowed list.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[allowed_instance_hrefs][add][]",
						Description: `Hrefs for instances to add into allowed list.`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[name]",
						Description: `The new name for the Volume.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume",
						Description: ``,
						Type:        "*VolumeParam2",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":                        "Associated cloud",
			"current_volume_attachment":    "Associated volume attachment. Describes where the volume is attached to and the attachment parameters.",
			"datacenter":                   "Associated datacenter/Zone",
			"deployment":                   "Containing Deployment",
			"image":                        "Associated image",
			"parent_volume_snapshot":       "The volume snapshot from which the volume was created, if any",
			"placement_group":              "Associated placement group",
			"recurring_volume_attachments": "Associated recurring volume attachments",
			"self":             "Href of itself",
			"volume_snapshots": "Associated volume snapshots",
			"volume_type":      "Associated volume type",
		},
	},
	"VolumeAttachment": &metadata.Resource{
		Name:        "VolumeAttachment",
		Description: `A VolumeAttachment represents a relationship between a volume and an instance.`,
		Identifier:  "application/vnd.rightscale.volume_attachment",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "create",
				Description: `Creates a new volume attachment.
Required parameters:
	volume_attachment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/instances/%s/volume_attachments",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_attachment",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachment$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_attachment[settings][delete_on_termination]",
						Description: `Setting to 'true' will schedule volume deletion if instance was terminated, default value is 'false'`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_attachment[instance_href]",
						Description: `The href of the instance to which the volume will be attached. Mutually exclusive with server_href.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_attachment[volume_href]",
						Description: `The href of the volume to be attached.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_attachment[server_href]",
						Description: `The href of the server to which the volume will be attached. Mutually exclusive with instance_href.Note: the Server must have a current_instance.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_attachment[device]",
						Description: `The device location where the volume will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_attachment",
						Description: ``,
						Type:        "*VolumeAttachmentParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "destroy",
				Description: `Deletes a given volume attachment.
Optional parameters:
	force: Specifies whether to force the detachment of a volume.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/instances/%s/volume_attachments/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volume_attachments/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_attachment",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachment$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "force",
						Description: `Specifies whether to force the detachment of a volume.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "force",
						Description: `Specifies whether to force the detachment of a volume.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all volume attachments.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/volume_attachments",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"instance_href", "resource_uid", "volume_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"instance_href", "resource_uid", "volume_href"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single volume attachment.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/volume_attachments/%s",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_attachments/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_attachment",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachment$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":    "Associated cloud",
			"instance": "Associated instance",
			"self":     "Href of itself",
			"volume":   "Associated volume",
		},
	},
	"VolumeSnapshot": &metadata.Resource{
		Name: "VolumeSnapshot",
		Description: `A VolumeSnapshot represents a Cloud storage volume at a particular point in time. One can create a snapshot regardless of whether or not a volume is attached to an Instance. When a snapshot is created,
various meta data is retained such as a Created At timestamp, a unique Resource UID (e.g. vol-52EF05A9), the Volume Owner and Visibility (e.g. private or public).
Snapshots consist of a series of data blocks that are incrementally saved.`,
		Identifier: "application/vnd.rightscale.volume_snapshot",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "copy",
				Description: `No description provided for copy.
Required parameters:
	volume_snapshot_copy`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_snapshots/%s/copy",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots/([^/]+)/copy$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volume_snapshots/%s/copy",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/copy$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_snapshot_copy[description]",
						Description: `The description of the Volume Snapshot to be copied.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot_copy[cloud_href]",
						Description: `The href of the destination cloud.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot_copy[name]",
						Description: `The name of the Volume Snapshot to be copied.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_snapshot_copy",
						Description: ``,
						Type:        "*VolumeSnapshotCopy",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Creates a new volume_snapshot.
Required parameters:
	volume_snapshot`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_snapshots",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/clouds/%s/volume_snapshots",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_snapshot[parent_volume_href]",
						Description: `The href of the Volume from which the Volume Snapshot will be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot[deployment_href]",
						Description: `The href of the Deployment that owns this Volume Snapshot.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot[description]",
						Description: `The description for the Volume Snapshot to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot[name]",
						Description: `The name for the Volume Snapshot to be created.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_snapshot",
						Description: ``,
						Type:        "*VolumeSnapshotParam",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    true,
					},
				},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given volume_snapshot.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_snapshots/%s",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/clouds/%s/volume_snapshots/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index",
				Description: `Lists all volume_snapshots.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_snapshots",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_snapshots",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"deployment_href", "description", "name", "parent_volume_href", "resource_uid", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"deployment_href", "description", "name", "parent_volume_href", "resource_uid", "state"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single volume_snapshot.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volumes/%s/volume_snapshots/%s",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots/([^/]+)$`),
					},
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_snapshots/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud":                        "Associated cloud",
			"deployment":                   "Containing Deployment",
			"parent_volume":                "The volume from which the snapshot was created.",
			"recurring_volume_attachments": "Associated recurring volume attachments",
			"self": "Href of itself",
		},
	},
	"VolumeType": &metadata.Resource{
		Name:        "VolumeType",
		Description: `A VolumeType describes the type of volume, particularly the size.`,
		Identifier:  "application/vnd.rightscale.volume_type",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Lists Volume Types.
Optional parameters:
	filter
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_types",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_types$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"name", "resource_uid"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Displays information about a single Volume Type.
Optional parameters:
	view`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/clouds/%s/volume_types/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_types/([^/]+)$`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
			},
		},
		Links: map[string]string{
			"cloud": "Associated cloud",
			"self":  "Href of itself",
		},
	},
}
