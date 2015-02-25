//************************************************************************//
//                     rsc - RightScale API command line tool
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
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var api_metadata = map[string]*metadata.Resource{
	"Accounts": &metadata.Resource{
		Name:        "Accounts",
		Description: `        Resources in RightScale generally belong to accounts`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/accounts`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/accounts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/accounts/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "int",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Clouds": &metadata.Resource{
		Name:        "Clouds",
		Description: `        Clouds provide remote resources for things like storage and compute.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/clouds`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "int",
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Datacenters": &metadata.Resource{
		Name:        "Datacenters",
		Description: `        Datacenters are cloud resources that give you the ability to place         resources in isolated locations...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/datacenters",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/datacenters`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/datacenters",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/datacenters`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/datacenters/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/datacenters/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/datacenters/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/datacenters/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Deployments": &metadata.Resource{
		Name:        "Deployments",
		Description: `        Deployments provide a way to group resources that logically belong         together...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `          List all Deployments in an Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/deployments`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[ids]",
						Description: `Comma separated list of Deployment IDs`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"ids, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show a single Deployment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/deployments/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `ID of the Deployment to show`,
						Type:        "int",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Images": &metadata.Resource{
		Name:        "Images",
		Description: `        Images define the initial Operating System and root disk contents         for new instances...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Images for the given Cloud.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/images",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/images`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[filter]",
						Description: `              Filter images by attribute. A filter takes the form <attribute><operator><value>.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/images/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/images/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/images/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/images/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"InstanceTypes": &metadata.Resource{
		Name:        "InstanceTypes",
		Description: `        An InstanceType represents a basic hardware configuration for an         Instance...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/instance_types",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/instance_types`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instance_types",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instance_types`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/instance_types/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/instance_types/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instance_types/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instance_types/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Instances": &metadata.Resource{
		Name:        "Instances",
		Description: `        Instances represent an entity that is runnable in the cloud.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `          List all Instances in an account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/instances",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/instances`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[filter]",
						Description: `              Filter instances by attribute. A filter takes the form <attribute><operator><value>.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[limit]",
						Description: `The maximum number of resources to return for this index`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `Filter by the given cloud ID`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[ids]",
						Description: `Comma separated list of Instance RsIds`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"filter, limit, ids, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/instances/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/instances/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"IpAddressBindings": &metadata.Resource{
		Name:        "IpAddressBindings",
		Description: `        An IpAddressBinding represents an abstraction for binding an IpAddress         to an instance...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ip_address_bindings",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/ip_address_bindings`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_address_bindings",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_address_bindings`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ip_address_bindings/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/ip_address_bindings/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_address_bindings/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_address_bindings/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"IpAddresses": &metadata.Resource{
		Name:        "IpAddresses",
		Description: `        An IpAddress provides an abstraction for IPv4 addresses bindable to         Instance resources running in a Cloud...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ip_addresses",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/ip_addresses`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_addresses`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ip_addresses/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/ip_addresses/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_addresses/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"MultiCloudImages": &metadata.Resource{
		Name:        "MultiCloudImages",
		Description: `        A MultiCloudImage is a RightScale component that functions as a pointer         to machine images in specific clouds (e...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/multi_cloud_images`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/multi_cloud_images/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/multi_cloud_images/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "int",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"NetworkInterfaceAttachments": &metadata.Resource{
		Name:        "NetworkInterfaceAttachments",
		Description: `        NetworkInterfaceAttachments represent an attachment between a         NetworkInterface and another resource...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_interface_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/network_interface_attachments`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_interface_attachments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/network_interface_attachments/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"NetworkInterfaces": &metadata.Resource{
		Name:        "NetworkInterfaces",
		Description: `        Just like their physical counterparts, NetworkInterfaces join other         resources to a network...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_interfaces",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/network_interfaces`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_interfaces/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/network_interfaces/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Networks": &metadata.Resource{
		Name:        "Networks",
		Description: `        A Network is a logical grouping of network devices.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/networks",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/networks`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/networks/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/networks/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"SecurityGroups": &metadata.Resource{
		Name:        "SecurityGroups",
		Description: `        Security Groups represent network security profiles that contain lists         of firewall rules for different ports and source IP addresses, as well         as trust relationships between security groups...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/security_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/security_groups`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/security_groups",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/security_groups`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/security_groups",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances/([^/]+)/security_groups`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[instance_id]",
						Description: `The Instance with which to scope the index`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/security_groups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/security_groups/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/security_groups/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"ServerArrays": &metadata.Resource{
		Name:        "ServerArrays",
		Description: `        A server array represents a logical group of instances and allows to         resize(grow/shrink) that group based on certain elasticity parameters...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/server_arrays`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/server_arrays/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "int",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"ServerTemplates": &metadata.Resource{
		Name:        "ServerTemplates",
		Description: `        ServerTemplates allow you to pre-configure servers by starting from a         base image and adding scripts that run during the boot, operational,         and shutdown phases...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/server_templates`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/server_templates/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "int",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Servers": &metadata.Resource{
		Name:        "Servers",
		Description: `        Servers represent the notion of a server/machine from RightScale's         perspective...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/servers`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/servers/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "int",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"SshKeys": &metadata.Resource{
		Name:        "SshKeys",
		Description: `        Ssh Keys represent a created SSH Key that exists in the cloud.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ssh_keys",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/ssh_keys`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ssh_keys",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ssh_keys`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ssh_keys/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/ssh_keys/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ssh_keys/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ssh_keys/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Subnets": &metadata.Resource{
		Name:        "Subnets",
		Description: `        A Subnet is a logical grouping of network devices`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/subnets",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/api/subnets`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/subnets",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/subnets`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/subnets",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances/([^/]+)/subnets`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[instance_id]",
						Description: `The Instance with which to scope the index`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/subnets/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/subnets/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/subnets/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/subnets/([^/]+)`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "params[id]",
						Description: `The identifier of the resource`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[cloud_id]",
						Description: `The identifier of Cloud this resource resides in`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "params[view]",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
}
