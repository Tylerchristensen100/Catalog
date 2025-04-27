package database

import (
	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func GetAllSchools(app *internal.App) []models.School {
	var school []models.School

	app.DB.Find(&school)

	return school
}

func GetSchoolByID(app *internal.App, id string) models.School {
	var school models.School

	app.DB.First(&school, id)

	return school
}

func UpdateSchool(app *internal.App, school models.School) (models.School, error) {
	err := app.DB.Save(&school).Error

	return school, err
}

func CreateSchool(app *internal.App, school models.School) (models.School, error) {
	err := app.DB.Where(models.School{Name: school.Name, Code: school.Code}).FirstOrCreate(&school).Error

	return school, err
}
