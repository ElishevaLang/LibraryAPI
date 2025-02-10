package api

import (
	"encoding/json"
	"libraryapi/service"
	"libraryapi/models"
	"net/http"
	"github.com/gorilla/mux"
)

func AddAuthor(authorService *service.AuthorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		err := json.NewDecoder(r.Body).Decode(&author)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err = authorService.AddAuthor(author)
		if err != nil {
			http.Error(w, "Error adding author", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAuthorByID(authorService *service.AuthorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		author, err := authorService.GetAuthorByID(vars["id"])
		if err != nil {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(author)
	}
}

func UpdateAuthor(authorService *service.AuthorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var updatedAuthor models.Author
		err := json.NewDecoder(r.Body).Decode(&updatedAuthor)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err = authorService.UpdateAuthor(vars["id"], updatedAuthor.Name)
		if err != nil {
			http.Error(w, "Error updating author", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteAuthor(authorService *service.AuthorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		err := authorService.DeleteAuthor(vars["id"])
		if err != nil {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
// SearchAuthorsByName מחפש מחברים לפי שם
func SearchAuthorsByName(authorService *service.AuthorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("name")
		if query == "" {
			http.Error(w, "Query parameter 'name' is required", http.StatusBadRequest)
			return
		}
		authors := authorService.SearchAuthorsByName(query)
		if len(authors) == 0 {
			http.Error(w, "No authors found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	}
}


// GetAllAuthors מחזיר את כל המחברים
func GetAllAuthors(authorService *service.AuthorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors := authorService.GetAllAuthors()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	}
}

