package database

import (
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func GetAllCourses(app *internal.App) []models.Course {
	var courses []models.Course
	if testing.Testing() {
		app.DB.Preload("Major").Order("id ASC").Find(&courses)
	} else {
		app.DB.Preload("Major").Order("major_code ASC").Find(&courses)
	}

	return courses
}

func GetCourseByID(app *internal.App, id string) models.Course {
	var course models.Course
	app.DB.Preload("Major").First(&course, id)

	return course
}

func UpdateCourse(app *internal.App, course models.Course) (models.Course, error) {
	err := app.DB.Save(&course).Error

	app.DB.Preload("Major").Where("id = ?", course.ID).Find(&course)

	return course, err
}

func CreateCourse(app *internal.App, course models.Course) (models.Course, error) {
	err := app.DB.Where(
		models.Course{Name: course.Name, MajorCode: course.MajorCode,
			Description: course.Description, CreditHours: course.CreditHours,
		}).FirstOrCreate(&course).Error

	return course, err
}
