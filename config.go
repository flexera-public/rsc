package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

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

// Create configuration file and save it to file at given path
func createConfig(path string) error {
	var config, _ = LoadConfig(path)
	var tokenDef, accountDef, endpointDef string
	if config != nil {
		PromptWarning("Found existing configuration file %v, overwrite? (y/N): ", path)
		var yn string
		fmt.Scanf("%s", &yn)
		if yn != "y" {
			PrintSuccess("Exiting")
			return nil
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
			return fmt.Errorf("Account ID must be an integer, got '%s'.", newAccount)
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

	_, err := rsapi15.New(config.Account, config.Token, config.Endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("Config test failed: %s", err.Error())
	}

	return nil
}
