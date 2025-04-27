package auth_handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"catalog.tylerChristensen/internal/database/mock_db"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestLogin(t *testing.T) {
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
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := Login(&app)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status < 300 && status >= 400 {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	location, err := url.Parse(rr.Header().Get("Location"))
	assert.NoError(t, err)

	query := location.Query()
	state := query.Get("state")
	challenge := query.Get("code_challenge")

	assert.NotEmpty(t, state)
	assert.NotEmpty(t, challenge)

}
