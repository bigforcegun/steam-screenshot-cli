package steam

import "path/filepath"

func buildGamePath(userID string, gameID string, isFile bool) string {
	if isFile {
		path, err := filepath.Abs("./spec/html/game-index.html")
		if err != nil {
			panic(err)
		}
		return "file://" + path
	}

	return "https://steamcommunity.com/id/" + userID + "/screenshots/?appid=0&sort=newestfirst&browsefilter=myfiles&view=grid#scrollTop=0"
}

func buildUserPath(userID string, isFile bool) string {
	if isFile {
		path, err := filepath.Abs("./spec/html/game-index.html")
		if err != nil {
			panic(err)
		}
		return "file://" + path
	}

	return "https://steamcommunity.com/id/" + userID + "/screenshots/?appid=0&sort=newestfirst&browsefilter=myfiles&view=grid#scrollTop=0"
}
