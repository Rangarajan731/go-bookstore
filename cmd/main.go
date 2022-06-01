package main

import (
	"log"
	"net/http"

	"github.com/Rangarajan731/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln(err)
	}

}
