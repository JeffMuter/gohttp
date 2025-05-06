package main

import (
	"gohttp/database"
	"gohttp/pages"
	"gohttp/router"
	"log"
	"net/http"
)

func main() {
	// Initialize templates before setting up the router
	err := pages.InitTemplates()
	if err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}

	// Initialize database
	err = database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	router := router.Router()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
