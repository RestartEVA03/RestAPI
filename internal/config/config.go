package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage-path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address      string        `yaml:"addres" env-default:"localhost:8080"`
	Timeout      time.Duration `yaml:"timeout" env_default:"4s"`
	IddleTimeout time.Duration `yaml:"idle_timeout" end_default:"60s"`
}

func Must_LoadConf() *Config {
	configPath := os.Getenv("../../config/local.yaml")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var conf Config

	if err := cleanenv.ReadConfig(configPath, &conf); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &conf
}
