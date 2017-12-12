package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/gboddin/ghcli/repo"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghcli"
	app.Usage = "Github command line interface"
	app.Commands = []cli.Command{
		{
			Name:        "repositories",
			Aliases:     []string{"repo","r"},
			Usage:       "commands for repositories",
			Subcommands: []cli.Command{
				repo.ListCmd,
			},
		},
	}
	app.Run(os.Args)
}

