package main

import (
	"code.google.com/p/gopass"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rightscale/rsclient/display"
	"github.com/rightscale/rsclient/rsapi15"
	"io/ioutil"
	"log"
	"os"
)

// Command line client entry point.
func main() {
	app := cli.NewApp()
	app.Name = "rsclient"
	app.Usage = "RightScale API client"

	app.Commands = []cli.Command{
		{
			Name:  "setup",
			Usage: "setup config file, defaults to $HOME/.rsclient, use '--config' to override",
			Description: `Build a rsclient config file and test it. Multiple config files can be created by specifying
   different values for the --config flag.`,
			Flags:  []cli.Flag{cfgFlag(), dumpFlag(), dbgFlag()},
			Action: createConfig,
		}, {
			Name:  "test",
			Usage: "test connectivity to RightScale API endpoint",
			Description: `First login using credential stored in specified config file (~/.rsclient by default) then perform
   a request to RightScale APIs`,
			Flags:  []cli.Flag{cfgFlag(), dumpFlag(), dbgFlag()},
			Action: testConfig,
		}, {
			Name:        "api15",
			Usage:       "use RightScale API 1.5",
			Description: `Setup client to use API 1.5 - http://reference.rightscale.com/api1.5/index.html`,
			Subcommands: rsapi15.Commands(),
		},
	}

	app.Run(os.Args)
}

func cfgFlag() cli.StringFlag {
	return cli.StringFlag{"config, c", fmt.Sprintf("%v/.rswind", os.Getenv("HOME")), "path to rsclient config file", ""}
}

func dumpFlag() cli.BoolFlag {
	return cli.BoolFlag{"dump_request_response, dump", "print HTTP request and response", ""}
}

func dbgFlag() cli.BoolFlag {
	return cli.BoolFlag{"debug, dbg", "print debug information", ""}
}

// Global logger
var logger *log.Logger = log.New(os.Stdout, "[rsclient]", 0)

// Read existing configuration file
func readConfig(path string) (map[string]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config map[string]string
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	if config["password"] == "" {
		return nil, errors.New("missing password")
	}
	config["password"], err = decrypt(config["password"])
	return config, err
}

// Create configuration file
func createConfig(c *cli.Context) {
	path := c.String("config")
	if path == "" {
		path = fmt.Sprintf("%v/.rsclient", os.Getenv("HOME"))
	}
	config, err := readConfig(path)
	var user, userDef, pwd, loginEndpoint, account, accountDef string
	if config != nil {
		display.PromptWarning("Found existing configuration file %v, overwrite? (y/N): ", path)
		var yn string
		fmt.Scanf("%s", &yn)
		if yn != "y" {
			display.PrintSuccess("Exiting")
			return
		} else {
			user = config["user"]
			userDef = fmt.Sprintf(" (%v)", user)
			account = config["account"]
			accountDef = fmt.Sprintf(" (%v)", account)
			loginEndpoint = config["loginEndpoint"]
		}
	}

	fmt.Printf("API Username%v: ", userDef)
	var newUser string
	fmt.Scanf("%s", &newUser)
	if newUser != "" {
		user = newUser
	}

	pwd, err = gopass.GetPass("Password: ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Account id%v: ", accountDef)
	var newAccount string
	fmt.Scanf("%s", &newAccount)
	if newAccount != "" {
		account = newAccount
	}

	password, err := encrypt(pwd)
	if err != nil {
		display.PrintError("!!Failed to encrypt password: %v", err)
		return
	}
	config = map[string]string{
		"user":     user,
		"password": password,
		"account":  account,
	}
	if loginEndpoint != "" {
		config["loginEndpoint"] = loginEndpoint
	}
	bytes, err := json.Marshal(config)
	if err != nil {
		display.PrintError("!!Failed to serialize config: %v", err)
		return
	}
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		display.PrintError("!!Failed to write config file: %v", err)
		return
	}
	testConfig(c)
}

func testConfig(c *cli.Context) {
	display.PrintTitle("Testing connectivity...")
	client, err := createClient(c)
	if err != nil {
		return
	}
	display.PrintSuccess("Logged in successfully against %v, now contacting RightScale...",
		client.LoginEndpoint())
	_, err = client.ListClouds()
	if err != nil {
		display.PrintError("!!Failed to contact RightScale: %v", err.Error())
		return
	} else {
		display.PrintSuccess("rsclient configured successfully!\n")
		return
	}
}

func createClient(c *cli.Context) (rsapi15.Client, error) {
	path := c.String("config")
	if path == "" {
		path = os.Getenv("HOME") + "/.rsclient"
	}
	config, err := readConfig(path)
	if err != nil {
		display.PrintError("!!Failed to read rsclient config, please run 'rsclient setup'\n")
		return nil, err
	}
	dump := c.Bool("dump_request_response")
	client := rsapi15.NewClient(config["user"], config["password"], config["account"],
		config["login"], dump)
	if err != nil {
		display.PrintError("!!Failed to login (%v)\n", err)
		return nil, err
	}
	return client, nil
}
