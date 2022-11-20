package config

import (
	"strings"

	"github.com/mhkarimi1383/simple-store/types"
)

var config *types.Config

func SetConfig(cfg *types.Config) {
	cfg.BasePath = strings.TrimSuffix(cfg.BasePath, "/")
	config = cfg
}

func GetConfig() types.Config {
	return *config
}
