package storage

import (
	"errors"
	"libraryapi/models"
	"sync"
)

// Store מנהל את אחסון הספרים בזיכרון
type Store struct {
	mu    sync.Mutex
	books map[string]models.Book
}

// NewStore מאתחל את ה-Store
func NewStore() *Store {
	return &Store{
		books: make(map[string]models.Book),
	}
}

// AddBook מוסיף ספר
func (s *Store) AddBook(book models.Book) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.books[book.ID] = book
}

// GetBooks מחזיר את כל הספרים
func (s *Store) GetBooks() []models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []models.Book
	for _, book := range s.books {
		result = append(result, book)
	}
	return result
}

// GetBookByID מחזיר ספר לפי ID
func (s *Store) GetBookByID(id string) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	book, exists := s.books[id]
	if !exists {
		return models.Book{}, errors.New("book not found")
	}
	return book, nil
}

// DeleteBook מוחק ספר לפי ID
func (s *Store) DeleteBook(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(s.books, id)
	return nil
}

// GetBooksByAuthor מחפש ספרים לפי שם המחבר
func (s *Store) GetBooksByAuthor(author string) []models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []models.Book
	for _, book := range s.books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}