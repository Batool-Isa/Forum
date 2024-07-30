package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/utils"
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
			utils.ErrorHandler(w, r, http.StatusBadRequest)
			return
		}
		postID := r.FormValue("post_id")
		commentText := r.FormValue("comment")

		uid, err := GetLoggedUser(r)
		if err != nil {
			utils.ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		postIDInt, err := strconv.Atoi(postID)
		if err != nil {
			utils.ErrorHandler(w, r, http.StatusBadRequest)
			return
		}
		database.InsertComment(commentText, uid, postIDInt)
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postIDInt), http.StatusSeeOther)
		return
	}
	utils.ErrorHandler(w, r, http.StatusMethodNotAllowed)
}
