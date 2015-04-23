// This example illustrates how to use CM 1.5 client to build a bash aliases file
// to login to various systems on various environments. The reference for the API can be found at
// http://reference.rightscale.com/api1.5/index.html.
package main // import "gopkg.in/rightscale/rsc.v1/cm15/examples/rsssh"

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strings"

	"gopkg.in/rightscale/rsc.v1/cm15"
	"gopkg.in/rightscale/rsc.v1/rsapi"
)

// Top level configuration
type Config struct {
	SshOptions   string                       // SSH options to use in the ssh command in the built aliases
	SshUser      string                       // The SSH user name. Default is rightscale
	OutputFile   string                       // The output file location. By default rightscale_aliases
	Environments map[string]EnvironmentDetail // Environment Detail
}

// Detail of a particular environment
type EnvironmentDetail struct {
	Account      int               // RightScale account ID
	ServerArrays map[string]string // Details about Server Arrays
	Servers      map[string]string // Details about Servers
}

// The SSH configuration used to build the aliases
type SshConfig struct {
	Name      string
	IpAddress string
}

func main() {
	// 1. Retrieve login and endpoint information
	email := flag.String("e", "", "Login email")
	pwd := flag.String("p", "", "Login password")
	host := flag.String("h", "us-3.rightscale.com", "RightScale API host")
	conf := flag.String("c", "", "Configuration file")
	flag.Parse()
	if *email == "" {
		fail("Login email required")
	}
	if *pwd == "" {
		fail("Login password required")
	}
	if *host == "" {
		fail("Host required")
	}
	if *conf == "" {
		usr, _ := user.Current()
		*conf = usr.HomeDir + "/.rsssh"
	}

	// Read the configuration
	content, err := ioutil.ReadFile(*conf)
	if err != nil {
		fail("Failed to read configuration file: %v\n", err.Error())
	}
	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		fail("Failed to Unmarshal configuration file: %v\n", err.Error())
	}
	if config.OutputFile == "" {
		config.OutputFile = "rightscale_aliases"
	}
	if config.SshUser == "" {
		config.SshUser = "rightscale"
	}
	logger := log.New(os.Stdout, "", 0)
	var sshConfig = make([]SshConfig, 0)

	for envName, envDetail := range config.Environments {
		if envDetail.Account == 0 {
			fail("No account specified for environment: %v\n", envName)
		}
		// Create a new client for every environment because they can be in different accounts.
		auth := rsapi.LoginAuthenticator{Username: *email, Password: *pwd, Client: http.DefaultClient}
		client, err := cm15.New(envDetail.Account, *host, &auth, logger, nil)
		if err != nil {
			fail("failed to create client: %s", err)
		}
		fetchDetails(client, envName, envDetail, &sshConfig)
	}

	aliases := buildAliases(sshConfig, config.SshOptions, config.SshUser)
	ioutil.WriteFile(config.OutputFile, []byte(aliases), 0644)
}

// Fetch details about all servers and server arrays in an environment
func fetchDetails(client *cm15.Api, envName string, envDetail EnvironmentDetail, sshConfig *[]SshConfig) {
	for nickname, name := range envDetail.ServerArrays {
		// Obtain the resource
		instances := serverArray(client, name)
		// Obtain the IP address of the first instance (only one instance is considered here -- for now)
		for _, instance := range instances {
			ipAddress := instance.PublicIpAddresses[0]
			number := getInstanceNumber(instance.Name)
			*sshConfig = append(*sshConfig, SshConfig{Name: envName + "_" + nickname + number, IpAddress: ipAddress})
		}
	}
	for nickname, name := range envDetail.Servers {
		instance := server(client, name)
		ipAddress := instance.PublicIpAddresses[0]
		*sshConfig = append(*sshConfig, SshConfig{Name: envName + "_" + nickname, IpAddress: ipAddress})
	}
}

// Get the instance number of an instance in a server array. This assumes that the server array instance is not renamed
// from the default name. For example, if the instance name is "App Server #4", this will return "#4".
func getInstanceNumber(name string) string {
	re, _ := regexp.Compile(`#\d+$`)
	matches := re.FindStringSubmatch(name)
	if len(matches) == 0 {
		return ""
	} else {
		return matches[len(matches)-1]
	}
}

// Builds the aliases string based on the SSH configuration of all servers and server arrays in all environments.
func buildAliases(sshConfig []SshConfig, sshOptions, sshUser string) string {
	var aliases string
	for _, conf := range sshConfig {
		aliases = aliases + fmt.Sprintf("alias %v='ssh %v %v@%v'\n", conf.Name, sshOptions, sshUser, conf.IpAddress)
	}
	return aliases
}

// Makes a GET call on the given server array and returns all its current instances.
func serverArray(client *cm15.Api, name string) []*cm15.Instance {
	serverArrayLocator := client.ServerArrayLocator("/api/server_arrays")
	serverArrays, err := serverArrayLocator.Index(rsapi.ApiParams{"view": "default", "filter": []string{"name==" + name}})
	if err != nil {
		fail("Failed to retrieve server array: %v\n", err.Error())
	}
	if len(serverArrays) == 0 {
		fail("Could not find server array with name: %v\n", name)
	} else if len(serverArrays) != 1 {
		fail("More than one server array found with name: %v\n", name)
	}
	array := serverArrays[0]
	var instancesHref string
	for _, l := range array.Links {
		if l["rel"] == "current_instances" {
			instancesHref = l["href"]
			break
		}
	}
	instanceLocator := client.InstanceLocator(instancesHref)
	instances, err := instanceLocator.Index(rsapi.ApiParams{})
	if err != nil {
		fail("Failed to retrieve current instances of the server array: %v\n", err.Error())
	}
	if len(instances) == 0 {
		fail("No instances found in server array: %v\n", name)
	}
	return instances
}

// Makes a GET call on the given server and returns the current instance of the server.
func server(client *cm15.Api, name string) *cm15.Instance {
	serverLocator := client.ServerLocator("/api/servers")
	servers, err := serverLocator.Index(rsapi.ApiParams{"view": "instance_detail", "filter": []string{"name==" + name}})
	if err != nil {
		fail("Failed to retrieve server: %v\n", err.Error())
	}
	if len(servers) == 0 {
		fail("Could not find server with name: %v\n", name)
	} else if len(servers) != 1 {
		fail("More than one server found with name: %v\n", name)
	}
	return servers[0].CurrentInstance
}

// Print error message and exit with code 1
func fail(format string, v ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Println(fmt.Sprintf(format, v...))
	os.Exit(1)
}
