package auth_handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"catalog.tylerChristensen/internal/database/mock_db"
	"golang.org/x/oauth2"
)

func TestCallback(t *testing.T) {
	DB, _ := mock_db.SetupTestDB(t)
	app := mock_db.SetupTestApp(DB)
	app.Config.Oauth = oauth2.Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://auth.example.com/authorize",
			TokenURL: "http://token.example.com/token",
		},
		RedirectURL: "http://redirect.example.com/callback",
		Scopes:      []string{"testScope1", "testScope2"},
	}

	mockState := "11111"
	verifierMap[mockState] = "test-verifier"

	url := "/callback?code=12345&state=" + mockState
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := Callback(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status < 300 && status >= 400 {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	cookieHeaders := rr.Result().Cookies()
	if len(cookieHeaders) != 1 {
		t.Fatalf("Expected 1 cookie header, got %d", len(cookieHeaders))
	}
	cookie := cookieHeaders[0]
	if cookie.Name != "access_token" {
		t.Errorf("Cookie name is not 'access_token': %s", cookie.Name)
	}
	if cookie.Value == "" {
		t.Error("Cookie value is empty")
	}

}
