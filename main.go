package main

import (
	"log"
	"os"

	"github.com/arthurkay/lucid-cli/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	lucid := &cli.App{
		Name:        "lucid-cli CLI",
		Usage:       "lucid-cli Framework",
		Version:     "v0.0.0-alpha",
		HideVersion: true,
		Flags:       []cli.Flag{},
		Commands:    cmd.Commands,
		Authors: []*cli.Author{
			{
				Name:  "Arthur Kalikiti",
				Email: "arthur@kalikiti.net",
			},
		},
	}

	err := lucid.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to start lucid-cli CLI because, %v", err)
	}
}
