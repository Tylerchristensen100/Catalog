package api

import (
	"net/http"

	"catalog.tylerChristensen/internal"
)

func GETGeneric(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"message": "Welcome to the Catalog API", "documentation": "./api/docs"}`))
	}
}
