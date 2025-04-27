package auth_handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/constants"
)

type logoutModel struct {
	AccessToken string `json:"access_token"`
	Logout      bool   `json:"logout"`
}

func Logout(app *internal.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessTokenCookie, _ := r.Cookie(constants.AccessTokenKey)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var logoutData logoutModel
		err = json.Unmarshal(body, &logoutData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		if accessTokenCookie.Value != logoutData.AccessToken {
			app.Log.Error("Access token in cookie does not match access token in request body")
		}

		if logoutData.Logout {
			http.SetCookie(w, &http.Cookie{
				Name:     constants.AccessTokenKey,
				Value:    "",
				Expires:  time.Now().Add(-1 * time.Hour),
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
				Secure:   true,
			})

			statusCode, err := auth.RevokeAccessToken(logoutData.AccessToken)
			if err != nil {
				app.Log.Error("Error revoking access token", "error", err)
				http.Error(w, "Error revoking access token", http.StatusInternalServerError)
				return
			}
			if statusCode != nil {
				app.Log.Info("Access token revoked", "status_code", *statusCode)
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Logout successful"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Logout Failed for unknown reason"))
	}
}
