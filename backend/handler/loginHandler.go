package handler

import (
	"Forum/backend/database"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		user, err := database.GetUser(username)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
		} 
			
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		CreateSession(w, user.Uid)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	RenderTemplate(w, "login.html", nil)
}
