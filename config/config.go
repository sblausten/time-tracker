package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Db  DBConfig
	Env string
}

type DBConfig struct {
	Name                string
	LocalAddress        string
	DockerServerAddress string
}

func (c Config) Build() Config {
	config := loadFrom("application-config.json")

	env, envIsPresent := os.LookupEnv("ENV")
	if !envIsPresent || env == "" {
		env = "dev"
	}

	config.Env = env
	log.Println("Running with env:", env)

	return config
}

func (d DBConfig) GetAddress(config Config) string {
	switch config.Env {
	case "test":
		return config.Db.DockerServerAddress
	default:
		return config.Db.LocalAddress
	}
}

func loadFrom(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open config file: ", err)
	}
	defer file.Close()

	configuration := Config{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Cannot decode json config: ", err)
	}

	return configuration
}
