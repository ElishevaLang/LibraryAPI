package service
import (
	"libraryapi/models"
	"libraryapi/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAddAuthorService(t *testing.T) {
	// יצירת MockStore
	mockStore := new(MockStore)
	author := models.Author{ID: "1", Name: "John Doe"}

	// תיאום ההתנהגות של ה-Mock
	mockStore.On("AddBook", mock.AnythingOfType("models.Book")).Return(nil)

	// יצירת שירות חדש עם ה-MockStore
	bookService := service.NewBookService(mockStore)

	// הוספת ספר דרך השירות
	bookService.AddBook(models.Book{ID: "1", Title: "Book Title", Author: "John Doe", PublishedYear: 2021})

	// בדיקה שהקריאה ל-AddBook התבצעה כמצופה
	mockStore.AssertExpectations(t)
}

func TestGetBooksService(t *testing.T) {
	// יצירת MockStore
	mockStore := new(MockStore)
	mockStore.On("GetBooks").Return([]models.Book{
		{ID: "1", Title: "Book 1", Author: "Author 1", PublishedYear: 2021},
		{ID: "2", Title: "Book 2", Author: "Author 2", PublishedYear: 2020},
	})

	// יצירת שירות חדש עם ה-MockStore
	bookService := service.NewBookService(mockStore)

	// קריאת ספרים מהשירות
	books := bookService.GetBooks()

	// בדיקה שמספר הספרים שהוחזר תואם לצפוי
	assert.Len(t, books, 2)
	mockStore.AssertExpectations(t)
}
