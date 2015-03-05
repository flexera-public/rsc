rsc - A generic RightScale API client
==========================================
rsc provides both a go package and command line tool for interacting with various RightScale APIs. 
Both API 1.5 and API 1.6 are supported at this time.

- Master: [![Build Status](https://travis-ci.org/rightscale/rsc.svg?branch=master)](https://travis-ci.org/rightscale/rsc)
  ![Code Coverage](https://s3.amazonaws.com/rs-code-coverage/rsc/cc_badge_master.svg)

Command Line Tool
-----------------
The command line tool uses subcommands to interact with each API, for example `rsc api15`
for sending requests to API 1.5 etc. `api15` is currently the default client (so `rsc` is
equivalent to `rsc api15`).

The general shape of a command line is:

```
$ rsc [GLOBAL] [API] ACTION HREF [PARAM=VALUE]
```
where `GLOBAL` is an optional list of global flags, `API` is `api15` or `api16`, `ACTION` is the 
API action to perform (i.e `index`, `show`, `update`, etc.), `HREF` is the resource or resource
collection href (i.e. `/api/servers`, `/api/servers/1` etc.) and `PARAM` and `VALUE` are the names
and values of the action parameters.

The list of available global flags is:
```
  --help            Show help.
  --version         Show application version.
  -c, --config="/Users/raphael/.rsc"
                    path to rsc config file
  -a, --account=ACCOUNT
                    RightScale account ID
  -h, --host=HOST   RightScale API host (e.g. 'us-3.rightscale.com')
  -k, --key=KEY     OAuth access token or API key
  --rl10            Proxy requests through RightLink 10 (exclusive with '--key')
  --x1=X1           Extract single value using given JSON:select expression
  --xm=XM           Extract zero, one or multiple values using given JSON:select expression and return space separated list (useful for bash scripts)
  --xj=XJ           Extract zero, one or multiple values using given JSON:select expression and return JSON
  --xh=XH           Extract header with given name
  -n, --noRedirect  Do not follow redirect responses
  --fetch           Fetch resource with href present in 'Location' header
  --dump            Dump HTTP request and response for debugging
  --pp              Pretty print response body
```

The names of the action and its parameters follow exactly the API documentation. For example, with
API 1.5 the command line to create a volume would look like:
```
$ rsc create clouds/1/volumes \
volume[name]="My New Volume" \
volume[size]="10" \
volume[datacenter_href]="/api/clouds/1/datacenters/5K443K2CF8NS6" \
volume[volume_type_href]="/api/clouds/1/volume_types/BDVEN383N1EN2"
```
The `/api/` prefix for hrefs is optional so the following lists all clouds:
```
$ rsc index clouds
```

Setup and Config
----------------
rsc has a top level `setup` command which can be used to create a rsc config file. The config file
contains the RightScale account ID, API host and (encrypted) API key or token so that these flags 
don't have to be provided each time the tool is invoked.

By default the config file is created in `$HOME/.rsc`, the location can be overridden using the
`--config` global flag. Multiple configs may be created to allow for different environments or
users. The config file itself is a simple JSON file that can be edited manually (apart from the
token value that needs to be encrypted by rsc).

Built-in Help
-------------
The `--help` flag is available on all commands. It displays contextual help, for example:
```
$ rsc index --help
usage: rsc [<flags>] index <href> [<params>]

Lists all resources of given type in account.

Args:
  <href>      API Resource or resource collection href on which to act, e.g. '/api/servers'
  [<params>]  Action parameters in the form QUERY=VALUE, e.g. 'server[name]=server42'
```
Or:
```
$ rsc api15 index clouds --help
usage: rsc [<flags>] api15 index /api/clouds [<cloud index params>]

Cloud index params:
filter[]=[]string
    <optional>

view=string
    <optional, [default|extended]>
```

Go Package
----------
Each API is encapsulated in a different package: package `rsapi15` for API 1.5, package `rsapi16`
for API 1.6, etc.

The packages contain "resource locators", one per resource exposed by the underlying API.

Locators are instantiated using factory methods exposed by the Api object. The factory methods
accept the collection or resource href and return the corresponding locator. For example the
following creates a cloud locator:
```go
var cloudLocator = api.CloudLocator("/api/clouds/1")
```
Locators expose one method for each action supported by the underlying collection or resource. For
example the clouds collection locator `CloudLocator` exposes an `Index()` and a `Show()` method.
Locator methods may return resources which are structs that expose the underlying resource 
attributes. For example the `CloudLocator` `Index()` method returns an array of `Cloud` resource.
A cloud resource is defined as:
```go
// Represents a Cloud (within the context of the account in the session).
type Cloud struct {
	Capabilities []string            `json:"capabilities,omitempty"`
	CloudType    string              `json:"cloud_type,omitempty"`
	Description  string              `json:"description,omitempty"`
	DisplayName  string              `json:"display_name,omitempty"`
	Links        []map[string]string `json:"links,omitempty"`
	Name         string              `json:"name,omitempty"`
}
```
and the `Index()` method is defined as:
```go
// GET /api/clouds
// Lists the clouds available to this account.
// -- Optional parameters:
// filter
// view
func (loc *CloudLocator) Index(options rsapi.ApiParams) ([]*Cloud, error)
```
The following code would invoke the `Index()` method:
```go
var clouds, err = api.CloudLocator("/api/clouds").Index(rsapi.ApiParams{})
```
`Create` actions all return a locator so that fetching the corresponding resource is convenient:
```go
var volumesLocator = api.VolumeLocator("/api/clouds/1/volumes")
var params rsapi15.VolumeParam{} // Code that sets parameters omitted for brevity
if loc, err := volumeLocator.Create(&params); err == nil {
	volume, err := loc.Show(rsapi.ApiParams{})
	// ... check error, use volume etc.
}
```

Adding support for new apis / Updating existing clients
-------------------------------------------------------
Most of the code of rsc is automatically generated from the JSON descriptions of the APIs.
There are currently two generators: the API 1.5 generator and the praxis generator. The 'generators'
directory contains the source for these generators.

The JSON metadata for APIs wrapped by rsc lives in a sub-directory, one per API. This directory
also contains the code generated by the generator from the JSON. It defines the corresponding go
package content as well as the command line tool subcommands.

Adding support for a new API consists of:
  1. Create a subdirectory whose name matches the name of the go package (`rsapi15`, `ss`, etc.).
  2. Put the JSON metadata describing the API in that directory.
  3. Add a `go generate` directive that invokes the generator against the JSON to one of the package
     source file.

The JSON metadata for praxis apps is the JSON generated by the `rake praxis:api_docs` command, the
JSON for API 1.5 is also automatically generated using `rake api:doc:generate`. Updating a client
to the latest version of an API thus consists of updating the corresponding JSON and rebuilding the
client (the `go generate` directives will take care of updating the generated code).
