package ivideon

import (
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/ivideon"
	"strings"
)

func Init() {
	streams.HandleFunc("ivideon", func(url string) (core.Producer, error) {
		id := strings.Replace(url[8:], "/", ":", 1)
		prod := ivideon.NewClient(id)
		if err := prod.Dial(); err != nil {
			return nil, err
		}
		return prod, nil
	})
}
