package config

import (
	goenv "github.com/Netflix/go-env"
	"log"
)

type environment struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBName     string `env:"DB_NAME"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	APIPort    int    `env:"API_PORT"`
}

var ENV environment

func PopulateEnv() {
	_, err := goenv.UnmarshalFromEnviron(&ENV)
	if err != nil {
		log.Fatal(err)
	}
}
