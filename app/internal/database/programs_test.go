package database

import (
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func TestGetAllPrograms(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testPrograms := generateTestPrograms()
	var programs []models.Program

	err := app.DB.Find(&programs).Error
	if err != nil {
		t.Errorf("Failed to find programs: %v", err)
	}

	if len(programs) != len(testPrograms) {
		t.Errorf("Expected to find %d programs, but found %d", len(testPrograms), len(programs))
	}

	t.Logf("Found %d programs", len(programs))
}

func TestGetProgramByID(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	var program models.Program

	app.DB.Where(models.Program{ID: 1}).Find(&program)

	if program.Name != "Computer Science" {
		t.Errorf("Expected program name to be 'Computer Science', but got '%s'", program.Name)
	}

	if program.School.ID != 1 {
		t.Errorf("Expected program school to be 1, but got %d", program.School.ID)
	}

	t.Logf("Found program: %s", program.Name)
}

func TestUpdateProgram(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	var program models.Program

	app.DB.Where(models.Program{ID: 1}).Find(&program)

	program.Name = "New Name"
	program.School.ID = 2

	updatedProgram, err := UpdateProgram(&app, program)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updatedProgram.Name != "New Name" {
		t.Errorf("Expected program name to be 'New Name', but got '%s'", updatedProgram.Name)
	}

	if updatedProgram.School.ID != 2 {
		t.Errorf("Expected program school to be 2, but got %d", updatedProgram.School.ID)
	}

	t.Logf("Updated program: %s", updatedProgram.Name)
}

func TestCreateProgram(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testProgram := models.Program{
		Name: "New Program",
		School: models.School{
			ID: 1,
		},
	}

	program, err := CreateProgram(&app, testProgram)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if program.Name != "New Program" {
		t.Errorf("Expected program name to be 'New Program', but got '%s'", program.Name)
	}

	if program.School.ID != 1 {
		t.Errorf("Expected program school to be 1, but got %d", program.School.ID)
	}

	t.Logf("Created program: %s", program.Name)
}
