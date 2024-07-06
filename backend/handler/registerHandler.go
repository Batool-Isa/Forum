package handler

import (
	"Forum/backend/database"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	//email := r.Form.Get("email")
	
	allUsers, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}

	for _, user := range allUsers {
		if user == username {
			http.Error(w, "Username already exist", http.StatusInternalServerError)
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	database.InsertUser(username, string(hashedPassword), "dummy@gmail.com")

	user, err := database.GetUser(username)
	if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    } 

	CreateSession(w, user.Uid)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}