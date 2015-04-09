# rsc - A generic RightScale API client [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/rightscale/rsc/blob/master/LICENSE)

`rsc` provides both a command line tool and a go package for interacting with the RightScale APIs.
The currently supported APIs are the RightScale Cloud Management API 1.5 and 1.6 APIs as well as
the RightScale Self-Service 1.0 APIs (latest version for this product).

- Master: [![Build Status](https://travis-ci.org/rightscale/rsc.svg?branch=master)](https://travis-ci.org/rightscale/rsc)
[![Coverage](https://s3.amazonaws.com/rs-code-coverage/rsc/cc_badge_master.svg)](https://gocover.io/github.com/rightscale/rsc)
- 1.0.rc0: [![Build Status](https://travis-ci.org/rightscale/rsc.svg?branch=1.0.rc0)](https://travis-ci.org/rightscale/rsc)
[![Coverage](https://s3.amazonaws.com/rs-code-coverage/rsc/cc_badge_1.0.rc0.svg)](https://gocover.io/github.com/rightscale/rsc)

`rsc` can be used in one of two ways:

* As a command line tool, see [below](#usage) for usage.
* As a way to make API requests to RightScale programmatically from go code.

Please use [github issues](https://github.com/rightscale/rsc/issues) to report problems.

Jump to:
* [Command line usage](#usage)
* [Go package usage](#go)
* [Development & Contributing](#contributing)
* [Command Line Help and Cookbook](COOKBOOK.md)

## <a name="reference"></a>API References

* [RightScale CM API 1.5](http://reference.rightscale.com/api1.5/index.html)
* [RightScale CM API 1.6](http://dev-api-docs.rightscale.com/#/1.6/controller/V1_6::Accounts)
* [RightScale SS 1.0 Designer](https://s3.amazonaws.com/rs_api_docs/selfservice/designer/index.html#/1.0/controller/V1::Controller::Schedule)
* [RightScale SS 1.0 Catalog](https://s3.amazonaws.com/rs_api_docs/selfservice/catalog/index.html#/1.0/controller/V1::Controller::AccountPreference)
* [RightScale SS 1.0 Manager](https://s3.amazonaws.com/rs_api_docs/selfservice/manager/index.html#/1.0/controller/V1::Controller::ScheduledAction)

## <a name="usage"></a>Command Line Tool

The command line tool uses subcommands to interact with each API. Use `rsc cm15` to send requests
to the RightScale Cloud Management API 1.5, `rsc cm16` to send requests to the RightScale Cloud
Management API 1.6 and `rsc ss` to send requests to the RightScale Self-Service API 1.0.

### Download & Install

Download a statically linked binary and run it as follows:
```
$ curl https://rightscale-binaries.s3.amazonaws.com/rsbin/rsc/1.0.rc0/rsc-linux-amd64.tgz |\
  tar -zxf - -O ./rsc/rsc > rsc
$ chmod +x ./rsc
$ ./rsc --version
rsc master - 2015-03-07 10:50:29 - 5c43698e47d615c0e9ecac8757430a8ad4c34c75
```

- MacOS: `https://rightscale-binaries.s3.amazonaws.com/rsbin/rsc/1.0.rc0/rsc-darwin-amd64.tgz`
- Windows: `https://rightscale-binaries.s3.amazonaws.com/rsbin/rsc/1.0.rc0/rsc-windows-amd64.zip`
- ODroid/RasPi/armhf: `https://rightscale-binaries.s3.amazonaws.com/rsbin/rsc/1.0.rc0/rsc-linux-arm.tgz`
See further down in the README for building from source.

### Command line

The general shape of a command line is:

```
$ rsc [GLOBAL] [API] ACTION HREF [PARAM=VALUE]
```
where:
- `GLOBAL` is an optional list of global flags,
- `API` is `cm15`, `cm16` or `ss`,
- `ACTION` is the API action to perform (i.e `index`, `show`, `update`, etc.),
- `HREF` is the resource or resource collection href (i.e. `/api/servers`, `/api/servers/1` etc.), and
- `PARAM` and `VALUE` are the names and values of the action parameters (e.g. `view=extended`).

For example:
```
$ rsc --rl10 cm15 show /api/servers 'filter[]=name==LB'
```

The sections below cover each option in order.

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
  --x1=X1           Extract single value using JSON:select
  --xm=XM           Extract zero, one or more values using JSON:select and return space separated list
  --xj=XJ           Extract zero, one or more values using JSON:select and return JSON
  --xh=XH           Extract header with given name
  -n, --noRedirect  Do not follow redirect responses
  --fetch           Fetch resource with href present in 'Location' header
  --dump=DUMP       Dump HTTP request and response. Possible values are 'debug' or 'json'.
  --pp              Pretty print response body
```

### Authentication

`rsc` can read the API host, account and authentication information from a config file instead of the
command line. See [Setup and Config](#config) below.

Authentication can be done in one of the following ways: 
* Using a RightScale user email and password (`--email` and `--password`)
* Using a OAuth token retrieved from the _API Credentials_ entry of the _Settings_ menu (`--key`)
* Using an API instance token to make API calls from a RightScale instance (`--key`)
* Using the RightLink10 proxy to make API calls from a RightScale instance running
  RightLink10 (`--rl10`)

### Extracting values from responses

The `--x1`, `--xm` and `--xj` flags make it possible to extract values from the response using a
JSON select expression (see [http://jsonselect.org/](http://jsonselect.org/)). For example:
```
$ rsc --xm .name cm15 index clouds
```
extracts the names of each cloud from the response and prints the result as a space separated list which
is convenient to consume from bash scripts.

For additional help on extracting values see the [Command Line Help and Cookbook](COOKBOOK.md).

### Actions and Parameters

The names of the actions available for a given API or a given API resource can be listed with the
special `actions` action:
```
$ rsc cm15 actions
```
(output of example above not shown for brevity)
To get the actions available on a resource specify the resource href as in:
```
$ rsc cm15 actions /api/clouds/1/volumes
Action  Href                              Resource
======= ================================= ========
create  /api/clouds/:cloud_id             Volume
------- --------------------------------- --------
destroy /api/clouds/:cloud_id/volumes/:id Volume
------- --------------------------------- --------
index   /api/clouds/:cloud_id             Volume
------- --------------------------------- --------
show    /api/clouds/:cloud_id/volumes/:id Volume
```
Parameters use url form encoding to represent nested data structures used in request bodies. For example
using the RightScale CM API 1.5 the command line to create a volume is:
```
$ rsc --pp --fetch cm15 create /api/clouds/1/volumes \
  'volume[name]=My New Volume' \
  'volume[size]=10' \
  'volume[datacenter_href]=/api/clouds/1/datacenters/5K443K2CF8NS6' \
  'volume[volume_type_href]=/api/clouds/1/volume_types/BDVEN383N1EN2'
```
The `--pp` and `--fetch` options above are optional, `--fetch` makes a subsequent API call to retrieve
the newly created resource and `--pp` pretty prints the response.
Use the name of the fields followed by `[]` to represent arrays:
```
$ rsc cm16 index /api/deployments 'filter[]=description==awesome deployment' \
  'filter[]=name==app servers'
```
The `/api/` prefix for CM API 1.5 and CM API 1.6 hrefs is optional so the following lists all
deployments in the account using the default view:
```
$ rsc cm16 index deployments
```

### <a name="config"></a>Command Line Tool Setup and Config
`rsc` has a top level `setup` command which creates a rsc config file. The config file contains the
RightScale account ID, API host, user email and (encrypted) password so that these flags don't have
to be provided each time the tool is invoked.

By default the config file is created in `$HOME/.rsc`, the location can be overridden using the
`--config` global flag. Multiple configs may be created to allow for different environments or
users. Use the `--config` flag when invoking the tool to specify the location of the config file if it's
not the default. The file itself is a simple JSON file that can be edited manually (apart from the
password value that needs to be encrypted by `rsc`).

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
$ rsc cm15 index clouds --help
usage: rsc [<flags>] index index /api/clouds [<cloud index params>]

Cloud index params:
filter[]=[]string
    <optional, [cloud_type|description|name]>

view=string
    <optional, [default|extended]>
```
The help lists the valid values for views and filters for example. It also indicates which flags are mandatory.

-----
## <a name="go"></a>Go Package

The other use case for `rsc` is making programmatic API requests to the
RightScale platform. Each API client code is encapsulated in a different
sub-package: package `cm15` for CM API 1.5, package `cm16`for CM API
1.6 and package `ss` for Self-service APIs.

### Client Creation

Each API client package defines an `Api` struct that represents the API
client. Clients are created using one of three factory methods: `New`,
`NewRL10` or `FromCommandLine`. The latter is used by the top level
`main` package to create clients from the values provided on the command
line. `NewRL10` is meant to be called by code that runs on a RightScale
server running the RightLink10 agent. It configures the client to
talk to the APIs through the proxy exposed by RightLink10. Overall
the most flexible function is `New` which accepts an API host name,
a RightScale account ID, authentication information and an optional
logger and low-level HTTP client. As an example the following creates a
CM API 1.5 client that connects to `us-3.rightscale.com` using a OAuth
refresh token for authentication, no logger and the default HTTP client:

```go
refreshToken := ... // Retrieve refresh token
auth := rsapi.OAuthAuthenticator{RefreshToken: refreshToken}
accountId := 123 // Retrieve account ID
client, err := cm15.New(accountId, "us-3.rightscale.com", &auth, nil, nil)
```

### Locators

Once a client has been created resources can be retrieved using resource locators. The code defines one
resource locator type per resource exposed by the underlying API.

Locators are instantiated using factory methods exposed by the client object. The factory methods
accept the collection or resource href and return the corresponding locator. For example the
following creates a cloud locator:
```go
var cloudLocator = client.CloudLocator("/api/clouds/1")
```
Locators expose one method for each action supported by the underlying collection or resource. For
example the clouds collection locator `CloudLocator` exposes an `Index()` and a `Show()` method.
Locator methods may return resources which are structs that expose the underlying resource 
attributes. For example the `CloudLocator` `Index()` method returns an array of `Cloud` resource.
A cloud resource is defined as:
```go
// Represents a Cloud (within the context of the account in the session).
type Cloud struct {
	Capabilities []map[string]interface{} `json:"capabilities,omitempty"`
	CloudType    string                   `json:"cloud_type,omitempty"`
	Description  string                   `json:"description,omitempty"`
	DisplayName  string                   `json:"display_name,omitempty"`
	Links        []map[string]string      `json:"links,omitempty"`
	Name         string                   `json:"name,omitempty"`
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
The following code would invoke the `Index()` method using the default view and no filter to make the API request:
```go
var clouds, err = api.CloudLocator("/api/clouds").Index(rsapi.ApiParams{})
```
`Create` actions all return a locator so that fetching the corresponding resource is easy:
```go
var volumesLocator = client.VolumeLocator("/api/clouds/1/volumes")
var params cm15.VolumeParam{} // Code that sets parameters omitted for brevity
loc, err := volumeLocator.Create(&params)
if err == nil {
	volume, err := loc.Show(rsapi.ApiParams{})
	// ... check error, use volume etc.
}
```

### Using the Generic Methods

So far we've seen how you can interact with the APIs using strongly
typed methods which are handy if you need to make specific API calls
from your code. However these methods don't work well if you need to
write a generic client that may need to make any arbitrary API call
given a resource, an action and generic parameters.

For this use case each API client package also contains a generic `Do`
method which accepts the name of a resource, the name of an action,
the href of the resource and a map of generic parameters (in the form of
`map[string]interface{}`):
```go
func (a *Api) Do(resource, action, href string, params rsapi.ApiParams) (*http.Response, error)
```
The `ss` package `Do` method additionally requires the name of the SS service to send requests to (i.e.
`designer`, `catalog` or `manager`). This is because there is not one RightScale Self-Service API but
three as noted in the [API reference](#reference) section above.

### Common code

The package `rsapi` contains common code for all client packages. It
also defines an Api struct that each client embeds as an anonymous field
and leverages for all common code. One such method that may be of use in
your code is `LoadResponse` that simply unmarshals the response body JSON
and returns the result. If the response contains a `Location` header (all
`create` actions return one) then the function returns a map containing
the value of the location under the `"Location"` key. The signature of
`LoadResponse` is:
```go
func (a *Api) LoadResponse(resp *http.Response) (interface{}, error)
```
The `rsapi` package also includes authenticators which signs API requests
by adding the required auth headers (cookie in the case of email/password
authentication, OAuth header in the case of OAuth and custom header in
the case of RightLink10). Finally it contains common code used by all
the clients to parse the command line.

-----
## <a name="contributing"></a>Development & Contributing

### Building

To build `rsc` from source, first run `godep restore` to retrieve all its dependencies and once that's done
run `make`.
```
godep restore && make
```
### Code generation

Part of the `rsc` source code (the vast majority in terms of lines of code) is automatically 
generated from API metadata. There are currently two code generators: `api15gen` consumes the
RightScale CM API 1.5 metadata hosted [here](http://reference.rightscale.com/api1.5/api_data.json) and
`praxisgen` consumes the metadata for any [praxis](http://praxis-framework.io/) application (for example
for the RightScale CM API 1.6).

The source code for the code generator tools lives under the `gen` directory.
Once the tools are compiled and installed they can be invoked using `go generate`,
see [http://blog.golang.org/generate](http://blog.golang.org/generate)
for information on how `go generate` works. The `go generate` comments live in the top level file
`generate.go`:
```go
//go:generate api15gen -metadata=cm15 -output=cm15
//go:generate praxisgen -metadata=cm16/api_docs -output=cm16 -pkg=cm16 -target=1.6 -client=Api
//go:generate praxisgen -metadata=ss/ssd/restful_doc -output=ss/ssd -pkg=ssd -target=1.0 -client=Api
//go:generate praxisgen -metadata=ss/ssc/restful_doc -output=ss/ssc -pkg=ssc -target=1.0 -client=Api
//go:generate praxisgen -metadata=ss/ssm/restful_doc -output=ss/ssm -pkg=ssm -target=1.0 -client=Api
```
When invoked the `api15gen` and `praxisgen` tools generate the `codegen_client.go` and `codegen_metadata.go`
for each API client in their directory.

The Makefile included in the source code contains a `rsc` target (the default target) that can be
invoked to generate the code and build the command line tool.

### Other Directories

The `ss` package consists of 3 sub-packages, one per self-service API: The `ssd` sub-package
contains the code for making requests to the RightScale Self-Service designer API, the `ssc` to the
RightScale Self-Service catalog API and finally `ssm` to the RightScale Self-Service manager API.
These are sub-packages instead of top-level packages so that they can be wrapped and exposed through
a single entry point (both for the command line and the go client methods).

The `metadata` directory defines the data types generated by `api15gen` and `praxisgen`. These
include data types that describe resources, actions and action parameters. These data structures
are used by the generated code and by the command line parsing code to provide contextual help
for example.

Finally, the `cmd` package contains common code used by `rsc` and each API command line processing
code to access command line flags.

### Contributing

PRs are always welcome, use github issues to report problems.

#### Adding Support to a New RightScale API - or any Praxis application

To add support for a new API:
  1. Create a subdirectory whose name matches the name of the go package (`cm15`, `ss`, etc.).
  2. Put the JSON metadata describing the API in that directory (`api_docs` directory).
  3. Add a `go generate` directive that invokes the generator against the JSON to `generate.go`.
  4. Generate the code and add the corresponding command line parsing and `Do` methods (see
     `commands.go`, `http.go` and `cm16.go` in the cm16 directory for an example).
  5. Add the corresponding sub-command to `rsc` (see the top level `command_line.go` file).

The JSON metadata for praxis apps is the JSON generated by the `rake praxis:api_docs` command.
Updating a client to the latest version of an API thus consists of updating the corresponding JSON
and rebuilding the client (the `go generate` directives will take care of updating the generated
code).

## License

The `rsc` source code is subject the MIT license,
see the [LICENSE](https://github.com/rightscale/rsc/edit/master/LICENSE) file.
