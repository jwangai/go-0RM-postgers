package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jack.go/pkg/models"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) { //function must be followed by http requests
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)     //throw status if it works
	json.NewEncoder(w).Encode(books) //encode meaning show all books
}
