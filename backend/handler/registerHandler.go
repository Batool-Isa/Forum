package handler

import (
	"Forum/backend/database"
	"Forum/backend/middleware"
	"Forum/backend/utils"
	"log"
	"net/http"
	"regexp"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type FormData struct {
	Username string
	Email    string
	Errors   map[string]string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	session := middleware.FromContext(r.Context())
	if session != nil {
		utils.ErrorHandler(w, r, http.StatusSeeOther)
		return
	}
	formData := FormData{
		Username: "",
		Email:    "",
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	confirmPassword := r.Form.Get("confirm-password")
	email := r.Form.Get("email")
	errors := make(map[string]string)

	allUsers, err := database.GetAllUsers()
	if err != nil {
		log.Println("Error getting users")
	}

	errors = ValidateUser(username, email, password, confirmPassword)

	for _, user := range allUsers {
		if user[0] == username {
			errors["username"] = "Username already exists"
			break
		}
		if user[1] == email {
			errors["email"] = "Email already exists"
			break
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if len(errors) > 0 {
		formData.Username = username
		formData.Email = email
		utils.RenderTemplate(w, r, "login.html", formData, errors)
		return
	}

	err = database.InsertUser(username, string(hashedPassword), email)
	if err != nil {
		errors["email"] = "Email already exists"
	}

	user, err := database.GetUser(username)
	if err != nil {
		utils.ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	
	err1 := CreateSession(w, user.Uid)
	if err1 != nil {
		utils.ErrorHandler(w, r, http.StatusConflict)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// Inputs Validation
func ValidateUser(username string, email string, password string, confirmPass string) map[string]string {
	validationErrors := make(map[string]string)

	if !isValidUsername(username) {
		validationErrors["username"] = "Username should be 3-20 alphanumeric, _, -"
	}
	if !isValidEmail(email) {
		validationErrors["email"] = "Invalid email format"
	}
	if len(password) < 8 && len(password) > 1 {
		validationErrors["password"] = "Password must be at least 8 characters long"
	}
	if password != confirmPass {
		fmt.Println("doesn't match")
		validationErrors["confPassword"] = "Password and Confirm Password does not match"
	}

	return validationErrors
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return regex.MatchString(username) 
}