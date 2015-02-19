package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/rightscale/rsc/rsapi15"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Command line global flags
type CommandLine struct {
	Commands            []string // Commands and sub-commands to be run (e.g. ["api15", "index"])
	ConfigPath          string   // Path to rsc config file, defaults to $HOME/.rsc
	Account             int      // RightScale account, optional
	Host                string   // API hostname, optional
	Token               string   // Auth token, required unless RL10 is true
	RL10                bool     // Whether to send requests using the RL10 proxy
	NoRedirect          bool     // Whether to disable auto-redirects
	FetchResource       bool     // Whether to fetch resource returned in 'Location' header
	ExtractOneSelect    string   // JSON select expression to extract single value from response, optional
	ExtractSelector     string   // JSON select expression to extract zero or more values from response, optional
	ExtractSelectorJson string   // JSON select expression to extract zero or more values from response, extracted values are displayed using JSON encoding, optional
	ExtractHeader       string   // Name of header to extract from response, optional
	Dump                bool     // Whether to dump raw HTTP request and response to stdout
	Pretty              bool     // Whether to display response body or extract values using pretty printer
	ShowHelp            bool     // Whether to show help for action flags
}

// Command line client entry point.
func main() {
	app := kingpin.New("rsc", "A RightScale API client")
	app.Version("0.1.0")

	// Register API15 subcommands and flags
	rsapi15.RegisterCommands(app)

	var parsed, err = ParseCommandLine(app)
	if err != nil {
		PrintFatal(err.Error())
	}

	// Execute appropriate command
	var resp *http.Response
	switch parsed.Commands[0] {

	case "setup":
		err = CreateConfig(parsed.ConfigPath)
		if err != nil {
			PrintFatal(err.Error())
		}

	case "api15":
		var client *rsapi15.Api15
		var httpClient *http.Client
		if parsed.NoRedirect {
			httpClient = &http.Client{
				CheckRedirect: func(*http.Request, []*http.Request) error {
					return fmt.Errorf("Client configured to prevent redirection")
				},
			}
		} else {
			httpClient = http.DefaultClient
		}
		if parsed.RL10 {
			client, err = rsapi15.NewRL10(nil, httpClient)
		} else {
			client, err = rsapi15.New(parsed.Account, parsed.Token, parsed.Host, nil, httpClient)
		}
		if err != nil {
			PrintFatal("Failed to create API session: %v", err.Error())
		}
		if parsed.ShowHelp {
			var err = client.ShowHelp(parsed.Commands)
			if err != nil {
				PrintFatal(err.Error())
			}
			return
		} else {
			if parsed.Token == "" && !parsed.RL10 {
				PrintFatal("Missing OAuth token, use '-token TOKEN' or 'setup'")
			}
			client.DumpRequestResponse = parsed.Dump
			client.FetchLocationResource = parsed.FetchResource
			resp, err = client.RunCommand(parsed.Commands)
		}

	default:
		PrintFatal("Unknown command '%s'", strings.Join(parsed.Commands, " "))
	}

	// Handle command results
	if err != nil {
		PrintFatal(err.Error())
	}
	if resp == nil {
		return // No results, just exit (may need to revise)
	}

	// Handle command output (apply any extraction)
	displayer, err := NewDisplayer(resp)
	if err != nil {
		PrintFatal(err.Error())
	}
	var moreThanOneError = false
	if parsed.ExtractOneSelect != "" {
		err = displayer.ApplySingleExtract(parsed.ExtractOneSelect)
		moreThanOneError = strings.Contains(err.Error(), "returned more than one value") // Ugh, there has to be a better way
	} else if parsed.ExtractSelector != "" {
		err = displayer.ApplyExtract(parsed.ExtractSelector, false)
	} else if parsed.ExtractSelectorJson != "" {
		err = displayer.ApplyExtract(parsed.ExtractSelectorJson, true)
	} else if parsed.ExtractHeader != "" {
		err = displayer.ApplyHeaderExtract(parsed.ExtractHeader)
	}
	if err != nil {
		PrintFatal(err.Error())
	}
	if parsed.Pretty {
		displayer.Pretty()
	}

	// We're done, print output and figure out correct exit code
	fmt.Printf(displayer.Output())
	var exitStatus = 0
	switch {
	case moreThanOneError:
		exitStatus = 6
	case resp.StatusCode == 401:
		exitStatus = 1
	case resp.StatusCode == 403:
		exitStatus = 3
	case resp.StatusCode == 404:
		exitStatus = 4
	case resp.StatusCode > 399 && resp.StatusCode < 500:
		exitStatus = 2
	case resp.StatusCode > 499:
		exitStatus = 5
	}
	os.Exit(exitStatus)
}

// Retrieve global flag values
func ParseCommandLine(app *kingpin.Application) (*CommandLine, error) {
	var flags = CommandLine{}
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

	app.Command("setup",
		"create config file, defaults to $HOME/.rsc, use '--config' to override")

	var cmd, err = app.Parse(os.Args[1:])
	if err != nil {
		// This is a bit hacky: basically doing `rsc api15 index clouds --help` results
		// in a command line that kingpin is unable to parse. So capture the `--help` and
		// retry parsing without it.
		var help = os.Args[len(os.Args)-1]
		var lastArgIndex = len(os.Args)
		if help == "--help" || help == "-h" || help == "-help" || help == "-?" {
			flags.ShowHelp = true
			lastArgIndex -= 1
			cmd, err = app.Parse(os.Args[1:lastArgIndex])
		}
	}
	var account int
	var token, host string
	if config, err := LoadConfig(flags.ConfigPath); err == nil {
		account = config.Account
		token = config.Token
		host = config.Host
	}
	if flags.Account == 0 {
		flags.Account = account
	}
	if flags.Token == "" {
		flags.Token = token
	}
	if flags.Host == "" {
		flags.Host = host
	}
	if err != nil {
		return nil, err
	}
	flags.Commands = strings.Split(cmd, " ")

	return &flags, nil
}

// Make it possible to change where fmt.Scanf and fmt.Printf read and write so callers can be
// tested.

// Controls where fmt.(F)Printf should write, defaults to stdout.
var out io.Writer = os.Stdout

// Controls where fmt.(F)Scanf should read, defaults to stdin.
var in io.Reader = os.Stdin

// SetOutput changes where functions print, mainly useful for testing
func SetOutput(o io.Writer) {
	out = o
}

// SetInput changes where prompt functions scans, mainly useful for testing
func SetInput(i io.Reader) {
	in = i
}
