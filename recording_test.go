// Copyright (c) 2015 RightScale, Inc. - see LICENSE

package main

// Omega: Alt+937

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/rightscale/rsc/recording"
)

// Iterate through all recorded test cases and play them back
var _ = Describe("Recorded request", func() {

	// Open the recording file
	f, err := os.Open("recorder/recording.json")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Cannot open recording: %s\n", err.Error())
		os.Exit(1)
	}
	decoder := json.NewDecoder(f)
	// Iterate through test cases
	for {
		var testCase recording.Recording
		err := decoder.Decode(&testCase)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "JSON decode: %s\n", err.Error())
		}
		// Perform the test by running main() with the command line args set
		It(strings.Join(testCase.CmdArgs, " "), func() {
			//server := httptest.NewTLSServer(http.HandlerFunc(handler))
			server := ghttp.NewServer()
			defer server.Close()

			// construct list of verifiers
			url := regexp.MustCompile(`https?://[^/]+(/[^?]+)\??(.*)`).
				FindStringSubmatch(testCase.RR.URI)
			//fmt.Fprintf(os.Stderr, "URL: %#v\n", url)
			handlers := []http.HandlerFunc{
				ghttp.VerifyRequest(testCase.RR.Verb, url[1], url[2]),
			}
			if len(testCase.RR.ReqBody) > 0 {
				handlers = append(handlers,
					ghttp.VerifyJSON(testCase.RR.ReqBody))
			}
			for k := range testCase.RR.ReqHeader {
				handlers = append(handlers,
					ghttp.VerifyHeaderKV(k, testCase.RR.ReqHeader.Get(k)))
			}
			respHeader := make(http.Header)
			for k, v := range testCase.RR.RespHeader {
				respHeader[k] = v
			}
			handlers = append(handlers,
				ghttp.RespondWith(testCase.RR.Status, testCase.RR.RespBody,
					respHeader))
			server.AppendHandlers(ghttp.CombineHandlers(handlers...))

			os.Args = append([]string{
				"rsc", "--noAuth", "--dump", "debug",
				"--host", strings.TrimPrefix(server.URL(), "http://")},
				testCase.CmdArgs...)
			//fmt.Fprintf(os.Stderr, "testing \"%s\"\n", strings.Join(os.Args, `" "`))

			// capture stdout and intercept calls to osExit
			stdoutBuf := bytes.Buffer{}
			SetOutput(&stdoutBuf)
			exitCode := 99
			osExit = func(code int) { exitCode = code }
			stderrBuf := bytes.Buffer{}
			SetErrorOutput(&stderrBuf)

			main()

			// Verify that stdout and the exit code are correct
			//fmt.Fprintf(os.Stderr, "Exit %d %d\n", exitCode, testCase.ExitCode)
			//fmt.Fprintf(os.Stderr, "stdout got <<%q>>\n  expected <<%q>>\n",
			//	stdoutBuf.String(), testCase.Stdout)
			//fmt.Fprintf(os.Stderr, "stderr got <<%q>>\n", stderrBuf.String())
			Ω(exitCode).Should(Equal(testCase.ExitCode), "Exit code doesn't match")
			Ω(stdoutBuf.String()).Should(Equal(testCase.Stdout), "Stdout doesn't match")
		})
	}

})
