package basket

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	RedisURI      string `env:"REDIS_URI" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`
}

var Config config

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("cannot parse env variables err is: %v", err)
	}
}
