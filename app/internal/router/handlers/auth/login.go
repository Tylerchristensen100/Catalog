package auth_handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"catalog.tylerChristensen/internal"
	"golang.org/x/oauth2"
)

var (
	verifierMap = make(map[string]string)
)

func Login(app *internal.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		verifier := oauth2.GenerateVerifier()
		state := generateState(app)
		verifierMap[state] = verifier

		url := app.Config.Oauth.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))

		http.Redirect(w, r, url, http.StatusFound)
	}
}

func generateState(app *internal.App) string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		app.Log.Error("Failed to generate random state", "error", err)
		return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	}
	return base64.StdEncoding.EncodeToString(b)
}
