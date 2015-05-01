package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Printer", func() {
	var (
		testOut *bytes.Buffer
		testIn  io.Reader
		format  string
	)

	JustBeforeEach(func() {
		SetOutput(testOut)
		SetErrorOutput(testOut)
		SetInput(testIn)
	})

	AfterEach(func() {
		SetOutput(os.Stdout)
		SetErrorOutput(os.Stderr)
		SetInput(os.Stdin)
	})

	Context("when given a single argument", func() {
		BeforeEach(func() {
			format = "A simple string"
			testOut = &bytes.Buffer{}
			testIn = bytes.NewReader([]byte("y\n"))
		})

		It("prompts confirmation", func() {
			y := PromptConfirmation(format)
			Ω(y).Should(Equal("y"))
			Ω(testOut.String()).Should(Equal(format))
		})

		It("prints subtitles", func() {
			PrintSubtitle(format)
			Ω(testOut.String()).Should(ContainSubstring(format + "\n"))
		})

		It("prints titles", func() {
			PrintTitle(format)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(format))
		})

		It("prints success", func() {
			PrintSuccess(format)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(format))
		})

		It("prints errors", func() {
			PrintError(format)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(format))
		})

		It("prints fatals", func() {
			var exitCode int
			osExit = func(code int) { exitCode = code }
			PrintFatal(format)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(format))
			Ω(exitCode).Should(Equal(1))
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
			var y = PromptConfirmation(format, suffix)
			Ω(y).Should(Equal("y"))
			Ω(testOut.String()).Should(Equal(formatWithSuffix))
		})

		It("prints subtitles", func() {
			PrintSubtitle(format, suffix)
			Ω(testOut.String()).Should(ContainSubstring(formatWithSuffix + "\n"))
		})

		It("prints titles", func() {
			PrintTitle(format, suffix)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(formatWithSuffix))
		})

		It("prints success", func() {
			PrintSuccess(format, suffix)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(formatWithSuffix))
		})

		It("prints errors", func() {
			PrintError(format, suffix)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(formatWithSuffix))
		})

		It("prints fatals", func() {
			var exitCode int
			osExit = func(code int) { exitCode = code }
			PrintFatal(format, suffix)
			var rawOut = testOut.String()
			Ω(rawOut).Should(ContainSubstring(formatWithSuffix))
			Ω(exitCode).Should(Equal(1))
		})
	})

})
