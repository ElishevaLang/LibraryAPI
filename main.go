package main

import (
	"libraryapi/api"
	"libraryapi/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewStore()
	router := api.SetupRoutes(store)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}