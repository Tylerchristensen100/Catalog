package admin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"catalog.tylerChristensen/internal/database/mock_db"
	"catalog.tylerChristensen/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestPOSTSchools(t *testing.T) {
	DB, _ := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testSchool := models.School{
		Name: "Test School",
		Code: "TS",
	}

	form := url.Values{}
	form.Add("name", testSchool.Name)
	form.Add("code", testSchool.Code)

	req, err := http.NewRequest("POST", "/api/admin/school", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := POSTSchool(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var actualCourse models.Course
	err = json.Unmarshal(rr.Body.Bytes(), &actualCourse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, testSchool.Name, actualCourse.Name)
	assert.Equal(t, testSchool.Code, actualCourse.Code)
	assert.NotZero(t, actualCourse.ID, "Expected ID to be set after creation")
}
