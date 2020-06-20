package main

import (
	"go_test_learning/internal/app/server"
	"log"
	"net/http"
)

func main() {
	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	apiServer, err := server.New(config)
	if err != nil {
		log.Fatal(err)
	}

	router := server.NewRouter(apiServer)

	if err = http.ListenAndServe(config.BindAddr, router); err != nil {
		log.Fatal(err)
	}
}
