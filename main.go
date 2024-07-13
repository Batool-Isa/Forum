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
	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/error", handler.ErrorHandler)
	http.HandleFunc("/create_post", handler.CreateHandler)
	http.HandleFunc("/like", handler.LikePost)
    http.HandleFunc("/dislike", handler.DislikePost)

	// Serve static files
	fs := http.FileServer(http.Dir("templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Create tables
	database.CreateTables()
	log.Println("Database setup complete")

	database.AddDummyData()
	database.ShowData()

	fmt.Println("Server started at http://localhost:3030/")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("Error starting server at 3000", err)
	}
}
