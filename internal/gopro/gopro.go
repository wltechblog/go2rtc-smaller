package gopro

import (
	"net/http"

	"github.com/wltechblog/go2rtc-smaller/internal/api"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/gopro"
)

func Init() {
	streams.HandleFunc("gopro", handleGoPro)

	api.HandleFunc("api/gopro", apiGoPro)
}

func handleGoPro(rawURL string) (core.Producer, error) {
	return gopro.Dial(rawURL)
}

func apiGoPro(w http.ResponseWriter, r *http.Request) {
	var items []*api.Source

	for _, host := range gopro.Discovery() {
		items = append(items, &api.Source{Name: host, URL: "gopro://" + host})
	}

	api.ResponseSources(w, items)
}
