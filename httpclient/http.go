package httpclient // import "gopkg.in/rightscale/rsc.v6/httpclient"

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"

	"gopkg.in/rightscale/rsc.v6/log"
	"gopkg.in/rightscale/rsc.v6/recording"
)

const (
	// NoDump is the default value for DumpFormat.
	NoDump Format = 1 << iota
	// Debug formats the dumps in human readable format, the use of this flag is exclusive with
	// JSON.
	Debug
	// JSON formats the dumps in JSON, the use of this flag is exclusive with Debug.
	JSON
	// Verbose enables the dumps for all requests and auth headers.
	Verbose
	// Record causes the dumps to be written to the recorder file descriptor (used by tests).
	Record
)

const (
	noRedirectError = "no redirect"
	requestIdHeader = "X-Request-Id"
)

var (
	// DumpFormat dictates how HTTP requests and responses are logged: NoDump prevents logging
	// altogether, Debug generates logs in human readable format and JSON in JSON format.
	// Verbose causes all headers to be logged - including sensitive ones.
	DumpFormat Format

	// Insecure dictates whether HTTP (true) or HTTPS (false) should be used to connect to the
	// API endpoints.
	Insecure bool

	// NoCertCheck dictates whether the SSL handshakes should bypass X509 certificate
	// validation (true) or not (false).
	NoCertCheck bool

	// ResponseHeaderTimeout if non-zero, specifies the amount of
	// time to wait in seconds for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	ResponseHeaderTimeout = 300 * time.Second

	// HiddenHeaders lists headers that should not be logged unless DumpFormat is Verbose.
	HiddenHeaders = map[string]bool{"Authorization": true, "Cookie": true}
)

// For tests
var (
	OsStderr io.Writer = os.Stderr
)

type (
	// HTTPClient makes it easier to stub HTTP clients for testing.
	HTTPClient interface {
		// Do makes a regular http request and returns the response/error.
		Do(req *http.Request) (*http.Response, error)

		// DoWithContext performs a request and is context-aware.
		DoWithContext(ctx context.Context, req *http.Request) (*http.Response, error)

		// DoHidden prevents logging, useful for requests made during authorization.
		DoHidden(req *http.Request) (*http.Response, error)

		// DoHiddenWithContext prevents logging and performs a context-aware request.
		DoHiddenWithContext(ctx context.Context, req *http.Request) (*http.Response, error)
	}

	// Format is the request/response dump format.
	Format int

	// HTTP client that optionally dumps requests and responses.
	// This client also disables the default http client redirect handling.
	dumpClient struct {
		Client *http.Client
		Format Format
	}
)

// Default DumpFormat to NoDump
func init() {
	DumpFormat = NoDump
}

// New returns an HTTP client using the settings specified by this package variables.
func New() HTTPClient {
	return &dumpClient{Client: newRawClient(false)}
}

// NewNoRedirect returns an HTTP client that does not follow redirects.
func NewNoRedirect() HTTPClient {
	return &dumpClient{Client: newRawClient(true)}
}

// ShortToken creates a 6 bytes unique string.
// Not meant to be cryptographically unique but good enough for logs.
func ShortToken() string {
	b := make([]byte, 6)
	io.ReadFull(rand.Reader, b)
	return base64.StdEncoding.EncodeToString(b)
}

// newRawClient creates an http package Client taking into account both the parameters and package
// variables.
func newRawClient(noredirect bool) *http.Client {
	tr := http.Transport{ResponseHeaderTimeout: ResponseHeaderTimeout, Proxy: http.ProxyFromEnvironment}
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: NoCertCheck}
	c := http.Client{Transport: &tr}
	if noredirect {
		c.CheckRedirect = func(*http.Request, []*http.Request) error {
			return fmt.Errorf(noRedirectError)
		}
	}
	return &c
}

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

// IsRecord is a convenience wrapper that returns true if the Record bit is set on the flag.
func (f Format) IsRecord() bool {
	return f&Record != 0
}

// DoHidden is equivalent to Do with the exception that nothing gets logged unless DumpFormat is
// set to Verbose.
func (d *dumpClient) DoHidden(req *http.Request) (*http.Response, error) {
	return d.doImp(req, true, nil)
}

// Do dumps the request, makes the request and dumps the response as specified by DumpFormat.
func (d *dumpClient) Do(req *http.Request) (*http.Response, error) {
	return d.doImp(req, false, nil)
}

func (d *dumpClient) DoWithContext(ctx context.Context, req *http.Request) (*http.Response, error) {
	return d.doImp(req, true, ctx)
}

func (d *dumpClient) DoHiddenWithContext(ctx context.Context, req *http.Request) (*http.Response, error) {
	return d.doImp(req, false, ctx)
}

// doImp actually performs the HTTP request logging according to the various settings.
func (d *dumpClient) doImp(req *http.Request, hidden bool, ctx context.Context) (*http.Response, error) {
	if req.URL.Scheme == "" {
		if Insecure {
			req.URL.Scheme = "http"
		} else {
			req.URL.Scheme = "https"
		}
	}
	req.Header.Set("User-Agent", UA)

	var reqBody []byte
	startedAt := time.Now()

	// prefer the X-Request-Id header as request token for logging, if present.
	id := req.Header.Get(requestIdHeader)
	if id == "" {
		id = ShortToken()
	}
	log.Info("started", "id", id, req.Method, req.URL.String())
	hide := (DumpFormat == NoDump) || (hidden && !DumpFormat.IsVerbose())
	if !hide {
		startedAt = time.Now()
		reqBody = dumpRequest(req)
	}
	var resp *http.Response
	var err error
	if ctx == nil {
		resp, err = d.Client.Do(req)
	} else {
		resp, err = ctxhttp.Do(ctx, d.getClientWithoutTimeout(), req)
	}
	if urlError, ok := err.(*url.Error); ok {
		if urlError.Err.Error() == noRedirectError {
			err = nil
		}
	}
	if err != nil {
		return nil, err
	}
	if !hide {
		dumpResponse(resp, req, reqBody)
	}
	log.Info("completed", "id", id, "status", resp.Status, "time", time.Since(startedAt).String())

	return resp, nil
}

// getClientWithoutTimeout returns a modified client that doesn't have the ResponseHeaderTimeout field set
// in its Transport.
func (d *dumpClient) getClientWithoutTimeout() *http.Client {
	// Get a copy of the client and modify as multiple concurrent go routines can be using this client.
	client := *d.Client
	tr, ok := client.Transport.(*http.Transport)
	if ok {
		// note that the http.Transport struct has internal mutex fields that are
		// not safe to copy. we have to be selective in copying fields.
		trCopy := &http.Transport{
			Proxy:                  tr.Proxy,
			DialContext:            tr.DialContext,
			Dial:                   tr.Dial,
			DialTLS:                tr.DialTLS,
			TLSClientConfig:        tr.TLSClientConfig,
			TLSHandshakeTimeout:    tr.TLSHandshakeTimeout,
			DisableKeepAlives:      tr.DisableKeepAlives,
			DisableCompression:     tr.DisableCompression,
			MaxIdleConns:           tr.MaxIdleConns,
			MaxIdleConnsPerHost:    tr.MaxIdleConnsPerHost,
			IdleConnTimeout:        tr.IdleConnTimeout,
			ResponseHeaderTimeout:  0, // explicitly zeroed-out
			ExpectContinueTimeout:  tr.ExpectContinueTimeout,
			TLSNextProto:           tr.TLSNextProto,
			MaxResponseHeaderBytes: tr.MaxResponseHeaderBytes,
		}
		tr = trCopy
	} else {
		// note that the following code has a known issue in that it depends on the
		// current value of the NoCertCheck global. if that global changes after
		// creation of this client then the behavior is undefined.
		tr = &http.Transport{Proxy: http.ProxyFromEnvironment}
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: NoCertCheck}
	}
	client.Transport = tr
	return &client
}

// Dump request if needed.
// Return request serialized as JSON if DumpFormat is JSON, nil otherwise.
func dumpRequest(req *http.Request) []byte {
	if DumpFormat == NoDump {
		return nil
	}
	reqBody, err := dumpReqBody(req)
	if err != nil {
		log.Error("Failed to load request body for dump", "error", err.Error())
	}
	if DumpFormat.IsDebug() {
		var buffer bytes.Buffer
		buffer.WriteString(req.Method + " " + req.URL.String() + "\n")
		writeHeaders(&buffer, req.Header)
		if reqBody != nil {
			buffer.WriteString("\n")
			buffer.Write(reqBody)
			buffer.WriteString("\n")
		}
		fmt.Fprint(OsStderr, buffer.String())
	} else if DumpFormat.IsJSON() {
		return reqBody
	}
	return nil
}

// dumpResponse dumps the response and optionally the request (in case of JSON format) according to
// DumpFormat.
// It also checks whether the special recorder pipe is opened and if so writes the dump to it.
func dumpResponse(resp *http.Response, req *http.Request, reqBody []byte) {
	if DumpFormat == NoDump {
		return
	}
	respBody, _ := dumpRespBody(resp)
	if DumpFormat.IsDebug() {
		var buffer bytes.Buffer
		buffer.WriteString("==> " + resp.Proto + " " + resp.Status + "\n")
		writeHeaders(&buffer, resp.Header)
		if respBody != nil {
			buffer.WriteString("\n")
			buffer.Write(respBody)
			buffer.WriteString("\n")
		}
		fmt.Fprint(OsStderr, buffer.String())
	} else if DumpFormat.IsJSON() {
		reqHeaders := make(http.Header)
		filterHeaders(req.Header, func(name string, value []string) {
			reqHeaders[name] = value
		})
		respHeaders := make(http.Header)
		filterHeaders(resp.Header, func(name string, value []string) {
			respHeaders[name] = value
		})
		dumped := recording.RequestResponse{
			Verb:       req.Method,
			URI:        req.URL.String(),
			ReqHeader:  reqHeaders,
			ReqBody:    string(reqBody),
			Status:     resp.StatusCode,
			RespHeader: respHeaders,
			RespBody:   string(respBody),
		}
		b, err := json.MarshalIndent(dumped, "", "    ")
		if err != nil {
			log.Error("Failed to dump request content", "error", err.Error())
			return
		}
		if DumpFormat.IsRecord() {
			f := os.NewFile(10, "fd10")
			_, err = f.Stat()
			if err == nil {
				// fd 10 is open, dump to it (used by recorder)
				fmt.Fprintf(f, "%s\n", string(b))
			}
		}
		fmt.Fprint(OsStderr, string(b))
	}
}

// writeHeaders is a helper function that writes the given HTTP headers to the given buffer as
// human readable strings. If DumpFormat is not Verbose then writeHeaders filters out headers whose
// names are keys of HiddenHeaders.
func writeHeaders(buffer *bytes.Buffer, headers http.Header) {
	filterHeaders(headers, func(name string, value []string) {
		buffer.WriteString(name)
		buffer.WriteString(": ")
		buffer.WriteString(strings.Join(value, ", "))
		buffer.WriteString("\n")
	})
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

// headerIterator is a HTTP header iterator.
type headerIterator func(name string, value []string)

// filterHeaders iterates through the headers skipping hidden headers unless DumpFormat is Verbose.
// It calls the given iterator for each header name/value pair. The values are serialized as
// strings.
func filterHeaders(headers http.Header, iterator headerIterator) {
	for k, v := range headers {
		if !DumpFormat.IsVerbose() {
			if _, ok := HiddenHeaders[k]; ok {
				continue
			}
		}
		iterator(k, v)
	}
}
