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

func TestPUTProgram(t *testing.T) {
	DB, testData := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testData.Programs[0].GradLevelID = testData.GradLevels[0].ID
	testData.Programs[0].SchoolID = testData.Schools[0].ID
	testData.Programs[0].Name = "Updated Program Name"
	testProgram := testData.Programs[0]

	form := url.Values{}
	form.Add("id", fmt.Sprintf("%d", testProgram.ID))
	form.Add("name", testProgram.Name)
	form.Add("grad_level", fmt.Sprintf("%d", testProgram.GradLevelID))
	form.Add("school", fmt.Sprintf("%d", testProgram.SchoolID))
	form.Add("major_code", testProgram.MajorCode)
	form.Add("online", fmt.Sprintf("%d", testProgram.Online))
	form.Add("description", testProgram.Description)
	form.Add("cip", fmt.Sprintf("%d", testProgram.Cip))
	form.Add("campus", fmt.Sprintf("%d", testProgram.Campus))
	form.Add("program_type", testProgram.ProgramType)

	req, err := http.NewRequest("PUT", "/api/admin/program/"+fmt.Sprintf("%d", testProgram.ID), bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	router := http.NewServeMux()
	router.HandleFunc("/api/admin/program/{id}", PUTProgram(&app))
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

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
