package handler

import (
	"Forum/backend/database"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
			RenderTemplate(w, "login.html", errors)
            return
		}

		CreateSession(w, user.Uid)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	RenderTemplate(w, "login.html")

}
