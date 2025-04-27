package database

import (
	"math/rand"
	"testing"
	"time"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func TestGetAllCips(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testCips := generateTestCIPs()

	var cips []models.Cip
	err := app.DB.Find(&cips).Error
	if err != nil {
		t.Errorf("Failed to find CIPs: %v", err)
	}

	if len(cips) < len(testCips) {
		t.Error("Expected to find CIPs, but found none")
	}

	if len(cips) != len(testCips) {
		t.Errorf("Expected to find %d CIPs, but found %d", len(testCips), len(cips))
	}

	t.Logf("Found %d CIPs", len(cips))
}

func TestGetCipById(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}

	var cip models.Cip
	err := app.DB.Find(&cip).Where(models.Cip{ID: 1}).Error
	if err != nil {
		t.Errorf("Failed to find CIP with ID 1: %v", err)
	}
	if cip.Cip != 12.34567 {
		t.Errorf("Expected to find CIP code 12.34567, but found %f", cip.Cip)
	}

	t.Logf("Found CIP with ID 1: %s", cip.Name)
}

func TestGetCipByCode(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}

	var cip models.Cip
	err := app.DB.Find(&cip).Where(models.Cip{Cip: 12.34567}).Error
	if err != nil {
		t.Errorf("Failed to find CIP with code 12.34567: %v", err)
	}
	if cip.Cip != 12.34567 {
		t.Errorf("Expected to find CIP code 12.34567, but found %f", cip.Cip)
	}

	t.Logf("Found CIP with code 12.34567: %s", cip.Name)
}

func TestUpdateCip(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}

	var cip models.Cip
	err := app.DB.Find(&cip).Where(models.Cip{ID: 1}).Error
	if err != nil {
		t.Errorf("Failed to find CIP with ID 1: %v", err)
	}

	cip.Name = "New Name"
	cip.Cip = 12.34567

	updatedCip, err := UpdateCip(&app, cip)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if updatedCip.Name != "New Name" {
		t.Errorf("Expected CIP name to be 'New Name', but got '%s'", updatedCip.Name)
	}

	if updatedCip.Cip != 12.34567 {
		t.Errorf("Expected CIP code to be '12.34567', but got '%f'", updatedCip.Cip)
	}
}

func TestCreateCip(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))

	testCip := models.Cip{
		Name: "New CIP",
		Cip:  rand.Float64() * (20 - 5),
	}

	createdCip, err := CreateCip(&app, testCip)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if createdCip.Name != testCip.Name {
		t.Errorf("Expected CIP name to be 'New CIP', but got '%s'", createdCip.Name)
	}

	if createdCip.Cip != testCip.Cip {
		t.Errorf("Expected CIP code to be '12.34567', but got '%f'", createdCip.Cip)
	}
}
