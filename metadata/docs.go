// Package metadata contains data structured produced by all API generators and consumed
// by all API clients and command line tools.
// These data structures include information about each resource exposed by the API.
// They describe the actions exposed by each resource in a way that a command line
// tool can easily consume. The idea is to expose one command per action, each
// action data structure thus describes the information required to define the
// action command line flags.
// The data structures also expose methods used by the client to parse the resource
// hrefs and extracts the to build the action URL in a generic way.
package metadata
