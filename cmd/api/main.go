package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/yamato0204/grpc-card-server/internal/server"
)

func main() {

	var conf server.Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.ListenAndServe(&conf); err != nil {
		log.Panicf("failed to listen and serve: %+v", err)
	}
}
