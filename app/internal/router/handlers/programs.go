package handlers

import (
	"net/http"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/models"
)

func GetPrograms(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")

		programs := database.GetAllPrograms(app)

		app.Templ.ExecuteTemplate(res, "programs-page", programs)
	}
}

func GetProgramsByName(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		name := req.PathValue("name")
		var program models.Program
		app.DB.Where("name = ?", name).First(&program)

		res.Header().Set("Content-Type", "text/html")
		app.Templ.ExecuteTemplate(res, "program-page", program)
	}
}
