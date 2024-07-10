package handler

import (
	"Forum/backend/database"
	"fmt"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		RenderTemplate(w, "create_post.html")
		return
	}
	if r.Method == "POST" {
		r.ParseForm()
		title := r.Form.Get("title")
		content := r.Form.Get("content")
		category := r.Form["category"]

		fmt.Println(category)

		database.InsertPost(1, title, content, category)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
