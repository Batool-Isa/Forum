package main

import (
	"Forum/backend/database"
	"Forum/backend/handler"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB("forum.db")
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/error", handler.ErrorHandler)
	http.HandleFunc("/create_post", handler.CreateHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Create tables
	database.CreateTables()
	log.Println("Database setup complete")

	database.AddDummyData()
	database.ShowData()

	fmt.Println("Server started at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server at 8080", err)
	}
}
