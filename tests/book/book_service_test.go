package api

import (
	"libraryapi/models"
	"libraryapi/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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


func TestGetBookByID(t *testing.T) {
	mockStore := new(MockStore)
	mockStore.On("GetBookByID", "1").Return(&models.Book{ID: "1", Title: "Book 1", Author: "Author 1", PublishedYear: 2021}, nil)

	bookService := service.NewBookService(mockStore)
	book, err := bookService.GetBookByID("1")

	assert.Nil(t, err)
	assert.Equal(t, "Book 1", book.Title)
	mockStore.AssertExpectations(t)
}

