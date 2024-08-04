package handler

import (
	"Forum/backend/database"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Create a new session
var (
	// Map for storing session information
	sessions = make(map[string]bool)
	// Mutex to synchronize access to session information
	sessionMutex = &sync.Mutex{}
)

func CreateSession(w http.ResponseWriter, u_id int) error{

	dbSession, s_err := database.GetSessionByUser(u_id)
	if s_err == nil && dbSession.Timestamp.After(time.Now()) {
		fmt.Println("User has already a valid session")
		s_err = errors.New("user has already")
		return s_err
	} 


	sessionID, err := generateSessionID()

	if err != nil {
		fmt.Println("Error generating session ID: ", err)
		return err
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
	//	MaxAge: 600000, // Set the expiration time to 600 seconds (10 minute)
	})

	database.InsertSession(sessionID, u_id)
	fmt.Println(sessionID)
	return nil
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

// key type is unexported to prevent collisions
type key int

const (
	SessionKey key = iota
)

// sessionhandler is a handler function that checks for session expiry
func Sessionhandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session cookie
		sessionCookie, err := r.Cookie("session_id")
		if err != nil {
			// No session cookie, assume session has expired
			//utils.ErrorHandler(w, r, http.StatusUnauthorized)
			fmt.Println("No active session cookie")
			next.ServeHTTP(w, r)
			return
		}

		// Fetch the session using GetSession function
		session, err := database.GetSession(sessionCookie.Value)
		if err != nil {
			utils.ErrorHandler(w, r, http.StatusUnauthorized)
			// Handle errors (e.g., session expired or not found)
			fmt.Println("Session error:", err)
			next.ServeHTTP(w, r)
			return
		}

		// Check if the session is expired
		if isSessionExpired(session.Session) {
			fmt.Println("Session has expired")
			next.ServeHTTP(w, r)
			return
		}

		// Store session in context
		ctx := context.WithValue(r.Context(), SessionKey, session)
		fmt.Println("Session is active")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// isSessionExpired checks if the given session has expired
func isSessionExpired(sessionID string) bool {

	// Retrieve the session data from your session store
	session, err := database.GetSession(sessionID)

	fmt.Println(time.Now().After(session.Timestamp))
	if err != nil {
		return true // No session
	}

	// Check if the session has expired based on the expiration time
	return time.Now().After(session.Timestamp)
}

// FromContext retrieves the session from the context
func FromContext(ctx context.Context) *structs.Session {
	val := ctx.Value(SessionKey)
	if val == nil {
		return nil
	}
	session, ok := val.(structs.Session)
	if !ok {
		return nil
	}
	return &session
}
