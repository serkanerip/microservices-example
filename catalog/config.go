package catalog

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	MongoURI string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
}

var Config config

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("cannot parse env variables err is: %v", err)
	}
}
