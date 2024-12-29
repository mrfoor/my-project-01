package main

import (
	"log"
	"net/http"

	"github.com/akhil/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"                  // Removed duplicate import
	_ "github.com/jinzhu/gorm/dialects/mysql" // Corrected import path for MySQL dialect
)

func main() {
	r := mux.NewRouter()                                      // Create a new router
	routes.RegisterStoreRoutes(r)                             // Register book routes (fixed function name)
	http.Handle("/", r)                                       // Attach the router to the default handler
	log.Println("Server is running on http://localhost:9010") // Log server start
	log.Fatal(http.ListenAndServe(":9010", r))                // Start the server and listen on port 9010
}
