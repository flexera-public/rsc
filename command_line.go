package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/rightscale/rsc/ca"
	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/cm16"
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rl10"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss"
	"gopkg.in/alecthomas/kingpin.v2"
)

// ParseCommandLine retrieves the command and top level flag values.
func ParseCommandLine(app *kingpin.Application) (*cmd.CommandLine, error) {
	// 1. Register all commands
	app.Command("setup", "create config file, defaults to $HOME/.rsc, use '--config' to override")
	app.Command("json", "apply jsonselect expression to STDIN")
	RegisterClientCommands(app)

	// 2. Parse flags
	cmdLine := cmd.CommandLine{}
	app.Flag("config", "path to rsc config file").Short('c').Default(path.Join(os.Getenv("HOME"), ".rsc")).StringVar(&cmdLine.ConfigPath)
	app.Flag("retry", "Number of retry attempts for non-successful API responses (500, 503, and timeouts only)").Short('R').Default("0").IntVar(&cmdLine.Retry)
	app.Flag("account", "RightScale account ID").Short('a').IntVar(&cmdLine.Account)
	app.Flag("host", "RightScale login endpoint (e.g. 'us-3.rightscale.com')").Short('h').StringVar(&cmdLine.Host)
	app.Flag("email", "Login email, use --email and --password or use --refreshToken, --accessToken, --apiToken or --rl10").StringVar(&cmdLine.Username)
	app.Flag("pwd", "Login password, use --email and --password or use --refreshToken, --accessToken, --apiToken or --rl10").StringVar(&cmdLine.Password)
	app.Flag("refreshToken", "OAuth refresh token, use --email and --password or use --refreshToken, --accessToken, --apiToken or --rl10").Short('r').StringVar(&cmdLine.OAuthToken)
	app.Flag("accessToken", "OAuth access token, use --email and --password or use --refreshToken, --accessToken, --apiToken or --rl10").Short('s').StringVar(&cmdLine.OAuthAccessToken)
	app.Flag("apiToken", "Instance API token, use --email and --password or use --refreshToken, --accessToken, --apiToken or --rl10").Short('p').StringVar(&cmdLine.APIToken)
	app.Flag("rl10", "Proxy requests through RightLink 10 agent, use --email and --password or use --refreshToken, --accessToken, --apiToken or --rl10").BoolVar(&cmdLine.RL10)
	app.Flag("noAuth", "Make unauthenticated requests, used for testing").BoolVar(&cmdLine.NoAuth)
	app.Flag("timeout", "Set the request timeout, defaults to 300s").Short('t').Default("300").IntVar(&cmdLine.Timeout)
	app.Flag("x1", "Extract single value using JSON:select").StringVar(&cmdLine.ExtractOneSelect)
	app.Flag("xm", "Extract zero, one or more values using JSON:select and return newline separated list").StringVar(&cmdLine.ExtractSelector)
	app.Flag("xj", "Extract zero, one or more values using JSON:select and return JSON").StringVar(&cmdLine.ExtractSelectorJSON)
	app.Flag("xh", "Extract header with given name").StringVar(&cmdLine.ExtractHeader)
	app.Flag("fetch", "Fetch resource with href present in 'Location' header").BoolVar(&cmdLine.FetchResource)
	app.Flag("dump", "Dump HTTP request and response. Possible values are 'debug' or 'json'.").EnumVar(&cmdLine.Dump, "debug", "json", "record")
	app.Flag("verbose", "Dump HTTP request and response including auth requests and headers, enables --dump=debug by default, use --dump=json to switch format").Short('v').BoolVar(&cmdLine.Verbose)
	app.Flag("pp", "Pretty print response body").BoolVar(&cmdLine.Pretty)

	// Keep around for a few releases for backwards compatibility
	app.Flag("key", "OAuth refresh token, use --email and --password or use --refreshToken, --accessToken, --apiToken or --rl10").Short('k').Hidden().StringVar(&cmdLine.OAuthToken)

	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"--help"}
	}
	// This is a bit hacky: basically doing `rsc api15 index clouds --help` results
	// in a command line that kingpin hijacks. So capture the `--help` try parsing
	// without it so we can print our own help.
	lastArgIndex := len(args)
	help := args[lastArgIndex-1]
	var cmd string
	var err error
	if help == "--help" || help == "-h" || help == "-help" || help == "-?" {
		cmdLine.ShowHelp = true
		lastArgIndex--
		cmd, err = app.Parse(args[:lastArgIndex])
	} else {
		cmd, err = app.Parse(args)
	}
	if err != nil {
		return nil, err
	}

	// 3. Complement with defaults from config at given path
	if !cmdLine.NoAuth {
		if config, err := LoadConfig(cmdLine.ConfigPath); err == nil {
			if cmdLine.OAuthAccessToken == "" && cmdLine.OAuthToken == "" {
				if cmdLine.Account == 0 {
					cmdLine.Account = config.Account
				}
				if cmdLine.Username == "" {
					cmdLine.Username = config.Email
				}
				if cmdLine.Password == "" {
					cmdLine.Password = config.Password
				}
			}
			if cmdLine.Host == "" {
				cmdLine.Host = config.LoginHost
			}
			if cmdLine.OAuthToken == "" {
				cmdLine.OAuthToken = config.RefreshToken
			}
		}
	}
	cmdLine.Command = cmd

	// 4. Special RL10 case (auth is handled differently)
	if strings.Split(cmdLine.Command, " ")[0] == "rl10" {
		cmdLine.RL10 = true
	}

	// 6. Validate we have everything we need
	validateCommandLine(&cmdLine)

	// 7. We're done
	return &cmdLine, nil
}

// Make sure all the required information is there
func validateCommandLine(cmdLine *cmd.CommandLine) {
	if cmdLine.Command == "setup" ||
		cmdLine.Command == "actions" ||
		cmdLine.Command == "json" ||
		cmdLine.ShowHelp ||
		cmdLine.RL10 {
		return
	}
	if cmdLine.Account == 0 && cmdLine.OAuthToken == "" && cmdLine.OAuthAccessToken == "" && cmdLine.APIToken == "" && !cmdLine.NoAuth {
		kingpin.Fatalf("missing --account option")
	}
	if cmdLine.Host == "" {
		kingpin.Fatalf("missing --host option")
	}
	if cmdLine.Password == "" && cmdLine.OAuthToken == "" && cmdLine.OAuthAccessToken == "" && cmdLine.APIToken == "" && !cmdLine.NoAuth {
		kingpin.Fatalf("missing login info, use --email and --password or use --key, --apiToken or --rl10")
	}
}

// Update the code below when adding new clients. This is the only place that needs to be changed.

// List all client commands below
const (
	// Cm15Command is the command for API 1.5 client.
	Cm15Command = "cm15"

	// Cm16Command is the command for API 1.6 client.
	Cm16Command = "cm16"

	// SsCommand is the command for SS client.
	SsCommand = "ss"

	// Rl10Command is the command for RL10 client.
	Rl10Command = "rl10"

	// CaCommand is the command for CA client.
	CaCommand = "ca"
)

// APIClient instantiates a client with the given name from command line arguments.
func APIClient(name string, cmdLine *cmd.CommandLine) (cmd.CommandClient, error) {
	switch name {
	case Cm15Command:
		return cm15.FromCommandLine(cmdLine)
	case Cm16Command:
		return cm16.FromCommandLine(cmdLine)
	case SsCommand:
		return ss.FromCommandLine(cmdLine)
	case Rl10Command:
		return rl10.FromCommandLine(cmdLine)
	case CaCommand:
		return ca.FromCommandLine(cmdLine)
	default:
		return nil, fmt.Errorf("No client for '%s'", name)
	}
}

// RegisterClientCommands registers all API client commands.
func RegisterClientCommands(app *kingpin.Application) {
	cm15Cmd := app.Command(Cm15Command, cm15.APIName)
	registrar := rsapi.Registrar{APICmd: cm15Cmd}
	cm15.RegisterCommands(&registrar)

	cm16Cmd := app.Command(Cm16Command, cm16.APIName)
	registrar = rsapi.Registrar{APICmd: cm16Cmd}
	cm16.RegisterCommands(&registrar)

	ssCmd := app.Command(SsCommand, ss.APIName)
	registrar = rsapi.Registrar{APICmd: ssCmd}
	ss.RegisterCommands(&registrar)

	rl10Cmd := app.Command(Rl10Command, rl10.APIName)
	registrar = rsapi.Registrar{APICmd: rl10Cmd}
	rl10.RegisterCommands(&registrar)

	caCmd := app.Command(CaCommand, ca.APIName)
	registrar = rsapi.Registrar{APICmd: caCmd}
	ca.RegisterCommands(&registrar)
}
