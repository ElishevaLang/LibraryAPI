package service

import (
	"libraryapi/models"
	"libraryapi/storage"
)

type AuthorService struct {
	store *storage.Store
}

func NewAuthorService(store *storage.Store) *AuthorService {
	return &AuthorService{store: store}
}

func (s *AuthorService) AddAuthor(author models.Author) error {
	s.store.AddAuthor(author)
	return nil
}

func (s *AuthorService) GetAuthorByID(id string) (*models.Author, error) {
	return s.store.GetAuthorByID(id)
}

func (s *AuthorService) UpdateAuthor(id string, newName string) error {
	return s.store.UpdateAuthor(id, newName)
}

func (s *AuthorService) DeleteAuthor(id string) error {
	return s.store.DeleteAuthor(id)
}

// func (s *AuthorService) SearchAuthorsByName(query string) []models.Author {
// 	return s.store.SearchAuthorsByName(query)
// }

//================================================================
func (s *AuthorService) SearchAuthorsByName(query string) []models.Author {
	return s.store.SearchAuthorsByName(query)
}
//================================================================

func (s *AuthorService) GetAllAuthors() []models.Author {
	return s.store.GetAllAuthors()
}


