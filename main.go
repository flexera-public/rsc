package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rightscale/rsc/rsapi15"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Command line client entry point.
func main() {
	app := kingpin.New("rsc", "A RightScale API client")
	app.Version("0.1.0")

	var cmdLine, err = ParseCommandLine(app)
	if err != nil {
		PrintFatal(err.Error())
	}

	// Execute appropriate command
	var resp *http.Response
	var topCommand = strings.Split(cmdLine.Command, " ")[0]
	if !IsClientCommand(topCommand) && topCommand != "setup" {
		topCommand = DefaultClientCommand
	}
	switch topCommand {
	case "setup":
		err = CreateConfig(cmdLine.ConfigPath)

	case "api15":
		var client, err2 = rsapi15.FromCommandLine(cmdLine)
		if err2 != nil {
			err = err2
		} else if client != nil {
			resp, err = client.RunCommand(cmdLine.Command)
		}
	}

	// Handle command results
	if err != nil {
		PrintFatal(err.Error())
	}
	if resp == nil {
		return // No results, just exit (e.g. printed help)
	}

	// Handle command output (apply any extraction)
	displayer, err := NewDisplayer(resp)
	if err != nil {
		PrintFatal(err.Error())
	}
	var moreThanOneError = false
	if cmdLine.ExtractOneSelect != "" {
		err = displayer.ApplySingleExtract(cmdLine.ExtractOneSelect)
		moreThanOneError = strings.Contains(err.Error(), "returned more than one value") // Ugh, there has to be a better way
	} else if cmdLine.ExtractSelector != "" {
		err = displayer.ApplyExtract(cmdLine.ExtractSelector, false)
	} else if cmdLine.ExtractSelectorJson != "" {
		err = displayer.ApplyExtract(cmdLine.ExtractSelectorJson, true)
	} else if cmdLine.ExtractHeader != "" {
		err = displayer.ApplyHeaderExtract(cmdLine.ExtractHeader)
	}
	if err != nil {
		PrintFatal(err.Error())
	}
	if cmdLine.Pretty {
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
