//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=grs/docs/api -output=grs -pkg=grs -target=2.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package grs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// API Version
const APIVersion = "2.0"

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

/******  AccessRule ******/

// Represents a set of access control statements that grant roles to
// principals (i.e. Users or Groups) within a given scope.
type AccessRule struct {
}

//===== Locator

// AccessRuleLocator exposes the AccessRule resource actions.
type AccessRuleLocator struct {
	Href
	api *API
}

// AccessRuleLocator builds a locator from the given href.
func (api *API) AccessRuleLocator(href string) *AccessRuleLocator {
	return &AccessRuleLocator{Href(href), api}
}

//===== Actions

// GET /grs/access_rules
//
// Lists all AccessRules.
func (loc *AccessRuleLocator) Index(options rsapi.APIParams) ([]*AccessRule, error) {
	var res []*AccessRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var principalHrefOpt = options["principal_href"]
	if principalHrefOpt != nil {
		params["principal_href"] = principalHrefOpt
	}
	var scopeHrefOpt = options["scope_href"]
	if scopeHrefOpt != nil {
		params["scope_href"] = scopeHrefOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AccessRule", "index")
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

// GET /grs/access_rules/:id
//
// Shows a single AccessRule.
func (loc *AccessRuleLocator) Show(options rsapi.APIParams) (*AccessRule, error) {
	var res *AccessRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AccessRule", "show")
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

// POST /grs/access_rules
//
// Creates an AccessRule.
func (loc *AccessRuleLocator) Create(principalHref string, roles []string, scopeHref string) (*AccessRuleLocator, error) {
	var res *AccessRuleLocator
	if principalHref == "" {
		return res, fmt.Errorf("principalHref is required")
	}
	if len(roles) == 0 {
		return res, fmt.Errorf("roles is required")
	}
	if scopeHref == "" {
		return res, fmt.Errorf("scopeHref is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"principal_href": principalHref,
		"roles":          roles,
		"scope_href":     scopeHref,
	}
	uri, err := loc.ActionPath("AccessRule", "create")
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
		return &AccessRuleLocator{Href(location), loc.api}, nil
	}
}

// PUT /grs/access_rules
//
// Replaces all AccessRules for a given scope.
func (loc *AccessRuleLocator) Replace(payload []*PayloadStruct, options rsapi.APIParams) error {
	if len(payload) == 0 {
		return fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var principalHrefOpt = options["principal_href"]
	if principalHrefOpt != nil {
		params["principal_href"] = principalHrefOpt
	}
	var scopeHrefOpt = options["scope_href"]
	if scopeHrefOpt != nil {
		params["scope_href"] = scopeHrefOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("AccessRule", "replace")
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

// DELETE /grs/access_rules/:id
//
// Deletes an AccessRule.
func (loc *AccessRuleLocator) Delete() (*AccessRule, error) {
	var res *AccessRule
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("AccessRule", "delete")
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

/******  Group ******/

// A Group represents a set of Users. It serves as a convenient bucket for
// granting identical privileges to large number of Users and for
// user-reporting purposes.
type Group struct {
}

//===== Locator

// GroupLocator exposes the Group resource actions.
type GroupLocator struct {
	Href
	api *API
}

// GroupLocator builds a locator from the given href.
func (api *API) GroupLocator(href string) *GroupLocator {
	return &GroupLocator{Href(href), api}
}

//===== Actions

// GET /grs/groups
//
// Lists all Groups.
func (loc *GroupLocator) Index(options rsapi.APIParams) ([]*Group, error) {
	var res []*Group
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Group", "index")
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

// GET /grs/groups/:id
//
// Shows a single Group.
func (loc *GroupLocator) Show(options rsapi.APIParams) (*Group, error) {
	var res *Group
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Group", "show")
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

// POST /grs/groups
//
// Creates a new Group.
func (loc *GroupLocator) Create(name string, org string, options rsapi.APIParams) (*GroupLocator, error) {
	var res *GroupLocator
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if org == "" {
		return res, fmt.Errorf("org is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name": name,
		"org":  org,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("Group", "create")
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
		return &GroupLocator{Href(location), loc.api}, nil
	}
}

// PATCH /grs/groups/:id
//
// Updates a Group.
func (loc *GroupLocator) Update(org string, options rsapi.APIParams) error {
	if org == "" {
		return fmt.Errorf("org is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"org": org,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	uri, err := loc.ActionPath("Group", "update")
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

// DELETE /grs/groups/:id
//
// Deletes a Group. This action deletes a Group only if the Group does not
// contain any Users or ProvisioningRules. If the Group has Users or
// ProvisioningRules, they must be deleted from the Group before trying to
// delete the Group.
func (loc *GroupLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Group", "delete")
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

/******  GroupUser ******/

// User sub-collection of a Group
type GroupUser struct {
}

//===== Locator

// GroupUserLocator exposes the GroupUser resource actions.
type GroupUserLocator struct {
	Href
	api *API
}

// GroupUserLocator builds a locator from the given href.
func (api *API) GroupUserLocator(href string) *GroupUserLocator {
	return &GroupUserLocator{Href(href), api}
}

//===== Actions

// GET /grs/groups/:group_id/users
//
// Lists all Users in a Group.
func (loc *GroupUserLocator) Index(options rsapi.APIParams) ([]*User, error) {
	var res []*User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("GroupUser", "index")
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

// GET /grs/groups/:group_id/users/:id
//
// Shows a User in a Group.
func (loc *GroupUserLocator) Show(options rsapi.APIParams) (*User, error) {
	var res *User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("GroupUser", "show")
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

// POST /grs/groups/:group_id/users
//
// Adds a User to the Group.
func (loc *GroupUserLocator) Create(payload string) (*GroupUserLocator, error) {
	var res *GroupUserLocator
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("GroupUser", "create")
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
		return &GroupUserLocator{Href(location), loc.api}, nil
	}
}

// PUT /grs/groups/:group_id/users
//
// Replaces the Users associated with the Group.
func (loc *GroupUserLocator) Replace(payload []string) error {
	if len(payload) == 0 {
		return fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("GroupUser", "replace")
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

// DELETE /grs/groups/:group_id/users(/:id)?
//
// Deletes one (or all) User memberships from the Group.
func (loc *GroupUserLocator) Delete(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var idOpt = options["id"]
	if idOpt != nil {
		params["id"] = idOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("GroupUser", "delete")
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

/******  IdentityProvider ******/

// An IdentityProvider represents a SAML identity provider (IdP)
// that is linked to your organization and trusted by the RightScale
// dashboard to authenticate your organization's users.
type IdentityProvider struct {
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

// GET /grs/identity_providers
//
// Lists all IdentityProviders.
func (loc *IdentityProviderLocator) Index(options rsapi.APIParams) ([]*IdentityProvider, error) {
	var res []*IdentityProvider
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
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

// GET /grs/identity_providers/:id
//
// Shows a single IdentityProvider.
func (loc *IdentityProviderLocator) Show(options rsapi.APIParams) (*IdentityProvider, error) {
	var res *IdentityProvider
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
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

/******  Org ******/

// Represents an organizational unit. An Org may have a parent
// Org and many child Orgs, thereby allowing a hierarchy of Orgs
// to be defined.
type Org struct {
}

//===== Locator

// OrgLocator exposes the Org resource actions.
type OrgLocator struct {
	Href
	api *API
}

// OrgLocator builds a locator from the given href.
func (api *API) OrgLocator(href string) *OrgLocator {
	return &OrgLocator{Href(href), api}
}

//===== Actions

// GET /grs/orgs/:id
//
// Shows a single Org.
func (loc *OrgLocator) Show(options rsapi.APIParams) (*Org, error) {
	var res *Org
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Org", "show")
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

// POST /grs/orgs
//
// Creates an Org.
func (loc *OrgLocator) Create(displayName string, name string, options rsapi.APIParams) (*OrgLocator, error) {
	var res *OrgLocator
	if displayName == "" {
		return res, fmt.Errorf("displayName is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"display_name": displayName,
		"name":         name,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var parentOrgOpt = options["parent_org"]
	if parentOrgOpt != nil {
		p["parent_org"] = parentOrgOpt
	}
	uri, err := loc.ActionPath("Org", "create")
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
		return &OrgLocator{Href(location), loc.api}, nil
	}
}

// PATCH /grs/orgs/:id
//
// Updates an Org.
func (loc *OrgLocator) Update(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var displayNameOpt = options["display_name"]
	if displayNameOpt != nil {
		p["display_name"] = displayNameOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var parentOrgOpt = options["parent_org"]
	if parentOrgOpt != nil {
		p["parent_org"] = parentOrgOpt
	}
	uri, err := loc.ActionPath("Org", "update")
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

// GET /grs/orgs/:id/child_orgs
//
// Lists all the child Orgs for an Org.
func (loc *OrgLocator) ChildOrgs(options rsapi.APIParams) ([]*Org, error) {
	var res []*Org
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var recurseOpt = options["recurse"]
	if recurseOpt != nil {
		params["recurse"] = recurseOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Org", "child_orgs")
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

/******  OrgGroup ******/

// Group sub-collection of an Org.
type OrgGroup struct {
}

//===== Locator

// OrgGroupLocator exposes the OrgGroup resource actions.
type OrgGroupLocator struct {
	Href
	api *API
}

// OrgGroupLocator builds a locator from the given href.
func (api *API) OrgGroupLocator(href string) *OrgGroupLocator {
	return &OrgGroupLocator{Href(href), api}
}

//===== Actions

// GET /grs/orgs/:org_id/groups
//
// Lists all Groups in an Org.
func (loc *OrgGroupLocator) Index(options rsapi.APIParams) ([]*Group, error) {
	var res []*Group
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var recurseOpt = options["recurse"]
	if recurseOpt != nil {
		params["recurse"] = recurseOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("OrgGroup", "index")
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

/******  OrgProject ******/

// Project sub-collection of an Org.
type OrgProject struct {
}

//===== Locator

// OrgProjectLocator exposes the OrgProject resource actions.
type OrgProjectLocator struct {
	Href
	api *API
}

// OrgProjectLocator builds a locator from the given href.
func (api *API) OrgProjectLocator(href string) *OrgProjectLocator {
	return &OrgProjectLocator{Href(href), api}
}

//===== Actions

// GET /grs/orgs/:org_id/projects
//
// Lists all Projects in an Org.
func (loc *OrgProjectLocator) Index(options rsapi.APIParams) ([]*Project, error) {
	var res []*Project
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var recurseOpt = options["recurse"]
	if recurseOpt != nil {
		params["recurse"] = recurseOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("OrgProject", "index")
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

// GET /grs/orgs/:org_id/projects/:id
//
// Shows a single Project scoped to its parent Org.
func (loc *OrgProjectLocator) Show(options rsapi.APIParams) (*Project, error) {
	var res *Project
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var recurseOpt = options["recurse"]
	if recurseOpt != nil {
		params["recurse"] = recurseOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("OrgProject", "show")
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

// POST /grs/orgs/:org_id/projects
//
// Creates a new Project.
func (loc *OrgProjectLocator) Create(name string, org string, options rsapi.APIParams) (*OrgProjectLocator, error) {
	var res *OrgProjectLocator
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if org == "" {
		return res, fmt.Errorf("org is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name": name,
		"org":  org,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("OrgProject", "create")
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
		return &OrgProjectLocator{Href(location), loc.api}, nil
	}
}

// PATCH /grs/orgs/:org_id/projects/:id
//
// Updates a Project.
func (loc *OrgProjectLocator) Update(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("OrgProject", "update")
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

/******  OrgRole ******/

// Roles sub-collection of an Org.
type OrgRole struct {
}

//===== Locator

// OrgRoleLocator exposes the OrgRole resource actions.
type OrgRoleLocator struct {
	Href
	api *API
}

// OrgRoleLocator builds a locator from the given href.
func (api *API) OrgRoleLocator(href string) *OrgRoleLocator {
	return &OrgRoleLocator{Href(href), api}
}

//===== Actions

// GET /grs/orgs/:org_id/roles
//
// No description provided for index.
func (loc *OrgRoleLocator) Index(options rsapi.APIParams) ([]*Role, error) {
	var res []*Role
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("OrgRole", "index")
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

/******  OrgUser ******/

// User sub-collection of an Org.
type OrgUser struct {
}

//===== Locator

// OrgUserLocator exposes the OrgUser resource actions.
type OrgUserLocator struct {
	Href
	api *API
}

// OrgUserLocator builds a locator from the given href.
func (api *API) OrgUserLocator(href string) *OrgUserLocator {
	return &OrgUserLocator{Href(href), api}
}

//===== Actions

// GET /grs/orgs/:org_id/users
//
// Lists all Users in an Org.
func (loc *OrgUserLocator) Index(options rsapi.APIParams) ([]*User, error) {
	var res []*User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var recurseOpt = options["recurse"]
	if recurseOpt != nil {
		params["recurse"] = recurseOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("OrgUser", "index")
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

// GET /grs/orgs/:org_id/users/:id
//
// Shows a User in an Org.
func (loc *OrgUserLocator) Show(options rsapi.APIParams) (*User, error) {
	var res *User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("OrgUser", "show")
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

// POST /grs/orgs/:org_id/users
//
// Creates a User affiliation with the Org.
func (loc *OrgUserLocator) Create(payload string) (*OrgUserLocator, error) {
	var res *OrgUserLocator
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("OrgUser", "create")
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
		return &OrgUserLocator{Href(location), loc.api}, nil
	}
}

// PUT /grs/orgs/:org_id/users
//
// Replaces the User affiliations in an Org. If this action should delete
// a previous User affiliation, it will only do so if the User is not a
// member of any Groups belonging to the Org. If the User is still a
// member of Groups belonging to the Org, the Group membership must
// be deleted before trying to delete the User affiliation in an Org via
// the replace action.
func (loc *OrgUserLocator) Replace(payload []string) error {
	if len(payload) == 0 {
		return fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("OrgUser", "replace")
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

// DELETE /grs/orgs/:org_id/users/:id
//
// Deletes a User affiliation in an Org. This action deletes a User
// affiliation only if the User is not a member of any Groups belonging
// to the Org. If the User is still a member of Groups belonging to the
// Org, the Group membership must be deleted before trying to delete the
// User affiliation in an Org.
func (loc *OrgUserLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("OrgUser", "delete")
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

/******  Project ******/

// A Project represents a container for related resources. It
// belongs to an Org and can contain many cloud, design, and
// other types of resources.
type Project struct {
}

//===== Locator

// ProjectLocator exposes the Project resource actions.
type ProjectLocator struct {
	Href
	api *API
}

// ProjectLocator builds a locator from the given href.
func (api *API) ProjectLocator(href string) *ProjectLocator {
	return &ProjectLocator{Href(href), api}
}

//===== Actions

// GET /grs/projects
//
// Lists all Projects in an Org.
func (loc *ProjectLocator) Index(options rsapi.APIParams) ([]*Project, error) {
	var res []*Project
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var recurseOpt = options["recurse"]
	if recurseOpt != nil {
		params["recurse"] = recurseOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Project", "index")
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

// GET /grs/projects/:id
//
// Shows a single Project.
func (loc *ProjectLocator) Show(orgId string, options rsapi.APIParams) (*Project, error) {
	var res *Project
	if orgId == "" {
		return res, fmt.Errorf("orgId is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{
		"org_id": orgId,
	}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Project", "show")
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

// POST /grs/projects
//
// Creates a new Project.
func (loc *ProjectLocator) Create(displayName string, name string, org string, options rsapi.APIParams) (*ProjectLocator, error) {
	var res *ProjectLocator
	if displayName == "" {
		return res, fmt.Errorf("displayName is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if org == "" {
		return res, fmt.Errorf("org is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"display_name": displayName,
		"name":         name,
		"org":          org,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("Project", "create")
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
		return &ProjectLocator{Href(location), loc.api}, nil
	}
}

// PATCH /grs/projects/:id
//
// Updates a Project.
func (loc *ProjectLocator) Update(displayName string, name string, options rsapi.APIParams) error {
	if displayName == "" {
		return fmt.Errorf("displayName is required")
	}
	if name == "" {
		return fmt.Errorf("name is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"display_name": displayName,
		"name":         name,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("Project", "update")
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

/******  ProvisioningRule ******/

// A relationship table that links ProvisioningTemplates to Groups
type ProvisioningRule struct {
}

//===== Locator

// ProvisioningRuleLocator exposes the ProvisioningRule resource actions.
type ProvisioningRuleLocator struct {
	Href
	api *API
}

// ProvisioningRuleLocator builds a locator from the given href.
func (api *API) ProvisioningRuleLocator(href string) *ProvisioningRuleLocator {
	return &ProvisioningRuleLocator{Href(href), api}
}

//===== Actions

// GET /grs/provisioning_rules
//
// Lists all ProvisioningRules.
func (loc *ProvisioningRuleLocator) Index(options rsapi.APIParams) ([]*ProvisioningRule, error) {
	var res []*ProvisioningRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningRule", "index")
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

// GET /grs/provisioning_rules/:id
//
// Shows a single ProvisioningRule.
func (loc *ProvisioningRuleLocator) Show(options rsapi.APIParams) (*ProvisioningRule, error) {
	var res *ProvisioningRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningRule", "show")
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

// POST /grs/provisioning_rules
//
// Creates a ProvisioningRule.
func (loc *ProvisioningRuleLocator) Create(group string, provisioningTemplate string, regex string, options rsapi.APIParams) (*ProvisioningRuleLocator, error) {
	var res *ProvisioningRuleLocator
	if group == "" {
		return res, fmt.Errorf("group is required")
	}
	if provisioningTemplate == "" {
		return res, fmt.Errorf("provisioningTemplate is required")
	}
	if regex == "" {
		return res, fmt.Errorf("regex is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"group":                 group,
		"provisioning_template": provisioningTemplate,
		"regex":                 regex,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("ProvisioningRule", "create")
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
		return &ProvisioningRuleLocator{Href(location), loc.api}, nil
	}
}

// PATCH /grs/provisioning_rules/:id
//
// Updates a ProvisioningRule.
func (loc *ProvisioningRuleLocator) Update(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var regexOpt = options["regex"]
	if regexOpt != nil {
		p["regex"] = regexOpt
	}
	uri, err := loc.ActionPath("ProvisioningRule", "update")
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

// DELETE /grs/provisioning_rules/:id
//
// Deletes a ProvisioningRule.
func (loc *ProvisioningRuleLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningRule", "delete")
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

/******  ProvisioningTemplate ******/

// Represents a set of rules applied to Users or Groups for a
// given IdentityProvider.
type ProvisioningTemplate struct {
}

//===== Locator

// ProvisioningTemplateLocator exposes the ProvisioningTemplate resource actions.
type ProvisioningTemplateLocator struct {
	Href
	api *API
}

// ProvisioningTemplateLocator builds a locator from the given href.
func (api *API) ProvisioningTemplateLocator(href string) *ProvisioningTemplateLocator {
	return &ProvisioningTemplateLocator{Href(href), api}
}

//===== Actions

// GET /grs/identity_providers/:identity_provider_id/provisioning_templates
//
// Lists all ProvisioningTemplates for an IdentityProvider.
func (loc *ProvisioningTemplateLocator) Index(options rsapi.APIParams) ([]*ProvisioningTemplate, error) {
	var res []*ProvisioningTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningTemplate", "index")
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

// GET /grs/identity_providers/:identity_provider_id/provisioning_templates/:id
//
// Shows a single ProvisioningTemplate for an IdentityProvider.
func (loc *ProvisioningTemplateLocator) Show(options rsapi.APIParams) (*ProvisioningTemplate, error) {
	var res *ProvisioningTemplate
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningTemplate", "show")
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

// POST /grs/identity_providers/:identity_provider_id/provisioning_templates
//
// Creates a ProvisioningTemplate for an IdentityProvider.
func (loc *ProvisioningTemplateLocator) Create(name string, xsl string, options rsapi.APIParams) (*ProvisioningTemplateLocator, error) {
	var res *ProvisioningTemplateLocator
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if xsl == "" {
		return res, fmt.Errorf("xsl is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"name": name,
		"xsl":  xsl,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("ProvisioningTemplate", "create")
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
		return &ProvisioningTemplateLocator{Href(location), loc.api}, nil
	}
}

// PATCH /grs/identity_providers/:identity_provider_id/provisioning_templates/:id
//
// Updates a ProvisioningTemplate for an IdentityProvider.
func (loc *ProvisioningTemplateLocator) Update(options rsapi.APIParams) error {
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
	var xslOpt = options["xsl"]
	if xslOpt != nil {
		p["xsl"] = xslOpt
	}
	uri, err := loc.ActionPath("ProvisioningTemplate", "update")
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

// DELETE /grs/identity_providers/:identity_provider_id/provisioning_templates/:id
//
// Deletes a ProvisioningTemplate. This action deletes a
// ProvisioningTemplate only if it does not have any ProvisioningRules.
// If the ProvisioningTemplate has ProvisioningRules associated with it,
// the ProvisioningRule associations must be deleted before trying to
// delete the ProvisioningTemplate.
func (loc *ProvisioningTemplateLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningTemplate", "delete")
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

// PUT /grs/identity_providers/:identity_provider_id/provisioning_templates/:id/actions/make_default
//
// Marks a ProvisioningTemplate as the default for the IdentityProvider.
func (loc *ProvisioningTemplateLocator) MakeDefault() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningTemplate", "make_default")
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

/******  ProvisioningTemplateRule ******/

// ProvisioningRules sub-collection of a ProvisioningTemplate.
type ProvisioningTemplateRule struct {
}

//===== Locator

// ProvisioningTemplateRuleLocator exposes the ProvisioningTemplateRule resource actions.
type ProvisioningTemplateRuleLocator struct {
	Href
	api *API
}

// ProvisioningTemplateRuleLocator builds a locator from the given href.
func (api *API) ProvisioningTemplateRuleLocator(href string) *ProvisioningTemplateRuleLocator {
	return &ProvisioningTemplateRuleLocator{Href(href), api}
}

//===== Actions

// GET /grs/identity_providers/:identity_provider_id/provisioning_templates/:provisioning_template_id/rules
//
// Lists all ProvisioningRules for the given ProvisioningTemplate.
func (loc *ProvisioningTemplateRuleLocator) Index(options rsapi.APIParams) ([]*ProvisioningRule, error) {
	var res []*ProvisioningRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningTemplateRule", "index")
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

// GET /grs/identity_providers/:identity_provider_id/provisioning_templates/:provisioning_template_id/rules/:id
//
// Shows a ProvisioningRule for the given ProvisioningTemplate.
func (loc *ProvisioningTemplateRuleLocator) Show(options rsapi.APIParams) (*ProvisioningRule, error) {
	var res *ProvisioningRule
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningTemplateRule", "show")
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

// POST /grs/identity_providers/:identity_provider_id/provisioning_templates/:provisioning_template_id/rules
//
// Creates a ProvisioningRule for the given ProvisioningTemplate.
func (loc *ProvisioningTemplateRuleLocator) Create(group string, regex string, options rsapi.APIParams) (*ProvisioningTemplateRuleLocator, error) {
	var res *ProvisioningTemplateRuleLocator
	if group == "" {
		return res, fmt.Errorf("group is required")
	}
	if regex == "" {
		return res, fmt.Errorf("regex is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"group": group,
		"regex": regex,
	}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	uri, err := loc.ActionPath("ProvisioningTemplateRule", "create")
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
		return &ProvisioningTemplateRuleLocator{Href(location), loc.api}, nil
	}
}

// PATCH /grs/identity_providers/:identity_provider_id/provisioning_templates/:provisioning_template_id/rules/:id
//
// Updates a ProvisioningRule for the given ProvisioningTemplate.
func (loc *ProvisioningTemplateRuleLocator) Update(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var regexOpt = options["regex"]
	if regexOpt != nil {
		p["regex"] = regexOpt
	}
	uri, err := loc.ActionPath("ProvisioningTemplateRule", "update")
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

// DELETE /grs/identity_providers/:identity_provider_id/provisioning_templates/:provisioning_template_id/rules/:id
//
// Deletes a ProvisioningRule for the given ProvisioningTemplate.
func (loc *ProvisioningTemplateRuleLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ProvisioningTemplateRule", "delete")
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

/******  User ******/

// A User represents an individual RightScale user.
type User struct {
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

// GET /grs/users/:id
//
// Shows a single User.
func (loc *UserLocator) Show(options rsapi.APIParams) (*User, error) {
	var res *User
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
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

// POST /grs/users
//
// Creates a User.
func (loc *UserLocator) Create(email string, firstName string, lastName string, options rsapi.APIParams) (*UserLocator, error) {
	var res *UserLocator
	if email == "" {
		return res, fmt.Errorf("email is required")
	}
	if firstName == "" {
		return res, fmt.Errorf("firstName is required")
	}
	if lastName == "" {
		return res, fmt.Errorf("lastName is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
	}
	var companyOpt = options["company"]
	if companyOpt != nil {
		p["company"] = companyOpt
	}
	var passwordOpt = options["password"]
	if passwordOpt != nil {
		p["password"] = passwordOpt
	}
	var phoneOpt = options["phone"]
	if phoneOpt != nil {
		p["phone"] = phoneOpt
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

// PATCH /grs/users/:id
//
// Updates a User.
func (loc *UserLocator) Update(options rsapi.APIParams) error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{}
	var companyOpt = options["company"]
	if companyOpt != nil {
		p["company"] = companyOpt
	}
	var emailOpt = options["email"]
	if emailOpt != nil {
		p["email"] = emailOpt
	}
	var firstNameOpt = options["first_name"]
	if firstNameOpt != nil {
		p["first_name"] = firstNameOpt
	}
	var lastNameOpt = options["last_name"]
	if lastNameOpt != nil {
		p["last_name"] = lastNameOpt
	}
	var phoneOpt = options["phone"]
	if phoneOpt != nil {
		p["phone"] = phoneOpt
	}
	var timezoneNameOpt = options["timezone_name"]
	if timezoneNameOpt != nil {
		p["timezone_name"] = timezoneNameOpt
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

/******  UserGroup ******/

// Group sub-collection of a user
type UserGroup struct {
}

//===== Locator

// UserGroupLocator exposes the UserGroup resource actions.
type UserGroupLocator struct {
	Href
	api *API
}

// UserGroupLocator builds a locator from the given href.
func (api *API) UserGroupLocator(href string) *UserGroupLocator {
	return &UserGroupLocator{Href(href), api}
}

//===== Actions

// GET /grs/users/:user_id/groups
//
// Lists all Groups of which the User is a member.
func (loc *UserGroupLocator) Index(options rsapi.APIParams) ([]*Group, error) {
	var res []*Group
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserGroup", "index")
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

// GET /grs/users/:user_id/groups/:id
//
// Shows a Group of which the User is a member.
func (loc *UserGroupLocator) Show(options rsapi.APIParams) (*Group, error) {
	var res *Group
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserGroup", "show")
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

// POST /grs/users/:user_id/groups
//
// Creates a Group membership for the User.
func (loc *UserGroupLocator) Create(payload string) (*UserGroupLocator, error) {
	var res *UserGroupLocator
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("UserGroup", "create")
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
		return &UserGroupLocator{Href(location), loc.api}, nil
	}
}

// PUT /grs/users/:user_id/groups
//
// Replaces the Group membership(s) for the User.
func (loc *UserGroupLocator) Replace(payload []string, options rsapi.APIParams) error {
	if len(payload) == 0 {
		return fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var orgHrefOpt = options["org_href"]
	if orgHrefOpt != nil {
		params["org_href"] = orgHrefOpt
	}
	var samlOnlyOpt = options["saml_only"]
	if samlOnlyOpt != nil {
		params["saml_only"] = samlOnlyOpt
	}
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("UserGroup", "replace")
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

// DELETE /grs/users/:user_id/groups(/:id)?
//
// Deletes one (or all) Group membership(s) for the User.
func (loc *UserGroupLocator) Delete(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var idOpt = options["id"]
	if idOpt != nil {
		params["id"] = idOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserGroup", "delete")
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

/******  UserIdentity ******/

// Represents a User's principal identity in the context of a
// specific identity provider. Known to our platform by a SAML
// NameID (also known as SAML subject), which is an opaque
// string, often an email address, but sometimes an LDAP DN,
// ActiveDirectory username, random GUID, or other identifier.
type UserIdentity struct {
}

//===== Locator

// UserIdentityLocator exposes the UserIdentity resource actions.
type UserIdentityLocator struct {
	Href
	api *API
}

// UserIdentityLocator builds a locator from the given href.
func (api *API) UserIdentityLocator(href string) *UserIdentityLocator {
	return &UserIdentityLocator{Href(href), api}
}

//===== Actions

// GET /grs/user_identities
//
// Lists all UserIdentities.
func (loc *UserIdentityLocator) Index(options rsapi.APIParams) ([]*UserIdentity, error) {
	var res []*UserIdentity
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserIdentity", "index")
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

// GET /grs/user_identities/:id
//
// Shows a single UserIdentity.
func (loc *UserIdentityLocator) Show(options rsapi.APIParams) (*UserIdentity, error) {
	var res *UserIdentity
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserIdentity", "show")
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

/******  UserOrg ******/

// Org sub-collection of a user.
type UserOrg struct {
}

//===== Locator

// UserOrgLocator exposes the UserOrg resource actions.
type UserOrgLocator struct {
	Href
	api *API
}

// UserOrgLocator builds a locator from the given href.
func (api *API) UserOrgLocator(href string) *UserOrgLocator {
	return &UserOrgLocator{Href(href), api}
}

//===== Actions

// GET /grs/users/:user_id/orgs
//
// Lists all Orgs with which the User is affiliated.
func (loc *UserOrgLocator) Index(options rsapi.APIParams) ([]*Org, error) {
	var res []*Org
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserOrg", "index")
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

// GET /grs/users/:user_id/orgs/:id
//
// Shows an Org with which the User is affiliated.
func (loc *UserOrgLocator) Show(options rsapi.APIParams) (*Org, error) {
	var res *Org
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserOrg", "show")
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

// POST /grs/users/:user_id/orgs
//
// Creates an affiliation for a User to an Org.
func (loc *UserOrgLocator) Create(payload string) (*UserOrgLocator, error) {
	var res *UserOrgLocator
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("UserOrg", "create")
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
		return &UserOrgLocator{Href(location), loc.api}, nil
	}
}

// PUT /grs/users/:user_id/orgs
//
// Replaces all Org affiliations for a User. If this action should delete
// a previous Org affiliation, it will only do so if the User is not a
// member of any Groups belonging to the Org. If the User is still a
// member of Groups belonging to the to the Org, the Group membership must
// be deleted before trying to delete the Org affiliation for a User via
// the replace action.
func (loc *UserOrgLocator) Replace(payload []string) error {
	if len(payload) == 0 {
		return fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("UserOrg", "replace")
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

// DELETE /grs/users/:user_id/orgs/:id
//
// Deletes an Org affiliation for a User. This action deletes an Org
// affiliation only if the User is not a member of any Groups belonging
// to the Org. If the User is still a member of Groups belonging to the
// Org, the Group membership must be deleted before trying to delete the
// Org affiliation for a User.
func (loc *UserOrgLocator) Delete() error {
	var params rsapi.APIParams
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserOrg", "delete")
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

/******  UserUserIdentity ******/

// UserIdentity sub-collection of a User, which defines relationships between
// the User and its IdentityProviders.
type UserUserIdentity struct {
}

//===== Locator

// UserUserIdentityLocator exposes the UserUserIdentity resource actions.
type UserUserIdentityLocator struct {
	Href
	api *API
}

// UserUserIdentityLocator builds a locator from the given href.
func (api *API) UserUserIdentityLocator(href string) *UserUserIdentityLocator {
	return &UserUserIdentityLocator{Href(href), api}
}

//===== Actions

// GET /grs/users/:user_id/identities
//
// Lists all UserIdentities for the given User.
func (loc *UserUserIdentityLocator) Index(options rsapi.APIParams) ([]*UserIdentity, error) {
	var res []*UserIdentity
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserUserIdentity", "index")
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

// GET /grs/users/:user_id/identities/:id
//
// Shows a UserIdentity for the given User.
func (loc *UserUserIdentityLocator) Show(options rsapi.APIParams) (*UserIdentity, error) {
	var res *UserIdentity
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var fieldsOpt = options["fields"]
	if fieldsOpt != nil {
		params["fields"] = fieldsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("UserUserIdentity", "show")
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

// PUT /grs/users/:user_id/identities
//
// Replaces the UserIdentities for the given User.
// Currently, only one identity per user is allowed, and the
// replace action will fail if more than one identity is provided
// at a time. This restriction could be lifted in the future.
func (loc *UserUserIdentityLocator) Replace(payload []*PayloadStruct) error {
	if len(payload) == 0 {
		return fmt.Errorf("payload is required")
	}
	var params rsapi.APIParams
	var p rsapi.APIParams
	p = rsapi.APIParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("UserUserIdentity", "replace")
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

type AccessRuleLinks struct {
	Principal string `json:"principal,omitempty"`
	Scope     string `json:"scope,omitempty"`
}

type AccessRuleParam struct {
	Href  string           `json:"href,omitempty"`
	Id    int              `json:"id,omitempty"`
	Kind  string           `json:"kind,omitempty"`
	Links *AccessRuleLinks `json:"links,omitempty"`
	Roles []string         `json:"roles,omitempty"`
}

type GroupLinks struct {
	Org   *OrgParam                       `json:"org,omitempty"`
	Users *UserGroupUserCollectionSummary `json:"users,omitempty"`
}

type GroupOrgGroupCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type GroupParam struct {
	Description string              `json:"description,omitempty"`
	Href        string              `json:"href,omitempty"`
	Id          int                 `json:"id,omitempty"`
	Kind        string              `json:"kind,omitempty"`
	Links       *GroupLinks         `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Timestamps  *ResourceTimestamps `json:"timestamps,omitempty"`
	Users       []*UserParam        `json:"users,omitempty"`
}

type GroupUserGroupCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type IdentityProviderLinks struct {
	Org                   *OrgParam                                                  `json:"org,omitempty"`
	ProvisioningTemplates *ProvisioningTemplateProvisioningTemplateCollectionSummary `json:"provisioning_templates,omitempty"`
}

type IdentityProviderParam struct {
	DefaultProvisioningTemplate *ProvisioningTemplateParam `json:"default_provisioning_template,omitempty"`
	Description                 string                     `json:"description,omitempty"`
	DiscoveryHint               string                     `json:"discovery_hint,omitempty"`
	EmailDomains                []string                   `json:"email_domains,omitempty"`
	Href                        string                     `json:"href,omitempty"`
	Id                          int                        `json:"id,omitempty"`
	Kind                        string                     `json:"kind,omitempty"`
	Links                       *IdentityProviderLinks     `json:"links,omitempty"`
	Name                        string                     `json:"name,omitempty"`
	ProtocolDetails             map[string]interface{}     `json:"protocol_details,omitempty"`
	ProtocolType                string                     `json:"protocol_type,omitempty"`
	ProtocolUid                 string                     `json:"protocol_uid,omitempty"`
	Timestamps                  *ResourceTimestamps        `json:"timestamps,omitempty"`
}

type OrgChildOrgCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type OrgLinks struct {
	ChildOrgs *OrgChildOrgCollectionSummary       `json:"child_orgs,omitempty"`
	Groups    *GroupOrgGroupCollectionSummary     `json:"groups,omitempty"`
	ParentOrg *OrgParam                           `json:"parent_org,omitempty"`
	Projects  *ProjectOrgProjectCollectionSummary `json:"projects,omitempty"`
	Users     *UserOrgUserCollectionSummary       `json:"users,omitempty"`
}

type OrgParam struct {
	Description string                            `json:"description,omitempty"`
	DisplayName string                            `json:"display_name,omitempty"`
	Href        string                            `json:"href,omitempty"`
	Id          string                            `json:"id,omitempty"`
	Kind        string                            `json:"kind,omitempty"`
	Legacy      *ReturnGroupsLinksOrgLegacyStruct `json:"legacy,omitempty"`
	Links       *OrgLinks                         `json:"links,omitempty"`
	Name        string                            `json:"name,omitempty"`
	Timestamps  *ResourceTimestamps               `json:"timestamps,omitempty"`
	Users       []*UserParam                      `json:"users,omitempty"`
}

type OrgUserOrgCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type Privilege struct {
	Action     string              `json:"action,omitempty"`
	Id         string              `json:"id,omitempty"`
	Kind       string              `json:"kind,omitempty"`
	Resource   string              `json:"resource,omitempty"`
	Scopes     []string            `json:"scopes,omitempty"`
	Suite      string              `json:"suite,omitempty"`
	Timestamps *ResourceTimestamps `json:"timestamps,omitempty"`
}

type ProjectLinks struct {
	Org *OrgParam `json:"org,omitempty"`
}

type ProjectOrgProjectCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type ProjectParam struct {
	Description string              `json:"description,omitempty"`
	DisplayName string              `json:"display_name,omitempty"`
	Href        string              `json:"href,omitempty"`
	Id          string              `json:"id,omitempty"`
	Kind        string              `json:"kind,omitempty"`
	Legacy      *ReturnLegacyStruct `json:"legacy,omitempty"`
	Links       *ProjectLinks       `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
	Timestamps  *ResourceTimestamps `json:"timestamps,omitempty"`
}

type ProvisioningRuleLinks struct {
	Group                *GroupParam                `json:"group,omitempty"`
	ProvisioningTemplate *ProvisioningTemplateParam `json:"provisioning_template,omitempty"`
}

type ProvisioningRuleParam struct {
	Description string                 `json:"description,omitempty"`
	Href        string                 `json:"href,omitempty"`
	Id          int                    `json:"id,omitempty"`
	Kind        string                 `json:"kind,omitempty"`
	Links       *ProvisioningRuleLinks `json:"links,omitempty"`
	Regex       string                 `json:"regex,omitempty"`
	Timestamps  *ResourceTimestamps    `json:"timestamps,omitempty"`
}

type ProvisioningRuleProvisioningRuleCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type ProvisioningTemplateLinks struct {
	ProvisioningRules *ProvisioningRuleProvisioningRuleCollectionSummary `json:"provisioning_rules,omitempty"`
}

type ProvisioningTemplateParam struct {
	Description       string                     `json:"description,omitempty"`
	Href              string                     `json:"href,omitempty"`
	Id                int                        `json:"id,omitempty"`
	Kind              string                     `json:"kind,omitempty"`
	Links             *ProvisioningTemplateLinks `json:"links,omitempty"`
	Name              string                     `json:"name,omitempty"`
	ProvisioningRules []*ProvisioningRuleParam   `json:"provisioning_rules,omitempty"`
	Timestamps        *ResourceTimestamps        `json:"timestamps,omitempty"`
	Xsl               string                     `json:"xsl,omitempty"`
}

type ProvisioningTemplateProvisioningTemplateCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type ResourceTimestamps struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ReturnGroupsLinksOrgLegacyStruct struct {
	AccountId  int    `json:"account_id,omitempty"`
	AccountUrl string `json:"account_url,omitempty"`
}

type ReturnLegacyStruct struct {
	AccountId  int    `json:"account_id,omitempty"`
	AccountUrl string `json:"account_url,omitempty"`
}

type Role struct {
	Id         int                 `json:"id,omitempty"`
	Kind       string              `json:"kind,omitempty"`
	Links      *RoleLinks          `json:"links,omitempty"`
	Name       string              `json:"name,omitempty"`
	Privileges []*Privilege        `json:"privileges,omitempty"`
	Scopes     []string            `json:"scopes,omitempty"`
	Timestamps *ResourceTimestamps `json:"timestamps,omitempty"`
}

type RoleLinks struct {
	Org *OrgParam `json:"org,omitempty"`
}

type UserGroupUserCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type UserIdentityLinks struct {
	IdentityProvider *IdentityProviderParam `json:"identity_provider,omitempty"`
	User             *UserParam             `json:"user,omitempty"`
}

type UserIdentityParam struct {
	Href         string              `json:"href,omitempty"`
	Id           int                 `json:"id,omitempty"`
	Kind         string              `json:"kind,omitempty"`
	Links        *UserIdentityLinks  `json:"links,omitempty"`
	PrincipalUid string              `json:"principal_uid,omitempty"`
	Timestamps   *ResourceTimestamps `json:"timestamps,omitempty"`
}

type UserLinks struct {
	Groups *GroupUserGroupCollectionSummary `json:"groups,omitempty"`
	Orgs   *OrgUserOrgCollectionSummary     `json:"orgs,omitempty"`
}

type UserOrgUserCollectionSummary struct {
	Href string `json:"href,omitempty"`
}

type UserParam struct {
	Company      string              `json:"company,omitempty"`
	Email        string              `json:"email,omitempty"`
	FirstName    string              `json:"first_name,omitempty"`
	Groups       []*GroupParam       `json:"groups,omitempty"`
	Href         string              `json:"href,omitempty"`
	Id           int                 `json:"id,omitempty"`
	Kind         string              `json:"kind,omitempty"`
	LastName     string              `json:"last_name,omitempty"`
	Links        *UserLinks          `json:"links,omitempty"`
	Orgs         []*OrgParam         `json:"orgs,omitempty"`
	Phone        string              `json:"phone,omitempty"`
	Timestamps   *ResourceTimestamps `json:"timestamps,omitempty"`
	TimezoneName string              `json:"timezone_name,omitempty"`
}
