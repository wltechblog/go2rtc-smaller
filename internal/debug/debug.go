package debug

import (
	"github.com/wltechblog/go2rtc-smaller/internal/api"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
)

func Init() {
	api.HandleFunc("api/stack", stackHandler)

	streams.HandleFunc("null", nullHandler)
}

func nullHandler(string) (core.Producer, error) {
	return nil, nil
}
