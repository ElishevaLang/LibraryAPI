package api

import (
	"encoding/json"
	"libraryapi/models"
	"libraryapi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetBooks(bookService *service.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author") 
		log.Printf("Received request: GET /books?author=%s", author)

		var books []models.Book
		if author != "" {
			books = bookService.GetBooksByAuthor(author) 
			log.Printf("Searching books by author: %s, found: %d books", author, len(books))
		} else {
			books = bookService.GetBooks() 
			log.Printf("Fetching all books, total found: %d", len(books))
		}

		if len(books) == 0 {
			log.Println("No books found")
			http.Error(w, "No books found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(books)
		log.Println("Books sent successfully")
	}
}

func GetBook(bookService *service.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Received request: GET /books/%s", vars["id"])

		book, err := bookService.GetBookByID(vars["id"])
		if err != nil {
			log.Printf("Book with ID %s not found", vars["id"])
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(book)
		log.Printf("Book with ID %s retrieved successfully", vars["id"])
	}
}

func AddBook(bookService *service.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request: POST /books")

		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Println("Invalid request payload")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		book.ID = uuid.New().String()
		bookService.AddBook(book)

		log.Printf("Book added successfully: ID=%s, Title=%s", book.ID, book.Title)
		w.WriteHeader(http.StatusCreated)
	}
}
func DeleteBook(bookService *service.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Received request: DELETE /books/%s", vars["id"])

		err := bookService.DeleteBook(vars["id"])
		if err != nil {
			log.Printf("Failed to delete book: ID %s not found", vars["id"])
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		log.Printf("Book with ID %s deleted successfully", vars["id"])
		w.WriteHeader(http.StatusNoContent)
	}
}
func UpdateBook(bookService *service.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Received request: PUT /books/%s", vars["id"])

		var updatedBook models.Book
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			log.Println("Invalid request payload")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		existingBook, err := bookService.GetBookByID(vars["id"])
		if err != nil {
			log.Printf("Book with ID %s not found", vars["id"])
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		existingBook.Title = updatedBook.Title
		existingBook.Author = updatedBook.Author
		existingBook.PublishedYear = updatedBook.PublishedYear

		bookService.UpdateBook(*existingBook)

		log.Printf("Book with ID %s updated successfully", vars["id"])
		w.WriteHeader(http.StatusOK)
	}
}
func GetSortedBooksByAsc(bookService *service.BookService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books := bookService.GetSortedBooksByAsc()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}

func GetBooksByPublishYear(bookService *service.BookService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        yearStr := r.URL.Query().Get("year")
        year, err := strconv.Atoi(yearStr)
        if err != nil {
            http.Error(w, "Invalid year parameter", http.StatusBadRequest)
            return
        }

        books := bookService.GetBooksByPublishYear(year)
        json.NewEncoder(w).Encode(books)
    }
}
