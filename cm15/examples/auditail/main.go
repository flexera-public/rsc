// This example illustrates how to setup a CM 1.5 client to tail the audit entry
// summary. The reference for the API can be found at
// http://reference.rightscale.com/api1.5/index.html.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/rsapi"
)

func main() {
	// 1. Retrieve login and endpoint information
	email := flag.String("e", "", "Login email")
	pwd := flag.String("p", "", "Login password")
	account := flag.Int("a", 0, "Account id")
	host := flag.String("h", "us-3.rightscale.com", "RightScale API host")
	filterEmail := flag.String("f", "", "Audit user email for filtering audit entries")
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
	if *filterEmail == "" {
		*filterEmail = *email
	}

	// 2. Setup client using basic auth
	auth := rsapi.NewBasicAuthenticator(*email, *pwd, *account)
	client := cm15.New(*host, auth)
	if err := client.CanAuthenticate(); err != nil {
		fail("invalid credentials: %s", err)
	}
	ticker := time.NewTicker(time.Second)

	// 3. Make an initial API call and retrieve the audit entries
	entries, err := fetchAuditEntries(client, *filterEmail)
	oldEntries := entries
	if err != nil {
		fail("Failed to retrieve audit entries: %v\n", err.Error())
	}
	printAudits(entries)

	// 4. Periodically retrieve audit entries and print the new ones
	for {
		select {
		case <-ticker.C:
			entries, err := fetchAuditEntries(client, *filterEmail)
			if err != nil {
				fail("Failed to retrieve audit entries: %v\n", err.Error())
			}
			printAudits(extractUnique(oldEntries, entries))
			oldEntries = entries
		}
	}
}

// Make an API call and fetch the audit entries matching specified criteria
func fetchAuditEntries(client *cm15.API, filterEmail string) ([]*cm15.AuditEntry, error) {
	auditLocator := client.AuditEntryLocator("/api/audit_entries")
	var apiParams = rsapi.APIParams{"filter": []string{"user_email==" + filterEmail}}
	auditEntries, err := auditLocator.Index(
		tomorrow(),  // End date
		"100",       // Limit
		yesterday(), // Start date
		apiParams,
	)
	if err != nil {
		return auditEntries, err
	}
	return auditEntries, nil
}

// Returns tommorrow's date in the format used by RightScale API
func tomorrow() string {
	return formatTime(time.Now().UTC().AddDate(0, 0, 1))
}

// Returns yesterday's date in the format used by RightScale API
func yesterday() string {
	return formatTime(time.Now().UTC().AddDate(0, 0, -1))
}

// Returns time in RightScale API supported format
func formatTime(tm time.Time) string {
	year, month, date := tm.Date()
	return time.Date(year, month, date, 0, 0, 0, 0, time.UTC).Format("2006/01/02 15:04:05 -0700")
}

// Prints the audit entries to console
func printAudits(entries []*cm15.AuditEntry) {
	for _, a := range entries {
		fmt.Printf("[%v] <%v>: %v\n", a.UpdatedAt, a.UserEmail, a.Summary)
	}
}

// Extract unique audit entries from the newly received list by comparing the href of audit entries
// in the old list.
func extractUnique(oldEntries, newEntries []*cm15.AuditEntry) []*cm15.AuditEntry {
	var uniqueEntries = make([]*cm15.AuditEntry, 0)
	var oldHrefs = make([]string, len(oldEntries))
	for i, e := range oldEntries {
		oldHrefs[i] = getHref(e)
	}
	for _, newEntry := range newEntries {
		if !stringInSlice(getHref(newEntry), oldHrefs) {
			uniqueEntries = append(uniqueEntries, newEntry)
		}
	}
	return uniqueEntries
}

// Get the href of an audit entry from the Links attribute by inspecting the self link
func getHref(entry *cm15.AuditEntry) string {
	var href string
	for _, link := range entry.Links {
		if link["rel"] == "self" {
			href = link["href"]
			break
		}
	}
	return href
}

// Find if a string exists in a slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Print error message and exit with code 1
func fail(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Println(fmt.Sprintf(format, v...))
	os.Exit(1)
}
