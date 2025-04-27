package admin

import (
	"encoding/json"
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/models"
)

func POSTSchool(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed POST School endpoint", "username", authCtx.Username)
		err := req.ParseMultipartForm(10 << 20)

		if err != nil {
			http.Error(res, "Error parsing form", http.StatusBadRequest)
			app.Log.Error("Error parsing form", "error", err)
			return
		}

		code := req.FormValue("code")
		if code == "" {
			http.Error(res, "Code code is required", http.StatusBadRequest)
			return
		}

		name := req.FormValue("name")
		if name == "" {
			http.Error(res, "Name is required", http.StatusBadRequest)
			return
		}

		school := models.School{
			Name: name,
			Code: code,
		}

		s, err := database.CreateSchool(app, school)
		if err != nil {
			http.Error(res, "Error creating School", http.StatusInternalServerError)
			app.Log.Error("Error creating School", "error", err)
			return
		}

		json, err := json.Marshal(s)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			app.Log.Error("Error marshalling JSON", "error", err)
			return
		}

		res.WriteHeader(http.StatusCreated)
		res.Header().Set("Content-Type", "application/json")
		res.Write(json)
	}
}
