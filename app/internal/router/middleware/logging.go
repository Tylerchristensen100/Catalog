package middleware

import (
	"net/http"

	"catalog.tylerChristensen/internal"
)

func Logger(app *internal.App, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(res, req)

		app.Log.Info("Request:", req.Method, req.URL.Path)
	})
}
