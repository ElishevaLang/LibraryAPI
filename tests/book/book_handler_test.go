package api

import (
	"bytes"
	"encoding/json"
	"libraryapi/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for BookService
type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) GetBooks() []models.Book {
	args := m.Called()
	return args.Get(0).([]models.Book)
}

func (m *MockBookService) GetBookByID(id string) (*models.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Book), args.Error(1)
}

func (m *MockBookService) AddBook(book models.Book) {
	m.Called(book)
}

func (m *MockBookService) DeleteBook(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookService) UpdateBook(updatedBook models.Book) error {
	args := m.Called(updatedBook)
	return args.Error(0)
}

func (m *MockBookService) GetSortedBooksByAsc() []models.Book {
	args := m.Called()
	return args.Get(0).([]models.Book)
}

func (m *MockBookService) GetBooksByPublishYear(year int) []models.Book {
	args := m.Called(year)
	return args.Get(0).([]models.Book)
}

func TestGetBooks(t *testing.T) {
	mockService := new(MockBookService)
	mockService.On("GetBooks").Return([]models.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", PublishedYear: 2021},
		{ID: "2", Title: "Book 2", Author: "Author 2", PublishedYear: 2020},
	})

	router := mux.NewRouter()
	router.HandleFunc("/books", GetBooks(mockService)).Methods("GET")

	req := httptest.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []models.Book
	err := json.NewDecoder(rr.Body).Decode(&books)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, books, 2)
	mockService.AssertExpectations(t)
}

func TestAddBook(t *testing.T) {
	mockService := new(MockBookService)
	book := models.Book{ID: "1", Title: "New Book", Author: "New Author", PublishedYear: 2021}

	mockService.On("AddBook", book).Return()

	reqBody, _ := json.Marshal(book)
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/books", AddBook(mockService)).Methods("POST")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	mockService := new(MockBookService)
	mockService.On("DeleteBook", "1").Return(nil)

	req, err := http.NewRequest("DELETE", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/books/{id}", DeleteBook(mockService)).Methods("DELETE")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateBook(t *testing.T) {
	mockService := new(MockBookService)
	updatedBook := models.Book{ID: "1", Title: "Updated Book", Author: "Updated Author", PublishedYear: 2021}

	mockService.On("UpdateBook", updatedBook).Return(nil)

	reqBody, _ := json.Marshal(updatedBook)
	req, err := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/books/{id}", UpdateBook(mockService)).Methods("PUT")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}
