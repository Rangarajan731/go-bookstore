package routes

import (
	"github.com/Rangarajan731/go-bookstore/pkg/handler"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/book", handler.AddBook).Methods("POST")
	router.HandleFunc("/book/{id}", handler.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", handler.DeleteBook).Methods("DELETE")
}
