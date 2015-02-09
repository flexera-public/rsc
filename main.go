package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/raphael/kingpin"
	"github.com/rightscale/rsclient/rsapi15"
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
	var config *ClientConfig
	err = json.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	config.Token, err = decrypt(config.Token)
	return config, err
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

	account := app.Flag("account", "RightScale account ID, overrides config").Int()
	endpoint := app.Flag("endpoint", "RightScale API endpoint (e.g. 'us-3.rightscale.com'), overrides config").String()
	token := app.Flag("token", "Oauth access token (overrides config)").String()

	pretty := app.Flag("pretty", "Pretty print response body").Short('p').Bool()
	extract := app.Flag("extract", "Extract value(s) from response media type at given path, path consists of dot separated attribute names, e.g. \"security_groups.href\"").Short('e').String()

	setup := app.Command("setup",
		"setup config file, defaults to $HOME/.rsclient, use '--config' to override")
	cfgPath := setup.Flag("config", "path to rsclient config file").Default(fmt.Sprintf("%v/.rsclient", os.Getenv("HOME"))).String()

	rsapi15.RegisterCommands(app)

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	if cmd == "setup" {
		createConfig(*cfgPath)
	} else {
		config, err := LoadConfig(*cfgPath)
		if account != nil {
			config.Account = *account
		}
		if endpoint != nil {
			config.Endpoint = *endpoint
		}
		if token != nil {
			config.Token = *token
		}
		client, err := rsapi15.NewClient(config.Token, config.Account, config.Endpoint, 0)
		if err != nil {
			PrintError("Failed to create API session: %v", err.Error())
		} else {
			res := client.RunCommand(cmd)
			pp := pretty != nil && *pretty
			if extract != nil {
				paths := strings.Split(extract, ".")
				for res != nil && len(paths) > 0 {
					res = res[paths[0]]
					paths = paths[1:]
				}
				if res == nil {
					PrintError("No value at path '%s', response was:", extract)
					pp = true // Force pretty print
				}
			}
			var err error
			var output string
			if pp {
				output, err := json.MarshalIndent(res, "", "  ")
			} else {
				output, err := json.Marshal(res)
			}
			if err != nil {
				PrintError("Failed to JSON encode response\n%+v", res)
				return
			}
			fmt.Println(output)
		}
	}
}

// Create configuration file
func createConfig(path string) {
	config, _ := LoadConfig(path)
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

	_, err := rsapi15.NewClient(config.Token, config.Account, config.Endpoint, 0)
	if err != nil {
		PrintError("Failed to contact RightScale: %v", err.Error())
	} else {
		PrintSuccess("rsclient configured successfully!")
	}
}
