// Package recording defines the data structures used to record and replay invokations of "rsc" for
// testing.
package recording

import "net/http"

// Recording represents a single "rsc" invokation recording.
type Recording struct {
	CmdArgs  []string        // command line arguments
	ExitCode int             // Exit code
	Stdout   string          // Exit print
	RR       RequestResponse // back-end request/response
}

// RequestResponse contains an HTTP Request and the corresponding response details.
type RequestResponse struct {
	Verb, URI  string      // request http verb and full uri with query string
	ReqHeader  http.Header // headers before std additions, such as user-agent
	ReqBody    string      // not []byte so that json.Marshal doesn't produce base64
	Status     int         // numerical response status
	RespHeader http.Header // full response headers
	RespBody   string      // not []byte so that json.Marshal doesn't produce base64
}
