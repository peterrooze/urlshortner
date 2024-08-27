package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"urlshortner/internal/api"
	"urlshortner/internal/shortener"
	"urlshortner/internal/storage"
)

func main() {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	store, err := storage.NewPostgresStorage(connString)
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL storage: %v", err)
	}
	defer store.Close()

	shortener := shortener.NewShortener(store)
	router := api.SetupRoutes(shortener)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
