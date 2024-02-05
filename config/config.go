package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var ContextKey = struct{ name string }{name: "config"}

type Config struct {
	App          AppConfiguration
	DB           DBConfiguration
	BuildVersion string
	BuildDate    string
}

type AppConfiguration struct {
	Name    string
	Address string
}
type DBConfiguration struct {
	URL string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return &config, nil
}
