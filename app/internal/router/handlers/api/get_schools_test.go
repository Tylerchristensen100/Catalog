package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"catalog.tylerChristensen/internal/database/mock_db"
)

func TestGETSchools(t *testing.T) {
	DB, testData := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	req, err := http.NewRequest("GET", "/api/schools", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := GETSchools(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := testData.Schools
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to marshal expected Schools: %v", err)
	}
	if rr.Body.String() != string(expectedJSON) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expectedJSON))
	}
}
