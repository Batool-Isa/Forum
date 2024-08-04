package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session := middleware.FromContext(r.Context())
	if session != nil {
		utils.ErrorHandler(w, r, http.StatusSeeOther)
		return
	}
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
			formData.Username = username
			utils.RenderTemplate(w, r, "login.html", formData, errors)
			return
		}

		err1 := CreateSession(w, user.Uid)
		if err1 != nil {
            utils.ErrorHandler(w, r, http.StatusConflict)
            return
        }
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	utils.RenderTemplate(w, r, "login.html", session)
}
