package middleware

import (
	"Forum/backend/database"
	"fmt"
	"net/http"
	"time"
)


// sessionMiddleware is a middleware function that checks for session expiry
func SessionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the session cookie
        session, err := r.Cookie("session_id")
        if err != nil {
            // No session cookie, assume session has expired
            fmt.Println("No active session cookie")
            return
        }
        // Check if the session has expired
        if isSessionExpired(session.Value) {
            fmt.Println("session has expired")
            return
        }
        // Session is valid, pass the request to the next handler
        next.ServeHTTP(w, r)
    })
}

// isSessionExpired checks if the given session has expired
func isSessionExpired(sessionID string) bool {

    // Retrieve the session data from your session store
    session, err := database.GetSession(sessionID)
	fmt.Println(session)

    if err != nil {
        return true	// No session
    }

    // Check if the session has expired based on the expiration time
    return time.Now().After(session.Timestamp)
}

