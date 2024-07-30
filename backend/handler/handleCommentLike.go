package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/structs"
	"fmt"
	"net/http"
	"strconv"
	// "structs"
)

func LikeComment(w http.ResponseWriter, r *http.Request) {
	sessionValue, ok := r.Context().Value(middleware.SessionKey).(structs.Session)
	if !ok {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	fmt.Println("active",sessionValue)
	if r.Method != http.MethodPost {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("post_id")
	postIDInt, err := strconv.Atoi(postID)
		if err != nil {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

	commentId := r.FormValue("comment_id")
	if commentId == "" {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	cid, err := strconv.Atoi(commentId)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	uid, err := GetLoggedUser(r)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Insert into dislikes table
	database.InsertCommentLikes(cid, uid)
	database.DeleteCommentDislike(cid, uid)
	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postIDInt), http.StatusSeeOther)
}
