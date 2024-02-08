package expr

import (
	"errors"

	"github.com/wltechblog/go2rtc-smaller/internal/app"
	"github.com/wltechblog/go2rtc-smaller/internal/streams"
	"github.com/wltechblog/go2rtc-smaller/pkg/expr"
)

func Init() {
	log := app.GetLogger("expr")

	streams.RedirectFunc("expr", func(url string) (string, error) {
		v, err := expr.Run(url[5:])
		if err != nil {
			return "", err
		}

		log.Debug().Msgf("[expr] url=%s", url)

		if url = v.(string); url == "" {
			return "", errors.New("expr: result is empty")
		}

		return url, nil
	})
}
