package handler

import (
	"net/http"

	"github.com/MrSaveliy/vk-filmoteca/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serveces *service.Service) *Handler {
	return &Handler{services: serveces}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-up", h.signUpHandler)
	mux.HandleFunc("/auth/sign-in", h.signInHandler)

	mux.HandleFunc("/role/create", h.createRole)

	mux.HandleFunc("/movies/create", h.createMovie)
	mux.HandleFunc("/movies/", h.getMovie)

	mux.HandleFunc("/actors/create", h.createActor)
	mux.HandleFunc("/actors/", h.getActor)

	return mux
}
