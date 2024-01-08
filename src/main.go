package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"log/slog"
	"os"
	"steam-screenshot-cli/src/commands"
	"strconv"
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
		//Usage:     "Show and download you open steam screenshots",
		UsageText: "Show and download you open steam screenshots",
		Commands: []*cli.Command{
			{
				Name:     "games",
				Category: "Games",
				Aliases:  []string{"g"},
				Usage:    "List all user games that have screenshots",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "user",
						Aliases:  []string{"u"},
						Usage:    "Specify public steam user id like in url https://steamcommunity.com/id/`USERID`/screenshots/",
						Required: true,
						Action: func(ctx *cli.Context, v string) error {
							if len(v) <= 4 {
								return fmt.Errorf("flag --user value %v must be more than 4 letters", v)
							}
							return nil
						},
					},
				},
				Action: func(cCtx *cli.Context) error {
					//userID := cCtx.Args().First()
					userID := cCtx.String("user")

					if userID != "" {
						slog.Info("👨‍💻 userID provided", userID)
						commands.ListUserGames(userID)

					} else {
						slog.Error("🤷‍♂️ no userID provided - stopping...")
					}
					return nil
				},
			},
			{
				Name:     "list-game",
				Category: "Games",
				Aliases:  []string{"lg"},
				Usage:    "Sync all user screenshots to output folder",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "user",
						Aliases:  []string{"u"},
						Usage:    "Specify public steam user id like in url https://steamcommunity.com/id/`USERID`/screenshots/",
						Required: true,
						Action: func(ctx *cli.Context, v string) error {
							if len(v) <= 4 {
								return fmt.Errorf("flag --user value %v must be more than 4 letters", v)
							}
							return nil
						},
					},
					&cli.StringFlag{
						Name:     "game",
						Aliases:  []string{"g"},
						Usage:    "Specify public steam gameID like in url https://steamcommunity.com/id/USERID/screenshots/?appid=`GAMEID`; ",
						Required: true,
						Action: func(ctx *cli.Context, v string) error {
							_, err := strconv.Atoi(v)
							if err != nil {
								return fmt.Errorf("flag --game value %v must be integer", v)
							}
							return nil
						},
					},
				},
				Action: func(cCtx *cli.Context) error {
					userID := cCtx.String("user")
					gameID := cCtx.String("game")

					commands.ListGameScreenshots(userID, gameID)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// fuck need 1.22 golang
// https://github.com/golang/go/commit/3188758653fc7d2b229e234273d41878ddfdd5f2
func setSlogLevel() {
	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)
}
