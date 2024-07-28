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

func DislikeComment(w http.ResponseWriter, r *http.Request) {
	sessionValue, ok := r.Context().Value(middleware.SessionKey).(structs.Session)
	if !ok {
		fmt.Println("Unable to retrieve session")
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}
	fmt.Println("acctttiiivvveee",sessionValue)
	fmt.Println("DislikeComment")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	commentId := r.FormValue("comment_id")
	fmt.Println(commentId)
	if commentId == "" {
		http.Error(w, "Missing comment_id", http.StatusBadRequest)
		return
	}

	postID := r.FormValue("post_id")
	postIDInt, err := strconv.Atoi(postID)
	fmt.Println("Post ID:", postIDInt)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	cid, err := strconv.Atoi(commentId)
	if err != nil {
		http.Error(w, "Invalid commentId", http.StatusBadRequest)
		return
	}

	uid, err := GetLoggedUser(r)
	if err != nil {
		http.Error(w, "Unable to retrieve user ID", http.StatusInternalServerError)

		return
	}


	// Insert into dislikes table
	database.InsertCommentDislikes(cid, uid)
	database.DeleteCommentLike(cid, uid)
	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postIDInt), http.StatusSeeOther)
}
