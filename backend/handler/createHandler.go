package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/utils"
	"net/http"
	"fmt"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	session := middleware.FromContext(r.Context())
	if session == nil {
		utils.ErrorHandler(w, r, http.StatusUnauthorized)
		//http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("http.Redi1354151jrect")
		return
	}
	if r.Method == "GET" {
		utils.RenderTemplate(w, r, "create_post.html", session)
		return
	}
	if r.Method == "POST" {
		r.ParseForm()
		title := r.Form.Get("title")
		content := r.Form.Get("content")
		category := r.Form["category"]

		uid, err := GetLoggedUser(r)
		if err != nil {
			utils.ErrorHandler(w, r, http.StatusInternalServerError)
			//http.Error(w, "Unable to retrieve user ID", http.StatusInternalServerError)

			return
		}

		database.InsertPost(uid, title, content, category)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	utils.ErrorHandler(w, r, http.StatusMethodNotAllowed)

}
