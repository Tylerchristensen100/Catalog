package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"catalog.tylerChristensen/internal/database/mock_db"
	"catalog.tylerChristensen/internal/models"
)

func TestPUTCip(t *testing.T) {
	DB, testData := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testData.CIPs[0].Cip = 12345.67
	testData.CIPs[0].Jobs = "Software Developer, Data Scientist"

	testCip := testData.CIPs[0]

	form := url.Values{}
	form.Add("id", fmt.Sprintf("%d", testCip.ID))
	form.Add("cip", fmt.Sprintf("%f", testCip.Cip))
	form.Add("name", testCip.Name)
	form.Add("description", testCip.Description)
	form.Add("jobs", testCip.Jobs)

	req, err := http.NewRequest("PUT", "/api/admin/cip/"+fmt.Sprintf("%d", testCip.ID), bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	router := http.NewServeMux()
	router.HandleFunc("/api/admin/cip/{id}", PUTCip(&app))
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var actualCip models.Cip
	err = json.Unmarshal(rr.Body.Bytes(), &actualCip)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, testCip.ID, actualCip.ID)
	assert.Equal(t, testCip.Cip, actualCip.Cip)
	assert.Equal(t, testCip.Name, actualCip.Name)
	assert.Equal(t, testCip.Description, actualCip.Description)
	assert.Equal(t, testCip.Jobs, actualCip.Jobs)
}
