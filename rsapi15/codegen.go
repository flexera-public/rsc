package rsapi15

import (
	"github.com/codegangsta/cli"
)

func (c *Client) ListClouds() ([]Cloud, error) {
	return []Cloud{}, nil
}

func (c *Client) CreateOauth2(p *Params) (interface{}, error) {
	return nil, nil
}

// command line tool subcommands
func Commands(c *Client) []cli.Command {
	return []cli.Command{
		{
			Name:  "cloud",
			Usage: "...",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "...",
					Action: func(co *cli.Context) {
						c.ListClouds()
					},
				},
			},
		},
	}
}

type Cloud struct{}
type Instance struct{}
