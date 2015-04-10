package rsapi  // import "gopkg.in/rightscale/rsc.v1-unstable/rsapi"

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/rightscale/rsc.v1-unstable/cmd"      // import "gopkg.in/rightscale/rsc.v1-unstable/cmd"
	"gopkg.in/rightscale/rsc.v1-unstable/metadata" // import "gopkg.in/rightscale/rsc.v1-unstable/metadata"
)

// RightScale client
// Instances of this struct should be created through `New`, `NewRL10` or `FromCommandLine`.
type Api struct {
	AccountId             int           // Account in which client is currently operating
	Auth                  Authenticator // Authenticator, signs requests for auth
	Logger                *log.Logger   // Optional logger, if specified requests and responses get logged
	Host                  string        // API host, e.g. "us-3.rightscale.com"
	Client                HttpClient    // Underlying http client
	Unsecure              bool          // Whether HTTP should be used instead of HTTPS (used by RL10 proxied requests)
	DumpRequestResponse   Format        // Whether to dump HTTP requests and responses to STDOUT, and if so in which format
	FetchLocationResource bool          // Whether to fetch resource pointed by Location header
	Metadata              ApiMetadata   // Generated API metadata
}

// Request/response dump format
type Format int

const (
	NoDump Format = iota
	Debug
	Json
)

// Api metadata consists of resource metadata indexed by resource name
type ApiMetadata map[string]*metadata.Resource

// Generic API parameter type, used to specify optional parameters for example
type ApiParams map[string]interface{}

// Use interface instead of raw http.Client to ease testing
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New returns a API client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(accountId int, host string, auth Authenticator, logger *log.Logger, client HttpClient) (*Api, error) {
	if client == nil {
		client = http.DefaultClient
	}
	if auth != nil {
		var err error
		host, err = auth.ResolveHost(host, accountId)
		if err != nil {
			return nil, err
		}
	}
	return &Api{
		AccountId: accountId,
		Auth:      auth,
		Logger:    logger,
		Host:      host,
		Client:    client,
	}, nil
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
	auth := RL10Authenticator{secret}
	host := "localhost:" + port
	return &Api{
		Auth:     &auth,
		Logger:   logger,
		Host:     host,
		Client:   client,
		Unsecure: true,
	}, nil
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
	var err error
	if cmdLine.RL10 {
		client, err = NewRL10(nil, httpClient)
	} else if cmdLine.Token != "" {
		auth := OAuthAuthenticator{
			RefreshToken: cmdLine.Token,
			AccessToken:  "",
			RefreshAt:    time.Now().Add(-time.Duration(5) * time.Minute),
			Client:       httpClient,
		}
		client, err = New(cmdLine.Account, cmdLine.Host, &auth, nil, httpClient)
	} else if cmdLine.Username != "" && cmdLine.Password != "" {
		auth := LoginAuthenticator{
			Username:  cmdLine.Username,
			Password:  cmdLine.Password,
			RefreshAt: time.Now().Add(-time.Duration(5) * time.Minute),
			Client:    httpClient,
		}
		client, err = New(cmdLine.Account, cmdLine.Host, &auth, nil, httpClient)
	} else {
		// No auth, used by tests
		client, err = New(cmdLine.Account, cmdLine.Host, nil, nil, httpClient)
		client.Unsecure = true
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to create API session: %v", err)
	}
	if !cmdLine.ShowHelp && !cmdLine.NoAuth {
		if cmdLine.Token == "" && cmdLine.Username == "" && !cmdLine.RL10 {
			return nil, fmt.Errorf("Missing authentication information, use '--email EMAIL --password PWD', '--token TOKEN' or 'setup'")
		}
		client.DumpRequestResponse = NoDump
		if cmdLine.Dump == "json" {
			client.DumpRequestResponse = Json
		} else if cmdLine.Dump == "debug" {
			client.DumpRequestResponse = Debug
		}
		client.FetchLocationResource = cmdLine.FetchResource
	}
	return client, nil
}
