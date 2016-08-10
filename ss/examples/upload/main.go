// This upload example illustrates how to upload a CAT to the self-service designer service.
// The reference for the API can be found at http://reference.rightscale.com/selfservice/manager/index.html#/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rightscale/rsc/httpclient"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssd"
)

// For testing
var osStdout io.Writer = os.Stdout

func main() {
	// 1. Retrieve login and endpoint information
	email := flag.String("e", "", "Login email")
	pwd := flag.String("p", "", "Login password")
	account := flag.Int("a", 0, "Account id")
	host := flag.String("h", "us-3.rightscale.com", "RightScale API host")
	insecure := flag.Bool("insecure", false, "Bypass certificate check, used for testing")
	cat := flag.String("cat", "", "Path to CAT file")
	flag.Parse()
	if *email == "" {
		fail("Login email required")
	}
	if *pwd == "" {
		fail("Login password required")
	}
	if *account == 0 {
		fail("Account id required")
	}
	if *host == "" {
		fail("Host required")
	}
	if *insecure {
		httpclient.NoCertCheck = true
	}
	if *cat == "" {
		fail("Path to CAT required")
	}
	if _, err := os.Stat(*cat); err != nil {
		fail("Can't find file %s", *cat)
	}

	// 2. Setup client using basic auth
	auth := rsapi.NewBasicAuthenticator(*email, *pwd, *account)
	ssAuth := rsapi.NewSSAuthenticator(auth, *account)
	client := ssd.New(*host, ssAuth)
	if err := client.CanAuthenticate(); err != nil {
		fail("invalid credentials: %s", err)
	}

	// 3. Make execution index call using expanded view
	file, err := os.Open(*cat)
	if err != nil {
		fail("Could not open %s", *cat)
	}
	name := filepath.Base(*cat)
	l := client.TemplateLocator(fmt.Sprintf("/api/designer/collections/%d/templates", *account))
	upload := rsapi.FileUpload{Name: "source", Filename: name, Reader: file}
	t, err := l.Create(&upload)
	if err != nil {
		fail("failed to create template: %s", err)
	}
	created, err := t.Show(rsapi.APIParams{})
	if err != nil {
		fail("failed to retrieve created template: %s", err)
	}
	fmt.Fprintf(osStdout, "%s created at %s by %s\n",
		created.Name, created.Timestamps.CreatedAt, created.CreatedBy.Name)
}

// Print error message and exit with code 1
// Make it overridable for testing
var fail = func(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Println(fmt.Sprintf(format, v...))
	os.Exit(1)
}
