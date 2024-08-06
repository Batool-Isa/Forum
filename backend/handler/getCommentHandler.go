package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/utils"
	"Forum/backend/structs"
	"fmt"
	"net/http"
	"strconv"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(middleware.SessionKey).(structs.Session)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		w.WriteHeader(http.StatusForbidden)
		return
		// utils.ErrorHandler(w, r, http.StatusForbidden)

		//http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		// return
	}
	if r.Method != http.MethodPost {
		utils.ErrorHandler(w, r, http.StatusMethodNotAllowed)

		//http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
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
		inserterr := database.InsertComment(commentText, uid, postIDInt)
		if inserterr != nil {
			utils.ErrorHandler(w, r, http.StatusBadRequest)
		}
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postIDInt), http.StatusSeeOther)
		return
	}
	utils.ErrorHandler(w, r, http.StatusMethodNotAllowed)
}
