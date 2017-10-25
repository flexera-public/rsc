v6.6.0 / 2017-10-25
-------------------
* Add action/priority to security group rules.
* Add source/destination cidr IPs/group names and source start/end port to security group rules. These are currently
  only available for AzureRM.
* Add cloud_specific_params[service_account] to instances

v6.5.0 / 2017-07-19
-------------------
* Add support for AzureRM Managed Disk.
* Fix a bug command line resource href matching where actions whose href pattern partially matched the href were
  considered a full match resulting in the wrong resource action being selected.

v6.4.1 / 2017-06-20
-------------------
* Fix bug introduced with storing refresh token in `~/.rsc` where it took precedence over the `-r`/`--refresh-token`
  flag rather than the flag taking precendence over the config file value like it should.

v6.4.0 / 2017-06-14
-------------------
* Build with Go 1.8 which fixes JSON output to prefer decimal notation rather than scientific notation which was messing
  up some IDs (see [encoding/json](https://golang.org/doc/go1.8#encoding_json) in the Go 1.8 Release Notes).
* Add `httpclient.NewPB` constructor which will create an `HTTPClient` from a `ParamBlock` rather than using package
  global variables which can be unsafe with multiple Goroutines.
* Use `context` instead of `golang.org/x/net/context`.
* Support refresh token in `~/.rsc` configuration file for the command line tool.

v6.3.0 / 2017-03-02
------------------
* Added ResourceGroup resource
* Added MultiCloudImageMatcher resource
* Updated fields for Subnet, Instance, ServerArray

v6.2.0 / 2016-08-29
-------------------
* rl10: Added /rll/proc/log_level.
* Add destroy action for alerts.
* Fields for Instance are updated.

v6.1.0 / 2016-08-09
--------
* SelfService action paths have been updated to the latest version of the SelfService API. This means that each path has been prefixed with one of "/api/catalog", "/api/design", or "/api/manager" depending on the service it is associated with. The built-in documentation has been updated with the new paths.
* The old SelfService action paths are now deprecated. While they still work as before they will no longer be supported in the next major release of RSC.
* Added support for the EndUser resource in the SelfService Catalog API.

v6.0.0 / 2016-06-28
--------
* Fields for Instance, Route, SshKey, Volume, User, and Permission types are updated.
* Instances#set_custom_lodgement is removed.
* InstanceCustomLodgement is removed.
* Added RightScript#delete, Servers#disable_runnable_bindings, Servers#enable_runnable_bindings, ServerArrays#disable_runnable_bindings, and ServerArrays#enable_runnable_bindings, and VolumeSnapshots#copy.
* Sessions#index now accepts a view argument and the supported views are: default and whoami.
* Parameters for Instances#create, Instances#update, Routes#create, Servers#create, Servers#update, ServerArrays#create, and Volumes#create updated.
* rl10: Add managed login actions

v5.0.3 / 2016-04-21
-------------------
* Add --retry flag. Specifies number of retry attempts for non-successful API responses(500, 503, and timeouts only)
* cm15: Add RightScriptAttachment actions. Add RightScript delete and update_source actions.
* rl10: Add Docker integration actions.

v5.0.2 / 2016-02-05
-------------------
* Using TokenAuthenticator now sends the account ID in the X-Account header for all requests
* Handle Unicode byte order mark in JSON input since Windows PowerShell variables can include them
* Do not use a timeout when the request is associated with a context
* Support new custom monitoring plugin actions in RL10 API
* Increase the default timeout for the rsc command line client to 5 minutes (300 seconds)

v5.0.1 / 2016-01-07
-------------------
* Update gomega to fix HasOccurred
* Fix ambiguous handling of some command line arguments
* Add --timeout flag
* Add link information to resource metadata

v5.0.0 / 2015-10-15
-------------------
* Fix issue with instance session auth
* Fix optional command line parameters being required
* Properly generate code and metadata for "aliased" CM 1.5 resource actions
* [break] Fix name of array query parameters in metadata (add "[]" suffix)

v4.0.1 / 2015-09-18
-------------------
* Support new TSS control actions in RL10 API
* cm15: Fix typo in server terminate action (thank you @manuelfelipe)
* cm15: Fix local disk size attribute for InstanceType (thank you @manuelfelipe)

v4.0.0 / 2015-08-25
-------------------
* [break] Fix issues reported by `golint`, e.g. `ApiParams` => `APIParams`
* [break] Fix return types for APIs that don't return resources
* Add TSS control actions to RL10 client

v3.1.1 / 2015-08-24
-------------------
* Update to latest kingpin v2
* Fix issue with cm15 package including server creation

v3.1.0 / 2015-08-14
-------------------
* Proxy support: honor HTTP_PROXY and HTTPS_PROXY environment variables
* Add cloud specific attributes to servers and instances
* Fix issue with request timing logging

v3.0.2 / 2015-07-21
-------------------
* Add support for file uploads (self-service templates creation)
* Fix praxisgen bug causing incorrect command line for ss
* Fix request ID logging

v3.0.1 / 2015-07-17
-------------------
* Fix request logging so it doesn't require dumping as well
* Fix dump output formatting
* Add proper error handling for invalid HTTP status codes
* Fix auditail example

v3.0.0 / 2015-07-16
-------------------
* Add support for Cloud Analytics APIs (via `ca` subcommand and package)
* Make it possible to specify account when using OAuth (fixes CM 1.6 auth)
* Add `Locator` method to all resources
* Refactor logging, expose logger via new `log` package
* Refactor how low level HTTP client is created, expose config via new `httpclient` package
* Fix issues with some CM 1.5 resource attribute types
* Make request timeouts configurable in `httpclient` package
* Make certificate validation optional (see `httpclient.NoCertCheck`)
* BREAK: remove need for providing http.Client instance when creating API client

v2.0.2 / 2015-06-09
-------------------
* Fix issues with authentication using the "ss" command line
* Remove need to pass HTTP scheme to Authenticator.CanAuthenticate

v2.0.1 / 2015-06-02
-------------------
* Remove automatic credential validation on creation
* Add 'CanAuthenticate' method to clients that makes credential test API request
* Add 'Insecure' method to clients that force use of HTTP instead of HTTPS
* Add '--verbose [-v]' flag that dumps auth requests and headers
* Properly handle redirects returned when creating new sessions on incorrect host
* Help is now printed on stdout instead of stderr

v2.0.0 / 2015-05-28
-------------------
* Remove the need to pass account ID to token based authenticators
* Add consistent credentials testing across all authenticators upon creation
* Add ability to pass arguments when running recipes through RL10

v1.0.9 / 2015-05-13
-------------------
* Fix passing of inputs on command line
* Add support for RL10 tss get and set hostname APIs

v1.0.8 / 2015-05-09
-------------------
* Add `--accessToken` flag
* Rename `--key` flag into `--refreshToken` (`--key` still works)
* Add documentation around authentication

v1.0.7 / 2015-05-07
-------------------
* Fix bug when specifying multiple array elements on command line

v1.0.6 / 2015-05-06
-------------------
* Update Self-Service client to match latest release
* Fix command line parsing of array of maps payload attributes

v1.0.5 / 2015-04-30
-------------------
* Add new instance API token authenticator and corresponding "--apiToken" flag
* Fix cm15 package ServerArrayLocator.CurrentInstances()
* Fix crashes in go-jsonselect
* Add tests to package examples

v1.0.4 / 2015-04-23
-------------------
* Fix handling of array arguments on the command line

v1.0.3 / 2015-04-20
-------------------
* Add rsssh example
* Change datetime attributes of cm15 package resources to use `*RubyTime` instead of `RubyTime`
* Change `CurrentInstance` attribute type in cm15 package to use `*Instance` instead of `Instance`

v1.0.2 / 2015-04-19
-------------------
* Update README, add basic example documentation
* Fix `CloudSpecificAttributes` and `DatacenterPolicy` attribute types in cm15 package

v1.0.1 / 2015-04-17
-------------------
* Fix crashes in go-jsonselect triggered with e.g. `rsc --x1 ':root ~ .name' ...`

v1.0.0 / 2015-04-15
-------------------
* Initial release
