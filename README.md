# rsc - An experimental generic RightScale API client

rsc provides both a go package and a command line tool for interacting with various RightScale APIs. 
The RightScale API 1.5, RightScale API 1.6 and Self-Service APIs are supported at this time.

- Master: [![Build Status](https://travis-ci.org/rightscale/rsc.svg?branch=master)](https://travis-ci.org/rightscale/rsc)
  ![Code Coverage](https://s3.amazonaws.com/rs-code-coverage/rsc/cc_badge_master.svg)

rsc can be used in one of three ways:

* As a command line tool, the section below describes that usage.
* As a way to make specific API requests programmatically. In this use case the request parameters
  and response use types that are specific to the request.
* As a way to build higher level generic API clients. The `Do` methods exposed for each API make
  it possible to write generic code that can make any API request given a resource and action name.

*NOTE*: rsc is work in progress. The test coverage is currently very limited, use at your
own risk. Please use github issues to report problems.

## API References

* [RightScale API 1.5](http://reference.rightscale.com/api1.5/index.html)
* [RightScale API 1.6](http://dev-api-docs.rightscale.com/#/1.6/controller/V1_6::Accounts)
* [Self-Service Designer](https://s3.amazonaws.com/rs_api_docs/selfservice/designer/index.html#/1.0/controller/V1::Controller::Schedule)
* [Self-Service Catalog](https://s3.amazonaws.com/rs_api_docs/selfservice/catalog/index.html#/1.0/controller/V1::Controller::AccountPreference)
* [Self-Service Manager](https://s3.amazonaws.com/rs_api_docs/selfservice/manager/index.html#/1.0/controller/V1::Controller::ScheduledAction)

## Command Line Tool

The command line tool uses subcommands to interact with each API. Use `rsc api15` to send requests
to the RightScale API 1.5, `rsc api16` to send requests to the RightScale API 1.6 and `rsc ss` to
send requests to the Self-Service APIs. `api15` is currently the default client so `rsc` is
equivalent to `rsc api15`.

The general shape of a command line is:

```
$ rsc [GLOBAL] [API] ACTION HREF [PARAM=VALUE]
```
where `GLOBAL` is an optional list of global flags, `API` is `api15`, `api16` or `ss`, `ACTION` is
the API action to perform (i.e `index`, `show`, `update`, etc.), `HREF` is the resource or resource
collection href (i.e. `/api/servers`, `/api/servers/1` etc.) and `PARAM` and `VALUE` are the names
and values of the action parameters (e.g. `view=extended`).

### Global Flags

The list of global flags is:
```
  --help            Show help.
  --version         Show application version.
  -c, --config="/Users/raphael/.rsc"
                    path to rsc config file
  -a, --account=ACCOUNT
                    RightScale account ID
  -h, --host=HOST   RightScale login endpoint (e.g. 'us-3.rightscale.com')
  --email=EMAIL     Login email, use --email and --password or use --key or --rl10
  --pwd=PWD         Login password, use --email and --password or use --key or --rl10
  -k, --key=KEY     OAuth access token or API key, use --email and --password or use --key or --rl10
  --rl10            Proxy requests through RL 10 agent, use --email and --password or use --key or --rl10
  --hrefs           List all known href patterns for selected API or resource
  --x1=X1           Extract single value using JSON:select
  --xm=XM           Extract zero, one or more values using JSON:select and return space separated list
  --xj=XJ           Extract zero, one or more values using JSON:select and return JSON
  --xh=XH           Extract header with given name
  -n, --noRedirect  Do not follow redirect responses
  --fetch           Fetch resource with href present in 'Location' header
  --dump            Dump HTTP request and response for debugging
  --pp              Pretty print response body
```
A few key points:

`rsc` can read the host, account and authentication information from a config file instead of the
command line. See [Setup and Config](#config) below.

Authentication can be done in one of three ways: 
* Using a RightScale user email and password (`--email` and `--password`)
* Using a OAuth token retrieved from the _API Credentials_ entry of the _Settings_ menu (`--key`)
* Using an API instance token to make API calls from a RightScale instance (`--key`)
* Using the RightLink 10 proxy to make API calls from a RightScale instance running RightLink 10 (`--rl10`)

The `--x1`, `--xm` and `--xj` flags make it possible to extract values from the response using a
JSON select expression (see [http://jsonselect.org/](http://jsonselect.org/)). For example:
```
$ rsc --xm .name index clouds
```
extracts the names of each cloud from the response and prints the result as a space separated list.

### Actions and Parameters

The names of the action and its parameters follow the API documentation. Parameter names use form
encoding to represent nested data structures. For example with API 1.5 the command line to create
a volume is:
```
$ rsc --pp --fetch create clouds/1/volumes \
volume[name]="My New Volume" \
volume[size]="10" \
volume[datacenter_href]="/api/clouds/1/datacenters/5K443K2CF8NS6" \
volume[volume_type_href]="/api/clouds/1/volume_types/BDVEN383N1EN2"
```
The `/api/` prefix for API 1.5 and API 1.6 hrefs is optional so the following lists all clouds:
```
$ rsc index clouds
```
### <a name="config"></a>Setup and Config
rsc has a top level `setup` command which creates a rsc config file. The config file contains the
RightScale account ID, API host, user email and (encrypted) password so that these flags don't have
to be provided each time the tool is invoked.

By default the config file is created in `$HOME/.rsc`, the location can be overridden using the
`--config` global flag. Multiple configs may be created to allow for different environments or
users. The config file itself is a simple JSON file that can be edited manually (apart from the
password value that needs to be encrypted by rsc).

### Built-in Help
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
usage: rsc [<flags>] index index /api/clouds [<cloud index params>]

Cloud index params:
filter[]=[]string
    <optional, [cloud_type|description|name]>

view=string
    <optional, [default|extended]>
```
Another useful flag is `--hrefs` which lists the valid Href pattern for the given resource or all
resources if no resource is specified. For example:
```
rsc --hrefs index clouds
Method Href Pattern    Resource.Action
GET    /api/clouds     Cloud.index
GET    /api/clouds/:id Cloud.show
```

## Go Package

Each API is encapsulated in a different package: package `rsapi15` for API 1.5, package `rsapi16`
for API 1.6 and package `ss` for Self-service.

### Using the Discrete Methods

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

### Using the Generic Methods

Each package also contains a generic `Do` method which sends requests for a given resource and
action name together with generic parameters:
```go
func (a *Api) Do(resource, action, href string, params rsapi.ApiParams) (*http.Response, error)
```
The `ss` package `Do` method also accepts the name of the SS service to send requests to (i.e.
`designer`, `catalog` or `manager`).

Package `rsapi` which contains common code for all packages defines the `LoadResponse` function
that can load the response body:
```go
func (a *Api) LoadResponse(resp *http.Response) (interface{}, error)
```
## Building: API Metadata and Code Generation

Part of the `rsc` source code (the vast majority in terms of lines of code) is automatically 
generated from API metadata. There are currently two code generators: `api15gen` consumes the
RightScale API 1.5 metadata hosted [here](http://reference.rightscale.com/api1.5/api_data.json) and
`praxisgen` consumes the metadata for any praxis application (for example for API 1.6. see
[https://s3.amazonaws.com/rs_api_docs/selfservice/manager/docs/1.0/resources/V1::Controller::Execution.json](https://s3.amazonaws.com/rs_api_docs/selfservice/manager/docs/1.0/resources/V1::Controller::Execution.json).

The code generator tools source code lives under the `gen` directory. Once the tools are compiled
and installed they can be invoked using `go generate`, see [http://blog.golang.org/generate](http://blog.golang.org/generate)
for information on how `go generate` works. The `go generate` comments live in the top level file
`generate.go`.

The Makefile included in the source code contains a `rsc` target that can be invoked to generate
the code and build the command line tool.

## Directory Structure

There is one package / sub-directory per API wrapped by rsc, currently these are the `rsapi15`,
`rsapi16` and `ss` directories. These directories contain the code generated by the generator from
the JSON: the `codegen_client.go` and `codegen_metadata.go` files.
It also contains code that expose the API sub-commands and generic `Do` client method.

The `ss` package consists of 3 sub-packages, one per self-service API: The `ssd` sub-package
contains the code for making requests to the designer API, the `ssc` to the catalog API and finally
`ssm` to the manager. These are sub-packages instead of top-level packages so that they can be
wrapped and exposed through a single entry point (both for the command line and the generic client
method).

The `rsapi` directory contains code that is shared by all API packages. This code includes common
command line parsing algorithms as well as low level HTTP request handling. It also includes
authenticators which sign API requests by adding the required auth headers (cookie in the case of
email/password authentication, OAuth header in the case of OAuth and custom header in the case of
RightLink 10).

The `metadata` directory defines the data types generated by `api15gen` and `praxisgen`. These
include data types that describe resources, actions and action parameters. These data structures
are used by the generated code and by the command line parsing code to provide contextual help
for example.

Finally, the `cmd` package contains common code used by `rsc` and each API command line processing
code to access command line flags.

### Contributing (internal)

To add support for a new API:
  1. Create a subdirectory whose name matches the name of the go package (`rsapi15`, `ss`, etc.).
  2. Put the JSON metadata describing the API in that directory.
  3. Add a `go generate` directive that invokes the generator against the JSON to `generate.go`.
  4. Generate the code and add the corresponding command line parsing and `Do` methods (see
     `commands.go`, `http.go` and `rsapi16.go` in the rsapi16 directory for an example).
  5. Add the corresponding sub-command to rsc (see the top level `command_line.go` file).

The JSON metadata for praxis apps is the JSON generated by the `rake praxis:api_docs` command.
Updating a client to the latest version of an API thus consists of updating the corresponding JSON
and rebuilding the client (the `go generate` directives will take care of updating the generated
code).
