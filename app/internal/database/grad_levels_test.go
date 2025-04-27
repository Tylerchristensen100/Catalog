package database

import (
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func TestGetAllGradLevels(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testGradLevels := generateTestGradLevels()

	var levels []models.GradLevel
	err := app.DB.Find(&levels).Error
	if err != nil {
		t.Errorf("Failed to find GradLevels: %v", err)
	}

	if len(levels) < len(testGradLevels) {
		t.Error("Expected to find GradLevels, but found none")
	}

	if len(levels) != len(testGradLevels) {
		t.Errorf("Expected to find %d GradLevels, but found %d", len(testGradLevels), len(levels))
	}

	t.Logf("Found %d GradLevels", len(levels))
}
