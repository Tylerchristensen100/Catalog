package database

import (
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func TestGetAllSchools(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testSchools := generateTestSchools()
	var school []models.School
	app.DB.Find(&school)

	if len(school) != len(testSchools) {
		t.Errorf("Expected to find %d schools, but found %d", len(testSchools), len(school))
	}

	t.Logf("Found %d schools", len(school))
}

func TestGetSchoolByID(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	var school models.School
	app.DB.First(&school, models.School{ID: 1})

	if school.Code != "CET" {
		t.Errorf("Expected school code to be 'CET', but got '%s'", school.Code)
	}

	if school.Name != "College of Engineering & Technology" {
		t.Errorf("Expected school name to be 'College of Engineering & Technology', but got '%s'", school.Name)
	}

	t.Logf("Found school: %s", school.Name)
}

func TestUpdateSchool(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	var school models.School
	app.DB.First(&school, models.School{ID: 1})

	school.Name = "New Name"
	school.Code = "New Code"

	updatedSchool, err := UpdateSchool(&app, school)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updatedSchool.Name != "New Name" {
		t.Errorf("Expected school name to be 'New Name', but got '%s'", updatedSchool.Name)
	}

	if updatedSchool.Code != "New Code" {
		t.Errorf("Expected school code to be 'New Code', but got '%s'", updatedSchool.Code)
	}
}

func TestCreateSchool(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testSchool := models.School{
		Name: "Test School",
		Code: "TS",
	}

	school, err := CreateSchool(&app, testSchool)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if school.Name != testSchool.Name {
		t.Errorf("Expected school name to be '%s', but got '%s'", testSchool.Name, school.Name)
	}

	if school.Code != testSchool.Code {
		t.Errorf("Expected school code to be '%s', but got '%s'", testSchool.Code, school.Code)
	}
}
