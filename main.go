package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/rightscale/rsclient/rsapi15"
	"gopkg.in/alecthomas/kingpin.v1"
)

// Basic configuration settings required by all clients
type ClientConfig struct {
	Account  int    // RightScale account ID
	Endpoint string // RightScale API endpoint, e.g. "us-3.rightscale.com"
	Token    string // RightScale API refresh token
}

// LoadConfig loads the client configuration from disk
func LoadConfig(path string) (*ClientConfig, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config ClientConfig
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	config.Token, err = decrypt(config.Token)
	return &config, err
}

// Save config encrypts the token and persists the config to file
func (cfg *ClientConfig) Save(path string) error {
	encrypted, err := encrypt(cfg.Token)
	if err != nil {
		return fmt.Errorf("Failed to encrypt token: %s", err.Error())
	}
	cfg.Token = encrypted
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Failed to serialize config: %s", err.Error())
	}
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write config file: %s", err.Error())
	}
	return nil
}

// Command line client entry point.
func main() {
	app := kingpin.New("rsclient", "A RightScale API client")
	app.Version("0.1.0")

	var cfgPathFlag = app.Flag("config", "path to rsclient config file").Default(path.Join(os.Getenv("HOME"), ".rsclient")).String()
	var accountFlag = app.Flag("account", "RightScale account ID").Int()
	var endpointFlag = app.Flag("endpoint", "RightScale API endpoint (e.g. 'us-3.rightscale.com')").String()
	var tokenFlag = app.Flag("token", "OAuth access token").String()

	var prettyFlag = app.Flag("pretty", "Pretty print response body").Short('p').Bool()
	//extract := app.Flag("extract", "Extract value(s) from response media type at given path, path consists of dot separated attribute names, e.g. \"security_groups.href\"").Short('e').String()

	app.Command("setup",
		"create config file, defaults to $HOME/.rsclient, use '--config' to override")

	rsapi15.RegisterCommands(app)

	var cmd = kingpin.MustParse(app.Parse(os.Args[1:]))

	if cmd == "setup" {
		createConfig(*cfgPathFlag)
	} else {
		var account int
		var token, endpoint string
		if config, err := LoadConfig(*cfgPathFlag); err == nil {
			account = config.Account
			token = config.Token
			endpoint = config.Endpoint
		}
		if *accountFlag > 0 {
			account = *accountFlag
		}
		if *tokenFlag != "" {
			token = *tokenFlag
		}
		if *endpointFlag != "" {
			endpoint = *endpointFlag
		}
		if token == "" {
			PrintError("Missing OAuth token, use '-token TOKEN' or 'setup'")
		}
		client, err := rsapi15.New(account, token, endpoint, nil, nil)
		if err != nil {
			PrintError("Failed to create API session: %v", err.Error())
		} else {
			res := client.RunCommand(cmd)
			pp := prettyFlag != nil && *prettyFlag
			var err error
			var output []byte
			if pp {
				output, err = json.MarshalIndent(res, "", "  ")
			} else {
				output, err = json.Marshal(res)
			}
			if err != nil {
				PrintError("Failed to JSON encode response\n%+v", res)
				return
			}
			fmt.Println(string(output))
		}
	}
}

// Create configuration file
func createConfig(path string) {
	var config, err = LoadConfig(path)
	var tokenDef, accountDef, endpointDef string
	if config != nil {
		PromptWarning("Found existing configuration file %v, overwrite? (y/N): ", path)
		var yn string
		fmt.Scanf("%s", &yn)
		if yn != "y" {
			PrintSuccess("Exiting")
			return
		}
		accountDef = fmt.Sprintf(" (%v)", config.Account)
		tokenDef = " (leave blank to leave unchanged)"
		if config.Endpoint == "" {
			config.Endpoint = "my.rightscale.com"
		}
		endpointDef = fmt.Sprintf(" (%v)", config.Endpoint)
	} else {
		config = &ClientConfig{}
	}

	fmt.Printf("Account id%v: ", accountDef)
	var newAccount string
	fmt.Scanf("%s", &newAccount)
	if newAccount != "" {
		a, err := strconv.Atoi(newAccount)
		if err != nil {
			PrintError("Account ID must be an integer, exiting.")
			return
		}
		config.Account = a
	}

	fmt.Printf("Refresh Token%v: ", tokenDef)
	var newToken string
	fmt.Scanf("%s", &newToken)
	if newToken != "" {
		config.Token = newToken
	}

	fmt.Printf("API Endpoint%v: ", endpointDef)
	var newEndpoint string
	fmt.Scanf("%s", &newEndpoint)
	if newEndpoint != "" {
		config.Endpoint = newEndpoint
	}

	config.Save(path)

	_, err = rsapi15.New(config.Account, config.Token, config.Endpoint, nil, nil)
	if err != nil {
		PrintError("Failed to contact RightScale: %v", err.Error())
	} else {
		PrintSuccess("rsclient configured successfully!")
	}
}
