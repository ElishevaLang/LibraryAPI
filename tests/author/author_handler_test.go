package service

import (
	"bytes"
	"encoding/json"
	"libraryapi/models"
	"libraryapi/api"
	"libraryapi/service_test"  // מייבא את ה-Mock של AuthorService
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// דוגמת בדיקה של ה-API
func TestAddAuthorAPI(t *testing.T) {
	mockService := new(service_test.MockAuthorService)
	author := models.Author{ID: "1", Name: "John Doe"}

	mockService.On("AddAuthor", author).Return(nil)

	reqBody, _ := json.Marshal(author)
	req, err := http.NewRequest("POST", "/authors", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/authors", api.AddAuthor(mockService)).Methods("POST")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}
