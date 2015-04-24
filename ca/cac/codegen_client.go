//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=ca/cac/docs/api -output=ca/cac -pkg=cac -target=1.0 -client=Api
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

/******  Account ******/

// Accounts act as a container for clouds credentials and other RightScale concepts such as
// Deployments or ServerArrays. Users with the `enterprise_manager` permission in an account can create
// child accounts. This resource is not included in the public docs.
type Account struct {
}

//===== Locator

// Account resource locator, exposes resource actions.
type AccountLocator struct {
	UrlResolver
	api *Api
}

// Account resource locator factory
func (api *Api) AccountLocator(href string) *AccountLocator {
	return &AccountLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/accounts
//
// Create a new child account.
func (loc *AccountLocator) Create(options rsapi.ApiParams) (*AccountLocator, error) {
	var res *AccountLocator
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var dunnoOpt = options["dunno"]
	if dunnoOpt != nil {
		payloadParams["dunno"] = dunnoOpt
	}
	uri, err := loc.Url("Account", "create")
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
		return &AccountLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/accounts
//
// List all accounts.
func (loc *AccountLocator) Index(options rsapi.ApiParams) (*Account, error) {
	var res *Account
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Account", "index")
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

// GET /api/accounts/:id
//
// Show a specific account.
func (loc *AccountLocator) Show(options rsapi.ApiParams) (*Account, error) {
	var res *Account
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Account", "show")
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

/******  AnalysisSnapshot ******/

// AnalysisSnapshots can be used to generate unique links to share data using filters over a date range.
type AnalysisSnapshot struct {
}

//===== Locator

// AnalysisSnapshot resource locator, exposes resource actions.
type AnalysisSnapshotLocator struct {
	UrlResolver
	api *Api
}

// AnalysisSnapshot resource locator factory
func (api *Api) AnalysisSnapshotLocator(href string) *AnalysisSnapshotLocator {
	return &AnalysisSnapshotLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/analysis_snapshots
//
// Create a new AnalysisSnapshot.
func (loc *AnalysisSnapshotLocator) Create(endTime time.Time, granularity string, startTime time.Time, options rsapi.ApiParams) (*AnalysisSnapshotLocator, error) {
	var res *AnalysisSnapshotLocator
	if granularity == "" {
		return res, fmt.Errorf("granularity is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"end_time":    endTime,
		"granularity": granularity,
		"start_time":  startTime,
	}
	var excludedTagTypesOpt = options["excluded_tag_types"]
	if excludedTagTypesOpt != nil {
		payloadParams["excluded_tag_types"] = excludedTagTypesOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		payloadParams["filters"] = filtersOpt
	}
	var isComparisonOpt = options["is_comparison"]
	if isComparisonOpt != nil {
		payloadParams["is_comparison"] = isComparisonOpt
	}
	var metricsOpt = options["metrics"]
	if metricsOpt != nil {
		payloadParams["metrics"] = metricsOpt
	}
	var moduleStatesOpt = options["module_states"]
	if moduleStatesOpt != nil {
		payloadParams["module_states"] = moduleStatesOpt
	}
	uri, err := loc.Url("AnalysisSnapshot", "create")
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
		return &AnalysisSnapshotLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/analysis_snapshots/:uuid
//
// Show a specific AnalysisSnapshot.
func (loc *AnalysisSnapshotLocator) Show(options rsapi.ApiParams) (*AnalysisSnapshot, error) {
	var res *AnalysisSnapshot
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("AnalysisSnapshot", "show")
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

/******  BudgetAlert ******/

// Enable you to set a monthly spend budget and be alerted via email when this is exceeded,
// based on either actual or forecasted spend. These emails include links to AnalysisSnapshots, which are
// generated automatically by us.
type BudgetAlert struct {
}

//===== Locator

// BudgetAlert resource locator, exposes resource actions.
type BudgetAlertLocator struct {
	UrlResolver
	api *Api
}

// BudgetAlert resource locator factory
func (api *Api) BudgetAlertLocator(href string) *BudgetAlertLocator {
	return &BudgetAlertLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/budget_alerts
//
// Create a new BudgetAlert.
func (loc *BudgetAlertLocator) Create(budget *BudgetStruct, frequency string, name string, type_ string, options rsapi.ApiParams) (*BudgetAlertLocator, error) {
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
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"budget":    budget,
		"frequency": frequency,
		"name":      name,
		"type":      type_,
	}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		payloadParams["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		payloadParams["attach_csv"] = attachCsvOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		payloadParams["filters"] = filtersOpt
	}
	uri, err := loc.Url("BudgetAlert", "create")
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
		return &BudgetAlertLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/budget_alerts
//
// List all BudgetAlerts.
func (loc *BudgetAlertLocator) Index(options rsapi.ApiParams) (*BudgetAlert, error) {
	var res *BudgetAlert
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("BudgetAlert", "index")
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

// GET /api/budget_alerts/:id
//
// Show a specific BudgetAlert.
func (loc *BudgetAlertLocator) Show(options rsapi.ApiParams) (*BudgetAlert, error) {
	var res *BudgetAlert
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("BudgetAlert", "show")
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

// PATCH /api/budget_alerts/:id
//
// Update the provided attributes of a BudgetAlert.
func (loc *BudgetAlertLocator) Update(options rsapi.ApiParams) (*BudgetAlert, error) {
	var res *BudgetAlert
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		payloadParams["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		payloadParams["attach_csv"] = attachCsvOpt
	}
	var budgetOpt = options["budget"]
	if budgetOpt != nil {
		payloadParams["budget"] = budgetOpt
	}
	var frequencyOpt = options["frequency"]
	if frequencyOpt != nil {
		payloadParams["frequency"] = frequencyOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var type_Opt = options["type"]
	if type_Opt != nil {
		payloadParams["type"] = type_Opt
	}
	uri, err := loc.Url("BudgetAlert", "update")
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

// DELETE /api/budget_alerts/:id
//
// Delete a BudgetAlert.
func (loc *BudgetAlertLocator) Destroy() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("BudgetAlert", "destroy")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  CloudBill ******/

// Enables you to get details about cloud bills. Only Amazon Web Services is supported for now.
type CloudBill struct {
}

//===== Locator

// CloudBill resource locator, exposes resource actions.
type CloudBillLocator struct {
	UrlResolver
	api *Api
}

// CloudBill resource locator factory
func (api *Api) CloudBillLocator(href string) *CloudBillLocator {
	return &CloudBillLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/cloud_bills/actions/filter_options
//
// Gets the filter options which can be used for filtering the cloud bill breakdown calls.
func (loc *CloudBillLocator) FilterOptions(endTime time.Time, filterTypes []string, startTime time.Time, options rsapi.ApiParams) (*Filter, error) {
	var res *Filter
	if len(filterTypes) == 0 {
		return res, fmt.Errorf("filterTypes is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":       endTime,
		"filter_types[]": filterTypes,
		"start_time":     startTime,
	}
	var cloudBillFiltersOpt = options["cloud_bill_filters"]
	if cloudBillFiltersOpt != nil {
		queryParams["cloud_bill_filters[]"] = cloudBillFiltersOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("CloudBill", "filter_options")
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

/******  CloudBillMetric ******/

// Enables you to get breakdowns of your cloud bill costs. Only Amazon Web Services is supported for now.
type CloudBillMetric struct {
}

//===== Locator

// CloudBillMetric resource locator, exposes resource actions.
type CloudBillMetricLocator struct {
	UrlResolver
	api *Api
}

// CloudBillMetric resource locator factory
func (api *Api) CloudBillMetricLocator(href string) *CloudBillMetricLocator {
	return &CloudBillMetricLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/cloud_bill_metrics/actions/grouped_time_series
//
// Calculates the time series of costs for cloud bills in a time period grouped into monthly
// time buckets and groups them into specified breakdown categories, e.g. show me cost of my
// cloud bills per month during the last year grouped by product.
func (loc *CloudBillMetricLocator) GroupedTimeSeries(endTime time.Time, group [][]string, startTime time.Time, options rsapi.ApiParams) (*TimeSeriesMetricsResult, error) {
	var res *TimeSeriesMetricsResult
	if len(group) == 0 {
		return res, fmt.Errorf("group is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"group[]":    group,
		"start_time": startTime,
	}
	var cloudBillFiltersOpt = options["cloud_bill_filters"]
	if cloudBillFiltersOpt != nil {
		queryParams["cloud_bill_filters[]"] = cloudBillFiltersOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("CloudBillMetric", "grouped_time_series")
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

/******  CurrentUser ******/

// Represents the currently logged-in user. This resource is not included in the public docs.
type CurrentUser struct {
}

//===== Locator

// CurrentUser resource locator, exposes resource actions.
type CurrentUserLocator struct {
	UrlResolver
	api *Api
}

// CurrentUser resource locator factory
func (api *Api) CurrentUserLocator(href string) *CurrentUserLocator {
	return &CurrentUserLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/current_user
//
// Show the user's details.
func (loc *CurrentUserLocator) Show(options rsapi.ApiParams) (*CurrentUser, error) {
	var res *CurrentUser
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("CurrentUser", "show")
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

// PATCH /api/current_user
//
// Update the user's details.
func (loc *CurrentUserLocator) Update(password string, options rsapi.ApiParams) (*CurrentUser, error) {
	var res *CurrentUser
	if password == "" {
		return res, fmt.Errorf("password is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"password": password,
	}
	var companyOpt = options["company"]
	if companyOpt != nil {
		payloadParams["company"] = companyOpt
	}
	var emailOpt = options["email"]
	if emailOpt != nil {
		payloadParams["email"] = emailOpt
	}
	var firstNameOpt = options["first_name"]
	if firstNameOpt != nil {
		payloadParams["first_name"] = firstNameOpt
	}
	var lastNameOpt = options["last_name"]
	if lastNameOpt != nil {
		payloadParams["last_name"] = lastNameOpt
	}
	var newPasswordOpt = options["new_password"]
	if newPasswordOpt != nil {
		payloadParams["new_password"] = newPasswordOpt
	}
	var phoneOpt = options["phone"]
	if phoneOpt != nil {
		payloadParams["phone"] = phoneOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		payloadParams["timezone"] = timezoneOpt
	}
	uri, err := loc.Url("CurrentUser", "update")
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
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"aws_access_key_id":     awsAccessKeyId,
		"aws_account_number":    awsAccountNumber,
		"aws_secret_access_key": awsSecretAccessKey,
		"cloud_vendor_name":     cloudVendorName,
	}
	uri, err := loc.Url("CurrentUser", "cloud_accounts")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// GET /api/current_user/actions/onboarding_status
//
// Gets the onboarding status of the user.
func (loc *CurrentUserLocator) OnboardingStatus(options rsapi.ApiParams) (*UserOnboardingStatus, error) {
	var res *UserOnboardingStatus
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("CurrentUser", "onboarding_status")
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

// GET /api/current_user/actions/environment
//
// Gets various environment settings.
func (loc *CurrentUserLocator) Environment() (*UserEnvironment, error) {
	var res *UserEnvironment
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("CurrentUser", "environment")
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

/******  Instance ******/

// Enables you to get instance details, including the cost of individual instances.
type Instance struct {
}

//===== Locator

// Instance resource locator, exposes resource actions.
type InstanceLocator struct {
	UrlResolver
	api *Api
}

// Instance resource locator factory
func (api *Api) InstanceLocator(href string) *InstanceLocator {
	return &InstanceLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/instances
//
// Gets instances that overlap with the requested time period.
func (loc *InstanceLocator) Index(endTime time.Time, startTime time.Time, options rsapi.ApiParams) (*Instance, error) {
	var res *Instance
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Instance", "index")
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

// GET /api/instances/actions/count
//
// Gets the count of instances that overlap with the requested time period.
func (loc *InstanceLocator) Count(endTime time.Time, startTime time.Time, options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Instance", "count")
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
	res = string(respBody)
	return res, err
}

// GET /api/instances/actions/exist
//
// Checks if any instances overlap with the requested time period.
func (loc *InstanceLocator) Exist(options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var endTimeOpt = options["end_time"]
	if endTimeOpt != nil {
		queryParams["end_time"] = endTimeOpt
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var startTimeOpt = options["start_time"]
	if startTimeOpt != nil {
		queryParams["start_time"] = startTimeOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Instance", "exist")
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
	res = string(respBody)
	return res, err
}

// GET /api/instances/actions/export
//
// Exports the instances that overlap with the requested time period in CSV format.
func (loc *InstanceLocator) Export(endTime time.Time, startTime time.Time, options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Instance", "export")
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
	res = string(respBody)
	return res, err
}

// GET /api/instances/actions/filter_options
//
// Gets the filter options for instances that overlap with the requested time period.
func (loc *InstanceLocator) FilterOptions(endTime time.Time, filterTypes []string, startTime time.Time, options rsapi.ApiParams) (*Filter, error) {
	var res *Filter
	if len(filterTypes) == 0 {
		return res, fmt.Errorf("filterTypes is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":       endTime,
		"filter_types[]": filterTypes,
		"start_time":     startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var searchTermOpt = options["search_term"]
	if searchTermOpt != nil {
		queryParams["search_term"] = searchTermOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Instance", "filter_options")
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

/******  InstanceCombination ******/

// InstanceCombinations represent instances that make-up a Scenario.
// Note that, when making create and update calls, a Pattern can only be applied to an InstanceCombination once.
type InstanceCombination struct {
}

//===== Locator

// InstanceCombination resource locator, exposes resource actions.
type InstanceCombinationLocator struct {
	UrlResolver
	api *Api
}

// InstanceCombination resource locator factory
func (api *Api) InstanceCombinationLocator(href string) *InstanceCombinationLocator {
	return &InstanceCombinationLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/scenarios/:scenario_id/instance_combinations
//
// Create a new InstanceCombination.
func (loc *InstanceCombinationLocator) Create(cloudName string, cloudVendorName string, instanceTypeName string, monthlyUsageOption string, platform string, quantity int, options rsapi.ApiParams) (*InstanceCombinationLocator, error) {
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
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"cloud_name":           cloudName,
		"cloud_vendor_name":    cloudVendorName,
		"instance_type_name":   instanceTypeName,
		"monthly_usage_option": monthlyUsageOption,
		"platform":             platform,
		"quantity":             quantity,
	}
	var datacenterNameOpt = options["datacenter_name"]
	if datacenterNameOpt != nil {
		payloadParams["datacenter_name"] = datacenterNameOpt
	}
	var monthlyUsageHoursOpt = options["monthly_usage_hours"]
	if monthlyUsageHoursOpt != nil {
		payloadParams["monthly_usage_hours"] = monthlyUsageHoursOpt
	}
	var patternsOpt = options["patterns"]
	if patternsOpt != nil {
		payloadParams["patterns"] = patternsOpt
	}
	uri, err := loc.Url("InstanceCombination", "create")
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
		return &InstanceCombinationLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/scenarios/:scenario_id/instance_combinations/:id
//
// Show a specific InstanceCombination.
func (loc *InstanceCombinationLocator) Show(options rsapi.ApiParams) (*InstanceCombination, error) {
	var res *InstanceCombination
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceCombination", "show")
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

// PATCH /api/scenarios/:scenario_id/instance_combinations/:id
//
// Update the provided attributes of an InstanceCombination.
func (loc *InstanceCombinationLocator) Update(options rsapi.ApiParams) (*InstanceCombination, error) {
	var res *InstanceCombination
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var cloudNameOpt = options["cloud_name"]
	if cloudNameOpt != nil {
		payloadParams["cloud_name"] = cloudNameOpt
	}
	var cloudVendorNameOpt = options["cloud_vendor_name"]
	if cloudVendorNameOpt != nil {
		payloadParams["cloud_vendor_name"] = cloudVendorNameOpt
	}
	var datacenterNameOpt = options["datacenter_name"]
	if datacenterNameOpt != nil {
		payloadParams["datacenter_name"] = datacenterNameOpt
	}
	var instanceTypeNameOpt = options["instance_type_name"]
	if instanceTypeNameOpt != nil {
		payloadParams["instance_type_name"] = instanceTypeNameOpt
	}
	var monthlyUsageHoursOpt = options["monthly_usage_hours"]
	if monthlyUsageHoursOpt != nil {
		payloadParams["monthly_usage_hours"] = monthlyUsageHoursOpt
	}
	var monthlyUsageOptionOpt = options["monthly_usage_option"]
	if monthlyUsageOptionOpt != nil {
		payloadParams["monthly_usage_option"] = monthlyUsageOptionOpt
	}
	var patternsOpt = options["patterns"]
	if patternsOpt != nil {
		payloadParams["patterns"] = patternsOpt
	}
	var platformOpt = options["platform"]
	if platformOpt != nil {
		payloadParams["platform"] = platformOpt
	}
	var quantityOpt = options["quantity"]
	if quantityOpt != nil {
		payloadParams["quantity"] = quantityOpt
	}
	uri, err := loc.Url("InstanceCombination", "update")
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

// DELETE /api/scenarios/:scenario_id/instance_combinations/:id
//
// Delete an InstanceCombination.
func (loc *InstanceCombinationLocator) Destroy() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceCombination", "destroy")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// GET /api/scenarios/:scenario_id/instance_combinations/:id/actions/reserved_instance_prices
//
// Returns pricing details for the various reserved instances that can be purchased for this InstanceCombination.
func (loc *InstanceCombinationLocator) ReservedInstancePrices(options rsapi.ApiParams) (*ReservedInstancePurchase, error) {
	var res *ReservedInstancePurchase
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceCombination", "reserved_instance_prices")
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

/******  InstanceMetric ******/

// Enables you to get aggregated metrics from instances, such as total_cost or lowest_instance_count.
type InstanceMetric struct {
}

//===== Locator

// InstanceMetric resource locator, exposes resource actions.
type InstanceMetricLocator struct {
	UrlResolver
	api *Api
}

// InstanceMetric resource locator factory
func (api *Api) InstanceMetricLocator(href string) *InstanceMetricLocator {
	return &InstanceMetricLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/instance_metrics/actions/overall
//
// Calculates the overall metrics for instance usages in a time period, e.g. show me the
// total cost of all my instances during the last month.
func (loc *InstanceMetricLocator) Overall(endTime time.Time, metrics []string, startTime time.Time, options rsapi.ApiParams) (*MetricsResult, error) {
	var res *MetricsResult
	if len(metrics) == 0 {
		return res, fmt.Errorf("metrics is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"metrics[]":  metrics,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceMetric", "overall")
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

// GET /api/instance_metrics/actions/grouped_overall
//
// Calculates the overall metrics for instance usages in a time period and groups them into
// specified breakdown categories, e.g. show me the total cost of all my instances during the
// last month grouped by different accounts.
func (loc *InstanceMetricLocator) GroupedOverall(endTime time.Time, group []string, metrics []string, startTime time.Time, options rsapi.ApiParams) (*MetricsResult, error) {
	var res *MetricsResult
	if len(group) == 0 {
		return res, fmt.Errorf("group is required")
	}
	if len(metrics) == 0 {
		return res, fmt.Errorf("metrics is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"group[]":    group,
		"metrics[]":  metrics,
		"start_time": startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceMetric", "grouped_overall")
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

// GET /api/instance_metrics/actions/time_series
//
// Calculates the metrics time series for instance usages in a time period allowing different
// time buckets (hour, 3 days, month, etc.), e.g. show me the lowest instance count of my
// instances per day during the last month.
func (loc *InstanceMetricLocator) TimeSeries(endTime time.Time, granularity string, metrics []string, startTime time.Time, options rsapi.ApiParams) (*TimeSeriesMetricsResult, error) {
	var res *TimeSeriesMetricsResult
	if granularity == "" {
		return res, fmt.Errorf("granularity is required")
	}
	if len(metrics) == 0 {
		return res, fmt.Errorf("metrics is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":    endTime,
		"granularity": granularity,
		"metrics[]":   metrics,
		"start_time":  startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var intervalOpt = options["interval"]
	if intervalOpt != nil {
		queryParams["interval"] = intervalOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceMetric", "time_series")
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

// GET /api/instance_metrics/actions/grouped_time_series
//
// Calculates the metrics time series for instance usages in a time period allowing different
// time buckets (hour, 3 days, month, etc.) and groups them into specified breakdown
// categories, e.g. show me the lowest instance count of my instances per day during the last
// month grouped by accounts.
func (loc *InstanceMetricLocator) GroupedTimeSeries(endTime time.Time, granularity string, group []string, metrics []string, startTime time.Time, options rsapi.ApiParams) (*TimeSeriesMetricsResult, error) {
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
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":    endTime,
		"granularity": granularity,
		"group[]":     group,
		"metrics[]":   metrics,
		"start_time":  startTime,
	}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var intervalOpt = options["interval"]
	if intervalOpt != nil {
		queryParams["interval"] = intervalOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceMetric", "grouped_time_series")
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

// GET /api/instance_metrics/actions/current_count
//
// Returns the count of currently running instances.
func (loc *InstanceMetricLocator) CurrentCount(options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var instanceFiltersOpt = options["instance_filters"]
	if instanceFiltersOpt != nil {
		queryParams["instance_filters[]"] = instanceFiltersOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceMetric", "current_count")
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

// InstanceUsagePeriod resource locator, exposes resource actions.
type InstanceUsagePeriodLocator struct {
	UrlResolver
	api *Api
}

// InstanceUsagePeriod resource locator factory
func (api *Api) InstanceUsagePeriodLocator(href string) *InstanceUsagePeriodLocator {
	return &InstanceUsagePeriodLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/instance_usage_periods
//
// Gets the instance usage periods of instances.
func (loc *InstanceUsagePeriodLocator) Index(instanceUsagePeriodFilters []*Filter, options rsapi.ApiParams) (*InstanceUsagePeriod, error) {
	var res *InstanceUsagePeriod
	if len(instanceUsagePeriodFilters) == 0 {
		return res, fmt.Errorf("instanceUsagePeriodFilters is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"instance_usage_period_filters[]": instanceUsagePeriodFilters,
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("InstanceUsagePeriod", "index")
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

/******  Pattern ******/

// Patterns describe operations in usage, and can be applied to InstanceCombinations in Scenarios to model changes in the cost.
// A pattern can only be applied to an InstanceCombination once.
type Pattern struct {
}

//===== Locator

// Pattern resource locator, exposes resource actions.
type PatternLocator struct {
	UrlResolver
	api *Api
}

// Pattern resource locator factory
func (api *Api) PatternLocator(href string) *PatternLocator {
	return &PatternLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/patterns
//
// Create a new Pattern.
func (loc *PatternLocator) Create(months string, name string, operation string, type_ string, value float64, years string, options rsapi.ApiParams) (*PatternLocator, error) {
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
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"months":    months,
		"name":      name,
		"operation": operation,
		"type":      type_,
		"value":     value,
		"years":     years,
	}
	var summaryOpt = options["summary"]
	if summaryOpt != nil {
		payloadParams["summary"] = summaryOpt
	}
	uri, err := loc.Url("Pattern", "create")
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
		return &PatternLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/patterns
//
// List all Patterns.
func (loc *PatternLocator) Index(options rsapi.ApiParams) (*Pattern, error) {
	var res *Pattern
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Pattern", "index")
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

// GET /api/patterns/:id
//
// Show a specific Pattern.
func (loc *PatternLocator) Show(options rsapi.ApiParams) (*Pattern, error) {
	var res *Pattern
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Pattern", "show")
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

// PATCH /api/patterns/:id
//
// Update the provided attributes of a Pattern.
func (loc *PatternLocator) Update(options rsapi.ApiParams) (*Pattern, error) {
	var res *Pattern
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var monthsOpt = options["months"]
	if monthsOpt != nil {
		payloadParams["months"] = monthsOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var operationOpt = options["operation"]
	if operationOpt != nil {
		payloadParams["operation"] = operationOpt
	}
	var summaryOpt = options["summary"]
	if summaryOpt != nil {
		payloadParams["summary"] = summaryOpt
	}
	var type_Opt = options["type"]
	if type_Opt != nil {
		payloadParams["type"] = type_Opt
	}
	var valueOpt = options["value"]
	if valueOpt != nil {
		payloadParams["value"] = valueOpt
	}
	var yearsOpt = options["years"]
	if yearsOpt != nil {
		payloadParams["years"] = yearsOpt
	}
	uri, err := loc.Url("Pattern", "update")
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

// DELETE /api/patterns/:id
//
// Delete a Pattern.
func (loc *PatternLocator) Destroy() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Pattern", "destroy")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /api/patterns/actions/create_defaults
//
// Create the following commonly used default Patterns: Increase by 2% every month,
// Increase by 5% every month, Increase by 10% every month, Increase by 15% every month,
// Increase by 500% during Nov - Dec, Increase by 200% during Jan - Feb, Decrease by 2% every month,
// Decrease by 5% every month, Decrease by 10% every month, Decrease by 15% every month, Add 1 every month.
func (loc *PatternLocator) CreateDefaults(options rsapi.ApiParams) (*Pattern, error) {
	var res *Pattern
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Pattern", "create_defaults")
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

/******  ReservedInstance ******/

// Enables you to get details of existing AWS ReservedInstances and some metrics about their utilization.
type ReservedInstance struct {
}

//===== Locator

// ReservedInstance resource locator, exposes resource actions.
type ReservedInstanceLocator struct {
	UrlResolver
	api *Api
}

// ReservedInstance resource locator factory
func (api *Api) ReservedInstanceLocator(href string) *ReservedInstanceLocator {
	return &ReservedInstanceLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/reserved_instances
//
// Gets Reserved Instances that overlap with the requested time period.
func (loc *ReservedInstanceLocator) Index(endTime time.Time, startTime time.Time, options rsapi.ApiParams) (*ReservedInstance, error) {
	var res *ReservedInstance
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		queryParams["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstance", "index")
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

// GET /api/reserved_instances/actions/count
//
// Gets the count of Reserved Instances that overlap with the requested time period.
func (loc *ReservedInstanceLocator) Count(endTime time.Time, startTime time.Time, options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		queryParams["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstance", "count")
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
	res = string(respBody)
	return res, err
}

// GET /api/reserved_instances/actions/exist
//
// Checks if any Reserved Instances overlap with the requested time period.
func (loc *ReservedInstanceLocator) Exist(options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var endTimeOpt = options["end_time"]
	if endTimeOpt != nil {
		queryParams["end_time"] = endTimeOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		queryParams["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var startTimeOpt = options["start_time"]
	if startTimeOpt != nil {
		queryParams["start_time"] = startTimeOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstance", "exist")
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
	res = string(respBody)
	return res, err
}

// GET /api/reserved_instances/actions/export
//
// Exports the Reserved Instances that overlap with the requested time period in CSV format.
func (loc *ReservedInstanceLocator) Export(endTime time.Time, startTime time.Time, options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		queryParams["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstance", "export")
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
	res = string(respBody)
	return res, err
}

// GET /api/reserved_instances/actions/filter_options
//
// Gets the filter options for Reserved Instances that overlap with the requested time period.
func (loc *ReservedInstanceLocator) FilterOptions(endTime time.Time, startTime time.Time, options rsapi.ApiParams) (*Filter, error) {
	var res *Filter
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"end_time":   endTime,
		"start_time": startTime,
	}
	var filterTypesOpt = options["filter_types"]
	if filterTypesOpt != nil {
		queryParams["filter_types[]"] = filterTypesOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		queryParams["offset"] = offsetOpt
	}
	var orderOpt = options["order"]
	if orderOpt != nil {
		queryParams["order[]"] = orderOpt
	}
	var reservedInstanceFiltersOpt = options["reserved_instance_filters"]
	if reservedInstanceFiltersOpt != nil {
		queryParams["reserved_instance_filters[]"] = reservedInstanceFiltersOpt
	}
	var searchTermOpt = options["search_term"]
	if searchTermOpt != nil {
		queryParams["search_term"] = searchTermOpt
	}
	var timezoneOpt = options["timezone"]
	if timezoneOpt != nil {
		queryParams["timezone"] = timezoneOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstance", "filter_options")
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

/******  ReservedInstancePurchase ******/

// ReservedInstancePurchases can be applied to InstanceCombinations in Scenarios to model changes in the cost. These are not actually purchased in the cloud and are only used for cost simulation purposes.
type ReservedInstancePurchase struct {
}

//===== Locator

// ReservedInstancePurchase resource locator, exposes resource actions.
type ReservedInstancePurchaseLocator struct {
	UrlResolver
	api *Api
}

// ReservedInstancePurchase resource locator factory
func (api *Api) ReservedInstancePurchaseLocator(href string) *ReservedInstancePurchaseLocator {
	return &ReservedInstancePurchaseLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases
//
// Create a new ReservedInstancePurchase. This is not actually purchased in the cloud and is only used for cost simulation purposes.
func (loc *ReservedInstancePurchaseLocator) Create(autoRenew bool, duration int, offeringType string, quantity int, startDate time.Time, options rsapi.ApiParams) (*ReservedInstancePurchaseLocator, error) {
	var res *ReservedInstancePurchaseLocator
	if offeringType == "" {
		return res, fmt.Errorf("offeringType is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"auto_renew":    autoRenew,
		"duration":      duration,
		"offering_type": offeringType,
		"quantity":      quantity,
		"start_date":    startDate,
	}
	uri, err := loc.Url("ReservedInstancePurchase", "create")
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
		return &ReservedInstancePurchaseLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases
//
// List all ReservedInstancePurchases for the InstanceCombination.
func (loc *ReservedInstancePurchaseLocator) Index(options rsapi.ApiParams) (*ReservedInstancePurchase, error) {
	var res *ReservedInstancePurchase
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstancePurchase", "index")
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

// GET /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases/:id
//
// Show a specific ReservedInstancePurchase.
func (loc *ReservedInstancePurchaseLocator) Show(options rsapi.ApiParams) (*ReservedInstancePurchase, error) {
	var res *ReservedInstancePurchase
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstancePurchase", "show")
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

// PATCH /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases/:id
//
// Update the provided attributes of a ReservedInstancePurchase.
func (loc *ReservedInstancePurchaseLocator) Update(options rsapi.ApiParams) (*ReservedInstancePurchase, error) {
	var res *ReservedInstancePurchase
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var autoRenewOpt = options["auto_renew"]
	if autoRenewOpt != nil {
		payloadParams["auto_renew"] = autoRenewOpt
	}
	var durationOpt = options["duration"]
	if durationOpt != nil {
		payloadParams["duration"] = durationOpt
	}
	var offeringTypeOpt = options["offering_type"]
	if offeringTypeOpt != nil {
		payloadParams["offering_type"] = offeringTypeOpt
	}
	var quantityOpt = options["quantity"]
	if quantityOpt != nil {
		payloadParams["quantity"] = quantityOpt
	}
	var startDateOpt = options["start_date"]
	if startDateOpt != nil {
		payloadParams["start_date"] = startDateOpt
	}
	uri, err := loc.Url("ReservedInstancePurchase", "update")
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

// DELETE /api/scenarios/:scenario_id/instance_combinations/:instance_combination_id/reserved_instance_purchases/:id
//
// Delete a ReservedInstancePurchase. This is not actually deleted in the cloud and is only used for cost simulation purposes.
func (loc *ReservedInstancePurchaseLocator) Destroy() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ReservedInstancePurchase", "destroy")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
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

// Scenario resource locator, exposes resource actions.
type ScenarioLocator struct {
	UrlResolver
	api *Api
}

// Scenario resource locator factory
func (api *Api) ScenarioLocator(href string) *ScenarioLocator {
	return &ScenarioLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/scenarios
//
// Create a new Scenario.
func (loc *ScenarioLocator) Create(snapshotTimestamp time.Time, options rsapi.ApiParams) (*ScenarioLocator, error) {
	var res *ScenarioLocator
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"snapshot_timestamp": snapshotTimestamp,
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		payloadParams["filters"] = filtersOpt
	}
	var isBlankOpt = options["is_blank"]
	if isBlankOpt != nil {
		payloadParams["is_blank"] = isBlankOpt
	}
	var isPersistedOpt = options["is_persisted"]
	if isPersistedOpt != nil {
		payloadParams["is_persisted"] = isPersistedOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var privateCloudInstanceCountOpt = options["private_cloud_instance_count"]
	if privateCloudInstanceCountOpt != nil {
		payloadParams["private_cloud_instance_count"] = privateCloudInstanceCountOpt
	}
	uri, err := loc.Url("Scenario", "create")
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
		return &ScenarioLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/scenarios
//
// List all Scenarios.
func (loc *ScenarioLocator) Index(options rsapi.ApiParams) (*Scenario, error) {
	var res *Scenario
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var includeNonPersistedOpt = options["include_non_persisted"]
	if includeNonPersistedOpt != nil {
		queryParams["include_non_persisted"] = includeNonPersistedOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Scenario", "index")
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

// GET /api/scenarios/:id
//
// Show a specific Scenario.
func (loc *ScenarioLocator) Show(options rsapi.ApiParams) (*Scenario, error) {
	var res *Scenario
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Scenario", "show")
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

// PATCH /api/scenarios/:id
//
// Update the provided attributes of a Scenario.
func (loc *ScenarioLocator) Update(options rsapi.ApiParams) (*Scenario, error) {
	var res *Scenario
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var isPersistedOpt = options["is_persisted"]
	if isPersistedOpt != nil {
		payloadParams["is_persisted"] = isPersistedOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var privateCloudInstanceCountOpt = options["private_cloud_instance_count"]
	if privateCloudInstanceCountOpt != nil {
		payloadParams["private_cloud_instance_count"] = privateCloudInstanceCountOpt
	}
	var snapshotTimestampOpt = options["snapshot_timestamp"]
	if snapshotTimestampOpt != nil {
		payloadParams["snapshot_timestamp"] = snapshotTimestampOpt
	}
	uri, err := loc.Url("Scenario", "update")
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

// DELETE /api/scenarios/:id
//
// Delete a Scenario.
func (loc *ScenarioLocator) Destroy() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Scenario", "destroy")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// GET /api/scenarios/:id/actions/forecast
//
// Run a simulation to generate a 3-year forecast showing the `average_instance_count`, `instance_upfront_cost`,
// `instance_usage_cost` and `instance_recurring_cost` metrics. This call might get major changes so it's best to avoid using it currently.
// If there are missing prices for any of the InstanceCombinations then these metrics will be excluded from the results for that InstanceCombination.
func (loc *ScenarioLocator) Forecast(options rsapi.ApiParams) (*TimeSeriesMetricsResult, error) {
	var res *TimeSeriesMetricsResult
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Scenario", "forecast")
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

/******  ScheduledReport ******/

// ScheduledReports are emailed to you, and include usage, cost, and the change from the previous reporting period.
// These emails include links to AnalysisSnapshots, which are generated automatically by us.
type ScheduledReport struct {
}

//===== Locator

// ScheduledReport resource locator, exposes resource actions.
type ScheduledReportLocator struct {
	UrlResolver
	api *Api
}

// ScheduledReport resource locator factory
func (api *Api) ScheduledReportLocator(href string) *ScheduledReportLocator {
	return &ScheduledReportLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/scheduled_reports
//
// Create a new ScheduledReport.
func (loc *ScheduledReportLocator) Create(frequency string, name string, options rsapi.ApiParams) (*ScheduledReportLocator, error) {
	var res *ScheduledReportLocator
	if frequency == "" {
		return res, fmt.Errorf("frequency is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"frequency": frequency,
		"name":      name,
	}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		payloadParams["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		payloadParams["attach_csv"] = attachCsvOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		payloadParams["filters"] = filtersOpt
	}
	uri, err := loc.Url("ScheduledReport", "create")
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
		return &ScheduledReportLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/scheduled_reports
//
// List all ScheduledReports.
func (loc *ScheduledReportLocator) Index(options rsapi.ApiParams) (*ScheduledReport, error) {
	var res *ScheduledReport
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledReport", "index")
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

// GET /api/scheduled_reports/:id
//
// Show a specific ScheduledReport.
func (loc *ScheduledReportLocator) Show(options rsapi.ApiParams) (*ScheduledReport, error) {
	var res *ScheduledReport
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledReport", "show")
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

// PATCH /api/scheduled_reports/:id
//
// Update the provided attributes of a ScheduledReport.
func (loc *ScheduledReportLocator) Update(options rsapi.ApiParams) (*ScheduledReport, error) {
	var res *ScheduledReport
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var additionalEmailsOpt = options["additional_emails"]
	if additionalEmailsOpt != nil {
		payloadParams["additional_emails"] = additionalEmailsOpt
	}
	var attachCsvOpt = options["attach_csv"]
	if attachCsvOpt != nil {
		payloadParams["attach_csv"] = attachCsvOpt
	}
	var frequencyOpt = options["frequency"]
	if frequencyOpt != nil {
		payloadParams["frequency"] = frequencyOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	uri, err := loc.Url("ScheduledReport", "update")
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

// DELETE /api/scheduled_reports/:id
//
// Delete a ScheduledReport.
func (loc *ScheduledReportLocator) Destroy() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledReport", "destroy")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /api/scheduled_reports/actions/create_defaults
//
// Create the default Scheduled Report: a weekly report with no filters
func (loc *ScheduledReportLocator) CreateDefaults(options rsapi.ApiParams) (*ScheduledReport, error) {
	var res *ScheduledReport
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("ScheduledReport", "create_defaults")
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

/******  TempInstancePrice ******/

// This is a temporary API call that can be used by the Cloud Analytics UI until the
// Pricing Service is live, at which point this API call will be deleted. This is not included in the public docs.
type TempInstancePrice struct {
}

//===== Locator

// TempInstancePrice resource locator, exposes resource actions.
type TempInstancePriceLocator struct {
	UrlResolver
	api *Api
}

// TempInstancePrice resource locator factory
func (api *Api) TempInstancePriceLocator(href string) *TempInstancePriceLocator {
	return &TempInstancePriceLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/temp_instance_prices
//
// Returns a JSON blob with all prices for Scenario Builder.
func (loc *TempInstancePriceLocator) Index() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("TempInstancePrice", "index")
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
	res = string(respBody)
	return res, err
}

/******  User ******/

// Users can have various permissions on multiple accounts. Users with admin permissions in an account
// can modify that account's users. This resource is not included in the public docs.
type User struct {
}

//===== Locator

// User resource locator, exposes resource actions.
type UserLocator struct {
	UrlResolver
	api *Api
}

// User resource locator factory
func (api *Api) UserLocator(href string) *UserLocator {
	return &UserLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/users
//
// Create a new user with the requested permissions in the requested accounts, and emails
// them the login details. Returns an error if the user already exists.
func (loc *UserLocator) Create(accounts []*UserAccounts, email string, options rsapi.ApiParams) (*UserLocator, error) {
	var res *UserLocator
	if len(accounts) == 0 {
		return res, fmt.Errorf("accounts is required")
	}
	if email == "" {
		return res, fmt.Errorf("email is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"accounts": accounts,
		"email":    email,
	}
	uri, err := loc.Url("User", "create")
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
		return &UserLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/users
//
// List all users.
func (loc *UserLocator) Index(options rsapi.ApiParams) (*User, error) {
	var res *User
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("User", "index")
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

// GET /api/users/:id
//
// Show a specific user.
func (loc *UserLocator) Show(options rsapi.ApiParams) (*User, error) {
	var res *User
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("User", "show")
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

// PATCH /api/users/:id
//
// Update a specific user's account permissions.
// This cannot be used to update other user parameters such as their name or password.
func (loc *UserLocator) Update(options rsapi.ApiParams) (*User, error) {
	var res *User
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var accountsOpt = options["accounts"]
	if accountsOpt != nil {
		payloadParams["accounts"] = accountsOpt
	}
	uri, err := loc.Url("User", "update")
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

// POST /api/users/actions/invite
//
// Invites a user to the requested account and gives them the required permissions
// so they can add/edit cloud credentials, the user is created if they don't already exist.
// This is used during new user onboarding as the user who signs-up might not be the person who has
// the cloud credentials required to connect their clouds to RightScale.
func (loc *UserLocator) Invite(options rsapi.ApiParams) (*User, error) {
	var res *User
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var accountIdOpt = options["account_id"]
	if accountIdOpt != nil {
		payloadParams["account_id"] = accountIdOpt
	}
	var emailOpt = options["email"]
	if emailOpt != nil {
		payloadParams["email"] = emailOpt
	}
	var messageOpt = options["message"]
	if messageOpt != nil {
		payloadParams["message"] = messageOpt
	}
	uri, err := loc.Url("User", "invite")
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

/******  UserSetting ******/

// Used by the Cloud Analytics UI to keep track of various UI states.
type UserSetting struct {
}

//===== Locator

// UserSetting resource locator, exposes resource actions.
type UserSettingLocator struct {
	UrlResolver
	api *Api
}

// UserSetting resource locator factory
func (api *Api) UserSettingLocator(href string) *UserSettingLocator {
	return &UserSettingLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/user_settings
//
// List the UserSettings.
func (loc *UserSettingLocator) Show(options rsapi.ApiParams) (*UserSetting, error) {
	var res *UserSetting
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserSetting", "show")
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

// PATCH /api/user_settings
//
// Update the provided attributes of UserSettings.
func (loc *UserSettingLocator) Update(options rsapi.ApiParams) (*UserSetting, error) {
	var res *UserSetting
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var dateRangeOpt = options["date_range"]
	if dateRangeOpt != nil {
		payloadParams["date_range"] = dateRangeOpt
	}
	var dismissedDialogsOpt = options["dismissed_dialogs"]
	if dismissedDialogsOpt != nil {
		payloadParams["dismissed_dialogs"] = dismissedDialogsOpt
	}
	var excludedTagTypesOpt = options["excluded_tag_types"]
	if excludedTagTypesOpt != nil {
		payloadParams["excluded_tag_types"] = excludedTagTypesOpt
	}
	var filtersOpt = options["filters"]
	if filtersOpt != nil {
		payloadParams["filters"] = filtersOpt
	}
	var granularityOpt = options["granularity"]
	if granularityOpt != nil {
		payloadParams["granularity"] = granularityOpt
	}
	var mainMenuVisibilityOpt = options["main_menu_visibility"]
	if mainMenuVisibilityOpt != nil {
		payloadParams["main_menu_visibility"] = mainMenuVisibilityOpt
	}
	var metricsOpt = options["metrics"]
	if metricsOpt != nil {
		payloadParams["metrics"] = metricsOpt
	}
	var moduleStatesOpt = options["module_states"]
	if moduleStatesOpt != nil {
		payloadParams["module_states"] = moduleStatesOpt
	}
	var onboardingStatusOpt = options["onboarding_status"]
	if onboardingStatusOpt != nil {
		payloadParams["onboarding_status"] = onboardingStatusOpt
	}
	var selectedCloudVendorNamesOpt = options["selected_cloud_vendor_names"]
	if selectedCloudVendorNamesOpt != nil {
		payloadParams["selected_cloud_vendor_names"] = selectedCloudVendorNamesOpt
	}
	var sortingOpt = options["sorting"]
	if sortingOpt != nil {
		payloadParams["sorting"] = sortingOpt
	}
	var tableColumnVisibilityOpt = options["table_column_visibility"]
	if tableColumnVisibilityOpt != nil {
		payloadParams["table_column_visibility"] = tableColumnVisibilityOpt
	}
	uri, err := loc.Url("UserSetting", "update")
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
	CreatedAt                   time.Time      `json:"created_at,omitempty"`
	CreatedBy                   string         `json:"created_by,omitempty"`
	EndTime                     time.Time      `json:"end_time,omitempty"`
	ExcludedTagTypes            []string       `json:"excluded_tag_types,omitempty"`
	Filters                     []*Filter      `json:"filters,omitempty"`
	Granularity                 string         `json:"granularity,omitempty"`
	Href                        string         `json:"href,omitempty"`
	IsComparison                bool           `json:"is_comparison,omitempty"`
	Kind                        string         `json:"kind,omitempty"`
	Metrics                     []string       `json:"metrics,omitempty"`
	MissingAccessToSomeAccounts bool           `json:"missing_access_to_some_accounts,omitempty"`
	ModuleStates                []*ModuleState `json:"module_states,omitempty"`
	StartTime                   time.Time      `json:"start_time,omitempty"`
	UpdatedAt                   time.Time      `json:"updated_at,omitempty"`
	Uuid                        string         `json:"uuid,omitempty"`
}

type BudgetAlertParam struct {
	AdditionalEmails []string            `json:"additional_emails,omitempty"`
	AttachCsv        bool                `json:"attach_csv,omitempty"`
	Budget           *ReturnBudgetStruct `json:"budget,omitempty"`
	CreatedAt        time.Time           `json:"created_at,omitempty"`
	Filters          []*Filter           `json:"filters,omitempty"`
	Frequency        string              `json:"frequency,omitempty"`
	Href             string              `json:"href,omitempty"`
	Id               int                 `json:"id,omitempty"`
	Kind             string              `json:"kind,omitempty"`
	Name             string              `json:"name,omitempty"`
	Type_            string              `json:"type,omitempty"`
	UpdatedAt        time.Time           `json:"updated_at,omitempty"`
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
	Company   string    `json:"company,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	Id        int       `json:"id,omitempty"`
	Kind      string    `json:"kind,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Timezone  string    `json:"timezone,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type DateRangeStruct struct {
	EndTime      time.Time `json:"end_time,omitempty"`
	IsComparison bool      `json:"is_comparison,omitempty"`
	StartTime    time.Time `json:"start_time,omitempty"`
	Type_        string    `json:"type,omitempty"`
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
	CreatedAt                 time.Time                        `json:"created_at,omitempty"`
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
	UpdatedAt                 time.Time                        `json:"updated_at,omitempty"`
}

type InstanceParam struct {
	AccountId                         int       `json:"account_id,omitempty"`
	AccountName                       string    `json:"account_name,omitempty"`
	CloudId                           int       `json:"cloud_id,omitempty"`
	CloudName                         string    `json:"cloud_name,omitempty"`
	CloudVendorName                   string    `json:"cloud_vendor_name,omitempty"`
	DatacenterKey                     string    `json:"datacenter_key,omitempty"`
	DatacenterName                    string    `json:"datacenter_name,omitempty"`
	DeploymentId                      int       `json:"deployment_id,omitempty"`
	DeploymentName                    string    `json:"deployment_name,omitempty"`
	EstimatedCostForPeriod            float64   `json:"estimated_cost_for_period,omitempty"`
	EstimatedManagedRcuCountForPeriod float64   `json:"estimated_managed_rcu_count_for_period,omitempty"`
	IncarnatorId                      int       `json:"incarnator_id,omitempty"`
	IncarnatorType                    string    `json:"incarnator_type,omitempty"`
	InstanceEndAt                     time.Time `json:"instance_end_at,omitempty"`
	InstanceKey                       string    `json:"instance_key,omitempty"`
	InstanceName                      string    `json:"instance_name,omitempty"`
	InstanceRsid                      string    `json:"instance_rsid,omitempty"`
	InstanceStartAt                   time.Time `json:"instance_start_at,omitempty"`
	InstanceTypeKey                   string    `json:"instance_type_key,omitempty"`
	InstanceTypeName                  string    `json:"instance_type_name,omitempty"`
	InstanceUid                       string    `json:"instance_uid,omitempty"`
	Kind                              string    `json:"kind,omitempty"`
	Platform                          string    `json:"platform,omitempty"`
	ProvisionedByUserEmail            string    `json:"provisioned_by_user_email,omitempty"`
	ProvisionedByUserId               int       `json:"provisioned_by_user_id,omitempty"`
	ServerTemplateId                  int       `json:"server_template_id,omitempty"`
	ServerTemplateName                string    `json:"server_template_name,omitempty"`
	State                             string    `json:"state,omitempty"`
	Tags                              []*Tag    `json:"tags,omitempty"`
	TotalUsageHours                   float64   `json:"total_usage_hours,omitempty"`
}

type InstanceUsagePeriodParam struct {
	EstimatedCost            float64   `json:"estimated_cost,omitempty"`
	EstimatedManagedRcuCount float64   `json:"estimated_managed_rcu_count,omitempty"`
	HourlyPrice              float64   `json:"hourly_price,omitempty"`
	InstanceKey              string    `json:"instance_key,omitempty"`
	InstanceTypeName         string    `json:"instance_type_name,omitempty"`
	Kind                     string    `json:"kind,omitempty"`
	PricingType              string    `json:"pricing_type,omitempty"`
	RcuRate                  float64   `json:"rcu_rate,omitempty"`
	ReservationUid           string    `json:"reservation_uid,omitempty"`
	UsageEndAt               time.Time `json:"usage_end_at,omitempty"`
	UsageStartAt             time.Time `json:"usage_start_at,omitempty"`
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
	CreatedAt time.Time        `json:"created_at,omitempty"`
	Href      string           `json:"href,omitempty"`
	Id        int              `json:"id,omitempty"`
	Kind      string           `json:"kind,omitempty"`
	Months    string           `json:"months,omitempty"`
	Name      string           `json:"name,omitempty"`
	Operation string           `json:"operation,omitempty"`
	Scenarios []*ScenarioParam `json:"scenarios,omitempty"`
	Summary   string           `json:"summary,omitempty"`
	Type_     string           `json:"type,omitempty"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
	Value     float64          `json:"value,omitempty"`
	Years     string           `json:"years,omitempty"`
}

type ReservedInstanceParam struct {
	AccountId             int       `json:"account_id,omitempty"`
	AccountName           string    `json:"account_name,omitempty"`
	CloudId               int       `json:"cloud_id,omitempty"`
	CloudName             string    `json:"cloud_name,omitempty"`
	CloudVendorName       string    `json:"cloud_vendor_name,omitempty"`
	CostSaved             float64   `json:"cost_saved,omitempty"`
	DatacenterKey         string    `json:"datacenter_key,omitempty"`
	DatacenterName        string    `json:"datacenter_name,omitempty"`
	Duration              int       `json:"duration,omitempty"`
	EndTime               time.Time `json:"end_time,omitempty"`
	InstanceCount         int       `json:"instance_count,omitempty"`
	InstanceTypeKey       string    `json:"instance_type_key,omitempty"`
	InstanceTypeName      string    `json:"instance_type_name,omitempty"`
	Kind                  string    `json:"kind,omitempty"`
	OfferingType          string    `json:"offering_type,omitempty"`
	Platform              string    `json:"platform,omitempty"`
	ReservationUid        string    `json:"reservation_uid,omitempty"`
	StartTime             time.Time `json:"start_time,omitempty"`
	State                 string    `json:"state,omitempty"`
	Tenancy               string    `json:"tenancy,omitempty"`
	UnusedRecurringCost   float64   `json:"unused_recurring_cost,omitempty"`
	UtilizationPercentage float64   `json:"utilization_percentage,omitempty"`
}

type ReservedInstancePurchaseLinks struct {
	InstanceCombination *InstanceCombinationParam `json:"instance_combination,omitempty"`
}

type ReservedInstancePurchaseParam struct {
	AutoRenew           bool                           `json:"auto_renew,omitempty"`
	CreatedAt           time.Time                      `json:"created_at,omitempty"`
	Duration            int                            `json:"duration,omitempty"`
	Href                string                         `json:"href,omitempty"`
	Id                  int                            `json:"id,omitempty"`
	InstanceCombination *InstanceCombinationParam      `json:"instance_combination,omitempty"`
	Kind                string                         `json:"kind,omitempty"`
	Links               *ReservedInstancePurchaseLinks `json:"links,omitempty"`
	OfferingType        string                         `json:"offering_type,omitempty"`
	Quantity            int                            `json:"quantity,omitempty"`
	StartDate           time.Time                      `json:"start_date,omitempty"`
	UpdatedAt           time.Time                      `json:"updated_at,omitempty"`
}

type ReturnBudgetStruct struct {
	Amount float64 `json:"amount,omitempty"`
	Period string  `json:"period,omitempty"`
}

type ReturnCurrentUserStruct struct {
	BetaEnabled                          bool      `json:"beta_enabled,omitempty"`
	CanSeeCostAndRcuMetrics              bool      `json:"can_see_cost_and_rcu_metrics,omitempty"`
	CanSeeManagedRcus                    bool      `json:"can_see_managed_rcus,omitempty"`
	CanSeeUnmanagedRcus                  bool      `json:"can_see_unmanaged_rcus,omitempty"`
	Company                              string    `json:"company,omitempty"`
	Email                                string    `json:"email,omitempty"`
	FirstLoginAt                         time.Time `json:"first_login_at,omitempty"`
	FirstName                            string    `json:"first_name,omitempty"`
	HasAdminOnAnyAccount                 bool      `json:"has_admin_on_any_account,omitempty"`
	HasCloudAnalyticsEnabledAccounts     bool      `json:"has_cloud_analytics_enabled_accounts,omitempty"`
	HasNonIpWhitelistedAccountsWithAdmin bool      `json:"has_non_ip_whitelisted_accounts_with_admin,omitempty"`
	HasOnlyExpiredAccounts               bool      `json:"has_only_expired_accounts,omitempty"`
	Id                                   int       `json:"id,omitempty"`
	IsCloudAnalyticsOnly                 bool      `json:"is_cloud_analytics_only,omitempty"`
	IsRightscaleEmployee                 bool      `json:"is_rightscale_employee,omitempty"`
	IsSelfserviceUser                    bool      `json:"is_selfservice_user,omitempty"`
	IsTeamUser                           bool      `json:"is_team_user,omitempty"`
	LastName                             string    `json:"last_name,omitempty"`
	NotificationMessage                  string    `json:"notification_message,omitempty"`
	Phone                                string    `json:"phone,omitempty"`
	SelfserviceUrl                       string    `json:"selfservice_url,omitempty"`
	Timezone                             string    `json:"timezone,omitempty"`
	TimezoneOffsetSeconds                int       `json:"timezone_offset_seconds,omitempty"`
	TrialEndDate                         time.Time `json:"trial_end_date,omitempty"`
}

type ReturnGoogleAnalyticsStruct struct {
	AccountId  string `json:"account_id,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}

type ReturnUserSettingsDateRangeStruct struct {
	EndTime      time.Time `json:"end_time,omitempty"`
	IsComparison bool      `json:"is_comparison,omitempty"`
	StartTime    time.Time `json:"start_time,omitempty"`
	Type_        string    `json:"type,omitempty"`
}

type ScenarioParam struct {
	CreatedAt                 time.Time                   `json:"created_at,omitempty"`
	Filters                   []*Filter                   `json:"filters,omitempty"`
	HistoricMetricsResults    []*TimeSeriesMetricsResult  `json:"historic_metrics_results,omitempty"`
	Href                      string                      `json:"href,omitempty"`
	Id                        int                         `json:"id,omitempty"`
	InstanceCombinations      []*InstanceCombinationParam `json:"instance_combinations,omitempty"`
	IsPersisted               bool                        `json:"is_persisted,omitempty"`
	Kind                      string                      `json:"kind,omitempty"`
	Name                      string                      `json:"name,omitempty"`
	PrivateCloudInstanceCount int                         `json:"private_cloud_instance_count,omitempty"`
	SnapshotTimestamp         time.Time                   `json:"snapshot_timestamp,omitempty"`
	UpdatedAt                 time.Time                   `json:"updated_at,omitempty"`
}

type ScheduledReportParam struct {
	AdditionalEmails []string  `json:"additional_emails,omitempty"`
	AttachCsv        bool      `json:"attach_csv,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	Filters          []*Filter `json:"filters,omitempty"`
	Frequency        string    `json:"frequency,omitempty"`
	Href             string    `json:"href,omitempty"`
	Id               int       `json:"id,omitempty"`
	Kind             string    `json:"kind,omitempty"`
	Name             string    `json:"name,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
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
	Timestamp time.Time        `json:"timestamp,omitempty"`
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
