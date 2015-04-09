package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/etcinit/nexii/util"
)

// GetCommand outputs the value of a key
func GetCommand(c *cli.Context) {
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

	if val, ok := response.Values[c.Args().First()]; ok {
		fmt.Print(val)
		return
	}

	fmt.Println("Key not found or provided")
}
