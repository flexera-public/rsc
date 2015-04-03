package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/rightscale/rsc/recording"
)

const output = "recording_new.json"

var exitRegexp = regexp.MustCompile(`exit status (\d+)`)

func main() {
	args := []string{"--dump=json"}
	args = append(args, os.Args[1:]...)
	cmd := exec.Command("rsc", args...)
	_, args = extractArg("--dump", args)
	_, args = extractArg("--host", args)
	_, args = extractArg("--key", args)
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	err := cmd.Run()
	exitCode := 0
	if err != nil {
		exCode := exitRegexp.FindStringSubmatch(err.Error())
		if len(exCode) == 2 {
			exitCode, err = strconv.Atoi(exCode[1])
			if err != nil {
				fail("Invalid exit code '%s': %s", exCode[1], err.Error())
			}
		} else {
			fail("%s %s failed: %s\n%s\n", "rsc", strings.Join(args, " "), err, out.String())
		}
	}
	var rr recording.RequestResponse
	err = json.Unmarshal(outErr.Bytes(), &rr)
	if err != nil {
		fail("load rsc output: %s - output was:\n%s\n", err, outErr.String())
	}
	rr.ReqHeader.Del("Authorization")
	rr.ReqHeader.Del("User-Agent")
	rr.RespHeader.Del("Cache-Control")
	rr.RespHeader.Del("Connection")
	rr.RespHeader.Del("Set-Cookie")
	rr.RespHeader.Del("Strict-Transport-Security")
	rr.RespHeader.Del("X-Request-Uuid")
	record := recording.Recording{
		CmdArgs:  args,
		ExitCode: exitCode,
		Stdout:   out.String(),
		RR:       rr,
	}
	js, err := json.MarshalIndent(record, "", "    ")
	if err != nil {
		fail("failed to serialize record: %s", err)
	}
	write(js)
	// Echo output of rsc so calling script can use it
	fmt.Print(out.String())
}

// Extract command line argument with given name and return remaining arguments
func extractArg(name string, args []string) (string, []string) {
	var val string
	var newArgs []string
	var skip bool
	for i, a := range args {
		if skip {
			continue
		}
		if strings.Contains(a, "=") {
			elems := strings.SplitN(a, "=", 2)
			if elems[0] == name {
				val = elems[1]
			} else {
				newArgs = append(newArgs, a)
			}
		} else if a == name && len(args) > (i+1) {
			val = args[i+1]
			skip = true
		} else {
			newArgs = append(newArgs, a)
		}
	}
	return val, newArgs
}

// Helper function that appends a string to output file
func write(b []byte) {
	f, err := os.OpenFile(output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fail("failed to open output file")
	}
	f.Write(b)
	f.WriteString("\n")
	f.Close()
}

// Print error message and exit with status code 1
func fail(msg string, format ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", format...)
	os.Exit(1)
}
