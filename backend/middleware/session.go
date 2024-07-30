package middleware

import (
	"Forum/backend/database"
	//"Forum/backend/handler"
	"Forum/backend/utils"
	"Forum/backend/structs"
	"context"
	"fmt"
	"net/http"
	"time"
)

// key type is unexported to prevent collisions
type key int

const (
	SessionKey key = iota
)

// sessionMiddleware is a middleware function that checks for session expiry
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session cookie
		sessionCookie, err := r.Cookie("session_id")
		if err != nil {
			// No session cookie, assume session has expired
			// fmt.Println("No active session cookie")
			utils.ErrorHandler(w, r, http.StatusUnauthorized)
			next.ServeHTTP(w, r)
			return
		}

		// Fetch the session using GetSession function
		session, err := database.GetSession(sessionCookie.Value)
		if err != nil {
			// Handle errors (e.g., session expired or not found)
			// utils.ErrorHandler(w, r, http.StatusUnauthorized) // the error is not being handled in sessionHandler.go
			next.ServeHTTP(w, r)
			return
		}

		// Check if the session is expired
		if isSessionExpired(session.Session) {
			// fmt.Println("Session has expired")
			utils.ErrorHandler(w, r, http.StatusUnauthorized) //CHECK THIS ERROR
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
