package main

import (
	"log"

	config "github.com/gleonz/golang-boilerplate/configuration"
	app "github.com/gleonz/golang-boilerplate/internal/application"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
