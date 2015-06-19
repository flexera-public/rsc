//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=ss/ssm/restful_doc -output=ss/ssm -pkg=ssm -target=1.0 -client=Api
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

// Url resolver produces an action URL and HTTP method from its name and a given resource href.
// The algorithm consists of first extracting the variables from the href and then substituing them
// in the action path. If there are more than one action paths then the algorithm picks the one that
// can substitute the most variables.
type UrlResolver string

func (r *UrlResolver) Url(rName, aName string) (*metadata.ActionPath, error) {
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
	return action.Url(vars)
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
	ApiResources         []*Resource            `json:"api_resources,omitempty"`
	AvailableActions     []string               `json:"available_actions,omitempty"`
	AvailableOperations  []*OperationDefinition `json:"available_operations,omitempty"`
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	Cost                 *CostStruct            `json:"cost,omitempty"`
	CreatedBy            *User                  `json:"created_by,omitempty"`
	CurrentSchedule      string                 `json:"current_schedule,omitempty"`
	Deployment           string                 `json:"deployment,omitempty"`
	DeploymentUrl        string                 `json:"deployment_url,omitempty"`
	Description          string                 `json:"description,omitempty"`
	EndsAt               *time.Time             `json:"ends_at,omitempty"`
	Href                 string                 `json:"href,omitempty"`
	Id                   string                 `json:"id,omitempty"`
	Kind                 string                 `json:"kind,omitempty"`
	LatestNotifications  []*Notification        `json:"latest_notifications,omitempty"`
	LaunchedFrom         *LaunchedFrom          `json:"launched_from,omitempty"`
	LaunchedFromSummary  map[string]interface{} `json:"launched_from_summary,omitempty"`
	Links                *ExecutionLinks        `json:"links,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	NextAction           *ScheduledAction       `json:"next_action,omitempty"`
	Outputs              []*Output              `json:"outputs,omitempty"`
	RunningOperations    []*Operation           `json:"running_operations,omitempty"`
	ScheduleRequired     bool                   `json:"schedule_required,omitempty"`
	Scheduled            bool                   `json:"scheduled,omitempty"`
	Schedules            []*Schedule            `json:"schedules,omitempty"`
	Source               string                 `json:"source,omitempty"`
	Status               string                 `json:"status,omitempty"`
	Timestamps           *TimestampsStruct      `json:"timestamps,omitempty"`
}

//===== Locator

// Execution resource locator, exposes resource actions.
type ExecutionLocator struct {
	UrlResolver
	api *Api
}

// Execution resource locator factory
func (api *Api) ExecutionLocator(href string) *ExecutionLocator {
	return &ExecutionLocator{UrlResolver(href), api}
}

//===== Actions

// GET /projects/:project_id/executions
//
// List information about the Executions, or use a filter to only return certain Executions. A view can be used for various levels of detail.
func (loc *ExecutionLocator) Index(options rsapi.ApiParams) ([]*Execution, error) {
	var res []*Execution
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter[]"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids[]"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "index")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /projects/:project_id/executions/:id
//
// Show details for a given Execution. A view can be used for various levels of detail.
func (loc *ExecutionLocator) Show(options rsapi.ApiParams) (*Execution, error) {
	var res *Execution
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "show")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /projects/:project_id/executions
//
// Create a new execution from a CAT, a compiled CAT, an Application in the Catalog, or a Template in Designer
func (loc *ExecutionLocator) Create(options rsapi.ApiParams) (*ExecutionLocator, error) {
	var res *ExecutionLocator
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var applicationHrefOpt = options["application_href"]
	if applicationHrefOpt != nil {
		payloadParams["application_href"] = applicationHrefOpt
	}
	var compiledCatOpt = options["compiled_cat"]
	if compiledCatOpt != nil {
		payloadParams["compiled_cat"] = compiledCatOpt
	}
	var currentScheduleOpt = options["current_schedule"]
	if currentScheduleOpt != nil {
		payloadParams["current_schedule"] = currentScheduleOpt
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		payloadParams["description"] = descriptionOpt
	}
	var endsAtOpt = options["ends_at"]
	if endsAtOpt != nil {
		payloadParams["ends_at"] = endsAtOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var options_Opt = options["options"]
	if options_Opt != nil {
		payloadParams["options"] = options_Opt
	}
	var scheduleRequiredOpt = options["schedule_required"]
	if scheduleRequiredOpt != nil {
		payloadParams["schedule_required"] = scheduleRequiredOpt
	}
	var scheduledActionsOpt = options["scheduled_actions"]
	if scheduledActionsOpt != nil {
		payloadParams["scheduled_actions"] = scheduledActionsOpt
	}
	var schedulesOpt = options["schedules"]
	if schedulesOpt != nil {
		payloadParams["schedules"] = schedulesOpt
	}
	var sourceOpt = options["source"]
	if sourceOpt != nil {
		payloadParams["source"] = sourceOpt
	}
	var templateHrefOpt = options["template_href"]
	if templateHrefOpt != nil {
		payloadParams["template_href"] = templateHrefOpt
	}
	uri, err := loc.Url("Execution", "create")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ExecutionLocator{UrlResolver(location), loc.api}, nil
	}
}

// PATCH /projects/:project_id/executions/:id
//
// Updates an execution end date or selected schedule.
func (loc *ExecutionLocator) Patch(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var currentScheduleOpt = options["current_schedule"]
	if currentScheduleOpt != nil {
		payloadParams["current_schedule"] = currentScheduleOpt
	}
	var endsAtOpt = options["ends_at"]
	if endsAtOpt != nil {
		payloadParams["ends_at"] = endsAtOpt
	}
	uri, err := loc.Url("Execution", "patch")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /projects/:project_id/executions/:id
//
// No description provided for delete.
func (loc *ExecutionLocator) Delete(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var forceOpt = options["force"]
	if forceOpt != nil {
		queryParams["force"] = forceOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /projects/:project_id/executions
//
// Delete several executions from the database. Note: if an execution has not successfully been terminated, there may still be associated cloud resources running.
func (loc *ExecutionLocator) MultiDelete(ids []string, options rsapi.ApiParams) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var forceOpt = options["force"]
	if forceOpt != nil {
		queryParams["force"] = forceOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "multi_delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// GET /projects/:project_id/executions/:id/download
//
// Download the CAT source for the execution.
func (loc *ExecutionLocator) Download(apiVersion string) error {
	if apiVersion == "" {
		return fmt.Errorf("apiVersion is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"api_version": apiVersion,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "download")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/:id/actions/launch
//
// Launch an Execution.
func (loc *ExecutionLocator) Launch() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "launch")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/:id/actions/start
//
// Start an Execution.
func (loc *ExecutionLocator) Start() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "start")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/:id/actions/stop
//
// Stop an Execution.
func (loc *ExecutionLocator) Stop() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "stop")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/:id/actions/terminate
//
// Terminate an Execution.
func (loc *ExecutionLocator) Terminate() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "terminate")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/actions/launch
//
// Launch several Executions.
func (loc *ExecutionLocator) MultiLaunch(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "multi_launch")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/actions/start
//
// Start several Executions.
func (loc *ExecutionLocator) MultiStart(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "multi_start")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/actions/stop
//
// Stop several Executions.
func (loc *ExecutionLocator) MultiStop(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "multi_stop")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/actions/terminate
//
// Terminate several Executions.
func (loc *ExecutionLocator) MultiTerminate(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "multi_terminate")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/:id/actions/run
//
// Runs an Operation on an Execution.
func (loc *ExecutionLocator) Run(name string, options rsapi.ApiParams) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"name": name,
	}
	var configurationOptionsOpt = options["configuration_options"]
	if configurationOptionsOpt != nil {
		payloadParams["configuration_options"] = configurationOptionsOpt
	}
	uri, err := loc.Url("Execution", "run")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/executions/actions/run
//
// Runs an Operation on several Executions.
func (loc *ExecutionLocator) MultiRun(ids []string, name string, options rsapi.ApiParams) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	if name == "" {
		return fmt.Errorf("name is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"name": name,
	}
	var configurationOptionsOpt = options["configuration_options"]
	if configurationOptionsOpt != nil {
		payloadParams["configuration_options"] = configurationOptionsOpt
	}
	uri, err := loc.Url("Execution", "multi_run")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
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

//===== Locator

// Notification resource locator, exposes resource actions.
type NotificationLocator struct {
	UrlResolver
	api *Api
}

// Notification resource locator factory
func (api *Api) NotificationLocator(href string) *NotificationLocator {
	return &NotificationLocator{UrlResolver(href), api}
}

//===== Actions

// GET /projects/:project_id/notifications
//
// List the most recent 50 Notifications. Use the filter parameter to specify specify Executions.
func (loc *NotificationLocator) Index(options rsapi.ApiParams) ([]*Notification, error) {
	var res []*Notification
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter[]"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids[]"] = idsOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Notification", "index")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /projects/:project_id/notifications/:id
//
// Get details for a specific Notification
func (loc *NotificationLocator) Show() (*Notification, error) {
	var res *Notification
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Notification", "show")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
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
// When a CloudApp is launched, a sequence of Operations is run as [explained here](http://support.rightscale.com/12-Guides/Self-Service/25_Cloud_Application_Template_Language) in the Operations section
// While a CloudApp is running, users may launch any custom Operations as defined in the CAT.
// Once a CAT is Terminated, a sequence of Operations is run as [explained here](http://support.rightscale.com/12-Guides/Self-Service/25_Cloud_Application_Template_Language#Operations) in the Operations section
type Operation struct {
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	CreatedBy            *User                  `json:"created_by,omitempty"`
	Execution            *Execution             `json:"execution,omitempty"`
	Href                 string                 `json:"href,omitempty"`
	Id                   string                 `json:"id,omitempty"`
	Kind                 string                 `json:"kind,omitempty"`
	Links                *OperationLinks        `json:"links,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Status               *StatusStruct          `json:"status,omitempty"`
	Timestamps           *TimestampsStruct      `json:"timestamps,omitempty"`
}

//===== Locator

// Operation resource locator, exposes resource actions.
type OperationLocator struct {
	UrlResolver
	api *Api
}

// Operation resource locator factory
func (api *Api) OperationLocator(href string) *OperationLocator {
	return &OperationLocator{UrlResolver(href), api}
}

//===== Actions

// GET /projects/:project_id/operations
//
// Get the list of 50 most recent Operations (usually filtered by Execution).
func (loc *OperationLocator) Index(options rsapi.ApiParams) ([]*Operation, error) {
	var res []*Operation
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter[]"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids[]"] = idsOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Operation", "index")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /projects/:project_id/operations/:id
//
// Get the details for a specific Operation
func (loc *OperationLocator) Show(options rsapi.ApiParams) (*Operation, error) {
	var res *Operation
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Operation", "show")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /projects/:project_id/operations
//
// Trigger an Operation to run by specifying the Execution ID and the name of the Operation.
func (loc *OperationLocator) Create(executionId string, name string, options rsapi.ApiParams) (*OperationLocator, error) {
	var res *OperationLocator
	if executionId == "" {
		return res, fmt.Errorf("executionId is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"execution_id": executionId,
		"name":         name,
	}
	var options_Opt = options["options"]
	if options_Opt != nil {
		payloadParams["options"] = options_Opt
	}
	uri, err := loc.Url("Operation", "create")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &OperationLocator{UrlResolver(location), loc.api}, nil
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

//===== Locator

// ScheduledAction resource locator, exposes resource actions.
type ScheduledActionLocator struct {
	UrlResolver
	api *Api
}

// ScheduledAction resource locator factory
func (api *Api) ScheduledActionLocator(href string) *ScheduledActionLocator {
	return &ScheduledActionLocator{UrlResolver(href), api}
}

//===== Actions

// GET /projects/:project_id/scheduled_actions
//
// List ScheduledAction resources in the project. The list can be filtered to a given execution.
func (loc *ScheduledActionLocator) Index(options rsapi.ApiParams) ([]*ScheduledAction, error) {
	var res []*ScheduledAction
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter[]"] = filterOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledAction", "index")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /projects/:project_id/scheduled_actions/:id
//
// Retrieve given ScheduledAction resource.
func (loc *ScheduledActionLocator) Show() (*ScheduledAction, error) {
	var res *ScheduledAction
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledAction", "show")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /projects/:project_id/scheduled_actions
//
// Create a new ScheduledAction resource.
func (loc *ScheduledActionLocator) Create(action string, executionId string, firstOccurrence *time.Time, options rsapi.ApiParams) (*ScheduledActionLocator, error) {
	var res *ScheduledActionLocator
	if action == "" {
		return res, fmt.Errorf("action is required")
	}
	if executionId == "" {
		return res, fmt.Errorf("executionId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"action":           action,
		"execution_id":     executionId,
		"first_occurrence": firstOccurrence,
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var operationOpt = options["operation"]
	if operationOpt != nil {
		payloadParams["operation"] = operationOpt
	}
	var recurrenceOpt = options["recurrence"]
	if recurrenceOpt != nil {
		payloadParams["recurrence"] = recurrenceOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		payloadParams["timezone"] = timezoneOpt
	}
	uri, err := loc.Url("ScheduledAction", "create")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ScheduledActionLocator{UrlResolver(location), loc.api}, nil
	}
}

// PATCH /projects/:project_id/scheduled_actions/:id
//
// Updates the 'next_occurrence' property of a ScheduledAction.
func (loc *ScheduledActionLocator) Patch(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var nextOccurrenceOpt = options["next_occurrence"]
	if nextOccurrenceOpt != nil {
		payloadParams["next_occurrence"] = nextOccurrenceOpt
	}
	uri, err := loc.Url("ScheduledAction", "patch")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /projects/:project_id/scheduled_actions/:id
//
// Delete a ScheduledAction.
func (loc *ScheduledActionLocator) Delete() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledAction", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /projects/:project_id/scheduled_actions/:id/actions/skip
//
// Skips the requested number of ScheduledAction occurrences. If no count is provided, one occurrence is skipped. On success, the next_occurrence view of the updated ScheduledAction is returned.
func (loc *ScheduledActionLocator) Skip(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var countOpt = options["count"]
	if countOpt != nil {
		payloadParams["count"] = countOpt
	}
	uri, err := loc.Url("ScheduledAction", "skip")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
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

type CompiledCAT struct {
	Conditions         map[string]interface{} `json:"conditions,omitempty"`
	Definitions        map[string]interface{} `json:"definitions,omitempty"`
	LongDescription    string                 `json:"long_description,omitempty"`
	Mappings           map[string]interface{} `json:"mappings,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Namespaces         []*Namespace           `json:"namespaces,omitempty"`
	Operations         map[string]interface{} `json:"operations,omitempty"`
	Outputs            map[string]interface{} `json:"outputs,omitempty"`
	Parameters         map[string]interface{} `json:"parameters,omitempty"`
	Permissions        map[string]interface{} `json:"permissions,omitempty"`
	RequiredParameters []string               `json:"required_parameters,omitempty"`
	Resources          map[string]interface{} `json:"resources,omitempty"`
	RsCaVer            int                    `json:"rs_ca_ver,omitempty"`
	ShortDescription   string                 `json:"short_description,omitempty"`
	Source             string                 `json:"source,omitempty"`
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
	ApiResources         []*Resource                                   `json:"api_resources,omitempty"`
	AvailableActions     []string                                      `json:"available_actions,omitempty"`
	AvailableOperations  []*OperationDefinition                        `json:"available_operations,omitempty"`
	ConfigurationOptions []*ConfigurationOption                        `json:"configuration_options,omitempty"`
	Cost                 *LatestNotificationsExecutionCostStruct       `json:"cost,omitempty"`
	CreatedBy            *User                                         `json:"created_by,omitempty"`
	CurrentSchedule      string                                        `json:"current_schedule,omitempty"`
	Deployment           string                                        `json:"deployment,omitempty"`
	DeploymentUrl        string                                        `json:"deployment_url,omitempty"`
	Description          string                                        `json:"description,omitempty"`
	EndsAt               *time.Time                                    `json:"ends_at,omitempty"`
	Href                 string                                        `json:"href,omitempty"`
	Id                   string                                        `json:"id,omitempty"`
	Kind                 string                                        `json:"kind,omitempty"`
	LatestNotifications  []*NotificationParam                          `json:"latest_notifications,omitempty"`
	LaunchedFrom         *LaunchedFrom                                 `json:"launched_from,omitempty"`
	LaunchedFromSummary  map[string]interface{}                        `json:"launched_from_summary,omitempty"`
	Links                *ExecutionLinks                               `json:"links,omitempty"`
	Name                 string                                        `json:"name,omitempty"`
	NextAction           *ScheduledActionParam                         `json:"next_action,omitempty"`
	Outputs              []*Output                                     `json:"outputs,omitempty"`
	RunningOperations    []*OperationParam                             `json:"running_operations,omitempty"`
	ScheduleRequired     bool                                          `json:"schedule_required,omitempty"`
	Scheduled            bool                                          `json:"scheduled,omitempty"`
	Schedules            []*Schedule                                   `json:"schedules,omitempty"`
	Source               string                                        `json:"source,omitempty"`
	Status               string                                        `json:"status,omitempty"`
	Timestamps           *LatestNotificationsExecutionTimestampsStruct `json:"timestamps,omitempty"`
}

type LatestNotificationsExecutionCostStruct struct {
	Unit      string     `json:"unit,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Value     string     `json:"value,omitempty"`
}

type LatestNotificationsExecutionNextActionOperationStruct struct {
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	Name                 string                 `json:"name,omitempty"`
}

type LatestNotificationsExecutionNextActionTimestampsStruct struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type LatestNotificationsExecutionRunningOperationsStatusStruct struct {
	Percent int     `json:"percent,omitempty"`
	Summary string  `json:"summary,omitempty"`
	Tasks   []*Task `json:"tasks,omitempty"`
}

type LatestNotificationsExecutionRunningOperationsStatusTasksStatusStruct struct {
	Percent int    `json:"percent,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type LatestNotificationsExecutionRunningOperationsTimestampsStruct struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

type LatestNotificationsExecutionTimestampsStruct struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	LaunchedAt   *time.Time `json:"launched_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

type LatestNotificationsTimestampsStruct struct {
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
	Category   string                               `json:"category,omitempty"`
	Execution  *ExecutionParam                      `json:"execution,omitempty"`
	Href       string                               `json:"href,omitempty"`
	Id         string                               `json:"id,omitempty"`
	Kind       string                               `json:"kind,omitempty"`
	Links      *NotificationLinks                   `json:"links,omitempty"`
	Message    string                               `json:"message,omitempty"`
	Read       bool                                 `json:"read,omitempty"`
	Timestamps *LatestNotificationsTimestampsStruct `json:"timestamps,omitempty"`
}

type OperationDefinition struct {
	Description string       `json:"description,omitempty"`
	Name        string       `json:"name,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty"`
}

type OperationLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
}

type OperationParam struct {
	ConfigurationOptions []*ConfigurationOption                                         `json:"configuration_options,omitempty"`
	CreatedBy            *User                                                          `json:"created_by,omitempty"`
	Execution            *ExecutionParam                                                `json:"execution,omitempty"`
	Href                 string                                                         `json:"href,omitempty"`
	Id                   string                                                         `json:"id,omitempty"`
	Kind                 string                                                         `json:"kind,omitempty"`
	Links                *OperationLinks                                                `json:"links,omitempty"`
	Name                 string                                                         `json:"name,omitempty"`
	Status               *LatestNotificationsExecutionRunningOperationsStatusStruct     `json:"status,omitempty"`
	Timestamps           *LatestNotificationsExecutionRunningOperationsTimestampsStruct `json:"timestamps,omitempty"`
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
	Action                string                                                  `json:"action,omitempty"`
	CreatedBy             *User                                                   `json:"created_by,omitempty"`
	Execution             *ExecutionParam                                         `json:"execution,omitempty"`
	ExecutionSchedule     bool                                                    `json:"execution_schedule,omitempty"`
	FirstOccurrence       *time.Time                                              `json:"first_occurrence,omitempty"`
	Href                  string                                                  `json:"href,omitempty"`
	Id                    string                                                  `json:"id,omitempty"`
	Kind                  string                                                  `json:"kind,omitempty"`
	Links                 *ScheduledActionLinks                                   `json:"links,omitempty"`
	Name                  string                                                  `json:"name,omitempty"`
	NextOccurrence        *time.Time                                              `json:"next_occurrence,omitempty"`
	Operation             *LatestNotificationsExecutionNextActionOperationStruct  `json:"operation,omitempty"`
	Recurrence            string                                                  `json:"recurrence,omitempty"`
	RecurrenceDescription string                                                  `json:"recurrence_description,omitempty"`
	Timestamps            *LatestNotificationsExecutionNextActionTimestampsStruct `json:"timestamps,omitempty"`
	Timezone              string                                                  `json:"timezone,omitempty"`
}

type StatusStruct struct {
	Percent int     `json:"percent,omitempty"`
	Summary string  `json:"summary,omitempty"`
	Tasks   []*Task `json:"tasks,omitempty"`
}

type Task struct {
	Label  string                                                                `json:"label,omitempty"`
	Name   string                                                                `json:"name,omitempty"`
	Status *LatestNotificationsExecutionRunningOperationsStatusTasksStatusStruct `json:"status,omitempty"`
}

type TimestampsStruct struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	LaunchedAt   *time.Time `json:"launched_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

type TimestampsStruct2 struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}
