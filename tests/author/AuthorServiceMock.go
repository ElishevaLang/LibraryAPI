package service 
import (
	"libraryapi/models"
	"github.com/stretchr/testify/mock"
)

// AddAuthor - mock של פונקציה להוספת סופר
func (m *MockAuthorService) AddAuthor(author models.Author) error {
	args := m.Called(author)
	return args.Error(0)
}

// GetAuthorByID - mock של פונקציה לקבלת סופר לפי ID
func (m *MockAuthorService) GetAuthorByID(id string) (*models.Author, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Author), args.Error(1)
}

// UpdateAuthor - mock של פונקציה לעדכון סופר
func (m *MockAuthorService) UpdateAuthor(id string, newName string) error {
	args := m.Called(id, newName)
	return args.Error(0)
}

// DeleteAuthor - mock של פונקציה למחוק סופר
func (m *MockAuthorService) DeleteAuthor(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetAllAuthors - mock של פונקציה לקבלת כל המחברים
func (m *MockAuthorService) GetAllAuthors() []models.Author {
	args := m.Called()
	return args.Get(0).([]models.Author)
}

// SearchAuthorsByName - mock של פונקציה לחיפוש סופרים לפי שם
func (m *MockAuthorService) SearchAuthorsByName(name string) []models.Author {
	args := m.Called(name)
	return args.Get(0).([]models.Author)
}
