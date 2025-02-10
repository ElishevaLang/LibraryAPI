package storage

import (
	"libraryapi/models"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) AddBook(book models.Book) {
	m.Called(book)
}

func (m *MockStore) GetBooks() []models.Book {
	args := m.Called()
	return args.Get(0).([]models.Book)
}

func (m *MockStore) GetBookByID(id string) (*models.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Book), args.Error(1)
}

func (m *MockStore) DeleteBook(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockStore) UpdateBook(updatedBook models.Book) error {
	args := m.Called(updatedBook)
	return args.Error(0)
}

func (m *MockStore) GetSortedBooksByAsc() []models.Book {
	args := m.Called()
	return args.Get(0).([]models.Book)
}

func (m *MockStore) GetBooksByPublishYear(year int) []models.Book {
	args := m.Called(year)
	return args.Get(0).([]models.Book)
}
