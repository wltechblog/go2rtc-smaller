package hls

import (
	"io"
	"net/url"

	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/mpegts"
)

func OpenURL(u *url.URL, body io.ReadCloser) (core.Producer, error) {
	rd, err := NewReader(u, body)
	if err != nil {
		return nil, err
	}
	return mpegts.Open(rd)
}
