package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerPort int    `yaml:"server_port" env-default:"8080"`
	MongoHost  string `yaml:"mongo_host" env-default:"mongodb://localhost:27017"`
}

func ProvideConfig() (*Config, error) {
	cfg := Config{}
	// TODO: actually read from yaml file
	// err := cleanenv.ReadConfig("./../../config.yaml", &cfg)
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &cfg, err
	}
	return &cfg, nil
}
