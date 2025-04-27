package database

import (
	"fmt"
	"testing"

	"catalog.tylerChristensen/internal/models"
	"golang.org/x/exp/rand"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	AutomigrateDB(db)
	generateTestData(db)

	t.Cleanup(func() {
		db.Migrator().DropTable(&models.Cip{}, &models.Course{}, &models.GradLevel{}, &models.School{}, &models.Major{}, &models.Program{}, &models.User{})
	})
	return db
}

func generateTestData(db *gorm.DB) {
	testCips := generateTestCIPs()
	db.Create(&testCips)

	testCourses := generateTestCourses()
	db.Create(&testCourses)

	testGradLevels := generateTestGradLevels()
	db.Create(&testGradLevels)

	testSchools := generateTestSchools()
	db.Create(&testSchools)

	testMajors := generateTestMajors()
	db.Create(&testMajors)

	testPrograms := generateTestPrograms()
	db.AutoMigrate(&models.Program{}) // Ensure the Program table is created before inserting data
	err := db.Create(&testPrograms).Error
	if err != nil {
		fmt.Printf("Error creating test programs: %v\n", err)
	}

	testUsers := generateTestUsers()
	err = db.Create(&testUsers).Error
	if err != nil {
		fmt.Printf("Error creating test users: %v\n", err)
	}
}

func ptr(s string) *string {
	return &s
}

func generateTestCIPs() []models.Cip {
	return []models.Cip{
		{Cip: 12.34567, Name: "CIP 1", Description: "Description 1", Jobs: "Job 1, Job 2"},
		{Cip: 98.76543, Name: "CIP 2", Description: "Description 2", Jobs: "Job 3"},
		{Cip: 10.00000, Name: "CIP 3", Description: "Description 3", Jobs: "Job 4, Job 5, Job 6"},
		{Cip: 1.23456, Name: "CIP 4", Description: "Description 4", Jobs: "Job 7"},
	}
}

func generateTestCourses() []models.Course {
	return []models.Course{
		{Name: "Introduction to Programming", MajorCode: "CS", Code: "101", CreditHours: 3, CreatedBy: 1},
		{Name: "Data Structures and Algorithms", MajorCode: "CS", Code: "201", CreditHours: 4, CreatedBy: 1},
		{Name: "Calculus I", MajorCode: "MA", Code: "150", CreditHours: 4, CreatedBy: 2},
		{Name: "Linear Algebra", MajorCode: "MA", Code: "220", CreditHours: 3, CreatedBy: 2},
		{Name: "Principles of Economics", MajorCode: "EC", Code: "101", CreditHours: 3, CreatedBy: 4},
	}
}

func generateTestGradLevels() []models.GradLevel {
	return []models.GradLevel{
		{Level: ptr("Undergraduate")},
		{Level: ptr("Graduate")},
	}
}

func generateTestMajors() []models.Major {
	return []models.Major{
		{Code: "CS", Name: "Computer Science"},
		{Code: "MBA", Name: "Master of Business Administration"},
		{Code: "DS", Name: "Data Science"},
		{Code: "WD", Name: "Web Development"},
	}
}

func generateTestPrograms() []models.Program {
	return []models.Program{
		{
			Name:        "Computer Science",
			GradLevelID: 1,
			GradLevel:   models.GradLevel{Level: ptr("Undergraduate")},
			ProgramType: "Degree",
			School:      models.School{ID: 1, Name: "College of Engineering & Technology"},
			MajorCode:   "CS",
			Online:      1,
			Campus:      0,
			Description: "A comprehensive CS program.",
			Cip:         11,
			CreatedBy:   1,
		},
		{
			Name:        "Master of Business Administration",
			GradLevel:   models.GradLevel{Level: ptr("Graduate")},
			ProgramType: "Degree",
			School:      models.School{ID: 2, Name: "Woodbury school of business"},
			MajorCode:   "MBA",
			Online:      0,
			Campus:      1,
			Description: "A challenging MBA program.",
			Cip:         52,
			CreatedBy:   2,
		},
		{
			Name:        "Data Science Certificate",
			GradLevel:   models.GradLevel{Level: ptr("Certificate")},
			ProgramType: "Certificate",
			School:      models.School{ID: 1, Name: "College of Engineering & Technology"},
			MajorCode:   "DS",
			Online:      1,
			Campus:      0,
			Description: "A focused data science program.",
			Cip:         27,
			CreatedBy:   1,
		},
		{
			Name:        "Web Development Bootcamp",
			GradLevel:   models.GradLevel{Level: ptr("Certificate")},
			ProgramType: "Bootcamp",
			School:      models.School{ID: 1, Name: "College of Engineering & Technology"},
			MajorCode:   "WD",
			Online:      1,
			Campus:      0,
			Description: "Intensive web dev training.",
			Cip:         15,
			CreatedBy:   1,
		},
	}
}

func generateTestUsers() []models.User {
	rand.Seed(42)
	return []models.User{
		{Username: "user1", Roles: "admin,faculty", ClientID: rand.Int31()},
		{Username: "user2", Roles: "faculty", ClientID: rand.Int31()},
		{Username: "user3", Roles: "faculty", ClientID: rand.Int31()},
		{Username: "user4", Roles: "undefined", ClientID: rand.Int31()},
		{Username: "user5", Roles: "admin", ClientID: rand.Int31()},
	}
}

func generateTestSchools() []models.School {
	return []models.School{
		{Code: "CET", Name: "College of Engineering & Technology"},
		{Code: "WSB", Name: "School of Business"},
		{Code: "SOA", Name: "School of the Arts"},
	}
}
