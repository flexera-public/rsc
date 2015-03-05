//************************************************************************//
//                     RightScale API client
//
// Generated
// Mar 4, 2015 at 10:20pm (PST)
// Command:
// $ praxisgen -metadata=ss/ssc/restful_doc -output=ss/ssc -pkg=ssc -target=1.0 -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssc

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

/******  AccountPreference ******/

// The AccountPreference resource stores preferences that apply account-wide, such as UI customization settings and other settings.
// The Self-Service portal uses some of these preferences in the portal itself, and this resource allows you to extend the settings
// to use in your own integration.
type AccountPreference struct {
	CreatedBy  *User             `json:"created_by,omitempty"`
	GroupName  string            `json:"group_name,omitempty"`
	Href       string            `json:"href,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Name       string            `json:"name,omitempty"`
	Timestamps *TimestampsStruct `json:"timestamps,omitempty"`
	Value      string            `json:"value,omitempty"`
}

//===== Locator

// AccountPreference resource locator, exposes resource actions.
type AccountPreferenceLocator struct {
	UrlResolver
	api *Api
}

// AccountPreference resource locator factory
func (api *Api) AccountPreferenceLocator(href string) *AccountPreferenceLocator {
	return &AccountPreferenceLocator{UrlResolver(href), api}
}

//===== Actions

// GET /accounts/:account_id/account_preferences
// List the AccountPreferences for this account.
func (loc *AccountPreferenceLocator) Index(accountId string, options rsapi.ApiParams) ([]*AccountPreference, error) {
	var res []*AccountPreference
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var groupOpt = options["group"]
	if groupOpt != nil {
		queryParams["group"] = groupOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("AccountPreference", "index")
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

// GET /accounts/:account_id/account_preferences/:name
// Get details for a particular AccountPreference
func (loc *AccountPreferenceLocator) Show(accountId string, name string) (*AccountPreference, error) {
	var res *AccountPreference
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("AccountPreference", "show")
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

// POST /accounts/:account_id/account_preferences
// Create a new AccountPreference or update an existing AccountPreference with the new value
func (loc *AccountPreferenceLocator) Create(accountId string) (*AccountPreferenceLocator, error) {
	var res *AccountPreferenceLocator
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("AccountPreference", "create")
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
		return &AccountPreferenceLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /accounts/:account_id/account_preferences/:name
// Delete an AccountPreference
func (loc *AccountPreferenceLocator) Delete(accountId string, name string) error {
	if accountId == "" {
		return fmt.Errorf("accountId is required")
	}
	if name == "" {
		return fmt.Errorf("name is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("AccountPreference", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  Application ******/

// An Application is an element in the Catalog that can be launched by users. Applications are generally created by uploading CAT
// files to the Designer and publishing them to the Catalog, though they can also be created via API calls to the Catalog directly without
// going through Designer. If an Application was created from Designer through the publish action, it contains a link back to the Template
// resource in Designer.
// In the Self-Service portal, an Application is equivalent to an item in the Catalog. Most users have access to these Application resources
// and can launch them to create Executions in the Manager application.
type Application struct {
	CompiledCat        string            `json:"compiled_cat,omitempty"`
	CreatedBy          *User             `json:"created_by,omitempty"`
	Href               string            `json:"href,omitempty"`
	Id                 string            `json:"id,omitempty"`
	Kind               string            `json:"kind,omitempty"`
	LongDescription    string            `json:"long_description,omitempty"`
	Name               string            `json:"name,omitempty"`
	Parameters         []*Parameter      `json:"parameters,omitempty"`
	RequiredParameters []string          `json:"required_parameters,omitempty"`
	ScheduleRequired   bool              `json:"schedule_required,omitempty"`
	Schedules          []*Schedule       `json:"schedules,omitempty"`
	ShortDescription   string            `json:"short_description,omitempty"`
	TemplateInfo       *TemplateInfo     `json:"template_info,omitempty"`
	Timestamps         *TimestampsStruct `json:"timestamps,omitempty"`
}

//===== Locator

// Application resource locator, exposes resource actions.
type ApplicationLocator struct {
	UrlResolver
	api *Api
}

// Application resource locator factory
func (api *Api) ApplicationLocator(href string) *ApplicationLocator {
	return &ApplicationLocator{UrlResolver(href), api}
}

//===== Actions

// GET /catalogs/:catalog_id/applications
// List the Applications available in the specified Catalog.
func (loc *ApplicationLocator) Index(catalogId string, options rsapi.ApiParams) ([]*Application, error) {
	var res []*Application
	if catalogId == "" {
		return res, fmt.Errorf("catalogId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids"] = idsOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "index")
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

// GET /catalogs/:catalog_id/applications/:id
// Show detailed information about a given Application.
func (loc *ApplicationLocator) Show(catalogId string, id string, options rsapi.ApiParams) (*Application, error) {
	var res *Application
	if catalogId == "" {
		return res, fmt.Errorf("catalogId is required")
	}
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "show")
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

// POST /catalogs/:catalog_id/applications
// Create a new Application in the Catalog.
func (loc *ApplicationLocator) Create(catalogId string) (*ApplicationLocator, error) {
	var res *ApplicationLocator
	if catalogId == "" {
		return res, fmt.Errorf("catalogId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "create")
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
		return &ApplicationLocator{UrlResolver(location), loc.api}, nil
	}
}

// PUT /catalogs/:catalog_id/applications/:id
// Update the content of an existing Application.
func (loc *ApplicationLocator) Update(catalogId string, id string) error {
	if catalogId == "" {
		return fmt.Errorf("catalogId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "update")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// PUT /catalogs/:catalog_id/applications
// Update the content of multiple Applications.
func (loc *ApplicationLocator) MultiUpdate(catalogId string) error {
	if catalogId == "" {
		return fmt.Errorf("catalogId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "multi_update")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /catalogs/:catalog_id/applications/:id
// Delete an Application from the Catalog
func (loc *ApplicationLocator) Delete(catalogId string, id string) error {
	if catalogId == "" {
		return fmt.Errorf("catalogId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /catalogs/:catalog_id/applications
// Delete multiple Applications from the Catalog
func (loc *ApplicationLocator) MultiDelete(catalogId string, ids []string) error {
	if catalogId == "" {
		return fmt.Errorf("catalogId is required")
	}
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids": ids,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "multi_delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// GET /catalogs/:catalog_id/applications/:id/download
// Download the underlying CAT source of an Application.
func (loc *ApplicationLocator) Download(apiVersion string, catalogId string, id string) error {
	if apiVersion == "" {
		return fmt.Errorf("apiVersion is required")
	}
	if catalogId == "" {
		return fmt.Errorf("catalogId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"api_version": apiVersion,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "download")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /catalogs/:catalog_id/applications/:id/actions/launch
// Launches an Application by creating an Execution with ScheduledActions as needed to match the optional Schedule provided.
func (loc *ApplicationLocator) Launch(catalogId string, id string) error {
	if catalogId == "" {
		return fmt.Errorf("catalogId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Application", "launch")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  NotificationRule ******/

// A notification rule describes which notification should be created
// when events occur in the system. Events may be generated when an
// execution status changes or when an operation fails for example.
// A rule has a source which can be a specific resource or a group of
// resources (described via a link-like syntax), a target which
// corresponds to a user (for now) and a minimum severity used to filter
// out events with lower severities.
type NotificationRule struct {
	AccountId   string            `json:"account_id,omitempty"`
	Href        string            `json:"href,omitempty"`
	Id          string            `json:"id,omitempty"`
	Kind        string            `json:"kind,omitempty"`
	MinSeverity string            `json:"min_severity,omitempty"`
	Priority    int               `json:"priority,omitempty"`
	Source      string            `json:"source,omitempty"`
	Target      string            `json:"target,omitempty"`
	Timestamps  *TimestampsStruct `json:"timestamps,omitempty"`
}

//===== Locator

// NotificationRule resource locator, exposes resource actions.
type NotificationRuleLocator struct {
	UrlResolver
	api *Api
}

// NotificationRule resource locator factory
func (api *Api) NotificationRuleLocator(href string) *NotificationRuleLocator {
	return &NotificationRuleLocator{UrlResolver(href), api}
}

//===== Actions

// GET /accounts/:account_id/notification_rules
// List all notification rules, potentially filtering by a collection of resources.
func (loc *NotificationRuleLocator) Index(accountId string, source string, options rsapi.ApiParams) ([]*NotificationRule, error) {
	var res []*NotificationRule
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	if source == "" {
		return res, fmt.Errorf("source is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"source": source,
	}
	var targetsOpt = options["targets"]
	if targetsOpt != nil {
		queryParams["targets"] = targetsOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("NotificationRule", "index")
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

// POST /accounts/:account_id/notification_rules
// Create one notification rule for a specific target and source.
// The source must be unique in the scope of target and account.
func (loc *NotificationRuleLocator) Create(accountId string, options rsapi.ApiParams) (*NotificationRuleLocator, error) {
	var res *NotificationRuleLocator
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("NotificationRule", "create")
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
		return &NotificationRuleLocator{UrlResolver(location), loc.api}, nil
	}
}

// PATCH /accounts/:account_id/notification_rules/:id
// Change min severity of existing rule
func (loc *NotificationRuleLocator) Patch(accountId string, id string) error {
	if accountId == "" {
		return fmt.Errorf("accountId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("NotificationRule", "patch")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// GET /accounts/:account_id/notification_rules/:id
// Show one notification rule.
func (loc *NotificationRuleLocator) Show(accountId string, id string) (*NotificationRule, error) {
	var res *NotificationRule
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("NotificationRule", "show")
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

// DELETE /accounts/:account_id/notification_rules/:id
// Delete one notification rule.
func (loc *NotificationRuleLocator) Delete(accountId string, id string) error {
	if accountId == "" {
		return fmt.Errorf("accountId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("NotificationRule", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /accounts/:account_id/notification_rules
// Delete one or more notification rules by id or source and target.
func (loc *NotificationRuleLocator) MultiDelete(accountId string) error {
	if accountId == "" {
		return fmt.Errorf("accountId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("NotificationRule", "multi_delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  UserPreference ******/

// The UserPreference resource stores preferences on a per user basis, such as default notification preference.
// The Self-Service portal uses these preferences in the portal.
type UserPreference struct {
	CreatedBy          *User               `json:"created_by,omitempty"`
	Href               string              `json:"href,omitempty"`
	Id                 string              `json:"id,omitempty"`
	Kind               string              `json:"kind,omitempty"`
	Timestamps         *TimestampsStruct   `json:"timestamps,omitempty"`
	UserId             int                 `json:"user_id,omitempty"`
	UserPreferenceInfo *UserPreferenceInfo `json:"user_preference_info,omitempty"`
	Value              string              `json:"value,omitempty"`
}

//===== Locator

// UserPreference resource locator, exposes resource actions.
type UserPreferenceLocator struct {
	UrlResolver
	api *Api
}

// UserPreference resource locator factory
func (api *Api) UserPreferenceLocator(href string) *UserPreferenceLocator {
	return &UserPreferenceLocator{UrlResolver(href), api}
}

//===== Actions

// GET /accounts/:account_id/user_preferences
// List the UserPreference for users in this account.
// Only administrators and infrastructure users may request the preferences of other users.
// Users who are not members of the admin role need to specify a filter with their ID in order to retrieve their preferences.
func (loc *UserPreferenceLocator) Index(accountId string, options rsapi.ApiParams) ([]*UserPreference, error) {
	var res []*UserPreference
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserPreference", "index")
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

// GET /accounts/:account_id/user_preferences/:id
// Get details for a particular UserPreference
func (loc *UserPreferenceLocator) Show(accountId string, id string, options rsapi.ApiParams) (*UserPreference, error) {
	var res *UserPreference
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserPreference", "show")
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

// POST /accounts/:account_id/user_preferences
// Create a new UserPreference.
// Multiple resources can be created at once with a multipart request.
// Values are validated with the corresponding UserPreferenceInfo.
func (loc *UserPreferenceLocator) Create(accountId string) (*UserPreferenceLocator, error) {
	var res *UserPreferenceLocator
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserPreference", "create")
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
		return &UserPreferenceLocator{UrlResolver(location), loc.api}, nil
	}
}

// PATCH /accounts/:account_id/user_preferences/:id
// Update the value of a UserPreference.
// Multiple values may be updated using a multipart request.
// Values are validated with the corresponding UserPreferenceInfo.
func (loc *UserPreferenceLocator) Update(accountId string, id string) error {
	if accountId == "" {
		return fmt.Errorf("accountId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserPreference", "update")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /accounts/:account_id/user_preferences/:id
// Delete a UserPreference
func (loc *UserPreferenceLocator) Delete(accountId string, id string) error {
	if accountId == "" {
		return fmt.Errorf("accountId is required")
	}
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserPreference", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  UserPreferenceInfo ******/

// The UserPreferenceInfo resource defines the available user preferences supported by the system.
// It is also used to validate values saved in UserPreference.
type UserPreferenceInfo struct {
	Category        string            `json:"category,omitempty"`
	DefaultValue    string            `json:"default_value,omitempty"`
	DisplayName     string            `json:"display_name,omitempty"`
	HelpText        string            `json:"help_text,omitempty"`
	Href            string            `json:"href,omitempty"`
	Id              string            `json:"id,omitempty"`
	Kind            string            `json:"kind,omitempty"`
	Name            string            `json:"name,omitempty"`
	ValueConstraint []string          `json:"value_constraint,omitempty"`
	ValueRange      *ValueRangeStruct `json:"value_range,omitempty"`
	ValueType       string            `json:"value_type,omitempty"`
}

//===== Locator

// UserPreferenceInfo resource locator, exposes resource actions.
type UserPreferenceInfoLocator struct {
	UrlResolver
	api *Api
}

// UserPreferenceInfo resource locator factory
func (api *Api) UserPreferenceInfoLocator(href string) *UserPreferenceInfoLocator {
	return &UserPreferenceInfoLocator{UrlResolver(href), api}
}

//===== Actions

// GET /accounts/:account_id/user_preference_infos
// List the UserPreferenceInfo.
func (loc *UserPreferenceInfoLocator) Index(accountId string, options rsapi.ApiParams) ([]*UserPreferenceInfo, error) {
	var res []*UserPreferenceInfo
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserPreferenceInfo", "index")
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

// GET /accounts/:account_id/user_preference_infos/:id
// Get details for a particular UserPreferenceInfo
func (loc *UserPreferenceInfoLocator) Show(accountId string, id string) (*UserPreferenceInfo, error) {
	var res *UserPreferenceInfo
	if accountId == "" {
		return res, fmt.Errorf("accountId is required")
	}
	if id == "" {
		return res, fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("UserPreferenceInfo", "show")
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

/****** Parameter Data Types ******/

type Parameter struct {
	Default     string                      `json:"default,omitempty"`
	Description string                      `json:"description,omitempty"`
	Name        string                      `json:"name,omitempty"`
	Operations  []string                    `json:"operations,omitempty"`
	Type_       string                      `json:"type,omitempty"`
	Ui          *ParametersUiStruct         `json:"ui,omitempty"`
	Validation  *ParametersValidationStruct `json:"validation,omitempty"`
}

type ParametersUiStruct struct {
	Category string `json:"category,omitempty"`
	Index    int    `json:"index,omitempty"`
	Label    string `json:"label,omitempty"`
}

type ParametersValidationStruct struct {
	AllowedPattern        string   `json:"allowed_pattern,omitempty"`
	AllowedValues         []string `json:"allowed_values,omitempty"`
	ConstraintDescription string   `json:"constraint_description,omitempty"`
	MaxLength             int      `json:"max_length,omitempty"`
	MaxValue              int      `json:"max_value,omitempty"`
	MinLength             int      `json:"min_length,omitempty"`
	MinValue              int      `json:"min_value,omitempty"`
	NoEcho                bool     `json:"no_echo,omitempty"`
}

type Recurrence struct {
	Hour   int    `json:"hour,omitempty"`
	Minute int    `json:"minute,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type Schedule struct {
	CreatedFrom     string      `json:"created_from,omitempty"`
	Description     string      `json:"description,omitempty"`
	Name            string      `json:"name,omitempty"`
	StartRecurrence *Recurrence `json:"start_recurrence,omitempty"`
	StopRecurrence  *Recurrence `json:"stop_recurrence,omitempty"`
}

type TemplateInfo struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

type TimestampsStruct struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type UserPreferenceInfoParam struct {
	Category        string                              `json:"category,omitempty"`
	DefaultValue    string                              `json:"default_value,omitempty"`
	DisplayName     string                              `json:"display_name,omitempty"`
	HelpText        string                              `json:"help_text,omitempty"`
	Href            string                              `json:"href,omitempty"`
	Id              string                              `json:"id,omitempty"`
	Kind            string                              `json:"kind,omitempty"`
	Name            string                              `json:"name,omitempty"`
	ValueConstraint []string                            `json:"value_constraint,omitempty"`
	ValueRange      *UserPreferenceInfoValueRangeStruct `json:"value_range,omitempty"`
	ValueType       string                              `json:"value_type,omitempty"`
}

type UserPreferenceInfoValueRangeStruct struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}

type ValueRangeStruct struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}
