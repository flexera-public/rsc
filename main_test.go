package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/rightscale/rsc/cmd"
	"gopkg.in/alecthomas/kingpin.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type timeoutError struct{ error }

func (e *timeoutError) Error() string   { return "i/o timeout" }
func (e *timeoutError) Timeout() bool   { return true }
func (e *timeoutError) Temporary() bool { return true }

var _ = Describe("Main", func() {

	Context("Creating response from JSON input", func() {
		var (
			resp     *http.Response
			json     = []byte(`{"foo":"bar","baz":42,"m":[{"a":1},{"a":2}]}`)
			bom      = []byte{0xef, 0xbb, 0xbf}
			respBody []byte
			err      error
		)

		JustBeforeEach(func() {
			resp = CreateJSONResponse(json)
			respBody, err = ioutil.ReadAll(resp.Body)
		})

		Context("Creates valid response from JSON without BOM", func() {
			It("creates valid response", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(bytes.HasPrefix(respBody, bom)).Should(BeFalse())
				Ω(respBody).Should(Equal(json))
			})
		})

		Context("Creates valid response from JSON with BOM", func() {
			BeforeEach(func() {
				json = append(bom, json...)
			})
			It("creates valid response", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(bytes.HasPrefix(respBody, bom)).Should(BeFalse())
				Ω(respBody).Should(Equal(bytes.TrimPrefix(json, bom)))
			})
		})

	})

	Context("retry option", func() {
		var app = kingpin.New("rsc", "rsc - tests")
		var retries = 6
		var cmdLine = cmd.CommandLine{Retry: retries}
		var origDoAPIRequest func(string, *cmd.CommandLine) (*http.Response, error)
		BeforeEach(func() {
			origDoAPIRequest = doAPIRequest
		})
		AfterEach(func() {
			doAPIRequest = origDoAPIRequest
		})
		It("retries if error on API call occurs", func() {
			counter := 0
			doAPIRequest = func(string, *cmd.CommandLine) (*http.Response, error) {
				counter += 1
				return nil, &timeoutError{}
			}
			ExecuteCommand(app, &cmdLine)
			Ω(counter).Should(Equal(1 + retries))
		})

		It("doesn't retry more than necessary", func() {
			counter := 0
			doAPIRequest = func(string, *cmd.CommandLine) (*http.Response, error) {
				counter += 1
				if counter < 3 {
					return &http.Response{StatusCode: 503}, nil
				} else {
					return nil, nil
				}
			}
			_, err := ExecuteCommand(app, &cmdLine)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(counter).Should(BeNumerically("<", 1+retries))
		})
	})

})
