package handler

import (
	"Forum/backend/database"
	"fmt"
	"net/http"
	"strconv"
	// "structs"
)

func DislikePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LikePost")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("post_id")
	fmt.Println(postID)
	if postID == "" {
		http.Error(w, "Missing post_id", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "Invalid post_id", http.StatusBadRequest)
		return
	}

	uid, err := GetLoggedUser(r)
	if err != nil {
		http.Error(w, "Unable to retrieve user ID", http.StatusInternalServerError)
		
		return
	}

	// Insert into dislikes table
	database.InsertDislikes(pid, uid)
	database.DeleteLike(pid, uid)
	database.UpdatePost(pid)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
