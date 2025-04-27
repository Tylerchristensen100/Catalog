package auth_handlers

import (
	"encoding/json"
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
)

func User(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		auth := auth.Context(req.Context())

		if auth == nil {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("Unauthorized"))
			return
		}

		json, err := json.Marshal(auth)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			return
		}

		res.Header().Set("Content-Type", "application/json")

		res.Write(json)
	}
}
