package steam

import (
	"github.com/gocolly/colly/v2"
	"log/slog"
)

const (
	tableScreenshotSelector = "div#sharedfiles_filterselect_app_filterable div"
	tableFileLinkSelector   = "a.profile_media_item.modalContentLink"
)

func GetGameScreenshots(userID string, gameID string) []SteamTableScreenshot {
	var screenshots []SteamTableScreenshot

	path := buildGamePath(userID, gameID, false)

	c := BuildCollector(path)

	c.OnHTML(tableFileLinkSelector, func(e *colly.HTMLElement) {
		screenshot := SteamTableScreenshot{}
		screenshot.URL = e.Attr("href")
		screenshot.ID = GetFileIDFromUrl(screenshot.URL)
		slog.Debug("🕸parsed table screenshot", "file", screenshot)
		screenshots = append(screenshots, screenshot)
	})

	c.Visit(path)
	c.Wait()
	return screenshots
}
