package main

import (
	"os"
	"path"

	"github.com/rightscale/rsc/cmds"
	"github.com/rightscale/rsc/rsapi15"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Command and top level flags
// API clients register additional sub-commands with their own flags
type CommandLine struct {
	Command             string // Command to be run (e.g. "api15 index")
	ConfigPath          string // Path to rsc config file, defaults to $HOME/.rsc
	Account             int    // RightScale account, optional
	Host                string // API hostname, optional
	Token               string // Auth token, required unless RL10 is true
	RL10                bool   // Whether to send requests using the RL10 proxy
	NoRedirect          bool   // Whether to disable auto-redirects
	FetchResource       bool   // Whether to fetch resource returned in 'Location' header
	ExtractOneSelect    string // JSON select expression to extract single value from response, optional
	ExtractSelector     string // JSON select expression to extract zero or more values from response, optional
	ExtractSelectorJson string // JSON select expression to extract zero or more values from response, extracted values are displayed using JSON encoding, optional
	ExtractHeader       string // Name of header to extract from response, optional
	Dump                bool   // Whether to dump raw HTTP request and response to stdout
	Pretty              bool   // Whether to display response body or extract values using pretty printer
	ShowHelp            bool   // Whether to show help for action flags
}

// Retrieve command and top level flag values
func ParseCommandLine(app *kingpin.Application) (*cmds.CommandLine, error) {
	// 1. Register all commands
	app.Command("setup",
		"create config file, defaults to $HOME/.rsc, use '--config' to override")
	RegisterClientCommands(app)

	// 2. Parse flags
	var flags = cmds.CommandLine{}
	app.Flag("config", "path to rsc config file").Short('c').Default(path.Join(os.Getenv("HOME"), ".rsc")).StringVar(&flags.ConfigPath)
	app.Flag("account", "RightScale account ID").Short('a').IntVar(&flags.Account)
	app.Flag("host", "RightScale API host (e.g. 'us-3.rightscale.com')").Short('h').StringVar(&flags.Host)
	app.Flag("key", "OAuth access token or API key").Short('k').StringVar(&flags.Token)
	app.Flag("rl10", "Proxy requests through RightLink 10 (exclusive with '--key')").BoolVar(&flags.RL10)
	app.Flag("x1", "Extract single value using given JSON:select expression").StringVar(&flags.ExtractOneSelect)
	app.Flag("xm", "Extract zero, one or multiple values using given JSON:select expression and return space separated list (useful for bash scripts)").StringVar(&flags.ExtractSelector)
	app.Flag("xj", "Extract zero, one or multiple values using given JSON:select expression and return JSON").StringVar(&flags.ExtractSelectorJson)
	app.Flag("xh", "Extract header with given name").StringVar(&flags.ExtractHeader)
	app.Flag("noRedirect", "Do not follow redirect responses").Short('n').BoolVar(&flags.NoRedirect)
	app.Flag("fetch", "Fetch resource with href present in 'Location' header").BoolVar(&flags.FetchResource)
	app.Flag("dump", "Dump HTTP request and response for debugging").BoolVar(&flags.Dump)
	app.Flag("pp", "Pretty print response body").BoolVar(&flags.Pretty)

	var args = os.Args[1:]
	var cmd, err = app.Parse(args)
	if err != nil {
		// This is a bit hacky: basically doing `rsc api15 index clouds --help` results
		// in a command line that kingpin is unable to parse. So capture the `--help` and
		// retry parsing without it.
		var lastArgIndex = len(args)
		var help = args[lastArgIndex-1]
		if help == "--help" || help == "-h" || help == "-help" || help == "-?" {
			flags.ShowHelp = true
			lastArgIndex -= 1
			cmd, err = app.Parse(args[:lastArgIndex])
			if err != nil {
				return nil, err
			}
		}
	}

	// 3. Complement with defaults from config at given path
	if config, err := LoadConfig(flags.ConfigPath); err == nil {
		if flags.Account == 0 {
			flags.Account = config.Account
		}
		if flags.Token == "" {
			flags.Token = config.Token
		}
		if flags.Host == "" {
			flags.Host = config.Host
		}
	}

	// 4. Set command and we're done
	flags.Command = cmd
	return &flags, nil
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
)

var (
	// List of all supported client commands
	AllCommands = []string{Api15Command}

	// Default command
	DefaultClientCommand string
)

// Register all API client commands
func RegisterClientCommands(app *kingpin.Application) {
	var api15Cmd = app.Command(Api15Command, "RightScale API 1.5 client")
	rsapi15.RegisterCommands(api15Cmd)

	// Register default client commands, hard coded for now (only one client at the moment...)
	DefaultClientCommand = Api15Command
	rsapi15.RegisterCommands(app)
}
