package auth

import (
	"context"
	"net/http"
	"strings"
	"time"
)

const (
	authHeader = "Authorization"
	bearer     = "Bearer "
)

type authContextKey string

const AuthKey authContextKey = "auth"

func RequireAuthorization(roles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get(authHeader)
			if accessToken == "" {
				http.Error(w, "Missing access token", http.StatusBadRequest)
				return
			}

			jwt := strings.TrimPrefix(accessToken, bearer)
			authInfo, err := VerifyJWT(jwt)

			if err != nil {
				http.SetCookie(w, &http.Cookie{
					Name:     "access_token",
					Value:    "",
					Expires:  time.Now().Add(-1 * time.Hour),
					HttpOnly: true,
					SameSite: http.SameSiteStrictMode,
					Secure:   true,
				})

				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if len(roles) > 0 {
				authorized := false
				for _, role := range roles {
					for _, userRole := range authInfo.Roles {
						if role == userRole {
							authorized = true
							break
						}
					}
				}
				if !authorized {
					http.Error(w, "You don't have the proper permissions to access this resource", http.StatusForbidden)
					return
				}

			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, AuthKey, authInfo)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func Context(ctx context.Context) *AuthInfo {
	auth, ok := ctx.Value(AuthKey).(*AuthInfo)
	if !ok {
		return nil
	}
	return auth
}
