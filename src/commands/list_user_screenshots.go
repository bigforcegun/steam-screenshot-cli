package commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/rodaine/table"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"steam-screenshot-cli/src/steam"
)

func ListUserScreenshots(userID string) {
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

func ListUserGames(userID string) {
	games := steam.GetUserGameList(userID)
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

	c.OnHTML(".detailsStatsContainerRight", func(e *colly.HTMLElement) {
		ch := e.DOM.Children()
		date := ch.Eq(1).Text()
		//steamFile.CreatedAt = e.Text
		steamFile.CreatedAt = date
	})

	c.OnHTML("img.screenshotEnlargeable", func(e *colly.HTMLElement) {
		steamFile.ImageUrl = e.Attr("src")
	})

	c.Visit("file://" + file_name)

	c.Wait()

	slog.Info(":🗄 file:", steamFile)

	saveSteamFile(steamFile)

}

func saveSteamFile(sfile steam.SteamFile) {
	var sourceDirectory = "./"
	var gameDirectory = sfile.GameName
	var target = filepath.Join(sourceDirectory, gameDirectory)
	var realFile = downloadFileToDirectory(sfile.ImageUrl, target, sfile)
	updateFileMeta(*realFile, sfile)
}

func downloadFileToDirectory(url string, target string, sfile steam.SteamFile) *os.File {

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	fileName := sfile.FileName()
	newFilePath := filepath.Join(target, fileName)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.MkdirAll(target, 0755) // FIXME: как бы тут дефолт создавать с umask
	}
	file, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
	return file
}

func updateFileMeta(file os.File, sfile steam.SteamFile) {
	createdTime := sfile.FileCreatedAt()
	// currentTime := time.Now().Local()
	//Set both access time and modified time of the file to the current time
	err := os.Chtimes(file.Name(), createdTime, createdTime)
	if err != nil {
		fmt.Println(err)
	}
}
