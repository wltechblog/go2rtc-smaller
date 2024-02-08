package hls

import (
	"errors"
	"time"

	"github.com/wltechblog/go2rtc-smaller/internal/api"
	"github.com/wltechblog/go2rtc-smaller/internal/api/ws"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/mp4"
	"github.com/wltechblog/go2rtc-smaller/pkg/tcp"
)

func handlerWSHLS(tr *ws.Transport, msg *ws.Message) error {
	stream := streams.GetOrPatch(tr.Request.URL.Query())
	if stream == nil {
		return errors.New(api.StreamNotFound)
	}

	codecs := msg.String()
	medias := mp4.ParseCodecs(codecs, true)
	cons := mp4.NewConsumer(medias)
	cons.Type = "HLS/fMP4 consumer"
	cons.RemoteAddr = tcp.RemoteAddr(tr.Request)
	cons.UserAgent = tr.Request.UserAgent()

	log.Trace().Msgf("[hls] new ws consumer codecs=%s", codecs)

	if err := stream.AddConsumer(cons); err != nil {
		log.Error().Err(err).Caller().Send()
		return err
	}

	session := NewSession(cons)

	session.alive = time.AfterFunc(keepalive, func() {
		sessionsMu.Lock()
		delete(sessions, session.id)
		sessionsMu.Unlock()

		stream.RemoveConsumer(cons)
	})

	sessionsMu.Lock()
	sessions[session.id] = session
	sessionsMu.Unlock()

	go session.Run()

	main := session.Main()
	tr.Write(&ws.Message{Type: "hls", Value: string(main)})

	return nil
}
