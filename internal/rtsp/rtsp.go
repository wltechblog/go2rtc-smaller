package rtsp

import (
	"io"
	"net"
	"net/url"

	"github.com/rs/zerolog"
	"github.com/wltechblog/go2rtc-smaller/internal/app"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/rtsp"
	"github.com/wltechblog/go2rtc-smaller/pkg/tcp"
)

func Init() {
	var conf struct {
		Mod struct {
			Listen       string `yaml:"listen" json:"listen"`
			Username     string `yaml:"username" json:"-"`
			Password     string `yaml:"password" json:"-"`
			DefaultQuery string `yaml:"default_query" json:"default_query"`
			PacketSize   uint16 `yaml:"pkt_size"`
		} `yaml:"rtsp"`
	}

	// default config
	conf.Mod.Listen = ":8554"
	conf.Mod.DefaultQuery = "video&audio"

	app.LoadConfig(&conf)
	app.Info["rtsp"] = conf.Mod

	log = app.GetLogger("rtsp")

	// RTSP server support
	address := conf.Mod.Listen
	if address == "" {
		return
	}

	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error().Err(err).Msg("[rtsp] listen")
		return
	}

	_, Port, _ = net.SplitHostPort(address)

	log.Info().Str("addr", address).Msg("[rtsp] listen")

	if query, err := url.ParseQuery(conf.Mod.DefaultQuery); err == nil {
		defaultMedias = ParseQuery(query)
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}

			c := rtsp.NewServer(conn)
			c.PacketSize = conf.Mod.PacketSize
			// skip check auth for localhost
			if conf.Mod.Username != "" && !conn.RemoteAddr().(*net.TCPAddr).IP.IsLoopback() {
				c.Auth(conf.Mod.Username, conf.Mod.Password)
			}
			go tcpHandler(c)
		}
	}()
}

type Handler func(conn *rtsp.Conn) bool

func HandleFunc(handler Handler) {
	handlers = append(handlers, handler)
}

var Port string

// internal

var log zerolog.Logger
var handlers []Handler
var defaultMedias []*core.Media


func tcpHandler(conn *rtsp.Conn) {
	var name string
	var closer func()

	trace := log.Trace().Enabled()

	conn.Listen(func(msg any) {
		if trace {
			switch msg := msg.(type) {
			case *tcp.Request:
				log.Trace().Msgf("[rtsp] server request:\n%s", msg)
			case *tcp.Response:
				log.Trace().Msgf("[rtsp] server response:\n%s", msg)
			}
		}

		switch msg {
		case rtsp.MethodDescribe:
			if len(conn.URL.Path) == 0 {
				log.Warn().Msg("[rtsp] server empty URL on DESCRIBE")
				return
			}

			name = conn.URL.Path[1:]

			stream := streams.Get(name)
			if stream == nil {
				return
			}

			log.Debug().Str("stream", name).Msg("[rtsp] new consumer")

			conn.SessionName = app.UserAgent

			query := conn.URL.Query()
			conn.Medias = ParseQuery(query)
			if conn.Medias == nil {
				for _, media := range defaultMedias {
					conn.Medias = append(conn.Medias, media.Clone())
				}
			}

			if s := query.Get("pkt_size"); s != "" {
				conn.PacketSize = uint16(core.Atoi(s))
			}

			if err := stream.AddConsumer(conn); err != nil {
				log.Warn().Err(err).Str("stream", name).Msg("[rtsp]")
				return
			}

			closer = func() {
				stream.RemoveConsumer(conn)
			}

		case rtsp.MethodAnnounce:
			if len(conn.URL.Path) == 0 {
				log.Warn().Msg("[rtsp] server empty URL on ANNOUNCE")
				return
			}

			name = conn.URL.Path[1:]

			stream := streams.Get(name)
			if stream == nil {
				return
			}

			log.Debug().Str("stream", name).Msg("[rtsp] new producer")

			stream.AddProducer(conn)

			closer = func() {
				stream.RemoveProducer(conn)
			}
		}
	})

	if err := conn.Accept(); err != nil {
		if err != io.EOF {
			log.Warn().Err(err).Caller().Send()
		}
		if closer != nil {
			closer()
		}
		_ = conn.Close()
		return
	}

	for _, handler := range handlers {
		if handler(conn) {
			return
		}
	}

	if closer != nil {
		if err := conn.Handle(); err != nil {
			log.Debug().Msgf("[rtsp] handle=%s", err)
		}

		closer()

		log.Debug().Str("stream", name).Msg("[rtsp] disconnect")
	}

	_ = conn.Close()
}

func ParseQuery(query map[string][]string) []*core.Media {
	if v := query["mp4"]; v != nil {
		return []*core.Media{
			{
				Kind:      core.KindVideo,
				Direction: core.DirectionSendonly,
				Codecs: []*core.Codec{
					{Name: core.CodecH264},
					{Name: core.CodecH265},
				},
			},
			{
				Kind:      core.KindAudio,
				Direction: core.DirectionSendonly,
				Codecs: []*core.Codec{
					{Name: core.CodecAAC},
				},
			},
		}
	}

	return core.ParseQuery(query)
}
