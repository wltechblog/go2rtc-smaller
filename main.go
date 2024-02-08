package main

import (
	"github.com/wltechblog/go2rtc-smaller/internal/api"
	"github.com/wltechblog/go2rtc-smaller/internal/api/ws"
	"github.com/wltechblog/go2rtc-smaller/internal/app"
	"github.com/wltechblog/go2rtc-smaller/internal/exec"
	"github.com/wltechblog/go2rtc-smaller/internal/hls"
	"github.com/wltechblog/go2rtc-smaller/internal/mjpeg"
	"github.com/wltechblog/go2rtc-smaller/internal/mp4"
	"github.com/wltechblog/go2rtc-smaller/internal/rtsp"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/internal/webrtc"
	"github.com/wltechblog/go2rtc-smaller/pkg/shell"
)

func main() {
	// 1. Core modules: app, api/ws, streams

	app.Init() // init config and logs

	api.Init() // init API before all others
	ws.Init()  // init WS API endpoint

	streams.Init() // streams module

	// 2. Main sources and servers

	rtsp.Init()   // rtsp source, RTSP server
	webrtc.Init() // webrtc source, WebRTC server

	// 3. Main API

	mp4.Init()   // MP4 API
	hls.Init()   // HLS API
	mjpeg.Init() // MJPEG API

	// 4. Other sources and servers

//	hass.Init()       // hass source, Hass API server
//	onvif.Init()      // onvif source, ONVIF API server
//	webtorrent.Init() // webtorrent source, WebTorrent module

	// 5. Other sources

//	rtmp.Init()     // rtmp source
	exec.Init()     // exec source
//	ffmpeg.Init()   // ffmpeg source
//	echo.Init()     // echo source
//	ivideon.Init()  // ivideon source
//	http.Init()     // http/tcp source
//	dvrip.Init()    // dvrip source
//	tapo.Init()     // tapo source
//	isapi.Init()    // isapi source
//	mpegts.Init()   // mpegts passive source
//	roborock.Init() // roborock source
//	homekit.Init()  // homekit source
//	nest.Init()     // nest source
//	expr.Init()     // expr source
//	gopro.Init()    // gopro source

	// 6. Helper modules

//	ngrok.Init() // ngrok module
//	srtp.Init()  // SRTP server
//	debug.Init() // debug API

	// 7. Go

	shell.RunUntilSignal()
}
