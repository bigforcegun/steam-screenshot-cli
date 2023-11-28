package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"steam-screenshot-cli/src/commands"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "ss",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					userID := cCtx.Args().First()
					commands.ShowScreenshots(userID)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
