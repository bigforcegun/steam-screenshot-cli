package steam

import (
	"github.com/gocolly/colly/v2"
	"log/slog"
	"net/http"
	"strings"
)

func BuildCollector(path string) *colly.Collector {
	c := colly.NewCollector()
	if strings.Contains(path, "file://") {
		t := &http.Transport{}
		t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
		c.WithTransport(t)
	}

	c.OnRequest(func(r *colly.Request) {
		slog.Debug("🌐", "Visiting -> ", r.URL)
	})

	return c
}
