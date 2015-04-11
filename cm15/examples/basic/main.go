// This basic example illustrates how to setup a CM 1.5 client to make a simple
// API call. The reference for the API can be found at
// http://reference.rightscale.com/api1.5/index.html.
package main  // import "gopkg.in/rightscale/rsc.v1-unstable/cm15/examples/basic"

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/rightscale/rsc.v1-unstable/cm15"
	"gopkg.in/rightscale/rsc.v1-unstable/rsapi"
)

func main() {
	// 1. Retrieve login and endpoint information
	email := flag.String("e", "", "Login email")
	pwd := flag.String("p", "", "Login password")
	account := flag.Int("a", 0, "Account id")
	host := flag.String("h", "us-3.rightscale.com", "RightScale API host")
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
	logger := log.New(os.Stdout, "", 0)
	auth := rsapi.LoginAuthenticator{Username: *email, Password: *pwd, Client: http.DefaultClient}
	client, err := cm15.New(*account, *host, &auth, logger, nil)
	if err != nil {
		fail("failed to create client: %s", err)
	}

	// 3. Make cloud index call using extended view
	l := client.CloudLocator("/api/clouds")
	clouds, err := l.Index(rsapi.ApiParams{"view": "extended"})
	if err != nil {
		fail("failed to list clouds: %s", err)
	}

	// 4. Print cloud capabilities
	for i, c := range clouds {
		fmt.Println("")
		fmt.Printf("%d. %s\n", i+1, c.Name)
		for _, ca := range c.Capabilities {
			fmt.Printf("%s: %v\n", ca["name"], ca["value"])
		}
	}
}

// Print error message and exit with code 1
func fail(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Println(fmt.Sprintf(format, v...))
	os.Exit(1)
}
