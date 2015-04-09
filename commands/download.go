package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/codegangsta/cli"
	"github.com/etcinit/nexii/util"
)

// DownloadCommand saves the value of a key as file with the same name
func DownloadCommand(c *cli.Context) {
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

	key := c.Args().First()
	filename := key

	if c.String("output") != "" {
		filename = c.String("output")
	}

	if val, ok := response.Values[key]; ok {
		// write whole the body
		err = ioutil.WriteFile(filename, []byte(val), 0644)
		if err != nil {
			panic(err)
		}
		return
	}

	fmt.Println("Key not found or provided")
}
