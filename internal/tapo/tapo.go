package tapo

import (
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/kasa"
	"github.com/wltechblog/go2rtc-smaller/pkg/tapo"
)

func Init() {
	streams.HandleFunc("kasa", func(url string) (core.Producer, error) {
		return kasa.Dial(url)
	})

	streams.HandleFunc("tapo", func(url string) (core.Producer, error) {
		return tapo.Dial(url)
	})
}
