package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/rightscale/rsclient/rsapi15"
)

// Command line client entry point.
func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "rsclient"
	app.Usage = "RightScale API client"
	app.Author = "RightScale"
	app.Email = "support@rightscale.com"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name: "setup",
			Usage: "setup config file, defaults to $HOME/.rsclient, use '--config' " +
				"to override",
			Description: "Build a rsclient config file and test it. Multiple config " +
				"files can be created by specifying different values for the " +
				"--config flag.",
			Flags: []cli.Flag{
				cli.StringFlag{"config, c", fmt.Sprintf("%v/.rsclient",
					os.Getenv("HOME")), "path to rsclient config file", ""}},
			Action: createConfig,
		}, {
			Name:  "api15",
			Usage: "use RightScale API 1.5",
			Description: "Setup client to use API 1.5 " +
				"- http://reference.rightscale.com/api1.5/index.html",
			Subcommands: rsapi15.Commands(),
		},
	}

	app.Run(os.Args)
}

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
	if config["token"] == "" {
		return nil, errors.New("missing refresh token")
	}
	config["token"], err = decrypt(config["token"])
	return config, err
}

// Create configuration file
func createConfig(c *cli.Context) {
	path := c.String("config")
	if path == "" {
		path = fmt.Sprintf("%v/.rsclient", os.Getenv("HOME"))
	}
	config, err := readConfig(path)
	var token, account, accountDef string
	if config != nil {
		PromptWarning("Found existing configuration file %v, overwrite? (y/N): ", path)
		var yn string
		fmt.Scanf("%s", &yn)
		if yn != "y" {
			PrintSuccess("Exiting")
			return
		} else {
			account = config["account"]
			accountDef = fmt.Sprintf(" (%v)", account)
			token = config["token"]
			tokenDef = " (leave blank to leave unchanged)"
		}
	}

	fmt.Printf("Account id%v: ", accountDef)
	var newAccount string
	fmt.Scanf("%s", &newAccount)
	if newAccount != "" {
		account = newAccount
	}

	fmt.Printf("Refresh Token%v: ", tokenDef)
	var newToken string
	fmt.Scanf("%s", &newToken)
	if newToken != "" {
		token = newToken
	}

	token, err := encrypt(token)
	if err != nil {
		PrintError("Failed to encrypt token: %v", err)
		return
	}
	config = map[string]string{
		"account": account,
		"token":   token,
	}
	bytes, err := json.Marshal(config)
	if err != nil {
		PrintError("Failed to serialize config: %v", err)
		return
	}
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		PrintError("Failed to write config file: %v", err)
		return
	}
	client, err := rsapi15.NewClient()
	if err != nil {
		PrintError("Failed to contact RightScale: %v", err.Error())
	} else {
		PrintSuccess("rsclient configured successfully!")
	}
}
