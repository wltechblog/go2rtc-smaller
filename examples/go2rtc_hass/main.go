package main

import (
	"github.com/wltechblog/go2rtc-smaller/internal/api"
	"github.com/wltechblog/go2rtc-smaller/internal/app"
	"github.com/wltechblog/go2rtc-smaller/internal/hass"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	api.Init()

	hass.Init()

	shell.RunUntilSignal()
}
