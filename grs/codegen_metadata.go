//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
// $ praxisgen -metadata=grs/docs/api -output=grs -pkg=grs -target=2.0 -client=API
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package grs

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var GenMetadata = map[string]*metadata.Resource{
	"AccessRule": &metadata.Resource{
		Name: "AccessRule",
		Description: `Represents a set of access control statements that grant roles to
principals (i.e. Users or Groups) within a given scope.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all AccessRules.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/access_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/access_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of AccessRule IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "principal_href",
						Description: `List only the AccessRules associated directly with the referenced principal.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name: "scope_href",
						Description: `List only the AccessRules scoped to the reference scope.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of AccessRule IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "principal_href",
						Description: `List only the AccessRules associated directly with the referenced principal.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name: "scope_href",
						Description: `List only the AccessRules scoped to the reference scope.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single AccessRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/access_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/access_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates an AccessRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/access_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/access_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "principal_href",
						Description: `Principal (Group or User) href to which rule will apply`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "roles[]",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scope_href",
						Description: `Scope href to which the rule is scoped`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "principal_href",
						Description: `Principal (Group or User) href to which rule will apply`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "roles",
						Description: `List of roles associated with the AccessRule`,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "scope_href",
						Description: `Scope href to which the rule is scoped`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "replace",
				Description: `Replaces all AccessRules for a given scope.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/grs/access_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/access_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "principal_href",
						Description: `Replace only the AccessRules associated directly with the referenced principal.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name: "scope_href",
						Description: `Replace only the AccessRules scoped to the reference scope.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "PayloadStruct[]",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name: "principal_href",
						Description: `Replace only the AccessRules associated directly with the referenced principal.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name: "scope_href",
						Description: `Replace only the AccessRules scoped to the reference scope.
NOTE: The request params must have at least the 'scope\_href',
the 'principal\_href', or both.`,
						Type:      "string",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]*PayloadStruct",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Deletes an AccessRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/access_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/access_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Group": &metadata.Resource{
		Name: "Group",
		Description: `A Group represents a set of Users. It serves as a convenient bucket for
granting identical privileges to large number of Users and for
user-reporting purposes.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Groups.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/groups`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Group IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Group IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/groups`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description for the Group`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `URL-friendly name of the Group`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile(`(?-mix:^[\w.-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `Org under which the new Group must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description for the Group`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `URL-friendly name of the Group`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
						Regexp:      regexp.MustCompile(`(?-mix:^[\w.-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `Org under which the new Group must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description for the Group`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `New URL-friendly name for the Group. Name must be unique within an Org.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						Regexp:      regexp.MustCompile(`(?-mix:^[\w.-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `New Org under which the Group will belong`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description for the Group`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `New URL-friendly name for the Group. Name must be unique within an Org.`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
						Regexp:      regexp.MustCompile(`(?-mix:^[\w.-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `New Org under which the Group will belong`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Deletes a Group. This action deletes a Group only if the Group does not
contain any Users or ProvisioningRules. If the Group has Users or
ProvisioningRules, they must be deleted from the Group before trying to
delete the Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/groups/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"GroupUser": &metadata.Resource{
		Name:        "GroupUser",
		Description: `User sub-collection of a Group`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Users in a Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/groups/%s/users",
						Variables:  []string{"group_id"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)/users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of User IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of User IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a User in a Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/groups/%s/users/%s",
						Variables:  []string{"group_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)/users/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Adds a User to the Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/groups/%s/users",
						Variables:  []string{"group_id"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)/users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "replace",
				Description: `Replaces the Users associated with the Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/grs/groups/%s/users",
						Variables:  []string{"group_id"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)/users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string[]",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Deletes one (or all) User memberships from the Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/groups/%s/users(/%s",
						Variables:  []string{"group_id", "id)?"},
						Regexp:     regexp.MustCompile(`/grs/groups/([^/]+)/users\(/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `ID of the User to delete`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `ID of the User to delete`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"IdentityProvider": &metadata.Resource{
		Name: "IdentityProvider",
		Description: `An IdentityProvider represents a SAML identity provider (IdP)
that is linked to your organization and trusted by the RightScale
dashboard to authenticate your organization's users.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all IdentityProviders.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/identity_providers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/identity_providers`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of IdentityProvider IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of IdentityProvider IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single IdentityProvider.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/identity_providers/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},
		},
	},
	"Org": &metadata.Resource{
		Name: "Org",
		Description: `Represents an organizational unit. An Org may have a parent
Org and many child Orgs, thereby allowing a hierarchy of Orgs
to be defined.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Shows a single Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/orgs",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/orgs`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Org`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Org; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Org. Only letters, numbers, and basic
punctuation are allowed. This will be used to construct Org ID.
Name must be unique among all the children of the parent Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "parent_org",
						Description: `Parent Org under which the new Org must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Org`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Org; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Org. Only letters, numbers, and basic
punctuation are allowed. This will be used to construct Org ID.
Name must be unique among all the children of the parent Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "parent_org",
						Description: `Parent Org under which the new Org must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/orgs/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Org`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Org; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Org. Only letters, numbers, and basic
punctuation are allowed. This will be used to construct Org ID.
Name must be unique among all the children of the parent Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
					&metadata.ActionParam{
						Name: "parent_org",
						Description: `Parent Org under which the new Org must be created. If this
attribute is set to nil, the Org is set to be a top-level Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Org`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Org; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Org. Only letters, numbers, and basic
punctuation are allowed. This will be used to construct Org ID.
Name must be unique among all the children of the parent Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
					&metadata.ActionParam{
						Name: "parent_org",
						Description: `Parent Org under which the new Org must be created. If this
attribute is set to nil, the Org is set to be a top-level Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: false,
						NonBlank:  false,
					},
				},
			},

			&metadata.Action{
				Name:        "child_orgs",
				Description: `Lists all the child Orgs for an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s/child_orgs",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/child_orgs`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of child Org IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child Orgs should be traversed and their
child Orgs returned as well`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of child Org IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child Orgs should be traversed and their
child Orgs returned as well`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},
		},
	},
	"OrgGroup": &metadata.Resource{
		Name:        "OrgGroup",
		Description: `Group sub-collection of an Org.`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Groups in an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s/groups",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/groups`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Group IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child Orgs should be traversed and their
Groups returned as well`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Group IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child Orgs should be traversed and their
Groups returned as well`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},
		},
	},
	"OrgProject": &metadata.Resource{
		Name:        "OrgProject",
		Description: `Project sub-collection of an Org.`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Projects in an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s/projects",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/projects`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Project IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Indicates whether child Orgs should be traversed, and the
Projects for the child Orgs should be returned as well`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Project IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Indicates whether child Orgs should be traversed, and the
Projects for the child Orgs should be returned as well`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single Project scoped to its parent Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s/projects/%s",
						Variables:  []string{"org_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/projects/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Indicates whether child Orgs should be traversed, and the
Projects for the child Orgs should be returned as well`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Indicates whether child Orgs should be traversed, and the
Projects for the child Orgs should be returned as well`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new Project.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/orgs/%s/projects",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/projects`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of the Project to create`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `Org under which the Project must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `Name of the Project to create`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `Org under which the Project must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a Project.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/orgs/%s/projects/%s",
						Variables:  []string{"org_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/projects/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"OrgRole": &metadata.Resource{
		Name:        "OrgRole",
		Description: `Roles sub-collection of an Org.`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `No description provided for index.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s/roles",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/roles`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Role IDs`,
						Type:        "string",
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
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Role IDs`,
						Type:        "string",
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
					},
				},
			},
		},
	},
	"OrgUser": &metadata.Resource{
		Name:        "OrgUser",
		Description: `User sub-collection of an Org.`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Users in an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s/users",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of User IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child Orgs should be traversed and their
Users returned as well`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of User IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child Orgs should be traversed and their
Users returned as well`,
						Type:      "bool",
						Location:  metadata.QueryParam,
						Mandatory: false,
						NonBlank:  false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a User in an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/orgs/%s/users/%s",
						Variables:  []string{"org_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/users/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a User affiliation with the Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/orgs/%s/users",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "replace",
				Description: `Replaces the User affiliations in an Org. If this action should delete
a previous User affiliation, it will only do so if the User is not a
member of any Groups belonging to the Org. If the User is still a
member of Groups belonging to the Org, the Group membership must
be deleted before trying to delete the User affiliation in an Org via
the replace action.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/grs/orgs/%s/users",
						Variables:  []string{"org_id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string[]",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Deletes a User affiliation in an Org. This action deletes a User
affiliation only if the User is not a member of any Groups belonging
to the Org. If the User is still a member of Groups belonging to the
Org, the Group membership must be deleted before trying to delete the
User affiliation in an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/orgs/%s/users/%s",
						Variables:  []string{"org_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/orgs/([^/]+)/users/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"Project": &metadata.Resource{
		Name: "Project",
		Description: `A Project represents a container for related resources. It
belongs to an Org and can contain many cloud, design, and
other types of resources.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Projects in an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/projects",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/projects`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma-separated list of Project IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child orgs should be traversed and the
projects returned as well`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma-separated list of Project IDs`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "recurse",
						Description: `Flag indicating whether child orgs should be traversed and the
projects returned as well`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single Project.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/projects/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/projects/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "org_id",
						Description: `Org ID`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "org_id",
						Description: `Org ID`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `The view to use to render this resource`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new Project.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/projects",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/projects`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Project; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Project. Only letters, numbers and basic
punctuation are allowed. Used to construct Project ID. Must be
unique among all Projects within the Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `Org under which the Project must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Project; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Project. Only letters, numbers and basic
punctuation are allowed. Used to construct Project ID. Must be
unique among all Projects within the Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
					&metadata.ActionParam{
						Name:        "org",
						Description: `Org under which the Project must be created`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a Project.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/projects/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/projects/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Project; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Project. Only letters, numbers and basic
punctuation are allowed. Used to construct Project ID. Must be
unique among all Projects within the Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Long-form description of the Project`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "display_name",
						Description: `Human-readable name for the Project; suitable for presentation in UIs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name: "name",
						Description: `URL-friendly name for the Project. Only letters, numbers and basic
punctuation are allowed. Used to construct Project ID. Must be
unique among all Projects within the Org.`,
						Type:      "string",
						Location:  metadata.PayloadParam,
						Mandatory: true,
						NonBlank:  false,
						Regexp:    regexp.MustCompile(`(?-mix:^[\w._-]+$)`),
					},
				},
			},
		},
	},
	"ProvisioningRule": &metadata.Resource{
		Name:        "ProvisioningRule",
		Description: `A relationship table that links ProvisioningTemplates to Groups`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all ProvisioningRules.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/provisioning_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/provisioning_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of ProvisioningRule IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of ProvisioningRule IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single ProvisioningRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/provisioning_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/provisioning_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a ProvisioningRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/provisioning_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/provisioning_rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "group",
						Description: `Group to which the ProvisioningRule belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "provisioning_template",
						Description: `ProvisioningTemplate to which the ProvisioningRule belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `Regex for ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "group",
						Description: `Group to which the ProvisioningRule belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "provisioning_template",
						Description: `ProvisioningTemplate to which the ProvisioningRule belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `Regex for ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a ProvisioningRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/provisioning_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/provisioning_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description of the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `New ProvisioningRule regex`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description of the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `New ProvisioningRule regex`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Deletes a ProvisioningRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/provisioning_rules/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/provisioning_rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"ProvisioningTemplate": &metadata.Resource{
		Name: "ProvisioningTemplate",
		Description: `Represents a set of rules applied to Users or Groups for a
given IdentityProvider.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all ProvisioningTemplates for an IdentityProvider.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates",
						Variables:  []string{"identity_provider_id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of ProvisioningTemplate IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of ProvisioningTemplate IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single ProvisioningTemplate for an IdentityProvider.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s",
						Variables:  []string{"identity_provider_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a ProvisioningTemplate for an IdentityProvider.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates",
						Variables:  []string{"identity_provider_id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `User-friendly name for the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "xsl",
						Description: `XSL definition for ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `User-friendly name for the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "xsl",
						Description: `XSL definition for ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a ProvisioningTemplate for an IdentityProvider.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s",
						Variables:  []string{"identity_provider_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description for the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `New user-friendly name for the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "xsl",
						Description: `New XSL definition for ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description for the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "name",
						Description: `New user-friendly name for the ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "xsl",
						Description: `New XSL definition for ProvisioningTemplate`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Deletes a ProvisioningTemplate. This action deletes a
ProvisioningTemplate only if it does not have any ProvisioningRules.
If the ProvisioningTemplate has ProvisioningRules associated with it,
the ProvisioningRule associations must be deleted before trying to
delete the ProvisioningTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s",
						Variables:  []string{"identity_provider_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},

			&metadata.Action{
				Name:        "make_default",
				Description: `Marks a ProvisioningTemplate as the default for the IdentityProvider.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s/actions/make_default",
						Variables:  []string{"identity_provider_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)/actions/make_default`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"ProvisioningTemplateRule": &metadata.Resource{
		Name:        "ProvisioningTemplateRule",
		Description: `ProvisioningRules sub-collection of a ProvisioningTemplate.`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all ProvisioningRules for the given ProvisioningTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s/rules",
						Variables:  []string{"identity_provider_id", "provisioning_template_id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)/rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of ProvisioningRule IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of ProvisioningRule IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a ProvisioningRule for the given ProvisioningTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s/rules/%s",
						Variables:  []string{"identity_provider_id", "provisioning_template_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)/rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a ProvisioningRule for the given ProvisioningTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s/rules",
						Variables:  []string{"identity_provider_id", "provisioning_template_id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)/rules`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "group",
						Description: `Group to which the ProvisioningRule belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `Regex for ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `Description of the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "group",
						Description: `Group to which the ProvisioningRule belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `Regex for ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a ProvisioningRule for the given ProvisioningTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s/rules/%s",
						Variables:  []string{"identity_provider_id", "provisioning_template_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)/rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description for the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `New Regex for ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "description",
						Description: `New description for the ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "regex",
						Description: `New Regex for ProvisioningRule`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Deletes a ProvisioningRule for the given ProvisioningTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/identity_providers/%s/provisioning_templates/%s/rules/%s",
						Variables:  []string{"identity_provider_id", "provisioning_template_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/identity_providers/([^/]+)/provisioning_templates/([^/]+)/rules/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"User": &metadata.Resource{
		Name:        "User",
		Description: `A User represents an individual RightScale user.`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Shows a single User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/users/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/users",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/users`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "company",
						Description: `Company to which the User belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "email",
						Description: `Email of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "first_name",
						Description: `First name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "last_name",
						Description: `Last name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "password",
						Description: `Password of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "phone",
						Description: `Phone number of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "company",
						Description: `Company to which the User belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "email",
						Description: `Email of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "first_name",
						Description: `First name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "last_name",
						Description: `Last name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "password",
						Description: `Password of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "phone",
						Description: `Phone number of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PATCH",
						Pattern:    "/grs/users/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "company",
						Description: `New company name to which the User belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "email",
						Description: `New Email of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "first_name",
						Description: `New first name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "last_name",
						Description: `New last name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "phone",
						Description: `New phone number for the user`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "timezone_name",
						Description: `New timezone to which the User belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "company",
						Description: `New company name to which the User belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "email",
						Description: `New Email of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "first_name",
						Description: `New first name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "last_name",
						Description: `New last name of the User`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "phone",
						Description: `New phone number for the user`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "timezone_name",
						Description: `New timezone to which the User belongs`,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"UserGroup": &metadata.Resource{
		Name:        "UserGroup",
		Description: `Group sub-collection of a user`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Groups of which the User is a member.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/users/%s/groups",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/groups`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Group IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Group IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a Group of which the User is a member.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/users/%s/groups/%s",
						Variables:  []string{"user_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/groups/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a Group membership for the User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/users/%s/groups",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/groups`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "replace",
				Description: `Replaces the Group membership(s) for the User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/grs/users/%s/groups",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/groups`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "org_href",
						Description: `Replace only the Groups that belongs to a given Org`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "saml_only",
						Description: `Replace only SAML-based Groups, if set to true`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string[]",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "org_href",
						Description: `Replace only the Groups that belongs to a given Org`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "saml_only",
						Description: `Replace only SAML-based Groups, if set to true`,
						Type:        "bool",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name:        "delete",
				Description: `Deletes one (or all) Group membership(s) for the User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/users/%s/groups(/%s",
						Variables:  []string{"user_id", "id)?"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/groups\(/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `ID of the Group to delete`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "id",
						Description: `ID of the Group to delete`,
						Type:        "int",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
				},
			},
		},
	},
	"UserIdentity": &metadata.Resource{
		Name: "UserIdentity",
		Description: `Represents a User's principal identity in the context of a
specific identity provider. Known to our platform by a SAML
NameID (also known as SAML subject), which is an opaque
string, often an email address, but sometimes an LDAP DN,
ActiveDirectory username, random GUID, or other identifier.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all UserIdentities.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/user_identities",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`/grs/user_identities`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of UserIdentity IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of UserIdentity IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single UserIdentity.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/user_identities/%s",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`/grs/user_identities/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},
		},
	},
	"UserOrg": &metadata.Resource{
		Name:        "UserOrg",
		Description: `Org sub-collection of a user.`,
		Identifier:  "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Orgs with which the User is affiliated.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/users/%s/orgs",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/orgs`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Org IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of Org IDs`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows an Org with which the User is affiliated.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/users/%s/orgs/%s",
						Variables:  []string{"user_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/orgs/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "extended", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates an affiliation for a User to an Org.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "POST",
						Pattern:    "/grs/users/%s/orgs",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/orgs`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "replace",
				Description: `Replaces all Org affiliations for a User. If this action should delete
a previous Org affiliation, it will only do so if the User is not a
member of any Groups belonging to the Org. If the User is still a
member of Groups belonging to the to the Org, the Group membership must
be deleted before trying to delete the Org affiliation for a User via
the replace action.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/grs/users/%s/orgs",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/orgs`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "string[]",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},

			&metadata.Action{
				Name: "delete",
				Description: `Deletes an Org affiliation for a User. This action deletes an Org
affiliation only if the User is not a member of any Groups belonging
to the Org. If the User is still a member of Groups belonging to the
Org, the Group membership must be deleted before trying to delete the
Org affiliation for a User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "DELETE",
						Pattern:    "/grs/users/%s/orgs/%s",
						Variables:  []string{"user_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/orgs/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{},
				APIParams:    []*metadata.ActionParam{},
			},
		},
	},
	"UserUserIdentity": &metadata.Resource{
		Name: "UserUserIdentity",
		Description: `UserIdentity sub-collection of a User, which defines relationships between
the User and its IdentityProviders.`,
		Identifier: "",
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all UserIdentities for the given User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/users/%s/identities",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/identities`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of User IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "filter",
						Description: `Filter by attribute`,
						Type:        "string",
						Location:    metadata.QueryParam,
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "ids",
						Description: `Comma separated list of User IDs`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a UserIdentity for the given User.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "GET",
						Pattern:    "/grs/users/%s/identities/%s",
						Variables:  []string{"user_id", "id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/identities/([^/]+)`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "fields",
						Description: `Selection of fields to render.`,
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
						ValidValues: []string{"default", "link"},
					},
				},
			},

			&metadata.Action{
				Name: "replace",
				Description: `Replaces the UserIdentities for the given User.
Currently, only one identity per user is allowed, and the
replace action will fail if more than one identity is provided
at a time. This restriction could be lifted in the future.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HTTPMethod: "PUT",
						Pattern:    "/grs/users/%s/identities",
						Variables:  []string{"user_id"},
						Regexp:     regexp.MustCompile(`/grs/users/([^/]+)/identities`),
					},
				},
				CommandFlags: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]string",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				Payload: "PayloadStruct[]",
				APIParams: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "payload",
						Description: ``,
						Type:        "[]*PayloadStruct",
						Location:    metadata.PayloadParam,
						Mandatory:   true,
						NonBlank:    false,
					},
				},
			},
		},
	},
}
