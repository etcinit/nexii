package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/etcinit/nexii/util"
)

// InfoCommand lists information about the application associated with the
// provided token.
func InfoCommand(c *cli.Context) {
	client, err := util.MakeClient(c)

	if err != nil {
		fmt.Println(err)
		return
	}

	response, errs := client.Fetch()

	if errs != nil {
		fmt.Println(errs)
		return
	}

	fmt.Println("ID:", response.Application.ID)
	fmt.Println("Name:", response.Application.Name)
	fmt.Println("Description:", response.Application.Description)
}
