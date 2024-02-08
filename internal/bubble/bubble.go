package bubble

import (
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/bubble"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
)

func Init() {
	streams.HandleFunc("bubble", handle)
}

func handle(url string) (core.Producer, error) {
	conn := bubble.NewClient(url)
	if err := conn.Dial(); err != nil {
		return nil, err
	}
	return conn, nil
}
