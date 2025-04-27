package docs_handlers

import (
	"net/http"
	"path/filepath"

	"catalog.tylerChristensen/internal"
)

func SwaggerHTML(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := filepath.Abs("./internal/docs/swagger/index.html")
		if err != nil {
			app.Log.Error(err.Error())
		}

		http.ServeFile(w, r, html)

	}
}
