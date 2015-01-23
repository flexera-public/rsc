//go:generate api15gen ./api_data.json
package rsapi15

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
func NewClient(token string, accountId int) (*Client, error) {
	c := &Client{accountId: accountId, refreshToken: token}
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

// Refresh access token
func (c *Client) refresh() error {
	session, err := c.CreateOauth2(&Params{
		"refresh_token": c.refreshToken,
		"account_id":    c.accountId,
		"grant_type":    "refresh_token",
	})
	if err != nil {
		return err
	}
	s := session.(map[string]string)
	c.accessToken = s["access_token"]
	d, err := time.ParseDuration(s["expires_in"] + "s")
	if err != nil {
		return err
	}
	c.refreshAt = time.Now().Add(d / 2)

	return nil
}

// Before request callback, check token freshness, sign request and log it if needed
func (c *Client) beforeRequest(r *http.Request, id string) {
	r.Header.Add("X-Account", strconv.Itoa(c.accountId))
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	if c.refreshAt.After(time.Now()) {
		err := c.refresh()
		if err != nil && c.logger != nil {
			c.logger.Printf("ERROR: failed to refresh token: %s, will retry with next API call.", err.Error())
		}
	}
	if c.logger != nil {
		c.logger.Printf("[%s] %s %s", id, r.Method, r.URL.String())
	}
}

// After request callback, log if needed
// TBD: add some verbose flag that enables logging request and response bodies?
func (c *Client) afterRequest(r *http.Response, id string) {
	if c.logger != nil {
		c.logger.Printf("[%s] %s", id, r.Status)
	}
}
