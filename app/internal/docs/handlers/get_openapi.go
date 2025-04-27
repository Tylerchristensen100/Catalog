package docs_handlers

import (
	"net/http"
	"path/filepath"

	"catalog.tylerChristensen/internal"
)

func OpenAPI(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := filepath.Abs("./internal/docs/openapi.yaml")
		if err != nil {
			app.Log.Error(err.Error())
		}

		http.ServeFile(w, r, html)

	}
}
