package admin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"fmt"

	"catalog.tylerChristensen/internal/database/mock_db"
	"catalog.tylerChristensen/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestPOSTProgram(t *testing.T) {
	DB, testData := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testProgram := models.Program{
		Name:        "Computer Science",
		MajorCode:   "CS",
		Description: "A program focused on computer science principles.",
		GradLevelID: testData.GradLevels[0].ID,
		ProgramType: "Undergraduate",
		Cip:         testData.CIPs[0].ID,
		SchoolID:    testData.Schools[0].ID,
		Online:      1,
		Campus:      1,
	}

	form := url.Values{}
	form.Add("name", testProgram.Name)
	form.Add("grad_level", fmt.Sprintf("%d", testProgram.GradLevelID))
	form.Add("school", fmt.Sprintf("%d", testProgram.SchoolID))
	form.Add("major_code", testProgram.MajorCode)
	form.Add("online", fmt.Sprintf("%d", testProgram.Online))
	form.Add("description", testProgram.Description)
	form.Add("cip", fmt.Sprintf("%d", testProgram.Cip))
	form.Add("campus", fmt.Sprintf("%d", testProgram.Campus))
	form.Add("program_type", testProgram.ProgramType)

	req, err := http.NewRequest("POST", "/api/admin/program", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := POSTProgram(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	t.Logf("Response: %s", rr.Body.String())

	var actualProgram models.Program
	err = json.Unmarshal(rr.Body.Bytes(), &actualProgram)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, testProgram.Name, actualProgram.Name)
	assert.Equal(t, testProgram.MajorCode, actualProgram.MajorCode)
	assert.Equal(t, testProgram.Description, actualProgram.Description)
	assert.Equal(t, testProgram.Cip, actualProgram.Cip)
	assert.Equal(t, testProgram.SchoolID, actualProgram.School.ID)
	assert.Equal(t, testProgram.Online, actualProgram.Online)
	assert.Equal(t, testProgram.Campus, actualProgram.Campus)
	assert.NotZero(t, actualProgram.ID, "Expected ID to be set after creation")

}
