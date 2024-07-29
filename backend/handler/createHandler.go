package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"fmt"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	session := middleware.FromContext(r.Context())
	if session == nil {
		ErrorHandler(w, r, http.StatusUnauthorized)
		//http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		RenderTemplate(w, r, "create_post.html", session)
		return
	}
	if r.Method == "POST" {
		r.ParseForm()
		title := r.Form.Get("title")
		content := r.Form.Get("content")
		category := r.Form["category"]

		fmt.Println(category)
		uid, err := GetLoggedUser(r)
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			//http.Error(w, "Unable to retrieve user ID", http.StatusInternalServerError)

			return
		}

		database.InsertPost(uid, title, content, category)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	ErrorHandler(w, r, http.StatusMethodNotAllowed)

}
