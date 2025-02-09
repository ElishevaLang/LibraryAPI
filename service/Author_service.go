package service


type AuthorService struct {
	authors map[string]string // AuthorID, AuthorName
}

func NewAuthorService() *AuthorService {
	return &AuthorService{authors: make(map[string]string)}
}

func (s *AuthorService) AddAuthor(id string, name string) {
	s.authors[id] = name
}

func (s *AuthorService) GetAuthorName(id string) string {
	return s.authors[id]
}