package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// Basic configuration settings required by all clients
type ClientConfig struct {
	Account int    // RightScale account ID
	Host    string // RightScale API host, e.g. "us-3.rightscale.com"
	Token   string // RightScale API refresh token
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
	config.Token, err = Decrypt(config.Token)
	return &config, err
}

// Save config encrypts the token and persists the config to file
func (cfg *ClientConfig) Save(path string) error {
	encrypted, err := Encrypt(cfg.Token)
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
func CreateConfig(path string) error {
	config, _ := LoadConfig(path)
	var tokenDef, accountDef, hostDef string
	if config != nil {
		yn := PromptConfirmation("Found existing configuration file %v, overwrite? (y/N): ", path)
		if yn != "y" {
			PrintSuccess("Exiting")
			return nil
		}
		accountDef = fmt.Sprintf(" (%v)", config.Account)
		tokenDef = " (leave blank to leave unchanged)"
		if config.Host == "" {
			config.Host = "my.rightscale.com"
		}
		hostDef = fmt.Sprintf(" (%v)", config.Host)
	} else {
		config = &ClientConfig{}
	}

	fmt.Fprintf(out, "Account id%v: ", accountDef)
	var newAccount string
	fmt.Fscanln(in, &newAccount)
	if newAccount != "" {
		a, err := strconv.Atoi(newAccount)
		if err != nil {
			return fmt.Errorf("Account ID must be an integer, got '%s'.", newAccount)
		}
		config.Account = a
	}

	fmt.Fprintf(out, "Refresh Token%v: ", tokenDef)
	var newToken string
	fmt.Fscanln(in, &newToken)
	if newToken != "" {
		config.Token = newToken
	}

	fmt.Fprintf(out, "API Host%v: ", hostDef)
	var newHost string
	fmt.Fscanln(in, &newHost)
	if newHost != "" {
		config.Host = newHost
	}

	err := config.Save(path)
	if err != nil {
		return fmt.Errorf("Failed to save config: %s", err.Error())
	}

	return nil
}
