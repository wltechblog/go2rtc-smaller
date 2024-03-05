package unix

import (
	"net"
	"sync"

	"github.com/rs/zerolog"
	"github.com/wltechblog/go2rtc-smaller/internal/app"
	"github.com/wltechblog/go2rtc-smaller/internal/rtsp"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/magic"
	pkg "github.com/wltechblog/go2rtc-smaller/pkg/rtsp"
	"github.com/wltechblog/go2rtc-smaller/pkg/shell"
)

func Init() {
	rtsp.HandleFunc(func(conn *pkg.Conn) bool {
		waitersMu.Lock()
		waiter := waiters[conn.URL.Path]
		waitersMu.Unlock()

		if waiter == nil {
			return false
		}

		// unblocking write to channel
		select {
		case waiter <- conn:
			return true
		default:
			return false
		}
	})

	streams.HandleFunc("unix", unixHandle)

	log = app.GetLogger("unix")
}

func unixHandle(url string) (core.Producer, error) {

	args := shell.QuoteSplit(url[5:]) // remove `unix:`
	filename := args[0]

	return handleUnix(filename)
}

func handleUnix(filename string) (core.Producer, error) {

	sock,err:=net.Dial("unix", filename)
	if err != nil {
		log.Printf("Error opening unix socket %s: %v", filename, err)
		return nil, err
	}



	prod, err := magic.Open(sock)
	if err != nil {
		_ = sock.Close()
	}

	return prod, err
}

// internal

var log zerolog.Logger
var waiters = map[string]chan core.Producer{}
var waitersMu sync.Mutex
