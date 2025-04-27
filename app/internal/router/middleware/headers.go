package middleware

import (
	"net/http"

	"catalog.tylerChristensen/internal"
)

func Headers(app *internal.App, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Vary", "origin")
		res.Header().Add("Vary", "Access-Control-Request-Method")
		res.Header().Add("Vary", "Access-Control-Request-Headers")

		origin := req.Header.Get("Origin")

		if origin != "" {
			for i := range app.Config.TrustedOrigins {
				if origin == app.Config.TrustedOrigins[i] {
					res.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		res.Header().Set("referrer-policy", "origin-when-cross-origin")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Add("Server", "Go")
		next.ServeHTTP(res, req)
	})
}
