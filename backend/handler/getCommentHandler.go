package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	session := middleware.FromContext(r.Context())
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		postID := r.FormValue("post_id")
		commentText := r.FormValue("comment")
		fmt.Println(commentText)

		uid, err := GetLoggedUser(r)
		if err != nil {
			http.Error(w, "Unable to retrieve user ID", http.StatusInternalServerError)
			return
		}
		postIDInt, err := strconv.Atoi(postID)
		fmt.Println("Post ID:", postIDInt)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}
		database.InsertComment(commentText, uid, postIDInt)
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postIDInt), http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
