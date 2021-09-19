package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinpollet/go-docker-multi-stage-build/pkg"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "greet",
		Usage: "greets the given name",
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				name = "World"
			}

			fmt.Println(pkg.Greet(name))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}