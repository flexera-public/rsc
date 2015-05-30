package rsapi

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
)

// RightScale client
// Instances of this struct should be created through `New`, `NewRL10` or `FromCommandLine`.
type Api struct {
	Auth                  Authenticator // Authenticator, signs requests for auth
	Logger                *log.Logger   // Optional logger, if specified requests and responses get logged
	Host                  string        // API host, e.g. "us-3.rightscale.com"
	Client                HttpClient    // Underlying http client (not used for authentication requests as these necessitate special redirect handling)
	DumpRequestResponse   Format        // Whether to dump HTTP requests and responses to STDOUT, and if so in which format
	FetchLocationResource bool          // Whether to fetch resource pointed by Location header
	Metadata              ApiMetadata   // Generated API metadata

	insecure bool // Whether HTTP should be used instead of HTTPS (used by RL10 proxied requests)
	// Use Insecure method to set to true.
}

// Request/response dump format
type Format int

const (
	NoDump  Format = 1 << iota // No dump
	Debug                      // Dump in human readable format, exclusive with JSON
	JSON                       // Dump in JSON format, exclusive with Debug
	Verbose                    // Dump auth requests and headers as well
)

// IsDebug is a convenience wrapper that returns true if the Debug bit is set on the flag.
func (f Format) IsDebug() bool {
	return f&Debug != 0
}

// IsJSON is a convenience wrapper that returns true if the JSON bit is set on the flag.
func (f Format) IsJSON() bool {
	return f&JSON != 0
}

// IsVerbose is a convenience wrapper that returns true if the Verbose bit is set on the flag.
func (f Format) IsVerbose() bool {
	return f&Verbose != 0
}

// Api metadata consists of resource metadata indexed by resource name
type ApiMetadata map[string]*metadata.Resource

// Generic API parameter type, used to specify optional parameters for example
type ApiParams map[string]interface{}

// Use interface instead of raw http.Client to ease testing
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New returns a API client that uses the given authenticator.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(host string, auth Authenticator, logger *log.Logger, client HttpClient) *Api {
	if client == nil {
		client = http.DefaultClient
	}
	a := &Api{
		Auth:   auth,
		Logger: logger,
		Host:   host,
		Client: client,
	}
	if auth != nil {
		auth.SetHost(a.FullHost())
	}
	return a
}

// NewRL10 returns a API client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the client returned by New in
// all other regards.
func NewRL10(logger *log.Logger, client HttpClient) (*Api, error) {
	rllConfig, err := os.Open(RllSecret)
	if err != nil {
		return nil, fmt.Errorf("Failed to load RLL config: %s", err)
	}
	defer rllConfig.Close()
	var port string
	var secret string
	scanner := bufio.NewScanner(rllConfig)
	for scanner.Scan() {
		line := scanner.Text()
		elems := strings.Split(line, "=")
		if len(elems) != 2 {
			return nil, fmt.Errorf("Invalid RLL configuration line '%s'", line)
		}
		switch elems[0] {
		case "RS_RLL_PORT":
			port = elems[1]
			if _, err := strconv.Atoi(elems[1]); err != nil {
				return nil, fmt.Errorf("Invalid port value '%s'", port)
			}
		case "RS_RLL_SECRET":
			secret = elems[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed to load RLL config: %s", err)
	}
	host := "localhost:" + port
	auth := NewRL10Authenticator(secret)
	auth.SetHost(host)
	api := &Api{
		Auth:   auth,
		Logger: logger,
		Host:   host,
		Client: client,
	}
	api.Insecure()
	return api, nil
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	var client *Api
	var httpClient *http.Client
	if cmdLine.NoRedirect {
		httpClient = &http.Client{
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return fmt.Errorf("Client configured to prevent redirection")
			},
		}
	} else {
		httpClient = http.DefaultClient
	}
	if cmdLine.RL10 {
		var err error
		if client, err = NewRL10(nil, httpClient); err != nil {
			return nil, err
		}
	} else if cmdLine.OAuthToken != "" {
		auth := NewOAuthAuthenticator(cmdLine.OAuthToken)
		client = New(cmdLine.Host, auth, nil, httpClient)
	} else if cmdLine.OAuthAccessToken != "" {
		auth := NewTokenAuthenticator(cmdLine.OAuthAccessToken)
		client = New(cmdLine.Host, auth, nil, httpClient)
	} else if cmdLine.APIToken != "" {
		auth := NewInstanceAuthenticator(cmdLine.APIToken, cmdLine.Account)
		client = New(cmdLine.Host, auth, nil, httpClient)
	} else if cmdLine.Username != "" && cmdLine.Password != "" {
		auth := NewBasicAuthenticator(cmdLine.Username, cmdLine.Password, cmdLine.Account)
		client = New(cmdLine.Host, auth, nil, httpClient)
	} else {
		// No auth, used by tests
		client = New(cmdLine.Host, nil, nil, httpClient)
		client.Insecure()
	}
	if !cmdLine.ShowHelp && !cmdLine.NoAuth {
		if cmdLine.OAuthToken == "" && cmdLine.OAuthAccessToken == "" && cmdLine.APIToken == "" && cmdLine.Username == "" && !cmdLine.RL10 {
			return nil, fmt.Errorf("Missing authentication information, use '--email EMAIL --password PWD', '--token TOKEN' or 'setup'")
		}
		format := NoDump
		if cmdLine.Verbose || cmdLine.Dump == "debug" {
			format = Debug
		}
		if cmdLine.Dump == "json" {
			format = JSON
		}
		if cmdLine.Verbose {
			format |= Verbose
		}
		client.EnableDump(format)
		client.FetchLocationResource = cmdLine.FetchResource
	}
	return client, nil
}

// EnableDump sets the dump format for all API requests made by the client.
func (a *Api) EnableDump(format Format) {
	a.DumpRequestResponse = format
	a.Auth.EnableDump(format)
}

// CanAuthenticate() makes a test authenticated request to the RightScale API and returns an error
// if it fails.
func (a *Api) CanAuthenticate() error {
	res := a.Auth.CanAuthenticate(a.FullHost())
	return res
}

// Force the use of HTTP instead of HTTPS for all requests.
func (a *Api) Insecure() {
	a.insecure = true
	if a.Auth != nil {
		// Reset host used for authentication so it's prefixed with "http://"
		a.Auth.SetHost(a.FullHost())
	}
}

// FullHost returns the scheme prefixed hostname and can be used to instantiate http.Client
func (a *Api) FullHost() string {
	scheme := "https://"
	if a.insecure {
		scheme = "http://"
	}
	return scheme + a.Host
}
