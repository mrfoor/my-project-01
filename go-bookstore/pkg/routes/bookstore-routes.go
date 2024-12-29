package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterStoreRoutes sets up the routes for book-related endpoints
var RegisterStoreRoutes = func(router *mux.Router) {
	// Corrected function name for CreateBook
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	// Updated to avoid duplicate routes for the same HTTP method
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")               // Retrieves all books
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   // Retrieves a book by ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    // Updates a book by ID
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") // Deletes a book by ID
}
