package controllers

import (
	"encoding/json" // Import for encoding and decoding JSON
	"fmt"
	"net/http"
	"strconv"

	"github.com/akhil/go-bookstore/pkg/models"
	"github.com/akhil/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

// Global variable for books
var NewBook models.Book

// Function to retrieve all books
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()                   // Call the method to retrieve all books
	res, _ := json.Marshal(newBooks)                   // Serialize the books into JSON
	w.Header().Set("Content-Type", "application/json") // Set the response content type to JSON
	w.WriteHeader(http.StatusOK)                       // Status 200 OK
	w.Write(res)                                       // Send the response
}

// Function to retrieve a book by its ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // Retrieve URL variables
	bookId := vars["bookId"]                  // Get the book ID from the variables
	ID, err := strconv.ParseInt(bookId, 0, 0) // Convert the book ID to an integer
	if err != nil {
		fmt.Println("Error while parsing bookId") // Handle parsing errors
		w.WriteHeader(http.StatusBadRequest)      // Status 400 Bad Request
		return
	}

	bookDetails, _ := models.GetBookById(ID) // Retrieve the book details
	res, _ := json.Marshal(bookDetails)      // Serialize the details into JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Status 200 OK
	w.Write(res)                 // Send the response
}

// Function to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}   // Instantiate a new Book model
	utils.ParseBody(r, CreateBook) // Parse the request body
	b := CreateBook.CreateBook()   // Add the new book to the database
	res, _ := json.Marshal(b)      // Serialize the created book into JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Status 200 OK
	w.Write(res)                 // Send the response
}

// Function to delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // Retrieve URL variables
	bookId := vars["bookId"]                  // Get the book ID
	ID, err := strconv.ParseInt(bookId, 0, 0) // Convert the book ID to an integer
	if err != nil {
		fmt.Println("Error while parsing bookId") // Handle parsing errors
		w.WriteHeader(http.StatusBadRequest)      // Status 400 Bad Request
		return
	}

	book := models.DeleteBook(ID) // Delete the book by its ID
	res, _ := json.Marshal(book)  // Serialize the result into JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Status 200 OK
	w.Write(res)                 // Send the response
}

// Function to update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}              // Instantiate a new Book model for updates
	utils.ParseBody(r, UpdateBook)            // Parse the request body
	vars := mux.Vars(r)                       // Retrieve URL variables
	bookId := vars["bookId"]                  // Corrected: "booId" to "bookId"
	ID, err := strconv.ParseInt(bookId, 0, 0) // Convert the book ID to an integer
	if err != nil {
		fmt.Println("Error while parsing bookId") // Handle parsing errors
		w.WriteHeader(http.StatusBadRequest)      // Status 400 Bad Request
		return
	}

	bookDetails, db := models.GetBookById(ID) // Retrieve book details and the associated database instance
	if UpdateBook.Name != "" {                // If a new name is provided
		bookDetails.Name = UpdateBook.Name // Update the book name
	}

	if UpdateBook.Publication != "" { // If a new publication value is provided
		bookDetails.Publication = UpdateBook.Publication // Update the publication field
	}

	db.Save(&bookDetails)               // Save the updated details in the database
	res, _ := json.Marshal(bookDetails) // Serialize the updated details into JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Status 200 OK
	w.Write(res)                 // Send the response
}
