package main

import (
	"os"
	"path"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi15"
	"github.com/rightscale/rsc/rsapi16"
	"github.com/rightscale/rsc/ss"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Retrieve command and top level flag values
func ParseCommandLine(app *kingpin.Application) (*cmd.CommandLine, error) {
	// 1. Register all commands
	app.Command("setup",
		"create config file, defaults to $HOME/.rsc, use '--config' to override")
	RegisterClientCommands(app)

	// 2. Parse flags
	var cmdLine = cmd.CommandLine{}
	app.Flag("config", "path to rsc config file").Short('c').Default(path.Join(os.Getenv("HOME"), ".rsc")).StringVar(&cmdLine.ConfigPath)
	app.Flag("account", "RightScale account ID").Short('a').IntVar(&cmdLine.Account)
	app.Flag("host", "RightScale API host (e.g. 'us-3.rightscale.com')").Short('h').StringVar(&cmdLine.Host)
	app.Flag("key", "OAuth access token or API key").Short('k').StringVar(&cmdLine.Token)
	app.Flag("rl10", "Proxy requests through RightLink 10 (exclusive with '--key')").BoolVar(&cmdLine.RL10)
	app.Flag("x1", "Extract single value using given JSON:select expression").StringVar(&cmdLine.ExtractOneSelect)
	app.Flag("xm", "Extract zero, one or multiple values using given JSON:select expression and return space separated list (useful for bash scripts)").StringVar(&cmdLine.ExtractSelector)
	app.Flag("xj", "Extract zero, one or multiple values using given JSON:select expression and return JSON").StringVar(&cmdLine.ExtractSelectorJson)
	app.Flag("xh", "Extract header with given name").StringVar(&cmdLine.ExtractHeader)
	app.Flag("noRedirect", "Do not follow redirect responses").Short('n').BoolVar(&cmdLine.NoRedirect)
	app.Flag("fetch", "Fetch resource with href present in 'Location' header").BoolVar(&cmdLine.FetchResource)
	app.Flag("dump", "Dump HTTP request and response for debugging").BoolVar(&cmdLine.Dump)
	app.Flag("pp", "Pretty print response body").BoolVar(&cmdLine.Pretty)

	var args = os.Args[1:]
	var cmd, err = app.Parse(args)
	if err != nil {
		// This is a bit hacky: basically doing `rsc api15 index clouds --help` results
		// in a command line that kingpin is unable to parse. So capture the `--help` and
		// retry parsing without it.
		var lastArgIndex = len(args)
		var help = args[lastArgIndex-1]
		if help == "--help" || help == "-h" || help == "-help" || help == "-?" {
			cmdLine.ShowHelp = true
			lastArgIndex -= 1
			cmd, err = app.Parse(args[:lastArgIndex])
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 3. Complement with defaults from config at given path
	if config, err := LoadConfig(cmdLine.ConfigPath); err == nil {
		if cmdLine.Account == 0 {
			cmdLine.Account = config.Account
		}
		if cmdLine.Token == "" {
			cmdLine.Token = config.Token
		}
		if cmdLine.Host == "" {
			cmdLine.Host = config.Host
		}
	}

	// 4. Set command and we're done
	cmdLine.Command = cmd
	return &cmdLine, nil
}

// Check whether given command corresponds to a API client command
func IsClientCommand(cmd string) bool {
	for _, c := range AllCommands {
		if cmd == c {
			return true
		}
	}
	return false
}

// Update the code below when adding new clients. This is the only place that needs to be changed.

// List all client commands below
const (
	// Command for API 1.5 client
	Api15Command = "api15"

	// Command for API 1.6 client
	Api16Command = "api16"

	// Command for SS client
	SSCommand = "ss"
)

var (
	// List of all supported client commands
	AllCommands = []string{Api15Command, Api16Command, SSCommand}

	// Default command
	DefaultClientCommand string
)

// Register all API client commands
func RegisterClientCommands(app *kingpin.Application) {
	var api15Cmd = app.Command(Api15Command, "RightScale API 1.5 client")
	rsapi15.RegisterCommands(api15Cmd)

	var api16Cmd = app.Command(Api16Command, "RightScale API 1.6 client")
	rsapi16.RegisterCommands(api16Cmd)

	var ssCmd = app.Command(SSCommand, "Self-Service client")
	ss.RegisterCommands(ssCmd)

	// Register default client commands, hard coded for now (only one client at the moment...)
	DefaultClientCommand = Api15Command
	rsapi15.RegisterCommands(app)
}
