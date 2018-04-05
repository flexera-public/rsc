package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"testing"
)

func TestUploadExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Upload Example Suite")
}

var _ = Describe("Upload Example", func() {
	var server *ghttp.Server
	var out bytes.Buffer

	// Override default failure handler so it doesn't kill the process
	fail = func(format string, v ...interface{}) {
		Fail(fmt.Sprintf(format, v...))
	}

	BeforeEach(func() {
		server = ghttp.NewTLSServer()
		os.Args = []string{"upload", "-insecure", "-e=dummy", "-p=dummy", "-a=42",
			"-h=" + server.URL()[8:], "--cat=test.cat"}
		osStdout = &out
	})

	AfterEach(func() {
		server.Close()
	})

	It("creates a cat", func() {
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/api/sessions"),
				ghttp.RespondWith(204, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/catalog/new_session"),
				ghttp.RespondWith(303, "", http.Header(map[string][]string{"Location": []string{"foo"}})),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/catalog/accounts/42/user_preferences"),
				ghttp.RespondWith(200, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/api/designer/collections/42/templates"),
				ghttp.RespondWith(201, "", http.Header(map[string][]string{"Location": []string{"/api/designer/collections/42/templates/1"}})),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/designer/collections/42/templates/1"),
				ghttp.RespondWith(201, responseBody),
			),
		)
		main()
		Ω(out.String()).Should(HavePrefix(output))
	})

})

// Additional handler that can be used to debug requests
func trace(rw http.ResponseWriter, req *http.Request) {
	d, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(d) + "\n")
}

const (
	output = "template_8_template_tests__2015-04-29_21:05:26.663 created at 2015-04-29 21:05:45.163"

	responseBody = `{
	   "kind":"self_service#template",
	   "id":"554147a973656c670c1f0000",
	   "name":"template_8_template_tests__2015-04-29_21:05:26.663",
	   "filename":"template.rb",
	   "href":"/api/designer/collections/62656/templates/554147a973656c670c1f0000",
	   "short_description":"upload of Template",
	   "created_by":{
	      "id":10662,
	      "name":"Joe",
	      "email":"joe@rightscale.com"
	   },
	   "timestamps":{
	      "created_at":"2015-04-29T21:05:45.163+00:00",
	      "updated_at":"2015-04-29T21:05:45.163+00:00",
	      "published_at":null
	   }
	}`
)
