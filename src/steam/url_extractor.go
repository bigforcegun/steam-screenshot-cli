package steam

import (
	"fmt"
	"net/url"
)

func GetFileIDFromUrl(urlStr string) string {
	myUrl, _ := url.Parse(urlStr)
	params, _ := url.ParseQuery(myUrl.RawQuery)
	fmt.Println(params)
	fileID := params.Get("id")
	return fileID
}
