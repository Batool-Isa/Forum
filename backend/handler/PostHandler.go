package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/structs"
	"log"
	"net/http"
	"strconv"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	session := middleware.FromContext(r.Context())
	postId := r.URL.Query().Get("id")

	if postId == "" {
		log.Println("Post ID is missing in the query parameters")
		http.Error(w, "Bad Request: Post ID is missing", http.StatusBadRequest)
		return
	}

	post_id, err := strconv.Atoi(postId)
	if err != nil {
		log.Printf("Invalid Post ID '%s': %v", postId, err)
		http.Error(w, "Bad Request: Invalid Post ID", http.StatusBadRequest)
		return
	}

	log.Printf("Fetching post with ID: %d", post_id)
	post, err := database.GetPostById(post_id)
	if err != nil {
		log.Printf("Error fetching post with ID %d: %v", post_id, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Post    structs.Post
		Session *structs.Session
	}{
		Post:    post,
		Session: session,
	}
	RenderTemplate(w, "comment.html", data)
}
