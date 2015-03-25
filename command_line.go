package main

import (
	"fmt"
	"os"
	"path"

	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/cm16"
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Retrieve command and top level flag values
func ParseCommandLine(app *kingpin.Application) (*cmd.CommandLine, error) {
	// 1. Register all commands
	app.Command("setup", "create config file, defaults to $HOME/.rsc, use '--config' to override")
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
	app.Flag("x1", "Extract single value using JSON:select").StringVar(&cmdLine.ExtractOneSelect)
	app.Flag("xm", "Extract zero, one or more values using JSON:select and return space separated list").StringVar(&cmdLine.ExtractSelector)
	app.Flag("xj", "Extract zero, one or more values using JSON:select and return JSON").StringVar(&cmdLine.ExtractSelectorJson)
	app.Flag("xh", "Extract header with given name").StringVar(&cmdLine.ExtractHeader)
	app.Flag("noRedirect", "Do not follow redirect responses").Short('n').BoolVar(&cmdLine.NoRedirect)
	app.Flag("fetch", "Fetch resource with href present in 'Location' header").BoolVar(&cmdLine.FetchResource)
	d := &cmdLine.Dump // strange, not sure why Kingpin forces that
	app.Flag("dump", "Dump HTTP request and response. Possible values are 'debug' or 'json'.").EnumVar(&d, "debug", "json")
	app.Flag("pp", "Pretty print response body").BoolVar(&cmdLine.Pretty)

	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"--help"}
	}
	cmd, err := app.Parse(args)
	if err != nil && len(args) > 0 {
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

	// 4. Validate we have everything we need
	validateCommandLine(&cmdLine)

	// 5. Set command and we're done
	cmdLine.Command = cmd
	return &cmdLine, nil
}

// Make sure all the required information is there
func validateCommandLine(cmd *cmd.CommandLine) {
	if cmd.Account == 0 {
		kingpin.Fatalf("missing --account option")
	}
	if cmd.Host == "" {
		kingpin.Fatalf("missing --host option")
	}
	if cmd.Password == "" && cmd.Token == "" && !cmd.RL10 {
		kingpin.Fatalf("missing login info, use --email and --password or use --key or --rl10")
	}
}

// Update the code below when adding new clients. This is the only place that needs to be changed.

// List all client commands below
const (
	// Command for API 1.5 client
	Cm15Command = "cm15"

	// Command for API 1.6 client
	Cm16Command = "cm16"

	// Command for SS client
	SsCommand = "ss"
)

// Instantiate client with given name from command line arguments
func ApiClient(name string, cmdLine *cmd.CommandLine) (cmd.CommandClient, error) {
	switch name {
	case Cm15Command:
		return cm15.FromCommandLine(cmdLine)
	case Cm16Command:
		return cm16.FromCommandLine(cmdLine)
	case SsCommand:
		return ss.FromCommandLine(cmdLine)
	default:
		return nil, fmt.Errorf("No client for '%s'", name)
	}
}

// Register all API client commands
func RegisterClientCommands(app *kingpin.Application) {
	cm15Cmd := app.Command(Cm15Command, cm15.ApiName)
	registrar := rsapi.Registrar{ApiCmd: cm15Cmd}
	cm15.RegisterCommands(&registrar)

	cm16Cmd := app.Command(Cm16Command, cm16.ApiName)
	registrar = rsapi.Registrar{ApiCmd: cm16Cmd}
	cm16.RegisterCommands(&registrar)

	ssCmd := app.Command(SsCommand, ss.ApiName)
	registrar = rsapi.Registrar{ApiCmd: ssCmd}
	ss.RegisterCommands(&registrar)
}
