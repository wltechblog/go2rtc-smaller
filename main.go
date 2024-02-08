package main

import (
	"github.com/AlexxIT/go2rtc/internal/api"
	"github.com/AlexxIT/go2rtc/internal/api/ws"
	"github.com/AlexxIT/go2rtc/internal/app"
	"github.com/AlexxIT/go2rtc/internal/debug"
	"github.com/AlexxIT/go2rtc/internal/exec"
	"github.com/AlexxIT/go2rtc/internal/hls"
	"github.com/AlexxIT/go2rtc/internal/mjpeg"
	"github.com/AlexxIT/go2rtc/internal/mp4"
	"github.com/AlexxIT/go2rtc/internal/ngrok"
	"github.com/AlexxIT/go2rtc/internal/rtsp"
	"github.com/AlexxIT/go2rtc/internal/streams"
	"github.com/AlexxIT/go2rtc/internal/webrtc"
	"github.com/AlexxIT/go2rtc/pkg/shell"
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

	ngrok.Init() // ngrok module
//	srtp.Init()  // SRTP server
	debug.Init() // debug API

	// 7. Go

	shell.RunUntilSignal()
}
