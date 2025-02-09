package api

import (
	"libraryapi/storage"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes initializes API routes
func SetupRoutes(store *storage.Store) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", GetBooks(store)).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook(store)).Methods("GET")
	router.HandleFunc("/books", AddBook(store)).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook(store)).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}