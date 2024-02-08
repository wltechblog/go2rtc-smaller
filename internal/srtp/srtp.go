package srtp

import (
	"github.com/wltechblog/go2rtc-smaller/internal/app"
	"github.com/wltechblog/go2rtc-smaller/pkg/srtp"
)

func Init() {
	var cfg struct {
		Mod struct {
			Listen string `yaml:"listen"`
		} `yaml:"srtp"`
	}

	// default config
	cfg.Mod.Listen = ":8443"

	// load config from YAML
	app.LoadConfig(&cfg)

	if cfg.Mod.Listen == "" {
		return
	}

	// create SRTP server (endpoint) for receiving video from HomeKit cameras
	Server = srtp.NewServer(cfg.Mod.Listen)
}

var Server *srtp.Server
