//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ api15gen -metadata=cm15 -output=cm15
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package cm15

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// API Version
const APIVersion = "1.5"

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

type Account struct {
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Permissions []Permission        `json:"permissions,omitempty"`
	Products    []string            `json:"products,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Account) Locator(api *API) *AccountLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.AccountLocator(l["href"])
		}
	}
	return nil
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

// GET /api/accounts/:id
//
// Show information about a single Account.
func (loc *AccountLocator) Show() (*Account, error) {
	var res *Account
	var params rsapi.APIParams
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

/******  AccountGroup ******/

// An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.
type AccountGroup struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *AccountGroup) Locator(api *API) *AccountGroupLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.AccountGroupLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// AccountGroupLocator exposes the AccountGroup resource actions.
type AccountGroupLocator struct {
	Href
	api *API
}

// AccountGroupLocator builds a locator from the given href.
func (api *API) AccountGroupLocator(href string) *AccountGroupLocator {
	return &AccountGroupLocator{Href(href), api}
}

//===== Actions

// GET /api/account_groups
//
// Lists the AccountGroups owned by this Account.
// Optional parameters:
// filter
// view
func (loc *AccountGroupLocator) Index(options rsapi.APIParams) ([]*AccountGroup, error) {
	var res []*AccountGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AccountGroup", "index")
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

// GET /api/account_groups/:id
//
// Show information about a single AccountGroup.
// Optional parameters:
// view
func (loc *AccountGroupLocator) Show(options rsapi.APIParams) (*AccountGroup, error) {
	var res *AccountGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AccountGroup", "show")
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

/******  Alert ******/

// An Alert represents an AlertSpec bound to a running Instance.
type Alert struct {
	Actions       []map[string]string `json:"actions,omitempty"`
	CreatedAt     *RubyTime           `json:"created_at,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	QuenchedUntil *RubyTime           `json:"quenched_until,omitempty"`
	Status        string              `json:"status,omitempty"`
	TriggeredAt   *RubyTime           `json:"triggered_at,omitempty"`
	UpdatedAt     *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Alert) Locator(api *API) *AlertLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.AlertLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// AlertLocator exposes the Alert resource actions.
type AlertLocator struct {
	Href
	api *API
}

// AlertLocator builds a locator from the given href.
func (api *API) AlertLocator(href string) *AlertLocator {
	return &AlertLocator{Href(href), api}
}

//===== Actions

// DELETE /api/clouds/:cloud_id/instances/:instance_id/alerts/:id
// DELETE /api/servers/:server_id/alerts/:id
// DELETE /api/server_arrays/:server_array_id/alerts/:id
// DELETE /api/deployments/:deployment_id/alerts/:id
// DELETE /api/alerts/:id
//
// Destroys the Alert.
func (loc *AlertLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Alert", "destroy")
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

// POST /api/clouds/:cloud_id/instances/:instance_id/alerts/:id/disable
// POST /api/servers/:server_id/alerts/:id/disable
// POST /api/server_arrays/:server_array_id/alerts/:id/disable
// POST /api/deployments/:deployment_id/alerts/:id/disable
// POST /api/alerts/:id/disable
//
// Disables the Alert indefinitely. Idempotent.
func (loc *AlertLocator) Disable() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Alert", "disable")
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

// POST /api/clouds/:cloud_id/instances/:instance_id/alerts/:id/enable
// POST /api/servers/:server_id/alerts/:id/enable
// POST /api/server_arrays/:server_array_id/alerts/:id/enable
// POST /api/deployments/:deployment_id/alerts/:id/enable
// POST /api/alerts/:id/enable
//
// Enables the Alert indefinitely. Idempotent.
func (loc *AlertLocator) Enable() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Alert", "enable")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/alerts
// GET /api/servers/:server_id/alerts
// GET /api/server_arrays/:server_array_id/alerts
// GET /api/deployments/:deployment_id/alerts
// GET /api/alerts
//
// Lists all Alerts.
// Optional parameters:
// filter
// view
func (loc *AlertLocator) Index(options rsapi.APIParams) ([]*Alert, error) {
	var res []*Alert
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Alert", "index")
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

// POST /api/clouds/:cloud_id/instances/:instance_id/alerts/:id/quench
// POST /api/servers/:server_id/alerts/:id/quench
// POST /api/server_arrays/:server_array_id/alerts/:id/quench
// POST /api/deployments/:deployment_id/alerts/:id/quench
// POST /api/alerts/:id/quench
//
// Suppresses the Alert from being triggered for a given time period. Idempotent.
// Required parameters:
// duration: The time period in seconds to suppress Alert from being triggered.
func (loc *AlertLocator) Quench(duration string) (string, error) {
	var res string
	if duration == "" {
		return res, fmt.Errorf("duration is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"duration": duration,
	}
	uri, err := loc.ActionPath("Alert", "quench")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/alerts/:id
// GET /api/servers/:server_id/alerts/:id
// GET /api/server_arrays/:server_array_id/alerts/:id
// GET /api/deployments/:deployment_id/alerts/:id
// GET /api/alerts/:id
//
// Shows the attributes of a specified Alert.
// Optional parameters:
// view
func (loc *AlertLocator) Show(options rsapi.APIParams) (*Alert, error) {
	var res *Alert
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Alert", "show")
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

/******  AlertSpec ******/

// An AlertSpec defines the conditions under which an Alert is triggered and escalated.
// Condition sentence: if &lt;file&gt;.&lt;variable&gt; &lt;condition&gt; '&lt;threshold&gt;' for &lt;duration&gt; min then escalate to '&lt;escalation_name&gt;'.
type AlertSpec struct {
	Actions        []map[string]string `json:"actions,omitempty"`
	Condition      string              `json:"condition,omitempty"`
	CreatedAt      *RubyTime           `json:"created_at,omitempty"`
	Description    string              `json:"description,omitempty"`
	Duration       int                 `json:"duration,omitempty"`
	EscalationName string              `json:"escalation_name,omitempty"`
	File           string              `json:"file,omitempty"`
	Links          []map[string]string `json:"links,omitempty"`
	Name           string              `json:"name,omitempty"`
	Threshold      string              `json:"threshold,omitempty"`
	UpdatedAt      *RubyTime           `json:"updated_at,omitempty"`
	Variable       string              `json:"variable,omitempty"`
	VoteTag        string              `json:"vote_tag,omitempty"`
	VoteType       string              `json:"vote_type,omitempty"`
}

// Locator returns a locator for the given resource
func (r *AlertSpec) Locator(api *API) *AlertSpecLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.AlertSpecLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// AlertSpecLocator exposes the AlertSpec resource actions.
type AlertSpecLocator struct {
	Href
	api *API
}

// AlertSpecLocator builds a locator from the given href.
func (api *API) AlertSpecLocator(href string) *AlertSpecLocator {
	return &AlertSpecLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances/:instance_id/alert_specs
// POST /api/servers/:server_id/alert_specs
// POST /api/server_arrays/:server_array_id/alert_specs
// POST /api/server_templates/:server_template_id/alert_specs
// POST /api/alert_specs
//
// Creates a new AlertSpec with the given parameters.
// Required parameters:
// alert_spec
func (loc *AlertSpecLocator) Create(alertSpec *AlertSpecParam) (*AlertSpecLocator, error) {
	var res *AlertSpecLocator
	if alertSpec == nil {
		return res, fmt.Errorf("alertSpec is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"alert_spec": alertSpec,
	}
	uri, err := loc.ActionPath("AlertSpec", "create")
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
		return &AlertSpecLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/instances/:instance_id/alert_specs/:id
// DELETE /api/servers/:server_id/alert_specs/:id
// DELETE /api/server_arrays/:server_array_id/alert_specs/:id
// DELETE /api/server_templates/:server_template_id/alert_specs/:id
// DELETE /api/alert_specs/:id
//
// Deletes a given AlertSpec.
func (loc *AlertSpecLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AlertSpec", "destroy")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/alert_specs
// GET /api/servers/:server_id/alert_specs
// GET /api/server_arrays/:server_array_id/alert_specs
// GET /api/server_templates/:server_template_id/alert_specs
// GET /api/alert_specs
//
// No description provided for index.
// Optional parameters:
// filter
// view
// with_inherited: Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.
func (loc *AlertSpecLocator) Index(options rsapi.APIParams) ([]*AlertSpec, error) {
	var res []*AlertSpec
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var withInheritedOpt = options["with_inherited"]
	if withInheritedOpt != nil {
		p["with_inherited"] = withInheritedOpt
	}
	uri, err := loc.ActionPath("AlertSpec", "index")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/alert_specs/:id
// GET /api/servers/:server_id/alert_specs/:id
// GET /api/server_arrays/:server_array_id/alert_specs/:id
// GET /api/server_templates/:server_template_id/alert_specs/:id
// GET /api/alert_specs/:id
//
// No description provided for show.
// Optional parameters:
// view
func (loc *AlertSpecLocator) Show(options rsapi.APIParams) (*AlertSpec, error) {
	var res *AlertSpec
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AlertSpec", "show")
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

// PUT /api/clouds/:cloud_id/instances/:instance_id/alert_specs/:id
// PUT /api/servers/:server_id/alert_specs/:id
// PUT /api/server_arrays/:server_array_id/alert_specs/:id
// PUT /api/server_templates/:server_template_id/alert_specs/:id
// PUT /api/alert_specs/:id
//
// Updates an AlertSpec with the given parameters.
// Required parameters:
// alert_spec
func (loc *AlertSpecLocator) Update(alertSpec *AlertSpecParam2) error {
	if alertSpec == nil {
		return fmt.Errorf("alertSpec is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"alert_spec": alertSpec,
	}
	uri, err := loc.ActionPath("AlertSpec", "update")
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

/******  AuditEntry ******/

// An Audit Entry can be used to track various activities of a resource.
type AuditEntry struct {
	Actions    []map[string]string `json:"actions,omitempty"`
	DetailSize int                 `json:"detail_size,omitempty"`
	Links      []map[string]string `json:"links,omitempty"`
	Summary    string              `json:"summary,omitempty"`
	UpdatedAt  *RubyTime           `json:"updated_at,omitempty"`
	UserEmail  string              `json:"user_email,omitempty"`
}

// Locator returns a locator for the given resource
func (r *AuditEntry) Locator(api *API) *AuditEntryLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.AuditEntryLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// AuditEntryLocator exposes the AuditEntry resource actions.
type AuditEntryLocator struct {
	Href
	api *API
}

// AuditEntryLocator builds a locator from the given href.
func (api *API) AuditEntryLocator(href string) *AuditEntryLocator {
	return &AuditEntryLocator{Href(href), api}
}

//===== Actions

// POST /api/audit_entries/:id/append
//
// Updates the summary and appends more details to a given AuditEntry. Each audit entry detail is stored
// as one chunk, the offset determines the location of that chunk within the overall audit entry details section.
// For example, if you create an AuditEntry and append "DEF" at offset 10, and later append
// "ABC" at offset 9, the overall audit entry details will be "ABCDEF". Use the \n character to
// separate details by new lines.
// Optional parameters:
// detail: The details to be appended to the audit entry record.
// notify: The event notification category. Defaults to 'None'.
// offset: The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.
// summary: The updated summary for the audit entry, maximum length is 255 characters.
func (loc *AuditEntryLocator) Append(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var detailOpt = options["detail"]
	if detailOpt != nil {
		p["detail"] = detailOpt
	}
	var notifyOpt = options["notify"]
	if notifyOpt != nil {
		p["notify"] = notifyOpt
	}
	var offsetOpt = options["offset"]
	if offsetOpt != nil {
		p["offset"] = offsetOpt
	}
	var summaryOpt = options["summary"]
	if summaryOpt != nil {
		p["summary"] = summaryOpt
	}
	uri, err := loc.ActionPath("AuditEntry", "append")
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

// POST /api/audit_entries
//
// Creates a new AuditEntry with the given parameters.
// Required parameters:
// audit_entry
// Optional parameters:
// notify: The event notification category. Defaults to 'None'.
// user_email: The email of the user (who created/triggered the audit entry). Only usable with instance role.
func (loc *AuditEntryLocator) Create(auditEntry *AuditEntryParam, options rsapi.APIParams) (*AuditEntryLocator, error) {
	var res *AuditEntryLocator
	if auditEntry == nil {
		return res, fmt.Errorf("auditEntry is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"audit_entry": auditEntry,
	}
	var notifyOpt = options["notify"]
	if notifyOpt != nil {
		p["notify"] = notifyOpt
	}
	var userEmailOpt = options["user_email"]
	if userEmailOpt != nil {
		p["user_email"] = userEmailOpt
	}
	uri, err := loc.ActionPath("AuditEntry", "create")
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
		return &AuditEntryLocator{Href(location), loc.api}, nil
	}
}

// GET /api/audit_entries/:id/detail
//
// shows the details of a given AuditEntry.
// Note that the media type of the response is simply text.
func (loc *AuditEntryLocator) Detail() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AuditEntry", "detail")
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

// GET /api/audit_entries
//
// Lists AuditEntries of the account. Due to the potentially large number of audit entries, a start and end date must
// be provided during an index call to limit the search. The format of the dates must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
// 2011/07/11 00:00:00 +0000.
// A maximum of 1000 records will be returned by each index call.
// Using the available filters, one can select or group which audit entries to retrieve.
// Required parameters:
// end_date: The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).
// limit: Limit the audit entries to this number. The limit should >= 1 and <= 1000
// start_date: The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000
// Optional parameters:
// filter
// view
func (loc *AuditEntryLocator) Index(endDate string, limit string, startDate string, options rsapi.APIParams) ([]*AuditEntry, error) {
	var res []*AuditEntry
	if endDate == "" {
		return res, fmt.Errorf("endDate is required")
	}
	if limit == "" {
		return res, fmt.Errorf("limit is required")
	}
	if startDate == "" {
		return res, fmt.Errorf("startDate is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"end_date":   endDate,
		"limit":      limit,
		"start_date": startDate,
	}
	uri, err := loc.ActionPath("AuditEntry", "index")
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

// GET /api/audit_entries/:id
//
// Lists the attributes of a given audit entry.
// Optional parameters:
// view
func (loc *AuditEntryLocator) Show(options rsapi.APIParams) (*AuditEntry, error) {
	var res *AuditEntry
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AuditEntry", "show")
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

// PUT /api/audit_entries/:id
//
// Updates the summary of a given AuditEntry.
// Required parameters:
// audit_entry
// Optional parameters:
// notify: The event notification category. Defaults to 'None'.
func (loc *AuditEntryLocator) Update(auditEntry *AuditEntryParam2, options rsapi.APIParams) error {
	if auditEntry == nil {
		return fmt.Errorf("auditEntry is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"audit_entry": auditEntry,
	}
	var notifyOpt = options["notify"]
	if notifyOpt != nil {
		p["notify"] = notifyOpt
	}
	uri, err := loc.ActionPath("AuditEntry", "update")
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

/******  Backup ******/

type Backup struct {
	Actions             []map[string]string `json:"actions,omitempty"`
	Committed           bool                `json:"committed,omitempty"`
	Completed           bool                `json:"completed,omitempty"`
	CreatedAt           *RubyTime           `json:"created_at,omitempty"`
	Description         string              `json:"description,omitempty"`
	FromMaster          bool                `json:"from_master,omitempty"`
	Lineage             string              `json:"lineage,omitempty"`
	Links               []map[string]string `json:"links,omitempty"`
	Name                string              `json:"name,omitempty"`
	VolumeSnapshotCount int                 `json:"volume_snapshot_count,omitempty"`
	VolumeSnapshots     []VolumeSnapshot    `json:"volume_snapshots,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Backup) Locator(api *API) *BackupLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.BackupLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// BackupLocator exposes the Backup resource actions.
type BackupLocator struct {
	Href
	api *API
}

// BackupLocator builds a locator from the given href.
func (api *API) BackupLocator(href string) *BackupLocator {
	return &BackupLocator{Href(href), api}
}

//===== Actions

// POST /api/backups/cleanup
//
// Deletes old backups that meet the given criteria. For example, if a user calls cleanup with keep monthlies set to 12,
// then the latest backup for each month, for 12 months, will be kept.
// All backups belong to a particular 'lineage'. Backups are not constrained to a specific cloud or a specific deployment.
// A lineage is account-specific. Hence, backups having the same lineage but belonging to different clouds are still considered
// for cleanup.
// If backups specific to a single cloud should be cleaned up, see the cloud_href parameter.
// Definitions:
// A snapshot is completed if its status is "available"
// A snapshot is committed if it has a tag "rs_backup:committed=true"
// A snapshot belongs to a backup "X" if it has a tag "rs_backup:backup_id=X"
// A snapshot is part of a backup with size "Y" if it has a tag "rs_backup:count=Y"
// A snapshot's position in a backup is "Z" if it has a tag "rs_backup:position=Z"
// Backups are of 3 types:
// Perfect backup: A backup which is completed (all the snapshots are completed) and committed (all the snapshots are committed) and the number of snapshots it found is equal to the number
// in the "rs_backup:count=" tag on each of the Snapshots.
// Imperfect backup: A backup which is not committed or if the number of snapshots it found is not equal to the number in the "rs_backup:count=" tag on each of the snapshots.
// Partial Perfect backup: A snapshot which is neither perfect nor imperfect
// An imperfect backup is picked up for cleanup only if there exists a perfect backup with a newer created_at timestamp.
// No constraints will be applied on such imperfect backups and all of them will be destroyed.
// For all the perfect backups, the constraints of keep_last and dailies etc. will be applied.
// The algorithm for choosing the perfect backups to keep is simple. It is the union of those set of backups if each of those conditions are applied
// independently. i.e backups_to_keep = backups_to_keep(keep_last) U backups_to_keep(dailies) U backups_to_keep(weeklies) U backups_to_keep(monthlies) U backups_to_keep(yearlies)
// Hence, it is important to "commit" a backup to make it eligible for cleanup.
// Required parameters:
// keep_last: The number of backups that should be kept.
// lineage: The lineage of the backups that are to be cleaned-up.
// Optional parameters:
// cloud_href: Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.
// dailies: The number of daily backups(the latest one in each day) that should be kept.
// monthlies: The number of monthly backups(the latest one in each month) that should be kept.
// weeklies: The number of weekly backups(the latest one in each week) that should be kept.
// yearlies: The number of yearly backups(the latest one in each year) that should be kept.
func (loc *BackupLocator) Cleanup(keepLast string, lineage string, options rsapi.APIParams) error {
	if keepLast == "" {
		return fmt.Errorf("keepLast is required")
	}
	if lineage == "" {
		return fmt.Errorf("lineage is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"keep_last": keepLast,
		"lineage":   lineage,
	}
	var cloudHrefOpt = options["cloud_href"]
	if cloudHrefOpt != nil {
		p["cloud_href"] = cloudHrefOpt
	}
	var dailiesOpt = options["dailies"]
	if dailiesOpt != nil {
		p["dailies"] = dailiesOpt
	}
	var monthliesOpt = options["monthlies"]
	if monthliesOpt != nil {
		p["monthlies"] = monthliesOpt
	}
	var weekliesOpt = options["weeklies"]
	if weekliesOpt != nil {
		p["weeklies"] = weekliesOpt
	}
	var yearliesOpt = options["yearlies"]
	if yearliesOpt != nil {
		p["yearlies"] = yearliesOpt
	}
	uri, err := loc.ActionPath("Backup", "cleanup")
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

// POST /api/backups
//
// Takes in an array of volume_attachment_hrefs and takes a snapshot of each.
// The volume_attachment_hrefs must belong to the same instance.
// Required parameters:
// backup
func (loc *BackupLocator) Create(backup *BackupParam) (*BackupLocator, error) {
	var res *BackupLocator
	if backup == nil {
		return res, fmt.Errorf("backup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"backup": backup,
	}
	uri, err := loc.ActionPath("Backup", "create")
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
		return &BackupLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/backups/:id
//
// Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.
func (loc *BackupLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Backup", "destroy")
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

// GET /api/backups
//
// Lists all of the backups with the given lineage tag. Filters can be used to search for a particular backup. If the
// 'latest_before' filter is set, only one backup is returned (the latest backup before the given timestamp).
// To get the latest completed backup, the 'completed' filter should be set to 'true' and the 'latest_before' filter
// should be set to the current timestamp. The format of the timestamp must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
// 2011/07/11 00:00:00 +0000.
// To get the latest completed backup just before, say 25 June 2009, then the 'completed' filter
// should be set to 'true' and the 'latest_before' filter should be set to 2009/06/25 00:00:00 +0000.
// Required parameters:
// lineage: Backups belonging to this lineage.
// Optional parameters:
// filter
func (loc *BackupLocator) Index(lineage string, options rsapi.APIParams) ([]*Backup, error) {
	var res []*Backup
	if lineage == "" {
		return res, fmt.Errorf("lineage is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"lineage": lineage,
	}
	uri, err := loc.ActionPath("Backup", "index")
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

// POST /api/backups/:id/restore
//
// Restores the given Backup.
// This call will:
// create the required number of Volumes from the volume_snapshots_hrefs in the given Backup,
// attach them to the given Instance at the device specified in the Snapshot. If the devices are already being used
// on the Instance, the Task will denote that the restore has failed.
// Required parameters:
// instance_href: The instance href that the backup will be restored to.
// Optional parameters:
// backup
func (loc *BackupLocator) Restore(instanceHref string, options rsapi.APIParams) error {
	if instanceHref == "" {
		return fmt.Errorf("instanceHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"instance_href": instanceHref,
	}
	var backupOpt = options["backup"]
	if backupOpt != nil {
		p["backup"] = backupOpt
	}
	uri, err := loc.ActionPath("Backup", "restore")
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

// GET /api/backups/:id
//
// Lists the attributes of a given backup
func (loc *BackupLocator) Show() (*Backup, error) {
	var res *Backup
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Backup", "show")
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

// PUT /api/backups/:id
//
// Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.
// Required parameters:
// backup
func (loc *BackupLocator) Update(backup *BackupParam3) error {
	if backup == nil {
		return fmt.Errorf("backup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"backup": backup,
	}
	uri, err := loc.ActionPath("Backup", "update")
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

/******  ChildAccount ******/

type ChildAccount struct {
}

//===== Locator

// ChildAccountLocator exposes the ChildAccount resource actions.
type ChildAccountLocator struct {
	Href
	api *API
}

// ChildAccountLocator builds a locator from the given href.
func (api *API) ChildAccountLocator(href string) *ChildAccountLocator {
	return &ChildAccountLocator{Href(href), api}
}

//===== Actions

// POST /api/child_accounts
//
// Create an enterprise ChildAccount for this Account. The User will by default get an 'admin' role
// on the ChildAccount to enable him/her to add, delete Users and Permissions.
// For more information on the valid values for 'cluster_href', refer to the following:
// http://support.rightscale.com/12-Guides/RightScale_API_1.5/Examples/ChildAccounts/Create
// Required parameters:
// child_account
func (loc *ChildAccountLocator) Create(childAccount *ChildAccountParam) (*ChildAccountLocator, error) {
	var res *ChildAccountLocator
	if childAccount == nil {
		return res, fmt.Errorf("childAccount is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"child_account": childAccount,
	}
	uri, err := loc.ActionPath("ChildAccount", "create")
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
		return &ChildAccountLocator{Href(location), loc.api}, nil
	}
}

// GET /api/child_accounts
//
// Lists the enterprise ChildAccounts available for this Account.
// Optional parameters:
// filter
func (loc *ChildAccountLocator) Index(options rsapi.APIParams) ([]*Account, error) {
	var res []*Account
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ChildAccount", "index")
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

// PUT /api/accounts/:id
// PUT /api/child_accounts/:id
//
// Update an enterprise ChildAccount for this Account.
// Required parameters:
// child_account
func (loc *ChildAccountLocator) Update(childAccount *ChildAccountParam2) error {
	if childAccount == nil {
		return fmt.Errorf("childAccount is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"child_account": childAccount,
	}
	uri, err := loc.ActionPath("ChildAccount", "update")
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

/******  Cloud ******/

// Represents a Cloud (within the context of the account in the session).
type Cloud struct {
	Capabilities []map[string]interface{} `json:"capabilities,omitempty"`
	CloudType    string                   `json:"cloud_type,omitempty"`
	Description  string                   `json:"description,omitempty"`
	DisplayName  string                   `json:"display_name,omitempty"`
	Links        []map[string]string      `json:"links,omitempty"`
	Name         string                   `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Cloud) Locator(api *API) *CloudLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.CloudLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// CloudLocator exposes the Cloud resource actions.
type CloudLocator struct {
	Href
	api *API
}

// CloudLocator builds a locator from the given href.
func (api *API) CloudLocator(href string) *CloudLocator {
	return &CloudLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds
//
// Lists the clouds available to this account.
// Optional parameters:
// filter
// view
func (loc *CloudLocator) Index(options rsapi.APIParams) ([]*Cloud, error) {
	var res []*Cloud
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Cloud", "index")
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

// GET /api/clouds/:id
//
// Show information about a single cloud.
// Optional parameters:
// view
func (loc *CloudLocator) Show(options rsapi.APIParams) (*Cloud, error) {
	var res *Cloud
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Cloud", "show")
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

/******  CloudAccount ******/

// Represents a Cloud Account (An association between the account and a cloud).
type CloudAccount struct {
	CreatedAt *RubyTime           `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *CloudAccount) Locator(api *API) *CloudAccountLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.CloudAccountLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// CloudAccountLocator exposes the CloudAccount resource actions.
type CloudAccountLocator struct {
	Href
	api *API
}

// CloudAccountLocator builds a locator from the given href.
func (api *API) CloudAccountLocator(href string) *CloudAccountLocator {
	return &CloudAccountLocator{Href(href), api}
}

//===== Actions

// POST /api/cloud_accounts
//
// Create a CloudAccount by passing in the respective credentials for each cloud.
// For more information on the specific parameters for each cloud, refer to the following:
// http://docs.rightscale.com/api/api_1.5_examples/cloudaccounts.html
// Required parameters:
// cloud_account
func (loc *CloudAccountLocator) Create(cloudAccount *CloudAccountParam) (*CloudAccountLocator, error) {
	var res *CloudAccountLocator
	if cloudAccount == nil {
		return res, fmt.Errorf("cloudAccount is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"cloud_account": cloudAccount,
	}
	uri, err := loc.ActionPath("CloudAccount", "create")
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
		return &CloudAccountLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/cloud_accounts/:id
//
// Delete a CloudAccount.
func (loc *CloudAccountLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CloudAccount", "destroy")
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

// GET /api/cloud_accounts
//
// Lists the CloudAccounts (non-aws) available to this Account.
func (loc *CloudAccountLocator) Index() ([]*CloudAccount, error) {
	var res []*CloudAccount
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CloudAccount", "index")
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

// GET /api/cloud_accounts/:id
//

func (loc *CloudAccountLocator) Show() (*CloudAccount, error) {
	var res *CloudAccount
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CloudAccount", "show")
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

/******  Cookbook ******/

// Represents a given instance of a single cookbook.
type Cookbook struct {
	Actions           []map[string]string `json:"actions,omitempty"`
	CreatedAt         *RubyTime           `json:"created_at,omitempty"`
	DownloadUrl       string              `json:"download_url,omitempty"`
	Id                int                 `json:"id,omitempty"`
	Links             []map[string]string `json:"links,omitempty"`
	Metadata          string              `json:"metadata,omitempty"`
	Name              string              `json:"name,omitempty"`
	Namespace         string              `json:"namespace,omitempty"`
	SourceInfoSummary string              `json:"source_info_summary,omitempty"`
	State             string              `json:"state,omitempty"`
	UpdatedAt         *RubyTime           `json:"updated_at,omitempty"`
	Version           string              `json:"version,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Cookbook) Locator(api *API) *CookbookLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.CookbookLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// CookbookLocator exposes the Cookbook resource actions.
type CookbookLocator struct {
	Href
	api *API
}

// CookbookLocator builds a locator from the given href.
func (api *API) CookbookLocator(href string) *CookbookLocator {
	return &CookbookLocator{Href(href), api}
}

//===== Actions

// DELETE /api/cookbooks/:id
//
// Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.
func (loc *CookbookLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Cookbook", "destroy")
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

// POST /api/cookbooks/:id/follow
//
// Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.
// Required parameters:
// value: Indicates if this action should follow (true) or unfollow (false) a Cookbook.
func (loc *CookbookLocator) Follow(value string) error {
	if value == "" {
		return fmt.Errorf("value is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"value": value,
	}
	uri, err := loc.ActionPath("Cookbook", "follow")
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

// POST /api/cookbooks/:id/freeze
//
// Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.
// Required parameters:
// value: Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.
func (loc *CookbookLocator) Freeze(value string) error {
	if value == "" {
		return fmt.Errorf("value is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"value": value,
	}
	uri, err := loc.ActionPath("Cookbook", "freeze")
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

// GET /api/cookbooks
//
// Lists the Cookbooks available to this account.
// The extended_designer view is only available to accounts with the designer permission.
// Optional parameters:
// filter
// view
func (loc *CookbookLocator) Index(options rsapi.APIParams) ([]*Cookbook, error) {
	var res []*Cookbook
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Cookbook", "index")
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

// POST /api/cookbooks/:id/obsolete
//
// Marks a Cookbook as obsolete (or un-obsolete).
// Required parameters:
// value: Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.
func (loc *CookbookLocator) Obsolete(value string) error {
	if value == "" {
		return fmt.Errorf("value is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"value": value,
	}
	uri, err := loc.ActionPath("Cookbook", "obsolete")
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

// GET /api/cookbooks/:id
//
// Show information about a single Cookbook.
// The extended_designer view is only available to accounts with the designer permission.
// Optional parameters:
// view
func (loc *CookbookLocator) Show(options rsapi.APIParams) (*Cookbook, error) {
	var res *Cookbook
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Cookbook", "show")
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

/******  CookbookAttachment ******/

// Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.
type CookbookAttachment struct {
	Actions    []map[string]string `json:"actions,omitempty"`
	Dependency bool                `json:"dependency,omitempty"`
	Id         string              `json:"id,omitempty"`
	Links      []map[string]string `json:"links,omitempty"`
}

// Locator returns a locator for the given resource
func (r *CookbookAttachment) Locator(api *API) *CookbookAttachmentLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.CookbookAttachmentLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// CookbookAttachmentLocator exposes the CookbookAttachment resource actions.
type CookbookAttachmentLocator struct {
	Href
	api *API
}

// CookbookAttachmentLocator builds a locator from the given href.
func (api *API) CookbookAttachmentLocator(href string) *CookbookAttachmentLocator {
	return &CookbookAttachmentLocator{Href(href), api}
}

//===== Actions

// POST /api/cookbooks/:cookbook_id/cookbook_attachments
// POST /api/server_templates/:server_template_id/cookbook_attachments
// POST /api/cookbook_attachments
//
// Attach a cookbook to a given resource.
// Optional parameters:
// cookbook_attachment
func (loc *CookbookAttachmentLocator) Create(options rsapi.APIParams) (*CookbookAttachmentLocator, error) {
	var res *CookbookAttachmentLocator
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var cookbookAttachmentOpt = options["cookbook_attachment"]
	if cookbookAttachmentOpt != nil {
		p["cookbook_attachment"] = cookbookAttachmentOpt
	}
	uri, err := loc.ActionPath("CookbookAttachment", "create")
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
		return &CookbookAttachmentLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/cookbooks/:cookbook_id/cookbook_attachments/:id
// DELETE /api/server_templates/:server_template_id/cookbook_attachments/:id
// DELETE /api/cookbook_attachments/:id
//
// Detach a cookbook from a given resource.
func (loc *CookbookAttachmentLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CookbookAttachment", "destroy")
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

// GET /api/cookbooks/:cookbook_id/cookbook_attachments
// GET /api/server_templates/:server_template_id/cookbook_attachments
// GET /api/cookbook_attachments
//
// Lists Cookbook Attachments.
// Optional parameters:
// view
func (loc *CookbookAttachmentLocator) Index(options rsapi.APIParams) ([]*CookbookAttachment, error) {
	var res []*CookbookAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CookbookAttachment", "index")
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

// POST /api/server_templates/:server_template_id/cookbook_attachments/multi_attach
// POST /api/cookbook_attachments/multi_attach
//
// Attach multiple cookbooks to a given resource.
// Required parameters:
// cookbook_attachments
func (loc *CookbookAttachmentLocator) MultiAttach(cookbookAttachments *CookbookAttachments) error {
	if cookbookAttachments == nil {
		return fmt.Errorf("cookbookAttachments is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"cookbook_attachments": cookbookAttachments,
	}
	uri, err := loc.ActionPath("CookbookAttachment", "multi_attach")
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

// POST /api/server_templates/:server_template_id/cookbook_attachments/multi_detach
// POST /api/cookbook_attachments/multi_detach
//
// Detach multiple cookbooks from a given resource.
// Required parameters:
// cookbook_attachments
func (loc *CookbookAttachmentLocator) MultiDetach(cookbookAttachments *CookbookAttachments2) error {
	if cookbookAttachments == nil {
		return fmt.Errorf("cookbookAttachments is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"cookbook_attachments": cookbookAttachments,
	}
	uri, err := loc.ActionPath("CookbookAttachment", "multi_detach")
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

// GET /api/cookbooks/:cookbook_id/cookbook_attachments/:id
// GET /api/server_templates/:server_template_id/cookbook_attachments/:id
// GET /api/cookbook_attachments/:id
//
// Displays information about a single cookbook attachment to a ServerTemplate.
// Optional parameters:
// view
func (loc *CookbookAttachmentLocator) Show(options rsapi.APIParams) (*CookbookAttachment, error) {
	var res *CookbookAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("CookbookAttachment", "show")
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

/******  Credential ******/

// A Credential provides an abstraction for sensitive textual information,
// such as passphrases or cloud credentials, whose value should be encrypted
// when stored in the database and not generally shown in the UI or through the
// API. Credentials may then be used as inputs of type "Cred" in RightScripts
// or Chef recipes. NOTE: Credential values may be updated through the API, but
// values cannot be retrieved via the API.
type Credential struct {
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
	Value       string              `json:"value,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Credential) Locator(api *API) *CredentialLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.CredentialLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// CredentialLocator exposes the Credential resource actions.
type CredentialLocator struct {
	Href
	api *API
}

// CredentialLocator builds a locator from the given href.
func (api *API) CredentialLocator(href string) *CredentialLocator {
	return &CredentialLocator{Href(href), api}
}

//===== Actions

// POST /api/credentials
//
// Creates a new Credential with the given parameters.
// Required parameters:
// credential
func (loc *CredentialLocator) Create(credential *CredentialParam) (*CredentialLocator, error) {
	var res *CredentialLocator
	if credential == nil {
		return res, fmt.Errorf("credential is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"credential": credential,
	}
	uri, err := loc.ActionPath("Credential", "create")
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
		return &CredentialLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/credentials/:id
//
// Deletes a Credential.
func (loc *CredentialLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Credential", "destroy")
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

// GET /api/credentials
//
// Lists the Credentials available to this account.
// Optional parameters:
// filter
// view
func (loc *CredentialLocator) Index(options rsapi.APIParams) ([]*Credential, error) {
	var res []*Credential
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Credential", "index")
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

// GET /api/credentials/:id
//
// Show information about a single Credential. Credential values may be retrieved using the "sensitive" view by users with "admin" role only.
// Optional parameters:
// view
func (loc *CredentialLocator) Show(options rsapi.APIParams) (*Credential, error) {
	var res *Credential
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Credential", "show")
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

// PUT /api/credentials/:id
//
// Updates attributes of a Credential.
// Required parameters:
// credential
func (loc *CredentialLocator) Update(credential *CredentialParam) error {
	if credential == nil {
		return fmt.Errorf("credential is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"credential": credential,
	}
	uri, err := loc.ActionPath("Credential", "update")
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

/******  Datacenter ******/

// Datacenters represent isolated facilities within a cloud. The level and type of isolation is cloud dependent.
// While Datacenters in large public clouds might correspond to different physical buildings, with different power,
// internet links...etc., Datacenters within the context of a private cloud might simply correspond to having different network providers.
// Spreading servers across distinct Datacenters helps minimize outages.
type Datacenter struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Datacenter) Locator(api *API) *DatacenterLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.DatacenterLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// DatacenterLocator exposes the Datacenter resource actions.
type DatacenterLocator struct {
	Href
	api *API
}

// DatacenterLocator builds a locator from the given href.
func (api *API) DatacenterLocator(href string) *DatacenterLocator {
	return &DatacenterLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/datacenters
//
// Lists all Datacenters for a particular cloud.
// Optional parameters:
// filter
// view
func (loc *DatacenterLocator) Index(options rsapi.APIParams) ([]*Datacenter, error) {
	var res []*Datacenter
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Datacenter", "index")
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

// GET /api/clouds/:cloud_id/datacenters/:id
//
// Displays information about a single Datacenter.
// Optional parameters:
// view
func (loc *DatacenterLocator) Show(options rsapi.APIParams) (*Datacenter, error) {
	var res *Datacenter
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Datacenter", "show")
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

/******  Deployment ******/

// Deployments represent logical groupings of related assets such as servers, server arrays, default configuration settings...etc.
type Deployment struct {
	Actions        []map[string]string `json:"actions,omitempty"`
	Description    string              `json:"description,omitempty"`
	Inputs         []map[string]string `json:"inputs,omitempty"`
	Links          []map[string]string `json:"links,omitempty"`
	Locked         bool                `json:"locked,omitempty"`
	Name           string              `json:"name,omitempty"`
	ServerTagScope string              `json:"server_tag_scope,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Deployment) Locator(api *API) *DeploymentLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.DeploymentLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// DeploymentLocator exposes the Deployment resource actions.
type DeploymentLocator struct {
	Href
	api *API
}

// DeploymentLocator builds a locator from the given href.
func (api *API) DeploymentLocator(href string) *DeploymentLocator {
	return &DeploymentLocator{Href(href), api}
}

//===== Actions

// POST /api/deployments/:id/clone
//
// Clones a given deployment.
// Optional parameters:
// deployment
func (loc *DeploymentLocator) Clone(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var deploymentOpt = options["deployment"]
	if deploymentOpt != nil {
		p["deployment"] = deploymentOpt
	}
	uri, err := loc.ActionPath("Deployment", "clone")
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

// POST /api/deployments
//
// Creates a new deployment with the given parameters.
// Required parameters:
// deployment
func (loc *DeploymentLocator) Create(deployment *DeploymentParam2) (*DeploymentLocator, error) {
	var res *DeploymentLocator
	if deployment == nil {
		return res, fmt.Errorf("deployment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"deployment": deployment,
	}
	uri, err := loc.ActionPath("Deployment", "create")
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
		return &DeploymentLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/deployments/:id
//
// Deletes a given deployment.
func (loc *DeploymentLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Deployment", "destroy")
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

// GET /api/deployments
//
// Lists deployments of the account.
// Using the available filters, one can select or group which deployments to retrieve.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// Optional parameters:
// filter
// view
func (loc *DeploymentLocator) Index(options rsapi.APIParams) ([]*Deployment, error) {
	var res []*Deployment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Deployment", "index")
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

// POST /api/deployments/:id/lock
//
// Locks a given deployment. Idempotent.
// Locking prevents servers from being deleted or moved from the deployment.
// Other actions such as adding servers or renaming the deployment are still allowed.
func (loc *DeploymentLocator) Lock() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Deployment", "lock")
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

// GET /api/deployments/:id/servers
//
// Lists the servers belonging to this deployment. This call is equivalent to servers#index call, where the servers returned will
// automatically be filtered by this deployment. See servers#index for details on other options and parameters.
// Optional parameters:
// filter
// view
func (loc *DeploymentLocator) Servers(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Deployment", "servers")
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

// GET /api/deployments/:id
//
// Lists the attributes of a given deployment.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// Optional parameters:
// view
func (loc *DeploymentLocator) Show(options rsapi.APIParams) (*Deployment, error) {
	var res *Deployment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Deployment", "show")
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

// POST /api/deployments/:id/unlock
//
// Unlocks a given deployment. Idempotent.
func (loc *DeploymentLocator) Unlock() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Deployment", "unlock")
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

// PUT /api/deployments/:id
//
// Updates attributes of a given deployment.
// Required parameters:
// deployment
func (loc *DeploymentLocator) Update(deployment *DeploymentParam3) error {
	if deployment == nil {
		return fmt.Errorf("deployment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"deployment": deployment,
	}
	uri, err := loc.ActionPath("Deployment", "update")
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

/******  HealthCheck ******/

type HealthCheck struct {
}

//===== Locator

// HealthCheckLocator exposes the HealthCheck resource actions.
type HealthCheckLocator struct {
	Href
	api *API
}

// HealthCheckLocator builds a locator from the given href.
func (api *API) HealthCheckLocator(href string) *HealthCheckLocator {
	return &HealthCheckLocator{Href(href), api}
}

//===== Actions

// GET /api/health-check/
//
// Check health of RightApi controllers
func (loc *HealthCheckLocator) Index() ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("HealthCheck", "index")
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

/******  IdentityProvider ******/

// An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account,
// and is trusted by the RightScale dashboard to authenticate your enterprise's end users.
// To register an Identity Provider, contact your account manager.
type IdentityProvider struct {
	Actions       []map[string]string `json:"actions,omitempty"`
	CreatedAt     *RubyTime           `json:"created_at,omitempty"`
	DiscoveryHint string              `json:"discovery_hint,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	Name          string              `json:"name,omitempty"`
	UpdatedAt     *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *IdentityProvider) Locator(api *API) *IdentityProviderLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.IdentityProviderLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// IdentityProviderLocator exposes the IdentityProvider resource actions.
type IdentityProviderLocator struct {
	Href
	api *API
}

// IdentityProviderLocator builds a locator from the given href.
func (api *API) IdentityProviderLocator(href string) *IdentityProviderLocator {
	return &IdentityProviderLocator{Href(href), api}
}

//===== Actions

// GET /api/identity_providers
//
// Lists the identity providers associated with this enterprise account.
// Optional parameters:
// filter
// view
func (loc *IdentityProviderLocator) Index(options rsapi.APIParams) ([]*IdentityProvider, error) {
	var res []*IdentityProvider
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IdentityProvider", "index")
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

// GET /api/identity_providers/:id
//
// Show the specified identity provider, if associated with this enterprise account.
// Optional parameters:
// view
func (loc *IdentityProviderLocator) Show(options rsapi.APIParams) (*IdentityProvider, error) {
	var res *IdentityProvider
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IdentityProvider", "show")
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

/******  Image ******/

// Images represent base VM image existing in a cloud. An image will define the initial Operating System and root disk contents
// for a new Instance to have, and therefore it represents the basic starting point for creating a new one.
type Image struct {
	Actions            []map[string]string `json:"actions,omitempty"`
	CpuArchitecture    string              `json:"cpu_architecture,omitempty"`
	Description        string              `json:"description,omitempty"`
	ImageType          string              `json:"image_type,omitempty"`
	Links              []map[string]string `json:"links,omitempty"`
	Name               string              `json:"name,omitempty"`
	OsPlatform         string              `json:"os_platform,omitempty"`
	ResourceUid        string              `json:"resource_uid,omitempty"`
	RootDeviceStorage  string              `json:"root_device_storage,omitempty"`
	VirtualizationType string              `json:"virtualization_type,omitempty"`
	Visibility         string              `json:"visibility,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Image) Locator(api *API) *ImageLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.ImageLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// ImageLocator exposes the Image resource actions.
type ImageLocator struct {
	Href
	api *API
}

// ImageLocator builds a locator from the given href.
func (api *API) ImageLocator(href string) *ImageLocator {
	return &ImageLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/images
//
// Lists all Images for the given Cloud.
// Optional parameters:
// filter
// view
func (loc *ImageLocator) Index(options rsapi.APIParams) ([]*Image, error) {
	var res []*Image
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Image", "index")
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

// GET /api/clouds/:cloud_id/images/:id
//
// Shows information about a single Image.
// Optional parameters:
// view
func (loc *ImageLocator) Show(options rsapi.APIParams) (*Image, error) {
	var res *Image
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Image", "show")
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

/******  Input ******/

// Inputs help extract dynamic information, usually specified at runtime, from repeatable configuration operations that can be codified.
// Inputs are variables defined in and used by RightScripts/Recipes. The two main attributes of an input are 'name' and 'value'. The 'name'
// identifies the input and the 'value', although a string encodes what type it is. It could be a text encoded as 'text:myvalue' or a credential
// encoded as 'cred:MY_CRED' or a key etc. Please see support.rightscale.com for more info on input hierarchies and their different types.
type Input struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

//===== Locator

// InputLocator exposes the Input resource actions.
type InputLocator struct {
	Href
	api *API
}

// InputLocator builds a locator from the given href.
func (api *API) InputLocator(href string) *InputLocator {
	return &InputLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instances/:instance_id/inputs
// GET /api/deployments/:deployment_id/inputs
// GET /api/server_templates/:server_template_id/inputs
//
// Retrieves the full list of existing inputs of the specified resource.
// Optional parameters:
// view
func (loc *InputLocator) Index(options rsapi.APIParams) ([]*Input, error) {
	var res []*Input
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Input", "index")
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

// PUT /api/clouds/:cloud_id/instances/:instance_id/inputs/multi_update
// PUT /api/deployments/:deployment_id/inputs/multi_update
// PUT /api/server_templates/:server_template_id/inputs/multi_update
//
// Performs a bulk update of inputs on the specified resource.
// If an input exists with the same name, its value will be updated. If an input does not
// exist with a specified name, it will be ignored.
// Input values are represented as strings.
// There are two notations for inputs:
// 1.0 notation - The deprecated notation used in API 1.0 and in 1.5
// 2.0 notation - The new notation that is partially supported in API 1.5, and will
// be the only notation supported in 2.0
// Although the two notations are similar, they have a few important differences, in particular:
// With 2.0 notation, values MUST begin with a prefix identifying their type, followed
// by a colon (example: "text:foo"). With 1.0 notation, unprefixed values are generally
// taken to be text-type.
// With 2.0 notation, a sentinel value "inherit" is used to express that an input
// should use an inherited value. With 1.0 notation the empty string was used to express
// the same thing. (Due to requirement 1, empty string is no longer a valid input.)
// With 2.0 notation, each element of an array is an entire input value; arrays can
// contain cred, env, or even other arrays. With 1.0 notation, array elements are
// implicitly text values and there is no way to specify anything else.Note that the UI
// does not support complex-valued arrays; please use this feature with caution!
// The following types of inputs are supported:
// Type
// Format
// 1.0 Example(s)
// 2.0 Example(s)
// Text string
// &lt;value&gt; (1.0 only)text:&lt;value&gt;
// footext:footext:multi word value
// text:footext:multi word value
// Blank string(input is present but its value is empty-string)
// text:blank (2.0 only)
// text:
// blank
// Ignore (input is not present, input will inherit)
// ignore$ignore (1.0 only)ignore:$ignore (1.0 only)
// ignore$ignoreignore:$ignore
// ignore
// Dynamically-substituted environment value
// env:&lt;value&gt;env:&lt;component&gt;:&lt;value&gt;
// env:MY_ENV_VARenv:my_server:MY_ENV_VAR
// env:MY_ENV_VARenv:my_server:MY_ENV_VAR
// Credential value
// cred:&lt;value&gt;
// cred:abcd1234wxyz
// cred:abcd1234wxyz
// Private SSH key
// key:&lt;value&gt;key:&lt;value&gt;:&lt;cloud_id&gt;
// key:1234abcd5678key:1234abcd5678:1
// key:1234abcd5678key:1234abcd5678:1
// Array of values
// array:&lt;value&gt;,... (1.0 only)array:["&lt;type&gt;:&lt;value&gt;",...] (2.0 only)
// array:x,y(NOTE: 1.0 only supports text inputs for arrays)
// array:["text:v1","text:v2"]array:["text:x","env:server_x:MY_VAR"]
// Note that in the case of array inputs, the portion after the colon must be
// valid JSON. In particular, when enclosing the input within double-quotes
// (e.g. for use in cURL or Ruby), the double-quotes must be escaped.
// Single-quotes may not be used within the array input, since they are not
// valid for JSON strings.
// The legacy format for providing inputs is as an array of name-value pairs
// (ex: -d inputs[][name]="MY_INPUT" -d inputs[][value]="text:foobar"), however
// the new format is supported for inputs provided as a hash
// (ex: -d inputs[MY_INPUT]="text:foobar").
// If the old format is used, the input is parsed using 1.0 semantics.
// If the new format is used, the input is parsed using the new 2.0 semantics.
// Required parameters:
// inputs
func (loc *InputLocator) MultiUpdate(inputs map[string]interface{}) error {
	if len(inputs) == 0 {
		return fmt.Errorf("inputs is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"inputs": inputs,
	}
	uri, err := loc.ActionPath("Input", "multi_update")
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

/******  Instance ******/

// Instances represent an entity that is runnable in the cloud.
// An instance of type "next" is a container of information that expresses how to configure a future instance when we decide
// to launch or start it.
// A "next" instance generally only exists in the RightScale realm, and usually doesn't have any corresponding representation
// existing in the cloud. However, if an instance is not of type "next", it will generally represent an existing running
// (or provisioned) virtual machine existing in the cloud.
type Instance struct {
	Actions                  []map[string]string    `json:"actions,omitempty"`
	AdminPassword            string                 `json:"admin_password,omitempty"`
	AssociatePublicIpAddress bool                   `json:"associate_public_ip_address,omitempty"`
	CloudSpecificAttributes  map[string]interface{} `json:"cloud_specific_attributes,omitempty"`
	CreatedAt                *RubyTime              `json:"created_at,omitempty"`
	Description              string                 `json:"description,omitempty"`
	InheritedSources         map[string]string      `json:"inherited_sources,omitempty"`
	Inputs                   []map[string]string    `json:"inputs,omitempty"`
	IpForwardingEnabled      bool                   `json:"ip_forwarding_enabled,omitempty"`
	Links                    []map[string]string    `json:"links,omitempty"`
	Locked                   bool                   `json:"locked,omitempty"`
	MonitoringCollectorHttp  string                 `json:"monitoring_collector_http,omitempty"`
	MonitoringCollectorUdp   string                 `json:"monitoring_collector_udp,omitempty"`
	MonitoringId             string                 `json:"monitoring_id,omitempty"`
	MonitoringServer         string                 `json:"monitoring_server,omitempty"`
	Name                     string                 `json:"name,omitempty"`
	OsPlatform               string                 `json:"os_platform,omitempty"`
	PricingType              string                 `json:"pricing_type,omitempty"`
	PrivateDnsNames          []string               `json:"private_dns_names,omitempty"`
	PrivateIpAddresses       []string               `json:"private_ip_addresses,omitempty"`
	PublicDnsNames           []string               `json:"public_dns_names,omitempty"`
	PublicIpAddresses        []string               `json:"public_ip_addresses,omitempty"`
	ResourceUid              string                 `json:"resource_uid,omitempty"`
	RsProvisioned            bool                   `json:"rs_provisioned,omitempty"`
	SecurityGroups           []SecurityGroup        `json:"security_groups,omitempty"`
	State                    string                 `json:"state,omitempty"`
	Subnets                  []Subnet               `json:"subnets,omitempty"`
	TerminatedAt             *RubyTime              `json:"terminated_at,omitempty"`
	UpdatedAt                *RubyTime              `json:"updated_at,omitempty"`
	UserData                 string                 `json:"user_data,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Instance) Locator(api *API) *InstanceLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.InstanceLocator(l["href"])
		}
	}
	return nil
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

// POST /api/clouds/:cloud_id/instances
//
// Creates and launches a raw instance using the provided parameters.
// Required parameters:
// instance
// Optional parameters:
// api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
func (loc *InstanceLocator) Create(instance *InstanceParam, options rsapi.APIParams) (*InstanceLocator, error) {
	var res *InstanceLocator
	if instance == nil {
		return res, fmt.Errorf("instance is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"instance": instance,
	}
	var apiBehaviorOpt = options["api_behavior"]
	if apiBehaviorOpt != nil {
		p["api_behavior"] = apiBehaviorOpt
	}
	uri, err := loc.ActionPath("Instance", "create")
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
		return &InstanceLocator{Href(location), loc.api}, nil
	}
}

// GET /api/clouds/:cloud_id/instances
// GET /api/clouds/:cloud_id/instances
// GET /api/server_arrays/:server_array_id/current_instances
//
// Lists instances of a given cloud, server array.
// Using the available filters, it is possible to craft powerful queries about which instances to retrieve.
// For example, one can easily list:
// * instances that have names that contain "app"
// * all instances of a given deployment
// * instances belonging to a given server array (i.e., have the same parent_url)
// To see the instances of a server array including the next_instance, use the URL "/api/clouds/:cloud_id/instances" with the filter "parent_href==/api/server_arrays/XX". To list only the running
// instances of a server array, use the URL "/api/server_arrays/:server_array_id/current_instances"
// The 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// Optional parameters:
// filter
// view
func (loc *InstanceLocator) Index(options rsapi.APIParams) ([]*Instance, error) {
	var res []*Instance
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
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

// POST /api/clouds/:cloud_id/instances/:id/launch
// POST /api/servers/:server_id/launch
// POST /api/server_arrays/:server_array_id/launch
//
// Launches an instance using the parameters that this instance has been configured with.
// Note that this action can only be performed in "next" instances, and not on instances that are already running.
// Optional parameters:
// api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
// count: For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error
// inputs
func (loc *InstanceLocator) Launch(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var apiBehaviorOpt = options["api_behavior"]
	if apiBehaviorOpt != nil {
		p["api_behavior"] = apiBehaviorOpt
	}
	var countOpt = options["count"]
	if countOpt != nil {
		p["count"] = countOpt
	}
	var inputsOpt = options["inputs"]
	if inputsOpt != nil {
		p["inputs"] = inputsOpt
	}
	uri, err := loc.ActionPath("Instance", "launch")
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

// POST /api/clouds/:cloud_id/instances/:id/lock
//

func (loc *InstanceLocator) Lock() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "lock")
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

// POST /api/clouds/:cloud_id/instances/multi_run_executable
// POST /api/server_arrays/:server_array_id/multi_run_executable
//
// Runs a script or a recipe in the running instances.
// This is an asynchronous function, which returns immediately after queuing the executable for execution.
// Status of the execution can be tracked at the URL returned in the "Location" header.
// Optional parameters:
// filter
// ignore_lock: Specifies the ability to ignore the lock(s) on the Instance(s).
// inputs
// recipe_name: The name of the recipe to be run.
// right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.
func (loc *InstanceLocator) MultiRunExecutable(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var ignoreLockOpt = options["ignore_lock"]
	if ignoreLockOpt != nil {
		p["ignore_lock"] = ignoreLockOpt
	}
	var inputsOpt = options["inputs"]
	if inputsOpt != nil {
		p["inputs"] = inputsOpt
	}
	var recipeNameOpt = options["recipe_name"]
	if recipeNameOpt != nil {
		p["recipe_name"] = recipeNameOpt
	}
	var rightScriptHrefOpt = options["right_script_href"]
	if rightScriptHrefOpt != nil {
		p["right_script_href"] = rightScriptHrefOpt
	}
	uri, err := loc.ActionPath("Instance", "multi_run_executable")
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

// POST /api/clouds/:cloud_id/instances/multi_terminate
// POST /api/server_arrays/:server_array_id/multi_terminate
//
// Terminates running instances.
// Either a filter or the parameter 'terminate_all' must be provided.
// Optional parameters:
// filter
// terminate_all: Specifies the ability to terminate all instances.
func (loc *InstanceLocator) MultiTerminate(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var terminateAllOpt = options["terminate_all"]
	if terminateAllOpt != nil {
		p["terminate_all"] = terminateAllOpt
	}
	uri, err := loc.ActionPath("Instance", "multi_terminate")
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

// POST /api/clouds/:cloud_id/instances/:id/reboot
// POST /api/servers/:server_id/reboot
//
// Reboot a running instance.
// Note that this action can only succeed if the instance is running. One cannot reboot instances of type "next".
func (loc *InstanceLocator) Reboot() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "reboot")
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

// POST /api/clouds/:cloud_id/instances/:id/run_executable
//
// Runs a script or a recipe in the running instance.
// This is an asynchronous function, which returns immediately after queuing the executable for execution.
// Status of the execution can be tracked at the URL returned in the "Location" header.
// Note that this can only be performed on running instances.
// Optional parameters:
// ignore_lock: Specifies the ability to ignore the lock on the Instance.
// inputs
// recipe_name: The name of the recipe to run.
// right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.
func (loc *InstanceLocator) RunExecutable(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var ignoreLockOpt = options["ignore_lock"]
	if ignoreLockOpt != nil {
		p["ignore_lock"] = ignoreLockOpt
	}
	var inputsOpt = options["inputs"]
	if inputsOpt != nil {
		p["inputs"] = inputsOpt
	}
	var recipeNameOpt = options["recipe_name"]
	if recipeNameOpt != nil {
		p["recipe_name"] = recipeNameOpt
	}
	var rightScriptHrefOpt = options["right_script_href"]
	if rightScriptHrefOpt != nil {
		p["right_script_href"] = rightScriptHrefOpt
	}
	uri, err := loc.ActionPath("Instance", "run_executable")
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

// GET /api/clouds/:cloud_id/instances/:id
//
// Shows attributes of a single instance.
// The 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// Optional parameters:
// view
func (loc *InstanceLocator) Show(options rsapi.APIParams) (*Instance, error) {
	var res *Instance
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "show")
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

// POST /api/clouds/:cloud_id/instances/:id/start
//
// Starts an instance that has been stopped, resuming it to its previously saved volume state.
// After an instance is started, the reference to your instance will have a different id.
// The new id can be found by performing an index query with the appropriate filters on the
// Instances resource, performing a show action on the Server resource for Server Instances, or
// performing a current_instances action on the ServerArray resource for ServerArray Instances.
func (loc *InstanceLocator) Start() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "start")
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

// POST /api/clouds/:cloud_id/instances/:id/stop
//
// Stores the instance's current volume state to resume later using the 'start' action.
// After an instance is stopped, the reference to your instance will have a different id.
// The new id can be found by performing an index query with the appropriate filters on the
// Instances resource, performing a show action on the Server resource for Server Instances, or performing a current_instances action on the ServerArray resource for ServerArray Instances.
func (loc *InstanceLocator) Stop() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "stop")
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

// POST /api/clouds/:cloud_id/instances/:id/terminate
// POST /api/servers/:server_id/terminate
//
// Terminates a running instance.
// Note that this action can succeed only if the instance is running. One cannot terminate instances of type "next".
func (loc *InstanceLocator) Terminate() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "terminate")
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

// POST /api/clouds/:cloud_id/instances/:id/unlock
//

func (loc *InstanceLocator) Unlock() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "unlock")
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

// PUT /api/clouds/:cloud_id/instances/:id
//
// Updates attributes of a single instance.
// Required parameters:
// instance
func (loc *InstanceLocator) Update(instance *InstanceParam2) error {
	if instance == nil {
		return fmt.Errorf("instance is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"instance": instance,
	}
	uri, err := loc.ActionPath("Instance", "update")
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

/******  InstanceType ******/

type InstanceType struct {
	Actions         []map[string]string `json:"actions,omitempty"`
	CpuArchitecture string              `json:"cpu_architecture,omitempty"`
	CpuCount        int                 `json:"cpu_count,omitempty"`
	CpuSpeed        string              `json:"cpu_speed,omitempty"`
	Description     string              `json:"description,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	LocalDiskSize   int                 `json:"local_disk_size,omitempty"`
	LocalDisks      int                 `json:"local_disks,omitempty"`
	Memory          string              `json:"memory,omitempty"`
	Name            string              `json:"name,omitempty"`
	ResourceUid     string              `json:"resource_uid,omitempty"`
}

// Locator returns a locator for the given resource
func (r *InstanceType) Locator(api *API) *InstanceTypeLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.InstanceTypeLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// InstanceTypeLocator exposes the InstanceType resource actions.
type InstanceTypeLocator struct {
	Href
	api *API
}

// InstanceTypeLocator builds a locator from the given href.
func (api *API) InstanceTypeLocator(href string) *InstanceTypeLocator {
	return &InstanceTypeLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instance_types
//
// Lists instance types.
// Optional parameters:
// filter
// view
func (loc *InstanceTypeLocator) Index(options rsapi.APIParams) ([]*InstanceType, error) {
	var res []*InstanceType
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceType", "index")
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

// GET /api/clouds/:cloud_id/instance_types/:id
//
// Displays information about a single Instance type.
// Optional parameters:
// view
func (loc *InstanceTypeLocator) Show(options rsapi.APIParams) (*InstanceType, error) {
	var res *InstanceType
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceType", "show")
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

/******  IpAddress ******/

// An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.
type IpAddress struct {
	Address   string              `json:"address,omitempty"`
	CreatedAt *RubyTime           `json:"created_at,omitempty"`
	Domain    string              `json:"domain,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	Name      string              `json:"name,omitempty"`
	UpdatedAt *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *IpAddress) Locator(api *API) *IpAddressLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.IpAddressLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// IpAddressLocator exposes the IpAddress resource actions.
type IpAddressLocator struct {
	Href
	api *API
}

// IpAddressLocator builds a locator from the given href.
func (api *API) IpAddressLocator(href string) *IpAddressLocator {
	return &IpAddressLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/ip_addresses
//
// Creates a new IpAddress with the given parameters.
// Required parameters:
// ip_address
func (loc *IpAddressLocator) Create(ipAddress *IpAddressParam) (*IpAddressLocator, error) {
	var res *IpAddressLocator
	if ipAddress == nil {
		return res, fmt.Errorf("ipAddress is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"ip_address": ipAddress,
	}
	uri, err := loc.ActionPath("IpAddress", "create")
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
		return &IpAddressLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/ip_addresses/:id
//
// Deletes a given IpAddress.
func (loc *IpAddressLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddress", "destroy")
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

// GET /api/clouds/:cloud_id/ip_addresses
//
// Lists the IpAddresses available to this account.
// Optional parameters:
// filter
func (loc *IpAddressLocator) Index(options rsapi.APIParams) ([]*IpAddress, error) {
	var res []*IpAddress
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddress", "index")
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

// GET /api/clouds/:cloud_id/ip_addresses/:id
//
// Show information about a single IpAddress.
func (loc *IpAddressLocator) Show() (*IpAddress, error) {
	var res *IpAddress
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddress", "show")
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

// PUT /api/clouds/:cloud_id/ip_addresses/:id
//
// Updates attributes of a given IpAddress.
// Required parameters:
// ip_address
func (loc *IpAddressLocator) Update(ipAddress *IpAddressParam2) error {
	if ipAddress == nil {
		return fmt.Errorf("ipAddress is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"ip_address": ipAddress,
	}
	uri, err := loc.ActionPath("IpAddress", "update")
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

/******  IpAddressBinding ******/

// An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.
// The IpAddress is bound immediately for a current instance, or on launch for a next instance.
// It also allows specifying port forwarding rules for that particular IpAddress and Instance pair.
type IpAddressBinding struct {
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	PrivatePort string              `json:"private_port,omitempty"`
	Protocol    string              `json:"protocol,omitempty"`
	PublicPort  string              `json:"public_port,omitempty"`
	Recurring   bool                `json:"recurring,omitempty"`
}

// Locator returns a locator for the given resource
func (r *IpAddressBinding) Locator(api *API) *IpAddressBindingLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.IpAddressBindingLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// IpAddressBindingLocator exposes the IpAddressBinding resource actions.
type IpAddressBindingLocator struct {
	Href
	api *API
}

// IpAddressBindingLocator builds a locator from the given href.
func (api *API) IpAddressBindingLocator(href string) *IpAddressBindingLocator {
	return &IpAddressBindingLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings
// POST /api/clouds/:cloud_id/ip_address_bindings
//
// Creates an ip address binding which attaches a specified IpAddress resource
// to a specified instance, and also allows for configuration of port forwarding
// rules. If the instance specified is a current (running) instance, a one-time
// IpAddressBinding will be created. If the instance is a next instance, then
// a recurring IpAddressBinding is created, which will cause the IpAddress to
// be bound each time the incarnator boots.
// Required parameters:
// ip_address_binding
func (loc *IpAddressBindingLocator) Create(ipAddressBinding *IpAddressBindingParam) (*IpAddressBindingLocator, error) {
	var res *IpAddressBindingLocator
	if ipAddressBinding == nil {
		return res, fmt.Errorf("ipAddressBinding is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"ip_address_binding": ipAddressBinding,
	}
	uri, err := loc.ActionPath("IpAddressBinding", "create")
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
		return &IpAddressBindingLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id
// DELETE /api/clouds/:cloud_id/ip_address_bindings/:id
//
// No description provided for destroy.
func (loc *IpAddressBindingLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddressBinding", "destroy")
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

// GET /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings
// GET /api/clouds/:cloud_id/ip_address_bindings
//
// Lists the ip address bindings available to this account.
// Optional parameters:
// filter
func (loc *IpAddressBindingLocator) Index(options rsapi.APIParams) ([]*IpAddressBinding, error) {
	var res []*IpAddressBinding
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddressBinding", "index")
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

// GET /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id
// GET /api/clouds/:cloud_id/ip_address_bindings/:id
//
// Show information about a single ip address binding.
func (loc *IpAddressBindingLocator) Show() (*IpAddressBinding, error) {
	var res *IpAddressBinding
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddressBinding", "show")
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

/******  MonitoringMetric ******/

// A monitoring metric is a stream of data that is captured in an instance. Metrics can be monitored, graphed and can be used as the basis for triggering alerts.
type MonitoringMetric struct {
	Actions   []map[string]string `json:"actions,omitempty"`
	GraphHref string              `json:"graph_href,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	Plugin    string              `json:"plugin,omitempty"`
	View      string              `json:"view,omitempty"`
}

// Locator returns a locator for the given resource
func (r *MonitoringMetric) Locator(api *API) *MonitoringMetricLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.MonitoringMetricLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// MonitoringMetricLocator exposes the MonitoringMetric resource actions.
type MonitoringMetricLocator struct {
	Href
	api *API
}

// MonitoringMetricLocator builds a locator from the given href.
func (api *API) MonitoringMetricLocator(href string) *MonitoringMetricLocator {
	return &MonitoringMetricLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id/data
//
// Gives the raw monitoring data for a particular metric. The response will include different variables
// associated with that metric and the data points for each of those variables.
// To get the data for a certain duration, for e.g. for the last 10 minutes(600 secs), provide the variables
// start="-600" and end="0".
// Required parameters:
// end: An integer number of seconds from current time. e.g. -150 or 0
// start: An integer number of seconds from current time. e.g. -300
func (loc *MonitoringMetricLocator) Data(end string, start string) (map[string]interface{}, error) {
	var res map[string]interface{}
	if end == "" {
		return res, fmt.Errorf("end is required")
	}
	if start == "" {
		return res, fmt.Errorf("start is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"end":   end,
		"start": start,
	}
	uri, err := loc.ActionPath("MonitoringMetric", "data")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics
//
// Lists the monitoring metrics available for the instance and their corresponding graph hrefs.
// Making a request to the graph_href will return a png image corresponding to that monitoring metric.
// Optional parameters:
// filter
// period: The time scale for which the graph is generated. Default is 'day'
// size: The size of the graph to be generated. Default is 'small'.
// title: The title of the graph.
// tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.
func (loc *MonitoringMetricLocator) Index(options rsapi.APIParams) ([]*MonitoringMetric, error) {
	var res []*MonitoringMetric
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var periodOpt = options["period"]
	if periodOpt != nil {
		p["period"] = periodOpt
	}
	var sizeOpt = options["size"]
	if sizeOpt != nil {
		p["size"] = sizeOpt
	}
	var titleOpt = options["title"]
	if titleOpt != nil {
		p["title"] = titleOpt
	}
	var tzOpt = options["tz"]
	if tzOpt != nil {
		p["tz"] = tzOpt
	}
	uri, err := loc.ActionPath("MonitoringMetric", "index")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id
//
// Shows attributes of a single monitoring metric.
// Making a request to the graph_href will return a png image corresponding to that monitoring metric.
// Optional parameters:
// period: The time scale for which the graph is generated. Default is 'day'.
// size: The size of the graph to be generated. Default is 'small'.
// title: The title of the graph.
// tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.
func (loc *MonitoringMetricLocator) Show(options rsapi.APIParams) (*MonitoringMetric, error) {
	var res *MonitoringMetric
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var periodOpt = options["period"]
	if periodOpt != nil {
		p["period"] = periodOpt
	}
	var sizeOpt = options["size"]
	if sizeOpt != nil {
		p["size"] = sizeOpt
	}
	var titleOpt = options["title"]
	if titleOpt != nil {
		p["title"] = titleOpt
	}
	var tzOpt = options["tz"]
	if tzOpt != nil {
		p["tz"] = tzOpt
	}
	uri, err := loc.ActionPath("MonitoringMetric", "show")
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

/******  MultiCloudImage ******/

// A MultiCloudImage is a RightScale component that functions as a pointer to machine images in specific clouds
// (e.g. AWS US-East, Rackspace). Each ServerTemplate can reference many MultiCloudImages that defines which
// image should be used when a server is launched in a particular cloud.
type MultiCloudImage struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Revision    int                 `json:"revision,omitempty"`
}

// Locator returns a locator for the given resource
func (r *MultiCloudImage) Locator(api *API) *MultiCloudImageLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.MultiCloudImageLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// MultiCloudImageLocator exposes the MultiCloudImage resource actions.
type MultiCloudImageLocator struct {
	Href
	api *API
}

// MultiCloudImageLocator builds a locator from the given href.
func (api *API) MultiCloudImageLocator(href string) *MultiCloudImageLocator {
	return &MultiCloudImageLocator{Href(href), api}
}

//===== Actions

// POST /api/multi_cloud_images/:id/clone
//
// Clones a given MultiCloudImage.
// Required parameters:
// multi_cloud_image
func (loc *MultiCloudImageLocator) Clone(multiCloudImage *MultiCloudImageParam) error {
	if multiCloudImage == nil {
		return fmt.Errorf("multiCloudImage is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"multi_cloud_image": multiCloudImage,
	}
	uri, err := loc.ActionPath("MultiCloudImage", "clone")
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

// POST /api/multi_cloud_images/:id/commit
//
// Commits a given MultiCloudImage. Only HEAD revisions can be committed.
// Required parameters:
// commit_message: The message associated with the commit.
func (loc *MultiCloudImageLocator) Commit(commitMessage string) error {
	if commitMessage == "" {
		return fmt.Errorf("commitMessage is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"commit_message": commitMessage,
	}
	uri, err := loc.ActionPath("MultiCloudImage", "commit")
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

// POST /api/server_templates/:server_template_id/multi_cloud_images
// POST /api/multi_cloud_images
//
// Creates a new MultiCloudImage with the given parameters.
// Required parameters:
// multi_cloud_image
func (loc *MultiCloudImageLocator) Create(multiCloudImage *MultiCloudImageParam) (*MultiCloudImageLocator, error) {
	var res *MultiCloudImageLocator
	if multiCloudImage == nil {
		return res, fmt.Errorf("multiCloudImage is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"multi_cloud_image": multiCloudImage,
	}
	uri, err := loc.ActionPath("MultiCloudImage", "create")
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
		return &MultiCloudImageLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/server_templates/:server_template_id/multi_cloud_images/:id
// DELETE /api/multi_cloud_images/:id
//
// Deletes a given MultiCloudImage.
func (loc *MultiCloudImageLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImage", "destroy")
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

// GET /api/server_templates/:server_template_id/multi_cloud_images
// GET /api/multi_cloud_images
//
// Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.
// Optional parameters:
// filter
func (loc *MultiCloudImageLocator) Index(options rsapi.APIParams) ([]*MultiCloudImage, error) {
	var res []*MultiCloudImage
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImage", "index")
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

// GET /api/server_templates/:server_template_id/multi_cloud_images/:id
// GET /api/multi_cloud_images/:id
//
// Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.
func (loc *MultiCloudImageLocator) Show() (*MultiCloudImage, error) {
	var res *MultiCloudImage
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImage", "show")
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

// PUT /api/server_templates/:server_template_id/multi_cloud_images/:id
// PUT /api/multi_cloud_images/:id
//
// Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).
// Currently, the attributes you can update are only the 'direct' attributes of a server template.
// Required parameters:
// multi_cloud_image
func (loc *MultiCloudImageLocator) Update(multiCloudImage *MultiCloudImageParam) error {
	if multiCloudImage == nil {
		return fmt.Errorf("multiCloudImage is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"multi_cloud_image": multiCloudImage,
	}
	uri, err := loc.ActionPath("MultiCloudImage", "update")
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

/******  MultiCloudImageMatcher ******/

// A MultiCloudImageMatcher generates MultiCloudImageSettings for all clouds of a
// given cloud type. For now, only one type of matcher is supported
// (fingerprint). Fingerprint will match images based upon a checksum as returned
// by the cloud and is supported CloudStack, OpenStack, and vSphere clouds. Pass
// in an example image with an image_href from which to generate the fingerprint.
type MultiCloudImageMatcher struct {
	Actions       []map[string]string `json:"actions,omitempty"`
	CloudType     string              `json:"cloud_type,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	MatchCriteria map[string]string   `json:"match_criteria,omitempty"`
	MatchType     string              `json:"match_type,omitempty"`
	UserData      string              `json:"user_data,omitempty"`
}

// Locator returns a locator for the given resource
func (r *MultiCloudImageMatcher) Locator(api *API) *MultiCloudImageMatcherLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.MultiCloudImageMatcherLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// MultiCloudImageMatcherLocator exposes the MultiCloudImageMatcher resource actions.
type MultiCloudImageMatcherLocator struct {
	Href
	api *API
}

// MultiCloudImageMatcherLocator builds a locator from the given href.
func (api *API) MultiCloudImageMatcherLocator(href string) *MultiCloudImageMatcherLocator {
	return &MultiCloudImageMatcherLocator{Href(href), api}
}

//===== Actions

// POST /api/multi_cloud_images/:multi_cloud_image_id/matchers
//
// Creates a new setting matcher for an existing MultiCloudImage.
// Required parameters:
// multi_cloud_image_matcher
func (loc *MultiCloudImageMatcherLocator) Create(multiCloudImageMatcher *MultiCloudImageMatcherParam) (*MultiCloudImageMatcherLocator, error) {
	var res *MultiCloudImageMatcherLocator
	if multiCloudImageMatcher == nil {
		return res, fmt.Errorf("multiCloudImageMatcher is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"multi_cloud_image_matcher": multiCloudImageMatcher,
	}
	uri, err := loc.ActionPath("MultiCloudImageMatcher", "create")
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
		return &MultiCloudImageMatcherLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/multi_cloud_images/:multi_cloud_image_id/matchers/:id
//
// Deletes a MultiCloudImage setting matcher.
func (loc *MultiCloudImageMatcherLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImageMatcher", "destroy")
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

// GET /api/multi_cloud_images/:multi_cloud_image_id/matchers
//
// Lists the MultiCloudImage setting matchers.
func (loc *MultiCloudImageMatcherLocator) Index() ([]*MultiCloudImageMatcher, error) {
	var res []*MultiCloudImageMatcher
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImageMatcher", "index")
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

// POST /api/multi_cloud_images/:multi_cloud_image_id/matchers/:id/rematch
//
// Generates new MultiCloudImageSettings based upon match_criteria. Returns hash of created/updated/destroyed settings.
func (loc *MultiCloudImageMatcherLocator) Rematch() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImageMatcher", "rematch")
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

// GET /api/multi_cloud_images/:multi_cloud_image_id/matchers/:id
//
// Show information about a single MultiCloudImage setting matcher.
func (loc *MultiCloudImageMatcherLocator) Show() (*MultiCloudImageMatcher, error) {
	var res *MultiCloudImageMatcher
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImageMatcher", "show")
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

/******  MultiCloudImageSetting ******/

// A MultiCloudImageSetting defines which
// settings should be used when a server is launched in a cloud.
type MultiCloudImageSetting struct {
	Actions  []map[string]string `json:"actions,omitempty"`
	Links    []map[string]string `json:"links,omitempty"`
	UserData string              `json:"user_data,omitempty"`
}

// Locator returns a locator for the given resource
func (r *MultiCloudImageSetting) Locator(api *API) *MultiCloudImageSettingLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.MultiCloudImageSettingLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// MultiCloudImageSettingLocator exposes the MultiCloudImageSetting resource actions.
type MultiCloudImageSettingLocator struct {
	Href
	api *API
}

// MultiCloudImageSettingLocator builds a locator from the given href.
func (api *API) MultiCloudImageSettingLocator(href string) *MultiCloudImageSettingLocator {
	return &MultiCloudImageSettingLocator{Href(href), api}
}

//===== Actions

// POST /api/multi_cloud_images/:multi_cloud_image_id/settings
//
// Creates a new setting for an existing MultiCloudImage.
// Required parameters:
// multi_cloud_image_setting
func (loc *MultiCloudImageSettingLocator) Create(multiCloudImageSetting *MultiCloudImageSettingParam) (*MultiCloudImageSettingLocator, error) {
	var res *MultiCloudImageSettingLocator
	if multiCloudImageSetting == nil {
		return res, fmt.Errorf("multiCloudImageSetting is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"multi_cloud_image_setting": multiCloudImageSetting,
	}
	uri, err := loc.ActionPath("MultiCloudImageSetting", "create")
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
		return &MultiCloudImageSettingLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/multi_cloud_images/:multi_cloud_image_id/settings/:id
//
// Deletes a MultiCloudImage setting.
func (loc *MultiCloudImageSettingLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImageSetting", "destroy")
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

// GET /api/multi_cloud_images/:multi_cloud_image_id/settings
//
// Lists the MultiCloudImage settings.
// Optional parameters:
// filter
func (loc *MultiCloudImageSettingLocator) Index(options rsapi.APIParams) ([]*MultiCloudImageSetting, error) {
	var res []*MultiCloudImageSetting
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImageSetting", "index")
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

// GET /api/multi_cloud_images/:multi_cloud_image_id/settings/:id
//
// Show information about a single MultiCloudImage setting.
func (loc *MultiCloudImageSettingLocator) Show() (*MultiCloudImageSetting, error) {
	var res *MultiCloudImageSetting
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImageSetting", "show")
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

// PUT /api/multi_cloud_images/:multi_cloud_image_id/settings/:id
//
// Updates a settings for a MultiCloudImage.
// Required parameters:
// multi_cloud_image_setting
func (loc *MultiCloudImageSettingLocator) Update(multiCloudImageSetting *MultiCloudImageSettingParam) error {
	if multiCloudImageSetting == nil {
		return fmt.Errorf("multiCloudImageSetting is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"multi_cloud_image_setting": multiCloudImageSetting,
	}
	uri, err := loc.ActionPath("MultiCloudImageSetting", "update")
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

/******  Network ******/

// A Network is a logical grouping of network devices.
type Network struct {
	Actions         []map[string]string `json:"actions,omitempty"`
	CidrBlock       string              `json:"cidr_block,omitempty"`
	Description     string              `json:"description,omitempty"`
	InstanceTenancy string              `json:"instance_tenancy,omitempty"`
	IsDefault       bool                `json:"is_default,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	ResourceUid     string              `json:"resource_uid,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Network) Locator(api *API) *NetworkLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.NetworkLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// NetworkLocator exposes the Network resource actions.
type NetworkLocator struct {
	Href
	api *API
}

// NetworkLocator builds a locator from the given href.
func (api *API) NetworkLocator(href string) *NetworkLocator {
	return &NetworkLocator{Href(href), api}
}

//===== Actions

// POST /api/networks
//
// Creates a new network.
// Required parameters:
// network
func (loc *NetworkLocator) Create(network *NetworkParam) (*NetworkLocator, error) {
	var res *NetworkLocator
	if network == nil {
		return res, fmt.Errorf("network is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network": network,
	}
	uri, err := loc.ActionPath("Network", "create")
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
		return &NetworkLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/networks/:id
//
// Deletes the given network(s).
func (loc *NetworkLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Network", "destroy")
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

// GET /api/networks
//
// Lists networks in this account.
// Optional parameters:
// filter
func (loc *NetworkLocator) Index(options rsapi.APIParams) ([]*Network, error) {
	var res []*Network
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Network", "index")
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

// GET /api/networks/:id
//
// Shows attributes of a single network.
func (loc *NetworkLocator) Show() (*Network, error) {
	var res *Network
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Network", "show")
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

// PUT /api/networks/:id
//
// Updates the given network.
// Required parameters:
// network
func (loc *NetworkLocator) Update(network *NetworkParam2) error {
	if network == nil {
		return fmt.Errorf("network is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network": network,
	}
	uri, err := loc.ActionPath("Network", "update")
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

/******  NetworkGateway ******/

// A NetworkGateway is an interface that allows traffic to be routed between networks.
type NetworkGateway struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	Type        string              `json:"type,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *NetworkGateway) Locator(api *API) *NetworkGatewayLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.NetworkGatewayLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// NetworkGatewayLocator exposes the NetworkGateway resource actions.
type NetworkGatewayLocator struct {
	Href
	api *API
}

// NetworkGatewayLocator builds a locator from the given href.
func (api *API) NetworkGatewayLocator(href string) *NetworkGatewayLocator {
	return &NetworkGatewayLocator{Href(href), api}
}

//===== Actions

// POST /api/network_gateways
//
// Create a new NetworkGateway.
// Required parameters:
// network_gateway
func (loc *NetworkGatewayLocator) Create(networkGateway *NetworkGatewayParam) (*NetworkGatewayLocator, error) {
	var res *NetworkGatewayLocator
	if networkGateway == nil {
		return res, fmt.Errorf("networkGateway is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network_gateway": networkGateway,
	}
	uri, err := loc.ActionPath("NetworkGateway", "create")
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
		return &NetworkGatewayLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/network_gateways/:id
//
// Delete an existing NetworkGateway.
func (loc *NetworkGatewayLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkGateway", "destroy")
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

// GET /api/network_gateways
//
// Lists the NetworkGateways available to this account.
// Optional parameters:
// filter
func (loc *NetworkGatewayLocator) Index(options rsapi.APIParams) ([]*NetworkGateway, error) {
	var res []*NetworkGateway
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkGateway", "index")
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

// GET /api/network_gateways/:id
//
// Show information about a single NetworkGateway.
func (loc *NetworkGatewayLocator) Show() (*NetworkGateway, error) {
	var res *NetworkGateway
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkGateway", "show")
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

// PUT /api/network_gateways/:id
//
// Update an existing NetworkGateway.
// Required parameters:
// network_gateway
func (loc *NetworkGatewayLocator) Update(networkGateway *NetworkGatewayParam2) error {
	if networkGateway == nil {
		return fmt.Errorf("networkGateway is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network_gateway": networkGateway,
	}
	uri, err := loc.ActionPath("NetworkGateway", "update")
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

/******  NetworkOptionGroup ******/

// A key/value pair hash containing options for configuring a Network.
// The key/value pairs are stored in the "options" parameter.
// Keys correspond to the type of option to set, and values correspond
// to the value of the particular option being set.
// Option keys that are supported vary depending on cloud -- please consult
// your particular cloud's documentation for available option keys.
type NetworkOptionGroup struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Options     map[string]string   `json:"options,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Type        string              `json:"type,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *NetworkOptionGroup) Locator(api *API) *NetworkOptionGroupLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.NetworkOptionGroupLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// NetworkOptionGroupLocator exposes the NetworkOptionGroup resource actions.
type NetworkOptionGroupLocator struct {
	Href
	api *API
}

// NetworkOptionGroupLocator builds a locator from the given href.
func (api *API) NetworkOptionGroupLocator(href string) *NetworkOptionGroupLocator {
	return &NetworkOptionGroupLocator{Href(href), api}
}

//===== Actions

// POST /api/network_option_groups
//
// Create a new NetworkOptionGroup.
// Required parameters:
// network_option_group
func (loc *NetworkOptionGroupLocator) Create(networkOptionGroup *NetworkOptionGroupParam) (*NetworkOptionGroupLocator, error) {
	var res *NetworkOptionGroupLocator
	if networkOptionGroup == nil {
		return res, fmt.Errorf("networkOptionGroup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network_option_group": networkOptionGroup,
	}
	uri, err := loc.ActionPath("NetworkOptionGroup", "create")
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
		return &NetworkOptionGroupLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/network_option_groups/:id
//
// Delete an existing NetworkOptionGroup.
func (loc *NetworkOptionGroupLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkOptionGroup", "destroy")
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

// GET /api/network_option_groups
//
// List NetworkOptionGroups available in this account.
// Optional parameters:
// filter
func (loc *NetworkOptionGroupLocator) Index(options rsapi.APIParams) ([]*NetworkOptionGroup, error) {
	var res []*NetworkOptionGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkOptionGroup", "index")
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

// GET /api/network_option_groups/:id
//
// Show information about a single NetworkOptionGroup.
func (loc *NetworkOptionGroupLocator) Show() (*NetworkOptionGroup, error) {
	var res *NetworkOptionGroup
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkOptionGroup", "show")
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

// PUT /api/network_option_groups/:id
//
// Update an existing NetworkOptionGroup.
// Required parameters:
// network_option_group
func (loc *NetworkOptionGroupLocator) Update(networkOptionGroup *NetworkOptionGroupParam2) error {
	if networkOptionGroup == nil {
		return fmt.Errorf("networkOptionGroup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network_option_group": networkOptionGroup,
	}
	uri, err := loc.ActionPath("NetworkOptionGroup", "update")
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

/******  NetworkOptionGroupAttachment ******/

// Resource for attaching NetworkOptionGroups to Networks.
// A single NetworkOptionGroup can be attached to many Networks.
// A Network/Subnet can have many NetworkOptionGroups attached, as long as the
// NetworkOptionGroups each have different types.
// This resource describes the attachment details between a particular
// NetworkOptionGroup and Network.
// Amazon currently only supports attaching NetworkOptionGroups to Networks.
// Other clouds in the future may support attaching to Subnets.
type NetworkOptionGroupAttachment struct {
	Actions            []map[string]string `json:"actions,omitempty"`
	CreatedAt          *RubyTime           `json:"created_at,omitempty"`
	Links              []map[string]string `json:"links,omitempty"`
	NetworkOptionGroup string              `json:"network_option_group,omitempty"`
	ResourceUid        string              `json:"resource_uid,omitempty"`
	UpdatedAt          *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *NetworkOptionGroupAttachment) Locator(api *API) *NetworkOptionGroupAttachmentLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.NetworkOptionGroupAttachmentLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// NetworkOptionGroupAttachmentLocator exposes the NetworkOptionGroupAttachment resource actions.
type NetworkOptionGroupAttachmentLocator struct {
	Href
	api *API
}

// NetworkOptionGroupAttachmentLocator builds a locator from the given href.
func (api *API) NetworkOptionGroupAttachmentLocator(href string) *NetworkOptionGroupAttachmentLocator {
	return &NetworkOptionGroupAttachmentLocator{Href(href), api}
}

//===== Actions

// POST /api/network_option_group_attachments
//
// Create a new NetworkOptionGroupAttachment.
// Required parameters:
// network_option_group_attachment
func (loc *NetworkOptionGroupAttachmentLocator) Create(networkOptionGroupAttachment *NetworkOptionGroupAttachmentParam) (*NetworkOptionGroupAttachmentLocator, error) {
	var res *NetworkOptionGroupAttachmentLocator
	if networkOptionGroupAttachment == nil {
		return res, fmt.Errorf("networkOptionGroupAttachment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network_option_group_attachment": networkOptionGroupAttachment,
	}
	uri, err := loc.ActionPath("NetworkOptionGroupAttachment", "create")
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
		return &NetworkOptionGroupAttachmentLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/network_option_group_attachments/:id
//
// Delete an existing NetworkOptionGroupAttachment.
func (loc *NetworkOptionGroupAttachmentLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkOptionGroupAttachment", "destroy")
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

// GET /api/network_option_group_attachments
//
// List NetworkOptionGroupAttachments in this account.
// Optional parameters:
// filter
// view
func (loc *NetworkOptionGroupAttachmentLocator) Index(options rsapi.APIParams) ([]*NetworkOptionGroupAttachment, error) {
	var res []*NetworkOptionGroupAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkOptionGroupAttachment", "index")
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

// GET /api/network_option_group_attachments/:id
//
// Show information about a single NetworkOptionGroupAttachment.
// Optional parameters:
// view
func (loc *NetworkOptionGroupAttachmentLocator) Show(options rsapi.APIParams) (*NetworkOptionGroupAttachment, error) {
	var res *NetworkOptionGroupAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkOptionGroupAttachment", "show")
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

// PUT /api/network_option_group_attachments/:id
//
// Update an existing NetworkOptionGroupAttachment.
// Required parameters:
// network_option_group_attachment
func (loc *NetworkOptionGroupAttachmentLocator) Update(networkOptionGroupAttachment *NetworkOptionGroupAttachmentParam2) error {
	if networkOptionGroupAttachment == nil {
		return fmt.Errorf("networkOptionGroupAttachment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"network_option_group_attachment": networkOptionGroupAttachment,
	}
	uri, err := loc.ActionPath("NetworkOptionGroupAttachment", "update")
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

/******  Oauth2 ******/

type Oauth2 struct {
}

//===== Locator

// Oauth2Locator exposes the Oauth2 resource actions.
type Oauth2Locator struct {
	Href
	api *API
}

// Oauth2Locator builds a locator from the given href.
func (api *API) Oauth2Locator(href string) *Oauth2Locator {
	return &Oauth2Locator{Href(href), api}
}

//===== Actions

// POST /api/oauth2/
//
// Perform an OAuth 2.0 token_refresh operation to obtain an access token that
// can be used in lieu of an API session cookie. (In other words, creates a
// session using OAuth 2.0).
// Note that an API-Version header is required with your request, and that the server
// may respond with a 301 Moved Permanently if you include an account_id parameter and
// your account is hosted in another RightScale cluster.
// The request parameters and response format are all as per the OAuth 2.0
// Internet Draft standard v23. In brief:
// - Successful responses include an access token, an expires-in timestamp, and a token type
// - The token type is always "bearer"
// - To use a bearer token, include header "Authorization: Bearer " with your API requests
// - The client must refresh the access token before it expires
// # Example Request using Curl (with prettified response):
// curl -i -H X-API-Version:1.5 -x POST https://my.rightscale.com/api/oauth2 -d "grant_type=refresh_token" -d "refresh_token=abcd1234deadbeef"
// {
// "access_token": "xyzzy",
// "expires_in":   3600,
// "token_type":   "bearer"
// }
// Required parameters:
// grant_type: Type of grant.
// Optional parameters:
// account_id: The client's account ID (only needed for instance agent clients).
// client_id: The client ID (only needed for confidential clients).
// client_secret: The client secret (only needed for confidential clients).
// r_s_version: The RightAgent protocol version the client conforms to (only needed for instance agent clients).
// refresh_token: The refresh token obtained from OAuth grant.
// right_link_version: The RightLink gem version the client conforms to (only needed for instance agent clients).
func (loc *Oauth2Locator) Create(grantType string, options rsapi.APIParams) (map[string]interface{}, error) {
	var res map[string]interface{}
	if grantType == "" {
		return res, fmt.Errorf("grantType is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"grant_type": grantType,
	}
	var accountIdOpt = options["account_id"]
	if accountIdOpt != nil {
		p["account_id"] = accountIdOpt
	}
	var clientIdOpt = options["client_id"]
	if clientIdOpt != nil {
		p["client_id"] = clientIdOpt
	}
	var clientSecretOpt = options["client_secret"]
	if clientSecretOpt != nil {
		p["client_secret"] = clientSecretOpt
	}
	var rsVersionOpt = options["r_s_version"]
	if rsVersionOpt != nil {
		p["r_s_version"] = rsVersionOpt
	}
	var refreshTokenOpt = options["refresh_token"]
	if refreshTokenOpt != nil {
		p["refresh_token"] = refreshTokenOpt
	}
	var rightLinkVersionOpt = options["right_link_version"]
	if rightLinkVersionOpt != nil {
		p["right_link_version"] = rightLinkVersionOpt
	}
	uri, err := loc.ActionPath("Oauth2", "create")
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

/******  Permission ******/

// Please note that API 1.5 does not support operations on Governance Groups
// or Orgs and only allows management of the following CM Roles:
// admin, actor, observer,
// aws_architect, publisher,
// designer, signup_wiz,
// enterprise_manager, server_login,
// library, security_manager,
// instance, server_superuser,
// infrastructure, ss_end_user,
// ss_designer, ss_observer
// Moreover, this API allows management of only roles granted directly
// on an account, to an individual user.
type Permission struct {
	Actions   []map[string]string `json:"actions,omitempty"`
	CreatedAt *RubyTime           `json:"created_at,omitempty"`
	DeletedAt *RubyTime           `json:"deleted_at,omitempty"`
	Id        string              `json:"id,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	RoleTitle string              `json:"role_title,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Permission) Locator(api *API) *PermissionLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.PermissionLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// PermissionLocator exposes the Permission resource actions.
type PermissionLocator struct {
	Href
	api *API
}

// PermissionLocator builds a locator from the given href.
func (api *API) PermissionLocator(href string) *PermissionLocator {
	return &PermissionLocator{Href(href), api}
}

//===== Actions

// POST /api/permissions
//
// Create a permission, thereby granting some user a particular role
// with respect to the current account.
// The 'observer' role has a special status; it must be granted before
// a user is eligible for any other permission in a given account.
// When provisioning users, always create the observer permission FIRST;
// creating any other permission before it will result in an error.
// For more information about the roles available and the privileges
// they confer, please refer to the following page of the RightScale
// support portal:
// http://support.rightscale.com/15-References/Lists/List_of_User_Roles
// Required parameters:
// permission
func (loc *PermissionLocator) Create(permission *PermissionParam) (*PermissionLocator, error) {
	var res *PermissionLocator
	if permission == nil {
		return res, fmt.Errorf("permission is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"permission": permission,
	}
	uri, err := loc.ActionPath("Permission", "create")
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
		return &PermissionLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/permissions/:id
//
// Destroy a permission, thereby revoking a user's role with respect
// to the current account.
// The 'observer' role has a special status; it cannot be revoked if
// a user has any other roles, because other roles become useless without
// being able to read data pertaining to the account.
// When deprovisioning user, always destroy the observer permission LAST;
// destroying it while the user has other permissions will result in an error.
func (loc *PermissionLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Permission", "destroy")
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

// GET /api/permissions
//
// List all permissions for all users of the current account.
// Optional parameters:
// filter
func (loc *PermissionLocator) Index(options rsapi.APIParams) ([]*Permission, error) {
	var res []*Permission
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Permission", "index")
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

// GET /api/permissions/:id
//
// Show information about a single permission.
func (loc *PermissionLocator) Show() (*Permission, error) {
	var res *Permission
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Permission", "show")
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

/******  PlacementGroup ******/

type PlacementGroup struct {
	Actions                 []map[string]string    `json:"actions,omitempty"`
	CloudSpecificAttributes map[string]interface{} `json:"cloud_specific_attributes,omitempty"`
	CreatedAt               *RubyTime              `json:"created_at,omitempty"`
	Description             string                 `json:"description,omitempty"`
	Links                   []map[string]string    `json:"links,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	ResourceUid             string                 `json:"resource_uid,omitempty"`
	State                   string                 `json:"state,omitempty"`
	UpdatedAt               *RubyTime              `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *PlacementGroup) Locator(api *API) *PlacementGroupLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.PlacementGroupLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// PlacementGroupLocator exposes the PlacementGroup resource actions.
type PlacementGroupLocator struct {
	Href
	api *API
}

// PlacementGroupLocator builds a locator from the given href.
func (api *API) PlacementGroupLocator(href string) *PlacementGroupLocator {
	return &PlacementGroupLocator{Href(href), api}
}

//===== Actions

// POST /api/placement_groups
//
// Creates a PlacementGroup.
// Required parameters:
// placement_group
func (loc *PlacementGroupLocator) Create(placementGroup *PlacementGroupParam) (*PlacementGroupLocator, error) {
	var res *PlacementGroupLocator
	if placementGroup == nil {
		return res, fmt.Errorf("placementGroup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"placement_group": placementGroup,
	}
	uri, err := loc.ActionPath("PlacementGroup", "create")
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
		return &PlacementGroupLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/placement_groups/:id
//
// Destroys a PlacementGroup.
func (loc *PlacementGroupLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PlacementGroup", "destroy")
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

// GET /api/placement_groups
//
// Lists all PlacementGroups in an account.
// Optional parameters:
// filter
// view
func (loc *PlacementGroupLocator) Index(options rsapi.APIParams) ([]*PlacementGroup, error) {
	var res []*PlacementGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PlacementGroup", "index")
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

// GET /api/placement_groups/:id
//
// Shows information about a single PlacementGroup.
// Optional parameters:
// view
func (loc *PlacementGroupLocator) Show(options rsapi.APIParams) (*PlacementGroup, error) {
	var res *PlacementGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PlacementGroup", "show")
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

/******  Preference ******/

// A Preference is a user and account-specific setting. Preferences are used in many part of the RightScale platform and can be used for custom purposes if desired.
type Preference struct {
	Actions   []map[string]string `json:"actions,omitempty"`
	Contents  string              `json:"contents,omitempty"`
	CreatedAt *RubyTime           `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Preference) Locator(api *API) *PreferenceLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.PreferenceLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// PreferenceLocator exposes the Preference resource actions.
type PreferenceLocator struct {
	Href
	api *API
}

// PreferenceLocator builds a locator from the given href.
func (api *API) PreferenceLocator(href string) *PreferenceLocator {
	return &PreferenceLocator{Href(href), api}
}

//===== Actions

// DELETE /api/preferences/:id
//
// Deletes the given preference.
func (loc *PreferenceLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Preference", "destroy")
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

// GET /api/preferences
//
// Lists all preferences.
// Optional parameters:
// filter
func (loc *PreferenceLocator) Index(options rsapi.APIParams) ([]*Preference, error) {
	var res []*Preference
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Preference", "index")
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

// GET /api/preferences/:id
//
// Shows a single preference.
func (loc *PreferenceLocator) Show() (*Preference, error) {
	var res *Preference
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Preference", "show")
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

// PUT /api/preferences/:id
//
// If 'id' is known, updates preference with given contents.
// Otherwise, creates new preference.
// Note: If create, will return '201 Created' and the location of the new preference.
// Required parameters:
// preference
func (loc *PreferenceLocator) Update(preference *PreferenceParam) error {
	if preference == nil {
		return fmt.Errorf("preference is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"preference": preference,
	}
	uri, err := loc.ActionPath("Preference", "update")
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

/******  Publication ******/

// A Publication is a revisioned component shared with a set of Account Groups.
// If shared with your account, it can be imported in to your account.
type Publication struct {
	Actions       []map[string]string    `json:"actions,omitempty"`
	CommitMessage map[string]interface{} `json:"commit_message,omitempty"`
	ContentType   string                 `json:"content_type,omitempty"`
	CreatedAt     *RubyTime              `json:"created_at,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Links         []map[string]string    `json:"links,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Publisher     string                 `json:"publisher,omitempty"`
	Revision      int                    `json:"revision,omitempty"`
	RevisionNotes string                 `json:"revision_notes,omitempty"`
	UpdatedAt     *RubyTime              `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Publication) Locator(api *API) *PublicationLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.PublicationLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// PublicationLocator exposes the Publication resource actions.
type PublicationLocator struct {
	Href
	api *API
}

// PublicationLocator builds a locator from the given href.
func (api *API) PublicationLocator(href string) *PublicationLocator {
	return &PublicationLocator{Href(href), api}
}

//===== Actions

// POST /api/publications/:id/import
//
// Imports the given publication and its subordinates to this account.
// Only non-HEAD revisions that are shared with the account can be imported.
func (loc *PublicationLocator) Import() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Publication", "import")
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

// GET /api/publications
//
// Lists the publications available to this account. Only non-HEAD revisions are possible.
// Optional parameters:
// filter
// view
func (loc *PublicationLocator) Index(options rsapi.APIParams) ([]*Publication, error) {
	var res []*Publication
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Publication", "index")
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

// GET /api/publications/:id
//
// Show information about a single publication. Only non-HEAD revisions are possible.
// Optional parameters:
// view
func (loc *PublicationLocator) Show(options rsapi.APIParams) (*Publication, error) {
	var res *Publication
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Publication", "show")
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

/******  PublicationLineage ******/

// A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.
// It is shared among all revisions of a Publication within the marketplace.
// Publication Lineages are different than lineages that exist within an account.
type PublicationLineage struct {
	Actions          []map[string]string `json:"actions,omitempty"`
	CommentsEmailed  bool                `json:"comments_emailed,omitempty"`
	CommentsEnabled  bool                `json:"comments_enabled,omitempty"`
	CreatedAt        *RubyTime           `json:"created_at,omitempty"`
	Links            []map[string]string `json:"links,omitempty"`
	LongDescription  string              `json:"long_description,omitempty"`
	Name             string              `json:"name,omitempty"`
	ShortDescription string              `json:"short_description,omitempty"`
	UpdatedAt        *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *PublicationLineage) Locator(api *API) *PublicationLineageLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.PublicationLineageLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// PublicationLineageLocator exposes the PublicationLineage resource actions.
type PublicationLineageLocator struct {
	Href
	api *API
}

// PublicationLineageLocator builds a locator from the given href.
func (api *API) PublicationLineageLocator(href string) *PublicationLineageLocator {
	return &PublicationLineageLocator{Href(href), api}
}

//===== Actions

// GET /api/publication_lineages/:id
//
// Show information about a single publication lineage. Only non-HEAD revisions are possible.
// Optional parameters:
// view
func (loc *PublicationLineageLocator) Show(options rsapi.APIParams) (*PublicationLineage, error) {
	var res *PublicationLineage
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("PublicationLineage", "show")
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

/******  RecurringVolumeAttachment ******/

// A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a Server/ServerArray the next time an instance is launched.
type RecurringVolumeAttachment struct {
	Actions      []map[string]string `json:"actions,omitempty"`
	CreatedAt    *RubyTime           `json:"created_at,omitempty"`
	Device       string              `json:"device,omitempty"`
	DeviceId     string              `json:"device_id,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	Name         string              `json:"name,omitempty"`
	RunnableType string              `json:"runnable_type,omitempty"`
	Size         string              `json:"size,omitempty"`
	Status       string              `json:"status,omitempty"`
	StorageType  string              `json:"storage_type,omitempty"`
	UpdatedAt    *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *RecurringVolumeAttachment) Locator(api *API) *RecurringVolumeAttachmentLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RecurringVolumeAttachmentLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RecurringVolumeAttachmentLocator exposes the RecurringVolumeAttachment resource actions.
type RecurringVolumeAttachmentLocator struct {
	Href
	api *API
}

// RecurringVolumeAttachmentLocator builds a locator from the given href.
func (api *API) RecurringVolumeAttachmentLocator(href string) *RecurringVolumeAttachmentLocator {
	return &RecurringVolumeAttachmentLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/recurring_volume_attachments
// POST /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments
// POST /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments
//
// Creates a new recurring volume attachment.
// Required parameters:
// recurring_volume_attachment
func (loc *RecurringVolumeAttachmentLocator) Create(recurringVolumeAttachment *RecurringVolumeAttachmentParam) (*RecurringVolumeAttachmentLocator, error) {
	var res *RecurringVolumeAttachmentLocator
	if recurringVolumeAttachment == nil {
		return res, fmt.Errorf("recurringVolumeAttachment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"recurring_volume_attachment": recurringVolumeAttachment,
	}
	uri, err := loc.ActionPath("RecurringVolumeAttachment", "create")
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
		return &RecurringVolumeAttachmentLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/recurring_volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments/:id
//
// Deletes a given recurring volume attachment.
func (loc *RecurringVolumeAttachmentLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RecurringVolumeAttachment", "destroy")
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

// GET /api/clouds/:cloud_id/recurring_volume_attachments
// GET /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments
// GET /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments
//
// Lists all recurring volume attachments.
// Optional parameters:
// filter
// view
func (loc *RecurringVolumeAttachmentLocator) Index(options rsapi.APIParams) ([]*RecurringVolumeAttachment, error) {
	var res []*RecurringVolumeAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RecurringVolumeAttachment", "index")
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

// GET /api/clouds/:cloud_id/recurring_volume_attachments/:id
// GET /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments/:id
// GET /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments/:id
//
// Displays information about a single recurring volume attachment.
// Optional parameters:
// view
func (loc *RecurringVolumeAttachmentLocator) Show(options rsapi.APIParams) (*RecurringVolumeAttachment, error) {
	var res *RecurringVolumeAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RecurringVolumeAttachment", "show")
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

/******  Repository ******/

// A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.
// RightScale currently supports the following types of repositores: git, svn, and URLs of compressed files (tar, tgz, gzip).
type Repository struct {
	Actions         []map[string]string `json:"actions,omitempty"`
	AssetCounts     int                 `json:"asset_counts,omitempty"`
	AssetPaths      []string            `json:"asset_paths,omitempty"`
	CommitReference string              `json:"commit_reference,omitempty"`
	CreatedAt       *RubyTime           `json:"created_at,omitempty"`
	Credentials     map[string]string   `json:"credentials,omitempty"`
	Description     string              `json:"description,omitempty"`
	FetchStatus     string              `json:"fetch_status,omitempty"`
	Id              string              `json:"id,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	ReadOnly        bool                `json:"read_only,omitempty"`
	Source          string              `json:"source,omitempty"`
	SourceType      string              `json:"source_type,omitempty"`
	UpdatedAt       *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Repository) Locator(api *API) *RepositoryLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RepositoryLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RepositoryLocator exposes the Repository resource actions.
type RepositoryLocator struct {
	Href
	api *API
}

// RepositoryLocator builds a locator from the given href.
func (api *API) RepositoryLocator(href string) *RepositoryLocator {
	return &RepositoryLocator{Href(href), api}
}

//===== Actions

// POST /api/repositories/:id/cookbook_import
//
// Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.
// Required parameters:
// asset_hrefs: Hrefs of the assets that should be imported.
// Optional parameters:
// follow: A flag indicating whether imported cookbooks should be followed.
// namespace: The namespace to import into.
// repository_commit_reference: Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.
// with_dependencies: A flag indicating whether dependencies should automatically be imported.
func (loc *RepositoryLocator) CookbookImport(assetHrefs []string, options rsapi.APIParams) error {
	if len(assetHrefs) == 0 {
		return fmt.Errorf("assetHrefs is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"asset_hrefs": assetHrefs,
	}
	var followOpt = options["follow"]
	if followOpt != nil {
		p["follow"] = followOpt
	}
	var namespaceOpt = options["namespace"]
	if namespaceOpt != nil {
		p["namespace"] = namespaceOpt
	}
	var repositoryCommitReferenceOpt = options["repository_commit_reference"]
	if repositoryCommitReferenceOpt != nil {
		p["repository_commit_reference"] = repositoryCommitReferenceOpt
	}
	var withDependenciesOpt = options["with_dependencies"]
	if withDependenciesOpt != nil {
		p["with_dependencies"] = withDependenciesOpt
	}
	uri, err := loc.ActionPath("Repository", "cookbook_import")
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

// POST /api/repositories/:id/cookbook_import_preview
//
// Retrieves a preview of the effects of a Cookbook import.
// NOTE: This action is for RightScale internal use only. The response is
// free-form JSON with no associated mediatype.
// DO NOT USE, THIS ACTION IS SUBJECT TO CHANGE AT ANYTIME.
// Required parameters:
// asset_hrefs: Hrefs of the assets that should be imported.
// namespace: The namespace to import into.
func (loc *RepositoryLocator) CookbookImportPreview(assetHrefs []string, namespace string) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	if len(assetHrefs) == 0 {
		return res, fmt.Errorf("assetHrefs is required")
	}
	if namespace == "" {
		return res, fmt.Errorf("namespace is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"asset_hrefs": assetHrefs,
		"namespace":   namespace,
	}
	uri, err := loc.ActionPath("Repository", "cookbook_import_preview")
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

// POST /api/repositories
//
// Creates a Repository.
// The following types of inputs are supported for the credential fields:
// Type
// Format
// Example(s)
// Text string
// text:&lt;value&gt;
// text:-----BEGIN RSA PRIVATE KEY-----text:secret
// Credential value
// cred:&lt;value&gt;
// cred:my ssh keycred:svn_1_password
// Required parameters:
// repository
func (loc *RepositoryLocator) Create(repository *RepositoryParam) (*RepositoryLocator, error) {
	var res *RepositoryLocator
	if repository == nil {
		return res, fmt.Errorf("repository is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"repository": repository,
	}
	uri, err := loc.ActionPath("Repository", "create")
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
		return &RepositoryLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/repositories/:id
//
// Deletes the specified Repositories.
func (loc *RepositoryLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Repository", "destroy")
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

// GET /api/repositories
//
// Lists all Repositories for this Account.
// Optional parameters:
// filter
// view
func (loc *RepositoryLocator) Index(options rsapi.APIParams) ([]*Repository, error) {
	var res []*Repository
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Repository", "index")
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

// POST /api/repositories/:id/refetch
//
// Refetches all RepositoryAssets associated with the Repository.
// Note that a refetch simply updates RightScale's view of the contents of the repository.
// You must perform an import to use the assets in your design objects (or use the auto import parameter).
// Optional parameters:
// auto_import: Whether cookbooks should automatically be imported after repositories are fetched.
func (loc *RepositoryLocator) Refetch(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var autoImportOpt = options["auto_import"]
	if autoImportOpt != nil {
		p["auto_import"] = autoImportOpt
	}
	uri, err := loc.ActionPath("Repository", "refetch")
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

// POST /api/repositories/resolve
//
// Show a list of repositories that have imported cookbooks with the given names.
// This operation returns a list of repositories that would later satisfy a call
// to the swap_repository
// action on a ServerTemplate.
// Optional parameters:
// imported_cookbook_name: A list of cookbook names that were imported by the repository.
func (loc *RepositoryLocator) Resolve(options rsapi.APIParams) ([]*Repository, error) {
	var res []*Repository
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var importedCookbookNameOpt = options["imported_cookbook_name"]
	if importedCookbookNameOpt != nil {
		p["imported_cookbook_name"] = importedCookbookNameOpt
	}
	uri, err := loc.ActionPath("Repository", "resolve")
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

// GET /api/repositories/:id
//
// Shows a specified Repository.
// Optional parameters:
// view
func (loc *RepositoryLocator) Show(options rsapi.APIParams) (*Repository, error) {
	var res *Repository
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Repository", "show")
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

// PUT /api/repositories/:id
//
// Updates a specified Repository.
// The following types of inputs are supported for the credential fields:
// Type
// Format
// Example(s)
// Text string
// text:&lt;value&gt;
// text:-----BEGIN RSA PRIVATE KEY-----text:secret
// Credential value
// cred:&lt;value&gt;
// cred:my ssh keycred:svn_1_password
// Required parameters:
// repository
func (loc *RepositoryLocator) Update(repository *RepositoryParam2) error {
	if repository == nil {
		return fmt.Errorf("repository is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"repository": repository,
	}
	uri, err := loc.ActionPath("Repository", "update")
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

/******  RepositoryAsset ******/

// A RepositoryAsset represents an item discovered in a Repository. These assets represent only a view of the Repository
// the last time it was scraped. In order to use these assets, you must import them into your account.
type RepositoryAsset struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Id          string              `json:"id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Version     string              `json:"version,omitempty"`
}

// Locator returns a locator for the given resource
func (r *RepositoryAsset) Locator(api *API) *RepositoryAssetLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RepositoryAssetLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RepositoryAssetLocator exposes the RepositoryAsset resource actions.
type RepositoryAssetLocator struct {
	Href
	api *API
}

// RepositoryAssetLocator builds a locator from the given href.
func (api *API) RepositoryAssetLocator(href string) *RepositoryAssetLocator {
	return &RepositoryAssetLocator{Href(href), api}
}

//===== Actions

// GET /api/repositories/:repository_id/repository_assets
//
// List a repository's current assets.
// Repository assests are the cookbook details that were scraped from a
// given repository.
// Optional parameters:
// view
func (loc *RepositoryAssetLocator) Index(options rsapi.APIParams) ([]*RepositoryAsset, error) {
	var res []*RepositoryAsset
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RepositoryAsset", "index")
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

// GET /api/repositories/:repository_id/repository_assets/:id
//
// Show information about a single asset.
// A repository assest are the cookbook details that were scraped from a
// repository.
// Optional parameters:
// view
func (loc *RepositoryAssetLocator) Show(options rsapi.APIParams) (*RepositoryAsset, error) {
	var res *RepositoryAsset
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RepositoryAsset", "show")
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

/******  ResourceGroup ******/

type ResourceGroup struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *ResourceGroup) Locator(api *API) *ResourceGroupLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.ResourceGroupLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// ResourceGroupLocator exposes the ResourceGroup resource actions.
type ResourceGroupLocator struct {
	Href
	api *API
}

// ResourceGroupLocator builds a locator from the given href.
func (api *API) ResourceGroupLocator(href string) *ResourceGroupLocator {
	return &ResourceGroupLocator{Href(href), api}
}

//===== Actions

// POST /api/resource_groups
//
// Creates a ResourceGroup.
// Required parameters:
// resource_group
func (loc *ResourceGroupLocator) Create(resourceGroup *ResourceGroupParam) (*ResourceGroupLocator, error) {
	var res *ResourceGroupLocator
	if resourceGroup == nil {
		return res, fmt.Errorf("resourceGroup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"resource_group": resourceGroup,
	}
	uri, err := loc.ActionPath("ResourceGroup", "create")
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
		return &ResourceGroupLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/resource_groups/:id
//
// Destroys a ResourceGroup.
func (loc *ResourceGroupLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ResourceGroup", "destroy")
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

// GET /api/resource_groups
//
// Lists all ResourceGroups in an account.
// Optional parameters:
// filter
func (loc *ResourceGroupLocator) Index(options rsapi.APIParams) ([]*ResourceGroup, error) {
	var res []*ResourceGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ResourceGroup", "index")
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

// GET /api/resource_groups/:id
//
// Shows information about a single ResourceGroup.
func (loc *ResourceGroupLocator) Show() (*ResourceGroup, error) {
	var res *ResourceGroup
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ResourceGroup", "show")
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

// PUT /api/resource_groups/:id
//
// Updates attributes of a given ResourceGroup.
// Required parameters:
// resource_group
func (loc *ResourceGroupLocator) Update(resourceGroup *ResourceGroupParam2) error {
	if resourceGroup == nil {
		return fmt.Errorf("resourceGroup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"resource_group": resourceGroup,
	}
	uri, err := loc.ActionPath("ResourceGroup", "update")
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

/******  RightScript ******/

// A RightScript is an executable piece of code that can be run on a server
// during the boot, operational, or decommission phases.
// All revisions of
// a RightScript belong to a RightScript lineage that is exposed by the
// "lineage" attribute (NOTE: This attribute is merely a string to locate
// all revisions of a RightScript and NOT a working URL).
type RightScript struct {
	CreatedAt   *RubyTime                `json:"created_at,omitempty"`
	Description string                   `json:"description,omitempty"`
	Id          string                   `json:"id,omitempty"`
	Inputs      []map[string]interface{} `json:"inputs,omitempty"`
	Lineage     string                   `json:"lineage,omitempty"`
	Links       []map[string]string      `json:"links,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Packages    string                   `json:"packages,omitempty"`
	Revision    int                      `json:"revision,omitempty"`
	Source      string                   `json:"source,omitempty"`
	UpdatedAt   *RubyTime                `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *RightScript) Locator(api *API) *RightScriptLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RightScriptLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RightScriptLocator exposes the RightScript resource actions.
type RightScriptLocator struct {
	Href
	api *API
}

// RightScriptLocator builds a locator from the given href.
func (api *API) RightScriptLocator(href string) *RightScriptLocator {
	return &RightScriptLocator{Href(href), api}
}

//===== Actions

// POST /api/right_scripts/:id/commit
//
// Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.
// Required parameters:
// right_script
func (loc *RightScriptLocator) Commit(rightScript *RightScriptParam) error {
	if rightScript == nil {
		return fmt.Errorf("rightScript is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"right_script": rightScript,
	}
	uri, err := loc.ActionPath("RightScript", "commit")
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

// POST /api/right_scripts
//
// No description provided for create.
// Required parameters:
// right_script
func (loc *RightScriptLocator) Create(rightScript *RightScriptParam2) (*RightScriptLocator, error) {
	var res *RightScriptLocator
	if rightScript == nil {
		return res, fmt.Errorf("rightScript is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"right_script": rightScript,
	}
	uri, err := loc.ActionPath("RightScript", "create")
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
		return &RightScriptLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/right_scripts/:id
//
// No description provided for destroy.
func (loc *RightScriptLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RightScript", "destroy")
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

// GET /api/right_scripts
//
// Lists RightScripts.
// Optional parameters:
// filter
// latest_only: Whether or not to return only the latest version for each lineage.
// view
func (loc *RightScriptLocator) Index(options rsapi.APIParams) ([]*RightScript, error) {
	var res []*RightScript
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var latestOnlyOpt = options["latest_only"]
	if latestOnlyOpt != nil {
		p["latest_only"] = latestOnlyOpt
	}
	uri, err := loc.ActionPath("RightScript", "index")
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

// GET /api/right_scripts/:id
//
// Displays information about a single RightScript.
// Optional parameters:
// view
func (loc *RightScriptLocator) Show(options rsapi.APIParams) (*RightScript, error) {
	var res *RightScript
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RightScript", "show")
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

// GET /api/right_scripts/:id/source
//
// Returns the script source for a RightScript
func (loc *RightScriptLocator) ShowSource() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RightScript", "show_source")
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

// PUT /api/right_scripts/:id
//
// Updates RightScript name/description
// Required parameters:
// right_script
func (loc *RightScriptLocator) Update(rightScript *RightScriptParam3) error {
	if rightScript == nil {
		return fmt.Errorf("rightScript is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"right_script": rightScript,
	}
	uri, err := loc.ActionPath("RightScript", "update")
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

// PUT /api/right_scripts/:id/source
//
// Updates the source of the given RightScript
// Required parameters:
// filename: The file name to update the RightScript source with.
func (loc *RightScriptLocator) UpdateSource(filename *rsapi.SourceUpload) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"filename": filename,
	}
	uri, err := loc.ActionPath("RightScript", "update_source")
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

/******  RightScriptAttachment ******/

type RightScriptAttachment struct {
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Digest      string              `json:"digest,omitempty"`
	DownloadUrl string              `json:"download_url,omitempty"`
	Filename    string              `json:"filename,omitempty"`
	Id          string              `json:"id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Size        string              `json:"size,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *RightScriptAttachment) Locator(api *API) *RightScriptAttachmentLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RightScriptAttachmentLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RightScriptAttachmentLocator exposes the RightScriptAttachment resource actions.
type RightScriptAttachmentLocator struct {
	Href
	api *API
}

// RightScriptAttachmentLocator builds a locator from the given href.
func (api *API) RightScriptAttachmentLocator(href string) *RightScriptAttachmentLocator {
	return &RightScriptAttachmentLocator{Href(href), api}
}

//===== Actions

// POST /api/right_scripts/:right_script_id/attachments
//
// Uploads the RightScript attachment links it to the RightScript. Create expects HTTP request to be formatted as multipart mime.
// Required parameters:
// right_script_attachment
func (loc *RightScriptAttachmentLocator) Create(rightScriptAttachment *RightScriptAttachmentParam) (*RightScriptAttachmentLocator, error) {
	var res *RightScriptAttachmentLocator
	if rightScriptAttachment == nil {
		return res, fmt.Errorf("rightScriptAttachment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"right_script_attachment": rightScriptAttachment,
	}
	uri, err := loc.ActionPath("RightScriptAttachment", "create")
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
		return &RightScriptAttachmentLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/right_scripts/:right_script_id/attachments/:id
//

func (loc *RightScriptAttachmentLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RightScriptAttachment", "destroy")
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

// GET /api/right_scripts/:right_script_id/attachments
//
// Lists RightScript attachments.
// Optional parameters:
// filter
// view
func (loc *RightScriptAttachmentLocator) Index(options rsapi.APIParams) ([]*RightScriptAttachment, error) {
	var res []*RightScriptAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RightScriptAttachment", "index")
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

// GET /api/right_scripts/:right_script_id/attachments/:id
//
// Displays information about a single RightScript attachment.
// Optional parameters:
// view
func (loc *RightScriptAttachmentLocator) Show(options rsapi.APIParams) (*RightScriptAttachment, error) {
	var res *RightScriptAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RightScriptAttachment", "show")
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

// PUT /api/right_scripts/:right_script_id/attachments/:id
//
// Uploads and updates existing attachment in a RightScript. Update expects HTTP request to formatted as multipart mime.
// Required parameters:
// right_script_attachment
func (loc *RightScriptAttachmentLocator) Update(rightScriptAttachment *RightScriptAttachmentParam) error {
	if rightScriptAttachment == nil {
		return fmt.Errorf("rightScriptAttachment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"right_script_attachment": rightScriptAttachment,
	}
	uri, err := loc.ActionPath("RightScriptAttachment", "update")
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

/******  Route ******/

// A Route defines how networking traffic should be routed from one
// destination to another. See next_hop_type for available endpoint targets.
type Route struct {
	CloudSpecificAttributes map[string]interface{} `json:"cloud_specific_attributes,omitempty"`
	CreatedAt               *RubyTime              `json:"created_at,omitempty"`
	Description             string                 `json:"description,omitempty"`
	DestinationCidrBlock    string                 `json:"destination_cidr_block,omitempty"`
	Links                   []map[string]string    `json:"links,omitempty"`
	NextHopIp               string                 `json:"next_hop_ip,omitempty"`
	NextHopType             string                 `json:"next_hop_type,omitempty"`
	NextHopUrl              string                 `json:"next_hop_url,omitempty"`
	ResourceUid             string                 `json:"resource_uid,omitempty"`
	State                   string                 `json:"state,omitempty"`
	UpdatedAt               *RubyTime              `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Route) Locator(api *API) *RouteLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RouteLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RouteLocator exposes the Route resource actions.
type RouteLocator struct {
	Href
	api *API
}

// RouteLocator builds a locator from the given href.
func (api *API) RouteLocator(href string) *RouteLocator {
	return &RouteLocator{Href(href), api}
}

//===== Actions

// POST /api/routes
// POST /api/route_tables/:route_table_id/routes
//
// Create a new Route.
// Required parameters:
// route
func (loc *RouteLocator) Create(route *RouteParam) (*RouteLocator, error) {
	var res *RouteLocator
	if route == nil {
		return res, fmt.Errorf("route is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"route": route,
	}
	uri, err := loc.ActionPath("Route", "create")
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
		return &RouteLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/routes/:id
// DELETE /api/route_tables/:route_table_id/routes/:id
//
// Delete an existing Route.
func (loc *RouteLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Route", "destroy")
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

// GET /api/routes
// GET /api/route_tables/:route_table_id/routes
//
// List Routes available in this account.
// Optional parameters:
// filter
func (loc *RouteLocator) Index(options rsapi.APIParams) ([]*Route, error) {
	var res []*Route
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Route", "index")
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

// GET /api/routes/:id
// GET /api/route_tables/:route_table_id/routes/:id
//
// Show information about a single Route.
func (loc *RouteLocator) Show() (*Route, error) {
	var res *Route
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Route", "show")
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

// PUT /api/routes/:id
// PUT /api/route_tables/:route_table_id/routes/:id
//
// Update an existing Route.
// Required parameters:
// route
func (loc *RouteLocator) Update(route *RouteParam2) error {
	if route == nil {
		return fmt.Errorf("route is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"route": route,
	}
	uri, err := loc.ActionPath("Route", "update")
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

/******  RouteTable ******/

// Grouped listing of Routes
type RouteTable struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Routes      []Route             `json:"routes,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *RouteTable) Locator(api *API) *RouteTableLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RouteTableLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RouteTableLocator exposes the RouteTable resource actions.
type RouteTableLocator struct {
	Href
	api *API
}

// RouteTableLocator builds a locator from the given href.
func (api *API) RouteTableLocator(href string) *RouteTableLocator {
	return &RouteTableLocator{Href(href), api}
}

//===== Actions

// POST /api/route_tables
//
// Create a new RouteTable.
// Required parameters:
// route_table
func (loc *RouteTableLocator) Create(routeTable *RouteTableParam) (*RouteTableLocator, error) {
	var res *RouteTableLocator
	if routeTable == nil {
		return res, fmt.Errorf("routeTable is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"route_table": routeTable,
	}
	uri, err := loc.ActionPath("RouteTable", "create")
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
		return &RouteTableLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/route_tables/:id
//
// Delete an existing RouteTable.
func (loc *RouteTableLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RouteTable", "destroy")
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

// GET /api/route_tables
//
// List RouteTables available in this account.
// Optional parameters:
// filter
// view
func (loc *RouteTableLocator) Index(options rsapi.APIParams) ([]*RouteTable, error) {
	var res []*RouteTable
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RouteTable", "index")
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

// GET /api/route_tables/:id
//
// Show information about a single RouteTable.
// Optional parameters:
// view
func (loc *RouteTableLocator) Show(options rsapi.APIParams) (*RouteTable, error) {
	var res *RouteTable
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RouteTable", "show")
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

// PUT /api/route_tables/:id
//
// Update an existing RouteTable.
// Required parameters:
// route_table
func (loc *RouteTableLocator) Update(routeTable *RouteTableParam2) error {
	if routeTable == nil {
		return fmt.Errorf("routeTable is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"route_table": routeTable,
	}
	uri, err := loc.ActionPath("RouteTable", "update")
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

/******  RunnableBinding ******/

// A RunnableBinding represents an item in a runlist of a ServerTemplate. These items could be
// RightScript or Chef recipes, and could be associated with any one of the three runlists of a
// ServerTemplate (boot, operational, decommission).
type RunnableBinding struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	Id          string              `json:"id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Position    int                 `json:"position,omitempty"`
	Recipe      string              `json:"recipe,omitempty"`
	RightScript RightScript         `json:"right_script,omitempty"`
	Sequence    string              `json:"sequence,omitempty"`
}

// Locator returns a locator for the given resource
func (r *RunnableBinding) Locator(api *API) *RunnableBindingLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.RunnableBindingLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// RunnableBindingLocator exposes the RunnableBinding resource actions.
type RunnableBindingLocator struct {
	Href
	api *API
}

// RunnableBindingLocator builds a locator from the given href.
func (api *API) RunnableBindingLocator(href string) *RunnableBindingLocator {
	return &RunnableBindingLocator{Href(href), api}
}

//===== Actions

// POST /api/server_templates/:server_template_id/runnable_bindings
//
// Bind an executable to the given ServerTemplate.
// An executable may be either a RightScript or Chef Recipe.
// The resource must be editable.
// Required parameters:
// runnable_binding
func (loc *RunnableBindingLocator) Create(runnableBinding *RunnableBindingParam) (*RunnableBindingLocator, error) {
	var res *RunnableBindingLocator
	if runnableBinding == nil {
		return res, fmt.Errorf("runnableBinding is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"runnable_binding": runnableBinding,
	}
	uri, err := loc.ActionPath("RunnableBinding", "create")
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
		return &RunnableBindingLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/server_templates/:server_template_id/runnable_bindings/:id
//
// Unbind an executable from the given resource.
// The resource must be editable.
func (loc *RunnableBindingLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RunnableBinding", "destroy")
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

// GET /api/server_templates/:server_template_id/runnable_bindings
//
// Lists the executables bound to the ServerTemplate.
// An excutable may be either a RightScript or Chef Recipe.
// Optional parameters:
// view
func (loc *RunnableBindingLocator) Index(options rsapi.APIParams) ([]*RunnableBinding, error) {
	var res []*RunnableBinding
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RunnableBinding", "index")
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

// PUT /api/server_templates/:server_template_id/runnable_bindings/multi_update
//
// Update attributes for multiple bound executables.
// The resource must be editable.
// Required parameters:
// runnable_bindings
func (loc *RunnableBindingLocator) MultiUpdate(runnableBindings []*RunnableBindings) error {
	if len(runnableBindings) == 0 {
		return fmt.Errorf("runnableBindings is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"runnable_bindings": runnableBindings,
	}
	uri, err := loc.ActionPath("RunnableBinding", "multi_update")
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

// GET /api/server_templates/:server_template_id/runnable_bindings/:id
//
// Show information about a single executable binding.
// An excutable may be either a RightScript or Chef Recipe.
// Optional parameters:
// view
func (loc *RunnableBindingLocator) Show(options rsapi.APIParams) (*RunnableBinding, error) {
	var res *RunnableBinding
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("RunnableBinding", "show")
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

/******  Scheduler ******/

// Provide RightLink with the ability to schedule script executions on instances
type Scheduler struct {
}

//===== Locator

// SchedulerLocator exposes the Scheduler resource actions.
type SchedulerLocator struct {
	Href
	api *API
}

// SchedulerLocator builds a locator from the given href.
func (api *API) SchedulerLocator(href string) *SchedulerLocator {
	return &SchedulerLocator{Href(href), api}
}

//===== Actions

// POST /api/right_net/scheduler/schedule_recipe
//
// Schedules a chef recipe for execution on the current instance
// Optional parameters:
// arguments: Serialized recipe execution arguments values keyed by name
// audit_id: Optional, reuse audit if specified
// audit_period: RunlistPolicy audit period
// formal_values: Formal input parameter values
// policy: RunlistPolicy policy name
// recipe: Chef recipe name, overridden by recipe_id
// recipe_id: ServerTemplateChefRecipe ID
// thread: RunlistPolicy thread name
func (loc *SchedulerLocator) ScheduleRecipe(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		p["arguments"] = argumentsOpt
	}
	var auditIdOpt = options["audit_id"]
	if auditIdOpt != nil {
		p["audit_id"] = auditIdOpt
	}
	var auditPeriodOpt = options["audit_period"]
	if auditPeriodOpt != nil {
		p["audit_period"] = auditPeriodOpt
	}
	var formalValuesOpt = options["formal_values"]
	if formalValuesOpt != nil {
		p["formal_values"] = formalValuesOpt
	}
	var policyOpt = options["policy"]
	if policyOpt != nil {
		p["policy"] = policyOpt
	}
	var recipeOpt = options["recipe"]
	if recipeOpt != nil {
		p["recipe"] = recipeOpt
	}
	var recipeIdOpt = options["recipe_id"]
	if recipeIdOpt != nil {
		p["recipe_id"] = recipeIdOpt
	}
	var threadOpt = options["thread"]
	if threadOpt != nil {
		p["thread"] = threadOpt
	}
	uri, err := loc.ActionPath("Scheduler", "schedule_recipe")
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

// POST /api/right_net/scheduler/schedule_right_script
//
// Schedules a RightScript for execution on the current instance
// Optional parameters:
// arguments: Serialized script execution arguments values keyed by name
// audit_id: Optional, reuse audit if specified
// audit_period: RunlistPolicy audit period
// formal_values: Formal input parameter values
// policy: RunlistPolicy policy name
// right_script: RightScript name, overridden by right_script_id
// right_script_id: RightScript ID
// thread: RunlistPolicy thread name
func (loc *SchedulerLocator) ScheduleRightScript(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		p["arguments"] = argumentsOpt
	}
	var auditIdOpt = options["audit_id"]
	if auditIdOpt != nil {
		p["audit_id"] = auditIdOpt
	}
	var auditPeriodOpt = options["audit_period"]
	if auditPeriodOpt != nil {
		p["audit_period"] = auditPeriodOpt
	}
	var formalValuesOpt = options["formal_values"]
	if formalValuesOpt != nil {
		p["formal_values"] = formalValuesOpt
	}
	var policyOpt = options["policy"]
	if policyOpt != nil {
		p["policy"] = policyOpt
	}
	var rightScriptOpt = options["right_script"]
	if rightScriptOpt != nil {
		p["right_script"] = rightScriptOpt
	}
	var rightScriptIdOpt = options["right_script_id"]
	if rightScriptIdOpt != nil {
		p["right_script_id"] = rightScriptIdOpt
	}
	var threadOpt = options["thread"]
	if threadOpt != nil {
		p["thread"] = threadOpt
	}
	uri, err := loc.ActionPath("Scheduler", "schedule_right_script")
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

/******  SecurityGroup ******/

// Security Groups represent network security profiles that contain lists of firewall rules for different ports and source IP addresses, as well as
// trust relationships amongst different security groups.
type SecurityGroup struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Href        string              `json:"href,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

// Locator returns a locator for the given resource
func (r *SecurityGroup) Locator(api *API) *SecurityGroupLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.SecurityGroupLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// SecurityGroupLocator exposes the SecurityGroup resource actions.
type SecurityGroupLocator struct {
	Href
	api *API
}

// SecurityGroupLocator builds a locator from the given href.
func (api *API) SecurityGroupLocator(href string) *SecurityGroupLocator {
	return &SecurityGroupLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/security_groups
//
// Create a security group.
// Required parameters:
// security_group
func (loc *SecurityGroupLocator) Create(securityGroup *SecurityGroupParam) (*SecurityGroupLocator, error) {
	var res *SecurityGroupLocator
	if securityGroup == nil {
		return res, fmt.Errorf("securityGroup is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"security_group": securityGroup,
	}
	uri, err := loc.ActionPath("SecurityGroup", "create")
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
		return &SecurityGroupLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/security_groups/:id
//
// Delete security group(s)
func (loc *SecurityGroupLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroup", "destroy")
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

// GET /api/clouds/:cloud_id/security_groups
//
// Lists Security Groups.
// Optional parameters:
// filter
// view
func (loc *SecurityGroupLocator) Index(options rsapi.APIParams) ([]*SecurityGroup, error) {
	var res []*SecurityGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroup", "index")
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

// GET /api/clouds/:cloud_id/security_groups/:id
//
// Displays information about a single Security Group.
// Optional parameters:
// view
func (loc *SecurityGroupLocator) Show(options rsapi.APIParams) (*SecurityGroup, error) {
	var res *SecurityGroup
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroup", "show")
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

/******  SecurityGroupRule ******/

type SecurityGroupRule struct {
	Action               string              `json:"action,omitempty"`
	Actions              []map[string]string `json:"actions,omitempty"`
	CidrIps              string              `json:"cidr_ips,omitempty"`
	Description          string              `json:"description,omitempty"`
	DestinationCidrIps   string              `json:"destination_cidr_ips,omitempty"`
	DestinationGroupName string              `json:"destination_group_name,omitempty"`
	Direction            string              `json:"direction,omitempty"`
	EndPort              string              `json:"end_port,omitempty"`
	GroupName            string              `json:"group_name,omitempty"`
	GroupOwner           string              `json:"group_owner,omitempty"`
	GroupUid             string              `json:"group_uid,omitempty"`
	Href                 string              `json:"href,omitempty"`
	IcmpCode             string              `json:"icmp_code,omitempty"`
	IcmpType             string              `json:"icmp_type,omitempty"`
	Links                []map[string]string `json:"links,omitempty"`
	Priority             int                 `json:"priority,omitempty"`
	Protocol             string              `json:"protocol,omitempty"`
	SourceCidrIps        string              `json:"source_cidr_ips,omitempty"`
	SourceEndPort        int                 `json:"source_end_port,omitempty"`
	SourceGroupName      string              `json:"source_group_name,omitempty"`
	SourceStartPort      int                 `json:"source_start_port,omitempty"`
	SourceType           string              `json:"source_type,omitempty"`
	StartPort            string              `json:"start_port,omitempty"`
}

// Locator returns a locator for the given resource
func (r *SecurityGroupRule) Locator(api *API) *SecurityGroupRuleLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.SecurityGroupRuleLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// SecurityGroupRuleLocator exposes the SecurityGroupRule resource actions.
type SecurityGroupRuleLocator struct {
	Href
	api *API
}

// SecurityGroupRuleLocator builds a locator from the given href.
func (api *API) SecurityGroupRuleLocator(href string) *SecurityGroupRuleLocator {
	return &SecurityGroupRuleLocator{Href(href), api}
}

//===== Actions

// POST /api/security_group_rules
// POST /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules
//
// Create a security group rule for a security group.
// The following flavors are supported:
// 1. group-based TCP/UDP
// 2. group-based ICMP
// 3. CIDR-based TCP/UDP
// 4. CIDR-based ICMP
// Required parameters:
// security_group_rule
func (loc *SecurityGroupRuleLocator) Create(securityGroupRule *SecurityGroupRuleParam) (*SecurityGroupRuleLocator, error) {
	var res *SecurityGroupRuleLocator
	if securityGroupRule == nil {
		return res, fmt.Errorf("securityGroupRule is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"security_group_rule": securityGroupRule,
	}
	uri, err := loc.ActionPath("SecurityGroupRule", "create")
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
		return &SecurityGroupRuleLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/security_group_rules/:id
// DELETE /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id
//
// Delete security group rule(s)
func (loc *SecurityGroupRuleLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroupRule", "destroy")
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

// GET /api/security_group_rules
// GET /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules
//
// Lists SecurityGroupRules.
// Optional parameters:
// view
func (loc *SecurityGroupRuleLocator) Index(options rsapi.APIParams) ([]*SecurityGroupRule, error) {
	var res []*SecurityGroupRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroupRule", "index")
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

// GET /api/security_group_rules/:id
// GET /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id
//
// Displays information about a single SecurityGroupRule.
// Optional parameters:
// view
func (loc *SecurityGroupRuleLocator) Show(options rsapi.APIParams) (*SecurityGroupRule, error) {
	var res *SecurityGroupRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroupRule", "show")
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

// PUT /api/security_group_rules/:id
// PUT /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id
//
// Required parameters:
// security_group_rule
func (loc *SecurityGroupRuleLocator) Update(securityGroupRule *SecurityGroupRuleParam2) error {
	if securityGroupRule == nil {
		return fmt.Errorf("securityGroupRule is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"security_group_rule": securityGroupRule,
	}
	uri, err := loc.ActionPath("SecurityGroupRule", "update")
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

/******  Server ******/

// Servers represent the notion of a server/machine from the RightScale's perspective. A Server, does not always
// have a corresponding VM running or provisioned in a cloud. Some clouds use the word "servers" to refer to created VM's. These allocated VM's
// are not called Servers in the RightScale API, they are called Instances.
// A Server always has a next_instance association, which will define the configuration to apply to a new instance when
// the server is launched or started (starting servers is not yet supported through this API).
// Once a Server is launched/started a current_instance relationship will exist.
// Accessing the current_instance of a server results in immediate runtime modification of this running server.
// Changes to the next_instance association prepares the
// configuration for the next instance launch/start (therefore they have no effect until such operation is performed).
type Server struct {
	Actions         []map[string]string `json:"actions,omitempty"`
	CreatedAt       *RubyTime           `json:"created_at,omitempty"`
	CurrentInstance *Instance           `json:"current_instance,omitempty"`
	Description     string              `json:"description,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	NextInstance    *Instance           `json:"next_instance,omitempty"`
	Optimized       bool                `json:"optimized,omitempty"`
	State           string              `json:"state,omitempty"`
	UpdatedAt       *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Server) Locator(api *API) *ServerLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.ServerLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// ServerLocator exposes the Server resource actions.
type ServerLocator struct {
	Href
	api *API
}

// ServerLocator builds a locator from the given href.
func (api *API) ServerLocator(href string) *ServerLocator {
	return &ServerLocator{Href(href), api}
}

//===== Actions

// POST /api/servers/:id/clone
//
// Clones a given server.
func (loc *ServerLocator) Clone() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "clone")
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

// POST /api/servers
// POST /api/deployments/:deployment_id/servers
//
// Creates a new server, and configures its corresponding "next" instance with the received parameters.
// Required parameters:
// server
func (loc *ServerLocator) Create(server *ServerParam) (*ServerLocator, error) {
	var res *ServerLocator
	if server == nil {
		return res, fmt.Errorf("server is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server": server,
	}
	uri, err := loc.ActionPath("Server", "create")
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
		return &ServerLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/servers/:id
// DELETE /api/deployments/:deployment_id/servers/:id
//
// Deletes a given server.
func (loc *ServerLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "destroy")
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

// POST /api/servers/:id/disable_runnable_bindings
//
// Disables a list of runnable bindings associated with a given server.
// Optional parameters:
// runnable_binding_hrefs: List of Runnable Bindings.
func (loc *ServerLocator) DisableRunnableBindings(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var runnableBindingHrefsOpt = options["runnable_binding_hrefs"]
	if runnableBindingHrefsOpt != nil {
		p["runnable_binding_hrefs"] = runnableBindingHrefsOpt
	}
	uri, err := loc.ActionPath("Server", "disable_runnable_bindings")
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

// POST /api/servers/:id/enable_runnable_bindings
//
// Enables a list of runnable bindings associated with a given server.
// Optional parameters:
// runnable_binding_hrefs: List of Runnable Bindings.
func (loc *ServerLocator) EnableRunnableBindings(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var runnableBindingHrefsOpt = options["runnable_binding_hrefs"]
	if runnableBindingHrefsOpt != nil {
		p["runnable_binding_hrefs"] = runnableBindingHrefsOpt
	}
	uri, err := loc.ActionPath("Server", "enable_runnable_bindings")
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

// GET /api/servers
// GET /api/deployments/:deployment_id/servers
//
// Lists servers.
// By using the available filters, it is possible to retrieve servers that have common characteristics.
// For example, one can list:
// * servers that have names that contain "app_server"
// * all servers of a given deployment
// For more filters, please see the 'index' action on 'Instances' resource as most of the attributes belong to
// a 'current_instance' than to a server.
// Optional parameters:
// filter
// view
func (loc *ServerLocator) Index(options rsapi.APIParams) ([]*Server, error) {
	var res []*Server
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "index")
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

// POST /api/servers/:id/launch
//
// Launches the "next" instance of this server. This function is equivalent to invoking the launch action on the
// URL of this servers next_instance. See Instances#launch for details.
// Optional parameters:
// api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
// count: For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error
// inputs
func (loc *ServerLocator) Launch(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var apiBehaviorOpt = options["api_behavior"]
	if apiBehaviorOpt != nil {
		p["api_behavior"] = apiBehaviorOpt
	}
	var countOpt = options["count"]
	if countOpt != nil {
		p["count"] = countOpt
	}
	var inputsOpt = options["inputs"]
	if inputsOpt != nil {
		p["inputs"] = inputsOpt
	}
	uri, err := loc.ActionPath("Server", "launch")
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

// GET /api/servers/:id
// GET /api/deployments/:deployment_id/servers/:id
//
// Shows the information of a single server.
// Optional parameters:
// view
func (loc *ServerLocator) Show(options rsapi.APIParams) (*Server, error) {
	var res *Server
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "show")
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

// POST /api/servers/:id/terminate
//
// Terminates the current instance of this server. This function is equivalent to invoking the terminate action on the
// URL of this servers current_instance. See Instances#terminate for details.
func (loc *ServerLocator) Terminate() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "terminate")
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

// POST /api/servers/:id/unwrap
// POST /api/deployments/:deployment_id/servers/:id/unwrap
//
// No description provided for unwrap.
func (loc *ServerLocator) Unwrap() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "unwrap")
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

// PUT /api/servers/:id
// PUT /api/deployments/:deployment_id/servers/:id
//
// Updates attributes of a single server.
// Required parameters:
// server
func (loc *ServerLocator) Update(server *ServerParam2) error {
	if server == nil {
		return fmt.Errorf("server is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server": server,
	}
	uri, err := loc.ActionPath("Server", "update")
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

// POST /api/servers/wrap_instance
// POST /api/deployments/:deployment_id/servers/wrap_instance
//
// Wrap an existing instance and set current instance for new server
// Required parameters:
// server
func (loc *ServerLocator) WrapInstance(server *ServerParam3) error {
	if server == nil {
		return fmt.Errorf("server is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server": server,
	}
	uri, err := loc.ActionPath("Server", "wrap_instance")
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

/******  ServerArray ******/

// A server array represents a logical group of instances and allows to resize(grow/shrink) that group based on certain elasticity parameters.
// A server array just like a server always has a next_instance association, which will define the configuration to apply when a new instance is launched.
// But unlike a server which has a current_instance relationship, the server array has a
// current_instances relationship that gives the information about
// all the running instances in the array. Changes to the next_instance association prepares the configuration for the next instance that is to be launched
// in the array and will therefore not affect any of the currently running instances.
type ServerArray struct {
	Actions          []map[string]string    `json:"actions,omitempty"`
	ArrayType        string                 `json:"array_type,omitempty"`
	CurrentInstances []Instance             `json:"current_instances,omitempty"`
	DatacenterPolicy []DatacenterPolicy     `json:"datacenter_policy,omitempty"`
	Description      string                 `json:"description,omitempty"`
	ElasticityParams map[string]interface{} `json:"elasticity_params,omitempty"`
	InstancesCount   int                    `json:"instances_count,omitempty"`
	Links            []map[string]string    `json:"links,omitempty"`
	Name             string                 `json:"name,omitempty"`
	NextInstance     *Instance              `json:"next_instance,omitempty"`
	Optimized        bool                   `json:"optimized,omitempty"`
	State            string                 `json:"state,omitempty"`
}

// Locator returns a locator for the given resource
func (r *ServerArray) Locator(api *API) *ServerArrayLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.ServerArrayLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// ServerArrayLocator exposes the ServerArray resource actions.
type ServerArrayLocator struct {
	Href
	api *API
}

// ServerArrayLocator builds a locator from the given href.
func (api *API) ServerArrayLocator(href string) *ServerArrayLocator {
	return &ServerArrayLocator{Href(href), api}
}

//===== Actions

// POST /api/server_arrays/:id/clone
//
// Clones a given server array.
func (loc *ServerArrayLocator) Clone() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerArray", "clone")
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

// POST /api/server_arrays
// POST /api/deployments/:deployment_id/server_arrays
//
// Creates a new server array, and configures its corresponding "next" instance with the received parameters.
// Required parameters:
// server_array
func (loc *ServerArrayLocator) Create(serverArray *ServerArrayParam) (*ServerArrayLocator, error) {
	var res *ServerArrayLocator
	if serverArray == nil {
		return res, fmt.Errorf("serverArray is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server_array": serverArray,
	}
	uri, err := loc.ActionPath("ServerArray", "create")
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
		return &ServerArrayLocator{Href(location), loc.api}, nil
	}
}

// GET /api/server_arrays/:id/current_instances
//
// List the running instances belonging to the server array. See Instances#index for details.
// This action is slightly different from invoking the index action on the Instances resource with the filter "parent_href == /api/server_arrays/XX" because the
// latter will include 'next_instance' as well.
// Optional parameters:
// filter
// view
func (loc *ServerArrayLocator) CurrentInstances(options rsapi.APIParams) ([]*Instance, error) {
	var res []*Instance
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerArray", "current_instances")
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

// DELETE /api/server_arrays/:id
// DELETE /api/deployments/:deployment_id/server_arrays/:id
//
// Deletes a given server array.
func (loc *ServerArrayLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerArray", "destroy")
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

// POST /api/server_arrays/:id/disable_runnable_bindings
//
// Disables a list of runnable bindings associated with a given server.
// Optional parameters:
// runnable_binding_hrefs: List of Runnable Bindings.
func (loc *ServerArrayLocator) DisableRunnableBindings(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var runnableBindingHrefsOpt = options["runnable_binding_hrefs"]
	if runnableBindingHrefsOpt != nil {
		p["runnable_binding_hrefs"] = runnableBindingHrefsOpt
	}
	uri, err := loc.ActionPath("ServerArray", "disable_runnable_bindings")
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

// POST /api/server_arrays/:id/enable_runnable_bindings
//
// Enables a list of runnable bindings associated with a given server.
// Optional parameters:
// runnable_binding_hrefs: List of Runnable Bindings.
func (loc *ServerArrayLocator) EnableRunnableBindings(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var runnableBindingHrefsOpt = options["runnable_binding_hrefs"]
	if runnableBindingHrefsOpt != nil {
		p["runnable_binding_hrefs"] = runnableBindingHrefsOpt
	}
	uri, err := loc.ActionPath("ServerArray", "enable_runnable_bindings")
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

// GET /api/server_arrays
// GET /api/deployments/:deployment_id/server_arrays
//
// Lists server arrays.
// By using the available filters, it is possible to retrieve server arrays that have common characteristics.
// For example, one can list:
// * arrays that have names that contain "my_server_array"
// * all arrays of a given deployment
// Optional parameters:
// filter
// view
func (loc *ServerArrayLocator) Index(options rsapi.APIParams) ([]*ServerArray, error) {
	var res []*ServerArray
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerArray", "index")
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

// POST /api/server_arrays/:id/launch
//
// Launches a new instance in the server array with the configuration defined in the 'next_instance'. This function is equivalent to invoking the launch action on the
// URL of this server_array's next_instance. See Instances#launch for details.
// Optional parameters:
// api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
// count: For Server Arrays, will launch the specified number of instances into the ServerArray. Attempting to call this action on non-server array objects will result in a parameter error
// inputs
func (loc *ServerArrayLocator) Launch(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var apiBehaviorOpt = options["api_behavior"]
	if apiBehaviorOpt != nil {
		p["api_behavior"] = apiBehaviorOpt
	}
	var countOpt = options["count"]
	if countOpt != nil {
		p["count"] = countOpt
	}
	var inputsOpt = options["inputs"]
	if inputsOpt != nil {
		p["inputs"] = inputsOpt
	}
	uri, err := loc.ActionPath("ServerArray", "launch")
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

// POST /api/server_arrays/:id/multi_run_executable
//
// Run an executable on all instances of this array. This function is equivalent to invoking the "multi_run_executable" action on the instances resource
// (Instances#multi_run_executable with the filter "parent_href == /api/server_arrays/XX"). To run an executable on a subset of the instances of the array, provide additional filters. To run an executable
// a single instance, invoke the action "run_executable" directly on the instance (see Instances#run_executable)
// Optional parameters:
// filter
// ignore_lock: Specifies the ability to ignore the lock(s) on the Instance(s).
// inputs
// recipe_name: The name of the recipe to be run.
// right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.
func (loc *ServerArrayLocator) MultiRunExecutable(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var ignoreLockOpt = options["ignore_lock"]
	if ignoreLockOpt != nil {
		p["ignore_lock"] = ignoreLockOpt
	}
	var inputsOpt = options["inputs"]
	if inputsOpt != nil {
		p["inputs"] = inputsOpt
	}
	var recipeNameOpt = options["recipe_name"]
	if recipeNameOpt != nil {
		p["recipe_name"] = recipeNameOpt
	}
	var rightScriptHrefOpt = options["right_script_href"]
	if rightScriptHrefOpt != nil {
		p["right_script_href"] = rightScriptHrefOpt
	}
	uri, err := loc.ActionPath("ServerArray", "multi_run_executable")
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

// POST /api/server_arrays/:id/multi_terminate
//
// Terminate all instances of this array. This function is equivalent to invoking the "multi_terminate" action on the instances resource ( Instances#multi_terminate with
// the filter "parent_href == /api/server_arrays/XX"). To terminate a subset of the instances of the array, provide additional filters. To terminate a single instance,
// invoke the action "terminate" directly on the instance (see Instances#terminate)
// Optional parameters:
// filter
// terminate_all: Specifies the ability to terminate all instances.
func (loc *ServerArrayLocator) MultiTerminate(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var terminateAllOpt = options["terminate_all"]
	if terminateAllOpt != nil {
		p["terminate_all"] = terminateAllOpt
	}
	uri, err := loc.ActionPath("ServerArray", "multi_terminate")
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

// GET /api/server_arrays/:id
// GET /api/deployments/:deployment_id/server_arrays/:id
//
// Shows the information of a single server array.
// Optional parameters:
// view
func (loc *ServerArrayLocator) Show(options rsapi.APIParams) (*ServerArray, error) {
	var res *ServerArray
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerArray", "show")
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

// PUT /api/server_arrays/:id
// PUT /api/deployments/:deployment_id/server_arrays/:id
//
// Updates attributes of a single server array.
// Required parameters:
// server_array
func (loc *ServerArrayLocator) Update(serverArray *ServerArrayParam2) error {
	if serverArray == nil {
		return fmt.Errorf("serverArray is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server_array": serverArray,
	}
	uri, err := loc.ActionPath("ServerArray", "update")
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

/******  ServerTemplate ******/

// ServerTemplates allow you to pre-configure servers by starting from a base image and adding scripts that run during the boot,
// operational, and shutdown phases. A ServerTemplate is a description of how a new instance will be configured when it is
// provisioned by your cloud provider.
// All revisions of a ServerTemplate belong to a ServerTemplate lineage that is exposed by the "lineage" attribute.
// (NOTE: This attribute is merely a string to locate all revisions of a ServerTemplate and NOT a working URL)
type ServerTemplate struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Inputs      []map[string]string `json:"inputs,omitempty"`
	Lineage     string              `json:"lineage,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Revision    int                 `json:"revision,omitempty"`
}

// Locator returns a locator for the given resource
func (r *ServerTemplate) Locator(api *API) *ServerTemplateLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.ServerTemplateLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// ServerTemplateLocator exposes the ServerTemplate resource actions.
type ServerTemplateLocator struct {
	Href
	api *API
}

// ServerTemplateLocator builds a locator from the given href.
func (api *API) ServerTemplateLocator(href string) *ServerTemplateLocator {
	return &ServerTemplateLocator{Href(href), api}
}

//===== Actions

// POST /api/server_templates/:id/clone
//
// Clones a given ServerTemplate.
// Required parameters:
// server_template
func (loc *ServerTemplateLocator) Clone(serverTemplate *ServerTemplateParam) error {
	if serverTemplate == nil {
		return fmt.Errorf("serverTemplate is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server_template": serverTemplate,
	}
	uri, err := loc.ActionPath("ServerTemplate", "clone")
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

// POST /api/server_templates/:id/commit
//
// Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.
// Required parameters:
// commit_head_dependencies: Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.
// commit_message: The message associated with the commit.
// freeze_repositories: Freeze the repositories.
func (loc *ServerTemplateLocator) Commit(commitHeadDependencies string, commitMessage string, freezeRepositories string) error {
	if commitHeadDependencies == "" {
		return fmt.Errorf("commitHeadDependencies is required")
	}
	if commitMessage == "" {
		return fmt.Errorf("commitMessage is required")
	}
	if freezeRepositories == "" {
		return fmt.Errorf("freezeRepositories is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"commit_head_dependencies": commitHeadDependencies,
		"commit_message":           commitMessage,
		"freeze_repositories":      freezeRepositories,
	}
	uri, err := loc.ActionPath("ServerTemplate", "commit")
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

// POST /api/server_templates
//
// Creates a new ServerTemplate with the given parameters.
// Required parameters:
// server_template
func (loc *ServerTemplateLocator) Create(serverTemplate *ServerTemplateParam) (*ServerTemplateLocator, error) {
	var res *ServerTemplateLocator
	if serverTemplate == nil {
		return res, fmt.Errorf("serverTemplate is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server_template": serverTemplate,
	}
	uri, err := loc.ActionPath("ServerTemplate", "create")
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
		return &ServerTemplateLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/server_templates/:id
//
// Deletes a given ServerTemplate.
func (loc *ServerTemplateLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplate", "destroy")
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

// POST /api/server_templates/:id/detect_changes_in_head
//
// Identifies RightScripts attached to the resource that differ from their HEAD.
// If the attached revision of the RightScript is the HEAD, then this will indicate
// a difference between it and the latest committed revision in the same lineage.
func (loc *ServerTemplateLocator) DetectChangesInHead() ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplate", "detect_changes_in_head")
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

// GET /api/server_templates
//
// Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// Optional parameters:
// filter
// view
func (loc *ServerTemplateLocator) Index(options rsapi.APIParams) ([]*ServerTemplate, error) {
	var res []*ServerTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplate", "index")
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

// POST /api/server_templates/:id/publish
//
// Publishes a given ServerTemplate and its subordinates.
// Only non-HEAD revisions that are owned by the account can be published.
// Required parameters:
// account_group_hrefs: List of hrefs of account groups to publish to.
// descriptions
// Optional parameters:
// allow_comments: Allow users to leave comments on this ServerTemplate.
// categories: List of Categories.
// email_comments: Email me when a user comments on this ServerTemplate.
func (loc *ServerTemplateLocator) Publish(accountGroupHrefs []string, descriptions *Descriptions, options rsapi.APIParams) error {
	if len(accountGroupHrefs) == 0 {
		return fmt.Errorf("accountGroupHrefs is required")
	}
	if descriptions == nil {
		return fmt.Errorf("descriptions is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"account_group_hrefs": accountGroupHrefs,
		"descriptions":        descriptions,
	}
	var allowCommentsOpt = options["allow_comments"]
	if allowCommentsOpt != nil {
		p["allow_comments"] = allowCommentsOpt
	}
	var categoriesOpt = options["categories"]
	if categoriesOpt != nil {
		p["categories"] = categoriesOpt
	}
	var emailCommentsOpt = options["email_comments"]
	if emailCommentsOpt != nil {
		p["email_comments"] = emailCommentsOpt
	}
	uri, err := loc.ActionPath("ServerTemplate", "publish")
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

// POST /api/server_templates/:id/resolve
//
// Enumerates all attached cookbooks, missing dependencies and bound executables.
// Version constraints on missing dependencies and the state of the Chef Recipes;
// whether or not the cookbook or recipe itself could be found among the
// attachments, will also be reported.
func (loc *ServerTemplateLocator) Resolve() ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplate", "resolve")
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

// GET /api/server_templates/:id
//
// Show information about a single ServerTemplate. HEAD revisions have a revision of 0.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// Optional parameters:
// view
func (loc *ServerTemplateLocator) Show(options rsapi.APIParams) (*ServerTemplate, error) {
	var res *ServerTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplate", "show")
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

// POST /api/server_templates/:id/swap_repository
//
// In-place replacement of attached cookbooks from a given repository.
// For each attached cookbook coming from the source repository, replace it by
// attaching a cookbook of identical name coming from the target repository.
// In order for the operation to be successful, all attachments that came from the
// source repository must exist in the target repository.
// If multiple cookbooks of a given name exist in the target repository, preference
// is given by the following order (top most being the highest preference):
// Name & Version Match / Primary Namespace
// Name & Version Match / Alternate Namespace
// Name Match / Primary Namespace
// Name Match / Alternate Namespace
// If multiple cookbooks still have the same preference for the replacement, the operation is
// indeterministic.
// Required parameters:
// source_repository_href: The repository whose cookbook attachments are to be replaced.
// target_repository_href: The repository whose cookbook attachments are to be utilized.
func (loc *ServerTemplateLocator) SwapRepository(sourceRepositoryHref string, targetRepositoryHref string) error {
	if sourceRepositoryHref == "" {
		return fmt.Errorf("sourceRepositoryHref is required")
	}
	if targetRepositoryHref == "" {
		return fmt.Errorf("targetRepositoryHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"source_repository_href": sourceRepositoryHref,
		"target_repository_href": targetRepositoryHref,
	}
	uri, err := loc.ActionPath("ServerTemplate", "swap_repository")
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

// PUT /api/server_templates/:id
//
// Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).
// Currently, the attributes you can update are only the 'direct' attributes of a server template. To
// manage multi cloud images of a ServerTemplate, please see the resource 'ServerTemplateMultiCloudImages'.
// Required parameters:
// server_template
func (loc *ServerTemplateLocator) Update(serverTemplate *ServerTemplateParam) error {
	if serverTemplate == nil {
		return fmt.Errorf("serverTemplate is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server_template": serverTemplate,
	}
	uri, err := loc.ActionPath("ServerTemplate", "update")
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

/******  ServerTemplateMultiCloudImage ******/

// This resource represents links between ServerTemplates and MultiCloud Images and enables you to effectively
// add/delete MultiCloudImages to ServerTemplates and make them the default one.
type ServerTemplateMultiCloudImage struct {
	Actions   []map[string]string `json:"actions,omitempty"`
	CreatedAt *RubyTime           `json:"created_at,omitempty"`
	IsDefault bool                `json:"is_default,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *ServerTemplateMultiCloudImage) Locator(api *API) *ServerTemplateMultiCloudImageLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.ServerTemplateMultiCloudImageLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// ServerTemplateMultiCloudImageLocator exposes the ServerTemplateMultiCloudImage resource actions.
type ServerTemplateMultiCloudImageLocator struct {
	Href
	api *API
}

// ServerTemplateMultiCloudImageLocator builds a locator from the given href.
func (api *API) ServerTemplateMultiCloudImageLocator(href string) *ServerTemplateMultiCloudImageLocator {
	return &ServerTemplateMultiCloudImageLocator{Href(href), api}
}

//===== Actions

// POST /api/server_template_multi_cloud_images
//
// Creates a new ServerTemplateMultiCloudImage with the given parameters.
// Required parameters:
// server_template_multi_cloud_image
func (loc *ServerTemplateMultiCloudImageLocator) Create(serverTemplateMultiCloudImage *ServerTemplateMultiCloudImageParam) (*ServerTemplateMultiCloudImageLocator, error) {
	var res *ServerTemplateMultiCloudImageLocator
	if serverTemplateMultiCloudImage == nil {
		return res, fmt.Errorf("serverTemplateMultiCloudImage is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"server_template_multi_cloud_image": serverTemplateMultiCloudImage,
	}
	uri, err := loc.ActionPath("ServerTemplateMultiCloudImage", "create")
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
		return &ServerTemplateMultiCloudImageLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/server_template_multi_cloud_images/:id
//
// Deletes a given ServerTemplateMultiCloudImage.
func (loc *ServerTemplateMultiCloudImageLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplateMultiCloudImage", "destroy")
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

// GET /api/server_template_multi_cloud_images
//
// Lists the ServerTemplateMultiCloudImages available to this account.
// Optional parameters:
// filter
// view
func (loc *ServerTemplateMultiCloudImageLocator) Index(options rsapi.APIParams) ([]*ServerTemplateMultiCloudImage, error) {
	var res []*ServerTemplateMultiCloudImage
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplateMultiCloudImage", "index")
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

// POST /api/server_template_multi_cloud_images/:id/make_default
//
// Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.
func (loc *ServerTemplateMultiCloudImageLocator) MakeDefault() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplateMultiCloudImage", "make_default")
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

// GET /api/server_template_multi_cloud_images/:id
//
// Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.
// Optional parameters:
// view
func (loc *ServerTemplateMultiCloudImageLocator) Show(options rsapi.APIParams) (*ServerTemplateMultiCloudImage, error) {
	var res *ServerTemplateMultiCloudImage
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplateMultiCloudImage", "show")
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

/******  Session ******/

// The sessions resource is in charge of creating API sessions that are bound to a given account. The sequence for login into the API is the following:
// * Perform a POST request to /api/sessions ('create' action) to my.rightscale.com or to any more specific hosts saved from previous sessions.
// * If the targeted host is not appropriate for the specific account being accessed it will return a 302 http code with a URL with which the client must retry the same POST request.
// * If the targeted host is the right one and the login is successful, it will return a 204 http code, along with two cookies that will need to be saved and passed in any subsequent API request.
// * If there is an authentication or authorization problem with the POST request an error (typically 401 or 422 ) may be returned at any point in the above sequence.
// * If the session expires, it will return a 403 http code with a "Session cookie is expired or invalid" message.
// Note that all API calls irrespective of the resource it is acting on, should pass a header "X_API_VERSION" with the value "1.5".
type Session struct {
	Actions []map[string]string `json:"actions,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
	Message string              `json:"message,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Session) Locator(api *API) *SessionLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.SessionLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// SessionLocator exposes the Session resource actions.
type SessionLocator struct {
	Href
	api *API
}

// SessionLocator builds a locator from the given href.
func (api *API) SessionLocator(href string) *SessionLocator {
	return &SessionLocator{Href(href), api}
}

//===== Actions

// GET /api/sessions/accounts
//
// List all the accounts that a user has access to.
// This call may be executed outside of an existing session. Doing so requires passing a username and password in the
// request body. The idea is that it should be possible to list accounts that can be used to create a session.
// Upon successfully authenticating the credentials, the system will return a 200 OK code and return the list of accounts.
// If an 302 redirect code is returned, the client is responsible of re-issuing the GET request against the content of the received Location header, passing the exact same parameters again.
// Example Request using Curl (not using an existing session):
// curl -i -H X_API_VERSION:1.5 -X GET -d email='email@me.com' -d password='mypassword' https://my.rightscale.com/api/sessions/accounts
// Example Request using Curl (using an existing session):
// curl -i -H X_API_VERSION:1.5 -X GET -b mycookies https://my.rightscale.com/api/sessions/accounts
// Optional parameters:
// email: The email to login with if not using existing session.
// password: The corresponding password.
// view: Extended view shows account permissions and products
func (loc *SessionLocator) Accounts(options rsapi.APIParams) ([]*Account, error) {
	var res []*Account
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var emailOpt = options["email"]
	if emailOpt != nil {
		p["email"] = emailOpt
	}
	var passwordOpt = options["password"]
	if passwordOpt != nil {
		p["password"] = passwordOpt
	}
	uri, err := loc.ActionPath("Session", "accounts")
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

// GET /api/sessions
//
// Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what
// features are available within its privileges.
// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions
// Optional parameters:
// view: Whoami view provides links to the logged-in principal and the account being accessed
func (loc *SessionLocator) Index(options rsapi.APIParams) ([]*Session, error) {
	var res []*Session
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Session", "index")
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

// GET /api/sessions/instance
//
// Shows the full attributes of the instance (that has the token used to log-in).
// This call can be used by an instance to get it's own details.
// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions/instance
func (loc *SessionLocator) IndexInstanceSession() (*Instance, error) {
	var res *Instance
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Session", "index_instance_session")
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

/******  SshKey ******/

// Ssh Keys represent a created SSH Key that exists in the cloud.
// An ssh key might also contain the private part of the key, and can be used to login to instances launched with it.
type SshKey struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Material    string              `json:"material,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

// Locator returns a locator for the given resource
func (r *SshKey) Locator(api *API) *SshKeyLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.SshKeyLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// SshKeyLocator exposes the SshKey resource actions.
type SshKeyLocator struct {
	Href
	api *API
}

// SshKeyLocator builds a locator from the given href.
func (api *API) SshKeyLocator(href string) *SshKeyLocator {
	return &SshKeyLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/ssh_keys
//
// Creates a new ssh key.
// Required parameters:
// ssh_key
func (loc *SshKeyLocator) Create(sshKey *SshKeyParam) (*SshKeyLocator, error) {
	var res *SshKeyLocator
	if sshKey == nil {
		return res, fmt.Errorf("sshKey is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"ssh_key": sshKey,
	}
	uri, err := loc.ActionPath("SshKey", "create")
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
		return &SshKeyLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/ssh_keys/:id
//
// Deletes a given ssh key.
func (loc *SshKeyLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SshKey", "destroy")
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

// GET /api/clouds/:cloud_id/ssh_keys
//
// Lists ssh keys.
// Optional parameters:
// filter
// view
func (loc *SshKeyLocator) Index(options rsapi.APIParams) ([]*SshKey, error) {
	var res []*SshKey
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SshKey", "index")
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

// GET /api/clouds/:cloud_id/ssh_keys/:id
//
// Displays information about a single ssh key.
// Optional parameters:
// view
func (loc *SshKeyLocator) Show(options rsapi.APIParams) (*SshKey, error) {
	var res *SshKey
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SshKey", "show")
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

/******  Subnet ******/

// A Subnet is a logical grouping of network devices. An Instance can have many
// Subnets.
type Subnet struct {
	CidrBlock   string              `json:"cidr_block,omitempty"`
	Description string              `json:"description,omitempty"`
	IsDefault   bool                `json:"is_default,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	Visibility  string              `json:"visibility,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Subnet) Locator(api *API) *SubnetLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.SubnetLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// SubnetLocator exposes the Subnet resource actions.
type SubnetLocator struct {
	Href
	api *API
}

// SubnetLocator builds a locator from the given href.
func (api *API) SubnetLocator(href string) *SubnetLocator {
	return &SubnetLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances/:instance_id/subnets
// POST /api/clouds/:cloud_id/subnets
//
// Creates a new subnet.
// Required parameters:
// subnet
func (loc *SubnetLocator) Create(subnet *SubnetParam) (*SubnetLocator, error) {
	var res *SubnetLocator
	if subnet == nil {
		return res, fmt.Errorf("subnet is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"subnet": subnet,
	}
	uri, err := loc.ActionPath("Subnet", "create")
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
		return &SubnetLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/instances/:instance_id/subnets/:id
// DELETE /api/clouds/:cloud_id/subnets/:id
//
// Deletes the given subnet(s).
func (loc *SubnetLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Subnet", "destroy")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/subnets
// GET /api/clouds/:cloud_id/subnets
//
// Lists subnets of a given cloud.
// Optional parameters:
// filter
func (loc *SubnetLocator) Index(options rsapi.APIParams) ([]*Subnet, error) {
	var res []*Subnet
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Subnet", "index")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/subnets/:id
// GET /api/clouds/:cloud_id/subnets/:id
//
// Shows attributes of a single subnet.
func (loc *SubnetLocator) Show() (*Subnet, error) {
	var res *Subnet
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Subnet", "show")
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

// PUT /api/clouds/:cloud_id/instances/:instance_id/subnets/:id
// PUT /api/clouds/:cloud_id/subnets/:id
//
// Updates name and description for the current subnet.
// Required parameters:
// subnet
func (loc *SubnetLocator) Update(subnet *SubnetParam2) error {
	if subnet == nil {
		return fmt.Errorf("subnet is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"subnet": subnet,
	}
	uri, err := loc.ActionPath("Subnet", "update")
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

/******  Tag ******/

// A tag or machine tag is a useful way of attaching useful metadata to an object/resource.
// Tags are commonly used as an extra label or identifier.
// For example, you might want to add a tag to an EBS Snapshot or AMI so that you can find it more quickly.
type Tag struct {
}

//===== Locator

// TagLocator exposes the Tag resource actions.
type TagLocator struct {
	Href
	api *API
}

// TagLocator builds a locator from the given href.
func (api *API) TagLocator(href string) *TagLocator {
	return &TagLocator{Href(href), api}
}

//===== Actions

// POST /api/tags/by_resource
//
// Get tags for a list of resource hrefs.
// The hrefs can belong to various resource types and the tags for a non-existent href will be empty.
// Required parameters:
// resource_hrefs: Hrefs of the resources for which tags are to be returned.
func (loc *TagLocator) ByResource(resourceHrefs []string) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	if len(resourceHrefs) == 0 {
		return res, fmt.Errorf("resourceHrefs is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"resource_hrefs": resourceHrefs,
	}
	uri, err := loc.ActionPath("Tag", "by_resource")
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

// POST /api/tags/by_tag
//
// Search for resources having a list of tags in a specific resource_type.
// The search criteria can contain plain tags ("my_db_server"), machine tags ("server:db=true"), or
// namespace &amp; predicate wildcards ("server:db=*"). The result set includes links to the resources.
// If match_all is "true", then the search is an "AND" operation -- only resources having all the
// tags are returned. If match_all has any other value or is missing, the search is performed
// as an "OR" operation.
// PLEASE NOTE the behavior of the include_tags_with_prefix parameter: if it is absent,
// or blank, then you will find resources that match your query but receive no information about
// the tags that apply to those resources. (Your search will also complete much more quickly in
// this case.)
// For example, a search with tag[]="server:db=true" and include_tags_with_prefix="backup:"
// will return resources that are tagged as a DB server, and also return all "backup" related tags
// for every matching resource.
// Required parameters:
// resource_type: Search among a single resource type.
// tags: The tags which must be present on the resource.
// Optional parameters:
// include_tags_with_prefix: If included, all tags matching this prefix will be returned. If not included, no tags will be returned.
// match_all: If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.
// with_deleted: If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.
func (loc *TagLocator) ByTag(resourceType string, tags []string, options rsapi.APIParams) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	if resourceType == "" {
		return res, fmt.Errorf("resourceType is required")
	}
	if len(tags) == 0 {
		return res, fmt.Errorf("tags is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"resource_type": resourceType,
		"tags":          tags,
	}
	var includeTagsWithPrefixOpt = options["include_tags_with_prefix"]
	if includeTagsWithPrefixOpt != nil {
		p["include_tags_with_prefix"] = includeTagsWithPrefixOpt
	}
	var matchAllOpt = options["match_all"]
	if matchAllOpt != nil {
		p["match_all"] = matchAllOpt
	}
	var withDeletedOpt = options["with_deleted"]
	if withDeletedOpt != nil {
		p["with_deleted"] = withDeletedOpt
	}
	uri, err := loc.ActionPath("Tag", "by_tag")
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

// POST /api/tags/multi_add
//
// Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.
// The hrefs can belong to various resource types. If a resource for a href could not be found, an
// error is returned and no tags are added for any resource.
// No error will be raised if the resource already has the tag(s) you are trying to add.
// Note: At this point, tags on 'next_instance' are not supported and one has to add tags to the 'server'.
// Required parameters:
// resource_hrefs: Hrefs of the resources for which the tags are to be added.
// tags: Tags to be added.
func (loc *TagLocator) MultiAdd(resourceHrefs []string, tags []string) error {
	if len(resourceHrefs) == 0 {
		return fmt.Errorf("resourceHrefs is required")
	}
	if len(tags) == 0 {
		return fmt.Errorf("tags is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"resource_hrefs": resourceHrefs,
		"tags":           tags,
	}
	uri, err := loc.ActionPath("Tag", "multi_add")
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

// POST /api/tags/multi_delete
//
// Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.
// The hrefs can belong to various resource types. If a resource for a href could not be found, an
// error is returned and no tags are deleted for any resource.
// Note that no error will be raised if the resource does not have the tag(s) you are trying to delete.
// Required parameters:
// resource_hrefs: Hrefs of the resources for which tags are to be deleted.
// tags: Tags to be deleted.
func (loc *TagLocator) MultiDelete(resourceHrefs []string, tags []string) error {
	if len(resourceHrefs) == 0 {
		return fmt.Errorf("resourceHrefs is required")
	}
	if len(tags) == 0 {
		return fmt.Errorf("tags is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"resource_hrefs": resourceHrefs,
		"tags":           tags,
	}
	uri, err := loc.ActionPath("Tag", "multi_delete")
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

/******  Task ******/

// Tasks represent processes that happen (or have happened) asynchronously within the context of an Instance.
// An example of a type of task is an operational script that runs in an instance.
// Task resources can be returned by certain API calls, such as Instances.run_executable, Backups.restore, and others.
type Task struct {
	Actions []map[string]string `json:"actions,omitempty"`
	Detail  string              `json:"detail,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
	Summary string              `json:"summary,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Task) Locator(api *API) *TaskLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.TaskLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// TaskLocator exposes the Task resource actions.
type TaskLocator struct {
	Href
	api *API
}

// TaskLocator builds a locator from the given href.
func (api *API) TaskLocator(href string) *TaskLocator {
	return &TaskLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instances/:instance_id/live/tasks/:id
// GET /api/server_arrays/:server_array_id/live/tasks/:id
//
// Displays information of a given task within the context of an instance.
// Optional parameters:
// view
func (loc *TaskLocator) Show(options rsapi.APIParams) (*Task, error) {
	var res *Task
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Task", "show")
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

/******  User ******/

type User struct {
	Actions      []map[string]string `json:"actions,omitempty"`
	Company      string              `json:"company,omitempty"`
	CreatedAt    *RubyTime           `json:"created_at,omitempty"`
	Email        string              `json:"email,omitempty"`
	FirstName    string              `json:"first_name,omitempty"`
	LastName     string              `json:"last_name,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	LoginName    string              `json:"login_name,omitempty"`
	Phone        string              `json:"phone,omitempty"`
	PrincipalUid string              `json:"principal_uid,omitempty"`
	TimezoneName string              `json:"timezone_name,omitempty"`
	UpdatedAt    *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *User) Locator(api *API) *UserLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.UserLocator(l["href"])
		}
	}
	return nil
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
// Create a user. If a user already exists with the same email, that user will be returned.
// Creating a user alone will not enable the user to access this account. You have to create
// 'permissions' for that user before it can be used. Performing a 'show' on a new user
// will fail unless you immediately create an 'observer' permission on the current account.
// Note that information about users and their permissions must be propagated globally across all
// RightScale clusters, and this can take some time (less than 60 seconds under normal circumstances)
// so the users you create may not be able to login for a minute or two after you create them. However,
// you may create or destroy permissions for newly-created users with no delay.
// To create a user that will login using password authentication, include the 'password' parameter
// with your request.
// To create an SSO-enabled user, you must specify the identity_provider that will be vouching for
// this user's identity, as well as the principal_uid (SAML NameID or OpenID identity URL) that
// the identity provider will assert for this user. Identity providers should be specified by
// their API href; you can obtain a list of the identity providers available to your account by
// invoking the 'index' action of the identity_providers API resource.
// Required parameters:
// user
func (loc *UserLocator) Create(user *UserParam) (*UserLocator, error) {
	var res *UserLocator
	if user == nil {
		return res, fmt.Errorf("user is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"user": user,
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
// List the users available to the account the user is logged in to. Therefore, to list the users of
// a child account, the user has to login to the child account first.
// Optional parameters:
// filter
func (loc *UserLocator) Index(options rsapi.APIParams) ([]*User, error) {
	var res []*User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
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
// Show information about a single user.
func (loc *UserLocator) Show() (*User, error) {
	var res *User
	var params rsapi.APIParams
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

// PUT /api/users/:id
//
// Update a user's contact information, change their password, or update their SSO settings.
// In order to update a user record, one of the following criteria must be met:
// 1. You've authenticated and are the user being modified, and you provide a valid current_password.
// 2. You're an admin and the user is linked to your enterprise SSO provider.
// 3. You're an admin and the user's email matches the email_domain of your enterprise SSO provider.
// In other words: you can update yourself if you know your own password, you can update
// yourself or others if you're an admin and they're linked to your SSO provider, and you can update any user
// if you're an admin and their email address is known to belong to your organization.
// For information about enabling canonical email domain ownership for your enterprise, please
// talk to your RightScale account manager or contact our support team.
// To update a user's contact information, simply pass the desired values for email, first_name,
// and so forth.
// To update a user's password, provide the desired new_password.
// To set or update a user's SSO information, you may provide a just a principal_uid (to maintain the
// user's existing identity provider) or you may provide an identity_provider_href and a
// principal_uid (to switch identity providers as well as specify a new user identity).
// In the context of SAML, principal_uid is equivalent to the SAML NameID or Subject claim.
// RightScale cannot predict or influence the NameID value that your SAML IdP will send to us for
// Required parameters:
// user
func (loc *UserLocator) Update(user *UserParam2) error {
	if user == nil {
		return fmt.Errorf("user is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"user": user,
	}
	uri, err := loc.ActionPath("User", "update")
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

/******  UserData ******/

type UserData struct {
}

//===== Locator

// UserDataLocator exposes the UserData resource actions.
type UserDataLocator struct {
	Href
	api *API
}

// UserDataLocator builds a locator from the given href.
func (api *API) UserDataLocator(href string) *UserDataLocator {
	return &UserDataLocator{Href(href), api}
}

//===== Actions

// GET /api/user_data/
//
// No description provided for show.
func (loc *UserDataLocator) Show() (map[string]interface{}, error) {
	var res map[string]interface{}
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserData", "show")
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

/******  Volume ******/

// A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).
type Volume struct {
	Actions                 []map[string]string    `json:"actions,omitempty"`
	CloudSpecificAttributes map[string]interface{} `json:"cloud_specific_attributes,omitempty"`
	CreatedAt               *RubyTime              `json:"created_at,omitempty"`
	Description             string                 `json:"description,omitempty"`
	Iops                    string                 `json:"iops,omitempty"`
	Links                   []map[string]string    `json:"links,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	PlacementGroup          `json:"placement_group,omitempty"`
	ResourceUid             string    `json:"resource_uid,omitempty"`
	Size                    int       `json:"size,omitempty"`
	Status                  string    `json:"status,omitempty"`
	UpdatedAt               *RubyTime `json:"updated_at,omitempty"`
	VolumeType              string    `json:"volume_type,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Volume) Locator(api *API) *VolumeLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.VolumeLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// VolumeLocator exposes the Volume resource actions.
type VolumeLocator struct {
	Href
	api *API
}

// VolumeLocator builds a locator from the given href.
func (api *API) VolumeLocator(href string) *VolumeLocator {
	return &VolumeLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/volumes
//
// Creates a new volume.
// Required parameters:
// volume
func (loc *VolumeLocator) Create(volume *VolumeParam) (*VolumeLocator, error) {
	var res *VolumeLocator
	if volume == nil {
		return res, fmt.Errorf("volume is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"volume": volume,
	}
	uri, err := loc.ActionPath("Volume", "create")
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
		return &VolumeLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/volumes/:id
//
// Deletes a given volume.
func (loc *VolumeLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Volume", "destroy")
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

// GET /api/clouds/:cloud_id/volumes
//
// Lists volumes.
// Optional parameters:
// filter
// view
func (loc *VolumeLocator) Index(options rsapi.APIParams) ([]*Volume, error) {
	var res []*Volume
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Volume", "index")
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

// GET /api/clouds/:cloud_id/volumes/:id
//
// Displays information about a single volume.
// Optional parameters:
// view
func (loc *VolumeLocator) Show(options rsapi.APIParams) (*Volume, error) {
	var res *Volume
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Volume", "show")
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

// PUT /api/clouds/:cloud_id/volumes/:id
//
// No description provided for update.
// Required parameters:
// volume
func (loc *VolumeLocator) Update(volume *VolumeParam2) error {
	if volume == nil {
		return fmt.Errorf("volume is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"volume": volume,
	}
	uri, err := loc.ActionPath("Volume", "update")
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

/******  VolumeAttachment ******/

// A VolumeAttachment represents a relationship between a volume and an instance.
type VolumeAttachment struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Device      string              `json:"device,omitempty"`
	DeviceId    string              `json:"device_id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *VolumeAttachment) Locator(api *API) *VolumeAttachmentLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.VolumeAttachmentLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// VolumeAttachmentLocator exposes the VolumeAttachment resource actions.
type VolumeAttachmentLocator struct {
	Href
	api *API
}

// VolumeAttachmentLocator builds a locator from the given href.
func (api *API) VolumeAttachmentLocator(href string) *VolumeAttachmentLocator {
	return &VolumeAttachmentLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances/:instance_id/volume_attachments
// POST /api/clouds/:cloud_id/volume_attachments
// POST /api/clouds/:cloud_id/volumes/:volume_id/volume_attachments
// POST /api/clouds/:cloud_id/volumes/:volume_id/volume_attachment
//
// Creates a new volume attachment.
// Required parameters:
// volume_attachment
func (loc *VolumeAttachmentLocator) Create(volumeAttachment *VolumeAttachmentParam) (*VolumeAttachmentLocator, error) {
	var res *VolumeAttachmentLocator
	if volumeAttachment == nil {
		return res, fmt.Errorf("volumeAttachment is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"volume_attachment": volumeAttachment,
	}
	uri, err := loc.ActionPath("VolumeAttachment", "create")
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
		return &VolumeAttachmentLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volumes/:volume_id/volume_attachments
// DELETE /api/clouds/:cloud_id/volumes/:volume_id/volume_attachment
//
// Deletes a given volume attachment.
// Optional parameters:
// force: Specifies whether to force the detachment of a volume.
func (loc *VolumeAttachmentLocator) Destroy(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var forceOpt = options["force"]
	if forceOpt != nil {
		p["force"] = forceOpt
	}
	uri, err := loc.ActionPath("VolumeAttachment", "destroy")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/volume_attachments
// GET /api/clouds/:cloud_id/volume_attachments
//
// Lists all volume attachments.
// Optional parameters:
// filter
// view
func (loc *VolumeAttachmentLocator) Index(options rsapi.APIParams) ([]*VolumeAttachment, error) {
	var res []*VolumeAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("VolumeAttachment", "index")
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

// GET /api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id
// GET /api/clouds/:cloud_id/volume_attachments/:id
// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_attachments
// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_attachment
//
// Displays information about a single volume attachment.
// Optional parameters:
// view
func (loc *VolumeAttachmentLocator) Show(options rsapi.APIParams) (*VolumeAttachment, error) {
	var res *VolumeAttachment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("VolumeAttachment", "show")
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

/******  VolumeSnapshot ******/

// A VolumeSnapshot represents a Cloud storage volume at a particular point in time. One can create a snapshot regardless of whether or not a volume is attached to an Instance. When a snapshot is created,
// various meta data is retained such as a Created At timestamp, a unique Resource UID (e.g. vol-52EF05A9), the Volume Owner and Visibility (e.g. private or public).
// Snapshots consist of a series of data blocks that are incrementally saved.
type VolumeSnapshot struct {
	Actions                 []map[string]string    `json:"actions,omitempty"`
	CloudSpecificAttributes map[string]interface{} `json:"cloud_specific_attributes,omitempty"`
	CreatedAt               *RubyTime              `json:"created_at,omitempty"`
	Description             string                 `json:"description,omitempty"`
	Links                   []map[string]string    `json:"links,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	Progress                string                 `json:"progress,omitempty"`
	ResourceUid             string                 `json:"resource_uid,omitempty"`
	Size                    string                 `json:"size,omitempty"`
	State                   string                 `json:"state,omitempty"`
	UpdatedAt               *RubyTime              `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *VolumeSnapshot) Locator(api *API) *VolumeSnapshotLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.VolumeSnapshotLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// VolumeSnapshotLocator exposes the VolumeSnapshot resource actions.
type VolumeSnapshotLocator struct {
	Href
	api *API
}

// VolumeSnapshotLocator builds a locator from the given href.
func (api *API) VolumeSnapshotLocator(href string) *VolumeSnapshotLocator {
	return &VolumeSnapshotLocator{Href(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id/copy
// POST /api/clouds/:cloud_id/volume_snapshots/:id/copy
//
// No description provided for copy.
// Required parameters:
// volume_snapshot_copy
func (loc *VolumeSnapshotLocator) Copy(volumeSnapshotCopy *VolumeSnapshotCopy) error {
	if volumeSnapshotCopy == nil {
		return fmt.Errorf("volumeSnapshotCopy is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"volume_snapshot_copy": volumeSnapshotCopy,
	}
	uri, err := loc.ActionPath("VolumeSnapshot", "copy")
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

// POST /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots
// POST /api/clouds/:cloud_id/volume_snapshots
//
// Creates a new volume_snapshot.
// Required parameters:
// volume_snapshot
func (loc *VolumeSnapshotLocator) Create(volumeSnapshot *VolumeSnapshotParam) (*VolumeSnapshotLocator, error) {
	var res *VolumeSnapshotLocator
	if volumeSnapshot == nil {
		return res, fmt.Errorf("volumeSnapshot is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"volume_snapshot": volumeSnapshot,
	}
	uri, err := loc.ActionPath("VolumeSnapshot", "create")
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
		return &VolumeSnapshotLocator{Href(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id
// DELETE /api/clouds/:cloud_id/volume_snapshots/:id
//
// Deletes a given volume_snapshot.
func (loc *VolumeSnapshotLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("VolumeSnapshot", "destroy")
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

// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots
// GET /api/clouds/:cloud_id/volume_snapshots
//
// Lists all volume_snapshots.
// Optional parameters:
// filter
// view
func (loc *VolumeSnapshotLocator) Index(options rsapi.APIParams) ([]*VolumeSnapshot, error) {
	var res []*VolumeSnapshot
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("VolumeSnapshot", "index")
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

// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id
// GET /api/clouds/:cloud_id/volume_snapshots/:id
//
// Displays information about a single volume_snapshot.
// Optional parameters:
// view
func (loc *VolumeSnapshotLocator) Show(options rsapi.APIParams) (*VolumeSnapshot, error) {
	var res *VolumeSnapshot
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("VolumeSnapshot", "show")
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

/******  VolumeType ******/

// A VolumeType describes the type of volume, particularly the size.
type VolumeType struct {
	Actions     []map[string]string `json:"actions,omitempty"`
	CreatedAt   *RubyTime           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Size        string              `json:"size,omitempty"`
	UpdatedAt   *RubyTime           `json:"updated_at,omitempty"`
}

// Locator returns a locator for the given resource
func (r *VolumeType) Locator(api *API) *VolumeTypeLocator {
	for _, l := range r.Links {
		if l["rel"] == "self" {
			return api.VolumeTypeLocator(l["href"])
		}
	}
	return nil
}

//===== Locator

// VolumeTypeLocator exposes the VolumeType resource actions.
type VolumeTypeLocator struct {
	Href
	api *API
}

// VolumeTypeLocator builds a locator from the given href.
func (api *API) VolumeTypeLocator(href string) *VolumeTypeLocator {
	return &VolumeTypeLocator{Href(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/volume_types
//
// Lists Volume Types.
// Optional parameters:
// filter
// view
func (loc *VolumeTypeLocator) Index(options rsapi.APIParams) ([]*VolumeType, error) {
	var res []*VolumeType
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("VolumeType", "index")
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

// GET /api/clouds/:cloud_id/volume_types/:id
//
// Displays information about a single Volume Type.
// Optional parameters:
// view
func (loc *VolumeTypeLocator) Show(options rsapi.APIParams) (*VolumeType, error) {
	var res *VolumeType
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("VolumeType", "show")
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

type AlertSpecParam struct {
	Condition      string `json:"condition,omitempty"`
	Description    string `json:"description,omitempty"`
	Duration       string `json:"duration,omitempty"`
	EscalationName string `json:"escalation_name,omitempty"`
	File           string `json:"file,omitempty"`
	Name           string `json:"name,omitempty"`
	SubjectHref    string `json:"subject_href,omitempty"`
	Threshold      string `json:"threshold,omitempty"`
	Variable       string `json:"variable,omitempty"`
	VoteTag        string `json:"vote_tag,omitempty"`
	VoteType       string `json:"vote_type,omitempty"`
}

type AlertSpecParam2 struct {
	Condition      string `json:"condition,omitempty"`
	Description    string `json:"description,omitempty"`
	Duration       string `json:"duration,omitempty"`
	EscalationName string `json:"escalation_name,omitempty"`
	File           string `json:"file,omitempty"`
	Name           string `json:"name,omitempty"`
	Threshold      string `json:"threshold,omitempty"`
	Variable       string `json:"variable,omitempty"`
	VoteTag        string `json:"vote_tag,omitempty"`
	VoteType       string `json:"vote_type,omitempty"`
}

type AlertSpecificParams struct {
	DecisionThreshold  string `json:"decision_threshold,omitempty"`
	VotersTagPredicate string `json:"voters_tag_predicate,omitempty"`
}

type AllowedInstanceHrefs struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

type AssetPaths struct {
	Cookbooks []string `json:"cookbooks,omitempty"`
}

type AuditEntryParam struct {
	AuditeeHref string `json:"auditee_href,omitempty"`
	Detail      string `json:"detail,omitempty"`
	Summary     string `json:"summary,omitempty"`
}

type AuditEntryParam2 struct {
	Offset  int    `json:"offset,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type BackupParam struct {
	Description           string   `json:"description,omitempty"`
	FromMaster            string   `json:"from_master,omitempty"`
	Lineage               string   `json:"lineage,omitempty"`
	Name                  string   `json:"name,omitempty"`
	VolumeAttachmentHrefs []string `json:"volume_attachment_hrefs,omitempty"`
}

type BackupParam2 struct {
	Description    string `json:"description,omitempty"`
	Iops           string `json:"iops,omitempty"`
	Name           string `json:"name,omitempty"`
	Size           string `json:"size,omitempty"`
	VolumeTypeHref string `json:"volume_type_href,omitempty"`
}

type BackupParam3 struct {
	Committed string `json:"committed,omitempty"`
}

type Bounds struct {
	MaxCount string `json:"max_count,omitempty"`
	MinCount string `json:"min_count,omitempty"`
}

type ChildAccountParam struct {
	ClusterHref string `json:"cluster_href,omitempty"`
	Name        string `json:"name,omitempty"`
}

type ChildAccountParam2 struct {
	Name string `json:"name,omitempty"`
}

type CloudAccountParam struct {
	CloudHref string                 `json:"cloud_href,omitempty"`
	Creds     map[string]interface{} `json:"creds,omitempty"`
	Token     string                 `json:"token,omitempty"`
}

type CloudSpecificAttributes struct {
	AdminUsername                    string `json:"admin_username,omitempty"`
	AutomaticInstanceStoreMapping    string `json:"automatic_instance_store_mapping,omitempty"`
	AvailabilitySet                  string `json:"availability_set,omitempty"`
	CreateBootVolume                 string `json:"create_boot_volume,omitempty"`
	CreateDefaultPortForwardingRules string `json:"create_default_port_forwarding_rules,omitempty"`
	DeleteBootVolume                 string `json:"delete_boot_volume,omitempty"`
	DiskGb                           int    `json:"disk_gb,omitempty"`
	EbsOptimized                     string `json:"ebs_optimized,omitempty"`
	IamInstanceProfile               string `json:"iam_instance_profile,omitempty"`
	KeepAliveId                      string `json:"keep_alive_id,omitempty"`
	KeepAliveUrl                     string `json:"keep_alive_url,omitempty"`
	LocalSsdCount                    string `json:"local_ssd_count,omitempty"`
	LocalSsdInterface                string `json:"local_ssd_interface,omitempty"`
	MaxSpotPrice                     string `json:"max_spot_price,omitempty"`
	MemoryMb                         int    `json:"memory_mb,omitempty"`
	Metadata                         string `json:"metadata,omitempty"`
	NumCores                         int    `json:"num_cores,omitempty"`
	PlacementTenancy                 string `json:"placement_tenancy,omitempty"`
	Preemptible                      string `json:"preemptible,omitempty"`
	PricingType                      string `json:"pricing_type,omitempty"`
	RootVolumePerformance            string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                   string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid                string `json:"root_volume_type_uid,omitempty"`
	ServiceAccount                   string `json:"service_account,omitempty"`
}

type CloudSpecificAttributes2 struct {
	AdminUsername                    string `json:"admin_username,omitempty"`
	AutomaticInstanceStoreMapping    string `json:"automatic_instance_store_mapping,omitempty"`
	AvailabilitySet                  string `json:"availability_set,omitempty"`
	CreateBootVolume                 string `json:"create_boot_volume,omitempty"`
	CreateDefaultPortForwardingRules string `json:"create_default_port_forwarding_rules,omitempty"`
	DeleteBootVolume                 string `json:"delete_boot_volume,omitempty"`
	DiskGb                           int    `json:"disk_gb,omitempty"`
	IamInstanceProfile               string `json:"iam_instance_profile,omitempty"`
	KeepAliveId                      string `json:"keep_alive_id,omitempty"`
	KeepAliveUrl                     string `json:"keep_alive_url,omitempty"`
	LocalSsdCount                    string `json:"local_ssd_count,omitempty"`
	LocalSsdInterface                string `json:"local_ssd_interface,omitempty"`
	MaxSpotPrice                     string `json:"max_spot_price,omitempty"`
	MemoryMb                         int    `json:"memory_mb,omitempty"`
	Metadata                         string `json:"metadata,omitempty"`
	NumCores                         int    `json:"num_cores,omitempty"`
	PlacementTenancy                 string `json:"placement_tenancy,omitempty"`
	Preemptible                      string `json:"preemptible,omitempty"`
	PricingType                      string `json:"pricing_type,omitempty"`
	RootVolumePerformance            string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                   string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid                string `json:"root_volume_type_uid,omitempty"`
	ServiceAccount                   string `json:"service_account,omitempty"`
}

type CloudSpecificAttributes3 struct {
	AccountType string `json:"account_type,omitempty"`
}

type CloudSpecificAttributes4 struct {
	InstanceTags []string `json:"instance_tags,omitempty"`
	Priority     int      `json:"priority,omitempty"`
}

type CloudSpecificAttributes5 struct {
	AdminUsername                    string `json:"admin_username,omitempty"`
	AutomaticInstanceStoreMapping    string `json:"automatic_instance_store_mapping,omitempty"`
	AvailabilitySet                  string `json:"availability_set,omitempty"`
	CreateBootVolume                 string `json:"create_boot_volume,omitempty"`
	CreateDefaultPortForwardingRules string `json:"create_default_port_forwarding_rules,omitempty"`
	DeleteBootVolume                 string `json:"delete_boot_volume,omitempty"`
	DiskGb                           int    `json:"disk_gb,omitempty"`
	IamInstanceProfile               string `json:"iam_instance_profile,omitempty"`
	KeepAliveId                      string `json:"keep_alive_id,omitempty"`
	KeepAliveUrl                     string `json:"keep_alive_url,omitempty"`
	LocalSsdCount                    string `json:"local_ssd_count,omitempty"`
	LocalSsdInterface                string `json:"local_ssd_interface,omitempty"`
	MaxSpotPrice                     string `json:"max_spot_price,omitempty"`
	MemoryMb                         int    `json:"memory_mb,omitempty"`
	Metadata                         string `json:"metadata,omitempty"`
	NumCores                         int    `json:"num_cores,omitempty"`
	PlacementTenancy                 string `json:"placement_tenancy,omitempty"`
	Preemptible                      string `json:"preemptible,omitempty"`
	PricingType                      string `json:"pricing_type,omitempty"`
	RootVolumePerformance            string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                   string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid                string `json:"root_volume_type_uid,omitempty"`
	ServiceAccount                   string `json:"service_account,omitempty"`
}

type CloudSpecificAttributes6 struct {
	AdminUsername                    string `json:"admin_username,omitempty"`
	AutomaticInstanceStoreMapping    string `json:"automatic_instance_store_mapping,omitempty"`
	AvailabilitySet                  string `json:"availability_set,omitempty"`
	CreateBootVolume                 string `json:"create_boot_volume,omitempty"`
	CreateDefaultPortForwardingRules string `json:"create_default_port_forwarding_rules,omitempty"`
	DeleteBootVolume                 string `json:"delete_boot_volume,omitempty"`
	DiskGb                           int    `json:"disk_gb,omitempty"`
	IamInstanceProfile               string `json:"iam_instance_profile,omitempty"`
	KeepAliveId                      string `json:"keep_alive_id,omitempty"`
	KeepAliveUrl                     string `json:"keep_alive_url,omitempty"`
	LocalSsdCount                    string `json:"local_ssd_count,omitempty"`
	LocalSsdInterface                string `json:"local_ssd_interface,omitempty"`
	MaxSpotPrice                     string `json:"max_spot_price,omitempty"`
	MemoryMb                         int    `json:"memory_mb,omitempty"`
	Metadata                         string `json:"metadata,omitempty"`
	NumCores                         int    `json:"num_cores,omitempty"`
	PlacementTenancy                 string `json:"placement_tenancy,omitempty"`
	Preemptible                      string `json:"preemptible,omitempty"`
	PricingType                      string `json:"pricing_type,omitempty"`
	RootVolumePerformance            string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                   string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid                string `json:"root_volume_type_uid,omitempty"`
	ServiceAccount                   string `json:"service_account,omitempty"`
}

type CloudSpecificAttributes7 struct {
	CreateBootVolume string `json:"create_boot_volume,omitempty"`
	DeleteBootVolume string `json:"delete_boot_volume,omitempty"`
}

type CookbookAttachmentParam struct {
	CookbookHref       string `json:"cookbook_href,omitempty"`
	ServerTemplateHref string `json:"server_template_href,omitempty"`
}

type CookbookAttachments struct {
	CookbookHrefs      []string `json:"cookbook_hrefs,omitempty"`
	ServerTemplateHref string   `json:"server_template_href,omitempty"`
}

type CookbookAttachments2 struct {
	CookbookAttachmentHrefs []string `json:"cookbook_attachment_hrefs,omitempty"`
}

type CredentialParam struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
}

type Credentials struct {
	Password string `json:"password,omitempty"`
	SshKey   string `json:"ssh_key,omitempty"`
	Username string `json:"username,omitempty"`
}

type DatacenterPolicy struct {
	DatacenterHref string `json:"datacenter_href,omitempty"`
	Max            string `json:"max,omitempty"`
	Weight         string `json:"weight,omitempty"`
}

type DeploymentParam struct {
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	ServerTagScope string `json:"server_tag_scope,omitempty"`
}

type DeploymentParam2 struct {
	Description       string `json:"description,omitempty"`
	Name              string `json:"name,omitempty"`
	ResourceGroupHref string `json:"resource_group_href,omitempty"`
	ServerTagScope    string `json:"server_tag_scope,omitempty"`
}

type DeploymentParam3 struct {
	Description       string `json:"description,omitempty"`
	Name              string `json:"name,omitempty"`
	ResourceGroupHref string `json:"resource_group_href,omitempty"`
	ServerTagScope    string `json:"server_tag_scope,omitempty"`
}

type Descriptions struct {
	Long  string `json:"long,omitempty"`
	Notes string `json:"notes,omitempty"`
	Short string `json:"short,omitempty"`
}

type ElasticityParams struct {
	AlertSpecificParams *AlertSpecificParams `json:"alert_specific_params,omitempty"`
	Bounds              *Bounds              `json:"bounds,omitempty"`
	Pacing              *Pacing              `json:"pacing,omitempty"`
	QueueSpecificParams *QueueSpecificParams `json:"queue_specific_params,omitempty"`
	Schedule            []*Schedule          `json:"schedule,omitempty"`
}

type InstanceParam struct {
	AssociatePublicIpAddress string                   `json:"associate_public_ip_address,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                   `json:"datacenter_href,omitempty"`
	DeploymentHref           string                   `json:"deployment_href,omitempty"`
	ImageHref                string                   `json:"image_href,omitempty"`
	InstanceTypeHref         string                   `json:"instance_type_href,omitempty"`
	IpForwardingEnabled      string                   `json:"ip_forwarding_enabled,omitempty"`
	KernelImageHref          string                   `json:"kernel_image_href,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	PlacementGroupHref       string                   `json:"placement_group_href,omitempty"`
	RamdiskImageHref         string                   `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                 `json:"security_group_hrefs,omitempty"`
	SshKeyHref               string                   `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                 `json:"subnet_hrefs,omitempty"`
	UserData                 string                   `json:"user_data,omitempty"`
}

type InstanceParam2 struct {
	AssociatePublicIpAddress string                    `json:"associate_public_ip_address,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes2 `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                    `json:"datacenter_href,omitempty"`
	DeploymentHref           string                    `json:"deployment_href,omitempty"`
	ImageHref                string                    `json:"image_href,omitempty"`
	InstanceTypeHref         string                    `json:"instance_type_href,omitempty"`
	IpForwardingEnabled      string                    `json:"ip_forwarding_enabled,omitempty"`
	KernelImageHref          string                    `json:"kernel_image_href,omitempty"`
	MultiCloudImageHref      string                    `json:"multi_cloud_image_href,omitempty"`
	Name                     string                    `json:"name,omitempty"`
	PrivateIpAddress         string                    `json:"private_ip_address,omitempty"`
	RamdiskImageHref         string                    `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                  `json:"security_group_hrefs,omitempty"`
	ServerTemplateHref       string                    `json:"server_template_href,omitempty"`
	SshKeyHref               string                    `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                  `json:"subnet_hrefs,omitempty"`
	UserData                 string                    `json:"user_data,omitempty"`
}

type InstanceParam3 struct {
	AssociatePublicIpAddress string                    `json:"associate_public_ip_address,omitempty"`
	CloudHref                string                    `json:"cloud_href,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes5 `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                    `json:"datacenter_href,omitempty"`
	ImageHref                string                    `json:"image_href,omitempty"`
	Inputs                   map[string]interface{}    `json:"inputs,omitempty"`
	InstanceTypeHref         string                    `json:"instance_type_href,omitempty"`
	IpForwardingEnabled      string                    `json:"ip_forwarding_enabled,omitempty"`
	KernelImageHref          string                    `json:"kernel_image_href,omitempty"`
	MultiCloudImageHref      string                    `json:"multi_cloud_image_href,omitempty"`
	PlacementGroupHref       string                    `json:"placement_group_href,omitempty"`
	RamdiskImageHref         string                    `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                  `json:"security_group_hrefs,omitempty"`
	ServerTemplateHref       string                    `json:"server_template_href,omitempty"`
	SshKeyHref               string                    `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                  `json:"subnet_hrefs,omitempty"`
	UserData                 string                    `json:"user_data,omitempty"`
}

type InstanceParam4 struct {
	AssociatePublicIpAddress string                    `json:"associate_public_ip_address,omitempty"`
	CloudHref                string                    `json:"cloud_href,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes6 `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                    `json:"datacenter_href,omitempty"`
	ImageHref                string                    `json:"image_href,omitempty"`
	Inputs                   map[string]interface{}    `json:"inputs,omitempty"`
	InstanceTypeHref         string                    `json:"instance_type_href,omitempty"`
	IpForwardingEnabled      string                    `json:"ip_forwarding_enabled,omitempty"`
	KernelImageHref          string                    `json:"kernel_image_href,omitempty"`
	MultiCloudImageHref      string                    `json:"multi_cloud_image_href,omitempty"`
	PlacementGroupHref       string                    `json:"placement_group_href,omitempty"`
	PrivateIpAddress         string                    `json:"private_ip_address,omitempty"`
	RamdiskImageHref         string                    `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                  `json:"security_group_hrefs,omitempty"`
	ServerTemplateHref       string                    `json:"server_template_href,omitempty"`
	SshKeyHref               string                    `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                  `json:"subnet_hrefs,omitempty"`
	UserData                 string                    `json:"user_data,omitempty"`
}

type InstanceParam5 struct {
	CloudSpecificAttributes *CloudSpecificAttributes7 `json:"cloud_specific_attributes,omitempty"`
}

type InstanceParam6 struct {
	Href                string                 `json:"href,omitempty"`
	Inputs              map[string]interface{} `json:"inputs,omitempty"`
	MultiCloudImageHref string                 `json:"multi_cloud_image_href,omitempty"`
	ServerTemplateHref  string                 `json:"server_template_href,omitempty"`
}

type IpAddressBindingParam struct {
	InstanceHref        string `json:"instance_href,omitempty"`
	PrivatePort         string `json:"private_port,omitempty"`
	Protocol            string `json:"protocol,omitempty"`
	PublicIpAddressHref string `json:"public_ip_address_href,omitempty"`
	PublicPort          string `json:"public_port,omitempty"`
	ServerHref          string `json:"server_href,omitempty"`
}

type IpAddressParam struct {
	DeploymentHref string `json:"deployment_href,omitempty"`
	Domain         string `json:"domain,omitempty"`
	Name           string `json:"name,omitempty"`
	NetworkHref    string `json:"network_href,omitempty"`
}

type IpAddressParam2 struct {
	DeploymentHref string `json:"deployment_href,omitempty"`
	Name           string `json:"name,omitempty"`
}

type ItemAge struct {
	Algorithm string `json:"algorithm,omitempty"`
	MaxAge    string `json:"max_age,omitempty"`
	Regexp    string `json:"regexp,omitempty"`
}

type MultiCloudImageMatcherParam struct {
	ImageHref string `json:"image_href,omitempty"`
	UserData  string `json:"user_data,omitempty"`
}

type MultiCloudImageParam struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type MultiCloudImageSettingParam struct {
	CloudHref        string `json:"cloud_href,omitempty"`
	ImageHref        string `json:"image_href,omitempty"`
	InstanceTypeHref string `json:"instance_type_href,omitempty"`
	KernelImageHref  string `json:"kernel_image_href,omitempty"`
	RamdiskImageHref string `json:"ramdisk_image_href,omitempty"`
	UserData         string `json:"user_data,omitempty"`
}

type NetworkGatewayParam struct {
	CloudHref   string `json:"cloud_href,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Type_       string `json:"type,omitempty"`
}

type NetworkGatewayParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	NetworkHref string `json:"network_href,omitempty"`
}

type NetworkOptionGroupAttachmentParam struct {
	NetworkHref            string `json:"network_href,omitempty"`
	NetworkOptionGroupHref string `json:"network_option_group_href,omitempty"`
}

type NetworkOptionGroupAttachmentParam2 struct {
	NetworkOptionGroupHref string `json:"network_option_group_href,omitempty"`
}

type NetworkOptionGroupParam struct {
	CloudHref   string                 `json:"cloud_href,omitempty"`
	Description string                 `json:"description,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Options_    map[string]interface{} `json:"options,omitempty"`
	Type_       string                 `json:"type,omitempty"`
}

type NetworkOptionGroupParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type NetworkParam struct {
	CidrBlock       string `json:"cidr_block,omitempty"`
	CloudHref       string `json:"cloud_href,omitempty"`
	DeploymentHref  string `json:"deployment_href,omitempty"`
	Description     string `json:"description,omitempty"`
	InstanceTenancy string `json:"instance_tenancy,omitempty"`
	Name            string `json:"name,omitempty"`
}

type NetworkParam2 struct {
	DeploymentHref string `json:"deployment_href,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	RouteTableHref string `json:"route_table_href,omitempty"`
}

type Pacing struct {
	ResizeCalmTime string `json:"resize_calm_time,omitempty"`
	ResizeDownBy   string `json:"resize_down_by,omitempty"`
	ResizeUpBy     string `json:"resize_up_by,omitempty"`
}

type PermissionParam struct {
	RoleTitle string `json:"role_title,omitempty"`
	UserHref  string `json:"user_href,omitempty"`
}

type PlacementGroupParam struct {
	CloudHref               string                    `json:"cloud_href,omitempty"`
	CloudSpecificAttributes *CloudSpecificAttributes3 `json:"cloud_specific_attributes,omitempty"`
	DeploymentHref          string                    `json:"deployment_href,omitempty"`
	Description             string                    `json:"description,omitempty"`
	Name                    string                    `json:"name,omitempty"`
}

type PreferenceParam struct {
	Contents string `json:"contents,omitempty"`
}

type ProtocolDetails struct {
	EndPort   string `json:"end_port,omitempty"`
	IcmpCode  string `json:"icmp_code,omitempty"`
	IcmpType  string `json:"icmp_type,omitempty"`
	StartPort string `json:"start_port,omitempty"`
}

type QueueSize struct {
	ItemsPerInstance string `json:"items_per_instance,omitempty"`
}

type QueueSpecificParams struct {
	CollectAuditEntries string     `json:"collect_audit_entries,omitempty"`
	ItemAge             *ItemAge   `json:"item_age,omitempty"`
	QueueSize           *QueueSize `json:"queue_size,omitempty"`
}

type RecurringVolumeAttachmentParam struct {
	Device         string                 `json:"device,omitempty"`
	RunnableHref   string                 `json:"runnable_href,omitempty"`
	Settings       map[string]interface{} `json:"settings,omitempty"`
	StorageHref    string                 `json:"storage_href,omitempty"`
	VolumeTypeHref string                 `json:"volume_type_href,omitempty"`
}

type RepositoryParam struct {
	AssetPaths      *AssetPaths  `json:"asset_paths,omitempty"`
	AutoImport      string       `json:"auto_import,omitempty"`
	CommitReference string       `json:"commit_reference,omitempty"`
	Credentials     *Credentials `json:"credentials,omitempty"`
	Description     string       `json:"description,omitempty"`
	Name            string       `json:"name,omitempty"`
	Source          string       `json:"source,omitempty"`
	SourceType      string       `json:"source_type,omitempty"`
}

type RepositoryParam2 struct {
	AssetPaths      *AssetPaths  `json:"asset_paths,omitempty"`
	CommitReference string       `json:"commit_reference,omitempty"`
	Credentials     *Credentials `json:"credentials,omitempty"`
	Description     string       `json:"description,omitempty"`
	Name            string       `json:"name,omitempty"`
	Source          string       `json:"source,omitempty"`
	SourceType      string       `json:"source_type,omitempty"`
}

type ResourceGroupParam struct {
	CloudHref      string `json:"cloud_href,omitempty"`
	DeploymentHref string `json:"deployment_href,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
}

type ResourceGroupParam2 struct {
	DeploymentHref string `json:"deployment_href,omitempty"`
	Description    string `json:"description,omitempty"`
}

type RightScriptAttachmentParam struct {
	Content  *rsapi.FileUpload `json:"content,omitempty"`
	Filename string            `json:"filename,omitempty"`
}

type RightScriptParam struct {
	CommitMessage string `json:"commit_message,omitempty"`
}

type RightScriptParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Packages    string `json:"packages,omitempty"`
	Source      string `json:"source,omitempty"`
}

type RightScriptParam3 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Packages    string `json:"packages,omitempty"`
	Source      string `json:"source,omitempty"`
}

type RouteParam struct {
	CloudSpecificAttributes *CloudSpecificAttributes4 `json:"cloud_specific_attributes,omitempty"`
	Description             string                    `json:"description,omitempty"`
	DestinationCidrBlock    string                    `json:"destination_cidr_block,omitempty"`
	NextHopHref             string                    `json:"next_hop_href,omitempty"`
	NextHopIp               string                    `json:"next_hop_ip,omitempty"`
	NextHopType             string                    `json:"next_hop_type,omitempty"`
	NextHopUrl              string                    `json:"next_hop_url,omitempty"`
	RouteTableHref          string                    `json:"route_table_href,omitempty"`
}

type RouteParam2 struct {
	Description          string `json:"description,omitempty"`
	DestinationCidrBlock string `json:"destination_cidr_block,omitempty"`
	NextHopHref          string `json:"next_hop_href,omitempty"`
	NextHopIp            string `json:"next_hop_ip,omitempty"`
	NextHopType          string `json:"next_hop_type,omitempty"`
}

type RouteTableParam struct {
	CloudHref   string `json:"cloud_href,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	NetworkHref string `json:"network_href,omitempty"`
}

type RouteTableParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type RunnableBindingParam struct {
	Position        string `json:"position,omitempty"`
	Recipe          string `json:"recipe,omitempty"`
	RightScriptHref string `json:"right_script_href,omitempty"`
	Sequence        string `json:"sequence,omitempty"`
}

type RunnableBindings struct {
	Id              string `json:"id,omitempty"`
	Position        string `json:"position,omitempty"`
	Recipe          string `json:"recipe,omitempty"`
	RightScriptHref string `json:"right_script_href,omitempty"`
	Sequence        string `json:"sequence,omitempty"`
}

type Schedule struct {
	Day      string `json:"day,omitempty"`
	MaxCount string `json:"max_count,omitempty"`
	MinCount string `json:"min_count,omitempty"`
	Time     string `json:"time,omitempty"`
}

type SecurityGroupParam struct {
	DeploymentHref string `json:"deployment_href,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	NetworkHref    string `json:"network_href,omitempty"`
}

type SecurityGroupRuleParam struct {
	Action            string           `json:"action,omitempty"`
	CidrIps           string           `json:"cidr_ips,omitempty"`
	Direction         string           `json:"direction,omitempty"`
	GroupName         string           `json:"group_name,omitempty"`
	GroupOwner        string           `json:"group_owner,omitempty"`
	Priority          int              `json:"priority,omitempty"`
	Protocol          string           `json:"protocol,omitempty"`
	ProtocolDetails   *ProtocolDetails `json:"protocol_details,omitempty"`
	SecurityGroupHref string           `json:"security_group_href,omitempty"`
	SourceType        string           `json:"source_type,omitempty"`
}

type SecurityGroupRuleParam2 struct {
	Description string `json:"description,omitempty"`
}

type ServerArrayParam struct {
	ArrayType        string              `json:"array_type,omitempty"`
	DatacenterPolicy []*DatacenterPolicy `json:"datacenter_policy,omitempty"`
	DeploymentHref   string              `json:"deployment_href,omitempty"`
	Description      string              `json:"description,omitempty"`
	ElasticityParams *ElasticityParams   `json:"elasticity_params,omitempty"`
	Instance         *InstanceParam3     `json:"instance,omitempty"`
	Name             string              `json:"name,omitempty"`
	Optimized        string              `json:"optimized,omitempty"`
	State            string              `json:"state,omitempty"`
}

type ServerArrayParam2 struct {
	ArrayType        string              `json:"array_type,omitempty"`
	DatacenterPolicy []*DatacenterPolicy `json:"datacenter_policy,omitempty"`
	DeploymentHref   string              `json:"deployment_href,omitempty"`
	Description      string              `json:"description,omitempty"`
	ElasticityParams *ElasticityParams   `json:"elasticity_params,omitempty"`
	Name             string              `json:"name,omitempty"`
	Optimized        string              `json:"optimized,omitempty"`
	State            string              `json:"state,omitempty"`
}

type ServerParam struct {
	DeploymentHref string          `json:"deployment_href,omitempty"`
	Description    string          `json:"description,omitempty"`
	Instance       *InstanceParam4 `json:"instance,omitempty"`
	Name           string          `json:"name,omitempty"`
	Optimized      string          `json:"optimized,omitempty"`
}

type ServerParam2 struct {
	AutomaticInstanceStoreMapping string          `json:"automatic_instance_store_mapping,omitempty"`
	Description                   string          `json:"description,omitempty"`
	Instance                      *InstanceParam5 `json:"instance,omitempty"`
	Name                          string          `json:"name,omitempty"`
	Optimized                     string          `json:"optimized,omitempty"`
	RootVolumeSize                string          `json:"root_volume_size,omitempty"`
}

type ServerParam3 struct {
	DeploymentHref string          `json:"deployment_href,omitempty"`
	Description    string          `json:"description,omitempty"`
	Instance       *InstanceParam6 `json:"instance,omitempty"`
	Name           string          `json:"name,omitempty"`
}

type ServerTemplateMultiCloudImageParam struct {
	MultiCloudImageHref string `json:"multi_cloud_image_href,omitempty"`
	ServerTemplateHref  string `json:"server_template_href,omitempty"`
}

type ServerTemplateParam struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Settings struct {
	DeleteOnTermination string `json:"delete_on_termination,omitempty"`
}

type SshKeyParam struct {
	Name string `json:"name,omitempty"`
}

type SubnetParam struct {
	CidrBlock      string `json:"cidr_block,omitempty"`
	DatacenterHref string `json:"datacenter_href,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	NetworkHref    string `json:"network_href,omitempty"`
}

type SubnetParam2 struct {
	CidrBlock      string `json:"cidr_block,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	RouteTableHref string `json:"route_table_href,omitempty"`
}

type UserParam struct {
	Company              string `json:"company,omitempty"`
	Email                string `json:"email,omitempty"`
	FirstName            string `json:"first_name,omitempty"`
	IdentityProviderHref string `json:"identity_provider_href,omitempty"`
	LastName             string `json:"last_name,omitempty"`
	Password             string `json:"password,omitempty"`
	Phone                string `json:"phone,omitempty"`
	PrincipalUid         string `json:"principal_uid,omitempty"`
	TimezoneName         string `json:"timezone_name,omitempty"`
}

type UserParam2 struct {
	Company              string `json:"company,omitempty"`
	CurrentEmail         string `json:"current_email,omitempty"`
	CurrentPassword      string `json:"current_password,omitempty"`
	FirstName            string `json:"first_name,omitempty"`
	IdentityProviderHref string `json:"identity_provider_href,omitempty"`
	LastName             string `json:"last_name,omitempty"`
	LoginName            string `json:"login_name,omitempty"`
	NewEmail             string `json:"new_email,omitempty"`
	NewPassword          string `json:"new_password,omitempty"`
	Phone                string `json:"phone,omitempty"`
	PrincipalUid         string `json:"principal_uid,omitempty"`
	TimezoneName         string `json:"timezone_name,omitempty"`
}

type VolumeAttachmentParam struct {
	Device       string    `json:"device,omitempty"`
	InstanceHref string    `json:"instance_href,omitempty"`
	ServerHref   string    `json:"server_href,omitempty"`
	Settings     *Settings `json:"settings,omitempty"`
	VolumeHref   string    `json:"volume_href,omitempty"`
}

type VolumeParam struct {
	DatacenterHref           string `json:"datacenter_href,omitempty"`
	DeploymentHref           string `json:"deployment_href,omitempty"`
	Description              string `json:"description,omitempty"`
	Encrypted                string `json:"encrypted,omitempty"`
	ImageHref                string `json:"image_href,omitempty"`
	Iops                     string `json:"iops,omitempty"`
	Name                     string `json:"name,omitempty"`
	ParentVolumeSnapshotHref string `json:"parent_volume_snapshot_href,omitempty"`
	PlacementGroupHref       string `json:"placement_group_href,omitempty"`
	Size                     string `json:"size,omitempty"`
	VolumeTypeHref           string `json:"volume_type_href,omitempty"`
}

type VolumeParam2 struct {
	AllowedInstanceHrefs *AllowedInstanceHrefs `json:"allowed_instance_hrefs,omitempty"`
	Name                 string                `json:"name,omitempty"`
}

type VolumeSnapshotCopy struct {
	CloudHref   string `json:"cloud_href,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type VolumeSnapshotParam struct {
	DeploymentHref   string `json:"deployment_href,omitempty"`
	Description      string `json:"description,omitempty"`
	Name             string `json:"name,omitempty"`
	ParentVolumeHref string `json:"parent_volume_href,omitempty"`
}
