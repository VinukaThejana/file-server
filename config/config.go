// Package config .
// Used to load all the configurations for the functionality of the application
package config

import (
	"fmt"
	"strings"

	"github.com/VinukaThejana/go-utils/logger"
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

var log logger.Logger

// Config is a struct that contains configurations that are used throughout the
// application
type Config struct {
	Port string `mapstructure:"PORT"`
	Path string `mapstructure:"PATH"`
}

// Load is a fucntion that is used to load all the configurations from the relevant
// sources
func (c *Config) Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/file-server")

	if err := viper.ReadInConfig(); err != nil {
		if err != nil {
			if strings.HasPrefix(err.Error(), "Config File \"config\" Not Found in") {
				viper.SetDefault("PORT", ":8080")
				viper.SetDefault("PATH", fmt.Sprintf("%s/", xdg.Home))
			}
		} else {
			log.Errorf(err, nil)
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Errorf(err, nil)
	}
}
