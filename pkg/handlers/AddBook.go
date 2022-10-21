package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"jack.go/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	//1. read the request body
	//2. append to the book mocks
	//3. send 201 created response

	//1
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)
	//2 adding book to the database

	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	//3

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created.")
}
