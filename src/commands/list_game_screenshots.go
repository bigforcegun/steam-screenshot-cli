package commands

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"steam-screenshot-cli/src/steam"
)

func ListGameScreenshots(userID string, gameID string) {
	screenshots := steam.GetGameScreenshots(userID, gameID)
	//TODO: обработка пустого списка и хуевой игры / пользователя / ответа / whatewer
	printScreenshotsTable(screenshots)
}

func printScreenshotsTable(screenshots []steam.SteamTableScreenshot) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "URL")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, screenshot := range screenshots {
		tbl.AddRow(screenshot.ID, screenshot.URL)
	}

	tbl.Print()

}
