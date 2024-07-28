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
		fmt.Println("Unable to retrieve session")
		http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
		return
	}
	fmt.Println("acctttiiivvveee",sessionValue)
	fmt.Println("LikeComment")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("post_id")
	postIDInt, err := strconv.Atoi(postID)
		fmt.Println("Post ID:", postIDInt)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

	commentId := r.FormValue("comment_id")
	fmt.Println(commentId)
	if commentId == "" {
		http.Error(w, "Missing comment_id", http.StatusBadRequest)
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
	fmt.Println("THE USER IS")
	fmt.Println(uid)

	// Insert into dislikes table
	database.InsertCommentLikes(cid, uid)
	database.DeleteCommentDislike(cid, uid)
	http.Redirect(w, r, fmt.Sprintf("/post?id=%d", postIDInt), http.StatusSeeOther)
}
