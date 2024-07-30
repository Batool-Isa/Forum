package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"fmt"
	"net/http"
	"strconv"
	// "structs"
)

func DislikeComment(w http.ResponseWriter, r *http.Request) {
	sessionValue, ok := r.Context().Value(middleware.SessionKey).(structs.Session)
	if !ok {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	fmt.Println("active", sessionValue)
	if r.Method != http.MethodPost {
		utils.ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	commentId := r.FormValue("comment_id")
	if commentId == "" {
		utils.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	postID := r.FormValue("post_id")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	cid, err := strconv.Atoi(commentId)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	uid, err := GetLoggedUser(r)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)

		return
	}

	// Insert into dislikes table
	database.InsertCommentDislikes(cid, uid)
	database.DeleteCommentLike(cid, uid)
	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postIDInt), http.StatusSeeOther)
}
