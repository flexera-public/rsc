//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ praxisgen -metadata=cm16/api_docs -output=cm16 -pkg=cm16 -target=1.6 -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package cm16

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"Account": &metadata.Resource{
		Name:        "Account",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/accounts/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/accounts/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Cloud": &metadata.Resource{
		Name:        "Cloud",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Datacenter": &metadata.Resource{
		Name:        "Datacenter",
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
						Pattern:    "/api/clouds/%s/datacenters",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/datacenters`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/datacenters/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/datacenters/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/datacenters/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/datacenters/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Deployment": &metadata.Resource{
		Name:        "Deployment",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Deployment IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "link"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Deployment IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show a single Deployment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/deployments/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "link"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "link"},
					},
				},
			},
		},
	},
	"Image": &metadata.Resource{
		Name:        "Image",
		Description: `        Images define the initial Operating System and root disk contents         for new instances...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Images for the given Cloud.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/images",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/images`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter",
						Description: `              Filter images by attribute. A filter takes the form <attribute><operator><value>.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter",
						Description: `              Filter images by attribute. A filter takes the form <attribute><operator><value>.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/images/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/images/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/images/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/images/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Instance": &metadata.Resource{
		Name:        "Instance",
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
						Pattern:    "/api/clouds/%s/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter",
						Description: `              Filter instances by attribute. A filter takes the form <attribute><operator><value>.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Instance RsIds`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `The maximum number of resources to return for this index`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "tiny", "link"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter",
						Description: `              Filter instances by attribute. A filter takes the form <attribute><operator><value>.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Instance RsIds`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `The maximum number of resources to return for this index`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "tiny", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/instances/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/instances/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "tiny"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full", "tiny"},
					},
				},
			},
		},
	},
	"InstanceType": &metadata.Resource{
		Name:        "InstanceType",
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
						Pattern:    "/api/clouds/%s/instance_types",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instance_types`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/instance_types/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/instance_types/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/instance_types/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instance_types/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"IpAddress": &metadata.Resource{
		Name:        "IpAddress",
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
						Pattern:    "/api/clouds/%s/ip_addresses",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_addresses`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ip_addresses/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/ip_addresses/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_addresses/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_addresses/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"IpAddressBinding": &metadata.Resource{
		Name:        "IpAddressBinding",
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
						Pattern:    "/api/clouds/%s/ip_address_bindings",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_address_bindings`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ip_address_bindings/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/ip_address_bindings/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/ip_address_bindings/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ip_address_bindings/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"MultiCloudImage": &metadata.Resource{
		Name:        "MultiCloudImage",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/multi_cloud_images/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/multi_cloud_images/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Network": &metadata.Resource{
		Name:        "Network",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/networks/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/networks/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"NetworkInterface": &metadata.Resource{
		Name:        "NetworkInterface",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_interfaces/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/network_interfaces/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"NetworkInterfaceAttachment": &metadata.Resource{
		Name:        "NetworkInterfaceAttachment",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_interface_attachments/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/network_interface_attachments/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"SecurityGroup": &metadata.Resource{
		Name:        "SecurityGroup",
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
						Pattern:    "/api/clouds/%s/security_groups",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/security_groups`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/security_groups",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances/([^/]+)/security_groups`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/security_groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/security_groups/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/security_groups/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/security_groups/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Server": &metadata.Resource{
		Name:        "Server",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/servers/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
			},
		},
	},
	"ServerArray": &metadata.Resource{
		Name:        "ServerArray",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/server_arrays/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "full"},
					},
				},
			},
		},
	},
	"ServerTemplate": &metadata.Resource{
		Name:        "ServerTemplate",
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
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "embedded"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "embedded"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/server_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "embedded"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "embedded"},
					},
				},
			},
		},
	},
	"SshKey": &metadata.Resource{
		Name:        "SshKey",
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
						Pattern:    "/api/clouds/%s/ssh_keys",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ssh_keys`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/ssh_keys/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/ssh_keys/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/ssh_keys/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/ssh_keys/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
	"Subnet": &metadata.Resource{
		Name:        "Subnet",
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
						Pattern:    "/api/clouds/%s/subnets",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/subnets`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/instances/%s/subnets",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/instances/([^/]+)/subnets`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `      Currently not implemented.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/subnets/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/api/subnets/([^/]+)`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/%s/subnets/%s",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`/api/clouds/([^/]+)/subnets/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
				ApiParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default"},
					},
				},
			},
		},
	},
}
