package main_test

import (
	"bytes"
	"fmt"
	"io"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc"
)

var _ = Describe("Printer", func() {
	var (
		testOut *bytes.Buffer
		testIn  io.Reader
		format  string
	)

	JustBeforeEach(func() {
		main.SetOutput(testOut)
		main.SetErrorOutput(testOut)
		main.SetInput(testIn)
	})

	AfterEach(func() {
		main.SetOutput(os.Stdout)
		main.SetErrorOutput(os.Stderr)
		main.SetInput(os.Stdin)
	})

	Context("when given a single argument", func() {
		BeforeEach(func() {
			format = "A simple string"
			testOut = &bytes.Buffer{}
			testIn = bytes.NewReader([]byte("y\n"))
		})

		It("prompts confirmation", func() {
			y := main.PromptConfirmation(format)
			Ω(y).Should(Equal("y"))
			Ω(testOut.String()).Should(Equal(format))
		})

		It("prints subtitles", func() {
			main.PrintSubtitle(format)
			Ω(testOut.String()).Should(ContainSubstring(format + "\n"))
		})

		It("prints titles", func() {
			main.PrintTitle(format)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(format))
		})

		It("prints success", func() {
			main.PrintSuccess(format)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(format))
		})

		It("prints errors", func() {
			main.PrintError(format)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(format))
		})
	})

	Context("when given a format", func() {
		var (
			suffix           string
			formatWithSuffix string
		)

		BeforeEach(func() {
			suffix = "foo"
			format = "A simple string with suffix '%s'"
			formatWithSuffix = fmt.Sprintf(format, suffix)
			testOut = &bytes.Buffer{}
			testIn = bytes.NewReader([]byte("y\n"))
		})

		It("prompts confirmation", func() {
			var y = main.PromptConfirmation(format, suffix)
			Ω(y).Should(Equal("y"))
			Ω(testOut.String()).Should(Equal(formatWithSuffix))
		})

		It("prints subtitles", func() {
			main.PrintSubtitle(format, suffix)
			Ω(testOut.String()).Should(ContainSubstring(formatWithSuffix + "\n"))
		})

		It("prints titles", func() {
			main.PrintTitle(format, suffix)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(formatWithSuffix))
		})

		It("prints success", func() {
			main.PrintSuccess(format, suffix)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(formatWithSuffix))
		})

		It("prints errors", func() {
			main.PrintError(format, suffix)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(formatWithSuffix))
		})
	})

})
