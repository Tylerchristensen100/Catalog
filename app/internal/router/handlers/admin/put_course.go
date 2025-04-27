package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/helpers"
	"catalog.tylerChristensen/internal/models"
)

func PUTCourse(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed PUT Course endpoint", "username", authCtx.Username)

		id, err := strconv.ParseInt(req.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(res, "Error parsing ID", http.StatusBadRequest)
			app.Log.Error("Error parsing ID", "error", err)
			return
		}

		req.ParseMultipartForm(10 << 20)

		app.Log.Info(req.FormValue("name"))
		app.Log.Info(req.FormValue("major_code"))
		app.Log.Info(req.FormValue("code"))
		app.Log.Info(req.FormValue("credit_hours"))

		var currentCourse models.Course
		err = app.DB.Find(&currentCourse, models.Course{ID: int32(id)}).Error
		if err != nil {
			http.Error(res, "Error getting Course", http.StatusInternalServerError)
			app.Log.Error("Error getting Course", "error", err)
			return
		}

		name := req.FormValue("name")
		if name == "" {
			name = currentCourse.Name
		}

		majorCode := req.FormValue("major_code")
		if majorCode == "" {
			majorCode = currentCourse.MajorCode
		}

		creditHours, err := strconv.ParseInt(req.FormValue("credit_hours"), 10, 32)
		if creditHours < 0 || err != nil {
			http.Error(res, "Credit Hours is invalid", http.StatusBadRequest)
			return
		}

		createdUser, err := helpers.GetUserFromUsername(app, authCtx)
		if err != nil {
			http.Error(res, "Error parsing user ID", http.StatusBadRequest)
			app.Log.Error("Error parsing user ID", "error", err)
			return
		}

		code := req.FormValue("code")

		updatedCourse := models.Course{
			ID:          int32(id),
			Name:        name,
			MajorCode:   majorCode,
			Code:        code,
			CreditHours: int32(creditHours),
			CreatedBy:   int32(createdUser.ID),
		}

		c, err := database.UpdateCourse(app, updatedCourse)
		if err != nil {
			http.Error(res, "Error updating Course", http.StatusInternalServerError)
			app.Log.Error("Error updating Course", "error", err)
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
