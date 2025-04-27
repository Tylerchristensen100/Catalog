package admin

import (
	"encoding/json"
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database"
)

func GETProgram(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed GET Program endpoint", "username", authCtx.Username)

		programs := database.GetAllPrograms(app)

		json, err := json.Marshal(programs)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			app.Log.Error("Error marshalling JSON", "error", err)
			return
		}
		res.Header().Set("Content-Type", "application/json")

		res.Write(json)
	}
}
