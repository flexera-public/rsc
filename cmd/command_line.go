package cmd

// Command and top level flags
// API clients register additional sub-commands with their own flags
// This data structure is created by rsc and given to each API client command line tool for
// processing.
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
	ShowHrefs           bool   // Whether to print all known API href patterns
}
