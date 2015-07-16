//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=ss/ssd/restful_doc -output=ss/ssd -pkg=ssd -target=1.0 -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

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
	return action.Url(vars)
}

/******  Schedule ******/

// A Schedule represents a recurring period during which a CloudApp should be running. It must have a unique name and an optional description. The recurrence rules follow the [Recurrence Rule format](https://tools.ietf.org/html/rfc5545#section-3.8.5.3).
// Multiple Schedules can be associated with a Template when published to the Catalog. Users will be able to launch the resulting CloudApp with one of the associated schedule. Updating or deleting a Schedule will not affect CloudApps that were published with that Schedule.
type Schedule struct {
	CreatedBy       *User             `json:"created_by,omitempty"`
	Description     string            `json:"description,omitempty"`
	Href            string            `json:"href,omitempty"`
	Id              string            `json:"id,omitempty"`
	Kind            string            `json:"kind,omitempty"`
	Name            string            `json:"name,omitempty"`
	StartRecurrence *Recurrence       `json:"start_recurrence,omitempty"`
	StopRecurrence  *Recurrence       `json:"stop_recurrence,omitempty"`
	Timestamps      *TimestampsStruct `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Schedule) Locator(api *Api) *ScheduleLocator {
	return api.ScheduleLocator(r.Href)
}

//===== Locator

// ScheduleLocator exposes the Schedule resource actions.
type ScheduleLocator struct {
	Href
	api *Api
}

// ScheduleLocator builds a locator from the given href.
func (api *Api) ScheduleLocator(href string) *ScheduleLocator {
	return &ScheduleLocator{Href(href), api}
}

//===== Actions

// GET /collections/:collection_id/schedules
//
// List the schedules available in Designer.
func (loc *ScheduleLocator) Index() ([]*Schedule, error) {
	var res []*Schedule
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Schedule", "index")
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

// GET /collections/:collection_id/schedules/:id
//
// Show detailed information about a given Schedule.
func (loc *ScheduleLocator) Show() (*Schedule, error) {
	var res *Schedule
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Schedule", "show")
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

// POST /collections/:collection_id/schedules
//
// Create a new Schedule.
func (loc *ScheduleLocator) Create(name string, startRecurrence *Recurrence, stopRecurrence *Recurrence, options rsapi.ApiParams) (*ScheduleLocator, error) {
	var res *ScheduleLocator
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if startRecurrence == nil {
		return res, fmt.Errorf("startRecurrence is required")
	}
	if stopRecurrence == nil {
		return res, fmt.Errorf("stopRecurrence is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"name":             name,
		"start_recurrence": startRecurrence,
		"stop_recurrence":  stopRecurrence,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		payloadParams["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("Schedule", "create")
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
		return &ScheduleLocator{Href(location), loc.api}, nil
	}
}

// PATCH /collections/:collection_id/schedules/:id
//
// Update one or more attributes of an existing Schedule.
// Note: updating a Schedule in Designer doesn't update it in the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.
func (loc *ScheduleLocator) Update(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		payloadParams["description"] = descriptionOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var startRecurrenceOpt = options["start_recurrence"]
	if startRecurrenceOpt != nil {
		payloadParams["start_recurrence"] = startRecurrenceOpt
	}
	var stopRecurrenceOpt = options["stop_recurrence"]
	if stopRecurrenceOpt != nil {
		payloadParams["stop_recurrence"] = stopRecurrenceOpt
	}
	uri, err := loc.ActionPath("Schedule", "update")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /collections/:collection_id/schedules/:id
//
// Delete a Schedule from the system.
// Note: deleting a Schedule from Designer doesn't remove it from the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.
func (loc *ScheduleLocator) Delete() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Schedule", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /collections/:collection_id/schedules
//
// Delete multiple Schedules from the system in bulk.
// Note: deleting a Schedule from Designer doesn't remove it from the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.
func (loc *ScheduleLocator) MultiDelete(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Schedule", "multi_delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  Template ******/

// A Template represent a CloudApplication Template (CAT) that has been uploaded to this design collection.
// For information on the syntax of a CAT file, please see the [CAT Designers Guide](http://support.rightscale.com/12-Guides/Self-Service/20_Cloud_Application_Template_%28CAT%29_Designers_Guide) on the RightScale Support
// site.
// A CAT file is compiled by Self-Service to make it ready for publication and subsequent launch by users. To
// test your CAT file syntax, you can call the compile action with the source content. In order to
// Publish your CAT to the Catalog where users can launch it, it must be uploaded to Designer first, and then
// published to the Catalog.
// CAT files are uniquely identified by the name of the CloudApplication, which is specified as the "name"
// attribute inside of a CAT file.
type Template struct {
	ApplicationInfo    *ApplicationInfo  `json:"application_info,omitempty"`
	CompiledCat        string            `json:"compiled_cat,omitempty"`
	CreatedBy          *User             `json:"created_by,omitempty"`
	Filename           string            `json:"filename,omitempty"`
	Href               string            `json:"href,omitempty"`
	Id                 string            `json:"id,omitempty"`
	Kind               string            `json:"kind,omitempty"`
	LongDescription    string            `json:"long_description,omitempty"`
	Name               string            `json:"name,omitempty"`
	Parameters         []*Parameter      `json:"parameters,omitempty"`
	PublishedBy        *User             `json:"published_by,omitempty"`
	RequiredParameters []string          `json:"required_parameters,omitempty"`
	ShortDescription   string            `json:"short_description,omitempty"`
	Source             string            `json:"source,omitempty"`
	Timestamps         *TimestampsStruct `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Template) Locator(api *Api) *TemplateLocator {
	return api.TemplateLocator(r.Href)
}

//===== Locator

// TemplateLocator exposes the Template resource actions.
type TemplateLocator struct {
	Href
	api *Api
}

// TemplateLocator builds a locator from the given href.
func (api *Api) TemplateLocator(href string) *TemplateLocator {
	return &TemplateLocator{Href(href), api}
}

//===== Actions

// GET /collections/:collection_id/templates
//
// List the templates available in Designer along with some general details.
func (loc *TemplateLocator) Index(options rsapi.ApiParams) ([]*Template, error) {
	var res []*Template
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids[]"] = idsOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Template", "index")
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

// GET /collections/:collection_id/templates/:id
//
// Show detailed information about a given Template. Use the views specified below for more information.
func (loc *TemplateLocator) Show(options rsapi.ApiParams) (*Template, error) {
	var res *Template
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Template", "show")
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

// POST /collections/:collection_id/templates
//
// Create a new Template by uploading its content to Designer.
func (loc *TemplateLocator) Create(source *FileUpload) (*TemplateLocator, error) {
	var res *TemplateLocator
	if source == nil {
		return res, fmt.Errorf("source is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"source": source,
	}
	uri, err := loc.ActionPath("Template", "create")
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
		return &TemplateLocator{Href(location), loc.api}, nil
	}
}

// PUT /collections/:collection_id/templates/:id
//
// Update the content of an existing Template (a Template with the same "name" value in the CAT).
func (loc *TemplateLocator) Update(source *FileUpload) error {
	if source == nil {
		return fmt.Errorf("source is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"source": source,
	}
	uri, err := loc.ActionPath("Template", "update")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /collections/:collection_id/templates/:id
//
// Delete a Template from the system. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.
func (loc *TemplateLocator) Delete() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Template", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// DELETE /collections/:collection_id/templates
//
// Delete multiple Templates from the system in bulk. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.
func (loc *TemplateLocator) MultiDelete(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"ids[]": ids,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Template", "multi_delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// GET /collections/:collection_id/templates/:id/download
//
// Download the source of a Template.
func (loc *TemplateLocator) Download(apiVersion string) error {
	if apiVersion == "" {
		return fmt.Errorf("apiVersion is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"api_version": apiVersion,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Template", "download")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /collections/:collection_id/templates/actions/compile
//
// Compile the Template, but don't save it to Designer. Useful for debugging a CAT file while you are still authoring it.
func (loc *TemplateLocator) Compile(source string) error {
	if source == "" {
		return fmt.Errorf("source is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"source": source,
	}
	uri, err := loc.ActionPath("Template", "compile")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /collections/:collection_id/templates/actions/publish
//
// Publish the given Template to the Catalog so that users can launch it.
func (loc *TemplateLocator) Publish(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"id": id,
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		payloadParams["name"] = nameOpt
	}
	var overriddenApplicationHrefOpt = options["overridden_application_href"]
	if overriddenApplicationHrefOpt != nil {
		payloadParams["overridden_application_href"] = overriddenApplicationHrefOpt
	}
	var scheduleRequiredOpt = options["schedule_required"]
	if scheduleRequiredOpt != nil {
		payloadParams["schedule_required"] = scheduleRequiredOpt
	}
	var schedulesOpt = options["schedules"]
	if schedulesOpt != nil {
		payloadParams["schedules"] = schedulesOpt
	}
	var shortDescriptionOpt = options["short_description"]
	if shortDescriptionOpt != nil {
		payloadParams["short_description"] = shortDescriptionOpt
	}
	uri, err := loc.ActionPath("Template", "publish")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

// POST /collections/:collection_id/templates/actions/unpublish
//
// Remove a publication from the Catalog by specifying its associated Template.
func (loc *TemplateLocator) Unpublish(id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"id": id,
	}
	uri, err := loc.ActionPath("Template", "unpublish")
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

type ApplicationInfo struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

type FileUpload struct {
	Filename string `json:"filename,omitempty"`
	Head     string `json:"head,omitempty"`
	Name     string `json:"name,omitempty"`
	Tempfile string `json:"tempfile,omitempty"`
	Type_    string `json:"type,omitempty"`
}

type Parameter struct {
	Default     interface{}                 `json:"default,omitempty"`
	Description string                      `json:"description,omitempty"`
	Name        string                      `json:"name,omitempty"`
	Operations  []interface{}               `json:"operations,omitempty"`
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
	AllowedPattern        string        `json:"allowed_pattern,omitempty"`
	AllowedValues         []interface{} `json:"allowed_values,omitempty"`
	ConstraintDescription string        `json:"constraint_description,omitempty"`
	MaxLength             int           `json:"max_length,omitempty"`
	MaxValue              int           `json:"max_value,omitempty"`
	MinLength             int           `json:"min_length,omitempty"`
	MinValue              int           `json:"min_value,omitempty"`
	NoEcho                bool          `json:"no_echo,omitempty"`
}

type Recurrence struct {
	Hour   int    `json:"hour,omitempty"`
	Minute int    `json:"minute,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type TimestampsStruct struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type TimestampsStruct2 struct {
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}
