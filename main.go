package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/log"
	"gopkg.in/alecthomas/kingpin.v2"

	// phoney reference to make Godep pull this in for the code generators
	_ "bitbucket.org/pkg/inflect"
)

// Command line client entry point.
func main() {
	app := kingpin.New("rsc", "A RightScale API client")
	app.Writer(os.Stdout)
	app.Version(VV)

	cmdLine, err := ParseCommandLine(app)
	if err != nil {
		line := strings.Join(os.Args, " ")
		PrintFatal(line + ": " + err.Error())
	}

	// Execute appropriate command
	app.Writer(errOut)
	log.Interactive()
	var resp *http.Response
	topCommand := strings.Split(cmdLine.Command, " ")[0]
	switch topCommand {
	case "setup":
		err = CreateConfig(cmdLine.ConfigPath)
	case "json":
		var b []byte
		b, err = ioutil.ReadAll(os.Stdin)
		if err == nil {
			resp = &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
			}
		}
	default:
		var client cmd.CommandClient
		client, err = APIClient(topCommand, cmdLine)
		if err == nil {
			resp, err = runCommand(client, cmdLine)
		}
	}

	// Handle command results
	if err != nil {
		PrintFatal(err.Error())
	}
	if resp == nil {
		return // No results, just exit (e.g. setup, printed help...)
	}

	var notExactlyOneError bool
	displayer, err := NewDisplayer(resp)
	if err != nil {
		PrintFatal(err.Error())
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Let user know if something went wrong
		fmt.Fprintln(errOut, resp.Status)
		if len(displayer.body) > 0 {
			fmt.Fprintln(errOut, displayer.body)
		}
	} else if cmdLine.ExtractOneSelect != "" {
		err = displayer.ApplySingleExtract(cmdLine.ExtractOneSelect)
		if err != nil {
			notExactlyOneError = strings.Contains(err.Error(),
				"instead of one value") // Ugh, there has to be a better way
			PrintError(err.Error())
		}
		fmt.Fprint(out, displayer.Output())
	} else {
		if cmdLine.ExtractSelector != "" {
			err = displayer.ApplyExtract(cmdLine.ExtractSelector, false)
		} else if cmdLine.ExtractSelectorJSON != "" {
			err = displayer.ApplyExtract(cmdLine.ExtractSelectorJSON, true)
		} else if cmdLine.ExtractHeader != "" {
			err = displayer.ApplyHeaderExtract(cmdLine.ExtractHeader)
		}
		if err != nil {
			PrintFatal(err.Error())
		} else if cmdLine.Pretty {
			displayer.Pretty()
		}
		fmt.Fprint(out, displayer.Output())
	}
	// Figure out exit code
	exitStatus := 0
	switch {
	case notExactlyOneError:
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
	//fmt.Fprintf(os.Stderr, "exitStatus=%d\n", exitStatus)
	osExit(exitStatus)
}

// Helper that runs command line with give command client
func runCommand(client cmd.CommandClient, cmdLine *cmd.CommandLine) (resp *http.Response, err error) {
	cmds := strings.Split(cmdLine.Command, " ")
	if cmdLine.ShowHelp {
		err = client.ShowCommandHelp(cmdLine.Command)
	} else if len(cmds) > 1 && cmds[1] == "actions" {
		err = client.ShowAPIActions(cmdLine.Command)
	} else {
		resp, err = client.RunCommand(cmdLine.Command)
	}
	return
}
