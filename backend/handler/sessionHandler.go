package handler
import (
	"Forum/backend/database"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"

)

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
	  MaxAge: 600000, // Set the expiration time to 600 seconds (10 minute)
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

func GetLoggedUser(r *http.Request) (int, error) {
	sessionCookie, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println("No active session cookie")
		return 0, err
	}

	// Fetch the session using GetSession function
	session, err := database.GetSession(sessionCookie.Value)
	if err != nil {
		// Handle errors (e.g., session expired or not found)
		fmt.Println("Session error:", err)
		return 0, err
	}
	
	return session.UserID, nil
}