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

func POSTCourse(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed POST Course endpoint", "username", authCtx.Username)

		err := req.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(res, "Error parsing form", http.StatusBadRequest)
			app.Log.Error("Error parsing form", "error", err)
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

		majorCode := req.FormValue("major_code")
		if majorCode == "" {
			http.Error(res, "Major Code is required", http.StatusBadRequest)
			return
		}

		creditHours, err := strconv.ParseInt(req.FormValue("credit_hours"), 10, 32)
		if creditHours < 0 || err != nil {
			http.Error(res, "Credit Hours is required", http.StatusBadRequest)
			return
		}

		createdUser, err := helpers.GetUserFromUsername(app, authCtx)
		if err != nil {
			http.Error(res, "Error parsing user ID", http.StatusBadRequest)
			app.Log.Error("Error parsing user ID", "error", err)
			return
		}

		code := req.FormValue("code")
		prerequisites := req.FormValue("prerequisites")

		course := models.Course{
			Name:          name,
			MajorCode:     majorCode,
			Code:          code,
			CreditHours:   int32(creditHours),
			Description:   description,
			Prerequisites: prerequisites,
			CreatedBy:     int32(createdUser.ID),
		}

		c, err := database.CreateCourse(app, course)
		if err != nil {
			http.Error(res, "Error creating course", http.StatusInternalServerError)
			app.Log.Error("Error creating course", "error", err)
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
