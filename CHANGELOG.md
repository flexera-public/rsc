v1.0.7 / 2015-05-7
------------------
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
