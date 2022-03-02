package mux_controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/atultherajput/go_crash_course/database/grom/dao"
	"github.com/atultherajput/go_crash_course/models"
	"github.com/gorilla/mux"
)

var Handler dao.Handler

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Books table
	Handler.Add(book)

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)

	var response = models.JsonResponse{Type: "success", Message: "The book has been inserted successfully!"}
	json.NewEncoder(w).Encode(response)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	//Find all books
	books := Handler.GetAll()

	w.WriteHeader(http.StatusOK)

	var response = models.JsonResponse{Type: "success", Data: books}
	json.NewEncoder(w).Encode(response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Find book by Id
	book := Handler.Get(id)

	w.WriteHeader(http.StatusOK)

	var response = models.JsonResponse{Type: "success", Data: book}
	json.NewEncoder(w).Encode(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	//Update DB
	Handler.Update(id, updatedBook)

	w.WriteHeader(http.StatusOK)

	var response = models.JsonResponse{Type: "success", Message: "The book has been upadted successfully!"}
	json.NewEncoder(w).Encode(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete that book
	Handler.Delete(id)

	w.WriteHeader(http.StatusOK)

	var response = models.JsonResponse{Type: "success", Message: "The book has been deleted successfully!"}
	json.NewEncoder(w).Encode(response)
}
