package rsapi

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/recording"
)

// Log request, dump its content if required then make request and log response and dump it.
func (a *Api) PerformRequest(req *http.Request) (*http.Response, error) {
	if a.insecure {
		req.URL.Scheme = "http"
	} else {
		req.URL.Scheme = "https"
	}
	var id string
	var startedAt time.Time
	if a.Logger != nil {
		startedAt = time.Now()
		b := make([]byte, 6)
		io.ReadFull(rand.Reader, b)
		id = base64.StdEncoding.EncodeToString(b)
		a.Logger.Printf("[%s] %s %s", id, req.Method, req.URL.String())
	}
	req.Header.Set("User-Agent", UA)
	var reqBody []byte
	if !a.DumpRequestResponse.IsVerbose() {
		reqBody = dumpRequest(a.DumpRequestResponse, req)
	}
	// Sign last so auth headers don't get printed or logged
	if a.Auth != nil {
		if err := a.Auth.Sign(req); err != nil {
			return nil, err
		}
	}
	if a.DumpRequestResponse.IsVerbose() {
		reqBody = dumpRequest(a.DumpRequestResponse, req)
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	dumpResponse(a.DumpRequestResponse, resp, req, reqBody)
	if a.Logger != nil {
		d := time.Since(startedAt)
		a.Logger.Printf("[%s] %s in %s", id, resp.Status, d.String())
	}

	return resp, err
}

// IdentifyParams organizes the given params in two groups: the payload params and the query params.
func IdentifyParams(a *metadata.Action, params ApiParams) (payloadParams ApiParams, queryParams ApiParams) {
	payloadParamNames := a.PayloadParamNames()
	payloadParams = make(ApiParams)
	for _, n := range payloadParamNames {
		if p, ok := params[n]; ok {
			payloadParams[n] = p
		}
	}
	queryParamNames := a.QueryParamNames()
	queryParams = make(ApiParams)
	for _, n := range queryParamNames {
		if p, ok := params[n]; ok {
			queryParams[n] = p
		}
	}
	return payloadParams, queryParams
}

// Deserialize JSON response into generic object.
// If the response has a "Location" header then the returned object is a map with one key "Location"
// containing the value of the header.
func (a *Api) LoadResponse(resp *http.Response) (interface{}, error) {
	defer resp.Body.Close()
	var respBody interface{}
	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err)
	}
	if len(jsonResp) > 0 {
		err = json.Unmarshal(jsonResp, &respBody)
		if err != nil {
			return nil, fmt.Errorf("Failed to load response (%s)", err)
		}
	}
	// Special case for "Location" header, assume that if there is a location there is no body
	loc := resp.Header.Get("Location")
	if len(loc) > 0 {
		var bodyMap = make(map[string]interface{})
		bodyMap["Location"] = loc
		respBody = interface{}(bodyMap)
	}
	return respBody, err
}

// Dump request if needed.
// Return request serialized as JSON if dump format is JSON, nil otherwise.
func dumpRequest(format Format, req *http.Request) []byte {
	if format == NoDump {
		return nil
	}
	if format.IsDebug() {
		b, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			bod := strings.TrimSpace(string(b))
			fmt.Fprintf(os.Stderr, "%s\n", bod)
		} else {
			fmt.Fprintf(os.Stderr, "** Failed to dump request content - %s\n", err)
		}
	} else if format.IsJSON() {
		reqBody, err := dumpReqBody(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "** Failed to load request body for dump: %s\n", err)
		}
		return reqBody
	}
	return nil
}

// dumpResponse dumps the response and optionally the request (in case of JSON format) according to
// the format.
// It also checks whether the special recorder pipe is opened and if so writes the dump to it.
func dumpResponse(format Format, resp *http.Response, req *http.Request, reqBody []byte) {
	if format == NoDump {
		return
	}
	if format.IsDebug() {
		b, err := httputil.DumpResponse(resp, false)
		if err == nil {
			fmt.Fprintf(os.Stderr, "==> %s\n", b)
		} else {
			fmt.Fprintf(os.Stderr, "** Failed to dump response content - %s\n", err)
		}
	} else if format.IsJSON() {
		respBody, err := dumpRespBody(resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "** Failed to load response body for dump: %s\n", err)
		}
		dumped := recording.RequestResponse{
			Verb:       req.Method,
			Uri:        req.URL.String(),
			ReqHeader:  req.Header,
			ReqBody:    string(reqBody),
			Status:     resp.StatusCode,
			RespHeader: resp.Header,
			RespBody:   string(respBody),
		}
		b, err := json.MarshalIndent(dumped, "", "    ")
		if err == nil {
			fd := os.Stderr
			f := os.NewFile(10, "fd10")
			_, err := f.Stat()
			if err == nil {
				// fd 10 is open, dump to it (used by recorder)
				fd = f
			}
			fmt.Fprintf(fd, "%s\n", string(b))
		} else {
			fmt.Fprintf(os.Stderr, "** Failed to dump request content - %s\n", err)
		}
	}
}

// Dump request body, strongly inspired from httputil.DumpRequest
func dumpReqBody(req *http.Request) ([]byte, error) {
	if req.Body == nil {
		return nil, nil
	}
	var save io.ReadCloser
	var err error
	save, req.Body, err = drainBody(req.Body)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	var dest io.Writer = &b
	chunked := len(req.TransferEncoding) > 0 && req.TransferEncoding[0] == "chunked"
	if chunked {
		dest = httputil.NewChunkedWriter(dest)
	}
	_, err = io.Copy(dest, req.Body)
	if chunked {
		dest.(io.Closer).Close()
		io.WriteString(&b, "\r\n")
	}
	req.Body = save
	return b.Bytes(), err
}

// Dump response body, strongly inspired from httputil.DumpResponse
func dumpRespBody(resp *http.Response) ([]byte, error) {
	if resp.Body == nil {
		return nil, nil
	}
	var b bytes.Buffer
	savecl := resp.ContentLength
	var save io.ReadCloser
	var err error
	save, resp.Body, err = drainBody(resp.Body)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&b, resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = save
	resp.ContentLength = savecl
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// One of the copies, say from b to r2, could be avoided by using a more
// elaborate trick where the other copy is made during Request/Response.Write.
// This would complicate things too much, given that these functions are for
// debugging only.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, nil, err
	}
	if err = b.Close(); err != nil {
		return nil, nil, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
