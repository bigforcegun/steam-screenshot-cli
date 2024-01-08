package commands

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"steam-screenshot-cli/src/steam"
)

func ListUserGames(userID string) {
	games := steam.GetUserGames(userID)
	printGameTable(games)
}

func printGameTable(games []steam.SteamGame) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "ID")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, game := range games {
		tbl.AddRow(game.Name, game.ID)
	}

	tbl.Print()

}
