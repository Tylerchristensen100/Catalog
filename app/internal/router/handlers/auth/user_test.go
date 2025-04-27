package auth_handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database/mock_db"
)

func TestUser(t *testing.T) {
	DB, _ := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)

	userInfo := &auth.AuthInfo{
		Name:     "Test User",
		Username: "test_user",
		Roles:    []string{"admin"},
		Verified: true,
	}

	req, err := http.NewRequest("GET", "/api/admin/user", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx, auth.AuthKey, userInfo)

	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := User(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, err := json.Marshal(userInfo)
	if err != nil {
		t.Fatalf("Failed to marshal expected response: %v", err)
	}

	if rr.Body.String() != string(expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}
