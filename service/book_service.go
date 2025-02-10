package service

import (
	"libraryapi/models"
	"libraryapi/storage"
)

type BookService struct {
	store *storage.Store
}

func NewBookService(store *storage.Store) *BookService {
	return &BookService{store: store}
}

func (s *BookService) AddBook(book models.Book) {
	s.store.AddBook(book)
}

func (s *BookService) GetBooks() []models.Book {
	return s.store.GetBooks()
}

func (s *BookService) GetBookByID(id string) (*models.Book, error) {
	return s.store.GetBookByID(id)
}

func (s *BookService) DeleteBook(id string) error {
	return s.store.DeleteBook(id)
}

func (s *BookService) GetBooksByAuthor(author string) []models.Book {
	return s.store.GetBooksByAuthor(author)
}

func (s *BookService) UpdateBook(updatedBook models.Book) error {
	// אם הספר לא נמצא בסטור, תחפש אותו
	existingBook, err := s.store.GetBookByID(updatedBook.ID)
	if err != nil {
		return err
	}
	existingBook.Title = updatedBook.Title
	existingBook.Author = updatedBook.Author
	existingBook.PublishedYear = updatedBook.PublishedYear

	return s.store.UpdateBook(*existingBook)
}
func (s *BookService) GetSortedBooksByAsc() []models.Book {
    return s.store.GetSortedBooksByAsc()
}
func (s *BookService) GetBooksByPublishYear(year int) []models.Book {
    return s.store.GetBooksByPublishYear(year)
}


