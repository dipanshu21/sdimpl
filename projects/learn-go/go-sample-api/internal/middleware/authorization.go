package middleware

import (
	"api/api"
	"api/internal/tools"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var ErrUnauthorized = errors.New("invalid or missing authorization token or username")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		username := r.URL.Query().Get("username")
		if authHeader == "" || username == "" {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		// Here you would typically validate the token.
		// For simplicity, we assume any non-empty token is valid.
		var database tools.DatabaseInterface
		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := database.GetUserLoginDetails(username)

		if loginDetails == nil || loginDetails.Token != authHeader {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
