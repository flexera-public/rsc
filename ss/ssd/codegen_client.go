//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=ss/ssd/restful_doc -output=ss/ssd -pkg=ssd -target=1.0 -client=API
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
func (r *Schedule) Locator(api *API) *ScheduleLocator {
	return api.ScheduleLocator(r.Href)
}

//===== Locator

// ScheduleLocator exposes the Schedule resource actions.
type ScheduleLocator struct {
	Href
	api *API
}

// ScheduleLocator builds a locator from the given href.
func (api *API) ScheduleLocator(href string) *ScheduleLocator {
	return &ScheduleLocator{Href(href), api}
}

//===== Actions

// GET /api/designer/collections/:collection_id/schedules
//
// List the schedules available in Designer.
func (loc *ScheduleLocator) Index() ([]*Schedule, error) {
	var res []*Schedule
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Schedule", "index")
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

// GET /api/designer/collections/:collection_id/schedules/:id
//
// Show detailed information about a given Schedule.
func (loc *ScheduleLocator) Show() (*Schedule, error) {
	var res *Schedule
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Schedule", "show")
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

// POST /api/designer/collections/:collection_id/schedules
//
// Create a new Schedule.
func (loc *ScheduleLocator) Create(name string, startRecurrence *Recurrence, stopRecurrence *Recurrence, options rsapi.APIParams) (*ScheduleLocator, error) {
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
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name":             name,
		"start_recurrence": startRecurrence,
		"stop_recurrence":  stopRecurrence,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("Schedule", "create")
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
		return &ScheduleLocator{Href(location), loc.api}, nil
	}
}

// PATCH /api/designer/collections/:collection_id/schedules/:id
//
// Update one or more attributes of an existing Schedule.
// Note: updating a Schedule in Designer doesn't update it in the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.
func (loc *ScheduleLocator) Update(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var startRecurrenceOpt = options["start_recurrence"]
	if startRecurrenceOpt != nil {
		p["start_recurrence"] = startRecurrenceOpt
	}
	var stopRecurrenceOpt = options["stop_recurrence"]
	if stopRecurrenceOpt != nil {
		p["stop_recurrence"] = stopRecurrenceOpt
	}
	uri, err := loc.ActionPath("Schedule", "update")
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

// DELETE /api/designer/collections/:collection_id/schedules/:id
//
// Delete a Schedule from the system.
// Note: deleting a Schedule from Designer doesn't remove it from the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.
func (loc *ScheduleLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Schedule", "delete")
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

// DELETE /api/designer/collections/:collection_id/schedules
//
// Delete multiple Schedules from the system in bulk.
// Note: deleting a Schedule from Designer doesn't remove it from the applications that were published with it to the Catalog or affect running CloudApps with that Schedule.
func (loc *ScheduleLocator) MultiDelete(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Schedule", "multi_delete")
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

/******  Template ******/

// A Template represent a CloudApplication Template (CAT) that has been uploaded to this design collection.
// For information on the syntax of a CAT file, please see the [CAT File Language Reference](http://docs.rightscale.com/ss/reference/ss_CAT_file_language.html) on the RightScale Docs
// site.
// A CAT file is compiled by Self-Service to make it ready for publication and subsequent launch by users. To
// test your CAT file syntax, you can call the compile action with the source content. In order to
// Publish your CAT to the Catalog where users can launch it, it must be uploaded to Designer first, and then
// published to the Catalog.
// CAT files are uniquely identified by the name of the CloudApplication, which is specified as the "name"
// attribute inside of a CAT file.
type Template struct {
	ApplicationInfo    *ApplicationInfo  `json:"application_info,omitempty"`
	CompilationHref    string            `json:"compilation_href,omitempty"`
	CompiledCat        string            `json:"compiled_cat,omitempty"`
	CreatedBy          *User             `json:"created_by,omitempty"`
	Dependencies       []*CatDependency  `json:"dependencies,omitempty"`
	Dependents         []*CatDependency  `json:"dependents,omitempty"`
	Filename           string            `json:"filename,omitempty"`
	Href               string            `json:"href,omitempty"`
	Id                 string            `json:"id,omitempty"`
	Imports            []string          `json:"imports,omitempty"`
	Kind               string            `json:"kind,omitempty"`
	LongDescription    string            `json:"long_description,omitempty"`
	Name               string            `json:"name,omitempty"`
	Package            string            `json:"package,omitempty"`
	Parameters         []*Parameter      `json:"parameters,omitempty"`
	PublishedBy        *User             `json:"published_by,omitempty"`
	RequiredParameters []string          `json:"required_parameters,omitempty"`
	RsCaVer            int               `json:"rs_ca_ver,omitempty"`
	ShortDescription   string            `json:"short_description,omitempty"`
	Source             string            `json:"source,omitempty"`
	SourceHref         string            `json:"source_href,omitempty"`
	Stale              bool              `json:"stale,omitempty"`
	Timestamps         *TimestampsStruct `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Template) Locator(api *API) *TemplateLocator {
	return api.TemplateLocator(r.Href)
}

//===== Locator

// TemplateLocator exposes the Template resource actions.
type TemplateLocator struct {
	Href
	api *API
}

// TemplateLocator builds a locator from the given href.
func (api *API) TemplateLocator(href string) *TemplateLocator {
	return &TemplateLocator{Href(href), api}
}

//===== Actions

// GET /api/designer/collections/:collection_id/templates
//
// List the templates available in Designer along with some general details.
func (loc *TemplateLocator) Index(options rsapi.APIParams) ([]*Template, error) {
	var res []*Template
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids[]"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Template", "index")
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

// GET /api/designer/collections/:collection_id/templates/:id
//
// Show detailed information about a given Template. Use the views specified below for more information.
func (loc *TemplateLocator) Show(options rsapi.APIParams) (*Template, error) {
	var res *Template
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Template", "show")
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

// POST /api/designer/collections/:collection_id/templates
//
// Create a new Template by uploading its content to Designer.
func (loc *TemplateLocator) Create(source *rsapi.FileUpload) (*TemplateLocator, error) {
	var res *TemplateLocator
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"source": source,
	}
	uri, err := loc.ActionPath("Template", "create")
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
		return &TemplateLocator{Href(location), loc.api}, nil
	}
}

// POST /api/designer/collections/:collection_id/templates/actions/create_from_compilation
//
// Create a new Template from a previously compiled CAT.
func (loc *TemplateLocator) CreateFromCompilation(compilationHref string, options rsapi.APIParams) error {
	if compilationHref == "" {
		return fmt.Errorf("compilationHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"compilation_href": compilationHref,
	}
	var filenameOpt = options["filename"]
	if filenameOpt != nil {
		p["filename"] = filenameOpt
	}
	uri, err := loc.ActionPath("Template", "create_from_compilation")
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

// PUT /api/designer/collections/:collection_id/templates/:id
//
// Update the content of an existing Template (a Template with the same "name" value in the CAT).
func (loc *TemplateLocator) Update(source *rsapi.FileUpload) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"source": source,
	}
	uri, err := loc.ActionPath("Template", "update")
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

// POST /api/designer/collections/:collection_id/templates/:id/actions/update_from_compilation
//
// Update a Template from a previously compiled CAT.
func (loc *TemplateLocator) UpdateFromCompilation(compilationHref string, options rsapi.APIParams) error {
	if compilationHref == "" {
		return fmt.Errorf("compilationHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"compilation_href": compilationHref,
	}
	var filenameOpt = options["filename"]
	if filenameOpt != nil {
		p["filename"] = filenameOpt
	}
	uri, err := loc.ActionPath("Template", "update_from_compilation")
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

// DELETE /api/designer/collections/:collection_id/templates/:id
//
// Delete a Template from the system. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.
func (loc *TemplateLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Template", "delete")
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

// DELETE /api/designer/collections/:collection_id/templates
//
// Delete multiple Templates from the system in bulk. Note: deleting a Template from Designer doesn't remove it from the Catalog if it has already been published -- see the "unpublish" action.
func (loc *TemplateLocator) MultiDelete(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"ids[]": ids,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Template", "multi_delete")
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

// GET /api/designer/collections/:collection_id/templates/:id/download
//
// Download the source of a Template.
func (loc *TemplateLocator) Download(apiVersion string) error {
	if apiVersion == "" {
		return fmt.Errorf("apiVersion is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"api_version": apiVersion,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Template", "download")
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

// POST /api/designer/collections/:collection_id/templates/actions/compile
//
// Compile the Template, but don't save it to Designer. Useful for debugging a CAT file while you are still authoring it.
func (loc *TemplateLocator) Compile(source string) error {
	if source == "" {
		return fmt.Errorf("source is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"source": source,
	}
	uri, err := loc.ActionPath("Template", "compile")
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

// POST /api/designer/collections/:collection_id/templates/actions/dependencies
//
// Lists the Templates which the provided CAT source or Template directly or indirectly depend upon
func (loc *TemplateLocator) Dependencies(options rsapi.APIParams) (*Template, error) {
	var res *Template
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var sourceOpt = options["source"]
	if sourceOpt != nil {
		p["source"] = sourceOpt
	}
	var templateIdOpt = options["template_id"]
	if templateIdOpt != nil {
		p["template_id"] = templateIdOpt
	}
	uri, err := loc.ActionPath("Template", "dependencies")
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

// GET /api/designer/collections/:collection_id/templates/actions/dependents
//
// List the Dependents templates available in Designer for the given package, even if no template actually define the package.
func (loc *TemplateLocator) Dependents(package_ string, options rsapi.APIParams) (*Template, error) {
	var res *Template
	if package_ == "" {
		return res, fmt.Errorf("package is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"package": package_,
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Template", "dependents")
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

// POST /api/designer/collections/:collection_id/templates/actions/publish
//
// Publish the given Template to the Catalog so that users can launch it.
func (loc *TemplateLocator) Publish(id string, options rsapi.APIParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"id": id,
	}
	var longDescriptionOpt = options["long_description"]
	if longDescriptionOpt != nil {
		p["long_description"] = longDescriptionOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var overriddenApplicationHrefOpt = options["overridden_application_href"]
	if overriddenApplicationHrefOpt != nil {
		p["overridden_application_href"] = overriddenApplicationHrefOpt
	}
	var scheduleRequiredOpt = options["schedule_required"]
	if scheduleRequiredOpt != nil {
		p["schedule_required"] = scheduleRequiredOpt
	}
	var schedulesOpt = options["schedules"]
	if schedulesOpt != nil {
		p["schedules"] = schedulesOpt
	}
	var shortDescriptionOpt = options["short_description"]
	if shortDescriptionOpt != nil {
		p["short_description"] = shortDescriptionOpt
	}
	uri, err := loc.ActionPath("Template", "publish")
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

// POST /api/designer/collections/:collection_id/templates/actions/unpublish
//
// Remove a publication from the Catalog by specifying its associated Template.
func (loc *TemplateLocator) Unpublish(id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"id": id,
	}
	uri, err := loc.ActionPath("Template", "unpublish")
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

/****** Parameter Data Types ******/

type ApplicationInfo struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

type CatDependency struct {
	Alias        string     `json:"alias,omitempty"`
	CompiledAt   *time.Time `json:"compiled_at,omitempty"`
	Name         string     `json:"name,omitempty"`
	Package      string     `json:"package,omitempty"`
	RsCaVer      int        `json:"rs_ca_ver,omitempty"`
	SourceHref   string     `json:"source_href,omitempty"`
	TemplateHref string     `json:"template_href,omitempty"`
	TemplateId   string     `json:"template_id,omitempty"`
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
	CompiledAt  *time.Time `json:"compiled_at,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}
