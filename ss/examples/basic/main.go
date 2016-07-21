// This basic example illustrates how to setup a Self-Service client to make a simple
// API call. The reference for the API can be found at
// http://reference.rightscale.com/selfservice/manager/index.html#/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/rightscale/rsc/httpclient"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss"
	"github.com/rightscale/rsc/ss/ssm"
)

// For testing
var osStdout io.Writer = os.Stdout

func main() {
	// 1. Retrieve login and endpoint information
	email := flag.String("e", "", "Login email")
	pwd := flag.String("p", "", "Login password")
	account := flag.Int("a", 0, "Account id")
	host := flag.String("h", "us-3.rightscale.com", "RightScale API host")
	insecure := flag.Bool("insecure", false, "Use HTTP instead of HTTPS - used for testing")
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

	// 2. Setup client using basic auth
	auth := rsapi.NewBasicAuthenticator(*email, *pwd, *account)
	ssAuth := rsapi.NewSSAuthenticator(auth, *account)
	client := ssm.New(*host, ssAuth)
	if *insecure {
		httpclient.Insecure = true
	}
	if err := client.CanAuthenticate(); err != nil {
		fail("invalid credentials: %s", err)
	}

	// 3. Make execution index call using expanded view
	l := client.ExecutionLocator(fmt.Sprintf("/api/manager/projects/%d/executions", *account))
	execs, err := l.Index(rsapi.APIParams{})
	if err != nil {
		fail("failed to list executions: %s", err)
	}
	sort.Sort(ByName(execs))

	// 4. Print executions launch from
	w := new(tabwriter.Writer)
	w.Init(osStdout, 5, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Name\tState\tBy\tLink")
	fmt.Fprintln(w, "-----\t-----\t-----\t-----")
	for _, e := range execs {
		link := fmt.Sprintf("https://%s/api/manager/projects/#/executions/%s", ss.HostFromLogin(client.Host),
			e.Id)
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%s", e.Name, e.Status, e.CreatedBy.Email, link))
	}
	w.Flush()
}

// ByName makes it possible to sort executions by name.
type ByName []*ssm.Execution

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }

// Print error message and exit with code 1
// Make it overridable for testing
var fail = func(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Println(fmt.Sprintf(format, v...))
	os.Exit(1)
}
