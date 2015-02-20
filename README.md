rsc - A generic RightScale API client
==========================================
rsc provides both a go package and command line tool for interacting with various RightScale APIs. 
Only API 1.5 is supported at this time.

Command Line Tool
-----------------
The command line tool uses subcommands to interact with each API, for example `rsc api15`
for sending requests to API 1.5 etc. `api15` is currently the default client (so `rsc` is
equivalent to `rsc api15`)

The general shape of a command line is:

```bash
rsc [GLOBAL] ACTION HREF [PARAM=VALUE]
```
where GLOBAL is an optional list of global flags, ACTION is the API action to perform (i.e "index",
"show", "update", etc.), HREF is the resource or resource collection href (i.e. "/api/servers",
"/api/servers/1" etc.) and PARAM and VALUE are the names and values of the action parameters.

The list of available global flags is:
```bash
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
```bash
rsc create clouds/1/volumes \
volume[name]="My New Volume" \
volume[size]=10 \
volume[datacenter_href]="/api/clouds/1/datacenters/5K443K2CF8NS6" \
volume[volume_type_href]="/api/clouds/1/volume_types/BDVEN383N1EN2"
```
The `/api/` prefix for hrefs is optional so the following lists all clouds:
```bash
rsc index clouds
```

Setup and Config
----------------
rsc has a top level `setup` command which can be used to create a rsc config file. The config file
contains the RightScale account ID, API host and (encrypted) API key or token so that these flags 
don't have to be provided each time the tool is invoked.

By default the config file is created in $HOME/.rsc, the location can be overridden using the
`--config` global flag. Multiple configs may be created to allow for different environments or
users. The config file itself is a simple JSON file that can be edited manually (apart from the
token value that needs to be encrypted by rsc).

Built-in Help
-------------
The "--help" flag is available on all commands. It displays contextual help, for example:
```bash
./rsc index --help
usage: rsc [<flags>] index [<flags>] <href>

Lists all resources of given type in account.

Flags:
  -P, --params=QUERY=VALUE
    Action parameters in the form QUERY=VALUE, e.g. '-P server[name]=server42'

Args:
  <href>  API Resource or resource collection href on which to act, e.g. '/api/servers'
```
Or:
```bash
/rsc api15 index clouds --help
usage: rsc [<flags>] api15 index [<Cloud.index flags>] /api/clouds
<Cloud.index flags>:
--filter
    <[]string, optional>
--view
    <string, optional>
```

Go Package
----------
Each API is encapsulated in a different package: package 'rsapi15' for API 1.5, package 'rsapi16'
for API 1.6, etc.

The packages contain "resource locators", one or two per resource exposed by the underlying API.
For each resource there may be a collection locator - e.g. `CloudsLocator` - if the API exposes
collection level actions for this resource type, and a resource locator - e.g. `CloudLocator` - if
the API exposes resource level actions for this resource type.

Locators can be retrieved using factory methods exposed by the Api15 object. The factory methods
accept the collection or resource href and return the corresponding locator. For example the
following creates a cloud locator:
```go
var cloudLocator = api.CloudLocator("/api/clouds/1")
```
Locators expose one method for each action supported by the underlying collection or resource. For
example the clouds collection locator `CloudsLocator` exposes an `Index` method while the
`CloudLocator` resource locator exposes a `Show` method. These methods may return resources which
are structs that expose the underlying resource attributes. For example the `CloudsLocator` `Index`
method returns an array of `Cloud` resource. A cloud resource is defined as:
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
and the `index` method is defined as:
```go
// GET /api/clouds
// Lists the clouds available to this account.
// -- Optional parameters:
// 	filter
// 	view
func (loc *CloudsLocator) Index(options ApiParams) ([]*Cloud, error)
```
The following code would invoke the `Index` method:
```go
var clouds = api.CloudsLocator("/api/clouds").Index(nil)
```
`Create` actions all return a locator so that fetching the corresponding resource can be done in
one expression:
```go
var params rsapi15.VolumeParam{} // Code that initializes parameters omitted for brevity
var volumesLocator = api.VolumesLocator("/api/clouds/1/volumes")
var volume = volumesLocator.Create(&params).Show(ApiParams{})
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
  1. Create a subdirectory whose name matches the name of the go package ('api15', 'ss', etc.).
  2. Put the JSON metadata describing the API in that directory.
  3. Add a `go generate` directive that invokes the generator against the JSON to one of the package
     source file.

The JSON metadata for praxis apps is the JSON generated by the `rake praxis:api_docs` command, the
JSON for API 1.5 is also automatically generated using `rake api:doc:generate`. Updating a client
to the latest version of an API thus consists of updating the corresponding JSON and rebuilding the
client (the `go generate` directives will take care of updating the generated code).
