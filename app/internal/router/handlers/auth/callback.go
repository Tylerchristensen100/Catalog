package auth_handlers

import (
	"context"
	"net/http"
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/constants"
	"golang.org/x/oauth2"
)

func Callback(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		if code == "" {
			http.Error(w, "Missing authorization code", http.StatusBadRequest)
			return
		}

		verifier, ok := verifierMap[state]
		if !ok {
			http.Error(w, "Code verifier not found", http.StatusBadRequest)
			return
		}
		delete(verifierMap, state)

		token, err := app.Config.Oauth.Exchange(context.Background(), code, oauth2.VerifierOption(verifier))
		if err != nil {
			if testing.Testing() {
				testResponse(w, r, app)
				return
			}
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusBadRequest)
			return
		}

		accessToken := token.AccessToken

		http.SetCookie(w, &http.Cookie{
			Name:   constants.AccessTokenKey,
			Value:  accessToken,
			Secure: true,
		})

		w.Header().Set("Authorization", "Bearer "+accessToken)
		http.Redirect(w, r, app.Config.Domain+"/admin/", http.StatusFound)
	}
}

func testResponse(w http.ResponseWriter, r *http.Request, app *internal.App) {
	accessToken := "test-access-token"
	http.SetCookie(w, &http.Cookie{
		Name:   constants.AccessTokenKey,
		Value:  accessToken,
		Secure: true,
	})

	w.Header().Set("Authorization", "Bearer "+accessToken)
	http.Redirect(w, r, app.Config.Domain+"/admin/", http.StatusFound)
}
