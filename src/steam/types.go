package steam

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type SteamFile struct {
	ID        string
	GameName  string
	CreatedAt string
	ImageUrl  string
}

func (f SteamFile) FileCreatedAt() time.Time {
	timeString := f.CreatedAt
	theTime, err := time.Parse("15 Jan @ 12:52am", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	fmt.Println("The time is", theTime)
	return theTime
	//return time.Now().Local()
}

func (f SteamFile) FileName() string {
	//return path.Base(f.ImageUrl)
	u, err := url.Parse(f.ImageUrl)
	if err != nil {
		panic(err)
	}
	paths := strings.Split(u.Path, "/")
	realPath := paths[len(paths)-2]
	return realPath + ".png"
}
