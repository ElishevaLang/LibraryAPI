package main

import (
	"libraryapi/api"
	"libraryapi/service"
	"libraryapi/storage"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// אתחול מחסן הנתונים (שכולל גם ספרים וגם מחברים)
	store := storage.NewStore()

	bookService := service.NewBookService(store)
	authorService := service.NewAuthorService(store)

	router := mux.NewRouter()

	api.SetupRoutes(router, bookService, authorService)


	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}






	


