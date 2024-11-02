package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Adejare77/bookStore/internal/models"
	"github.com/Adejare77/bookStore/internal/utils"
)

type BookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Response(book models.Book) BookResponse {
	return BookResponse{
		ID:          book.ID,
		Title:       *book.Title,
		Author:      *book.Author,
		Publication: *book.Publication,
	}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "Server Error -- Unable to Fetch Data", http.StatusInternalServerError)
		return
	}
	responses := []BookResponse{}
	if len(books) != 0 {
		for _, book := range books {
			responses = append(responses, Response(book))
		}
	}
	if err := json.NewEncoder(w).Encode(responses); err != nil {
		http.Error(w, "Unable to Encode", http.StatusBadRequest)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookId := utils.StringToInt(w, r)

	book, err := models.GetBookById(uint64(bookId))

	if err != nil {
		if err.Error() == "Invalid Book ID" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.NewCustomError(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"Error": "Server Error"})
		}
		return
	}

	json.NewEncoder(w).Encode(Response(book))
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	bookId := utils.StringToInt(w, r)

	if err := models.DeleteBookById(bookId); err != nil {
		if err.Error() == "Invalid Book ID" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.NewCustomError(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Server Error"})
		}
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Book Deleted Successfully"})
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// var body *models.Book = &models.Book{}
	body := &models.Book{}

	if err := utils.ParseBody(r, body); err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
	}

	bookId := utils.StringToInt(w, r)

	if err := models.UpdateBookById(bookId, *body); err != nil {
		if err.Error() == "Invalid Book ID" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.NewCustomError(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Server Error"})
			http.Error(w, "BadRequest", http.StatusBadRequest)
		}
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/books/%d", bookId), http.StatusSeeOther)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	body := &models.Book{}

	if err := utils.ParseBody(r, body); err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
	}
	if err := body.CreateBook(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error Posting Book")
		if strings.Contains(err.Error(), "Duplicate entry") {
			// w.WriteHeader(http.StatusConflict)
			// json.NewEncoder(w).Encode(map[string]string{"error": "Duplicate Entry"})
			http.Error(w, `{"error": "Duplicate Entry"}`, http.StatusConflict)
		} else if strings.Contains(err.Error(), "'author' cannot be null") {
			json.NewEncoder(w).Encode(map[string]string{"error": "author field missing"})
		} else if strings.Contains(err.Error(), "'title' cannot be null") {
			json.NewEncoder(w).Encode(map[string]string{"error": "title field missing"})
		} else if strings.Contains(err.Error(), "'publication' cannot be null") {
			json.NewEncoder(w).Encode(map[string]string{"error": "publication field missing"})
		} else {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(Response(*body))
}
