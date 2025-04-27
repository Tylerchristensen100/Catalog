package database

import (
	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func GetAllCips(app *internal.App) []models.Cip {
	var cips []models.Cip

	app.DB.Find(&cips)

	return cips
}

func GetCipById(app *internal.App, id string) models.Cip {
	var cip models.Cip

	app.DB.First(&cip, id)

	return cip
}

func GetCipByCode(app *internal.App, code float64) models.Cip {
	var cip models.Cip

	app.DB.First(&cip, models.Cip{Cip: code})

	return cip
}

func UpdateCip(app *internal.App, cip models.Cip) (models.Cip, error) {
	err := app.DB.Save(&cip).Error

	return cip, err
}

func CreateCip(app *internal.App, cip models.Cip) (models.Cip, error) {
	err := app.DB.Where(models.Cip{Name: cip.Name, Cip: cip.Cip}).FirstOrCreate(&cip).Error

	return cip, err
}
