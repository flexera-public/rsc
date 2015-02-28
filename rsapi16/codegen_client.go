//************************************************************************//
//                     RightScale API client
//
// Generated
// Feb 27, 2015 at 9:39pm (PST)
// Command:
// $ praxisgen -metadata=api_docs -output=. -pkg=rsapi16 -target=1.6 -client=Api16
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi16

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// Url resolver produces an action URL and HTTP method from its name and a given resource href.
// The algorithm consists of first extracting the variables from the href and then substituing them
// in the action path. If there are more than one action paths then the algorithm picks the one that
// can substitute the most variables.
type UrlResolver string

func (r *UrlResolver) Url(rName, aName string) (*metadata.ActionPath, error) {
	var res, ok = GenMetadata[rName]
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

// Resources in RightScale generally belong to accounts. Users can have
// any number of accounts, but when performing an action, a user is
// operating under a particular account.
type Account struct {
	Href string `json:"href,omitempty"`
	Id   int    `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

//===== Locator

// Account resource locator, exposes resource actions.
type AccountLocator struct {
	UrlResolver
	api *Api16
}

// Account resource locator factory
func (api *Api16) AccountLocator(href string) *AccountLocator {
	return &AccountLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/accounts
// Currently not implemented.
func (loc *AccountLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Account", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/accounts/:id
// Currently not implemented.
func (loc *AccountLocator) Show(id int, options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Account", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// Cloud resource locator, exposes resource actions.
type CloudLocator struct {
	UrlResolver
	api *Api16
}

// Cloud resource locator factory
func (api *Api16) CloudLocator(href string) *CloudLocator {
	return &CloudLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds
// Currently not implemented.
func (loc *CloudLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Cloud", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:id
// Currently not implemented.
func (loc *CloudLocator) Show(id int, options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Cloud", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// Datacenter resource locator, exposes resource actions.
type DatacenterLocator struct {
	UrlResolver
	api *Api16
}

// Datacenter resource locator factory
func (api *Api16) DatacenterLocator(href string) *DatacenterLocator {
	return &DatacenterLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/datacenters
// GET /api/clouds/:cloud_id/datacenters
// Currently not implemented.
func (loc *DatacenterLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Datacenter", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/datacenters/:id
// GET /api/clouds/:cloud_id/datacenters/:id
// Currently not implemented.
func (loc *DatacenterLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Datacenter", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// Deployment resource locator, exposes resource actions.
type DeploymentLocator struct {
	UrlResolver
	api *Api16
}

// Deployment resource locator factory
func (api *Api16) DeploymentLocator(href string) *DeploymentLocator {
	return &DeploymentLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/deployments
// List all Deployments in an Account.
func (loc *DeploymentLocator) Index(options rsapi.ApiParams) (*Deployment, error) {
	var res *Deployment
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids"] = idsOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Deployment", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
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

// GET /api/deployments/:id
// Show a single Deployment
func (loc *DeploymentLocator) Show(id int, options rsapi.ApiParams) (*Deployment, error) {
	var res *Deployment
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Deployment", "show")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
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

//===== Locator

// Image resource locator, exposes resource actions.
type ImageLocator struct {
	UrlResolver
	api *Api16
}

// Image resource locator factory
func (api *Api16) ImageLocator(href string) *ImageLocator {
	return &ImageLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/images
// Lists all Images for the given Cloud.
func (loc *ImageLocator) Index(options rsapi.ApiParams) (*Image, error) {
	var res *Image
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Image", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
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

// GET /api/images/:id
// GET /api/clouds/:cloud_id/images/:id
// Currently not implemented.
func (loc *ImageLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Image", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// InstanceType resource locator, exposes resource actions.
type InstanceTypeLocator struct {
	UrlResolver
	api *Api16
}

// InstanceType resource locator factory
func (api *Api16) InstanceTypeLocator(href string) *InstanceTypeLocator {
	return &InstanceTypeLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/instance_types
// GET /api/clouds/:cloud_id/instance_types
// Currently not implemented.
func (loc *InstanceTypeLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("InstanceType", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/instance_types/:id
// GET /api/clouds/:cloud_id/instance_types/:id
// Currently not implemented.
func (loc *InstanceTypeLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("InstanceType", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// Instance resource locator, exposes resource actions.
type InstanceLocator struct {
	UrlResolver
	api *Api16
}

// Instance resource locator factory
func (api *Api16) InstanceLocator(href string) *InstanceLocator {
	return &InstanceLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/instances
// GET /api/clouds/:cloud_id/instances
// List all Instances in an account.
func (loc *InstanceLocator) Index(options rsapi.ApiParams) (*Instance, error) {
	var res *Instance
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		queryParams["filter"] = filterOpt
	}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		queryParams["ids"] = idsOpt
	}
	var limitOpt = options["limit"]
	if limitOpt != nil {
		queryParams["limit"] = limitOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Instance", "index")
	if err != nil {
		return res, err
	}
	var resp, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
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

// GET /api/instances/:id
// GET /api/clouds/:cloud_id/instances/:id
// Currently not implemented.
func (loc *InstanceLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Instance", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// IpAddressBinding resource locator, exposes resource actions.
type IpAddressBindingLocator struct {
	UrlResolver
	api *Api16
}

// IpAddressBinding resource locator factory
func (api *Api16) IpAddressBindingLocator(href string) *IpAddressBindingLocator {
	return &IpAddressBindingLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/ip_address_bindings
// GET /api/clouds/:cloud_id/ip_address_bindings
// Currently not implemented.
func (loc *IpAddressBindingLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("IpAddressBinding", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/ip_address_bindings/:id
// GET /api/clouds/:cloud_id/ip_address_bindings/:id
// Currently not implemented.
func (loc *IpAddressBindingLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("IpAddressBinding", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// IpAddress resource locator, exposes resource actions.
type IpAddressLocator struct {
	UrlResolver
	api *Api16
}

// IpAddress resource locator factory
func (api *Api16) IpAddressLocator(href string) *IpAddressLocator {
	return &IpAddressLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/ip_addresses
// GET /api/clouds/:cloud_id/ip_addresses
// Currently not implemented.
func (loc *IpAddressLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("IpAddress", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/ip_addresses/:id
// GET /api/clouds/:cloud_id/ip_addresses/:id
// Currently not implemented.
func (loc *IpAddressLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("IpAddress", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// MultiCloudImage resource locator, exposes resource actions.
type MultiCloudImageLocator struct {
	UrlResolver
	api *Api16
}

// MultiCloudImage resource locator factory
func (api *Api16) MultiCloudImageLocator(href string) *MultiCloudImageLocator {
	return &MultiCloudImageLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/multi_cloud_images
// Currently not implemented.
func (loc *MultiCloudImageLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("MultiCloudImage", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/multi_cloud_images/:id
// Currently not implemented.
func (loc *MultiCloudImageLocator) Show(id int, options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("MultiCloudImage", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// NetworkInterfaceAttachment resource locator, exposes resource actions.
type NetworkInterfaceAttachmentLocator struct {
	UrlResolver
	api *Api16
}

// NetworkInterfaceAttachment resource locator factory
func (api *Api16) NetworkInterfaceAttachmentLocator(href string) *NetworkInterfaceAttachmentLocator {
	return &NetworkInterfaceAttachmentLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/network_interface_attachments
// Currently not implemented.
func (loc *NetworkInterfaceAttachmentLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("NetworkInterfaceAttachment", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/network_interface_attachments/:id
// Currently not implemented.
func (loc *NetworkInterfaceAttachmentLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("NetworkInterfaceAttachment", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// NetworkInterface resource locator, exposes resource actions.
type NetworkInterfaceLocator struct {
	UrlResolver
	api *Api16
}

// NetworkInterface resource locator factory
func (api *Api16) NetworkInterfaceLocator(href string) *NetworkInterfaceLocator {
	return &NetworkInterfaceLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/network_interfaces
// Currently not implemented.
func (loc *NetworkInterfaceLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("NetworkInterface", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/network_interfaces/:id
// Currently not implemented.
func (loc *NetworkInterfaceLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("NetworkInterface", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// Network resource locator, exposes resource actions.
type NetworkLocator struct {
	UrlResolver
	api *Api16
}

// Network resource locator factory
func (api *Api16) NetworkLocator(href string) *NetworkLocator {
	return &NetworkLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/networks
// Currently not implemented.
func (loc *NetworkLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Network", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/networks/:id
// Currently not implemented.
func (loc *NetworkLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Network", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// SecurityGroup resource locator, exposes resource actions.
type SecurityGroupLocator struct {
	UrlResolver
	api *Api16
}

// SecurityGroup resource locator factory
func (api *Api16) SecurityGroupLocator(href string) *SecurityGroupLocator {
	return &SecurityGroupLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/security_groups
// GET /api/clouds/:cloud_id/security_groups
// GET /api/clouds/:cloud_id/instances/:instance_id/security_groups
// Currently not implemented.
func (loc *SecurityGroupLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("SecurityGroup", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/security_groups/:id
// GET /api/clouds/:cloud_id/security_groups/:id
// Currently not implemented.
func (loc *SecurityGroupLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("SecurityGroup", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// ServerArray resource locator, exposes resource actions.
type ServerArrayLocator struct {
	UrlResolver
	api *Api16
}

// ServerArray resource locator factory
func (api *Api16) ServerArrayLocator(href string) *ServerArrayLocator {
	return &ServerArrayLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/server_arrays
// Currently not implemented.
func (loc *ServerArrayLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("ServerArray", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_arrays/:id
// Currently not implemented.
func (loc *ServerArrayLocator) Show(id int, options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("ServerArray", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// ServerTemplate resource locator, exposes resource actions.
type ServerTemplateLocator struct {
	UrlResolver
	api *Api16
}

// ServerTemplate resource locator factory
func (api *Api16) ServerTemplateLocator(href string) *ServerTemplateLocator {
	return &ServerTemplateLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/server_templates
// Currently not implemented.
func (loc *ServerTemplateLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("ServerTemplate", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_templates/:id
// Currently not implemented.
func (loc *ServerTemplateLocator) Show(id int, options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("ServerTemplate", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// Server resource locator, exposes resource actions.
type ServerLocator struct {
	UrlResolver
	api *Api16
}

// Server resource locator factory
func (api *Api16) ServerLocator(href string) *ServerLocator {
	return &ServerLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/servers
// Currently not implemented.
func (loc *ServerLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Server", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/servers/:id
// Currently not implemented.
func (loc *ServerLocator) Show(id int, options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Server", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// SshKey resource locator, exposes resource actions.
type SshKeyLocator struct {
	UrlResolver
	api *Api16
}

// SshKey resource locator factory
func (api *Api16) SshKeyLocator(href string) *SshKeyLocator {
	return &SshKeyLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/ssh_keys
// GET /api/clouds/:cloud_id/ssh_keys
// Currently not implemented.
func (loc *SshKeyLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("SshKey", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/ssh_keys/:id
// GET /api/clouds/:cloud_id/ssh_keys/:id
// Currently not implemented.
func (loc *SshKeyLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("SshKey", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
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

//===== Locator

// Subnet resource locator, exposes resource actions.
type SubnetLocator struct {
	UrlResolver
	api *Api16
}

// Subnet resource locator factory
func (api *Api16) SubnetLocator(href string) *SubnetLocator {
	return &SubnetLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/subnets
// GET /api/clouds/:cloud_id/subnets
// GET /api/clouds/:cloud_id/instances/:instance_id/subnets
// Currently not implemented.
func (loc *SubnetLocator) Index(options rsapi.ApiParams) error {
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Subnet", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/subnets/:id
// GET /api/clouds/:cloud_id/subnets/:id
// Currently not implemented.
func (loc *SubnetLocator) Show(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		queryParams["view"] = viewOpt
	}
	var payloadParams rsapi.ApiParams
	var uri, err = loc.Url("Subnet", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err2 != nil {
		return err2
	}
	return nil
}

/****** Parameter Data Types ******/

type Account2 struct {
	Href *string `json:"href,omitempty"`
	Id   *int    `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Cloud2 struct {
	CloudType   *string `json:"cloud_type,omitempty"`
	Description *string `json:"description,omitempty"`
	Href        *string `json:"href,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Datacenter2 struct {
	Description *string          `json:"description,omitempty"`
	Href        *string          `json:"href,omitempty"`
	Id          *string          `json:"id,omitempty"`
	Kind        *string          `json:"kind,omitempty"`
	LegacyId    *int             `json:"legacy_id,omitempty"`
	Links       *DatacenterLinks `json:"links,omitempty"`
	Name        *string          `json:"name,omitempty"`
}

type DatacenterLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
}

type Deployment2 struct {
	Description  *string          `json:"description,omitempty"`
	Href         *string          `json:"href,omitempty"`
	Id           *int             `json:"id,omitempty"`
	Instances    []*Instance2     `json:"instances,omitempty"`
	Kind         *string          `json:"kind,omitempty"`
	Links        *DeploymentLinks `json:"links,omitempty"`
	Locked       *bool            `json:"locked,omitempty"`
	Name         *string          `json:"name,omitempty"`
	ServerArrays []*ServerArray2  `json:"server_arrays,omitempty"`
	Servers      []*Server2       `json:"servers,omitempty"`
	Tags         []string         `json:"tags,omitempty"`
}

type DeploymentLinks struct {
	Account *Account2 `json:"account,omitempty"`
}

type Image2 struct {
	CpuArchitecture    *string     `json:"cpu_architecture,omitempty"`
	Description        *string     `json:"description,omitempty"`
	Href               *string     `json:"href,omitempty"`
	Id                 *string     `json:"id,omitempty"`
	ImageType          *string     `json:"image_type,omitempty"`
	InheritedSource    *string     `json:"inherited_source,omitempty"`
	Kind               *string     `json:"kind,omitempty"`
	LegacyId           *int        `json:"legacy_id,omitempty"`
	Links              *ImageLinks `json:"links,omitempty"`
	Name               *string     `json:"name,omitempty"`
	OsPlatform         *string     `json:"os_platform,omitempty"`
	ResourceUid        *string     `json:"resource_uid,omitempty"`
	RootDeviceStorage  *string     `json:"root_device_storage,omitempty"`
	VirtualizationType *string     `json:"virtualization_type,omitempty"`
	Visibility         *string     `json:"visibility,omitempty"`
}

type ImageLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
}

type Incarnator struct {
	Href  *string `json:"href,omitempty"`
	Id    *int    `json:"id,omitempty"`
	Kind  *string `json:"kind,omitempty"`
	Name  *string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
}

type Instance2 struct {
	Actions            []string                   `json:"actions,omitempty"`
	Description        *string                    `json:"description,omitempty"`
	Href               *string                    `json:"href,omitempty"`
	Id                 *string                    `json:"id,omitempty"`
	IpAddresses        []*IpAddress2              `json:"ip_addresses,omitempty"`
	IsNext             *bool                      `json:"is_next,omitempty"`
	Kind               *string                    `json:"kind,omitempty"`
	LegacyId           *int                       `json:"legacy_id,omitempty"`
	Links              *InstanceLinks             `json:"links,omitempty"`
	Locked             *bool                      `json:"locked,omitempty"`
	MonitoringId       *string                    `json:"monitoring_id,omitempty"`
	MonitoringServer   *string                    `json:"monitoring_server,omitempty"`
	MonitoringToken    *string                    `json:"monitoring_token,omitempty"`
	Name               *string                    `json:"name,omitempty"`
	Networks           []*Network2                `json:"networks,omitempty"`
	OsPlatform         *string                    `json:"os_platform,omitempty"`
	PrivateDnsNames    []string                   `json:"private_dns_names,omitempty"`
	PrivateIpAddresses []string                   `json:"private_ip_addresses,omitempty"`
	PublicDnsNames     []string                   `json:"public_dns_names,omitempty"`
	PublicIpAddresses  []string                   `json:"public_ip_addresses,omitempty"`
	ResourceUid        *string                    `json:"resource_uid,omitempty"`
	SecurityGroups     *SecurityGroupCollection   `json:"security_groups,omitempty"`
	ServerTemplate     *ServerTemplate2           `json:"server_template,omitempty"`
	SshHost            *string                    `json:"ssh_host,omitempty"`
	State              *string                    `json:"state,omitempty"`
	Subnets            *SubnetCollection          `json:"subnets,omitempty"`
	Tags               []string                   `json:"tags,omitempty"`
	Timestamps         *InstancesTimestampsStruct `json:"timestamps,omitempty"`
}

type InstanceLinks struct {
	Account                 *Account2                `json:"account,omitempty"`
	Cloud                   *Cloud2                  `json:"cloud,omitempty"`
	ComputedImage           *Image2                  `json:"computed_image,omitempty"`
	ComputedMultiCloudImage *MultiCloudImage2        `json:"computed_multi_cloud_image,omitempty"`
	Datacenter              *Datacenter2             `json:"datacenter,omitempty"`
	Deployment              *Deployment2             `json:"deployment,omitempty"`
	Image                   *Image2                  `json:"image,omitempty"`
	Incarnator              *Incarnator              `json:"incarnator,omitempty"`
	InstanceType            *InstanceType2           `json:"instance_type,omitempty"`
	MultiCloudImage         *MultiCloudImage2        `json:"multi_cloud_image,omitempty"`
	SecurityGroups          *SecurityGroupCollection `json:"security_groups,omitempty"`
	SshKey                  *SshKey2                 `json:"ssh_key,omitempty"`
	Subnets                 *SubnetCollection        `json:"subnets,omitempty"`
}

type InstanceSummaryStruct struct {
	HealthyCount       *int `json:"healthy_count,omitempty"`
	NotTerminatedCount *int `json:"not_terminated_count,omitempty"`
	TotalCount         *int `json:"total_count,omitempty"`
	UnhealthyCount     *int `json:"unhealthy_count,omitempty"`
}

type InstanceType2 struct {
	CpuCount    *int               `json:"cpu_count,omitempty"`
	CpuSpeed    *string            `json:"cpu_speed,omitempty"`
	Description *string            `json:"description,omitempty"`
	Href        *string            `json:"href,omitempty"`
	Id          *string            `json:"id,omitempty"`
	Kind        *string            `json:"kind,omitempty"`
	LegacyId    *int               `json:"legacy_id,omitempty"`
	Links       *InstanceTypeLinks `json:"links,omitempty"`
	Memory      *string            `json:"memory,omitempty"`
	Name        *string            `json:"name,omitempty"`
}

type InstanceTypeLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
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

type IpAddress2 struct {
	Address  *string         `json:"address,omitempty"`
	Href     *string         `json:"href,omitempty"`
	Id       *string         `json:"id,omitempty"`
	Kind     *string         `json:"kind,omitempty"`
	LegacyId *int            `json:"legacy_id,omitempty"`
	Links    *IpAddressLinks `json:"links,omitempty"`
	Name     *string         `json:"name,omitempty"`
}

type IpAddressBinding2 struct {
	Href        *string                `json:"href,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	LegacyId    *int                   `json:"legacy_id,omitempty"`
	Links       *IpAddressBindingLinks `json:"links,omitempty"`
	PrivatePort *int                   `json:"private_port,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
	PublicPort  *int                   `json:"public_port,omitempty"`
}

type IpAddressBindingLinks struct {
	Cloud     *Cloud2     `json:"cloud,omitempty"`
	Instance  *Instance2  `json:"instance,omitempty"`
	IpAddress *IpAddress2 `json:"ip_address,omitempty"`
}

type IpAddressLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
}

type MultiCloudImage2 struct {
	Description     *string `json:"description,omitempty"`
	Href            *string `json:"href,omitempty"`
	Id              *int    `json:"id,omitempty"`
	InheritedSource *string `json:"inherited_source,omitempty"`
	Kind            *string `json:"kind,omitempty"`
	Name            *string `json:"name,omitempty"`
	Version         *int    `json:"version,omitempty"`
}

type Network2 struct {
	Description *string       `json:"description,omitempty"`
	Href        *string       `json:"href,omitempty"`
	Id          *string       `json:"id,omitempty"`
	Kind        *string       `json:"kind,omitempty"`
	LegacyId    *int          `json:"legacy_id,omitempty"`
	Links       *NetworkLinks `json:"links,omitempty"`
	Name        *string       `json:"name,omitempty"`
}

type NetworkInterface2 struct {
	Description *string                `json:"description,omitempty"`
	Href        *string                `json:"href,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Links       *NetworkInterfaceLinks `json:"links,omitempty"`
}

type NetworkInterfaceAttachment2 struct {
	Href  *string                          `json:"href,omitempty"`
	Id    *string                          `json:"id,omitempty"`
	Kind  *string                          `json:"kind,omitempty"`
	Links *NetworkInterfaceAttachmentLinks `json:"links,omitempty"`
}

type NetworkInterfaceAttachmentLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
}

type NetworkInterfaceLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
}

type NetworkLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
}

type SecurityGroup2 struct {
	Description *string             `json:"description,omitempty"`
	Href        *string             `json:"href,omitempty"`
	Id          *string             `json:"id,omitempty"`
	Kind        *string             `json:"kind,omitempty"`
	LegacyId    *int                `json:"legacy_id,omitempty"`
	Links       *SecurityGroupLinks `json:"links,omitempty"`
	Name        *string             `json:"name,omitempty"`
}

type SecurityGroupCollection struct {
	Count *int    `json:"count,omitempty"`
	Href  *string `json:"href,omitempty"`
}

type SecurityGroupLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
}

type Server2 struct {
	Actions         []string     `json:"actions,omitempty"`
	CurrentInstance *Instance2   `json:"current_instance,omitempty"`
	Description     *string      `json:"description,omitempty"`
	Href            *string      `json:"href,omitempty"`
	Id              *int         `json:"id,omitempty"`
	Instance        *Instance2   `json:"instance,omitempty"`
	Kind            *string      `json:"kind,omitempty"`
	Links           *ServerLinks `json:"links,omitempty"`
	Name            *string      `json:"name,omitempty"`
	NextInstance    *Instance2   `json:"next_instance,omitempty"`
	Tags            []string     `json:"tags,omitempty"`
}

type ServerArray2 struct {
	Actions         []string                           `json:"actions,omitempty"`
	Description     *string                            `json:"description,omitempty"`
	Href            *string                            `json:"href,omitempty"`
	Id              *int                               `json:"id,omitempty"`
	InstanceSummary *ServerArraysInstanceSummaryStruct `json:"instance_summary,omitempty"`
	Kind            *string                            `json:"kind,omitempty"`
	Links           *ServerArrayLinks                  `json:"links,omitempty"`
	Name            *string                            `json:"name,omitempty"`
	NextInstance    *Instance2                         `json:"next_instance,omitempty"`
	State           *string                            `json:"state,omitempty"`
	Tags            []string                           `json:"tags,omitempty"`
}

type ServerArrayLinks struct {
	Account          *Account2    `json:"account,omitempty"`
	Cloud            *Cloud2      `json:"cloud,omitempty"`
	CurrentInstances []*Instance2 `json:"current_instances,omitempty"`
	NextInstance     *Instance2   `json:"next_instance,omitempty"`
}

type ServerArraysInstanceSummaryStruct struct {
	HealthyCount       *int `json:"healthy_count,omitempty"`
	NotTerminatedCount *int `json:"not_terminated_count,omitempty"`
	TotalCount         *int `json:"total_count,omitempty"`
	UnhealthyCount     *int `json:"unhealthy_count,omitempty"`
}

type ServerLinks struct {
	Account         *Account2  `json:"account,omitempty"`
	Cloud           *Cloud2    `json:"cloud,omitempty"`
	CurrentInstance *Instance2 `json:"current_instance,omitempty"`
	NextInstance    *Instance2 `json:"next_instance,omitempty"`
}

type ServerTemplate2 struct {
	Description *string `json:"description,omitempty"`
	Href        *string `json:"href,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	Name        *string `json:"name,omitempty"`
	Version     *int    `json:"version,omitempty"`
}

type SshKey2 struct {
	Fingerprint *string `json:"fingerprint,omitempty"`
	Href        *string `json:"href,omitempty"`
	Id          *string `json:"id,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	LegacyId    *int    `json:"legacy_id,omitempty"`
	ResourceUid *string `json:"resource_uid,omitempty"`
}

type Subnet2 struct {
	Description *string      `json:"description,omitempty"`
	Href        *string      `json:"href,omitempty"`
	Id          *string      `json:"id,omitempty"`
	Kind        *string      `json:"kind,omitempty"`
	LegacyId    *int         `json:"legacy_id,omitempty"`
	Links       *SubnetLinks `json:"links,omitempty"`
	Name        *string      `json:"name,omitempty"`
}

type SubnetCollection struct {
	Count *int    `json:"count,omitempty"`
	Href  *string `json:"href,omitempty"`
}

type SubnetLinks struct {
	Cloud *Cloud2 `json:"cloud,omitempty"`
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
