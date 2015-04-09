package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/etcinit/nexii/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "nexii"
	app.Usage = "A Nexus Configuration Server CLI client"
	app.Author = "Eduardo Trujillo <ed@chromabits.com>"
	app.Version = "0.1.0"
	app.Action = func(c *cli.Context) {
		println("See 'nexii help' for more information")
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "endpoint",
			Value:  "",
			Usage:  "Nexus server endpoint",
			EnvVar: "NEXUS_SERVER,NEXUS_ENDPOINT",
		},
		cli.StringFlag{
			Name:   "token",
			Value:  "",
			Usage:  "Nexus server API token",
			EnvVar: "NEXUS_APIKEY,NEXUS_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List all keys",
			Action:  commands.ListCommand,
		},
		{
			Name:    "get",
			Aliases: []string{"g", "cat"},
			Usage:   "Get and output a specific key",
			Action:  commands.GetCommand,
		},
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "Get and save a specific key",
			Action:  commands.DownloadCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Value: "",
					Usage: "Override output file name",
				},
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i", "app"},
			Usage:   "Get information about the application",
			Action:  commands.InfoCommand,
		},
	}

	app.Run(os.Args)
}
