package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// Define your configuration fields here
}

func LoadConfig() (*Config, error) {
	// Load configuration from environment variables, files, etc.
	// Example using viper:
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	// Parse configuration into Config struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
