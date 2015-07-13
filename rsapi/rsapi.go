package rsapi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/httpclient"
	"github.com/rightscale/rsc/metadata"
)

type (
	// RightScale client
	// Instances of this struct should be created through `New`, `NewRL10` or `FromCommandLine`.
	Api struct {
		Auth                  Authenticator         // Authenticator, signs requests for auth
		Host                  string                // API host, e.g. "us-3.rightscale.com"
		Client                httpclient.HTTPClient // Underlying http client (not used for authentication requests as these necessitate special redirect handling)
		FetchLocationResource bool                  // Whether to fetch resource pointed by Location header
		Metadata              ApiMetadata           // Generated API metadata

		insecure bool // Whether HTTP should be used instead of HTTPS (used by RL10 proxied requests)
		// Use Insecure method to set to true.
	}
)

// Api metadata consists of resource metadata indexed by resource name
type ApiMetadata map[string]*metadata.Resource

// Generic API parameter type, used to specify optional parameters for example
type ApiParams map[string]interface{}

// New returns a API client that uses the given authenticator.
// host may be blank in which case client attempts to resolve it using auth.
func New(host string, auth Authenticator) *Api {
	client := httpclient.New()
	if strings.HasPrefix(host, "http://") {
		host = host[7:]
	} else if strings.HasPrefix(host, "https://") {
		host = host[8:]
	}
	a := &Api{
		Auth:   auth,
		Host:   host,
		Client: client,
	}
	if auth != nil {
		auth.SetHost(host)
	}
	return a
}

// NewRL10 returns a API client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the client returned by New in
// all other regards.
func NewRL10() (*Api, error) {
	client := httpclient.New()
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
		Host:   host,
		Client: client,
	}
	httpclient.Insecure = true
	return api, nil
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	var client *Api
	ss := strings.HasPrefix(cmdLine.Command, "ss")
	if cmdLine.RL10 {
		var err error
		if client, err = NewRL10(); err != nil {
			return nil, err
		}
	} else if cmdLine.OAuthToken != "" {
		auth := NewOAuthAuthenticator(cmdLine.OAuthToken)
		if ss {
			auth = NewSSAuthenticator(auth, cmdLine.Account)
		}
		client = New(cmdLine.Host, auth)
	} else if cmdLine.OAuthAccessToken != "" {
		auth := NewTokenAuthenticator(cmdLine.OAuthAccessToken)
		if ss {
			auth = NewSSAuthenticator(auth, cmdLine.Account)
		}
		client = New(cmdLine.Host, auth)
	} else if cmdLine.APIToken != "" {
		auth := NewInstanceAuthenticator(cmdLine.APIToken, cmdLine.Account)
		if ss {
			auth = NewSSAuthenticator(auth, cmdLine.Account)
		}
		client = New(cmdLine.Host, auth)
	} else if cmdLine.Username != "" && cmdLine.Password != "" {
		auth := NewBasicAuthenticator(cmdLine.Username, cmdLine.Password, cmdLine.Account)
		if ss {
			auth = NewSSAuthenticator(auth, cmdLine.Account)
		}
		client = New(cmdLine.Host, auth)
	} else {
		// No auth, used by tests
		client = New(cmdLine.Host, nil)
		httpclient.Insecure = true
	}
	if !cmdLine.ShowHelp && !cmdLine.NoAuth {
		if cmdLine.OAuthToken == "" && cmdLine.OAuthAccessToken == "" && cmdLine.APIToken == "" && cmdLine.Username == "" && !cmdLine.RL10 {
			return nil, fmt.Errorf("Missing authentication information, use '--email EMAIL --password PWD', '--token TOKEN' or 'setup'")
		}
		if cmdLine.Verbose || cmdLine.Dump == "debug" {
			httpclient.DumpFormat = httpclient.Debug
		}
		if cmdLine.Dump == "json" {
			httpclient.DumpFormat = httpclient.JSON
		}
		if cmdLine.Dump == "record" {
			httpclient.DumpFormat = httpclient.JSON | httpclient.Record
		}
		if cmdLine.Verbose {
			httpclient.DumpFormat |= httpclient.Verbose
		}
		client.FetchLocationResource = cmdLine.FetchResource
	}
	return client, nil
}

// CanAuthenticate() makes a test authenticated request to the RightScale API and returns an error
// if it fails.
func (a *Api) CanAuthenticate() error {
	res := a.Auth.CanAuthenticate(a.Host)
	return res
}
