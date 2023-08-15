package main

import (
	"log"

	"github.com/WBear29/go-restapi-server-template/config"
	"github.com/WBear29/go-restapi-server-template/internal/app"
)

var version string

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error while reading config: %v", err)
	}
	log.Printf("Starting: %v version: %v env: %v", cfg.App.Name, version, cfg.App.Env)

	// Run App
	app.Run(cfg, version)
}
