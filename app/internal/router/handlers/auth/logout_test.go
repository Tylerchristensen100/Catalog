package auth_handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"catalog.tylerChristensen/internal/constants"
	"catalog.tylerChristensen/internal/database/mock_db"
)

func TestLogout(t *testing.T) {
	DB, _ := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	mockData := logoutModel{
		AccessToken: "test_access_token",
		Logout:      true,
	}

	jsonData, err := json.Marshal(mockData)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req, err := http.NewRequest("GET", "/api/logout", bytes.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  constants.AccessTokenKey,
		Value: "test_access_token",
	})

	rr := httptest.NewRecorder()
	handler := Logout(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	cookieHeaders := rr.Result().Cookies()
	if len(cookieHeaders) > 1 {
		t.Errorf("Expected less than 1 cookie header")
	}

}
