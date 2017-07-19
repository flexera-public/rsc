# rsc - A generic RightScale API client
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/rightscale/rsc/blob/master/LICENSE) [![Godoc](https://godoc.org/github.com/rightscale/rsc?status.svg)](http://godoc.org/github.com/rightscale/rsc)

Master
[![Build Status](https://travis-ci.org/rightscale/rsc.svg?branch=master)](https://travis-ci.org/rightscale/rsc)

v6.5.0
[![Build Status](https://travis-ci.org/rightscale/rsc.svg?branch=v6.5.0)](https://travis-ci.org/rightscale/rsc)

`rsc` provides both a command line tool and a go package for interacting with the RightScale APIs.
The currently supported APIs are the RightScale Cloud Management API 1.5 and 1.6 APIs, the
RightScale Self-Service 1.0 APIs (latest version for this product) and the RightLink10 APIs exposed
by the RightLink10 agent. The RightScale APIs reference can be found on the
[RightScale docs](http://docs.rightscale.com/api/) site.

`rsc` can be used in one of two ways:

* as a command line tool
* as a way to make API requests to RightScale programmatically from Go code

## <a name="usage"></a>Command Line Tool

The command line tool uses subcommands to interact with each API. Use `rsc cm15` to send requests
to the RightScale Cloud Management API 1.5, `rsc cm16` to send requests to the RightScale Cloud
Management API 1.6, `rsc ss` to send requests to the RightScale Self-Service API 1.0 and `rsc rl10`
to send requests to RightLink10.

### Installation

The `rsc` command line tool is a statically linked binary making installation a breeze. There is
no dependency on any runtime library. Just download the correct version for your OS and
architecture and you're good to go.

The latest stable versions can be download from:
- MacOS X: `https://binaries.rightscale.com/rsbin/rsc/v6/rsc-darwin-amd64.tgz`
- Windows: `https://binaries.rightscale.com/rsbin/rsc/v6/rsc-windows-amd64.zip`
- Linux: `https://binaries.rightscale.com/rsbin/rsc/v6/rsc-linux-amd64.tgz`
- ODroid/RasPi/armhf: `https://binaries.rightscale.com/rsbin/rsc/v6/rsc-linux-arm.tgz`

As an example the following downloads and runs the MacOS X version:
```
$ curl https://binaries.rightscale.com/rsbin/rsc/v6/rsc-darwin-amd64.tgz | tar -zxf - -O rsc/rsc > rsc
$ chmod +x ./rsc
$ ./rsc --version
rsc v6.5.0 - 2017-07-19 00:34:20 - 308f4fcbf12da4b94d6fd683c25dd86deb35237e
```

#### Versioning

- To download the latest stable use the links with 'v6' in them.
- To download a specific version, replace the 'v6' by the exact version, such as 'v6.5.0'.
- All versions with the same major number (e.g. 'v6') are intended to be "upward" compatible.
- The 'v6' links download a specific version, so `rsc --version` will print something like 'v6.5.0'
  and not 'v6'.
- The latest dev version is 'master'.

### Command Line

The general shape of a command line is:

```
$ rsc [GLOBAL] [API] ACTION HREF [PARAM=VALUE]
```
where:
- `GLOBAL` is an optional list of global flags,
- `API` is `cm15`, `cm16`, `ss` or `rl10`
- `ACTION` is the API action to perform (i.e `index`, `show`, `update`, etc.),
- `HREF` is the resource or resource collection href (i.e. `/api/servers`, `/api/servers/1` etc.), and
- `PARAM` and `VALUE` are the names and values of the action parameters (e.g. `view=extended`).

For example:
```
$ rsc cm15 show /api/servers 'filter[]=name==LB'
```
The sections below cover each option in order.

### Global Flags

The list of global flags is:
```
  --help           Show help.
  --version        Show application version.
  -c, --config="/home/raphael/.rsc"
                   path to rsc config file
  -R, --retry=0    Number of retry attempts for non-successful API responses (500, 503, and timeouts only)
  -a, --account=ACCOUNT
                   RightScale account ID
  -h, --host=HOST  RightScale login endpoint (e.g. 'us-3.rightscale.com')
  --email=EMAIL    Login email, use --email and --pwd or use --refreshToken, --accessToken, --apiToken or --rl10
  --pwd=PWD        Login password, use --email and --pwd or use --refreshToken, --accessToken, --apiToken or --rl10
  -r, --refreshToken=REFRESHTOKEN
                   OAuth refresh token, use --email and --pwd or use --refreshToken, --accessToken, --apiToken or --rl10
  -s, --accessToken=ACCESSTOKEN
                   OAuth access token, use --email and --pwd or use --refreshToken, --accessToken, --apiToken or --rl10
  -p, --apiToken=APITOKEN
                   Instance API token, use --email and --pwd or use --refreshToken, --accessToken, --apiToken or --rl10
  --rl10           Proxy requests through RightLink 10 agent, use --email and --pwd or use --refreshToken, --accessToken, --apiToken or --rl10
  --noAuth         Make unauthenticated requests, used for testing
  --x1=X1          Extract single value using JSON:select
  --xm=XM          Extract zero, one or more values using JSON:select and return newline separated list
  --xj=XJ          Extract zero, one or more values using JSON:select and return JSON
  --xh=XH          Extract header with given name
  --fetch          Fetch resource with href present in 'Location' header
  --dump=DUMP      Dump HTTP request and response. Possible values are 'debug' or 'json'.
  -v, --verbose    Dump HTTP request and response including auth requests and headers, enables --dump=debug by default, use --dump=json to switch format
  --pp             Pretty print response body
```

### Authentication

There are 3 different mechanisms for authenticating with the RightScale platform.

#### Basic Authentication

Basic authentication uses an email and password to create a session against the Cloud Management
APIs (a.k.a. API 1.5). Creating a session returns a cookie. The session cookie can be used to make
authenticated requests againsts all accounts the authenticated user has access to. While the session
is created by making an API request to the CM 1.5 APIs the resulting cookie can be used to make API
calls against all the RightScale APIs enabled for the user.

`rsc` supports basic authentication via the `--email` and `--pwd` flags. When using `rsc` as a
Go package authentication is done once and the same cookie is used for all API requests made with
the same client. The package also takes care of refreshing the cookie before the session expires.

Below is an example listing all clouds available in a given account using the `rsc` command line
tool with basic authentication:
```
rsc --account $ACCOUNT --email $EMAIL --pwd $PASSWORD --host $HOST cm15 index clouds
```
The example assumes that the ACCOUNT, EMAIL, PASSWORD and HOST environment variables contain the
account id, user email, password and RightScale API host respectively.

At the time of writing there are two possible values for the RightScale API host:
`us-3.rightscale.com` or `us-4.rightscale.com`. When using basic authentication it is also
possible to use `my.rightscale.com` in which case the server will redirect to the appropriate
host. The correct value for the host can be retrieved by logging in into the RightScale Cloud
Management dashboard: the URL contains the appropriate host for the account being logged into.

#### OAuth Authentication

RightScale supports the [OAuth 2.0 Authorization Framework](http://tools.ietf.org/html/rfc6749)
where a refresh token can be exchanged with a temporary access token to make authenticated requests.
The refresh token is created using the RightScale Cloud Management dashboard from the *Settings >
Account Settings > API Credentials* menu. Access tokens are then created using the
[`OAuth2`](http://reference.rightscale.com/api1.5/resources/ResourceOauth2.html#create) resource
exposed by the Cloud Management 1.5 APIs. Note that refresh and access tokens are
*account specific*. This has two consequences: refresh tokens must be retrieved for each account and
it is not necessary to specify the account when making an API request with an access token.

`rsc` supports OAuth authentication via the `--refreshToken` and `--accessToken` flags. If the
refresh token is used then `rsc` takes care of creating an access token and uses the access token
to make the final API request. This results in two API requests which may not be optimal in certain
scenarios (e.g. scripts calling `rsc` to make multiple API requests). For these scenarios `rsc` can
also use a pre-existing access token directly. That access token can be created using the OAuth2
resource.

Below is an example listing all clouds available in the account using the refresh token retrieved
from the RightScale Cloud Management dashboard:
```
rsc --refreshToken $REFRESH --host $HOST cm15 index clouds
```
Note that in this case the account doesn't need to be specified on the command line, it is
inferred from the token.

Here is another example that first creates an access token explicitly then uses that token to list
all clouds:
```
export ACCESS=`rsc --x1 .access_token --refreshToken $REFRESH cm15 create oauth2 grant_type=refresh_token refresh_token=$REFRESH`
rsc --accessToken $ACCESS --host $HOST cm15 index clouds
```
The example above uses the `--x1` flag to extract the access token from the response. Extracting
data from responses is described in the [Extracting values from responses](#extract) section below.

#### Instance Facing APIs

The final mechanism for authenticating against the RightScale APIs consists of using an instance
specific token to make API requests from a RightScale managed instance. This token is written to
the instance user data on launch by the RightScale platform. There are actually two different ways
scripts running on RightScale managed instances can authenticate:

* The instance API token can be used to create a session which returns a cookie. This method is
  similar to basic authentication except that a token is used instead of email and password.
* [RightLink 10](http://docs.rightscale.com/rl/about.html) also runs a HTTP proxy that can be used
  to make authenticated requests without requiring any credentials from the client.

Here is an example using the instance API token to list all clouds:
```
rsc --apiToken $TOKEN --host $HOST cm15 index clouds
```
And here is another example running on a RightLink 10 enabled instance:
```
rsc --rl10 cm15 index clouds
```

#### Storing Client Credentials

The `setup` command can be used to create a configuration file that contains the host, account id,
user email and password so that these don't need to be specified each time. All these settings or
a subset may be stored (i.e. the password doesn't have to be stored if that's not desirable).

By default the config file is created in `$HOME/.rsc`, the location can be overridden using the
`--config` global flag. Multiple configs may be created to allow for different environments or
users. Use the `--config` flag when invoking the tool to specify the location of the config file if it's
not the default.

The configuration file is a simple JSON file that lists the fields defined during setup. The
password is encrypted before being stored although it is a two way encryption scheme and so is not
meant to be a truly secure mechanism but rather a way to avoid having the password stored in plain
text on disk.
```
rsc setup
Account id: 12345
Login email: myemail@mycompany.com
Login password: 12345abc
API Login host: us-3.rightscale.com
```
The values stored in the configuration can be overridden using command line flags so that for
example a different account can be specified:
```
rsc --account 234 cm15 index clouds
```
### <a name=extract></a>Extracting Values From Responses

The `--x1`, `--xm` and `--xj` flags make it possible to extract values from the response using a
JSON select expression (see [https://github.com/lloyd/JSONSelect](https://github.com/lloyd/JSONSelect)). For example:
```
$ rsc --xm .name cm15 index /api/clouds
```
extracts the names of each cloud from the response and prints the result as a newline separated list
which is convenient to consume from bash scripts:
```
$ declare -a "clouds=(`rsc --xm=.cloud_type cm15 index clouds`)"; set | egrep '^clouds'
clouds=([0]="amazon" [1]="open_stack_v2" [2]="cloud_stack" [3]="rackspace_next_gen" [4]="google" [5]="azure" [6]="soft_layer" [7]="vscale")
```

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
Parameters use URL form encoding to represent nested data structures used in request bodies. For example
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
deployments:
```
$ rsc cm16 index deployments
```

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
The help lists the valid values for views and filters for example.
It also indicates which flags are mandatory.

-----
## <a name="go"></a>Go Package

The other use case for `rsc` is making programmatic API requests to the
RightScale platform. Each API client code is encapsulated in a different
sub-package: package `cm15` for CM API 1.5, package `cm16`for CM API
1.6, package `ss` for Self-service APIs and package `rl10` for RightLink10 requests.

### Installation

`rsc` uses gopkg.in for versioning, this means that you can download the released `rsc` packages
as follows:
```
go get gopkg.in/rightscale/rsc.v6
```
and import then in your code with:
```go
import "gopkg.in/rightscale/rsc.v6"
```
If you intend on contributing, just want to play around with the code or feel adventurous you can
download and use the beelding edge version from github which corresponds to the master branch:
```
$ go get github.com/rightscale/rsc
```
```go
import "github.com/rightscale/rsc
```


### Client Creation

Each API client package defines an `API` struct that represents the API
client. Clients are created using one of three factory methods: `New`,
`NewRL10` or `FromCommandLine`. The latter is used by the top level
`main` package to create clients from the values provided on the command
line. `NewRL10` is meant to be called by code that runs on a RightScale
server running the RightLink10 agent. It configures the client to
talk to the APIs through the proxy exposed by RightLink10. Overall
the most flexible function is `New` which accepts an API host name,
a RightScale account ID, authentication information and an optional
low-level HTTP client. As an example the following creates a
CM API 1.5 client that connects to `us-3.rightscale.com` using a OAuth
refresh token for authentication and the default HTTP client:
```go
// Retrieve refresh tokens from the RightScale dashboard Settings/API Credentials menu
refreshToken := "3e040efed9a83ac758f3b1cbdfa041b905742169"
// Corresponding RightScale account ID
accountID := 60073
auth := rsapi.NewOAuthAuthenticator(refreshToken, accountID)
accountId := 123
client := cm15.New("us-3.rightscale.com", &auth, nil)
```
You may test the credentials using the `CanAuthenticate` method:
```go
if err := client.CanAuthenticate(); err != nil {
	// Woops creds are not working
	return err
}
```
This method makes a test API request so is expensive, the idea is to call it once then use the client
to make a series of requests.

### Logging

The `log` package exposes a `Logger` variable of type `log15.Logger`. This logger is used
by rsc to log API requests made to external services. Configure the logger handler
to enable logging. The [log15](https://godoc.org/github.com/inconshreveable/log15) package comes
with a number of default handlers including a file and a syslog handlers.

Configuring the `log` package to log to the file `/var/log/myapp.log` would look like:
```go
handler, err := log15.FileHandler("/var/log/myapp.log", log15.LogfmtFormat())
if err != nil {
	return err
}
log.Logger.SetHandler(handler)
```
Consult the [log15](https://godoc.org/github.com/inconshreveable/log15) GoDocs for more information.

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
// Optional parameters:
// filter
// view
func (loc *CloudLocator) Index(options rsapi.APIParams) ([]*Cloud, error)
```
The following code would invoke the `Index()` method using the default view and no filter to make the API request:
```go
var clouds, err = api.CloudLocator("/api/clouds").Index(rsapi.APIParams{})
```
`Create` actions all return a locator so that fetching the corresponding resource is easy:
```go
var volumesLocator = client.VolumeLocator("/api/clouds/1/volumes")
var params cm15.VolumeParam{} // Code that sets parameters omitted for brevity
loc, err := volumeLocator.Create(&params)
if err == nil {
	volume, err := loc.Show(rsapi.APIParams{})
	// ... check error, use volume etc.
}
```
It is also possible to create a locator directly from a resource by using the resource `Locator`
method:
```
clouds, err := client.CloudLocator("/api/clouds").Index()
if err == nil {
	first := clouds[0].Locator(client)
	first.Show() // first is a CloudLocator instance
}
```

### Using the Generic Methods

So far we've seen how you can interact with the APIs using strongly
typed methods which are handy if you need to make specific API calls
from your code. However these methods don't work well if you need to
write a generic client that may need to make any arbitrary API call
given the names of a resource and an action and generic parameters.

For this use case each API client package also contains a generic `BuildRequest`
method which accepts the name of a resource, the name of an action,
the href of the resource and a map of generic parameters (in the form of
`map[string]interface{}`):
```go
func (a *API) BuildRequest(resource, action, href string, params rsapi.APIParams) (*http.Request, error)
```
The client also exposes a `PerformRequest` method that makes the request and optionally
dumps the request body and response to STDERR for debugging:
```go
func (a *API) PerformRequest(req *http.Request) (*http.Response, error)
```
The `httpclient` package exposes a `DumpFormat` variable that controls how much logging is done
when HTTP requests are done. The default consists of logging the request method and URL as well
as the response code and timing information. Setting `DumpFormat` to `httpclient.Debug` causes
the request and response bodies to get logged as well using the Debug log level.

### Common code

The package `rsapi` contains common code for all client packages. It
also defines an API struct that each client embeds as an anonymous field
and leverages for all common code. One such method that may be of use in
your code is `LoadResponse` that simply unmarshals the response body JSON
and returns the result. If the response contains a `Location` header (all
`create` actions return one) then the function returns a map containing
the value of the location under the `"Location"` key. The signature of
`LoadResponse` is:
```go
func (a *API) LoadResponse(resp *http.Response) (interface{}, error)
```
The `rsapi` package also includes authenticators which signs API requests
by adding the required auth headers (cookie in the case of email/password
authentication, OAuth header in the case of OAuth and custom header in
the case of RightLink10). Finally it contains common code used by all
the clients to parse the command line.

### Examples

The repository includes the following Go examples:

* [Cloud Management 1.5 Basic Example](https://github.com/rightscale/rsc/tree/master/cm15/examples/basic)
* [Cloud Management 1.5 Auditail Example](https://github.com/rightscale/rsc/tree/master/cm15/examples/auditail)
* [Cloud Management 1.5 RS SSH Example](https://github.com/rightscale/rsc/tree/master/cm15/examples/rsssh)
* [Self-Service Basic Example](https://github.com/rightscale/rsc/tree/master/ss/examples/basic)

### Reference Documentation

The documentation for each package is hosted on godoc.org:

* [Cloud Management 1.5](http://godoc.org/github.com/rightscale/rsc/cm15)
* [Cloud Management 1.6](http://godoc.org/github.com/rightscale/rsc/cm16)
* [Self-Service Designer](http://godoc.org/github.com/rightscale/rsc/ss/ssd)
* [Self-Service Catalog](http://godoc.org/github.com/rightscale/rsc/ss/ssc)
* [Self-Service Manager](http://godoc.org/github.com/rightscale/rsc/ss/ssm)
* [RightLink10](http://godoc.org/github.com/rightscale/rsc/rl10)

-----
## <a name="contributing"></a>Development & Contributing

### Building

The following make targets are useful:
- `make depend` installs required tools
- `make` builds a binary for your local OS
- `make build` builds binaries for Linux, OS-X and Windows
- `make test` runs the test suite

#### Your own build of the latest release version

The simple option is `go get gopkg.in/rightscale/rsc.v6`, this will use the checked-in
code-generated files.

The more involved option is:
```
mkdir -p $GOPATH/src/gopkg.in/rightscale
cd $GOPATH/src/gopkg.in/rightscale
git clone https://github.com/rightscale/rsc.git rsc.v6
cd rsc.v6
git checkout v6.2.0
make depend
make
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
`generate.go`.

When invoked the `api15gen` and `praxisgen` tools generate the `codegen_client.go` and `codegen_metadata.go`
for each API client in their directory.

The Makefile takes care of running `go generate` prior to building `rsc`.

#### Adding Support to a New RightScale API - or any Praxis application

As noted above `praxisgen` can be used to generate client code for any Praxis API. The steps
involved are:

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
