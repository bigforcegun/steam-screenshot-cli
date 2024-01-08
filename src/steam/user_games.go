package steam

import (
	"github.com/gocolly/colly/v2"
)

const (
	gameListSelector = "div#sharedfiles_filterselect_app_filterable div"
)

func GetUserGames(userID string) []SteamGame {
	var games []SteamGame

	path := buildUserPath(userID, false)

	c := BuildCollector(path)

	c.OnHTML(gameListSelector, func(e *colly.HTMLElement) {
		steamGame := SteamGame{}
		onClickText := e.Attr("onclick")
		steamGame.Name = e.Text
		steamGame.ID = parseGameID(onClickText)
		games = append(games, steamGame)
	})

	c.Visit(path)
	c.Wait()
	return games
}
