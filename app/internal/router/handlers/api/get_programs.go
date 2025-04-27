package api

import (
	"encoding/json"
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/database"
)

func GETPrograms(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		programs := database.GetAllPrograms(app)

		json, err := json.Marshal(programs)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			return
		}
		res.Header().Set("Content-Type", "application/json")

		res.Write(json)
	}
}
