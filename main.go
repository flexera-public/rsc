package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi15"
	"github.com/rightscale/rsc/rsapi16"
	"github.com/rightscale/rsc/ss"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Command line client entry point.
func main() {
	app := kingpin.New("rsc", "A RightScale API client")
	app.Version("0.3.0")

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
	var client cmd.CommandClient
	switch topCommand {
	case "setup":
		err = CreateConfig(cmdLine.ConfigPath)
	case "api15":
		client, err = rsapi15.FromCommandLine(cmdLine)
		if err == nil {
			resp, err = runCommand(client, cmdLine)
		}
	case "api16":
		client, err = rsapi16.FromCommandLine(cmdLine)
		if err == nil {
			resp, err = runCommand(client, cmdLine)
		}
	case "ss":
		client, err = ss.FromCommandLine(cmdLine)
		if err == nil {
			resp, err = runCommand(client, cmdLine)
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
		if err != nil {
			moreThanOneError = strings.Contains(err.Error(), "returned more than one value") // Ugh, there has to be a better way
		}
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

// Helper that runs command line with give command client
func runCommand(client cmd.CommandClient, cmdLine *cmd.CommandLine) (resp *http.Response, err error) {
	if cmdLine.ShowHelp {
		err = client.ShowCommandHelp(cmdLine.Command)
	} else {
		resp, err = client.RunCommand(cmdLine.Command)
	}
	return
}
