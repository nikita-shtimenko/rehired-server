package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Database `yaml:"database"`
	}

	App struct {
		Env     string `env-required:"true" yaml:"env" env:"ENV"`
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Database struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"DATABASE_POOL_MAX"`
		URL     string `env-required:"true"                 env:"DATABASE_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./internal/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
