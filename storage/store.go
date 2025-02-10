package storage

import (
	"errors"
	"libraryapi/models"
	"strings" 
	"sync"
	"sort"
)

type Store struct {
	mu      sync.RWMutex
	books   map[string]models.Book
	authors map[string]models.Author 
}

func NewStore() *Store {
	return &Store{
		books:   make(map[string]models.Book),
		authors: make(map[string]models.Author), 
	}
}

func (s *Store) AddBook(book models.Book) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.books[book.ID] = book
}

func (s *Store) GetBooks() []models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var books []models.Book
	for _, book := range s.books {
		books = append(books, book)
	}
	return books
}

func (s *Store) GetBooksByAuthor(author string) []models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var books []models.Book
	for _, book := range s.books {
		if strings.Contains(strings.ToLower(book.Author), strings.ToLower(author)) {
			books = append(books, book)
		}
	}
	return books
}

func (s *Store) GetBookByID(id string) (*models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	book, exists := s.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return &book, nil
}

func (s *Store) DeleteBook(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(s.books, id)
	return nil
}

func (s *Store) UpdateBook(updatedBook models.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.books[updatedBook.ID]
	if !exists {
		return errors.New("book not found")
	}
	s.books[updatedBook.ID] = updatedBook
	return nil
}

func (s *Store) GetSortedBooksByAsc() []models.Book {
    s.mu.RLock()
    defer s.mu.RUnlock()

    books := make([]models.Book, 0, len(s.books))
    for _, book := range s.books {
        books = append(books, book)
    }

    sort.Slice(books, func(i, j int) bool {
        return books[i].Title < books[j].Title
    })

    return books
}

func (s *Store) GetBooksByPublishYear(year int) []models.Book {
    s.mu.RLock()
    defer s.mu.RUnlock()
    var books []models.Book
    for _, book := range s.books {
        if book.PublishedYear == year {
            books = append(books, book)
        }
    }
    return books
}

func (s *Store) AddAuthor(author models.Author) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.authors[author.ID] = author
}

func (s *Store) GetAllAuthors() []models.Author {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var authors []models.Author
	for _, author := range s.authors {
		authors = append(authors, author)
	}
	return authors
}

func (s *Store) GetAuthorByID(id string) (*models.Author, error) { 
	defer s.mu.RUnlock()
	author, exists := s.authors[id]
	if !exists {
		return nil, errors.New("author not found")
	}
	return &author, nil
}

func (s *Store) UpdateAuthor(id string, newName string) error { 
	s.mu.Lock()
	defer s.mu.Unlock()
	author, exists := s.authors[id]
	if !exists {
		return errors.New("author not found")
	}
	author.Name = newName
	s.authors[id] = author
	return nil
}

func (s *Store) DeleteAuthor(id string) error { // עדכון ל- string
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.authors[id]; !exists {
		return errors.New("author not found")
	}
	delete(s.authors, id)
	return nil
}


func (s *Store) SearchAuthorsByName(query string) []models.Author {
	s.mu.Lock()
	defer s.mu.Unlock()
	var authors []models.Author
	for _, author := range s.authors { 
		if contains(author.Name, query) {
			authors = append(authors, author)
		}
	}
	return authors
}

func contains(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}


