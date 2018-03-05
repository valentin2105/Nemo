// Package middlewares provides common middleware handlers.
package middlewares

import (
	"net/http"

	"context"

	"github.com/gorilla/sessions"
)

type key uint8

// SetSessionStore - store session
func SetSessionStore(sessionStore sessions.Store) func(http.Handler) http.Handler {
	var key key
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			req = req.WithContext(context.WithValue(req.Context(), key, sessionStore))
			next.ServeHTTP(res, req)
		})
	}
}

// MustLogin is a middleware that checks existence of current user.
func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		sessionStore := req.Context().Value("sessionStore").(sessions.Store)
		session, _ := sessionStore.Get(req, "Nemo-session")
		userRowInterface := session.Values["user"]

		if userRowInterface == nil {
			http.Redirect(res, req, "/login", 302)
			return
		}

		next.ServeHTTP(res, req)
	})
}
