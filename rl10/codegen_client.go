//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=rl10/docs/api -output=rl10 -pkg=rl10 -target=unversioned -client=Api
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

/******  DebugCookbookPath ******/

// Manipulate debug cookbook directory location
type DebugCookbookPath struct {
}

//===== Locator

// DebugCookbookPath resource locator, exposes resource actions.
type DebugCookbookPathLocator struct {
	UrlResolver
	api *Api
}

// DebugCookbookPath resource locator factory
func (api *Api) DebugCookbookPathLocator(href string) *DebugCookbookPathLocator {
	return &DebugCookbookPathLocator{UrlResolver(href), api}
}

//===== Actions

// GET /rll/debug/cookbook
//
// Retrieve debug cookbook directory location
func (loc *DebugCookbookPathLocator) Show() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("DebugCookbookPath", "show")
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

// PUT /rll/debug/cookbook
//
// Set debug cookbook directory location
func (loc *DebugCookbookPathLocator) Update(path string) (string, error) {
	var res string
	if path == "" {
		return res, fmt.Errorf("path is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"path": path,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("DebugCookbookPath", "update")
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

// DELETE /rll/debug/cookbook
//
// Remove debug cookbook directory location
func (loc *DebugCookbookPathLocator) Delete() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("DebugCookbookPath", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  Env ******/

// Manipulate global script environment variables
type Env struct {
}

//===== Locator

// Env resource locator, exposes resource actions.
type EnvLocator struct {
	UrlResolver
	api *Api
}

// Env resource locator factory
func (api *Api) EnvLocator(href string) *EnvLocator {
	return &EnvLocator{UrlResolver(href), api}
}

//===== Actions

// GET /rll/env
//
// Retrieve all environment variables
func (loc *EnvLocator) Index() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Env", "index")
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

// GET /rll/env/:name
//
// Retrieve environment variable value
func (loc *EnvLocator) Show() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Env", "show")
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

// PUT /rll/env/:name
//
// Set environment variable value
func (loc *EnvLocator) Update(payload string) (string, error) {
	var res string
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"payload": payload,
	}
	uri, err := loc.Url("Env", "update")
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

// DELETE /rll/env/:name
//
// Delete environment variable
func (loc *EnvLocator) Delete() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Env", "delete")
	if err != nil {
		return err
	}
	_, err = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	return nil
}

/******  Proc ******/

// List of process variables, such as version, identity, and protocol_version
type Proc struct {
}

//===== Locator

// Proc resource locator, exposes resource actions.
type ProcLocator struct {
	UrlResolver
	api *Api
}

// Proc resource locator factory
func (api *Api) ProcLocator(href string) *ProcLocator {
	return &ProcLocator{UrlResolver(href), api}
}

//===== Actions

// GET /rll/proc
//
// List all process variables
func (loc *ProcLocator) Index() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Proc", "index")
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

// GET /rll/proc/:name
//
// Retrieve process variable value
func (loc *ProcLocator) Show() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Proc", "show")
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

/******  Rl10 ******/

// Miscellaneous RightLink 10 local requests
type Rl10 struct {
}

//===== Locator

// Rl10 resource locator, exposes resource actions.
type Rl10Locator struct {
	UrlResolver
	api *Api
}

// Rl10 resource locator factory
func (api *Api) Rl10Locator(href string) *Rl10Locator {
	return &Rl10Locator{UrlResolver(href), api}
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
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"exec": exec,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Rl10", "upgrade")
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

// POST /rll/run/recipe
//
// Run git-based scripts (as recipes) synchronously
func (loc *Rl10Locator) RunRecipe(recipe string, options rsapi.ApiParams) (string, error) {
	var res string
	if recipe == "" {
		return res, fmt.Errorf("recipe is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"recipe": recipe,
	}
	var jsonOpt = options["json"]
	if jsonOpt != nil {
		queryParams["json"] = jsonOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Rl10", "run_recipe")
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

// POST /rll/run/right_script
//
// Run RightScripts synchronously
func (loc *Rl10Locator) RunRightScript(options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		queryParams["arguments"] = argumentsOpt
	}
	var rightScriptOpt = options["right_script"]
	if rightScriptOpt != nil {
		queryParams["right_script"] = rightScriptOpt
	}
	var rightScriptIdOpt = options["right_script_id"]
	if rightScriptIdOpt != nil {
		queryParams["right_script_id"] = rightScriptIdOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.Url("Rl10", "run_right_script")
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

/****** Parameter Data Types ******/
