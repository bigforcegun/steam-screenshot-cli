package steam

import (
	"net/url"
	"regexp"
	"strings"
)

func GetFileIDFromUrl(urlStr string) string {
	myUrl, _ := url.Parse(urlStr)
	params, _ := url.ParseQuery(myUrl.RawQuery)
	fileID := params.Get("id")
	return fileID
}

// "javascript:SelectSharedFilesContentFilter({ 'appid': '1368030' });"
func parseGameID(onClickText string) string {
	onclickRegex, _ := regexp.Compile("\\d+")
	gameID := onclickRegex.FindString(onClickText)
	return strings.Trim(gameID, " ")
}
