package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Address string `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() *Config{
	configPath, exists := os.LookupEnv("CONFIG_PATH")

    if !exists {
    	log.Fatalf("CONFIG PATH is not set: %s", configPath)
   }

   var cfg Config

   err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}