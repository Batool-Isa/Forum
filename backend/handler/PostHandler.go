package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/structs"
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	session := middleware.FromContext(r.Context())
	postId := r.URL.Query().Get("id")

	if postId == "" {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	post_id, err := strconv.Atoi(postId)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	log.Printf("Fetching post with ID: %d", post_id)
	post, err := database.GetPostById(post_id)
	if err != nil {
		if err==sql.ErrNoRows{
			ErrorHandler(w, r, http.StatusNotFound)
			return
		}
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	data := struct {
		Post    structs.Post
		Session *structs.Session
	}{
		Post:    post,
		Session: session,
	}
	RenderTemplate(w, r, "comment.html", data)
}
