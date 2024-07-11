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

	RenderTemplate(w, "login.html")

}


// Create a new session
var (
	// Map for storing session information
	sessions = make(map[string]bool)
	// Mutex to synchronize access to session information
	sessionMutex = &sync.Mutex{}
)

func CreateSession(w http.ResponseWriter, u_id int) {
	sessionID, err := generateSessionID()

	if err != nil {
		fmt.Println("Error generating session ID: ", err)
		return
	 }

	 // Save the session on the server side
	 sessionMutex.Lock()
	 sessions[sessionID] = true
	 for id := range sessions {
	  fmt.Println("loginHandler: Current session ID is ", id)
	 }
	 sessionMutex.Unlock()
	
	 // Send the session ID to the client as a Cookie
	 http.SetCookie(w, &http.Cookie{
	  Name:   "session_id",
	  Value:  sessionID,
	  Path:   "/",
	  MaxAge: 60, // Set the expiration time to 60 seconds (1 minute)
	 })

	database.InsertSession(sessionID, u_id)
	fmt.Println(sessionID)
}

func generateSessionID() (string, error) {
    // Generate 32 random bytes
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }

    // Convert the bytes to a hexadecimal string
    return hex.EncodeToString(b), nil
}


