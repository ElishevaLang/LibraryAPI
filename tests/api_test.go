package tests

import (
	"bytes"
	"encoding/json"
	"libraryapi/api"
	"libraryapi/models"
	"net/http"
	"net/http/httptest"
	"testing"
    "libraryapi/storage"
)

// TestAddBook בודק הוספת ספר חדש
func TestAddBook(t *testing.T) {
	store := storage.NewStore()
	router := api.SetupRoutes(store)

	book := models.Book{ID: "1", Title: "Test Book", Author: "John Doe", PublishedYear: 2021}
	body, _ := json.Marshal(book)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, rr.Code)
	}
}

// TestGetBooksByAuthor בודק חיפוש לפי מחבר
func TestGetBooksByAuthor(t *testing.T) {
	store := storage.NewStore()
	store.AddBook(models.Book{ID: "1", Title: "Go Programming", Author: "Alice", PublishedYear: 2020})
	store.AddBook(models.Book{ID: "2", Title: "Python Basics", Author: "Bob", PublishedYear: 2021})

	router := api.SetupRoutes(store)

	req, _ := http.NewRequest("GET", "/books?author=Alice", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var books []models.Book
	json.Unmarshal(rr.Body.Bytes(), &books)

	if len(books) != 1 || books[0].Author != "Alice" {
		t.Errorf("Expected 1 book by Alice but got %d", len(books))
	}
}

func TestDeleteBook(t *testing.T) {
	store := storage.NewStore()
	store.AddBook(models.Book{ID: "1", Title: "Test Book", Author: "John Doe", PublishedYear: 2021})
	router := api.SetupRoutes(store)

	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got %d", http.StatusNoContent, rr.Code)
	}

	// Test deleting non-existent book
	req, _ = http.NewRequest("DELETE", "/books/999", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, rr.Code)
	}
}