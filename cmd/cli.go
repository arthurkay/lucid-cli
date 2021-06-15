package cmd

import (
	"github.com/urfave/cli/v2"
)

var Commands []*cli.Command

func init() {
	Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Creates lucid-cli framework boilerplate",
			UsageText: `Creating the framework's boiler plate template without giving a directory will initialise the application code in the current directory.
			If you need to define your own path, you'll need to pass in the -d flag`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "directory",
					Aliases: []string{"d"},
				},
			},
			Action: func(c *cli.Context) error {
				InitProject(c)
				return nil
			},
		},
	}
}
