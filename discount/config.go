package discount

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	PostgresUser     string `env:"PostgresUser" envDefault:"postgres"`
	PostgresPassword string `env:"PostgresPassword" envDefault:"postgres"`
	PostgresAddr     string `env:"PostgresAddr" envDefault:"localhost:5432"`
	PostgresDB       string `env:"PostgresDB" envDefault:"test"`
}

var Config config

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("cannot parse env variables err is: %v", err)
	}
}
