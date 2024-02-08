package mpegts

import (
	"net/http"

	"github.com/wltechblog/go2rtc-smaller/internal/api"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/aac"
	"github.com/wltechblog/go2rtc-smaller/pkg/tcp"
	"github.com/rs/zerolog/log"
)

func apiStreamAAC(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Query().Get("src")
	stream := streams.Get(src)
	if stream == nil {
		http.Error(w, api.StreamNotFound, http.StatusNotFound)
		return
	}

	cons := aac.NewConsumer()
	cons.RemoteAddr = tcp.RemoteAddr(r)
	cons.UserAgent = r.UserAgent()

	if err := stream.AddConsumer(cons); err != nil {
		log.Error().Err(err).Caller().Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "audio/aac")

	_, _ = cons.WriteTo(w)

	stream.RemoveConsumer(cons)
}
