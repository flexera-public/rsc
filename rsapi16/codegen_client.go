//************************************************************************//
//                     RightScale API client
//
// Generated
// Feb 25, 2015 at 12:35pm (PST)
// Command:
// $ praxisgen -metadata=../../rsapi16/api_docs -output=../../rsapi16 -pkg=rsapi16 -target=1.6
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
)

// Helper function that merges optional parameters into payload
func mergeOptionals(params, options ApiParams) ApiParams {
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

/******  Accounts ******/

// Resources in RightScale generally belong to accounts. Users can have
// any number of accounts, but when performing an action, a user is
// operating under a particular account.
type Accounts struct {
	Href string `json:"href,omitempty"`
	Id   int    `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

//===== Locator

// Accounts resource locator, exposes resource actions.
type AccountsLocator struct {
	UrlResolver
	api *Api15
}

// Accounts resource locator factory
func (api *Api15) AccountsLocator(href string) *AccountsLocator {
	return &AccountsLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/accounts
// Currently not implemented.
func (loc *AccountsLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("Accounts", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/accounts/:id
// Currently not implemented.
func (loc *AccountsLocator) Show(id int, options ApiParams) error {
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Accounts", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Clouds ******/

// Clouds provide remote resources for things like storage and compute.
// You must have registered a cloud within your account in order to use
// it.
type Clouds struct {
	CloudType   string `json:"cloud_type,omitempty"`
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          int    `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
}

//===== Locator

// Clouds resource locator, exposes resource actions.
type CloudsLocator struct {
	UrlResolver
	api *Api15
}

// Clouds resource locator factory
func (api *Api15) CloudsLocator(href string) *CloudsLocator {
	return &CloudsLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds
// Currently not implemented.
func (loc *CloudsLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("Clouds", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/clouds/:id
// Currently not implemented.
func (loc *CloudsLocator) Show(id int, options ApiParams) error {
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Clouds", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Datacenters ******/

// Datacenters are cloud resources that give you the ability to place
// resources in isolated locations. A carefully designed system placed in
// multiple datacenters can provide fault tolerance when one datacenter
// has a problem.
type Datacenters struct {
	Description string           `json:"description,omitempty"`
	Href        string           `json:"href,omitempty"`
	Id          string           `json:"id,omitempty"`
	Kind        string           `json:"kind,omitempty"`
	LegacyId    int              `json:"legacy_id,omitempty"`
	Links       *DatacenterLinks `json:"links,omitempty"`
	Name        string           `json:"name,omitempty"`
}

//===== Locator

// Datacenters resource locator, exposes resource actions.
type DatacentersLocator struct {
	UrlResolver
	api *Api15
}

// Datacenters resource locator factory
func (api *Api15) DatacentersLocator(href string) *DatacentersLocator {
	return &DatacentersLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/datacenters
// GET /api/clouds/:cloud_id/datacenters
// Currently not implemented.
func (loc *DatacentersLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("Datacenters", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/datacenters/:id
// GET /api/clouds/:cloud_id/datacenters/:id
// Currently not implemented.
func (loc *DatacentersLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Datacenters", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Deployments ******/

// Deployments provide a way to group resources that logically belong
// together.
type Deployments struct {
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

// Deployments resource locator, exposes resource actions.
type DeploymentsLocator struct {
	UrlResolver
	api *Api15
}

// Deployments resource locator factory
func (api *Api15) DeploymentsLocator(href string) *DeploymentsLocator {
	return &DeploymentsLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/deployments
// List all Deployments in an Account.
func (loc *DeploymentsLocator) Index(options ApiParams) (*Deployment, error) {
	var res *Deployment
	var params = mergeOptionals(ApiParams{}, options)
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
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// GET /api/deployments/:id
// Show a single Deployment
func (loc *DeploymentsLocator) Show(id int, options ApiParams) (*Deployment, error) {
	var res *Deployment
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
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

/******  Images ******/

// Images define the initial Operating System and root disk contents
// for new instances.
type Images struct {
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

// Images resource locator, exposes resource actions.
type ImagesLocator struct {
	UrlResolver
	api *Api15
}

// Images resource locator factory
func (api *Api15) ImagesLocator(href string) *ImagesLocator {
	return &ImagesLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/clouds/:cloud_id/images
// Lists all Images for the given Cloud.
func (loc *ImagesLocator) Index(options ApiParams) (*Image, error) {
	var res *Image
	var params = mergeOptionals(ApiParams{}, options)
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
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// GET /api/images/:id
// GET /api/clouds/:cloud_id/images/:id
// Currently not implemented.
func (loc *ImagesLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Images", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  InstanceTypes ******/

// An InstanceType represents a basic hardware configuration for an
// Instance.
// Combining all possible configurations of hardware into a smaller,
// well-known set of options makes instances easier to manage, and allows
// better allocation efficiency into physical hosts.
type InstanceTypes struct {
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

// InstanceTypes resource locator, exposes resource actions.
type InstanceTypesLocator struct {
	UrlResolver
	api *Api15
}

// InstanceTypes resource locator factory
func (api *Api15) InstanceTypesLocator(href string) *InstanceTypesLocator {
	return &InstanceTypesLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/instance_types
// GET /api/clouds/:cloud_id/instance_types
// Currently not implemented.
func (loc *InstanceTypesLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("InstanceTypes", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/instance_types/:id
// GET /api/clouds/:cloud_id/instance_types/:id
// Currently not implemented.
func (loc *InstanceTypesLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("InstanceTypes", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Instances ******/

// Instances represent an entity that is runnable in the cloud.
// An instance of type "next" is a container of information that expresses
// how to configure a future instance when we decide to launch or start
// it. A "next" instance generally only exists in the RightScale realm,
// and usually doesn't have any corresponding representation existing in
// the cloud. However, if an instance is not of type "next", it will
// generally represent an existing running (or provisioned) virtual
// machine existing in the cloud.
type Instances struct {
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
	Timestamps         *StrandedAtStruct        `json:"timestamps,omitempty"`
}

//===== Locator

// Instances resource locator, exposes resource actions.
type InstancesLocator struct {
	UrlResolver
	api *Api15
}

// Instances resource locator factory
func (api *Api15) InstancesLocator(href string) *InstancesLocator {
	return &InstancesLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/instances
// GET /api/clouds/:cloud_id/instances
// List all Instances in an account.
func (loc *InstancesLocator) Index(options ApiParams) (*Instance, error) {
	var res *Instance
	var params = mergeOptionals(ApiParams{}, options)
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
	var err4 = json.Unmarshal(respBody, res)
	return res, err4
}

// GET /api/instances/:id
// GET /api/clouds/:cloud_id/instances/:id
// Currently not implemented.
func (loc *InstancesLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Instances", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  IpAddressBindings ******/

// An IpAddressBinding represents an abstraction for binding an IpAddress
// to an instance. The IpAddress is bound immediately for a current
// instance, or on launch for a next instance.
type IpAddressBindings struct {
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

// IpAddressBindings resource locator, exposes resource actions.
type IpAddressBindingsLocator struct {
	UrlResolver
	api *Api15
}

// IpAddressBindings resource locator factory
func (api *Api15) IpAddressBindingsLocator(href string) *IpAddressBindingsLocator {
	return &IpAddressBindingsLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/ip_address_bindings
// GET /api/clouds/:cloud_id/ip_address_bindings
// Currently not implemented.
func (loc *IpAddressBindingsLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("IpAddressBindings", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/ip_address_bindings/:id
// GET /api/clouds/:cloud_id/ip_address_bindings/:id
// Currently not implemented.
func (loc *IpAddressBindingsLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("IpAddressBindings", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  IpAddresses ******/

// An IpAddress provides an abstraction for IPv4 addresses bindable to
// Instance resources running in a Cloud.
type IpAddresses struct {
	Address  string          `json:"address,omitempty"`
	Href     string          `json:"href,omitempty"`
	Id       string          `json:"id,omitempty"`
	Kind     string          `json:"kind,omitempty"`
	LegacyId int             `json:"legacy_id,omitempty"`
	Links    *IpAddressLinks `json:"links,omitempty"`
	Name     string          `json:"name,omitempty"`
}

//===== Locator

// IpAddresses resource locator, exposes resource actions.
type IpAddressesLocator struct {
	UrlResolver
	api *Api15
}

// IpAddresses resource locator factory
func (api *Api15) IpAddressesLocator(href string) *IpAddressesLocator {
	return &IpAddressesLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/ip_addresses
// GET /api/clouds/:cloud_id/ip_addresses
// Currently not implemented.
func (loc *IpAddressesLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("IpAddresses", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/ip_addresses/:id
// GET /api/clouds/:cloud_id/ip_addresses/:id
// Currently not implemented.
func (loc *IpAddressesLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("IpAddresses", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  MultiCloudImages ******/

// A MultiCloudImage is a RightScale component that functions as a pointer
// to machine images in specific clouds (e.g. AWS US-East, Rackspace).
// Each ServerTemplate can reference many MultiCloudImages that define
// which image should be used when a server is launched in a particular
// cloud.
type MultiCloudImages struct {
	Description     string `json:"description,omitempty"`
	Href            string `json:"href,omitempty"`
	Id              int    `json:"id,omitempty"`
	InheritedSource string `json:"inherited_source,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Version         int    `json:"version,omitempty"`
}

//===== Locator

// MultiCloudImages resource locator, exposes resource actions.
type MultiCloudImagesLocator struct {
	UrlResolver
	api *Api15
}

// MultiCloudImages resource locator factory
func (api *Api15) MultiCloudImagesLocator(href string) *MultiCloudImagesLocator {
	return &MultiCloudImagesLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/multi_cloud_images
// Currently not implemented.
func (loc *MultiCloudImagesLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("MultiCloudImages", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/multi_cloud_images/:id
// Currently not implemented.
func (loc *MultiCloudImagesLocator) Show(id int, options ApiParams) error {
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("MultiCloudImages", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  NetworkInterfaceAttachments ******/

// NetworkInterfaceAttachments represent an attachment between a
// NetworkInterface and another resource.
type NetworkInterfaceAttachments struct {
	Href  string                           `json:"href,omitempty"`
	Id    string                           `json:"id,omitempty"`
	Kind  string                           `json:"kind,omitempty"`
	Links *NetworkInterfaceAttachmentLinks `json:"links,omitempty"`
}

//===== Locator

// NetworkInterfaceAttachments resource locator, exposes resource actions.
type NetworkInterfaceAttachmentsLocator struct {
	UrlResolver
	api *Api15
}

// NetworkInterfaceAttachments resource locator factory
func (api *Api15) NetworkInterfaceAttachmentsLocator(href string) *NetworkInterfaceAttachmentsLocator {
	return &NetworkInterfaceAttachmentsLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/network_interface_attachments
// Currently not implemented.
func (loc *NetworkInterfaceAttachmentsLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("NetworkInterfaceAttachments", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/network_interface_attachments/:id
// Currently not implemented.
func (loc *NetworkInterfaceAttachmentsLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("NetworkInterfaceAttachments", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  NetworkInterfaces ******/

// Just like their physical counterparts, NetworkInterfaces join other
// resources to a network.
type NetworkInterfaces struct {
	Description string                 `json:"description,omitempty"`
	Href        string                 `json:"href,omitempty"`
	Id          string                 `json:"id,omitempty"`
	Kind        string                 `json:"kind,omitempty"`
	Links       *NetworkInterfaceLinks `json:"links,omitempty"`
}

//===== Locator

// NetworkInterfaces resource locator, exposes resource actions.
type NetworkInterfacesLocator struct {
	UrlResolver
	api *Api15
}

// NetworkInterfaces resource locator factory
func (api *Api15) NetworkInterfacesLocator(href string) *NetworkInterfacesLocator {
	return &NetworkInterfacesLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/network_interfaces
// Currently not implemented.
func (loc *NetworkInterfacesLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("NetworkInterfaces", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/network_interfaces/:id
// Currently not implemented.
func (loc *NetworkInterfacesLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("NetworkInterfaces", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Networks ******/

// A Network is a logical grouping of network devices.
type Networks struct {
	Description string        `json:"description,omitempty"`
	Href        string        `json:"href,omitempty"`
	Id          string        `json:"id,omitempty"`
	Kind        string        `json:"kind,omitempty"`
	LegacyId    int           `json:"legacy_id,omitempty"`
	Links       *NetworkLinks `json:"links,omitempty"`
	Name        string        `json:"name,omitempty"`
}

//===== Locator

// Networks resource locator, exposes resource actions.
type NetworksLocator struct {
	UrlResolver
	api *Api15
}

// Networks resource locator factory
func (api *Api15) NetworksLocator(href string) *NetworksLocator {
	return &NetworksLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/networks
// Currently not implemented.
func (loc *NetworksLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("Networks", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/networks/:id
// Currently not implemented.
func (loc *NetworksLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Networks", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  SecurityGroups ******/

// Security Groups represent network security profiles that contain lists
// of firewall rules for different ports and source IP addresses, as well
// as trust relationships between security groups.
type SecurityGroups struct {
	Description string              `json:"description,omitempty"`
	Href        string              `json:"href,omitempty"`
	Id          string              `json:"id,omitempty"`
	Kind        string              `json:"kind,omitempty"`
	LegacyId    int                 `json:"legacy_id,omitempty"`
	Links       *SecurityGroupLinks `json:"links,omitempty"`
	Name        string              `json:"name,omitempty"`
}

//===== Locator

// SecurityGroups resource locator, exposes resource actions.
type SecurityGroupsLocator struct {
	UrlResolver
	api *Api15
}

// SecurityGroups resource locator factory
func (api *Api15) SecurityGroupsLocator(href string) *SecurityGroupsLocator {
	return &SecurityGroupsLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/security_groups
// GET /api/clouds/:cloud_id/security_groups
// GET /api/clouds/:cloud_id/instances/:instance_id/security_groups
// Currently not implemented.
func (loc *SecurityGroupsLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("SecurityGroups", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/security_groups/:id
// GET /api/clouds/:cloud_id/security_groups/:id
// Currently not implemented.
func (loc *SecurityGroupsLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("SecurityGroups", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  ServerArrays ******/

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
type ServerArrays struct {
	Actions         []string          `json:"actions,omitempty"`
	Description     string            `json:"description,omitempty"`
	Href            string            `json:"href,omitempty"`
	Id              int               `json:"id,omitempty"`
	InstanceSummary *TotalCountStruct `json:"instance_summary,omitempty"`
	Kind            string            `json:"kind,omitempty"`
	Links           *ServerArrayLinks `json:"links,omitempty"`
	Name            string            `json:"name,omitempty"`
	NextInstance    *Instance         `json:"next_instance,omitempty"`
	State           string            `json:"state,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
}

//===== Locator

// ServerArrays resource locator, exposes resource actions.
type ServerArraysLocator struct {
	UrlResolver
	api *Api15
}

// ServerArrays resource locator factory
func (api *Api15) ServerArraysLocator(href string) *ServerArraysLocator {
	return &ServerArraysLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/server_arrays
// Currently not implemented.
func (loc *ServerArraysLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("ServerArrays", "index")
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
// Currently not implemented.
func (loc *ServerArraysLocator) Show(id int, options ApiParams) error {
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("ServerArrays", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  ServerTemplates ******/

// ServerTemplates allow you to pre-configure servers by starting from a
// base image and adding scripts that run during the boot, operational,
// and shutdown phases. A ServerTemplate is a description of how a new
// instance will be configured when it is provisioned by your cloud
// provider.
// All revisions of a ServerTemplate belong to a ServerTemplate lineage
// that is exposed by the "lineage" attribute. (NOTE: This attribute is
// merely a string to locate all revisions of a ServerTemplate and NOT a
// working URL)
type ServerTemplates struct {
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          int    `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	Version     int    `json:"version,omitempty"`
}

//===== Locator

// ServerTemplates resource locator, exposes resource actions.
type ServerTemplatesLocator struct {
	UrlResolver
	api *Api15
}

// ServerTemplates resource locator factory
func (api *Api15) ServerTemplatesLocator(href string) *ServerTemplatesLocator {
	return &ServerTemplatesLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/server_templates
// Currently not implemented.
func (loc *ServerTemplatesLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("ServerTemplates", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/server_templates/:id
// Currently not implemented.
func (loc *ServerTemplatesLocator) Show(id int, options ApiParams) error {
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("ServerTemplates", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Servers ******/

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
type Servers struct {
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

// Servers resource locator, exposes resource actions.
type ServersLocator struct {
	UrlResolver
	api *Api15
}

// Servers resource locator factory
func (api *Api15) ServersLocator(href string) *ServersLocator {
	return &ServersLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/servers
// Currently not implemented.
func (loc *ServersLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("Servers", "index")
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
// Currently not implemented.
func (loc *ServersLocator) Show(id int, options ApiParams) error {
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Servers", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  SshKeys ******/

// Ssh Keys represent a created SSH Key that exists in the cloud.
type SshKeys struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Href        string `json:"href,omitempty"`
	Id          string `json:"id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	LegacyId    int    `json:"legacy_id,omitempty"`
	ResourceUid string `json:"resource_uid,omitempty"`
}

//===== Locator

// SshKeys resource locator, exposes resource actions.
type SshKeysLocator struct {
	UrlResolver
	api *Api15
}

// SshKeys resource locator factory
func (api *Api15) SshKeysLocator(href string) *SshKeysLocator {
	return &SshKeysLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/ssh_keys
// GET /api/clouds/:cloud_id/ssh_keys
// Currently not implemented.
func (loc *SshKeysLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("SshKeys", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/ssh_keys/:id
// GET /api/clouds/:cloud_id/ssh_keys/:id
// Currently not implemented.
func (loc *SshKeysLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("SshKeys", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/******  Subnets ******/

// A Subnet is a logical grouping of network devices. An Instance can have
// many Subnets.
type Subnets struct {
	Description string       `json:"description,omitempty"`
	Href        string       `json:"href,omitempty"`
	Id          string       `json:"id,omitempty"`
	Kind        string       `json:"kind,omitempty"`
	LegacyId    int          `json:"legacy_id,omitempty"`
	Links       *SubnetLinks `json:"links,omitempty"`
	Name        string       `json:"name,omitempty"`
}

//===== Locator

// Subnets resource locator, exposes resource actions.
type SubnetsLocator struct {
	UrlResolver
	api *Api15
}

// Subnets resource locator factory
func (api *Api15) SubnetsLocator(href string) *SubnetsLocator {
	return &SubnetsLocator{UrlResolver(href), api}
}

//===== Actions

// GET /api/subnets
// GET /api/clouds/:cloud_id/subnets
// GET /api/clouds/:cloud_id/instances/:instance_id/subnets
// Currently not implemented.
func (loc *SubnetsLocator) Index(options ApiParams) error {
	var params = mergeOptionals(ApiParams{}, options)
	var uri, err = loc.Url("Subnets", "index")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

// GET /api/subnets/:id
// GET /api/clouds/:cloud_id/subnets/:id
// Currently not implemented.
func (loc *SubnetsLocator) Show(id string, options ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params = mergeOptionals(ApiParams{
		"id": id,
	}, options)
	var uri, err = loc.Url("Subnets", "show")
	if err != nil {
		return err
	}
	var _, err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return err2
	}
	return nil
}

/****** Parameter Data Types ******/

type Account struct {
	Href *string `json:"href,omitempty"`
	Id   *int    `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Cloud struct {
	CloudType   *string `json:"cloud_type,omitempty"`
	Description *string `json:"description,omitempty"`
	Href        *string `json:"href,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Datacenter struct {
	Description *string          `json:"description,omitempty"`
	Href        *string          `json:"href,omitempty"`
	Id          *string          `json:"id,omitempty"`
	Kind        *string          `json:"kind,omitempty"`
	LegacyId    *int             `json:"legacy_id,omitempty"`
	Links       *DatacenterLinks `json:"links,omitempty"`
	Name        *string          `json:"name,omitempty"`
}

type DatacenterLinks struct {
	Cloud *Cloud `json:"cloud,omitempty"`
}

type Deployment struct {
	Description  *string          `json:"description,omitempty"`
	Href         *string          `json:"href,omitempty"`
	Id           *int             `json:"id,omitempty"`
	Instances    []*Instance      `json:"instances,omitempty"`
	Kind         *string          `json:"kind,omitempty"`
	Links        *DeploymentLinks `json:"links,omitempty"`
	Locked       *bool            `json:"locked,omitempty"`
	Name         *string          `json:"name,omitempty"`
	ServerArrays []*ServerArray   `json:"server_arrays,omitempty"`
	Servers      []*Server        `json:"servers,omitempty"`
	Tags         []string         `json:"tags,omitempty"`
}

type DeploymentLinks struct {
	Account *Account `json:"account,omitempty"`
}

type Image struct {
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
	Cloud *Cloud `json:"cloud,omitempty"`
}

type Incarnator struct {
	Id    *int    `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Href  *string `json:"href,omitempty"`
	Kind  *string `json:"kind,omitempty"`
	State *string `json:"state,omitempty"`
}

type Instance struct {
	Actions            []string                 `json:"actions,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	Href               *string                  `json:"href,omitempty"`
	Id                 *string                  `json:"id,omitempty"`
	IpAddresses        []*IpAddress             `json:"ip_addresses,omitempty"`
	IsNext             *bool                    `json:"is_next,omitempty"`
	Kind               *string                  `json:"kind,omitempty"`
	LegacyId           *int                     `json:"legacy_id,omitempty"`
	Links              *InstanceLinks           `json:"links,omitempty"`
	Locked             *bool                    `json:"locked,omitempty"`
	MonitoringId       *string                  `json:"monitoring_id,omitempty"`
	MonitoringServer   *string                  `json:"monitoring_server,omitempty"`
	MonitoringToken    *string                  `json:"monitoring_token,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	Networks           []*Network               `json:"networks,omitempty"`
	OsPlatform         *string                  `json:"os_platform,omitempty"`
	PrivateDnsNames    []string                 `json:"private_dns_names,omitempty"`
	PrivateIpAddresses []string                 `json:"private_ip_addresses,omitempty"`
	PublicDnsNames     []string                 `json:"public_dns_names,omitempty"`
	PublicIpAddresses  []string                 `json:"public_ip_addresses,omitempty"`
	ResourceUid        *string                  `json:"resource_uid,omitempty"`
	SecurityGroups     *SecurityGroupCollection `json:"security_groups,omitempty"`
	ServerTemplate     *ServerTemplate          `json:"server_template,omitempty"`
	SshHost            *string                  `json:"ssh_host,omitempty"`
	State              *string                  `json:"state,omitempty"`
	Subnets            *SubnetCollection        `json:"subnets,omitempty"`
	Tags               []string                 `json:"tags,omitempty"`
	Timestamps         *BootedAtStruct          `json:"timestamps,omitempty"`
}

type InstanceLinks struct {
	Image                   *Image                   `json:"image,omitempty"`
	MultiCloudImage         *MultiCloudImage         `json:"multi_cloud_image,omitempty"`
	Subnets                 *SubnetCollection        `json:"subnets,omitempty"`
	ComputedMultiCloudImage *MultiCloudImage         `json:"computed_multi_cloud_image,omitempty"`
	Account                 *Account                 `json:"account,omitempty"`
	Incarnator              *Incarnator              `json:"incarnator,omitempty"`
	Deployment              *Deployment              `json:"deployment,omitempty"`
	SshKey                  *SshKey                  `json:"ssh_key,omitempty"`
	SecurityGroups          *SecurityGroupCollection `json:"security_groups,omitempty"`
	ComputedImage           *Image                   `json:"computed_image,omitempty"`
	Datacenter              *Datacenter              `json:"datacenter,omitempty"`
	InstanceType            *InstanceType            `json:"instance_type,omitempty"`
	Cloud                   *Cloud                   `json:"cloud,omitempty"`
}

type InstanceType struct {
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
	Cloud *Cloud `json:"cloud,omitempty"`
}

type IpAddress struct {
	Address  *string         `json:"address,omitempty"`
	Href     *string         `json:"href,omitempty"`
	Id       *string         `json:"id,omitempty"`
	Kind     *string         `json:"kind,omitempty"`
	LegacyId *int            `json:"legacy_id,omitempty"`
	Links    *IpAddressLinks `json:"links,omitempty"`
	Name     *string         `json:"name,omitempty"`
}

type IpAddressBinding struct {
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
	Cloud     *Cloud     `json:"cloud,omitempty"`
	Instance  *Instance  `json:"instance,omitempty"`
	IpAddress *IpAddress `json:"ip_address,omitempty"`
}

type IpAddressLinks struct {
	Cloud *Cloud `json:"cloud,omitempty"`
}

type MultiCloudImage struct {
	Description     *string `json:"description,omitempty"`
	Href            *string `json:"href,omitempty"`
	Id              *int    `json:"id,omitempty"`
	InheritedSource *string `json:"inherited_source,omitempty"`
	Kind            *string `json:"kind,omitempty"`
	Name            *string `json:"name,omitempty"`
	Version         *int    `json:"version,omitempty"`
}

type Network struct {
	Description *string       `json:"description,omitempty"`
	Href        *string       `json:"href,omitempty"`
	Id          *string       `json:"id,omitempty"`
	Kind        *string       `json:"kind,omitempty"`
	LegacyId    *int          `json:"legacy_id,omitempty"`
	Links       *NetworkLinks `json:"links,omitempty"`
	Name        *string       `json:"name,omitempty"`
}

type NetworkInterface struct {
	Description *string                `json:"description,omitempty"`
	Href        *string                `json:"href,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Links       *NetworkInterfaceLinks `json:"links,omitempty"`
}

type NetworkInterfaceAttachment struct {
	Href  *string                          `json:"href,omitempty"`
	Id    *string                          `json:"id,omitempty"`
	Kind  *string                          `json:"kind,omitempty"`
	Links *NetworkInterfaceAttachmentLinks `json:"links,omitempty"`
}

type NetworkInterfaceAttachmentLinks struct {
	Cloud *Cloud `json:"cloud,omitempty"`
}

type NetworkInterfaceLinks struct {
	Cloud *Cloud `json:"cloud,omitempty"`
}

type NetworkLinks struct {
	Cloud *Cloud `json:"cloud,omitempty"`
}

type SecurityGroup struct {
	Description *string             `json:"description,omitempty"`
	Href        *string             `json:"href,omitempty"`
	Id          *string             `json:"id,omitempty"`
	Kind        *string             `json:"kind,omitempty"`
	LegacyId    *int                `json:"legacy_id,omitempty"`
	Links       *SecurityGroupLinks `json:"links,omitempty"`
	Name        *string             `json:"name,omitempty"`
}

type SecurityGroupCollection struct {
	Href  *string `json:"href,omitempty"`
	Count *int    `json:"count,omitempty"`
}

type SecurityGroupLinks struct {
	Cloud *Cloud `json:"cloud,omitempty"`
}

type Server struct {
	Actions         []string     `json:"actions,omitempty"`
	CurrentInstance *Instance    `json:"current_instance,omitempty"`
	Description     *string      `json:"description,omitempty"`
	Href            *string      `json:"href,omitempty"`
	Id              *int         `json:"id,omitempty"`
	Instance        *Instance    `json:"instance,omitempty"`
	Kind            *string      `json:"kind,omitempty"`
	Links           *ServerLinks `json:"links,omitempty"`
	Name            *string      `json:"name,omitempty"`
	NextInstance    *Instance    `json:"next_instance,omitempty"`
	Tags            []string     `json:"tags,omitempty"`
}

type ServerArray struct {
	Actions         []string          `json:"actions,omitempty"`
	Description     *string           `json:"description,omitempty"`
	Href            *string           `json:"href,omitempty"`
	Id              *int              `json:"id,omitempty"`
	InstanceSummary *TotalCountStruct `json:"instance_summary,omitempty"`
	Kind            *string           `json:"kind,omitempty"`
	Links           *ServerArrayLinks `json:"links,omitempty"`
	Name            *string           `json:"name,omitempty"`
	NextInstance    *Instance         `json:"next_instance,omitempty"`
	State           *string           `json:"state,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
}

type ServerArrayLinks struct {
	NextInstance     *Instance   `json:"next_instance,omitempty"`
	CurrentInstances []*Instance `json:"current_instances,omitempty"`
	Account          *Account    `json:"account,omitempty"`
	Cloud            *Cloud      `json:"cloud,omitempty"`
}

type ServerLinks struct {
	Account         *Account  `json:"account,omitempty"`
	Cloud           *Cloud    `json:"cloud,omitempty"`
	NextInstance    *Instance `json:"next_instance,omitempty"`
	CurrentInstance *Instance `json:"current_instance,omitempty"`
}

type ServerTemplate struct {
	Description *string `json:"description,omitempty"`
	Href        *string `json:"href,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	Name        *string `json:"name,omitempty"`
	Version     *int    `json:"version,omitempty"`
}

type SshKey struct {
	Fingerprint *string `json:"fingerprint,omitempty"`
	Href        *string `json:"href,omitempty"`
	Id          *string `json:"id,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	LegacyId    *int    `json:"legacy_id,omitempty"`
	ResourceUid *string `json:"resource_uid,omitempty"`
}

type Subnet struct {
	Description *string      `json:"description,omitempty"`
	Href        *string      `json:"href,omitempty"`
	Id          *string      `json:"id,omitempty"`
	Kind        *string      `json:"kind,omitempty"`
	LegacyId    *int         `json:"legacy_id,omitempty"`
	Links       *SubnetLinks `json:"links,omitempty"`
	Name        *string      `json:"name,omitempty"`
}

type SubnetCollection struct {
	Href  *string `json:"href,omitempty"`
	Count *int    `json:"count,omitempty"`
}

type SubnetLinks struct {
	Cloud *Cloud `json:"cloud,omitempty"`
}

type TotalCountStruct struct {
	HealthyCount       *int `json:"healthy_count,omitempty"`
	UnhealthyCount     *int `json:"unhealthy_count,omitempty"`
	NotTerminatedCount *int `json:"not_terminated_count,omitempty"`
	TotalCount         *int `json:"total_count,omitempty"`
}
