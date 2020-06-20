package server

import (
	"flag"
	"github.com/BurntSushi/toml"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	MongoDBURL  string `toml:"mongodb_url"`
	MongoDBName string `toml:"mongodb_name"`
}

func NewConfig() (*Config, error) {
	configPath := configPath()
	config := defaultConfig()
	_, err := toml.DecodeFile(configPath, &config)
	return config, err
}

func configPath() string {
	var configPath string
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.Parse()
	return configPath
}

func defaultConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		MongoDBURL:  "mongodb://localhost:27017",
		MongoDBName: "testing",
	}
}
