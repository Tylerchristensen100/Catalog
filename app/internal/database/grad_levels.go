package database

import (
	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func GetAllGradLevels(app *internal.App) []models.GradLevel {
	var gradLevels []models.GradLevel

	app.DB.Find(&gradLevels)

	return gradLevels
}
