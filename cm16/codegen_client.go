//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=cm16/api_docs -output=cm16 -pkg=cm16 -target=1.6 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package cm16

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// API Version
const APIVersion = "1.6"

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

// Resources in RightScale generally belong to accounts. Users can have
// any number of accounts, but when performing an action, a user is
// operating under a particular account.
type Account struct {
	Href string `json:"href,omitempty"`
	Id   int    `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Account) Locator(api *API) *AccountLocator {
	return api.AccountLocator(r.Href)
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

// GET /api/accounts
//
// Currently not implemented.
func (loc *AccountLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Account", "index")
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

// GET /api/accounts/:id
//
// Currently not implemented.
func (loc *AccountLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Account", "show")
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

// Clouds provide remote resources for things like storage and compute.
// You must have registered a cloud within your account in order to use
// it.
type Cloud struct {
	CloudType   string `json:"cloud_type,omitempty"`
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          int    `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Cloud) Locator(api *API) *CloudLocator {
	return api.CloudLocator(r.Href)
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
// Currently not implemented.
func (loc *CloudLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Cloud", "index")
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

// GET /api/clouds/:id
//
// Currently not implemented.
func (loc *CloudLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Cloud", "show")
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

// Datacenters are cloud resources that give you the ability to place
// resources in isolated locations. A carefully designed system placed in
// multiple datacenters can provide fault tolerance when one datacenter
// has a problem.
type Datacenter struct {
	Description string           `json:"description,omitempty"`
	Href        string           `json:"href,omitempty"`
	Id          string           `json:"id,omitempty"`
	Kind        string           `json:"kind,omitempty"`
	LegacyId    int              `json:"legacy_id,omitempty"`
	Links       *DatacenterLinks `json:"links,omitempty"`
	Name        string           `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Datacenter) Locator(api *API) *DatacenterLocator {
	return api.DatacenterLocator(r.Href)
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

// GET /api/datacenters
// GET /api/clouds/:cloud_id/datacenters
//
// Currently not implemented.
func (loc *DatacenterLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Datacenter", "index")
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

// GET /api/datacenters/:id
// GET /api/clouds/:cloud_id/datacenters/:id
//
// Currently not implemented.
func (loc *DatacenterLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Datacenter", "show")
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

/******  Deployment ******/

// Deployments provide a way to group resources that logically belong
// together.
type Deployment struct {
	Description  string           `json:"description,omitempty"`
	Href         string           `json:"href,omitempty"`
	Id           int              `json:"id,omitempty"`
	Instances    []*Instance      `json:"instances,omitempty"`
	Kind         string           `json:"kind,omitempty"`
	Links        *DeploymentLinks `json:"links,omitempty"`
	Locked       bool             `json:"locked,omitempty"`
	Name         string           `json:"name,omitempty"`
	ServerArrays []*ServerArray   `json:"server_arrays,omitempty"`
	Servers      []*Server        `json:"servers,omitempty"`
	Tags         []string         `json:"tags,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Deployment) Locator(api *API) *DeploymentLocator {
	return api.DeploymentLocator(r.Href)
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

// GET /api/deployments
//
// List all Deployments in an Account.
func (loc *DeploymentLocator) Index(options rsapi.APIParams) ([]*Deployment, error) {
	var res []*Deployment
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
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

// GET /api/deployments/:id
//
// Show a single Deployment
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

/******  Image ******/

// Images define the initial Operating System and root disk contents
// for new instances.
type Image struct {
	CpuArchitecture    string      `json:"cpu_architecture,omitempty"`
	Description        string      `json:"description,omitempty"`
	Href               string      `json:"href,omitempty"`
	Id                 string      `json:"id,omitempty"`
	ImageType          string      `json:"image_type,omitempty"`
	InheritedSource    string      `json:"inherited_source,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	LegacyId           int         `json:"legacy_id,omitempty"`
	Links              *ImageLinks `json:"links,omitempty"`
	Name               string      `json:"name,omitempty"`
	OsPlatform         string      `json:"os_platform,omitempty"`
	ResourceUid        string      `json:"resource_uid,omitempty"`
	RootDeviceStorage  string      `json:"root_device_storage,omitempty"`
	VirtualizationType string      `json:"virtualization_type,omitempty"`
	Visibility         string      `json:"visibility,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Image) Locator(api *API) *ImageLocator {
	return api.ImageLocator(r.Href)
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
func (loc *ImageLocator) Index(options rsapi.APIParams) ([]*Image, error) {
	var res []*Image
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
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

// GET /api/images/:id
// GET /api/clouds/:cloud_id/images/:id
//
// Currently not implemented.
func (loc *ImageLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Image", "show")
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
// An instance of type "next" is a container of information that expresses
// how to configure a future instance when we decide to launch or start
// it. A "next" instance generally only exists in the RightScale realm,
// and usually doesn't have any corresponding representation existing in
// the cloud. However, if an instance is not of type "next", it will
// generally represent an existing running (or provisioned) virtual
// machine existing in the cloud.
type Instance struct {
	Actions            []string                 `json:"actions,omitempty"`
	Description        string                   `json:"description,omitempty"`
	Href               string                   `json:"href,omitempty"`
	Id                 string                   `json:"id,omitempty"`
	IpAddresses        []*IpAddress             `json:"ip_addresses,omitempty"`
	IsNext             bool                     `json:"is_next,omitempty"`
	Kind               string                   `json:"kind,omitempty"`
	LegacyId           int                      `json:"legacy_id,omitempty"`
	Links              *InstanceLinks           `json:"links,omitempty"`
	Locked             bool                     `json:"locked,omitempty"`
	MonitoringId       string                   `json:"monitoring_id,omitempty"`
	MonitoringServer   string                   `json:"monitoring_server,omitempty"`
	MonitoringToken    string                   `json:"monitoring_token,omitempty"`
	Name               string                   `json:"name,omitempty"`
	Networks           []*Network               `json:"networks,omitempty"`
	OsPlatform         string                   `json:"os_platform,omitempty"`
	PrivateDnsNames    []string                 `json:"private_dns_names,omitempty"`
	PrivateIpAddresses []string                 `json:"private_ip_addresses,omitempty"`
	PublicDnsNames     []string                 `json:"public_dns_names,omitempty"`
	PublicIpAddresses  []string                 `json:"public_ip_addresses,omitempty"`
	ResourceUid        string                   `json:"resource_uid,omitempty"`
	SecurityGroups     *SecurityGroupCollection `json:"security_groups,omitempty"`
	ServerTemplate     *ServerTemplate          `json:"server_template,omitempty"`
	SshHost            string                   `json:"ssh_host,omitempty"`
	State              string                   `json:"state,omitempty"`
	Subnets            *SubnetCollection        `json:"subnets,omitempty"`
	Tags               []string                 `json:"tags,omitempty"`
	Timestamps         *TimestampsStruct        `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Instance) Locator(api *API) *InstanceLocator {
	return api.InstanceLocator(r.Href)
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

// GET /api/instances
// GET /api/clouds/:cloud_id/instances
//
// List all Instances in an account.
func (loc *InstanceLocator) Index(options rsapi.APIParams) ([]*Instance, error) {
	var res []*Instance
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids"] = idsOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		params["limit"] = limitOpt
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

// GET /api/instances/:id
// GET /api/clouds/:cloud_id/instances/:id
//
// Currently not implemented.
func (loc *InstanceLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Instance", "show")
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

// An InstanceType represents a basic hardware configuration for an
// Instance.
// Combining all possible configurations of hardware into a smaller,
// well-known set of options makes instances easier to manage, and allows
// better allocation efficiency into physical hosts.
type InstanceType struct {
	CpuCount    int                `json:"cpu_count,omitempty"`
	CpuSpeed    string             `json:"cpu_speed,omitempty"`
	Description string             `json:"description,omitempty"`
	Href        string             `json:"href,omitempty"`
	Id          string             `json:"id,omitempty"`
	Kind        string             `json:"kind,omitempty"`
	LegacyId    int                `json:"legacy_id,omitempty"`
	Links       *InstanceTypeLinks `json:"links,omitempty"`
	Memory      string             `json:"memory,omitempty"`
	Name        string             `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *InstanceType) Locator(api *API) *InstanceTypeLocator {
	return api.InstanceTypeLocator(r.Href)
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

// GET /api/instance_types
// GET /api/clouds/:cloud_id/instance_types
//
// Currently not implemented.
func (loc *InstanceTypeLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceType", "index")
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

// GET /api/instance_types/:id
// GET /api/clouds/:cloud_id/instance_types/:id
//
// Currently not implemented.
func (loc *InstanceTypeLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("InstanceType", "show")
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

/******  IpAddress ******/

// An IpAddress provides an abstraction for IPv4 addresses bindable to
// Instance resources running in a Cloud.
type IpAddress struct {
	Address  string          `json:"address,omitempty"`
	Href     string          `json:"href,omitempty"`
	Id       string          `json:"id,omitempty"`
	Kind     string          `json:"kind,omitempty"`
	LegacyId int             `json:"legacy_id,omitempty"`
	Links    *IpAddressLinks `json:"links,omitempty"`
	Name     string          `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *IpAddress) Locator(api *API) *IpAddressLocator {
	return api.IpAddressLocator(r.Href)
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

// GET /api/ip_addresses
// GET /api/clouds/:cloud_id/ip_addresses
//
// Currently not implemented.
func (loc *IpAddressLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddress", "index")
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

// GET /api/ip_addresses/:id
// GET /api/clouds/:cloud_id/ip_addresses/:id
//
// Currently not implemented.
func (loc *IpAddressLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddress", "show")
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

// An IpAddressBinding represents an abstraction for binding an IpAddress
// to an instance. The IpAddress is bound immediately for a current
// instance, or on launch for a next instance.
type IpAddressBinding struct {
	Href        string                 `json:"href,omitempty"`
	Id          string                 `json:"id,omitempty"`
	Kind        string                 `json:"kind,omitempty"`
	LegacyId    int                    `json:"legacy_id,omitempty"`
	Links       *IpAddressBindingLinks `json:"links,omitempty"`
	PrivatePort int                    `json:"private_port,omitempty"`
	Protocol    string                 `json:"protocol,omitempty"`
	PublicPort  int                    `json:"public_port,omitempty"`
}

// Locator returns a locator for the given resource
func (r *IpAddressBinding) Locator(api *API) *IpAddressBindingLocator {
	return api.IpAddressBindingLocator(r.Href)
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

// GET /api/ip_address_bindings
// GET /api/clouds/:cloud_id/ip_address_bindings
//
// Currently not implemented.
func (loc *IpAddressBindingLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddressBinding", "index")
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

// GET /api/ip_address_bindings/:id
// GET /api/clouds/:cloud_id/ip_address_bindings/:id
//
// Currently not implemented.
func (loc *IpAddressBindingLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("IpAddressBinding", "show")
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

/******  MultiCloudImage ******/

// A MultiCloudImage is a RightScale component that functions as a pointer
// to machine images in specific clouds (e.g. AWS US-East, Rackspace).
// Each ServerTemplate can reference many MultiCloudImages that define
// which image should be used when a server is launched in a particular
// cloud.
type MultiCloudImage struct {
	Description     string `json:"description,omitempty"`
	Href            string `json:"href,omitempty"`
	Id              int    `json:"id,omitempty"`
	InheritedSource string `json:"inherited_source,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Version         int    `json:"version,omitempty"`
}

// Locator returns a locator for the given resource
func (r *MultiCloudImage) Locator(api *API) *MultiCloudImageLocator {
	return api.MultiCloudImageLocator(r.Href)
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

// GET /api/multi_cloud_images
//
// Currently not implemented.
func (loc *MultiCloudImageLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImage", "index")
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

// GET /api/multi_cloud_images/:id
//
// Currently not implemented.
func (loc *MultiCloudImageLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("MultiCloudImage", "show")
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
	Description string        `json:"description,omitempty"`
	Href        string        `json:"href,omitempty"`
	Id          string        `json:"id,omitempty"`
	Kind        string        `json:"kind,omitempty"`
	LegacyId    int           `json:"legacy_id,omitempty"`
	Links       *NetworkLinks `json:"links,omitempty"`
	Name        string        `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Network) Locator(api *API) *NetworkLocator {
	return api.NetworkLocator(r.Href)
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

// GET /api/networks
//
// Currently not implemented.
func (loc *NetworkLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Network", "index")
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

// GET /api/networks/:id
//
// Currently not implemented.
func (loc *NetworkLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Network", "show")
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

/******  NetworkInterface ******/

// Just like their physical counterparts, NetworkInterfaces join other
// resources to a network.
type NetworkInterface struct {
	Description string                 `json:"description,omitempty"`
	Href        string                 `json:"href,omitempty"`
	Id          string                 `json:"id,omitempty"`
	Kind        string                 `json:"kind,omitempty"`
	Links       *NetworkInterfaceLinks `json:"links,omitempty"`
}

// Locator returns a locator for the given resource
func (r *NetworkInterface) Locator(api *API) *NetworkInterfaceLocator {
	return api.NetworkInterfaceLocator(r.Href)
}

//===== Locator

// NetworkInterfaceLocator exposes the NetworkInterface resource actions.
type NetworkInterfaceLocator struct {
	Href
	api *API
}

// NetworkInterfaceLocator builds a locator from the given href.
func (api *API) NetworkInterfaceLocator(href string) *NetworkInterfaceLocator {
	return &NetworkInterfaceLocator{Href(href), api}
}

//===== Actions

// GET /api/network_interfaces
//
// Currently not implemented.
func (loc *NetworkInterfaceLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkInterface", "index")
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

// GET /api/network_interfaces/:id
//
// Currently not implemented.
func (loc *NetworkInterfaceLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkInterface", "show")
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

/******  NetworkInterfaceAttachment ******/

// NetworkInterfaceAttachments represent an attachment between a
// NetworkInterface and another resource.
type NetworkInterfaceAttachment struct {
	Href  string                           `json:"href,omitempty"`
	Id    string                           `json:"id,omitempty"`
	Kind  string                           `json:"kind,omitempty"`
	Links *NetworkInterfaceAttachmentLinks `json:"links,omitempty"`
}

// Locator returns a locator for the given resource
func (r *NetworkInterfaceAttachment) Locator(api *API) *NetworkInterfaceAttachmentLocator {
	return api.NetworkInterfaceAttachmentLocator(r.Href)
}

//===== Locator

// NetworkInterfaceAttachmentLocator exposes the NetworkInterfaceAttachment resource actions.
type NetworkInterfaceAttachmentLocator struct {
	Href
	api *API
}

// NetworkInterfaceAttachmentLocator builds a locator from the given href.
func (api *API) NetworkInterfaceAttachmentLocator(href string) *NetworkInterfaceAttachmentLocator {
	return &NetworkInterfaceAttachmentLocator{Href(href), api}
}

//===== Actions

// GET /api/network_interface_attachments
//
// Currently not implemented.
func (loc *NetworkInterfaceAttachmentLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkInterfaceAttachment", "index")
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

// GET /api/network_interface_attachments/:id
//
// Currently not implemented.
func (loc *NetworkInterfaceAttachmentLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("NetworkInterfaceAttachment", "show")
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

// Security Groups represent network security profiles that contain lists
// of firewall rules for different ports and source IP addresses, as well
// as trust relationships between security groups.
type SecurityGroup struct {
	Description string              `json:"description,omitempty"`
	Href        string              `json:"href,omitempty"`
	Id          string              `json:"id,omitempty"`
	Kind        string              `json:"kind,omitempty"`
	LegacyId    int                 `json:"legacy_id,omitempty"`
	Links       *SecurityGroupLinks `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *SecurityGroup) Locator(api *API) *SecurityGroupLocator {
	return api.SecurityGroupLocator(r.Href)
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

// GET /api/security_groups
// GET /api/clouds/:cloud_id/security_groups
// GET /api/clouds/:cloud_id/instances/:instance_id/security_groups
//
// Currently not implemented.
func (loc *SecurityGroupLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroup", "index")
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

// GET /api/security_groups/:id
// GET /api/clouds/:cloud_id/security_groups/:id
//
// Currently not implemented.
func (loc *SecurityGroupLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SecurityGroup", "show")
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

// Servers represent the notion of a server/machine from RightScale's
// perspective. A Server, does not always have a corresponding VM running
// or provisioned in a cloud. Some clouds use the word "servers" to refer
// to created VMs. These allocated VMs are not called Servers in the
// RightScale API, they are called Instances.
// A Server always has a next_instance association, which will define the
// configuration to apply to a new instance when the server is launched or
// started (starting servers is not yet supported through this API). Once
// a Server is launched/started, a current_instance relationship will
// exist.  Accessing the current_instance of a server results in immediate
// runtime modification of this running server. Changes to the
// next_instance association prepares the configuration for the next
// instance launch/start (therefore they have no effect until such
// operation is performed).
type Server struct {
	Actions         []string     `json:"actions,omitempty"`
	CurrentInstance *Instance    `json:"current_instance,omitempty"`
	Description     string       `json:"description,omitempty"`
	Href            string       `json:"href,omitempty"`
	Id              int          `json:"id,omitempty"`
	Instance        *Instance    `json:"instance,omitempty"`
	Kind            string       `json:"kind,omitempty"`
	Links           *ServerLinks `json:"links,omitempty"`
	Name            string       `json:"name,omitempty"`
	NextInstance    *Instance    `json:"next_instance,omitempty"`
	Tags            []string     `json:"tags,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Server) Locator(api *API) *ServerLocator {
	return api.ServerLocator(r.Href)
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

// GET /api/servers
//
// Currently not implemented.
func (loc *ServerLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "index")
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
//
// Currently not implemented.
func (loc *ServerLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Server", "show")
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

// A server array represents a logical group of instances and allows to
// resize(grow/shrink) that group based on certain elasticity parameters.
// A server array just like a server always has a next_instance
// association, which will define the configuration to apply when a new
// instance is launched. But unlike a server which has a current_instance
// relationship, the server array has a current_instances relationship
// that gives the information about all the running instances in the
// array.  Changes to the next_instance association prepares the
// configuration for the next instance that is to be launched in the array
// and will therefore not affect any of the currently running instances.
type ServerArray struct {
	Actions         []string               `json:"actions,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Href            string                 `json:"href,omitempty"`
	Id              int                    `json:"id,omitempty"`
	InstanceSummary *InstanceSummaryStruct `json:"instance_summary,omitempty"`
	Kind            string                 `json:"kind,omitempty"`
	Links           *ServerArrayLinks      `json:"links,omitempty"`
	Name            string                 `json:"name,omitempty"`
	NextInstance    *Instance              `json:"next_instance,omitempty"`
	State           string                 `json:"state,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
}

// Locator returns a locator for the given resource
func (r *ServerArray) Locator(api *API) *ServerArrayLocator {
	return api.ServerArrayLocator(r.Href)
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

// GET /api/server_arrays
//
// Currently not implemented.
func (loc *ServerArrayLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerArray", "index")
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
//
// Currently not implemented.
func (loc *ServerArrayLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerArray", "show")
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

// ServerTemplates allow you to pre-configure servers by starting from a
// base image and adding scripts that run during the boot, operational,
// and shutdown phases. A ServerTemplate is a description of how a new
// instance will be configured when it is provisioned by your cloud
// provider.
// All revisions of a ServerTemplate belong to a ServerTemplate lineage
// that is exposed by the "lineage" attribute. (NOTE: This attribute is
// merely a string to locate all revisions of a ServerTemplate and NOT a
// working URL)
type ServerTemplate struct {
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          int    `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	Version     int    `json:"version,omitempty"`
}

// Locator returns a locator for the given resource
func (r *ServerTemplate) Locator(api *API) *ServerTemplateLocator {
	return api.ServerTemplateLocator(r.Href)
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

// GET /api/server_templates
//
// Currently not implemented.
func (loc *ServerTemplateLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplate", "index")
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

// GET /api/server_templates/:id
//
// Currently not implemented.
func (loc *ServerTemplateLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("ServerTemplate", "show")
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

/******  SshKey ******/

// Ssh Keys represent a created SSH Key that exists in the cloud.
type SshKey struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          string `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	LegacyId    int    `json:"legacy_id,omitempty"`
	ResourceUid string `json:"resource_uid,omitempty"`
}

// Locator returns a locator for the given resource
func (r *SshKey) Locator(api *API) *SshKeyLocator {
	return api.SshKeyLocator(r.Href)
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

// GET /api/ssh_keys
// GET /api/clouds/:cloud_id/ssh_keys
//
// Currently not implemented.
func (loc *SshKeyLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SshKey", "index")
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

// GET /api/ssh_keys/:id
// GET /api/clouds/:cloud_id/ssh_keys/:id
//
// Currently not implemented.
func (loc *SshKeyLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("SshKey", "show")
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

/******  Subnet ******/

// A Subnet is a logical grouping of network devices. An Instance can have
// many Subnets.
type Subnet struct {
	Description string       `json:"description,omitempty"`
	Href        string       `json:"href,omitempty"`
	Id          string       `json:"id,omitempty"`
	Kind        string       `json:"kind,omitempty"`
	LegacyId    int          `json:"legacy_id,omitempty"`
	Links       *SubnetLinks `json:"links,omitempty"`
	Name        string       `json:"name,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Subnet) Locator(api *API) *SubnetLocator {
	return api.SubnetLocator(r.Href)
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

// GET /api/subnets
// GET /api/clouds/:cloud_id/subnets
// GET /api/clouds/:cloud_id/instances/:instance_id/subnets
//
// Currently not implemented.
func (loc *SubnetLocator) Index(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Subnet", "index")
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

// GET /api/subnets/:id
// GET /api/clouds/:cloud_id/subnets/:id
//
// Currently not implemented.
func (loc *SubnetLocator) Show(options rsapi.APIParams) error {
	var params rsapi.APIParams
	params = rsapi.APIParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.APIParams
	uri, err := loc.ActionPath("Subnet", "show")
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

type AccountParam struct {
	Href string `json:"href,omitempty"`
	Id   int    `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

type CloudParam struct {
	CloudType   string `json:"cloud_type,omitempty"`
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          int    `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
}

type DatacenterLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type DatacenterParam struct {
	Description string           `json:"description,omitempty"`
	Href        string           `json:"href,omitempty"`
	Id          string           `json:"id,omitempty"`
	Kind        string           `json:"kind,omitempty"`
	LegacyId    int              `json:"legacy_id,omitempty"`
	Links       *DatacenterLinks `json:"links,omitempty"`
	Name        string           `json:"name,omitempty"`
}

type DeploymentLinks struct {
	Account *AccountParam `json:"account,omitempty"`
}

type DeploymentParam struct {
	Description  string              `json:"description,omitempty"`
	Href         string              `json:"href,omitempty"`
	Id           int                 `json:"id,omitempty"`
	Instances    []*InstanceParam    `json:"instances,omitempty"`
	Kind         string              `json:"kind,omitempty"`
	Links        *DeploymentLinks    `json:"links,omitempty"`
	Locked       bool                `json:"locked,omitempty"`
	Name         string              `json:"name,omitempty"`
	ServerArrays []*ServerArrayParam `json:"server_arrays,omitempty"`
	Servers      []*ServerParam      `json:"servers,omitempty"`
	Tags         []string            `json:"tags,omitempty"`
}

type ImageLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type ImageParam struct {
	CpuArchitecture    string      `json:"cpu_architecture,omitempty"`
	Description        string      `json:"description,omitempty"`
	Href               string      `json:"href,omitempty"`
	Id                 string      `json:"id,omitempty"`
	ImageType          string      `json:"image_type,omitempty"`
	InheritedSource    string      `json:"inherited_source,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	LegacyId           int         `json:"legacy_id,omitempty"`
	Links              *ImageLinks `json:"links,omitempty"`
	Name               string      `json:"name,omitempty"`
	OsPlatform         string      `json:"os_platform,omitempty"`
	ResourceUid        string      `json:"resource_uid,omitempty"`
	RootDeviceStorage  string      `json:"root_device_storage,omitempty"`
	VirtualizationType string      `json:"virtualization_type,omitempty"`
	Visibility         string      `json:"visibility,omitempty"`
}

type Incarnator struct {
	Href  string `json:"href,omitempty"`
	Id    int    `json:"id,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
}

type InstanceLinks struct {
	Account                 *AccountParam            `json:"account,omitempty"`
	Cloud                   *CloudParam              `json:"cloud,omitempty"`
	ComputedImage           *ImageParam              `json:"computed_image,omitempty"`
	ComputedMultiCloudImage *MultiCloudImageParam    `json:"computed_multi_cloud_image,omitempty"`
	Datacenter              *DatacenterParam         `json:"datacenter,omitempty"`
	Deployment              *DeploymentParam         `json:"deployment,omitempty"`
	Image                   *ImageParam              `json:"image,omitempty"`
	Incarnator              *Incarnator              `json:"incarnator,omitempty"`
	InstanceType            *InstanceTypeParam       `json:"instance_type,omitempty"`
	MultiCloudImage         *MultiCloudImageParam    `json:"multi_cloud_image,omitempty"`
	SecurityGroups          *SecurityGroupCollection `json:"security_groups,omitempty"`
	SshKey                  *SshKeyParam             `json:"ssh_key,omitempty"`
	Subnets                 *SubnetCollection        `json:"subnets,omitempty"`
}

type InstanceParam struct {
	Actions            []string                   `json:"actions,omitempty"`
	Description        string                     `json:"description,omitempty"`
	Href               string                     `json:"href,omitempty"`
	Id                 string                     `json:"id,omitempty"`
	IpAddresses        []*IpAddressParam          `json:"ip_addresses,omitempty"`
	IsNext             bool                       `json:"is_next,omitempty"`
	Kind               string                     `json:"kind,omitempty"`
	LegacyId           int                        `json:"legacy_id,omitempty"`
	Links              *InstanceLinks             `json:"links,omitempty"`
	Locked             bool                       `json:"locked,omitempty"`
	MonitoringId       string                     `json:"monitoring_id,omitempty"`
	MonitoringServer   string                     `json:"monitoring_server,omitempty"`
	MonitoringToken    string                     `json:"monitoring_token,omitempty"`
	Name               string                     `json:"name,omitempty"`
	Networks           []*NetworkParam            `json:"networks,omitempty"`
	OsPlatform         string                     `json:"os_platform,omitempty"`
	PrivateDnsNames    []string                   `json:"private_dns_names,omitempty"`
	PrivateIpAddresses []string                   `json:"private_ip_addresses,omitempty"`
	PublicDnsNames     []string                   `json:"public_dns_names,omitempty"`
	PublicIpAddresses  []string                   `json:"public_ip_addresses,omitempty"`
	ResourceUid        string                     `json:"resource_uid,omitempty"`
	SecurityGroups     *SecurityGroupCollection   `json:"security_groups,omitempty"`
	ServerTemplate     *ServerTemplateParam       `json:"server_template,omitempty"`
	SshHost            string                     `json:"ssh_host,omitempty"`
	State              string                     `json:"state,omitempty"`
	Subnets            *SubnetCollection          `json:"subnets,omitempty"`
	Tags               []string                   `json:"tags,omitempty"`
	Timestamps         *InstancesTimestampsStruct `json:"timestamps,omitempty"`
}

type InstanceSummaryStruct struct {
	HealthyCount       int `json:"healthy_count,omitempty"`
	NotTerminatedCount int `json:"not_terminated_count,omitempty"`
	TotalCount         int `json:"total_count,omitempty"`
	UnhealthyCount     int `json:"unhealthy_count,omitempty"`
}

type InstanceTypeLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type InstanceTypeParam struct {
	CpuCount    int                `json:"cpu_count,omitempty"`
	CpuSpeed    string             `json:"cpu_speed,omitempty"`
	Description string             `json:"description,omitempty"`
	Href        string             `json:"href,omitempty"`
	Id          string             `json:"id,omitempty"`
	Kind        string             `json:"kind,omitempty"`
	LegacyId    int                `json:"legacy_id,omitempty"`
	Links       *InstanceTypeLinks `json:"links,omitempty"`
	Memory      string             `json:"memory,omitempty"`
	Name        string             `json:"name,omitempty"`
}

type InstancesLinksDeploymentServerArraysInstanceSummaryStruct struct {
	HealthyCount       int `json:"healthy_count,omitempty"`
	NotTerminatedCount int `json:"not_terminated_count,omitempty"`
	TotalCount         int `json:"total_count,omitempty"`
	UnhealthyCount     int `json:"unhealthy_count,omitempty"`
}

type InstancesTimestampsStruct struct {
	BootedAt      *time.Time `json:"booted_at,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	OperationalAt *time.Time `json:"operational_at,omitempty"`
	PendingAt     *time.Time `json:"pending_at,omitempty"`
	StrandedAt    *time.Time `json:"stranded_at,omitempty"`
	TerminatedAt  *time.Time `json:"terminated_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

type IpAddressBindingLinks struct {
	Cloud     *CloudParam     `json:"cloud,omitempty"`
	Instance  *InstanceParam  `json:"instance,omitempty"`
	IpAddress *IpAddressParam `json:"ip_address,omitempty"`
}

type IpAddressLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type IpAddressParam struct {
	Address  string          `json:"address,omitempty"`
	Href     string          `json:"href,omitempty"`
	Id       string          `json:"id,omitempty"`
	Kind     string          `json:"kind,omitempty"`
	LegacyId int             `json:"legacy_id,omitempty"`
	Links    *IpAddressLinks `json:"links,omitempty"`
	Name     string          `json:"name,omitempty"`
}

type MultiCloudImageParam struct {
	Description     string `json:"description,omitempty"`
	Href            string `json:"href,omitempty"`
	Id              int    `json:"id,omitempty"`
	InheritedSource string `json:"inherited_source,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Version         int    `json:"version,omitempty"`
}

type NetworkInterfaceAttachmentLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type NetworkInterfaceLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type NetworkLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type NetworkParam struct {
	Description string        `json:"description,omitempty"`
	Href        string        `json:"href,omitempty"`
	Id          string        `json:"id,omitempty"`
	Kind        string        `json:"kind,omitempty"`
	LegacyId    int           `json:"legacy_id,omitempty"`
	Links       *NetworkLinks `json:"links,omitempty"`
	Name        string        `json:"name,omitempty"`
}

type SecurityGroupCollection struct {
	Count int    `json:"count,omitempty"`
	Href  string `json:"href,omitempty"`
}

type SecurityGroupLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type ServerArrayLinks struct {
	Account          *AccountParam    `json:"account,omitempty"`
	Cloud            *CloudParam      `json:"cloud,omitempty"`
	CurrentInstances []*InstanceParam `json:"current_instances,omitempty"`
	NextInstance     *InstanceParam   `json:"next_instance,omitempty"`
}

type ServerArrayParam struct {
	Actions         []string                                                   `json:"actions,omitempty"`
	Description     string                                                     `json:"description,omitempty"`
	Href            string                                                     `json:"href,omitempty"`
	Id              int                                                        `json:"id,omitempty"`
	InstanceSummary *InstancesLinksDeploymentServerArraysInstanceSummaryStruct `json:"instance_summary,omitempty"`
	Kind            string                                                     `json:"kind,omitempty"`
	Links           *ServerArrayLinks                                          `json:"links,omitempty"`
	Name            string                                                     `json:"name,omitempty"`
	NextInstance    *InstanceParam                                             `json:"next_instance,omitempty"`
	State           string                                                     `json:"state,omitempty"`
	Tags            []string                                                   `json:"tags,omitempty"`
}

type ServerLinks struct {
	Account         *AccountParam  `json:"account,omitempty"`
	Cloud           *CloudParam    `json:"cloud,omitempty"`
	CurrentInstance *InstanceParam `json:"current_instance,omitempty"`
	NextInstance    *InstanceParam `json:"next_instance,omitempty"`
}

type ServerParam struct {
	Actions         []string       `json:"actions,omitempty"`
	CurrentInstance *InstanceParam `json:"current_instance,omitempty"`
	Description     string         `json:"description,omitempty"`
	Href            string         `json:"href,omitempty"`
	Id              int            `json:"id,omitempty"`
	Instance        *InstanceParam `json:"instance,omitempty"`
	Kind            string         `json:"kind,omitempty"`
	Links           *ServerLinks   `json:"links,omitempty"`
	Name            string         `json:"name,omitempty"`
	NextInstance    *InstanceParam `json:"next_instance,omitempty"`
	Tags            []string       `json:"tags,omitempty"`
}

type ServerTemplateParam struct {
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          int    `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	Version     int    `json:"version,omitempty"`
}

type SshKeyParam struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          string `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	LegacyId    int    `json:"legacy_id,omitempty"`
	ResourceUid string `json:"resource_uid,omitempty"`
}

type SubnetCollection struct {
	Count int    `json:"count,omitempty"`
	Href  string `json:"href,omitempty"`
}

type SubnetLinks struct {
	Cloud *CloudParam `json:"cloud,omitempty"`
}

type TimestampsStruct struct {
	BootedAt      *time.Time `json:"booted_at,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	OperationalAt *time.Time `json:"operational_at,omitempty"`
	PendingAt     *time.Time `json:"pending_at,omitempty"`
	StrandedAt    *time.Time `json:"stranded_at,omitempty"`
	TerminatedAt  *time.Time `json:"terminated_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
