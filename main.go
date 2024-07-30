package main

import (
	"Forum/backend/database"
	"Forum/backend/handler"
	"Forum/backend/middleware"
	"Forum/backend/utils"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	init_err := database.InitDB("forum.db")
	if init_err != nil {
		utils.ErrorHandler(nil, nil, http.StatusInternalServerError)
	}

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
	create_err := database.CreateTables()
	if create_err != nil {
		utils.ErrorHandler(nil, nil, http.StatusInternalServerError)
	}
	

	database.AddDummyData()
	database.CleanUpPosts()

	fmt.Println("Server started at http://localhost:3090/")
	err := http.ListenAndServe(":3090", nil)
	if err != nil {
		log.Fatal("Error starting server at 3090", err)
	}
}
