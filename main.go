package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"gopkg.in/alecthomas/kingpin.v1"

	// phoney reference to make Godep pull this in for the code generators
	_ "bitbucket.org/pkg/inflect"
)

// Command line client entry point.
func main() {
	app := kingpin.New("rsc", "A RightScale API client")
	app.Version(VV)

	cmdLine, err := ParseCommandLine(app)
	if err != nil {
		PrintFatal(err.Error())
	}

	// Execute appropriate command
	var resp *http.Response
	topCommand := strings.Split(cmdLine.Command, " ")[0]
	if topCommand == "setup" {
		err = CreateConfig(cmdLine.ConfigPath)
	} else {
		var client cmd.CommandClient
		client, err = ApiClient(topCommand, cmdLine)
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
	moreThanOneError := false
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
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Let user know if something went wrong
		fmt.Fprintf(os.Stderr, "%d: %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	fmt.Fprintf(out, displayer.Output())
	exitStatus := 0
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
	osExit(exitStatus)
}

// Helper that runs command line with give command client
func runCommand(client cmd.CommandClient, cmdLine *cmd.CommandLine) (resp *http.Response, err error) {
	cmds := strings.Split(cmdLine.Command, " ")
	if cmdLine.ShowHelp {
		err = client.ShowCommandHelp(cmdLine.Command)
	} else if len(cmds) > 1 && cmds[1] == "actions" {
		err = client.ShowApiActions(cmdLine.Command)
	} else {
		resp, err = client.RunCommand(cmdLine.Command)
	}
	return
}
