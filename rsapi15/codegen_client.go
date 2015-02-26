//************************************************************************//
//                     RightScale API client
//
// Generated
// Feb 25, 2015 at 5:09pm (PST)
// Command:
// $ ./api15gen -metadata=../../rsapi15 -output=../../rsapi15
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// Helper function that merges optional parameters into payload
func mergeOptionals(params, options rsapi.ApiParams) rsapi.ApiParams {
	for name, value := range options {
		params[name] = value
	}
	return params
}

// Url resolver produces an action URL and HTTP method from its name and a given resource href.
// The algorithm consists of first extracting the variables from the href and then substituing them
// in the action path. If there are more than one action paths then the algorithm picks the one that
// can substitute the most variables.
type UrlResolver string

func (r *UrlResolver) Url(rName, aName string) (*metadata.ActionPath, error) {
	var res, ok = api_metadata[rName]
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
	var vars, err = res.ExtractVariables(string(*r))
	if err != nil {
		return nil, err
	}
	return action.Url(vars)
}

/******  Account ******/

type Account struct {
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Permissions []Permission        `json:"permissions,omitempty"`
	Products    []string            `json:"products,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// Account resource locator, exposes resource actions.
type AccountLocator struct {
	UrlResolver
	api *Api15
}

// Account resource locator factory
func (api *Api15) AccountLocator(href string) *AccountLocator {
	return &AccountLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/accounts/:id
// Show information about a single Account.
func (loc *AccountLocator) Show() (*Account, error) {
	var res *Account
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Accounts", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  AccountGroup ******/

// An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.
type AccountGroup struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// AccountGroup resource locator, exposes resource actions.
type AccountGroupLocator struct {
	UrlResolver
	api *Api15
}

// AccountGroup resource locator factory
func (api *Api15) AccountGroupLocator(href string) *AccountGroupLocator {
	return &AccountGroupLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/account_groups
// Lists the AccountGroups owned by this Account.
// -- Optional parameters:
// filter
// view
func (loc *AccountGroupLocator) Index(options rsapi.ApiParams) ([]*AccountGroup, error) {
	var res []*AccountGroup
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("AccountGroups", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/account_groups/:id
// Show information about a single AccountGroup.
// -- Optional parameters:
// view
func (loc *AccountGroupLocator) Show(options rsapi.ApiParams) (*AccountGroup, error) {
	var res *AccountGroup
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("AccountGroups", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Alert ******/

// An Alert represents an AlertSpec bound to a running Instance.
type Alert struct {
	Actions       []string            `json:"actions,omitempty"`
	CreatedAt     time.Time           `json:"created_at,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	QuenchedUntil time.Time           `json:"quenched_until,omitempty"`
	Status        string              `json:"status,omitempty"`
	TriggeredAt   time.Time           `json:"triggered_at,omitempty"`
	UpdatedAt     time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// Alert resource locator, exposes resource actions.
type AlertLocator struct {
	UrlResolver
	api *Api15
}

// Alert resource locator factory
func (api *Api15) AlertLocator(href string) *AlertLocator {
	return &AlertLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances/:instance_id/alerts/:id/disable
// POST /api/servers/:server_id/alerts/:id/disable
// POST /api/server_arrays/:server_array_id/alerts/:id/disable
// POST /api/deployments/:deployment_id/alerts/:id/disable
// POST /api/alerts/:id/disable
// Disables the Alert indefinitely. Idempotent.
func (loc *AlertLocator) Disable() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Alerts", "disable")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:instance_id/alerts/:id/enable
// POST /api/servers/:server_id/alerts/:id/enable
// POST /api/server_arrays/:server_array_id/alerts/:id/enable
// POST /api/deployments/:deployment_id/alerts/:id/enable
// POST /api/alerts/:id/enable
// Enables the Alert indefinitely. Idempotent.
func (loc *AlertLocator) Enable() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Alerts", "enable")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/instances/:instance_id/alerts
// GET /api/servers/:server_id/alerts
// GET /api/server_arrays/:server_array_id/alerts
// GET /api/deployments/:deployment_id/alerts
// GET /api/alerts
// Lists all Alerts.
// -- Optional parameters:
// filter
// view
func (loc *AlertLocator) Index(options rsapi.ApiParams) ([]*Alert, error) {
	var res []*Alert
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Alerts", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/clouds/:cloud_id/instances/:instance_id/alerts/:id/quench
// POST /api/servers/:server_id/alerts/:id/quench
// POST /api/server_arrays/:server_array_id/alerts/:id/quench
// POST /api/deployments/:deployment_id/alerts/:id/quench
// POST /api/alerts/:id/quench
// Suppresses the Alert from being triggered for a given time period. Idempotent.
// duration: The time period in seconds to suppress Alert from being triggered.
func (loc *AlertLocator) Quench(duration string) (string, error) {
	var res string
	if duration == "" {
		return res, fmt.Errorf("duration is required")
	}
	var params = rsapi.ApiParams{
		"duration": duration,
	}
	var uri, err = loc.Url("Alerts", "quench")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/instances/:instance_id/alerts/:id
// GET /api/servers/:server_id/alerts/:id
// GET /api/server_arrays/:server_array_id/alerts/:id
// GET /api/deployments/:deployment_id/alerts/:id
// GET /api/alerts/:id
// Shows the attributes of a specified Alert.
// -- Optional parameters:
// view
func (loc *AlertLocator) Show(options rsapi.ApiParams) (*Alert, error) {
	var res *Alert
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Alerts", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  AlertSpec ******/

// An AlertSpec defines the conditions under which an Alert is triggered and escalated.
// Condition sentence: if &lt;file&gt;.&lt;variable&gt; &lt;condition&gt; '&lt;threshold&gt;' for &lt;duration&gt; min then escalate to '&lt;escalation_name&gt;'.
type AlertSpec struct {
	Actions        []string            `json:"actions,omitempty"`
	Condition      string              `json:"condition,omitempty"`
	CreatedAt      time.Time           `json:"created_at,omitempty"`
	Description    string              `json:"description,omitempty"`
	Duration       int                 `json:"duration,omitempty"`
	EscalationName string              `json:"escalation_name,omitempty"`
	File           string              `json:"file,omitempty"`
	Links          []map[string]string `json:"links,omitempty"`
	Name           string              `json:"name,omitempty"`
	Threshold      string              `json:"threshold,omitempty"`
	UpdatedAt      time.Time           `json:"updated_at,omitempty"`
	Variable       string              `json:"variable,omitempty"`
	VoteTag        string              `json:"vote_tag,omitempty"`
	VoteType       string              `json:"vote_type,omitempty"`
}

//===== Locator

// AlertSpec resource locator, exposes resource actions.
type AlertSpecLocator struct {
	UrlResolver
	api *Api15
}

// AlertSpec resource locator factory
func (api *Api15) AlertSpecLocator(href string) *AlertSpecLocator {
	return &AlertSpecLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/servers/:server_id/alert_specs
// POST /api/server_arrays/:server_array_id/alert_specs
// POST /api/server_templates/:server_template_id/alert_specs
// POST /api/alert_specs
// Creates a new AlertSpec with the given parameters.
func (loc *AlertSpecLocator) Create(alertSpec *AlertSpecParam) (*AlertSpecLocator, error) {
	var res *AlertSpecLocator
	if alertSpec == nil {
		return res, fmt.Errorf("alertSpec is required")
	}
	var params = rsapi.ApiParams{
		"alert_spec": alertSpec,
	}
	var uri, err = loc.Url("AlertSpecs", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &AlertSpecLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/servers/:server_id/alert_specs/:id
// DELETE /api/server_arrays/:server_array_id/alert_specs/:id
// DELETE /api/server_templates/:server_template_id/alert_specs/:id
// DELETE /api/alert_specs/:id
// Deletes a given AlertSpec.
func (loc *AlertSpecLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("AlertSpecs", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/servers/:server_id/alert_specs
// GET /api/server_arrays/:server_array_id/alert_specs
// GET /api/server_templates/:server_template_id/alert_specs
// GET /api/alert_specs
// No description provided for index.
// -- Optional parameters:
// filter
// view
// with_inherited: Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.
func (loc *AlertSpecLocator) Index(options rsapi.ApiParams) ([]*AlertSpec, error) {
	var res []*AlertSpec
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("AlertSpecs", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/servers/:server_id/alert_specs/:id
// GET /api/server_arrays/:server_array_id/alert_specs/:id
// GET /api/server_templates/:server_template_id/alert_specs/:id
// GET /api/alert_specs/:id
// No description provided for show.
// -- Optional parameters:
// view
func (loc *AlertSpecLocator) Show(options rsapi.ApiParams) (*AlertSpec, error) {
	var res *AlertSpec
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("AlertSpecs", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/servers/:server_id/alert_specs/:id
// PUT /api/server_arrays/:server_array_id/alert_specs/:id
// PUT /api/server_templates/:server_template_id/alert_specs/:id
// PUT /api/alert_specs/:id
// Updates an AlertSpec with the given parameters.
func (loc *AlertSpecLocator) Update(alertSpec *AlertSpecParam2) error {
	if alertSpec == nil {
		return fmt.Errorf("alertSpec is required")
	}
	var params = rsapi.ApiParams{
		"alert_spec": alertSpec,
	}
	var uri, err = loc.Url("AlertSpecs", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  AuditEntry ******/

// An Audit Entry can be used to track various activities of a resource.
type AuditEntry struct {
	Actions    []string            `json:"actions,omitempty"`
	DetailSize int                 `json:"detail_size,omitempty"`
	Links      []map[string]string `json:"links,omitempty"`
	Summary    string              `json:"summary,omitempty"`
	UpdatedAt  time.Time           `json:"updated_at,omitempty"`
	UserEmail  string              `json:"user_email,omitempty"`
}

//===== Locator

// AuditEntry resource locator, exposes resource actions.
type AuditEntryLocator struct {
	UrlResolver
	api *Api15
}

// AuditEntry resource locator factory
func (api *Api15) AuditEntryLocator(href string) *AuditEntryLocator {
	return &AuditEntryLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/audit_entries/:id/append
// Updates the summary and appends more details to a given AuditEntry. Each audit entry detail is stored
// as one chunk, the offset determines the location of that chunk within the overall audit entry details section.
// For example, if you create an AuditEntry and append "DEF" at offset 10, and later append
// "ABC" at offset 9, the overall audit entry details will be "ABCDEF". Use the \n character to
// separate details by new lines.
// -- Optional parameters:
// detail: The details to be appended to the audit entry record.
// notify: The event notification category. Defaults to 'None'.
// offset: The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.
// summary: The updated summary for the audit entry, maximum length is 255 characters.
func (loc *AuditEntryLocator) Append(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("AuditEntries", "append")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/audit_entries
// Creates a new AuditEntry with the given parameters.
// -- Optional parameters:
// notify: The event notification category. Defaults to 'None'.
// user_email: The email of the user (who created/triggered the audit entry). Only usable with instance role.
func (loc *AuditEntryLocator) Create(auditEntry *AuditEntryParam, options rsapi.ApiParams) (*AuditEntryLocator, error) {
	var res *AuditEntryLocator
	if auditEntry == nil {
		return res, fmt.Errorf("auditEntry is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"audit_entry": auditEntry,
	}, options)
	var uri, err = loc.Url("AuditEntries", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &AuditEntryLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/audit_entries/:id/detail
// shows the details of a given AuditEntry.
// Note that the media type of the response is simply text.
func (loc *AuditEntryLocator) Detail() (string, error) {
	var res string
	var params = map[string]interface{}{}
	var uri, err = loc.Url("AuditEntries", "detail")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/audit_entries
// Lists AuditEntries of the account. Due to the potentially large number of audit entries, a start and end date must
// be provided during an index call to limit the search. The format of the dates must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
// 2011/07/11 00:00:00 +0000.
// A maximum of 1000 records will be returned by each index call.
// Using the available filters, one can select or group which audit entries to retrieve.
// endDate: The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).
// limit: Limit the audit entries to this number. The limit should >= 1 and <= 1000
// startDate: The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000
// -- Optional parameters:
// filter
// view
func (loc *AuditEntryLocator) Index(endDate string, limit string, startDate string, options rsapi.ApiParams) ([]*AuditEntry, error) {
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
	var params = mergeOptionals(rsapi.ApiParams{
		"end_date":   endDate,
		"limit":      limit,
		"start_date": startDate,
	}, options)
	var uri, err = loc.Url("AuditEntries", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/audit_entries/:id
// Lists the attributes of a given audit entry.
// -- Optional parameters:
// view
func (loc *AuditEntryLocator) Show(options rsapi.ApiParams) (*AuditEntry, error) {
	var res *AuditEntry
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("AuditEntries", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/audit_entries/:id
// Updates the summary of a given AuditEntry.
// -- Optional parameters:
// notify: The event notification category. Defaults to 'None'.
func (loc *AuditEntryLocator) Update(auditEntry *AuditEntryParam2, options rsapi.ApiParams) error {
	if auditEntry == nil {
		return fmt.Errorf("auditEntry is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"audit_entry": auditEntry,
	}, options)
	var uri, err = loc.Url("AuditEntries", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Backup ******/

type Backup struct {
	Actions             []string            `json:"actions,omitempty"`
	Committed           bool                `json:"committed,omitempty"`
	Completed           bool                `json:"completed,omitempty"`
	CreatedAt           time.Time           `json:"created_at,omitempty"`
	Description         string              `json:"description,omitempty"`
	FromMaster          bool                `json:"from_master,omitempty"`
	Lineage             string              `json:"lineage,omitempty"`
	Links               []map[string]string `json:"links,omitempty"`
	Name                string              `json:"name,omitempty"`
	VolumeSnapshotCount int                 `json:"volume_snapshot_count,omitempty"`
	VolumeSnapshots     []VolumeSnapshot    `json:"volume_snapshots,omitempty"`
}

//===== Locator

// Backup resource locator, exposes resource actions.
type BackupLocator struct {
	UrlResolver
	api *Api15
}

// Backup resource locator factory
func (api *Api15) BackupLocator(href string) *BackupLocator {
	return &BackupLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/backups/cleanup
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
// keepLast: The number of backups that should be kept.
// lineage: The lineage of the backups that are to be cleaned-up.
// -- Optional parameters:
// cloud_href: Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.
// dailies: The number of daily backups(the latest one in each day) that should be kept.
// monthlies: The number of monthly backups(the latest one in each month) that should be kept.
// weeklies: The number of weekly backups(the latest one in each week) that should be kept.
// yearlies: The number of yearly backups(the latest one in each year) that should be kept.
func (loc *BackupLocator) Cleanup(keepLast string, lineage string, options rsapi.ApiParams) error {
	if keepLast == "" {
		return fmt.Errorf("keepLast is required")
	}
	if lineage == "" {
		return fmt.Errorf("lineage is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"keep_last": keepLast,
		"lineage":   lineage,
	}, options)
	var uri, err = loc.Url("Backups", "cleanup")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/backups
// Takes in an array of volumeattachmenthrefs and takes a snapshot of each.
// The volumeattachmenthrefs must belong to the same instance.
func (loc *BackupLocator) Create(backup *BackupParam) (*BackupLocator, error) {
	var res *BackupLocator
	if backup == nil {
		return res, fmt.Errorf("backup is required")
	}
	var params = rsapi.ApiParams{
		"backup": backup,
	}
	var uri, err = loc.Url("Backups", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &BackupLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/backups/:id
// Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.
func (loc *BackupLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Backups", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/backups
// Lists all of the backups with the given lineage tag. Filters can be used to search for a particular backup. If the
// 'latest_before' filter is set, only one backup is returned (the latest backup before the given timestamp).
// To get the latest completed backup, the 'completed' filter should be set to 'true' and the 'latest_before' filter
// should be set to the current timestamp. The format of the timestamp must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
// 2011/07/11 00:00:00 +0000.
// To get the latest completed backup just before, say 25 June 2009, then the 'completed' filter
// should be set to 'true' and the 'latest_before' filter should be set to 2009/06/25 00:00:00 +0000.
// lineage: Backups belonging to this lineage.
// -- Optional parameters:
// filter
func (loc *BackupLocator) Index(lineage string, options rsapi.ApiParams) ([]*Backup, error) {
	var res []*Backup
	if lineage == "" {
		return res, fmt.Errorf("lineage is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"lineage": lineage,
	}, options)
	var uri, err = loc.Url("Backups", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/backups/:id/restore
// Restores the given Backup.
// This call will:
// create the required number of Volumes from the volume_snapshots_hrefs in the given Backup,
// attach them to the given Instance at the device specified in the Snapshot. If the devices are already being used
// on the Instance, the Task will denote that the restore has failed.
// instanceHref: The instance href that the backup will be restored to.
// -- Optional parameters:
// backup
func (loc *BackupLocator) Restore(instanceHref string, options rsapi.ApiParams) error {
	if instanceHref == "" {
		return fmt.Errorf("instanceHref is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"instance_href": instanceHref,
	}, options)
	var uri, err = loc.Url("Backups", "restore")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/backups/:id
// Lists the attributes of a given backup
func (loc *BackupLocator) Show() (*Backup, error) {
	var res *Backup
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Backups", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/backups/:id
// Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.
func (loc *BackupLocator) Update(backup *BackupParam2) error {
	if backup == nil {
		return fmt.Errorf("backup is required")
	}
	var params = rsapi.ApiParams{
		"backup": backup,
	}
	var uri, err = loc.Url("Backups", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  ChildAccount ******/

//===== Locator

// ChildAccount resource locator, exposes resource actions.
type ChildAccountLocator struct {
	UrlResolver
	api *Api15
}

// ChildAccount resource locator factory
func (api *Api15) ChildAccountLocator(href string) *ChildAccountLocator {
	return &ChildAccountLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/child_accounts
// Create an enterprise ChildAccount for this Account. The User will by default get an 'admin' role
// on the ChildAccount to enable him/her to add, delete Users and Permissions.
// For more information on the valid values for 'cluster_href', refer to the following:
// http://support.rightscale.com/12-Guides/RightScale_API_1.5/Examples/ChildAccounts/Create
func (loc *ChildAccountLocator) Create(childAccount *ChildAccountParam) (*ChildAccountLocator, error) {
	var res *ChildAccountLocator
	if childAccount == nil {
		return res, fmt.Errorf("childAccount is required")
	}
	var params = rsapi.ApiParams{
		"child_account": childAccount,
	}
	var uri, err = loc.Url("ChildAccounts", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ChildAccountLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/child_accounts
// Lists the enterprise ChildAccounts available for this Account.
// -- Optional parameters:
// filter
func (loc *ChildAccountLocator) Index(options rsapi.ApiParams) ([]*Account, error) {
	var res []*Account
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("ChildAccounts", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// PUT /api/accounts/:id
// PUT /api/child_accounts/:id
// Update an enterprise ChildAccount for this Account.
func (loc *ChildAccountLocator) Update(childAccount *ChildAccountParam2) error {
	if childAccount == nil {
		return fmt.Errorf("childAccount is required")
	}
	var params = rsapi.ApiParams{
		"child_account": childAccount,
	}
	var uri, err = loc.Url("ChildAccounts", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Cloud ******/

// Represents a Cloud (within the context of the account in the session).
type Cloud struct {
	Capabilities []string            `json:"capabilities,omitempty"`
	CloudType    string              `json:"cloud_type,omitempty"`
	Description  string              `json:"description,omitempty"`
	DisplayName  string              `json:"display_name,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	Name         string              `json:"name,omitempty"`
}

//===== Locator

// Cloud resource locator, exposes resource actions.
type CloudLocator struct {
	UrlResolver
	api *Api15
}

// Cloud resource locator factory
func (api *Api15) CloudLocator(href string) *CloudLocator {
	return &CloudLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds
// Lists the clouds available to this account.
// -- Optional parameters:
// filter
// view
func (loc *CloudLocator) Index(options rsapi.ApiParams) ([]*Cloud, error) {
	var res []*Cloud
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Clouds", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:id
// Show information about a single cloud.
// -- Optional parameters:
// view
func (loc *CloudLocator) Show(options rsapi.ApiParams) (*Cloud, error) {
	var res *Cloud
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Clouds", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  CloudAccount ******/

// Represents a Cloud Account (An association between the account and a cloud).
type CloudAccount struct {
	CreatedAt time.Time           `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// CloudAccount resource locator, exposes resource actions.
type CloudAccountLocator struct {
	UrlResolver
	api *Api15
}

// CloudAccount resource locator factory
func (api *Api15) CloudAccountLocator(href string) *CloudAccountLocator {
	return &CloudAccountLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/cloud_accounts
// Create a CloudAccount by passing in the respective credentials for each cloud.
// For more information on the specific parameters for each cloud, refer to the following:
// http://support.rightscale.com/12-Guides/RightScale_API_1.5/Examples/Cloud_Accounts/Create_Cloud_Accounts
func (loc *CloudAccountLocator) Create(cloudAccount *CloudAccountParam) (*CloudAccountLocator, error) {
	var res *CloudAccountLocator
	if cloudAccount == nil {
		return res, fmt.Errorf("cloudAccount is required")
	}
	var params = rsapi.ApiParams{
		"cloud_account": cloudAccount,
	}
	var uri, err = loc.Url("CloudAccounts", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &CloudAccountLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/cloud_accounts/:id
// Delete a CloudAccount.
func (loc *CloudAccountLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("CloudAccounts", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/cloud_accounts
// Lists the CloudAccounts (non-aws) available to this Account.
func (loc *CloudAccountLocator) Index() ([]*CloudAccount, error) {
	var res []*CloudAccount
	var params = map[string]interface{}{}
	var uri, err = loc.Url("CloudAccounts", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/cloud_accounts/:id

func (loc *CloudAccountLocator) Show() (*CloudAccount, error) {
	var res *CloudAccount
	var params = map[string]interface{}{}
	var uri, err = loc.Url("CloudAccounts", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Cookbook ******/

// Represents a given instance of a single cookbook.
type Cookbook struct {
	Actions           []string            `json:"actions,omitempty"`
	CreatedAt         time.Time           `json:"created_at,omitempty"`
	DownloadUrl       string              `json:"download_url,omitempty"`
	Id                string              `json:"id,omitempty"`
	Links             []map[string]string `json:"links,omitempty"`
	Metadata          string              `json:"metadata,omitempty"`
	Name              string              `json:"name,omitempty"`
	Namespace         string              `json:"namespace,omitempty"`
	SourceInfoSummary string              `json:"source_info_summary,omitempty"`
	State             string              `json:"state,omitempty"`
	UpdatedAt         time.Time           `json:"updated_at,omitempty"`
	Version           string              `json:"version,omitempty"`
}

//===== Locator

// Cookbook resource locator, exposes resource actions.
type CookbookLocator struct {
	UrlResolver
	api *Api15
}

// Cookbook resource locator factory
func (api *Api15) CookbookLocator(href string) *CookbookLocator {
	return &CookbookLocator{UrlResolver(href), api}
}

//===== Actions

// DELETE /api/cookbooks/:id
// Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.
func (loc *CookbookLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Cookbooks", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/cookbooks/:id/follow
// Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.
// value: Indicates if this action should follow (true) or unfollow (false) a Cookbook.
func (loc *CookbookLocator) Follow(value string) error {
	if value == "" {
		return fmt.Errorf("value is required")
	}
	var params = rsapi.ApiParams{
		"value": value,
	}
	var uri, err = loc.Url("Cookbooks", "follow")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/cookbooks/:id/freeze
// Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.
// value: Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.
func (loc *CookbookLocator) Freeze(value string) error {
	if value == "" {
		return fmt.Errorf("value is required")
	}
	var params = rsapi.ApiParams{
		"value": value,
	}
	var uri, err = loc.Url("Cookbooks", "freeze")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/cookbooks
// Lists the Cookbooks available to this account.
// The extended_designer view is only available to accounts with the designer permission.
// -- Optional parameters:
// filter
// view
func (loc *CookbookLocator) Index(options rsapi.ApiParams) ([]*Cookbook, error) {
	var res []*Cookbook
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Cookbooks", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/cookbooks/:id/obsolete
// Marks a Cookbook as obsolete (or un-obsolete).
// value: Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.
func (loc *CookbookLocator) Obsolete(value string) error {
	if value == "" {
		return fmt.Errorf("value is required")
	}
	var params = rsapi.ApiParams{
		"value": value,
	}
	var uri, err = loc.Url("Cookbooks", "obsolete")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/cookbooks/:id
// Show information about a single Cookbook.
// The extended_designer view is only available to accounts with the designer permission.
// -- Optional parameters:
// view
func (loc *CookbookLocator) Show(options rsapi.ApiParams) (*Cookbook, error) {
	var res *Cookbook
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Cookbooks", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  CookbookAttachment ******/

// Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.
type CookbookAttachment struct {
	Actions    []string            `json:"actions,omitempty"`
	Dependency bool                `json:"dependency,omitempty"`
	Id         string              `json:"id,omitempty"`
	Links      []map[string]string `json:"links,omitempty"`
}

//===== Locator

// CookbookAttachment resource locator, exposes resource actions.
type CookbookAttachmentLocator struct {
	UrlResolver
	api *Api15
}

// CookbookAttachment resource locator factory
func (api *Api15) CookbookAttachmentLocator(href string) *CookbookAttachmentLocator {
	return &CookbookAttachmentLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/cookbooks/:cookbook_id/cookbook_attachments
// POST /api/server_templates/:server_template_id/cookbook_attachments
// POST /api/cookbook_attachments
// Attach a cookbook to a given resource.
// -- Optional parameters:
// cookbook_attachment
func (loc *CookbookAttachmentLocator) Create(options rsapi.ApiParams) (*CookbookAttachmentLocator, error) {
	var res *CookbookAttachmentLocator
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("CookbookAttachments", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &CookbookAttachmentLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/cookbooks/:cookbook_id/cookbook_attachments/:id
// DELETE /api/server_templates/:server_template_id/cookbook_attachments/:id
// DELETE /api/cookbook_attachments/:id
// Detach a cookbook from a given resource.
func (loc *CookbookAttachmentLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("CookbookAttachments", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/cookbooks/:cookbook_id/cookbook_attachments
// GET /api/server_templates/:server_template_id/cookbook_attachments
// GET /api/cookbook_attachments
// Lists Cookbook Attachments.
// -- Optional parameters:
// view
func (loc *CookbookAttachmentLocator) Index(options rsapi.ApiParams) ([]*CookbookAttachment, error) {
	var res []*CookbookAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("CookbookAttachments", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/server_templates/:server_template_id/cookbook_attachments/multi_attach
// POST /api/cookbook_attachments/multi_attach
// Attach multiple cookbooks to a given resource.
func (loc *CookbookAttachmentLocator) MultiAttach(cookbookAttachments *CookbookAttachments) error {
	if cookbookAttachments == nil {
		return fmt.Errorf("cookbookAttachments is required")
	}
	var params = rsapi.ApiParams{
		"cookbook_attachments": cookbookAttachments,
	}
	var uri, err = loc.Url("CookbookAttachments", "multi_attach")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_templates/:server_template_id/cookbook_attachments/multi_detach
// POST /api/cookbook_attachments/multi_detach
// Detach multiple cookbooks from a given resource.
func (loc *CookbookAttachmentLocator) MultiDetach(cookbookAttachments *CookbookAttachments2) error {
	if cookbookAttachments == nil {
		return fmt.Errorf("cookbookAttachments is required")
	}
	var params = rsapi.ApiParams{
		"cookbook_attachments": cookbookAttachments,
	}
	var uri, err = loc.Url("CookbookAttachments", "multi_detach")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/cookbooks/:cookbook_id/cookbook_attachments/:id
// GET /api/server_templates/:server_template_id/cookbook_attachments/:id
// GET /api/cookbook_attachments/:id
// Displays information about a single cookbook attachment to a ServerTemplate.
// -- Optional parameters:
// view
func (loc *CookbookAttachmentLocator) Show(options rsapi.ApiParams) (*CookbookAttachment, error) {
	var res *CookbookAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("CookbookAttachments", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Credential ******/

// A Credential provides an abstraction for sensitive textual information,
// such as passphrases or cloud credentials, whose value should be encrypted
// when stored in the database and not generally shown in the UI or through the
// API. Credentials may then be used as inputs of type "Cred" in RightScripts
// or Chef recipes. NOTE: Credential values may be updated through the API, but
// values cannot be retrieved via the API.
type Credential struct {
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
	Value       string              `json:"value,omitempty"`
}

//===== Locator

// Credential resource locator, exposes resource actions.
type CredentialLocator struct {
	UrlResolver
	api *Api15
}

// Credential resource locator factory
func (api *Api15) CredentialLocator(href string) *CredentialLocator {
	return &CredentialLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/credentials
// Creates a new Credential with the given parameters.
func (loc *CredentialLocator) Create(credential *CredentialParam) (*CredentialLocator, error) {
	var res *CredentialLocator
	if credential == nil {
		return res, fmt.Errorf("credential is required")
	}
	var params = rsapi.ApiParams{
		"credential": credential,
	}
	var uri, err = loc.Url("Credentials", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &CredentialLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/credentials/:id
// Deletes a Credential.
func (loc *CredentialLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Credentials", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/credentials
// Lists the Credentials available to this account.
// -- Optional parameters:
// filter
// view
func (loc *CredentialLocator) Index(options rsapi.ApiParams) ([]*Credential, error) {
	var res []*Credential
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Credentials", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/credentials/:id
// Show information about a single Credential. NOTE: Credential values may be updated through the API, but values cannot be retrieved via the API.
// -- Optional parameters:
// view
func (loc *CredentialLocator) Show(options rsapi.ApiParams) (*Credential, error) {
	var res *Credential
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Credentials", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/credentials/:id
// Updates attributes of a Credential.
func (loc *CredentialLocator) Update(credential *CredentialParam) error {
	if credential == nil {
		return fmt.Errorf("credential is required")
	}
	var params = rsapi.ApiParams{
		"credential": credential,
	}
	var uri, err = loc.Url("Credentials", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Datacenter ******/

// Datacenters represent isolated facilities within a cloud. The level and type of isolation is cloud dependent.
// While Datacenters in large public clouds might correspond to different physical buildings, with different power,
// internet links...etc., Datacenters within the context of a private cloud might simply correspond to having different network providers.
// Spreading servers across distinct Datacenters helps minimize outages.
type Datacenter struct {
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

//===== Locator

// Datacenter resource locator, exposes resource actions.
type DatacenterLocator struct {
	UrlResolver
	api *Api15
}

// Datacenter resource locator factory
func (api *Api15) DatacenterLocator(href string) *DatacenterLocator {
	return &DatacenterLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/datacenters
// Lists all Datacenters for a particular cloud.
// -- Optional parameters:
// filter
// view
func (loc *DatacenterLocator) Index(options rsapi.ApiParams) ([]*Datacenter, error) {
	var res []*Datacenter
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Datacenters", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/datacenters/:id
// Displays information about a single Datacenter.
// -- Optional parameters:
// view
func (loc *DatacenterLocator) Show(options rsapi.ApiParams) (*Datacenter, error) {
	var res *Datacenter
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Datacenters", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Deployment ******/

// Deployments represent logical groupings of related assets such as servers, server arrays, default configuration settings...etc.
type Deployment struct {
	Actions        []string            `json:"actions,omitempty"`
	Description    string              `json:"description,omitempty"`
	Inputs         []map[string]string `json:"inputs,omitempty"`
	Links          []map[string]string `json:"links,omitempty"`
	Locked         bool                `json:"locked,omitempty"`
	Name           string              `json:"name,omitempty"`
	ServerTagScope string              `json:"server_tag_scope,omitempty"`
}

//===== Locator

// Deployment resource locator, exposes resource actions.
type DeploymentLocator struct {
	UrlResolver
	api *Api15
}

// Deployment resource locator factory
func (api *Api15) DeploymentLocator(href string) *DeploymentLocator {
	return &DeploymentLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/deployments/:id/clone
// Clones a given deployment.
// -- Optional parameters:
// deployment
func (loc *DeploymentLocator) Clone(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Deployments", "clone")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/deployments
// Creates a new deployment with the given parameters.
func (loc *DeploymentLocator) Create(deployment *DeploymentParam) (*DeploymentLocator, error) {
	var res *DeploymentLocator
	if deployment == nil {
		return res, fmt.Errorf("deployment is required")
	}
	var params = rsapi.ApiParams{
		"deployment": deployment,
	}
	var uri, err = loc.Url("Deployments", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &DeploymentLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/deployments/:id
// Deletes a given deployment.
func (loc *DeploymentLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Deployments", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/deployments
// Lists deployments of the account.
// Using the available filters, one can select or group which deployments to retrieve.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// filter
// view
func (loc *DeploymentLocator) Index(options rsapi.ApiParams) ([]*Deployment, error) {
	var res []*Deployment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Deployments", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/deployments/:id/lock
// Locks a given deployment. Idempotent.
// Locking prevents servers from being deleted or moved from the deployment.
// Other actions such as adding servers or renaming the deployment are still allowed.
func (loc *DeploymentLocator) Lock() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Deployments", "lock")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/deployments/:id/servers
// Lists the servers belonging to this deployment. This call is equivalent to servers#index call, where the servers returned will
// automatically be filtered by this deployment. See servers#index for details on other options and parameters.
func (loc *DeploymentLocator) Servers() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Deployments", "servers")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/deployments/:id
// Lists the attributes of a given deployment.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// view
func (loc *DeploymentLocator) Show(options rsapi.ApiParams) (*Deployment, error) {
	var res *Deployment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Deployments", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// POST /api/deployments/:id/unlock
// Unlocks a given deployment. Idempotent.
func (loc *DeploymentLocator) Unlock() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Deployments", "unlock")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// PUT /api/deployments/:id
// Updates attributes of a given deployment.
func (loc *DeploymentLocator) Update(deployment *DeploymentParam) error {
	if deployment == nil {
		return fmt.Errorf("deployment is required")
	}
	var params = rsapi.ApiParams{
		"deployment": deployment,
	}
	var uri, err = loc.Url("Deployments", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  HealthCheck ******/

//===== Locator

// HealthCheck resource locator, exposes resource actions.
type HealthCheckLocator struct {
	UrlResolver
	api *Api15
}

// HealthCheck resource locator factory
func (api *Api15) HealthCheckLocator(href string) *HealthCheckLocator {
	return &HealthCheckLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/health-check/
// Check health of RightApi controllers
func (loc *HealthCheckLocator) Index() ([]*map[string]string, error) {
	var res []*map[string]string
	var params = map[string]interface{}{}
	var uri, err = loc.Url("HealthCheck", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

/******  IdentityProvider ******/

// An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account,
// and is trusted by the RightScale dashboard to authenticate your enterprise's end users.
// To register an Identity Provider, contact your account manager.
type IdentityProvider struct {
	Actions       []string            `json:"actions,omitempty"`
	CreatedAt     time.Time           `json:"created_at,omitempty"`
	DiscoveryHint string              `json:"discovery_hint,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	Name          string              `json:"name,omitempty"`
	UpdatedAt     time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// IdentityProvider resource locator, exposes resource actions.
type IdentityProviderLocator struct {
	UrlResolver
	api *Api15
}

// IdentityProvider resource locator factory
func (api *Api15) IdentityProviderLocator(href string) *IdentityProviderLocator {
	return &IdentityProviderLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/identity_providers
// Lists the identity providers associated with this enterprise account.
// -- Optional parameters:
// filter
// view
func (loc *IdentityProviderLocator) Index(options rsapi.ApiParams) ([]*IdentityProvider, error) {
	var res []*IdentityProvider
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("IdentityProviders", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/identity_providers/:id
// Show the specified identity provider, if associated with this enterprise account.
// -- Optional parameters:
// view
func (loc *IdentityProviderLocator) Show(options rsapi.ApiParams) (*IdentityProvider, error) {
	var res *IdentityProvider
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("IdentityProviders", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Image ******/

// Images represent base VM image existing in a cloud. An image will define the initial Operating System and root disk contents
// for a new Instance to have, and therefore it represents the basic starting point for creating a new one.
type Image struct {
	Actions            []string            `json:"actions,omitempty"`
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

//===== Locator

// Image resource locator, exposes resource actions.
type ImageLocator struct {
	UrlResolver
	api *Api15
}

// Image resource locator factory
func (api *Api15) ImageLocator(href string) *ImageLocator {
	return &ImageLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/images
// Lists all Images for the given Cloud.
// -- Optional parameters:
// filter
// view
func (loc *ImageLocator) Index(options rsapi.ApiParams) ([]*Image, error) {
	var res []*Image
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Images", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/images/:id
// Shows information about a single Image.
// -- Optional parameters:
// view
func (loc *ImageLocator) Show(options rsapi.ApiParams) (*Image, error) {
	var res *Image
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Images", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
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

// Input resource locator, exposes resource actions.
type InputLocator struct {
	UrlResolver
	api *Api15
}

// Input resource locator factory
func (api *Api15) InputLocator(href string) *InputLocator {
	return &InputLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instances/:instance_id/inputs
// GET /api/deployments/:deployment_id/inputs
// GET /api/server_templates/:server_template_id/inputs
// Retrieves the full list of existing inputs of the specified resource.
// -- Optional parameters:
// view
func (loc *InputLocator) Index(options rsapi.ApiParams) ([]*Input, error) {
	var res []*Input
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Inputs", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// PUT /api/clouds/:cloud_id/instances/:instance_id/inputs/multi_update
// PUT /api/deployments/:deployment_id/inputs/multi_update
// PUT /api/server_templates/:server_template_id/inputs/multi_update
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
// Ignore (input is not present)
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
func (loc *InputLocator) MultiUpdate(inputs map[string]string) error {
	if len(inputs) == 0 {
		return fmt.Errorf("inputs is required")
	}
	var params = rsapi.ApiParams{
		"inputs": inputs,
	}
	var uri, err = loc.Url("Inputs", "multi_update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
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
	Actions                  []string            `json:"actions,omitempty"`
	AssociatePublicIpAddress bool                `json:"associate_public_ip_address,omitempty"`
	CloudSpecificAttributes  map[string]string   `json:"cloud_specific_attributes,omitempty"`
	CreatedAt                time.Time           `json:"created_at,omitempty"`
	Description              string              `json:"description,omitempty"`
	InheritedSources         []string            `json:"inherited_sources,omitempty"`
	Inputs                   []map[string]string `json:"inputs,omitempty"`
	IpForwardingEnabled      bool                `json:"ip_forwarding_enabled,omitempty"`
	Links                    []map[string]string `json:"links,omitempty"`
	Locked                   bool                `json:"locked,omitempty"`
	MonitoringId             string              `json:"monitoring_id,omitempty"`
	MonitoringServer         string              `json:"monitoring_server,omitempty"`
	Name                     string              `json:"name,omitempty"`
	OsPlatform               string              `json:"os_platform,omitempty"`
	PricingType              string              `json:"pricing_type,omitempty"`
	PrivateDnsNames          []string            `json:"private_dns_names,omitempty"`
	PrivateIpAddresses       []string            `json:"private_ip_addresses,omitempty"`
	PublicDnsNames           []string            `json:"public_dns_names,omitempty"`
	PublicIpAddresses        []string            `json:"public_ip_addresses,omitempty"`
	ResourceUid              string              `json:"resource_uid,omitempty"`
	SecurityGroups           []SecurityGroup     `json:"security_groups,omitempty"`
	State                    string              `json:"state,omitempty"`
	Subnets                  []Subnet            `json:"subnets,omitempty"`
	TerminatedAt             time.Time           `json:"terminated_at,omitempty"`
	UpdatedAt                time.Time           `json:"updated_at,omitempty"`
	UserData                 string              `json:"user_data,omitempty"`
}

//===== Locator

// Instance resource locator, exposes resource actions.
type InstanceLocator struct {
	UrlResolver
	api *Api15
}

// Instance resource locator factory
func (api *Api15) InstanceLocator(href string) *InstanceLocator {
	return &InstanceLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances
// Creates and launches a raw instance using the provided parameters.
func (loc *InstanceLocator) Create(instance *InstanceParam) (*InstanceLocator, error) {
	var res *InstanceLocator
	if instance == nil {
		return res, fmt.Errorf("instance is required")
	}
	var params = rsapi.ApiParams{
		"instance": instance,
	}
	var uri, err = loc.Url("Instances", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &InstanceLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/clouds/:cloud_id/instances
// GET /api/clouds/:cloud_id/instances
// GET /api/server_arrays/:server_array_id/current_instances
// Lists instances of a given cloud, server array.
// Using the available filters, it is possible to craft powerful queries about which instances to retrieve.
// For example, one can easily list:
// instances that have names that contain "app"
// all instances of a given deployment
// instances belonging to a given server array (i.e., have the same parent_url)
// To see the instances of a server array including the next_instance, use the URL "/api/clouds/:cloud_id/instances" with the filter "parent_href==/api/server_arrays/XX". To list only the running
// instances of a server array, use the URL "/api/server_arrays/:server_array_id/current_instances"
// The 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// filter
// view
func (loc *InstanceLocator) Index(options rsapi.ApiParams) ([]*Instance, error) {
	var res []*Instance
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Instances", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/clouds/:cloud_id/instances/:id/launch
// POST /api/servers/:server_id/launch
// POST /api/server_arrays/:server_array_id/launch
// Launches an instance using the parameters that this instance has been configured with.
// Note that this action can only be performed in "next" instances, and not on instances that are already running.
// -- Optional parameters:
// api_behavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
// inputs
func (loc *InstanceLocator) Launch(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Instances", "launch")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:id/lock

func (loc *InstanceLocator) Lock() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Instances", "lock")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/multi_run_executable
// POST /api/server_arrays/:server_array_id/multi_run_executable
// Runs a script or a recipe in the running instances.
// This is an asynchronous function, which returns immediately after queuing the executable for execution.
// Status of the execution can be tracked at the URL returned in the "Location" header.
// -- Optional parameters:
// filter
// ignore_lock: Specifies the ability to ignore the lock(s) on the Instance(s).
// inputs
// recipe_name: The name of the recipe to be run.
// right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.
func (loc *InstanceLocator) MultiRunExecutable(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Instances", "multi_run_executable")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/multi_terminate
// POST /api/server_arrays/:server_array_id/multi_terminate
// Terminates running instances.
// Either a filter or the parameter 'terminate_all' must be provided.
// -- Optional parameters:
// filter
// terminate_all: Specifies the ability to terminate all instances.
func (loc *InstanceLocator) MultiTerminate(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Instances", "multi_terminate")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:id/reboot
// POST /api/servers/:server_id/reboot
// Reboot a running instance.
// Note that this action can only succeed if the instance is running. One cannot reboot instances of type "next".
func (loc *InstanceLocator) Reboot() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Instances", "reboot")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:id/run_executable
// Runs a script or a recipe in the running instance.
// This is an asynchronous function, which returns immediately after queuing the executable for execution.
// Status of the execution can be tracked at the URL returned in the "Location" header.
// Note that this can only be performed on running instances.
// -- Optional parameters:
// ignore_lock: Specifies the ability to ignore the lock on the Instance.
// inputs
// recipe_name: The name of the recipe to run.
// right_script_href: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.
func (loc *InstanceLocator) RunExecutable(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Instances", "run_executable")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:id/set_custom_lodgement
// This method is deprecated.  Please use InstanceCustomLodgement.
// quantity: At least one name/value pair must be specified. Currently, a maximum of 2 name/value pairs is supported.
// timeframe: The timeframe (either a month or a single day) for which the quantity value is valid (currently for the PDT timezone only).
func (loc *InstanceLocator) SetCustomLodgement(quantity []*Quantity, timeframe string) error {
	if len(quantity) == 0 {
		return fmt.Errorf("quantity is required")
	}
	if timeframe == "" {
		return fmt.Errorf("timeframe is required")
	}
	var params = rsapi.ApiParams{
		"quantity":  quantity,
		"timeframe": timeframe,
	}
	var uri, err = loc.Url("Instances", "set_custom_lodgement")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/instances/:id
// Shows attributes of a single instance.
// The 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// view
func (loc *InstanceLocator) Show(options rsapi.ApiParams) (*Instance, error) {
	var res *Instance
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Instances", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// POST /api/clouds/:cloud_id/instances/:id/start
// Starts an instance that has been stopped, resuming it to its previously saved volume state.
// After an instance is started, the reference to your instance will have a different id.
// The new id can be found by performing an index query with the appropriate filters on the
// Instances resource, performing a show action on the Server resource for Server Instances, or
// performing a current_instances action on the ServerArray resource for ServerArray Instances.
func (loc *InstanceLocator) Start() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Instances", "start")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:id/stop
// Stores the instance's current volume state to resume later using the 'start' action.
// After an instance is stopped, the reference to your instance will have a different id.
// The new id can be found by performing an index query with the appropriate filters on the
// Instances resource, performing a show action on the Server resource for Server Instances, or performing a current_instances action on the ServerArray resource for ServerArray Instances.
func (loc *InstanceLocator) Stop() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Instances", "stop")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:id/terminate
// POST /api/servers/:server_id/terminate
// Terminates a running instance.
// Note that this action can succeed only if the instance is running. One cannot terminate instances of type "next".
func (loc *InstanceLocator) Terminate() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Instances", "terminate")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/clouds/:cloud_id/instances/:id/unlock

func (loc *InstanceLocator) Unlock() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Instances", "unlock")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// PUT /api/clouds/:cloud_id/instances/:id
// Updates attributes of a single instance.
func (loc *InstanceLocator) Update(instance *InstanceParam2) error {
	if instance == nil {
		return fmt.Errorf("instance is required")
	}
	var params = rsapi.ApiParams{
		"instance": instance,
	}
	var uri, err = loc.Url("Instances", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  InstanceCustomLodgement ******/

// An InstanceCustomLodgement represents a way to create custom reports about a specific instance with a user defined quantity.  Replaces the legacy Instances#setcustomlodgement interface.
type InstanceCustomLodgement struct {
	AccountOwner                         string                   `json:"account_owner,omitempty"`
	Actions                              []string                 `json:"actions,omitempty"`
	EndAt                                time.Time                `json:"end_at,omitempty"`
	Links                                []map[string]string      `json:"links,omitempty"`
	Quantity                             []map[string]interface{} `json:"quantity,omitempty"`
	ResourceBillingEndAt                 time.Time                `json:"resource_billing_end_at,omitempty"`
	ResourceBillingStartAt               time.Time                `json:"resource_billing_start_at,omitempty"`
	ResourceInstanceType                 string                   `json:"resource_instance_type,omitempty"`
	ResourceLaunchedBy                   string                   `json:"resource_launched_by,omitempty"`
	ResourceTemplateLibraryHref          string                   `json:"resource_template_library_href,omitempty"`
	ResourceTemplateName                 string                   `json:"resource_template_name,omitempty"`
	ResourceTemplatePublishedByAccountId string                   `json:"resource_template_published_by_account_id,omitempty"`
	ResourceUid                          string                   `json:"resource_uid,omitempty"`
	StartAt                              time.Time                `json:"start_at,omitempty"`
}

//===== Locator

// InstanceCustomLodgement resource locator, exposes resource actions.
type InstanceCustomLodgementLocator struct {
	UrlResolver
	api *Api15
}

// InstanceCustomLodgement resource locator factory
func (api *Api15) InstanceCustomLodgementLocator(href string) *InstanceCustomLodgementLocator {
	return &InstanceCustomLodgementLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements
// Create a lodgement with the quantity and timeframe specified.
// quantity: At least one name/value pair must be specified. Currently, a maximum of 2 name/value pairs is supported.
// timeframe: The time-frame (either a month "YYYY_MM" or a single day "YYYY_MM_DD") for which the quantity value is valid (currently for the PDT timezone only).
func (loc *InstanceCustomLodgementLocator) Create(quantity []*Quantity, timeframe string) (*InstanceCustomLodgementLocator, error) {
	var res *InstanceCustomLodgementLocator
	if len(quantity) == 0 {
		return res, fmt.Errorf("quantity is required")
	}
	if timeframe == "" {
		return res, fmt.Errorf("timeframe is required")
	}
	var params = rsapi.ApiParams{
		"quantity":  quantity,
		"timeframe": timeframe,
	}
	var uri, err = loc.Url("InstanceCustomLodgements", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &InstanceCustomLodgementLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id
// Destroy the specified lodgement.
func (loc *InstanceCustomLodgementLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("InstanceCustomLodgements", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements
// List InstanceCustomLodgements of a given cloud and instance.
// -- Optional parameters:
// view
func (loc *InstanceCustomLodgementLocator) Index(options rsapi.ApiParams) ([]*InstanceCustomLodgement, error) {
	var res []*InstanceCustomLodgement
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("InstanceCustomLodgements", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id
// Show the specified lodgement.
func (loc *InstanceCustomLodgementLocator) Show() (*InstanceCustomLodgement, error) {
	var res *InstanceCustomLodgement
	var params = map[string]interface{}{}
	var uri, err = loc.Url("InstanceCustomLodgements", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id
// Update a lodgement with the quantity specified.
// quantity: At least one name/value pair must be specified. Currently, a maximum of 2 name/value pairs is supported.
func (loc *InstanceCustomLodgementLocator) Update(quantity []*Quantity) error {
	if len(quantity) == 0 {
		return fmt.Errorf("quantity is required")
	}
	var params = rsapi.ApiParams{
		"quantity": quantity,
	}
	var uri, err = loc.Url("InstanceCustomLodgements", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  InstanceType ******/

type InstanceType struct {
	Actions         []string            `json:"actions,omitempty"`
	CpuArchitecture string              `json:"cpu_architecture,omitempty"`
	CpuCount        int                 `json:"cpu_count,omitempty"`
	CpuSpeed        string              `json:"cpu_speed,omitempty"`
	Description     string              `json:"description,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	LocalDiskSize   string              `json:"local_disk_size,omitempty"`
	LocalDisks      int                 `json:"local_disks,omitempty"`
	Memory          string              `json:"memory,omitempty"`
	Name            string              `json:"name,omitempty"`
	ResourceUid     string              `json:"resource_uid,omitempty"`
}

//===== Locator

// InstanceType resource locator, exposes resource actions.
type InstanceTypeLocator struct {
	UrlResolver
	api *Api15
}

// InstanceType resource locator factory
func (api *Api15) InstanceTypeLocator(href string) *InstanceTypeLocator {
	return &InstanceTypeLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instance_types
// Lists instance types.
// -- Optional parameters:
// filter
// view
func (loc *InstanceTypeLocator) Index(options rsapi.ApiParams) ([]*InstanceType, error) {
	var res []*InstanceType
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("InstanceTypes", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/instance_types/:id
// Displays information about a single Instance type.
// -- Optional parameters:
// view
func (loc *InstanceTypeLocator) Show(options rsapi.ApiParams) (*InstanceType, error) {
	var res *InstanceType
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("InstanceTypes", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  IpAddress ******/

// An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.
type IpAddress struct {
	Address   string              `json:"address,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	Domain    string              `json:"domain,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	Name      string              `json:"name,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// IpAddress resource locator, exposes resource actions.
type IpAddressLocator struct {
	UrlResolver
	api *Api15
}

// IpAddress resource locator factory
func (api *Api15) IpAddressLocator(href string) *IpAddressLocator {
	return &IpAddressLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/ip_addresses
// Creates a new IpAddress with the given parameters.
func (loc *IpAddressLocator) Create(ipAddress *IpAddressParam) (*IpAddressLocator, error) {
	var res *IpAddressLocator
	if ipAddress == nil {
		return res, fmt.Errorf("ipAddress is required")
	}
	var params = rsapi.ApiParams{
		"ip_address": ipAddress,
	}
	var uri, err = loc.Url("IpAddresses", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &IpAddressLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/ip_addresses/:id
// Deletes a given IpAddress.
func (loc *IpAddressLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("IpAddresses", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/ip_addresses
// Lists the IpAddresses available to this account.
// -- Optional parameters:
// filter
func (loc *IpAddressLocator) Index(options rsapi.ApiParams) ([]*IpAddress, error) {
	var res []*IpAddress
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("IpAddresses", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/ip_addresses/:id
// Show information about a single IpAddress.
func (loc *IpAddressLocator) Show() (*IpAddress, error) {
	var res *IpAddress
	var params = map[string]interface{}{}
	var uri, err = loc.Url("IpAddresses", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/clouds/:cloud_id/ip_addresses/:id
// Updates attributes of a given IpAddress.
func (loc *IpAddressLocator) Update(ipAddress *IpAddressParam2) error {
	if ipAddress == nil {
		return fmt.Errorf("ipAddress is required")
	}
	var params = rsapi.ApiParams{
		"ip_address": ipAddress,
	}
	var uri, err = loc.Url("IpAddresses", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  IpAddressBinding ******/

// An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.
// The IpAddress is bound immediately for a current instance, or on launch for a next instance.
// It also allows specifying port forwarding rules for that particular IpAddress and Instance pair.
type IpAddressBinding struct {
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	PrivatePort string              `json:"private_port,omitempty"`
	Protocol    string              `json:"protocol,omitempty"`
	PublicPort  string              `json:"public_port,omitempty"`
	Recurring   bool                `json:"recurring,omitempty"`
}

//===== Locator

// IpAddressBinding resource locator, exposes resource actions.
type IpAddressBindingLocator struct {
	UrlResolver
	api *Api15
}

// IpAddressBinding resource locator factory
func (api *Api15) IpAddressBindingLocator(href string) *IpAddressBindingLocator {
	return &IpAddressBindingLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings
// POST /api/clouds/:cloud_id/ip_address_bindings
// Creates an ip address binding which attaches a specified IpAddress resource
// to a specified instance, and also allows for configuration of port forwarding
// rules. If the instance specified is a current (running) instance, a one-time
// IpAddressBinding will be created. If the instance is a next instance, then
// a recurring IpAddressBinding is created, which will cause the IpAddress to
// be bound each time the incarnator boots.
func (loc *IpAddressBindingLocator) Create(ipAddressBinding *IpAddressBindingParam) (*IpAddressBindingLocator, error) {
	var res *IpAddressBindingLocator
	if ipAddressBinding == nil {
		return res, fmt.Errorf("ipAddressBinding is required")
	}
	var params = rsapi.ApiParams{
		"ip_address_binding": ipAddressBinding,
	}
	var uri, err = loc.Url("IpAddressBindings", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &IpAddressBindingLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id
// DELETE /api/clouds/:cloud_id/ip_address_bindings/:id
// No description provided for destroy.
func (loc *IpAddressBindingLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("IpAddressBindings", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings
// GET /api/clouds/:cloud_id/ip_address_bindings
// Lists the ip address bindings available to this account.
// -- Optional parameters:
// filter
func (loc *IpAddressBindingLocator) Index(options rsapi.ApiParams) ([]*IpAddressBinding, error) {
	var res []*IpAddressBinding
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("IpAddressBindings", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id
// GET /api/clouds/:cloud_id/ip_address_bindings/:id
// Show information about a single ip address binding.
func (loc *IpAddressBindingLocator) Show() (*IpAddressBinding, error) {
	var res *IpAddressBinding
	var params = map[string]interface{}{}
	var uri, err = loc.Url("IpAddressBindings", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  MonitoringMetric ******/

// A monitoring metric is a stream of data that is captured in an instance. Metrics can be monitored, graphed and can be used as the basis for triggering alerts.
type MonitoringMetric struct {
	Actions   []string            `json:"actions,omitempty"`
	GraphHref string              `json:"graph_href,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	Plugin    string              `json:"plugin,omitempty"`
	View      string              `json:"view,omitempty"`
}

//===== Locator

// MonitoringMetric resource locator, exposes resource actions.
type MonitoringMetricLocator struct {
	UrlResolver
	api *Api15
}

// MonitoringMetric resource locator factory
func (api *Api15) MonitoringMetricLocator(href string) *MonitoringMetricLocator {
	return &MonitoringMetricLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id/data
// Gives the raw monitoring data for a particular metric. The response will include different variables
// associated with that metric and the data points for each of those variables.
// To get the data for a certain duration, for e.g. for the last 10 minutes(600 secs), provide the variables
// start="-600" and end="0".
// end: An integer number of seconds from current time. e.g. -150 or 0
// start: An integer number of seconds from current time. e.g. -300
func (loc *MonitoringMetricLocator) Data(end string, start string) (map[string]string, error) {
	var res map[string]string
	if end == "" {
		return res, fmt.Errorf("end is required")
	}
	if start == "" {
		return res, fmt.Errorf("start is required")
	}
	var params = rsapi.ApiParams{
		"end":   end,
		"start": start,
	}
	var uri, err = loc.Url("MonitoringMetrics", "data")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics
// Lists the monitoring metrics available for the instance and their corresponding graph hrefs.
// Making a request to the graph_href will return a png image corresponding to that monitoring metric.
// -- Optional parameters:
// filter
// period: The time scale for which the graph is generated. Default is 'day'
// size: The size of the graph to be generated. Default is 'small'.
// title: The title of the graph.
// tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.
func (loc *MonitoringMetricLocator) Index(options rsapi.ApiParams) ([]*MonitoringMetric, error) {
	var res []*MonitoringMetric
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("MonitoringMetrics", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id
// Shows attributes of a single monitoring metric.
// Making a request to the graph_href will return a png image corresponding to that monitoring metric.
// -- Optional parameters:
// period: The time scale for which the graph is generated. Default is 'day'.
// size: The size of the graph to be generated. Default is 'small'.
// title: The title of the graph.
// tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.
func (loc *MonitoringMetricLocator) Show(options rsapi.ApiParams) (*MonitoringMetric, error) {
	var res *MonitoringMetric
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("MonitoringMetrics", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  MultiCloudImage ******/

// A MultiCloudImage is a RightScale component that functions as a pointer to machine images in specific clouds
// (e.g. AWS US-East, Rackspace). Each ServerTemplate can reference many MultiCloudImages that defines which
// image should be used when a server is launched in a particular cloud.
type MultiCloudImage struct {
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Revision    string              `json:"revision,omitempty"`
}

//===== Locator

// MultiCloudImage resource locator, exposes resource actions.
type MultiCloudImageLocator struct {
	UrlResolver
	api *Api15
}

// MultiCloudImage resource locator factory
func (api *Api15) MultiCloudImageLocator(href string) *MultiCloudImageLocator {
	return &MultiCloudImageLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/multi_cloud_images/:id/clone
// Clones a given MultiCloudImage.
func (loc *MultiCloudImageLocator) Clone(multiCloudImage *MultiCloudImageParam) error {
	if multiCloudImage == nil {
		return fmt.Errorf("multiCloudImage is required")
	}
	var params = rsapi.ApiParams{
		"multi_cloud_image": multiCloudImage,
	}
	var uri, err = loc.Url("MultiCloudImages", "clone")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/multi_cloud_images/:id/commit
// Commits a given MultiCloudImage. Only HEAD revisions can be committed.
// commitMessage: The message associated with the commit.
func (loc *MultiCloudImageLocator) Commit(commitMessage string) error {
	if commitMessage == "" {
		return fmt.Errorf("commitMessage is required")
	}
	var params = rsapi.ApiParams{
		"commit_message": commitMessage,
	}
	var uri, err = loc.Url("MultiCloudImages", "commit")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_templates/:server_template_id/multi_cloud_images
// POST /api/multi_cloud_images
// Creates a new MultiCloudImage with the given parameters.
func (loc *MultiCloudImageLocator) Create(multiCloudImage *MultiCloudImageParam) (*MultiCloudImageLocator, error) {
	var res *MultiCloudImageLocator
	if multiCloudImage == nil {
		return res, fmt.Errorf("multiCloudImage is required")
	}
	var params = rsapi.ApiParams{
		"multi_cloud_image": multiCloudImage,
	}
	var uri, err = loc.Url("MultiCloudImages", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &MultiCloudImageLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/server_templates/:server_template_id/multi_cloud_images/:id
// DELETE /api/multi_cloud_images/:id
// Deletes a given MultiCloudImage.
func (loc *MultiCloudImageLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("MultiCloudImages", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_templates/:server_template_id/multi_cloud_images
// GET /api/multi_cloud_images
// Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.
// -- Optional parameters:
// filter
func (loc *MultiCloudImageLocator) Index(options rsapi.ApiParams) ([]*MultiCloudImage, error) {
	var res []*MultiCloudImage
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("MultiCloudImages", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/server_templates/:server_template_id/multi_cloud_images/:id
// GET /api/multi_cloud_images/:id
// Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.
func (loc *MultiCloudImageLocator) Show() (*MultiCloudImage, error) {
	var res *MultiCloudImage
	var params = map[string]interface{}{}
	var uri, err = loc.Url("MultiCloudImages", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/server_templates/:server_template_id/multi_cloud_images/:id
// PUT /api/multi_cloud_images/:id
// Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).
// Currently, the attributes you can update are only the 'direct' attributes of a server template.
func (loc *MultiCloudImageLocator) Update(multiCloudImage *MultiCloudImageParam) error {
	if multiCloudImage == nil {
		return fmt.Errorf("multiCloudImage is required")
	}
	var params = rsapi.ApiParams{
		"multi_cloud_image": multiCloudImage,
	}
	var uri, err = loc.Url("MultiCloudImages", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  MultiCloudImageSetting ******/

// A MultiCloudImageSetting defines which
// settings should be used when a server is launched in a cloud.
type MultiCloudImageSetting struct {
	Actions []string            `json:"actions,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
}

//===== Locator

// MultiCloudImageSetting resource locator, exposes resource actions.
type MultiCloudImageSettingLocator struct {
	UrlResolver
	api *Api15
}

// MultiCloudImageSetting resource locator factory
func (api *Api15) MultiCloudImageSettingLocator(href string) *MultiCloudImageSettingLocator {
	return &MultiCloudImageSettingLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/multi_cloud_images/:multi_cloud_image_id/settings
// Creates a new setting for an existing MultiCloudImage.
func (loc *MultiCloudImageSettingLocator) Create(multiCloudImageSetting *MultiCloudImageSettingParam) (*MultiCloudImageSettingLocator, error) {
	var res *MultiCloudImageSettingLocator
	if multiCloudImageSetting == nil {
		return res, fmt.Errorf("multiCloudImageSetting is required")
	}
	var params = rsapi.ApiParams{
		"multi_cloud_image_setting": multiCloudImageSetting,
	}
	var uri, err = loc.Url("MultiCloudImageSettings", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &MultiCloudImageSettingLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/multi_cloud_images/:multi_cloud_image_id/settings/:id
// Deletes a MultiCloudImage setting.
func (loc *MultiCloudImageSettingLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("MultiCloudImageSettings", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/multi_cloud_images/:multi_cloud_image_id/settings
// Lists the MultiCloudImage settings.
// -- Optional parameters:
// filter
func (loc *MultiCloudImageSettingLocator) Index(options rsapi.ApiParams) ([]*MultiCloudImageSetting, error) {
	var res []*MultiCloudImageSetting
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("MultiCloudImageSettings", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/multi_cloud_images/:multi_cloud_image_id/settings/:id
// Show information about a single MultiCloudImage setting.
func (loc *MultiCloudImageSettingLocator) Show() (*MultiCloudImageSetting, error) {
	var res *MultiCloudImageSetting
	var params = map[string]interface{}{}
	var uri, err = loc.Url("MultiCloudImageSettings", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/multi_cloud_images/:multi_cloud_image_id/settings/:id
// Updates a settings for a MultiCLoudImage.
func (loc *MultiCloudImageSettingLocator) Update(multiCloudImageSetting *MultiCloudImageSettingParam) error {
	if multiCloudImageSetting == nil {
		return fmt.Errorf("multiCloudImageSetting is required")
	}
	var params = rsapi.ApiParams{
		"multi_cloud_image_setting": multiCloudImageSetting,
	}
	var uri, err = loc.Url("MultiCloudImageSettings", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Network ******/

// A Network is a logical grouping of network devices.
type Network struct {
	Actions         []string            `json:"actions,omitempty"`
	CidrBlock       string              `json:"cidr_block,omitempty"`
	Description     string              `json:"description,omitempty"`
	InstanceTenancy string              `json:"instance_tenancy,omitempty"`
	IsDefault       bool                `json:"is_default,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	ResourceUid     string              `json:"resource_uid,omitempty"`
}

//===== Locator

// Network resource locator, exposes resource actions.
type NetworkLocator struct {
	UrlResolver
	api *Api15
}

// Network resource locator factory
func (api *Api15) NetworkLocator(href string) *NetworkLocator {
	return &NetworkLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/networks
// Creates a new network.
func (loc *NetworkLocator) Create(network *NetworkParam) (*NetworkLocator, error) {
	var res *NetworkLocator
	if network == nil {
		return res, fmt.Errorf("network is required")
	}
	var params = rsapi.ApiParams{
		"network": network,
	}
	var uri, err = loc.Url("Networks", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &NetworkLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/networks/:id
// Deletes the given network(s).
func (loc *NetworkLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Networks", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/networks
// Lists networks in this account.
// -- Optional parameters:
// filter
func (loc *NetworkLocator) Index(options rsapi.ApiParams) ([]*Network, error) {
	var res []*Network
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Networks", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/networks/:id
// Shows attributes of a single network.
func (loc *NetworkLocator) Show() (*Network, error) {
	var res *Network
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Networks", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/networks/:id
// Updates the given network.
func (loc *NetworkLocator) Update(network *NetworkParam2) error {
	if network == nil {
		return fmt.Errorf("network is required")
	}
	var params = rsapi.ApiParams{
		"network": network,
	}
	var uri, err = loc.Url("Networks", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  NetworkGateway ******/

// A NetworkGateway is an interface that allows traffic to be routed between networks.
type NetworkGateway struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	Type        string              `json:"type,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// NetworkGateway resource locator, exposes resource actions.
type NetworkGatewayLocator struct {
	UrlResolver
	api *Api15
}

// NetworkGateway resource locator factory
func (api *Api15) NetworkGatewayLocator(href string) *NetworkGatewayLocator {
	return &NetworkGatewayLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/network_gateways
// Create a new NetworkGateway.
func (loc *NetworkGatewayLocator) Create(networkGateway *NetworkGatewayParam) (*NetworkGatewayLocator, error) {
	var res *NetworkGatewayLocator
	if networkGateway == nil {
		return res, fmt.Errorf("networkGateway is required")
	}
	var params = rsapi.ApiParams{
		"network_gateway": networkGateway,
	}
	var uri, err = loc.Url("NetworkGateways", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &NetworkGatewayLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/network_gateways/:id
// Delete an existing NetworkGateway.
func (loc *NetworkGatewayLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("NetworkGateways", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/network_gateways
// Lists the NetworkGateways available to this account.
// -- Optional parameters:
// filter
func (loc *NetworkGatewayLocator) Index(options rsapi.ApiParams) ([]*NetworkGateway, error) {
	var res []*NetworkGateway
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("NetworkGateways", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/network_gateways/:id
// Show information about a single NetworkGateway.
func (loc *NetworkGatewayLocator) Show() (*NetworkGateway, error) {
	var res *NetworkGateway
	var params = map[string]interface{}{}
	var uri, err = loc.Url("NetworkGateways", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/network_gateways/:id
// Update an existing NetworkGateway.
func (loc *NetworkGatewayLocator) Update(networkGateway *NetworkGatewayParam2) error {
	if networkGateway == nil {
		return fmt.Errorf("networkGateway is required")
	}
	var params = rsapi.ApiParams{
		"network_gateway": networkGateway,
	}
	var uri, err = loc.Url("NetworkGateways", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
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
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Options     map[string]string   `json:"options,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Type        string              `json:"type,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// NetworkOptionGroup resource locator, exposes resource actions.
type NetworkOptionGroupLocator struct {
	UrlResolver
	api *Api15
}

// NetworkOptionGroup resource locator factory
func (api *Api15) NetworkOptionGroupLocator(href string) *NetworkOptionGroupLocator {
	return &NetworkOptionGroupLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/network_option_groups
// Create a new NetworkOptionGroup.
func (loc *NetworkOptionGroupLocator) Create(networkOptionGroup *NetworkOptionGroupParam) (*NetworkOptionGroupLocator, error) {
	var res *NetworkOptionGroupLocator
	if networkOptionGroup == nil {
		return res, fmt.Errorf("networkOptionGroup is required")
	}
	var params = rsapi.ApiParams{
		"network_option_group": networkOptionGroup,
	}
	var uri, err = loc.Url("NetworkOptionGroups", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &NetworkOptionGroupLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/network_option_groups/:id
// Delete an existing NetworkOptionGroup.
func (loc *NetworkOptionGroupLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("NetworkOptionGroups", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/network_option_groups
// List NetworkOptionGroups available in this account.
// -- Optional parameters:
// filter
func (loc *NetworkOptionGroupLocator) Index(options rsapi.ApiParams) ([]*NetworkOptionGroup, error) {
	var res []*NetworkOptionGroup
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("NetworkOptionGroups", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/network_option_groups/:id
// Show information about a single NetworkOptionGroup.
func (loc *NetworkOptionGroupLocator) Show() (*NetworkOptionGroup, error) {
	var res *NetworkOptionGroup
	var params = map[string]interface{}{}
	var uri, err = loc.Url("NetworkOptionGroups", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/network_option_groups/:id
// Update an existing NetworkOptionGroup.
func (loc *NetworkOptionGroupLocator) Update(networkOptionGroup *NetworkOptionGroupParam2) error {
	if networkOptionGroup == nil {
		return fmt.Errorf("networkOptionGroup is required")
	}
	var params = rsapi.ApiParams{
		"network_option_group": networkOptionGroup,
	}
	var uri, err = loc.Url("NetworkOptionGroups", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
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
	Actions            []string            `json:"actions,omitempty"`
	CreatedAt          time.Time           `json:"created_at,omitempty"`
	Links              []map[string]string `json:"links,omitempty"`
	NetworkOptionGroup string              `json:"network_option_group,omitempty"`
	ResourceUid        string              `json:"resource_uid,omitempty"`
	UpdatedAt          time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// NetworkOptionGroupAttachment resource locator, exposes resource actions.
type NetworkOptionGroupAttachmentLocator struct {
	UrlResolver
	api *Api15
}

// NetworkOptionGroupAttachment resource locator factory
func (api *Api15) NetworkOptionGroupAttachmentLocator(href string) *NetworkOptionGroupAttachmentLocator {
	return &NetworkOptionGroupAttachmentLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/network_option_group_attachments
// Create a new NetworkOptionGroupAttachment.
func (loc *NetworkOptionGroupAttachmentLocator) Create(networkOptionGroupAttachment *NetworkOptionGroupAttachmentParam) (*NetworkOptionGroupAttachmentLocator, error) {
	var res *NetworkOptionGroupAttachmentLocator
	if networkOptionGroupAttachment == nil {
		return res, fmt.Errorf("networkOptionGroupAttachment is required")
	}
	var params = rsapi.ApiParams{
		"network_option_group_attachment": networkOptionGroupAttachment,
	}
	var uri, err = loc.Url("NetworkOptionGroupAttachments", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &NetworkOptionGroupAttachmentLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/network_option_group_attachments/:id
// Delete an existing NetworkOptionGroupAttachment.
func (loc *NetworkOptionGroupAttachmentLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("NetworkOptionGroupAttachments", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/network_option_group_attachments
// List NetworkOptionGroupAttachments in this account.
// -- Optional parameters:
// filter
// view
func (loc *NetworkOptionGroupAttachmentLocator) Index(options rsapi.ApiParams) ([]*NetworkOptionGroupAttachment, error) {
	var res []*NetworkOptionGroupAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("NetworkOptionGroupAttachments", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/network_option_group_attachments/:id
// Show information about a single NetworkOptionGroupAttachment.
// -- Optional parameters:
// view
func (loc *NetworkOptionGroupAttachmentLocator) Show(options rsapi.ApiParams) (*NetworkOptionGroupAttachment, error) {
	var res *NetworkOptionGroupAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("NetworkOptionGroupAttachments", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/network_option_group_attachments/:id
// Update an existing NetworkOptionGroupAttachment.
func (loc *NetworkOptionGroupAttachmentLocator) Update(networkOptionGroupAttachment *NetworkOptionGroupAttachmentParam2) error {
	if networkOptionGroupAttachment == nil {
		return fmt.Errorf("networkOptionGroupAttachment is required")
	}
	var params = rsapi.ApiParams{
		"network_option_group_attachment": networkOptionGroupAttachment,
	}
	var uri, err = loc.Url("NetworkOptionGroupAttachments", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Oauth2 ******/

// Note that all API calls irrespective of the resource it is acting on, should pass a header
// "X-Api-Version" with the value "1.5".
// This is an OAuth 2.0 token endpoint that can be used to perform token-refresh operations and obtain
// a bearer access token, which can be used in lieu of a session cookie when interacting with API
// resources.
// This is not an API resource; in order to maintain compatibility with OAuth 2.0, it does not conform
// to the conventions established by other RightScale API resources. However, an API-version header is
// still required when interacting with the OAuth endpoint.
// OAuth 2.0 endpoints always use the POST verb, accept a www-urlencoded request body (similarly to a
// browser form submission) and the OAuth action is indicated by the "grant_type" parameter. This
// endpoint supports the following OAuth 2.0 operations:
// refresh_token - for end-user login using a previously-negotiated OAuth grant
// client_credentials - for instance login using API credentials transmitted via user-data
// RightScale's OAuth implementation has two proprietary aspects that you should be aware of:
// clients MUST transmit an X-Api-Version header with every OAuth request
// clients MAY transmit an account_id parameter as part of their POST form data
// If you choose to post an account_id, then the API may respond with a 301 redirect if your account
// is hosted in another RightScale cluster. If you omit this parameter and your account is hosted
// elsewhere, then you will simply receive a 400 Bad Request (because your grant is not known to
// this cluster).
// For more information on how to use OAuth 2.0 with RightScale, refer to the following:
// http://support.rightscale.com/12-Guides/03-RightScale_API/OAuth
// http://tools.ietf.org/html/draft-ietf-oauth-v2-23

//===== Locator

// Oauth2 resource locator, exposes resource actions.
type Oauth2Locator struct {
	UrlResolver
	api *Api15
}

// Oauth2 resource locator factory
func (api *Api15) Oauth2Locator(href string) *Oauth2Locator {
	return &Oauth2Locator{UrlResolver(href), api}
}

//===== Actions

// POST /api/oauth2/
// Perform an OAuth 2.0 token_refresh operation to obtain an access token that
// can be used in lieu of an API session cookie. (In other words, creates a
// session using OAuth 2.0).
// Note that an API-Version header is required with your request, and that the server
// may respond with a 301 Moved Permanently if you include an account_id parameter and
// your account is hosted in another RightScale cluster.
// The request parameters and response format are all as per the OAuth 2.0
// Internet Draft standard v23. In brief:
// Successful responses include an access token, an expires-in timestamp, and a token type
// The token type is always "bearer"
// To use a bearer token, include header "Authorization: Bearer " with your API requests
// The client must refresh the access token before it expires
// # Example Request using Curl (with prettified response):
// curl -i -H X-API-Version:1.5 -x POST https://my.rightscale.com/api/oauth2 -d "grant_type=refresh_token" -d "refresh_token=abcd1234deadbeef"
// {
// "access_token": "xyzzy",
// "expires_in":   3600,
// "token_type":   "bearer"
// }
// grantType: Type of grant.
// -- Optional parameters:
// account_id: The client's account ID (only needed for instance agent clients).
// client_id: The client ID (only needed for confidential clients).
// client_secret: The client secret (only needed for confidential clients).
// r_s_version: The RightAgent protocol version the client conforms to (only needed for instance agent clients).
// refresh_token: The refresh token obtained from OAuth grant.
// right_link_version: The RightLink gem version the client conforms to (only needed for instance agent clients).
func (loc *Oauth2Locator) Create(grantType string, options rsapi.ApiParams) (map[string]interface{}, error) {
	var res map[string]interface{}
	if grantType == "" {
		return res, fmt.Errorf("grantType is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"grant_type": grantType,
	}, options)
	var uri, err = loc.Url("Oauth2", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

/******  Permission ******/

type Permission struct {
	Actions   []string            `json:"actions,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	RoleTitle string              `json:"role_title,omitempty"`
}

//===== Locator

// Permission resource locator, exposes resource actions.
type PermissionLocator struct {
	UrlResolver
	api *Api15
}

// Permission resource locator factory
func (api *Api15) PermissionLocator(href string) *PermissionLocator {
	return &PermissionLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/permissions
// Create a permission, thereby granting some user a particular role
// with respect to the current account.
// The 'observer' role has a special status; it must be granted before
// a user is eligible for any other permission in a given account.
// When provisioning users, always create the observer permission FIRST;
// creating any other permission before it will result in an error.
// For more information about the roles available and the privileges
// they confer, please refer to the following page of the RightScale
// support portal:
// http://support.rightscale.com/15-References/Lists/ListofUser_Roles
func (loc *PermissionLocator) Create(permission *PermissionParam) (*PermissionLocator, error) {
	var res *PermissionLocator
	if permission == nil {
		return res, fmt.Errorf("permission is required")
	}
	var params = rsapi.ApiParams{
		"permission": permission,
	}
	var uri, err = loc.Url("Permissions", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &PermissionLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/permissions/:id
// Destroy a permission, thereby revoking a user's role with respect
// to the current account.
// The 'observer' role has a special status; it cannot be revoked if
// a user has any other roles, because other roles become useless without
// being able to read data pertaining to the account.
// When deprovisioning user, always destroy the observer permission LAST;
// destroying it while the user has other permissions will result in an error.
func (loc *PermissionLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Permissions", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/permissions
// List all permissions for all users of the current acount.
// -- Optional parameters:
// filter
func (loc *PermissionLocator) Index(options rsapi.ApiParams) ([]*Permission, error) {
	var res []*Permission
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Permissions", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/permissions/:id
// Show information about a single permission.
func (loc *PermissionLocator) Show() (*Permission, error) {
	var res *Permission
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Permissions", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  PlacementGroup ******/

type PlacementGroup struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// PlacementGroup resource locator, exposes resource actions.
type PlacementGroupLocator struct {
	UrlResolver
	api *Api15
}

// PlacementGroup resource locator factory
func (api *Api15) PlacementGroupLocator(href string) *PlacementGroupLocator {
	return &PlacementGroupLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/placement_groups
// Creates a PlacementGroup.
func (loc *PlacementGroupLocator) Create(placementGroup *PlacementGroupParam) (*PlacementGroupLocator, error) {
	var res *PlacementGroupLocator
	if placementGroup == nil {
		return res, fmt.Errorf("placementGroup is required")
	}
	var params = rsapi.ApiParams{
		"placement_group": placementGroup,
	}
	var uri, err = loc.Url("PlacementGroups", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &PlacementGroupLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/placement_groups/:id
// Destroys a PlacementGroup.
func (loc *PlacementGroupLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("PlacementGroups", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/placement_groups
// Lists all PlacementGroups in an account.
// -- Optional parameters:
// filter
// view
func (loc *PlacementGroupLocator) Index(options rsapi.ApiParams) ([]*PlacementGroup, error) {
	var res []*PlacementGroup
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("PlacementGroups", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/placement_groups/:id
// Shows information about a single PlacementGroup.
// -- Optional parameters:
// view
func (loc *PlacementGroupLocator) Show(options rsapi.ApiParams) (*PlacementGroup, error) {
	var res *PlacementGroup
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("PlacementGroups", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Preference ******/

// A Preference is a user and account-specific setting. Preferences are used in many part of the RightScale platform and can be used for custom purposes if desired.
type Preference struct {
	Actions   []string            `json:"actions,omitempty"`
	Contents  string              `json:"contents,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// Preference resource locator, exposes resource actions.
type PreferenceLocator struct {
	UrlResolver
	api *Api15
}

// Preference resource locator factory
func (api *Api15) PreferenceLocator(href string) *PreferenceLocator {
	return &PreferenceLocator{UrlResolver(href), api}
}

//===== Actions

// DELETE /api/preferences/:id
// Deletes the given preference.
func (loc *PreferenceLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Preferences", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/preferences
// Lists all preferences.
// -- Optional parameters:
// filter
func (loc *PreferenceLocator) Index(options rsapi.ApiParams) ([]*Preference, error) {
	var res []*Preference
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Preferences", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/preferences/:id
// Shows a single preference.
func (loc *PreferenceLocator) Show() (*Preference, error) {
	var res *Preference
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Preferences", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/preferences/:id
// If 'id' is known, updates preference with given contents.
// Otherwise, creates new preference.
// Note: If create, will return '201 Created' and the location of the new preference.
func (loc *PreferenceLocator) Update(preference *PreferenceParam) error {
	if preference == nil {
		return fmt.Errorf("preference is required")
	}
	var params = rsapi.ApiParams{
		"preference": preference,
	}
	var uri, err = loc.Url("Preferences", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Publication ******/

// A Publication is a revisioned component shared with a set of Account Groups.
// If shared with your account, it can be imported in to your account.
type Publication struct {
	Actions       []string            `json:"actions,omitempty"`
	CommitMessage string              `json:"commit_message,omitempty"`
	ContentType   string              `json:"content_type,omitempty"`
	CreatedAt     time.Time           `json:"created_at,omitempty"`
	Description   string              `json:"description,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	Name          string              `json:"name,omitempty"`
	Publisher     string              `json:"publisher,omitempty"`
	Revision      string              `json:"revision,omitempty"`
	RevisionNotes string              `json:"revision_notes,omitempty"`
	UpdatedAt     time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// Publication resource locator, exposes resource actions.
type PublicationLocator struct {
	UrlResolver
	api *Api15
}

// Publication resource locator factory
func (api *Api15) PublicationLocator(href string) *PublicationLocator {
	return &PublicationLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/publications/:id/import
// Imports the given publication and its subordinates to this account.
// Only non-HEAD revisions that are shared with the account can be imported.
func (loc *PublicationLocator) Import() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Publications", "import")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/publications
// Lists the publications available to this account. Only non-HEAD revisions are possible.
// -- Optional parameters:
// filter
// view
func (loc *PublicationLocator) Index(options rsapi.ApiParams) ([]*Publication, error) {
	var res []*Publication
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Publications", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/publications/:id
// Show information about a single publication. Only non-HEAD revisions are possible.
// -- Optional parameters:
// view
func (loc *PublicationLocator) Show(options rsapi.ApiParams) (*Publication, error) {
	var res *Publication
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Publications", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  PublicationLineage ******/

// A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.
// It is shared among all revisions of a Publication within the marketplace.
// Publication Lineages are different than lineages that exist within an account.
type PublicationLineage struct {
	Actions          []string            `json:"actions,omitempty"`
	CommentsEmailed  bool                `json:"comments_emailed,omitempty"`
	CommentsEnabled  bool                `json:"comments_enabled,omitempty"`
	CreatedAt        time.Time           `json:"created_at,omitempty"`
	Links            []map[string]string `json:"links,omitempty"`
	LongDescription  string              `json:"long_description,omitempty"`
	Name             string              `json:"name,omitempty"`
	ShortDescription string              `json:"short_description,omitempty"`
	UpdatedAt        time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// PublicationLineage resource locator, exposes resource actions.
type PublicationLineageLocator struct {
	UrlResolver
	api *Api15
}

// PublicationLineage resource locator factory
func (api *Api15) PublicationLineageLocator(href string) *PublicationLineageLocator {
	return &PublicationLineageLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/publication_lineages/:id
// Show information about a single publication lineage. Only non-HEAD revisions are possible.
// -- Optional parameters:
// view
func (loc *PublicationLineageLocator) Show(options rsapi.ApiParams) (*PublicationLineage, error) {
	var res *PublicationLineage
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("PublicationLineages", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  RecurringVolumeAttachment ******/

// A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a Server/ServerArray the next time an instance is launched.
type RecurringVolumeAttachment struct {
	Actions      []string            `json:"actions,omitempty"`
	CreatedAt    time.Time           `json:"created_at,omitempty"`
	Device       string              `json:"device,omitempty"`
	DeviceId     string              `json:"device_id,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	Name         string              `json:"name,omitempty"`
	RunnableType string              `json:"runnable_type,omitempty"`
	Size         string              `json:"size,omitempty"`
	Status       string              `json:"status,omitempty"`
	StorageType  string              `json:"storage_type,omitempty"`
	UpdatedAt    time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// RecurringVolumeAttachment resource locator, exposes resource actions.
type RecurringVolumeAttachmentLocator struct {
	UrlResolver
	api *Api15
}

// RecurringVolumeAttachment resource locator factory
func (api *Api15) RecurringVolumeAttachmentLocator(href string) *RecurringVolumeAttachmentLocator {
	return &RecurringVolumeAttachmentLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/recurring_volume_attachments
// POST /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments
// POST /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments
// Creates a new recurring volume attachment.
func (loc *RecurringVolumeAttachmentLocator) Create(recurringVolumeAttachment *RecurringVolumeAttachmentParam) (*RecurringVolumeAttachmentLocator, error) {
	var res *RecurringVolumeAttachmentLocator
	if recurringVolumeAttachment == nil {
		return res, fmt.Errorf("recurringVolumeAttachment is required")
	}
	var params = rsapi.ApiParams{
		"recurring_volume_attachment": recurringVolumeAttachment,
	}
	var uri, err = loc.Url("RecurringVolumeAttachments", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &RecurringVolumeAttachmentLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/recurring_volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments/:id
// Deletes a given recurring volume attachment.
func (loc *RecurringVolumeAttachmentLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("RecurringVolumeAttachments", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/recurring_volume_attachments
// GET /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments
// GET /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments
// Lists all recurring volume attachments.
// -- Optional parameters:
// filter
// view
func (loc *RecurringVolumeAttachmentLocator) Index(options rsapi.ApiParams) ([]*RecurringVolumeAttachment, error) {
	var res []*RecurringVolumeAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RecurringVolumeAttachments", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/recurring_volume_attachments/:id
// GET /api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments/:id
// GET /api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments/:id
// Displays information about a single recurring volume attachment.
// -- Optional parameters:
// view
func (loc *RecurringVolumeAttachmentLocator) Show(options rsapi.ApiParams) (*RecurringVolumeAttachment, error) {
	var res *RecurringVolumeAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RecurringVolumeAttachments", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Repository ******/

// A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.
// RightScale currently supports the following types of repositores: git, svn, and URLs of compressed files (tar, tgz, gzip).
type Repository struct {
	Actions         []string            `json:"actions,omitempty"`
	AssetCounts     int                 `json:"asset_counts,omitempty"`
	AssetPaths      []string            `json:"asset_paths,omitempty"`
	CommitReference string              `json:"commit_reference,omitempty"`
	CreatedAt       time.Time           `json:"created_at,omitempty"`
	Credentials     map[string]string   `json:"credentials,omitempty"`
	Description     string              `json:"description,omitempty"`
	FetchStatus     string              `json:"fetch_status,omitempty"`
	Id              string              `json:"id,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	ReadOnly        bool                `json:"read_only,omitempty"`
	Source          string              `json:"source,omitempty"`
	SourceType      string              `json:"source_type,omitempty"`
	UpdatedAt       time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// Repository resource locator, exposes resource actions.
type RepositoryLocator struct {
	UrlResolver
	api *Api15
}

// Repository resource locator factory
func (api *Api15) RepositoryLocator(href string) *RepositoryLocator {
	return &RepositoryLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/repositories/:id/cookbook_import
// Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.
// assetHrefs: Hrefs of the assets that should be imported.
// -- Optional parameters:
// follow: A flag indicating whether imported cookbooks should be followed.
// namespace: The namespace to import into.
// repository_commit_reference: Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.
// with_dependencies: A flag indicating whether dependencies should automatically be imported.
func (loc *RepositoryLocator) CookbookImport(assetHrefs []string, options rsapi.ApiParams) error {
	if len(assetHrefs) == 0 {
		return fmt.Errorf("assetHrefs is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"asset_hrefs": assetHrefs,
	}, options)
	var uri, err = loc.Url("Repositories", "cookbook_import")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/repositories/:id/cookbook_import_preview
// Retrieves a preview of the effects of a Cookbook import.
// NOTE: This action is for RightScale internal use only. The response is
// free-form JSON with no associated mediatype.
// DO NOT USE, THIS ACTION IS SUBJECT TO CHANGE AT ANYTIME.
// assetHrefs: Hrefs of the assets that should be imported.
// namespace: The namespace to import into.
func (loc *RepositoryLocator) CookbookImportPreview(assetHrefs []string, namespace string) ([]*map[string]string, error) {
	var res []*map[string]string
	if len(assetHrefs) == 0 {
		return res, fmt.Errorf("assetHrefs is required")
	}
	if namespace == "" {
		return res, fmt.Errorf("namespace is required")
	}
	var params = rsapi.ApiParams{
		"asset_hrefs": assetHrefs,
		"namespace":   namespace,
	}
	var uri, err = loc.Url("Repositories", "cookbook_import_preview")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/repositories
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
func (loc *RepositoryLocator) Create(repository *RepositoryParam) (*RepositoryLocator, error) {
	var res *RepositoryLocator
	if repository == nil {
		return res, fmt.Errorf("repository is required")
	}
	var params = rsapi.ApiParams{
		"repository": repository,
	}
	var uri, err = loc.Url("Repositories", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &RepositoryLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/repositories/:id
// Deletes the specified Repositories.
func (loc *RepositoryLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Repositories", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/repositories
// Lists all Repositories for this Account.
// -- Optional parameters:
// filter
// view
func (loc *RepositoryLocator) Index(options rsapi.ApiParams) ([]*Repository, error) {
	var res []*Repository
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Repositories", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/repositories/:id/refetch
// Refetches all RepositoryAssets associated with the Repository.
// Note that a refetch simply updates RightScale's view of the contents of the repository.
// You must perform an import to use the assets in your design objects (or use the auto import parameter).
// -- Optional parameters:
// auto_import: Whether cookbooks should automatically be imported after repositories are fetched.
func (loc *RepositoryLocator) Refetch(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Repositories", "refetch")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/repositories/resolve
// Show a list of repositories that have imported cookbooks with the given names.
// This operation returns a list of repositories that would later satisfy a call
// to the swap_repository
// action on a ServerTemplate.
// -- Optional parameters:
// imported_cookbook_name: A list of cookbook names that were imported by the repository.
func (loc *RepositoryLocator) Resolve(options rsapi.ApiParams) ([]*Repository, error) {
	var res []*Repository
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Repositories", "resolve")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/repositories/:id
// Shows a specified Repository.
// -- Optional parameters:
// view
func (loc *RepositoryLocator) Show(options rsapi.ApiParams) (*Repository, error) {
	var res *Repository
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Repositories", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/repositories/:id
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
func (loc *RepositoryLocator) Update(repository *RepositoryParam2) error {
	if repository == nil {
		return fmt.Errorf("repository is required")
	}
	var params = rsapi.ApiParams{
		"repository": repository,
	}
	var uri, err = loc.Url("Repositories", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  RepositoryAsset ******/

// A RepositoryAsset represents an item discovered in a Repository. These assets represent only a view of the Repository
// the last time it was scraped. In order to use these assets, you must import them into your account.
type RepositoryAsset struct {
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Id          string              `json:"id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Version     string              `json:"version,omitempty"`
}

//===== Locator

// RepositoryAsset resource locator, exposes resource actions.
type RepositoryAssetLocator struct {
	UrlResolver
	api *Api15
}

// RepositoryAsset resource locator factory
func (api *Api15) RepositoryAssetLocator(href string) *RepositoryAssetLocator {
	return &RepositoryAssetLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/repositories/:repository_id/repository_assets
// List a repository's current assets.
// Repository assests are the cookbook details that were scraped from a
// given repository.
// -- Optional parameters:
// view
func (loc *RepositoryAssetLocator) Index(options rsapi.ApiParams) ([]*RepositoryAsset, error) {
	var res []*RepositoryAsset
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RepositoryAssets", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/repositories/:repository_id/repository_assets/:id
// Show information about a single asset.
// A repository assest are the cookbook details that were scraped from a
// repository.
// -- Optional parameters:
// view
func (loc *RepositoryAssetLocator) Show(options rsapi.ApiParams) (*RepositoryAsset, error) {
	var res *RepositoryAsset
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RepositoryAssets", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  RightScript ******/

// A RightScript is an executable piece of code that can be run on a server
// during the boot, operational, or decommission phases.
// All revisions of
// a RightScript belong to a RightScript lineage that is exposed by the
// "lineage" attribute (NOTE: This attribute is merely a string to locate
// all revisions of a RightScript and NOT a working URL).
type RightScript struct {
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Id          string              `json:"id,omitempty"`
	Lineage     string              `json:"lineage,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Revision    string              `json:"revision,omitempty"`
	Source      string              `json:"source,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// RightScript resource locator, exposes resource actions.
type RightScriptLocator struct {
	UrlResolver
	api *Api15
}

// RightScript resource locator factory
func (api *Api15) RightScriptLocator(href string) *RightScriptLocator {
	return &RightScriptLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/right_scripts/:id/commit
// Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.
func (loc *RightScriptLocator) Commit(rightScript *RightScriptParam) error {
	if rightScript == nil {
		return fmt.Errorf("rightScript is required")
	}
	var params = rsapi.ApiParams{
		"right_script": rightScript,
	}
	var uri, err = loc.Url("RightScripts", "commit")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/right_scripts
// Lists RightScripts.
// -- Optional parameters:
// filter
// latest_only: Whether or not to return only the latest version for each lineage.
// view
func (loc *RightScriptLocator) Index(options rsapi.ApiParams) ([]*RightScript, error) {
	var res []*RightScript
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RightScripts", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/right_scripts/:id
// Displays information about a single RightScript.
func (loc *RightScriptLocator) Show() (*RightScript, error) {
	var res *RightScript
	var params = map[string]interface{}{}
	var uri, err = loc.Url("RightScripts", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// GET /api/right_scripts/:id/source
// Returns the script source for a RightScript
func (loc *RightScriptLocator) ShowSource() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("RightScripts", "show_source")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// PUT /api/right_scripts/:id
// Updates RightScript name/description
func (loc *RightScriptLocator) Update(rightScript *RightScriptParam2) error {
	if rightScript == nil {
		return fmt.Errorf("rightScript is required")
	}
	var params = rsapi.ApiParams{
		"right_script": rightScript,
	}
	var uri, err = loc.Url("RightScripts", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// PUT /api/right_scripts/:id/source
// Updates the source of the given RightScript
func (loc *RightScriptLocator) UpdateSource() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("RightScripts", "update_source")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Route ******/

// A Route defines how networking traffic should be routed from one
// destination to another. See nexthoptype for available endpoint targets.
type Route struct {
	CreatedAt            time.Time           `json:"created_at,omitempty"`
	Description          string              `json:"description,omitempty"`
	DestinationCidrBlock string              `json:"destination_cidr_block,omitempty"`
	Links                []map[string]string `json:"links,omitempty"`
	NextHopIp            string              `json:"next_hop_ip,omitempty"`
	NextHopType          string              `json:"next_hop_type,omitempty"`
	ResourceUid          string              `json:"resource_uid,omitempty"`
	State                string              `json:"state,omitempty"`
	UpdatedAt            time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// Route resource locator, exposes resource actions.
type RouteLocator struct {
	UrlResolver
	api *Api15
}

// Route resource locator factory
func (api *Api15) RouteLocator(href string) *RouteLocator {
	return &RouteLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/routes
// POST /api/route_tables/:route_table_id/routes
// Create a new Route.
func (loc *RouteLocator) Create(route *RouteParam) (*RouteLocator, error) {
	var res *RouteLocator
	if route == nil {
		return res, fmt.Errorf("route is required")
	}
	var params = rsapi.ApiParams{
		"route": route,
	}
	var uri, err = loc.Url("Routes", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &RouteLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/routes/:id
// DELETE /api/route_tables/:route_table_id/routes/:id
// Delete an existing Route.
func (loc *RouteLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Routes", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/routes
// GET /api/route_tables/:route_table_id/routes
// List Routes available in this account.
// -- Optional parameters:
// filter
func (loc *RouteLocator) Index(options rsapi.ApiParams) ([]*Route, error) {
	var res []*Route
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Routes", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/routes/:id
// GET /api/route_tables/:route_table_id/routes/:id
// Show information about a single Route.
func (loc *RouteLocator) Show() (*Route, error) {
	var res *Route
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Routes", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/routes/:id
// PUT /api/route_tables/:route_table_id/routes/:id
// Update an existing Route.
func (loc *RouteLocator) Update(route *RouteParam2) error {
	if route == nil {
		return fmt.Errorf("route is required")
	}
	var params = rsapi.ApiParams{
		"route": route,
	}
	var uri, err = loc.Url("Routes", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  RouteTable ******/

// Grouped listing of Routes
type RouteTable struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Routes      []Route             `json:"routes,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// RouteTable resource locator, exposes resource actions.
type RouteTableLocator struct {
	UrlResolver
	api *Api15
}

// RouteTable resource locator factory
func (api *Api15) RouteTableLocator(href string) *RouteTableLocator {
	return &RouteTableLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/route_tables
// Create a new RouteTable.
func (loc *RouteTableLocator) Create(routeTable *RouteTableParam) (*RouteTableLocator, error) {
	var res *RouteTableLocator
	if routeTable == nil {
		return res, fmt.Errorf("routeTable is required")
	}
	var params = rsapi.ApiParams{
		"route_table": routeTable,
	}
	var uri, err = loc.Url("RouteTables", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &RouteTableLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/route_tables/:id
// Delete an existing RouteTable.
func (loc *RouteTableLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("RouteTables", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/route_tables
// List RouteTables available in this account.
// -- Optional parameters:
// filter
// view
func (loc *RouteTableLocator) Index(options rsapi.ApiParams) ([]*RouteTable, error) {
	var res []*RouteTable
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RouteTables", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/route_tables/:id
// Show information about a single RouteTable.
// -- Optional parameters:
// view
func (loc *RouteTableLocator) Show(options rsapi.ApiParams) (*RouteTable, error) {
	var res *RouteTable
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RouteTables", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/route_tables/:id
// Update an existing RouteTable.
func (loc *RouteTableLocator) Update(routeTable *RouteTableParam2) error {
	if routeTable == nil {
		return fmt.Errorf("routeTable is required")
	}
	var params = rsapi.ApiParams{
		"route_table": routeTable,
	}
	var uri, err = loc.Url("RouteTables", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  RunnableBinding ******/

// A RunnableBinding represents an item in a runlist of a ServerTemplate. These items could be
// RightScript or Chef recipes, and could be associated with any one of the three runlists of a
// ServerTemplate (boot, operational, decommission).
type RunnableBinding struct {
	Actions     []string            `json:"actions,omitempty"`
	Id          string              `json:"id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Position    int                 `json:"position,omitempty"`
	Recipe      string              `json:"recipe,omitempty"`
	RightScript string              `json:"right_script,omitempty"`
	Sequence    string              `json:"sequence,omitempty"`
}

//===== Locator

// RunnableBinding resource locator, exposes resource actions.
type RunnableBindingLocator struct {
	UrlResolver
	api *Api15
}

// RunnableBinding resource locator factory
func (api *Api15) RunnableBindingLocator(href string) *RunnableBindingLocator {
	return &RunnableBindingLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/server_templates/:server_template_id/runnable_bindings
// Bind an executable to the given ServerTemplate.
// An executable may be either a RightScript or Chef Recipe.
// The resource must be editable.
func (loc *RunnableBindingLocator) Create(runnableBinding *RunnableBindingParam) (*RunnableBindingLocator, error) {
	var res *RunnableBindingLocator
	if runnableBinding == nil {
		return res, fmt.Errorf("runnableBinding is required")
	}
	var params = rsapi.ApiParams{
		"runnable_binding": runnableBinding,
	}
	var uri, err = loc.Url("RunnableBindings", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &RunnableBindingLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/server_templates/:server_template_id/runnable_bindings/:id
// Unbind an executable from the given resource.
// The resource must be editable.
func (loc *RunnableBindingLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("RunnableBindings", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_templates/:server_template_id/runnable_bindings
// Lists the executables bound to the ServerTemplate.
// An excutable may be either a RightScript or Chef Recipe.
// -- Optional parameters:
// view
func (loc *RunnableBindingLocator) Index(options rsapi.ApiParams) ([]*RunnableBinding, error) {
	var res []*RunnableBinding
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RunnableBindings", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// PUT /api/server_templates/:server_template_id/runnable_bindings/multi_update
// Update attributes for multiple bound executables.
// The resource must be editable.
func (loc *RunnableBindingLocator) MultiUpdate(runnableBindings []*RunnableBindings) error {
	if len(runnableBindings) == 0 {
		return fmt.Errorf("runnableBindings is required")
	}
	var params = rsapi.ApiParams{
		"runnable_bindings": runnableBindings,
	}
	var uri, err = loc.Url("RunnableBindings", "multi_update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_templates/:server_template_id/runnable_bindings/:id
// Show information about a single executable binding.
// An excutable may be either a RightScript or Chef Recipe.
// -- Optional parameters:
// view
func (loc *RunnableBindingLocator) Show(options rsapi.ApiParams) (*RunnableBinding, error) {
	var res *RunnableBinding
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("RunnableBindings", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  SecurityGroup ******/

// Security Groups represent network security profiles that contain lists of firewall rules for different ports and source IP addresses, as well as
// trust relationships amongst different security groups.
type SecurityGroup struct {
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Href        string              `json:"href,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

//===== Locator

// SecurityGroup resource locator, exposes resource actions.
type SecurityGroupLocator struct {
	UrlResolver
	api *Api15
}

// SecurityGroup resource locator factory
func (api *Api15) SecurityGroupLocator(href string) *SecurityGroupLocator {
	return &SecurityGroupLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/security_groups
// Create a security group.
func (loc *SecurityGroupLocator) Create(securityGroup *SecurityGroupParam) (*SecurityGroupLocator, error) {
	var res *SecurityGroupLocator
	if securityGroup == nil {
		return res, fmt.Errorf("securityGroup is required")
	}
	var params = rsapi.ApiParams{
		"security_group": securityGroup,
	}
	var uri, err = loc.Url("SecurityGroups", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &SecurityGroupLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/security_groups/:id
// Delete security group(s)
func (loc *SecurityGroupLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("SecurityGroups", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/security_groups
// Lists Security Groups.
// -- Optional parameters:
// filter
// view
func (loc *SecurityGroupLocator) Index(options rsapi.ApiParams) ([]*SecurityGroup, error) {
	var res []*SecurityGroup
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("SecurityGroups", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/security_groups/:id
// Displays information about a single Security Group.
// -- Optional parameters:
// view
func (loc *SecurityGroupLocator) Show(options rsapi.ApiParams) (*SecurityGroup, error) {
	var res *SecurityGroup
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("SecurityGroups", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  SecurityGroupRule ******/

type SecurityGroupRule struct {
	Actions     []string            `json:"actions,omitempty"`
	CidrIps     []string            `json:"cidr_ips,omitempty"`
	Description string              `json:"description,omitempty"`
	Direction   string              `json:"direction,omitempty"`
	EndPort     string              `json:"end_port,omitempty"`
	GroupName   string              `json:"group_name,omitempty"`
	GroupOwner  string              `json:"group_owner,omitempty"`
	GroupUid    string              `json:"group_uid,omitempty"`
	Href        string              `json:"href,omitempty"`
	IcmpCode    string              `json:"icmp_code,omitempty"`
	IcmpType    string              `json:"icmp_type,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Protocol    string              `json:"protocol,omitempty"`
	SourceType  string              `json:"source_type,omitempty"`
	StartPort   string              `json:"start_port,omitempty"`
}

//===== Locator

// SecurityGroupRule resource locator, exposes resource actions.
type SecurityGroupRuleLocator struct {
	UrlResolver
	api *Api15
}

// SecurityGroupRule resource locator factory
func (api *Api15) SecurityGroupRuleLocator(href string) *SecurityGroupRuleLocator {
	return &SecurityGroupRuleLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/security_group_rules
// POST /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules
// Create a security group rule for a security group.
// The following flavors are supported:
// group-based TCP/UDP
// group-based ICMP
// CIDR-based TCP/UDP
// CIDR-based ICMP
func (loc *SecurityGroupRuleLocator) Create(securityGroupRule *SecurityGroupRuleParam) (*SecurityGroupRuleLocator, error) {
	var res *SecurityGroupRuleLocator
	if securityGroupRule == nil {
		return res, fmt.Errorf("securityGroupRule is required")
	}
	var params = rsapi.ApiParams{
		"security_group_rule": securityGroupRule,
	}
	var uri, err = loc.Url("SecurityGroupRules", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &SecurityGroupRuleLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/security_group_rules/:id
// DELETE /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id
// Delete security group rule(s)
func (loc *SecurityGroupRuleLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("SecurityGroupRules", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/security_group_rules
// GET /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules
// Lists SecurityGroupRules.
// -- Optional parameters:
// view
func (loc *SecurityGroupRuleLocator) Index(options rsapi.ApiParams) ([]*SecurityGroupRule, error) {
	var res []*SecurityGroupRule
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("SecurityGroupRules", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/security_group_rules/:id
// GET /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id
// Displays information about a single SecurityGroupRule.
// -- Optional parameters:
// view
func (loc *SecurityGroupRuleLocator) Show(options rsapi.ApiParams) (*SecurityGroupRule, error) {
	var res *SecurityGroupRule
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("SecurityGroupRules", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/security_group_rules/:id
// PUT /api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id

func (loc *SecurityGroupRuleLocator) Update(securityGroupRule *SecurityGroupRuleParam2) error {
	if securityGroupRule == nil {
		return fmt.Errorf("securityGroupRule is required")
	}
	var params = rsapi.ApiParams{
		"security_group_rule": securityGroupRule,
	}
	var uri, err = loc.Url("SecurityGroupRules", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
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
	Actions         []string            `json:"actions,omitempty"`
	CreatedAt       time.Time           `json:"created_at,omitempty"`
	CurrentInstance Instance            `json:"current_instance,omitempty"`
	Description     string              `json:"description,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	NextInstance    *Instance           `json:"next_instance,omitempty"`
	Optimized       bool                `json:"optimized,omitempty"`
	State           string              `json:"state,omitempty"`
	UpdatedAt       time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// Server resource locator, exposes resource actions.
type ServerLocator struct {
	UrlResolver
	api *Api15
}

// Server resource locator factory
func (api *Api15) ServerLocator(href string) *ServerLocator {
	return &ServerLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/servers/:id/clone
// Clones a given server.
func (loc *ServerLocator) Clone() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Servers", "clone")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/servers
// POST /api/deployments/:deployment_id/servers
// Creates a new server, and configures its corresponding "next" instance with the received parameters.
func (loc *ServerLocator) Create(server *ServerParam) (*ServerLocator, error) {
	var res *ServerLocator
	if server == nil {
		return res, fmt.Errorf("server is required")
	}
	var params = rsapi.ApiParams{
		"server": server,
	}
	var uri, err = loc.Url("Servers", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ServerLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/servers/:id
// DELETE /api/deployments/:deployment_id/servers/:id
// Deletes a given server.
func (loc *ServerLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Servers", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/servers
// GET /api/deployments/:deployment_id/servers
// Lists servers.
// By using the available filters, it is possible to retrieve servers that have common characteristics.
// For example, one can list:
// servers that have names that contain "app_server"
// all servers of a given deployment
// For more filters, please see the 'index' action on 'Instances' resource as most of the attributes belong to
// a 'current_instance' than to a server.
// -- Optional parameters:
// filter
// view
func (loc *ServerLocator) Index(options rsapi.ApiParams) ([]*Server, error) {
	var res []*Server
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Servers", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/servers/:id/launch
// Launches the "next" instance of this server. This function is equivalent to invoking the launch action on the
// URL of this servers next_instance. See Instances#launch for details.
func (loc *ServerLocator) Launch() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Servers", "launch")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/servers/:id
// GET /api/deployments/:deployment_id/servers/:id
// Shows the information of a single server.
// -- Optional parameters:
// view
func (loc *ServerLocator) Show(options rsapi.ApiParams) (*Server, error) {
	var res *Server
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Servers", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// POST /api/servers/:id/teminate
// Terminates the current instance of this server. This function is equivalent to invoking the terminate action on the
// URL of this servers current_instance. See Instances#terminate for details.
func (loc *ServerLocator) Terminate() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Servers", "terminate")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// PUT /api/servers/:id
// PUT /api/deployments/:deployment_id/servers/:id
// Updates attributes of a single server.
func (loc *ServerLocator) Update(server *ServerParam2) error {
	if server == nil {
		return fmt.Errorf("server is required")
	}
	var params = rsapi.ApiParams{
		"server": server,
	}
	var uri, err = loc.Url("Servers", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/servers/wrap_instance
// POST /api/deployments/:deployment_id/servers/wrap_instance
// Wrap an existing instance and set current instance for new server
func (loc *ServerLocator) WrapInstance(server *ServerParam2) error {
	if server == nil {
		return fmt.Errorf("server is required")
	}
	var params = rsapi.ApiParams{
		"server": server,
	}
	var uri, err = loc.Url("Servers", "wrap_instance")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
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
	Actions          []string               `json:"actions,omitempty"`
	ArrayType        string                 `json:"array_type,omitempty"`
	CurrentInstances []Instance             `json:"current_instances,omitempty"`
	DatacenterPolicy string                 `json:"datacenter_policy,omitempty"`
	Description      string                 `json:"description,omitempty"`
	ElasticityParams map[string]interface{} `json:"elasticity_params,omitempty"`
	InstancesCount   int                    `json:"instances_count,omitempty"`
	Links            []map[string]string    `json:"links,omitempty"`
	Name             string                 `json:"name,omitempty"`
	NextInstance     *Instance              `json:"next_instance,omitempty"`
	Optimized        bool                   `json:"optimized,omitempty"`
	State            string                 `json:"state,omitempty"`
}

//===== Locator

// ServerArray resource locator, exposes resource actions.
type ServerArrayLocator struct {
	UrlResolver
	api *Api15
}

// ServerArray resource locator factory
func (api *Api15) ServerArrayLocator(href string) *ServerArrayLocator {
	return &ServerArrayLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/server_arrays/:id/clone
// Clones a given server array.
func (loc *ServerArrayLocator) Clone() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerArrays", "clone")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_arrays
// POST /api/deployments/:deployment_id/server_arrays
// Creates a new server array, and configures its corresponding "next" instance with the received parameters.
func (loc *ServerArrayLocator) Create(serverArray *ServerArrayParam) (*ServerArrayLocator, error) {
	var res *ServerArrayLocator
	if serverArray == nil {
		return res, fmt.Errorf("serverArray is required")
	}
	var params = rsapi.ApiParams{
		"server_array": serverArray,
	}
	var uri, err = loc.Url("ServerArrays", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ServerArrayLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/server_arrays/:id/current_instances
// List the running instances belonging to the server array. See Instances#index for details.
// This action is slightly different from invoking the index action on the Instances resource with the filter "parent_href == /api/server_arrays/XX" because the
// latter will include 'next_instance' as well.
func (loc *ServerArrayLocator) CurrentInstances() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerArrays", "current_instances")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// DELETE /api/server_arrays/:id
// DELETE /api/deployments/:deployment_id/server_arrays/:id
// Deletes a given server array.
func (loc *ServerArrayLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerArrays", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_arrays
// GET /api/deployments/:deployment_id/server_arrays
// Lists server arrays.
// By using the available filters, it is possible to retrieve server arrays that have common characteristics.
// For example, one can list:
// arrays that have names that contain "my_server_array"
// all arrays of a given deployment
// -- Optional parameters:
// filter
// view
func (loc *ServerArrayLocator) Index(options rsapi.ApiParams) ([]*ServerArray, error) {
	var res []*ServerArray
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("ServerArrays", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/server_arrays/:id/launch
// Launches a new instance in the server array with the configuration defined in the 'next_instance'. This function is equivalent to invoking the launch action on the
// URL of this server_array's next_instance. See Instances#launch for details.
func (loc *ServerArrayLocator) Launch() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerArrays", "launch")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_arrays/:id/multi_run_executable
// Run an executable on all instances of this array. This function is equivalent to invoking the "multi_run_executable" action on the instances resource
// (Instances#multi_run_executable with the filter "parent_href == /api/server_arrays/XX"). To run an executable on a subset of the instances of the array, provide additional filters. To run an executable
// a single instance, invoke the action "run_executable" directly on the instance (see Instances#run_executable)
func (loc *ServerArrayLocator) MultiRunExecutable() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerArrays", "multi_run_executable")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_arrays/:id/multi_terminate
// Terminate all instances of this array. This function is equivalent to invoking the "multi_terminate" action on the instances resource ( Instances#multi_terminate with
// the filter "parent_href == /api/server_arrays/XX"). To terminate a subset of the instances of the array, provide additional filters. To terminate a single instance,
// invoke the action "terminate" directly on the instance (see Instances#terminate)
func (loc *ServerArrayLocator) MultiTerminate() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerArrays", "multi_terminate")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_arrays/:id
// GET /api/deployments/:deployment_id/server_arrays/:id
// Shows the information of a single server array.
// -- Optional parameters:
// view
func (loc *ServerArrayLocator) Show(options rsapi.ApiParams) (*ServerArray, error) {
	var res *ServerArray
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("ServerArrays", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/server_arrays/:id
// PUT /api/deployments/:deployment_id/server_arrays/:id
// Updates attributes of a single server array.
func (loc *ServerArrayLocator) Update(serverArray *ServerArrayParam2) error {
	if serverArray == nil {
		return fmt.Errorf("serverArray is required")
	}
	var params = rsapi.ApiParams{
		"server_array": serverArray,
	}
	var uri, err = loc.Url("ServerArrays", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
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
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Inputs      []map[string]string `json:"inputs,omitempty"`
	Lineage     string              `json:"lineage,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Revision    string              `json:"revision,omitempty"`
}

//===== Locator

// ServerTemplate resource locator, exposes resource actions.
type ServerTemplateLocator struct {
	UrlResolver
	api *Api15
}

// ServerTemplate resource locator factory
func (api *Api15) ServerTemplateLocator(href string) *ServerTemplateLocator {
	return &ServerTemplateLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/server_templates/:id/clone
// Clones a given ServerTemplate.
func (loc *ServerTemplateLocator) Clone(serverTemplate *ServerTemplateParam) error {
	if serverTemplate == nil {
		return fmt.Errorf("serverTemplate is required")
	}
	var params = rsapi.ApiParams{
		"server_template": serverTemplate,
	}
	var uri, err = loc.Url("ServerTemplates", "clone")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_templates/:id/commit
// Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.
// commitHeadDependencies: Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.
// commitMessage: The message associated with the commit.
// freezeRepositories: Freeze the repositories.
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
	var params = rsapi.ApiParams{
		"commit_head_dependencies": commitHeadDependencies,
		"commit_message":           commitMessage,
		"freeze_repositories":      freezeRepositories,
	}
	var uri, err = loc.Url("ServerTemplates", "commit")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_templates
// Creates a new ServerTemplate with the given parameters.
func (loc *ServerTemplateLocator) Create(serverTemplate *ServerTemplateParam) (*ServerTemplateLocator, error) {
	var res *ServerTemplateLocator
	if serverTemplate == nil {
		return res, fmt.Errorf("serverTemplate is required")
	}
	var params = rsapi.ApiParams{
		"server_template": serverTemplate,
	}
	var uri, err = loc.Url("ServerTemplates", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ServerTemplateLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/server_templates/:id
// Deletes a given ServerTemplate.
func (loc *ServerTemplateLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerTemplates", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_templates/:id/detect_changes_in_head
// Identifies RightScripts attached to the resource that differ from their HEAD.
// If the attached revision of the RightScript is the HEAD, then this will indicate
// a difference between it and the latest committed revision in the same lineage.
func (loc *ServerTemplateLocator) DetectChangesInHead() ([]*map[string]string, error) {
	var res []*map[string]string
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerTemplates", "detect_changes_in_head")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/server_templates
// Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// filter
// view
func (loc *ServerTemplateLocator) Index(options rsapi.ApiParams) ([]*ServerTemplate, error) {
	var res []*ServerTemplate
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("ServerTemplates", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/server_templates/:id/publish
// Publishes a given ServerTemplate and its subordinates.
// Only non-HEAD revisions that are owned by the account can be published.
// accountGroupHrefs: List of hrefs of account groups to publish to.
// -- Optional parameters:
// allow_comments: Allow users to leave comments on this ServerTemplate.
// categories: List of Categories.
// email_comments: Email me when a user comments on this ServerTemplate.
func (loc *ServerTemplateLocator) Publish(accountGroupHrefs []string, descriptions *Descriptions, options rsapi.ApiParams) error {
	if len(accountGroupHrefs) == 0 {
		return fmt.Errorf("accountGroupHrefs is required")
	}
	if descriptions == nil {
		return fmt.Errorf("descriptions is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"account_group_hrefs": accountGroupHrefs,
		"descriptions":        descriptions,
	}, options)
	var uri, err = loc.Url("ServerTemplates", "publish")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/server_templates/:id/resolve
// Enumerates all attached cookbooks, missing dependencies and bound executables.
// Version constraints on missing dependencies and the state of the Chef Recipes;
// whether or not the cookbook or recipe itself could be found among the
// attachments, will also be reported.
func (loc *ServerTemplateLocator) Resolve() ([]*map[string]string, error) {
	var res []*map[string]string
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerTemplates", "resolve")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/server_templates/:id
// Show information about a single ServerTemplate. HEAD revisions have a revision of 0.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// view
func (loc *ServerTemplateLocator) Show(options rsapi.ApiParams) (*ServerTemplate, error) {
	var res *ServerTemplate
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("ServerTemplates", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// POST /api/server_templates/:id/swap_repository
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
// sourceRepositoryHref: The repository whose cookbook attachments are to be replaced.
// targetRepositoryHref: The repository whose cookbook attachments are to be utilized.
func (loc *ServerTemplateLocator) SwapRepository(sourceRepositoryHref string, targetRepositoryHref string) error {
	if sourceRepositoryHref == "" {
		return fmt.Errorf("sourceRepositoryHref is required")
	}
	if targetRepositoryHref == "" {
		return fmt.Errorf("targetRepositoryHref is required")
	}
	var params = rsapi.ApiParams{
		"source_repository_href": sourceRepositoryHref,
		"target_repository_href": targetRepositoryHref,
	}
	var uri, err = loc.Url("ServerTemplates", "swap_repository")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// PUT /api/server_templates/:id
// Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).
// Currently, the attributes you can update are only the 'direct' attributes of a server template. To
// manage multi cloud images of a ServerTemplate, please see the resource 'ServerTemplateMultiCloudImages'.
func (loc *ServerTemplateLocator) Update(serverTemplate *ServerTemplateParam) error {
	if serverTemplate == nil {
		return fmt.Errorf("serverTemplate is required")
	}
	var params = rsapi.ApiParams{
		"server_template": serverTemplate,
	}
	var uri, err = loc.Url("ServerTemplates", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  ServerTemplateMultiCloudImage ******/

// This resource represents links between ServerTemplates and MultiCloud Images and enables you to effectively
// add/delete MultiCloudImages to ServerTemplates and make them the default one.
type ServerTemplateMultiCloudImage struct {
	Actions   []string            `json:"actions,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	IsDefault bool                `json:"is_default,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// ServerTemplateMultiCloudImage resource locator, exposes resource actions.
type ServerTemplateMultiCloudImageLocator struct {
	UrlResolver
	api *Api15
}

// ServerTemplateMultiCloudImage resource locator factory
func (api *Api15) ServerTemplateMultiCloudImageLocator(href string) *ServerTemplateMultiCloudImageLocator {
	return &ServerTemplateMultiCloudImageLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/server_template_multi_cloud_images
// Creates a new ServerTemplateMultiCloudImage with the given parameters.
func (loc *ServerTemplateMultiCloudImageLocator) Create(serverTemplateMultiCloudImage *ServerTemplateMultiCloudImageParam) (*ServerTemplateMultiCloudImageLocator, error) {
	var res *ServerTemplateMultiCloudImageLocator
	if serverTemplateMultiCloudImage == nil {
		return res, fmt.Errorf("serverTemplateMultiCloudImage is required")
	}
	var params = rsapi.ApiParams{
		"server_template_multi_cloud_image": serverTemplateMultiCloudImage,
	}
	var uri, err = loc.Url("ServerTemplateMultiCloudImages", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ServerTemplateMultiCloudImageLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/server_template_multi_cloud_images/:id
// Deletes a given ServerTemplateMultiCloudImage.
func (loc *ServerTemplateMultiCloudImageLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerTemplateMultiCloudImages", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_template_multi_cloud_images
// Lists the ServerTemplateMultiCloudImages available to this account.
// -- Optional parameters:
// filter
// view
func (loc *ServerTemplateMultiCloudImageLocator) Index(options rsapi.ApiParams) ([]*ServerTemplateMultiCloudImage, error) {
	var res []*ServerTemplateMultiCloudImage
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("ServerTemplateMultiCloudImages", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/server_template_multi_cloud_images/:id/make_default
// Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.
func (loc *ServerTemplateMultiCloudImageLocator) MakeDefault() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("ServerTemplateMultiCloudImages", "make_default")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_template_multi_cloud_images/:id
// Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.
// -- Optional parameters:
// view
func (loc *ServerTemplateMultiCloudImageLocator) Show(options rsapi.ApiParams) (*ServerTemplateMultiCloudImage, error) {
	var res *ServerTemplateMultiCloudImage
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("ServerTemplateMultiCloudImages", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Session ******/

// The sessions resource is in charge of creating API sessions that are bound to a given account. The sequence for login into the API is the following:
// Perform a POST request to /api/sessions ('create' action) to my.rightscale.com or to any more specific hosts saved from previous sessions.
// If the targeted host is not appropriate for the specific account being accessed it will return a 302 http code with a URL with which the client must retry the same POST request.
// If the targeted host is the right one and the login is successful, it will return a 204 http code, along with two cookies that will need to be saved and passed in any subsequent API request.
// If there is an authentication or authorization problem with the POST request an error (typically 401 or 422 ) may be returned at any point in the above sequence.
// If the session expires, it will return a 403 http code with a "Session cookie is expired or invalid" message.
// Note that all API calls irrespective of the resource it is acting on, should pass a header "X_API_VERSION" with the value "1.5".
type Session struct {
	Actions []string            `json:"actions,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
	Message string              `json:"message,omitempty"`
}

//===== Locator

// Session resource locator, exposes resource actions.
type SessionLocator struct {
	UrlResolver
	api *Api15
}

// Session resource locator factory
func (api *Api15) SessionLocator(href string) *SessionLocator {
	return &SessionLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/sessions/accounts
// List all the accounts that a user has access to.
// This call may be executed outside of an existing session. Doing so requires passing a username and password in the
// request body. The idea is that it should be possible to list accounts that can be used to create a session.
// Upon successfully authenticating the credentials, the system will return a 200 OK code and return the list of accounts.
// If an 302 redirect code is returned, the client is responsible of re-issuing the GET request against the content of the received Location header, passing the exact same parameters again.
// Example Request using Curl (not using an existing session):
// curl -i -H X_API_VERSION:1.5 -X GET -d email='email@me.com' -d password='mypassword' https://my.rightscale.com/api/sessions/accounts
// Example Request using Curl (using an existing session):
// curl -i -H X_API_VERSION:1.5 -X GET -b mycookies https://my.rightscale.com/api/sessions/accounts
// -- Optional parameters:
// email: The email to login with if not using existing session.
// password: The corresponding password.
// view: Extended view shows account permissions and products
func (loc *SessionLocator) Accounts(options rsapi.ApiParams) ([]*Account, error) {
	var res []*Account
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Sessions", "accounts")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/sessions
// Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what
// features are available within its privileges.
// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions
func (loc *SessionLocator) Index() ([]*Session, error) {
	var res []*Session
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Sessions", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/sessions/instance
// Shows the full attributes of the instance (that has the token used to log-in).
// This call can be used by an instance to get it's own details.
// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions/instance
func (loc *SessionLocator) IndexInstanceSession() (Instance, error) {
	var res Instance
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Sessions", "index_instance_session")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

/******  SshKey ******/

// Ssh Keys represent a created SSH Key that exists in the cloud.
// An ssh key might also contain the private part of the key, and can be used to login to instances launched with it.
type SshKey struct {
	Actions     []string            `json:"actions,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Material    string              `json:"material,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

//===== Locator

// SshKey resource locator, exposes resource actions.
type SshKeyLocator struct {
	UrlResolver
	api *Api15
}

// SshKey resource locator factory
func (api *Api15) SshKeyLocator(href string) *SshKeyLocator {
	return &SshKeyLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/ssh_keys
// Creates a new ssh key.
func (loc *SshKeyLocator) Create(sshKey *SshKeyParam) (*SshKeyLocator, error) {
	var res *SshKeyLocator
	if sshKey == nil {
		return res, fmt.Errorf("sshKey is required")
	}
	var params = rsapi.ApiParams{
		"ssh_key": sshKey,
	}
	var uri, err = loc.Url("SshKeys", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &SshKeyLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/ssh_keys/:id
// Deletes a given ssh key.
func (loc *SshKeyLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("SshKeys", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/ssh_keys
// Lists ssh keys.
// -- Optional parameters:
// filter
// view
func (loc *SshKeyLocator) Index(options rsapi.ApiParams) ([]*SshKey, error) {
	var res []*SshKey
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("SshKeys", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/ssh_keys/:id
// Displays information about a single ssh key.
// -- Optional parameters:
// view
func (loc *SshKeyLocator) Show(options rsapi.ApiParams) (*SshKey, error) {
	var res *SshKey
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("SshKeys", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
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

//===== Locator

// Subnet resource locator, exposes resource actions.
type SubnetLocator struct {
	UrlResolver
	api *Api15
}

// Subnet resource locator factory
func (api *Api15) SubnetLocator(href string) *SubnetLocator {
	return &SubnetLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances/:instance_id/subnets
// POST /api/clouds/:cloud_id/subnets
// Creates a new subnet.
func (loc *SubnetLocator) Create(subnet *SubnetParam) (*SubnetLocator, error) {
	var res *SubnetLocator
	if subnet == nil {
		return res, fmt.Errorf("subnet is required")
	}
	var params = rsapi.ApiParams{
		"subnet": subnet,
	}
	var uri, err = loc.Url("Subnets", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &SubnetLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/instances/:instance_id/subnets/:id
// DELETE /api/clouds/:cloud_id/subnets/:id
// Deletes the given subnet(s).
func (loc *SubnetLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Subnets", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/instances/:instance_id/subnets
// GET /api/clouds/:cloud_id/subnets
// Lists subnets of a given cloud.
// -- Optional parameters:
// filter
func (loc *SubnetLocator) Index(options rsapi.ApiParams) ([]*Subnet, error) {
	var res []*Subnet
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Subnets", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/instances/:instance_id/subnets/:id
// GET /api/clouds/:cloud_id/subnets/:id
// Shows attributes of a single subnet.
func (loc *SubnetLocator) Show() (*Subnet, error) {
	var res *Subnet
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Subnets", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/clouds/:cloud_id/instances/:instance_id/subnets/:id
// PUT /api/clouds/:cloud_id/subnets/:id
// Updates name and description for the current subnet.
func (loc *SubnetLocator) Update(subnet *SubnetParam2) error {
	if subnet == nil {
		return fmt.Errorf("subnet is required")
	}
	var params = rsapi.ApiParams{
		"subnet": subnet,
	}
	var uri, err = loc.Url("Subnets", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Tag ******/

// A tag or machine tag is a useful way of attaching useful metadata to an object/resource.
// Tags are commonly used as an extra label or identifier.
// For example, you might want to add a tag to an EBS Snapshot or AMI so that you can find it more quickly.

//===== Locator

// Tag resource locator, exposes resource actions.
type TagLocator struct {
	UrlResolver
	api *Api15
}

// Tag resource locator factory
func (api *Api15) TagLocator(href string) *TagLocator {
	return &TagLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/tags/by_resource
// Get tags for a list of resource hrefs.
// The hrefs can belong to various resource types and the tags for a non-existent href will be empty.
// resourceHrefs: Hrefs of the resources for which tags are to be returned.
func (loc *TagLocator) ByResource(resourceHrefs []string) ([]*map[string]string, error) {
	var res []*map[string]string
	if len(resourceHrefs) == 0 {
		return res, fmt.Errorf("resourceHrefs is required")
	}
	var params = rsapi.ApiParams{
		"resource_hrefs": resourceHrefs,
	}
	var uri, err = loc.Url("Tags", "by_resource")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/tags/by_tag
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
// resourceType: Search among a single resource type.
// tags: The tags which must be present on the resource.
// -- Optional parameters:
// include_tags_with_prefix: If included, all tags matching this prefix will be returned. If not included, no tags will be returned.
// match_all: If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.
// with_deleted: If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.
func (loc *TagLocator) ByTag(resourceType string, tags []string, options rsapi.ApiParams) ([]*map[string]string, error) {
	var res []*map[string]string
	if resourceType == "" {
		return res, fmt.Errorf("resourceType is required")
	}
	if len(tags) == 0 {
		return res, fmt.Errorf("tags is required")
	}
	var params = mergeOptionals(rsapi.ApiParams{
		"resource_type": resourceType,
		"tags":          tags,
	}, options)
	var uri, err = loc.Url("Tags", "by_tag")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// POST /api/tags/multi_add
// Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.
// The hrefs can belong to various resource types. If a resource for a href could not be found, an
// error is returned and no tags are added for any resource.
// No error will be raised if the resource already has the tag(s) you are trying to add.
// Note: At this point, tags on 'next_instance' are not supported and one has to add tags to the 'server'.
// resourceHrefs: Hrefs of the resources for which the tags are to be added.
// tags: Tags to be added.
func (loc *TagLocator) MultiAdd(resourceHrefs []string, tags []string) error {
	if len(resourceHrefs) == 0 {
		return fmt.Errorf("resourceHrefs is required")
	}
	if len(tags) == 0 {
		return fmt.Errorf("tags is required")
	}
	var params = rsapi.ApiParams{
		"resource_hrefs": resourceHrefs,
		"tags":           tags,
	}
	var uri, err = loc.Url("Tags", "multi_add")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// POST /api/tags/multi_delete
// Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.
// The hrefs can belong to various resource types. If a resource for a href could not be found, an
// error is returned and no tags are deleted for any resource.
// Note that no error will be raised if the resource does not have the tag(s) you are trying to delete.
// resourceHrefs: Hrefs of the resources for which tags are to be deleted.
// tags: Tags to be deleted.
func (loc *TagLocator) MultiDelete(resourceHrefs []string, tags []string) error {
	if len(resourceHrefs) == 0 {
		return fmt.Errorf("resourceHrefs is required")
	}
	if len(tags) == 0 {
		return fmt.Errorf("tags is required")
	}
	var params = rsapi.ApiParams{
		"resource_hrefs": resourceHrefs,
		"tags":           tags,
	}
	var uri, err = loc.Url("Tags", "multi_delete")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Task ******/

// Tasks represent processes that happen (or have happened) asynchronously within the context of an Instance.
// An example of a type of task is an operational script that runs in an instance.
// Task resources can be returned by certain API calls, such as Instances.run_executable, Backups.restore, and others.
type Task struct {
	Actions []string            `json:"actions,omitempty"`
	Detail  string              `json:"detail,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
	Summary string              `json:"summary,omitempty"`
}

//===== Locator

// Task resource locator, exposes resource actions.
type TaskLocator struct {
	UrlResolver
	api *Api15
}

// Task resource locator factory
func (api *Api15) TaskLocator(href string) *TaskLocator {
	return &TaskLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/instances/:instance_id/live/tasks/:id
// GET /api/server_arrays/:server_array_id/live/tasks/:id
// Displays information of a given task within the context of an instance.
// -- Optional parameters:
// view
func (loc *TaskLocator) Show(options rsapi.ApiParams) (*Task, error) {
	var res *Task
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Tasks", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  User ******/

// A User represents an individual RightScale user. Users must be given access to RightScale accounts in order to
// access RightScale resources.
type User struct {
	Actions      []string            `json:"actions,omitempty"`
	Company      string              `json:"company,omitempty"`
	CreatedAt    time.Time           `json:"created_at,omitempty"`
	Email        string              `json:"email,omitempty"`
	FirstName    string              `json:"first_name,omitempty"`
	LastName     string              `json:"last_name,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	Phone        string              `json:"phone,omitempty"`
	PrincipalUid string              `json:"principal_uid,omitempty"`
	TimezoneName string              `json:"timezone_name,omitempty"`
	UpdatedAt    time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// User resource locator, exposes resource actions.
type UserLocator struct {
	UrlResolver
	api *Api15
}

// User resource locator factory
func (api *Api15) UserLocator(href string) *UserLocator {
	return &UserLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/users
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
func (loc *UserLocator) Create(user *UserParam) (*UserLocator, error) {
	var res *UserLocator
	if user == nil {
		return res, fmt.Errorf("user is required")
	}
	var params = rsapi.ApiParams{
		"user": user,
	}
	var uri, err = loc.Url("Users", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &UserLocator{UrlResolver(location), loc.api}, nil
	}
}

// GET /api/users
// List the users available to the account the user is logged in to. Therefore, to list the users of
// a child account, the user has to login to the child account first.
// -- Optional parameters:
// filter
func (loc *UserLocator) Index(options rsapi.ApiParams) ([]*User, error) {
	var res []*User
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Users", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/users/:id
// Show information about a single user.
func (loc *UserLocator) Show() (*User, error) {
	var res *User
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Users", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// PUT /api/users/:id
// Update a user's contact information, change her password, or update SSO her settings. In order
// to update a user record, one of the following criteria must be met:
// You're logged in AS the user being modified and you provide a valid current_password.
// You're an admin and the user is linked to your enterprise SSO provider
// You're an admin and the user's email matches the email_domain of your enterprise SSO provider
// In other words: you can update yourself if you know your own password; you can update
// yourself or others if they're linked to your SSO providers, and you can update any user
// if her email address is known to belong to your organization.
// For information about enabling canonical email domain ownership for your enterprise, please
// talk to your RightScale account manager or contact our support team.
// To update a user's contact information, simply pass the desired values for email, first_name,
// and so forth.
// To update a user's password, provide a valid current_password and specify the desired
// new_password.
// To update a user's SSO information, you may provide a just a principal_uid (to maintain the
// user's existing identity provider) or you may provide an identity_provider_href and a
// principal_uid (to switch identity providers as well as specify a new user identity).
// In the context of SAML. principal_uid is equivalent to the SAML NameID or Subject claim;
// RightScale cannot predict or influence the NameID value that your SAML IdP will send to us for
func (loc *UserLocator) Update(user *UserParam2) error {
	if user == nil {
		return fmt.Errorf("user is required")
	}
	var params = rsapi.ApiParams{
		"user": user,
	}
	var uri, err = loc.Url("Users", "update")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  UserData ******/

//===== Locator

// UserData resource locator, exposes resource actions.
type UserDataLocator struct {
	UrlResolver
	api *Api15
}

// UserData resource locator factory
func (api *Api15) UserDataLocator(href string) *UserDataLocator {
	return &UserDataLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/user_data/
// No description provided for show.
func (loc *UserDataLocator) Show() (*map[string]string, error) {
	var res *map[string]string
	var params = map[string]interface{}{}
	var uri, err = loc.Url("UserDatas", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  Volume ******/

// A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).
type Volume struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Iops        string              `json:"iops,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Size        string              `json:"size,omitempty"`
	Status      string              `json:"status,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
	VolumeType  string              `json:"volume_type,omitempty"`
}

//===== Locator

// Volume resource locator, exposes resource actions.
type VolumeLocator struct {
	UrlResolver
	api *Api15
}

// Volume resource locator factory
func (api *Api15) VolumeLocator(href string) *VolumeLocator {
	return &VolumeLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/volumes
// Creates a new volume.
func (loc *VolumeLocator) Create(volume *VolumeParam) (*VolumeLocator, error) {
	var res *VolumeLocator
	if volume == nil {
		return res, fmt.Errorf("volume is required")
	}
	var params = rsapi.ApiParams{
		"volume": volume,
	}
	var uri, err = loc.Url("Volumes", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &VolumeLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/volumes/:id
// Deletes a given volume.
func (loc *VolumeLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("Volumes", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/volumes
// Lists volumes.
// -- Optional parameters:
// filter
// view
func (loc *VolumeLocator) Index(options rsapi.ApiParams) ([]*Volume, error) {
	var res []*Volume
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Volumes", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/volumes/:id
// Displays information about a single volume.
// -- Optional parameters:
// view
func (loc *VolumeLocator) Show(options rsapi.ApiParams) (*Volume, error) {
	var res *Volume
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("Volumes", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  VolumeAttachment ******/

// A VolumeAttachment represents a relationship between a volume and an instance.
type VolumeAttachment struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Device      string              `json:"device,omitempty"`
	DeviceId    string              `json:"device_id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// VolumeAttachment resource locator, exposes resource actions.
type VolumeAttachmentLocator struct {
	UrlResolver
	api *Api15
}

// VolumeAttachment resource locator factory
func (api *Api15) VolumeAttachmentLocator(href string) *VolumeAttachmentLocator {
	return &VolumeAttachmentLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/instances/:instance_id/volume_attachments
// POST /api/clouds/:cloud_id/volume_attachments
// POST /api/clouds/:cloud_id/volumes/:volume_id/volume_attachments
// POST /api/clouds/:cloud_id/volumes/:volume_id/volume_attachment
// Creates a new volume attachment.
func (loc *VolumeAttachmentLocator) Create(volumeAttachment *VolumeAttachmentParam) (*VolumeAttachmentLocator, error) {
	var res *VolumeAttachmentLocator
	if volumeAttachment == nil {
		return res, fmt.Errorf("volumeAttachment is required")
	}
	var params = rsapi.ApiParams{
		"volume_attachment": volumeAttachment,
	}
	var uri, err = loc.Url("VolumeAttachments", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &VolumeAttachmentLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volume_attachments/:id
// DELETE /api/clouds/:cloud_id/volumes/:volume_id/volume_attachments
// DELETE /api/clouds/:cloud_id/volumes/:volume_id/volume_attachment
// Deletes a given volume attachment.
// -- Optional parameters:
// force: Specifies whether to force the detachment of a volume.
func (loc *VolumeAttachmentLocator) Destroy(options rsapi.ApiParams) error {
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("VolumeAttachments", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/instances/:instance_id/volume_attachments
// GET /api/clouds/:cloud_id/volume_attachments
// Lists all volume attachments.
// -- Optional parameters:
// filter
// view
func (loc *VolumeAttachmentLocator) Index(options rsapi.ApiParams) ([]*VolumeAttachment, error) {
	var res []*VolumeAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("VolumeAttachments", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id
// GET /api/clouds/:cloud_id/volume_attachments/:id
// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_attachments
// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_attachment
// Displays information about a single volume attachment.
// -- Optional parameters:
// view
func (loc *VolumeAttachmentLocator) Show(options rsapi.ApiParams) (*VolumeAttachment, error) {
	var res *VolumeAttachment
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("VolumeAttachments", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  VolumeSnapshot ******/

// A VolumeSnapshot represents a Cloud storage volume at a particular point in time. One can create a snapshot regardless of whether or not a volume is attached to an Instance. When a snapshot is created,
// various meta data is retained such as a Created At timestamp, a unique Resource UID (e.g. vol-52EF05A9), the Volume Owner and Visibility (e.g. private or public).
// Snapshots consist of a series of data blocks that are incrementally saved.
type VolumeSnapshot struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Progress    string              `json:"progress,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Size        string              `json:"size,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// VolumeSnapshot resource locator, exposes resource actions.
type VolumeSnapshotLocator struct {
	UrlResolver
	api *Api15
}

// VolumeSnapshot resource locator factory
func (api *Api15) VolumeSnapshotLocator(href string) *VolumeSnapshotLocator {
	return &VolumeSnapshotLocator{UrlResolver(href), api}
}

//===== Actions

// POST /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots
// POST /api/clouds/:cloud_id/volume_snapshots
// Creates a new volume_snapshot.
func (loc *VolumeSnapshotLocator) Create(volumeSnapshot *VolumeSnapshotParam) (*VolumeSnapshotLocator, error) {
	var res *VolumeSnapshotLocator
	if volumeSnapshot == nil {
		return res, fmt.Errorf("volumeSnapshot is required")
	}
	var params = rsapi.ApiParams{
		"volume_snapshot": volumeSnapshot,
	}
	var uri, err = loc.Url("VolumeSnapshots", "create")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &VolumeSnapshotLocator{UrlResolver(location), loc.api}, nil
	}
}

// DELETE /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id
// DELETE /api/clouds/:cloud_id/volume_snapshots/:id
// Deletes a given volume_snapshot.
func (loc *VolumeSnapshotLocator) Destroy() error {
	var params = map[string]interface{}{}
	var uri, err = loc.Url("VolumeSnapshots", "destroy")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots
// GET /api/clouds/:cloud_id/volume_snapshots
// Lists all volume_snapshots.
// -- Optional parameters:
// filter
// view
func (loc *VolumeSnapshotLocator) Index(options rsapi.ApiParams) ([]*VolumeSnapshot, error) {
	var res []*VolumeSnapshot
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("VolumeSnapshots", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id
// GET /api/clouds/:cloud_id/volume_snapshots/:id
// Displays information about a single volume_snapshot.
// -- Optional parameters:
// view
func (loc *VolumeSnapshotLocator) Show(options rsapi.ApiParams) (*VolumeSnapshot, error) {
	var res *VolumeSnapshot
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("VolumeSnapshots", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/******  VolumeType ******/

// A VolumeType describes the type of volume, particularly the size.
type VolumeType struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Size        string              `json:"size,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
}

//===== Locator

// VolumeType resource locator, exposes resource actions.
type VolumeTypeLocator struct {
	UrlResolver
	api *Api15
}

// VolumeType resource locator factory
func (api *Api15) VolumeTypeLocator(href string) *VolumeTypeLocator {
	return &VolumeTypeLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/volume_types
// Lists Volume Types.
// -- Optional parameters:
// filter
// view
func (loc *VolumeTypeLocator) Index(options rsapi.ApiParams) ([]*VolumeType, error) {
	var res []*VolumeType
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("VolumeTypes", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, &res)
	return res, err4
}

// GET /api/clouds/:cloud_id/volume_types/:id
// Displays information about a single Volume Type.
// -- Optional parameters:
// view
func (loc *VolumeTypeLocator) Show(options rsapi.ApiParams) (*VolumeType, error) {
	var res *VolumeType
	var params = mergeOptionals(rsapi.ApiParams{}, options)
	var uri, err = loc.Url("VolumeTypes", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return res, err2
	}
	defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

/****** Parameter Data Types ******/

type AlertSpecParam struct {
	Condition      string  `json:"condition,omitempty"`
	Description    *string `json:"description,omitempty"`
	Duration       string  `json:"duration,omitempty"`
	EscalationName *string `json:"escalation_name,omitempty"`
	File           string  `json:"file,omitempty"`
	Name           string  `json:"name,omitempty"`
	SubjectHref    *string `json:"subject_href,omitempty"`
	Threshold      string  `json:"threshold,omitempty"`
	Variable       string  `json:"variable,omitempty"`
	VoteTag        *string `json:"vote_tag,omitempty"`
	VoteType       *string `json:"vote_type,omitempty"`
}

type AlertSpecParam2 struct {
	Condition      *string `json:"condition,omitempty"`
	Description    *string `json:"description,omitempty"`
	Duration       *string `json:"duration,omitempty"`
	EscalationName *string `json:"escalation_name,omitempty"`
	File           *string `json:"file,omitempty"`
	Name           *string `json:"name,omitempty"`
	Threshold      *string `json:"threshold,omitempty"`
	Variable       *string `json:"variable,omitempty"`
	VoteTag        *string `json:"vote_tag,omitempty"`
	VoteType       *string `json:"vote_type,omitempty"`
}

type AlertSpecificParams struct {
	DecisionThreshold  *string `json:"decision_threshold,omitempty"`
	VotersTagPredicate *string `json:"voters_tag_predicate,omitempty"`
}

type AssetPaths struct {
	Cookbooks []string `json:"cookbooks,omitempty"`
}

type AuditEntryParam struct {
	AuditeeHref string  `json:"auditee_href,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Summary     string  `json:"summary,omitempty"`
}

type AuditEntryParam2 struct {
	Offset  *int   `json:"offset,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type BackupParam struct {
	Description           *string  `json:"description,omitempty"`
	FromMaster            *string  `json:"from_master,omitempty"`
	Lineage               string   `json:"lineage,omitempty"`
	Name                  string   `json:"name,omitempty"`
	VolumeAttachmentHrefs []string `json:"volume_attachment_hrefs,omitempty"`
}

type BackupParam2 struct {
	Committed string `json:"committed,omitempty"`
}

type Bounds struct {
	MaxCount *string `json:"max_count,omitempty"`
	MinCount *string `json:"min_count,omitempty"`
}

type ChildAccountParam struct {
	ClusterHref *string `json:"cluster_href,omitempty"`
	Name        string  `json:"name,omitempty"`
}

type ChildAccountParam2 struct {
	Name *string `json:"name,omitempty"`
}

type CloudAccountParam struct {
	CloudHref *string           `json:"cloud_href,omitempty"`
	Creds     map[string]string `json:"creds,omitempty"`
	Token     *string           `json:"token,omitempty"`
}

type CloudSpecificAttributes struct {
	AutomaticInstanceStoreMapping *string `json:"automatic_instance_store_mapping,omitempty"`
	EbsOptimized                  *string `json:"ebs_optimized,omitempty"`
	IamInstanceProfile            *string `json:"iam_instance_profile,omitempty"`
	RootVolumePerformance         *string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                *string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid             *string `json:"root_volume_type_uid,omitempty"`
}

type CloudSpecificAttributes2 struct {
	AutomaticInstanceStoreMapping *string `json:"automatic_instance_store_mapping,omitempty"`
	IamInstanceProfile            *string `json:"iam_instance_profile,omitempty"`
	RootVolumePerformance         *string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                *string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid             *string `json:"root_volume_type_uid,omitempty"`
}

type CookbookAttachmentParam struct {
	CookbookHref       *string `json:"cookbook_href,omitempty"`
	ServerTemplateHref *string `json:"server_template_href,omitempty"`
}

type CookbookAttachments struct {
	CookbookHrefs      []string `json:"cookbook_hrefs,omitempty"`
	ServerTemplateHref *string  `json:"server_template_href,omitempty"`
}

type CookbookAttachments2 struct {
	CookbookAttachmentHrefs []string `json:"cookbook_attachment_hrefs,omitempty"`
}

type CredentialParam struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
	Value       string  `json:"value,omitempty"`
}

type Credentials struct {
	Password *string `json:"password,omitempty"`
	SshKey   *string `json:"ssh_key,omitempty"`
	Username *string `json:"username,omitempty"`
}

type DatacenterPolicy struct {
	DatacenterHref string `json:"datacenter_href,omitempty"`
	Max            string `json:"max,omitempty"`
	Weight         string `json:"weight,omitempty"`
}

type DeploymentParam struct {
	Description    *string `json:"description,omitempty"`
	Name           *string `json:"name,omitempty"`
	ServerTagScope *string `json:"server_tag_scope,omitempty"`
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
	AssociatePublicIpAddress *string                  `json:"associate_public_ip_address,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           *string                  `json:"datacenter_href,omitempty"`
	DeploymentHref           *string                  `json:"deployment_href,omitempty"`
	ImageHref                string                   `json:"image_href,omitempty"`
	InstanceTypeHref         string                   `json:"instance_type_href,omitempty"`
	KernelImageHref          *string                  `json:"kernel_image_href,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	PlacementGroupHref       *string                  `json:"placement_group_href,omitempty"`
	RamdiskImageHref         *string                  `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                 `json:"security_group_hrefs,omitempty"`
	SshKeyHref               *string                  `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                 `json:"subnet_hrefs,omitempty"`
	UserData                 *string                  `json:"user_data,omitempty"`
}

type InstanceParam2 struct {
	Href                string            `json:"href,omitempty"`
	Inputs              map[string]string `json:"inputs,omitempty"`
	MultiCloudImageHref *string           `json:"multi_cloud_image_href,omitempty"`
	ServerTemplateHref  string            `json:"server_template_href,omitempty"`
}

type IpAddressBindingParam struct {
	InstanceHref        string  `json:"instance_href,omitempty"`
	PrivatePort         *string `json:"private_port,omitempty"`
	Protocol            *string `json:"protocol,omitempty"`
	PublicIpAddressHref *string `json:"public_ip_address_href,omitempty"`
	PublicPort          *string `json:"public_port,omitempty"`
}

type IpAddressParam struct {
	DeploymentHref *string `json:"deployment_href,omitempty"`
	Domain         *string `json:"domain,omitempty"`
	Name           string  `json:"name,omitempty"`
	NetworkHref    *string `json:"network_href,omitempty"`
}

type IpAddressParam2 struct {
	DeploymentHref *string `json:"deployment_href,omitempty"`
	Name           string  `json:"name,omitempty"`
}

type ItemAge struct {
	Algorithm *string `json:"algorithm,omitempty"`
	MaxAge    *string `json:"max_age,omitempty"`
	Regexp    *string `json:"regexp,omitempty"`
}

type MultiCloudImageParam struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
}

type MultiCloudImageSettingParam struct {
	CloudHref        *string `json:"cloud_href,omitempty"`
	ImageHref        *string `json:"image_href,omitempty"`
	InstanceTypeHref *string `json:"instance_type_href,omitempty"`
	KernelImageHref  *string `json:"kernel_image_href,omitempty"`
	RamdiskImageHref *string `json:"ramdisk_image_href,omitempty"`
	UserData         *string `json:"user_data,omitempty"`
}

type NetworkGatewayParam struct {
	CloudHref   string  `json:"cloud_href,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
	Type_       string  `json:"type,omitempty"`
}

type NetworkGatewayParam2 struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	NetworkHref *string `json:"network_href,omitempty"`
}

type NetworkOptionGroupAttachmentParam struct {
	NetworkHref            string  `json:"network_href,omitempty"`
	NetworkOptionGroupHref *string `json:"network_option_group_href,omitempty"`
}

type NetworkOptionGroupAttachmentParam2 struct {
	NetworkOptionGroupHref *string `json:"network_option_group_href,omitempty"`
}

type NetworkOptionGroupParam struct {
	CloudHref   string            `json:"cloud_href,omitempty"`
	Description *string           `json:"description,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Options     map[string]string `json:"options,omitempty"`
	Type_       string            `json:"type,omitempty"`
}

type NetworkOptionGroupParam2 struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type NetworkParam struct {
	CidrBlock       *string `json:"cidr_block,omitempty"`
	CloudHref       string  `json:"cloud_href,omitempty"`
	Description     *string `json:"description,omitempty"`
	InstanceTenancy *string `json:"instance_tenancy,omitempty"`
	Name            *string `json:"name,omitempty"`
}

type NetworkParam2 struct {
	Description    *string `json:"description,omitempty"`
	Name           *string `json:"name,omitempty"`
	RouteTableHref *string `json:"route_table_href,omitempty"`
}

type Pacing struct {
	ResizeCalmTime *string `json:"resize_calm_time,omitempty"`
	ResizeDownBy   *string `json:"resize_down_by,omitempty"`
	ResizeUpBy     *string `json:"resize_up_by,omitempty"`
}

type PermissionParam struct {
	RoleTitle string `json:"role_title,omitempty"`
	UserHref  string `json:"user_href,omitempty"`
}

type PlacementGroupParam struct {
	CloudHref   string  `json:"cloud_href,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
}

type PreferenceParam struct {
	Contents string `json:"contents,omitempty"`
}

type ProtocolDetails struct {
	EndPort   *string `json:"end_port,omitempty"`
	IcmpCode  *string `json:"icmp_code,omitempty"`
	IcmpType  *string `json:"icmp_type,omitempty"`
	StartPort *string `json:"start_port,omitempty"`
}

type Quantity struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type QueueSize struct {
	ItemsPerInstance *string `json:"items_per_instance,omitempty"`
}

type QueueSpecificParams struct {
	CollectAuditEntries *string    `json:"collect_audit_entries,omitempty"`
	ItemAge             *ItemAge   `json:"item_age,omitempty"`
	QueueSize           *QueueSize `json:"queue_size,omitempty"`
}

type RecurringVolumeAttachmentParam struct {
	Device       string `json:"device,omitempty"`
	RunnableHref string `json:"runnable_href,omitempty"`
	StorageHref  string `json:"storage_href,omitempty"`
}

type RepositoryParam struct {
	AssetPaths      *AssetPaths  `json:"asset_paths,omitempty"`
	AutoImport      *string      `json:"auto_import,omitempty"`
	CommitReference *string      `json:"commit_reference,omitempty"`
	Credentials     *Credentials `json:"credentials,omitempty"`
	Description     *string      `json:"description,omitempty"`
	Name            string       `json:"name,omitempty"`
	Source          string       `json:"source,omitempty"`
	SourceType      string       `json:"source_type,omitempty"`
}

type RepositoryParam2 struct {
	AssetPaths      *AssetPaths  `json:"asset_paths,omitempty"`
	CommitReference *string      `json:"commit_reference,omitempty"`
	Credentials     *Credentials `json:"credentials,omitempty"`
	Description     *string      `json:"description,omitempty"`
	Name            *string      `json:"name,omitempty"`
	Source          *string      `json:"source,omitempty"`
	SourceType      *string      `json:"source_type,omitempty"`
}

type RightScriptParam struct {
	CommitMessage string `json:"commit_message,omitempty"`
}

type RightScriptParam2 struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type RouteParam struct {
	Description          *string `json:"description,omitempty"`
	DestinationCidrBlock string  `json:"destination_cidr_block,omitempty"`
	NextHopHref          *string `json:"next_hop_href,omitempty"`
	NextHopIp            *string `json:"next_hop_ip,omitempty"`
	NextHopType          string  `json:"next_hop_type,omitempty"`
	RouteTableHref       string  `json:"route_table_href,omitempty"`
}

type RouteParam2 struct {
	Description          *string `json:"description,omitempty"`
	DestinationCidrBlock *string `json:"destination_cidr_block,omitempty"`
	NextHopHref          *string `json:"next_hop_href,omitempty"`
	NextHopIp            *string `json:"next_hop_ip,omitempty"`
	NextHopType          *string `json:"next_hop_type,omitempty"`
}

type RouteTableParam struct {
	CloudHref   string  `json:"cloud_href,omitempty"`
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
	NetworkHref string  `json:"network_href,omitempty"`
}

type RouteTableParam2 struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type RunnableBindingParam struct {
	Position        *string `json:"position,omitempty"`
	Recipe          *string `json:"recipe,omitempty"`
	RightScriptHref *string `json:"right_script_href,omitempty"`
	Sequence        *string `json:"sequence,omitempty"`
}

type RunnableBindings struct {
	Id              string  `json:"id,omitempty"`
	Position        *string `json:"position,omitempty"`
	Recipe          *string `json:"recipe,omitempty"`
	RightScriptHref *string `json:"right_script_href,omitempty"`
	Sequence        *string `json:"sequence,omitempty"`
}

type Schedule struct {
	Day      string `json:"day,omitempty"`
	MaxCount string `json:"max_count,omitempty"`
	MinCount string `json:"min_count,omitempty"`
	Time     string `json:"time,omitempty"`
}

type SecurityGroupParam struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
	NetworkHref *string `json:"network_href,omitempty"`
}

type SecurityGroupRuleParam struct {
	CidrIps           *string          `json:"cidr_ips,omitempty"`
	Direction         *string          `json:"direction,omitempty"`
	GroupName         *string          `json:"group_name,omitempty"`
	GroupOwner        *string          `json:"group_owner,omitempty"`
	Protocol          string           `json:"protocol,omitempty"`
	ProtocolDetails   *ProtocolDetails `json:"protocol_details,omitempty"`
	SecurityGroupHref *string          `json:"security_group_href,omitempty"`
	SourceType        string           `json:"source_type,omitempty"`
}

type SecurityGroupRuleParam2 struct {
	Description *string `json:"description,omitempty"`
}

type ServerArrayParam struct {
	ArrayType        string              `json:"array_type,omitempty"`
	DatacenterPolicy []*DatacenterPolicy `json:"datacenter_policy,omitempty"`
	DeploymentHref   *string             `json:"deployment_href,omitempty"`
	Description      *string             `json:"description,omitempty"`
	ElasticityParams *ElasticityParams   `json:"elasticity_params,omitempty"`
	Instance         *InstanceParam2     `json:"instance,omitempty"`
	Name             string              `json:"name,omitempty"`
	Optimized        *string             `json:"optimized,omitempty"`
	State            string              `json:"state,omitempty"`
}

type ServerArrayParam2 struct {
	ArrayType        *string             `json:"array_type,omitempty"`
	DatacenterPolicy []*DatacenterPolicy `json:"datacenter_policy,omitempty"`
	DeploymentHref   *string             `json:"deployment_href,omitempty"`
	Description      *string             `json:"description,omitempty"`
	ElasticityParams *ElasticityParams   `json:"elasticity_params,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Optimized        *string             `json:"optimized,omitempty"`
	State            *string             `json:"state,omitempty"`
}

type ServerParam struct {
	DeploymentHref *string         `json:"deployment_href,omitempty"`
	Description    *string         `json:"description,omitempty"`
	Instance       *InstanceParam2 `json:"instance,omitempty"`
	Name           string          `json:"name,omitempty"`
	Optimized      *string         `json:"optimized,omitempty"`
}

type ServerParam2 struct {
	DeploymentHref *string         `json:"deployment_href,omitempty"`
	Description    *string         `json:"description,omitempty"`
	Instance       *InstanceParam2 `json:"instance,omitempty"`
	Name           *string         `json:"name,omitempty"`
}

type ServerTemplateMultiCloudImageParam struct {
	MultiCloudImageHref string `json:"multi_cloud_image_href,omitempty"`
	ServerTemplateHref  string `json:"server_template_href,omitempty"`
}

type ServerTemplateParam struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name,omitempty"`
}

type SshKeyParam struct {
	Name string `json:"name,omitempty"`
}

type SubnetParam struct {
	CidrBlock      string  `json:"cidr_block,omitempty"`
	DatacenterHref *string `json:"datacenter_href,omitempty"`
	Description    *string `json:"description,omitempty"`
	Name           *string `json:"name,omitempty"`
	NetworkHref    string  `json:"network_href,omitempty"`
}

type SubnetParam2 struct {
	Description    *string `json:"description,omitempty"`
	Name           *string `json:"name,omitempty"`
	RouteTableHref *string `json:"route_table_href,omitempty"`
}

type UserParam struct {
	Company              string  `json:"company,omitempty"`
	Email                string  `json:"email,omitempty"`
	FirstName            string  `json:"first_name,omitempty"`
	IdentityProviderHref *string `json:"identity_provider_href,omitempty"`
	LastName             string  `json:"last_name,omitempty"`
	Password             *string `json:"password,omitempty"`
	Phone                string  `json:"phone,omitempty"`
	PrincipalUid         *string `json:"principal_uid,omitempty"`
	TimezoneName         *string `json:"timezone_name,omitempty"`
}

type UserParam2 struct {
	Company              *string `json:"company,omitempty"`
	CurrentEmail         string  `json:"current_email,omitempty"`
	CurrentPassword      *string `json:"current_password,omitempty"`
	FirstName            *string `json:"first_name,omitempty"`
	IdentityProviderHref *string `json:"identity_provider_href,omitempty"`
	LastName             *string `json:"last_name,omitempty"`
	NewEmail             *string `json:"new_email,omitempty"`
	NewPassword          *string `json:"new_password,omitempty"`
	Phone                *string `json:"phone,omitempty"`
	PrincipalUid         *string `json:"principal_uid,omitempty"`
	TimezoneName         *string `json:"timezone_name,omitempty"`
}

type VolumeAttachmentParam struct {
	Device       *string `json:"device,omitempty"`
	InstanceHref *string `json:"instance_href,omitempty"`
	VolumeHref   *string `json:"volume_href,omitempty"`
}

type VolumeParam struct {
	DatacenterHref           *string `json:"datacenter_href,omitempty"`
	DeploymentHref           *string `json:"deployment_href,omitempty"`
	Description              *string `json:"description,omitempty"`
	Encrypted                *string `json:"encrypted,omitempty"`
	Iops                     *string `json:"iops,omitempty"`
	Name                     string  `json:"name,omitempty"`
	ParentVolumeSnapshotHref *string `json:"parent_volume_snapshot_href,omitempty"`
	PlacementGroupHref       *string `json:"placement_group_href,omitempty"`
	Size                     *string `json:"size,omitempty"`
	VolumeTypeHref           *string `json:"volume_type_href,omitempty"`
}

type VolumeSnapshotParam struct {
	DeploymentHref   *string `json:"deployment_href,omitempty"`
	Description      *string `json:"description,omitempty"`
	Name             string  `json:"name,omitempty"`
	ParentVolumeHref *string `json:"parent_volume_href,omitempty"`
}
