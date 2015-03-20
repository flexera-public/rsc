package cmd

import "net/http"

// Command and top level flags
// API clients register additional sub-commands with their own flags
// This data structure is created by rsc and given to each API client command line tool for
// processing.
type CommandLine struct {
	Command             string // Command to be run (e.g. "api15 index")
	ConfigPath          string // Path to rsc config file, defaults to $HOME/.rsc
	Account             int    // RightScale account, optional
	Host                string // API hostname, optional
	Token               string // Auth token, alternative to Username+Password or RL10
	Username            string // Login username, alternative to Token or RL10
	Password            string // Login pasword, alternative to Token or RL10
	RL10                bool   // Whether to send requests using the RL10 proxy
	NoRedirect          bool   // Whether to disable auto-redirects
	FetchResource       bool   // Whether to fetch resource returned in 'Location' header
	ExtractOneSelect    string // JSON select expression to extract single value from response, optional
	ExtractSelector     string // JSON select expression to extract zero or more values from response, optional
	ExtractSelectorJson string // JSON select expression to extract zero or more values from response, extracted values are displayed using JSON encoding, optional
	ExtractHeader       string // Name of header to extract from response, optional
	Dump                string // Whether to dump raw HTTP request and response to stdout (values are empty string - don't dump, "debug" or "json")
	Pretty              bool   // Whether to display response body or extract values using pretty printer
	ShowHelp            bool   // Whether to show help for action flags
}

// Common interface between rsc package and API client packages.
// Provide methods to show help about and run commands.
type CommandClient interface {
	ShowCommandHelp(cmdLine string) error              // Show contextual help
	ShowApiActions(cmdLine string) error               // Print API or resource actions
	RunCommand(cmdLine string) (*http.Response, error) // Run command
}
