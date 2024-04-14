package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		HTTP `yaml:"http"`
		PG   `yaml:"postgres"`
		JWT  `yaml:"jwt"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// PG -.
	PG struct {
		URL string `env-required:"true" yaml:"pg_url" env:"PG_URL"`
	}

	// JWT -.
	JWT struct {
		SecretKey      string `mapstructure:"secret_key" yaml:"secret_key"`
		AccessTokenTTL int64  `mapstructure:"access_token_ttl" yaml:"access_token_ttl"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
