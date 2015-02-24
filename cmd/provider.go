package cmd

import "gopkg.in/alecthomas/kingpin.v1"

// Common interface between main rsc package and API client packages that makes it possible for
// client packages to register command line arguments.
// Note that both the kingpin *Application and *CmdClause types implement the CommandProvider
// interface.
type CommandProvider interface {
	Command(name, help string) *kingpin.CmdClause
}
