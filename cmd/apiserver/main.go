package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"go_test_learning/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	// apiserver.exe -help
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err = s.Start(); err != nil {
		log.Fatal(err)
	}
}
