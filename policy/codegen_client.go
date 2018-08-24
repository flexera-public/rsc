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

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
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

/****** AppliedPolicyService ******/

// Show retrieves the details of an applied policy.
type AppliedPolicyService struct {
}

//===== Locator

// AppliedPolicyServiceLocator exposes the AppliedPolicyService resource actions.
type AppliedPolicyServiceLocator struct {
	Href
	api *API
}

// AppliedPolicyServiceLocator builds a locator from the given href.
func (api *API) AppliedPolicyServiceLocator(href string) *AppliedPolicyServiceLocator {
	return &AppliedPolicyServiceLocator{Href(href), api}
}

//===== Actions

// get /api/governance/projects/{project_id}/applied_policies
//
// Index retrieves the list of applied policies in a project.
func (loc *AppliedPolicyServiceLocator) Index(options rsapi.APIParams) (*AppliedPolicyList, error) {
	var res *AppliedPolicyList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicyService", "index")
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

// post /api/governance/projects/{project_id}/applied_policies
//
// Create applies a policy template to a given project. The applied policy will continually run until deleted.
func (loc *AppliedPolicyServiceLocator) Create(name string, templateHref string, options rsapi.APIParams) (*AppliedPolicy, error) {
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
	uri, err := loc.ActionPath("AppliedPolicyService", "create")
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

// get /api/governance/projects/{project_id}/applied_policies/{policy_id}
//
// Show retrieves the details of an applied policy.
func (loc *AppliedPolicyServiceLocator) Show(options rsapi.APIParams) (*AppliedPolicy, error) {
	var res *AppliedPolicy
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicyService", "show")
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

// delete /api/governance/projects/{project_id}/applied_policies/{policy_id}
//
// Delete stops and deletes an applied policy.
func (loc *AppliedPolicyServiceLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicyService", "delete")
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

// post /api/governance/projects/{project_id}/applied_policies/{policy_id}/evaluate
//
// Evaluate executes an applied policy evaluation on demand. It does not affect the normal execution schedule.
func (loc *AppliedPolicyServiceLocator) Evaluate() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicyService", "evaluate")
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

// get /api/governance/projects/{project_id}/applied_policies/{policy_id}/log
//
// ShowLog retrieves the last evaluation log of an applied policy.
func (loc *AppliedPolicyServiceLocator) ShowLog() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicyService", "show_log")
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

// get /api/governance/projects/{project_id}/applied_policies/{policy_id}/status
//
// ShowStatus retrieves the evaluation status details of an applied policy.
func (loc *AppliedPolicyServiceLocator) ShowStatus() (*AppliedPolicyStatus, error) {
	var res *AppliedPolicyStatus
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AppliedPolicyService", "show_status")
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

/****** IncidentService ******/

// Show retrieves the details of an incident.
type IncidentService struct {
}

//===== Locator

// IncidentServiceLocator exposes the IncidentService resource actions.
type IncidentServiceLocator struct {
	Href
	api *API
}

// IncidentServiceLocator builds a locator from the given href.
func (api *API) IncidentServiceLocator(href string) *IncidentServiceLocator {
	return &IncidentServiceLocator{Href(href), api}
}

//===== Actions

// get /api/governance/projects/{project_id}/incidents
//
// Index retrieves the list of incidents in a project.
func (loc *IncidentServiceLocator) Index(options rsapi.APIParams) (*IncidentList, error) {
	var res *IncidentList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentService", "index")
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

// get /api/governance/projects/{project_id}/incidents/{incident_id}
//
// Show retrieves the details of an incident.
func (loc *IncidentServiceLocator) Show(options rsapi.APIParams) (*Incident, error) {
	var res *Incident
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentService", "show")
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

// get /api/governance/projects/{project_id}/incidents/{incident_id}/escalations
//
// IndexEscalations retrieves the status details of all of the escalations for an incident.
func (loc *IncidentServiceLocator) IndexEscalations() (*Escalations, error) {
	var res *Escalations
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentService", "index_escalations")
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

// get /api/governance/projects/{project_id}/incidents/{incident_id}/resolutions
//
// IndexResolutions retrieves the status details of all of the resolutions for an incident.
func (loc *IncidentServiceLocator) IndexResolutions() (*Resolutions, error) {
	var res *Resolutions
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentService", "index_resolutions")
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

// put /api/governance/projects/{project_id}/incidents/{incident_id}/resolve
//
// Resolve resolves an incident by setting it to an inactive state, indicating that it has been addressed.
func (loc *IncidentServiceLocator) Resolve() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IncidentService", "resolve")
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

/****** PolicyTemplateService ******/

// Show retrieves the details of a policy template.
type PolicyTemplateService struct {
}

//===== Locator

// PolicyTemplateServiceLocator exposes the PolicyTemplateService resource actions.
type PolicyTemplateServiceLocator struct {
	Href
	api *API
}

// PolicyTemplateServiceLocator builds a locator from the given href.
func (api *API) PolicyTemplateServiceLocator(href string) *PolicyTemplateServiceLocator {
	return &PolicyTemplateServiceLocator{Href(href), api}
}

//===== Actions

// get /api/governance/projects/{project_id}/policy_templates
//
// IndexPolicyTemplates retrieves the list of policy templates in a project.
func (loc *PolicyTemplateServiceLocator) Index(options rsapi.APIParams) (*PolicyTemplateList, error) {
	var res *PolicyTemplateList
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyTemplateService", "index")
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

// post /api/governance/projects/{project_id}/policy_templates
//
// Upload uploads a policy template for a project, first compiling it. On failure, an array of syntax errors will be returned.
func (loc *PolicyTemplateServiceLocator) Upload(filename string, source *rsapi.FileUpload) (*PolicyTemplate, error) {
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
	uri, err := loc.ActionPath("PolicyTemplateService", "upload")
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

// post /api/governance/projects/{project_id}/policy_templates/compile
//
// Compile compiles a policy template for a project. This is only to be used for checking the syntax of a policy template; the results are not stored.
func (loc *PolicyTemplateServiceLocator) Compile(filename string, source *rsapi.FileUpload) error {
	if filename == "" {
		return fmt.Errorf("filename is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"filename": filename,
		"source":   source,
	}
	uri, err := loc.ActionPath("PolicyTemplateService", "compile")
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

// get /api/governance/projects/{project_id}/policy_templates/{template_id}
//
// Show retrieves the details of a policy template.
func (loc *PolicyTemplateServiceLocator) Show(options rsapi.APIParams) (*PolicyTemplate, error) {
	var res *PolicyTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyTemplateService", "show")
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

// put /api/governance/projects/{project_id}/policy_templates/{template_id}
//
// Update updates a policy template in place for a project, by replacing it. Any existing applied policies using the template will not be updated; they must be deleted and recreated again.
func (loc *PolicyTemplateServiceLocator) Update(filename string, source *rsapi.FileUpload) (*PolicyTemplate, error) {
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
	uri, err := loc.ActionPath("PolicyTemplateService", "update")
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

// delete /api/governance/projects/{project_id}/policy_templates/{template_id}
//
// Delete deletes a policy template from a project. Deleting a policy template will not delete any applied policies created from the template, they must be stopped explicitly.
func (loc *PolicyTemplateServiceLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PolicyTemplateService", "delete")
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

/****** PublishedTemplateService ******/

// Show retrieves the details of a published template.
type PublishedTemplateService struct {
}

//===== Locator

// PublishedTemplateServiceLocator exposes the PublishedTemplateService resource actions.
type PublishedTemplateServiceLocator struct {
	Href
	api *API
}

// PublishedTemplateServiceLocator builds a locator from the given href.
func (api *API) PublishedTemplateServiceLocator(href string) *PublishedTemplateServiceLocator {
	return &PublishedTemplateServiceLocator{Href(href), api}
}

//===== Actions

// get /api/governance/orgs/{org_id}/published_templates
//
// Index retrieves the list of published templates in an organization.
func (loc *PublishedTemplateServiceLocator) Index(options rsapi.APIParams) (*PublishedTemplateList, error) {
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
	uri, err := loc.ActionPath("PublishedTemplateService", "index")
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

// post /api/governance/orgs/{org_id}/published_templates
//
// Create creates an organization-scoped published template from a project-scoped policy template.
func (loc *PublishedTemplateServiceLocator) Create(templateHref string) error {
	if templateHref == "" {
		return fmt.Errorf("templateHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"template_href": templateHref,
	}
	uri, err := loc.ActionPath("PublishedTemplateService", "create")
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

// get /api/governance/orgs/{org_id}/published_templates/{template_id}
//
// Show retrieves the details of a published template.
func (loc *PublishedTemplateServiceLocator) Show(options rsapi.APIParams) (*PublishedTemplate, error) {
	var res *PublishedTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplateService", "show")
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

// put /api/governance/orgs/{org_id}/published_templates/{template_id}
//
// Update updates a published template in place for an organization, by replacing it. Any existing applied policies using the template will not be updated; they must be deleted and recreated again.
func (loc *PublishedTemplateServiceLocator) Update(templateHref string) error {
	if templateHref == "" {
		return fmt.Errorf("templateHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"template_href": templateHref,
	}
	uri, err := loc.ActionPath("PublishedTemplateService", "update")
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

// delete /api/governance/orgs/{org_id}/published_templates/{template_id}
//
// Delete deletes a published template from an organization. Deleting a published template will not delete any applied policies created from the template, they must be stopped explicitly.
func (loc *PublishedTemplateServiceLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplateService", "delete")
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

// post /api/governance/orgs/{org_id}/published_templates/{template_id}/hide
//
// Hide hides a RightScale built-in template from an organization.
func (loc *PublishedTemplateServiceLocator) Hide() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplateService", "hide")
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

// post /api/governance/orgs/{org_id}/published_templates/{template_id}/unhide
//
// Unhide unhides a RightScale built-in template from an organization.
func (loc *PublishedTemplateServiceLocator) Unhide() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublishedTemplateService", "unhide")
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

type AppliedPolicy struct {
	CreatedAt         *time.Time         `json:"created_at,omitempty"`
	CreatedBy         *User              `json:"created_by,omitempty"`
	Frequency         string             `json:"frequency,omitempty"`
	Href              string             `json:"href,omitempty"`
	Id                string             `json:"id,omitempty"`
	Kind              string             `json:"kind,omitempty"`
	Name              string             `json:"name,omitempty"`
	PolicyTemplate    *PolicyTemplate    `json:"policy_template,omitempty"`
	PublishedTemplate *PublishedTemplate `json:"published_template,omitempty"`
}

type AppliedPolicyCreate struct {
	Description  string                           `json:"description,omitempty"`
	DryRun       bool                             `json:"dry_run,omitempty"`
	Frequency    string                           `json:"frequency,omitempty"`
	Name         string                           `json:"name,omitempty"`
	Options      []*ConfigurationOptionCreateType `json:"options,omitempty"`
	TemplateHref string                           `json:"template_href,omitempty"`
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

type ConfigurationOption struct {
	Label string      `json:"label,omitempty"`
	Name  string      `json:"name,omitempty"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
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
	Error      string     `json:"error,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	Status     string     `json:"status,omitempty"`
	Type       string     `json:"type,omitempty"`
}

type Escalations struct {
	Escalations []*Escalation `json:"escalations,omitempty"`
	FinishedAt  *time.Time    `json:"finished_at,omitempty"`
	StartedAt   *time.Time    `json:"started_at,omitempty"`
	Status      string        `json:"status,omitempty"`
}

type Incident struct {
	ActionFailed       bool                   `json:"action_failed,omitempty"`
	AppliedPolicy      *AppliedPolicy         `json:"applied_policy,omitempty"`
	Category           string                 `json:"category,omitempty"`
	CreatedAt          *time.Time             `json:"created_at,omitempty"`
	DryRun             bool                   `json:"dry_run,omitempty"`
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

type IncidentList struct {
	Count       int         `json:"count,omitempty"`
	Items       []*Incident `json:"items,omitempty"`
	Kind        string      `json:"kind,omitempty"`
	NotModified string      `json:"not_modified,omitempty"`
}

type PolicyTemplate struct {
	Category         string                 `json:"category,omitempty"`
	CreatedAt        *time.Time             `json:"created_at,omitempty"`
	CreatedBy        *User                  `json:"created_by,omitempty"`
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
	UpdatedAt        *time.Time             `json:"updated_at,omitempty"`
	UpdatedBy        *User                  `json:"updated_by,omitempty"`
}

type PolicyTemplateCompile struct {
	Filename string            `json:"filename,omitempty"`
	Source   *rsapi.FileUpload `json:"source,omitempty"`
}

type PolicyTemplateList struct {
	Count       int               `json:"count,omitempty"`
	Items       []*PolicyTemplate `json:"items,omitempty"`
	Kind        string            `json:"kind,omitempty"`
	NotModified string            `json:"not_modified,omitempty"`
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

type PublishedTemplate struct {
	BuiltIn     bool       `json:"built_in,omitempty"`
	Fingerprint string     `json:"fingerprint,omitempty"`
	Href        string     `json:"href,omitempty"`
	Id          string     `json:"id,omitempty"`
	Kind        string     `json:"kind,omitempty"`
	Name        string     `json:"name,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	UpdatedBy   *User      `json:"updated_by,omitempty"`
}

type PublishedTemplateCreate struct {
	TemplateHref string `json:"template_href,omitempty"`
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
	Error      string     `json:"error,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	Status     string     `json:"status,omitempty"`
	Type       string     `json:"type,omitempty"`
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
