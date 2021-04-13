//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ goav2gen -metadata=policy/docs -output=policy -pkg=policy -version=1.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package policy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/rightscale/rsc.v9/metadata"
	"gopkg.in/rightscale/rsc.v9/rsapi"
)

// API Version
const APIVersion = "1.0"

// An Href contains the relative path to a resource or resource collection,
// e.g. "/api/servers/123" or "/api/servers".
type Href string

// ActionPath computes the path to the given resource action. For example given the href
// "/api/servers/123" calling ActionPath with resource "servers" and action "clone" returns the path
// "/api/servers/123/clone" and verb POST.
// The algorithm consists of extracting the variables from the href by looking up a matching
// pattern from the resource metadata. The variables are then substituted in the action path.
// If there are more than one pattern that match the href then the algorithm picks the one that can
// substitute the most variables.
func (r *Href) ActionPath(rName, aName string) (*metadata.ActionPath, error) {
	res, ok := GenMetadata[rName]
	if !ok {
		return nil, fmt.Errorf("No resource with name '%s'", rName)
	}
	var action *metadata.Action
	for _, a := range res.Actions {
		if a.Name == aName {
			action = a
			break
		}
	}
	if action == nil {
		return nil, fmt.Errorf("No action with name '%s' on %s", aName, rName)
	}
	vars, err := res.ExtractVariables(string(*r))
	if err != nil {
		return nil, err
	}
	return action.URL(vars)
}

/****** ActionStatus ******/

// Show retrieves the details of an action status.
// **
type ActionStatus struct {
	FinishedAt *time.Time             `json:"finished_at,omitempty"`
	Id         string                 `json:"id,omitempty"`
	Kind       string                 `json:"kind,omitempty"`
	Label      string                 `json:"label,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Options    []*ConfigurationOption `json:"options,omitempty"`
	RunBy      *User                  `json:"run_by,omitempty"`
	StartedAt  *time.Time             `json:"started_at,omitempty"`
	Status     string                 `json:"status,omitempty"`
}

//===== Locator

// ActionStatusLocator exposes the ActionStatus resource actions.
type ActionStatusLocator struct {
	Href
	api *API
}

// ActionStatusLocator builds a locator from the given href.
func (api *API) ActionStatusLocator(href string) *ActionStatusLocator {
	return &ActionStatusLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/projects/{project_id}/action_status
//
// Index returns a list of action statuses in a project.
// **
func (loc *ActionStatusLocator) Index(options rsapi.APIParams) (*ActionStatusList, error) {
	var res *ActionStatusList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var incidentIdOpt = options["incident_id"]
	if incidentIdOpt != nil {
		params["incident_id"] = incidentIdOpt
	}
	var appliedPolicyIdOpt = options["applied_policy_id"]
	if appliedPolicyIdOpt != nil {
		params["applied_policy_id"] = appliedPolicyIdOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ActionStatus", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/action_status/{id}
//
// Show retrieves the details of an action status.
// **
func (loc *ActionStatusLocator) Show(options rsapi.APIParams) (*ActionStatus, error) {
	var res *ActionStatus
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ActionStatus", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

/****** AppliedPolicy ******/

// Show retrieves the details of an applied policy.
// **
type AppliedPolicy struct {
	Category            string                 `json:"category,omitempty"`
	CreatedAt           *time.Time             `json:"created_at,omitempty"`
	CreatedBy           *User                  `json:"created_by,omitempty"`
	Credentials         map[string]interface{} `json:"credentials,omitempty"`
	Description         string                 `json:"description,omitempty"`
	DocLink             string                 `json:"doc_link,omitempty"`
	DryRun              bool                   `json:"dry_run,omitempty"`
	Error               string                 `json:"error,omitempty"`
	ErroredAt           *time.Time             `json:"errored_at,omitempty"`
	Frequency           string                 `json:"frequency,omitempty"`
	Href                string                 `json:"href,omitempty"`
	Id                  string                 `json:"id,omitempty"`
	IncidentAggregateId string                 `json:"incident_aggregate_id,omitempty"`
	Info                map[string]interface{} `json:"info,omitempty"`
	Kind                string                 `json:"kind,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Options             []*ConfigurationOption `json:"options,omitempty"`
	PolicyAggregateId   string                 `json:"policy_aggregate_id,omitempty"`
	PolicyTemplate      *PolicyTemplateLink    `json:"policy_template,omitempty"`
	Project             *Project               `json:"project,omitempty"`
	PublishedTemplate   *PublishedTemplateLink `json:"published_template,omitempty"`
	Scope               string                 `json:"scope,omitempty"`
	Severity            string                 `json:"severity,omitempty"`
	SkipApprovals       bool                   `json:"skip_approvals,omitempty"`
	Status              string                 `json:"status,omitempty"`
	UpdatedAt           *time.Time             `json:"updated_at,omitempty"`
}

//===== Locator

// AppliedPolicyLocator exposes the AppliedPolicy resource actions.
type AppliedPolicyLocator struct {
	Href
	api *API
}

// AppliedPolicyLocator builds a locator from the given href.
func (api *API) AppliedPolicyLocator(href string) *AppliedPolicyLocator {
	return &AppliedPolicyLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/projects/{project_id}/applied_policies
//
// Index retrieves the list of applied policies in a project.
// **
func (loc *AppliedPolicyLocator) Index(options rsapi.APIParams) (*AppliedPolicyList, error) {
	var res *AppliedPolicyList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		params["name"] = nameOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicy", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /api/governance/projects/{project_id}/applied_policies
//
// Create applies a policy template to a given project. The applied policy will continually run until deleted.
// **
func (loc *AppliedPolicyLocator) Create(name string, templateHref string, options rsapi.APIParams) (*AppliedPolicy, error) {
	var res *AppliedPolicy
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if templateHref == "" {
		return res, fmt.Errorf("templateHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name":          name,
		"template_href": templateHref,
	}
	var credentialsOpt = options["credentials"]
	if credentialsOpt != nil {
		p["credentials"] = credentialsOpt
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var dryRunOpt = options["dry_run"]
	if dryRunOpt != nil {
		p["dry_run"] = dryRunOpt
	}
	var frequencyOpt = options["frequency"]
	if frequencyOpt != nil {
		p["frequency"] = frequencyOpt
	}
	var optionsOpt = options["options"]
	if optionsOpt != nil {
		p["options"] = optionsOpt
	}
	var severityOpt = options["severity"]
	if severityOpt != nil {
		p["severity"] = severityOpt
	}
	var skipApprovalsOpt = options["skip_approvals"]
	if skipApprovalsOpt != nil {
		p["skip_approvals"] = skipApprovalsOpt
	}
	uri, err := loc.ActionPath("AppliedPolicy", "create")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/applied_policies/{policy_id}
//
// Show retrieves the details of an applied policy.
// **
func (loc *AppliedPolicyLocator) Show(options rsapi.APIParams) (*AppliedPolicy, error) {
	var res *AppliedPolicy
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicy", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// DELETE /api/governance/projects/{project_id}/applied_policies/{policy_id}
//
// Delete stops and deletes an applied policy.
// **
func (loc *AppliedPolicyLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicy", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// POST /api/governance/projects/{project_id}/applied_policies/{policy_id}/evaluate
//
// Evaluate executes an applied policy evaluation on demand. It does not affect the normal execution schedule.
// **
func (loc *AppliedPolicyLocator) Evaluate() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicy", "evaluate")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// GET /api/governance/projects/{project_id}/applied_policies/{policy_id}/log
//
// ShowLog retrieves the last evaluation log of an applied policy. *The content type is "text/markdown"*.
// **
func (loc *AppliedPolicyLocator) ShowLog() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicy", "show_log")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// GET /api/governance/projects/{project_id}/applied_policies/{policy_id}/status
//
// ShowStatus retrieves the evaluation status details of an applied policy.
// **
func (loc *AppliedPolicyLocator) ShowStatus() (*AppliedPolicyStatus, error) {
	var res *AppliedPolicyStatus
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicy", "show_status")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

/****** Approval ******/

type Approval struct {
}

//===== Locator

// ApprovalLocator exposes the Approval resource actions.
type ApprovalLocator struct {
	Href
	api *API
}

// ApprovalLocator builds a locator from the given href.
func (api *API) ApprovalLocator(href string) *ApprovalLocator {
	return &ApprovalLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/projects/{project_id}/approval_requests
//
// Index retrieves the list of approval requests in a project.
// **
func (loc *ApprovalLocator) Index(options rsapi.APIParams) (*ApprovalRequestList, error) {
	var res *ApprovalRequestList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var idOpt = options["id"]
	if idOpt != nil {
		params["id"] = idOpt
	}
	var subjectKindOpt = options["subject_kind"]
	if subjectKindOpt != nil {
		params["subject_kind"] = subjectKindOpt
	}
	var subjectHrefOpt = options["subject_href"]
	if subjectHrefOpt != nil {
		params["subject_href"] = subjectHrefOpt
	}
	var statusOpt = options["status"]
	if statusOpt != nil {
		params["status"] = statusOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Approval", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/approval_requests/{approval_request_id}
//
// Show retrieves the details of an approval request.
// **
func (loc *ApprovalLocator) Show(options rsapi.APIParams) (*ApprovalRequest, error) {
	var res *ApprovalRequest
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Approval", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /api/governance/projects/{project_id}/approval_requests/{approval_request_id}/approve
//
// Approve approves a single approval request.
// **
func (loc *ApprovalLocator) Approve(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var optionsOpt = options["options"]
	if optionsOpt != nil {
		p["options"] = optionsOpt
	}
	uri, err := loc.ActionPath("Approval", "approve")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// POST /api/governance/projects/{project_id}/approval_requests/{approval_request_id}/deny
//
// Deny denies a single approval request.
// **
func (loc *ApprovalLocator) Deny(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var commentOpt = options["comment"]
	if commentOpt != nil {
		p["comment"] = commentOpt
	}
	uri, err := loc.ActionPath("Approval", "deny")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

/****** ArchivedIncident ******/

// Show retrieves the details of an archived incident.
// **
type ArchivedIncident struct {
	ActionFailed       bool                   `json:"action_failed,omitempty"`
	AppliedPolicy      *AppliedPolicyLink     `json:"applied_policy,omitempty"`
	Category           string                 `json:"category,omitempty"`
	CreatedAt          *time.Time             `json:"created_at,omitempty"`
	DryRun             bool                   `json:"dry_run,omitempty"`
	Etag               string                 `json:"etag,omitempty"`
	Href               string                 `json:"href,omitempty"`
	Id                 string                 `json:"id,omitempty"`
	Kind               string                 `json:"kind,omitempty"`
	NotModified        string                 `json:"not_modified,omitempty"`
	Options            []*ConfigurationOption `json:"options,omitempty"`
	Project            *Project               `json:"project,omitempty"`
	ResolutionMessage  string                 `json:"resolution_message,omitempty"`
	ResolvedAt         *time.Time             `json:"resolved_at,omitempty"`
	ResolvedBy         *User                  `json:"resolved_by,omitempty"`
	Severity           string                 `json:"severity,omitempty"`
	State              string                 `json:"state,omitempty"`
	Summary            string                 `json:"summary,omitempty"`
	UpdatedAt          *time.Time             `json:"updated_at,omitempty"`
	ViolationDataCount int                    `json:"violation_data_count,omitempty"`
}

//===== Locator

// ArchivedIncidentLocator exposes the ArchivedIncident resource actions.
type ArchivedIncidentLocator struct {
	Href
	api *API
}

// ArchivedIncidentLocator builds a locator from the given href.
func (api *API) ArchivedIncidentLocator(href string) *ArchivedIncidentLocator {
	return &ArchivedIncidentLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/projects/{project_id}/archived_incidents
//
// Index retrieves the list of archived_incidents in a project.
// **
func (loc *ArchivedIncidentLocator) Index(options rsapi.APIParams) (*ArchivedIncidentList, error) {
	var res *ArchivedIncidentList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var stateOpt = options["state"]
	if stateOpt != nil {
		params["state"] = stateOpt
	}
	var appliedPolicyIdOpt = options["applied_policy_id"]
	if appliedPolicyIdOpt != nil {
		params["applied_policy_id"] = appliedPolicyIdOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ArchivedIncident", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/archived_incidents/{incident_id}
//
// Show retrieves the details of an archived incident.
// **
func (loc *ArchivedIncidentLocator) Show(options rsapi.APIParams) (*ArchivedIncident, error) {
	var res *ArchivedIncident
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ArchivedIncident", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/archived_incidents/{incident_id}/escalations
//
// IndexEscalations retrieves the status details of all of the escalations for an archived incident.
// **
func (loc *ArchivedIncidentLocator) IndexEscalations() (*Escalations, error) {
	var res *Escalations
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ArchivedIncident", "index_escalations")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/archived_incidents/{incident_id}/resolutions
//
// IndexResolutions retrieves the status details of all of the resolutions for an archived incident.
// **
func (loc *ArchivedIncidentLocator) IndexResolutions() (*Resolutions, error) {
	var res *Resolutions
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ArchivedIncident", "index_resolutions")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

/****** Incident ******/

// Show retrieves the details of an incident.
// **
type Incident struct {
	ActionFailed        bool                   `json:"action_failed,omitempty"`
	AppliedPolicy       *AppliedPolicyLink     `json:"applied_policy,omitempty"`
	Category            string                 `json:"category,omitempty"`
	CreatedAt           *time.Time             `json:"created_at,omitempty"`
	DryRun              bool                   `json:"dry_run,omitempty"`
	Etag                string                 `json:"etag,omitempty"`
	Href                string                 `json:"href,omitempty"`
	Id                  string                 `json:"id,omitempty"`
	IncidentAggregateId string                 `json:"incident_aggregate_id,omitempty"`
	Kind                string                 `json:"kind,omitempty"`
	NotModified         string                 `json:"not_modified,omitempty"`
	Options             []*ConfigurationOption `json:"options,omitempty"`
	Project             *Project               `json:"project,omitempty"`
	ResolutionMessage   string                 `json:"resolution_message,omitempty"`
	ResolvedAt          *time.Time             `json:"resolved_at,omitempty"`
	ResolvedBy          *User                  `json:"resolved_by,omitempty"`
	Severity            string                 `json:"severity,omitempty"`
	State               string                 `json:"state,omitempty"`
	Summary             string                 `json:"summary,omitempty"`
	UpdatedAt           *time.Time             `json:"updated_at,omitempty"`
	ViolationDataCount  int                    `json:"violation_data_count,omitempty"`
}

//===== Locator

// IncidentLocator exposes the Incident resource actions.
type IncidentLocator struct {
	Href
	api *API
}

// IncidentLocator builds a locator from the given href.
func (api *API) IncidentLocator(href string) *IncidentLocator {
	return &IncidentLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/projects/{project_id}/incidents
//
// Index retrieves the list of incidents in a project.
// **
func (loc *IncidentLocator) Index(options rsapi.APIParams) (*IncidentList, error) {
	var res *IncidentList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var stateOpt = options["state"]
	if stateOpt != nil {
		params["state"] = stateOpt
	}
	var appliedPolicyIdOpt = options["applied_policy_id"]
	if appliedPolicyIdOpt != nil {
		params["applied_policy_id"] = appliedPolicyIdOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Incident", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/incidents/{incident_id}
//
// Show retrieves the details of an incident.
// **
func (loc *IncidentLocator) Show(options rsapi.APIParams) (*Incident, error) {
	var res *Incident
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Incident", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /api/governance/projects/{project_id}/incidents/{incident_id}/actions/{action_id}/run_action
//
// RunAction executes any action listed in available_actions on an incident. It can run against all resources in an incident or only a selected amount, depending on passed in options. Actions will run in parallel.
// **
func (loc *IncidentLocator) RunAction(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var optionsOpt = options["options"]
	if optionsOpt != nil {
		p["options"] = optionsOpt
	}
	uri, err := loc.ActionPath("Incident", "run_action")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// GET /api/governance/projects/{project_id}/incidents/{incident_id}/escalations
//
// IndexEscalations retrieves the status details of all of the escalations for an incident. This API method is deprecated and will no longer be updated as of July 30, 2020. Please use the index_statuses method instead.
// **
func (loc *IncidentLocator) IndexEscalations() (*Escalations, error) {
	var res *Escalations
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Incident", "index_escalations")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/projects/{project_id}/incidents/{incident_id}/resolutions
//
// IndexResolutions retrieves the status details of all of the resolutions for an incident. This API method is deprecated and will no longer be updated as of July 30, 2020. Please use the index_statuses method instead.
// **
func (loc *IncidentLocator) IndexResolutions() (*Resolutions, error) {
	var res *Resolutions
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Incident", "index_resolutions")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// PUT /api/governance/projects/{project_id}/incidents/{incident_id}/resolve
//
// Resolve resolves an incident by setting it to an inactive state, indicating that it has been addressed.
// **
func (loc *IncidentLocator) Resolve() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Incident", "resolve")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

/****** IncidentAggregate ******/

// Show retrieves the details of an aggregate.
// **
type IncidentAggregate struct {
	ActionSummary   *ActionSummary           `json:"action_summary,omitempty"`
	Category        string                   `json:"category,omitempty"`
	Count           int                      `json:"count,omitempty"`
	CreatedAt       *time.Time               `json:"created_at,omitempty"`
	DryRun          bool                     `json:"dry_run,omitempty"`
	Href            string                   `json:"href,omitempty"`
	Id              string                   `json:"id,omitempty"`
	IncidentSummary *IncidentSummary         `json:"incident_summary,omitempty"`
	Items           []*IncidentAggregateItem `json:"items,omitempty"`
	Kind            string                   `json:"kind,omitempty"`
	NotModified     string                   `json:"not_modified,omitempty"`
	Org             *Org                     `json:"org,omitempty"`
	PolicyAggregate *PolicyAggregateLink     `json:"policy_aggregate,omitempty"`
	Severity        string                   `json:"severity,omitempty"`
	State           string                   `json:"state,omitempty"`
	UpdatedAt       *time.Time               `json:"updated_at,omitempty"`
}

//===== Locator

// IncidentAggregateLocator exposes the IncidentAggregate resource actions.
type IncidentAggregateLocator struct {
	Href
	api *API
}

// IncidentAggregateLocator builds a locator from the given href.
func (api *API) IncidentAggregateLocator(href string) *IncidentAggregateLocator {
	return &IncidentAggregateLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/orgs/{org_id}/incident_aggregates
//
// Index retrieves the list of incident_aggregates in an organization.
// **
func (loc *IncidentAggregateLocator) Index(options rsapi.APIParams) (*IncidentAggregateList, error) {
	var res *IncidentAggregateList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentAggregate", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/orgs/{org_id}/incident_aggregates/non_catalog
//
// ShowNonCatalog retrieves a list of incidents in the non-catalog policy aggregate. These incidents largely originate from dev/test environments.
// **
func (loc *IncidentAggregateLocator) ShowNonCatalog(options rsapi.APIParams) (*IncidentAggregateNonCatalog, error) {
	var res *IncidentAggregateNonCatalog
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentAggregate", "show_non_catalog")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/orgs/{org_id}/incident_aggregates/{incident_aggregate_id}
//
// Show retrieves the details of an aggregate.
// **
func (loc *IncidentAggregateLocator) Show(options rsapi.APIParams) (*IncidentAggregate, error) {
	var res *IncidentAggregate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentAggregate", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

/****** PolicyAggregate ******/

// Show retrieves the details of a policy aggregate.
// **
type PolicyAggregate struct {
	ActiveCount           int                    `json:"active_count,omitempty"`
	Category              string                 `json:"category,omitempty"`
	Count                 int                    `json:"count,omitempty"`
	CreatedAt             *time.Time             `json:"created_at,omitempty"`
	CreatedBy             *User                  `json:"created_by,omitempty"`
	Credentials           map[string]interface{} `json:"credentials,omitempty"`
	Description           string                 `json:"description,omitempty"`
	DocLink               string                 `json:"doc_link,omitempty"`
	DryRun                bool                   `json:"dry_run,omitempty"`
	ErrorCount            int                    `json:"error_count,omitempty"`
	Errors                map[string]interface{} `json:"errors,omitempty"`
	ExcludedProjectIds    []int                  `json:"excluded_project_ids,omitempty"`
	Frequency             string                 `json:"frequency,omitempty"`
	Href                  string                 `json:"href,omitempty"`
	Id                    string                 `json:"id,omitempty"`
	IncidentAggregateHref string                 `json:"incident_aggregate_href,omitempty"`
	Info                  map[string]interface{} `json:"info,omitempty"`
	Items                 []*PolicyAggregateItem `json:"items,omitempty"`
	Kind                  string                 `json:"kind,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	Options               []*ConfigurationOption `json:"options,omitempty"`
	Org                   *Org                   `json:"org,omitempty"`
	ProjectIds            []int                  `json:"project_ids,omitempty"`
	PublishedTemplate     *PublishedTemplateLink `json:"published_template,omitempty"`
	RunningProjectIds     []int                  `json:"running_project_ids,omitempty"`
	Severity              string                 `json:"severity,omitempty"`
	SkipApprovals         bool                   `json:"skip_approvals,omitempty"`
	Status                string                 `json:"status,omitempty"`
	UpdatedAt             *time.Time             `json:"updated_at,omitempty"`
}

//===== Locator

// PolicyAggregateLocator exposes the PolicyAggregate resource actions.
type PolicyAggregateLocator struct {
	Href
	api *API
}

// PolicyAggregateLocator builds a locator from the given href.
func (api *API) PolicyAggregateLocator(href string) *PolicyAggregateLocator {
	return &PolicyAggregateLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/orgs/{org_id}/policy_aggregates
//
// Index retrieves the list of policy aggregates in an org.
// **
func (loc *PolicyAggregateLocator) Index(options rsapi.APIParams) (*PolicyAggregateList, error) {
	var res *PolicyAggregateList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		params["name"] = nameOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyAggregate", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /api/governance/orgs/{org_id}/policy_aggregates
//
// Create applies a policy template to a given project. The policy aggregate will continually run until deleted.
// **
func (loc *PolicyAggregateLocator) Create(name string, templateHref string, options rsapi.APIParams) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if templateHref == "" {
		return fmt.Errorf("templateHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name":          name,
		"template_href": templateHref,
	}
	var allProjectsOpt = options["all_projects"]
	if allProjectsOpt != nil {
		p["all_projects"] = allProjectsOpt
	}
	var credentialsOpt = options["credentials"]
	if credentialsOpt != nil {
		p["credentials"] = credentialsOpt
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var dryRunOpt = options["dry_run"]
	if dryRunOpt != nil {
		p["dry_run"] = dryRunOpt
	}
	var frequencyOpt = options["frequency"]
	if frequencyOpt != nil {
		p["frequency"] = frequencyOpt
	}
	var optionsOpt = options["options"]
	if optionsOpt != nil {
		p["options"] = optionsOpt
	}
	var projectIdsOpt = options["project_ids"]
	if projectIdsOpt != nil {
		p["project_ids"] = projectIdsOpt
	}
	var severityOpt = options["severity"]
	if severityOpt != nil {
		p["severity"] = severityOpt
	}
	var skipApprovalsOpt = options["skip_approvals"]
	if skipApprovalsOpt != nil {
		p["skip_approvals"] = skipApprovalsOpt
	}
	uri, err := loc.ActionPath("PolicyAggregate", "create")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// GET /api/governance/orgs/{org_id}/policy_aggregates/non_catalog
//
// ShowNonCatalog retrieves applied policies that are not part of a regular aggregate. Only applied policies are applied from the project-scoped Template endpoint (as opposed to the org-wide Catalog) are part of this view. These template-based policies should largely be only used for development and testing purposes.
// **
func (loc *PolicyAggregateLocator) ShowNonCatalog(options rsapi.APIParams) (*PolicyAggregateNonCatalog, error) {
	var res *PolicyAggregateNonCatalog
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyAggregate", "show_non_catalog")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /api/governance/orgs/{org_id}/policy_aggregates/{policy_aggregate_id}
//
// Show retrieves the details of a policy aggregate.
// **
func (loc *PolicyAggregateLocator) Show(options rsapi.APIParams) (*PolicyAggregate, error) {
	var res *PolicyAggregate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyAggregate", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// DELETE /api/governance/orgs/{org_id}/policy_aggregates/{policy_aggregate_id}
//
// Delete asynchronously stops and deletes a policy aggregate. All individual applied policies in the aggregate will be stopped.
// **
func (loc *PolicyAggregateLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyAggregate", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

/****** PolicyTemplate ******/

// Show retrieves the details of a policy template.
// **
type PolicyTemplate struct {
	Category         string                 `json:"category,omitempty"`
	CreatedAt        *time.Time             `json:"created_at,omitempty"`
	CreatedBy        *User                  `json:"created_by,omitempty"`
	DefaultFrequency string                 `json:"default_frequency,omitempty"`
	DocLink          string                 `json:"doc_link,omitempty"`
	Fingerprint      string                 `json:"fingerprint,omitempty"`
	Href             string                 `json:"href,omitempty"`
	Id               string                 `json:"id,omitempty"`
	Info             map[string]interface{} `json:"info,omitempty"`
	Kind             string                 `json:"kind,omitempty"`
	Name             string                 `json:"name,omitempty"`
	ProjectId        int                    `json:"project_id,omitempty"`
	RequiredRoles    []string               `json:"required_roles,omitempty"`
	RsPtVer          int                    `json:"rs_pt_ver,omitempty"`
	Severity         string                 `json:"severity,omitempty"`
	ShortDescription string                 `json:"short_description,omitempty"`
	Tenancy          string                 `json:"tenancy,omitempty"`
	UpdatedAt        *time.Time             `json:"updated_at,omitempty"`
	UpdatedBy        *User                  `json:"updated_by,omitempty"`
}

//===== Locator

// PolicyTemplateLocator exposes the PolicyTemplate resource actions.
type PolicyTemplateLocator struct {
	Href
	api *API
}

// PolicyTemplateLocator builds a locator from the given href.
func (api *API) PolicyTemplateLocator(href string) *PolicyTemplateLocator {
	return &PolicyTemplateLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/projects/{project_id}/policy_templates
//
// IndexPolicyTemplates retrieves the list of policy templates in a project.
// **
func (loc *PolicyTemplateLocator) Index(options rsapi.APIParams) (*PolicyTemplateList, error) {
	var res *PolicyTemplateList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyTemplate", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /api/governance/projects/{project_id}/policy_templates
//
// Upload uploads a policy template for a project, first compiling it. On failure, an array of syntax errors will be returned.
// **
func (loc *PolicyTemplateLocator) Upload(filename string, source *rsapi.FileUpload) (*PolicyTemplate, error) {
	var res *PolicyTemplate
	if filename == "" {
		return res, fmt.Errorf("filename is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"filename": filename,
		"source":   source,
	}
	uri, err := loc.ActionPath("PolicyTemplate", "upload")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /api/governance/projects/{project_id}/policy_templates/compile
//
// Compile compiles a policy template for a project. This is only to be used for checking the syntax of a policy template; the results are not stored.
// **
func (loc *PolicyTemplateLocator) Compile(filename string, source *rsapi.FileUpload) error {
	if filename == "" {
		return fmt.Errorf("filename is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"filename": filename,
		"source":   source,
	}
	uri, err := loc.ActionPath("PolicyTemplate", "compile")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// GET /api/governance/projects/{project_id}/policy_templates/{template_id}
//
// Show retrieves the details of a policy template.
// **
func (loc *PolicyTemplateLocator) Show(options rsapi.APIParams) (*PolicyTemplate, error) {
	var res *PolicyTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyTemplate", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// PUT /api/governance/projects/{project_id}/policy_templates/{template_id}
//
// Update updates a policy template in place for a project, by replacing it. Any existing applied policies using the template will not be updated; they must be deleted and recreated again.
// **
func (loc *PolicyTemplateLocator) Update(filename string, source *rsapi.FileUpload) (*PolicyTemplate, error) {
	var res *PolicyTemplate
	if filename == "" {
		return res, fmt.Errorf("filename is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"filename": filename,
		"source":   source,
	}
	uri, err := loc.ActionPath("PolicyTemplate", "update")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// DELETE /api/governance/projects/{project_id}/policy_templates/{template_id}
//
// Delete deletes a policy template from a project. Deleting a policy template will not delete any applied policies created from the template, they must be stopped explicitly.
// **
func (loc *PolicyTemplateLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyTemplate", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// POST /api/governance/projects/{project_id}/policy_templates/{template_id}/retrieve_data
//
// Retrieve Data retrieves the data sources specified in a give policy template.
// **
func (loc *PolicyTemplateLocator) RetrieveData(options rsapi.APIParams) (interface{}, error) {
	var res interface{}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var credentialsOpt = options["credentials"]
	if credentialsOpt != nil {
		p["credentials"] = credentialsOpt
	}
	var namesOpt = options["names"]
	if namesOpt != nil {
		p["names"] = namesOpt
	}
	var optionsOpt = options["options"]
	if optionsOpt != nil {
		p["options"] = optionsOpt
	}
	uri, err := loc.ActionPath("PolicyTemplate", "retrieve_data")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

/****** PublishedTemplate ******/

// Show retrieves the details of a published template.
// **
type PublishedTemplate struct {
	BuiltIn                   bool                   `json:"built_in,omitempty"`
	Category                  string                 `json:"category,omitempty"`
	CreatedAt                 *time.Time             `json:"created_at,omitempty"`
	CreatedBy                 *User                  `json:"created_by,omitempty"`
	DefaultFrequency          string                 `json:"default_frequency,omitempty"`
	DocLink                   string                 `json:"doc_link,omitempty"`
	Fingerprint               string                 `json:"fingerprint,omitempty"`
	Hidden                    bool                   `json:"hidden,omitempty"`
	HiddenAt                  *time.Time             `json:"hidden_at,omitempty"`
	HiddenBy                  *User                  `json:"hidden_by,omitempty"`
	Href                      string                 `json:"href,omitempty"`
	Id                        string                 `json:"id,omitempty"`
	Info                      map[string]interface{} `json:"info,omitempty"`
	Kind                      string                 `json:"kind,omitempty"`
	Name                      string                 `json:"name,omitempty"`
	OrgId                     int                    `json:"org_id,omitempty"`
	PolicyTemplateFingerprint string                 `json:"policy_template_fingerprint,omitempty"`
	PolicyTemplateId          string                 `json:"policy_template_id,omitempty"`
	PolicyTemplateUrl         string                 `json:"policy_template_url,omitempty"`
	ProjectId                 int                    `json:"project_id,omitempty"`
	RequiredRoles             []string               `json:"required_roles,omitempty"`
	RsPtVer                   int                    `json:"rs_pt_ver,omitempty"`
	Severity                  string                 `json:"severity,omitempty"`
	ShortDescription          string                 `json:"short_description,omitempty"`
	Tenancy                   string                 `json:"tenancy,omitempty"`
	UpdatedAt                 *time.Time             `json:"updated_at,omitempty"`
	UpdatedBy                 *User                  `json:"updated_by,omitempty"`
}

//===== Locator

// PublishedTemplateLocator exposes the PublishedTemplate resource actions.
type PublishedTemplateLocator struct {
	Href
	api *API
}

// PublishedTemplateLocator builds a locator from the given href.
func (api *API) PublishedTemplateLocator(href string) *PublishedTemplateLocator {
	return &PublishedTemplateLocator{Href(href), api}
}

//===== Actions

// GET /api/governance/orgs/{org_id}/published_templates
//
// Index retrieves the list of published templates in an organization.
// **
func (loc *PublishedTemplateLocator) Index(options rsapi.APIParams) (*PublishedTemplateList, error) {
	var res *PublishedTemplateList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var showHiddenOpt = options["show_hidden"]
	if showHiddenOpt != nil {
		params["show_hidden"] = showHiddenOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplate", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /api/governance/orgs/{org_id}/published_templates
//
// Create creates an organization-scoped published template from a project-scoped policy template.
// **
func (loc *PublishedTemplateLocator) Create(templateHref string) error {
	if templateHref == "" {
		return fmt.Errorf("templateHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"template_href": templateHref,
	}
	uri, err := loc.ActionPath("PublishedTemplate", "create")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// GET /api/governance/orgs/{org_id}/published_templates/{template_id}
//
// Show retrieves the details of a published template.
// **
func (loc *PublishedTemplateLocator) Show(options rsapi.APIParams) (*PublishedTemplate, error) {
	var res *PublishedTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplate", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// PUT /api/governance/orgs/{org_id}/published_templates/{template_id}
//
// Update updates a published template in place for an organization, by replacing it. Any existing applied policies using the template will not be updated; they must be deleted and recreated again.
// **
func (loc *PublishedTemplateLocator) Update(templateHref string) error {
	if templateHref == "" {
		return fmt.Errorf("templateHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"template_href": templateHref,
	}
	uri, err := loc.ActionPath("PublishedTemplate", "update")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// DELETE /api/governance/orgs/{org_id}/published_templates/{template_id}
//
// Delete deletes a published template from an organization. Deleting a published template will not delete any applied policies created from the template, they must be stopped explicitly.
// **
func (loc *PublishedTemplateLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplate", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// POST /api/governance/orgs/{org_id}/published_templates/{template_id}/hide
//
// Hide hides a RightScale built-in template from an organization.
// **
func (loc *PublishedTemplateLocator) Hide() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplate", "hide")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

// POST /api/governance/orgs/{org_id}/published_templates/{template_id}/unhide
//
// Unhide unhides a RightScale built-in template from an organization.
// **
func (loc *PublishedTemplateLocator) Unhide() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplate", "unhide")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

/****** Data Types ******/

type ActionStatusList struct {
	Count       int             `json:"count,omitempty"`
	Items       []*ActionStatus `json:"items,omitempty"`
	Kind        string          `json:"kind,omitempty"`
	NotModified string          `json:"not_modified,omitempty"`
}

type ActionSummary struct {
	FailedCount  int `json:"failed_count,omitempty"`
	PendingCount int `json:"pending_count,omitempty"`
}

type AppliedPolicyCreate struct {
	Credentials   map[string]interface{}           `json:"credentials,omitempty"`
	Description   string                           `json:"description,omitempty"`
	DryRun        bool                             `json:"dry_run,omitempty"`
	Frequency     string                           `json:"frequency,omitempty"`
	Name          string                           `json:"name,omitempty"`
	Options       []*ConfigurationOptionCreateType `json:"options,omitempty"`
	Severity      string                           `json:"severity,omitempty"`
	SkipApprovals bool                             `json:"skip_approvals,omitempty"`
	TemplateHref  string                           `json:"template_href,omitempty"`
}

type AppliedPolicyLink struct {
	CreatedAt           *time.Time             `json:"created_at,omitempty"`
	CreatedBy           *User                  `json:"created_by,omitempty"`
	Frequency           string                 `json:"frequency,omitempty"`
	Href                string                 `json:"href,omitempty"`
	Id                  string                 `json:"id,omitempty"`
	IncidentAggregateId string                 `json:"incident_aggregate_id,omitempty"`
	Kind                string                 `json:"kind,omitempty"`
	Name                string                 `json:"name,omitempty"`
	PolicyAggregateId   string                 `json:"policy_aggregate_id,omitempty"`
	PolicyTemplate      *PolicyTemplateLink    `json:"policy_template,omitempty"`
	PublishedTemplate   *PublishedTemplateLink `json:"published_template,omitempty"`
}

type AppliedPolicyList struct {
	Count       int              `json:"count,omitempty"`
	Items       []*AppliedPolicy `json:"items,omitempty"`
	Kind        string           `json:"kind,omitempty"`
	NotModified string           `json:"not_modified,omitempty"`
}

type AppliedPolicyStatus struct {
	EvaluationError      string     `json:"evaluation_error,omitempty"`
	EvaluationErroredAt  *time.Time `json:"evaluation_errored_at,omitempty"`
	LastEvaluationFinish *time.Time `json:"last_evaluation_finish,omitempty"`
	LastEvaluationStart  *time.Time `json:"last_evaluation_start,omitempty"`
	NextEvaluationStart  *time.Time `json:"next_evaluation_start,omitempty"`
}

type ApprovalApprove struct {
	Options []*ConfigurationOptionCreateType `json:"options,omitempty"`
}

type ApprovalDeny struct {
	Comment string `json:"comment,omitempty"`
}

type ApprovalRequest struct {
	CreatedAt   *time.Time       `json:"created_at,omitempty"`
	Description string           `json:"description,omitempty"`
	Href        string           `json:"href,omitempty"`
	Id          string           `json:"id,omitempty"`
	Kind        string           `json:"kind,omitempty"`
	Label       string           `json:"label,omitempty"`
	ProjectId   int              `json:"project_id,omitempty"`
	Status      string           `json:"status,omitempty"`
	Subject     *ApprovalSubject `json:"subject,omitempty"`
	UpdatedAt   *time.Time       `json:"updated_at,omitempty"`
}

type ApprovalRequestExtended struct {
	ApprovedAt    *time.Time             `json:"approved_at,omitempty"`
	ApprovedBy    *User                  `json:"approved_by,omitempty"`
	CreatedAt     *time.Time             `json:"created_at,omitempty"`
	DenialComment string                 `json:"denial_comment,omitempty"`
	DeniedAt      *time.Time             `json:"denied_at,omitempty"`
	DeniedBy      *User                  `json:"denied_by,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Href          string                 `json:"href,omitempty"`
	Id            string                 `json:"id,omitempty"`
	Kind          string                 `json:"kind,omitempty"`
	Label         string                 `json:"label,omitempty"`
	Options       []*ConfigurationOption `json:"options,omitempty"`
	Parameters    map[string]interface{} `json:"parameters,omitempty"`
	ProjectId     int                    `json:"project_id,omitempty"`
	Status        string                 `json:"status,omitempty"`
	Subject       *ApprovalSubject       `json:"subject,omitempty"`
	UpdatedAt     *time.Time             `json:"updated_at,omitempty"`
}

type ApprovalRequestList struct {
	Count       int                `json:"count,omitempty"`
	Items       []*ApprovalRequest `json:"items,omitempty"`
	Kind        string             `json:"kind,omitempty"`
	NotModified string             `json:"not_modified,omitempty"`
}

type ApprovalSubject struct {
	Href string `json:"href,omitempty"`
	Kind string `json:"kind,omitempty"`
}

type ArchivedIncidentList struct {
	Count       int                 `json:"count,omitempty"`
	Items       []*ArchivedIncident `json:"items,omitempty"`
	Kind        string              `json:"kind,omitempty"`
	NotModified string              `json:"not_modified,omitempty"`
}

type ConfigurationOption struct {
	Label  string      `json:"label,omitempty"`
	Name   string      `json:"name,omitempty"`
	NoEcho bool        `json:"no_echo,omitempty"`
	Type   string      `json:"type,omitempty"`
	Value  interface{} `json:"value,omitempty"`
}

type ConfigurationOptionCreateType struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Escalation struct {
	Actions []*EscalationAction `json:"actions,omitempty"`
	Name    string              `json:"name,omitempty"`
	Status  string              `json:"status,omitempty"`
}

type EscalationAction struct {
	ApprovalRequest *ApprovalRequestExtended `json:"approval_request,omitempty"`
	Error           string                   `json:"error,omitempty"`
	FinishedAt      *time.Time               `json:"finished_at,omitempty"`
	ProcessHref     string                   `json:"process_href,omitempty"`
	StartedAt       *time.Time               `json:"started_at,omitempty"`
	Status          string                   `json:"status,omitempty"`
	Type            string                   `json:"type,omitempty"`
}

type Escalations struct {
	Escalations []*Escalation `json:"escalations,omitempty"`
	FinishedAt  *time.Time    `json:"finished_at,omitempty"`
	StartedAt   *time.Time    `json:"started_at,omitempty"`
	Status      string        `json:"status,omitempty"`
}

type IncidentAggregateIndex struct {
	ActionSummary   *ActionSummary       `json:"action_summary,omitempty"`
	Category        string               `json:"category,omitempty"`
	Count           int                  `json:"count,omitempty"`
	CreatedAt       *time.Time           `json:"created_at,omitempty"`
	DryRun          bool                 `json:"dry_run,omitempty"`
	Etag            string               `json:"etag,omitempty"`
	Href            string               `json:"href,omitempty"`
	Id              string               `json:"id,omitempty"`
	IncidentSummary *IncidentSummary     `json:"incident_summary,omitempty"`
	Kind            string               `json:"kind,omitempty"`
	NotModified     string               `json:"not_modified,omitempty"`
	Org             *Org                 `json:"org,omitempty"`
	PolicyAggregate *PolicyAggregateLink `json:"policy_aggregate,omitempty"`
	Severity        string               `json:"severity,omitempty"`
	State           string               `json:"state,omitempty"`
	UpdatedAt       *time.Time           `json:"updated_at,omitempty"`
}

type IncidentAggregateItem struct {
	ActionFailed       bool       `json:"action_failed,omitempty"`
	ActionPending      bool       `json:"action_pending,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	Id                 string     `json:"id,omitempty"`
	Kind               string     `json:"kind,omitempty"`
	Project            *Project   `json:"project,omitempty"`
	ResolutionMessage  string     `json:"resolution_message,omitempty"`
	ResolvedAt         *time.Time `json:"resolved_at,omitempty"`
	ResolvedBy         *User      `json:"resolved_by,omitempty"`
	State              string     `json:"state,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	Url                string     `json:"url,omitempty"`
	ViolationDataCount int        `json:"violation_data_count,omitempty"`
}

type IncidentAggregateList struct {
	Count       int                       `json:"count,omitempty"`
	Items       []*IncidentAggregateIndex `json:"items,omitempty"`
	Kind        string                    `json:"kind,omitempty"`
	NotModified string                    `json:"not_modified,omitempty"`
}

type IncidentAggregateNonCatalog struct {
	ActionSummary   *ActionSummary                     `json:"action_summary,omitempty"`
	Count           int                                `json:"count,omitempty"`
	Href            string                             `json:"href,omitempty"`
	IncidentSummary *IncidentSummary                   `json:"incident_summary,omitempty"`
	Items           []*IncidentAggregateNonCatalogItem `json:"items,omitempty"`
	Kind            string                             `json:"kind,omitempty"`
	NotModified     string                             `json:"not_modified,omitempty"`
	PolicyAggregate *PolicyAggregateNonCatalogLink     `json:"policy_aggregate,omitempty"`
	UpdatedAt       *time.Time                         `json:"updated_at,omitempty"`
}

type IncidentAggregateNonCatalogItem struct {
	ActionFailed       bool               `json:"action_failed,omitempty"`
	ActionPending      bool               `json:"action_pending,omitempty"`
	AppliedPolicy      *AppliedPolicyLink `json:"applied_policy,omitempty"`
	Category           string             `json:"category,omitempty"`
	CreatedAt          *time.Time         `json:"created_at,omitempty"`
	DryRun             bool               `json:"dry_run,omitempty"`
	Id                 string             `json:"id,omitempty"`
	Kind               string             `json:"kind,omitempty"`
	Project            *Project           `json:"project,omitempty"`
	ResolutionMessage  string             `json:"resolution_message,omitempty"`
	ResolvedAt         *time.Time         `json:"resolved_at,omitempty"`
	ResolvedBy         *User              `json:"resolved_by,omitempty"`
	Severity           string             `json:"severity,omitempty"`
	State              string             `json:"state,omitempty"`
	UpdatedAt          *time.Time         `json:"updated_at,omitempty"`
	Url                string             `json:"url,omitempty"`
	ViolationDataCount int                `json:"violation_data_count,omitempty"`
}

type IncidentList struct {
	Count       int         `json:"count,omitempty"`
	Items       []*Incident `json:"items,omitempty"`
	Kind        string      `json:"kind,omitempty"`
	NotModified string      `json:"not_modified,omitempty"`
}

type IncidentRunAction struct {
	Options []*ConfigurationOptionCreateType `json:"options,omitempty"`
}

type IncidentSummary struct {
	IncidentCount      int `json:"incident_count,omitempty"`
	ResolvedCount      int `json:"resolved_count,omitempty"`
	TriggeredCount     int `json:"triggered_count,omitempty"`
	ViolationDataCount int `json:"violation_data_count,omitempty"`
}

type Org struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PolicyAggregateCreate struct {
	AllProjects   bool                             `json:"all_projects,omitempty"`
	Credentials   map[string]interface{}           `json:"credentials,omitempty"`
	Description   string                           `json:"description,omitempty"`
	DryRun        bool                             `json:"dry_run,omitempty"`
	Frequency     string                           `json:"frequency,omitempty"`
	Name          string                           `json:"name,omitempty"`
	Options       []*ConfigurationOptionCreateType `json:"options,omitempty"`
	ProjectIds    []int                            `json:"project_ids,omitempty"`
	Severity      string                           `json:"severity,omitempty"`
	SkipApprovals bool                             `json:"skip_approvals,omitempty"`
	TemplateHref  string                           `json:"template_href,omitempty"`
}

type PolicyAggregateIndex struct {
	ActiveCount           int                    `json:"active_count,omitempty"`
	Category              string                 `json:"category,omitempty"`
	Count                 int                    `json:"count,omitempty"`
	CreatedAt             *time.Time             `json:"created_at,omitempty"`
	CreatedBy             *User                  `json:"created_by,omitempty"`
	Credentials           map[string]interface{} `json:"credentials,omitempty"`
	Description           string                 `json:"description,omitempty"`
	DocLink               string                 `json:"doc_link,omitempty"`
	DryRun                bool                   `json:"dry_run,omitempty"`
	ErrorCount            int                    `json:"error_count,omitempty"`
	Frequency             string                 `json:"frequency,omitempty"`
	Href                  string                 `json:"href,omitempty"`
	Id                    string                 `json:"id,omitempty"`
	IncidentAggregateHref string                 `json:"incident_aggregate_href,omitempty"`
	Info                  map[string]interface{} `json:"info,omitempty"`
	Kind                  string                 `json:"kind,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	Options               []*ConfigurationOption `json:"options,omitempty"`
	Org                   *Org                   `json:"org,omitempty"`
	PublishedTemplate     *PublishedTemplateLink `json:"published_template,omitempty"`
	RunningProjectIds     []int                  `json:"running_project_ids,omitempty"`
	Severity              string                 `json:"severity,omitempty"`
	SkipApprovals         bool                   `json:"skip_approvals,omitempty"`
	Status                string                 `json:"status,omitempty"`
	UpdatedAt             *time.Time             `json:"updated_at,omitempty"`
}

type PolicyAggregateItem struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Error     string     `json:"error,omitempty"`
	ErroredAt *time.Time `json:"errored_at,omitempty"`
	Id        string     `json:"id,omitempty"`
	Kind      string     `json:"kind,omitempty"`
	Project   *Project   `json:"project,omitempty"`
	Status    string     `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Url       string     `json:"url,omitempty"`
}

type PolicyAggregateLink struct {
	CreatedAt             *time.Time             `json:"created_at,omitempty"`
	CreatedBy             *User                  `json:"created_by,omitempty"`
	Href                  string                 `json:"href,omitempty"`
	Id                    string                 `json:"id,omitempty"`
	IncidentAggregateHref string                 `json:"incident_aggregate_href,omitempty"`
	Kind                  string                 `json:"kind,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	PublishedTemplate     *PublishedTemplateLink `json:"published_template,omitempty"`
}

type PolicyAggregateList struct {
	Count       int                     `json:"count,omitempty"`
	Items       []*PolicyAggregateIndex `json:"items,omitempty"`
	Kind        string                  `json:"kind,omitempty"`
	NotModified string                  `json:"not_modified,omitempty"`
}

type PolicyAggregateNonCatalog struct {
	ActiveCount           int                              `json:"active_count,omitempty"`
	Count                 int                              `json:"count,omitempty"`
	ErrorCount            int                              `json:"error_count,omitempty"`
	Href                  string                           `json:"href,omitempty"`
	IncidentAggregateHref string                           `json:"incident_aggregate_href,omitempty"`
	Items                 []*PolicyAggregateNonCatalogItem `json:"items,omitempty"`
	Kind                  string                           `json:"kind,omitempty"`
	RunningProjectIds     []int                            `json:"running_project_ids,omitempty"`
}

type PolicyAggregateNonCatalogItem struct {
	Category       string              `json:"category,omitempty"`
	CreatedAt      *time.Time          `json:"created_at,omitempty"`
	CreatedBy      *User               `json:"created_by,omitempty"`
	Description    string              `json:"description,omitempty"`
	DryRun         bool                `json:"dry_run,omitempty"`
	Error          string              `json:"error,omitempty"`
	ErroredAt      *time.Time          `json:"errored_at,omitempty"`
	Frequency      string              `json:"frequency,omitempty"`
	Id             string              `json:"id,omitempty"`
	Kind           string              `json:"kind,omitempty"`
	Name           string              `json:"name,omitempty"`
	PolicyTemplate *PolicyTemplateLink `json:"policy_template,omitempty"`
	Project        *Project            `json:"project,omitempty"`
	Severity       string              `json:"severity,omitempty"`
	Status         string              `json:"status,omitempty"`
	UpdatedAt      *time.Time          `json:"updated_at,omitempty"`
	Url            string              `json:"url,omitempty"`
}

type PolicyAggregateNonCatalogLink struct {
	Href                  string `json:"href,omitempty"`
	IncidentAggregateHref string `json:"incident_aggregate_href,omitempty"`
	Kind                  string `json:"kind,omitempty"`
}

type PolicyTemplateCompile struct {
	Filename string            `json:"filename,omitempty"`
	Source   *rsapi.FileUpload `json:"source,omitempty"`
}

type PolicyTemplateLink struct {
	Fingerprint string     `json:"fingerprint,omitempty"`
	Href        string     `json:"href,omitempty"`
	Id          string     `json:"id,omitempty"`
	Kind        string     `json:"kind,omitempty"`
	Name        string     `json:"name,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	UpdatedBy   *User      `json:"updated_by,omitempty"`
}

type PolicyTemplateList struct {
	Count       int               `json:"count,omitempty"`
	Items       []*PolicyTemplate `json:"items,omitempty"`
	Kind        string            `json:"kind,omitempty"`
	NotModified string            `json:"not_modified,omitempty"`
}

type PolicyTemplateRetrieveData struct {
	Credentials map[string]interface{}           `json:"credentials,omitempty"`
	Names       []string                         `json:"names,omitempty"`
	Options     []*ConfigurationOptionCreateType `json:"options,omitempty"`
}

type PolicyTemplateUpdate struct {
	Filename string            `json:"filename,omitempty"`
	Source   *rsapi.FileUpload `json:"source,omitempty"`
}

type PolicyTemplateUpload struct {
	Filename string            `json:"filename,omitempty"`
	Source   *rsapi.FileUpload `json:"source,omitempty"`
}

type Project struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	OrgId   int    `json:"org_id,omitempty"`
	OrgName string `json:"org_name,omitempty"`
}

type PublishedTemplateCreate struct {
	TemplateHref string `json:"template_href,omitempty"`
}

type PublishedTemplateLink struct {
	BuiltIn     bool       `json:"built_in,omitempty"`
	Fingerprint string     `json:"fingerprint,omitempty"`
	Href        string     `json:"href,omitempty"`
	Id          string     `json:"id,omitempty"`
	Kind        string     `json:"kind,omitempty"`
	Name        string     `json:"name,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	UpdatedBy   *User      `json:"updated_by,omitempty"`
}

type PublishedTemplateList struct {
	Count       int                  `json:"count,omitempty"`
	Items       []*PublishedTemplate `json:"items,omitempty"`
	Kind        string               `json:"kind,omitempty"`
	NotModified string               `json:"not_modified,omitempty"`
}

type PublishedTemplateUpdate struct {
	TemplateHref string `json:"template_href,omitempty"`
}

type Resolution struct {
	Actions []*ResolutionAction `json:"actions,omitempty"`
	Name    string              `json:"name,omitempty"`
	Status  string              `json:"status,omitempty"`
}

type ResolutionAction struct {
	ApprovalRequest *ApprovalRequestExtended `json:"approval_request,omitempty"`
	Error           string                   `json:"error,omitempty"`
	FinishedAt      *time.Time               `json:"finished_at,omitempty"`
	ProcessHref     string                   `json:"process_href,omitempty"`
	StartedAt       *time.Time               `json:"started_at,omitempty"`
	Status          string                   `json:"status,omitempty"`
	Type            string                   `json:"type,omitempty"`
}

type Resolutions struct {
	FinishedAt  *time.Time    `json:"finished_at,omitempty"`
	Resolutions []*Resolution `json:"resolutions,omitempty"`
	StartedAt   *time.Time    `json:"started_at,omitempty"`
	Status      string        `json:"status,omitempty"`
}

type User struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}
