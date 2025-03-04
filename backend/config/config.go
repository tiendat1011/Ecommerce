package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DbHost string `env:"DB_HOST" envDefault:"localhost"`
	DbPort string `env:"DB_PORT" envDefault:"27017"`
	DbUser string `env:"DB_USER" envDefault:"admin"`
	DbPass string `env:"DB_PASS" envDefault:"password"`
}

func Load() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
