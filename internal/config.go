package internal

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"local"  env-required:"true"`
	DbConfig DatabaseConfig `yaml:"db"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Path   string `yaml:"path"`
}

var cfg Config

func initConfig() Config {
	isTest := os.Getenv("TEST") == "true"

	if isTest {
		return Config{
			Env: "test",
			DbConfig: DatabaseConfig {
				Driver: "postgres",
				Path: "postgresql://secUREusER:StrongEnoughPassword)@51.250.26.59:5432/postgres?sslmode=disable",
			},
		}
	}
	configPath := "I:/dev/go-projects/bank-profile-service/configs/test.yaml"

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}

func ObtainConfig() Config {
	if (cfg == Config{}) {
		cfg = initConfig()
	}
	return cfg
}
