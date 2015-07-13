package httpclient_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/rightscale/rsc/httpclient"
)

// Helper function that creates a http.Header from a name value pair
func httpHeaders(name, value string) http.Header {
	m := map[string][]string{name: []string{value}}
	return http.Header(m)
}

var _ = Describe("HTTP client", func() {
	var client httpclient.HTTPClient
	var server *ghttp.Server
	var stderr bytes.Buffer
	var req *http.Request
	var resp *http.Response
	var useHidden bool

	BeforeEach(func() {
		httpclient.NoCertCheck = true
		httpclient.OsStderr = &stderr
		server = ghttp.NewTLSServer()
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("POST", "/redirect"),
			ghttp.RespondWith(303, "", httpHeaders("Location", server.URL()+"/ok")),
		))
		server.AppendHandlers(ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/ok"),
			ghttp.RespondWith(200, "OK"),
		))
	})

	JustBeforeEach(func() {
		stderr.Reset()
		var err error
		var body = strings.NewReader(`{"foo":"bar"}`)
		req, err = http.NewRequest("POST", server.URL()+"/redirect", body)
		Ω(err).ShouldNot(HaveOccurred())
		if useHidden {
			resp, err = client.DoHidden(req)
		} else {
			resp, err = client.Do(req)
		}
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Context("created with New", func() {

		BeforeEach(func() {
			client = httpclient.New()
		})

		It("follows redirects", func() {
			Ω(resp.StatusCode).Should(Equal(200))
			rb, err := ioutil.ReadAll(resp.Body)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(rb).Should(Equal([]byte("OK")))
		})

		It("does not dump by default", func() {
			Ω(stderr.String()).Should(BeEmpty())
		})

		Context("setting DumpFormat to debug", func() {
			BeforeEach(func() {
				httpclient.DumpFormat = httpclient.Debug
			})

			AfterEach(func() {
				httpclient.DumpFormat = httpclient.NoDump
			})

			It("dumps the request and response details", func() {
				Ω(stderr.String()).ShouldNot(BeEmpty())
				Ω(stderr.String()).Should(ContainSubstring("POST"))
				Ω(stderr.String()).Should(ContainSubstring(server.URL()))
				Ω(stderr.String()).Should(ContainSubstring("200 OK"))
			})

			Context("using DoHidden", func() {
				BeforeEach(func() {
					useHidden = true
				})

				AfterEach(func() {
					useHidden = false
				})

				It("does not dump", func() {
					Ω(stderr.String()).Should(BeEmpty())
				})
			})

		})

		Context("setting DumpFormat to JSON", func() {
			BeforeEach(func() {
				httpclient.DumpFormat = httpclient.JSON
			})

			AfterEach(func() {
				httpclient.DumpFormat = httpclient.NoDump
			})

			It("dumps the request and response details in JSON", func() {
				Ω(stderr.String()).ShouldNot(BeEmpty())
				var js map[string]interface{}
				err := json.Unmarshal(stderr.Bytes(), &js)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(js).ShouldNot(BeEmpty())
			})

			Context("using DoHidden", func() {
				BeforeEach(func() {
					useHidden = true
				})

				AfterEach(func() {
					useHidden = false
				})

				It("does not dump", func() {
					Ω(stderr.String()).Should(BeEmpty())
				})

				Context("using Verbose", func() {
					BeforeEach(func() {
						httpclient.DumpFormat = httpclient.JSON | httpclient.Verbose
					})

					AfterEach(func() {
						httpclient.DumpFormat = httpclient.NoDump
					})

					It("does dump", func() {
						Ω(stderr.String()).ShouldNot(BeEmpty())
						var js map[string]interface{}
						err := json.Unmarshal(stderr.Bytes(), &js)
						Ω(err).ShouldNot(HaveOccurred())
						Ω(js).ShouldNot(BeEmpty())
					})
				})

			})
		})
	})

	Context("created with NewNoRedirect", func() {
		BeforeEach(func() {
			client = httpclient.NewNoRedirect()
		})

		It("does not follow redirects", func() {
			Ω(resp.StatusCode).Should(Equal(303))
		})
	})

})
