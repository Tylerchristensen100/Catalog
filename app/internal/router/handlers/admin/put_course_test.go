package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"catalog.tylerChristensen/internal/database/mock_db"
	"catalog.tylerChristensen/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestPUTCourse(t *testing.T) {
	DB, testData := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testData.Courses[0].MajorCode = testData.Majors[0].Code
	testData.Courses[0].Name = "Updated Course Name"
	testData.Courses[0].CreditHours = 5
	testCourse := testData.Courses[0]

	form := url.Values{}
	form.Add("id", string(testCourse.ID))
	form.Add("name", testCourse.Name)
	form.Add("major_code", testCourse.MajorCode)
	form.Add("code", testCourse.Code)
	form.Add("credit_hours", fmt.Sprintf("%d", testCourse.CreditHours))
	form.Add("description", testCourse.Description)

	req, err := http.NewRequest("PUT", "/api/admin/course/"+fmt.Sprintf("%d", testCourse.ID), bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	router := http.NewServeMux()
	router.HandleFunc("/api/admin/course/{id}", PUTCourse(&app))
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

	assert.Equal(t, testCourse.ID, actualCourse.ID)
	assert.Equal(t, testCourse.Name, actualCourse.Name)
	assert.Equal(t, testCourse.MajorCode, actualCourse.MajorCode)
	assert.Equal(t, testCourse.Code, actualCourse.Code)
	assert.Equal(t, testCourse.CreditHours, actualCourse.CreditHours)
	assert.Equal(t, testCourse.Description, actualCourse.Description)

}
