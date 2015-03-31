package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/rightscale/rsc/recording"
)
import "os/exec"

var (
	output  = kingpin.Flag("output", "path to generated file").Short('o').Default("recording.json").String()
	key     = kingpin.Flag("key", "API key").Short('k').String()
	account = kingpin.Flag("account", "RightScale account").Short('a').Default("60073").Int()
	host    = kingpin.Flag("host", "RightScale host").Short('h').Default("us-3.rightscale.com").String()

	exitRegexp = regexp.MustCompile(`exit status (\d+)`)
)

func main() {
	kingpin.Parse()
	k := *key
	if k == "" {
		k = os.Getenv("RS_KEY")
	}
	if k == "" {
		kingpin.Fatalf("Please set RS_KEY to your API key")
	}
	os.RemoveAll(*output)
	write("[\n")
	for i, t := range TestCases {
		args := makeArgs(t, k)
		cmd := exec.Command("rsc", args...)
		// Remove key, dump and host from args so it doesn't get printed or recorded
		args = args[3:]
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
					kingpin.Fatalf("Invalid exit code '%s': %s", exCode[1], err.Error())
				}
			} else {
				kingpin.Fatalf("%s %s failed: %s\n%s\n", "rsc", strings.Join(args, " "), err, out.String())
			}
		}
		var rr recording.RequestResponse
		kingpin.FatalIfError(json.Unmarshal(outErr.Bytes(), &rr), "load rsc output")
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
		js, err := json.Marshal(record)
		if err != nil {
			kingpin.Fatalf("failed to serialize record: %s", err)
		}
		write(string(js))
		if i < len(TestCases)-1 {
			write(",\n")
		}
	}
	write("]\n")
}

// Helper function that appends a string to output file
func write(s string) {
	f, err := os.OpenFile(*output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	kingpin.FatalIfError(err, "failed to open output file")
	f.WriteString(s)
	f.Close()
}

// Build array of "rsc" command line arguments
func makeArgs(cmd, key string) []string {
	keyArg := fmt.Sprintf("--key=%s", key)
	hostArg := fmt.Sprintf("--host=%s", *host)
	accountArg := fmt.Sprintf("--account=%d", *account)
	raw := []string{keyArg, "--dump=json", hostArg, accountArg}
	return append(raw, strings.Split(cmd, " ")...)
}
