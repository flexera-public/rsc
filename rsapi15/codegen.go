package rsapi15

import (
	"github.com/codegangsta/cli"
	"github.com/rightscale/rsclient/display"
)

type Client interface {
	LoginEndpoint() string
	ListClouds() ([]Cloud, error)
}

type client struct {
	user     string
	password string
	account  string
	login    string
	dump     bool
}

func NewClient(user, pwd, account, login string, dump bool) *client {
	return &client{
		user:     user,
		password: pwd,
		account:  account,
		login:    login,
		dump:     dump,
	}
}

func (c *client) LoginEndpoint() string {
	return c.login
}

func (c *client) ListClouds() ([]Cloud, error) {
	return []Cloud{}, nil
}

// command line tool subcommands
func Commands() []cli.Command {
	return []cli.Command{}
}

func DisplayResource(res interface{}) {
	switch r := res.(type) {
	case *Instance:
		DisplayInstance(r)
	default:
		display.PrintError("Unknown resource type %T", res)
	}
}

type Cloud struct{}
type Instance struct{}

func DisplayInstance(i *Instance) {}
