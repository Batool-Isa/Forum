package main

import (
	"Forum/backend/database"
	"Forum/backend/handler"
	"Forum/backend/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, fileErr := os.OpenFile("myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if fileErr != nil {
        log.Fatal(fileErr)
    }

    log.SetOutput(file)

	database.InitDB("forum.db")

	http.Handle("/", middleware.OptionalSessionMiddleware(http.HandlerFunc(handler.IndexHandler)))
	http.Handle("/login", middleware.OptionalSessionMiddleware(http.HandlerFunc(handler.LoginHandler)))
	http.Handle("/register", middleware.OptionalSessionMiddleware(http.HandlerFunc(handler.RegisterHandler)))
	http.Handle("/create_post", middleware.SessionMiddleware(http.HandlerFunc(handler.CreateHandler)))
	http.Handle("/like", middleware.SessionMiddleware(http.HandlerFunc(handler.LikePost)))
	http.Handle("/dislike", middleware.SessionMiddleware(http.HandlerFunc(handler.DislikePost)))
	http.Handle("/logout", middleware.SessionMiddleware(http.HandlerFunc(handler.Logout)))
	http.Handle("/post", middleware.OptionalSessionMiddleware(http.HandlerFunc(handler.PostHandler)))
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
	database.CleanUpPosts()
	database.UpdateSession()

	fmt.Println("Server started at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server at 8080", err)
	}
}
