package models

import (
	"github.com/akhil/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book structure represents the book model
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"` // Corrected JSON tags: `json:"name"`
	Author      string `json:"author"`       // Corrected JSON tags
	Publication string `json:"publication"`  // Corrected JSON tags
}

// Initialize the database connection and auto-migrate the Book model
func init() {
	config.Connect()        // Establish connection to the database
	db = config.GetDB()     // Get the database instance
	db.AutoMigrate(&Book{}) // Automatically migrate the Book model
}

// CreateBook adds a new book to the database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // Mark the book record as new
	db.Create(&b)   // Create the record in the database
	return b        // Return the created book
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() []Book {
	var Books []Book // Create a slice to hold books
	db.Find(&Books)  // Retrieve all books and store them in the slice
	return Books     // Return the list of books
}

// GetBookById retrieves a book by its ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book                            // Create a Book variable to hold the result
	db := db.Where("ID = ?", Id).Find(&getBook) // Find the book by ID
	return &getBook, db                         // Return the book and the database instance
}

// DeleteBook deletes a book by its ID
func DeleteBook(ID int64) Book {
	var book Book                        // Create a Book variable to hold the result
	db.Where("ID = ?", ID).Delete(&book) // Delete the book by ID
	return book                          // Return the deleted book
}
