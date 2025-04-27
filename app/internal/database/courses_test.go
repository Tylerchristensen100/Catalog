package database

import (
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func TestGetAllCourses(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}

	testCourses := generateTestCourses()
	var courses []models.Course
	err := app.DB.Find(&courses).Error
	if err != nil {
		t.Errorf("Failed to find courses: %v", err)
	}

	if len(courses) < len(testCourses) {
		t.Errorf("Expected to over find %d courses, but found %d", len(testCourses), len(courses))
	}

	t.Logf("Found %d courses", len(courses))
}

func TestGetCourseByID(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testCourses := generateTestCourses()

	var course models.Course
	app.DB.First(&course, models.Course{ID: 1})

	if course.Name != testCourses[0].Name {
		t.Errorf("Expected course name to be 'Introduction to Programming', but got '%s'", course.Name)
	}

	if course.Code != testCourses[0].Code {
		t.Errorf("Expected course code to be 101, but got '%s'", course.Code)
	}

	if course.CreditHours != testCourses[0].CreditHours {
		t.Errorf("Expected course credit hours to be 3, but got %d", course.CreditHours)
	}

	if course.MajorCode != testCourses[0].MajorCode {
		t.Errorf("Expected course major code to be 'CS', but got '%s'", course.MajorCode)
	}

	t.Logf("Found course: %s", course.Name)
}

func TestUpdateCourse(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	var course models.Course
	app.DB.First(&course, models.Course{ID: 1})

	course.Name = "New Name"
	course.Code = "New Code"
	course.CreditHours = 4

	updatedCourse, err := UpdateCourse(&app, course)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updatedCourse.Name != "New Name" {
		t.Errorf("Expected course name to be 'New Name', but got '%s'", updatedCourse.Name)
	}

	if updatedCourse.Code != "New Code" {
		t.Errorf("Expected course code to be 'New Code', but got '%s'", updatedCourse.Code)
	}
}

func TestCreateCourse(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testCourse := models.Course{
		Name:        "Test Course",
		Code:        "TC101",
		CreditHours: 3,
		MajorCode:   "CS",
	}

	createdCourse, err := CreateCourse(&app, testCourse)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if createdCourse.Name != "Test Course" {
		t.Errorf("Expected course name to be 'Test Course', but got '%s'", createdCourse.Name)
	}

	if createdCourse.Code != "TC101" {
		t.Errorf("Expected course code to be 'TC101', but got '%s'", createdCourse.Code)
	}
}
