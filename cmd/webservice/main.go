package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"go-web-service/internal/handler"
)

func main() {
	r := chi.NewRouter()

	h := handler.NewHandler()

	r.Get("/hello", h.Hello)

	log.Println("starting server at localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
