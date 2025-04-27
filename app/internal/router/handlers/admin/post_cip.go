package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/models"
)

func POSTCip(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed POST CIP endpoint", "username", authCtx.Username)
		err := req.ParseForm()
		if err != nil {
			http.Error(res, "Error parsing form", http.StatusBadRequest)
			app.Log.Error("Error parsing form", "error", err)
			return
		}

		cipString := req.FormValue("cip")
		if cipString == "" {
			http.Error(res, "CIP code is required", http.StatusBadRequest)
			return
		}

		cipCode, err := strconv.ParseFloat(cipString, 64)
		if err != nil {
			http.Error(res, "CIP code must be a number", http.StatusBadRequest)
			app.Log.Error("CIP code must be a number", "error", err)
			return
		}

		name := req.FormValue("name")
		if name == "" {
			http.Error(res, "Name is required", http.StatusBadRequest)
			return
		}

		description := req.FormValue("description")
		if description == "" {
			http.Error(res, "Description is required", http.StatusBadRequest)
			return
		}
		jobs := req.FormValue("jobs")
		if jobs == "" {
			http.Error(res, "Jobs is required", http.StatusBadRequest)
			return
		}

		cip := models.Cip{
			Cip:         cipCode,
			Name:        name,
			Description: description,
			Jobs:        jobs,
		}

		c, err := database.CreateCip(app, cip)
		if err != nil {
			http.Error(res, "Error creating CIP", http.StatusInternalServerError)
			app.Log.Error("Error creating CIP", "error", err)
			return
		}

		json, err := json.Marshal(c)
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
