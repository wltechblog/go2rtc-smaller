package main

import (
	"github.com/wltechblog/go2rtc-smaller/internal/app"
	"github.com/wltechblog/go2rtc-smaller/internal/rtsp"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	rtsp.Init()

	shell.RunUntilSignal()
}
