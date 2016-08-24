//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=rl10/docs/api -output=rl10 -pkg=rl10 -target=unversioned -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rl10

import (
	"fmt"
	"io/ioutil"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// API Version
const APIVersion = "unversioned"

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

/******  DebugCookbookPath ******/

// Manipulate debug cookbook directory location
type DebugCookbookPath struct {
}

//===== Locator

// DebugCookbookPathLocator exposes the DebugCookbookPath resource actions.
type DebugCookbookPathLocator struct {
	Href
	api *API
}

// DebugCookbookPathLocator builds a locator from the given href.
func (api *API) DebugCookbookPathLocator(href string) *DebugCookbookPathLocator {
	return &DebugCookbookPathLocator{Href(href), api}
}

//===== Actions

// GET /rll/debug/cookbook
//
// Retrieve debug cookbook directory location
func (loc *DebugCookbookPathLocator) Show() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("DebugCookbookPath", "show")
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

// PUT /rll/debug/cookbook
//
// Set debug cookbook directory location
func (loc *DebugCookbookPathLocator) Update(path string) (string, error) {
	var res string
	if path == "" {
		return res, fmt.Errorf("path is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"path": path,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("DebugCookbookPath", "update")
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

// DELETE /rll/debug/cookbook
//
// Remove debug cookbook directory location
func (loc *DebugCookbookPathLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("DebugCookbookPath", "delete")
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

/******  DockerControl ******/

// Manipulate the Docker integration in RightLink 10
type DockerControl struct {
}

//===== Locator

// DockerControlLocator exposes the DockerControl resource actions.
type DockerControlLocator struct {
	Href
	api *API
}

// DockerControlLocator builds a locator from the given href.
func (api *API) DockerControlLocator(href string) *DockerControlLocator {
	return &DockerControlLocator{Href(href), api}
}

//===== Actions

// GET /rll/docker/control
//
// Show Docker integration features
func (loc *DockerControlLocator) Show() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("DockerControl", "show")
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

// PUT /rll/docker/control
//
// Enable/disable Docker integration features
func (loc *DockerControlLocator) Update(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var dockerHostOpt = options["docker_host"]
	if dockerHostOpt != nil {
		params["docker_host"] = dockerHostOpt
	}
	var enableDockerOpt = options["enable_docker"]
	if enableDockerOpt != nil {
		params["enable_docker"] = enableDockerOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("DockerControl", "update")
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

/******  Env ******/

// Manipulate global script environment variables
type Env struct {
}

//===== Locator

// EnvLocator exposes the Env resource actions.
type EnvLocator struct {
	Href
	api *API
}

// EnvLocator builds a locator from the given href.
func (api *API) EnvLocator(href string) *EnvLocator {
	return &EnvLocator{Href(href), api}
}

//===== Actions

// GET /rll/env
//
// Retrieve all environment variables
func (loc *EnvLocator) Index() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Env", "index")
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

// GET /rll/env/:name
//
// Retrieve environment variable value
func (loc *EnvLocator) Show() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Env", "show")
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

// PUT /rll/env/:name
//
// Set environment variable value
func (loc *EnvLocator) Update(payload string) (string, error) {
	var res string
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("Env", "update")
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

// DELETE /rll/env/:name
//
// Delete environment variable
func (loc *EnvLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Env", "delete")
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

/******  LoginControl ******/

// Manipulate login policy settings
type LoginControl struct {
}

//===== Locator

// LoginControlLocator exposes the LoginControl resource actions.
type LoginControlLocator struct {
	Href
	api *API
}

// LoginControlLocator builds a locator from the given href.
func (api *API) LoginControlLocator(href string) *LoginControlLocator {
	return &LoginControlLocator{Href(href), api}
}

//===== Actions

// GET /rll/login/control
//
// Show login policy features
func (loc *LoginControlLocator) Show() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("LoginControl", "show")
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

// PUT /rll/login/control
//
// Enable/disable login policy features
func (loc *LoginControlLocator) Update(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var enableLoginOpt = options["enable_login"]
	if enableLoginOpt != nil {
		params["enable_login"] = enableLoginOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("LoginControl", "update")
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

/******  Proc ******/

// List of process variables, such as version, identity, and protocol_version
type Proc struct {
}

//===== Locator

// ProcLocator exposes the Proc resource actions.
type ProcLocator struct {
	Href
	api *API
}

// ProcLocator builds a locator from the given href.
func (api *API) ProcLocator(href string) *ProcLocator {
	return &ProcLocator{Href(href), api}
}

//===== Actions

// GET /rll/proc
//
// List all process variables
func (loc *ProcLocator) Index() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Proc", "index")
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

// GET /rll/proc/:name
//
// Retrieve process variable value
func (loc *ProcLocator) Show() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Proc", "show")
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

// PUT /rll/proc/:name
//
// Set process variable value
func (loc *ProcLocator) Update(payload string) (string, error) {
	var res string
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("Proc", "update")
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

/******  Rl10 ******/

// Miscellaneous RightLink 10 local requests
type Rl10 struct {
}

//===== Locator

// Rl10Locator exposes the Rl10 resource actions.
type Rl10Locator struct {
	Href
	api *API
}

// Rl10Locator builds a locator from the given href.
func (api *API) Rl10Locator(href string) *Rl10Locator {
	return &Rl10Locator{Href(href), api}
}

//===== Actions

// POST /rll/upgrade
//
// Relaunch the RightLink process using a specified binary
func (loc *Rl10Locator) Upgrade(exec string) (string, error) {
	var res string
	if exec == "" {
		return res, fmt.Errorf("exec is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"exec": exec,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Rl10", "upgrade")
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

// POST /rll/run/recipe
//
// Run git-based scripts (as recipes) synchronously
func (loc *Rl10Locator) RunRecipe(recipe string, options rsapi.APIParams) (string, error) {
	var res string
	if recipe == "" {
		return res, fmt.Errorf("recipe is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"recipe": recipe,
	}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		params["arguments"] = argumentsOpt
	}
	var jsonOpt = options["json"]
	if jsonOpt != nil {
		params["json"] = jsonOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Rl10", "run_recipe")
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

// POST /rll/run/right_script
//
// Run RightScripts synchronously
func (loc *Rl10Locator) RunRightScript(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		params["arguments"] = argumentsOpt
	}
	var rightScriptOpt = options["right_script"]
	if rightScriptOpt != nil {
		params["right_script"] = rightScriptOpt
	}
	var rightScriptIdOpt = options["right_script_id"]
	if rightScriptIdOpt != nil {
		params["right_script_id"] = rightScriptIdOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Rl10", "run_right_script")
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

/******  TSS ******/

// Manipulate the TSS proxy (this is deprecated, please use the /rll/tss/control resource)
type TSS struct {
}

//===== Locator

// TSSLocator exposes the TSS resource actions.
type TSSLocator struct {
	Href
	api *API
}

// TSSLocator builds a locator from the given href.
func (api *API) TSSLocator(href string) *TSSLocator {
	return &TSSLocator{Href(href), api}
}

//===== Actions

// GET /rll/tss/hostname
//
// Get the TSS hostname to proxy (deprecated, RL10 knows the hostname)
func (loc *TSSLocator) GetHostname() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSS", "get_hostname")
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

// PUT /rll/tss/hostname
//
// Set the TSS hostname to proxy (deprecated, RL10 knows the hostname)
func (loc *TSSLocator) PutHostname(hostname string) (string, error) {
	var res string
	if hostname == "" {
		return res, fmt.Errorf("hostname is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"hostname": hostname,
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSS", "put_hostname")
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

/******  TSSControl ******/

// Manipulate monitoring (TSS) settings
type TSSControl struct {
}

//===== Locator

// TSSControlLocator exposes the TSSControl resource actions.
type TSSControlLocator struct {
	Href
	api *API
}

// TSSControlLocator builds a locator from the given href.
func (api *API) TSSControlLocator(href string) *TSSControlLocator {
	return &TSSControlLocator{Href(href), api}
}

//===== Actions

// GET /rll/tss/control
//
// Show monitoring features
func (loc *TSSControlLocator) Show() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSControl", "show")
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

// PUT /rll/tss/control
//
// Enable/disable monitoring features
func (loc *TSSControlLocator) Update(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var enableMonitoringOpt = options["enable_monitoring"]
	if enableMonitoringOpt != nil {
		params["enable_monitoring"] = enableMonitoringOpt
	}
	var tssIdOpt = options["tss_id"]
	if tssIdOpt != nil {
		params["tss_id"] = tssIdOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSControl", "update")
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

// PUT /rll/tss/control
//
// Control the TSS monitoring (deprecated, use the update action)
func (loc *TSSControlLocator) PutControl(options rsapi.APIParams) (string, error) {
	var res string
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var enableMonitoringOpt = options["enable_monitoring"]
	if enableMonitoringOpt != nil {
		params["enable_monitoring"] = enableMonitoringOpt
	}
	var tssIdOpt = options["tss_id"]
	if tssIdOpt != nil {
		params["tss_id"] = tssIdOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSControl", "put_control")
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

/******  TSSPlugin ******/

// TSS Custom Plugins
type TSSPlugin struct {
}

//===== Locator

// TSSPluginLocator exposes the TSSPlugin resource actions.
type TSSPluginLocator struct {
	Href
	api *API
}

// TSSPluginLocator builds a locator from the given href.
func (api *API) TSSPluginLocator(href string) *TSSPluginLocator {
	return &TSSPluginLocator{Href(href), api}
}

//===== Actions

// GET /rll/tss/exec
//
// Get TSS plugins list
func (loc *TSSPluginLocator) Index() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSPlugin", "index")
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

// PUT /rll/tss/exec/:name
//
// Add new TSS custom plugin
func (loc *TSSPluginLocator) Create(executable string, options rsapi.APIParams) (string, error) {
	var res string
	if executable == "" {
		return res, fmt.Errorf("executable is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"executable": executable,
	}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		params["arguments[]"] = argumentsOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSPlugin", "create")
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

// GET /rll/tss/exec/:name
//
// Get TSS plugin info
func (loc *TSSPluginLocator) Show() (string, error) {
	var res string
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSPlugin", "show")
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

// PUT /rll/tss/exec/:name
//
// Update TSS custom plugin
func (loc *TSSPluginLocator) Update(executable string, options rsapi.APIParams) error {
	if executable == "" {
		return fmt.Errorf("executable is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"executable": executable,
	}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		params["arguments[]"] = argumentsOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSPlugin", "update")
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

// DELETE /rll/tss/exec/:name
//
// Delete TSS plugin info
func (loc *TSSPluginLocator) Destroy() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("TSSPlugin", "destroy")
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
