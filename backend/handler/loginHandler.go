package handler

import (
	"Forum/backend/database"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	formData := FormData{
		Username: "",
	}
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		errors := make(map[string]string)

		user, err := database.GetUser(username)
		if err != nil {
			errors["user"] = "Invalid username or password"
		} 
			
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			errors["user"] = "Invalid username or password"
		}

		if len(errors) > 0 {
			fmt.Println(errors)
			formData.Username = username
			RenderTemplate(w, "login.html", formData, errors)
            return
		}

		CreateSession(w, user.Uid)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	RenderTemplate(w, "login.html", nil)
}
