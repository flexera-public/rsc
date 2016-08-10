//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=ca/cac/docs/api -output=ca/cac -pkg=cac -target=1.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package cac

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

/******  Account ******/

// Accounts act as a container for clouds credentials and other RightScale concepts such as
// Deployments or ServerArrays. Users with the `enterprise_manager` permission in an account can create
// child accounts. This resource is not included in the public docs.
type Account struct {
}

//===== Locator

// AccountLocator exposes the Account resource actions.
type AccountLocator struct {
	Href
	api *API
}

// AccountLocator builds a locator from the given href.
func (api *API) AccountLocator(href string) *AccountLocator {
	return &AccountLocator{Href(href), api}
}

//===== Actions

// POST /api/accounts
//
// Create a new child account.
func (loc *AccountLocator) Create(options rsapi.APIParams) (*AccountLocator, error) {
	var res *AccountLocator
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var dunnoOpt = options["dunno"]
	if dunnoOpt != nil {
		p["dunno"] = dunnoOpt
	}
	uri, err := loc.ActionPath("Account", "create")
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
		return &AccountLocator{Href(location), loc.api}, nil
	}
}

// GET /api/accounts
//
// List all accounts.
func (loc *AccountLocator) Index(options rsapi.APIParams) ([]*Account, error) {
	var res []*Account
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Account", "index")
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

// GET /api/accounts/:id
//
// Show a specific account.
func (loc *AccountLocator) Show(options rsapi.APIParams) (*Account, error) {
	var res *Account
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Account", "show")
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

/******  AnalysisSnapshot ******/

// AnalysisSnapshots can be used to generate unique links to share data using filters over a date range.
type AnalysisSnapshot struct {
}

//===== Locator

// AnalysisSnapshotLocator exposes the AnalysisSnapshot resource actions.
type AnalysisSnapshotLocator struct {
	Href
	api *API
}

// AnalysisSnapshotLocator builds a locator from the given href.
func (api *API) AnalysisSnapshotLocator(href string) *AnalysisSnapshotLocator {
	return &AnalysisSnapshotLocator{Href(href), api}
}

//===== Actions

// POST /api/analysis_snapshots
//
// Create a new AnalysisSnapshot.
func (loc *AnalysisSnapshotLocator) Create(endTime *time.Time, granularity string, startTime *time.Time, options rsapi.APIParams) (*AnalysisSnapshotLocator, error) {
	var res *AnalysisSnapshotLocator
	if granularity == "" {
		return res, fmt.Errorf("granularity is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"end_time":    endTime,
		"granularity": granularity,
		"start_time":  startTime,
	}
	var excludedTagTypesOpt = options["excluded_tag_types"]
	if excludedTagTypesOpt != nil {
		p["excluded_tag_types"] = excludedTagTypesOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		p["filters"] = filtersOpt
	}
	var isComparisonOpt = options["is_comparison"]
	if isComparisonOpt != nil {
		p["is_comparison"] = isComparisonOpt
	}
	var metricsOpt = options["metrics"]
	if metricsOpt != nil {
		p["metrics"] = metricsOpt
	}
	var moduleStatesOpt = options["module_states"]
	if moduleStatesOpt != nil {
		p["module_states"] = moduleStatesOpt
	}
	uri, err := loc.ActionPath("AnalysisSnapshot", "create")
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
		return &AnalysisSnapshotLocator{Href(location), loc.api}, nil
	}
}

// GET /api/analysis_snapshots/:uuid
//
// Show a specific AnalysisSnapshot.
func (loc *AnalysisSnapshotLocator) Show(options rsapi.APIParams) (*AnalysisSnapshot, error) {
	var res *AnalysisSnapshot
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AnalysisSnapshot", "show")
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

/******  BudgetAlert ******/

// Enable you to set a monthly spend budget and be alerted via email when this is exceeded,
// based on either actual or forecasted spend. These emails include links to AnalysisSnapshots, which are
// generated automatically by us.
type BudgetAlert struct {
}

//===== Locator

// BudgetAlertLocator exposes the BudgetAlert resource actions.
type BudgetAlertLocator struct {
	Href
	api *API
}

// BudgetAlertLocator builds a locator from the given href.
func (api *API) BudgetAlertLocator(href string) *BudgetAlertLocator {
	return &BudgetAlertLocator{Href(href), api}
}

//===== Actions

// POST /api/budget_alerts
//
// Create a new BudgetAlert.
func (loc *BudgetAlertLocator) Create(budget *BudgetStruct, frequency string, name string, type_ string, options rsapi.APIParams) (*BudgetAlertLocator, error) {
	var res *BudgetAlertLocator
	if budget == nil {
		return res, fmt.Errorf("budget is required")
	}
	if frequency == "" {
		return res, fmt.Errorf("frequency is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if type_ == "" {
		return res, fmt.Errorf("type_ is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"budget":    budget,
		"frequency": frequency,
		"name":      name,
		"type":      type_,
	}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		p["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		p["attach_csv"] = attachCsvOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		p["filters"] = filtersOpt
	}
	uri, err := loc.ActionPath("BudgetAlert", "create")
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
		return &BudgetAlertLocator{Href(location), loc.api}, nil
	}
}

// GET /api/budget_alerts
//
// List all BudgetAlerts.
func (loc *BudgetAlertLocator) Index(options rsapi.APIParams) ([]*BudgetAlert, error) {
	var res []*BudgetAlert
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("BudgetAlert", "index")
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

// GET /api/budget_alerts/:id
//
// Show a specific BudgetAlert.
func (loc *BudgetAlertLocator) Show(options rsapi.APIParams) (*BudgetAlert, error) {
	var res *BudgetAlert
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("BudgetAlert", "show")
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

// PATCH /api/budget_alerts/:id
//
// Update the provided attributes of a BudgetAlert.
func (loc *BudgetAlertLocator) Update(options rsapi.APIParams) (*BudgetAlert, error) {
	var res *BudgetAlert
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		p["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		p["attach_csv"] = attachCsvOpt
	}
	var budgetOpt = options["budget"]
	if budgetOpt != nil {
		p["budget"] = budgetOpt
	}
	var frequencyOpt = options["frequency"]
	if frequencyOpt != nil {
		p["frequency"] = frequencyOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var type_Opt = options["type"]
	if type_Opt != nil {
		p["type"] = type_Opt
	}
	uri, err := loc.ActionPath("BudgetAlert", "update")
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

// DELETE /api/budget_alerts/:id
//
// Delete a BudgetAlert.
func (loc *BudgetAlertLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("BudgetAlert", "destroy")
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

/******  CloudBill ******/

// Enables you to get details about cloud bills. Only Amazon Web Services is supported for now.
type CloudBill struct {
}

//===== Locator

// CloudBillLocator exposes the CloudBill resource actions.
type CloudBillLocator struct {
	Href
	api *API
}

// CloudBillLocator builds a locator from the given href.
func (api *API) CloudBillLocator(href string) *CloudBillLocator {
	return &CloudBillLocator{Href(href), api}
}

//===== Actions

// GET /api/cloud_bills/actions/filter_options
//
// Gets the filter options which can be used for filtering the cloud bill breakdown calls.
func (loc *CloudBillLocator) FilterOptions(endTime *time.Time, filterTypes []string, startTime *time.Time, options rsapi.APIParams) (*Filter, error) {
	var res *Filter
	if len(filterTypes) == 0 {
		return res, fmt.Errorf("filterTypes is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":       endTime,
		"filter_types[]": filterTypes,
		"start_time":     startTime,
	}
	var cloudBillFiltersOpt = options["cloud_bill_filters"]
	if cloudBillFiltersOpt != nil {
		params["cloud_bill_filters[]"] = cloudBillFiltersOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CloudBill", "filter_options")
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

/******  CloudBillMetric ******/

// Enables you to get breakdowns of your cloud bill costs. Only Amazon Web Services is supported for now.
type CloudBillMetric struct {
}

//===== Locator

// CloudBillMetricLocator exposes the CloudBillMetric resource actions.
type CloudBillMetricLocator struct {
	Href
	api *API
}

// CloudBillMetricLocator builds a locator from the given href.
func (api *API) CloudBillMetricLocator(href string) *CloudBillMetricLocator {
	return &CloudBillMetricLocator{Href(href), api}
}

//===== Actions

// GET /api/cloud_bill_metrics/actions/grouped_time_series
//
// Calculates the time series of costs for cloud bills in a time period grouped into monthly
// time buckets and groups them into specified breakdown categories, e.g. show me cost of my
// cloud bills per month during the last year grouped by product.
func (loc *CloudBillMetricLocator) GroupedTimeSeries(endTime *time.Time, group [][]string, startTime *time.Time, options rsapi.APIParams) (*TimeSeriesMetricsResult, error) {
	var res *TimeSeriesMetricsResult
	if len(group) == 0 {
		return res, fmt.Errorf("group is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"group[]":    group,
		"start_time": startTime,
	}
	var cloudBillFiltersOpt = options["cloud_bill_filters"]
	if cloudBillFiltersOpt != nil {
		params["cloud_bill_filters[]"] = cloudBillFiltersOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CloudBillMetric", "grouped_time_series")
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

/******  CurrentUser ******/

// Represents the currently logged-in user. This resource is not included in the public docs.
type CurrentUser struct {
}

//===== Locator

// CurrentUserLocator exposes the CurrentUser resource actions.
type CurrentUserLocator struct {
	Href
	api *API
}

// CurrentUserLocator builds a locator from the given href.
func (api *API) CurrentUserLocator(href string) *CurrentUserLocator {
	return &CurrentUserLocator{Href(href), api}
}

//===== Actions

// GET /api/current_user
//
// Show the user's details.
func (loc *CurrentUserLocator) Show(options rsapi.APIParams) (*CurrentUser, error) {
	var res *CurrentUser
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CurrentUser", "show")
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

// PATCH /api/current_user
//
// Update the user's details.
func (loc *CurrentUserLocator) Update(password string, options rsapi.APIParams) (*CurrentUser, error) {
	var res *CurrentUser
	if password == "" {
		return res, fmt.Errorf("password is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"password": password,
	}
	var companyOpt = options["company"]
	if companyOpt != nil {
		p["company"] = companyOpt
	}
	var emailOpt = options["email"]
	if emailOpt != nil {
		p["email"] = emailOpt
	}
	var firstNameOpt = options["first_name"]
	if firstNameOpt != nil {
		p["first_name"] = firstNameOpt
	}
	var lastNameOpt = options["last_name"]
	if lastNameOpt != nil {
		p["last_name"] = lastNameOpt
	}
	var newPasswordOpt = options["new_password"]
	if newPasswordOpt != nil {
		p["new_password"] = newPasswordOpt
	}
	var phoneOpt = options["phone"]
	if phoneOpt != nil {
		p["phone"] = phoneOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		p["timezone"] = timezoneOpt
	}
	uri, err := loc.ActionPath("CurrentUser", "update")
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

// POST /api/current_user/actions/cloud_accounts
//
// Creates a cloud account in the first available child account,
// or the account used to login if there are no available child accounts.
func (loc *CurrentUserLocator) CloudAccounts(awsAccessKeyId string, awsAccountNumber string, awsSecretAccessKey string, cloudVendorName string) error {
	if awsAccessKeyId == "" {
		return fmt.Errorf("awsAccessKeyId is required")
	}
	if awsAccountNumber == "" {
		return fmt.Errorf("awsAccountNumber is required")
	}
	if awsSecretAccessKey == "" {
		return fmt.Errorf("awsSecretAccessKey is required")
	}
	if cloudVendorName == "" {
		return fmt.Errorf("cloudVendorName is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"aws_access_key_id":     awsAccessKeyId,
		"aws_account_number":    awsAccountNumber,
		"aws_secret_access_key": awsSecretAccessKey,
		"cloud_vendor_name":     cloudVendorName,
	}
	uri, err := loc.ActionPath("CurrentUser", "cloud_accounts")
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

// GET /api/current_user/actions/onboarding_status
//
// Gets the onboarding status of the user.
func (loc *CurrentUserLocator) OnboardingStatus(options rsapi.APIParams) (*UserOnboardingStatus, error) {
	var res *UserOnboardingStatus
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CurrentUser", "onboarding_status")
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

// GET /api/current_user/actions/environment
//
// Gets various environment settings.
func (loc *CurrentUserLocator) Environment() (*UserEnvironment, error) {
	var res *UserEnvironment
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CurrentUser", "environment")
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

/******  Instance ******/

// Enables you to get instance details, including the cost of individual instances.
type Instance struct {
}

//===== Locator

// InstanceLocator exposes the Instance resource actions.
type InstanceLocator struct {
	Href
	api *API
}

// InstanceLocator builds a locator from the given href.
func (api *API) InstanceLocator(href string) *InstanceLocator {
	return &InstanceLocator{Href(href), api}
}

//===== Actions

// GET /api/instances
//
// Gets instances that overlap with the requested time period.
func (loc *InstanceLocator) Index(endTime *time.Time, startTime *time.Time, options rsapi.APIParams) ([]*Instance, error) {
	var res []*Instance
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "index")
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

// GET /api/instances/actions/count
//
// Gets the count of instances that overlap with the requested time period.
func (loc *InstanceLocator) Count(endTime *time.Time, startTime *time.Time, options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "count")
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

// GET /api/instances/actions/exist
//
// Checks if any instances overlap with the requested time period.
func (loc *InstanceLocator) Exist(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var endTimeOpt = options["end_time"]
	if endTimeOpt != nil {
		params["end_time"] = endTimeOpt
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var startTimeOpt = options["start_time"]
	if startTimeOpt != nil {
		params["start_time"] = startTimeOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "exist")
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

// GET /api/instances/actions/export
//
// Exports the instances that overlap with the requested time period in CSV format.
func (loc *InstanceLocator) Export(endTime *time.Time, startTime *time.Time, options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "export")
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

// GET /api/instances/actions/filter_options
//
// Gets the filter options for instances that overlap with the requested time period.
func (loc *InstanceLocator) FilterOptions(endTime *time.Time, filterTypes []string, startTime *time.Time, options rsapi.APIParams) (*Filter, error) {
	var res *Filter
	if len(filterTypes) == 0 {
		return res, fmt.Errorf("filterTypes is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":       endTime,
		"filter_types[]": filterTypes,
		"start_time":     startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var searchTermOpt = options["search_term"]
	if searchTermOpt != nil {
		params["search_term"] = searchTermOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "filter_options")
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

/******  InstanceCombination ******/

// InstanceCombinations represent instances that make-up a Scenario.
// Note that, when making create and update calls, a Pattern can only be applied to an InstanceCombination once.
type InstanceCombination struct {
}

//===== Locator

// InstanceCombinationLocator exposes the InstanceCombination resource actions.
type InstanceCombinationLocator struct {
	Href
	api *API
}

// InstanceCombinationLocator builds a locator from the given href.
func (api *API) InstanceCombinationLocator(href string) *InstanceCombinationLocator {
	return &InstanceCombinationLocator{Href(href), api}
}

//===== Actions

// POST /api/scenarios/:scenario_id/instance_combinations
//
// Create a new InstanceCombination.
func (loc *InstanceCombinationLocator) Create(cloudName string, cloudVendorName string, instanceTypeName string, monthlyUsageOption string, platform string, quantity int, options rsapi.APIParams) (*InstanceCombinationLocator, error) {
	var res *InstanceCombinationLocator
	if cloudName == "" {
		return res, fmt.Errorf("cloudName is required")
	}
	if cloudVendorName == "" {
		return res, fmt.Errorf("cloudVendorName is required")
	}
	if instanceTypeName == "" {
		return res, fmt.Errorf("instanceTypeName is required")
	}
	if monthlyUsageOption == "" {
		return res, fmt.Errorf("monthlyUsageOption is required")
	}
	if platform == "" {
		return res, fmt.Errorf("platform is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"cloud_name":           cloudName,
		"cloud_vendor_name":    cloudVendorName,
		"instance_type_name":   instanceTypeName,
		"monthly_usage_option": monthlyUsageOption,
		"platform":             platform,
		"quantity":             quantity,
	}
	var datacenterNameOpt = options["datacenter_name"]
	if datacenterNameOpt != nil {
		p["datacenter_name"] = datacenterNameOpt
	}
	var monthlyUsageHoursOpt = options["monthly_usage_hours"]
	if monthlyUsageHoursOpt != nil {
		p["monthly_usage_hours"] = monthlyUsageHoursOpt
	}
	var patternsOpt = options["patterns"]
	if patternsOpt != nil {
		p["patterns"] = patternsOpt
	}
	uri, err := loc.ActionPath("InstanceCombination", "create")
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
		return &InstanceCombinationLocator{Href(location), loc.api}, nil
	}
}

// GET /api/scenarios/:scenario_id/instance_combinations/:id
//
// Show a specific InstanceCombination.
func (loc *InstanceCombinationLocator) Show(options rsapi.APIParams) (*InstanceCombination, error) {
	var res *InstanceCombination
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceCombination", "show")
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

// PATCH /api/scenarios/:scenario_id/instance_combinations/:id
//
// Update the provided attributes of an InstanceCombination.
func (loc *InstanceCombinationLocator) Update(options rsapi.APIParams) (*InstanceCombination, error) {
	var res *InstanceCombination
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var cloudNameOpt = options["cloud_name"]
	if cloudNameOpt != nil {
		p["cloud_name"] = cloudNameOpt
	}
	var cloudVendorNameOpt = options["cloud_vendor_name"]
	if cloudVendorNameOpt != nil {
		p["cloud_vendor_name"] = cloudVendorNameOpt
	}
	var datacenterNameOpt = options["datacenter_name"]
	if datacenterNameOpt != nil {
		p["datacenter_name"] = datacenterNameOpt
	}
	var instanceTypeNameOpt = options["instance_type_name"]
	if instanceTypeNameOpt != nil {
		p["instance_type_name"] = instanceTypeNameOpt
	}
	var monthlyUsageHoursOpt = options["monthly_usage_hours"]
	if monthlyUsageHoursOpt != nil {
		p["monthly_usage_hours"] = monthlyUsageHoursOpt
	}
	var monthlyUsageOptionOpt = options["monthly_usage_option"]
	if monthlyUsageOptionOpt != nil {
		p["monthly_usage_option"] = monthlyUsageOptionOpt
	}
	var patternsOpt = options["patterns"]
	if patternsOpt != nil {
		p["patterns"] = patternsOpt
	}
	var platformOpt = options["platform"]
	if platformOpt != nil {
		p["platform"] = platformOpt
	}
	var quantityOpt = options["quantity"]
	if quantityOpt != nil {
		p["quantity"] = quantityOpt
	}
	uri, err := loc.ActionPath("InstanceCombination", "update")
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

// DELETE /api/scenarios/:scenario_id/instance_combinations/:id
//
// Delete an InstanceCombination.
func (loc *InstanceCombinationLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceCombination", "destroy")
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

// GET /api/scenarios/:scenario_id/instance_combinations/:id/actions/reserved_instance_prices
//
// Returns pricing details for the various reserved instances that can be purchased for this InstanceCombination.
func (loc *InstanceCombinationLocator) ReservedInstancePrices(options rsapi.APIParams) (*ReservedInstancePurchase, error) {
	var res *ReservedInstancePurchase
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceCombination", "reserved_instance_prices")
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

/******  InstanceMetric ******/

// Enables you to get aggregated metrics from instances, such as total_cost or lowest_instance_count.
type InstanceMetric struct {
}

//===== Locator

// InstanceMetricLocator exposes the InstanceMetric resource actions.
type InstanceMetricLocator struct {
	Href
	api *API
}

// InstanceMetricLocator builds a locator from the given href.
func (api *API) InstanceMetricLocator(href string) *InstanceMetricLocator {
	return &InstanceMetricLocator{Href(href), api}
}

//===== Actions

// GET /api/instance_metrics/actions/overall
//
// Calculates the overall metrics for instance usages in a time period, e.g. show me the
// total cost of all my instances during the last month.
func (loc *InstanceMetricLocator) Overall(endTime *time.Time, metrics []string, startTime *time.Time, options rsapi.APIParams) (*MetricsResult, error) {
	var res *MetricsResult
	if len(metrics) == 0 {
		return res, fmt.Errorf("metrics is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"metrics[]":  metrics,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceMetric", "overall")
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

// GET /api/instance_metrics/actions/grouped_overall
//
// Calculates the overall metrics for instance usages in a time period and groups them into
// specified breakdown categories, e.g. show me the total cost of all my instances during the
// last month grouped by different accounts.
func (loc *InstanceMetricLocator) GroupedOverall(endTime *time.Time, group []string, metrics []string, startTime *time.Time, options rsapi.APIParams) (*MetricsResult, error) {
	var res *MetricsResult
	if len(group) == 0 {
		return res, fmt.Errorf("group is required")
	}
	if len(metrics) == 0 {
		return res, fmt.Errorf("metrics is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"group[]":    group,
		"metrics[]":  metrics,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceMetric", "grouped_overall")
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

// GET /api/instance_metrics/actions/time_series
//
// Calculates the metrics time series for instance usages in a time period allowing different
// time buckets (hour, 3 days, month, etc.), e.g. show me the lowest instance count of my
// instances per day during the last month.
func (loc *InstanceMetricLocator) TimeSeries(endTime *time.Time, granularity string, metrics []string, startTime *time.Time, options rsapi.APIParams) (*TimeSeriesMetricsResult, error) {
	var res *TimeSeriesMetricsResult
	if granularity == "" {
		return res, fmt.Errorf("granularity is required")
	}
	if len(metrics) == 0 {
		return res, fmt.Errorf("metrics is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":    endTime,
		"granularity": granularity,
		"metrics[]":   metrics,
		"start_time":  startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var intervalOpt = options["interval"]
	if intervalOpt != nil {
		params["interval"] = intervalOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceMetric", "time_series")
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

// GET /api/instance_metrics/actions/grouped_time_series
//
// Calculates the metrics time series for instance usages in a time period allowing different
// time buckets (hour, 3 days, month, etc.) and groups them into specified breakdown
// categories, e.g. show me the lowest instance count of my instances per day during the last
// month grouped by accounts.
func (loc *InstanceMetricLocator) GroupedTimeSeries(endTime *time.Time, granularity string, group []string, metrics []string, startTime *time.Time, options rsapi.APIParams) (*TimeSeriesMetricsResult, error) {
	var res *TimeSeriesMetricsResult
	if granularity == "" {
		return res, fmt.Errorf("granularity is required")
	}
	if len(group) == 0 {
		return res, fmt.Errorf("group is required")
	}
	if len(metrics) == 0 {
		return res, fmt.Errorf("metrics is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":    endTime,
		"granularity": granularity,
		"group[]":     group,
		"metrics[]":   metrics,
		"start_time":  startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var intervalOpt = options["interval"]
	if intervalOpt != nil {
		params["interval"] = intervalOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceMetric", "grouped_time_series")
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

// GET /api/instance_metrics/actions/current_count
//
// Returns the count of currently running instances.
func (loc *InstanceMetricLocator) CurrentCount(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		params["instance_filters[]"] = instanceFiltersOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceMetric", "current_count")
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

/******  InstanceUsagePeriod ******/

// Enables you to get usage period details from instances. An instance can have many usage periods, which can
// be caused by stop/start actions or changes to the instance type etc. InstanceUsagePeriods are used internally to
// calculate aggregate InstanceMetrics.
type InstanceUsagePeriod struct {
}

//===== Locator

// InstanceUsagePeriodLocator exposes the InstanceUsagePeriod resource actions.
type InstanceUsagePeriodLocator struct {
	Href
	api *API
}

// InstanceUsagePeriodLocator builds a locator from the given href.
func (api *API) InstanceUsagePeriodLocator(href string) *InstanceUsagePeriodLocator {
	return &InstanceUsagePeriodLocator{Href(href), api}
}

//===== Actions

// GET /api/instance_usage_periods
//
// Gets the instance usage periods of instances.
func (loc *InstanceUsagePeriodLocator) Index(instanceUsagePeriodFilters []*Filter, options rsapi.APIParams) ([]*InstanceUsagePeriod, error) {
	var res []*InstanceUsagePeriod
	if len(instanceUsagePeriodFilters) == 0 {
		return res, fmt.Errorf("instanceUsagePeriodFilters is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"instance_usage_period_filters[]": instanceUsagePeriodFilters,
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceUsagePeriod", "index")
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

/******  Pattern ******/

// Patterns describe operations in usage, and can be applied to InstanceCombinations in Scenarios to model changes in the cost.
// A pattern can only be applied to an InstanceCombination once.
type Pattern struct {
}

//===== Locator

// PatternLocator exposes the Pattern resource actions.
type PatternLocator struct {
	Href
	api *API
}

// PatternLocator builds a locator from the given href.
func (api *API) PatternLocator(href string) *PatternLocator {
	return &PatternLocator{Href(href), api}
}

//===== Actions

// POST /api/patterns
//
// Create a new Pattern.
func (loc *PatternLocator) Create(months string, name string, operation string, type_ string, value float64, years string, options rsapi.APIParams) (*PatternLocator, error) {
	var res *PatternLocator
	if months == "" {
		return res, fmt.Errorf("months is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if operation == "" {
		return res, fmt.Errorf("operation is required")
	}
	if type_ == "" {
		return res, fmt.Errorf("type_ is required")
	}
	if years == "" {
		return res, fmt.Errorf("years is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"months":    months,
		"name":      name,
		"operation": operation,
		"type":      type_,
		"value":     value,
		"years":     years,
	}
	var summaryOpt = options["summary"]
	if summaryOpt != nil {
		p["summary"] = summaryOpt
	}
	uri, err := loc.ActionPath("Pattern", "create")
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
		return &PatternLocator{Href(location), loc.api}, nil
	}
}

// GET /api/patterns
//
// List all Patterns.
func (loc *PatternLocator) Index(options rsapi.APIParams) ([]*Pattern, error) {
	var res []*Pattern
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Pattern", "index")
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

// GET /api/patterns/:id
//
// Show a specific Pattern.
func (loc *PatternLocator) Show(options rsapi.APIParams) (*Pattern, error) {
	var res *Pattern
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Pattern", "show")
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

// PATCH /api/patterns/:id
//
// Update the provided attributes of a Pattern.
func (loc *PatternLocator) Update(options rsapi.APIParams) (*Pattern, error) {
	var res *Pattern
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var monthsOpt = options["months"]
	if monthsOpt != nil {
		p["months"] = monthsOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var operationOpt = options["operation"]
	if operationOpt != nil {
		p["operation"] = operationOpt
	}
	var summaryOpt = options["summary"]
	if summaryOpt != nil {
		p["summary"] = summaryOpt
	}
	var type_Opt = options["type"]
	if type_Opt != nil {
		p["type"] = type_Opt
	}
	var valueOpt = options["value"]
	if valueOpt != nil {
		p["value"] = valueOpt
	}
	var yearsOpt = options["years"]
	if yearsOpt != nil {
		p["years"] = yearsOpt
	}
	uri, err := loc.ActionPath("Pattern", "update")
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

// DELETE /api/patterns/:id
//
// Delete a Pattern.
func (loc *PatternLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Pattern", "destroy")
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

// POST /api/patterns/actions/create_defaults
//
// Create the following commonly used default Patterns: Increase by 2% every month,
// Increase by 5% every month, Increase by 10% every month, Increase by 15% every month,
// Increase by 500% during Nov - Dec, Increase by 200% during Jan - Feb, Decrease by 2% every month,
// Decrease by 5% every month, Decrease by 10% every month, Decrease by 15% every month, Add 1 every month.
func (loc *PatternLocator) CreateDefaults(options rsapi.APIParams) (*Pattern, error) {
	var res *Pattern
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Pattern", "create_defaults")
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

/******  ReservedInstance ******/

// Enables you to get details of existing AWS ReservedInstances and some metrics about their utilization.
type ReservedInstance struct {
}

//===== Locator

// ReservedInstanceLocator exposes the ReservedInstance resource actions.
type ReservedInstanceLocator struct {
	Href
	api *API
}

// ReservedInstanceLocator builds a locator from the given href.
func (api *API) ReservedInstanceLocator(href string) *ReservedInstanceLocator {
	return &ReservedInstanceLocator{Href(href), api}
}

//===== Actions

// GET /api/reserved_instances
//
// Gets Reserved Instances that overlap with the requested time period.
func (loc *ReservedInstanceLocator) Index(endTime *time.Time, startTime *time.Time, options rsapi.APIParams) ([]*ReservedInstance, error) {
	var res []*ReservedInstance
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		params["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstance", "index")
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

// GET /api/reserved_instances/actions/count
//
// Gets the count of Reserved Instances that overlap with the requested time period.
func (loc *ReservedInstanceLocator) Count(endTime *time.Time, startTime *time.Time, options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		params["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstance", "count")
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

// GET /api/reserved_instances/actions/exist
//
// Checks if any Reserved Instances overlap with the requested time period.
func (loc *ReservedInstanceLocator) Exist(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var endTimeOpt = options["end_time"]
	if endTimeOpt != nil {
		params["end_time"] = endTimeOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		params["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var startTimeOpt = options["start_time"]
	if startTimeOpt != nil {
		params["start_time"] = startTimeOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstance", "exist")
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

// GET /api/reserved_instances/actions/export
//
// Exports the Reserved Instances that overlap with the requested time period in CSV format.
func (loc *ReservedInstanceLocator) Export(endTime *time.Time, startTime *time.Time, options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		params["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstance", "export")
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

// GET /api/reserved_instances/actions/filter_options
//
// Gets the filter options for Reserved Instances that overlap with the requested time period.
func (loc *ReservedInstanceLocator) FilterOptions(endTime *time.Time, startTime *time.Time, options rsapi.APIParams) (*Filter, error) {
	var res *Filter
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var filterTypesOpt = options["filter_types"]
	if filterTypesOpt != nil {
		params["filter_types[]"] = filterTypesOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		params["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		params["order[]"] = orderOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		params["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var searchTermOpt = options["search_term"]
	if searchTermOpt != nil {
		params["search_term"] = searchTermOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		params["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstance", "filter_options")
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

/******  ReservedInstancePurchase ******/

// ReservedInstancePurchases can be applied to InstanceCombinations in Scenarios to model changes in the cost. These are not actually purchased in the cloud and are only used for cost simulation purposes.
type ReservedInstancePurchase struct {
}

//===== Locator

// ReservedInstancePurchaseLocator exposes the ReservedInstancePurchase resource actions.
type ReservedInstancePurchaseLocator struct {
	Href
	api *API
}

// ReservedInstancePurchaseLocator builds a locator from the given href.
func (api *API) ReservedInstancePurchaseLocator(href string) *ReservedInstancePurchaseLocator {
	return &ReservedInstancePurchaseLocator{Href(href), api}
}

//===== Actions

// POST /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases
//
// Create a new ReservedInstancePurchase. This is not actually purchased in the cloud and is only used for cost simulation purposes.
func (loc *ReservedInstancePurchaseLocator) Create(autoRenew bool, duration int, offeringType string, quantity int, startDate *time.Time, options rsapi.APIParams) (*ReservedInstancePurchaseLocator, error) {
	var res *ReservedInstancePurchaseLocator
	if offeringType == "" {
		return res, fmt.Errorf("offeringType is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"auto_renew":    autoRenew,
		"duration":      duration,
		"offering_type": offeringType,
		"quantity":      quantity,
		"start_date":    startDate,
	}
	uri, err := loc.ActionPath("ReservedInstancePurchase", "create")
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
		return &ReservedInstancePurchaseLocator{Href(location), loc.api}, nil
	}
}

// GET /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases
//
// List all ReservedInstancePurchases for the InstanceCombination.
func (loc *ReservedInstancePurchaseLocator) Index(options rsapi.APIParams) ([]*ReservedInstancePurchase, error) {
	var res []*ReservedInstancePurchase
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstancePurchase", "index")
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

// GET /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases/:id
//
// Show a specific ReservedInstancePurchase.
func (loc *ReservedInstancePurchaseLocator) Show(options rsapi.APIParams) (*ReservedInstancePurchase, error) {
	var res *ReservedInstancePurchase
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstancePurchase", "show")
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

// PATCH /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases/:id
//
// Update the provided attributes of a ReservedInstancePurchase.
func (loc *ReservedInstancePurchaseLocator) Update(options rsapi.APIParams) (*ReservedInstancePurchase, error) {
	var res *ReservedInstancePurchase
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var autoRenewOpt = options["auto_renew"]
	if autoRenewOpt != nil {
		p["auto_renew"] = autoRenewOpt
	}
	var durationOpt = options["duration"]
	if durationOpt != nil {
		p["duration"] = durationOpt
	}
	var offeringTypeOpt = options["offering_type"]
	if offeringTypeOpt != nil {
		p["offering_type"] = offeringTypeOpt
	}
	var quantityOpt = options["quantity"]
	if quantityOpt != nil {
		p["quantity"] = quantityOpt
	}
	var startDateOpt = options["start_date"]
	if startDateOpt != nil {
		p["start_date"] = startDateOpt
	}
	uri, err := loc.ActionPath("ReservedInstancePurchase", "update")
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

// DELETE /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases/:id
//
// Delete a ReservedInstancePurchase. This is not actually deleted in the cloud and is only used for cost simulation purposes.
func (loc *ReservedInstancePurchaseLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ReservedInstancePurchase", "destroy")
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

/******  Scenario ******/

// Scenarios can be used to model changes in cloud usage to forecast costs over a 3-year period.
// Use the forecast action to generate the results after you create a Scenario and add your InstanceCombinations,
// ReservedInstancePurchases and Patterns.
type Scenario struct {
}

//===== Locator

// ScenarioLocator exposes the Scenario resource actions.
type ScenarioLocator struct {
	Href
	api *API
}

// ScenarioLocator builds a locator from the given href.
func (api *API) ScenarioLocator(href string) *ScenarioLocator {
	return &ScenarioLocator{Href(href), api}
}

//===== Actions

// POST /api/scenarios
//
// Create a new Scenario.
func (loc *ScenarioLocator) Create(snapshotTimestamp *time.Time, options rsapi.APIParams) (*ScenarioLocator, error) {
	var res *ScenarioLocator
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"snapshot_timestamp": snapshotTimestamp,
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		p["filters"] = filtersOpt
	}
	var isBlankOpt = options["is_blank"]
	if isBlankOpt != nil {
		p["is_blank"] = isBlankOpt
	}
	var isPersistedOpt = options["is_persisted"]
	if isPersistedOpt != nil {
		p["is_persisted"] = isPersistedOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var privateCloudInstanceCountOpt = options["private_cloud_instance_count"]
	if privateCloudInstanceCountOpt != nil {
		p["private_cloud_instance_count"] = privateCloudInstanceCountOpt
	}
	uri, err := loc.ActionPath("Scenario", "create")
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
		return &ScenarioLocator{Href(location), loc.api}, nil
	}
}

// GET /api/scenarios
//
// List all Scenarios.
func (loc *ScenarioLocator) Index(options rsapi.APIParams) ([]*Scenario, error) {
	var res []*Scenario
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var includeNonPersistedOpt = options["include_non_persisted"]
	if includeNonPersistedOpt != nil {
		params["include_non_persisted"] = includeNonPersistedOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Scenario", "index")
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

// GET /api/scenarios/:id
//
// Show a specific Scenario.
func (loc *ScenarioLocator) Show(options rsapi.APIParams) (*Scenario, error) {
	var res *Scenario
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Scenario", "show")
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

// PATCH /api/scenarios/:id
//
// Update the provided attributes of a Scenario.
func (loc *ScenarioLocator) Update(options rsapi.APIParams) (*Scenario, error) {
	var res *Scenario
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var isPersistedOpt = options["is_persisted"]
	if isPersistedOpt != nil {
		p["is_persisted"] = isPersistedOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var privateCloudInstanceCountOpt = options["private_cloud_instance_count"]
	if privateCloudInstanceCountOpt != nil {
		p["private_cloud_instance_count"] = privateCloudInstanceCountOpt
	}
	var snapshotTimestampOpt = options["snapshot_timestamp"]
	if snapshotTimestampOpt != nil {
		p["snapshot_timestamp"] = snapshotTimestampOpt
	}
	uri, err := loc.ActionPath("Scenario", "update")
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

// DELETE /api/scenarios/:id
//
// Delete a Scenario.
func (loc *ScenarioLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Scenario", "destroy")
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

// GET /api/scenarios/:id/actions/forecast
//
// Run a simulation to generate a 3-year forecast showing the `average_instance_count`, `instance_upfront_cost`,
// `instance_usage_cost` and `instance_recurring_cost` metrics. This call might get major changes so it's best to avoid using it currently.
// If there are missing prices for any of the InstanceCombinations then these metrics will be excluded from the results for that InstanceCombination.
func (loc *ScenarioLocator) Forecast(options rsapi.APIParams) (*TimeSeriesMetricsResult, error) {
	var res *TimeSeriesMetricsResult
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Scenario", "forecast")
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

/******  ScheduledReport ******/

// ScheduledReports are emailed to you, and include usage, cost, and the change from the previous reporting period.
// These emails include links to AnalysisSnapshots, which are generated automatically by us.
type ScheduledReport struct {
}

//===== Locator

// ScheduledReportLocator exposes the ScheduledReport resource actions.
type ScheduledReportLocator struct {
	Href
	api *API
}

// ScheduledReportLocator builds a locator from the given href.
func (api *API) ScheduledReportLocator(href string) *ScheduledReportLocator {
	return &ScheduledReportLocator{Href(href), api}
}

//===== Actions

// POST /api/scheduled_reports
//
// Create a new ScheduledReport.
func (loc *ScheduledReportLocator) Create(frequency string, name string, options rsapi.APIParams) (*ScheduledReportLocator, error) {
	var res *ScheduledReportLocator
	if frequency == "" {
		return res, fmt.Errorf("frequency is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"frequency": frequency,
		"name":      name,
	}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		p["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		p["attach_csv"] = attachCsvOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		p["filters"] = filtersOpt
	}
	uri, err := loc.ActionPath("ScheduledReport", "create")
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
		return &ScheduledReportLocator{Href(location), loc.api}, nil
	}
}

// GET /api/scheduled_reports
//
// List all ScheduledReports.
func (loc *ScheduledReportLocator) Index(options rsapi.APIParams) ([]*ScheduledReport, error) {
	var res []*ScheduledReport
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ScheduledReport", "index")
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

// GET /api/scheduled_reports/:id
//
// Show a specific ScheduledReport.
func (loc *ScheduledReportLocator) Show(options rsapi.APIParams) (*ScheduledReport, error) {
	var res *ScheduledReport
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ScheduledReport", "show")
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

// PATCH /api/scheduled_reports/:id
//
// Update the provided attributes of a ScheduledReport.
func (loc *ScheduledReportLocator) Update(options rsapi.APIParams) (*ScheduledReport, error) {
	var res *ScheduledReport
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		p["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		p["attach_csv"] = attachCsvOpt
	}
	var frequencyOpt = options["frequency"]
	if frequencyOpt != nil {
		p["frequency"] = frequencyOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	uri, err := loc.ActionPath("ScheduledReport", "update")
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

// DELETE /api/scheduled_reports/:id
//
// Delete a ScheduledReport.
func (loc *ScheduledReportLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ScheduledReport", "destroy")
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

// POST /api/scheduled_reports/actions/create_defaults
//
// Create the default Scheduled Report: a weekly report with no filters
func (loc *ScheduledReportLocator) CreateDefaults(options rsapi.APIParams) (*ScheduledReport, error) {
	var res *ScheduledReport
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ScheduledReport", "create_defaults")
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

/******  TempInstancePrice ******/

// This is a temporary API call that can be used by the Cloud Analytics UI until the
// Pricing Service is live, at which point this API call will be deleted. This is not included in the public docs.
type TempInstancePrice struct {
}

//===== Locator

// TempInstancePriceLocator exposes the TempInstancePrice resource actions.
type TempInstancePriceLocator struct {
	Href
	api *API
}

// TempInstancePriceLocator builds a locator from the given href.
func (api *API) TempInstancePriceLocator(href string) *TempInstancePriceLocator {
	return &TempInstancePriceLocator{Href(href), api}
}

//===== Actions

// GET /api/temp_instance_prices
//
// Returns a JSON blob with all prices for Scenario Builder.
func (loc *TempInstancePriceLocator) Index() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TempInstancePrice", "index")
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

/******  User ******/

// Users can have various permissions on multiple accounts. Users with admin permissions in an account
// can modify that account's users. This resource is not included in the public docs.
type User struct {
}

//===== Locator

// UserLocator exposes the User resource actions.
type UserLocator struct {
	Href
	api *API
}

// UserLocator builds a locator from the given href.
func (api *API) UserLocator(href string) *UserLocator {
	return &UserLocator{Href(href), api}
}

//===== Actions

// POST /api/users
//
// Create a new user with the requested permissions in the requested accounts, and emails
// them the login details. Returns an error if the user already exists.
func (loc *UserLocator) Create(accounts []*UserAccounts, email string, options rsapi.APIParams) (*UserLocator, error) {
	var res *UserLocator
	if len(accounts) == 0 {
		return res, fmt.Errorf("accounts is required")
	}
	if email == "" {
		return res, fmt.Errorf("email is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"accounts": accounts,
		"email":    email,
	}
	uri, err := loc.ActionPath("User", "create")
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
		return &UserLocator{Href(location), loc.api}, nil
	}
}

// GET /api/users
//
// List all users.
func (loc *UserLocator) Index(options rsapi.APIParams) ([]*User, error) {
	var res []*User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("User", "index")
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

// GET /api/users/:id
//
// Show a specific user.
func (loc *UserLocator) Show(options rsapi.APIParams) (*User, error) {
	var res *User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("User", "show")
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

// PATCH /api/users/:id
//
// Update a specific user's account permissions.
// This cannot be used to update other user parameters such as their name or password.
func (loc *UserLocator) Update(options rsapi.APIParams) (*User, error) {
	var res *User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var accountsOpt = options["accounts"]
	if accountsOpt != nil {
		p["accounts"] = accountsOpt
	}
	uri, err := loc.ActionPath("User", "update")
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

// POST /api/users/actions/invite
//
// Invites a user to the requested account and gives them the required permissions
// so they can add/edit cloud credentials, the user is created if they don't already exist.
// This is used during new user onboarding as the user who signs-up might not be the person who has
// the cloud credentials required to connect their clouds to RightScale.
func (loc *UserLocator) Invite(options rsapi.APIParams) (*User, error) {
	var res *User
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var accountIdOpt = options["account_id"]
	if accountIdOpt != nil {
		p["account_id"] = accountIdOpt
	}
	var emailOpt = options["email"]
	if emailOpt != nil {
		p["email"] = emailOpt
	}
	var messageOpt = options["message"]
	if messageOpt != nil {
		p["message"] = messageOpt
	}
	uri, err := loc.ActionPath("User", "invite")
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

/******  UserSetting ******/

// Used by the Cloud Analytics UI to keep track of various UI states.
type UserSetting struct {
}

//===== Locator

// UserSettingLocator exposes the UserSetting resource actions.
type UserSettingLocator struct {
	Href
	api *API
}

// UserSettingLocator builds a locator from the given href.
func (api *API) UserSettingLocator(href string) *UserSettingLocator {
	return &UserSettingLocator{Href(href), api}
}

//===== Actions

// GET /api/user_settings
//
// List the UserSettings.
func (loc *UserSettingLocator) Show(options rsapi.APIParams) (*UserSetting, error) {
	var res *UserSetting
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserSetting", "show")
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

// PATCH /api/user_settings
//
// Update the provided attributes of UserSettings.
func (loc *UserSettingLocator) Update(options rsapi.APIParams) (*UserSetting, error) {
	var res *UserSetting
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var dateRangeOpt = options["date_range"]
	if dateRangeOpt != nil {
		p["date_range"] = dateRangeOpt
	}
	var dismissedDialogsOpt = options["dismissed_dialogs"]
	if dismissedDialogsOpt != nil {
		p["dismissed_dialogs"] = dismissedDialogsOpt
	}
	var excludedTagTypesOpt = options["excluded_tag_types"]
	if excludedTagTypesOpt != nil {
		p["excluded_tag_types"] = excludedTagTypesOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		p["filters"] = filtersOpt
	}
	var granularityOpt = options["granularity"]
	if granularityOpt != nil {
		p["granularity"] = granularityOpt
	}
	var mainMenuVisibilityOpt = options["main_menu_visibility"]
	if mainMenuVisibilityOpt != nil {
		p["main_menu_visibility"] = mainMenuVisibilityOpt
	}
	var metricsOpt = options["metrics"]
	if metricsOpt != nil {
		p["metrics"] = metricsOpt
	}
	var moduleStatesOpt = options["module_states"]
	if moduleStatesOpt != nil {
		p["module_states"] = moduleStatesOpt
	}
	var onboardingStatusOpt = options["onboarding_status"]
	if onboardingStatusOpt != nil {
		p["onboarding_status"] = onboardingStatusOpt
	}
	var selectedCloudVendorNamesOpt = options["selected_cloud_vendor_names"]
	if selectedCloudVendorNamesOpt != nil {
		p["selected_cloud_vendor_names"] = selectedCloudVendorNamesOpt
	}
	var sortingOpt = options["sorting"]
	if sortingOpt != nil {
		p["sorting"] = sortingOpt
	}
	var tableColumnVisibilityOpt = options["table_column_visibility"]
	if tableColumnVisibilityOpt != nil {
		p["table_column_visibility"] = tableColumnVisibilityOpt
	}
	uri, err := loc.ActionPath("UserSetting", "update")
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

/****** Parameter Data Types ******/

type AccountParam struct {
	CloudAccounts                             []*CloudAccount `json:"cloud_accounts,omitempty"`
	CloudAnalyticsEnabled                     bool            `json:"cloud_analytics_enabled,omitempty"`
	EnterpriseId                              int             `json:"enterprise_id,omitempty"`
	EnterpriseName                            string          `json:"enterprise_name,omitempty"`
	ExpiresIn                                 int             `json:"expires_in,omitempty"`
	Href                                      string          `json:"href,omitempty"`
	Id                                        int             `json:"id,omitempty"`
	IsCloudAnalyticsBlockedByBillingAdminOnly bool            `json:"is_cloud_analytics_blocked_by_billing_admin_only,omitempty"`
	IsEnterpriseParent                        bool            `json:"is_enterprise_parent,omitempty"`
	Kind                                      string          `json:"kind,omitempty"`
	Name                                      string          `json:"name,omitempty"`
	OwnerId                                   int             `json:"owner_id,omitempty"`
	ParentAccountId                           int             `json:"parent_account_id,omitempty"`
	ParentAccountName                         string          `json:"parent_account_name,omitempty"`
	PlanCode                                  string          `json:"plan_code,omitempty"`
	ShardId                                   int             `json:"shard_id,omitempty"`
	UserHasActor                              bool            `json:"user_has_actor,omitempty"`
	UserHasAdmin                              bool            `json:"user_has_admin,omitempty"`
	UserHasEnterpriseManager                  bool            `json:"user_has_enterprise_manager,omitempty"`
	UsesIpWhitelisting                        bool            `json:"uses_ip_whitelisting,omitempty"`
}

type AnalysisSnapshotParam struct {
	CreatedAt                   *time.Time     `json:"created_at,omitempty"`
	CreatedBy                   string         `json:"created_by,omitempty"`
	EndTime                     *time.Time     `json:"end_time,omitempty"`
	ExcludedTagTypes            []string       `json:"excluded_tag_types,omitempty"`
	Filters                     []*Filter      `json:"filters,omitempty"`
	Granularity                 string         `json:"granularity,omitempty"`
	Href                        string         `json:"href,omitempty"`
	IsComparison                bool           `json:"is_comparison,omitempty"`
	Kind                        string         `json:"kind,omitempty"`
	Metrics                     []string       `json:"metrics,omitempty"`
	MissingAccessToSomeAccounts bool           `json:"missing_access_to_some_accounts,omitempty"`
	ModuleStates                []*ModuleState `json:"module_states,omitempty"`
	StartTime                   *time.Time     `json:"start_time,omitempty"`
	UpdatedAt                   *time.Time     `json:"updated_at,omitempty"`
	Uuid                        string         `json:"uuid,omitempty"`
}

type BudgetAlertParam struct {
	AdditionalEmails []string            `json:"additional_emails,omitempty"`
	AttachCsv        bool                `json:"attach_csv,omitempty"`
	Budget           *ReturnBudgetStruct `json:"budget,omitempty"`
	CreatedAt        *time.Time          `json:"created_at,omitempty"`
	Filters          []*Filter           `json:"filters,omitempty"`
	Frequency        string              `json:"frequency,omitempty"`
	Href             string              `json:"href,omitempty"`
	Id               int                 `json:"id,omitempty"`
	Kind             string              `json:"kind,omitempty"`
	Name             string              `json:"name,omitempty"`
	Type_            string              `json:"type,omitempty"`
	UpdatedAt        *time.Time          `json:"updated_at,omitempty"`
}

type BudgetStruct struct {
	Amount float64 `json:"amount,omitempty"`
	Period string  `json:"period,omitempty"`
}

type CloudAccount struct {
	CloudId   int    `json:"cloud_id,omitempty"`
	CloudName string `json:"cloud_name,omitempty"`
	CloudType string `json:"cloud_type,omitempty"`
	Kind      string `json:"kind,omitempty"`
}

type CurrentUserParam struct {
	Company   string     `json:"company,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Email     string     `json:"email,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	Id        int        `json:"id,omitempty"`
	Kind      string     `json:"kind,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Timezone  string     `json:"timezone,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type DateRangeStruct struct {
	EndTime      *time.Time `json:"end_time,omitempty"`
	IsComparison bool       `json:"is_comparison,omitempty"`
	StartTime    *time.Time `json:"start_time,omitempty"`
	Type_        string     `json:"type,omitempty"`
}

type Filter struct {
	Kind            string `json:"kind,omitempty"`
	Label           string `json:"label,omitempty"`
	TagResourceType string `json:"tag_resource_type,omitempty"`
	Type_           string `json:"type,omitempty"`
	Value           string `json:"value,omitempty"`
}

type InstanceCombinationLinks struct {
	Scenario *ScenarioParam `json:"scenario,omitempty"`
}

type InstanceCombinationParam struct {
	CloudName                 string                           `json:"cloud_name,omitempty"`
	CloudVendorName           string                           `json:"cloud_vendor_name,omitempty"`
	CreatedAt                 *time.Time                       `json:"created_at,omitempty"`
	DatacenterName            string                           `json:"datacenter_name,omitempty"`
	Href                      string                           `json:"href,omitempty"`
	Id                        int                              `json:"id,omitempty"`
	InstanceTypeName          string                           `json:"instance_type_name,omitempty"`
	Kind                      string                           `json:"kind,omitempty"`
	Links                     *InstanceCombinationLinks        `json:"links,omitempty"`
	MonthlyUsageHours         int                              `json:"monthly_usage_hours,omitempty"`
	MonthlyUsageOption        string                           `json:"monthly_usage_option,omitempty"`
	Patterns                  []*PatternParam                  `json:"patterns,omitempty"`
	Platform                  string                           `json:"platform,omitempty"`
	Quantity                  int                              `json:"quantity,omitempty"`
	ReservedInstancePurchases []*ReservedInstancePurchaseParam `json:"reserved_instance_purchases,omitempty"`
	Scenario                  *ScenarioParam                   `json:"scenario,omitempty"`
	UpdatedAt                 *time.Time                       `json:"updated_at,omitempty"`
}

type InstanceParam struct {
	AccountId                         int        `json:"account_id,omitempty"`
	AccountName                       string     `json:"account_name,omitempty"`
	CloudId                           int        `json:"cloud_id,omitempty"`
	CloudName                         string     `json:"cloud_name,omitempty"`
	CloudVendorName                   string     `json:"cloud_vendor_name,omitempty"`
	DatacenterKey                     string     `json:"datacenter_key,omitempty"`
	DatacenterName                    string     `json:"datacenter_name,omitempty"`
	DeploymentId                      int        `json:"deployment_id,omitempty"`
	DeploymentName                    string     `json:"deployment_name,omitempty"`
	EstimatedCostForPeriod            float64    `json:"estimated_cost_for_period,omitempty"`
	EstimatedManagedRcuCountForPeriod float64    `json:"estimated_managed_rcu_count_for_period,omitempty"`
	IncarnatorId                      int        `json:"incarnator_id,omitempty"`
	IncarnatorType                    string     `json:"incarnator_type,omitempty"`
	InstanceEndAt                     *time.Time `json:"instance_end_at,omitempty"`
	InstanceKey                       string     `json:"instance_key,omitempty"`
	InstanceName                      string     `json:"instance_name,omitempty"`
	InstanceRsid                      string     `json:"instance_rsid,omitempty"`
	InstanceStartAt                   *time.Time `json:"instance_start_at,omitempty"`
	InstanceTypeKey                   string     `json:"instance_type_key,omitempty"`
	InstanceTypeName                  string     `json:"instance_type_name,omitempty"`
	InstanceUid                       string     `json:"instance_uid,omitempty"`
	Kind                              string     `json:"kind,omitempty"`
	Platform                          string     `json:"platform,omitempty"`
	ProvisionedByUserEmail            string     `json:"provisioned_by_user_email,omitempty"`
	ProvisionedByUserId               int        `json:"provisioned_by_user_id,omitempty"`
	ServerTemplateId                  int        `json:"server_template_id,omitempty"`
	ServerTemplateName                string     `json:"server_template_name,omitempty"`
	State                             string     `json:"state,omitempty"`
	Tags                              []*Tag     `json:"tags,omitempty"`
	TotalUsageHours                   float64    `json:"total_usage_hours,omitempty"`
}

type InstanceUsagePeriodParam struct {
	EstimatedCost            float64    `json:"estimated_cost,omitempty"`
	EstimatedManagedRcuCount float64    `json:"estimated_managed_rcu_count,omitempty"`
	HourlyPrice              float64    `json:"hourly_price,omitempty"`
	InstanceKey              string     `json:"instance_key,omitempty"`
	InstanceTypeName         string     `json:"instance_type_name,omitempty"`
	Kind                     string     `json:"kind,omitempty"`
	PricingType              string     `json:"pricing_type,omitempty"`
	RcuRate                  float64    `json:"rcu_rate,omitempty"`
	ReservationUid           string     `json:"reservation_uid,omitempty"`
	UsageEndAt               *time.Time `json:"usage_end_at,omitempty"`
	UsageStartAt             *time.Time `json:"usage_start_at,omitempty"`
}

type Metrics struct {
	AverageInstanceCount          float64 `json:"average_instance_count,omitempty"`
	HighestInstanceCount          float64 `json:"highest_instance_count,omitempty"`
	InstanceUsageCost             float64 `json:"instance_usage_cost,omitempty"`
	Kind                          string  `json:"kind,omitempty"`
	LowestInstanceCount           float64 `json:"lowest_instance_count,omitempty"`
	ManagedInstanceHours          float64 `json:"managed_instance_hours,omitempty"`
	ManagedInstanceRcuCount       float64 `json:"managed_instance_rcu_count,omitempty"`
	ReservedInstanceRecurringCost float64 `json:"reserved_instance_recurring_cost,omitempty"`
	ReservedInstanceUpfrontCost   float64 `json:"reserved_instance_upfront_cost,omitempty"`
	TotalCost                     float64 `json:"total_cost,omitempty"`
	UnmanagedInstanceHours        float64 `json:"unmanaged_instance_hours,omitempty"`
	UnmanagedInstanceRcuCount     float64 `json:"unmanaged_instance_rcu_count,omitempty"`
	WastedReservedInstanceCost    float64 `json:"wasted_reserved_instance_cost,omitempty"`
}

type MetricsResult struct {
	BreakdownMetricsResults []*MetricsResult       `json:"breakdown_metrics_results,omitempty"`
	Group                   map[string]interface{} `json:"group,omitempty"`
	Kind                    string                 `json:"kind,omitempty"`
	Metrics                 *Metrics               `json:"metrics,omitempty"`
}

type ModuleState struct {
	Active   bool   `json:"active,omitempty"`
	Expanded bool   `json:"expanded,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SortKey  string `json:"sort_key,omitempty"`
	Type_    string `json:"type,omitempty"`
}

type PatternParam struct {
	CreatedAt *time.Time       `json:"created_at,omitempty"`
	Href      string           `json:"href,omitempty"`
	Id        int              `json:"id,omitempty"`
	Kind      string           `json:"kind,omitempty"`
	Months    string           `json:"months,omitempty"`
	Name      string           `json:"name,omitempty"`
	Operation string           `json:"operation,omitempty"`
	Scenarios []*ScenarioParam `json:"scenarios,omitempty"`
	Summary   string           `json:"summary,omitempty"`
	Type_     string           `json:"type,omitempty"`
	UpdatedAt *time.Time       `json:"updated_at,omitempty"`
	Value     float64          `json:"value,omitempty"`
	Years     string           `json:"years,omitempty"`
}

type ReservedInstanceParam struct {
	AccountId             int        `json:"account_id,omitempty"`
	AccountName           string     `json:"account_name,omitempty"`
	CloudId               int        `json:"cloud_id,omitempty"`
	CloudName             string     `json:"cloud_name,omitempty"`
	CloudVendorName       string     `json:"cloud_vendor_name,omitempty"`
	CostSaved             float64    `json:"cost_saved,omitempty"`
	DatacenterKey         string     `json:"datacenter_key,omitempty"`
	DatacenterName        string     `json:"datacenter_name,omitempty"`
	Duration              int        `json:"duration,omitempty"`
	EndTime               *time.Time `json:"end_time,omitempty"`
	InstanceCount         int        `json:"instance_count,omitempty"`
	InstanceTypeKey       string     `json:"instance_type_key,omitempty"`
	InstanceTypeName      string     `json:"instance_type_name,omitempty"`
	Kind                  string     `json:"kind,omitempty"`
	OfferingType          string     `json:"offering_type,omitempty"`
	Platform              string     `json:"platform,omitempty"`
	ReservationUid        string     `json:"reservation_uid,omitempty"`
	StartTime             *time.Time `json:"start_time,omitempty"`
	State                 string     `json:"state,omitempty"`
	Tenancy               string     `json:"tenancy,omitempty"`
	UnusedRecurringCost   float64    `json:"unused_recurring_cost,omitempty"`
	UtilizationPercentage float64    `json:"utilization_percentage,omitempty"`
}

type ReservedInstancePurchaseLinks struct {
	InstanceCombination *InstanceCombinationParam `json:"instance_combination,omitempty"`
}

type ReservedInstancePurchaseParam struct {
	AutoRenew           bool                           `json:"auto_renew,omitempty"`
	CreatedAt           *time.Time                     `json:"created_at,omitempty"`
	Duration            int                            `json:"duration,omitempty"`
	Href                string                         `json:"href,omitempty"`
	Id                  int                            `json:"id,omitempty"`
	InstanceCombination *InstanceCombinationParam      `json:"instance_combination,omitempty"`
	Kind                string                         `json:"kind,omitempty"`
	Links               *ReservedInstancePurchaseLinks `json:"links,omitempty"`
	OfferingType        string                         `json:"offering_type,omitempty"`
	Quantity            int                            `json:"quantity,omitempty"`
	StartDate           *time.Time                     `json:"start_date,omitempty"`
	UpdatedAt           *time.Time                     `json:"updated_at,omitempty"`
}

type ReturnBudgetStruct struct {
	Amount float64 `json:"amount,omitempty"`
	Period string  `json:"period,omitempty"`
}

type ReturnCurrentUserStruct struct {
	BetaEnabled                          bool       `json:"beta_enabled,omitempty"`
	CanSeeCostAndRcuMetrics              bool       `json:"can_see_cost_and_rcu_metrics,omitempty"`
	CanSeeManagedRcus                    bool       `json:"can_see_managed_rcus,omitempty"`
	CanSeeUnmanagedRcus                  bool       `json:"can_see_unmanaged_rcus,omitempty"`
	Company                              string     `json:"company,omitempty"`
	Email                                string     `json:"email,omitempty"`
	FirstLoginAt                         *time.Time `json:"first_login_at,omitempty"`
	FirstName                            string     `json:"first_name,omitempty"`
	HasAdminOnAnyAccount                 bool       `json:"has_admin_on_any_account,omitempty"`
	HasCloudAnalyticsEnabledAccounts     bool       `json:"has_cloud_analytics_enabled_accounts,omitempty"`
	HasNonIpWhitelistedAccountsWithAdmin bool       `json:"has_non_ip_whitelisted_accounts_with_admin,omitempty"`
	HasOnlyExpiredAccounts               bool       `json:"has_only_expired_accounts,omitempty"`
	Id                                   int        `json:"id,omitempty"`
	IsCloudAnalyticsOnly                 bool       `json:"is_cloud_analytics_only,omitempty"`
	IsRightscaleEmployee                 bool       `json:"is_rightscale_employee,omitempty"`
	IsSelfserviceUser                    bool       `json:"is_selfservice_user,omitempty"`
	IsTeamUser                           bool       `json:"is_team_user,omitempty"`
	LastName                             string     `json:"last_name,omitempty"`
	NotificationMessage                  string     `json:"notification_message,omitempty"`
	Phone                                string     `json:"phone,omitempty"`
	SelfserviceUrl                       string     `json:"selfservice_url,omitempty"`
	Timezone                             string     `json:"timezone,omitempty"`
	TimezoneOffsetSeconds                int        `json:"timezone_offset_seconds,omitempty"`
	TrialEndDate                         *time.Time `json:"trial_end_date,omitempty"`
}

type ReturnGoogleAnalyticsStruct struct {
	AccountId  string `json:"account_id,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}

type ReturnUserSettingsDateRangeStruct struct {
	EndTime      *time.Time `json:"end_time,omitempty"`
	IsComparison bool       `json:"is_comparison,omitempty"`
	StartTime    *time.Time `json:"start_time,omitempty"`
	Type_        string     `json:"type,omitempty"`
}

type ScenarioParam struct {
	CreatedAt                 *time.Time                  `json:"created_at,omitempty"`
	Filters                   []*Filter                   `json:"filters,omitempty"`
	HistoricMetricsResults    []*TimeSeriesMetricsResult  `json:"historic_metrics_results,omitempty"`
	Href                      string                      `json:"href,omitempty"`
	Id                        int                         `json:"id,omitempty"`
	InstanceCombinations      []*InstanceCombinationParam `json:"instance_combinations,omitempty"`
	IsPersisted               bool                        `json:"is_persisted,omitempty"`
	Kind                      string                      `json:"kind,omitempty"`
	Name                      string                      `json:"name,omitempty"`
	PrivateCloudInstanceCount int                         `json:"private_cloud_instance_count,omitempty"`
	SnapshotTimestamp         *time.Time                  `json:"snapshot_timestamp,omitempty"`
	UpdatedAt                 *time.Time                  `json:"updated_at,omitempty"`
}

type ScheduledReportParam struct {
	AdditionalEmails []string   `json:"additional_emails,omitempty"`
	AttachCsv        bool       `json:"attach_csv,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	Filters          []*Filter  `json:"filters,omitempty"`
	Frequency        string     `json:"frequency,omitempty"`
	Href             string     `json:"href,omitempty"`
	Id               int        `json:"id,omitempty"`
	Kind             string     `json:"kind,omitempty"`
	Name             string     `json:"name,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

type Tag struct {
	Key          string `json:"key,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	Value        string `json:"value,omitempty"`
}

type TimeSeriesMetricsResult struct {
	Kind      string           `json:"kind,omitempty"`
	Results   []*MetricsResult `json:"results,omitempty"`
	Timestamp *time.Time       `json:"timestamp,omitempty"`
}

type UserAccounts struct {
	AccountId                           int      `json:"account_id,omitempty"`
	AccountName                         string   `json:"account_name,omitempty"`
	BillingAdminOnly                    bool     `json:"billing_admin_only,omitempty"`
	CloudAnalyticsAccountSettingEnabled bool     `json:"cloud_analytics_account_setting_enabled,omitempty"`
	CloudAnalyticsEnabled               bool     `json:"cloud_analytics_enabled,omitempty"`
	Kind                                string   `json:"kind,omitempty"`
	Permissions                         []string `json:"permissions,omitempty"`
}

type UserEnvironment struct {
	AnalyticsUiSha     string                       `json:"analytics_ui_sha,omitempty"`
	CloudManagementUrl string                       `json:"cloud_management_url,omitempty"`
	CurrentUser        *ReturnCurrentUserStruct     `json:"current_user,omitempty"`
	Environment        string                       `json:"environment,omitempty"`
	GoogleAnalytics    *ReturnGoogleAnalyticsStruct `json:"google_analytics,omitempty"`
	Kind               string                       `json:"kind,omitempty"`
	UseLocalAssets     bool                         `json:"use_local_assets,omitempty"`
	UserSettings       *UserSettingParam            `json:"user_settings,omitempty"`
}

type UserOnboardingStatus struct {
	AccountIdToAddFirstCloud int      `json:"account_id_to_add_first_cloud,omitempty"`
	AccountNames             []string `json:"account_names,omitempty"`
	Kind                     string   `json:"kind,omitempty"`
	Status                   string   `json:"status,omitempty"`
	UrlToAddFirstCloud       string   `json:"url_to_add_first_cloud,omitempty"`
}

type UserParam struct {
	Accounts                             []*UserAccounts `json:"accounts,omitempty"`
	Email                                string          `json:"email,omitempty"`
	HasAnyExpiredAccounts                bool            `json:"has_any_expired_accounts,omitempty"`
	HasAnyIpWhitelistedAccountsWithAdmin bool            `json:"has_any_ip_whitelisted_accounts_with_admin,omitempty"`
	Href                                 string          `json:"href,omitempty"`
	Id                                   int             `json:"id,omitempty"`
	Kind                                 string          `json:"kind,omitempty"`
}

type UserSettingParam struct {
	DateRange                *ReturnUserSettingsDateRangeStruct `json:"date_range,omitempty"`
	DismissedDialogs         map[string]interface{}             `json:"dismissed_dialogs,omitempty"`
	ExcludedTagTypes         []string                           `json:"excluded_tag_types,omitempty"`
	Filters                  []*Filter                          `json:"filters,omitempty"`
	Granularity              string                             `json:"granularity,omitempty"`
	Kind                     string                             `json:"kind,omitempty"`
	MainMenuVisibility       string                             `json:"main_menu_visibility,omitempty"`
	Metrics                  []string                           `json:"metrics,omitempty"`
	ModuleStates             []*ModuleState                     `json:"module_states,omitempty"`
	OnboardingStatus         string                             `json:"onboarding_status,omitempty"`
	SelectedCloudVendorNames map[string]interface{}             `json:"selected_cloud_vendor_names,omitempty"`
	Sorting                  map[string]interface{}             `json:"sorting,omitempty"`
	TableColumnVisibility    map[string]interface{}             `json:"table_column_visibility,omitempty"`
}
