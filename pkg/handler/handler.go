package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Rangarajan731/go-bookstore/pkg/models"
)

type defaultResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Book    *models.Book `json:"book"`
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")
	json.NewEncoder(w).Encode(models.GetBooks(r))
}

func GetBook(w http.ResponseWriter, r *http.Request) {

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")
	json.NewEncoder(w).Encode(models.GetBook(r))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")
	book := models.UpdateBook(r)
	json.NewEncoder(w).Encode(defaultResponse{"completed", "Updated Book", book})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")
	book := models.DeleteBook(r)
	json.NewEncoder(w).Encode(defaultResponse{"completed", "Deleted Book", book})
}
