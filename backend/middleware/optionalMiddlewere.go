package middleware

import (
	"Forum/backend/database"
	"context"
	"net/http"
	
)

// OptionalSessionMiddleware checks for session and adds it to context if available
func OptionalSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session_id")
		if err != nil {
			next.ServeHTTP(w, r) // No session, proceed without it
			return
		}

		session, err := database.GetSession(sessionCookie.Value)
		if err != nil || isSessionExpired(session.Session) {
			next.ServeHTTP(w, r) // Invalid or expired session, proceed without it
			return
		}

		// Store session in context
		ctx := context.WithValue(r.Context(), SessionKey, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


