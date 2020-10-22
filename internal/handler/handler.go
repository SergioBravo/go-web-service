package handler

import (
	"fmt"
	"go-web-service/internal/api"
	"net/http"
)

// Handler ...
type Handler struct {
	jokeClient api.Client
}

// NewHandler ...
func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

// Hello ...
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, joke.Joke)
}
