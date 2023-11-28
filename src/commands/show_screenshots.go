package commands

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log/slog"
	"net/http"
	"path/filepath"
	"steam-screenshot-cli/src/steam"
)

func ShowScreenshots(userID string) {
	/*
		ну карочи блеть

		зашли сюда - https://steamcommunity.com/id/bigforcegun/screenshots/#scrollTop=0
		получили скринты
		https://steamcommunity.com/sharedfiles/filedetails/?id=3095228393 - зашли по ID файла
		получили файл и мета инфу
		сохранили
		...
		профит

		вариант 2
		достать по API список игор
		https://steamcommunity.com/id/bigforcegun/screenshots/?appid=620#scrollTop=25
		по appid ходить по игорам и выкачивать скрины

		API не дает инфу по скринам, только по установленым игорам
	*/

	//url := "https://steamcommunity.com/id/bigforcegun/screenshots/?p=1&appid=539470&sort=newestfirst&browsefilter=myfiles&view=grid&privacy=30"

	//parseGamePageV1(url)
	parseFilePageV1()

	fmt.Println("go go to screenshots", userID)
}

func parseGamePageV1(url string) {
	// pages := []string{}
	var pages []string

	file_name, err := filepath.Abs("./spec/html/game-index.html")
	if err != nil {
		panic(err)
	}

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(t)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e)
		// e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("a.profile_media_item.modalContentLink", func(e *colly.HTMLElement) {
		fmt.Println("_________")
		steamFileUrl := e.Attr("href")
		slog.Info("🗄", "href", steamFileUrl)
		slog.Info("📝", "text", e.Text)
		fileID := steam.GetFileIDFromUrl(steamFileUrl)
		pages = append(pages, fileID)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//c.OnHTML("a", func(e *colly.HTMLElement) {
	//	c.Visit("file://" + dir + "/html" + e.Attr("href"))
	//})

	c.Visit("file://" + file_name)
	c.Wait()
	for i, p := range pages {
		fmt.Printf("%d : %s\n", i, p)
	}
	//c.Visit(url)
}

func parseFilePageV1() {

	var steamFile steam.SteamFile
	file_name, err := filepath.Abs("./spec/html/file-index.html")
	if err != nil {
		panic(err)
	}

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(t)

	c.OnHTML(".screenshotAppName", func(e *colly.HTMLElement) {
		steamFile.GameName = e.Text
	})

	c.OnHTML(".detailsStatRight", func(e *colly.HTMLElement) {
		steamFile.CreatedAt = e.Text
	})

	c.OnHTML("img.screenshotEnlargeable", func(e *colly.HTMLElement) {
		steamFile.PictureURL = e.Attr("src")
	})

	c.Visit("file://" + file_name)

	c.Wait()

	slog.Info(":🗄 file:", steamFile)
}
