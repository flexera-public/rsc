//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ goav2gen -metadata=policy/docs -output=policy -pkg=policy -version=1.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package policy

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"AppliedPolicy": &metadata.Resource{
		Name: "AppliedPolicy",
		Description: `Show retrieves the details of an applied policy.
**`,
		Identifier: "application/vnd.rightscale.applied_policy",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Index retrieves the list of applied policies in a project.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/applied_policies",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/applied_policies`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render applied policies`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name is a filter to only show applied policies that match this name.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render applied policies`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name is a filter to only show applied policies that match this name.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Create applies a policy template to a given project. The applied policy will continually run until deleted.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/projects/%s/applied_policies",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/applied_policies`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credentials",
						Description: `credentials is the map of name and credential used to launch the policy.`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `description provides a human readable description for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "dry_run",
						Description: `dry_run is a flag used for testing a policy so that an incident can be raised without performing an action.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "frequency",
						Description: `frequency specifies the frequency with which to run policy evaluations`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"15 minutes", "hourly", "daily", "weekly", "monthly"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name provides a name for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][name]",
						Description: `name of option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][value]",
						Description: `value of option`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "severity",
						Description: `severity is the severity level of the incident.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"low", "medium", "high", "critical"},
					},
					&metadata.ActionParam{
						Name:        "skip_approvals",
						Description: `skip_approvals means that any approval actions will be skipped and all actions automatically run.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the policy template or published template that is applied.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/(projects/[0-9]+/policy|orgs/[0-9]+/published)_templates/[0-9a-f]+$"),
					},
				},
				Payload: "AppliedPolicyCreate",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credentials",
						Description: `credentials is the map of name and credential used to launch the policy.`,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `description provides a human readable description for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "dry_run",
						Description: `dry_run is a flag used for testing a policy so that an incident can be raised without performing an action.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "frequency",
						Description: `frequency specifies the frequency with which to run policy evaluations`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"15 minutes", "hourly", "daily", "weekly", "monthly"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name provides a name for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options",
						Description: `options lists the configuration options used to parameterize the policy.`,
						Type:        "[]*ConfigurationOptionCreateType",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "severity",
						Description: `severity is the severity level of the incident.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"low", "medium", "high", "critical"},
					},
					&metadata.ActionParam{
						Name:        "skip_approvals",
						Description: `skip_approvals means that any approval actions will be skipped and all actions automatically run.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the policy template or published template that is applied.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/(projects/[0-9]+/policy|orgs/[0-9]+/published)_templates/[0-9a-f]+$"),
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of an applied policy.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/applied_policies/%s",
						Variables:  []string{"project_id", "policy_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/applied_policies/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render applied policy`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render applied policy`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "source"},
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Delete stops and deletes an applied policy.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/governance/projects/%s/applied_policies/%s",
						Variables:  []string{"project_id", "policy_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/applied_policies/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "evaluate",
				Description: `Evaluate executes an applied policy evaluation on demand. It does not affect the normal execution schedule.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/projects/%s/applied_policies/%s/evaluate",
						Variables:  []string{"project_id", "policy_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/applied_policies/([^/]+)/evaluate`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "show_log",
				Description: `ShowLog retrieves the last evaluation log of an applied policy. *The content type is "text/markdown"*.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/applied_policies/%s/log",
						Variables:  []string{"project_id", "policy_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/applied_policies/([^/]+)/log`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "show_status",
				Description: `ShowStatus retrieves the evaluation status details of an applied policy.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/applied_policies/%s/status",
						Variables:  []string{"project_id", "policy_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/applied_policies/([^/]+)/status`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Approval": &metadata.Resource{
		Name:        "Approval",
		Description: ``,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Index retrieves the list of approval requests in a project.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/approval_requests",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/approval_requests`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render approval requests.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Optional resource ID to filter. Multiple may be specified.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subject_kind",
						Description: `Subject kind to filter on.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subject_href",
						Description: `Subject HREF to filter on.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "status",
						Description: `Status of the approval to filter on. Multiple may be specified.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render approval requests.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
					&metadata.ActionParam{
						Name:        "id",
						Description: `Optional resource ID to filter. Multiple may be specified.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subject_kind",
						Description: `Subject kind to filter on.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subject_href",
						Description: `Subject HREF to filter on.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "status",
						Description: `Status of the approval to filter on. Multiple may be specified.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of an approval request.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/approval_requests/%s",
						Variables:  []string{"project_id", "approval_request_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/approval_requests/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render approval request.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render approval request.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "approve",
				Description: `Approve approves a single approval request.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/projects/%s/approval_requests/%s/approve",
						Variables:  []string{"project_id", "approval_request_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/approval_requests/([^/]+)/approve`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "options[][name]",
						Description: `name of option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][value]",
						Description: `value of option`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				Payload: "ApprovalApprove",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "options",
						Description: `options lists the configuration options used to parameterize the approval.`,
						Type:        "[]*ConfigurationOptionCreateType",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "deny",
				Description: `Deny denies a single approval request.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/projects/%s/approval_requests/%s/deny",
						Variables:  []string{"project_id", "approval_request_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/approval_requests/([^/]+)/deny`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "comment",
						Description: `A comment that explains the reason for denial`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				Payload: "ApprovalDeny",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "comment",
						Description: `A comment that explains the reason for denial`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"ArchivedIncident": &metadata.Resource{
		Name: "ArchivedIncident",
		Description: `Show retrieves the details of an archived incident.
**`,
		Identifier: "application/vnd.rightscale.archived_incident",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Index retrieves the list of archived_incidents in a project.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/archived_incidents",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/archived_incidents`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render archived_incidents`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
					&metadata.ActionParam{
						Name:        "state",
						Description: `state is a filter to only show archived_incidents that are in this state.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "applied_policy_id",
						Description: `applied_policy_id is a filter to only show archived_incidents that were created by a certain applied policy.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render archived_incidents`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
					&metadata.ActionParam{
						Name:        "state",
						Description: `state is a filter to only show archived_incidents that are in this state.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "applied_policy_id",
						Description: `applied_policy_id is a filter to only show archived_incidents that were created by a certain applied policy.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of an archived incident.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/archived_incidents/%s",
						Variables:  []string{"project_id", "incident_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/archived_incidents/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render archived incident`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render archived incident`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
			},

			&metadata.Action{
				Name: "index_escalations",
				Description: `IndexEscalations retrieves the status details of all of the escalations for an archived incident.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/archived_incidents/%s/escalations",
						Variables:  []string{"project_id", "incident_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/archived_incidents/([^/]+)/escalations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index_resolutions",
				Description: `IndexResolutions retrieves the status details of all of the resolutions for an archived incident.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/archived_incidents/%s/resolutions",
						Variables:  []string{"project_id", "incident_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/archived_incidents/([^/]+)/resolutions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Incident": &metadata.Resource{
		Name: "Incident",
		Description: `Show retrieves the details of an incident.
**`,
		Identifier: "application/vnd.rightscale.incident",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Index retrieves the list of incidents in a project.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/incidents",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/incidents`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incidents`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
					&metadata.ActionParam{
						Name:        "state",
						Description: `state is a filter to only show incidents that are in this state.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "applied_policy_id",
						Description: `applied_policy_id is a filter to only show incidents that were created by a certain applied policy.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incidents`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
					&metadata.ActionParam{
						Name:        "state",
						Description: `state is a filter to only show incidents that are in this state.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "applied_policy_id",
						Description: `applied_policy_id is a filter to only show incidents that were created by a certain applied policy.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of an incident.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/incidents/%s",
						Variables:  []string{"project_id", "incident_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/incidents/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
			},

			&metadata.Action{
				Name: "index_escalations",
				Description: `IndexEscalations retrieves the status details of all of the escalations for an incident.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/incidents/%s/escalations",
						Variables:  []string{"project_id", "incident_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/incidents/([^/]+)/escalations`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "index_resolutions",
				Description: `IndexResolutions retrieves the status details of all of the resolutions for an incident.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/incidents/%s/resolutions",
						Variables:  []string{"project_id", "incident_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/incidents/([^/]+)/resolutions`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "resolve",
				Description: `Resolve resolves an incident by setting it to an inactive state, indicating that it has been addressed.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/governance/projects/%s/incidents/%s/resolve",
						Variables:  []string{"project_id", "incident_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/incidents/([^/]+)/resolve`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"IncidentAggregate": &metadata.Resource{
		Name: "IncidentAggregate",
		Description: `Show retrieves the details of an aggregate.
**`,
		Identifier: "application/vnd.rightscale.incident_aggregate",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Index retrieves the list of incident_aggregates in an organization.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/incident_aggregates",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/incident_aggregates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident_aggregates`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident_aggregates`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name: "show_non_catalog",
				Description: `ShowNonCatalog retrieves a list of incidents in the non-catalog policy aggregate. These incidents largely originate from dev/test environments.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/incident_aggregates/non_catalog",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/incident_aggregates/non_catalog`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident aggregate`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "index"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident aggregate`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "index"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of an aggregate.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/incident_aggregates/%s",
						Variables:  []string{"org_id", "incident_aggregate_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/incident_aggregates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident aggregate`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "index", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render incident aggregate`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "index", "source"},
					},
				},
			},
		},
	},
	"PolicyAggregate": &metadata.Resource{
		Name: "PolicyAggregate",
		Description: `Show retrieves the details of a policy aggregate.
**`,
		Identifier: "application/vnd.rightscale.policy_aggregate",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Index retrieves the list of policy aggregates in an org.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/policy_aggregates",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/policy_aggregates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy aggregates.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name is a filter to only show policy aggregates that match this name.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy aggregates.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name is a filter to only show policy aggregates that match this name.`,
						Type:        "interface{}",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Create applies a policy template to a given project. The policy aggregate will continually run until deleted.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/orgs/%s/policy_aggregates",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/policy_aggregates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "all_projects",
						Description: `all_projects is a flag to specify the policy should be run on all projects in the org.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "credentials",
						Description: `credentials is the map of name to credential used to launch the policy.`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `description provides a human readable description for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "dry_run",
						Description: `dry_run is a flag used for testing a policy so that an incident can be raised without performing an action.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "frequency",
						Description: `frequency specifies the frequency with which to run policy evaluations`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"15 minutes", "hourly", "daily", "weekly", "monthly"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name provides a name for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][name]",
						Description: `name of option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][value]",
						Description: `value of option`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_ids[]",
						Description: ``,
						Type:        "int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "severity",
						Description: `severity is the severity level of the incident.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"low", "medium", "high", "critical"},
					},
					&metadata.ActionParam{
						Name:        "skip_approvals",
						Description: `skip_approvals means that any approval actions will be skipped and all actions automatically run.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the published template that is applied.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/orgs/[0-9]+/published_templates/[0-9a-f]+$"),
					},
				},
				Payload: "PolicyAggregateCreate",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "all_projects",
						Description: `all_projects is a flag to specify the policy should be run on all projects in the org.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "credentials",
						Description: `credentials is the map of name to credential used to launch the policy.`,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "description",
						Description: `description provides a human readable description for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "dry_run",
						Description: `dry_run is a flag used for testing a policy so that an incident can be raised without performing an action.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "frequency",
						Description: `frequency specifies the frequency with which to run policy evaluations`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"15 minutes", "hourly", "daily", "weekly", "monthly"},
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `name provides a name for this specific application of the policy.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options",
						Description: `options lists the configuration options used to parameterize the policy.`,
						Type:        "[]*ConfigurationOptionCreateType",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "project_ids",
						Description: `A list of project ids to include in this aggregate.`,
						Type:        "[]int",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "severity",
						Description: `severity is the severity level of the incident.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"low", "medium", "high", "critical"},
					},
					&metadata.ActionParam{
						Name:        "skip_approvals",
						Description: `skip_approvals means that any approval actions will be skipped and all actions automatically run.`,
						Type:        "bool",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the published template that is applied.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/orgs/[0-9]+/published_templates/[0-9a-f]+$"),
					},
				},
			},

			&metadata.Action{
				Name: "show_non_catalog",
				Description: `ShowNonCatalog retrieves applied policies that are not part of a regular aggregate. Only applied policies are applied from the project-scoped Template endpoint (as opposed to the org-wide Catalog) are part of this view. These template-based policies should largely be only used for development and testing purposes.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/policy_aggregates/non_catalog",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/policy_aggregates/non_catalog`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render applied policy items.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "index"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render applied policy items.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "index"},
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of a policy aggregate.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/policy_aggregates/%s",
						Variables:  []string{"org_id", "policy_aggregate_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/policy_aggregates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy aggregate`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy aggregate`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "source"},
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Delete asynchronously stops and deletes a policy aggregate. All individual applied policies in the aggregate will be stopped.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/governance/orgs/%s/policy_aggregates/%s",
						Variables:  []string{"org_id", "policy_aggregate_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/policy_aggregates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"PolicyTemplate": &metadata.Resource{
		Name: "PolicyTemplate",
		Description: `Show retrieves the details of a policy template.
**`,
		Identifier: "application/vnd.rightscale.policy_template",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `IndexPolicyTemplates retrieves the list of policy templates in a project.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/policy_templates",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/policy_templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy templates`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy templates`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
				},
			},

			&metadata.Action{
				Name: "upload",
				Description: `Upload uploads a policy template for a project, first compiling it. On failure, an array of syntax errors will be returned.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/projects/%s/policy_templates",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/policy_templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `filename is the name of the policy template source code file.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `source is the policy template source code.`,
						Type:        "file",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "PolicyTemplateUpload",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `filename is the name of the policy template source code file.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `source is the policy template source code.`,
						Type:        "*rsapi.FileUpload",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "compile",
				Description: `Compile compiles a policy template for a project. This is only to be used for checking the syntax of a policy template; the results are not stored.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/projects/%s/policy_templates/compile",
						Variables:  []string{"project_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/policy_templates/compile`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `filename is the name of the policy template source code file.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `source is the policy template source code.`,
						Type:        "file",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "PolicyTemplateCompile",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `filename is the name of the policy template source code file.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `source is the policy template source code.`,
						Type:        "*rsapi.FileUpload",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of a policy template.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/projects/%s/policy_templates/%s",
						Variables:  []string{"project_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/policy_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy template`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render policy template`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update updates a policy template in place for a project, by replacing it. Any existing applied policies using the template will not be updated; they must be deleted and recreated again.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/governance/projects/%s/policy_templates/%s",
						Variables:  []string{"project_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/policy_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `filename is the name of the policy template source code file.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `source is the policy template source code.`,
						Type:        "file",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "PolicyTemplateUpdate",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filename",
						Description: `filename is the name of the policy template source code file.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "source",
						Description: `source is the policy template source code.`,
						Type:        "*rsapi.FileUpload",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Delete deletes a policy template from a project. Deleting a policy template will not delete any applied policies created from the template, they must be stopped explicitly.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/governance/projects/%s/policy_templates/%s",
						Variables:  []string{"project_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/policy_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "retrieve_data",
				Description: `Retrieve Data retrieves the data sources specified in a give policy template.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/projects/%s/policy_templates/%s/retrieve_data",
						Variables:  []string{"project_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/projects/([^/]+)/policy_templates/([^/]+)/retrieve_data`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credentials",
						Description: `credentials is the map of name and credential used to launch the policy.`,
						Type:        "map",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "names[]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][name]",
						Description: `name of option`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options[][value]",
						Description: `value of option`,
						Type:        "interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				Payload: "PolicyTemplateRetrieveData",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credentials",
						Description: `credentials is the map of name and credential used to launch the policy.`,
						Type:        "map[string]interface{}",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "names",
						Description: `names is a filter to only retrieve datasources or resources that match the given names`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "options",
						Description: `options lists the configuration options used to parameterize the policy.`,
						Type:        "[]*ConfigurationOptionCreateType",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"PublishedTemplate": &metadata.Resource{
		Name: "PublishedTemplate",
		Description: `Show retrieves the details of a published template.
**`,
		Identifier: "application/vnd.rightscale.published_template",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name: "index",
				Description: `Index retrieves the list of published templates in an organization.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/published_templates",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/published_templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render published templates`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
					&metadata.ActionParam{
						Name:        "show_hidden",
						Description: `show_hidden will show templates that have been hidden if set to true.`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render published templates`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended"},
					},
					&metadata.ActionParam{
						Name:        "show_hidden",
						Description: `show_hidden will show templates that have been hidden if set to true.`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "create",
				Description: `Create creates an organization-scoped published template from a project-scoped policy template.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/orgs/%s/published_templates",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/published_templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the policy template to publish to the organization.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"),
					},
				},
				Payload: "PublishedTemplateCreate",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the policy template to publish to the organization.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"),
					},
				},
			},

			&metadata.Action{
				Name: "show",
				Description: `Show retrieves the details of a published template.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/api/governance/orgs/%s/published_templates/%s",
						Variables:  []string{"org_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/published_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render published template`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `View used to render published template`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "source"},
					},
				},
			},

			&metadata.Action{
				Name: "update",
				Description: `Update updates a published template in place for an organization, by replacing it. Any existing applied policies using the template will not be updated; they must be deleted and recreated again.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/api/governance/orgs/%s/published_templates/%s",
						Variables:  []string{"org_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/published_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the policy template to publish to the organization.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"),
					},
				},
				Payload: "PublishedTemplateUpdate",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "template_href",
						Description: `template_href is the href of the policy template to publish to the organization.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile("^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"),
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Delete deletes a published template from an organization. Deleting a published template will not delete any applied policies created from the template, they must be stopped explicitly.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/api/governance/orgs/%s/published_templates/%s",
						Variables:  []string{"org_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/published_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "hide",
				Description: `Hide hides a RightScale built-in template from an organization.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/orgs/%s/published_templates/%s/hide",
						Variables:  []string{"org_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/published_templates/([^/]+)/hide`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name: "unhide",
				Description: `Unhide unhides a RightScale built-in template from an organization.
**`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/api/governance/orgs/%s/published_templates/%s/unhide",
						Variables:  []string{"org_id", "template_id"},
						Regexp:     regexp.MustCompile(`/api/governance/orgs/([^/]+)/published_templates/([^/]+)/unhide`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
}
