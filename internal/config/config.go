package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env:"ADDRESS" env-default:"localhost:8080"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true" `
	HTTPServer  `yaml:"http_server" env-prefix:"HTTP_SERVER_"`
}

func MustLoad() *Config {
	var cfg Config

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "config/local.yaml"
	}

	//this will read the local.yaml and add to cfg struct
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %s: %v", err)
	}

	//
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("invalid environment config: %v", err)
	}

	return &cfg
}
