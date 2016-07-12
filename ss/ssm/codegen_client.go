//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=ss/ssm/restful_doc -output=ss/ssm -pkg=ssm -target=1.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssm

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

/******  Execution ******/

// An Execution is a launched instance of a CloudApp. Executions can be created from the catalog
// by launching an Application, from Designer by launching a Template, or directly in Manager
// by using the API and sending the CAT source or CAT Compiled source.
// Executions are represented in RightScale Cloud Management by a deployment -- the resources
// defined in the CAT are all created in the Deployment. Any action on a running CloudApp should
// be made on its Execution resource.
// Making changes to any resource directly in the CM deployment
// may result in undesired behavior since the Execution only refreshes certain information as a
// result of running an Operation on an Execution. For example, if a Server is replaced in CM
// instead of through Self-Service, the new Server's information won' be available in
// Self-Service.
type Execution struct {
	ApiResources            []*Resource            `json:"api_resources,omitempty"`
	AvailableActions        []string               `json:"available_actions,omitempty"`
	AvailableOperations     []*OperationDefinition `json:"available_operations,omitempty"`
	AvailableOperationsInfo []*OperationInfo       `json:"available_operations_info,omitempty"`
	CompilationHref         string                 `json:"compilation_href,omitempty"`
	ConfigurationOptions    []*ConfigurationOption `json:"configuration_options,omitempty"`
	Cost                    *CostStruct            `json:"cost,omitempty"`
	CreatedBy               *User                  `json:"created_by,omitempty"`
	CurrentSchedule         string                 `json:"current_schedule,omitempty"`
	Dependencies            []*CatDependency       `json:"dependencies,omitempty"`
	Deployment              string                 `json:"deployment,omitempty"`
	DeploymentUrl           string                 `json:"deployment_url,omitempty"`
	Description             string                 `json:"description,omitempty"`
	EndsAt                  *time.Time             `json:"ends_at,omitempty"`
	Href                    string                 `json:"href,omitempty"`
	Id                      string                 `json:"id,omitempty"`
	Kind                    string                 `json:"kind,omitempty"`
	LatestNotification      *Notification          `json:"latest_notification,omitempty"`
	LatestNotifications     []*Notification        `json:"latest_notifications,omitempty"`
	LaunchedFrom            *LaunchedFrom          `json:"launched_from,omitempty"`
	LaunchedFromSummary     map[string]interface{} `json:"launched_from_summary,omitempty"`
	Links                   *ExecutionLinks        `json:"links,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	NextAction              *ScheduledAction       `json:"next_action,omitempty"`
	Outputs                 []*Output              `json:"outputs,omitempty"`
	RunningOperations       []*Operation           `json:"running_operations,omitempty"`
	ScheduleRequired        bool                   `json:"schedule_required,omitempty"`
	Scheduled               bool                   `json:"scheduled,omitempty"`
	Schedules               []*Schedule            `json:"schedules,omitempty"`
	Source                  string                 `json:"source,omitempty"`
	Status                  string                 `json:"status,omitempty"`
	Timestamps              *TimestampsStruct      `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Execution) Locator(api *API) *ExecutionLocator {
	return api.ExecutionLocator(r.Href)
}

//===== Locator

// ExecutionLocator exposes the Execution resource actions.
type ExecutionLocator struct {
	Href
	api *API
}

// ExecutionLocator builds a locator from the given href.
func (api *API) ExecutionLocator(href string) *ExecutionLocator {
	return &ExecutionLocator{Href(href), api}
}

//===== Actions

// GET /api/manager/projects/:project_id/executions
//
// List information about the Executions, or use a filter to only return certain Executions. A view can be used for various levels of detail.
func (loc *ExecutionLocator) Index(options rsapi.APIParams) ([]*Execution, error) {
	var res []*Execution
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids[]"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "index")
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

// GET /api/manager/projects/:project_id/executions/:id
//
// Show details for a given Execution. A view can be used for various levels of detail.
func (loc *ExecutionLocator) Show(options rsapi.APIParams) (*Execution, error) {
	var res *Execution
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "show")
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

// POST /api/manager/projects/:project_id/executions
//
// Create a new execution from a CAT, a compiled CAT, an Application in the Catalog, or a Template in Designer
func (loc *ExecutionLocator) Create(options rsapi.APIParams) (*ExecutionLocator, error) {
	var res *ExecutionLocator
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var applicationHrefOpt = options["application_href"]
	if applicationHrefOpt != nil {
		p["application_href"] = applicationHrefOpt
	}
	var compilationHrefOpt = options["compilation_href"]
	if compilationHrefOpt != nil {
		p["compilation_href"] = compilationHrefOpt
	}
	var compiledCatOpt = options["compiled_cat"]
	if compiledCatOpt != nil {
		p["compiled_cat"] = compiledCatOpt
	}
	var currentScheduleOpt = options["current_schedule"]
	if currentScheduleOpt != nil {
		p["current_schedule"] = currentScheduleOpt
	}
	var deferLaunchOpt = options["defer_launch"]
	if deferLaunchOpt != nil {
		p["defer_launch"] = deferLaunchOpt
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var endsAtOpt = options["ends_at"]
	if endsAtOpt != nil {
		p["ends_at"] = endsAtOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var options_Opt = options["options"]
	if options_Opt != nil {
		p["options"] = options_Opt
	}
	var scheduleRequiredOpt = options["schedule_required"]
	if scheduleRequiredOpt != nil {
		p["schedule_required"] = scheduleRequiredOpt
	}
	var scheduledActionsOpt = options["scheduled_actions"]
	if scheduledActionsOpt != nil {
		p["scheduled_actions"] = scheduledActionsOpt
	}
	var schedulesOpt = options["schedules"]
	if schedulesOpt != nil {
		p["schedules"] = schedulesOpt
	}
	var sourceOpt = options["source"]
	if sourceOpt != nil {
		p["source"] = sourceOpt
	}
	var templateHrefOpt = options["template_href"]
	if templateHrefOpt != nil {
		p["template_href"] = templateHrefOpt
	}
	uri, err := loc.ActionPath("Execution", "create")
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
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ExecutionLocator{Href(location), loc.api}, nil
	}
}

// PATCH /api/manager/projects/:project_id/executions/:id
//
// Updates an execution end date or selected schedule.
func (loc *ExecutionLocator) Patch(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var currentScheduleOpt = options["current_schedule"]
	if currentScheduleOpt != nil {
		p["current_schedule"] = currentScheduleOpt
	}
	var endsAtOpt = options["ends_at"]
	if endsAtOpt != nil {
		p["ends_at"] = endsAtOpt
	}
	uri, err := loc.ActionPath("Execution", "patch")
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

// DELETE /api/manager/projects/:project_id/executions/:id
//
// No description provided for delete.
func (loc *ExecutionLocator) Delete(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var forceOpt = options["force"]
	if forceOpt != nil {
		params["force"] = forceOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "delete")
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

// DELETE /api/manager/projects/:project_id/executions
//
// Delete several executions from the database. Note: if an execution has not successfully been terminated, there may still be associated cloud resources running.
func (loc *ExecutionLocator) MultiDelete(ids []string, options rsapi.APIParams) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var forceOpt = options["force"]
	if forceOpt != nil {
		params["force"] = forceOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "multi_delete")
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

// GET /api/manager/projects/:project_id/executions/:id/download
//
// Download the CAT source for the execution.
func (loc *ExecutionLocator) Download(apiVersion string) error {
	if apiVersion == "" {
		return fmt.Errorf("apiVersion is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"api_version": apiVersion,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "download")
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

// POST /api/manager/projects/:project_id/executions/:id/actions/launch
//
// Launch an Execution.
func (loc *ExecutionLocator) Launch() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "launch")
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

// POST /api/manager/projects/:project_id/executions/:id/actions/start
//
// Start an Execution.
func (loc *ExecutionLocator) Start() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "start")
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

// POST /api/manager/projects/:project_id/executions/:id/actions/stop
//
// Stop an Execution.
func (loc *ExecutionLocator) Stop() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "stop")
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

// POST /api/manager/projects/:project_id/executions/:id/actions/terminate
//
// Terminate an Execution.
func (loc *ExecutionLocator) Terminate() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "terminate")
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

// POST /api/manager/projects/:project_id/executions/actions/launch
//
// Launch several Executions.
func (loc *ExecutionLocator) MultiLaunch(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "multi_launch")
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

// POST /api/manager/projects/:project_id/executions/actions/start
//
// Start several Executions.
func (loc *ExecutionLocator) MultiStart(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "multi_start")
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

// POST /api/manager/projects/:project_id/executions/actions/stop
//
// Stop several Executions.
func (loc *ExecutionLocator) MultiStop(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "multi_stop")
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

// POST /api/manager/projects/:project_id/executions/actions/terminate
//
// Terminate several Executions.
func (loc *ExecutionLocator) MultiTerminate(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Execution", "multi_terminate")
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

// POST /api/manager/projects/:project_id/executions/:id/actions/run
//
// Runs an Operation on an Execution.
func (loc *ExecutionLocator) Run(name string, options rsapi.APIParams) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name": name,
	}
	var configurationOptionsOpt = options["configuration_options"]
	if configurationOptionsOpt != nil {
		p["configuration_options"] = configurationOptionsOpt
	}
	uri, err := loc.ActionPath("Execution", "run")
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

// POST /api/manager/projects/:project_id/executions/actions/run
//
// Runs an Operation on several Executions.
func (loc *ExecutionLocator) MultiRun(ids []string, name string, options rsapi.APIParams) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	if name == "" {
		return fmt.Errorf("name is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name": name,
	}
	var configurationOptionsOpt = options["configuration_options"]
	if configurationOptionsOpt != nil {
		p["configuration_options"] = configurationOptionsOpt
	}
	uri, err := loc.ActionPath("Execution", "multi_run")
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

/******  Notification ******/

// The Notification resource represents a system notification that an action has occurred. Generally
// these Notifications are the start and completion of Operations. Currently notifications are only
// available via the API/UI and are not distributed externally to users.
type Notification struct {
	Category   string             `json:"category,omitempty"`
	Execution  *Execution         `json:"execution,omitempty"`
	Href       string             `json:"href,omitempty"`
	Id         string             `json:"id,omitempty"`
	Kind       string             `json:"kind,omitempty"`
	Links      *NotificationLinks `json:"links,omitempty"`
	Message    string             `json:"message,omitempty"`
	Read       bool               `json:"read,omitempty"`
	Timestamps *TimestampsStruct  `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Notification) Locator(api *API) *NotificationLocator {
	return api.NotificationLocator(r.Href)
}

//===== Locator

// NotificationLocator exposes the Notification resource actions.
type NotificationLocator struct {
	Href
	api *API
}

// NotificationLocator builds a locator from the given href.
func (api *API) NotificationLocator(href string) *NotificationLocator {
	return &NotificationLocator{Href(href), api}
}

//===== Actions

// GET /api/manager/projects/:project_id/notifications
//
// List the most recent 50 Notifications. Use the filter parameter to specify specify Executions.
func (loc *NotificationLocator) Index(options rsapi.APIParams) ([]*Notification, error) {
	var res []*Notification
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids[]"] = idsOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Notification", "index")
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

// GET /api/manager/projects/:project_id/notifications/:id
//
// Get details for a specific Notification
func (loc *NotificationLocator) Show() (*Notification, error) {
	var res *Notification
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Notification", "show")
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

/******  Operation ******/

// Operations represent actions that can be taken on an Execution.
// When a CloudApp is launched, a sequence of Operations is run as [explained here](http://docs.rightscale.com/ss/reference/ss_CAT_file_language.html#operations) in the Operations section
// While a CloudApp is running, users may launch any custom Operations as defined in the CAT.
// Once a CAT is Terminated, a sequence of Operations is run as [explained here](http://docs.rightscale.com/ss/reference/ss_CAT_file_language.html#operations) in the Operations section
type Operation struct {
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	CreatedBy            *User                  `json:"created_by,omitempty"`
	Execution            *Execution             `json:"execution,omitempty"`
	Href                 string                 `json:"href,omitempty"`
	Id                   string                 `json:"id,omitempty"`
	Kind                 string                 `json:"kind,omitempty"`
	Label                string                 `json:"label,omitempty"`
	Links                *OperationLinks        `json:"links,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Status               *StatusStruct          `json:"status,omitempty"`
	Timestamps           *TimestampsStruct      `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Operation) Locator(api *API) *OperationLocator {
	return api.OperationLocator(r.Href)
}

//===== Locator

// OperationLocator exposes the Operation resource actions.
type OperationLocator struct {
	Href
	api *API
}

// OperationLocator builds a locator from the given href.
func (api *API) OperationLocator(href string) *OperationLocator {
	return &OperationLocator{Href(href), api}
}

//===== Actions

// GET /api/manager/projects/:project_id/operations
//
// Get the list of 50 most recent Operations (usually filtered by Execution).
func (loc *OperationLocator) Index(options rsapi.APIParams) ([]*Operation, error) {
	var res []*Operation
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids[]"] = idsOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Operation", "index")
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

// GET /api/manager/projects/:project_id/operations/:id
//
// Get the details for a specific Operation
func (loc *OperationLocator) Show(options rsapi.APIParams) (*Operation, error) {
	var res *Operation
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Operation", "show")
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

// POST /api/manager/projects/:project_id/operations
//
// Trigger an Operation to run by specifying the Execution ID and the name of the Operation.
func (loc *OperationLocator) Create(executionId string, name string, options rsapi.APIParams) (*OperationLocator, error) {
	var res *OperationLocator
	if executionId == "" {
		return res, fmt.Errorf("executionId is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"execution_id": executionId,
		"name":         name,
	}
	var options_Opt = options["options"]
	if options_Opt != nil {
		p["options"] = options_Opt
	}
	uri, err := loc.ActionPath("Operation", "create")
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
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &OperationLocator{Href(location), loc.api}, nil
	}
}

/******  ScheduledAction ******/

// ScheduledActions describe a set of timed occurrences for an action to be run (at most once per day).
// Recurrence Rules are based off of the [RFC 5545](https://tools.ietf.org/html/rfc5545) iCal spec, and timezones are from the standard [tzinfo database](http://www.iana.org/time-zones).
// All DateTimes must be passed in [ISO-8601 format](https://en.wikipedia.org/wiki/ISO_8601)
type ScheduledAction struct {
	Action                string                `json:"action,omitempty"`
	CreatedBy             *User                 `json:"created_by,omitempty"`
	Execution             *Execution            `json:"execution,omitempty"`
	ExecutionSchedule     bool                  `json:"execution_schedule,omitempty"`
	FirstOccurrence       *time.Time            `json:"first_occurrence,omitempty"`
	Href                  string                `json:"href,omitempty"`
	Id                    string                `json:"id,omitempty"`
	Kind                  string                `json:"kind,omitempty"`
	Links                 *ScheduledActionLinks `json:"links,omitempty"`
	Name                  string                `json:"name,omitempty"`
	NextOccurrence        *time.Time            `json:"next_occurrence,omitempty"`
	Operation             *OperationStruct      `json:"operation,omitempty"`
	Recurrence            string                `json:"recurrence,omitempty"`
	RecurrenceDescription string                `json:"recurrence_description,omitempty"`
	Timestamps            *TimestampsStruct     `json:"timestamps,omitempty"`
	Timezone              string                `json:"timezone,omitempty"`
}

// Locator returns a locator for the given resource
func (r *ScheduledAction) Locator(api *API) *ScheduledActionLocator {
	return api.ScheduledActionLocator(r.Href)
}

//===== Locator

// ScheduledActionLocator exposes the ScheduledAction resource actions.
type ScheduledActionLocator struct {
	Href
	api *API
}

// ScheduledActionLocator builds a locator from the given href.
func (api *API) ScheduledActionLocator(href string) *ScheduledActionLocator {
	return &ScheduledActionLocator{Href(href), api}
}

//===== Actions

// GET /api/manager/projects/:project_id/scheduled_actions
//
// List ScheduledAction resources in the project. The list can be filtered to a given execution.
func (loc *ScheduledActionLocator) Index(options rsapi.APIParams) ([]*ScheduledAction, error) {
	var res []*ScheduledAction
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ScheduledAction", "index")
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

// GET /api/manager/projects/:project_id/scheduled_actions/:id
//
// Retrieve given ScheduledAction resource.
func (loc *ScheduledActionLocator) Show() (*ScheduledAction, error) {
	var res *ScheduledAction
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ScheduledAction", "show")
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

// POST /api/manager/projects/:project_id/scheduled_actions
//
// Create a new ScheduledAction resource.
func (loc *ScheduledActionLocator) Create(action string, executionId string, firstOccurrence *time.Time, options rsapi.APIParams) (*ScheduledActionLocator, error) {
	var res *ScheduledActionLocator
	if action == "" {
		return res, fmt.Errorf("action is required")
	}
	if executionId == "" {
		return res, fmt.Errorf("executionId is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"action":           action,
		"execution_id":     executionId,
		"first_occurrence": firstOccurrence,
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var operationOpt = options["operation"]
	if operationOpt != nil {
		p["operation"] = operationOpt
	}
	var recurrenceOpt = options["recurrence"]
	if recurrenceOpt != nil {
		p["recurrence"] = recurrenceOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		p["timezone"] = timezoneOpt
	}
	uri, err := loc.ActionPath("ScheduledAction", "create")
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
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ScheduledActionLocator{Href(location), loc.api}, nil
	}
}

// PATCH /api/manager/projects/:project_id/scheduled_actions/:id
//
// Updates the 'next_occurrence' property of a ScheduledAction.
func (loc *ScheduledActionLocator) Patch(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var nextOccurrenceOpt = options["next_occurrence"]
	if nextOccurrenceOpt != nil {
		p["next_occurrence"] = nextOccurrenceOpt
	}
	uri, err := loc.ActionPath("ScheduledAction", "patch")
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

// DELETE /api/manager/projects/:project_id/scheduled_actions/:id
//
// Delete a ScheduledAction.
func (loc *ScheduledActionLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ScheduledAction", "delete")
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

// POST /api/manager/projects/:project_id/scheduled_actions/:id/actions/skip
//
// Skips the requested number of ScheduledAction occurrences. If no count is provided, one occurrence is skipped. On success, the next_occurrence view of the updated ScheduledAction is returned.
func (loc *ScheduledActionLocator) Skip(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var countOpt = options["count"]
	if countOpt != nil {
		p["count"] = countOpt
	}
	uri, err := loc.ActionPath("ScheduledAction", "skip")
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

/****** Parameter Data Types ******/

type ApiResourcesValueStruct struct {
	Details map[string]interface{} `json:"details,omitempty"`
	Href    string                 `json:"href,omitempty"`
}

type AvailableOperationsParametersUiStruct struct {
	Category string `json:"category,omitempty"`
	Index    int    `json:"index,omitempty"`
	Label    string `json:"label,omitempty"`
}

type AvailableOperationsParametersValidationStruct struct {
	AllowedPattern        string        `json:"allowed_pattern,omitempty"`
	AllowedValues         []interface{} `json:"allowed_values,omitempty"`
	ConstraintDescription string        `json:"constraint_description,omitempty"`
	MaxLength             int           `json:"max_length,omitempty"`
	MaxValue              int           `json:"max_value,omitempty"`
	MinLength             int           `json:"min_length,omitempty"`
	MinValue              int           `json:"min_value,omitempty"`
	NoEcho                bool          `json:"no_echo,omitempty"`
}

type CatDependency struct {
	Alias      string `json:"alias,omitempty"`
	Name       string `json:"name,omitempty"`
	Package    string `json:"package,omitempty"`
	SourceHref string `json:"source_href,omitempty"`
}

type CompiledCAT struct {
	CatParserGemVersion string                   `json:"cat_parser_gem_version,omitempty"`
	CompilerVer         string                   `json:"compiler_ver,omitempty"`
	Conditions          map[string]interface{}   `json:"conditions,omitempty"`
	Definitions         map[string]interface{}   `json:"definitions,omitempty"`
	DependencyHashes    []map[string]interface{} `json:"dependency_hashes,omitempty"`
	Imports             map[string]interface{}   `json:"imports,omitempty"`
	LongDescription     string                   `json:"long_description,omitempty"`
	Mappings            map[string]interface{}   `json:"mappings,omitempty"`
	Name                string                   `json:"name,omitempty"`
	Namespaces          []*Namespace             `json:"namespaces,omitempty"`
	Operations          map[string]interface{}   `json:"operations,omitempty"`
	Outputs             map[string]interface{}   `json:"outputs,omitempty"`
	Package             string                   `json:"package,omitempty"`
	Parameters          map[string]interface{}   `json:"parameters,omitempty"`
	Permissions         map[string]interface{}   `json:"permissions,omitempty"`
	RequiredParameters  []string                 `json:"required_parameters,omitempty"`
	Resources           map[string]interface{}   `json:"resources,omitempty"`
	RsCaVer             int                      `json:"rs_ca_ver,omitempty"`
	ShortDescription    string                   `json:"short_description,omitempty"`
	Source              string                   `json:"source,omitempty"`
}

type ConfigurationOption struct {
	Name  string      `json:"name,omitempty"`
	Type_ string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type CostStruct struct {
	Unit      string     `json:"unit,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Value     string     `json:"value,omitempty"`
}

type ExecutionLink struct {
	Href string `json:"href,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ExecutionLinks struct {
	LatestNotifications *NotificationLatestNotificationsLink `json:"latest_notifications,omitempty"`
	RunningOperations   *OperationRunningOperationsLink      `json:"running_operations,omitempty"`
}

type ExecutionParam struct {
	ApiResources            []*Resource                                  `json:"api_resources,omitempty"`
	AvailableActions        []string                                     `json:"available_actions,omitempty"`
	AvailableOperations     []*OperationDefinition                       `json:"available_operations,omitempty"`
	AvailableOperationsInfo []*OperationInfo                             `json:"available_operations_info,omitempty"`
	CompilationHref         string                                       `json:"compilation_href,omitempty"`
	ConfigurationOptions    []*ConfigurationOption                       `json:"configuration_options,omitempty"`
	Cost                    *LatestNotificationExecutionCostStruct       `json:"cost,omitempty"`
	CreatedBy               *User                                        `json:"created_by,omitempty"`
	CurrentSchedule         string                                       `json:"current_schedule,omitempty"`
	Dependencies            []*CatDependency                             `json:"dependencies,omitempty"`
	Deployment              string                                       `json:"deployment,omitempty"`
	DeploymentUrl           string                                       `json:"deployment_url,omitempty"`
	Description             string                                       `json:"description,omitempty"`
	EndsAt                  *time.Time                                   `json:"ends_at,omitempty"`
	Href                    string                                       `json:"href,omitempty"`
	Id                      string                                       `json:"id,omitempty"`
	Kind                    string                                       `json:"kind,omitempty"`
	LatestNotification      *NotificationParam                           `json:"latest_notification,omitempty"`
	LatestNotifications     []*NotificationParam                         `json:"latest_notifications,omitempty"`
	LaunchedFrom            *LaunchedFrom                                `json:"launched_from,omitempty"`
	LaunchedFromSummary     map[string]interface{}                       `json:"launched_from_summary,omitempty"`
	Links                   *ExecutionLinks                              `json:"links,omitempty"`
	Name                    string                                       `json:"name,omitempty"`
	NextAction              *ScheduledActionParam                        `json:"next_action,omitempty"`
	Outputs                 []*Output                                    `json:"outputs,omitempty"`
	RunningOperations       []*OperationParam                            `json:"running_operations,omitempty"`
	ScheduleRequired        bool                                         `json:"schedule_required,omitempty"`
	Scheduled               bool                                         `json:"scheduled,omitempty"`
	Schedules               []*Schedule                                  `json:"schedules,omitempty"`
	Source                  string                                       `json:"source,omitempty"`
	Status                  string                                       `json:"status,omitempty"`
	Timestamps              *LatestNotificationExecutionTimestampsStruct `json:"timestamps,omitempty"`
}

type LatestNotificationExecutionCostStruct struct {
	Unit      string     `json:"unit,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Value     string     `json:"value,omitempty"`
}

type LatestNotificationExecutionNextActionOperationStruct struct {
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	Name                 string                 `json:"name,omitempty"`
}

type LatestNotificationExecutionNextActionTimestampsStruct struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type LatestNotificationExecutionRunningOperationsStatusStruct struct {
	Percent int     `json:"percent,omitempty"`
	Summary string  `json:"summary,omitempty"`
	Tasks   []*Task `json:"tasks,omitempty"`
}

type LatestNotificationExecutionRunningOperationsStatusTasksStatusStruct struct {
	Percent int    `json:"percent,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type LatestNotificationExecutionRunningOperationsTimestampsStruct struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

type LatestNotificationExecutionTimestampsStruct struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	LaunchedAt   *time.Time `json:"launched_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

type LatestNotificationTimestampsStruct struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type LaunchedFrom struct {
	Type_ string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Namespace struct {
	Name    string            `json:"name,omitempty"`
	Service *NamespaceService `json:"service,omitempty"`
	Types   []*NamespaceType  `json:"types,omitempty"`
}

type NamespaceService struct {
	Auth    string                 `json:"auth,omitempty"`
	Headers map[string]interface{} `json:"headers,omitempty"`
	Host    string                 `json:"host,omitempty"`
	Path    string                 `json:"path,omitempty"`
}

type NamespaceType struct {
	Delete    string                 `json:"delete,omitempty"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Path      string                 `json:"path,omitempty"`
	Provision string                 `json:"provision,omitempty"`
}

type NotificationLatestNotificationsLink struct {
	Href string `json:"href,omitempty"`
}

type NotificationLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
}

type NotificationParam struct {
	Category   string                              `json:"category,omitempty"`
	Execution  *ExecutionParam                     `json:"execution,omitempty"`
	Href       string                              `json:"href,omitempty"`
	Id         string                              `json:"id,omitempty"`
	Kind       string                              `json:"kind,omitempty"`
	Links      *NotificationLinks                  `json:"links,omitempty"`
	Message    string                              `json:"message,omitempty"`
	Read       bool                                `json:"read,omitempty"`
	Timestamps *LatestNotificationTimestampsStruct `json:"timestamps,omitempty"`
}

type OperationDefinition struct {
	Description string       `json:"description,omitempty"`
	Label       string       `json:"label,omitempty"`
	Name        string       `json:"name,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty"`
}

type OperationInfo struct {
	Description string `json:"description,omitempty"`
	Label       string `json:"label,omitempty"`
	Name        string `json:"name,omitempty"`
}

type OperationLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
}

type OperationParam struct {
	ConfigurationOptions []*ConfigurationOption                                        `json:"configuration_options,omitempty"`
	CreatedBy            *User                                                         `json:"created_by,omitempty"`
	Execution            *ExecutionParam                                               `json:"execution,omitempty"`
	Href                 string                                                        `json:"href,omitempty"`
	Id                   string                                                        `json:"id,omitempty"`
	Kind                 string                                                        `json:"kind,omitempty"`
	Label                string                                                        `json:"label,omitempty"`
	Links                *OperationLinks                                               `json:"links,omitempty"`
	Name                 string                                                        `json:"name,omitempty"`
	Status               *LatestNotificationExecutionRunningOperationsStatusStruct     `json:"status,omitempty"`
	Timestamps           *LatestNotificationExecutionRunningOperationsTimestampsStruct `json:"timestamps,omitempty"`
}

type OperationRunningOperationsLink struct {
	Href string `json:"href,omitempty"`
}

type OperationStruct struct {
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	Name                 string                 `json:"name,omitempty"`
}

type Output struct {
	Category    string      `json:"category,omitempty"`
	Description string      `json:"description,omitempty"`
	Index       int         `json:"index,omitempty"`
	Label       string      `json:"label,omitempty"`
	Name        string      `json:"name,omitempty"`
	Value       interface{} `json:"value,omitempty"`
}

type Parameter struct {
	Default     interface{}                                    `json:"default,omitempty"`
	Description string                                         `json:"description,omitempty"`
	Name        string                                         `json:"name,omitempty"`
	Operations  []interface{}                                  `json:"operations,omitempty"`
	Type_       string                                         `json:"type,omitempty"`
	Ui          *AvailableOperationsParametersUiStruct         `json:"ui,omitempty"`
	Validation  *AvailableOperationsParametersValidationStruct `json:"validation,omitempty"`
}

type Recurrence struct {
	Hour   int    `json:"hour,omitempty"`
	Minute int    `json:"minute,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type Resource struct {
	Name  string                   `json:"name,omitempty"`
	Type_ string                   `json:"type,omitempty"`
	Value *ApiResourcesValueStruct `json:"value,omitempty"`
}

type Schedule struct {
	CreatedFrom     string      `json:"created_from,omitempty"`
	Description     string      `json:"description,omitempty"`
	Name            string      `json:"name,omitempty"`
	StartRecurrence *Recurrence `json:"start_recurrence,omitempty"`
	StopRecurrence  *Recurrence `json:"stop_recurrence,omitempty"`
}

type ScheduledActionLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
}

type ScheduledActionParam struct {
	Action                string                                                 `json:"action,omitempty"`
	CreatedBy             *User                                                  `json:"created_by,omitempty"`
	Execution             *ExecutionParam                                        `json:"execution,omitempty"`
	ExecutionSchedule     bool                                                   `json:"execution_schedule,omitempty"`
	FirstOccurrence       *time.Time                                             `json:"first_occurrence,omitempty"`
	Href                  string                                                 `json:"href,omitempty"`
	Id                    string                                                 `json:"id,omitempty"`
	Kind                  string                                                 `json:"kind,omitempty"`
	Links                 *ScheduledActionLinks                                  `json:"links,omitempty"`
	Name                  string                                                 `json:"name,omitempty"`
	NextOccurrence        *time.Time                                             `json:"next_occurrence,omitempty"`
	Operation             *LatestNotificationExecutionNextActionOperationStruct  `json:"operation,omitempty"`
	Recurrence            string                                                 `json:"recurrence,omitempty"`
	RecurrenceDescription string                                                 `json:"recurrence_description,omitempty"`
	Timestamps            *LatestNotificationExecutionNextActionTimestampsStruct `json:"timestamps,omitempty"`
	Timezone              string                                                 `json:"timezone,omitempty"`
}

type StatusStruct struct {
	Percent int     `json:"percent,omitempty"`
	Summary string  `json:"summary,omitempty"`
	Tasks   []*Task `json:"tasks,omitempty"`
}

type Task struct {
	Label  string                                                               `json:"label,omitempty"`
	Name   string                                                               `json:"name,omitempty"`
	Status *LatestNotificationExecutionRunningOperationsStatusTasksStatusStruct `json:"status,omitempty"`
}

type TimestampsStruct struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	LaunchedAt   *time.Time `json:"launched_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

type TimestampsStruct2 struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type TimestampsStruct3 struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

type TimestampsStruct4 struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}
