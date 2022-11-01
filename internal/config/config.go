package config

import (
	"github.com/mhkarimi1383/simple-store/types"
	log "github.com/sirupsen/logrus"
	"strings"
)

var config *types.Config

func SetConfig(cfg *types.Config) {
	if strings.HasSuffix(cfg.BasePath, "/") {
		cfg.BasePath = strings.TrimSuffix(cfg.BasePath, "/")
	}
	config = cfg
}

func GetConfig() types.Config {
	log.Printf("%+v", config)
	return *config
}
