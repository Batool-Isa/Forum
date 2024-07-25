package handler

import (
	"Forum/backend/database"
	"Forum/backend/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category")

	var posts []structs.Post
	var err error

	if categoryID == "" || categoryID == "all" {
		posts, err = database.GetAllPosts()
	} else {
		catID, err := strconv.Atoi(categoryID)
		if err != nil {
			fmt.Println("herrrrrrrrrrrrrrrrrrrrr")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		posts, err = database.GetPostsByCategory(catID)
	}

	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
