package rsapi15

import (
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Resource command
type ResourceCmd struct {
	Name        string
	Description string
	Actions     []*ActionCmd
}

// Resource action subcommand
type ActionCmd struct {
	Name        string
	Description string
	Flags       []*ActionFlag
}

// Resource action subcommand flag
type ActionFlag struct {
	Name        string
	Description string
	Type        string
	IsPattern   bool
	Mandatory   bool
	NonBlank    bool
	Regexp      string
	ValidValues []string
}

// Register all commands with kinpin application
func RegisterCommands(app *kingpin.Application) {
	for _, command := range commands {
		var resCmd = app.Command(command.Name, command.Description)
		for _, action := range command.Actions {
			var actionCmd = resCmd.Command(action.Name, action.Description)
			for _, flag := range action.Flags {
				var cause *kingpin.FlagClause
				cause = actionCmd.Flag(flag.Name, flag.Description).Dispatch(buildParam)
				if flag.Mandatory {
					cause = cause.Required()
				}
				switch flag.Type {
				case "string":
					cause.String()
				case "[]string":
					cause.Strings()
				case "int":
					cause.Int()
				case "map":
					cause.StringMap()
				}

			}
		}

	}
}

func (a *Api15) RunCommand(cmd string) string {
	return "42"
}

func buildParam(c *kingpin.ParseContext) error {
	spew.Dump("CONTEXT:", c)
	return nil
}
