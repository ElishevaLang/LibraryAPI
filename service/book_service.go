package service

import (
	"encoding/json"
	"errors"
	"libraryapi/models"
	"os"
	"sync"
)

// BookService מנהל את הספרים ושומר אותם לקובץ JSON
type BookService struct {
	mu       sync.Mutex
	books    map[string]models.Book
	filename string
}

// NewBookService יוצר מופע חדש של השירות
func NewBookService(filename string) *BookService {
	service := &BookService{
		books:    make(map[string]models.Book),
		filename: filename,
	}
	service.loadBooks()
	return service
}

// loadBooks טוען ספרים מהקובץ
func (s *BookService) loadBooks() {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.Open(s.filename)
	if err != nil {
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&s.books)
}

// saveBooks שומר את הספרים לקובץ JSON
func (s *BookService) saveBooks() {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, _ := os.Create(s.filename)
	defer file.Close()
	json.NewEncoder(file).Encode(s.books)
}

// GetBooks מחזיר את כל הספרים
func (s *BookService) GetBooks() []models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	var books []models.Book
	for _, book := range s.books {
		books = append(books, book)
	}
	return books
}

// GetBookByID מחזיר ספר לפי מזהה
func (s *BookService) GetBookByID(id string) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	book, exists := s.books[id]
	if !exists {
		return models.Book{}, errors.New("book not found")
	}
	return book, nil
}

// AddBook מוסיף ספר
func (s *BookService) AddBook(book models.Book) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.books[book.ID] = book
	s.saveBooks()
}

// DeleteBook מוחק ספר לפי מזהה
func (s *BookService) DeleteBook(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(s.books, id)
	s.saveBooks()
	return nil
}