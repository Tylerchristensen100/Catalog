package database

import (
	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func GetAllPrograms(app *internal.App) []models.Program {
	var programs []models.Program

	app.DB.Preload("GradLevel").Preload("School").Order("name ASC").Find(&programs)

	return programs
}

func GetProgramByID(app *internal.App, id int) []models.Program {
	var programs []models.Program

	app.DB.Preload("GradLevel").Preload("School").Where("id = ?", id).Find(&programs)

	return programs
}

func UpdateProgram(app *internal.App, program models.Program) (models.Program, error) {
	err := app.DB.Save(&program).Error

	app.DB.Preload("GradLevel").Preload("School").Where("id = ?", program.ID).Find(&program)

	return program, err
}

func CreateProgram(app *internal.App, program models.Program) (models.Program, error) {
	err := app.DB.Where(
		models.Program{Name: program.Name, SchoolID: program.School.ID,
			Description: program.Description, GradLevelID: program.GradLevelID,
		}).FirstOrCreate(&program).Error

	return program, err
}
