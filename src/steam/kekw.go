package steam

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	gameListSelector = "div#sharedfiles_filterselect_app_filterable div"
)

func GetUserGameList(userID string) []SteamGame {
	var games []SteamGame

	path := buildPath(userID, false)

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

// "javascript:SelectSharedFilesContentFilter({ 'appid': '1368030' });"
func parseGameID(onClickText string) string {
	onclickRegex, _ := regexp.Compile("\\d+")
	gameID := onclickRegex.FindString(onClickText)
	return strings.Trim(gameID, " ")
}

func buildPath(userID string, isFile bool) string {
	if isFile {
		path, err := filepath.Abs("./spec/html/game-index.html")
		if err != nil {
			panic(err)
		}
		return "file://" + path
	}

	return "https://steamcommunity.com/id/" + userID + "/screenshots/?appid=0&sort=newestfirst&browsefilter=myfiles&view=grid#scrollTop=0"
}

func BuildCollector(path string) *colly.Collector {
	c := colly.NewCollector()
	if strings.Contains(path, "file://") {
		t := &http.Transport{}
		t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
		c.WithTransport(t)
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting -> ", r.URL)
	})

	return c
}
