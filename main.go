package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/rightscale/rsc/rsapi15"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Command line client entry point.
func main() {
	app := kingpin.New("rsc", "A RightScale API client")
	app.Version("0.1.0")

	// Register global commands and flags
	var cfgPathFlag = app.Flag("config", "path to rsc config file").Short('c').Default(path.Join(os.Getenv("HOME"), ".rsc")).String()
	var accountFlag = app.Flag("account", "RightScale account ID").Short('a').Int()
	var hostFlag = app.Flag("host", "RightScale API host (e.g. 'us-3.rightscale.com')").Short('e').String()
	var tokenFlag = app.Flag("token", "OAuth access token").Short('t').String()
	var rl10Flag = app.Flag("rl10", "Proxy requests through RightLink 10 (exclusive with '-token')").Bool()

	var extractOneFlag = app.Flag("x1", "Extract single value using given JSON:select expression").String()
	var extractMultipleFlag = app.Flag("xm", "Extract zero, one or multiple values using given JSON:select expression and return space separated list (useful for bash scripts)").String()
	var extractMultipleFlagJson = app.Flag("xj", "Extract zero, one or multiple values using given JSON:select expression and return JSON").String()
	var extractLinkFlag = app.Flag("xl", "Extract href of link found using given JSON:select expression").String()
	var extractHeaderFlag = app.Flag("xh", "Extract header with given name").String()
	var noRedirectFlag = app.Flag("noRedirect", "Do not follow redirect responses").Bool()
	var fetchFlag = app.Flag("fetch", "Fetch resource with href present in 'Location' header").Bool()
	var dumpRequestResponseFlag = app.Flag("dump", "Dump HTTP request and response for debugging").Bool()
	var prettyFlag = app.Flag("pp", "Pretty print response body").Bool()

	app.Command("setup",
		"create config file, defaults to $HOME/.rsc, use '--config' to override")

	// Register API15 subcommands and flags
	rsapi15.RegisterCommands(app)

	var cmd, err = app.Parse(os.Args[1:])
	var showHelp = false
	if err != nil {
		var help = os.Args[len(os.Args)-1]
		var lastArgIndex = len(os.Args)
		if help == "--help" || help == "-h" || help == "-help" || help == "-?" {
			showHelp = true
			lastArgIndex -= 1
			cmd, err = app.Parse(os.Args[1:lastArgIndex])
		}
	}
	if err != nil {
		PrintFatal(err.Error())
	}

	// Initialize global flag values
	var account int
	var token, host string
	if config, err := LoadConfig(*cfgPathFlag); err == nil {
		account = config.Account
		token = config.Token
		host = config.Host
	}
	if *accountFlag > 0 {
		account = *accountFlag
	}
	if *tokenFlag != "" {
		token = *tokenFlag
	}
	if *hostFlag != "" {
		host = *hostFlag
	}

	// Execute appropriate command
	var resp *http.Response
	switch strings.Split(cmd, " ")[0] {

	case "setup":
		err = createConfig(*cfgPathFlag)
		if err != nil {
			PrintFatal(err.Error())
		}

	case "api15":
		var client *rsapi15.Api15
		var httpClient *http.Client
		if *noRedirectFlag {
			httpClient = &http.Client{
				CheckRedirect: func(*http.Request, []*http.Request) error {
					return fmt.Errorf("Client configured to prevent redirection")
				},
			}
		} else {
			httpClient = http.DefaultClient
		}
		if *rl10Flag {
			client, err = rsapi15.NewRL10(nil, httpClient)
		} else {
			client, err = rsapi15.New(account, token, host, nil, httpClient)
		}
		if err != nil {
			PrintFatal("Failed to create API session: %v", err.Error())
		}
		if showHelp {
			var err = client.ShowHelp(cmd)
			if err != nil {
				PrintFatal(err.Error())
			}
			return
		} else {
			if token == "" && !*rl10Flag {
				PrintFatal("Missing OAuth token, use '-token TOKEN' or 'setup'")
			}
			client.DumpRequestResponse = *dumpRequestResponseFlag
			client.FetchLocationResource = *fetchFlag
			resp, err = client.RunCommand(cmd)
		}

	default:
		PrintFatal("Unknown command '%s'", cmd)
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
	if *extractOneFlag != "" {
		err = displayer.ApplySingleExtract(*extractOneFlag)
	} else if *extractMultipleFlag != "" {
		err = displayer.ApplyExtract(*extractMultipleFlag, false)
	} else if *extractMultipleFlagJson != "" {
		err = displayer.ApplyExtract(*extractMultipleFlagJson, true)
	} else if *extractLinkFlag != "" {
		err = displayer.ApplyLinkExtract(*extractLinkFlag)
	} else if *extractHeaderFlag != "" {
		err = displayer.ApplyHeaderExtract(*extractHeaderFlag)
	}
	if err != nil {
		PrintFatal(err.Error())
	}
	if *prettyFlag {
		displayer.Pretty()
	}

	// We're done
	fmt.Printf(displayer.Output())
}
