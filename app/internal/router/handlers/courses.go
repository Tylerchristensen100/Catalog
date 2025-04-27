package handlers

import (
	"net/http"
	"strings"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/models"
)

func GetCourses(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		courses := database.GetAllCourses(app)

		app.Templ.ExecuteTemplate(res, "courses-page", courses)
	}
}

func GetCoursesByCourseCode(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		courseCode := strings.Split(req.PathValue("courseCode"), "-")

		var course models.Course
		app.DB.Where("major_code", courseCode[0]).Where("code = ?", courseCode[1]).First(&course)

		if course == (models.Course{}) {
			http.Redirect(res, req, "/404", http.StatusFound)
		}

		res.Header().Set("Content-Type", "text/html")
		app.Templ.ExecuteTemplate(res, "course-page", course)
	}
}
