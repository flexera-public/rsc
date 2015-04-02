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
			fmt.Fprintf(os.Stderr, "Json decode: %s\n", err.Error())
		}
		// Perform the test by running main() with the command line args set
		It(strings.Join(testCase.CmdArgs, " "), func() {
			//server := httptest.NewTLSServer(http.HandlerFunc(handler))
			server := ghttp.NewServer()
			defer server.Close()

			// construct list of verifiers
			uri := regexp.MustCompile(`https?://[^/]+`).
				ReplaceAllString(testCase.RR.Uri, "")
			handlers := []http.HandlerFunc{
				ghttp.VerifyRequest(testCase.RR.Verb, uri),
			}
			if len(testCase.RR.ReqBody) > 0 {
				handlers = append(handlers,
					ghttp.VerifyJSON(testCase.RR.ReqBody))
			}
			for k, _ := range testCase.RR.ReqHeader {
				handlers = append(handlers,
					ghttp.VerifyHeaderKV(k, testCase.RR.ReqHeader.Get(k)))
			}
			handlers = append(handlers,
				ghttp.RespondWith(testCase.RR.Status, testCase.RR.RespBody,
					http.Header{"Content-Type": []string{"application/json"}}))
			server.AppendHandlers(ghttp.CombineHandlers(handlers...))

			os.Args = append([]string{
				"rsc", "--noAuth",
				"--host", strings.TrimPrefix(server.URL(), "http://")},
				testCase.CmdArgs...)
			fmt.Fprintf(os.Stderr, "testing \"%s\"\n", strings.Join(os.Args, `" "`))

			// capture stdout and intercept calls to osExit
			stdoutBuf := bytes.Buffer{}
			SetOutput(&stdoutBuf)
			exitCode := 99
			osExit = func(code int) { exitCode = code }

			main()

			// Verify that stdout and the exit code are correct
			//fmt.Fprintf(os.Stderr, "Exit %d %d\n", exitCode, testCase.ExitCode)
			Ω(exitCode).Should(Equal(testCase.ExitCode))
			//fmt.Fprintf(os.Stderr, "stdout got <<%s>> expected <<%s>>\n",
			//stdoutBuf.String(), testCase.Stdout)
			Ω(stdoutBuf.String()).Should(Equal(testCase.Stdout))
		})
	}

})
