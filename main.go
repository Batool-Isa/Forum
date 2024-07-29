package main

import (
	"Forum/backend/database"
	"Forum/backend/handler"
	"Forum/backend/middleware"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB("forum.db")
	http.Handle("/", middleware.SessionMiddleware(http.HandlerFunc(handler.IndexHandler)))
	http.Handle("/login", middleware.SessionMiddleware(http.HandlerFunc(handler.LoginHandler)))
	http.Handle("/register", middleware.SessionMiddleware(http.HandlerFunc(handler.RegisterHandler)))
	//http.Handle("/error", middleware.SessionMiddleware(http.HandlerFunc(handler.ErrorHandler)))
	http.Handle("/create_post", middleware.SessionMiddleware(http.HandlerFunc(handler.CreateHandler)))
	http.Handle("/like", middleware.SessionMiddleware(http.HandlerFunc(handler.LikePost)))
	http.Handle("/dislike", middleware.SessionMiddleware(http.HandlerFunc(handler.DislikePost)))
	http.Handle("/logout", middleware.SessionMiddleware(http.HandlerFunc(handler.Logout)))
	http.Handle("/post", middleware.SessionMiddleware(http.HandlerFunc(handler.PostHandler)))
	http.Handle("/add_comment", middleware.SessionMiddleware(http.HandlerFunc(handler.CommentHandler)))
	http.Handle("/like_comment", middleware.SessionMiddleware(http.HandlerFunc(handler.LikeComment)))
	http.Handle("/dislike_comment", middleware.SessionMiddleware(http.HandlerFunc(handler.DislikeComment)))


	// Serve static files
	fs := http.FileServer(http.Dir("templates/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Create tables
	database.CreateTables()
	log.Println("Database setup complete")

	database.AddDummyData()
	database.ShowData()

	fmt.Println("Server started at http://localhost:3090/")
	err := http.ListenAndServe(":3090", nil)
	if err != nil {
		log.Fatal("Error starting server at 3090", err)
	}
}
