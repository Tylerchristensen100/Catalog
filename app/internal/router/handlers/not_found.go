package handlers

import (
	"net/http"

	"catalog.tylerChristensen/internal"
)

func NotFound(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")

		app.Templ.ExecuteTemplate(res, "not-found-page", nil)
	}
}
