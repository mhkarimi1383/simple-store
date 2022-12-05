// Package config holds a singletone pointer of the configuration
package config

import (
	"strings"

	"github.com/mhkarimi1383/simple-store/types"
)

var config *types.Config

// SetConfig sets config for singletone use
func SetConfig(cfg *types.Config) {
	cfg.BasePath = strings.TrimSuffix(cfg.BasePath, "/")
	config = cfg
}

// GetConfig used to get configuration filled by SetConfig
func GetConfig() types.Config {
	return *config
}
