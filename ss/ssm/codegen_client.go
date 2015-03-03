//************************************************************************//
//                     RightScale API client
//
// Generated
// Mar 2, 2015 at 5:58pm (PST)
// Command:
// $ praxisgen -metadata=ssm/restful_doc -output=ssm -pkg=ssm -target=1.0 -client=Api
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
	Cost                 *CostStruct2           `json:"cost,omitempty"`
	CreatedBy            *User                  `json:"created_by,omitempty"`
	Deployment           string                 `json:"deployment,omitempty"`
	DeploymentUrl        string                 `json:"deployment_url,omitempty"`
	Description          string                 `json:"description,omitempty"`
	EndsAt               time.Time              `json:"ends_at,omitempty"`
	Href                 string                 `json:"href,omitempty"`
	Id                   string                 `json:"id,omitempty"`
	Kind                 string                 `json:"kind,omitempty"`
	LatestNotifications  []*Notification        `json:"latest_notifications,omitempty"`
	LaunchedFrom         *LaunchedFrom          `json:"launched_from,omitempty"`
	LaunchedFromSummary  map[string]string      `json:"launched_from_summary,omitempty"`
	Links                *ExecutionLinks        `json:"links,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	NextAction           *ScheduledAction       `json:"next_action,omitempty"`
	NextOperation        *ScheduledOperation    `json:"next_operation,omitempty"`
	Outputs              []*Output              `json:"outputs,omitempty"`
	RunningOperations    []*Operation           `json:"running_operations,omitempty"`
	Scheduled            bool                   `json:"scheduled,omitempty"`
	Source               string                 `json:"source,omitempty"`
	Status               string                 `json:"status,omitempty"`
	Timestamps           *TimestampsStruct2     `json:"timestamps,omitempty"`
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
// List information about the Executions, or use a filter to only return certain Executions. A view can be used for various levels of detail.
func (loc *ExecutionLocator) Index(projectId string, options rsapi.ApiParams) ([]Execution, error) {
	var res []Execution
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids"] = idsOpt
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
// Show details for a given Execution. A view can be used for various levels of detail.
func (loc *ExecutionLocator) Show(id string, projectId string, options rsapi.ApiParams) (Execution, error) {
	var res Execution
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
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
// Create a new execution from a CAT, a compiled CAT, an Application in the Catalog, or a Template in Designer
func (loc *ExecutionLocator) Create(projectId string) error {
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "create")
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
// No description provided for delete.
func (loc *ExecutionLocator) Delete(id string, projectId string, options rsapi.ApiParams) (*Execution, error) {
	var res *Execution
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var forceOpt = options["force"]
	if forceOpt != nil {
		queryParams["force"] = forceOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "delete")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// DELETE /projects/:project_id/executions
// Delete several executions from the database. Note: if an execution has not successfully been terminated, there may still be associated cloud resources running.
func (loc *ExecutionLocator) MultiDelete(ids []string, projectId string, options rsapi.ApiParams) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids": ids,
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
// Download the CAT source for the execution.
func (loc *ExecutionLocator) Download(apiVersion string, id string, projectId string) (*Execution, error) {
	var res *Execution
	if apiVersion == "" {
		return res, fmt.Errorf("apiVersion is required")
	}
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"api_version": apiVersion,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "download")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// POST /projects/:project_id/executions/:id/actions/launch
// Launch an Execution.
func (loc *ExecutionLocator) Launch(id string, projectId string) (*Execution, error) {
	var res *Execution
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "launch")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// POST /projects/:project_id/executions/:id/actions/start
// Start an Execution.
func (loc *ExecutionLocator) Start(id string, projectId string) (*Execution, error) {
	var res *Execution
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "start")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// POST /projects/:project_id/executions/:id/actions/stop
// Stop an Execution.
func (loc *ExecutionLocator) Stop(id string, projectId string) (*Execution, error) {
	var res *Execution
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "stop")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// POST /projects/:project_id/executions/:id/actions/terminate
// Terminate an Execution.
func (loc *ExecutionLocator) Terminate(id string, projectId string) (*Execution, error) {
	var res *Execution
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Execution", "terminate")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// POST /projects/:project_id/executions/actions/launch
// Launch several executions.
func (loc *ExecutionLocator) MultiLaunch(ids []string, projectId string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids": ids,
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
// Start several executions.
func (loc *ExecutionLocator) MultiStart(ids []string, projectId string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids": ids,
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
// Stop several executions.
func (loc *ExecutionLocator) MultiStop(ids []string, projectId string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids": ids,
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
// Terminate several executions.
func (loc *ExecutionLocator) MultiTerminate(ids []string, projectId string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids": ids,
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
	Timestamps *TimestampsStruct3 `json:"timestamps,omitempty"`
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
// List the most recent 50 Notifications. Use the filter parameter to specify specify Executions.
func (loc *NotificationLocator) Index(projectId string, options rsapi.ApiParams) ([]Notification, error) {
	var res []Notification
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids"] = idsOpt
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
// Get details for a specific Notification
func (loc *NotificationLocator) Show(id string, projectId string) (Notification, error) {
	var res Notification
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
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
	ConfigurationOptions []*ConfigurationOption                `json:"configuration_options,omitempty"`
	CreatedBy            *User                                 `json:"created_by,omitempty"`
	Execution            *Execution                            `json:"execution,omitempty"`
	Href                 string                                `json:"href,omitempty"`
	Id                   string                                `json:"id,omitempty"`
	Kind                 string                                `json:"kind,omitempty"`
	Links                *OperationLinks                       `json:"links,omitempty"`
	Name                 string                                `json:"name,omitempty"`
	Status               *StatusStruct                         `json:"status,omitempty"`
	Timestamps           *NextOperationLastRunTimestampsStruct `json:"timestamps,omitempty"`
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
// Get the list of 50 most recent Operations (usually filtered by Execution).
func (loc *OperationLocator) Index(projectId string, options rsapi.ApiParams) ([]Operation, error) {
	var res []Operation
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids"] = idsOpt
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
// Get the details for a specific Operation
func (loc *OperationLocator) Show(id string, projectId string, options rsapi.ApiParams) (Operation, error) {
	var res Operation
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
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
// Trigger an Operation to run by specifying the Execution ID and the name of the Operation.
func (loc *OperationLocator) Create(projectId string) (*OperationLocator, error) {
	var res *OperationLocator
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
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
	Action                string                      `json:"action,omitempty"`
	CreatedBy             *User                       `json:"created_by,omitempty"`
	Execution             *Execution                  `json:"execution,omitempty"`
	FirstOccurrence       time.Time                   `json:"first_occurrence,omitempty"`
	Href                  string                      `json:"href,omitempty"`
	Id                    string                      `json:"id,omitempty"`
	Kind                  string                      `json:"kind,omitempty"`
	Links                 *ScheduledActionLinks       `json:"links,omitempty"`
	Mandatory             bool                        `json:"mandatory,omitempty"`
	Name                  string                      `json:"name,omitempty"`
	NextOccurrence        time.Time                   `json:"next_occurrence,omitempty"`
	Recurrence            string                      `json:"recurrence,omitempty"`
	RecurrenceDescription string                      `json:"recurrence_description,omitempty"`
	Timestamps            *NextActionTimestampsStruct `json:"timestamps,omitempty"`
	Timezone              string                      `json:"timezone,omitempty"`
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
// List ScheduledAction resources in the project. The list can be filtered to a given execution.
func (loc *ScheduledActionLocator) Index(projectId string, options rsapi.ApiParams) ([]ScheduledAction, error) {
	var res []ScheduledAction
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
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
// Retrieve given ScheduledAction resource.
func (loc *ScheduledActionLocator) Show(id string, projectId string) (ScheduledAction, error) {
	var res ScheduledAction
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
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
// Create a new ScheduledAction resource.
func (loc *ScheduledActionLocator) Create(projectId string) (*ScheduledAction, error) {
	var res *ScheduledAction
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledAction", "create")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// PATCH /projects/:project_id/scheduled_actions/:id
// Update one or more ScheduledAction properties. If the ScheduledAction has the mandatory attribute set to true, the 'force' flag must be set in order to modify it. All ScheduledActions created through the UI are set to 'mandatory' by default. When the 'recurrence' is updated, the 'next_occurrence' will be modified accordingly unless it's also specified.
func (loc *ScheduledActionLocator) Patch(id string, projectId string) (*ScheduledAction, error) {
	var res *ScheduledAction
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledAction", "patch")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// DELETE /projects/:project_id/scheduled_actions/:id
// Delete a ScheduledAction. If the ScheduledAction has the mandatory attribute set to true, the 'force' flag must be set in order to delete it.
func (loc *ScheduledActionLocator) Delete(id string, projectId string) (*ScheduledAction, error) {
	var res *ScheduledAction
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledAction", "delete")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// POST /projects/:project_id/scheduled_actions/:id/actions/skip
// Skips the requested number of ScheduledAction occurrences. If no count is provided, one occurrence is skipped. On success, the next_occurrence view of the updated ScheduledAction is returned.
func (loc *ScheduledActionLocator) Skip(id string, projectId string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
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

/******  ScheduledOperation ******/

// ScheduledOperations describe a set of timed occurrences for an operation to be run (at most once per day).
// Recurrence Rules are based off of the [RFC 5545](https://tools.ietf.org/html/rfc5545) iCal spec, and timezones are from the standard [tzinfo database](http://www.iana.org/time-zones).
// All DateTimes must be passed in [ISO-8601 format](https://en.wikipedia.org/wiki/ISO_8601)
type ScheduledOperation struct {
	CreatedBy             *User                    `json:"created_by,omitempty"`
	Execution             *Execution               `json:"execution,omitempty"`
	FirstOccurrence       time.Time                `json:"first_occurrence,omitempty"`
	Href                  string                   `json:"href,omitempty"`
	Id                    string                   `json:"id,omitempty"`
	Kind                  string                   `json:"kind,omitempty"`
	LastRun               *Operation               `json:"last_run,omitempty"`
	Links                 *ScheduledOperationLinks `json:"links,omitempty"`
	Mandatory             bool                     `json:"mandatory,omitempty"`
	Name                  string                   `json:"name,omitempty"`
	NextOccurrence        time.Time                `json:"next_occurrence,omitempty"`
	Operation             *OperationStruct         `json:"operation,omitempty"`
	Recurrence            string                   `json:"recurrence,omitempty"`
	RecurrenceDescription string                   `json:"recurrence_description,omitempty"`
	Timestamps            *TimestampsStruct4       `json:"timestamps,omitempty"`
	Timezone              string                   `json:"timezone,omitempty"`
}

//===== Locator

// ScheduledOperation resource locator, exposes resource actions.
type ScheduledOperationLocator struct {
	UrlResolver
	api *Api
}

// ScheduledOperation resource locator factory
func (api *Api) ScheduledOperationLocator(href string) *ScheduledOperationLocator {
	return &ScheduledOperationLocator{UrlResolver(href), api}
}

//===== Actions

// GET /projects/:project_id/scheduled_operations
// List ScheduledOperation resources in the project. The list can be filtered to a given execution.
func (loc *ScheduledOperationLocator) Index(projectId string, options rsapi.ApiParams) ([]ScheduledOperation, error) {
	var res []ScheduledOperation
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledOperation", "index")
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

// GET /projects/:project_id/scheduled_operations/:id
// Retrieve given ScheduledOperation resource.
func (loc *ScheduledOperationLocator) Show(id string, projectId string) (ScheduledOperation, error) {
	var res ScheduledOperation
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledOperation", "show")
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

// POST /projects/:project_id/scheduled_operations
// Create a new ScheduledOperation resource.
func (loc *ScheduledOperationLocator) Create(projectId string) (*ScheduledOperation, error) {
	var res *ScheduledOperation
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledOperation", "create")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// PATCH /projects/:project_id/scheduled_operations/:id
// Update one or more ScheduledOperation properties. If the ScheduledOperation has the mandatory attribute set to true, the 'force' flag must be set in order to modify it. All ScheduledOperations created through the UI are set to 'mandatory' by default. When the 'recurrence' is updated, the 'next_occurrence' will be modified accordingly unless it's also specified.
func (loc *ScheduledOperationLocator) Patch(id string, projectId string) (*ScheduledOperation, error) {
	var res *ScheduledOperation
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledOperation", "patch")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// DELETE /projects/:project_id/scheduled_operations/:id
// Delete a ScheduledOperation. If the ScheduledOperation has the mandatory attribute set to true, the 'force' flag must be set in order to delete it.
func (loc *ScheduledOperationLocator) Delete(id string, projectId string) (*ScheduledOperation, error) {
	var res *ScheduledOperation
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	if projectId == "" {
		return res, fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledOperation", "delete")
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
	err = json.Unmarshal(respBody, res)
	return res, err
}

// POST /projects/:project_id/scheduled_operations/:id/actions/skip
// Skips the requested number of ScheduledOperation occurrences. If no count is provided, one occurrence is skipped. On success, the next_occurrence view of the updated ScheduledOperation is returned.
func (loc *ScheduledOperationLocator) Skip(id string, projectId string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	if projectId == "" {
		return fmt.Errorf("projectId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledOperation", "skip")
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
	Details map[string]string `json:"details,omitempty"`
	Href    *string           `json:"href,omitempty"`
}

type AvailableOperationsParametersUiStruct struct {
	Category *string `json:"category,omitempty"`
	Index    *int    `json:"index,omitempty"`
	Label    *string `json:"label,omitempty"`
}

type AvailableOperationsParametersValidationStruct struct {
	AllowedPattern        *string  `json:"allowed_pattern,omitempty"`
	AllowedValues         []string `json:"allowed_values,omitempty"`
	ConstraintDescription *string  `json:"constraint_description,omitempty"`
	MaxLength             *int     `json:"max_length,omitempty"`
	MaxValue              *int     `json:"max_value,omitempty"`
	MinLength             *int     `json:"min_length,omitempty"`
	MinValue              *int     `json:"min_value,omitempty"`
	NoEcho                *bool    `json:"no_echo,omitempty"`
}

type ConfigurationOption struct {
	Name  *string `json:"name,omitempty"`
	Type_ *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type CostStruct struct {
	Unit      *string    `json:"unit,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Value     *string    `json:"value,omitempty"`
}

type CostStruct2 struct {
	Unit      *string    `json:"unit,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Value     *string    `json:"value,omitempty"`
}

type Execution2 struct {
	ApiResources         []*Resource            `json:"api_resources,omitempty"`
	AvailableActions     []string               `json:"available_actions,omitempty"`
	AvailableOperations  []*OperationDefinition `json:"available_operations,omitempty"`
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	Cost                 *CostStruct            `json:"cost,omitempty"`
	CreatedBy            *User                  `json:"created_by,omitempty"`
	Deployment           *string                `json:"deployment,omitempty"`
	DeploymentUrl        *string                `json:"deployment_url,omitempty"`
	Description          *string                `json:"description,omitempty"`
	EndsAt               *time.Time             `json:"ends_at,omitempty"`
	Href                 *string                `json:"href,omitempty"`
	Id                   string                 `json:"id,omitempty"`
	Kind                 *string                `json:"kind,omitempty"`
	LatestNotifications  []*Notification2       `json:"latest_notifications,omitempty"`
	LaunchedFrom         *LaunchedFrom          `json:"launched_from,omitempty"`
	LaunchedFromSummary  map[string]string      `json:"launched_from_summary,omitempty"`
	Links                *ExecutionLinks        `json:"links,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	NextAction           *ScheduledAction2      `json:"next_action,omitempty"`
	NextOperation        *ScheduledOperation2   `json:"next_operation,omitempty"`
	Outputs              []*Output              `json:"outputs,omitempty"`
	RunningOperations    []*Operation2          `json:"running_operations,omitempty"`
	Scheduled            *bool                  `json:"scheduled,omitempty"`
	Source               *string                `json:"source,omitempty"`
	Status               *string                `json:"status,omitempty"`
	Timestamps           *TimestampsStruct      `json:"timestamps,omitempty"`
}

type ExecutionLink struct {
	Href *string `json:"href,omitempty"`
	Id   string  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ExecutionLinks struct {
	LatestNotifications *NotificationLatestNotificationsLink `json:"latest_notifications,omitempty"`
	RunningOperations   *OperationRunningOperationsLink      `json:"running_operations,omitempty"`
}

type LatestNotificationsTimestampsStruct struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type LaunchedFrom struct {
	Type_ *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type NextActionTimestampsStruct struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type NextOperationLastRunStatusStruct struct {
	Percent *int    `json:"percent,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Tasks   []*Task `json:"tasks,omitempty"`
}

type NextOperationLastRunStatusTasksStatusStruct struct {
	Percent *int    `json:"percent,omitempty"`
	Summary *string `json:"summary,omitempty"`
}

type NextOperationLastRunTimestampsStruct struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

type NextOperationOperationStruct struct {
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	Name                 string                 `json:"name,omitempty"`
}

type Notification2 struct {
	Category   *string                              `json:"category,omitempty"`
	Execution  *Execution2                          `json:"execution,omitempty"`
	Href       *string                              `json:"href,omitempty"`
	Id         string                               `json:"id,omitempty"`
	Kind       *string                              `json:"kind,omitempty"`
	Links      *NotificationLinks                   `json:"links,omitempty"`
	Message    *string                              `json:"message,omitempty"`
	Read       *bool                                `json:"read,omitempty"`
	Timestamps *LatestNotificationsTimestampsStruct `json:"timestamps,omitempty"`
}

type NotificationLatestNotificationsLink struct {
	Href *string `json:"href,omitempty"`
}

type NotificationLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
}

type Operation2 struct {
	ConfigurationOptions []*ConfigurationOption                `json:"configuration_options,omitempty"`
	CreatedBy            *User                                 `json:"created_by,omitempty"`
	Execution            *Execution2                           `json:"execution,omitempty"`
	Href                 *string                               `json:"href,omitempty"`
	Id                   string                                `json:"id,omitempty"`
	Kind                 *string                               `json:"kind,omitempty"`
	Links                *OperationLinks                       `json:"links,omitempty"`
	Name                 *string                               `json:"name,omitempty"`
	Status               *NextOperationLastRunStatusStruct     `json:"status,omitempty"`
	Timestamps           *NextOperationLastRunTimestampsStruct `json:"timestamps,omitempty"`
}

type OperationDefinition struct {
	Description *string      `json:"description,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty"`
}

type OperationLink struct {
	Href *string `json:"href,omitempty"`
	Id   string  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type OperationLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
}

type OperationRunningOperationsLink struct {
	Href *string `json:"href,omitempty"`
}

type OperationStruct struct {
	ConfigurationOptions []*ConfigurationOption `json:"configuration_options,omitempty"`
	Name                 string                 `json:"name,omitempty"`
}

type Output struct {
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	Index       *int    `json:"index,omitempty"`
	Label       *string `json:"label,omitempty"`
	Name        *string `json:"name,omitempty"`
	Value       *string `json:"value,omitempty"`
}

type Parameter struct {
	Default     *string                                        `json:"default,omitempty"`
	Description *string                                        `json:"description,omitempty"`
	Name        *string                                        `json:"name,omitempty"`
	Operations  []string                                       `json:"operations,omitempty"`
	Type_       *string                                        `json:"type,omitempty"`
	Ui          *AvailableOperationsParametersUiStruct         `json:"ui,omitempty"`
	Validation  *AvailableOperationsParametersValidationStruct `json:"validation,omitempty"`
}

type Resource struct {
	Name  *string                  `json:"name,omitempty"`
	Type_ *string                  `json:"type,omitempty"`
	Value *ApiResourcesValueStruct `json:"value,omitempty"`
}

type ScheduledAction2 struct {
	Action                *string                     `json:"action,omitempty"`
	CreatedBy             *User                       `json:"created_by,omitempty"`
	Execution             *Execution2                 `json:"execution,omitempty"`
	FirstOccurrence       *time.Time                  `json:"first_occurrence,omitempty"`
	Href                  *string                     `json:"href,omitempty"`
	Id                    string                      `json:"id,omitempty"`
	Kind                  *string                     `json:"kind,omitempty"`
	Links                 *ScheduledActionLinks       `json:"links,omitempty"`
	Mandatory             *bool                       `json:"mandatory,omitempty"`
	Name                  *string                     `json:"name,omitempty"`
	NextOccurrence        *time.Time                  `json:"next_occurrence,omitempty"`
	Recurrence            *string                     `json:"recurrence,omitempty"`
	RecurrenceDescription *string                     `json:"recurrence_description,omitempty"`
	Timestamps            *NextActionTimestampsStruct `json:"timestamps,omitempty"`
	Timezone              *string                     `json:"timezone,omitempty"`
}

type ScheduledActionLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
}

type ScheduledOperation2 struct {
	CreatedBy             *User                         `json:"created_by,omitempty"`
	Execution             *Execution2                   `json:"execution,omitempty"`
	FirstOccurrence       *time.Time                    `json:"first_occurrence,omitempty"`
	Href                  *string                       `json:"href,omitempty"`
	Id                    string                        `json:"id,omitempty"`
	Kind                  *string                       `json:"kind,omitempty"`
	LastRun               *Operation2                   `json:"last_run,omitempty"`
	Links                 *ScheduledOperationLinks      `json:"links,omitempty"`
	Mandatory             *bool                         `json:"mandatory,omitempty"`
	Name                  *string                       `json:"name,omitempty"`
	NextOccurrence        *time.Time                    `json:"next_occurrence,omitempty"`
	Operation             *NextOperationOperationStruct `json:"operation,omitempty"`
	Recurrence            *string                       `json:"recurrence,omitempty"`
	RecurrenceDescription *string                       `json:"recurrence_description,omitempty"`
	Timestamps            *NextActionTimestampsStruct   `json:"timestamps,omitempty"`
	Timezone              *string                       `json:"timezone,omitempty"`
}

type ScheduledOperationLinks struct {
	Execution *ExecutionLink `json:"execution,omitempty"`
	LastRun   *OperationLink `json:"last_run,omitempty"`
}

type StatusStruct struct {
	Percent *int    `json:"percent,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Tasks   []*Task `json:"tasks,omitempty"`
}

type Task struct {
	Label  *string                                      `json:"label,omitempty"`
	Name   *string                                      `json:"name,omitempty"`
	Status *NextOperationLastRunStatusTasksStatusStruct `json:"status,omitempty"`
}

type TimestampsStruct struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	LaunchedAt   *time.Time `json:"launched_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

type TimestampsStruct2 struct {
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	LaunchedAt   *time.Time `json:"launched_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
}

type TimestampsStruct3 struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type TimestampsStruct4 struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Email *string `json:"email,omitempty"`
	Id    int     `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
}
