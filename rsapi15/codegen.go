//************************************************************************//
//                     RightScale API 1.5 go client
//
// Generated Jan 29, 2015 at 3:37pm (PST)
// Command:
// $ api15gen -keep=T -metadata=../../rsapi15/api_data.json
// -attributes=../../rsapi15/attributes.json -output=../../rsapi15/codegen.go
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

// == AccountGroup ==

// An Account Group specifies which RightScale accounts will have access to
// import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.)
// from the MultiCloud Marketplace.
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
func (c *Client) IndexAccountGroups(filter []string, view string) ([]AccountGroup, error) {
	var res []AccountGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/account_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexAccountGroupsG(p *Params) ([]AccountGroup, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexAccountGroups(filter, view)
}

// GET api/account_groups/:id(.:format)?
// Show information about a single AccountGroup.
func (c *Client) ShowAccountGroup(view string, id string) (*AccountGroup, error) {
	var res *AccountGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/account_groups/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowAccountGroupG(p *Params) (*AccountGroup, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowAccountGroup(view, id)
}

// == Account ==

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
func (c *Client) ShowAccountG(p *Params) (*Account, error) {
	id := (*p)["id"].(string)
	return c.ShowAccount(id)
}

// == AlertSpec ==

// An AlertSpec defines the conditions under which an Alert is triggered
// and escalated. Condition sentence: if &lt;file&gt;.&lt;variable&gt;
// &lt;condition&gt; '&lt;threshold&gt;' for &lt;duration&gt; min then escalate
// to '&lt;escalation_name&gt;'.
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
func (c *Client) CreateAlertSpec(alertSpecVoteType string, alertSpecDuration string, alertSpecEscalationName string, alertSpecSubjectHref string, alertSpecThreshold string, alertSpecVariable string, alertSpecVoteTag string, alertSpecCondition string, alertSpecDescription string, alertSpecFile string, alertSpecName string, serverId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"alert_spec[vote_type]":       alertSpecVoteType,
		"alert_spec[duration]":        alertSpecDuration,
		"alert_spec[escalation_name]": alertSpecEscalationName,
		"alert_spec[subject_href]":    alertSpecSubjectHref,
		"alert_spec[threshold]":       alertSpecThreshold,
		"alert_spec[variable]":        alertSpecVariable,
		"alert_spec[vote_tag]":        alertSpecVoteTag,
		"alert_spec[condition]":       alertSpecCondition,
		"alert_spec[description]":     alertSpecDescription,
		"alert_spec[file]":            alertSpecFile,
		"alert_spec[name]":            alertSpecName}
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
func (c *Client) CreateAlertSpecG(p *Params) (Href, error) {
	alertSpecVoteType := (*p)["alertSpecVoteType"].(string)
	alertSpecDuration := (*p)["alertSpecDuration"].(string)
	alertSpecEscalationName := (*p)["alertSpecEscalationName"].(string)
	alertSpecSubjectHref := (*p)["alertSpecSubjectHref"].(string)
	alertSpecThreshold := (*p)["alertSpecThreshold"].(string)
	alertSpecVariable := (*p)["alertSpecVariable"].(string)
	alertSpecVoteTag := (*p)["alertSpecVoteTag"].(string)
	alertSpecCondition := (*p)["alertSpecCondition"].(string)
	alertSpecDescription := (*p)["alertSpecDescription"].(string)
	alertSpecFile := (*p)["alertSpecFile"].(string)
	alertSpecName := (*p)["alertSpecName"].(string)
	serverId := (*p)["serverId"].(string)
	return c.CreateAlertSpec(alertSpecVoteType, alertSpecDuration, alertSpecEscalationName, alertSpecSubjectHref, alertSpecThreshold, alertSpecVariable, alertSpecVoteTag, alertSpecCondition, alertSpecDescription, alertSpecFile, alertSpecName, serverId)
}

// DELETE api/servers/:server_id/alert_specs/:id(.:format)?
// Deletes a given AlertSpec.
func (c *Client) DestroyAlertSpec(serverId string, id string) error {
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
func (c *Client) DestroyAlertSpecG(p *Params) error {
	serverId := (*p)["serverId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyAlertSpec(serverId, id)
}

// GET api/servers/:server_id/alert_specs(.:format)?
// <no description>
func (c *Client) IndexAlertSpecs(view string, filter []string, withInherited string, serverId string) ([]AlertSpec, error) {
	var res []AlertSpec
	payload := map[string]interface{}{
		"with_inherited": withInherited}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers/"+serverId+"/alert_specs", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
	for _, v := range filter {
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
func (c *Client) IndexAlertSpecsG(p *Params) ([]AlertSpec, error) {
	view := (*p)["view"].(string)
	filter := (*p)["filter"].([]string)
	withInherited := (*p)["withInherited"].(string)
	serverId := (*p)["serverId"].(string)
	return c.IndexAlertSpecs(view, filter, withInherited, serverId)
}

// GET api/servers/:server_id/alert_specs/:id(.:format)?
// <no description>
func (c *Client) ShowAlertSpec(view string, serverId string, id string) (*AlertSpec, error) {
	var res *AlertSpec
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers/"+serverId+"/alert_specs/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowAlertSpecG(p *Params) (*AlertSpec, error) {
	view := (*p)["view"].(string)
	serverId := (*p)["serverId"].(string)
	id := (*p)["id"].(string)
	return c.ShowAlertSpec(view, serverId, id)
}

// PUT api/servers/:server_id/alert_specs/:id(.:format)?
// Updates an AlertSpec with the given parameters.
func (c *Client) UpdateAlertSpec(alertSpecVoteTag string, alertSpecCondition string, alertSpecEscalationName string, alertSpecFile string, alertSpecThreshold string, alertSpecVoteType string, alertSpecDescription string, alertSpecDuration string, alertSpecName string, alertSpecVariable string, serverId string, id string) error {
	payload := map[string]interface{}{
		"alert_spec[vote_tag]":        alertSpecVoteTag,
		"alert_spec[condition]":       alertSpecCondition,
		"alert_spec[escalation_name]": alertSpecEscalationName,
		"alert_spec[file]":            alertSpecFile,
		"alert_spec[threshold]":       alertSpecThreshold,
		"alert_spec[vote_type]":       alertSpecVoteType,
		"alert_spec[description]":     alertSpecDescription,
		"alert_spec[duration]":        alertSpecDuration,
		"alert_spec[name]":            alertSpecName,
		"alert_spec[variable]":        alertSpecVariable}
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
func (c *Client) UpdateAlertSpecG(p *Params) error {
	alertSpecVoteTag := (*p)["alertSpecVoteTag"].(string)
	alertSpecCondition := (*p)["alertSpecCondition"].(string)
	alertSpecEscalationName := (*p)["alertSpecEscalationName"].(string)
	alertSpecFile := (*p)["alertSpecFile"].(string)
	alertSpecThreshold := (*p)["alertSpecThreshold"].(string)
	alertSpecVoteType := (*p)["alertSpecVoteType"].(string)
	alertSpecDescription := (*p)["alertSpecDescription"].(string)
	alertSpecDuration := (*p)["alertSpecDuration"].(string)
	alertSpecName := (*p)["alertSpecName"].(string)
	alertSpecVariable := (*p)["alertSpecVariable"].(string)
	serverId := (*p)["serverId"].(string)
	id := (*p)["id"].(string)
	return c.UpdateAlertSpec(alertSpecVoteTag, alertSpecCondition, alertSpecEscalationName, alertSpecFile, alertSpecThreshold, alertSpecVoteType, alertSpecDescription, alertSpecDuration, alertSpecName, alertSpecVariable, serverId, id)
}

// == Alert ==

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
func (c *Client) DisableAlert(cloudId string, instanceId string, id string) error {
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
func (c *Client) DisableAlertG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.DisableAlert(cloudId, instanceId, id)
}

// POST api/clouds/:cloud_id/instances/:instance_id/alerts/:id/enable(.:format)?
// Enables the Alert indefinitely. Idempotent.
func (c *Client) EnableAlert(cloudId string, instanceId string, id string) error {
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
func (c *Client) EnableAlertG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.EnableAlert(cloudId, instanceId, id)
}

// GET api/clouds/:cloud_id/instances/:instance_id/alerts(.:format)?
// Lists all Alerts.
func (c *Client) IndexAlerts(filter []string, view string, cloudId string, instanceId string) ([]Alert, error) {
	var res []Alert
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/alerts", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexAlertsG(p *Params) ([]Alert, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.IndexAlerts(filter, view, cloudId, instanceId)
}

// POST api/clouds/:cloud_id/instances/:instance_id/alerts/:id/quench(.:format)?
// Suppresses the Alert from being triggered for a given time period.
// Idempotent.
func (c *Client) QuenchAlert(duration string, cloudId string, instanceId string, id string) (string, error) {
	var res string
	payload := map[string]interface{}{
		"duration": duration}
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
func (c *Client) QuenchAlertG(p *Params) (string, error) {
	duration := (*p)["duration"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.QuenchAlert(duration, cloudId, instanceId, id)
}

// GET api/clouds/:cloud_id/instances/:instance_id/alerts/:id(.:format)?
// Shows the attributes of a specified Alert.
func (c *Client) ShowAlert(view string, cloudId string, instanceId string, id string) (*Alert, error) {
	var res *Alert
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/alerts/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowAlertG(p *Params) (*Alert, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.ShowAlert(view, cloudId, instanceId, id)
}

// == AuditEntry ==

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
// Updates the summary and appends more details to a given AuditEntry. Each
// audit entry detail is stored as one chunk, the offset determines the location
// of that chunk within the overall audit entry details section. For example,
// if you create an AuditEntry and append "DEF" at offset 10, and later append
// "ABC" at offset 9, the overall audit entry details will be "ABCDEF". Use the
// \n character to separate details by new lines.
func (c *Client) AppendAuditEntry(detail string, notify string, offset int, summary string, id string) error {
	payload := map[string]interface{}{
		"detail":  detail,
		"notify":  notify,
		"offset":  offset,
		"summary": summary}
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
func (c *Client) AppendAuditEntryG(p *Params) error {
	detail := (*p)["detail"].(string)
	notify := (*p)["notify"].(string)
	offset := (*p)["offset"].(int)
	summary := (*p)["summary"].(string)
	id := (*p)["id"].(string)
	return c.AppendAuditEntry(detail, notify, offset, summary, id)
}

// POST api/audit_entries(.:format)?
// Creates a new AuditEntry with the given parameters.
func (c *Client) CreateAuditEntry(auditEntrySummary string, notify string, userEmail string, auditEntryAuditeeHref string, auditEntryDetail string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"audit_entry[summary]":      auditEntrySummary,
		"notify":                    notify,
		"user_email":                userEmail,
		"audit_entry[auditee_href]": auditEntryAuditeeHref,
		"audit_entry[detail]":       auditEntryDetail}
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
func (c *Client) CreateAuditEntryG(p *Params) (Href, error) {
	auditEntrySummary := (*p)["auditEntrySummary"].(string)
	notify := (*p)["notify"].(string)
	userEmail := (*p)["userEmail"].(string)
	auditEntryAuditeeHref := (*p)["auditEntryAuditeeHref"].(string)
	auditEntryDetail := (*p)["auditEntryDetail"].(string)
	return c.CreateAuditEntry(auditEntrySummary, notify, userEmail, auditEntryAuditeeHref, auditEntryDetail)
}

// GET api/audit_entries/:id/detail(.:format)?
// shows the details of a given AuditEntry. Note that the media type of the
// response is simply text.
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
func (c *Client) DetailAuditEntryG(p *Params) (string, error) {
	id := (*p)["id"].(string)
	return c.DetailAuditEntry(id)
}

// GET api/audit_entries(.:format)?
// Lists AuditEntries of the account. Due to the potentially large number of
// audit entries, a start and end date must be provided during an index call
// to limit the search. The format of the dates must be YYYY/MM/DD HH:MM:SS
// [+/-]ZZZZ e.g. 2011/07/11 00:00:00 +0000. A maximum of 1000 records will be
// returned by each index call.  Using the available filters, one can select or
// group which audit entries to retrieve.
func (c *Client) IndexAuditEntries(filter []string, view string, endDate string, limit string, startDate string) ([]AuditEntry, error) {
	var res []AuditEntry
	payload := map[string]interface{}{
		"end_date":   endDate,
		"limit":      limit,
		"start_date": startDate}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/audit_entries", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexAuditEntriesG(p *Params) ([]AuditEntry, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	endDate := (*p)["endDate"].(string)
	limit := (*p)["limit"].(string)
	startDate := (*p)["startDate"].(string)
	return c.IndexAuditEntries(filter, view, endDate, limit, startDate)
}

// GET api/audit_entries/:id(.:format)?
// Lists the attributes of a given audit entry.
func (c *Client) ShowAuditEntry(view string, id string) (*AuditEntry, error) {
	var res *AuditEntry
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/audit_entries/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowAuditEntryG(p *Params) (*AuditEntry, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowAuditEntry(view, id)
}

// PUT api/audit_entries/:id(.:format)?
// Updates the summary of a given AuditEntry.
func (c *Client) UpdateAuditEntry(notify string, auditEntryOffset int, auditEntrySummary string, id string) error {
	payload := map[string]interface{}{
		"notify":               notify,
		"audit_entry[offset]":  auditEntryOffset,
		"audit_entry[summary]": auditEntrySummary}
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
func (c *Client) UpdateAuditEntryG(p *Params) error {
	notify := (*p)["notify"].(string)
	auditEntryOffset := (*p)["auditEntryOffset"].(int)
	auditEntrySummary := (*p)["auditEntrySummary"].(string)
	id := (*p)["id"].(string)
	return c.UpdateAuditEntry(notify, auditEntryOffset, auditEntrySummary, id)
}

// == Backup ==

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
// Deletes old backups that meet the given criteria. For example, if a
// user calls cleanup with keep monthlies set to 12, then the latest backup
// for each month, for 12 months, will be kept.  All backups belong to a
// particular 'lineage'. Backups are not constrained to a specific cloud
// or a specific deployment. A lineage is account-specific. Hence, backups
// having the same lineage but belonging to different clouds are still
// considered for cleanup.  If backups specific to a single cloud should be
// cleaned up, see the cloud_href parameter.   Definitions: A snapshot is
// completed if its status is "available" A snapshot is committed if it has
// a tag "rs_backup:committed=true" A snapshot belongs to a backup "X" if
// it has a tag "rs_backup:backup_id=X" A snapshot is part of a backup with
// size "Y" if it has a tag "rs_backup:count=Y" A snapshot's position in a
// backup is "Z" if it has a tag "rs_backup:position=Z"  Backups are of 3
// types: Perfect backup: A backup which is completed (all the snapshots are
// completed) and committed (all the snapshots are committed) and the number
// of snapshots it found is equal to the number in the "rs_backup:count=" tag
// on each of the Snapshots. Imperfect backup: A backup which is not committed
// or if the number of snapshots it found is not equal to the number in the
// "rs_backup:count=" tag on each of the snapshots. Partial Perfect backup:
// A snapshot which is neither perfect nor imperfect  An imperfect backup is
// picked up for cleanup only if there exists a perfect backup with a newer
// created_at timestamp. No constraints will be applied on such imperfect
// backups and all of them will be destroyed.  For all the perfect backups,
// the constraints of keep_last and dailies etc. will be applied. The algorithm
// for choosing the perfect backups to keep is simple. It is the union of those
// set of backups if each of those conditions are applied independently. i.e
// backups_to_keep = backups_to_keep(keep_last) U backups_to_keep(dailies)
// U backups_to_keep(weeklies) U backups_to_keep(monthlies) U
// backups_to_keep(yearlies)  Hence, it is important to "commit" a backup to
// make it eligible for cleanup.
func (c *Client) CleanupBackup(monthlies string, weeklies string, yearlies string, cloudHref string, dailies string, keepLast string, lineage string) error {
	payload := map[string]interface{}{
		"monthlies":  monthlies,
		"weeklies":   weeklies,
		"yearlies":   yearlies,
		"cloud_href": cloudHref,
		"dailies":    dailies,
		"keep_last":  keepLast,
		"lineage":    lineage}
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
func (c *Client) CleanupBackupG(p *Params) error {
	monthlies := (*p)["monthlies"].(string)
	weeklies := (*p)["weeklies"].(string)
	yearlies := (*p)["yearlies"].(string)
	cloudHref := (*p)["cloudHref"].(string)
	dailies := (*p)["dailies"].(string)
	keepLast := (*p)["keepLast"].(string)
	lineage := (*p)["lineage"].(string)
	return c.CleanupBackup(monthlies, weeklies, yearlies, cloudHref, dailies, keepLast, lineage)
}

// POST api/backups(.:format)?
// Takes in an array of volumeattachmenthrefs and takes a snapshot of each. The
// volumeattachmenthrefs must belong to the same instance.
func (c *Client) CreateBackup(backupDescription string, backupFromMaster string, backupLineage string, backupName string, backupVolumeAttachmentHrefs []string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"backup[description]":             backupDescription,
		"backup[from_master]":             backupFromMaster,
		"backup[lineage]":                 backupLineage,
		"backup[name]":                    backupName,
		"backup[volume_attachment_hrefs]": backupVolumeAttachmentHrefs}
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
func (c *Client) CreateBackupG(p *Params) (Href, error) {
	backupDescription := (*p)["backupDescription"].(string)
	backupFromMaster := (*p)["backupFromMaster"].(string)
	backupLineage := (*p)["backupLineage"].(string)
	backupName := (*p)["backupName"].(string)
	backupVolumeAttachmentHrefs := (*p)["backupVolumeAttachmentHrefs"].([]string)
	return c.CreateBackup(backupDescription, backupFromMaster, backupLineage, backupName, backupVolumeAttachmentHrefs)
}

// DELETE api/backups/:id(.:format)?
// Deletes a given backup by deleting all of its snapshots, this call will
// succeed even if the backup has not completed.
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
func (c *Client) DestroyBackupG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyBackup(id)
}

// GET api/backups(.:format)?
// Lists all of the backups with the given lineage tag. Filters can be used to
// search for a particular backup. If the 'latest_before' filter is set, only
// one backup is returned (the latest backup before the given timestamp).  To
// get the latest completed backup, the 'completed' filter should be set to
// 'true' and the 'latest_before' filter should be set to the current timestamp.
// The format of the timestamp must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g.
// 2011/07/11 00:00:00 +0000.  To get the latest completed backup just before,
// say 25 June 2009, then the 'completed' filter should be set to 'true' and the
// 'latest_before' filter should be set to 2009/06/25 00:00:00 +0000.
func (c *Client) IndexBackups(filter []string, lineage string) ([]Backup, error) {
	var res []Backup
	payload := map[string]interface{}{
		"lineage": lineage}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/backups", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexBackupsG(p *Params) ([]Backup, error) {
	filter := (*p)["filter"].([]string)
	lineage := (*p)["lineage"].(string)
	return c.IndexBackups(filter, lineage)
}

// POST api/backups/:id/restore(.:format)?
// Restores the given Backup. This call will:   create the required number of
// Volumes from the volume_snapshots_hrefs in the given Backup, attach them to
// the given Instance at the device specified in the Snapshot. If the devices
// are already being used    on the Instance, the Task will denote that the
// restore has failed.
func (c *Client) RestoreBackup(backupVolumeTypeHref string, instanceHref string, backupDescription string, backupIops string, backupName string, backupSize string, id string) error {
	payload := map[string]interface{}{
		"backup[volume_type_href]": backupVolumeTypeHref,
		"instance_href":            instanceHref,
		"backup[description]":      backupDescription,
		"backup[iops]":             backupIops,
		"backup[name]":             backupName,
		"backup[size]":             backupSize}
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
func (c *Client) RestoreBackupG(p *Params) error {
	backupVolumeTypeHref := (*p)["backupVolumeTypeHref"].(string)
	instanceHref := (*p)["instanceHref"].(string)
	backupDescription := (*p)["backupDescription"].(string)
	backupIops := (*p)["backupIops"].(string)
	backupName := (*p)["backupName"].(string)
	backupSize := (*p)["backupSize"].(string)
	id := (*p)["id"].(string)
	return c.RestoreBackup(backupVolumeTypeHref, instanceHref, backupDescription, backupIops, backupName, backupSize, id)
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
func (c *Client) ShowBackupG(p *Params) (*Backup, error) {
	id := (*p)["id"].(string)
	return c.ShowBackup(id)
}

// PUT api/backups/:id(.:format)?
// Updates the committed tag for all of the VolumeSnapshots in the given Backup
// to the given value.
func (c *Client) UpdateBackup(backupCommitted string, id string) error {
	payload := map[string]interface{}{
		"backup[committed]": backupCommitted}
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
func (c *Client) UpdateBackupG(p *Params) error {
	backupCommitted := (*p)["backupCommitted"].(string)
	id := (*p)["id"].(string)
	return c.UpdateBackup(backupCommitted, id)
}

// == ChildAccount ==

type ChildAccount struct {
}

// POST api/child_accounts(.:format)?

func (c *Client) CreateChildAccount(childAccountClusterHref string, childAccountName string) (map[string]interface{}, error) {
	var res map[string]interface{}
	payload := map[string]interface{}{
		"child_account[cluster_href]": childAccountClusterHref,
		"child_account[name]":         childAccountName}
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
func (c *Client) CreateChildAccountG(p *Params) (map[string]interface{}, error) {
	childAccountClusterHref := (*p)["childAccountClusterHref"].(string)
	childAccountName := (*p)["childAccountName"].(string)
	return c.CreateChildAccount(childAccountClusterHref, childAccountName)
}

// GET api/child_accounts(.:format)?
// Lists the enterprise ChildAccounts available for this Account.
func (c *Client) IndexChildAccounts(filter []string) ([]map[string]string, error) {
	var res []map[string]string
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/child_accounts", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexChildAccountsG(p *Params) ([]map[string]string, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexChildAccounts(filter)
}

// PUT api/accounts/:id(.:format)?
// Update an enterprise ChildAccount for this Account.
func (c *Client) UpdateChildAccount(childAccountName string, id string) error {
	payload := map[string]interface{}{
		"child_account[name]": childAccountName}
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
func (c *Client) UpdateChildAccountG(p *Params) error {
	childAccountName := (*p)["childAccountName"].(string)
	id := (*p)["id"].(string)
	return c.UpdateChildAccount(childAccountName, id)
}

// == CloudAccount ==

// Represents a Cloud Account (An association between the account and a cloud).
type CloudAccount struct {
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

// POST api/cloud_accounts(.:format)?

func (c *Client) CreateCloudAccount(cloudAccountCloudHref string, cloudAccountCreds string, cloudAccountToken string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"cloud_account[cloud_href]": cloudAccountCloudHref,
		"cloud_account[creds][*]":   cloudAccountCreds,
		"cloud_account[token]":      cloudAccountToken}
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
func (c *Client) CreateCloudAccountG(p *Params) (Href, error) {
	cloudAccountCloudHref := (*p)["cloudAccountCloudHref"].(string)
	cloudAccountCreds := (*p)["cloudAccountCreds"].(string)
	cloudAccountToken := (*p)["cloudAccountToken"].(string)
	return c.CreateCloudAccount(cloudAccountCloudHref, cloudAccountCreds, cloudAccountToken)
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
func (c *Client) DestroyCloudAccountG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyCloudAccount(id)
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
func (c *Client) IndexCloudAccountsG(p *Params) ([]CloudAccount, error) {
	return c.IndexCloudAccounts()
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
func (c *Client) ShowCloudAccountG(p *Params) (*CloudAccount, error) {
	id := (*p)["id"].(string)
	return c.ShowCloudAccount(id)
}

// == Cloud ==

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
func (c *Client) IndexClouds(filter []string, view string) ([]Cloud, error) {
	var res []Cloud
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexCloudsG(p *Params) ([]Cloud, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexClouds(filter, view)
}

// GET api/clouds/:id(.:format)?
// Show information about a single cloud.
func (c *Client) ShowCloud(view string, id string) (*Cloud, error) {
	var res *Cloud
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowCloudG(p *Params) (*Cloud, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowCloud(view, id)
}

// == CookbookAttachment ==

// Cookbook Attachment is used to associate a particular cookbook with a
// ServerTemplate. A Cookbook Attachment must be in place before a recipe can be
// bound to a runlist using RunnableBinding.
type CookbookAttachment struct {
	Actions    []string            `json:"actions,omitempty"`
	Dependency bool                `json:"dependency,omitempty"`
	Id         string              `json:"id,omitempty"`
	Links      []map[string]string `json:"links,omitempty"`
}

// POST api/cookbooks/:cookbook_id/cookbook_attachments(.:format)?
// Attach a cookbook to a given resource.
func (c *Client) CreateCookbookAttachment(cookbookAttachmentCookbookHref string, cookbookAttachmentServerTemplateHref string, cookbookId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"cookbook_attachment[cookbook_href]":        cookbookAttachmentCookbookHref,
		"cookbook_attachment[server_template_href]": cookbookAttachmentServerTemplateHref}
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
func (c *Client) CreateCookbookAttachmentG(p *Params) (Href, error) {
	cookbookAttachmentCookbookHref := (*p)["cookbookAttachmentCookbookHref"].(string)
	cookbookAttachmentServerTemplateHref := (*p)["cookbookAttachmentServerTemplateHref"].(string)
	cookbookId := (*p)["cookbookId"].(string)
	return c.CreateCookbookAttachment(cookbookAttachmentCookbookHref, cookbookAttachmentServerTemplateHref, cookbookId)
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
func (c *Client) DestroyCookbookAttachmentG(p *Params) error {
	cookbookId := (*p)["cookbookId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyCookbookAttachment(cookbookId, id)
}

// GET api/cookbooks/:cookbook_id/cookbook_attachments(.:format)?
// Lists Cookbook Attachments.
func (c *Client) IndexCookbookAttachments(view string, cookbookId string) ([]CookbookAttachment, error) {
	var res []CookbookAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks/"+cookbookId+"/cookbook_attachments", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexCookbookAttachmentsG(p *Params) ([]CookbookAttachment, error) {
	view := (*p)["view"].(string)
	cookbookId := (*p)["cookbookId"].(string)
	return c.IndexCookbookAttachments(view, cookbookId)
}

// POST api/server_templates/:server_template_id/cookbook_attachments/multi_attach(.:format)?
// Attach multiple cookbooks to a given resource.
func (c *Client) MultiAttachCookbookAttachments(cookbookAttachmentsCookbookHrefs []string, cookbookAttachmentsServerTemplateHref string, serverTemplateId string) error {
	payload := map[string]interface{}{
		"cookbook_attachments[cookbook_hrefs]":       cookbookAttachmentsCookbookHrefs,
		"cookbook_attachments[server_template_href]": cookbookAttachmentsServerTemplateHref}
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
func (c *Client) MultiAttachCookbookAttachmentsG(p *Params) error {
	cookbookAttachmentsCookbookHrefs := (*p)["cookbookAttachmentsCookbookHrefs"].([]string)
	cookbookAttachmentsServerTemplateHref := (*p)["cookbookAttachmentsServerTemplateHref"].(string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	return c.MultiAttachCookbookAttachments(cookbookAttachmentsCookbookHrefs, cookbookAttachmentsServerTemplateHref, serverTemplateId)
}

// POST api/server_templates/:server_template_id/cookbook_attachments/multi_detach(.:format)?
// Detach multiple cookbooks from a given resource.
func (c *Client) MultiDetachCookbookAttachments(cookbookAttachmentsCookbookAttachmentHrefs []string, serverTemplateId string) error {
	payload := map[string]interface{}{
		"cookbook_attachments[cookbook_attachment_hrefs]": cookbookAttachmentsCookbookAttachmentHrefs}
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
func (c *Client) MultiDetachCookbookAttachmentsG(p *Params) error {
	cookbookAttachmentsCookbookAttachmentHrefs := (*p)["cookbookAttachmentsCookbookAttachmentHrefs"].([]string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	return c.MultiDetachCookbookAttachments(cookbookAttachmentsCookbookAttachmentHrefs, serverTemplateId)
}

// GET api/cookbooks/:cookbook_id/cookbook_attachments/:id(.:format)?
// Displays information about a single cookbook attachment to a ServerTemplate.
func (c *Client) ShowCookbookAttachment(view string, cookbookId string, id string) (*CookbookAttachment, error) {
	var res *CookbookAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks/"+cookbookId+"/cookbook_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowCookbookAttachmentG(p *Params) (*CookbookAttachment, error) {
	view := (*p)["view"].(string)
	cookbookId := (*p)["cookbookId"].(string)
	id := (*p)["id"].(string)
	return c.ShowCookbookAttachment(view, cookbookId, id)
}

// == Cookbook ==

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
// Destroys a Cookbook. Only available for cookbooks that have no Cookbook
// Attachments.
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
func (c *Client) DestroyCookbookG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyCookbook(id)
}

// POST api/cookbooks/:id/follow(.:format)?
// Follows (or unfollows) a Cookbook. Only available for cookbooks that are in
// the Alternate namespace.
func (c *Client) FollowCookbook(value string, id string) error {
	payload := map[string]interface{}{
		"value": value}
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
func (c *Client) FollowCookbookG(p *Params) error {
	value := (*p)["value"].(string)
	id := (*p)["id"].(string)
	return c.FollowCookbook(value, id)
}

// POST api/cookbooks/:id/freeze(.:format)?
// Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in
// the Primary namespace.
func (c *Client) FreezeCookbook(value string, id string) error {
	payload := map[string]interface{}{
		"value": value}
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
func (c *Client) FreezeCookbookG(p *Params) error {
	value := (*p)["value"].(string)
	id := (*p)["id"].(string)
	return c.FreezeCookbook(value, id)
}

// GET api/cookbooks(.:format)?
// Lists the Cookbooks available to this account.  The extended_designer view is
// only available to accounts with the designer permission.
func (c *Client) IndexCookbooks(filter []string, view string) ([]Cookbook, error) {
	var res []Cookbook
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexCookbooksG(p *Params) ([]Cookbook, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexCookbooks(filter, view)
}

// POST api/cookbooks/:id/obsolete(.:format)?
// Marks a Cookbook as obsolete (or un-obsolete).
func (c *Client) ObsoleteCookbook(value string, id string) error {
	payload := map[string]interface{}{
		"value": value}
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
func (c *Client) ObsoleteCookbookG(p *Params) error {
	value := (*p)["value"].(string)
	id := (*p)["id"].(string)
	return c.ObsoleteCookbook(value, id)
}

// GET api/cookbooks/:id(.:format)?
// Show information about a single Cookbook.  The extended_designer view is only
// available to accounts with the designer permission.
func (c *Client) ShowCookbook(view string, id string) (*Cookbook, error) {
	var res *Cookbook
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/cookbooks/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowCookbookG(p *Params) (*Cookbook, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowCookbook(view, id)
}

// == Credential ==

// A Credential provides an abstraction for sensitive textual information, such
// as passphrases or cloud credentials, whose value should be encrypted when
// stored in the database and not generally shown in the UI or through the API.
// Credentials may then be used as inputs of type "Cred" in RightScripts or Chef
// recipes. NOTE: Credential values may be updated through the API, but values
// cannot be retrieved via the API.
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
func (c *Client) CreateCredential(credentialName string, credentialValue string, credentialDescription string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"credential[name]":        credentialName,
		"credential[value]":       credentialValue,
		"credential[description]": credentialDescription}
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
func (c *Client) CreateCredentialG(p *Params) (Href, error) {
	credentialName := (*p)["credentialName"].(string)
	credentialValue := (*p)["credentialValue"].(string)
	credentialDescription := (*p)["credentialDescription"].(string)
	return c.CreateCredential(credentialName, credentialValue, credentialDescription)
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
func (c *Client) DestroyCredentialG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyCredential(id)
}

// GET api/credentials(.:format)?
// Lists the Credentials available to this account.
func (c *Client) IndexCredentials(filter []string, view string) ([]Credential, error) {
	var res []Credential
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/credentials", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexCredentialsG(p *Params) ([]Credential, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexCredentials(filter, view)
}

// GET api/credentials/:id(.:format)?
// Show information about a single Credential. NOTE: Credential values may be
// updated through the API, but values cannot be retrieved via the API.
func (c *Client) ShowCredential(view string, id string) (*Credential, error) {
	var res *Credential
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/credentials/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowCredentialG(p *Params) (*Credential, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowCredential(view, id)
}

// PUT api/credentials/:id(.:format)?
// Updates attributes of a Credential.
func (c *Client) UpdateCredential(credentialDescription string, credentialName string, credentialValue string, id string) error {
	payload := map[string]interface{}{
		"credential[description]": credentialDescription,
		"credential[name]":        credentialName,
		"credential[value]":       credentialValue}
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
func (c *Client) UpdateCredentialG(p *Params) error {
	credentialDescription := (*p)["credentialDescription"].(string)
	credentialName := (*p)["credentialName"].(string)
	credentialValue := (*p)["credentialValue"].(string)
	id := (*p)["id"].(string)
	return c.UpdateCredential(credentialDescription, credentialName, credentialValue, id)
}

// == Datacenter ==

// Datacenters represent isolated facilities within a cloud. The level and type
// of isolation is cloud dependent.  While Datacenters in large public clouds
// might correspond to different physical buildings, with different power,
// internet links...etc., Datacenters within the context of a private cloud
// might simply correspond to having different network providers.  Spreading
// servers across distinct Datacenters helps minimize outages.
type Datacenter struct {
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

// GET api/clouds/:cloud_id/datacenters(.:format)?
// Lists all Datacenters for a particular cloud.
func (c *Client) IndexDatacenters(filter []string, view string, cloudId string) ([]Datacenter, error) {
	var res []Datacenter
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/datacenters", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexDatacentersG(p *Params) ([]Datacenter, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexDatacenters(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/datacenters/:id(.:format)?
// Displays information about a single Datacenter.
func (c *Client) ShowDatacenter(view string, cloudId string, id string) (*Datacenter, error) {
	var res *Datacenter
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/datacenters/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowDatacenterG(p *Params) (*Datacenter, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowDatacenter(view, cloudId, id)
}

// == Deployment ==

// Deployments represent logical groupings of related assets such as servers,
// server arrays, default configuration settings...etc.
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
func (c *Client) CloneDeployment(deploymentDescription string, deploymentName string, deploymentServerTagScope string, id string) error {
	payload := map[string]interface{}{
		"deployment[description]":      deploymentDescription,
		"deployment[name]":             deploymentName,
		"deployment[server_tag_scope]": deploymentServerTagScope}
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
func (c *Client) CloneDeploymentG(p *Params) error {
	deploymentDescription := (*p)["deploymentDescription"].(string)
	deploymentName := (*p)["deploymentName"].(string)
	deploymentServerTagScope := (*p)["deploymentServerTagScope"].(string)
	id := (*p)["id"].(string)
	return c.CloneDeployment(deploymentDescription, deploymentName, deploymentServerTagScope, id)
}

// POST api/deployments(.:format)?
// Creates a new deployment with the given parameters.
func (c *Client) CreateDeployment(deploymentDescription string, deploymentName string, deploymentServerTagScope string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"deployment[description]":      deploymentDescription,
		"deployment[name]":             deploymentName,
		"deployment[server_tag_scope]": deploymentServerTagScope}
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
func (c *Client) CreateDeploymentG(p *Params) (Href, error) {
	deploymentDescription := (*p)["deploymentDescription"].(string)
	deploymentName := (*p)["deploymentName"].(string)
	deploymentServerTagScope := (*p)["deploymentServerTagScope"].(string)
	return c.CreateDeployment(deploymentDescription, deploymentName, deploymentServerTagScope)
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
func (c *Client) DestroyDeploymentG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyDeployment(id)
}

// GET api/deployments(.:format)?
// Lists deployments of the account.  Using the available filters, one can
// select or group which deployments to retrieve. The 'inputs_2_0' view is
// for retrieving inputs in 2.0 serialization (for more details please see
// Inputs#index.)
func (c *Client) IndexDeployments(filter []string, view string) ([]Deployment, error) {
	var res []Deployment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/deployments", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexDeploymentsG(p *Params) ([]Deployment, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexDeployments(filter, view)
}

// POST api/deployments/:id/lock(.:format)?
// Locks a given deployment. Idempotent. Locking prevents servers from being
// deleted or moved from the deployment. Other actions such as adding servers or
// renaming the deployment are still allowed.
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
func (c *Client) LockDeploymentG(p *Params) error {
	id := (*p)["id"].(string)
	return c.LockDeployment(id)
}

// GET api/deployments/:id/servers
// Lists the servers belonging to this deployment. This call is equivalent to
// servers#index call, where the servers returned will automatically be filtered
// by this deployment. See servers#index for details on other options and
// parameters.
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
func (c *Client) ServersDeploymentG(p *Params) error {
	id := (*p)["id"].(string)
	return c.ServersDeployment(id)
}

// GET api/deployments/:id(.:format)?
// Lists the attributes of a given deployment.  The 'inputs_2_0' view is
// for retrieving inputs in 2.0 serialization (for more details please see
// Inputs#index.)
func (c *Client) ShowDeployment(view string, id string) (*Deployment, error) {
	var res *Deployment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/deployments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowDeploymentG(p *Params) (*Deployment, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowDeployment(view, id)
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
func (c *Client) UnlockDeploymentG(p *Params) error {
	id := (*p)["id"].(string)
	return c.UnlockDeployment(id)
}

// PUT api/deployments/:id(.:format)?
// Updates attributes of a given deployment.
func (c *Client) UpdateDeployment(deploymentName string, deploymentServerTagScope string, deploymentDescription string, id string) error {
	payload := map[string]interface{}{
		"deployment[name]":             deploymentName,
		"deployment[server_tag_scope]": deploymentServerTagScope,
		"deployment[description]":      deploymentDescription}
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
func (c *Client) UpdateDeploymentG(p *Params) error {
	deploymentName := (*p)["deploymentName"].(string)
	deploymentServerTagScope := (*p)["deploymentServerTagScope"].(string)
	deploymentDescription := (*p)["deploymentDescription"].(string)
	id := (*p)["id"].(string)
	return c.UpdateDeployment(deploymentName, deploymentServerTagScope, deploymentDescription, id)
}

// == HealthCheck ==

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
func (c *Client) IndexHealthCheckG(p *Params) ([]map[string]string, error) {
	return c.IndexHealthCheck()
}

// == IdentityProvider ==

// An Identity Provider represents a SAML identity provider (IdP) that is linked
// to your RightScale Enterprise account, and is trusted by the RightScale
// dashboard to authenticate your enterprise's end users. To register an
// Identity Provider, contact your account manager.
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
func (c *Client) IndexIdentityProviders(filter []string, view string) ([]IdentityProvider, error) {
	var res []IdentityProvider
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/identity_providers", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexIdentityProvidersG(p *Params) ([]IdentityProvider, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexIdentityProviders(filter, view)
}

// GET api/identity_providers/:id(.:format)?
// Show the specified identity provider, if associated with this enterprise
// account.
func (c *Client) ShowIdentityProvider(view string, id string) (*IdentityProvider, error) {
	var res *IdentityProvider
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/identity_providers/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowIdentityProviderG(p *Params) (*IdentityProvider, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowIdentityProvider(view, id)
}

// == Image ==

// Images represent base VM image existing in a cloud. An image will define the
// initial Operating System and root disk contents  for a new Instance to have,
// and therefore it represents the basic starting point for creating a new one.
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
func (c *Client) IndexImages(filter []string, view string, cloudId string) ([]Image, error) {
	var res []Image
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/images", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexImagesG(p *Params) ([]Image, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexImages(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/images/:id(.:format)?
// Shows information about a single Image.
func (c *Client) ShowImage(view string, cloudId string, id string) (*Image, error) {
	var res *Image
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/images/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowImageG(p *Params) (*Image, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowImage(view, cloudId, id)
}

// == Input ==

// Inputs help extract dynamic information, usually specified at runtime,
// from repeatable configuration operations that can be codified. Inputs
// are variables defined in and used by RightScripts/Recipes. The two main
// attributes of an input are 'name' and 'value'. The 'name' identifies the
// input and the 'value', although a string encodes what type it is. It could
// be a text encoded as 'text:myvalue' or a credential encoded as 'cred:MY_CRED'
// or a key etc. Please see support.rightscale.com for more info on input
// hierarchies and their different types.
type Input struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// GET api/clouds/:cloud_id/instances/:instance_id/inputs(.:format)?
// Retrieves the full list of existing inputs of the specified resource.
func (c *Client) IndexInputs(view string, cloudId string, instanceId string) ([]Input, error) {
	var res []Input
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/inputs", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexInputsG(p *Params) ([]Input, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.IndexInputs(view, cloudId, instanceId)
}

// PUT api/clouds/:cloud_id/instances/:instance_id/inputs/multi_update(.:format)?
// Performs a bulk update of inputs on the specified resource.  If an input
// exists with the same name, its value will be updated. If an input does
// not exist with a specified name, it will be ignored.  Input values are
// represented as strings.  There are two notations for inputs:  1.0 notation
// - The deprecated notation used in API 1.0 and in 1.5     2.0 notation -
// The new notation that is partially supported in API 1.5, and will
//   be the only notation supported in 2.0  Although the two notations are
// similar, they have a few important differences, in particular:     With
// 2.0 notation, values MUST begin with a prefix identifying their type,
// followed   by a colon (example: "text:foo"). With 1.0 notation, unprefixed
// values are generally   taken to be text-type.   With 2.0 notation, a
// sentinel value "inherit" is used to express that an input   should use an
// inherited value. With 1.0 notation the empty string was used to express
//  the same thing. (Due to requirement 1, empty string is no longer a valid
// input.)   With 2.0 notation, each element of an array is an entire input
// value; arrays can   contain cred, env, or even other arrays. With 1.0
// notation, array elements are   implicitly text values and there is no way
// to specify anything else.Note that the UI   does not support complex-valued
// arrays; please use this feature with caution!   The following types of
// inputs are supported:    Type Format 1.0 Example(s) 2.0 Example(s)   Text
// string &lt;value&gt; (1.0 only)text:&lt;value&gt; footext:footext:multi
// word value text:footext:multi word value   Blank string(input is present
// but its value is empty-string) text:blank (2.0 only) text: blank
// Ignore (input is not present) ignore$ignore (1.0 only)ignore:$ignore
// (1.0 only) ignore$ignoreignore:$ignore ignore   Dynamically-substituted
// environment value env:&lt;value&gt;env:&lt;component&gt;:&lt;value&gt;
// env:MY_ENV_VARenv:my_server:MY_ENV_VAR env:MY_ENV_VARenv:my_server:MY_ENV_VAR
//   Credential value cred:&lt;value&gt; cred:abcd1234wxyz cred:abcd1234wxyz
//   Private SSH key key:&lt;value&gt;key:&lt;value&gt;:&lt;cloud_id&gt;
// key:1234abcd5678key:1234abcd5678:1 key:1234abcd5678key:1234abcd5678:1
//   Array of values array:&lt;value&gt;,... (1.0
// only)array:["&lt;type&gt;:&lt;value&gt;",...] (2.0 only)
// array:x,y(NOTE: 1.0 only supports text inputs for arrays)
// array:["text:v1","text:v2"]array:["text:x","env:server_x:MY_VAR"]    Note
// that in the case of array inputs, the portion after the colon must be valid
// JSON. In particular, when enclosing the input within double-quotes (e.g.
// for use in cURL or Ruby), the double-quotes must be escaped. Single-quotes
// may not be used within the array input, since they are not valid for JSON
// strings.  The legacy format for providing inputs is as an array of name-value
// pairs (ex: -d inputs[][name]="MY_INPUT" -d inputs[][value]="text:foobar"),
// however the new format is supported for inputs provided as a hash (ex: -d
// inputs[MY_INPUT]="text:foobar").  If the old format is used, the input is
// parsed using 1.0 semantics. If the new format is used, the input is parsed
// using the new 2.0 semantics.
func (c *Client) MultiUpdateInputs(inputsValue string, inputs string, inputsName string, cloudId string, instanceId string) error {
	payload := map[string]interface{}{
		"inputs[][value]": inputsValue,
		"inputs[*]":       inputs,
		"inputs[][name]":  inputsName}
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
func (c *Client) MultiUpdateInputsG(p *Params) error {
	inputsValue := (*p)["inputsValue"].(string)
	inputs := (*p)["inputs"].(string)
	inputsName := (*p)["inputsName"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.MultiUpdateInputs(inputsValue, inputs, inputsName, cloudId, instanceId)
}

// == InstanceCustomLodgement ==

// An InstanceCustomLodgement represents a way to create custom reports about
// a specific instance with a user defined quantity.  Replaces the legacy
// Instances#setcustomlodgement interface.
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
func (c *Client) CreateInstanceCustomLodgement(quantity []string, quantityName string, quantityValue string, timeframe string, cloudId string, instanceId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"quantity":          quantity,
		"quantity[][name]":  quantityName,
		"quantity[][value]": quantityValue,
		"timeframe":         timeframe}
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
func (c *Client) CreateInstanceCustomLodgementG(p *Params) (Href, error) {
	quantity := (*p)["quantity"].([]string)
	quantityName := (*p)["quantityName"].(string)
	quantityValue := (*p)["quantityValue"].(string)
	timeframe := (*p)["timeframe"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.CreateInstanceCustomLodgement(quantity, quantityName, quantityValue, timeframe, cloudId, instanceId)
}

// DELETE api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id(.:format)?
// Destroy the specified lodgement.
func (c *Client) DestroyInstanceCustomLodgement(cloudId string, instanceId string, id string) error {
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
func (c *Client) DestroyInstanceCustomLodgementG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyInstanceCustomLodgement(cloudId, instanceId, id)
}

// GET api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements(.:format)?
// List InstanceCustomLodgements of a given cloud and instance.
func (c *Client) IndexInstanceCustomLodgements(view string, cloudId string, instanceId string) ([]InstanceCustomLodgement, error) {
	var res []InstanceCustomLodgement
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/instance_custom_lodgements", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexInstanceCustomLodgementsG(p *Params) ([]InstanceCustomLodgement, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.IndexInstanceCustomLodgements(view, cloudId, instanceId)
}

// GET api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id(.:format)?
// Show the specified lodgement.
func (c *Client) ShowInstanceCustomLodgement(cloudId string, instanceId string, id string) (*InstanceCustomLodgement, error) {
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
func (c *Client) ShowInstanceCustomLodgementG(p *Params) (*InstanceCustomLodgement, error) {
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.ShowInstanceCustomLodgement(cloudId, instanceId, id)
}

// PUT api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id(.:format)?
// Update a lodgement with the quantity specified.
func (c *Client) UpdateInstanceCustomLodgement(quantity []string, quantityName string, quantityValue string, cloudId string, instanceId string, id string) error {
	payload := map[string]interface{}{
		"quantity":          quantity,
		"quantity[][name]":  quantityName,
		"quantity[][value]": quantityValue}
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
func (c *Client) UpdateInstanceCustomLodgementG(p *Params) error {
	quantity := (*p)["quantity"].([]string)
	quantityName := (*p)["quantityName"].(string)
	quantityValue := (*p)["quantityValue"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.UpdateInstanceCustomLodgement(quantity, quantityName, quantityValue, cloudId, instanceId, id)
}

// == InstanceType ==

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
func (c *Client) IndexInstanceTypes(filter []string, view string, cloudId string) ([]InstanceType, error) {
	var res []InstanceType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instance_types", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexInstanceTypesG(p *Params) ([]InstanceType, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexInstanceTypes(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/instance_types/:id(.:format)?
// Displays information about a single Instance type.
func (c *Client) ShowInstanceType(view string, cloudId string, id string) (*InstanceType, error) {
	var res *InstanceType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instance_types/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowInstanceTypeG(p *Params) (*InstanceType, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowInstanceType(view, cloudId, id)
}

// == Instance ==

// Instances represent an entity that is runnable in the cloud.  An instance
// of type "next" is a container of information that expresses how to configure
// a future instance when we decide to launch or start it. A "next" instance
// generally only exists in the RightScale realm, and usually doesn't have any
// corresponding representation existing in the cloud. However, if an instance
// is not of type "next", it will generally represent an existing running (or
// provisioned) virtual machine existing in the cloud.
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
func (c *Client) CreateInstance(instanceCloudSpecificAttributesEbsOptimized string, instanceCloudSpecificAttributesRootVolumeSize string, instanceCloudSpecificAttributesRootVolumeTypeUid string, instancePlacementGroupHref string, instanceAssociatePublicIpAddress string, instanceCloudSpecificAttributesRootVolumePerformance string, instanceDeploymentHref string, instanceSubnetHrefs []string, instanceUserData string, instanceSecurityGroupHrefs []string, instanceCloudSpecificAttributesAutomaticInstanceStoreMapping string, instanceCloudSpecificAttributesIamInstanceProfile string, instanceImageHref string, instanceInstanceTypeHref string, instanceName string, instanceRamdiskImageHref string, instanceSshKeyHref string, instanceDatacenterHref string, instanceKernelImageHref string, cloudId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"instance[cloud_specific_attributes][ebs_optimized]":                    instanceCloudSpecificAttributesEbsOptimized,
		"instance[cloud_specific_attributes][root_volume_size]":                 instanceCloudSpecificAttributesRootVolumeSize,
		"instance[cloud_specific_attributes][root_volume_type_uid]":             instanceCloudSpecificAttributesRootVolumeTypeUid,
		"instance[placement_group_href]":                                        instancePlacementGroupHref,
		"instance[associate_public_ip_address]":                                 instanceAssociatePublicIpAddress,
		"instance[cloud_specific_attributes][root_volume_performance]":          instanceCloudSpecificAttributesRootVolumePerformance,
		"instance[deployment_href]":                                             instanceDeploymentHref,
		"instance[subnet_hrefs]":                                                instanceSubnetHrefs,
		"instance[user_data]":                                                   instanceUserData,
		"instance[security_group_hrefs]":                                        instanceSecurityGroupHrefs,
		"instance[cloud_specific_attributes][automatic_instance_store_mapping]": instanceCloudSpecificAttributesAutomaticInstanceStoreMapping,
		"instance[cloud_specific_attributes][iam_instance_profile]":             instanceCloudSpecificAttributesIamInstanceProfile,
		"instance[image_href]":                                                  instanceImageHref,
		"instance[instance_type_href]":                                          instanceInstanceTypeHref,
		"instance[name]":                                                        instanceName,
		"instance[ramdisk_image_href]":                                          instanceRamdiskImageHref,
		"instance[ssh_key_href]":                                                instanceSshKeyHref,
		"instance[datacenter_href]":                                             instanceDatacenterHref,
		"instance[kernel_image_href]":                                           instanceKernelImageHref}
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
func (c *Client) CreateInstanceG(p *Params) (Href, error) {
	instanceCloudSpecificAttributesEbsOptimized := (*p)["instanceCloudSpecificAttributesEbsOptimized"].(string)
	instanceCloudSpecificAttributesRootVolumeSize := (*p)["instanceCloudSpecificAttributesRootVolumeSize"].(string)
	instanceCloudSpecificAttributesRootVolumeTypeUid := (*p)["instanceCloudSpecificAttributesRootVolumeTypeUid"].(string)
	instancePlacementGroupHref := (*p)["instancePlacementGroupHref"].(string)
	instanceAssociatePublicIpAddress := (*p)["instanceAssociatePublicIpAddress"].(string)
	instanceCloudSpecificAttributesRootVolumePerformance := (*p)["instanceCloudSpecificAttributesRootVolumePerformance"].(string)
	instanceDeploymentHref := (*p)["instanceDeploymentHref"].(string)
	instanceSubnetHrefs := (*p)["instanceSubnetHrefs"].([]string)
	instanceUserData := (*p)["instanceUserData"].(string)
	instanceSecurityGroupHrefs := (*p)["instanceSecurityGroupHrefs"].([]string)
	instanceCloudSpecificAttributesAutomaticInstanceStoreMapping := (*p)["instanceCloudSpecificAttributesAutomaticInstanceStoreMapping"].(string)
	instanceCloudSpecificAttributesIamInstanceProfile := (*p)["instanceCloudSpecificAttributesIamInstanceProfile"].(string)
	instanceImageHref := (*p)["instanceImageHref"].(string)
	instanceInstanceTypeHref := (*p)["instanceInstanceTypeHref"].(string)
	instanceName := (*p)["instanceName"].(string)
	instanceRamdiskImageHref := (*p)["instanceRamdiskImageHref"].(string)
	instanceSshKeyHref := (*p)["instanceSshKeyHref"].(string)
	instanceDatacenterHref := (*p)["instanceDatacenterHref"].(string)
	instanceKernelImageHref := (*p)["instanceKernelImageHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.CreateInstance(instanceCloudSpecificAttributesEbsOptimized, instanceCloudSpecificAttributesRootVolumeSize, instanceCloudSpecificAttributesRootVolumeTypeUid, instancePlacementGroupHref, instanceAssociatePublicIpAddress, instanceCloudSpecificAttributesRootVolumePerformance, instanceDeploymentHref, instanceSubnetHrefs, instanceUserData, instanceSecurityGroupHrefs, instanceCloudSpecificAttributesAutomaticInstanceStoreMapping, instanceCloudSpecificAttributesIamInstanceProfile, instanceImageHref, instanceInstanceTypeHref, instanceName, instanceRamdiskImageHref, instanceSshKeyHref, instanceDatacenterHref, instanceKernelImageHref, cloudId)
}

// GET api/clouds/:cloud_id/instances(.:format)?
// Lists instances of a given cloud, server array.  Using the available
// filters, it is possible to craft powerful queries about which instances
// to retrieve. For example, one can easily list:   instances that
// have names that contain "app" all instances of a given deployment
// instances belonging to a given server array (i.e., have the same
// parent_url)   To see the instances of a server array including
// the next_instance, use the URL "/api/clouds/:cloud_id/instances"
// with the filter "parent_href==/api/server_arrays/XX". To list
// only the running instances of a server array, use the URL
// "/api/server_arrays/:server_array_id/current_instances"  The
// 'full_inputs_2_0' view is for retrieving inputs in 2.0 serialization (for
// more details please see Inputs#index.)
func (c *Client) IndexInstances(view string, filter []string, cloudId string) ([]Instance, error) {
	var res []Instance
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
	for _, v := range filter {
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
func (c *Client) IndexInstancesG(p *Params) ([]Instance, error) {
	view := (*p)["view"].(string)
	filter := (*p)["filter"].([]string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexInstances(view, filter, cloudId)
}

// POST api/clouds/:cloud_id/instances/:id/launch(.:format)?
// Launches an instance using the parameters that this instance has been
// configured with.  Note that this action can only be performed in "next"
// instances, and not on instances that are already running.
func (c *Client) LaunchInstance(apiBehavior string, inputs string, inputsName string, inputsValue string, cloudId string, id string) error {
	payload := map[string]interface{}{
		"api_behavior":    apiBehavior,
		"inputs[*]":       inputs,
		"inputs[][name]":  inputsName,
		"inputs[][value]": inputsValue}
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
func (c *Client) LaunchInstanceG(p *Params) error {
	apiBehavior := (*p)["apiBehavior"].(string)
	inputs := (*p)["inputs"].(string)
	inputsName := (*p)["inputsName"].(string)
	inputsValue := (*p)["inputsValue"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.LaunchInstance(apiBehavior, inputs, inputsName, inputsValue, cloudId, id)
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
func (c *Client) LockInstanceG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.LockInstance(cloudId, id)
}

// POST api/clouds/:cloud_id/instances/multi_run_executable(.:format)?
// Runs a script or a recipe in the running instances.  This is an asynchronous
// function, which returns immediately after queuing the executable for
// execution. Status of the execution can be tracked at the URL returned in the
// "Location" header.
func (c *Client) MultiRunExecutableInstances(filter []string, ignoreLock string, inputs string, inputsName string, inputsValue string, recipeName string, rightScriptHref string, cloudId string) error {
	payload := map[string]interface{}{
		"ignore_lock":       ignoreLock,
		"inputs[*]":         inputs,
		"inputs[][name]":    inputsName,
		"inputs[][value]":   inputsValue,
		"recipe_name":       recipeName,
		"right_script_href": rightScriptHref}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/multi_run_executable", body)
	if err != nil {
		return err
	}
	for _, v := range filter {
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
func (c *Client) MultiRunExecutableInstancesG(p *Params) error {
	filter := (*p)["filter"].([]string)
	ignoreLock := (*p)["ignoreLock"].(string)
	inputs := (*p)["inputs"].(string)
	inputsName := (*p)["inputsName"].(string)
	inputsValue := (*p)["inputsValue"].(string)
	recipeName := (*p)["recipeName"].(string)
	rightScriptHref := (*p)["rightScriptHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.MultiRunExecutableInstances(filter, ignoreLock, inputs, inputsName, inputsValue, recipeName, rightScriptHref, cloudId)
}

// POST api/clouds/:cloud_id/instances/multi_terminate(.:format)?
// Terminates running instances. Either a filter or the parameter
// 'terminate_all' must be provided.
func (c *Client) MultiTerminateInstances(filter []string, terminateAll string, cloudId string) error {
	payload := map[string]interface{}{
		"terminate_all": terminateAll}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", c.endpoint+"/api/clouds/"+cloudId+"/instances/multi_terminate", body)
	if err != nil {
		return err
	}
	for _, v := range filter {
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
func (c *Client) MultiTerminateInstancesG(p *Params) error {
	filter := (*p)["filter"].([]string)
	terminateAll := (*p)["terminateAll"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.MultiTerminateInstances(filter, terminateAll, cloudId)
}

// POST api/clouds/:cloud_id/instances/:id/reboot(.:format)?
// Reboot a running instance.  Note that this action can only succeed if the
// instance is running. One cannot reboot instances of type "next".
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
func (c *Client) RebootInstanceG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.RebootInstance(cloudId, id)
}

// POST api/clouds/:cloud_id/instances/:id/run_executable(.:format)?
// Runs a script or a recipe in the running instance.  This is an asynchronous
// function, which returns immediately after queuing the executable for
// execution. Status of the execution can be tracked at the URL returned in the
// "Location" header. Note that this can only be performed on running instances.
func (c *Client) RunExecutableInstance(inputsValue string, recipeName string, rightScriptHref string, ignoreLock string, inputs string, inputsName string, cloudId string, id string) error {
	payload := map[string]interface{}{
		"inputs[][value]":   inputsValue,
		"recipe_name":       recipeName,
		"right_script_href": rightScriptHref,
		"ignore_lock":       ignoreLock,
		"inputs[*]":         inputs,
		"inputs[][name]":    inputsName}
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
func (c *Client) RunExecutableInstanceG(p *Params) error {
	inputsValue := (*p)["inputsValue"].(string)
	recipeName := (*p)["recipeName"].(string)
	rightScriptHref := (*p)["rightScriptHref"].(string)
	ignoreLock := (*p)["ignoreLock"].(string)
	inputs := (*p)["inputs"].(string)
	inputsName := (*p)["inputsName"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.RunExecutableInstance(inputsValue, recipeName, rightScriptHref, ignoreLock, inputs, inputsName, cloudId, id)
}

// POST api/clouds/:cloud_id/instances/:id/set_custom_lodgement(.:format)?
// This method is deprecated.  Please use InstanceCustomLodgement.
func (c *Client) SetCustomLodgementInstance(quantity []string, quantityName string, quantityValue string, timeframe string, cloudId string, id string) error {
	payload := map[string]interface{}{
		"quantity":          quantity,
		"quantity[][name]":  quantityName,
		"quantity[][value]": quantityValue,
		"timeframe":         timeframe}
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
func (c *Client) SetCustomLodgementInstanceG(p *Params) error {
	quantity := (*p)["quantity"].([]string)
	quantityName := (*p)["quantityName"].(string)
	quantityValue := (*p)["quantityValue"].(string)
	timeframe := (*p)["timeframe"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.SetCustomLodgementInstance(quantity, quantityName, quantityValue, timeframe, cloudId, id)
}

// GET api/clouds/:cloud_id/instances/:id(.:format)?
// Shows attributes of a single instance.  The 'full_inputs_2_0' view is
// for retrieving inputs in 2.0 serialization (for more details please see
// Inputs#index.)
func (c *Client) ShowInstance(view string, cloudId string, id string) (*Instance, error) {
	var res *Instance
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowInstanceG(p *Params) (*Instance, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowInstance(view, cloudId, id)
}

// POST api/clouds/:cloud_id/instances/:id/start(.:format)?
// Starts an instance that has been stopped, resuming it to its previously saved
// volume state.  After an instance is started, the reference to your instance
// will have a different id.  The new id can be found by performing an index
// query with the appropriate filters on the Instances resource, performing
// a show action on the Server resource for Server Instances, or performing
// a current_instances action on the ServerArray resource for ServerArray
// Instances.
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
func (c *Client) StartInstanceG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.StartInstance(cloudId, id)
}

// POST api/clouds/:cloud_id/instances/:id/stop(.:format)?
// Stores the instance's current volume state to resume later using the 'start'
// action.  After an instance is stopped, the reference to your instance
// will have a different id.  The new id can be found by performing an index
// query with the appropriate filters on the Instances resource, performing
// a show action on the Server resource for Server Instances, or performing
// a current_instances action on the ServerArray resource for ServerArray
// Instances.
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
func (c *Client) StopInstanceG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.StopInstance(cloudId, id)
}

// POST api/clouds/:cloud_id/instances/:id/terminate(.:format)?
// Terminates a running instance.  Note that this action can succeed only if the
// instance is running. One cannot terminate instances of type "next".
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
func (c *Client) TerminateInstanceG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.TerminateInstance(cloudId, id)
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
func (c *Client) UnlockInstanceG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.UnlockInstance(cloudId, id)
}

// PUT api/clouds/:cloud_id/instances/:id(.:format)?
// Updates attributes of a single instance.
func (c *Client) UpdateInstance(instanceCloudSpecificAttributesRootVolumePerformance string, instanceCloudSpecificAttributesRootVolumeTypeUid string, instanceUserData string, instanceName string, instanceRamdiskImageHref string, instanceSecurityGroupHrefs []string, instanceAssociatePublicIpAddress string, instanceKernelImageHref string, instanceMultiCloudImageHref string, instanceCloudSpecificAttributesRootVolumeSize string, instanceDeploymentHref string, instanceSubnetHrefs []string, instanceInstanceTypeHref string, instanceIpForwardingEnabled string, instanceServerTemplateHref string, instanceSshKeyHref string, instanceCloudSpecificAttributesAutomaticInstanceStoreMapping string, instanceCloudSpecificAttributesIamInstanceProfile string, instanceDatacenterHref string, instanceImageHref string, cloudId string, id string) error {
	payload := map[string]interface{}{
		"instance[cloud_specific_attributes][root_volume_performance]":          instanceCloudSpecificAttributesRootVolumePerformance,
		"instance[cloud_specific_attributes][root_volume_type_uid]":             instanceCloudSpecificAttributesRootVolumeTypeUid,
		"instance[user_data]":                                                   instanceUserData,
		"instance[name]":                                                        instanceName,
		"instance[ramdisk_image_href]":                                          instanceRamdiskImageHref,
		"instance[security_group_hrefs]":                                        instanceSecurityGroupHrefs,
		"instance[associate_public_ip_address]":                                 instanceAssociatePublicIpAddress,
		"instance[kernel_image_href]":                                           instanceKernelImageHref,
		"instance[multi_cloud_image_href]":                                      instanceMultiCloudImageHref,
		"instance[cloud_specific_attributes][root_volume_size]":                 instanceCloudSpecificAttributesRootVolumeSize,
		"instance[deployment_href]":                                             instanceDeploymentHref,
		"instance[subnet_hrefs]":                                                instanceSubnetHrefs,
		"instance[instance_type_href]":                                          instanceInstanceTypeHref,
		"instance[ip_forwarding_enabled]":                                       instanceIpForwardingEnabled,
		"instance[server_template_href]":                                        instanceServerTemplateHref,
		"instance[ssh_key_href]":                                                instanceSshKeyHref,
		"instance[cloud_specific_attributes][automatic_instance_store_mapping]": instanceCloudSpecificAttributesAutomaticInstanceStoreMapping,
		"instance[cloud_specific_attributes][iam_instance_profile]":             instanceCloudSpecificAttributesIamInstanceProfile,
		"instance[datacenter_href]":                                             instanceDatacenterHref,
		"instance[image_href]":                                                  instanceImageHref}
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
func (c *Client) UpdateInstanceG(p *Params) error {
	instanceCloudSpecificAttributesRootVolumePerformance := (*p)["instanceCloudSpecificAttributesRootVolumePerformance"].(string)
	instanceCloudSpecificAttributesRootVolumeTypeUid := (*p)["instanceCloudSpecificAttributesRootVolumeTypeUid"].(string)
	instanceUserData := (*p)["instanceUserData"].(string)
	instanceName := (*p)["instanceName"].(string)
	instanceRamdiskImageHref := (*p)["instanceRamdiskImageHref"].(string)
	instanceSecurityGroupHrefs := (*p)["instanceSecurityGroupHrefs"].([]string)
	instanceAssociatePublicIpAddress := (*p)["instanceAssociatePublicIpAddress"].(string)
	instanceKernelImageHref := (*p)["instanceKernelImageHref"].(string)
	instanceMultiCloudImageHref := (*p)["instanceMultiCloudImageHref"].(string)
	instanceCloudSpecificAttributesRootVolumeSize := (*p)["instanceCloudSpecificAttributesRootVolumeSize"].(string)
	instanceDeploymentHref := (*p)["instanceDeploymentHref"].(string)
	instanceSubnetHrefs := (*p)["instanceSubnetHrefs"].([]string)
	instanceInstanceTypeHref := (*p)["instanceInstanceTypeHref"].(string)
	instanceIpForwardingEnabled := (*p)["instanceIpForwardingEnabled"].(string)
	instanceServerTemplateHref := (*p)["instanceServerTemplateHref"].(string)
	instanceSshKeyHref := (*p)["instanceSshKeyHref"].(string)
	instanceCloudSpecificAttributesAutomaticInstanceStoreMapping := (*p)["instanceCloudSpecificAttributesAutomaticInstanceStoreMapping"].(string)
	instanceCloudSpecificAttributesIamInstanceProfile := (*p)["instanceCloudSpecificAttributesIamInstanceProfile"].(string)
	instanceDatacenterHref := (*p)["instanceDatacenterHref"].(string)
	instanceImageHref := (*p)["instanceImageHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.UpdateInstance(instanceCloudSpecificAttributesRootVolumePerformance, instanceCloudSpecificAttributesRootVolumeTypeUid, instanceUserData, instanceName, instanceRamdiskImageHref, instanceSecurityGroupHrefs, instanceAssociatePublicIpAddress, instanceKernelImageHref, instanceMultiCloudImageHref, instanceCloudSpecificAttributesRootVolumeSize, instanceDeploymentHref, instanceSubnetHrefs, instanceInstanceTypeHref, instanceIpForwardingEnabled, instanceServerTemplateHref, instanceSshKeyHref, instanceCloudSpecificAttributesAutomaticInstanceStoreMapping, instanceCloudSpecificAttributesIamInstanceProfile, instanceDatacenterHref, instanceImageHref, cloudId, id)
}

// == IpAddressBinding ==

// An IpAddressBinding represents an abstraction for binding an IpAddress to an
// instance. The IpAddress is bound immediately for a current instance, or on
// launch for a next instance. It also allows specifying port forwarding rules
// for that particular IpAddress and Instance pair.
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
// IpAddressBinding will be created. If the instance is a next instance, then a
// recurring IpAddressBinding is created, which will cause the IpAddress to be
// bound each time the incarnator boots.
func (c *Client) CreateIpAddressBinding(ipAddressBindingPublicIpAddressHref string, ipAddressBindingPublicPort string, ipAddressBindingInstanceHref string, ipAddressBindingPrivatePort string, ipAddressBindingProtocol string, cloudId string, ipAddressId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"ip_address_binding[public_ip_address_href]": ipAddressBindingPublicIpAddressHref,
		"ip_address_binding[public_port]":            ipAddressBindingPublicPort,
		"ip_address_binding[instance_href]":          ipAddressBindingInstanceHref,
		"ip_address_binding[private_port]":           ipAddressBindingPrivatePort,
		"ip_address_binding[protocol]":               ipAddressBindingProtocol}
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
func (c *Client) CreateIpAddressBindingG(p *Params) (Href, error) {
	ipAddressBindingPublicIpAddressHref := (*p)["ipAddressBindingPublicIpAddressHref"].(string)
	ipAddressBindingPublicPort := (*p)["ipAddressBindingPublicPort"].(string)
	ipAddressBindingInstanceHref := (*p)["ipAddressBindingInstanceHref"].(string)
	ipAddressBindingPrivatePort := (*p)["ipAddressBindingPrivatePort"].(string)
	ipAddressBindingProtocol := (*p)["ipAddressBindingProtocol"].(string)
	cloudId := (*p)["cloudId"].(string)
	ipAddressId := (*p)["ipAddressId"].(string)
	return c.CreateIpAddressBinding(ipAddressBindingPublicIpAddressHref, ipAddressBindingPublicPort, ipAddressBindingInstanceHref, ipAddressBindingPrivatePort, ipAddressBindingProtocol, cloudId, ipAddressId)
}

// DELETE api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id(.:format)?
// <no description>
func (c *Client) DestroyIpAddressBinding(cloudId string, ipAddressId string, id string) error {
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
func (c *Client) DestroyIpAddressBindingG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	ipAddressId := (*p)["ipAddressId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyIpAddressBinding(cloudId, ipAddressId, id)
}

// GET api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings(.:format)?
// Lists the ip address bindings available to this account.
func (c *Client) IndexIpAddressBindings(filter []string, cloudId string, ipAddressId string) ([]IpAddressBinding, error) {
	var res []IpAddressBinding
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses/"+ipAddressId+"/ip_address_bindings", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexIpAddressBindingsG(p *Params) ([]IpAddressBinding, error) {
	filter := (*p)["filter"].([]string)
	cloudId := (*p)["cloudId"].(string)
	ipAddressId := (*p)["ipAddressId"].(string)
	return c.IndexIpAddressBindings(filter, cloudId, ipAddressId)
}

// GET api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id(.:format)?
// Show information about a single ip address binding.
func (c *Client) ShowIpAddressBinding(cloudId string, ipAddressId string, id string) (*IpAddressBinding, error) {
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
func (c *Client) ShowIpAddressBindingG(p *Params) (*IpAddressBinding, error) {
	cloudId := (*p)["cloudId"].(string)
	ipAddressId := (*p)["ipAddressId"].(string)
	id := (*p)["id"].(string)
	return c.ShowIpAddressBinding(cloudId, ipAddressId, id)
}

// == IpAddress ==

// An IpAddress provides an abstraction for IPv4 addresses bindable to Instance
// resources running in a Cloud.
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
func (c *Client) CreateIpAddress(ipAddressDeploymentHref string, ipAddressDomain string, ipAddressName string, ipAddressNetworkHref string, cloudId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"ip_address[deployment_href]": ipAddressDeploymentHref,
		"ip_address[domain]":          ipAddressDomain,
		"ip_address[name]":            ipAddressName,
		"ip_address[network_href]":    ipAddressNetworkHref}
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
func (c *Client) CreateIpAddressG(p *Params) (Href, error) {
	ipAddressDeploymentHref := (*p)["ipAddressDeploymentHref"].(string)
	ipAddressDomain := (*p)["ipAddressDomain"].(string)
	ipAddressName := (*p)["ipAddressName"].(string)
	ipAddressNetworkHref := (*p)["ipAddressNetworkHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.CreateIpAddress(ipAddressDeploymentHref, ipAddressDomain, ipAddressName, ipAddressNetworkHref, cloudId)
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
func (c *Client) DestroyIpAddressG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyIpAddress(cloudId, id)
}

// GET api/clouds/:cloud_id/ip_addresses(.:format)?
// Lists the IpAddresses available to this account.
func (c *Client) IndexIpAddresses(filter []string, cloudId string) ([]IpAddress, error) {
	var res []IpAddress
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ip_addresses", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexIpAddressesG(p *Params) ([]IpAddress, error) {
	filter := (*p)["filter"].([]string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexIpAddresses(filter, cloudId)
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
func (c *Client) ShowIpAddressG(p *Params) (*IpAddress, error) {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowIpAddress(cloudId, id)
}

// PUT api/clouds/:cloud_id/ip_addresses/:id(.:format)?
// Updates attributes of a given IpAddress.
func (c *Client) UpdateIpAddress(ipAddressName string, ipAddressDeploymentHref string, cloudId string, id string) error {
	payload := map[string]interface{}{
		"ip_address[name]":            ipAddressName,
		"ip_address[deployment_href]": ipAddressDeploymentHref}
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
func (c *Client) UpdateIpAddressG(p *Params) error {
	ipAddressName := (*p)["ipAddressName"].(string)
	ipAddressDeploymentHref := (*p)["ipAddressDeploymentHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.UpdateIpAddress(ipAddressName, ipAddressDeploymentHref, cloudId, id)
}

// == MonitoringMetric ==

// A monitoring metric is a stream of data that is captured in an instance.
// Metrics can be monitored, graphed and can be used as the basis for triggering
// alerts.
type MonitoringMetric struct {
	Actions   []string            `json:"actions,omitempty"`
	GraphHref string              `json:"graph_href,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	Plugin    string              `json:"plugin,omitempty"`
	View      string              `json:"view,omitempty"`
}

// GET api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id/data(.:format)?
// Gives the raw monitoring data for a particular metric. The response will
// include different variables associated with that metric and the data points
// for each of those variables.  To get the data for a certain duration, for
// e.g. for the last 10 minutes(600 secs), provide the variables start="-600"
// and end="0".
func (c *Client) DataMonitoringMetric(end string, start string, cloudId string, instanceId string, id string) (map[string]string, error) {
	var res map[string]string
	payload := map[string]interface{}{
		"end":   end,
		"start": start}
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
func (c *Client) DataMonitoringMetricG(p *Params) (map[string]string, error) {
	end := (*p)["end"].(string)
	start := (*p)["start"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.DataMonitoringMetric(end, start, cloudId, instanceId, id)
}

// GET api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics(.:format)?
// Lists the monitoring metrics available for the instance and their
// corresponding graph hrefs. Making a request to the graph_href will return a
// png image corresponding to that monitoring metric.
func (c *Client) IndexMonitoringMetrics(filter []string, size string, title string, tz string, period string, cloudId string, instanceId string) ([]MonitoringMetric, error) {
	var res []MonitoringMetric
	payload := map[string]interface{}{
		"size":   size,
		"title":  title,
		"tz":     tz,
		"period": period}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/monitoring_metrics", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexMonitoringMetricsG(p *Params) ([]MonitoringMetric, error) {
	filter := (*p)["filter"].([]string)
	size := (*p)["size"].(string)
	title := (*p)["title"].(string)
	tz := (*p)["tz"].(string)
	period := (*p)["period"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.IndexMonitoringMetrics(filter, size, title, tz, period, cloudId, instanceId)
}

// GET api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id(.:format)?
// Shows attributes of a single monitoring metric. Making a request to the
// graph_href will return a png image corresponding to that monitoring metric.
func (c *Client) ShowMonitoringMetric(period string, size string, title string, tz string, cloudId string, instanceId string, id string) (*MonitoringMetric, error) {
	var res *MonitoringMetric
	payload := map[string]interface{}{
		"period": period,
		"size":   size,
		"title":  title,
		"tz":     tz}
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
func (c *Client) ShowMonitoringMetricG(p *Params) (*MonitoringMetric, error) {
	period := (*p)["period"].(string)
	size := (*p)["size"].(string)
	title := (*p)["title"].(string)
	tz := (*p)["tz"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.ShowMonitoringMetric(period, size, title, tz, cloudId, instanceId, id)
}

// == MultiCloudImageSetting ==

// A MultiCloudImageSetting defines which settings should be used when a server
// is launched in a cloud.
type MultiCloudImageSetting struct {
	Actions []string            `json:"actions,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
}

// POST api/multi_cloud_images/:multi_cloud_image_id/settings(.:format)?
// Creates a new setting for an existing MultiCloudImage.
func (c *Client) CreateMultiCloudImageSetting(multiCloudImageSettingCloudHref string, multiCloudImageSettingImageHref string, multiCloudImageSettingInstanceTypeHref string, multiCloudImageSettingKernelImageHref string, multiCloudImageSettingRamdiskImageHref string, multiCloudImageSettingUserData string, multiCloudImageId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"multi_cloud_image_setting[cloud_href]":         multiCloudImageSettingCloudHref,
		"multi_cloud_image_setting[image_href]":         multiCloudImageSettingImageHref,
		"multi_cloud_image_setting[instance_type_href]": multiCloudImageSettingInstanceTypeHref,
		"multi_cloud_image_setting[kernel_image_href]":  multiCloudImageSettingKernelImageHref,
		"multi_cloud_image_setting[ramdisk_image_href]": multiCloudImageSettingRamdiskImageHref,
		"multi_cloud_image_setting[user_data]":          multiCloudImageSettingUserData}
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
func (c *Client) CreateMultiCloudImageSettingG(p *Params) (Href, error) {
	multiCloudImageSettingCloudHref := (*p)["multiCloudImageSettingCloudHref"].(string)
	multiCloudImageSettingImageHref := (*p)["multiCloudImageSettingImageHref"].(string)
	multiCloudImageSettingInstanceTypeHref := (*p)["multiCloudImageSettingInstanceTypeHref"].(string)
	multiCloudImageSettingKernelImageHref := (*p)["multiCloudImageSettingKernelImageHref"].(string)
	multiCloudImageSettingRamdiskImageHref := (*p)["multiCloudImageSettingRamdiskImageHref"].(string)
	multiCloudImageSettingUserData := (*p)["multiCloudImageSettingUserData"].(string)
	multiCloudImageId := (*p)["multiCloudImageId"].(string)
	return c.CreateMultiCloudImageSetting(multiCloudImageSettingCloudHref, multiCloudImageSettingImageHref, multiCloudImageSettingInstanceTypeHref, multiCloudImageSettingKernelImageHref, multiCloudImageSettingRamdiskImageHref, multiCloudImageSettingUserData, multiCloudImageId)
}

// DELETE api/multi_cloud_images/:multi_cloud_image_id/settings/:id(.:format)?
// Deletes a MultiCloudImage setting.
func (c *Client) DestroyMultiCloudImageSetting(multiCloudImageId string, id string) error {
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
func (c *Client) DestroyMultiCloudImageSettingG(p *Params) error {
	multiCloudImageId := (*p)["multiCloudImageId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyMultiCloudImageSetting(multiCloudImageId, id)
}

// GET api/multi_cloud_images/:multi_cloud_image_id/settings(.:format)?
// Lists the MultiCloudImage settings.
func (c *Client) IndexMultiCloudImageSettings(filter []string, multiCloudImageId string) ([]MultiCloudImageSetting, error) {
	var res []MultiCloudImageSetting
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/multi_cloud_images/"+multiCloudImageId+"/settings", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexMultiCloudImageSettingsG(p *Params) ([]MultiCloudImageSetting, error) {
	filter := (*p)["filter"].([]string)
	multiCloudImageId := (*p)["multiCloudImageId"].(string)
	return c.IndexMultiCloudImageSettings(filter, multiCloudImageId)
}

// GET api/multi_cloud_images/:multi_cloud_image_id/settings/:id(.:format)?
// Show information about a single MultiCloudImage setting.
func (c *Client) ShowMultiCloudImageSetting(multiCloudImageId string, id string) (*MultiCloudImageSetting, error) {
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
func (c *Client) ShowMultiCloudImageSettingG(p *Params) (*MultiCloudImageSetting, error) {
	multiCloudImageId := (*p)["multiCloudImageId"].(string)
	id := (*p)["id"].(string)
	return c.ShowMultiCloudImageSetting(multiCloudImageId, id)
}

// PUT api/multi_cloud_images/:multi_cloud_image_id/settings/:id(.:format)?
// Updates a settings for a MultiCLoudImage.
func (c *Client) UpdateMultiCloudImageSetting(multiCloudImageSettingKernelImageHref string, multiCloudImageSettingRamdiskImageHref string, multiCloudImageSettingUserData string, multiCloudImageSettingCloudHref string, multiCloudImageSettingImageHref string, multiCloudImageSettingInstanceTypeHref string, multiCloudImageId string, id string) error {
	payload := map[string]interface{}{
		"multi_cloud_image_setting[kernel_image_href]":  multiCloudImageSettingKernelImageHref,
		"multi_cloud_image_setting[ramdisk_image_href]": multiCloudImageSettingRamdiskImageHref,
		"multi_cloud_image_setting[user_data]":          multiCloudImageSettingUserData,
		"multi_cloud_image_setting[cloud_href]":         multiCloudImageSettingCloudHref,
		"multi_cloud_image_setting[image_href]":         multiCloudImageSettingImageHref,
		"multi_cloud_image_setting[instance_type_href]": multiCloudImageSettingInstanceTypeHref}
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
func (c *Client) UpdateMultiCloudImageSettingG(p *Params) error {
	multiCloudImageSettingKernelImageHref := (*p)["multiCloudImageSettingKernelImageHref"].(string)
	multiCloudImageSettingRamdiskImageHref := (*p)["multiCloudImageSettingRamdiskImageHref"].(string)
	multiCloudImageSettingUserData := (*p)["multiCloudImageSettingUserData"].(string)
	multiCloudImageSettingCloudHref := (*p)["multiCloudImageSettingCloudHref"].(string)
	multiCloudImageSettingImageHref := (*p)["multiCloudImageSettingImageHref"].(string)
	multiCloudImageSettingInstanceTypeHref := (*p)["multiCloudImageSettingInstanceTypeHref"].(string)
	multiCloudImageId := (*p)["multiCloudImageId"].(string)
	id := (*p)["id"].(string)
	return c.UpdateMultiCloudImageSetting(multiCloudImageSettingKernelImageHref, multiCloudImageSettingRamdiskImageHref, multiCloudImageSettingUserData, multiCloudImageSettingCloudHref, multiCloudImageSettingImageHref, multiCloudImageSettingInstanceTypeHref, multiCloudImageId, id)
}

// == MultiCloudImage ==

// A MultiCloudImage is a RightScale component that functions as a pointer
// to machine images in specific clouds (e.g. AWS US-East, Rackspace). Each
// ServerTemplate can reference many MultiCloudImages that defines which image
// should be used when a server is launched in a particular cloud.
type MultiCloudImage struct {
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Revision    string              `json:"revision,omitempty"`
}

// POST api/multi_cloud_images/:id/clone(.:format)?
// Clones a given MultiCloudImage.
func (c *Client) CloneMultiCloudImage(multiCloudImageName string, multiCloudImageDescription string, id string) error {
	payload := map[string]interface{}{
		"multi_cloud_image[name]":        multiCloudImageName,
		"multi_cloud_image[description]": multiCloudImageDescription}
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
func (c *Client) CloneMultiCloudImageG(p *Params) error {
	multiCloudImageName := (*p)["multiCloudImageName"].(string)
	multiCloudImageDescription := (*p)["multiCloudImageDescription"].(string)
	id := (*p)["id"].(string)
	return c.CloneMultiCloudImage(multiCloudImageName, multiCloudImageDescription, id)
}

// POST api/multi_cloud_images/:id/commit(.:format)?
// Commits a given MultiCloudImage. Only HEAD revisions can be committed.
func (c *Client) CommitMultiCloudImage(commitMessage string, id string) error {
	payload := map[string]interface{}{
		"commit_message": commitMessage}
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
func (c *Client) CommitMultiCloudImageG(p *Params) error {
	commitMessage := (*p)["commitMessage"].(string)
	id := (*p)["id"].(string)
	return c.CommitMultiCloudImage(commitMessage, id)
}

// POST api/server_templates/:server_template_id/multi_cloud_images(.:format)?
// Creates a new MultiCloudImage with the given parameters.
func (c *Client) CreateMultiCloudImage(multiCloudImageDescription string, multiCloudImageName string, serverTemplateId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"multi_cloud_image[description]": multiCloudImageDescription,
		"multi_cloud_image[name]":        multiCloudImageName}
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
func (c *Client) CreateMultiCloudImageG(p *Params) (Href, error) {
	multiCloudImageDescription := (*p)["multiCloudImageDescription"].(string)
	multiCloudImageName := (*p)["multiCloudImageName"].(string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	return c.CreateMultiCloudImage(multiCloudImageDescription, multiCloudImageName, serverTemplateId)
}

// DELETE api/server_templates/:server_template_id/multi_cloud_images/:id(.:format)?
// Deletes a given MultiCloudImage.
func (c *Client) DestroyMultiCloudImage(serverTemplateId string, id string) error {
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
func (c *Client) DestroyMultiCloudImageG(p *Params) error {
	serverTemplateId := (*p)["serverTemplateId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyMultiCloudImage(serverTemplateId, id)
}

// GET api/server_templates/:server_template_id/multi_cloud_images(.:format)?
// Lists the MultiCloudImages available to this account. HEAD revisions have a
// revision of 0.
func (c *Client) IndexMultiCloudImages(filter []string, serverTemplateId string) ([]MultiCloudImage, error) {
	var res []MultiCloudImage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+serverTemplateId+"/multi_cloud_images", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexMultiCloudImagesG(p *Params) ([]MultiCloudImage, error) {
	filter := (*p)["filter"].([]string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	return c.IndexMultiCloudImages(filter, serverTemplateId)
}

// GET api/server_templates/:server_template_id/multi_cloud_images/:id(.:format)?
// Show information about a single MultiCloudImage. HEAD revisions have a
// revision of 0.
func (c *Client) ShowMultiCloudImage(serverTemplateId string, id string) (*MultiCloudImage, error) {
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
func (c *Client) ShowMultiCloudImageG(p *Params) (*MultiCloudImage, error) {
	serverTemplateId := (*p)["serverTemplateId"].(string)
	id := (*p)["id"].(string)
	return c.ShowMultiCloudImage(serverTemplateId, id)
}

// PUT api/server_templates/:server_template_id/multi_cloud_images/:id(.:format)?
// Updates attributes of a given MultiCloudImage. Only HEAD revisions can be
// updated (revision 0). Currently, the attributes you can update are only the
// 'direct' attributes of a server template.
func (c *Client) UpdateMultiCloudImage(multiCloudImageDescription string, multiCloudImageName string, serverTemplateId string, id string) error {
	payload := map[string]interface{}{
		"multi_cloud_image[description]": multiCloudImageDescription,
		"multi_cloud_image[name]":        multiCloudImageName}
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
func (c *Client) UpdateMultiCloudImageG(p *Params) error {
	multiCloudImageDescription := (*p)["multiCloudImageDescription"].(string)
	multiCloudImageName := (*p)["multiCloudImageName"].(string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	id := (*p)["id"].(string)
	return c.UpdateMultiCloudImage(multiCloudImageDescription, multiCloudImageName, serverTemplateId, id)
}

// == NetworkGateway ==

// A NetworkGateway is an interface that allows traffic to be routed between
// networks.
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
func (c *Client) CreateNetworkGateway(networkGatewayCloudHref string, networkGatewayDescription string, networkGatewayName string, networkGatewayType string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"network_gateway[cloud_href]":  networkGatewayCloudHref,
		"network_gateway[description]": networkGatewayDescription,
		"network_gateway[name]":        networkGatewayName,
		"network_gateway[type]":        networkGatewayType}
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
func (c *Client) CreateNetworkGatewayG(p *Params) (Href, error) {
	networkGatewayCloudHref := (*p)["networkGatewayCloudHref"].(string)
	networkGatewayDescription := (*p)["networkGatewayDescription"].(string)
	networkGatewayName := (*p)["networkGatewayName"].(string)
	networkGatewayType := (*p)["networkGatewayType"].(string)
	return c.CreateNetworkGateway(networkGatewayCloudHref, networkGatewayDescription, networkGatewayName, networkGatewayType)
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
func (c *Client) DestroyNetworkGatewayG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyNetworkGateway(id)
}

// GET api/network_gateways(.:format)?
// Lists the NetworkGateways available to this account.
func (c *Client) IndexNetworkGateways(filter []string) ([]NetworkGateway, error) {
	var res []NetworkGateway
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_gateways", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexNetworkGatewaysG(p *Params) ([]NetworkGateway, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexNetworkGateways(filter)
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
func (c *Client) ShowNetworkGatewayG(p *Params) (*NetworkGateway, error) {
	id := (*p)["id"].(string)
	return c.ShowNetworkGateway(id)
}

// PUT api/network_gateways/:id(.:format)?
// Update an existing NetworkGateway.
func (c *Client) UpdateNetworkGateway(networkGatewayDescription string, networkGatewayName string, networkGatewayNetworkHref string, id string) error {
	payload := map[string]interface{}{
		"network_gateway[description]":  networkGatewayDescription,
		"network_gateway[name]":         networkGatewayName,
		"network_gateway[network_href]": networkGatewayNetworkHref}
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
func (c *Client) UpdateNetworkGatewayG(p *Params) error {
	networkGatewayDescription := (*p)["networkGatewayDescription"].(string)
	networkGatewayName := (*p)["networkGatewayName"].(string)
	networkGatewayNetworkHref := (*p)["networkGatewayNetworkHref"].(string)
	id := (*p)["id"].(string)
	return c.UpdateNetworkGateway(networkGatewayDescription, networkGatewayName, networkGatewayNetworkHref, id)
}

// == NetworkOptionGroupAttachment ==

// Resource for attaching NetworkOptionGroups to Networks.  A single
// NetworkOptionGroup can be attached to many Networks. A Network/Subnet can
// have many NetworkOptionGroups attached, as long as the NetworkOptionGroups
// each have different types.  This resource describes the attachment details
// between a particular NetworkOptionGroup and Network.  Amazon currently only
// supports attaching NetworkOptionGroups to Networks. Other clouds in the
// future may support attaching to Subnets.
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
func (c *Client) CreateNetworkOptionGroupAttachment(networkOptionGroupAttachmentNetworkHref string, networkOptionGroupAttachmentNetworkOptionGroupHref string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"network_option_group_attachment[network_href]":              networkOptionGroupAttachmentNetworkHref,
		"network_option_group_attachment[network_option_group_href]": networkOptionGroupAttachmentNetworkOptionGroupHref}
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
func (c *Client) CreateNetworkOptionGroupAttachmentG(p *Params) (Href, error) {
	networkOptionGroupAttachmentNetworkHref := (*p)["networkOptionGroupAttachmentNetworkHref"].(string)
	networkOptionGroupAttachmentNetworkOptionGroupHref := (*p)["networkOptionGroupAttachmentNetworkOptionGroupHref"].(string)
	return c.CreateNetworkOptionGroupAttachment(networkOptionGroupAttachmentNetworkHref, networkOptionGroupAttachmentNetworkOptionGroupHref)
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
func (c *Client) DestroyNetworkOptionGroupAttachmentG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyNetworkOptionGroupAttachment(id)
}

// GET api/network_option_group_attachments(.:format)?
// List NetworkOptionGroupAttachments in this account.
func (c *Client) IndexNetworkOptionGroupAttachments(filter []string, view string) ([]NetworkOptionGroupAttachment, error) {
	var res []NetworkOptionGroupAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_option_group_attachments", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexNetworkOptionGroupAttachmentsG(p *Params) ([]NetworkOptionGroupAttachment, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexNetworkOptionGroupAttachments(filter, view)
}

// GET api/network_option_group_attachments/:id(.:format)?
// Show information about a single NetworkOptionGroupAttachment.
func (c *Client) ShowNetworkOptionGroupAttachment(view string, id string) (*NetworkOptionGroupAttachment, error) {
	var res *NetworkOptionGroupAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_option_group_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowNetworkOptionGroupAttachmentG(p *Params) (*NetworkOptionGroupAttachment, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowNetworkOptionGroupAttachment(view, id)
}

// PUT api/network_option_group_attachments/:id(.:format)?
// Update an existing NetworkOptionGroupAttachment.
func (c *Client) UpdateNetworkOptionGroupAttachment(networkOptionGroupAttachmentNetworkOptionGroupHref string, id string) error {
	payload := map[string]interface{}{
		"network_option_group_attachment[network_option_group_href]": networkOptionGroupAttachmentNetworkOptionGroupHref}
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
func (c *Client) UpdateNetworkOptionGroupAttachmentG(p *Params) error {
	networkOptionGroupAttachmentNetworkOptionGroupHref := (*p)["networkOptionGroupAttachmentNetworkOptionGroupHref"].(string)
	id := (*p)["id"].(string)
	return c.UpdateNetworkOptionGroupAttachment(networkOptionGroupAttachmentNetworkOptionGroupHref, id)
}

// == NetworkOptionGroup ==

// A key/value pair hash containing options for configuring a Network.  The
// key/value pairs are stored in the "options" parameter.  Keys correspond
// to the type of option to set, and values correspond to the value of the
// particular option being set.  Option keys that are supported vary depending
// on cloud -- please consult your particular cloud's documentation for
// available option keys.
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
func (c *Client) CreateNetworkOptionGroup(networkOptionGroupCloudHref string, networkOptionGroupDescription string, networkOptionGroupName string, networkOptionGroupOptions string, networkOptionGroupType string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"network_option_group[cloud_href]":  networkOptionGroupCloudHref,
		"network_option_group[description]": networkOptionGroupDescription,
		"network_option_group[name]":        networkOptionGroupName,
		"network_option_group[options][*]":  networkOptionGroupOptions,
		"network_option_group[type]":        networkOptionGroupType}
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
func (c *Client) CreateNetworkOptionGroupG(p *Params) (Href, error) {
	networkOptionGroupCloudHref := (*p)["networkOptionGroupCloudHref"].(string)
	networkOptionGroupDescription := (*p)["networkOptionGroupDescription"].(string)
	networkOptionGroupName := (*p)["networkOptionGroupName"].(string)
	networkOptionGroupOptions := (*p)["networkOptionGroupOptions"].(string)
	networkOptionGroupType := (*p)["networkOptionGroupType"].(string)
	return c.CreateNetworkOptionGroup(networkOptionGroupCloudHref, networkOptionGroupDescription, networkOptionGroupName, networkOptionGroupOptions, networkOptionGroupType)
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
func (c *Client) DestroyNetworkOptionGroupG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyNetworkOptionGroup(id)
}

// GET api/network_option_groups(.:format)?
// List NetworkOptionGroups available in this account.
func (c *Client) IndexNetworkOptionGroups(filter []string) ([]NetworkOptionGroup, error) {
	var res []NetworkOptionGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/network_option_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexNetworkOptionGroupsG(p *Params) ([]NetworkOptionGroup, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexNetworkOptionGroups(filter)
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
func (c *Client) ShowNetworkOptionGroupG(p *Params) (*NetworkOptionGroup, error) {
	id := (*p)["id"].(string)
	return c.ShowNetworkOptionGroup(id)
}

// PUT api/network_option_groups/:id(.:format)?
// Update an existing NetworkOptionGroup.
func (c *Client) UpdateNetworkOptionGroup(networkOptionGroupName string, networkOptionGroupOptions string, networkOptionGroupDescription string, id string) error {
	payload := map[string]interface{}{
		"network_option_group[name]":        networkOptionGroupName,
		"network_option_group[options][*]":  networkOptionGroupOptions,
		"network_option_group[description]": networkOptionGroupDescription}
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
func (c *Client) UpdateNetworkOptionGroupG(p *Params) error {
	networkOptionGroupName := (*p)["networkOptionGroupName"].(string)
	networkOptionGroupOptions := (*p)["networkOptionGroupOptions"].(string)
	networkOptionGroupDescription := (*p)["networkOptionGroupDescription"].(string)
	id := (*p)["id"].(string)
	return c.UpdateNetworkOptionGroup(networkOptionGroupName, networkOptionGroupOptions, networkOptionGroupDescription, id)
}

// == Network ==

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
func (c *Client) CreateNetwork(networkCidrBlock string, networkCloudHref string, networkDescription string, networkInstanceTenancy string, networkName string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"network[cidr_block]":       networkCidrBlock,
		"network[cloud_href]":       networkCloudHref,
		"network[description]":      networkDescription,
		"network[instance_tenancy]": networkInstanceTenancy,
		"network[name]":             networkName}
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
func (c *Client) CreateNetworkG(p *Params) (Href, error) {
	networkCidrBlock := (*p)["networkCidrBlock"].(string)
	networkCloudHref := (*p)["networkCloudHref"].(string)
	networkDescription := (*p)["networkDescription"].(string)
	networkInstanceTenancy := (*p)["networkInstanceTenancy"].(string)
	networkName := (*p)["networkName"].(string)
	return c.CreateNetwork(networkCidrBlock, networkCloudHref, networkDescription, networkInstanceTenancy, networkName)
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
func (c *Client) DestroyNetworkG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyNetwork(id)
}

// GET api/networks(.:format)?
// Lists networks in this account.
func (c *Client) IndexNetworks(filter []string) ([]Network, error) {
	var res []Network
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/networks", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexNetworksG(p *Params) ([]Network, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexNetworks(filter)
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
func (c *Client) ShowNetworkG(p *Params) (*Network, error) {
	id := (*p)["id"].(string)
	return c.ShowNetwork(id)
}

// PUT api/networks/:id(.:format)?
// Updates the given network.
func (c *Client) UpdateNetwork(networkName string, networkRouteTableHref string, networkDescription string, id string) error {
	payload := map[string]interface{}{
		"network[name]":             networkName,
		"network[route_table_href]": networkRouteTableHref,
		"network[description]":      networkDescription}
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
func (c *Client) UpdateNetworkG(p *Params) error {
	networkName := (*p)["networkName"].(string)
	networkRouteTableHref := (*p)["networkRouteTableHref"].(string)
	networkDescription := (*p)["networkDescription"].(string)
	id := (*p)["id"].(string)
	return c.UpdateNetwork(networkName, networkRouteTableHref, networkDescription, id)
}

// == Oauth2 ==

// Note that all API calls irrespective of the resource it is acting on,
// should pass a header "X-Api-Version" with the value "1.5".   This is
// an OAuth 2.0 token endpoint that can be used to perform token-refresh
// operations and obtain a bearer access token, which can be used in lieu of
// a session cookie when interacting with API resources.  This is not an API
// resource; in order to maintain compatibility with OAuth 2.0, it does not
// conform to the conventions established by other RightScale API resources.
// However, an API-version header is still required when interacting with the
// OAuth endpoint.  OAuth 2.0 endpoints always use the POST verb, accept a
// www-urlencoded request body (similarly to a browser form submission) and
// the OAuth action is indicated by the "grant_type" parameter. This endpoint
// supports the following OAuth 2.0 operations:   refresh_token - for end-user
// login using a previously-negotiated OAuth grant client_credentials - for
// instance login using API credentials transmitted via user-data   RightScale's
// OAuth implementation has two proprietary aspects that you should be aware
// of:   clients MUST transmit an X-Api-Version header with every OAuth request
// clients MAY transmit an account_id parameter as part of their POST form data
//   If you choose to post an account_id, then the API may respond with a 301
// redirect if your account is hosted in another RightScale cluster. If you omit
// this parameter and your account is hosted elsewhere, then you will simply
// receive a 400 Bad Request (because your grant is not known to this cluster).
//  For more information on how to use OAuth 2.0 with RightScale, refer to the
// following: http://support.rightscale.com/12-Guides/03-RightScale_API/OAuth
// http://tools.ietf.org/html/draft-ietf-oauth-v2-23
type Oauth2 struct {
}

// POST api/oauth2/
// Perform an OAuth 2.0 token_refresh operation to obtain an access token
// that can be used in lieu of an API session cookie. (In other words, creates
// a session using OAuth 2.0).  Note that an API-Version header is required
// with your request, and that the server may respond with a 301 Moved
// Permanently if you include an account_id parameter and your account is
// hosted in another RightScale cluster.  The request parameters and response
// format are all as per the OAuth 2.0 Internet Draft standard v23. In brief:
//   Successful responses include an access token, an expires-in timestamp,
// and a token type The token type is always "bearer" To use a bearer token,
// include header "Authorization: Bearer " with your API requests The client
// must refresh the access token before it expires    # Example Request
// using Curl (with prettified response): curl -i -H X-API-Version:1.5 -x
// POST https://my.rightscale.com/api/oauth2 -d "grant_type=refresh_token"
// -d "refresh_token=abcd1234deadbeef"  {   "access_token": "xyzzy",
// "expires_in":   3600,   "token_type":   "bearer" }
func (c *Client) CreateOauth2(accountId int, clientId string, clientSecret string, grantType string, rsVersion int, refreshToken string, rightLinkVersion string) (map[string]interface{}, error) {
	var res map[string]interface{}
	payload := map[string]interface{}{
		"account_id":         accountId,
		"client_id":          clientId,
		"client_secret":      clientSecret,
		"grant_type":         grantType,
		"r_s_version":        rsVersion,
		"refresh_token":      refreshToken,
		"right_link_version": rightLinkVersion}
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
func (c *Client) CreateOauth2G(p *Params) (map[string]interface{}, error) {
	accountId := (*p)["accountId"].(int)
	clientId := (*p)["clientId"].(string)
	clientSecret := (*p)["clientSecret"].(string)
	grantType := (*p)["grantType"].(string)
	rsVersion := (*p)["rsVersion"].(int)
	refreshToken := (*p)["refreshToken"].(string)
	rightLinkVersion := (*p)["rightLinkVersion"].(string)
	return c.CreateOauth2(accountId, clientId, clientSecret, grantType, rsVersion, refreshToken, rightLinkVersion)
}

// == Permission ==

type Permission struct {
	Actions   []string            `json:"actions,omitempty"`
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	RoleTitle string              `json:"role_title,omitempty"`
}

// POST api/permissions(.:format)?
// Create a permission, thereby granting some user a particular role with
// respect to the current account.  The 'observer' role has a special status; it
// must be granted before a user is eligible for any other permission in a given
// account.   When provisioning users, always create the observer permission
// FIRST;  creating any other permission before it will result in an error.  For
// more information about the roles available and the privileges they confer,
// please refer to the following page of the RightScale support portal:
// http://support.rightscale.com/15-References/Lists/ListofUser_Roles
func (c *Client) CreatePermission(permissionUserHref string, permissionRoleTitle string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"permission[user_href]":  permissionUserHref,
		"permission[role_title]": permissionRoleTitle}
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
func (c *Client) CreatePermissionG(p *Params) (Href, error) {
	permissionUserHref := (*p)["permissionUserHref"].(string)
	permissionRoleTitle := (*p)["permissionRoleTitle"].(string)
	return c.CreatePermission(permissionUserHref, permissionRoleTitle)
}

// DELETE api/permissions/:id(.:format)?
// Destroy a permission, thereby revoking a user's role with respect to the
// current account.  The 'observer' role has a special status; it cannot
// be revoked if a user has any other roles, because other roles become
// useless without being able to read data pertaining to the account.  When
// deprovisioning user, always destroy the observer permission LAST; destroying
// it while the user has other permissions will result in an error.
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
func (c *Client) DestroyPermissionG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyPermission(id)
}

// GET api/permissions(.:format)?
// List all permissions for all users of the current acount.
func (c *Client) IndexPermissions(filter []string) ([]Permission, error) {
	var res []Permission
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/permissions", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexPermissionsG(p *Params) ([]Permission, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexPermissions(filter)
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
func (c *Client) ShowPermissionG(p *Params) (*Permission, error) {
	id := (*p)["id"].(string)
	return c.ShowPermission(id)
}

// == PlacementGroup ==

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
func (c *Client) CreatePlacementGroup(placementGroupName string, placementGroupCloudHref string, placementGroupDescription string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"placement_group[name]":        placementGroupName,
		"placement_group[cloud_href]":  placementGroupCloudHref,
		"placement_group[description]": placementGroupDescription}
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
func (c *Client) CreatePlacementGroupG(p *Params) (Href, error) {
	placementGroupName := (*p)["placementGroupName"].(string)
	placementGroupCloudHref := (*p)["placementGroupCloudHref"].(string)
	placementGroupDescription := (*p)["placementGroupDescription"].(string)
	return c.CreatePlacementGroup(placementGroupName, placementGroupCloudHref, placementGroupDescription)
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
func (c *Client) DestroyPlacementGroupG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyPlacementGroup(id)
}

// GET api/placement_groups(.:format)?
// Lists all PlacementGroups in an account.
func (c *Client) IndexPlacementGroups(filter []string, view string) ([]PlacementGroup, error) {
	var res []PlacementGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/placement_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexPlacementGroupsG(p *Params) ([]PlacementGroup, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexPlacementGroups(filter, view)
}

// GET api/placement_groups/:id(.:format)?
// Shows information about a single PlacementGroup.
func (c *Client) ShowPlacementGroup(view string, id string) (*PlacementGroup, error) {
	var res *PlacementGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/placement_groups/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowPlacementGroupG(p *Params) (*PlacementGroup, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowPlacementGroup(view, id)
}

// == Preference ==

// A Preference is a user and account-specific setting. Preferences are used in
// many part of the RightScale platform and can be used for custom purposes if
// desired.
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
func (c *Client) DestroyPreferenceG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyPreference(id)
}

// GET api/preferences(.:format)?
// Lists all preferences.
func (c *Client) IndexPreferences(filter []string) ([]Preference, error) {
	var res []Preference
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/preferences", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexPreferencesG(p *Params) ([]Preference, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexPreferences(filter)
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
func (c *Client) ShowPreferenceG(p *Params) (*Preference, error) {
	id := (*p)["id"].(string)
	return c.ShowPreference(id)
}

// PUT api/preferences/:id(.:format)?
// If 'id' is known, updates preference with given contents. Otherwise, creates
// new preference. Note: If create, will return '201 Created' and the location
// of the new preference.
func (c *Client) UpdatePreference(preferenceContents string, id string) error {
	payload := map[string]interface{}{
		"preference[contents]": preferenceContents}
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
func (c *Client) UpdatePreferenceG(p *Params) error {
	preferenceContents := (*p)["preferenceContents"].(string)
	id := (*p)["id"].(string)
	return c.UpdatePreference(preferenceContents, id)
}

// == PublicationLineage ==

// A Publication Lineage contains lineage information for a Publication in the
// MultiCloudMarketplace. It is shared among all revisions of a Publication
// within the marketplace. Publication Lineages are different than lineages that
// exist within an account.
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
// Show information about a single publication lineage. Only non-HEAD revisions
// are possible.
func (c *Client) ShowPublicationLineage(view string, id string) (*PublicationLineage, error) {
	var res *PublicationLineage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/publication_lineages/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowPublicationLineageG(p *Params) (*PublicationLineage, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowPublicationLineage(view, id)
}

// == Publication ==

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
// Imports the given publication and its subordinates to this account. Only
// non-HEAD revisions that are shared with the account can be imported.
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
func (c *Client) ImportPublicationG(p *Params) error {
	id := (*p)["id"].(string)
	return c.ImportPublication(id)
}

// GET api/publications(.:format)?
// Lists the publications available to this account. Only non-HEAD revisions are
// possible.
func (c *Client) IndexPublications(filter []string, view string) ([]Publication, error) {
	var res []Publication
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/publications", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexPublicationsG(p *Params) ([]Publication, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexPublications(filter, view)
}

// GET api/publications/:id(.:format)?
// Show information about a single publication. Only non-HEAD revisions are
// possible.
func (c *Client) ShowPublication(view string, id string) (*Publication, error) {
	var res *Publication
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/publications/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowPublicationG(p *Params) (*Publication, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowPublication(view, id)
}

// == RecurringVolumeAttachment ==

// A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a
// Server/ServerArray the next time an instance is launched.
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
func (c *Client) CreateRecurringVolumeAttachment(recurringVolumeAttachmentDevice string, recurringVolumeAttachmentRunnableHref string, recurringVolumeAttachmentStorageHref string, cloudId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"recurring_volume_attachment[device]":        recurringVolumeAttachmentDevice,
		"recurring_volume_attachment[runnable_href]": recurringVolumeAttachmentRunnableHref,
		"recurring_volume_attachment[storage_href]":  recurringVolumeAttachmentStorageHref}
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
func (c *Client) CreateRecurringVolumeAttachmentG(p *Params) (Href, error) {
	recurringVolumeAttachmentDevice := (*p)["recurringVolumeAttachmentDevice"].(string)
	recurringVolumeAttachmentRunnableHref := (*p)["recurringVolumeAttachmentRunnableHref"].(string)
	recurringVolumeAttachmentStorageHref := (*p)["recurringVolumeAttachmentStorageHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.CreateRecurringVolumeAttachment(recurringVolumeAttachmentDevice, recurringVolumeAttachmentRunnableHref, recurringVolumeAttachmentStorageHref, cloudId)
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
func (c *Client) DestroyRecurringVolumeAttachmentG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyRecurringVolumeAttachment(cloudId, id)
}

// GET api/clouds/:cloud_id/recurring_volume_attachments(.:format)?
// Lists all recurring volume attachments.
func (c *Client) IndexRecurringVolumeAttachments(filter []string, view string, cloudId string) ([]RecurringVolumeAttachment, error) {
	var res []RecurringVolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/recurring_volume_attachments", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexRecurringVolumeAttachmentsG(p *Params) ([]RecurringVolumeAttachment, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexRecurringVolumeAttachments(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/recurring_volume_attachments/:id(.:format)?
// Displays information about a single recurring volume attachment.
func (c *Client) ShowRecurringVolumeAttachment(view string, cloudId string, id string) (*RecurringVolumeAttachment, error) {
	var res *RecurringVolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/recurring_volume_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowRecurringVolumeAttachmentG(p *Params) (*RecurringVolumeAttachment, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowRecurringVolumeAttachment(view, cloudId, id)
}

// == Repository ==

// A Repository is a location from which you can download and import design
// objects such as Chef cookbooks. Using this resource you can add and modify
// repository information and import assets discovered in the repository.
// RightScale currently supports the following types of repositores: git, svn,
// and URLs of compressed files (tar, tgz, gzip).
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
// Performs a Cookbook import, which allows you to use the specified cookbooks
// in your design objects.
func (c *Client) CookbookImportRepository(follow string, namespace string, repositoryCommitReference string, withDependencies string, assetHrefs []string, id string) error {
	payload := map[string]interface{}{
		"follow":                      follow,
		"namespace":                   namespace,
		"repository_commit_reference": repositoryCommitReference,
		"with_dependencies":           withDependencies,
		"asset_hrefs":                 assetHrefs}
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
func (c *Client) CookbookImportRepositoryG(p *Params) error {
	follow := (*p)["follow"].(string)
	namespace := (*p)["namespace"].(string)
	repositoryCommitReference := (*p)["repositoryCommitReference"].(string)
	withDependencies := (*p)["withDependencies"].(string)
	assetHrefs := (*p)["assetHrefs"].([]string)
	id := (*p)["id"].(string)
	return c.CookbookImportRepository(follow, namespace, repositoryCommitReference, withDependencies, assetHrefs, id)
}

// POST api/repositories/:id/cookbook_import_preview(.:format)?
// Retrieves a preview of the effects of a Cookbook import.  NOTE: This action
// is for RightScale internal use only. The response is free-form JSON with
// no associated mediatype.  DO NOT USE, THIS ACTION IS SUBJECT TO CHANGE AT
// ANYTIME.
func (c *Client) CookbookImportPreviewRepository(assetHrefs []string, namespace string, id string) ([]map[string]string, error) {
	var res []map[string]string
	payload := map[string]interface{}{
		"asset_hrefs": assetHrefs,
		"namespace":   namespace}
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
func (c *Client) CookbookImportPreviewRepositoryG(p *Params) ([]map[string]string, error) {
	assetHrefs := (*p)["assetHrefs"].([]string)
	namespace := (*p)["namespace"].(string)
	id := (*p)["id"].(string)
	return c.CookbookImportPreviewRepository(assetHrefs, namespace, id)
}

// POST api/repositories(.:format)?
// Creates a Repository.  The following types of inputs are supported for the
// credential fields:    Type Format Example(s)   Text string text:&lt;value&gt;
// text:-----BEGIN RSA PRIVATE KEY-----text:secret   Credential value
// cred:&lt;value&gt; cred:my ssh keycred:svn_1_password
func (c *Client) CreateRepository(repositoryCommitReference string, repositoryCredentialsPassword string, repositoryCredentialsSshKey string, repositoryName string, repositorySourceType string, repositoryAssetPathsCookbooks []string, repositoryCredentialsUsername string, repositoryDescription string, repositorySource string, repositoryAutoImport string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"repository[commit_reference]":       repositoryCommitReference,
		"repository[credentials][password]":  repositoryCredentialsPassword,
		"repository[credentials][ssh_key]":   repositoryCredentialsSshKey,
		"repository[name]":                   repositoryName,
		"repository[source_type]":            repositorySourceType,
		"repository[asset_paths][cookbooks]": repositoryAssetPathsCookbooks,
		"repository[credentials][username]":  repositoryCredentialsUsername,
		"repository[description]":            repositoryDescription,
		"repository[source]":                 repositorySource,
		"repository[auto_import]":            repositoryAutoImport}
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
func (c *Client) CreateRepositoryG(p *Params) (Href, error) {
	repositoryCommitReference := (*p)["repositoryCommitReference"].(string)
	repositoryCredentialsPassword := (*p)["repositoryCredentialsPassword"].(string)
	repositoryCredentialsSshKey := (*p)["repositoryCredentialsSshKey"].(string)
	repositoryName := (*p)["repositoryName"].(string)
	repositorySourceType := (*p)["repositorySourceType"].(string)
	repositoryAssetPathsCookbooks := (*p)["repositoryAssetPathsCookbooks"].([]string)
	repositoryCredentialsUsername := (*p)["repositoryCredentialsUsername"].(string)
	repositoryDescription := (*p)["repositoryDescription"].(string)
	repositorySource := (*p)["repositorySource"].(string)
	repositoryAutoImport := (*p)["repositoryAutoImport"].(string)
	return c.CreateRepository(repositoryCommitReference, repositoryCredentialsPassword, repositoryCredentialsSshKey, repositoryName, repositorySourceType, repositoryAssetPathsCookbooks, repositoryCredentialsUsername, repositoryDescription, repositorySource, repositoryAutoImport)
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
func (c *Client) DestroyRepositoryG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyRepository(id)
}

// GET api/repositories(.:format)?
// Lists all Repositories for this Account.
func (c *Client) IndexRepositories(filter []string, view string) ([]Repository, error) {
	var res []Repository
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexRepositoriesG(p *Params) ([]Repository, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexRepositories(filter, view)
}

// POST api/repositories/:id/refetch(.:format)?
// Refetches all RepositoryAssets associated with the Repository. Note that a
// refetch simply updates RightScale's view of the contents of the repository.
// You must perform an import to use the assets in your design objects (or use
// the auto import parameter).
func (c *Client) RefetchRepository(autoImport string, id string) error {
	payload := map[string]interface{}{
		"auto_import": autoImport}
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
func (c *Client) RefetchRepositoryG(p *Params) error {
	autoImport := (*p)["autoImport"].(string)
	id := (*p)["id"].(string)
	return c.RefetchRepository(autoImport, id)
}

// POST api/repositories/resolve(.:format)?
// Show a list of repositories that have imported cookbooks with the given
// names.  This operation returns a list of repositories that would later
// satisfy a call to the swap_repository action on a ServerTemplate.
func (c *Client) ResolveRepository(importedCookbookName []string) ([]Repository, error) {
	var res []Repository
	payload := map[string]interface{}{
		"imported_cookbook_name": importedCookbookName}
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
func (c *Client) ResolveRepositoryG(p *Params) ([]Repository, error) {
	importedCookbookName := (*p)["importedCookbookName"].([]string)
	return c.ResolveRepository(importedCookbookName)
}

// GET api/repositories/:id(.:format)?
// Shows a specified Repository.
func (c *Client) ShowRepository(view string, id string) (*Repository, error) {
	var res *Repository
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowRepositoryG(p *Params) (*Repository, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowRepository(view, id)
}

// PUT api/repositories/:id(.:format)?
// Updates a specified Repository.  The following types of inputs are
// supported for the credential fields:    Type Format Example(s)   Text
// string text:&lt;value&gt; text:-----BEGIN RSA PRIVATE KEY-----text:secret
// Credential value cred:&lt;value&gt; cred:my ssh keycred:svn_1_password
func (c *Client) UpdateRepository(repositorySource string, repositoryCredentialsPassword string, repositoryAssetPathsCookbooks []string, repositoryCommitReference string, repositoryCredentialsSshKey string, repositoryCredentialsUsername string, repositoryDescription string, repositoryName string, repositorySourceType string, id string) error {
	payload := map[string]interface{}{
		"repository[source]":                 repositorySource,
		"repository[credentials][password]":  repositoryCredentialsPassword,
		"repository[asset_paths][cookbooks]": repositoryAssetPathsCookbooks,
		"repository[commit_reference]":       repositoryCommitReference,
		"repository[credentials][ssh_key]":   repositoryCredentialsSshKey,
		"repository[credentials][username]":  repositoryCredentialsUsername,
		"repository[description]":            repositoryDescription,
		"repository[name]":                   repositoryName,
		"repository[source_type]":            repositorySourceType}
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
func (c *Client) UpdateRepositoryG(p *Params) error {
	repositorySource := (*p)["repositorySource"].(string)
	repositoryCredentialsPassword := (*p)["repositoryCredentialsPassword"].(string)
	repositoryAssetPathsCookbooks := (*p)["repositoryAssetPathsCookbooks"].([]string)
	repositoryCommitReference := (*p)["repositoryCommitReference"].(string)
	repositoryCredentialsSshKey := (*p)["repositoryCredentialsSshKey"].(string)
	repositoryCredentialsUsername := (*p)["repositoryCredentialsUsername"].(string)
	repositoryDescription := (*p)["repositoryDescription"].(string)
	repositoryName := (*p)["repositoryName"].(string)
	repositorySourceType := (*p)["repositorySourceType"].(string)
	id := (*p)["id"].(string)
	return c.UpdateRepository(repositorySource, repositoryCredentialsPassword, repositoryAssetPathsCookbooks, repositoryCommitReference, repositoryCredentialsSshKey, repositoryCredentialsUsername, repositoryDescription, repositoryName, repositorySourceType, id)
}

// == RepositoryAsset ==

// A RepositoryAsset represents an item discovered in a Repository. These assets
// represent only a view of the Repository the last time it was scraped. In
// order to use these assets, you must import them into your account.
type RepositoryAsset struct {
	Actions     []string            `json:"actions,omitempty"`
	Description string              `json:"description,omitempty"`
	Id          string              `json:"id,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Version     string              `json:"version,omitempty"`
}

// GET api/repositories/:repository_id/repository_assets(.:format)?
// List a repository's current assets.  Repository assests are the cookbook
// details that were scraped from a given repository.
func (c *Client) IndexRepositoryAssets(view string, repositoryId string) ([]RepositoryAsset, error) {
	var res []RepositoryAsset
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories/"+repositoryId+"/repository_assets", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexRepositoryAssetsG(p *Params) ([]RepositoryAsset, error) {
	view := (*p)["view"].(string)
	repositoryId := (*p)["repositoryId"].(string)
	return c.IndexRepositoryAssets(view, repositoryId)
}

// GET api/repositories/:repository_id/repository_assets/:id(.:format)?
// Show information about a single asset.  A repository assest are the cookbook
// details that were scraped from a repository.
func (c *Client) ShowRepositoryAsset(view string, repositoryId string, id string) (*RepositoryAsset, error) {
	var res *RepositoryAsset
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/repositories/"+repositoryId+"/repository_assets/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowRepositoryAssetG(p *Params) (*RepositoryAsset, error) {
	view := (*p)["view"].(string)
	repositoryId := (*p)["repositoryId"].(string)
	id := (*p)["id"].(string)
	return c.ShowRepositoryAsset(view, repositoryId, id)
}

// == RightScript ==

// A RightScript is an executable piece of code that can be run on a server
// during the boot, operational, or decommission phases.  All revisions of a
// RightScript belong to a RightScript lineage that is exposed by the "lineage"
// attribute (NOTE: This attribute is merely a string to locate all revisions of
// a RightScript and NOT a working URL).
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
// Commits the given RightScript. Only HEAD revisions (revision 0) can be
// committed.
func (c *Client) CommitRightScript(rightScriptCommitMessage string, id string) error {
	payload := map[string]interface{}{
		"right_script[commit_message]": rightScriptCommitMessage}
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
func (c *Client) CommitRightScriptG(p *Params) error {
	rightScriptCommitMessage := (*p)["rightScriptCommitMessage"].(string)
	id := (*p)["id"].(string)
	return c.CommitRightScript(rightScriptCommitMessage, id)
}

// GET api/right_scripts(.:format)?
// Lists RightScripts.
func (c *Client) IndexRightScripts(filter []string, view string, latestOnly string) ([]RightScript, error) {
	var res []RightScript
	payload := map[string]interface{}{
		"latest_only": latestOnly}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/right_scripts", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexRightScriptsG(p *Params) ([]RightScript, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	latestOnly := (*p)["latestOnly"].(string)
	return c.IndexRightScripts(filter, view, latestOnly)
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
func (c *Client) ShowRightScriptG(p *Params) (*RightScript, error) {
	id := (*p)["id"].(string)
	return c.ShowRightScript(id)
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
func (c *Client) ShowSourceRightScriptG(p *Params) error {
	id := (*p)["id"].(string)
	return c.ShowSourceRightScript(id)
}

// PUT api/right_scripts/:id(.:format)?
// Updates RightScript name/description
func (c *Client) UpdateRightScript(rightScriptDescription string, rightScriptName string, id string) error {
	payload := map[string]interface{}{
		"right_script[description]": rightScriptDescription,
		"right_script[name]":        rightScriptName}
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
func (c *Client) UpdateRightScriptG(p *Params) error {
	rightScriptDescription := (*p)["rightScriptDescription"].(string)
	rightScriptName := (*p)["rightScriptName"].(string)
	id := (*p)["id"].(string)
	return c.UpdateRightScript(rightScriptDescription, rightScriptName, id)
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
func (c *Client) UpdateSourceRightScriptG(p *Params) error {
	id := (*p)["id"].(string)
	return c.UpdateSourceRightScript(id)
}

// == RouteTable ==

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
func (c *Client) CreateRouteTable(routeTableCloudHref string, routeTableDescription string, routeTableName string, routeTableNetworkHref string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"route_table[cloud_href]":   routeTableCloudHref,
		"route_table[description]":  routeTableDescription,
		"route_table[name]":         routeTableName,
		"route_table[network_href]": routeTableNetworkHref}
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
func (c *Client) CreateRouteTableG(p *Params) (Href, error) {
	routeTableCloudHref := (*p)["routeTableCloudHref"].(string)
	routeTableDescription := (*p)["routeTableDescription"].(string)
	routeTableName := (*p)["routeTableName"].(string)
	routeTableNetworkHref := (*p)["routeTableNetworkHref"].(string)
	return c.CreateRouteTable(routeTableCloudHref, routeTableDescription, routeTableName, routeTableNetworkHref)
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
func (c *Client) DestroyRouteTableG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyRouteTable(id)
}

// GET api/route_tables(.:format)?
// List RouteTables available in this account.
func (c *Client) IndexRouteTables(filter []string, view string) ([]RouteTable, error) {
	var res []RouteTable
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/route_tables", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexRouteTablesG(p *Params) ([]RouteTable, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexRouteTables(filter, view)
}

// GET api/route_tables/:id(.:format)?
// Show information about a single RouteTable.
func (c *Client) ShowRouteTable(view string, id string) (*RouteTable, error) {
	var res *RouteTable
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/route_tables/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowRouteTableG(p *Params) (*RouteTable, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowRouteTable(view, id)
}

// PUT api/route_tables/:id(.:format)?
// Update an existing RouteTable.
func (c *Client) UpdateRouteTable(routeTableDescription string, routeTableName string, id string) error {
	payload := map[string]interface{}{
		"route_table[description]": routeTableDescription,
		"route_table[name]":        routeTableName}
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
func (c *Client) UpdateRouteTableG(p *Params) error {
	routeTableDescription := (*p)["routeTableDescription"].(string)
	routeTableName := (*p)["routeTableName"].(string)
	id := (*p)["id"].(string)
	return c.UpdateRouteTable(routeTableDescription, routeTableName, id)
}

// == Route ==

// A Route defines how networking traffic should be routed from one destination
// to another. See nexthoptype for available endpoint targets.
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
func (c *Client) CreateRoute(routeDescription string, routeDestinationCidrBlock string, routeNextHopHref string, routeNextHopIp string, routeNextHopType string, routeRouteTableHref string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"route[description]":            routeDescription,
		"route[destination_cidr_block]": routeDestinationCidrBlock,
		"route[next_hop_href]":          routeNextHopHref,
		"route[next_hop_ip]":            routeNextHopIp,
		"route[next_hop_type]":          routeNextHopType,
		"route[route_table_href]":       routeRouteTableHref}
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
func (c *Client) CreateRouteG(p *Params) (Href, error) {
	routeDescription := (*p)["routeDescription"].(string)
	routeDestinationCidrBlock := (*p)["routeDestinationCidrBlock"].(string)
	routeNextHopHref := (*p)["routeNextHopHref"].(string)
	routeNextHopIp := (*p)["routeNextHopIp"].(string)
	routeNextHopType := (*p)["routeNextHopType"].(string)
	routeRouteTableHref := (*p)["routeRouteTableHref"].(string)
	return c.CreateRoute(routeDescription, routeDestinationCidrBlock, routeNextHopHref, routeNextHopIp, routeNextHopType, routeRouteTableHref)
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
func (c *Client) DestroyRouteG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyRoute(id)
}

// GET api/routes(.:format)?
// List Routes available in this account.
func (c *Client) IndexRoutes(filter []string) ([]Route, error) {
	var res []Route
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/routes", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexRoutesG(p *Params) ([]Route, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexRoutes(filter)
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
func (c *Client) ShowRouteG(p *Params) (*Route, error) {
	id := (*p)["id"].(string)
	return c.ShowRoute(id)
}

// PUT api/routes/:id(.:format)?
// Update an existing Route.
func (c *Client) UpdateRoute(routeDestinationCidrBlock string, routeNextHopHref string, routeNextHopIp string, routeNextHopType string, routeDescription string, id string) error {
	payload := map[string]interface{}{
		"route[destination_cidr_block]": routeDestinationCidrBlock,
		"route[next_hop_href]":          routeNextHopHref,
		"route[next_hop_ip]":            routeNextHopIp,
		"route[next_hop_type]":          routeNextHopType,
		"route[description]":            routeDescription}
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
func (c *Client) UpdateRouteG(p *Params) error {
	routeDestinationCidrBlock := (*p)["routeDestinationCidrBlock"].(string)
	routeNextHopHref := (*p)["routeNextHopHref"].(string)
	routeNextHopIp := (*p)["routeNextHopIp"].(string)
	routeNextHopType := (*p)["routeNextHopType"].(string)
	routeDescription := (*p)["routeDescription"].(string)
	id := (*p)["id"].(string)
	return c.UpdateRoute(routeDestinationCidrBlock, routeNextHopHref, routeNextHopIp, routeNextHopType, routeDescription, id)
}

// == RunnableBinding ==

// A RunnableBinding represents an item in a runlist of a ServerTemplate.
// These items could be RightScript or Chef recipes, and could be associated
// with any one of the three runlists of a  ServerTemplate (boot, operational,
// decommission).
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
// Bind an executable to the given ServerTemplate.  An executable may be either
// a RightScript or Chef Recipe.  The resource must be editable.
func (c *Client) CreateRunnableBinding(runnableBindingPosition string, runnableBindingRecipe string, runnableBindingRightScriptHref string, runnableBindingSequence string, serverTemplateId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"runnable_binding[position]":          runnableBindingPosition,
		"runnable_binding[recipe]":            runnableBindingRecipe,
		"runnable_binding[right_script_href]": runnableBindingRightScriptHref,
		"runnable_binding[sequence]":          runnableBindingSequence}
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
func (c *Client) CreateRunnableBindingG(p *Params) (Href, error) {
	runnableBindingPosition := (*p)["runnableBindingPosition"].(string)
	runnableBindingRecipe := (*p)["runnableBindingRecipe"].(string)
	runnableBindingRightScriptHref := (*p)["runnableBindingRightScriptHref"].(string)
	runnableBindingSequence := (*p)["runnableBindingSequence"].(string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	return c.CreateRunnableBinding(runnableBindingPosition, runnableBindingRecipe, runnableBindingRightScriptHref, runnableBindingSequence, serverTemplateId)
}

// DELETE api/server_templates/:server_template_id/runnable_bindings/:id(.:format)?
// Unbind an executable from the given resource.  The resource must be editable.
func (c *Client) DestroyRunnableBinding(serverTemplateId string, id string) error {
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
func (c *Client) DestroyRunnableBindingG(p *Params) error {
	serverTemplateId := (*p)["serverTemplateId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyRunnableBinding(serverTemplateId, id)
}

// GET api/server_templates/:server_template_id/runnable_bindings(.:format)?
// Lists the executables bound to the ServerTemplate.  An excutable may be
// either a RightScript or Chef Recipe.
func (c *Client) IndexRunnableBindings(view string, serverTemplateId string) ([]RunnableBinding, error) {
	var res []RunnableBinding
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+serverTemplateId+"/runnable_bindings", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexRunnableBindingsG(p *Params) ([]RunnableBinding, error) {
	view := (*p)["view"].(string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	return c.IndexRunnableBindings(view, serverTemplateId)
}

// PUT api/server_templates/:server_template_id/runnable_bindings/multi_update(.:format)?
// Update attributes for multiple bound executables.  The resource must be
// editable.
func (c *Client) MultiUpdateRunnableBindings(runnableBindingsRecipe string, runnableBindingsRightScriptHref string, runnableBindingsSequence string, runnableBindings []string, runnableBindingsId string, runnableBindingsPosition string, serverTemplateId string) error {
	payload := map[string]interface{}{
		"runnable_bindings[][recipe]":            runnableBindingsRecipe,
		"runnable_bindings[][right_script_href]": runnableBindingsRightScriptHref,
		"runnable_bindings[][sequence]":          runnableBindingsSequence,
		"runnable_bindings":                      runnableBindings,
		"runnable_bindings[][id]":                runnableBindingsId,
		"runnable_bindings[][position]":          runnableBindingsPosition}
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
func (c *Client) MultiUpdateRunnableBindingsG(p *Params) error {
	runnableBindingsRecipe := (*p)["runnableBindingsRecipe"].(string)
	runnableBindingsRightScriptHref := (*p)["runnableBindingsRightScriptHref"].(string)
	runnableBindingsSequence := (*p)["runnableBindingsSequence"].(string)
	runnableBindings := (*p)["runnableBindings"].([]string)
	runnableBindingsId := (*p)["runnableBindingsId"].(string)
	runnableBindingsPosition := (*p)["runnableBindingsPosition"].(string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	return c.MultiUpdateRunnableBindings(runnableBindingsRecipe, runnableBindingsRightScriptHref, runnableBindingsSequence, runnableBindings, runnableBindingsId, runnableBindingsPosition, serverTemplateId)
}

// GET api/server_templates/:server_template_id/runnable_bindings/:id(.:format)?
// Show information about a single executable binding.  An excutable may be
// either a RightScript or Chef Recipe.
func (c *Client) ShowRunnableBinding(view string, serverTemplateId string, id string) (*RunnableBinding, error) {
	var res *RunnableBinding
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+serverTemplateId+"/runnable_bindings/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowRunnableBindingG(p *Params) (*RunnableBinding, error) {
	view := (*p)["view"].(string)
	serverTemplateId := (*p)["serverTemplateId"].(string)
	id := (*p)["id"].(string)
	return c.ShowRunnableBinding(view, serverTemplateId, id)
}

// == SecurityGroupRule ==

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
// Create a security group rule for a security group. The following flavors
// are supported:   group-based TCP/UDP group-based ICMP CIDR-based TCP/UDP
// CIDR-based ICMP
func (c *Client) CreateSecurityGroupRule(securityGroupRuleCidrIps string, securityGroupRuleDirection string, securityGroupRuleProtocol string, securityGroupRuleProtocolDetailsEndPort string, securityGroupRuleProtocolDetailsStartPort string, securityGroupRuleSourceType string, securityGroupRuleGroupName string, securityGroupRuleGroupOwner string, securityGroupRuleProtocolDetailsIcmpCode string, securityGroupRuleProtocolDetailsIcmpType string, securityGroupRuleSecurityGroupHref string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"security_group_rule[cidr_ips]":                     securityGroupRuleCidrIps,
		"security_group_rule[direction]":                    securityGroupRuleDirection,
		"security_group_rule[protocol]":                     securityGroupRuleProtocol,
		"security_group_rule[protocol_details][end_port]":   securityGroupRuleProtocolDetailsEndPort,
		"security_group_rule[protocol_details][start_port]": securityGroupRuleProtocolDetailsStartPort,
		"security_group_rule[source_type]":                  securityGroupRuleSourceType,
		"security_group_rule[group_name]":                   securityGroupRuleGroupName,
		"security_group_rule[group_owner]":                  securityGroupRuleGroupOwner,
		"security_group_rule[protocol_details][icmp_code]":  securityGroupRuleProtocolDetailsIcmpCode,
		"security_group_rule[protocol_details][icmp_type]":  securityGroupRuleProtocolDetailsIcmpType,
		"security_group_rule[security_group_href]":          securityGroupRuleSecurityGroupHref}
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
func (c *Client) CreateSecurityGroupRuleG(p *Params) (Href, error) {
	securityGroupRuleCidrIps := (*p)["securityGroupRuleCidrIps"].(string)
	securityGroupRuleDirection := (*p)["securityGroupRuleDirection"].(string)
	securityGroupRuleProtocol := (*p)["securityGroupRuleProtocol"].(string)
	securityGroupRuleProtocolDetailsEndPort := (*p)["securityGroupRuleProtocolDetailsEndPort"].(string)
	securityGroupRuleProtocolDetailsStartPort := (*p)["securityGroupRuleProtocolDetailsStartPort"].(string)
	securityGroupRuleSourceType := (*p)["securityGroupRuleSourceType"].(string)
	securityGroupRuleGroupName := (*p)["securityGroupRuleGroupName"].(string)
	securityGroupRuleGroupOwner := (*p)["securityGroupRuleGroupOwner"].(string)
	securityGroupRuleProtocolDetailsIcmpCode := (*p)["securityGroupRuleProtocolDetailsIcmpCode"].(string)
	securityGroupRuleProtocolDetailsIcmpType := (*p)["securityGroupRuleProtocolDetailsIcmpType"].(string)
	securityGroupRuleSecurityGroupHref := (*p)["securityGroupRuleSecurityGroupHref"].(string)
	return c.CreateSecurityGroupRule(securityGroupRuleCidrIps, securityGroupRuleDirection, securityGroupRuleProtocol, securityGroupRuleProtocolDetailsEndPort, securityGroupRuleProtocolDetailsStartPort, securityGroupRuleSourceType, securityGroupRuleGroupName, securityGroupRuleGroupOwner, securityGroupRuleProtocolDetailsIcmpCode, securityGroupRuleProtocolDetailsIcmpType, securityGroupRuleSecurityGroupHref)
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
func (c *Client) DestroySecurityGroupRuleG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroySecurityGroupRule(id)
}

// GET api/security_group_rules(.:format)?
// Lists SecurityGroupRules.
func (c *Client) IndexSecurityGroupRules(view string) ([]SecurityGroupRule, error) {
	var res []SecurityGroupRule
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/security_group_rules", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexSecurityGroupRulesG(p *Params) ([]SecurityGroupRule, error) {
	view := (*p)["view"].(string)
	return c.IndexSecurityGroupRules(view)
}

// GET api/security_group_rules/:id(.:format)?
// Displays information about a single SecurityGroupRule.
func (c *Client) ShowSecurityGroupRule(view string, id string) (*SecurityGroupRule, error) {
	var res *SecurityGroupRule
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/security_group_rules/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowSecurityGroupRuleG(p *Params) (*SecurityGroupRule, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowSecurityGroupRule(view, id)
}

// PUT api/security_group_rules/:id(.:format)?

func (c *Client) UpdateSecurityGroupRule(securityGroupRuleDescription string, id string) error {
	payload := map[string]interface{}{
		"security_group_rule[description]": securityGroupRuleDescription}
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
func (c *Client) UpdateSecurityGroupRuleG(p *Params) error {
	securityGroupRuleDescription := (*p)["securityGroupRuleDescription"].(string)
	id := (*p)["id"].(string)
	return c.UpdateSecurityGroupRule(securityGroupRuleDescription, id)
}

// == SecurityGroup ==

// Security Groups represent network security profiles that contain lists of
// firewall rules for different ports and source IP addresses, as well as trust
// relationships amongst different security groups.
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
func (c *Client) CreateSecurityGroup(securityGroupDescription string, securityGroupName string, securityGroupNetworkHref string, cloudId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"security_group[description]":  securityGroupDescription,
		"security_group[name]":         securityGroupName,
		"security_group[network_href]": securityGroupNetworkHref}
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
func (c *Client) CreateSecurityGroupG(p *Params) (Href, error) {
	securityGroupDescription := (*p)["securityGroupDescription"].(string)
	securityGroupName := (*p)["securityGroupName"].(string)
	securityGroupNetworkHref := (*p)["securityGroupNetworkHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.CreateSecurityGroup(securityGroupDescription, securityGroupName, securityGroupNetworkHref, cloudId)
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
func (c *Client) DestroySecurityGroupG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.DestroySecurityGroup(cloudId, id)
}

// GET api/clouds/:cloud_id/security_groups(.:format)?
// Lists Security Groups.
func (c *Client) IndexSecurityGroups(filter []string, view string, cloudId string) ([]SecurityGroup, error) {
	var res []SecurityGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/security_groups", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexSecurityGroupsG(p *Params) ([]SecurityGroup, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexSecurityGroups(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/security_groups/:id(.:format)?
// Displays information about a single Security Group.
func (c *Client) ShowSecurityGroup(view string, cloudId string, id string) (*SecurityGroup, error) {
	var res *SecurityGroup
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/security_groups/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowSecurityGroupG(p *Params) (*SecurityGroup, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowSecurityGroup(view, cloudId, id)
}

// == ServerArray ==

// A server array represents a logical group of instances and allows to
// resize(grow/shrink) that group based on certain elasticity parameters.  A
// server array just like a server always has a next_instance association, which
// will define the configuration to apply when a new instance is launched. But
// unlike a server which has a current_instance relationship, the server array
// has a current_instances relationship that gives the information about all
// the running instances in the array. Changes to the next_instance association
// prepares the configuration for the next instance that is to be launched
// in the array and will therefore not affect any of the currently running
// instances.
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
func (c *Client) CloneServerArrayG(p *Params) error {
	id := (*p)["id"].(string)
	return c.CloneServerArray(id)
}

// POST api/server_arrays(.:format)?
// Creates a new server array, and configures its corresponding "next" instance
// with the received parameters.
func (c *Client) CreateServerArray(serverArrayElasticityParamsPacingResizeUpBy string, serverArrayInstanceCloudSpecificAttributesIamInstanceProfile string, serverArrayInstanceSubnetHrefs []string, serverArrayInstanceUserData string, serverArrayDatacenterPolicy []string, serverArrayDatacenterPolicyDatacenterHref string, serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold string, serverArrayElasticityParamsScheduleMinCount string, serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping string, serverArrayInstanceInputsName string, serverArrayInstanceInputsValue string, serverArrayInstanceIpForwardingEnabled string, serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate string, serverArrayElasticityParamsBoundsMaxCount string, serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge string, serverArrayInstancePlacementGroupHref string, serverArrayInstanceCloudSpecificAttributesRootVolumePerformance string, serverArrayName string, serverArrayOptimized string, serverArrayState string, serverArrayArrayType string, serverArrayDeploymentHref string, serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance string, serverArrayInstanceCloudSpecificAttributesRootVolumeSize string, serverArrayInstanceRamdiskImageHref string, serverArrayDatacenterPolicyMax string, serverArrayElasticityParamsPacingResizeCalmTime string, serverArrayInstanceCloudHref string, serverArrayElasticityParamsScheduleTime string, serverArrayInstanceAssociatePublicIpAddress string, serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid string, serverArrayInstanceImageHref string, serverArrayInstanceKernelImageHref string, serverArrayElasticityParamsBoundsMinCount string, serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm string, serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp string, serverArrayInstanceSecurityGroupHrefs []string, serverArrayInstanceInstanceTypeHref string, serverArrayInstanceServerTemplateHref string, serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries string, serverArrayElasticityParamsSchedule []string, serverArrayInstanceDatacenterHref string, serverArrayDatacenterPolicyWeight string, serverArrayElasticityParamsScheduleDay string, serverArrayInstanceInputs string, serverArrayInstanceMultiCloudImageHref string, serverArrayInstanceSshKeyHref string, serverArrayDescription string, serverArrayElasticityParamsPacingResizeDownBy string, serverArrayElasticityParamsScheduleMaxCount string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"server_array[elasticity_params][pacing][resize_up_by]":                                  serverArrayElasticityParamsPacingResizeUpBy,
		"server_array[instance][cloud_specific_attributes][iam_instance_profile]":                serverArrayInstanceCloudSpecificAttributesIamInstanceProfile,
		"server_array[instance][subnet_hrefs]":                                                   serverArrayInstanceSubnetHrefs,
		"server_array[instance][user_data]":                                                      serverArrayInstanceUserData,
		"server_array[datacenter_policy]":                                                        serverArrayDatacenterPolicy,
		"server_array[datacenter_policy][][datacenter_href]":                                     serverArrayDatacenterPolicyDatacenterHref,
		"server_array[elasticity_params][alert_specific_params][decision_threshold]":             serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold,
		"server_array[elasticity_params][schedule][][min_count]":                                 serverArrayElasticityParamsScheduleMinCount,
		"server_array[instance][cloud_specific_attributes][automatic_instance_store_mapping]":    serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping,
		"server_array[instance][inputs][][name]":                                                 serverArrayInstanceInputsName,
		"server_array[instance][inputs][][value]":                                                serverArrayInstanceInputsValue,
		"server_array[instance][ip_forwarding_enabled]":                                          serverArrayInstanceIpForwardingEnabled,
		"server_array[elasticity_params][alert_specific_params][voters_tag_predicate]":           serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate,
		"server_array[elasticity_params][bounds][max_count]":                                     serverArrayElasticityParamsBoundsMaxCount,
		"server_array[elasticity_params][queue_specific_params][item_age][max_age]":              serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge,
		"server_array[instance][placement_group_href]":                                           serverArrayInstancePlacementGroupHref,
		"server_array[instance][cloud_specific_attributes][root_volume_performance]":             serverArrayInstanceCloudSpecificAttributesRootVolumePerformance,
		"server_array[name]":                                                                     serverArrayName,
		"server_array[optimized]":                                                                serverArrayOptimized,
		"server_array[state]":                                                                    serverArrayState,
		"server_array[array_type]":                                                               serverArrayArrayType,
		"server_array[deployment_href]":                                                          serverArrayDeploymentHref,
		"server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]": serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance,
		"server_array[instance][cloud_specific_attributes][root_volume_size]":                    serverArrayInstanceCloudSpecificAttributesRootVolumeSize,
		"server_array[instance][ramdisk_image_href]":                                             serverArrayInstanceRamdiskImageHref,
		"server_array[datacenter_policy][][max]":                                                 serverArrayDatacenterPolicyMax,
		"server_array[elasticity_params][pacing][resize_calm_time]":                              serverArrayElasticityParamsPacingResizeCalmTime,
		"server_array[instance][cloud_href]":                                                     serverArrayInstanceCloudHref,
		"server_array[elasticity_params][schedule][][time]":                                      serverArrayElasticityParamsScheduleTime,
		"server_array[instance][associate_public_ip_address]":                                    serverArrayInstanceAssociatePublicIpAddress,
		"server_array[instance][cloud_specific_attributes][root_volume_type_uid]":                serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid,
		"server_array[instance][image_href]":                                                     serverArrayInstanceImageHref,
		"server_array[instance][kernel_image_href]":                                              serverArrayInstanceKernelImageHref,
		"server_array[elasticity_params][bounds][min_count]":                                     serverArrayElasticityParamsBoundsMinCount,
		"server_array[elasticity_params][queue_specific_params][item_age][algorithm]":            serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm,
		"server_array[elasticity_params][queue_specific_params][item_age][regexp]":               serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp,
		"server_array[instance][security_group_hrefs]":                                           serverArrayInstanceSecurityGroupHrefs,
		"server_array[instance][instance_type_href]":                                             serverArrayInstanceInstanceTypeHref,
		"server_array[instance][server_template_href]":                                           serverArrayInstanceServerTemplateHref,
		"server_array[elasticity_params][queue_specific_params][collect_audit_entries]":          serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries,
		"server_array[elasticity_params][schedule]":                                              serverArrayElasticityParamsSchedule,
		"server_array[instance][datacenter_href]":                                                serverArrayInstanceDatacenterHref,
		"server_array[datacenter_policy][][weight]":                                              serverArrayDatacenterPolicyWeight,
		"server_array[elasticity_params][schedule][][day]":                                       serverArrayElasticityParamsScheduleDay,
		"server_array[instance][inputs][*]":                                                      serverArrayInstanceInputs,
		"server_array[instance][multi_cloud_image_href]":                                         serverArrayInstanceMultiCloudImageHref,
		"server_array[instance][ssh_key_href]":                                                   serverArrayInstanceSshKeyHref,
		"server_array[description]":                                                              serverArrayDescription,
		"server_array[elasticity_params][pacing][resize_down_by]":                                serverArrayElasticityParamsPacingResizeDownBy,
		"server_array[elasticity_params][schedule][][max_count]":                                 serverArrayElasticityParamsScheduleMaxCount}
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
func (c *Client) CreateServerArrayG(p *Params) (Href, error) {
	serverArrayElasticityParamsPacingResizeUpBy := (*p)["serverArrayElasticityParamsPacingResizeUpBy"].(string)
	serverArrayInstanceCloudSpecificAttributesIamInstanceProfile := (*p)["serverArrayInstanceCloudSpecificAttributesIamInstanceProfile"].(string)
	serverArrayInstanceSubnetHrefs := (*p)["serverArrayInstanceSubnetHrefs"].([]string)
	serverArrayInstanceUserData := (*p)["serverArrayInstanceUserData"].(string)
	serverArrayDatacenterPolicy := (*p)["serverArrayDatacenterPolicy"].([]string)
	serverArrayDatacenterPolicyDatacenterHref := (*p)["serverArrayDatacenterPolicyDatacenterHref"].(string)
	serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold := (*p)["serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold"].(string)
	serverArrayElasticityParamsScheduleMinCount := (*p)["serverArrayElasticityParamsScheduleMinCount"].(string)
	serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping := (*p)["serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping"].(string)
	serverArrayInstanceInputsName := (*p)["serverArrayInstanceInputsName"].(string)
	serverArrayInstanceInputsValue := (*p)["serverArrayInstanceInputsValue"].(string)
	serverArrayInstanceIpForwardingEnabled := (*p)["serverArrayInstanceIpForwardingEnabled"].(string)
	serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate := (*p)["serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate"].(string)
	serverArrayElasticityParamsBoundsMaxCount := (*p)["serverArrayElasticityParamsBoundsMaxCount"].(string)
	serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge := (*p)["serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge"].(string)
	serverArrayInstancePlacementGroupHref := (*p)["serverArrayInstancePlacementGroupHref"].(string)
	serverArrayInstanceCloudSpecificAttributesRootVolumePerformance := (*p)["serverArrayInstanceCloudSpecificAttributesRootVolumePerformance"].(string)
	serverArrayName := (*p)["serverArrayName"].(string)
	serverArrayOptimized := (*p)["serverArrayOptimized"].(string)
	serverArrayState := (*p)["serverArrayState"].(string)
	serverArrayArrayType := (*p)["serverArrayArrayType"].(string)
	serverArrayDeploymentHref := (*p)["serverArrayDeploymentHref"].(string)
	serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance := (*p)["serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance"].(string)
	serverArrayInstanceCloudSpecificAttributesRootVolumeSize := (*p)["serverArrayInstanceCloudSpecificAttributesRootVolumeSize"].(string)
	serverArrayInstanceRamdiskImageHref := (*p)["serverArrayInstanceRamdiskImageHref"].(string)
	serverArrayDatacenterPolicyMax := (*p)["serverArrayDatacenterPolicyMax"].(string)
	serverArrayElasticityParamsPacingResizeCalmTime := (*p)["serverArrayElasticityParamsPacingResizeCalmTime"].(string)
	serverArrayInstanceCloudHref := (*p)["serverArrayInstanceCloudHref"].(string)
	serverArrayElasticityParamsScheduleTime := (*p)["serverArrayElasticityParamsScheduleTime"].(string)
	serverArrayInstanceAssociatePublicIpAddress := (*p)["serverArrayInstanceAssociatePublicIpAddress"].(string)
	serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid := (*p)["serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid"].(string)
	serverArrayInstanceImageHref := (*p)["serverArrayInstanceImageHref"].(string)
	serverArrayInstanceKernelImageHref := (*p)["serverArrayInstanceKernelImageHref"].(string)
	serverArrayElasticityParamsBoundsMinCount := (*p)["serverArrayElasticityParamsBoundsMinCount"].(string)
	serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm := (*p)["serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm"].(string)
	serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp := (*p)["serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp"].(string)
	serverArrayInstanceSecurityGroupHrefs := (*p)["serverArrayInstanceSecurityGroupHrefs"].([]string)
	serverArrayInstanceInstanceTypeHref := (*p)["serverArrayInstanceInstanceTypeHref"].(string)
	serverArrayInstanceServerTemplateHref := (*p)["serverArrayInstanceServerTemplateHref"].(string)
	serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries := (*p)["serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries"].(string)
	serverArrayElasticityParamsSchedule := (*p)["serverArrayElasticityParamsSchedule"].([]string)
	serverArrayInstanceDatacenterHref := (*p)["serverArrayInstanceDatacenterHref"].(string)
	serverArrayDatacenterPolicyWeight := (*p)["serverArrayDatacenterPolicyWeight"].(string)
	serverArrayElasticityParamsScheduleDay := (*p)["serverArrayElasticityParamsScheduleDay"].(string)
	serverArrayInstanceInputs := (*p)["serverArrayInstanceInputs"].(string)
	serverArrayInstanceMultiCloudImageHref := (*p)["serverArrayInstanceMultiCloudImageHref"].(string)
	serverArrayInstanceSshKeyHref := (*p)["serverArrayInstanceSshKeyHref"].(string)
	serverArrayDescription := (*p)["serverArrayDescription"].(string)
	serverArrayElasticityParamsPacingResizeDownBy := (*p)["serverArrayElasticityParamsPacingResizeDownBy"].(string)
	serverArrayElasticityParamsScheduleMaxCount := (*p)["serverArrayElasticityParamsScheduleMaxCount"].(string)
	return c.CreateServerArray(serverArrayElasticityParamsPacingResizeUpBy, serverArrayInstanceCloudSpecificAttributesIamInstanceProfile, serverArrayInstanceSubnetHrefs, serverArrayInstanceUserData, serverArrayDatacenterPolicy, serverArrayDatacenterPolicyDatacenterHref, serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold, serverArrayElasticityParamsScheduleMinCount, serverArrayInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping, serverArrayInstanceInputsName, serverArrayInstanceInputsValue, serverArrayInstanceIpForwardingEnabled, serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate, serverArrayElasticityParamsBoundsMaxCount, serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge, serverArrayInstancePlacementGroupHref, serverArrayInstanceCloudSpecificAttributesRootVolumePerformance, serverArrayName, serverArrayOptimized, serverArrayState, serverArrayArrayType, serverArrayDeploymentHref, serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance, serverArrayInstanceCloudSpecificAttributesRootVolumeSize, serverArrayInstanceRamdiskImageHref, serverArrayDatacenterPolicyMax, serverArrayElasticityParamsPacingResizeCalmTime, serverArrayInstanceCloudHref, serverArrayElasticityParamsScheduleTime, serverArrayInstanceAssociatePublicIpAddress, serverArrayInstanceCloudSpecificAttributesRootVolumeTypeUid, serverArrayInstanceImageHref, serverArrayInstanceKernelImageHref, serverArrayElasticityParamsBoundsMinCount, serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm, serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp, serverArrayInstanceSecurityGroupHrefs, serverArrayInstanceInstanceTypeHref, serverArrayInstanceServerTemplateHref, serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries, serverArrayElasticityParamsSchedule, serverArrayInstanceDatacenterHref, serverArrayDatacenterPolicyWeight, serverArrayElasticityParamsScheduleDay, serverArrayInstanceInputs, serverArrayInstanceMultiCloudImageHref, serverArrayInstanceSshKeyHref, serverArrayDescription, serverArrayElasticityParamsPacingResizeDownBy, serverArrayElasticityParamsScheduleMaxCount)
}

// GET api/server_arrays/:id/current_instances
// List the running instances belonging to the server array. See Instances#index
// for details. This action is slightly different from invoking the
// index action on the Instances resource with the filter "parent_href ==
// /api/server_arrays/XX" because the latter will include 'next_instance' as
// well.
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
func (c *Client) CurrentInstancesServerArrayG(p *Params) error {
	id := (*p)["id"].(string)
	return c.CurrentInstancesServerArray(id)
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
func (c *Client) DestroyServerArrayG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyServerArray(id)
}

// GET api/server_arrays(.:format)?
// Lists server arrays.  By using the available filters, it is possible to
// retrieve server arrays that have common characteristics. For example, one can
// list:   arrays that have names that contain "my_server_array" all arrays of a
// given deployment
func (c *Client) IndexServerArrays(filter []string, view string) ([]ServerArray, error) {
	var res []ServerArray
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_arrays", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexServerArraysG(p *Params) ([]ServerArray, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexServerArrays(filter, view)
}

// POST api/server_arrays/:id/launch
// Launches a new instance in the server array with the configuration defined
// in the 'next_instance'. This function is equivalent to invoking the launch
// action on the URL of this server_array's next_instance. See Instances#launch
// for details.
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
func (c *Client) LaunchServerArrayG(p *Params) error {
	id := (*p)["id"].(string)
	return c.LaunchServerArray(id)
}

// POST api/server_arrays/:id/multi_run_executable
// Run an executable on all instances of this array. This function is
// equivalent to invoking the "multi_run_executable" action on the instances
// resource (Instances#multi_run_executable with the filter "parent_href ==
// /api/server_arrays/XX"). To run an executable on a subset of the instances
// of the array, provide additional filters. To run an executable a single
// instance, invoke the action "run_executable" directly on the instance (see
// Instances#run_executable)
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
func (c *Client) MultiRunExecutableServerArraysG(p *Params) error {
	id := (*p)["id"].(string)
	return c.MultiRunExecutableServerArrays(id)
}

// POST api/server_arrays/:id/multi_terminate
// Terminate all instances of this array. This function is equivalent
// to invoking the "multi_terminate" action on the instances resource
// ( Instances#multi_terminate with the filter "parent_href ==
// /api/server_arrays/XX"). To terminate a subset of the instances of the array,
// provide additional filters. To terminate a single instance, invoke the action
// "terminate" directly on the instance (see Instances#terminate)
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
func (c *Client) MultiTerminateServerArraysG(p *Params) error {
	id := (*p)["id"].(string)
	return c.MultiTerminateServerArrays(id)
}

// GET api/server_arrays/:id(.:format)?
// Shows the information of a single server array.
func (c *Client) ShowServerArray(view string, id string) (*ServerArray, error) {
	var res *ServerArray
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_arrays/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowServerArrayG(p *Params) (*ServerArray, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowServerArray(view, id)
}

// PUT api/server_arrays/:id(.:format)?
// Updates attributes of a single server array.
func (c *Client) UpdateServerArray(serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm string, serverArrayOptimized string, serverArrayDatacenterPolicyMax string, serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp string, serverArrayName string, serverArrayElasticityParamsScheduleMinCount string, serverArrayDatacenterPolicyDatacenterHref string, serverArrayElasticityParamsSchedule []string, serverArrayElasticityParamsScheduleTime string, serverArrayArrayType string, serverArrayDatacenterPolicy []string, serverArrayDatacenterPolicyWeight string, serverArrayElasticityParamsPacingResizeDownBy string, serverArrayElasticityParamsPacingResizeUpBy string, serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate string, serverArrayElasticityParamsBoundsMaxCount string, serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries string, serverArrayElasticityParamsScheduleDay string, serverArrayElasticityParamsScheduleMaxCount string, serverArrayDeploymentHref string, serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold string, serverArrayElasticityParamsBoundsMinCount string, serverArrayElasticityParamsPacingResizeCalmTime string, serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge string, serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance string, serverArrayState string, serverArrayDescription string, id string) error {
	payload := map[string]interface{}{
		"server_array[elasticity_params][queue_specific_params][item_age][algorithm]":            serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm,
		"server_array[optimized]":                                                                serverArrayOptimized,
		"server_array[datacenter_policy][][max]":                                                 serverArrayDatacenterPolicyMax,
		"server_array[elasticity_params][queue_specific_params][item_age][regexp]":               serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp,
		"server_array[name]":                                                                     serverArrayName,
		"server_array[elasticity_params][schedule][][min_count]":                                 serverArrayElasticityParamsScheduleMinCount,
		"server_array[datacenter_policy][][datacenter_href]":                                     serverArrayDatacenterPolicyDatacenterHref,
		"server_array[elasticity_params][schedule]":                                              serverArrayElasticityParamsSchedule,
		"server_array[elasticity_params][schedule][][time]":                                      serverArrayElasticityParamsScheduleTime,
		"server_array[array_type]":                                                               serverArrayArrayType,
		"server_array[datacenter_policy]":                                                        serverArrayDatacenterPolicy,
		"server_array[datacenter_policy][][weight]":                                              serverArrayDatacenterPolicyWeight,
		"server_array[elasticity_params][pacing][resize_down_by]":                                serverArrayElasticityParamsPacingResizeDownBy,
		"server_array[elasticity_params][pacing][resize_up_by]":                                  serverArrayElasticityParamsPacingResizeUpBy,
		"server_array[elasticity_params][alert_specific_params][voters_tag_predicate]":           serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate,
		"server_array[elasticity_params][bounds][max_count]":                                     serverArrayElasticityParamsBoundsMaxCount,
		"server_array[elasticity_params][queue_specific_params][collect_audit_entries]":          serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries,
		"server_array[elasticity_params][schedule][][day]":                                       serverArrayElasticityParamsScheduleDay,
		"server_array[elasticity_params][schedule][][max_count]":                                 serverArrayElasticityParamsScheduleMaxCount,
		"server_array[deployment_href]":                                                          serverArrayDeploymentHref,
		"server_array[elasticity_params][alert_specific_params][decision_threshold]":             serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold,
		"server_array[elasticity_params][bounds][min_count]":                                     serverArrayElasticityParamsBoundsMinCount,
		"server_array[elasticity_params][pacing][resize_calm_time]":                              serverArrayElasticityParamsPacingResizeCalmTime,
		"server_array[elasticity_params][queue_specific_params][item_age][max_age]":              serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge,
		"server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]": serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance,
		"server_array[state]":       serverArrayState,
		"server_array[description]": serverArrayDescription}
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
func (c *Client) UpdateServerArrayG(p *Params) error {
	serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm := (*p)["serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm"].(string)
	serverArrayOptimized := (*p)["serverArrayOptimized"].(string)
	serverArrayDatacenterPolicyMax := (*p)["serverArrayDatacenterPolicyMax"].(string)
	serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp := (*p)["serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp"].(string)
	serverArrayName := (*p)["serverArrayName"].(string)
	serverArrayElasticityParamsScheduleMinCount := (*p)["serverArrayElasticityParamsScheduleMinCount"].(string)
	serverArrayDatacenterPolicyDatacenterHref := (*p)["serverArrayDatacenterPolicyDatacenterHref"].(string)
	serverArrayElasticityParamsSchedule := (*p)["serverArrayElasticityParamsSchedule"].([]string)
	serverArrayElasticityParamsScheduleTime := (*p)["serverArrayElasticityParamsScheduleTime"].(string)
	serverArrayArrayType := (*p)["serverArrayArrayType"].(string)
	serverArrayDatacenterPolicy := (*p)["serverArrayDatacenterPolicy"].([]string)
	serverArrayDatacenterPolicyWeight := (*p)["serverArrayDatacenterPolicyWeight"].(string)
	serverArrayElasticityParamsPacingResizeDownBy := (*p)["serverArrayElasticityParamsPacingResizeDownBy"].(string)
	serverArrayElasticityParamsPacingResizeUpBy := (*p)["serverArrayElasticityParamsPacingResizeUpBy"].(string)
	serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate := (*p)["serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate"].(string)
	serverArrayElasticityParamsBoundsMaxCount := (*p)["serverArrayElasticityParamsBoundsMaxCount"].(string)
	serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries := (*p)["serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries"].(string)
	serverArrayElasticityParamsScheduleDay := (*p)["serverArrayElasticityParamsScheduleDay"].(string)
	serverArrayElasticityParamsScheduleMaxCount := (*p)["serverArrayElasticityParamsScheduleMaxCount"].(string)
	serverArrayDeploymentHref := (*p)["serverArrayDeploymentHref"].(string)
	serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold := (*p)["serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold"].(string)
	serverArrayElasticityParamsBoundsMinCount := (*p)["serverArrayElasticityParamsBoundsMinCount"].(string)
	serverArrayElasticityParamsPacingResizeCalmTime := (*p)["serverArrayElasticityParamsPacingResizeCalmTime"].(string)
	serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge := (*p)["serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge"].(string)
	serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance := (*p)["serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance"].(string)
	serverArrayState := (*p)["serverArrayState"].(string)
	serverArrayDescription := (*p)["serverArrayDescription"].(string)
	id := (*p)["id"].(string)
	return c.UpdateServerArray(serverArrayElasticityParamsQueueSpecificParamsItemAgeAlgorithm, serverArrayOptimized, serverArrayDatacenterPolicyMax, serverArrayElasticityParamsQueueSpecificParamsItemAgeRegexp, serverArrayName, serverArrayElasticityParamsScheduleMinCount, serverArrayDatacenterPolicyDatacenterHref, serverArrayElasticityParamsSchedule, serverArrayElasticityParamsScheduleTime, serverArrayArrayType, serverArrayDatacenterPolicy, serverArrayDatacenterPolicyWeight, serverArrayElasticityParamsPacingResizeDownBy, serverArrayElasticityParamsPacingResizeUpBy, serverArrayElasticityParamsAlertSpecificParamsVotersTagPredicate, serverArrayElasticityParamsBoundsMaxCount, serverArrayElasticityParamsQueueSpecificParamsCollectAuditEntries, serverArrayElasticityParamsScheduleDay, serverArrayElasticityParamsScheduleMaxCount, serverArrayDeploymentHref, serverArrayElasticityParamsAlertSpecificParamsDecisionThreshold, serverArrayElasticityParamsBoundsMinCount, serverArrayElasticityParamsPacingResizeCalmTime, serverArrayElasticityParamsQueueSpecificParamsItemAgeMaxAge, serverArrayElasticityParamsQueueSpecificParamsQueueSizeItemsPerInstance, serverArrayState, serverArrayDescription, id)
}

// == ServerTemplateMultiCloudImage ==

// This resource represents links between ServerTemplates and MultiCloud Images
// and enables you to effectively add/delete MultiCloudImages to ServerTemplates
// and make them the default one.
type ServerTemplateMultiCloudImage struct {
	Actions   []string            `json:"actions,omitempty"`
	CreatedAt *time.Time          `json:"created_at,omitempty"`
	IsDefault bool                `json:"is_default,omitempty"`
	Links     []map[string]string `json:"links,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

// POST api/server_template_multi_cloud_images(.:format)?
// Creates a new ServerTemplateMultiCloudImage with the given parameters.
func (c *Client) CreateServerTemplateMultiCloudImage(serverTemplateMultiCloudImageServerTemplateHref string, serverTemplateMultiCloudImageMultiCloudImageHref string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"server_template_multi_cloud_image[server_template_href]":   serverTemplateMultiCloudImageServerTemplateHref,
		"server_template_multi_cloud_image[multi_cloud_image_href]": serverTemplateMultiCloudImageMultiCloudImageHref}
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
func (c *Client) CreateServerTemplateMultiCloudImageG(p *Params) (Href, error) {
	serverTemplateMultiCloudImageServerTemplateHref := (*p)["serverTemplateMultiCloudImageServerTemplateHref"].(string)
	serverTemplateMultiCloudImageMultiCloudImageHref := (*p)["serverTemplateMultiCloudImageMultiCloudImageHref"].(string)
	return c.CreateServerTemplateMultiCloudImage(serverTemplateMultiCloudImageServerTemplateHref, serverTemplateMultiCloudImageMultiCloudImageHref)
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
func (c *Client) DestroyServerTemplateMultiCloudImageG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyServerTemplateMultiCloudImage(id)
}

// GET api/server_template_multi_cloud_images(.:format)?
// Lists the ServerTemplateMultiCloudImages available to this account.
func (c *Client) IndexServerTemplateMultiCloudImages(filter []string, view string) ([]ServerTemplateMultiCloudImage, error) {
	var res []ServerTemplateMultiCloudImage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_template_multi_cloud_images", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexServerTemplateMultiCloudImagesG(p *Params) ([]ServerTemplateMultiCloudImage, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexServerTemplateMultiCloudImages(filter, view)
}

// POST api/server_template_multi_cloud_images/:id/make_default(.:format)?
// Makes a given ServerTemplateMultiCloudImage the default for the
// ServerTemplate.
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
func (c *Client) MakeDefaultServerTemplateMultiCloudImageG(p *Params) error {
	id := (*p)["id"].(string)
	return c.MakeDefaultServerTemplateMultiCloudImage(id)
}

// GET api/server_template_multi_cloud_images/:id(.:format)?
// Show information about a single ServerTemplateMultiCloudImage which
// represents an association between a ServerTemplate and a MultiCloudImage.
func (c *Client) ShowServerTemplateMultiCloudImage(view string, id string) (*ServerTemplateMultiCloudImage, error) {
	var res *ServerTemplateMultiCloudImage
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_template_multi_cloud_images/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowServerTemplateMultiCloudImageG(p *Params) (*ServerTemplateMultiCloudImage, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowServerTemplateMultiCloudImage(view, id)
}

// == ServerTemplate ==

// ServerTemplates allow you to pre-configure servers by starting from a base
// image and adding scripts that run during the boot, operational, and shutdown
// phases. A ServerTemplate is a description of how a new instance will be
// configured when it is provisioned by your cloud provider.  All revisions of
// a ServerTemplate belong to a ServerTemplate lineage that is exposed by the
// "lineage" attribute. (NOTE: This attribute is merely a string to locate all
// revisions of a ServerTemplate and NOT a working URL)
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
func (c *Client) CloneServerTemplate(serverTemplateDescription string, serverTemplateName string, id string) error {
	payload := map[string]interface{}{
		"server_template[description]": serverTemplateDescription,
		"server_template[name]":        serverTemplateName}
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
func (c *Client) CloneServerTemplateG(p *Params) error {
	serverTemplateDescription := (*p)["serverTemplateDescription"].(string)
	serverTemplateName := (*p)["serverTemplateName"].(string)
	id := (*p)["id"].(string)
	return c.CloneServerTemplate(serverTemplateDescription, serverTemplateName, id)
}

// POST api/server_templates/:id/commit(.:format)?
// Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are
// owned by the account can be committed.
func (c *Client) CommitServerTemplate(commitHeadDependencies string, commitMessage string, freezeRepositories string, id string) error {
	payload := map[string]interface{}{
		"commit_head_dependencies": commitHeadDependencies,
		"commit_message":           commitMessage,
		"freeze_repositories":      freezeRepositories}
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
func (c *Client) CommitServerTemplateG(p *Params) error {
	commitHeadDependencies := (*p)["commitHeadDependencies"].(string)
	commitMessage := (*p)["commitMessage"].(string)
	freezeRepositories := (*p)["freezeRepositories"].(string)
	id := (*p)["id"].(string)
	return c.CommitServerTemplate(commitHeadDependencies, commitMessage, freezeRepositories, id)
}

// POST api/server_templates(.:format)?
// Creates a new ServerTemplate with the given parameters.
func (c *Client) CreateServerTemplate(serverTemplateDescription string, serverTemplateName string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"server_template[description]": serverTemplateDescription,
		"server_template[name]":        serverTemplateName}
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
func (c *Client) CreateServerTemplateG(p *Params) (Href, error) {
	serverTemplateDescription := (*p)["serverTemplateDescription"].(string)
	serverTemplateName := (*p)["serverTemplateName"].(string)
	return c.CreateServerTemplate(serverTemplateDescription, serverTemplateName)
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
func (c *Client) DestroyServerTemplateG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyServerTemplate(id)
}

// POST api/server_templates/:id/detect_changes_in_head(.:format)?
// Identifies RightScripts attached to the resource that differ from their HEAD.
//  If the attached revision of the RightScript is the HEAD, then this will
// indicate a difference between it and the latest committed revision in the
// same lineage.
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
func (c *Client) DetectChangesInHeadServerTemplateG(p *Params) ([]map[string]string, error) {
	id := (*p)["id"].(string)
	return c.DetectChangesInHeadServerTemplate(id)
}

// GET api/server_templates(.:format)?
// Lists the ServerTemplates available to this account. HEAD revisions have
// a revision of 0.  The 'inputs_2_0' view is for retrieving inputs in 2.0
// serialization (for more details please see Inputs#index.)
func (c *Client) IndexServerTemplates(filter []string, view string) ([]ServerTemplate, error) {
	var res []ServerTemplate
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexServerTemplatesG(p *Params) ([]ServerTemplate, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexServerTemplates(filter, view)
}

// POST api/server_templates/:id/publish(.:format)?
// Publishes a given ServerTemplate and its subordinates. Only non-HEAD
// revisions that are owned by the account can be published.
func (c *Client) PublishServerTemplate(descriptionsLong string, descriptionsNotes string, descriptionsShort string, emailComments string, accountGroupHrefs []string, allowComments string, categories []string, id string) error {
	payload := map[string]interface{}{
		"descriptions[long]":  descriptionsLong,
		"descriptions[notes]": descriptionsNotes,
		"descriptions[short]": descriptionsShort,
		"email_comments":      emailComments,
		"account_group_hrefs": accountGroupHrefs,
		"allow_comments":      allowComments,
		"categories":          categories}
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
func (c *Client) PublishServerTemplateG(p *Params) error {
	descriptionsLong := (*p)["descriptionsLong"].(string)
	descriptionsNotes := (*p)["descriptionsNotes"].(string)
	descriptionsShort := (*p)["descriptionsShort"].(string)
	emailComments := (*p)["emailComments"].(string)
	accountGroupHrefs := (*p)["accountGroupHrefs"].([]string)
	allowComments := (*p)["allowComments"].(string)
	categories := (*p)["categories"].([]string)
	id := (*p)["id"].(string)
	return c.PublishServerTemplate(descriptionsLong, descriptionsNotes, descriptionsShort, emailComments, accountGroupHrefs, allowComments, categories, id)
}

// POST api/server_templates/:id/resolve(.:format)?
// Enumerates all attached cookbooks, missing dependencies and bound
// executables.  Version constraints on missing dependencies and the state of
// the Chef Recipes; whether or not the cookbook or recipe itself could be found
// among the attachments, will also be reported.
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
func (c *Client) ResolveServerTemplateG(p *Params) ([]map[string]string, error) {
	id := (*p)["id"].(string)
	return c.ResolveServerTemplate(id)
}

// GET api/server_templates/:id(.:format)?
// Show information about a single ServerTemplate. HEAD revisions have a
// revision of 0.  The 'inputs_2_0' view is for retrieving inputs in 2.0
// serialization (for more details please see Inputs#index.)
func (c *Client) ShowServerTemplate(view string, id string) (*ServerTemplate, error) {
	var res *ServerTemplate
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/server_templates/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowServerTemplateG(p *Params) (*ServerTemplate, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowServerTemplate(view, id)
}

// POST api/server_templates/:id/swap_repository(.:format)?
// In-place replacement of attached cookbooks from a given repository.  For each
// attached cookbook coming from the source repository, replace it by attaching
// a cookbook of identical name coming from the target repository.  In order
// for the operation to be successful, all attachments that came from the source
// repository must exist in the target repository.  If multiple cookbooks of
// a given name exist in the target repository, preference is given by the
// following order (top most being the highest preference):     Name & Version
// Match / Primary Namespace   Name & Version Match / Alternate Namespace   Name
// Match / Primary Namespace   Name Match / Alternate Namespace   If multiple
// cookbooks still have the same preference for the replacement, the operation
// is indeterministic.
func (c *Client) SwapRepositoryServerTemplate(targetRepositoryHref string, sourceRepositoryHref string, id string) error {
	payload := map[string]interface{}{
		"target_repository_href": targetRepositoryHref,
		"source_repository_href": sourceRepositoryHref}
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
func (c *Client) SwapRepositoryServerTemplateG(p *Params) error {
	targetRepositoryHref := (*p)["targetRepositoryHref"].(string)
	sourceRepositoryHref := (*p)["sourceRepositoryHref"].(string)
	id := (*p)["id"].(string)
	return c.SwapRepositoryServerTemplate(targetRepositoryHref, sourceRepositoryHref, id)
}

// PUT api/server_templates/:id(.:format)?
// Updates attributes of a given ServerTemplate. Only HEAD revisions can be
// updated (revision 0). Currently, the attributes you can update are only the
// 'direct' attributes of a server template. To manage multi cloud images of a
// ServerTemplate, please see the resource 'ServerTemplateMultiCloudImages'.
func (c *Client) UpdateServerTemplate(serverTemplateDescription string, serverTemplateName string, id string) error {
	payload := map[string]interface{}{
		"server_template[description]": serverTemplateDescription,
		"server_template[name]":        serverTemplateName}
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
func (c *Client) UpdateServerTemplateG(p *Params) error {
	serverTemplateDescription := (*p)["serverTemplateDescription"].(string)
	serverTemplateName := (*p)["serverTemplateName"].(string)
	id := (*p)["id"].(string)
	return c.UpdateServerTemplate(serverTemplateDescription, serverTemplateName, id)
}

// == Server ==

// Servers represent the notion of a server/machine from the RightScale's
// perspective. A Server, does not always have a corresponding VM running
// or provisioned in a cloud. Some clouds use the word "servers" to refer to
// created VM's. These allocated VM's are not called Servers in the RightScale
// API, they are called Instances.  A Server always has a next_instance
// association, which will define the configuration to apply to a new instance
// when the server is launched or started (starting servers is not yet supported
// through this API). Once a Server is launched/started a current_instance
// relationship will exist. Accessing the current_instance of a server results
// in immediate runtime modification of this running server. Changes to the
// next_instance association prepares the configuration for the next instance
// launch/start (therefore they have no effect until such operation is
// performed).
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
func (c *Client) CloneServerG(p *Params) error {
	id := (*p)["id"].(string)
	return c.CloneServer(id)
}

// POST api/servers(.:format)?
// Creates a new server, and configures its corresponding "next" instance with
// the received parameters.
func (c *Client) CreateServer(serverInstanceDatacenterHref string, serverInstanceRamdiskImageHref string, serverInstanceServerTemplateHref string, serverInstanceAssociatePublicIpAddress string, serverInstanceMultiCloudImageHref string, serverInstanceSubnetHrefs []string, serverOptimized string, serverDescription string, serverInstanceInputs string, serverInstanceKernelImageHref string, serverInstanceCloudSpecificAttributesRootVolumeTypeUid string, serverInstanceCloudHref string, serverInstanceCloudSpecificAttributesRootVolumeSize string, serverInstanceInputsName string, serverInstancePlacementGroupHref string, serverDeploymentHref string, serverInstanceUserData string, serverInstanceInputsValue string, serverInstanceInstanceTypeHref string, serverInstanceSecurityGroupHrefs []string, serverInstanceImageHref string, serverInstanceIpForwardingEnabled string, serverInstanceCloudSpecificAttributesIamInstanceProfile string, serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping string, serverInstanceCloudSpecificAttributesRootVolumePerformance string, serverInstanceSshKeyHref string, serverName string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"server[instance][datacenter_href]":                                             serverInstanceDatacenterHref,
		"server[instance][ramdisk_image_href]":                                          serverInstanceRamdiskImageHref,
		"server[instance][server_template_href]":                                        serverInstanceServerTemplateHref,
		"server[instance][associate_public_ip_address]":                                 serverInstanceAssociatePublicIpAddress,
		"server[instance][multi_cloud_image_href]":                                      serverInstanceMultiCloudImageHref,
		"server[instance][subnet_hrefs]":                                                serverInstanceSubnetHrefs,
		"server[optimized]":                                                             serverOptimized,
		"server[description]":                                                           serverDescription,
		"server[instance][inputs][*]":                                                   serverInstanceInputs,
		"server[instance][kernel_image_href]":                                           serverInstanceKernelImageHref,
		"server[instance][cloud_specific_attributes][root_volume_type_uid]":             serverInstanceCloudSpecificAttributesRootVolumeTypeUid,
		"server[instance][cloud_href]":                                                  serverInstanceCloudHref,
		"server[instance][cloud_specific_attributes][root_volume_size]":                 serverInstanceCloudSpecificAttributesRootVolumeSize,
		"server[instance][inputs][][name]":                                              serverInstanceInputsName,
		"server[instance][placement_group_href]":                                        serverInstancePlacementGroupHref,
		"server[deployment_href]":                                                       serverDeploymentHref,
		"server[instance][user_data]":                                                   serverInstanceUserData,
		"server[instance][inputs][][value]":                                             serverInstanceInputsValue,
		"server[instance][instance_type_href]":                                          serverInstanceInstanceTypeHref,
		"server[instance][security_group_hrefs]":                                        serverInstanceSecurityGroupHrefs,
		"server[instance][image_href]":                                                  serverInstanceImageHref,
		"server[instance][ip_forwarding_enabled]":                                       serverInstanceIpForwardingEnabled,
		"server[instance][cloud_specific_attributes][iam_instance_profile]":             serverInstanceCloudSpecificAttributesIamInstanceProfile,
		"server[instance][cloud_specific_attributes][automatic_instance_store_mapping]": serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping,
		"server[instance][cloud_specific_attributes][root_volume_performance]":          serverInstanceCloudSpecificAttributesRootVolumePerformance,
		"server[instance][ssh_key_href]":                                                serverInstanceSshKeyHref,
		"server[name]":                                                                  serverName}
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
func (c *Client) CreateServerG(p *Params) (Href, error) {
	serverInstanceDatacenterHref := (*p)["serverInstanceDatacenterHref"].(string)
	serverInstanceRamdiskImageHref := (*p)["serverInstanceRamdiskImageHref"].(string)
	serverInstanceServerTemplateHref := (*p)["serverInstanceServerTemplateHref"].(string)
	serverInstanceAssociatePublicIpAddress := (*p)["serverInstanceAssociatePublicIpAddress"].(string)
	serverInstanceMultiCloudImageHref := (*p)["serverInstanceMultiCloudImageHref"].(string)
	serverInstanceSubnetHrefs := (*p)["serverInstanceSubnetHrefs"].([]string)
	serverOptimized := (*p)["serverOptimized"].(string)
	serverDescription := (*p)["serverDescription"].(string)
	serverInstanceInputs := (*p)["serverInstanceInputs"].(string)
	serverInstanceKernelImageHref := (*p)["serverInstanceKernelImageHref"].(string)
	serverInstanceCloudSpecificAttributesRootVolumeTypeUid := (*p)["serverInstanceCloudSpecificAttributesRootVolumeTypeUid"].(string)
	serverInstanceCloudHref := (*p)["serverInstanceCloudHref"].(string)
	serverInstanceCloudSpecificAttributesRootVolumeSize := (*p)["serverInstanceCloudSpecificAttributesRootVolumeSize"].(string)
	serverInstanceInputsName := (*p)["serverInstanceInputsName"].(string)
	serverInstancePlacementGroupHref := (*p)["serverInstancePlacementGroupHref"].(string)
	serverDeploymentHref := (*p)["serverDeploymentHref"].(string)
	serverInstanceUserData := (*p)["serverInstanceUserData"].(string)
	serverInstanceInputsValue := (*p)["serverInstanceInputsValue"].(string)
	serverInstanceInstanceTypeHref := (*p)["serverInstanceInstanceTypeHref"].(string)
	serverInstanceSecurityGroupHrefs := (*p)["serverInstanceSecurityGroupHrefs"].([]string)
	serverInstanceImageHref := (*p)["serverInstanceImageHref"].(string)
	serverInstanceIpForwardingEnabled := (*p)["serverInstanceIpForwardingEnabled"].(string)
	serverInstanceCloudSpecificAttributesIamInstanceProfile := (*p)["serverInstanceCloudSpecificAttributesIamInstanceProfile"].(string)
	serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping := (*p)["serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping"].(string)
	serverInstanceCloudSpecificAttributesRootVolumePerformance := (*p)["serverInstanceCloudSpecificAttributesRootVolumePerformance"].(string)
	serverInstanceSshKeyHref := (*p)["serverInstanceSshKeyHref"].(string)
	serverName := (*p)["serverName"].(string)
	return c.CreateServer(serverInstanceDatacenterHref, serverInstanceRamdiskImageHref, serverInstanceServerTemplateHref, serverInstanceAssociatePublicIpAddress, serverInstanceMultiCloudImageHref, serverInstanceSubnetHrefs, serverOptimized, serverDescription, serverInstanceInputs, serverInstanceKernelImageHref, serverInstanceCloudSpecificAttributesRootVolumeTypeUid, serverInstanceCloudHref, serverInstanceCloudSpecificAttributesRootVolumeSize, serverInstanceInputsName, serverInstancePlacementGroupHref, serverDeploymentHref, serverInstanceUserData, serverInstanceInputsValue, serverInstanceInstanceTypeHref, serverInstanceSecurityGroupHrefs, serverInstanceImageHref, serverInstanceIpForwardingEnabled, serverInstanceCloudSpecificAttributesIamInstanceProfile, serverInstanceCloudSpecificAttributesAutomaticInstanceStoreMapping, serverInstanceCloudSpecificAttributesRootVolumePerformance, serverInstanceSshKeyHref, serverName)
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
func (c *Client) DestroyServerG(p *Params) error {
	id := (*p)["id"].(string)
	return c.DestroyServer(id)
}

// GET api/servers(.:format)?
// Lists servers.  By using the available filters, it is possible to retrieve
// servers that have common characteristics. For example, one can list:
// servers that have names that contain "app_server" all servers of a given
// deployment   For more filters, please see the 'index' action on 'Instances'
// resource as most of the attributes belong to a 'current_instance' than to a
// server.
func (c *Client) IndexServers(filter []string, view string) ([]Server, error) {
	var res []Server
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexServersG(p *Params) ([]Server, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	return c.IndexServers(filter, view)
}

// POST api/servers/:id/launch
// Launches the "next" instance of this server. This function is equivalent
// to invoking the launch action on the URL of this servers next_instance. See
// Instances#launch for details.
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
func (c *Client) LaunchServerG(p *Params) error {
	id := (*p)["id"].(string)
	return c.LaunchServer(id)
}

// GET api/servers/:id(.:format)?
// Shows the information of a single server.
func (c *Client) ShowServer(view string, id string) (*Server, error) {
	var res *Server
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/servers/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowServerG(p *Params) (*Server, error) {
	view := (*p)["view"].(string)
	id := (*p)["id"].(string)
	return c.ShowServer(view, id)
}

// POST api/servers/:id/teminate
// Terminates the current instance of this server. This function is equivalent
// to invoking the terminate action on the URL of this servers current_instance.
// See Instances#terminate for details.
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
func (c *Client) TerminateServerG(p *Params) error {
	id := (*p)["id"].(string)
	return c.TerminateServer(id)
}

// PUT api/servers/:id(.:format)?
// Updates attributes of a single server.
func (c *Client) UpdateServer(serverAutomaticInstanceStoreMapping string, serverDescription string, serverName string, serverOptimized string, serverRootVolumeSize string, id string) error {
	payload := map[string]interface{}{
		"server[automatic_instance_store_mapping]": serverAutomaticInstanceStoreMapping,
		"server[description]":                      serverDescription,
		"server[name]":                             serverName,
		"server[optimized]":                        serverOptimized,
		"server[root_volume_size]":                 serverRootVolumeSize}
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
func (c *Client) UpdateServerG(p *Params) error {
	serverAutomaticInstanceStoreMapping := (*p)["serverAutomaticInstanceStoreMapping"].(string)
	serverDescription := (*p)["serverDescription"].(string)
	serverName := (*p)["serverName"].(string)
	serverOptimized := (*p)["serverOptimized"].(string)
	serverRootVolumeSize := (*p)["serverRootVolumeSize"].(string)
	id := (*p)["id"].(string)
	return c.UpdateServer(serverAutomaticInstanceStoreMapping, serverDescription, serverName, serverOptimized, serverRootVolumeSize, id)
}

// POST api/servers/wrap_instance(.:format)?
// Wrap an existing instance and set current instance for new server
func (c *Client) WrapInstanceServer(serverDescription string, serverInstanceHref string, serverInstanceInputs string, serverInstanceMultiCloudImageHref string, serverInstanceServerTemplateHref string, serverName string, serverDeploymentHref string) error {
	payload := map[string]interface{}{
		"server[description]":                      serverDescription,
		"server[instance][href]":                   serverInstanceHref,
		"server[instance][inputs][*]":              serverInstanceInputs,
		"server[instance][multi_cloud_image_href]": serverInstanceMultiCloudImageHref,
		"server[instance][server_template_href]":   serverInstanceServerTemplateHref,
		"server[name]":                             serverName,
		"server[deployment_href]":                  serverDeploymentHref}
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
func (c *Client) WrapInstanceServerG(p *Params) error {
	serverDescription := (*p)["serverDescription"].(string)
	serverInstanceHref := (*p)["serverInstanceHref"].(string)
	serverInstanceInputs := (*p)["serverInstanceInputs"].(string)
	serverInstanceMultiCloudImageHref := (*p)["serverInstanceMultiCloudImageHref"].(string)
	serverInstanceServerTemplateHref := (*p)["serverInstanceServerTemplateHref"].(string)
	serverName := (*p)["serverName"].(string)
	serverDeploymentHref := (*p)["serverDeploymentHref"].(string)
	return c.WrapInstanceServer(serverDescription, serverInstanceHref, serverInstanceInputs, serverInstanceMultiCloudImageHref, serverInstanceServerTemplateHref, serverName, serverDeploymentHref)
}

// == Session ==

// The sessions resource is in charge of creating API sessions that are
// bound to a given account. The sequence for login into the API is the
// following:   Perform a POST request to /api/sessions ('create' action) to
// my.rightscale.com or to any more specific hosts saved from previous sessions.
// If the targeted host is not appropriate for the specific account being
// accessed it will return a 302 http code with a URL with which the client must
// retry the same POST request. If the targeted host is the right one and the
// login is successful, it will return a 204 http code, along with two cookies
// that will need to be saved and passed in any subsequent API request. If there
// is an authentication or authorization problem with the POST request an error
// (typically 401 or 422 ) may be returned at any point in the above sequence.
// If the session expires, it will return a 403 http code with a "Session cookie
// is expired or invalid" message.    Note that all API calls irrespective of
// the resource it is acting on, should pass a header "X_API_VERSION" with the
// value "1.5".
type Session struct {
	Actions []string            `json:"actions,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
	Message string              `json:"message,omitempty"`
}

// GET api/sessions/accounts(.:format)?
// List all the accounts that a user has access to.  This call may be executed
// outside of an existing session. Doing so requires passing a username and
// password in the request body. The idea is that it should be possible to
// list accounts that can be used to create a session.  Upon successfully
// authenticating the credentials, the system will return a 200 OK code and
// return the list of accounts. If an 302 redirect code is returned, the
// client is responsible of re-issuing the GET request against the content
// of the received Location header, passing the exact same parameters again.
//   Example Request using Curl (not using an existing session): curl -i -H
// X_API_VERSION:1.5 -X GET -d email='email@me.com' -d password='mypassword'
// https://my.rightscale.com/api/sessions/accounts  Example Request using Curl
// (using an existing session): curl -i -H X_API_VERSION:1.5 -X GET -b mycookies
// https://my.rightscale.com/api/sessions/accounts
func (c *Client) AccountsSession(view string, email string, password string) ([]Account, error) {
	var res []Account
	payload := map[string]interface{}{
		"email":    email,
		"password": password}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/sessions/accounts", body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) AccountsSessionG(p *Params) ([]Account, error) {
	view := (*p)["view"].(string)
	email := (*p)["email"].(string)
	password := (*p)["password"].(string)
	return c.AccountsSession(view, email, password)
}

// POST api/sessions(.:format)?
// Creates API session scoped to a given account. (API login)  This call
// requires a form of authentication (user and password), as well as the account
// for which the session needs to be created. Upon successfully authenticating
// the credentials, the system will return a 204 code and set of two cookies
// that will serve as the credentials for the session. Both of these cookies
// must be passed in any of the subsequent requests for this session. If an
// 302 redirect code is returned, the client is responsible of re-issuing
// the POST request against the content of the received Location header,
// passing the exact same parameters again.   Example Request using Curl:
// curl -i -H X_API_VERSION:1.5 -c mycookies -X POST -d email='email@me.com'
// -d password='mypassword' -d account_href=/api/accounts/11
// https://my.rightscale.com/api/sessions
func (c *Client) CreateSession(email string, password string, accountHref string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"email":        email,
		"password":     password,
		"account_href": accountHref}
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
func (c *Client) CreateSessionG(p *Params) (Href, error) {
	email := (*p)["email"].(string)
	password := (*p)["password"].(string)
	accountHref := (*p)["accountHref"].(string)
	return c.CreateSession(email, password, accountHref)
}

// POST api/sessions/instance(.:format)?
// Creates API session scoped to a given account and instance.  This call
// requires a form of authentication (token), as well as the account for
// which the session needs to be created. Upon successfully authenticating
// the credentials, the system will return a 204 code and set of two cookies
// that will serve as the credentials for the session. Both of these cookies
// must be passed in any of the subsequent requests for this session. If an
// 302 redirect code is returned, the client is responsible of re-issuing
// the POST request against the content of the received Location header,
// passing the exact same parameters again.   Users can find their account
// ID and instance\_token from their instance's user_data: account ID regex:
// /RS_API_TOKEN=(\d+):/ instance_token regex: /RS_API_TOKEN=(?:\d+):(\w+)&/
//  Example Request using Curl: curl -i -H X_API_VERSION:1.5 -c mycookies
// -X POST -d instance_token='randomtoken' -d account_href=/api/accounts/11
// https://my.rightscale.com/api/sessions/instance
func (c *Client) CreateInstanceSessionSession(accountHref string, instanceToken string) error {
	payload := map[string]interface{}{
		"account_href":   accountHref,
		"instance_token": instanceToken}
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
func (c *Client) CreateInstanceSessionSessionG(p *Params) error {
	accountHref := (*p)["accountHref"].(string)
	instanceToken := (*p)["instanceToken"].(string)
	return c.CreateInstanceSessionSession(accountHref, instanceToken)
}

// GET api/sessions(.:format)?
// Returns a list of root resources so an authenticated session can use them
// as a starting point or a way to know what features are available within its
// privileges.   Example Request using Curl: curl -i -H X_API_VERSION:1.5 -b
// mycookies -X GET https://my.rightscale.com/api/sessions
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
func (c *Client) IndexSessionsG(p *Params) ([]Session, error) {
	return c.IndexSessions()
}

// GET api/sessions/instance(.:format)?
// Shows the full attributes of the instance (that has the token used to
// log-in). This call can be used by an instance to get it's own details.
// Example Request using Curl: curl -i -H X_API_VERSION:1.5 -b mycookies -X GET
// https://my.rightscale.com/api/sessions/instance
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
func (c *Client) IndexInstanceSessionSessionG(p *Params) (Instance, error) {
	return c.IndexInstanceSessionSession()
}

// == SshKey ==

// Ssh Keys represent a created SSH Key that exists in the cloud.  An ssh key
// might also contain the private part of the key, and can be used to login to
// instances launched with it.
type SshKey struct {
	Actions     []string            `json:"actions,omitempty"`
	Links       []map[string]string `json:"links,omitempty"`
	Material    string              `json:"material,omitempty"`
	ResourceUid string              `json:"resource_uid,omitempty"`
}

// POST api/clouds/:cloud_id/ssh_keys(.:format)?
// Creates a new ssh key.
func (c *Client) CreateSshKey(sshKeyName string, cloudId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"ssh_key[name]": sshKeyName}
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
func (c *Client) CreateSshKeyG(p *Params) (Href, error) {
	sshKeyName := (*p)["sshKeyName"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.CreateSshKey(sshKeyName, cloudId)
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
func (c *Client) DestroySshKeyG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.DestroySshKey(cloudId, id)
}

// GET api/clouds/:cloud_id/ssh_keys(.:format)?
// Lists ssh keys.
func (c *Client) IndexSshKeys(filter []string, view string, cloudId string) ([]SshKey, error) {
	var res []SshKey
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ssh_keys", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexSshKeysG(p *Params) ([]SshKey, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexSshKeys(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/ssh_keys/:id(.:format)?
// Displays information about a single ssh key.
func (c *Client) ShowSshKey(view string, cloudId string, id string) (*SshKey, error) {
	var res *SshKey
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/ssh_keys/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowSshKeyG(p *Params) (*SshKey, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowSshKey(view, cloudId, id)
}

// == Subnet ==

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
func (c *Client) CreateSubnet(subnetName string, subnetNetworkHref string, subnetCidrBlock string, subnetDatacenterHref string, subnetDescription string, cloudId string, instanceId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"subnet[name]":            subnetName,
		"subnet[network_href]":    subnetNetworkHref,
		"subnet[cidr_block]":      subnetCidrBlock,
		"subnet[datacenter_href]": subnetDatacenterHref,
		"subnet[description]":     subnetDescription}
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
func (c *Client) CreateSubnetG(p *Params) (Href, error) {
	subnetName := (*p)["subnetName"].(string)
	subnetNetworkHref := (*p)["subnetNetworkHref"].(string)
	subnetCidrBlock := (*p)["subnetCidrBlock"].(string)
	subnetDatacenterHref := (*p)["subnetDatacenterHref"].(string)
	subnetDescription := (*p)["subnetDescription"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.CreateSubnet(subnetName, subnetNetworkHref, subnetCidrBlock, subnetDatacenterHref, subnetDescription, cloudId, instanceId)
}

// DELETE api/clouds/:cloud_id/instances/:instance_id/subnets/:id(.:format)?
// Deletes the given subnet(s).
func (c *Client) DestroySubnet(cloudId string, instanceId string, id string) error {
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
func (c *Client) DestroySubnetG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.DestroySubnet(cloudId, instanceId, id)
}

// GET api/clouds/:cloud_id/instances/:instance_id/subnets(.:format)?
// Lists subnets of a given cloud.
func (c *Client) IndexSubnets(filter []string, cloudId string, instanceId string) ([]Subnet, error) {
	var res []Subnet
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/subnets", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexSubnetsG(p *Params) ([]Subnet, error) {
	filter := (*p)["filter"].([]string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.IndexSubnets(filter, cloudId, instanceId)
}

// GET api/clouds/:cloud_id/instances/:instance_id/subnets/:id(.:format)?
// Shows attributes of a single subnet.
func (c *Client) ShowSubnet(cloudId string, instanceId string, id string) (*Subnet, error) {
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
func (c *Client) ShowSubnetG(p *Params) (*Subnet, error) {
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.ShowSubnet(cloudId, instanceId, id)
}

// PUT api/clouds/:cloud_id/instances/:instance_id/subnets/:id(.:format)?
// Updates name and description for the current subnet.
func (c *Client) UpdateSubnet(subnetName string, subnetRouteTableHref string, subnetDescription string, cloudId string, instanceId string, id string) error {
	payload := map[string]interface{}{
		"subnet[name]":             subnetName,
		"subnet[route_table_href]": subnetRouteTableHref,
		"subnet[description]":      subnetDescription}
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
func (c *Client) UpdateSubnetG(p *Params) error {
	subnetName := (*p)["subnetName"].(string)
	subnetRouteTableHref := (*p)["subnetRouteTableHref"].(string)
	subnetDescription := (*p)["subnetDescription"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.UpdateSubnet(subnetName, subnetRouteTableHref, subnetDescription, cloudId, instanceId, id)
}

// == Tag ==

// A tag or machine tag is a useful way of attaching useful metadata to an
// object/resource. Tags are commonly used as an extra label or identifier. For
// example, you might want to add a tag to an EBS Snapshot or AMI so that you
// can find it more quickly.
type Tag struct {
}

// POST api/tags/by_resource(.:format)?
// Get tags for a list of resource hrefs. The hrefs can belong to various
// resource types and the tags for a non-existent href will be empty.
func (c *Client) ByResourceTag(resourceHrefs []string) ([]map[string]string, error) {
	var res []map[string]string
	payload := map[string]interface{}{
		"resource_hrefs": resourceHrefs}
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
func (c *Client) ByResourceTagG(p *Params) ([]map[string]string, error) {
	resourceHrefs := (*p)["resourceHrefs"].([]string)
	return c.ByResourceTag(resourceHrefs)
}

// POST api/tags/by_tag(.:format)?
// Search for resources having a list of tags in a specific resource_type.
// The search criteria can contain plain tags ("my_db_server"), machine tags
// ("server:db=true"), or namespace &amp; predicate wildcards ("server:db=*").
// The result set includes links to the resources.  If match_all is "true",
// then the search is an "AND" operation -- only resources having all the
// tags are returned. If match_all has any other value or is missing, the
// search is performed as an "OR" operation.  PLEASE NOTE the behavior of the
// include_tags_with_prefix parameter: if it is absent, or blank, then you will
// find resources that match your query but receive no information about the
// tags that apply to those resources. (Your search will also complete much more
// quickly in this case.)  For example, a search with tag[]="server:db=true"
// and include_tags_with_prefix="backup:" will return resources that are tagged
// as a DB server, and also return all "backup" related tags for every matching
// resource.
func (c *Client) ByTagTag(tags []string, withDeleted string, includeTagsWithPrefix string, matchAll string, resourceType string) ([]map[string]string, error) {
	var res []map[string]string
	payload := map[string]interface{}{
		"tags":                     tags,
		"with_deleted":             withDeleted,
		"include_tags_with_prefix": includeTagsWithPrefix,
		"match_all":                matchAll,
		"resource_type":            resourceType}
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
func (c *Client) ByTagTagG(p *Params) ([]map[string]string, error) {
	tags := (*p)["tags"].([]string)
	withDeleted := (*p)["withDeleted"].(string)
	includeTagsWithPrefix := (*p)["includeTagsWithPrefix"].(string)
	matchAll := (*p)["matchAll"].(string)
	resourceType := (*p)["resourceType"].(string)
	return c.ByTagTag(tags, withDeleted, includeTagsWithPrefix, matchAll, resourceType)
}

// POST api/tags/multi_add(.:format)?
// Add a list of tags to a list of hrefs. The tags must be either plain_tags or
// machine_tags. The hrefs can belong to various resource types. If a resource
// for a href could not be found, an error is returned and no tags are added for
// any resource.  No error will be raised if the resource already has the tag(s)
// you are trying to add.  Note: At this point, tags on 'next_instance' are not
// supported and one has to add tags to the 'server'.
func (c *Client) MultiAddTags(resourceHrefs []string, tags []string) error {
	payload := map[string]interface{}{
		"resource_hrefs": resourceHrefs,
		"tags":           tags}
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
func (c *Client) MultiAddTagsG(p *Params) error {
	resourceHrefs := (*p)["resourceHrefs"].([]string)
	tags := (*p)["tags"].([]string)
	return c.MultiAddTags(resourceHrefs, tags)
}

// POST api/tags/multi_delete(.:format)?
// Delete a list of tags on a list of hrefs. The tags must be either plain_tags
// or machine_tags. The hrefs can belong to various resource types. If a
// resource for a href could not be found, an error is returned and no tags are
// deleted for any resource.  Note that no error will be raised if the resource
// does not have the tag(s) you are trying to delete.
func (c *Client) MultiDeleteTags(resourceHrefs []string, tags []string) error {
	payload := map[string]interface{}{
		"resource_hrefs": resourceHrefs,
		"tags":           tags}
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
func (c *Client) MultiDeleteTagsG(p *Params) error {
	resourceHrefs := (*p)["resourceHrefs"].([]string)
	tags := (*p)["tags"].([]string)
	return c.MultiDeleteTags(resourceHrefs, tags)
}

// == Task ==

// Tasks represent processes that happen (or have happened) asynchronously
// within the context of an Instance.  An example of a type of task is an
// operational script that runs in an instance.  Task resources can be returned
// by certain API calls, such as Instances.run_executable, Backups.restore, and
// others.
type Task struct {
	Actions []string            `json:"actions,omitempty"`
	Detail  string              `json:"detail,omitempty"`
	Links   []map[string]string `json:"links,omitempty"`
	Summary string              `json:"summary,omitempty"`
}

// GET api/clouds/:cloud_id/instances/:instance_id/live/tasks/:id(.:format)?
// Displays information of a given task within the context of an instance.
func (c *Client) ShowTask(view string, cloudId string, instanceId string, id string) (*Task, error) {
	var res *Task
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/live/tasks/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowTaskG(p *Params) (*Task, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.ShowTask(view, cloudId, instanceId, id)
}

// == UserData ==

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
func (c *Client) ShowUserDataG(p *Params) (*map[string]string, error) {
	return c.ShowUserData()
}

// == User ==

// A User represents an individual RightScale user. Users must be given access
// to RightScale accounts in order to  access RightScale resources.
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
// Create a user. If a user already exists with the same email, that user will
// be returned.  Creating a user alone will not enable the user to access this
// account. You have to create 'permissions' for that user before it can be
// used. Performing a 'show' on a new user will fail unless you immediately
// create an 'observer' permission on the current account.  Note that
// information about users and their permissions must be propagated globally
// across all RightScale clusters, and this can take some time (less than 60
// seconds under normal circumstances) so the users you create may not be able
// to login for a minute or two after you create them. However, you may create
// or destroy permissions for newly-created users with no delay.  To create a
// user that will login using password authentication, include the 'password'
// parameter with your request.  To create an SSO-enabled user, you must specify
// the identity_provider that will be vouching for this user's identity, as well
// as the principal_uid (SAML NameID or OpenID identity URL) that the identity
// provider will assert for this user. Identity providers should be specified
// by their API href; you can obtain a list of the identity providers available
// to your account by invoking the 'index' action of the identity_providers API
// resource.
func (c *Client) CreateUser(userPrincipalUid string, userTimezoneName string, userFirstName string, userLastName string, userPassword string, userPhone string, userCompany string, userEmail string, userIdentityProviderHref string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"user[principal_uid]":          userPrincipalUid,
		"user[timezone_name]":          userTimezoneName,
		"user[first_name]":             userFirstName,
		"user[last_name]":              userLastName,
		"user[password]":               userPassword,
		"user[phone]":                  userPhone,
		"user[company]":                userCompany,
		"user[email]":                  userEmail,
		"user[identity_provider_href]": userIdentityProviderHref}
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
func (c *Client) CreateUserG(p *Params) (Href, error) {
	userPrincipalUid := (*p)["userPrincipalUid"].(string)
	userTimezoneName := (*p)["userTimezoneName"].(string)
	userFirstName := (*p)["userFirstName"].(string)
	userLastName := (*p)["userLastName"].(string)
	userPassword := (*p)["userPassword"].(string)
	userPhone := (*p)["userPhone"].(string)
	userCompany := (*p)["userCompany"].(string)
	userEmail := (*p)["userEmail"].(string)
	userIdentityProviderHref := (*p)["userIdentityProviderHref"].(string)
	return c.CreateUser(userPrincipalUid, userTimezoneName, userFirstName, userLastName, userPassword, userPhone, userCompany, userEmail, userIdentityProviderHref)
}

// GET api/users(.:format)?
// List the users available to the account the user is logged in to. Therefore,
// to list the users of a child account, the user has to login to the child
// account first.
func (c *Client) IndexUsers(filter []string) ([]User, error) {
	var res []User
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/users", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
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
func (c *Client) IndexUsersG(p *Params) ([]User, error) {
	filter := (*p)["filter"].([]string)
	return c.IndexUsers(filter)
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
func (c *Client) ShowUserG(p *Params) (*User, error) {
	id := (*p)["id"].(string)
	return c.ShowUser(id)
}

// PUT api/users/:id(.:format)?
// Update a user's contact information, change her password, or update SSO her
// settings. In order to update a user record, one of the following criteria
// must be met:   You're logged in AS the user being modified and you provide
// a valid current_password. You're an admin and the user is linked to your
// enterprise SSO provider You're an admin and the user's email matches the
// email_domain of your enterprise SSO provider   In other words: you can update
// yourself if you know your own password; you can update yourself or others
// if they're linked to your SSO providers, and you can update any user if
// her email address is known to belong to your organization.  For information
// about enabling canonical email domain ownership for your enterprise,
// please talk to your RightScale account manager or contact our support
// team.  To update a user's contact information, simply pass the desired
// values for email, first_name, and so forth.  To update a user's password,
// provide a valid current_password and specify the desired new_password.
// To update a user's SSO information, you may provide a just a principal_uid
// (to maintain the user's existing identity provider) or you may provide an
// identity_provider_href and a principal_uid (to switch identity providers as
// well as specify a new user identity).  In the context of SAML. principal_uid
// is equivalent to the SAML NameID or Subject claim; RightScale cannot predict
// or influence the NameID value that your SAML IdP will send to us for
func (c *Client) UpdateUser(userPhone string, userCurrentPassword string, userFirstName string, userIdentityProviderHref string, userLastName string, userNewEmail string, userNewPassword string, userCompany string, userCurrentEmail string, userPrincipalUid string, userTimezoneName string, id string) error {
	payload := map[string]interface{}{
		"user[phone]":                  userPhone,
		"user[current_password]":       userCurrentPassword,
		"user[first_name]":             userFirstName,
		"user[identity_provider_href]": userIdentityProviderHref,
		"user[last_name]":              userLastName,
		"user[new_email]":              userNewEmail,
		"user[new_password]":           userNewPassword,
		"user[company]":                userCompany,
		"user[current_email]":          userCurrentEmail,
		"user[principal_uid]":          userPrincipalUid,
		"user[timezone_name]":          userTimezoneName}
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
func (c *Client) UpdateUserG(p *Params) error {
	userPhone := (*p)["userPhone"].(string)
	userCurrentPassword := (*p)["userCurrentPassword"].(string)
	userFirstName := (*p)["userFirstName"].(string)
	userIdentityProviderHref := (*p)["userIdentityProviderHref"].(string)
	userLastName := (*p)["userLastName"].(string)
	userNewEmail := (*p)["userNewEmail"].(string)
	userNewPassword := (*p)["userNewPassword"].(string)
	userCompany := (*p)["userCompany"].(string)
	userCurrentEmail := (*p)["userCurrentEmail"].(string)
	userPrincipalUid := (*p)["userPrincipalUid"].(string)
	userTimezoneName := (*p)["userTimezoneName"].(string)
	id := (*p)["id"].(string)
	return c.UpdateUser(userPhone, userCurrentPassword, userFirstName, userIdentityProviderHref, userLastName, userNewEmail, userNewPassword, userCompany, userCurrentEmail, userPrincipalUid, userTimezoneName, id)
}

// == VolumeAttachment ==

// A VolumeAttachment represents a relationship between a volume and an
// instance.
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
func (c *Client) CreateVolumeAttachment(volumeAttachmentVolumeHref string, volumeAttachmentDevice string, volumeAttachmentInstanceHref string, cloudId string, instanceId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"volume_attachment[volume_href]":   volumeAttachmentVolumeHref,
		"volume_attachment[device]":        volumeAttachmentDevice,
		"volume_attachment[instance_href]": volumeAttachmentInstanceHref}
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
func (c *Client) CreateVolumeAttachmentG(p *Params) (Href, error) {
	volumeAttachmentVolumeHref := (*p)["volumeAttachmentVolumeHref"].(string)
	volumeAttachmentDevice := (*p)["volumeAttachmentDevice"].(string)
	volumeAttachmentInstanceHref := (*p)["volumeAttachmentInstanceHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.CreateVolumeAttachment(volumeAttachmentVolumeHref, volumeAttachmentDevice, volumeAttachmentInstanceHref, cloudId, instanceId)
}

// DELETE api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id(.:format)?
// Deletes a given volume attachment.
func (c *Client) DestroyVolumeAttachment(force string, cloudId string, instanceId string, id string) error {
	payload := map[string]interface{}{
		"force": force}
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
func (c *Client) DestroyVolumeAttachmentG(p *Params) error {
	force := (*p)["force"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyVolumeAttachment(force, cloudId, instanceId, id)
}

// GET api/clouds/:cloud_id/instances/:instance_id/volume_attachments(.:format)?
// Lists all volume attachments.
func (c *Client) IndexVolumeAttachments(filter []string, view string, cloudId string, instanceId string) ([]VolumeAttachment, error) {
	var res []VolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/volume_attachments", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexVolumeAttachmentsG(p *Params) ([]VolumeAttachment, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	return c.IndexVolumeAttachments(filter, view, cloudId, instanceId)
}

// GET api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id(.:format)?
// Displays information about a single volume attachment.
func (c *Client) ShowVolumeAttachment(view string, cloudId string, instanceId string, id string) (*VolumeAttachment, error) {
	var res *VolumeAttachment
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/instances/"+instanceId+"/volume_attachments/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowVolumeAttachmentG(p *Params) (*VolumeAttachment, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	instanceId := (*p)["instanceId"].(string)
	id := (*p)["id"].(string)
	return c.ShowVolumeAttachment(view, cloudId, instanceId, id)
}

// == VolumeSnapshot ==

// A VolumeSnapshot represents a Cloud storage volume at a particular point
// in time. One can create a snapshot regardless of whether or not a volume
// is attached to an Instance. When a snapshot is created, various meta data
// is retained such as a Created At timestamp, a unique Resource UID (e.g.
// vol-52EF05A9), the Volume Owner and Visibility (e.g. private or public).
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
func (c *Client) CreateVolumeSnapshot(volumeSnapshotName string, volumeSnapshotParentVolumeHref string, volumeSnapshotDeploymentHref string, volumeSnapshotDescription string, cloudId string, volumeId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"volume_snapshot[name]":               volumeSnapshotName,
		"volume_snapshot[parent_volume_href]": volumeSnapshotParentVolumeHref,
		"volume_snapshot[deployment_href]":    volumeSnapshotDeploymentHref,
		"volume_snapshot[description]":        volumeSnapshotDescription}
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
func (c *Client) CreateVolumeSnapshotG(p *Params) (Href, error) {
	volumeSnapshotName := (*p)["volumeSnapshotName"].(string)
	volumeSnapshotParentVolumeHref := (*p)["volumeSnapshotParentVolumeHref"].(string)
	volumeSnapshotDeploymentHref := (*p)["volumeSnapshotDeploymentHref"].(string)
	volumeSnapshotDescription := (*p)["volumeSnapshotDescription"].(string)
	cloudId := (*p)["cloudId"].(string)
	volumeId := (*p)["volumeId"].(string)
	return c.CreateVolumeSnapshot(volumeSnapshotName, volumeSnapshotParentVolumeHref, volumeSnapshotDeploymentHref, volumeSnapshotDescription, cloudId, volumeId)
}

// DELETE api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id(.:format)?
// Deletes a given volume_snapshot.
func (c *Client) DestroyVolumeSnapshot(cloudId string, volumeId string, id string) error {
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
func (c *Client) DestroyVolumeSnapshotG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	volumeId := (*p)["volumeId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyVolumeSnapshot(cloudId, volumeId, id)
}

// GET api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots(.:format)?
// Lists all volume_snapshots.
func (c *Client) IndexVolumeSnapshots(filter []string, view string, cloudId string, volumeId string) ([]VolumeSnapshot, error) {
	var res []VolumeSnapshot
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+volumeId+"/volume_snapshots", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexVolumeSnapshotsG(p *Params) ([]VolumeSnapshot, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	volumeId := (*p)["volumeId"].(string)
	return c.IndexVolumeSnapshots(filter, view, cloudId, volumeId)
}

// GET api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id(.:format)?
// Displays information about a single volume_snapshot.
func (c *Client) ShowVolumeSnapshot(view string, cloudId string, volumeId string, id string) (*VolumeSnapshot, error) {
	var res *VolumeSnapshot
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+volumeId+"/volume_snapshots/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowVolumeSnapshotG(p *Params) (*VolumeSnapshot, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	volumeId := (*p)["volumeId"].(string)
	id := (*p)["id"].(string)
	return c.ShowVolumeSnapshot(view, cloudId, volumeId, id)
}

// == VolumeType ==

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
func (c *Client) IndexVolumeTypes(filter []string, view string, cloudId string) ([]VolumeType, error) {
	var res []VolumeType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volume_types", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexVolumeTypesG(p *Params) ([]VolumeType, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexVolumeTypes(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/volume_types/:id(.:format)?
// Displays information about a single Volume Type.
func (c *Client) ShowVolumeType(view string, cloudId string, id string) (*VolumeType, error) {
	var res *VolumeType
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volume_types/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowVolumeTypeG(p *Params) (*VolumeType, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowVolumeType(view, cloudId, id)
}

// == Volume ==

// A Volume provides a highly reliable, efficient and persistent storage
// solution that can be mounted to a cloud instance (in the same datacenter /
// zone).
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
func (c *Client) CreateVolume(volumeName string, volumeParentVolumeSnapshotHref string, volumeEncrypted string, volumeIops string, volumeDeploymentHref string, volumeDescription string, volumePlacementGroupHref string, volumeSize string, volumeVolumeTypeHref string, volumeDatacenterHref string, cloudId string) (Href, error) {
	var res Href
	payload := map[string]interface{}{
		"volume[name]":                        volumeName,
		"volume[parent_volume_snapshot_href]": volumeParentVolumeSnapshotHref,
		"volume[encrypted]":                   volumeEncrypted,
		"volume[iops]":                        volumeIops,
		"volume[deployment_href]":             volumeDeploymentHref,
		"volume[description]":                 volumeDescription,
		"volume[placement_group_href]":        volumePlacementGroupHref,
		"volume[size]":                        volumeSize,
		"volume[volume_type_href]":            volumeVolumeTypeHref,
		"volume[datacenter_href]":             volumeDatacenterHref}
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
func (c *Client) CreateVolumeG(p *Params) (Href, error) {
	volumeName := (*p)["volumeName"].(string)
	volumeParentVolumeSnapshotHref := (*p)["volumeParentVolumeSnapshotHref"].(string)
	volumeEncrypted := (*p)["volumeEncrypted"].(string)
	volumeIops := (*p)["volumeIops"].(string)
	volumeDeploymentHref := (*p)["volumeDeploymentHref"].(string)
	volumeDescription := (*p)["volumeDescription"].(string)
	volumePlacementGroupHref := (*p)["volumePlacementGroupHref"].(string)
	volumeSize := (*p)["volumeSize"].(string)
	volumeVolumeTypeHref := (*p)["volumeVolumeTypeHref"].(string)
	volumeDatacenterHref := (*p)["volumeDatacenterHref"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.CreateVolume(volumeName, volumeParentVolumeSnapshotHref, volumeEncrypted, volumeIops, volumeDeploymentHref, volumeDescription, volumePlacementGroupHref, volumeSize, volumeVolumeTypeHref, volumeDatacenterHref, cloudId)
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
func (c *Client) DestroyVolumeG(p *Params) error {
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.DestroyVolume(cloudId, id)
}

// GET api/clouds/:cloud_id/volumes(.:format)?
// Lists volumes.
func (c *Client) IndexVolumes(filter []string, view string, cloudId string) ([]Volume, error) {
	var res []Volume
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes", body)
	if err != nil {
		return res, err
	}
	for _, v := range filter {
		req.URL.Query().Add("filter", v)
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) IndexVolumesG(p *Params) ([]Volume, error) {
	filter := (*p)["filter"].([]string)
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	return c.IndexVolumes(filter, view, cloudId)
}

// GET api/clouds/:cloud_id/volumes/:id(.:format)?
// Displays information about a single volume.
func (c *Client) ShowVolume(view string, cloudId string, id string) (*Volume, error) {
	var res *Volume
	b := []byte{}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("GET", c.endpoint+"/api/clouds/"+cloudId+"/volumes/"+id, body)
	if err != nil {
		return res, err
	}
	req.URL.Query().Set("view", view)
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
func (c *Client) ShowVolumeG(p *Params) (*Volume, error) {
	view := (*p)["view"].(string)
	cloudId := (*p)["cloudId"].(string)
	id := (*p)["id"].(string)
	return c.ShowVolume(view, cloudId, id)
}
