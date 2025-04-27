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

func TestPOSTCourse(t *testing.T) {
	DB, _ := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testCase := models.Course{
		Name:        "Valid Course",
		MajorCode:   "CS",
		Code:        "101",
		CreditHours: 3,
		Description: "A valid course description",
	}

	form := url.Values{}
	form.Add("name", testCase.Name)
	form.Add("major_code", testCase.MajorCode)
	form.Add("code", testCase.Code)
	form.Add("credit_hours", "3")
	form.Add("description", testCase.Description)

	req, err := http.NewRequest("POST", "/api/admin/course", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := POSTCourse(&app)

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

	assert.Equal(t, testCase.Name, actualCourse.Name)
	assert.Equal(t, testCase.MajorCode, actualCourse.MajorCode)
	assert.Equal(t, testCase.Code, actualCourse.Code)
	assert.Equal(t, testCase.CreditHours, actualCourse.CreditHours)
	assert.Equal(t, testCase.Description, actualCourse.Description)

}
