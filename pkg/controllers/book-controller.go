package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AutonuKro/go-book-management-api/pkg/models"
	"github.com/AutonuKro/go-book-management-api/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBook()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	book := createBook.CreateBook()
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		fmt.Println("error while parsing")
	}
	deletedBook := models.DeleteBook(ID)
	resposne, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resposne)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook models.Book
	utils.ParseBody(r, &updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/jso")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
