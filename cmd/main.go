package main

import (
	"log"
	"net/http"

	"github.com/theborzet/connection_project/internal/bybit/repository"
	"github.com/theborzet/connection_project/internal/handlers"
	"github.com/theborzet/connection_project/pkg/common/config"
)

func main() {
	config, err := config.LOadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client := &http.Client{}

	mu := http.NewServeMux()

	api := &repository.MyAPI{
		Api:    client,
		Config: config,
	}

	handler := handlers.NewHandler(api)

	handlers.RegistrationRouts(client, mu, handler)

	if err := http.ListenAndServe(config.Port, mu); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
