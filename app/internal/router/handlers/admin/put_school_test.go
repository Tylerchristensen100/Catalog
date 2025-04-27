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

func TestPUTSchools(t *testing.T) {
	DB, testData := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testData.Schools[0].Name = "Updated School Name"
	testData.Schools[0].Code = "TEST"
	testSchool := testData.Schools[0]

	form := url.Values{}
	form.Add("id", string(testSchool.ID))
	form.Add("name", testSchool.Name)
	form.Add("code", testSchool.Code)

	req, err := http.NewRequest("PUT", "/api/admin/school/"+testSchool.Code, bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	router := http.NewServeMux()
	router.HandleFunc("/api/admin/school/{code}", PUTSchool(&app))
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
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
