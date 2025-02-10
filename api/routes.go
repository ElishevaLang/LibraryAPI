package api

import (
	"libraryapi/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(router *mux.Router, bookService *service.BookService, authorService *service.AuthorService) {

	router.HandleFunc("/books", GetBooks(bookService)).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook(bookService)).Methods("GET")
	router.HandleFunc("/books", AddBook(bookService)).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook(bookService)).Methods("DELETE")
	router.HandleFunc("/books/{id}", UpdateBook(bookService)).Methods("PUT")
	router.HandleFunc("/books/sorted", GetSortedBooksByAsc(bookService)).Methods("GET")
	router.HandleFunc("/books/year", GetBooksByPublishYear(bookService)).Methods("GET")

	router.HandleFunc("/authors", AddAuthor(authorService)).Methods("POST")
	router.HandleFunc("/authors/{id}", GetAuthorByID(authorService)).Methods("GET")
	router.HandleFunc("/authors/{id}", UpdateAuthor(authorService)).Methods("PUT")
	router.HandleFunc("/authors/{id}", DeleteAuthor(authorService)).Methods("DELETE")
	router.HandleFunc("/authors/search", SearchAuthorsByName(authorService)).Methods("GET") // הוספת מסלול לחיפוש מחברים
	router.HandleFunc("/authors/all", GetAllAuthors(authorService)).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
