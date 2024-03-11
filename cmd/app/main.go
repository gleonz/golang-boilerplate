package main

import (
	"log"

	"github.com/gleonz/golang-boilerplate/config"
	"github.com/gleonz/golang-boilerplate/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
