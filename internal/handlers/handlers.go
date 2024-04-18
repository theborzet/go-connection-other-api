package handlers

import (
	"net/http"

	"github.com/theborzet/connection_project/internal/bybit/repository"
)

type Handler struct {
	repository repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

func RegistrationRouts(client *http.Client, mu *http.ServeMux, handler *Handler) {
	mu.HandleFunc("/username", handler.GetStatusHedg)
}
