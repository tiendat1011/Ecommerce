package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DbHost string `env:"DB_HOST" envDefault:"localhost"`
	DbPort string `env:"DB_PORT" envDefault:"27017"`
	DbUser string `env:"DB_USER" envDefault:"admin"`
	DbPass string `env:"DB_PASS" envDefault:"password"`
	DbName string `env:"DB_NAME" envDefault:"ecommerce"`
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
	JwtSecret string `env:"JWT_SECRET" envDefault:"iluvgolang"`
}

var Cfg *Config

func Load() error {
	Cfg = &Config{}
	if err := env.Parse(Cfg); err != nil {
		return err
	}

	return nil
}
