package util

import (
	"errors"
	"os"

	"github.com/codegangsta/cli"
	"github.com/etcinit/nexus-client-go"
)

// MakeClient builds a client using the CLI context, with environment vars
// as a fallback
func MakeClient(c *cli.Context) (*nexus.Client, error) {
	endpoint := os.Getenv("NEXUS_SERVER")
	token := os.Getenv("NEXUS_APIKEY")

	if c.String("endpoint") != "" {
		endpoint = c.String("endpoint")
	}
	if c.String("token") != "" {
		token = c.String("token")
	}

	if endpoint == "" || token == "" {
		return nil, errors.New("Missing client config. Make sure environment variables are set or arguments are provided.")
	}

	return nexus.NewClient(endpoint, token), nil
}
