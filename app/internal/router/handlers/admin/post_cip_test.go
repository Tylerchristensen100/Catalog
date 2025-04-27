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

func TestPOSTCip(t *testing.T) {
	DB, _ := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	testCip := models.Cip{
		ID:          1,
		Cip:         12.2688,
		Name:        "Test CIP",
		Description: "Test description",
		Jobs:        "Test jobs",
	}

	form := url.Values{}
	form.Add("cip", fmt.Sprintf("%f", testCip.Cip))
	form.Add("name", testCip.Name)
	form.Add("description", testCip.Description)
	form.Add("jobs", testCip.Jobs)

	req, err := http.NewRequest("POST", "/api/admin/cip", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := POSTCip(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var actualCip models.Cip
	err = json.Unmarshal(rr.Body.Bytes(), &actualCip)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, testCip.Cip, actualCip.Cip)
	assert.Equal(t, testCip.Name, actualCip.Name)
	assert.Equal(t, testCip.Description, actualCip.Description)
	assert.Equal(t, testCip.Jobs, actualCip.Jobs)
}
