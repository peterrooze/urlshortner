package api

import (
	"urlshortner/internal/shortener"

	"github.com/gorilla/mux"
)

func SetupRoutes(s *shortener.Shortener) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/shorten", shortenHandler(s)).Methods("POST")
	r.HandleFunc("/{shortURL}", redirectHandler(s)).Methods("GET")

	return r
}
