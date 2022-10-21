package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"jack.go/pkg/models"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	//read the dynamic id parameter

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//find the book by id
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//delete bok
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
