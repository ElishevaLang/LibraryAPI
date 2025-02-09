package api

import (
	"encoding/json"
	"libraryapi/models"
	"libraryapi/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

// GetBooks מחזיר את כל הספרים או מחפש לפי שם מחבר אם סופק פרמטר
func GetBooks(store *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author") // מקבל פרמטר של מחבר מה-URL
		log.Printf("Received request: GET /books?author=%s", author)

		var books []models.Book
		if author != "" {
			books = store.GetBooksByAuthor(author) // אם יש מחבר, מחפש ספרים של המחבר
			log.Printf("Searching books by author: %s, found: %d books", author, len(books))
		} else {
			books = store.GetBooks() // אם אין מחבר, מחזיר את כל הספרים
			log.Printf("Fetching all books, total found: %d", len(books))
		}

		// אם לא נמצאו ספרים
		if len(books) == 0 {
			log.Println("No books found")
			http.Error(w, "No books found", http.StatusNotFound)
			return
		}

		// מחזיר את הספרים כ-JSON
		json.NewEncoder(w).Encode(books)
		log.Println("Books sent successfully")
	}
}
// GetBook retrieves a book by ID
func GetBook(store *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Received request: GET /books/%s", vars["id"])

		book, err := store.GetBookByID(vars["id"])
		if err != nil {
			log.Printf("Book with ID %s not found", vars["id"])
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(book)
		log.Printf("Book with ID %s retrieved successfully", vars["id"])
	}
}

// AddBook adds a new book
func AddBook(store *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request: POST /books")

		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Println("Invalid request payload")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// יצירת ID ייחודי לספר
		book.ID = uuid.New().String()
		store.AddBook(book)

		log.Printf("Book added successfully: ID=%s, Title=%s", book.ID, book.Title)
		w.WriteHeader(http.StatusCreated)
	}
}

// DeleteBook removes a book by ID
func DeleteBook(store *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Received request: DELETE /books/%s", vars["id"])

		err := store.DeleteBook(vars["id"])
		if err != nil {
			log.Printf("Failed to delete book: ID %s not found", vars["id"])
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		log.Printf("Book with ID %s deleted successfully", vars["id"])
		w.WriteHeader(http.StatusNoContent)
	}
}