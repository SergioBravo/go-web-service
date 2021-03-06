package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"go-web-service/internal/api/jokes"
	"go-web-service/internal/config"
	"go-web-service/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	path := cfg.Host + ":" + cfg.Port

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	r := chi.NewRouter()

	h := handler.NewHandler(apiClient)

	r.Get("/hello", h.Hello)

	log.Println(fmt.Sprintf("starting server at %s", path))
	err = http.ListenAndServe(path, r)
	if err != nil {
		log.Fatal(err)
	}
}
