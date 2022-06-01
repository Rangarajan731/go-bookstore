package routes

import (
	"github.com/Rangarajan731/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/books/", routes.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", routes.GetBook).Methods("GET")
	router.HandleFunc("/book/", routes.AddBook).Methods("POST")
	router.HandleFunc("/book/{id}", routes.UpdateBook).Methods("PUT")
	router.handleFunc("/book/{id}", routes.DeleteBook).Methods("DELETE")
}
