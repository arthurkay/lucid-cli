package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func InitProject(c *cli.Context) {
	if c.String("directory") != "" {
		// Create project in the chosen directory
		err := os.Mkdir(c.String("directory"), 0775)
		if err != nil {
			fmt.Printf("%v", err)
		}
	} else {
		// Create the project in current directory
		fmt.Printf("Initialise the app in the current directory")
	}
}
