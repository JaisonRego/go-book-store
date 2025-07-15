package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaisonrego/bookstore/pkg/models"
	"github.com/jaisonrego/bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBook := models.GetAllBooks()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewBook)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error While Parsing")
	}
	book, _ := models.GetBookById(ID)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error While Parsing")
	}
	book := models.DeleteBook(ID)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error While Parsing")
	}
	bookDetails, db := models.GetBookById(ID)

	if bookDetails.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if bookDetails.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if bookDetails.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookDetails)
}
