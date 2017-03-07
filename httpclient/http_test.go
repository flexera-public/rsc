package httpclient_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/onsi/gomega/ghttp"
	"github.com/rightscale/rsc/httpclient"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Helper function that creates a http.Header from a name value pair
func httpHeaders(name, value string) http.Header {
	m := map[string][]string{name: []string{value}}
	return http.Header(m)
}

var _ = Describe("HTTP client", func() {
	var (
		client         httpclient.HTTPClient
		server         *ghttp.Server
		stderr         bytes.Buffer
		req            *http.Request
		resp           *http.Response
		oldNoCertCheck bool
		oldOsStderr    io.Writer
		useHidden      bool
		err            error
		expectError    bool
	)

	BeforeEach(func() {
		oldNoCertCheck = httpclient.NoCertCheck
		oldOsStderr = httpclient.OsStderr
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
		var body = strings.NewReader(`{"foo":"bar"}`)
		req, err = http.NewRequest("POST", server.URL()+"/redirect", body)
		Ω(err).ShouldNot(HaveOccurred())
		if useHidden {
			resp, err = client.DoHidden(req)
		} else {
			resp, err = client.Do(req)
		}
		if expectError {
			Ω(err).Should(HaveOccurred())
		} else {
			Ω(err).ShouldNot(HaveOccurred())
		}
	})

	AfterEach(func() {
		server.Close()
		server = nil
		client = nil
		stderr.Reset()
		req = nil
		resp = nil
		useHidden = false
		err = nil
		expectError = false
		httpclient.NoCertCheck = oldNoCertCheck
		httpclient.OsStderr = oldOsStderr
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

	Context("created with NewNoRedirect", func() {
		BeforeEach(func() {
			client = httpclient.NewNoRedirect()
		})

		It("does not follow redirects", func() {
			Ω(resp.StatusCode).Should(Equal(303))
		})
	})

	Context("created with NewPB", func() {
		Context("with NoRedirect=false", func() {
			var (
				pb            *httpclient.ParamBlock
				oldDumpFormat httpclient.Format
			)

			BeforeEach(func() {
				oldDumpFormat = httpclient.DumpFormat
				pb = &httpclient.ParamBlock{
					NoCertCheck: true,
					NoRedirect:  false,
				}
				client = httpclient.NewPB(pb)
			})

			AfterEach(func() {
				pb = nil
				httpclient.DumpFormat = oldDumpFormat
			})

			It("follows redirects", func() {
				Ω(resp.StatusCode).Should(Equal(200))
				rb, err := ioutil.ReadAll(resp.Body)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(rb).Should(Equal([]byte("OK")))
			})

			It("does not dump by default", func() {
				httpclient.DumpFormat = httpclient.Debug // ignored package variable
				Ω(stderr.String()).Should(BeEmpty())
			})

			Context("setting DumpFormat to debug", func() {
				BeforeEach(func() {
					// ignored package variables
					httpclient.DumpFormat = httpclient.NoDump
					httpclient.HiddenHeaders = map[string]bool{"User-Agent": true}

					pb.DumpFormat = httpclient.Debug
					client = httpclient.NewPB(pb)
				})

				Context("with default hidden headers", func() {
					It("dumps the request and response details with default hidden headers", func() {
						Ω(stderr.String()).ShouldNot(BeEmpty())
						Ω(stderr.String()).Should(ContainSubstring("POST"))
						Ω(stderr.String()).Should(ContainSubstring("User-Agent"))
						Ω(stderr.String()).Should(ContainSubstring(server.URL()))
						Ω(stderr.String()).Should(ContainSubstring("200 OK"))
					})
				})

				Context("with custom hidden headers", func() {
					BeforeEach(func() {
						pb.HiddenHeaders = map[string]bool{"User-Agent": true}
						client = httpclient.NewPB(pb)
					})

					It("dumps the request and response details with default hidden headers", func() {
						Ω(stderr.String()).ShouldNot(BeEmpty())
						Ω(stderr.String()).Should(ContainSubstring("POST"))
						Ω(stderr.String()).ShouldNot(ContainSubstring("User-Agent"))
						Ω(stderr.String()).Should(ContainSubstring(server.URL()))
						Ω(stderr.String()).Should(ContainSubstring("200 OK"))
					})
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
		})

		Context("with NoRedirect=true", func() {
			BeforeEach(func() {
				client = httpclient.NewPB(
					&httpclient.ParamBlock{
						NoCertCheck: true,
						NoRedirect:  true,
					})
			})

			It("does not follow redirects", func() {
				Ω(resp.StatusCode).Should(Equal(303))
			})
		})

		Context("with NoCertCheck=false", func() {
			BeforeEach(func() {
				httpclient.NoCertCheck = true // ignored package variable
				client = httpclient.NewPB(
					&httpclient.ParamBlock{
						NoCertCheck: false,
					})
				expectError = true
			})

			It("fails cert check", func() {
				Ω(err.Error()).Should(ContainSubstring("certificate signed by unknown authority"))
			})
		})
	})

	// CAN'T DO: Proxy support must be tested using integration tests. See:
	// https://go.googlesource.com/go/+/go1.4.2/src/net/http/transport.go#296 dealing
	// with httpProxyEnv instantions. The env is read only once ever and the methods
	// to reset this mechanism for testing purposes are internal to the http package.
	// Context("with proxy", func() {
	// 	BeforeEach(func() {
	// 	  os.Setenv("http_proxy", "http://1.2.3.4:90")
	// 		client = httpclient.New()
	// 	})
	//})

})
