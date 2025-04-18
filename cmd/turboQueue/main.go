package main

import (
	"context"
	"log"

	"github.com/appu900/TurboQueue/config"
	"github.com/appu900/TurboQueue/internal/api"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	server := api.NewServer(config)
	if err := server.Start(context.Background()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
