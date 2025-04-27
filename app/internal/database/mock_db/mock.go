package mock_db

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) (*gorm.DB, TestData) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	dropTables(db)

	database.AutomigrateDB(db)
	data := generateTestData(db)

	t.Cleanup(func() {
		dropTables(db)
	})
	return db, data
}

func SetupTestApp(db *gorm.DB) internal.App {
	ctx := context.Background()
	return internal.App{
		Log:     slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})),
		DB:      db,
		Context: ctx,
	}
}

func generateTestData(db *gorm.DB) TestData {
	testCips := GenerateTestCIPs()
	db.Create(&testCips)

	testMajors := GenerateTestMajors()
	db.Create(&testMajors)

	testCourses := GenerateTestCourses(testMajors)
	db.Create(&testCourses)

	testGradLevels := GenerateTestGradLevels()
	db.Create(&testGradLevels)

	testSchools := GenerateTestSchools()
	db.Create(&testSchools)

	testPrograms := GenerateTestPrograms(testGradLevels, testSchools)
	err := db.Create(&testPrograms).Error
	if err != nil {
		fmt.Printf("Error creating test programs: %v\n", err)
	}

	testUsers := GenerateTestUsers()
	err = db.Create(&testUsers).Error
	if err != nil {
		fmt.Printf("Error creating test users: %v\n", err)
	}
	return TestData{
		CIPs:       testCips,
		Courses:    testCourses,
		GradLevels: testGradLevels,
		Schools:    testSchools,
		Majors:     testMajors,
		Programs:   testPrograms,
		Users:      testUsers,
	}
}

func ptr(s string) *string {
	return &s
}

type TestData struct {
	CIPs       []models.Cip
	Courses    []models.Course
	GradLevels []models.GradLevel
	Schools    []models.School
	Majors     []models.Major
	Programs   []models.Program
	Users      []models.User
}

func dropTables(db *gorm.DB) {
	db.Migrator().DropTable(&models.Cip{}, &models.Course{}, &models.GradLevel{}, &models.School{}, &models.Major{}, &models.Program{}, &models.User{})

}
