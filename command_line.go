package main

import (
	"os"
	"path"

	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/cm16"
	"github.com/rightscale/rsc/cmd"
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
	cmdLine := cmd.CommandLine{}
	app.Flag("config", "path to rsc config file").Short('c').Default(path.Join(os.Getenv("HOME"), ".rsc")).StringVar(&cmdLine.ConfigPath)
	app.Flag("account", "RightScale account ID").Short('a').IntVar(&cmdLine.Account)
	app.Flag("host", "RightScale login endpoint (e.g. 'us-3.rightscale.com')").Short('h').StringVar(&cmdLine.Host)
	app.Flag("email", "Login email, use --email and --password or use --key or --rl10").StringVar(&cmdLine.Username)
	app.Flag("pwd", "Login password, use --email and --password or use --key or --rl10").StringVar(&cmdLine.Password)
	app.Flag("key", "OAuth access token or API key, use --email and --password or use --key or --rl10").Short('k').StringVar(&cmdLine.Token)
	app.Flag("rl10", "Proxy requests through RL 10 agent, use --email and --password or use --key or --rl10").BoolVar(&cmdLine.RL10)
	app.Flag("hrefs", "List all known href patterns for selected API or resource").BoolVar(&cmdLine.ShowHrefs)
	app.Flag("x1", "Extract single value using JSON:select").StringVar(&cmdLine.ExtractOneSelect)
	app.Flag("xm", "Extract zero, one or more values using JSON:select and return space separated list").StringVar(&cmdLine.ExtractSelector)
	app.Flag("xj", "Extract zero, one or more values using JSON:select and return JSON").StringVar(&cmdLine.ExtractSelectorJson)
	app.Flag("xh", "Extract header with given name").StringVar(&cmdLine.ExtractHeader)
	app.Flag("noRedirect", "Do not follow redirect responses").Short('n').BoolVar(&cmdLine.NoRedirect)
	app.Flag("fetch", "Fetch resource with href present in 'Location' header").BoolVar(&cmdLine.FetchResource)
	app.Flag("dump", "Dump HTTP request and response for debugging").BoolVar(&cmdLine.Dump)
	app.Flag("pp", "Pretty print response body").BoolVar(&cmdLine.Pretty)

	args := os.Args[1:]
	cmd, err := app.Parse(args)
	if err != nil {
		// This is a bit hacky: basically doing `rsc api15 index clouds --help` results
		// in a command line that kingpin is unable to parse. So capture the `--help` and
		// retry parsing without it.
		lastArgIndex := len(args)
		help := args[lastArgIndex-1]
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
		if cmdLine.Username == "" {
			cmdLine.Username = config.Email
		}
		if cmdLine.Password == "" {
			cmdLine.Password = config.Password
		}
		if cmdLine.Host == "" {
			cmdLine.Host = config.LoginHost
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
	Cm15Command = "cm15"

	// Command for API 1.6 client
	Cm16Command = "cm16"

	// Command for SS client
	SSCommand = "ss"
)

var (
	// List of all supported client commands
	AllCommands = []string{Cm15Command, Cm16Command, SSCommand}

	// Default command
	DefaultClientCommand string
)

// Register all API client commands
func RegisterClientCommands(app *kingpin.Application) {
	cm15Cmd := app.Command(Cm15Command, "RightScale CM API 1.5 client")
	cm15.RegisterCommands(cm15Cmd)

	cm16Cmd := app.Command(Cm16Command, "RightScale CM API 1.6 client")
	cm16.RegisterCommands(cm16Cmd)

	ssCmd := app.Command(SSCommand, "RightScale SS API 1.0 client")
	ss.RegisterCommands(ssCmd)
}
