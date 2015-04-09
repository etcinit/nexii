package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/etcinit/nexii/util"
)

// PingCommand sends a ping to the Nexus server with this machine's name and a
// message.
func PingCommand(c *cli.Context) {
	client, err := util.MakeClient(c)

	if err != nil {
		fmt.Println(err)
		return
	}

	errs := client.Ping(c.String("name"), c.String("message"))

	if errs != nil {
		fmt.Println(errs)
		return
	}
}
