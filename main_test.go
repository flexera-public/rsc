package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

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

})
