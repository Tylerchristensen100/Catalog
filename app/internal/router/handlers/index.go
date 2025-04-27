package handlers

import (
	"net/http"

	"catalog.tylerChristensen/internal"
)

func GetIndex(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Content-Type", "text/html")
		app.Templ.ExecuteTemplate(res, "index-page", nil)
	}
}
