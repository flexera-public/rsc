//************************************************************************//
//                     RightScale API 1.5 go client
//
// Generated Feb 4, 2015 at 6:48pm (PST)
// Command:
// $ api15gen -keep=T -metadata=../../rsapi15/api_data.json -attributes=../../rsapi15/attributes.json -output=../../rsapi15/codegen.go
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Href
type Href string

// Convenience type
type ApiParams map[string]interface{}

// Helper function that merges optional parameters into payload
func mergeOptionals(params, options ApiParams) ApiParams {
	for name, value := range options {
		params[name] = value
	}
	return params
}

/******  Account ******/

type Account struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Permissions []Permission        `json:"permissions,omitempty"`
	Products    []string            `json:"products,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// GET api/accounts/:id(.:format)?
// Show information about a single Account.
func (c *Client) ShowAccount(id string) (*Account, error) {
	var res *Account
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/accounts/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  AccountGroup ******/

// An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.
type AccountGroup struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// GET api/account_groups(.:format)?
// Lists the AccountGroups owned by this Account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexAccountGroups(options ApiParams) ([]AccountGroup, error) {
	var res []AccountGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/account_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/account_groups/:id(.:format)?
// Show information about a single AccountGroup.
// -- Optional parameters:
// 	view
func (c *Client) ShowAccountGroup(id string, options ApiParams) (*AccountGroup, error) {
	var res *AccountGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/account_groups/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Alert ******/

// An Alert represents an AlertSpec bound to a running Instance.
type Alert struct {
	Actions       []string            `json:"actions,omitempty"`
	CreatedAt     *time.Time          `json:"created_at,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	QuenchedUntil *time.Time          `json:"quenched_until,omitempty"`
	Status        string              `json:"status,omitempty"`
	TriggeredAt   *time.Time          `json:"triggered_at,omitempty"`
	UpdatedAt     *time.Time          `json:"updated_at,omitempty"`
}

// POST api/clouds/:cloud_id/instances/:instance_id/alerts/:id/disable(.:format)?
// Disables the Alert indefinitely. Idempotent.
func (c *Client) DisableAlert(cloudId string, id string, instanceId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/alerts/"+id+"/disable", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:instance_id/alerts/:id/enable(.:format)?
// Enables the Alert indefinitely. Idempotent.
func (c *Client) EnableAlert(cloudId string, id string, instanceId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/alerts/"+id+"/enable", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/instances/:instance_id/alerts(.:format)?
// Lists all Alerts.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexAlerts(cloudId string, instanceId string, options ApiParams) ([]Alert, error) {
	var res []Alert
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/alerts", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/clouds/:cloud_id/instances/:instance_id/alerts/:id/quench(.:format)?
// Suppresses the Alert from being triggered for a given time period. Idempotent.
// 	duration: The time period in seconds to suppress Alert from being triggered.
func (c *Client) QuenchAlert(cloudId string, duration string, id string, instanceId string) (string, error) {
	var res string
	payload := ApiParams{
		"duration": duration,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/alerts/"+id+"/quench", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/instances/:instance_id/alerts/:id(.:format)?
// Shows the attributes of a specified Alert.
// -- Optional parameters:
// 	view
func (c *Client) ShowAlert(cloudId string, id string, instanceId string, options ApiParams) (*Alert, error) {
	var res *Alert
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/alerts/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  AlertSpec ******/

// An AlertSpec defines the conditions under which an Alert is triggered and escalated.
// Condition sentence: if &lt;file&gt;.&lt;variable&gt; &lt;condition&gt; '&lt;threshold&gt;' for &lt;duration&gt; min then escalate to '&lt;escalation_name&gt;'.
type AlertSpec struct {
	Actions        []string            `json:"actions,omitempty"`
	Condition      string              `json:"condition,omitempty"`
	CreatedAt      *time.Time          `json:"created_at,omitempty"`
	Description    string              `json:"description,omitempty"`
	Duration       int                 `json:"duration,omitempty"`
	EscalationName string              `json:"escalation_name,omitempty"`
	File           string              `json:"file,omitempty"`
	Links          []map[string]string `json:"links,omitempty"`
	Name           string              `json:"name,omitempty"`
	Threshold      string              `json:"threshold,omitempty"`
	UpdatedAt      *time.Time          `json:"updated_at,omitempty"`
	Variable       string              `json:"variable,omitempty"`
	VoteTag        string              `json:"vote_tag,omitempty"`
	VoteType       string              `json:"vote_type,omitempty"`
}

// POST api/servers/:server_id/alert_specs(.:format)?
// Creates a new AlertSpec with the given parameters.
func (c *Client) CreateAlertSpec(alertSpec *AlertSpecParam, serverId string) (Href, error) {
	var res Href
	payload := ApiParams{
		"alert_spec": alertSpec,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/servers/"+serverId+"/alert_specs", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/servers/:server_id/alert_specs/:id(.:format)?
// Deletes a given AlertSpec.
func (c *Client) DestroyAlertSpec(id string, serverId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/servers/"+serverId+"/alert_specs/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/servers/:server_id/alert_specs(.:format)?
// <no description>
// -- Optional parameters:
// 	filter
// 	view
// 	withInherited: Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.
func (c *Client) IndexAlertSpecs(serverId string, options ApiParams) ([]AlertSpec, error) {
	var res []AlertSpec
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers/"+serverId+"/alert_specs", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/servers/:server_id/alert_specs/:id(.:format)?
// <no description>
// -- Optional parameters:
// 	view
func (c *Client) ShowAlertSpec(id string, serverId string, options ApiParams) (*AlertSpec, error) {
	var res *AlertSpec
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers/"+serverId+"/alert_specs/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/servers/:server_id/alert_specs/:id(.:format)?
// Updates an AlertSpec with the given parameters.
func (c *Client) UpdateAlertSpec(alertSpec *AlertSpecParam2, id string, serverId string) error {
	payload := ApiParams{
		"alert_spec": alertSpec,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/servers/"+serverId+"/alert_specs/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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
	UpdatedAt  *time.Time          `json:"updated_at,omitempty"`
	UserEmail  string              `json:"user_email,omitempty"`
}

// POST api/audit_entries/:id/append(.:format)?
// Updates the summary and appends more details to a given AuditEntry. Each audit entry detail is stored
// as one chunk, the offset determines the location of that chunk within the overall audit entry details section.
// For example, if you create an AuditEntry and append "DEF" at offset 10, and later append
// "ABC" at offset 9, the overall audit entry details will be "ABCDEF". Use the \n character to
// separate details by new lines.
// -- Optional parameters:
// 	detail: The details to be appended to the audit entry record.
// 	notify: The event notification category. Defaults to 'None'.
// 	offset: The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.
// 	summary: The updated summary for the audit entry, maximum length is 255 characters.
func (c *Client) AppendAuditEntry(id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/audit_entries/"+id+"/append", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/audit_entries(.:format)?
// Creates a new AuditEntry with the given parameters.
// -- Optional parameters:
// 	notify: The event notification category. Defaults to 'None'.
// 	userEmail: The email of the user (who created/triggered the audit entry). Only usable with instance role.
func (c *Client) CreateAuditEntry(auditEntry *AuditEntryParam, options ApiParams) (Href, error) {
	var res Href
	payload := mergeOptionals(ApiParams{
		"audit_entry": auditEntry,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/audit_entries", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// GET api/audit_entries/:id/detail(.:format)?
// shows the details of a given AuditEntry.
// Note that the media type of the response is simply text.
func (c *Client) DetailAuditEntry(id string) (string, error) {
	var res string
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/audit_entries/"+id+"/detail", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/audit_entries(.:format)?
// Lists AuditEntries of the account. Due to the potentially large number of audit entries, a start and end date must
// be provided during an index call to limit the search. The format of the dates must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
// 2011/07/11 00:00:00 +0000.
// A maximum of 1000 records will be returned by each index call.

// Using the available filters, one can select or group which audit entries to retrieve.
// 	endDate: The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).
// 	limit: Limit the audit entries to this number. The limit should >= 1 and <= 1000
// 	startDate: The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexAuditEntries(endDate string, limit string, startDate string, options ApiParams) ([]AuditEntry, error) {
	var res []AuditEntry
	payload := ApiParams{
		"end_date":   endDate,
		"limit":      limit,
		"start_date": startDate,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/audit_entries", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/audit_entries/:id(.:format)?
// Lists the attributes of a given audit entry.
// -- Optional parameters:
// 	view
func (c *Client) ShowAuditEntry(id string, options ApiParams) (*AuditEntry, error) {
	var res *AuditEntry
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/audit_entries/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/audit_entries/:id(.:format)?
// Updates the summary of a given AuditEntry.
// -- Optional parameters:
// 	notify: The event notification category. Defaults to 'None'.
func (c *Client) UpdateAuditEntry(auditEntry *AuditEntryParam2, id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{
		"audit_entry": auditEntry,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/audit_entries/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  Backup ******/

type Backup struct {
	Actions             []string            `json:"actions,omitempty"`
	Committed           bool                `json:"committed,omitempty"`
	Completed           bool                `json:"completed,omitempty"`
	CreatedAt           *time.Time          `json:"created_at,omitempty"`
	Description         string              `json:"description,omitempty"`
	FromMaster          bool                `json:"from_master,omitempty"`
	Lineage             string              `json:"lineage,omitempty"`
	Links               []map[string]string `json:"links,omitempty"`
	Name                string              `json:"name,omitempty"`
	VolumeSnapshotCount int                 `json:"volume_snapshot_count,omitempty"`
	VolumeSnapshots     []VolumeSnapshot    `json:"volume_snapshots,omitempty"`
}

// POST api/backups/cleanup(.:format)?
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

// 	keepLast: The number of backups that should be kept.
// 	lineage: The lineage of the backups that are to be cleaned-up.
// -- Optional parameters:
// 	cloudHref: Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.
// 	dailies: The number of daily backups(the latest one in each day) that should be kept.
// 	monthlies: The number of monthly backups(the latest one in each month) that should be kept.
// 	weeklies: The number of weekly backups(the latest one in each week) that should be kept.
// 	yearlies: The number of yearly backups(the latest one in each year) that should be kept.
func (c *Client) CleanupBackup(keepLast string, lineage string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{

		"keep_last": keepLast,
		"lineage":   lineage,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/backups/cleanup", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/backups(.:format)?
// Takes in an array of volumeattachmenthrefs and takes a snapshot of each.
// The volumeattachmenthrefs must belong to the same instance.
func (c *Client) CreateBackup(backup *BackupParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"backup": backup,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/backups", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/backups/:id(.:format)?
// Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.
func (c *Client) DestroyBackup(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/backups/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/backups(.:format)?
// Lists all of the backups with the given lineage tag. Filters can be used to search for a particular backup. If the
// 'latest_before' filter is set, only one backup is returned (the latest backup before the given timestamp).

// To get the latest completed backup, the 'completed' filter should be set to 'true' and the 'latest_before' filter
// should be set to the current timestamp. The format of the timestamp must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
// 2011/07/11 00:00:00 +0000.

// To get the latest completed backup just before, say 25 June 2009, then the 'completed' filter
// should be set to 'true' and the 'latest_before' filter should be set to 2009/06/25 00:00:00 +0000.
// 	lineage: Backups belonging to this lineage.
// -- Optional parameters:
// 	filter
func (c *Client) IndexBackups(lineage string, options ApiParams) ([]Backup, error) {
	var res []Backup
	payload := ApiParams{
		"lineage": lineage,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/backups", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/backups/:id/restore(.:format)?
// Restores the given Backup.
// This call will:

// create the required number of Volumes from the volume_snapshots_hrefs in the given Backup,
// attach them to the given Instance at the device specified in the Snapshot. If the devices are already being used
//    on the Instance, the Task will denote that the restore has failed.

// 	instanceHref: The instance href that the backup will be restored to.
// -- Optional parameters:
// 	backup
func (c *Client) RestoreBackup(id string, instanceHref string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{

		"instance_href": instanceHref,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/backups/"+id+"/restore", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/backups/:id(.:format)?
// Lists the attributes of a given backup
func (c *Client) ShowBackup(id string) (*Backup, error) {
	var res *Backup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/backups/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/backups/:id(.:format)?
// Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.
func (c *Client) UpdateBackup(backup *BackupParam3, id string) error {
	payload := ApiParams{
		"backup": backup,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/backups/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  ChildAccount ******/

type ChildAccount struct {
}

// POST api/child_accounts(.:format)?
// Create an enterprise ChildAccount for this Account. The User will by default get an 'admin' role
// on the ChildAccount to enable him/her to add, delete Users and Permissions.

// For more information on the valid values for 'cluster_href', refer to the following:

// http://support.rightscale.com/12-Guides/RightScale_API_1.5/Examples/ChildAccounts/Create

func (c *Client) CreateChildAccount(childAccount *ChildAccountParam) (map[string]interface{}, error) {
	var res map[string]interface{}
	payload := ApiParams{
		"child_account": childAccount,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/child_accounts", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/child_accounts(.:format)?
// Lists the enterprise ChildAccounts available for this Account.
// -- Optional parameters:
// 	filter
func (c *Client) IndexChildAccounts(options ApiParams) ([]map[string]string, error) {
	var res []map[string]string
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/child_accounts", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/accounts/:id(.:format)?
// Update an enterprise ChildAccount for this Account.
func (c *Client) UpdateChildAccount(childAccount *ChildAccountParam2, id string) error {
	payload := ApiParams{
		"child_account": childAccount,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/accounts/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// GET api/clouds(.:format)?
// Lists the clouds available to this account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexClouds(options ApiParams) ([]Cloud, error) {
	var res []Cloud
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:id(.:format)?
// Show information about a single cloud.
// -- Optional parameters:
// 	view
func (c *Client) ShowCloud(id string, options ApiParams) (*Cloud, error) {
	var res *Cloud
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  CloudAccount ******/

// Represents a Cloud Account (An association between the account and a cloud).
type CloudAccount struct {
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

// POST api/cloud_accounts(.:format)?
// Create a CloudAccount by passing in the respective credentials for each cloud.

// For more information on the specific parameters for each cloud, refer to the following:

// http://support.rightscale.com/12-Guides/RightScale_API_1.5/Examples/Cloud_Accounts/Create_Cloud_Accounts

func (c *Client) CreateCloudAccount(cloudAccount *CloudAccountParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"cloud_account": cloudAccount,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/cloud_accounts", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/cloud_accounts/:id(.:format)?
// Delete a CloudAccount.
func (c *Client) DestroyCloudAccount(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/cloud_accounts/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/cloud_accounts(.:format)?
// Lists the CloudAccounts (non-aws) available to this Account.
func (c *Client) IndexCloudAccounts() ([]CloudAccount, error) {
	var res []CloudAccount
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cloud_accounts", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/cloud_accounts/:id(.:format)?

func (c *Client) ShowCloudAccount(id string) (*CloudAccount, error) {
	var res *CloudAccount
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cloud_accounts/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Cookbook ******/

// Represents a given instance of a single cookbook.
type Cookbook struct {
	Actions           []string            `json:"actions,omitempty"`
	CreatedAt         *time.Time          `json:"created_at,omitempty"`
	DownloadUrl       string              `json:"download_url,omitempty"`
	Id                string              `json:"id,omitempty"`
	Links             []map[string]string `json:"links,omitempty"`
	Metadata          string              `json:"metadata,omitempty"`
	Name              string              `json:"name,omitempty"`
	Namespace         string              `json:"namespace,omitempty"`
	SourceInfoSummary string              `json:"source_info_summary,omitempty"`
	State             string              `json:"state,omitempty"`
	UpdatedAt         *time.Time          `json:"updated_at,omitempty"`
	Version           string              `json:"version,omitempty"`
}

// DELETE api/cookbooks/:id(.:format)?
// Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.
func (c *Client) DestroyCookbook(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/cookbooks/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/cookbooks/:id/follow(.:format)?
// Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.
// 	value: Indicates if this action should follow (true) or unfollow (false) a Cookbook.
func (c *Client) FollowCookbook(id string, value string) error {
	payload := ApiParams{
		"value": value,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/cookbooks/"+id+"/follow", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/cookbooks/:id/freeze(.:format)?
// Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.
// 	value: Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.
func (c *Client) FreezeCookbook(id string, value string) error {
	payload := ApiParams{
		"value": value,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/cookbooks/"+id+"/freeze", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/cookbooks(.:format)?
// Lists the Cookbooks available to this account.

// The extended_designer view is only available to accounts with the designer permission.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexCookbooks(options ApiParams) ([]Cookbook, error) {
	var res []Cookbook
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/cookbooks/:id/obsolete(.:format)?
// Marks a Cookbook as obsolete (or un-obsolete).
// 	value: Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.
func (c *Client) ObsoleteCookbook(id string, value string) error {
	payload := ApiParams{
		"value": value,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/cookbooks/"+id+"/obsolete", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/cookbooks/:id(.:format)?
// Show information about a single Cookbook.

// The extended_designer view is only available to accounts with the designer permission.
// -- Optional parameters:
// 	view
func (c *Client) ShowCookbook(id string, options ApiParams) (*Cookbook, error) {
	var res *Cookbook
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  CookbookAttachment ******/

// Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.
type CookbookAttachment struct {
	Actions    []string            `json:"actions,omitempty"`
	Dependency bool                `json:"dependency,omitempty"`
	Id         string              `json:"id,omitempty"`
	Links      []map[string]string `json:"links,omitempty"`
}

// POST api/cookbooks/:cookbook_id/cookbook_attachments(.:format)?
// Attach a cookbook to a given resource.
// -- Optional parameters:
// 	cookbookAttachment
func (c *Client) CreateCookbookAttachment(cookbookId string, options ApiParams) (Href, error) {
	var res Href
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/cookbooks/"+cookbookId+"/cookbook_attachments", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/cookbooks/:cookbook_id/cookbook_attachments/:id(.:format)?
// Detach a cookbook from a given resource.
func (c *Client) DestroyCookbookAttachment(cookbookId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/cookbooks/"+cookbookId+"/cookbook_attachments/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/cookbooks/:cookbook_id/cookbook_attachments(.:format)?
// Lists Cookbook Attachments.
// -- Optional parameters:
// 	view
func (c *Client) IndexCookbookAttachments(cookbookId string, options ApiParams) ([]CookbookAttachment, error) {
	var res []CookbookAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks/"+cookbookId+"/cookbook_attachments", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/server_templates/:server_template_id/cookbook_attachments/multi_attach(.:format)?
// Attach multiple cookbooks to a given resource.
func (c *Client) MultiAttachCookbookAttachments(cookbookAttachments *CookbookAttachments, serverTemplateId string) error {
	payload := ApiParams{
		"cookbook_attachments": cookbookAttachments,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+serverTemplateId+"/cookbook_attachments/multi_attach", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_templates/:server_template_id/cookbook_attachments/multi_detach(.:format)?
// Detach multiple cookbooks from a given resource.
func (c *Client) MultiDetachCookbookAttachments(cookbookAttachments *CookbookAttachments2, serverTemplateId string) error {
	payload := ApiParams{
		"cookbook_attachments": cookbookAttachments,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+serverTemplateId+"/cookbook_attachments/multi_detach", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/cookbooks/:cookbook_id/cookbook_attachments/:id(.:format)?
// Displays information about a single cookbook attachment to a ServerTemplate.
// -- Optional parameters:
// 	view
func (c *Client) ShowCookbookAttachment(cookbookId string, id string, options ApiParams) (*CookbookAttachment, error) {
	var res *CookbookAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks/"+cookbookId+"/cookbook_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Credential ******/

// A Credential provides an abstraction for sensitive textual information,
// such as passphrases or cloud credentials, whose value should be encrypted
// when stored in the database and not generally shown in the UI or through the
// API. Credentials may then be used as inputs of type "Cred" in RightScripts
// or Chef recipes. NOTE: Credential values may be updated through the API, but
// values cannot be retrieved via the API.
type Credential struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
	Value       string              `json:"value,omitempty"`
}

// POST api/credentials(.:format)?
// Creates a new Credential with the given parameters.
func (c *Client) CreateCredential(credential *CredentialParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"credential": credential,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/credentials", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/credentials/:id(.:format)?
// Deletes a Credential.
func (c *Client) DestroyCredential(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/credentials/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/credentials(.:format)?
// Lists the Credentials available to this account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexCredentials(options ApiParams) ([]Credential, error) {
	var res []Credential
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/credentials", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/credentials/:id(.:format)?
// Show information about a single Credential. NOTE: Credential values may be updated through the API, but values cannot be retrieved via the API.
// -- Optional parameters:
// 	view
func (c *Client) ShowCredential(id string, options ApiParams) (*Credential, error) {
	var res *Credential
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/credentials/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/credentials/:id(.:format)?
// Updates attributes of a Credential.
func (c *Client) UpdateCredential(credential *CredentialParam2, id string) error {
	payload := ApiParams{
		"credential": credential,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/credentials/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// GET api/clouds/:cloud_id/datacenters(.:format)?
// Lists all Datacenters for a particular cloud.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexDatacenters(cloudId string, options ApiParams) ([]Datacenter, error) {
	var res []Datacenter
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/datacenters", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/datacenters/:id(.:format)?
// Displays information about a single Datacenter.
// -- Optional parameters:
// 	view
func (c *Client) ShowDatacenter(cloudId string, id string, options ApiParams) (*Datacenter, error) {
	var res *Datacenter
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/datacenters/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/deployments/:id/clone(.:format)?
// Clones a given deployment.
// -- Optional parameters:
// 	deployment
func (c *Client) CloneDeployment(id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/deployments/"+id+"/clone", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/deployments(.:format)?
// Creates a new deployment with the given parameters.
func (c *Client) CreateDeployment(deployment *DeploymentParam2) (Href, error) {
	var res Href
	payload := ApiParams{
		"deployment": deployment,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/deployments", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/deployments/:id(.:format)?
// Deletes a given deployment.
func (c *Client) DestroyDeployment(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/deployments/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/deployments(.:format)?
// Lists deployments of the account.

// Using the available filters, one can select or group which deployments to retrieve.
// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexDeployments(options ApiParams) ([]Deployment, error) {
	var res []Deployment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/deployments", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/deployments/:id/lock(.:format)?
// Locks a given deployment. Idempotent.
// Locking prevents servers from being deleted or moved from the deployment.
// Other actions such as adding servers or renaming the deployment are still allowed.
func (c *Client) LockDeployment(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/deployments/"+id+"/lock", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/deployments/:id/servers
// Lists the servers belonging to this deployment. This call is equivalent to servers#index call, where the servers returned will
// automatically be filtered by this deployment. See servers#index for details on other options and parameters.
func (c *Client) ServersDeployment(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/deployments/"+id+"/servers", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/deployments/:id(.:format)?
// Lists the attributes of a given deployment.

// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// 	view
func (c *Client) ShowDeployment(id string, options ApiParams) (*Deployment, error) {
	var res *Deployment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/deployments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/deployments/:id/unlock(.:format)?
// Unlocks a given deployment. Idempotent.
func (c *Client) UnlockDeployment(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/deployments/"+id+"/unlock", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// PUT api/deployments/:id(.:format)?
// Updates attributes of a given deployment.
func (c *Client) UpdateDeployment(deployment *DeploymentParam, id string) error {
	payload := ApiParams{
		"deployment": deployment,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/deployments/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  HealthCheck ******/

type HealthCheck struct {
}

// GET api/health-check/
// Check health of RightApi controllers
func (c *Client) IndexHealthCheck() ([]map[string]string, error) {
	var res []map[string]string
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/health-check/", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  IdentityProvider ******/

// An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account,
// and is trusted by the RightScale dashboard to authenticate your enterprise's end users.
// To register an Identity Provider, contact your account manager.
type IdentityProvider struct {
	Actions       []string            `json:"actions,omitempty"`
	CreatedAt     *time.Time          `json:"created_at,omitempty"`
	DiscoveryHint string              `json:"discovery_hint,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	Name          string              `json:"name,omitempty"`
	UpdatedAt     *time.Time          `json:"updated_at,omitempty"`
}

// GET api/identity_providers(.:format)?
// Lists the identity providers associated with this enterprise account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexIdentityProviders(options ApiParams) ([]IdentityProvider, error) {
	var res []IdentityProvider
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/identity_providers", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/identity_providers/:id(.:format)?
// Show the specified identity provider, if associated with this enterprise account.
// -- Optional parameters:
// 	view
func (c *Client) ShowIdentityProvider(id string, options ApiParams) (*IdentityProvider, error) {
	var res *IdentityProvider
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/identity_providers/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/images(.:format)?
// Lists all Images for the given Cloud.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexImages(cloudId string, options ApiParams) ([]Image, error) {
	var res []Image
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/images", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/images/:id(.:format)?
// Shows information about a single Image.
// -- Optional parameters:
// 	view
func (c *Client) ShowImage(cloudId string, id string, options ApiParams) (*Image, error) {
	var res *Image
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/images/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Input ******/

// Inputs help extract dynamic information, usually specified at runtime, from repeatable configuration operations that can be codified.
// Inputs are variables defined in and used by RightScripts/Recipes. The two main attributes of an input are 'name' and 'value'. The 'name'
// identifies the input and the 'value', although a string encodes what type it is. It could be a text encoded as 'text:myvalue' or a credential
// encoded as 'cred:MY_CRED' or a key etc. Please see support.rightscale.com for more info on input hierarchies and their different types.
type Input struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// GET api/clouds/:cloud_id/instances/:instance_id/inputs(.:format)?
// Retrieves the full list of existing inputs of the specified resource.
// -- Optional parameters:
// 	view
func (c *Client) IndexInputs(cloudId string, instanceId string, options ApiParams) ([]Input, error) {
	var res []Input
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/inputs", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/clouds/:cloud_id/instances/:instance_id/inputs/multi_update(.:format)?
// Performs a bulk update of inputs on the specified resource.

// If an input exists with the same name, its value will be updated. If an input does not
// exist with a specified name, it will be ignored.

// Input values are represented as strings.

// There are two notations for inputs:

// 1.0 notation - The deprecated notation used in API 1.0 and in 1.5
//     2.0 notation - The new notation that is partially supported in API 1.5, and will
//         be the only notation supported in 2.0

// Although the two notations are similar, they have a few important differences, in particular:

//   With 2.0 notation, values MUST begin with a prefix identifying their type, followed
//   by a colon (example: "text:foo"). With 1.0 notation, unprefixed values are generally
//   taken to be text-type.
//   With 2.0 notation, a sentinel value "inherit" is used to express that an input
//   should use an inherited value. With 1.0 notation the empty string was used to express
//   the same thing. (Due to requirement 1, empty string is no longer a valid input.)
//   With 2.0 notation, each element of an array is an entire input value; arrays can
//   contain cred, env, or even other arrays. With 1.0 notation, array elements are
//   implicitly text values and there is no way to specify anything else.Note that the UI
//   does not support complex-valued arrays; please use this feature with caution!

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
func (c *Client) MultiUpdateInputs(cloudId string, inputs map[string]string, instanceId string) error {
	payload := ApiParams{
		"inputs": inputs,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/inputs/multi_update", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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
	CreatedAt                *time.Time          `json:"created_at,omitempty"`
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
	TerminatedAt             *time.Time          `json:"terminated_at,omitempty"`
	UpdatedAt                *time.Time          `json:"updated_at,omitempty"`
	UserData                 string              `json:"user_data,omitempty"`
}

// POST api/clouds/:cloud_id/instances(.:format)?
// Creates and launches a raw instance using the provided parameters.
func (c *Client) CreateInstance(cloudId string, instance *InstanceParam4) (Href, error) {
	var res Href
	payload := ApiParams{
		"instance": instance,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// GET api/clouds/:cloud_id/instances(.:format)?
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
// 	filter
// 	view
func (c *Client) IndexInstances(cloudId string, options ApiParams) ([]Instance, error) {
	var res []Instance
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/clouds/:cloud_id/instances/:id/launch(.:format)?
// Launches an instance using the parameters that this instance has been configured with.

// Note that this action can only be performed in "next" instances, and not on instances that are already running.
// -- Optional parameters:
// 	apiBehavior: When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'
// 	inputs
func (c *Client) LaunchInstance(cloudId string, id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/launch", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:id/lock(.:format)?

func (c *Client) LockInstance(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/lock", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/multi_run_executable(.:format)?
// Runs a script or a recipe in the running instances.

// This is an asynchronous function, which returns immediately after queuing the executable for execution.
// Status of the execution can be tracked at the URL returned in the "Location" header.
// -- Optional parameters:
// 	filter
// 	ignoreLock: Specifies the ability to ignore the lock(s) on the Instance(s).
// 	inputs
// 	recipeName: The name of the recipe to be run.
// 	rightScriptHref: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.
func (c *Client) MultiRunExecutableInstances(cloudId string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/multi_run_executable", body)
	if err != nil {
		return err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/multi_terminate(.:format)?
// Terminates running instances.
// Either a filter or the parameter 'terminate_all' must be provided.
// -- Optional parameters:
// 	filter
// 	terminateAll: Specifies the ability to terminate all instances.
func (c *Client) MultiTerminateInstances(cloudId string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/multi_terminate", body)
	if err != nil {
		return err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:id/reboot(.:format)?
// Reboot a running instance.

// Note that this action can only succeed if the instance is running. One cannot reboot instances of type "next".
func (c *Client) RebootInstance(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/reboot", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:id/run_executable(.:format)?
// Runs a script or a recipe in the running instance.

// This is an asynchronous function, which returns immediately after queuing the executable for execution.
// Status of the execution can be tracked at the URL returned in the "Location" header.
// Note that this can only be performed on running instances.
// -- Optional parameters:
// 	ignoreLock: Specifies the ability to ignore the lock on the Instance.
// 	inputs
// 	recipeName: The name of the recipe to run.
// 	rightScriptHref: The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.
func (c *Client) RunExecutableInstance(cloudId string, id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/run_executable", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:id/set_custom_lodgement(.:format)?
// This method is deprecated.  Please use InstanceCustomLodgement.
// 	quantity: At least one name/value pair must be specified. Currently, a maximum of 2 name/value pairs is supported.
// 	timeframe: The timeframe (either a month or a single day) for which the quantity value is valid (currently for the PDT timezone only).
func (c *Client) SetCustomLodgementInstance(cloudId string, id string, quantity []*Quantity, timeframe string) error {
	payload := ApiParams{
		"quantity":  quantity,
		"timeframe": timeframe,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/set_custom_lodgement", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/instances/:id(.:format)?
// Shows attributes of a single instance.

// The 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// 	view
func (c *Client) ShowInstance(cloudId string, id string, options ApiParams) (*Instance, error) {
	var res *Instance
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/clouds/:cloud_id/instances/:id/start(.:format)?
// Starts an instance that has been stopped, resuming it to its previously saved volume state.

// After an instance is started, the reference to your instance will have a different id.

// The new id can be found by performing an index query with the appropriate filters on the
// Instances resource, performing a show action on the Server resource for Server Instances, or
// performing a current_instances action on the ServerArray resource for ServerArray Instances.
func (c *Client) StartInstance(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/start", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:id/stop(.:format)?
// Stores the instance's current volume state to resume later using the 'start' action.

// After an instance is stopped, the reference to your instance will have a different id.

// The new id can be found by performing an index query with the appropriate filters on the
// Instances resource, performing a show action on the Server resource for Server Instances, or performing a current_instances action on the ServerArray resource for ServerArray Instances.
func (c *Client) StopInstance(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/stop", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:id/terminate(.:format)?
// Terminates a running instance.

// Note that this action can succeed only if the instance is running. One cannot terminate instances of type "next".
func (c *Client) TerminateInstance(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/terminate", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/clouds/:cloud_id/instances/:id/unlock(.:format)?

func (c *Client) UnlockInstance(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id+"/unlock", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// PUT api/clouds/:cloud_id/instances/:id(.:format)?
// Updates attributes of a single instance.
func (c *Client) UpdateInstance(cloudId string, id string, instance *InstanceParam5) error {
	payload := ApiParams{
		"instance": instance,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  InstanceCustomLodgement ******/

// An InstanceCustomLodgement represents a way to create custom reports about a specific instance with a user defined quantity.  Replaces the legacy Instances#setcustomlodgement interface.
type InstanceCustomLodgement struct {
	AccountOwner                         string                   `json:"account_owner,omitempty"`
	Actions                              []string                 `json:"actions,omitempty"`
	EndAt                                *time.Time               `json:"end_at,omitempty"`
	Links                                []map[string]string      `json:"links,omitempty"`
	Quantity                             []map[string]interface{} `json:"quantity,omitempty"`
	ResourceBillingEndAt                 *time.Time               `json:"resource_billing_end_at,omitempty"`
	ResourceBillingStartAt               *time.Time               `json:"resource_billing_start_at,omitempty"`
	ResourceInstanceType                 string                   `json:"resource_instance_type,omitempty"`
	ResourceLaunchedBy                   string                   `json:"resource_launched_by,omitempty"`
	ResourceTemplateLibraryHref          string                   `json:"resource_template_library_href,omitempty"`
	ResourceTemplateName                 string                   `json:"resource_template_name,omitempty"`
	ResourceTemplatePublishedByAccountId string                   `json:"resource_template_published_by_account_id,omitempty"`
	ResourceUid                          string                   `json:"resource_uid,omitempty"`
	StartAt                              *time.Time               `json:"start_at,omitempty"`
}

// POST api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements(.:format)?
// Create a lodgement with the quantity and timeframe specified.
// 	quantity: At least one name/value pair must be specified. Currently, a maximum of 2 name/value pairs is supported.
// 	timeframe: The time-frame (either a month "YYYY_MM" or a single day "YYYY_MM_DD") for which the quantity value is valid (currently for the PDT timezone only).
func (c *Client) CreateInstanceCustomLodgement(cloudId string, instanceId string, quantity []*Quantity, timeframe string) (Href, error) {
	var res Href
	payload := ApiParams{
		"quantity":  quantity,
		"timeframe": timeframe,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/instance_custom_lodgements", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id(.:format)?
// Destroy the specified lodgement.
func (c *Client) DestroyInstanceCustomLodgement(cloudId string, id string, instanceId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/instance_custom_lodgements/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements(.:format)?
// List InstanceCustomLodgements of a given cloud and instance.
// -- Optional parameters:
// 	view
func (c *Client) IndexInstanceCustomLodgements(cloudId string, instanceId string, options ApiParams) ([]InstanceCustomLodgement, error) {
	var res []InstanceCustomLodgement
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/instance_custom_lodgements", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id(.:format)?
// Show the specified lodgement.
func (c *Client) ShowInstanceCustomLodgement(cloudId string, id string, instanceId string) (*InstanceCustomLodgement, error) {
	var res *InstanceCustomLodgement
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/instance_custom_lodgements/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id(.:format)?
// Update a lodgement with the quantity specified.
// 	quantity: At least one name/value pair must be specified. Currently, a maximum of 2 name/value pairs is supported.
func (c *Client) UpdateInstanceCustomLodgement(cloudId string, id string, instanceId string, quantity []*Quantity2) error {
	payload := ApiParams{
		"quantity": quantity,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/instance_custom_lodgements/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// GET api/clouds/:cloud_id/instance_types(.:format)?
// Lists instance types.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexInstanceTypes(cloudId string, options ApiParams) ([]InstanceType, error) {
	var res []InstanceType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instance_types", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/instance_types/:id(.:format)?
// Displays information about a single Instance type.
// -- Optional parameters:
// 	view
func (c *Client) ShowInstanceType(cloudId string, id string, options ApiParams) (*InstanceType, error) {
	var res *InstanceType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instance_types/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  IpAddress ******/

// An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.
type IpAddress struct {
	Address   string              `json:"address,omitempty"`
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	Domain    string              `json:"domain,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	Name      string              `json:"name,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

// POST api/clouds/:cloud_id/ip_addresses(.:format)?
// Creates a new IpAddress with the given parameters.
func (c *Client) CreateIpAddress(cloudId string, ipAddress *IpAddressParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"ip_address": ipAddress,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/ip_addresses/:id(.:format)?
// Deletes a given IpAddress.
func (c *Client) DestroyIpAddress(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/ip_addresses(.:format)?
// Lists the IpAddresses available to this account.
// -- Optional parameters:
// 	filter
func (c *Client) IndexIpAddresses(cloudId string, options ApiParams) ([]IpAddress, error) {
	var res []IpAddress
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/ip_addresses/:id(.:format)?
// Show information about a single IpAddress.
func (c *Client) ShowIpAddress(cloudId string, id string) (*IpAddress, error) {
	var res *IpAddress
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/clouds/:cloud_id/ip_addresses/:id(.:format)?
// Updates attributes of a given IpAddress.
func (c *Client) UpdateIpAddress(cloudId string, id string, ipAddress *IpAddressParam2) error {
	payload := ApiParams{
		"ip_address": ipAddress,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  IpAddressBinding ******/

// An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.
// The IpAddress is bound immediately for a current instance, or on launch for a next instance.
// It also allows specifying port forwarding rules for that particular IpAddress and Instance pair.
type IpAddressBinding struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	PrivatePort string              `json:"private_port,omitempty"`
	Protocol    string              `json:"protocol,omitempty"`
	PublicPort  string              `json:"public_port,omitempty"`
	Recurring   bool                `json:"recurring,omitempty"`
}

// POST api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings(.:format)?
// Creates an ip address binding which attaches a specified IpAddress resource
// to a specified instance, and also allows for configuration of port forwarding
// rules. If the instance specified is a current (running) instance, a one-time
// IpAddressBinding will be created. If the instance is a next instance, then
// a recurring IpAddressBinding is created, which will cause the IpAddress to
// be bound each time the incarnator boots.
func (c *Client) CreateIpAddressBinding(cloudId string, ipAddressBinding *IpAddressBindingParam, ipAddressId string) (Href, error) {
	var res Href
	payload := ApiParams{
		"ip_address_binding": ipAddressBinding,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+ipAddressId+"/ip_address_bindings", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id(.:format)?
// <no description>
func (c *Client) DestroyIpAddressBinding(cloudId string, id string, ipAddressId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+ipAddressId+"/ip_address_bindings/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings(.:format)?
// Lists the ip address bindings available to this account.
// -- Optional parameters:
// 	filter
func (c *Client) IndexIpAddressBindings(cloudId string, ipAddressId string, options ApiParams) ([]IpAddressBinding, error) {
	var res []IpAddressBinding
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+ipAddressId+"/ip_address_bindings", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id(.:format)?
// Show information about a single ip address binding.
func (c *Client) ShowIpAddressBinding(cloudId string, id string, ipAddressId string) (*IpAddressBinding, error) {
	var res *IpAddressBinding
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+ipAddressId+"/ip_address_bindings/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  MonitoringMetric ******/

// A monitoring metric is a stream of data that is captured in an instance. Metrics can be monitored, graphed and can be used as the basis for triggering alerts.
type MonitoringMetric struct {
	Actions   []string            `json:"actions,omitempty"`
	GraphHref string              `json:"graph_href,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	Plugin    string              `json:"plugin,omitempty"`
	View      string              `json:"view,omitempty"`
}

// GET api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id/data(.:format)?
// Gives the raw monitoring data for a particular metric. The response will include different variables
// associated with that metric and the data points for each of those variables.

// To get the data for a certain duration, for e.g. for the last 10 minutes(600 secs), provide the variables
// start="-600" and end="0".
// 	end: An integer number of seconds from current time. e.g. -150 or 0
// 	start: An integer number of seconds from current time. e.g. -300
func (c *Client) DataMonitoringMetric(cloudId string, end string, id string, instanceId string, start string) (map[string]string, error) {
	var res map[string]string
	payload := ApiParams{
		"end":   end,
		"start": start,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/monitoring_metrics/"+id+"/data", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics(.:format)?
// Lists the monitoring metrics available for the instance and their corresponding graph hrefs.
// Making a request to the graph_href will return a png image corresponding to that monitoring metric.
// -- Optional parameters:
// 	filter
// 	period: The time scale for which the graph is generated. Default is 'day'
// 	size: The size of the graph to be generated. Default is 'small'.
// 	title: The title of the graph.
// 	tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.
func (c *Client) IndexMonitoringMetrics(cloudId string, instanceId string, options ApiParams) ([]MonitoringMetric, error) {
	var res []MonitoringMetric
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/monitoring_metrics", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id(.:format)?
// Shows attributes of a single monitoring metric.
// Making a request to the graph_href will return a png image corresponding to that monitoring metric.
// -- Optional parameters:
// 	period: The time scale for which the graph is generated. Default is 'day'.
// 	size: The size of the graph to be generated. Default is 'small'.
// 	title: The title of the graph.
// 	tz: The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences.
func (c *Client) ShowMonitoringMetric(cloudId string, id string, instanceId string, options ApiParams) (*MonitoringMetric, error) {
	var res *MonitoringMetric
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/monitoring_metrics/"+id, body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/multi_cloud_images/:id/clone(.:format)?
// Clones a given MultiCloudImage.
func (c *Client) CloneMultiCloudImage(id string, multiCloudImage *MultiCloudImageParam) error {
	payload := ApiParams{
		"multi_cloud_image": multiCloudImage,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/multi_cloud_images/"+id+"/clone", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/multi_cloud_images/:id/commit(.:format)?
// Commits a given MultiCloudImage. Only HEAD revisions can be committed.
// 	commitMessage: The message associated with the commit.
func (c *Client) CommitMultiCloudImage(commitMessage string, id string) error {
	payload := ApiParams{
		"commit_message": commitMessage,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/multi_cloud_images/"+id+"/commit", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_templates/:server_template_id/multi_cloud_images(.:format)?
// Creates a new MultiCloudImage with the given parameters.
func (c *Client) CreateMultiCloudImage(multiCloudImage *MultiCloudImageParam2, serverTemplateId string) (Href, error) {
	var res Href
	payload := ApiParams{
		"multi_cloud_image": multiCloudImage,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+serverTemplateId+"/multi_cloud_images", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/server_templates/:server_template_id/multi_cloud_images/:id(.:format)?
// Deletes a given MultiCloudImage.
func (c *Client) DestroyMultiCloudImage(id string, serverTemplateId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/server_templates/"+serverTemplateId+"/multi_cloud_images/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/server_templates/:server_template_id/multi_cloud_images(.:format)?
// Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.
// -- Optional parameters:
// 	filter
func (c *Client) IndexMultiCloudImages(serverTemplateId string, options ApiParams) ([]MultiCloudImage, error) {
	var res []MultiCloudImage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+serverTemplateId+"/multi_cloud_images", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/server_templates/:server_template_id/multi_cloud_images/:id(.:format)?
// Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.
func (c *Client) ShowMultiCloudImage(id string, serverTemplateId string) (*MultiCloudImage, error) {
	var res *MultiCloudImage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+serverTemplateId+"/multi_cloud_images/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/server_templates/:server_template_id/multi_cloud_images/:id(.:format)?
// Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).
// Currently, the attributes you can update are only the 'direct' attributes of a server template.
func (c *Client) UpdateMultiCloudImage(id string, multiCloudImage *MultiCloudImageParam, serverTemplateId string) error {
	payload := ApiParams{
		"multi_cloud_image": multiCloudImage,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/server_templates/"+serverTemplateId+"/multi_cloud_images/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// POST api/multi_cloud_images/:multi_cloud_image_id/settings(.:format)?
// Creates a new setting for an existing MultiCloudImage.
func (c *Client) CreateMultiCloudImageSetting(multiCloudImageId string, multiCloudImageSetting *MultiCloudImageSettingParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"multi_cloud_image_setting": multiCloudImageSetting,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/multi_cloud_images/"+multiCloudImageId+"/settings", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/multi_cloud_images/:multi_cloud_image_id/settings/:id(.:format)?
// Deletes a MultiCloudImage setting.
func (c *Client) DestroyMultiCloudImageSetting(id string, multiCloudImageId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/multi_cloud_images/"+multiCloudImageId+"/settings/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/multi_cloud_images/:multi_cloud_image_id/settings(.:format)?
// Lists the MultiCloudImage settings.
// -- Optional parameters:
// 	filter
func (c *Client) IndexMultiCloudImageSettings(multiCloudImageId string, options ApiParams) ([]MultiCloudImageSetting, error) {
	var res []MultiCloudImageSetting
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/multi_cloud_images/"+multiCloudImageId+"/settings", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/multi_cloud_images/:multi_cloud_image_id/settings/:id(.:format)?
// Show information about a single MultiCloudImage setting.
func (c *Client) ShowMultiCloudImageSetting(id string, multiCloudImageId string) (*MultiCloudImageSetting, error) {
	var res *MultiCloudImageSetting
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/multi_cloud_images/"+multiCloudImageId+"/settings/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/multi_cloud_images/:multi_cloud_image_id/settings/:id(.:format)?
// Updates a settings for a MultiCLoudImage.
func (c *Client) UpdateMultiCloudImageSetting(id string, multiCloudImageId string, multiCloudImageSetting *MultiCloudImageSettingParam2) error {
	payload := ApiParams{
		"multi_cloud_image_setting": multiCloudImageSetting,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/multi_cloud_images/"+multiCloudImageId+"/settings/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// POST api/networks(.:format)?
// Creates a new network.
func (c *Client) CreateNetwork(network *NetworkParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"network": network,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/networks", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/networks/:id(.:format)?
// Deletes the given network(s).
func (c *Client) DestroyNetwork(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/networks/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/networks(.:format)?
// Lists networks in this account.
// -- Optional parameters:
// 	filter
func (c *Client) IndexNetworks(options ApiParams) ([]Network, error) {
	var res []Network
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/networks", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/networks/:id(.:format)?
// Shows attributes of a single network.
func (c *Client) ShowNetwork(id string) (*Network, error) {
	var res *Network
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/networks/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/networks/:id(.:format)?
// Updates the given network.
func (c *Client) UpdateNetwork(id string, network *NetworkParam2) error {
	payload := ApiParams{
		"network": network,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/networks/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  NetworkGateway ******/

// A NetworkGateway is an interface that allows traffic to be routed between networks.
type NetworkGateway struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	Type        string              `json:"type,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// POST api/network_gateways(.:format)?
// Create a new NetworkGateway.
func (c *Client) CreateNetworkGateway(networkGateway *NetworkGatewayParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"network_gateway": networkGateway,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/network_gateways", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/network_gateways/:id(.:format)?
// Delete an existing NetworkGateway.
func (c *Client) DestroyNetworkGateway(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/network_gateways/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/network_gateways(.:format)?
// Lists the NetworkGateways available to this account.
// -- Optional parameters:
// 	filter
func (c *Client) IndexNetworkGateways(options ApiParams) ([]NetworkGateway, error) {
	var res []NetworkGateway
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_gateways", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/network_gateways/:id(.:format)?
// Show information about a single NetworkGateway.
func (c *Client) ShowNetworkGateway(id string) (*NetworkGateway, error) {
	var res *NetworkGateway
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_gateways/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/network_gateways/:id(.:format)?
// Update an existing NetworkGateway.
func (c *Client) UpdateNetworkGateway(id string, networkGateway *NetworkGatewayParam2) error {
	payload := ApiParams{
		"network_gateway": networkGateway,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/network_gateways/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Options     map[string]string   `json:"options,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Type        string              `json:"type,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// POST api/network_option_groups(.:format)?
// Create a new NetworkOptionGroup.
func (c *Client) CreateNetworkOptionGroup(networkOptionGroup *NetworkOptionGroupParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"network_option_group": networkOptionGroup,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/network_option_groups", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/network_option_groups/:id(.:format)?
// Delete an existing NetworkOptionGroup.
func (c *Client) DestroyNetworkOptionGroup(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/network_option_groups/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/network_option_groups(.:format)?
// List NetworkOptionGroups available in this account.
// -- Optional parameters:
// 	filter
func (c *Client) IndexNetworkOptionGroups(options ApiParams) ([]NetworkOptionGroup, error) {
	var res []NetworkOptionGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_option_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/network_option_groups/:id(.:format)?
// Show information about a single NetworkOptionGroup.
func (c *Client) ShowNetworkOptionGroup(id string) (*NetworkOptionGroup, error) {
	var res *NetworkOptionGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_option_groups/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/network_option_groups/:id(.:format)?
// Update an existing NetworkOptionGroup.
func (c *Client) UpdateNetworkOptionGroup(id string, networkOptionGroup *NetworkOptionGroupParam2) error {
	payload := ApiParams{
		"network_option_group": networkOptionGroup,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/network_option_groups/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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
	CreatedAt          *time.Time          `json:"created_at,omitempty"`
	Links              []map[string]string `json:"links,omitempty"`
	NetworkOptionGroup string              `json:"network_option_group,omitempty"`
	ResourceUid        string              `json:"resource_uid,omitempty"`
	UpdatedAt          *time.Time          `json:"updated_at,omitempty"`
}

// POST api/network_option_group_attachments(.:format)?
// Create a new NetworkOptionGroupAttachment.
func (c *Client) CreateNetworkOptionGroupAttachment(networkOptionGroupAttachment *NetworkOptionGroupAttachmentParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"network_option_group_attachment": networkOptionGroupAttachment,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/network_option_group_attachments", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/network_option_group_attachments/:id(.:format)?
// Delete an existing NetworkOptionGroupAttachment.
func (c *Client) DestroyNetworkOptionGroupAttachment(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/network_option_group_attachments/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/network_option_group_attachments(.:format)?
// List NetworkOptionGroupAttachments in this account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexNetworkOptionGroupAttachments(options ApiParams) ([]NetworkOptionGroupAttachment, error) {
	var res []NetworkOptionGroupAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_option_group_attachments", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/network_option_group_attachments/:id(.:format)?
// Show information about a single NetworkOptionGroupAttachment.
// -- Optional parameters:
// 	view
func (c *Client) ShowNetworkOptionGroupAttachment(id string, options ApiParams) (*NetworkOptionGroupAttachment, error) {
	var res *NetworkOptionGroupAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_option_group_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/network_option_group_attachments/:id(.:format)?
// Update an existing NetworkOptionGroupAttachment.
func (c *Client) UpdateNetworkOptionGroupAttachment(id string, networkOptionGroupAttachment *NetworkOptionGroupAttachmentParam2) error {
	payload := ApiParams{
		"network_option_group_attachment": networkOptionGroupAttachment,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/network_option_group_attachments/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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
type Oauth2 struct {
}

// POST api/oauth2/
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
//   "access_token": "xyzzy",
//   "expires_in":   3600,
//   "token_type":   "bearer"
// }

// 	grantType: Type of grant.
// -- Optional parameters:
// 	accountId: The client's account ID (only needed for instance agent clients).
// 	clientId: The client ID (only needed for confidential clients).
// 	clientSecret: The client secret (only needed for confidential clients).
// 	refreshToken: The refresh token obtained from OAuth grant.
// 	rightLinkVersion: The RightLink gem version the client conforms to (only needed for instance agent clients).
// 	rsVersion: The RightAgent protocol version the client conforms to (only needed for instance agent clients).
func (c *Client) CreateOauth2(grantType string, options ApiParams) (map[string]interface{}, error) {
	var res map[string]interface{}
	payload := mergeOptionals(ApiParams{

		"grant_type": grantType,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/oauth2/", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Permission ******/

type Permission struct {
	Actions   []string            `json:"actions,omitempty"`
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	RoleTitle string              `json:"role_title,omitempty"`
}

// POST api/permissions(.:format)?
// Create a permission, thereby granting some user a particular role
// with respect to the current account.

// The 'observer' role has a special status; it must be granted before
// a user is eligible for any other permission in a given account.

// When provisioning users, always create the observer permission FIRST;
// creating any other permission before it will result in an error.

// For more information about the roles available and the privileges
// they confer, please refer to the following page of the RightScale
// support portal:
//   http://support.rightscale.com/15-References/Lists/ListofUser_Roles
func (c *Client) CreatePermission(permission *PermissionParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"permission": permission,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/permissions", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/permissions/:id(.:format)?
// Destroy a permission, thereby revoking a user's role with respect
// to the current account.

// The 'observer' role has a special status; it cannot be revoked if
// a user has any other roles, because other roles become useless without
// being able to read data pertaining to the account.

// When deprovisioning user, always destroy the observer permission LAST;
// destroying it while the user has other permissions will result in an error.
func (c *Client) DestroyPermission(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/permissions/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/permissions(.:format)?
// List all permissions for all users of the current acount.
// -- Optional parameters:
// 	filter
func (c *Client) IndexPermissions(options ApiParams) ([]Permission, error) {
	var res []Permission
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/permissions", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/permissions/:id(.:format)?
// Show information about a single permission.
func (c *Client) ShowPermission(id string) (*Permission, error) {
	var res *Permission
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/permissions/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  PlacementGroup ******/

type PlacementGroup struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// POST api/placement_groups(.:format)?
// Creates a PlacementGroup.
func (c *Client) CreatePlacementGroup(placementGroup *PlacementGroupParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"placement_group": placementGroup,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/placement_groups", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/placement_groups/:id(.:format)?
// Destroys a PlacementGroup.
func (c *Client) DestroyPlacementGroup(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/placement_groups/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/placement_groups(.:format)?
// Lists all PlacementGroups in an account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexPlacementGroups(options ApiParams) ([]PlacementGroup, error) {
	var res []PlacementGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/placement_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/placement_groups/:id(.:format)?
// Shows information about a single PlacementGroup.
// -- Optional parameters:
// 	view
func (c *Client) ShowPlacementGroup(id string, options ApiParams) (*PlacementGroup, error) {
	var res *PlacementGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/placement_groups/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Preference ******/

// A Preference is a user and account-specific setting. Preferences are used in many part of the RightScale platform and can be used for custom purposes if desired.
type Preference struct {
	Actions   []string            `json:"actions,omitempty"`
	Contents  string              `json:"contents,omitempty"`
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

// DELETE api/preferences/:id(.:format)?
// Deletes the given preference.
func (c *Client) DestroyPreference(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/preferences/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/preferences(.:format)?
// Lists all preferences.
// -- Optional parameters:
// 	filter
func (c *Client) IndexPreferences(options ApiParams) ([]Preference, error) {
	var res []Preference
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/preferences", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/preferences/:id(.:format)?
// Shows a single preference.
func (c *Client) ShowPreference(id string) (*Preference, error) {
	var res *Preference
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/preferences/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/preferences/:id(.:format)?
// If 'id' is known, updates preference with given contents.
// Otherwise, creates new preference.
// Note: If create, will return '201 Created' and the location of the new preference.
func (c *Client) UpdatePreference(id string, preference *PreferenceParam) error {
	payload := ApiParams{
		"preference": preference,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/preferences/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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
	CreatedAt     *time.Time          `json:"created_at,omitempty"`
	Description   string              `json:"description,omitempty"`
	Links         []map[string]string `json:"links,omitempty"`
	Name          string              `json:"name,omitempty"`
	Publisher     string              `json:"publisher,omitempty"`
	Revision      string              `json:"revision,omitempty"`
	RevisionNotes string              `json:"revision_notes,omitempty"`
	UpdatedAt     *time.Time          `json:"updated_at,omitempty"`
}

// POST api/publications/:id/import(.:format)?
// Imports the given publication and its subordinates to this account.
// Only non-HEAD revisions that are shared with the account can be imported.
func (c *Client) ImportPublication(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/publications/"+id+"/import", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/publications(.:format)?
// Lists the publications available to this account. Only non-HEAD revisions are possible.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexPublications(options ApiParams) ([]Publication, error) {
	var res []Publication
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/publications", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/publications/:id(.:format)?
// Show information about a single publication. Only non-HEAD revisions are possible.
// -- Optional parameters:
// 	view
func (c *Client) ShowPublication(id string, options ApiParams) (*Publication, error) {
	var res *Publication
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/publications/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  PublicationLineage ******/

// A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.
// It is shared among all revisions of a Publication within the marketplace.
// Publication Lineages are different than lineages that exist within an account.
type PublicationLineage struct {
	Actions          []string            `json:"actions,omitempty"`
	CommentsEmailed  bool                `json:"comments_emailed,omitempty"`
	CommentsEnabled  bool                `json:"comments_enabled,omitempty"`
	CreatedAt        *time.Time          `json:"created_at,omitempty"`
	Links            []map[string]string `json:"links,omitempty"`
	LongDescription  string              `json:"long_description,omitempty"`
	Name             string              `json:"name,omitempty"`
	ShortDescription string              `json:"short_description,omitempty"`
	UpdatedAt        *time.Time          `json:"updated_at,omitempty"`
}

// GET api/publication_lineages/:id(.:format)?
// Show information about a single publication lineage. Only non-HEAD revisions are possible.
// -- Optional parameters:
// 	view
func (c *Client) ShowPublicationLineage(id string, options ApiParams) (*PublicationLineage, error) {
	var res *PublicationLineage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/publication_lineages/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  RecurringVolumeAttachment ******/

// A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a Server/ServerArray the next time an instance is launched.
type RecurringVolumeAttachment struct {
	Actions      []string            `json:"actions,omitempty"`
	CreatedAt    *time.Time          `json:"created_at,omitempty"`
	Device       string              `json:"device,omitempty"`
	DeviceId     string              `json:"device_id,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	Name         string              `json:"name,omitempty"`
	RunnableType string              `json:"runnable_type,omitempty"`
	Size         string              `json:"size,omitempty"`
	Status       string              `json:"status,omitempty"`
	StorageType  string              `json:"storage_type,omitempty"`
	UpdatedAt    *time.Time          `json:"updated_at,omitempty"`
}

// POST api/clouds/:cloud_id/recurring_volume_attachments(.:format)?
// Creates a new recurring volume attachment.
func (c *Client) CreateRecurringVolumeAttachment(cloudId string, recurringVolumeAttachment *RecurringVolumeAttachmentParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"recurring_volume_attachment": recurringVolumeAttachment,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/recurring_volume_attachments", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/recurring_volume_attachments/:id(.:format)?
// Deletes a given recurring volume attachment.
func (c *Client) DestroyRecurringVolumeAttachment(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/recurring_volume_attachments/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/recurring_volume_attachments(.:format)?
// Lists all recurring volume attachments.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexRecurringVolumeAttachments(cloudId string, options ApiParams) ([]RecurringVolumeAttachment, error) {
	var res []RecurringVolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/recurring_volume_attachments", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/recurring_volume_attachments/:id(.:format)?
// Displays information about a single recurring volume attachment.
// -- Optional parameters:
// 	view
func (c *Client) ShowRecurringVolumeAttachment(cloudId string, id string, options ApiParams) (*RecurringVolumeAttachment, error) {
	var res *RecurringVolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/recurring_volume_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Repository ******/

// A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.

// RightScale currently supports the following types of repositores: git, svn, and URLs of compressed files (tar, tgz, gzip).
type Repository struct {
	Actions         []string            `json:"actions,omitempty"`
	AssetCounts     int                 `json:"asset_counts,omitempty"`
	AssetPaths      []string            `json:"asset_paths,omitempty"`
	CommitReference string              `json:"commit_reference,omitempty"`
	CreatedAt       *time.Time          `json:"created_at,omitempty"`
	Credentials     map[string]string   `json:"credentials,omitempty"`
	Description     string              `json:"description,omitempty"`
	FetchStatus     string              `json:"fetch_status,omitempty"`
	Id              string              `json:"id,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	ReadOnly        bool                `json:"read_only,omitempty"`
	Source          string              `json:"source,omitempty"`
	SourceType      string              `json:"source_type,omitempty"`
	UpdatedAt       *time.Time          `json:"updated_at,omitempty"`
}

// POST api/repositories/:id/cookbook_import(.:format)?
// Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.
// 	assetHrefs: Hrefs of the assets that should be imported.
// -- Optional parameters:
// 	follow: A flag indicating whether imported cookbooks should be followed.
// 	namespace: The namespace to import into.
// 	repositoryCommitReference: Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.
// 	withDependencies: A flag indicating whether dependencies should automatically be imported.
func (c *Client) CookbookImportRepository(assetHrefs []string, id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{
		"asset_hrefs": assetHrefs,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/repositories/"+id+"/cookbook_import", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/repositories/:id/cookbook_import_preview(.:format)?
// Retrieves a preview of the effects of a Cookbook import.

// NOTE: This action is for RightScale internal use only. The response is
// free-form JSON with no associated mediatype.

// DO NOT USE, THIS ACTION IS SUBJECT TO CHANGE AT ANYTIME.
// 	assetHrefs: Hrefs of the assets that should be imported.
// 	namespace: The namespace to import into.
func (c *Client) CookbookImportPreviewRepository(assetHrefs []string, id string, namespace string) ([]map[string]string, error) {
	var res []map[string]string
	payload := ApiParams{
		"asset_hrefs": assetHrefs,
		"namespace":   namespace,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/repositories/"+id+"/cookbook_import_preview", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/repositories(.:format)?
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

func (c *Client) CreateRepository(repository *RepositoryParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"repository": repository,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/repositories", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/repositories/:id(.:format)?
// Deletes the specified Repositories.
func (c *Client) DestroyRepository(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/repositories/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/repositories(.:format)?
// Lists all Repositories for this Account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexRepositories(options ApiParams) ([]Repository, error) {
	var res []Repository
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/repositories/:id/refetch(.:format)?
// Refetches all RepositoryAssets associated with the Repository.
// Note that a refetch simply updates RightScale's view of the contents of the repository.
// You must perform an import to use the assets in your design objects (or use the auto import parameter).
// -- Optional parameters:
// 	autoImport: Whether cookbooks should automatically be imported after repositories are fetched.
func (c *Client) RefetchRepository(id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/repositories/"+id+"/refetch", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/repositories/resolve(.:format)?
// Show a list of repositories that have imported cookbooks with the given names.

// This operation returns a list of repositories that would later satisfy a call
// to the swap_repository
// action on a ServerTemplate.
// -- Optional parameters:
// 	importedCookbookName: A list of cookbook names that were imported by the repository.
func (c *Client) ResolveRepository(options ApiParams) ([]Repository, error) {
	var res []Repository
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/repositories/resolve", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/repositories/:id(.:format)?
// Shows a specified Repository.
// -- Optional parameters:
// 	view
func (c *Client) ShowRepository(id string, options ApiParams) (*Repository, error) {
	var res *Repository
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/repositories/:id(.:format)?
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

func (c *Client) UpdateRepository(id string, repository *RepositoryParam2) error {
	payload := ApiParams{
		"repository": repository,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/repositories/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// GET api/repositories/:repository_id/repository_assets(.:format)?
// List a repository's current assets.

// Repository assests are the cookbook details that were scraped from a
// given repository.
// -- Optional parameters:
// 	view
func (c *Client) IndexRepositoryAssets(repositoryId string, options ApiParams) ([]RepositoryAsset, error) {
	var res []RepositoryAsset
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories/"+repositoryId+"/repository_assets", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/repositories/:repository_id/repository_assets/:id(.:format)?
// Show information about a single asset.

// A repository assest are the cookbook details that were scraped from a
// repository.
// -- Optional parameters:
// 	view
func (c *Client) ShowRepositoryAsset(id string, repositoryId string, options ApiParams) (*RepositoryAsset, error) {
	var res *RepositoryAsset
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories/"+repositoryId+"/repository_assets/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  RightScript ******/

// A RightScript is an executable piece of code that can be run on a server
// during the boot, operational, or decommission phases.

// All revisions of
// a RightScript belong to a RightScript lineage that is exposed by the
// "lineage" attribute (NOTE: This attribute is merely a string to locate
// all revisions of a RightScript and NOT a working URL).
type RightScript struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Id          string              `json:"id,omitempty"`
	Lineage     string              `json:"lineage,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Revision    string              `json:"revision,omitempty"`
	Source      string              `json:"source,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// POST api/right_scripts/:id/commit(.:format)?
// Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.
func (c *Client) CommitRightScript(id string, rightScript *RightScriptParam) error {
	payload := ApiParams{
		"right_script": rightScript,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/right_scripts/"+id+"/commit", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/right_scripts(.:format)?
// Lists RightScripts.
// -- Optional parameters:
// 	filter
// 	latestOnly: Whether or not to return only the latest version for each lineage.
// 	view
func (c *Client) IndexRightScripts(options ApiParams) ([]RightScript, error) {
	var res []RightScript
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/right_scripts", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/right_scripts/:id(.:format)?
// Displays information about a single RightScript.
func (c *Client) ShowRightScript(id string) (*RightScript, error) {
	var res *RightScript
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/right_scripts/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/right_scripts/:id/source(.:format)?
// Returns the script source for a RightScript
func (c *Client) ShowSourceRightScript(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/right_scripts/"+id+"/source", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// PUT api/right_scripts/:id(.:format)?
// Updates RightScript name/description
func (c *Client) UpdateRightScript(id string, rightScript *RightScriptParam2) error {
	payload := ApiParams{
		"right_script": rightScript,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/right_scripts/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// PUT api/right_scripts/:id/source(.:format)?
// Updates the source of the given RightScript
func (c *Client) UpdateSourceRightScript(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/right_scripts/"+id+"/source", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  Route ******/

// A Route defines how networking traffic should be routed from one
// destination to another. See nexthoptype for available endpoint targets.
type Route struct {
	CreatedAt            *time.Time          `json:"created_at,omitempty"`
	Description          string              `json:"description,omitempty"`
	DestinationCidrBlock string              `json:"destination_cidr_block,omitempty"`
	Links                []map[string]string `json:"links,omitempty"`
	NextHopIp            string              `json:"next_hop_ip,omitempty"`
	NextHopType          string              `json:"next_hop_type,omitempty"`
	ResourceUid          string              `json:"resource_uid,omitempty"`
	State                string              `json:"state,omitempty"`
	UpdatedAt            *time.Time          `json:"updated_at,omitempty"`
}

// POST api/routes(.:format)?
// Create a new Route.
func (c *Client) CreateRoute(route *RouteParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"route": route,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/routes", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/routes/:id(.:format)?
// Delete an existing Route.
func (c *Client) DestroyRoute(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/routes/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/routes(.:format)?
// List Routes available in this account.
// -- Optional parameters:
// 	filter
func (c *Client) IndexRoutes(options ApiParams) ([]Route, error) {
	var res []Route
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/routes", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/routes/:id(.:format)?
// Show information about a single Route.
func (c *Client) ShowRoute(id string) (*Route, error) {
	var res *Route
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/routes/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/routes/:id(.:format)?
// Update an existing Route.
func (c *Client) UpdateRoute(id string, route *RouteParam2) error {
	payload := ApiParams{
		"route": route,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/routes/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  RouteTable ******/

// Grouped listing of Routes
type RouteTable struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Routes      []Route             `json:"routes,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// POST api/route_tables(.:format)?
// Create a new RouteTable.
func (c *Client) CreateRouteTable(routeTable *RouteTableParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"route_table": routeTable,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/route_tables", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/route_tables/:id(.:format)?
// Delete an existing RouteTable.
func (c *Client) DestroyRouteTable(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/route_tables/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/route_tables(.:format)?
// List RouteTables available in this account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexRouteTables(options ApiParams) ([]RouteTable, error) {
	var res []RouteTable
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/route_tables", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/route_tables/:id(.:format)?
// Show information about a single RouteTable.
// -- Optional parameters:
// 	view
func (c *Client) ShowRouteTable(id string, options ApiParams) (*RouteTable, error) {
	var res *RouteTable
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/route_tables/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/route_tables/:id(.:format)?
// Update an existing RouteTable.
func (c *Client) UpdateRouteTable(id string, routeTable *RouteTableParam2) error {
	payload := ApiParams{
		"route_table": routeTable,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/route_tables/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// POST api/server_templates/:server_template_id/runnable_bindings(.:format)?
// Bind an executable to the given ServerTemplate.

// An executable may be either a RightScript or Chef Recipe.

// The resource must be editable.
func (c *Client) CreateRunnableBinding(runnableBinding *RunnableBindingParam, serverTemplateId string) (Href, error) {
	var res Href
	payload := ApiParams{
		"runnable_binding": runnableBinding,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+serverTemplateId+"/runnable_bindings", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/server_templates/:server_template_id/runnable_bindings/:id(.:format)?
// Unbind an executable from the given resource.

// The resource must be editable.
func (c *Client) DestroyRunnableBinding(id string, serverTemplateId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/server_templates/"+serverTemplateId+"/runnable_bindings/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/server_templates/:server_template_id/runnable_bindings(.:format)?
// Lists the executables bound to the ServerTemplate.

// An excutable may be either a RightScript or Chef Recipe.
// -- Optional parameters:
// 	view
func (c *Client) IndexRunnableBindings(serverTemplateId string, options ApiParams) ([]RunnableBinding, error) {
	var res []RunnableBinding
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+serverTemplateId+"/runnable_bindings", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/server_templates/:server_template_id/runnable_bindings/multi_update(.:format)?
// Update attributes for multiple bound executables.

// The resource must be editable.
func (c *Client) MultiUpdateRunnableBindings(runnableBindings []*RunnableBindings, serverTemplateId string) error {
	payload := ApiParams{
		"runnable_bindings": runnableBindings,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/server_templates/"+serverTemplateId+"/runnable_bindings/multi_update", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/server_templates/:server_template_id/runnable_bindings/:id(.:format)?
// Show information about a single executable binding.

// An excutable may be either a RightScript or Chef Recipe.
// -- Optional parameters:
// 	view
func (c *Client) ShowRunnableBinding(id string, serverTemplateId string, options ApiParams) (*RunnableBinding, error) {
	var res *RunnableBinding
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+serverTemplateId+"/runnable_bindings/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/clouds/:cloud_id/security_groups(.:format)?
// Create a security group.
func (c *Client) CreateSecurityGroup(cloudId string, securityGroup *SecurityGroupParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"security_group": securityGroup,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/security_groups", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/security_groups/:id(.:format)?
// Delete security group(s)
func (c *Client) DestroySecurityGroup(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/security_groups/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/security_groups(.:format)?
// Lists Security Groups.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexSecurityGroups(cloudId string, options ApiParams) ([]SecurityGroup, error) {
	var res []SecurityGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/security_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/security_groups/:id(.:format)?
// Displays information about a single Security Group.
// -- Optional parameters:
// 	view
func (c *Client) ShowSecurityGroup(cloudId string, id string, options ApiParams) (*SecurityGroup, error) {
	var res *SecurityGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/security_groups/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/security_group_rules(.:format)?
// Create a security group rule for a security group.
// The following flavors are supported:

// group-based TCP/UDP
// group-based ICMP
// CIDR-based TCP/UDP
// CIDR-based ICMP

func (c *Client) CreateSecurityGroupRule(securityGroupRule *SecurityGroupRuleParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"security_group_rule": securityGroupRule,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/security_group_rules", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/security_group_rules/:id(.:format)?
// Delete security group rule(s)
func (c *Client) DestroySecurityGroupRule(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/security_group_rules/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/security_group_rules(.:format)?
// Lists SecurityGroupRules.
// -- Optional parameters:
// 	view
func (c *Client) IndexSecurityGroupRules(options ApiParams) ([]SecurityGroupRule, error) {
	var res []SecurityGroupRule
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/security_group_rules", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/security_group_rules/:id(.:format)?
// Displays information about a single SecurityGroupRule.
// -- Optional parameters:
// 	view
func (c *Client) ShowSecurityGroupRule(id string, options ApiParams) (*SecurityGroupRule, error) {
	var res *SecurityGroupRule
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/security_group_rules/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/security_group_rules/:id(.:format)?

func (c *Client) UpdateSecurityGroupRule(id string, securityGroupRule *SecurityGroupRuleParam2) error {
	payload := ApiParams{
		"security_group_rule": securityGroupRule,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/security_group_rules/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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
	CreatedAt       *time.Time          `json:"created_at,omitempty"`
	CurrentInstance Instance            `json:"current_instance,omitempty"`
	Description     string              `json:"description,omitempty"`
	Links           []map[string]string `json:"links,omitempty"`
	Name            string              `json:"name,omitempty"`
	NextInstance    *Instance           `json:"next_instance,omitempty"`
	Optimized       bool                `json:"optimized,omitempty"`
	State           string              `json:"state,omitempty"`
	UpdatedAt       *time.Time          `json:"updated_at,omitempty"`
}

// POST api/servers/:id/clone(.:format)?
// Clones a given server.
func (c *Client) CloneServer(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/servers/"+id+"/clone", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/servers(.:format)?
// Creates a new server, and configures its corresponding "next" instance with the received parameters.
func (c *Client) CreateServer(server *ServerParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"server": server,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/servers", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/servers/:id(.:format)?
// Deletes a given server.
func (c *Client) DestroyServer(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/servers/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/servers(.:format)?
// Lists servers.

// By using the available filters, it is possible to retrieve servers that have common characteristics.
// For example, one can list:

// servers that have names that contain "app_server"
// all servers of a given deployment

// For more filters, please see the 'index' action on 'Instances' resource as most of the attributes belong to
// a 'current_instance' than to a server.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexServers(options ApiParams) ([]Server, error) {
	var res []Server
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/servers/:id/launch
// Launches the "next" instance of this server. This function is equivalent to invoking the launch action on the
// URL of this servers next_instance. See Instances#launch for details.
func (c *Client) LaunchServer(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/servers/"+id+"/launch", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/servers/:id(.:format)?
// Shows the information of a single server.
// -- Optional parameters:
// 	view
func (c *Client) ShowServer(id string, options ApiParams) (*Server, error) {
	var res *Server
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/servers/:id/teminate
// Terminates the current instance of this server. This function is equivalent to invoking the terminate action on the
// URL of this servers current_instance. See Instances#terminate for details.
func (c *Client) TerminateServer(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/servers/"+id+"/teminate", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// PUT api/servers/:id(.:format)?
// Updates attributes of a single server.
func (c *Client) UpdateServer(id string, server *ServerParam2) error {
	payload := ApiParams{
		"server": server,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/servers/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/servers/wrap_instance(.:format)?
// Wrap an existing instance and set current instance for new server
func (c *Client) WrapInstanceServer(server *ServerParam3) error {
	payload := ApiParams{
		"server": server,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/servers/wrap_instance", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// POST api/server_arrays/:id/clone(.:format)?
// Clones a given server array.
func (c *Client) CloneServerArray(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_arrays/"+id+"/clone", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_arrays(.:format)?
// Creates a new server array, and configures its corresponding "next" instance with the received parameters.
func (c *Client) CreateServerArray(serverArray *ServerArrayParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"server_array": serverArray,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_arrays", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// GET api/server_arrays/:id/current_instances
// List the running instances belonging to the server array. See Instances#index for details.
// This action is slightly different from invoking the index action on the Instances resource with the filter "parent_href == /api/server_arrays/XX" because the
// latter will include 'next_instance' as well.
func (c *Client) CurrentInstancesServerArray(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_arrays/"+id+"/current_instances", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// DELETE api/server_arrays/:id(.:format)?
// Deletes a given server array.
func (c *Client) DestroyServerArray(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/server_arrays/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/server_arrays(.:format)?
// Lists server arrays.

// By using the available filters, it is possible to retrieve server arrays that have common characteristics.
// For example, one can list:

// arrays that have names that contain "my_server_array"
// all arrays of a given deployment

// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexServerArrays(options ApiParams) ([]ServerArray, error) {
	var res []ServerArray
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_arrays", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/server_arrays/:id/launch
// Launches a new instance in the server array with the configuration defined in the 'next_instance'. This function is equivalent to invoking the launch action on the
// URL of this server_array's next_instance. See Instances#launch for details.
func (c *Client) LaunchServerArray(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_arrays/"+id+"/launch", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_arrays/:id/multi_run_executable
// Run an executable on all instances of this array. This function is equivalent to invoking the "multi_run_executable" action on the instances resource
// (Instances#multi_run_executable with the filter "parent_href == /api/server_arrays/XX"). To run an executable on a subset of the instances of the array, provide additional filters. To run an executable
// a single instance, invoke the action "run_executable" directly on the instance (see Instances#run_executable)
func (c *Client) MultiRunExecutableServerArrays(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_arrays/"+id+"/multi_run_executable", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_arrays/:id/multi_terminate
// Terminate all instances of this array. This function is equivalent to invoking the "multi_terminate" action on the instances resource ( Instances#multi_terminate with
// the filter "parent_href == /api/server_arrays/XX"). To terminate a subset of the instances of the array, provide additional filters. To terminate a single instance,
// invoke the action "terminate" directly on the instance (see Instances#terminate)
func (c *Client) MultiTerminateServerArrays(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_arrays/"+id+"/multi_terminate", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/server_arrays/:id(.:format)?
// Shows the information of a single server array.
// -- Optional parameters:
// 	view
func (c *Client) ShowServerArray(id string, options ApiParams) (*ServerArray, error) {
	var res *ServerArray
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_arrays/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/server_arrays/:id(.:format)?
// Updates attributes of a single server array.
func (c *Client) UpdateServerArray(id string, serverArray *ServerArrayParam2) error {
	payload := ApiParams{
		"server_array": serverArray,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/server_arrays/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// POST api/server_templates/:id/clone(.:format)?
// Clones a given ServerTemplate.
func (c *Client) CloneServerTemplate(id string, serverTemplate *ServerTemplateParam) error {
	payload := ApiParams{
		"server_template": serverTemplate,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+id+"/clone", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_templates/:id/commit(.:format)?
// Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.
// 	commitHeadDependencies: Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.
// 	commitMessage: The message associated with the commit.
// 	freezeRepositories: Freeze the repositories.
func (c *Client) CommitServerTemplate(commitHeadDependencies string, commitMessage string, freezeRepositories string, id string) error {
	payload := ApiParams{
		"commit_head_dependencies": commitHeadDependencies,
		"commit_message":           commitMessage,
		"freeze_repositories":      freezeRepositories,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+id+"/commit", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_templates(.:format)?
// Creates a new ServerTemplate with the given parameters.
func (c *Client) CreateServerTemplate(serverTemplate *ServerTemplateParam2) (Href, error) {
	var res Href
	payload := ApiParams{
		"server_template": serverTemplate,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/server_templates/:id(.:format)?
// Deletes a given ServerTemplate.
func (c *Client) DestroyServerTemplate(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/server_templates/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_templates/:id/detect_changes_in_head(.:format)?
// Identifies RightScripts attached to the resource that differ from their HEAD.

// If the attached revision of the RightScript is the HEAD, then this will indicate
// a difference between it and the latest committed revision in the same lineage.
func (c *Client) DetectChangesInHeadServerTemplate(id string) ([]map[string]string, error) {
	var res []map[string]string
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+id+"/detect_changes_in_head", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/server_templates(.:format)?
// Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.

// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexServerTemplates(options ApiParams) ([]ServerTemplate, error) {
	var res []ServerTemplate
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/server_templates/:id/publish(.:format)?
// Publishes a given ServerTemplate and its subordinates.
// Only non-HEAD revisions that are owned by the account can be published.
// 	accountGroupHrefs: List of hrefs of account groups to publish to.
// -- Optional parameters:
// 	allowComments: Allow users to leave comments on this ServerTemplate.
// 	categories: List of Categories.
// 	emailComments: Email me when a user comments on this ServerTemplate.
func (c *Client) PublishServerTemplate(accountGroupHrefs []string, descriptions *Descriptions, id string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{
		"account_group_hrefs": accountGroupHrefs,

		"descriptions": descriptions,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+id+"/publish", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/server_templates/:id/resolve(.:format)?
// Enumerates all attached cookbooks, missing dependencies and bound executables.

// Version constraints on missing dependencies and the state of the Chef Recipes;
// whether or not the cookbook or recipe itself could be found among the
// attachments, will also be reported.
func (c *Client) ResolveServerTemplate(id string) ([]map[string]string, error) {
	var res []map[string]string
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+id+"/resolve", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/server_templates/:id(.:format)?
// Show information about a single ServerTemplate. HEAD revisions have a revision of 0.

// The 'inputs_2_0' view is for retrieving inputs in 2.0 serialization (for more
// details please see Inputs#index.)
// -- Optional parameters:
// 	view
func (c *Client) ShowServerTemplate(id string, options ApiParams) (*ServerTemplate, error) {
	var res *ServerTemplate
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/server_templates/:id/swap_repository(.:format)?
// In-place replacement of attached cookbooks from a given repository.

// For each attached cookbook coming from the source repository, replace it by
// attaching a cookbook of identical name coming from the target repository.

// In order for the operation to be successful, all attachments that came from the
// source repository must exist in the target repository.

// If multiple cookbooks of a given name exist in the target repository, preference
// is given by the following order (top most being the highest preference):

//   Name & Version Match / Primary Namespace
//   Name & Version Match / Alternate Namespace
//   Name Match / Primary Namespace
//   Name Match / Alternate Namespace

// If multiple cookbooks still have the same preference for the replacement, the operation is
// indeterministic.
// 	sourceRepositoryHref: The repository whose cookbook attachments are to be replaced.
// 	targetRepositoryHref: The repository whose cookbook attachments are to be utilized.
func (c *Client) SwapRepositoryServerTemplate(id string, sourceRepositoryHref string, targetRepositoryHref string) error {
	payload := ApiParams{
		"source_repository_href": sourceRepositoryHref,
		"target_repository_href": targetRepositoryHref,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_templates/"+id+"/swap_repository", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// PUT api/server_templates/:id(.:format)?
// Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).
// Currently, the attributes you can update are only the 'direct' attributes of a server template. To
// manage multi cloud images of a ServerTemplate, please see the resource 'ServerTemplateMultiCloudImages'.
func (c *Client) UpdateServerTemplate(id string, serverTemplate *ServerTemplateParam) error {
	payload := ApiParams{
		"server_template": serverTemplate,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/server_templates/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  ServerTemplateMultiCloudImage ******/

// This resource represents links between ServerTemplates and MultiCloud Images and enables you to effectively
// add/delete MultiCloudImages to ServerTemplates and make them the default one.
type ServerTemplateMultiCloudImage struct {
	Actions   []string            `json:"actions,omitempty"`
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	IsDefault bool                `json:"is_default,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

// POST api/server_template_multi_cloud_images(.:format)?
// Creates a new ServerTemplateMultiCloudImage with the given parameters.
func (c *Client) CreateServerTemplateMultiCloudImage(serverTemplateMultiCloudImage *ServerTemplateMultiCloudImageParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"server_template_multi_cloud_image": serverTemplateMultiCloudImage,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_template_multi_cloud_images", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/server_template_multi_cloud_images/:id(.:format)?
// Deletes a given ServerTemplateMultiCloudImage.
func (c *Client) DestroyServerTemplateMultiCloudImage(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/server_template_multi_cloud_images/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/server_template_multi_cloud_images(.:format)?
// Lists the ServerTemplateMultiCloudImages available to this account.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexServerTemplateMultiCloudImages(options ApiParams) ([]ServerTemplateMultiCloudImage, error) {
	var res []ServerTemplateMultiCloudImage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_template_multi_cloud_images", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/server_template_multi_cloud_images/:id/make_default(.:format)?
// Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.
func (c *Client) MakeDefaultServerTemplateMultiCloudImage(id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/server_template_multi_cloud_images/"+id+"/make_default", body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/server_template_multi_cloud_images/:id(.:format)?
// Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.
// -- Optional parameters:
// 	view
func (c *Client) ShowServerTemplateMultiCloudImage(id string, options ApiParams) (*ServerTemplateMultiCloudImage, error) {
	var res *ServerTemplateMultiCloudImage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_template_multi_cloud_images/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/sessions/accounts(.:format)?
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
// 	email: The email to login with if not using existing session.
// 	password: The corresponding password.
// 	view: Extended view shows account permissions and products
func (c *Client) AccountsSession(options ApiParams) ([]Account, error) {
	var res []Account
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/sessions/accounts", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/sessions(.:format)?
// Creates API session scoped to a given account. (API login)

// This call requires a form of authentication (user and password), as well as the account for which the session needs to be created.
// Upon successfully authenticating the credentials, the system will return a 204 code and set of two cookies that will serve as the credentials for the session. Both of these cookies
// must be passed in any of the subsequent requests for this session.
// If an 302 redirect code is returned, the client is responsible of re-issuing the POST request against the content of the received Location header, passing the exact same parameters again.

// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -c mycookies -X POST -d email='email@me.com' -d password='mypassword' -d account_href=/api/accounts/11 https://my.rightscale.com/api/sessions

// 	accountHref: The account href for which the session needs to be created.
// 	email: The email to login with.
// 	password: The corresponding password.
func (c *Client) CreateSession(accountHref string, email string, password string) (Href, error) {
	var res Href
	payload := ApiParams{
		"account_href": accountHref,
		"email":        email,
		"password":     password,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/sessions", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// POST api/sessions/instance(.:format)?
// Creates API session scoped to a given account and instance.

// This call requires a form of authentication (token), as well as the account for which the session needs to be created.
// Upon successfully authenticating the credentials, the system will return a 204 code and set of two cookies that will serve as the credentials for the session. Both of these cookies
// must be passed in any of the subsequent requests for this session.
// If an 302 redirect code is returned, the client is responsible of re-issuing the POST request against the content of the received Location header, passing the exact same parameters again.

// Users can find their account ID and instance\_token from their instance's user_data:
// account ID regex: /RS_API_TOKEN=(\d+):/
// instance_token regex: /RS_API_TOKEN=(?:\d+):(\w+)&/

// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -c mycookies -X POST -d instance_token='randomtoken' -d account_href=/api/accounts/11 https://my.rightscale.com/api/sessions/instance

// 	accountHref: The account href for which the session needs to be created.
// 	instanceToken: The instance token to login with.
func (c *Client) CreateInstanceSessionSession(accountHref string, instanceToken string) error {
	payload := ApiParams{
		"account_href":   accountHref,
		"instance_token": instanceToken,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/sessions/instance", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/sessions(.:format)?
// Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what
// features are available within its privileges.

// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions

func (c *Client) IndexSessions() ([]Session, error) {
	var res []Session
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/sessions", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/sessions/instance(.:format)?
// Shows the full attributes of the instance (that has the token used to log-in).
// This call can be used by an instance to get it's own details.

// Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -b mycookies -X GET https://my.rightscale.com/api/sessions/instance

func (c *Client) IndexInstanceSessionSession() (Instance, error) {
	var res Instance
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/sessions/instance", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  SshKey ******/

// Ssh Keys represent a created SSH Key that exists in the cloud.

// An ssh key might also contain the private part of the key, and can be used to login to instances launched with it.
type SshKey struct {
	Actions     []string            `json:"actions,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Material    string              `json:"material,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

// POST api/clouds/:cloud_id/ssh_keys(.:format)?
// Creates a new ssh key.
func (c *Client) CreateSshKey(cloudId string, sshKey *SshKeyParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"ssh_key": sshKey,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/ssh_keys", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/ssh_keys/:id(.:format)?
// Deletes a given ssh key.
func (c *Client) DestroySshKey(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/ssh_keys/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/ssh_keys(.:format)?
// Lists ssh keys.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexSshKeys(cloudId string, options ApiParams) ([]SshKey, error) {
	var res []SshKey
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ssh_keys", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/ssh_keys/:id(.:format)?
// Displays information about a single ssh key.
// -- Optional parameters:
// 	view
func (c *Client) ShowSshKey(cloudId string, id string, options ApiParams) (*SshKey, error) {
	var res *SshKey
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ssh_keys/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/clouds/:cloud_id/instances/:instance_id/subnets(.:format)?
// Creates a new subnet.
func (c *Client) CreateSubnet(cloudId string, instanceId string, subnet *SubnetParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"subnet": subnet,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/subnets", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/instances/:instance_id/subnets/:id(.:format)?
// Deletes the given subnet(s).
func (c *Client) DestroySubnet(cloudId string, id string, instanceId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/subnets/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/instances/:instance_id/subnets(.:format)?
// Lists subnets of a given cloud.
// -- Optional parameters:
// 	filter
func (c *Client) IndexSubnets(cloudId string, instanceId string, options ApiParams) ([]Subnet, error) {
	var res []Subnet
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/subnets", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/instances/:instance_id/subnets/:id(.:format)?
// Shows attributes of a single subnet.
func (c *Client) ShowSubnet(cloudId string, id string, instanceId string) (*Subnet, error) {
	var res *Subnet
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/subnets/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/clouds/:cloud_id/instances/:instance_id/subnets/:id(.:format)?
// Updates name and description for the current subnet.
func (c *Client) UpdateSubnet(cloudId string, id string, instanceId string, subnet *SubnetParam2) error {
	payload := ApiParams{
		"subnet": subnet,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/subnets/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  Tag ******/

// A tag or machine tag is a useful way of attaching useful metadata to an object/resource.
// Tags are commonly used as an extra label or identifier.
// For example, you might want to add a tag to an EBS Snapshot or AMI so that you can find it more quickly.
type Tag struct {
}

// POST api/tags/by_resource(.:format)?
// Get tags for a list of resource hrefs.
// The hrefs can belong to various resource types and the tags for a non-existent href will be empty.
// 	resourceHrefs: Hrefs of the resources for which tags are to be returned.
func (c *Client) ByResourceTag(resourceHrefs []string) ([]map[string]string, error) {
	var res []map[string]string
	payload := ApiParams{
		"resource_hrefs": resourceHrefs,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/tags/by_resource", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/tags/by_tag(.:format)?
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
// 	resourceType: Search among a single resource type.
// 	tags: The tags which must be present on the resource.
// -- Optional parameters:
// 	includeTagsWithPrefix: If included, all tags matching this prefix will be returned. If not included, no tags will be returned.
// 	matchAll: If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.
// 	withDeleted: If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.
func (c *Client) ByTagTag(resourceType string, tags []string, options ApiParams) ([]map[string]string, error) {
	var res []map[string]string
	payload := mergeOptionals(ApiParams{

		"resource_type": resourceType,
		"tags":          tags,
	}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/tags/by_tag", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// POST api/tags/multi_add(.:format)?
// Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.
// The hrefs can belong to various resource types. If a resource for a href could not be found, an
// error is returned and no tags are added for any resource.

// No error will be raised if the resource already has the tag(s) you are trying to add.

// Note: At this point, tags on 'next_instance' are not supported and one has to add tags to the 'server'.
// 	resourceHrefs: Hrefs of the resources for which the tags are to be added.
// 	tags: Tags to be added.
func (c *Client) MultiAddTags(resourceHrefs []string, tags []string) error {
	payload := ApiParams{
		"resource_hrefs": resourceHrefs,
		"tags":           tags,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/tags/multi_add", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// POST api/tags/multi_delete(.:format)?
// Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.
// The hrefs can belong to various resource types. If a resource for a href could not be found, an
// error is returned and no tags are deleted for any resource.

// Note that no error will be raised if the resource does not have the tag(s) you are trying to delete.
// 	resourceHrefs: Hrefs of the resources for which tags are to be deleted.
// 	tags: Tags to be deleted.
func (c *Client) MultiDeleteTags(resourceHrefs []string, tags []string) error {
	payload := ApiParams{
		"resource_hrefs": resourceHrefs,
		"tags":           tags,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/tags/multi_delete", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
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

// GET api/clouds/:cloud_id/instances/:instance_id/live/tasks/:id(.:format)?
// Displays information of a given task within the context of an instance.
// -- Optional parameters:
// 	view
func (c *Client) ShowTask(cloudId string, id string, instanceId string, options ApiParams) (*Task, error) {
	var res *Task
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/live/tasks/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  User ******/

// A User represents an individual RightScale user. Users must be given access to RightScale accounts in order to
// access RightScale resources.
type User struct {
	Actions      []string            `json:"actions,omitempty"`
	Company      string              `json:"company,omitempty"`
	CreatedAt    *time.Time          `json:"created_at,omitempty"`
	Email        string              `json:"email,omitempty"`
	FirstName    string              `json:"first_name,omitempty"`
	LastName     string              `json:"last_name,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	Phone        string              `json:"phone,omitempty"`
	PrincipalUid string              `json:"principal_uid,omitempty"`
	TimezoneName string              `json:"timezone_name,omitempty"`
	UpdatedAt    *time.Time          `json:"updated_at,omitempty"`
}

// POST api/users(.:format)?
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
func (c *Client) CreateUser(user *UserParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"user": user,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/users", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// GET api/users(.:format)?
// List the users available to the account the user is logged in to. Therefore, to list the users of
// a child account, the user has to login to the child account first.
// -- Optional parameters:
// 	filter
func (c *Client) IndexUsers(options ApiParams) ([]User, error) {
	var res []User
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/users", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/users/:id(.:format)?
// Show information about a single user.
func (c *Client) ShowUser(id string) (*User, error) {
	var res *User
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/users/"+id, body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// PUT api/users/:id(.:format)?
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
func (c *Client) UpdateUser(id string, user *UserParam2) error {
	payload := ApiParams{
		"user": user,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("PUT", c.endpoint+"/api/users/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

/******  UserData ******/

type UserData struct {
}

// GET api/user_data/
// <no description>
func (c *Client) ShowUserData() (*map[string]string, error) {
	var res *map[string]string
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/user_data/", body)
	if err != nil {
		return res, err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  Volume ******/

// A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).
type Volume struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Iops        string              `json:"iops,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Size        string              `json:"size,omitempty"`
	Status      string              `json:"status,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
	VolumeType  string              `json:"volume_type,omitempty"`
}

// POST api/clouds/:cloud_id/volumes(.:format)?
// Creates a new volume.
func (c *Client) CreateVolume(cloudId string, volume *VolumeParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"volume": volume,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/volumes", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/volumes/:id(.:format)?
// Deletes a given volume.
func (c *Client) DestroyVolume(cloudId string, id string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/volumes(.:format)?
// Lists volumes.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexVolumes(cloudId string, options ApiParams) ([]Volume, error) {
	var res []Volume
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/volumes/:id(.:format)?
// Displays information about a single volume.
// -- Optional parameters:
// 	view
func (c *Client) ShowVolume(cloudId string, id string, options ApiParams) (*Volume, error) {
	var res *Volume
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  VolumeAttachment ******/

// A VolumeAttachment represents a relationship between a volume and an instance.
type VolumeAttachment struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Device      string              `json:"device,omitempty"`
	DeviceId    string              `json:"device_id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// POST api/clouds/:cloud_id/instances/:instance_id/volume_attachments(.:format)?
// Creates a new volume attachment.
func (c *Client) CreateVolumeAttachment(cloudId string, instanceId string, volumeAttachment *VolumeAttachmentParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"volume_attachment": volumeAttachment,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/volume_attachments", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id(.:format)?
// Deletes a given volume attachment.
// -- Optional parameters:
// 	force: Specifies whether to force the detachment of a volume.
func (c *Client) DestroyVolumeAttachment(cloudId string, id string, instanceId string, options ApiParams) error {
	payload := mergeOptionals(ApiParams{}, options)
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/volume_attachments/"+id, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/instances/:instance_id/volume_attachments(.:format)?
// Lists all volume attachments.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexVolumeAttachments(cloudId string, instanceId string, options ApiParams) ([]VolumeAttachment, error) {
	var res []VolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/volume_attachments", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id(.:format)?
// Displays information about a single volume attachment.
// -- Optional parameters:
// 	view
func (c *Client) ShowVolumeAttachment(cloudId string, id string, instanceId string, options ApiParams) (*VolumeAttachment, error) {
	var res *VolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/volume_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  VolumeSnapshot ******/

// A VolumeSnapshot represents a Cloud storage volume at a particular point in time. One can create a snapshot regardless of whether or not a volume is attached to an Instance. When a snapshot is created,
// various meta data is retained such as a Created At timestamp, a unique Resource UID (e.g. vol-52EF05A9), the Volume Owner and Visibility (e.g. private or public).
// Snapshots consist of a series of data blocks that are incrementally saved.
type VolumeSnapshot struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Progress    string              `json:"progress,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Size        string              `json:"size,omitempty"`
	State       string              `json:"state,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// POST api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots(.:format)?
// Creates a new volume_snapshot.
func (c *Client) CreateVolumeSnapshot(cloudId string, volumeId string, volumeSnapshot *VolumeSnapshotParam) (Href, error) {
	var res Href
	payload := ApiParams{
		"volume_snapshot": volumeSnapshot,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+volumeId+"/volume_snapshots", body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return res, err
	}
	loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}
}

// DELETE api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id(.:format)?
// Deletes a given volume_snapshot.
func (c *Client) DestroyVolumeSnapshot(cloudId string, id string, volumeId string) error {
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("DELETE", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+volumeId+"/volume_snapshots/"+id, body)
	if err != nil {
		return err
	}
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return err
	}
	return nil
}

// GET api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots(.:format)?
// Lists all volume_snapshots.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexVolumeSnapshots(cloudId string, volumeId string, options ApiParams) ([]VolumeSnapshot, error) {
	var res []VolumeSnapshot
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+volumeId+"/volume_snapshots", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id(.:format)?
// Displays information about a single volume_snapshot.
// -- Optional parameters:
// 	view
func (c *Client) ShowVolumeSnapshot(cloudId string, id string, volumeId string, options ApiParams) (*VolumeSnapshot, error) {
	var res *VolumeSnapshot
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+volumeId+"/volume_snapshots/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

/******  VolumeType ******/

// A VolumeType describes the type of volume, particularly the size.
type VolumeType struct {
	Actions     []string            `json:"actions,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
	Size        string              `json:"size,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
}

// GET api/clouds/:cloud_id/volume_types(.:format)?
// Lists Volume Types.
// -- Optional parameters:
// 	filter
// 	view
func (c *Client) IndexVolumeTypes(cloudId string, options ApiParams) ([]VolumeType, error) {
	var res []VolumeType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volume_types", body)
	if err != nil {
		return res, err
	}
	for _, v := range options["filter"].([]string) {
		v = options["filter"].(string)
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

// GET api/clouds/:cloud_id/volume_types/:id(.:format)?
// Displays information about a single Volume Type.
// -- Optional parameters:
// 	view
func (c *Client) ShowVolumeType(cloudId string, id string, options ApiParams) (*VolumeType, error) {
	var res *VolumeType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volume_types/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", options["view"].(string))
	ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
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

type AlertSpecificParams2 struct {
	DecisionThreshold  string `json:"decision_threshold,omitempty"`
	VotersTagPredicate string `json:"voters_tag_predicate,omitempty"`
}

type AssetPaths struct {
	Cookbooks []string `json:"cookbooks,omitempty"`
}

type AssetPaths2 struct {
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

type Bounds2 struct {
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
	CloudHref string            `json:"cloud_href,omitempty"`
	Creds     map[string]string `json:"creds,omitempty"`
	Token     string            `json:"token,omitempty"`
}

type CloudSpecificAttributes struct {
	AutomaticInstanceStoreMapping string `json:"automatic_instance_store_mapping,omitempty"`
	IamInstanceProfile            string `json:"iam_instance_profile,omitempty"`
	RootVolumePerformance         string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid             string `json:"root_volume_type_uid,omitempty"`
}

type CloudSpecificAttributes2 struct {
	AutomaticInstanceStoreMapping string `json:"automatic_instance_store_mapping,omitempty"`
	IamInstanceProfile            string `json:"iam_instance_profile,omitempty"`
	RootVolumePerformance         string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid             string `json:"root_volume_type_uid,omitempty"`
}

type CloudSpecificAttributes3 struct {
	AutomaticInstanceStoreMapping string `json:"automatic_instance_store_mapping,omitempty"`
	EbsOptimized                  string `json:"ebs_optimized,omitempty"`
	IamInstanceProfile            string `json:"iam_instance_profile,omitempty"`
	RootVolumePerformance         string `json:"root_volume_performance,omitempty"`
	RootVolumeSize                string `json:"root_volume_size,omitempty"`
	RootVolumeTypeUid             string `json:"root_volume_type_uid,omitempty"`
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

type CredentialParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
}

type Credentials struct {
	Password string `json:"password,omitempty"`
	SshKey   string `json:"ssh_key,omitempty"`
	Username string `json:"username,omitempty"`
}

type Credentials2 struct {
	Password string `json:"password,omitempty"`
	SshKey   string `json:"ssh_key,omitempty"`
	Username string `json:"username,omitempty"`
}

type DatacenterPolicy struct {
	DatacenterHref string `json:"datacenter_href,omitempty"`
	Max            string `json:"max,omitempty"`
	Weight         string `json:"weight,omitempty"`
}

type DatacenterPolicy2 struct {
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
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	ServerTagScope string `json:"server_tag_scope,omitempty"`
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

type ElasticityParams2 struct {
	AlertSpecificParams *AlertSpecificParams2 `json:"alert_specific_params,omitempty"`
	Bounds              *Bounds2              `json:"bounds,omitempty"`
	Pacing              *Pacing2              `json:"pacing,omitempty"`
	QueueSpecificParams *QueueSpecificParams2 `json:"queue_specific_params,omitempty"`
	Schedule            []*Schedule2          `json:"schedule,omitempty"`
}

type InstanceParam struct {
	AssociatePublicIpAddress string                   `json:"associate_public_ip_address,omitempty"`
	CloudHref                string                   `json:"cloud_href,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                   `json:"datacenter_href,omitempty"`
	ImageHref                string                   `json:"image_href,omitempty"`
	Inputs                   map[string]string        `json:"inputs,omitempty"`
	InstanceTypeHref         string                   `json:"instance_type_href,omitempty"`
	IpForwardingEnabled      string                   `json:"ip_forwarding_enabled,omitempty"`
	KernelImageHref          string                   `json:"kernel_image_href,omitempty"`
	MultiCloudImageHref      string                   `json:"multi_cloud_image_href,omitempty"`
	PlacementGroupHref       string                   `json:"placement_group_href,omitempty"`
	RamdiskImageHref         string                   `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                 `json:"security_group_hrefs,omitempty"`
	ServerTemplateHref       string                   `json:"server_template_href,omitempty"`
	SshKeyHref               string                   `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                 `json:"subnet_hrefs,omitempty"`
	UserData                 string                   `json:"user_data,omitempty"`
}

type InstanceParam2 struct {
	AssociatePublicIpAddress string                    `json:"associate_public_ip_address,omitempty"`
	CloudHref                string                    `json:"cloud_href,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes2 `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                    `json:"datacenter_href,omitempty"`
	ImageHref                string                    `json:"image_href,omitempty"`
	Inputs                   map[string]string         `json:"inputs,omitempty"`
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

type InstanceParam3 struct {
	Href                string            `json:"href,omitempty"`
	Inputs              map[string]string `json:"inputs,omitempty"`
	MultiCloudImageHref string            `json:"multi_cloud_image_href,omitempty"`
	ServerTemplateHref  string            `json:"server_template_href,omitempty"`
}

type InstanceParam4 struct {
	AssociatePublicIpAddress string                    `json:"associate_public_ip_address,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes3 `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                    `json:"datacenter_href,omitempty"`
	DeploymentHref           string                    `json:"deployment_href,omitempty"`
	ImageHref                string                    `json:"image_href,omitempty"`
	InstanceTypeHref         string                    `json:"instance_type_href,omitempty"`
	KernelImageHref          string                    `json:"kernel_image_href,omitempty"`
	Name                     string                    `json:"name,omitempty"`
	PlacementGroupHref       string                    `json:"placement_group_href,omitempty"`
	RamdiskImageHref         string                    `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                  `json:"security_group_hrefs,omitempty"`
	SshKeyHref               string                    `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                  `json:"subnet_hrefs,omitempty"`
	UserData                 string                    `json:"user_data,omitempty"`
}

type InstanceParam5 struct {
	AssociatePublicIpAddress string                   `json:"associate_public_ip_address,omitempty"`
	CloudSpecificAttributes  *CloudSpecificAttributes `json:"cloud_specific_attributes,omitempty"`
	DatacenterHref           string                   `json:"datacenter_href,omitempty"`
	DeploymentHref           string                   `json:"deployment_href,omitempty"`
	ImageHref                string                   `json:"image_href,omitempty"`
	InstanceTypeHref         string                   `json:"instance_type_href,omitempty"`
	IpForwardingEnabled      string                   `json:"ip_forwarding_enabled,omitempty"`
	KernelImageHref          string                   `json:"kernel_image_href,omitempty"`
	MultiCloudImageHref      string                   `json:"multi_cloud_image_href,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	RamdiskImageHref         string                   `json:"ramdisk_image_href,omitempty"`
	SecurityGroupHrefs       []string                 `json:"security_group_hrefs,omitempty"`
	ServerTemplateHref       string                   `json:"server_template_href,omitempty"`
	SshKeyHref               string                   `json:"ssh_key_href,omitempty"`
	SubnetHrefs              []string                 `json:"subnet_hrefs,omitempty"`
	UserData                 string                   `json:"user_data,omitempty"`
}

type IpAddressBindingParam struct {
	InstanceHref        string `json:"instance_href,omitempty"`
	PrivatePort         string `json:"private_port,omitempty"`
	Protocol            string `json:"protocol,omitempty"`
	PublicIpAddressHref string `json:"public_ip_address_href,omitempty"`
	PublicPort          string `json:"public_port,omitempty"`
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

type ItemAge2 struct {
	Algorithm string `json:"algorithm,omitempty"`
	MaxAge    string `json:"max_age,omitempty"`
	Regexp    string `json:"regexp,omitempty"`
}

type MultiCloudImageParam struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type MultiCloudImageParam2 struct {
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

type MultiCloudImageSettingParam2 struct {
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
	CloudHref   string            `json:"cloud_href,omitempty"`
	Description string            `json:"description,omitempty"`
	Name        string            `json:"name,omitempty"`
	Options     map[string]string `json:"options,omitempty"`
	Type_       string            `json:"type,omitempty"`
}

type NetworkOptionGroupParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type NetworkParam struct {
	CidrBlock       string `json:"cidr_block,omitempty"`
	CloudHref       string `json:"cloud_href,omitempty"`
	Description     string `json:"description,omitempty"`
	InstanceTenancy string `json:"instance_tenancy,omitempty"`
	Name            string `json:"name,omitempty"`
}

type NetworkParam2 struct {
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	RouteTableHref string `json:"route_table_href,omitempty"`
}

type Pacing struct {
	ResizeCalmTime string `json:"resize_calm_time,omitempty"`
	ResizeDownBy   string `json:"resize_down_by,omitempty"`
	ResizeUpBy     string `json:"resize_up_by,omitempty"`
}

type Pacing2 struct {
	ResizeCalmTime string `json:"resize_calm_time,omitempty"`
	ResizeDownBy   string `json:"resize_down_by,omitempty"`
	ResizeUpBy     string `json:"resize_up_by,omitempty"`
}

type PermissionParam struct {
	RoleTitle string `json:"role_title,omitempty"`
	UserHref  string `json:"user_href,omitempty"`
}

type PlacementGroupParam struct {
	CloudHref   string `json:"cloud_href,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
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

type Quantity struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Quantity2 struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type QueueSize struct {
	ItemsPerInstance string `json:"items_per_instance,omitempty"`
}

type QueueSize2 struct {
	ItemsPerInstance string `json:"items_per_instance,omitempty"`
}

type QueueSpecificParams struct {
	CollectAuditEntries string     `json:"collect_audit_entries,omitempty"`
	ItemAge             *ItemAge   `json:"item_age,omitempty"`
	QueueSize           *QueueSize `json:"queue_size,omitempty"`
}

type QueueSpecificParams2 struct {
	CollectAuditEntries string      `json:"collect_audit_entries,omitempty"`
	ItemAge             *ItemAge2   `json:"item_age,omitempty"`
	QueueSize           *QueueSize2 `json:"queue_size,omitempty"`
}

type RecurringVolumeAttachmentParam struct {
	Device       string `json:"device,omitempty"`
	RunnableHref string `json:"runnable_href,omitempty"`
	StorageHref  string `json:"storage_href,omitempty"`
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
	AssetPaths      *AssetPaths2  `json:"asset_paths,omitempty"`
	CommitReference string        `json:"commit_reference,omitempty"`
	Credentials     *Credentials2 `json:"credentials,omitempty"`
	Description     string        `json:"description,omitempty"`
	Name            string        `json:"name,omitempty"`
	Source          string        `json:"source,omitempty"`
	SourceType      string        `json:"source_type,omitempty"`
}

type RightScriptParam struct {
	CommitMessage string `json:"commit_message,omitempty"`
}

type RightScriptParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type RouteParam struct {
	Description          string `json:"description,omitempty"`
	DestinationCidrBlock string `json:"destination_cidr_block,omitempty"`
	NextHopHref          string `json:"next_hop_href,omitempty"`
	NextHopIp            string `json:"next_hop_ip,omitempty"`
	NextHopType          string `json:"next_hop_type,omitempty"`
	RouteTableHref       string `json:"route_table_href,omitempty"`
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

type Schedule2 struct {
	Day      string `json:"day,omitempty"`
	MaxCount string `json:"max_count,omitempty"`
	MinCount string `json:"min_count,omitempty"`
	Time     string `json:"time,omitempty"`
}

type SecurityGroupParam struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	NetworkHref string `json:"network_href,omitempty"`
}

type SecurityGroupRuleParam struct {
	CidrIps           string           `json:"cidr_ips,omitempty"`
	Direction         string           `json:"direction,omitempty"`
	GroupName         string           `json:"group_name,omitempty"`
	GroupOwner        string           `json:"group_owner,omitempty"`
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
	Instance         *InstanceParam      `json:"instance,omitempty"`
	Name             string              `json:"name,omitempty"`
	Optimized        string              `json:"optimized,omitempty"`
	State            string              `json:"state,omitempty"`
}

type ServerArrayParam2 struct {
	ArrayType        string               `json:"array_type,omitempty"`
	DatacenterPolicy []*DatacenterPolicy2 `json:"datacenter_policy,omitempty"`
	DeploymentHref   string               `json:"deployment_href,omitempty"`
	Description      string               `json:"description,omitempty"`
	ElasticityParams *ElasticityParams2   `json:"elasticity_params,omitempty"`
	Name             string               `json:"name,omitempty"`
	Optimized        string               `json:"optimized,omitempty"`
	State            string               `json:"state,omitempty"`
}

type ServerParam struct {
	DeploymentHref string          `json:"deployment_href,omitempty"`
	Description    string          `json:"description,omitempty"`
	Instance       *InstanceParam2 `json:"instance,omitempty"`
	Name           string          `json:"name,omitempty"`
	Optimized      string          `json:"optimized,omitempty"`
}

type ServerParam2 struct {
	AutomaticInstanceStoreMapping string `json:"automatic_instance_store_mapping,omitempty"`
	Description                   string `json:"description,omitempty"`
	Name                          string `json:"name,omitempty"`
	Optimized                     string `json:"optimized,omitempty"`
	RootVolumeSize                string `json:"root_volume_size,omitempty"`
}

type ServerParam3 struct {
	DeploymentHref string          `json:"deployment_href,omitempty"`
	Description    string          `json:"description,omitempty"`
	Instance       *InstanceParam3 `json:"instance,omitempty"`
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

type ServerTemplateParam2 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
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
	NewEmail             string `json:"new_email,omitempty"`
	NewPassword          string `json:"new_password,omitempty"`
	Phone                string `json:"phone,omitempty"`
	PrincipalUid         string `json:"principal_uid,omitempty"`
	TimezoneName         string `json:"timezone_name,omitempty"`
}

type VolumeAttachmentParam struct {
	Device       string `json:"device,omitempty"`
	InstanceHref string `json:"instance_href,omitempty"`
	VolumeHref   string `json:"volume_href,omitempty"`
}

type VolumeParam struct {
	DatacenterHref           string `json:"datacenter_href,omitempty"`
	DeploymentHref           string `json:"deployment_href,omitempty"`
	Description              string `json:"description,omitempty"`
	Encrypted                string `json:"encrypted,omitempty"`
	Iops                     string `json:"iops,omitempty"`
	Name                     string `json:"name,omitempty"`
	ParentVolumeSnapshotHref string `json:"parent_volume_snapshot_href,omitempty"`
	PlacementGroupHref       string `json:"placement_group_href,omitempty"`
	Size                     string `json:"size,omitempty"`
	VolumeTypeHref           string `json:"volume_type_href,omitempty"`
}

type VolumeSnapshotParam struct {
	DeploymentHref   string `json:"deployment_href,omitempty"`
	Description      string `json:"description,omitempty"`
	Name             string `json:"name,omitempty"`
	ParentVolumeHref string `json:"parent_volume_href,omitempty"`
}
