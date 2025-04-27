package api

import (
	"encoding/json"
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/database"
)

func GETGradLevels(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		gradLevels := database.GetAllGradLevels(app)

		json, err := json.Marshal(gradLevels)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			return
		}

		res.Header().Set("Content-Type", "application/json")

		res.Write(json)
	}
}
