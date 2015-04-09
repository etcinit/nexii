package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/etcinit/nexii/util"
)

// ListCommand lists all the available keys
func ListCommand(c *cli.Context) {
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

	fmt.Printf("Received %d keys:\n", len(response.Values))

	for keyName := range response.Values {
		fmt.Println(keyName)
	}
}
