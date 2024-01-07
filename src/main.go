package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"steam-screenshot-cli/src/commands"
	"time"
)

func main() {
	app := &cli.App{
		Name:     "🚰 Steam Screenshot download cli tool",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "BigForceGun",
				Email: "bigforcegun@pm.me",
			},
		},
		Usage:     "demonstrate available API",
		UsageText: "contrive - demonstrating the available API",
		Commands: []*cli.Command{
			{
				Name:     "user",
				Category: "list",
				Aliases:  []string{"lu"},
				Usage:    "Sync all user screenshots to output folder",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "force", Aliases: []string{"f"}},
					&cli.StringFlag{Name: "game-id", Aliases: []string{"g"}},
				},
				Action: func(cCtx *cli.Context) error {
					userID := cCtx.Args().First()
					commands.ListUserScreenshots(userID)
					return nil
				},
			},
			{
				Name:     "games",
				Category: "list",
				Aliases:  []string{"lgs"},
				Usage:    "Sync all user screenshots to output folder",
				Action: func(cCtx *cli.Context) error {
					userID := cCtx.Args().First()
					commands.ListUserGames(userID)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
