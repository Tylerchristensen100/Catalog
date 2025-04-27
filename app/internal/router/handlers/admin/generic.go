package admin

import (
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
)

func Generic(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed generic endpoint", "id", authCtx.Username)

		res.WriteHeader(http.StatusNotFound)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"message": "Catalog Admin API's"}`))
	}
}
