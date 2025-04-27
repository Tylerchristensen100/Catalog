package handlers

import (
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/models"
)

func GetSchools(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		schools := database.GetAllSchools(app)

		app.Templ.ExecuteTemplate(res, "schools-page", schools)
	}
}

func GetSchoolsByCode(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		code := req.PathValue("code")
		var school models.School
		app.DB.Where("code = ?", code).First(&school)

		res.Header().Set("Content-Type", "text/html")
		app.Templ.ExecuteTemplate(res, "school-page", school)
	}
}
