package config

import "github.com/kelseyhightower/envconfig"

// NewConfig generates a new config struct
func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("KRUDER", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
