package handler

import (
	"Forum/backend/database"
	"fmt"
	// "Forum/backend/structs"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type FormData struct {
	Username string
	Email    string
	Errors   map[string]string
}


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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

	errors = ValidateUser(email, password, confirmPassword)

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
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	
	err = database.InsertUser(username, string(hashedPassword), email)
	if err != nil {
		errors["email"] = "Email already exists"
	}

	if len(errors) > 0 {
		formData.Username = username
		formData.Email = email
        RenderTemplate(w, "login.html", formData, errors)
        return
    }

	user, err := database.GetUser(username)
	if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    } 

	CreateSession(w, user.Uid)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}


// Inputs Validation
func ValidateUser(email string, password string, confirmPass string) map[string]string {
    validationErrors := make(map[string]string)

    // Validate email
	if !isValidEmail(email) {
		fmt.Println("Please enter a valid email")
        validationErrors["email"] = "Invalid email format"
    }

    // Validate password
	if len(password) < 8 {
		fmt.Println("Password must be at least 8 characters long")
        validationErrors["password"] = "Password must be at least 8 characters long"
    }

	if password != confirmPass {
		fmt.Println("Password and Confirm Password does not match")
		validationErrors["confPassword"] = "Password and Confirm Password does not match"
	}

    return validationErrors
}

func isValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}