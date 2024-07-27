package handler

import (
	"Forum/backend/database"
	"fmt"
	"net/http"
	"strconv"
	// "structs"
)

// func GetUserIDFromCookie(r *http.Request) (int, error) {
// 	cookie, err := r.Cookie("user_id")
// 	if err != nil {
// 		return 0, err
// 	}

// 	userID, err := strconv.Atoi(cookie.Value)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return userID, nil
// }

func LikePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LikePost")
	if r.Method != http.MethodPost {
		ErrorHandler(w,r,http.StatusMethodNotAllowed)
		//http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("post_id")
	fmt.Println(postID)
	if postID == "" {
		ErrorHandler(w,r,http.StatusBadRequest)
		//http.Error(w, "Missing post_id", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(postID)
	if err != nil {
		ErrorHandler(w,r,http.StatusBadRequest)
		//http.Error(w, "Invalid post_id", http.StatusBadRequest)
		return
	}

	uid, err := GetLoggedUser(r)
	if err != nil {
		ErrorHandler(w,r,http.StatusForbidden)
		//ErrorHandler(w,r,http.StatusInternalServerError)
		//http.Error(w, "Unable to retrieve user ID", http.StatusInternalServerError)
		
		return
	}


	
	// Insert into likes table
	database.InsertLikes(pid, uid)
	database.DeleteDislike(pid, uid)
	database.UpdatePost(pid)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

