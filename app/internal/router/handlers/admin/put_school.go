package admin

import (
	"encoding/json"
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/models"
)

func PUTSchool(app *internal.App) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed PUT School endpoint", "username", authCtx.Username)

		schoolCode := req.PathValue("code")
		if schoolCode == "" {
			http.Error(res, "Error parsing Code", http.StatusBadRequest)
			app.Log.Error("Error parsing Code", "error", schoolCode)
			return
		}

		req.ParseForm()

		var currentSchool models.School
		err := app.DB.Find(&currentSchool, models.School{Code: schoolCode}).Error
		if err != nil {
			http.Error(res, "Error getting Code", http.StatusInternalServerError)
			app.Log.Error("Error getting Code", "error", err)
			return
		}

		name := req.FormValue("name")
		if name == "" {
			name = currentSchool.Name
		}

		code := req.FormValue("code")
		if code == "" {
			code = currentSchool.Code
		}

		updatedSchool := models.School{
			ID:   currentSchool.ID,
			Name: name,
			Code: code,
		}

		s, err := database.UpdateSchool(app, updatedSchool)
		if err != nil {
			http.Error(res, "Error updating School", http.StatusInternalServerError)
			app.Log.Error("Error updating School", "error", err)
			return
		}
		json, err := json.Marshal(s)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			app.Log.Error("Error marshalling JSON", "error", err)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "application/json")
		res.Write(json)
	}
}
