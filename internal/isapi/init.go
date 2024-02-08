package isapi

import (
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/isapi"
)

func Init() {
	streams.HandleFunc("isapi", handle)
}

func handle(url string) (core.Producer, error) {
	conn, err := isapi.NewClient(url)
	if err != nil {
		return nil, err
	}
	if err = conn.Dial(); err != nil {
		return nil, err
	}
	return conn, nil
}
