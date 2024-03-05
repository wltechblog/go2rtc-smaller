package magic

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"

	"github.com/wltechblog/go2rtc-smaller/pkg/core"
	"github.com/wltechblog/go2rtc-smaller/pkg/h264/annexb"
	"github.com/wltechblog/go2rtc-smaller/pkg/magic/bitstream"
	"github.com/wltechblog/go2rtc-smaller/pkg/magic/mjpeg"
)

func Open(r io.Reader) (core.Producer, error) {
	rd := core.NewReadBuffer(r)

	b, err := rd.Peek(4)
	if err != nil {
		return nil, err
	}

	switch {
	case bytes.HasPrefix(b, []byte(annexb.StartCode)):
		return bitstream.Open(rd)

	case bytes.HasPrefix(b, []byte{0xFF, 0xD8}):
		return mjpeg.Open(rd)
/*
	case bytes.HasPrefix(b, []byte(flv.Signature)):
		return flv.Open(rd)

	case bytes.HasPrefix(b, []byte{0xFF, 0xF1}):
		return aac.Open(rd)

	case bytes.HasPrefix(b, []byte("--")):
		return multipart.Open(rd)

	case b[0] == mpegts.SyncByte:
		return mpegts.Open(rd)
*/
	}

	// support MJPEG with trash on start
	// https://github.com/wltechblog/go2rtc-smaller/issues/747
	if b, err = rd.Peek(4096); err != nil {
		return nil, err
	}

	if i := bytes.Index(b, []byte{0xFF, 0xD8, 0xFF, 0xDB}); i > 0 {
		_, _ = io.ReadFull(rd, make([]byte, i))
		return mjpeg.Open(rd)
	}

	return nil, errors.New("magic: unsupported header: " + hex.EncodeToString(b[:4]))
}
