package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Helper that creates test responses
func makeResponse(body string, headers map[string][]string) *http.Response {
	b := bytes.NewBufferString(body)
	return &http.Response{
		Body:   ioutil.NopCloser(b),
		Header: http.Header(headers),
	}
}

var _ = Describe("Displayer", func() {
	var (
		resp      *http.Response
		displayer *Displayer
		err       error
	)

	JustBeforeEach(func() {
		displayer, err = NewDisplayer(resp)
	})

	Context("with a response containing invalid JSON", func() {
		var invalidJSON = "{notvalid}"
		BeforeEach(func() {
			resp = makeResponse(invalidJSON, nil)
		})

		It("NewDisplayer contains the raw value", func() {
			Ω(displayer).ShouldNot(BeNil())
			Ω(displayer.RawOutput).Should(Equal(invalidJSON))
			Ω(err).Should(BeNil())
		})
	})

	Context("with a response containing valid JSON", func() {
		var (
			validJSON = `{"foo":"bar","baz":42,"m":[{"a":1},{"a":2}]}`
		)

		BeforeEach(func() {
			resp = makeResponse(validJSON, nil)
		})

		It("NewDisplayer returns a new displayer and no error", func() {
			Ω(displayer).ShouldNot(BeNil())
			Ω(err).Should(BeNil())
		})

		It("Applies json selects that extract a single value", func() {
			Ω(displayer).ShouldNot(BeNil())
			Ω(displayer.ApplySingleExtract(".baz")).ShouldNot(HaveOccurred())
			Ω(displayer.RawOutput).Should(Equal(float64(42)))
		})

		It("Applies json selects that extract multiple values in JSON", func() {
			Ω(displayer).ShouldNot(BeNil())
			Ω(displayer.ApplyExtract(".m .a", true)).ShouldNot(HaveOccurred())
			Ω(displayer.RawOutput).Should(Equal([]interface{}{float64(1), float64(2)}))
		})

		It("Applies json selects that extract multiple values not in JSON", func() {
			Ω(displayer).ShouldNot(BeNil())
			Ω(displayer.ApplyExtract(".m .a", false)).ShouldNot(HaveOccurred())
			Ω(displayer.RawOutput).Should(Equal("1\n2\n"))
		})

		Context("with the go value corresponding to the JSON data", func() {
			var (
				value interface{}
			)

			BeforeEach(func() {
				json.Unmarshal([]byte(validJSON), &value)
			})

			It("returns the output", func() {
				Ω(displayer).ShouldNot(BeNil())
				var output, _ = json.Marshal(value)
				Ω(displayer.Output()).Should(Equal(string(output)))
			})

			It("prettifies", func() {
				Ω(displayer).ShouldNot(BeNil())
				displayer.Pretty()
				var pretty, _ = json.MarshalIndent(value, "", "    ")
				Ω(displayer.Output()).Should(Equal(string(pretty) + "\n"))
			})
		})

	})

	Context("with a response containing headers", func() {
		var (
			headerValue = []string{"foo", "bar"}
			headers     = map[string][]string{"X-Test": headerValue}
		)

		BeforeEach(func() {
			resp = makeResponse("", headers)
		})

		It("ApplyHeaderExtract returns the extracted header value", func() {
			Ω(err).Should(BeNil())
			Ω(displayer).ShouldNot(BeNil())
			Ω(displayer.ApplyHeaderExtract("X-Test")).ShouldNot(HaveOccurred())
			Ω(displayer.RawOutput).Should(Equal("foo"))
		})

		It("ApplyHeaderExtract returns an error if header is missing", func() {
			Ω(err).Should(BeNil())
			Ω(displayer).ShouldNot(BeNil())
			Ω(displayer.ApplyHeaderExtract("X-NonExistent")).Should(HaveOccurred())
		})
	})
})
