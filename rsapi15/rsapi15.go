//go:generate api15gen ./api_data.json
package rsapi15

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

// RightScale API 1.5 client
// Exposes a couple of methods for each API call (one using a discrete list of parameters, the other
// using the generic Params structure, see Params).
// Instances of this struct should be created through `NewClient` which takes care of the initial
// authentication.
// The client can be attached to a logger using the `SetLogger` method.
type Client struct {
	accountId    int
	refreshToken string
	accessToken  string
	refreshAt    time.Time
	logger       *log.Logger
	endpoint     string
	client       http.Client
}

// APIs can be called in two ways: using a discrete interface where all parameters have to be
// provided (and are typed properly) or a generic interface using a map of interface{}. The later
// is useful to call APIs that support a lot of parameters but where only a few are needed for a
// given use case (e.g. session creation). For this case we provide the convenient Params struct
// to define the parameter values inline, e.g.
//    CreateSession(&Params{account_id: 71, token: "foo"})
type Params map[string]interface{}

// NewClient returns a API 1.5 client from a refresh token and an account id.
// It attempts to authenticate with RightScale and returns an error if that fails.
// If no endpoint is provided (i.e. `endpoint` argument is the empty string) then the correct
// endpoint is detected during initial auth (in other words, provide the endpoint to avoid the
// initial redirect and save time).
// The timeout semantic is the same as the `net/http` package `Client` type `timeout` field. In
// particular a timeout value of zero means no timeout.
func NewClient(token string, accountId int, endpoint string, timeout time.Duration) (*Client, error) {
	if endpoint == "" {
		endpoint = "https://my.rightscale.com"
	}
	client := http.Client{Timeout: timeout}
	c := &Client{
		accountId:    accountId,
		refreshToken: token,
		endpoint:     endpoint,
		client:       client,
	}
	client.CheckRedirect = c.handleRedirect
	if err := c.refresh(); err != nil {
		return nil, err
	} else {
		return c, nil
	}
}

// SetAccount sets the active account for all API calls.
func (c *Client) SetAccount(a int) {
	c.accountId = a
}

// Assign a logger to the client, the client logs requests and responses.
func (c *Client) SetLogger(l *log.Logger) {
	c.logger = l
}

// Update endpoint to match redirect
func (c *Client) handleRedirect(req *http.Request, via []*http.Request) error {
	elems := strings.SplitN(req.URL.String(), "/api", 2)
	c.endpoint = elems[0]
	return nil
}

// Refresh access token, handles redirection
func (c *Client) refresh() error {
	session, err := c.CreateOauth2(c.accountId, "", "", "refresh_token", c.refreshToken, "", 0)
	if err != nil {
		return err
	}
	c.accessToken = session["access_token"].(string)
	d, err := time.ParseDuration(session["expires_in"].(string) + "s")
	if err != nil {
		return err
	}
	c.refreshAt = time.Now().Add(d / 2)

	return nil
}

// Request context created by `beforeRequest`
// Used by `afterRequest` to log the time it took to process the request
type requestContext struct {
	id        string
	startedAt time.Time
}

// Before request callback, check token freshness, sign request and log it if needed
// Return short unique id used for logging, pass it back in afterRequest to log it together
// with response.
func (c *Client) beforeRequest(r *http.Request) *requestContext {
	r.Header.Add("X-Account", strconv.Itoa(c.accountId))
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	r.URL.Path = path.Join(c.endpoint, r.URL.Path)
	if c.refreshAt.After(time.Now()) {
		err := c.refresh()
		if err != nil && c.logger != nil {
			c.logger.Printf("ERROR: failed to refresh token: %s, will retry with next API call.", err.Error())
		}
	}
	b := make([]byte, 6)
	io.ReadFull(rand.Reader, b)
	id := base64.StdEncoding.EncodeToString(b)
	if c.logger != nil {
		c.logger.Printf("[%s] %s %s", id, r.Method, r.URL.String())
	}
	return &requestContext{id: id, startedAt: time.Now()}
}

// After request callback, log if needed
// TBD: add some verbose flag that enables logging request and response bodies?
func (c *Client) afterRequest(r *http.Response, ctx *requestContext) {
	d := time.Since(ctx.startedAt)
	if c.logger != nil {
		c.logger.Printf("[%s] %s in %s", ctx.id, r.Status, d.String())
	}
}
