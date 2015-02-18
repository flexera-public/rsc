package rsapi15

import (
	"fmt"
	"net/http"
	"strings"

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

// Flag values indexed by flag name indexed by command name
type FlagValues map[string]map[string]interface{}

// Map that holds flag values resulting from command line parsing
var flagValues FlagValues

// Register all commands with kinpin application
func RegisterCommands(app *kingpin.Application) {
	flagValues = make(FlagValues)
	for _, command := range commands {
		var resCmd = app.Command(command.Name, command.Description)
		for _, action := range command.Actions {
			var actionCmd = resCmd.Command(action.Name, action.Description)
			var actionFlagValues = make(map[string]interface{})
			var flagKey = fmt.Sprintf("%s %s", command.Name, action.Name)
			flagValues[flagKey] = actionFlagValues
			for _, flag := range action.Flags {
				var cause *kingpin.FlagClause
				cause = actionCmd.Flag(flag.Name, flag.Description)
				if flag.Mandatory {
					cause = cause.Required()
				}
				switch flag.Type {
				case "string":
					actionFlagValues[flag.Name] = cause.String()
				case "[]string":
					actionFlagValues[flag.Name] = cause.Strings()
				case "int":
					actionFlagValues[flag.Name] = cause.Int()
				case "map":
					actionFlagValues[flag.Name] = cause.StringMap()
				}

			}
		}

	}
}

// Actually run command
func (a *Api15) RunCommand(cmd, href string) (*http.Response, error) {
	var elems = strings.Split(cmd, " ")
	if len(elems) != 2 {
		return nil, fmt.Errorf("Invalid command '%s'", cmd)
	}
	return a.Do(href, elems[2], flagValues[cmd])
}
