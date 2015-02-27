package rsapi

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

// Log request, dump its content if required then make request and log response and dump it.
func (a *Api) PerformRequest(req *http.Request) (*http.Response, error) {
	var id string
	var startedAt time.Time
	if a.Logger != nil {
		startedAt = time.Now()
		b := make([]byte, 6)
		io.ReadFull(rand.Reader, b)
		id = base64.StdEncoding.EncodeToString(b)
		a.Logger.Printf("[%s] %s %s", id, req.Method, req.URL.String())
	}
	if a.DumpRequestResponse {
		var b, err = httputil.DumpRequest(req, true)
		if err == nil {
			fmt.Printf("REQUEST\n-------\n%s\n\n", b)
		}
	}
	// Sign last so auth headers don't get printed or logged
	if err := a.Auth.Sign(req, a.Host); err != nil {
		return nil, err
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if a.Logger != nil {
		d := time.Since(startedAt)
		a.Logger.Printf("[%s] %s in %s", id, resp.Status, d.String())
	}
	if a.DumpRequestResponse {
		var b, err = httputil.DumpResponse(resp, false)
		if err == nil {
			fmt.Printf("RESPONSE\n--------\n%s", b)
		}
	}

	return resp, err
}
