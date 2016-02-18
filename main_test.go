package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rightscale/rsc/cmd"
	"gopkg.in/alecthomas/kingpin.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
		var cmdLine = cmd.CommandLine{Retry: 6}
		It("retries if error on API call occurs", func() {
			counter := 0
			doAPIRequest = func(string, *cmd.CommandLine) (*http.Response, error) {
				counter += 1
				return nil, errors.New("test")
			}
			ExecuteCommand(app, &cmdLine)
			Ω(counter).Should(Equal(6))
		})

		It("doesn't retries more than necessary", func() {
			counter := 0
			doAPIRequest = func(string, *cmd.CommandLine) (*http.Response, error) {
				counter += 1
				if counter < 3 {
					return nil, errors.New("test")
				} else {
					return nil, nil
				}
			}
			_, err := ExecuteCommand(app, &cmdLine)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(counter).ShouldNot(Equal(6))
		})
	})

})
