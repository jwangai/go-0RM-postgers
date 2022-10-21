package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"jack.go/pkg/db"
	"jack.go/pkg/handlers"
)

func main() {
	DB := db.Init()           //initialize the db
	h := handlers.New(DB)     //pass the db to the handlers
	router := mux.NewRouter() //initialize the router

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
