package internal

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var cfg Config

type Config struct {
	Env      string         `yaml:"env" env-default:"local"  env-required:"true"`
	DbConfig DatabaseConfig `yaml:"db"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Path   string `yaml:"path"`
}


func Init(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", path)
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
}

func ObtainConfig() Config {
	return cfg
}
