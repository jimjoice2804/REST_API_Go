package Config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var cfg Config

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "local.yaml"
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
