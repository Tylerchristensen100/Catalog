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

func PUTCip(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed PUT CIP endpoint", "username", authCtx.Username)

		id, err := strconv.ParseInt(req.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(res, "Error parsing ID", http.StatusBadRequest)
			app.Log.Error("Error parsing ID", "error", err)
			return
		}

		req.ParseMultipartForm(10 << 20)

		var currentCip models.Cip
		err = app.DB.Find(&currentCip, models.Cip{ID: int32(id)}).Error
		if err != nil {
			http.Error(res, "Error getting CIP", http.StatusInternalServerError)
			app.Log.Error("Error getting CIP", "error", err)
			return
		}

		cipCode, err := strconv.ParseFloat(req.FormValue("cip"), 64)
		if err != nil {
			http.Error(res, "Error parsing CIP code", http.StatusBadRequest)
			app.Log.Error("Error parsing CIP code", "error", err)
			return
		}

		if cipCode <= 0 {
			cipCode = currentCip.Cip
		}

		name := req.FormValue("name")
		if name == "" {
			name = currentCip.Name
		}

		description := req.FormValue("description")
		if description == "" {
			description = currentCip.Description
		}

		jobs := req.FormValue("jobs")
		if jobs == "" {
			jobs = currentCip.Jobs
		}

		updatedCip := models.Cip{
			ID:          int32(id),
			Cip:         cipCode,
			Name:        name,
			Description: description,
			Jobs:        jobs,
		}

		c, err := database.UpdateCip(app, updatedCip)
		if err != nil {
			http.Error(res, "Error updating CIP", http.StatusInternalServerError)
			app.Log.Error("Error updating CIP", "error", err)
			return
		}
		json, err := json.Marshal(c)
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
