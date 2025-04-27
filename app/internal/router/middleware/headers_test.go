package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"catalog.tylerChristensen/internal"
)

func TestHeaders(t *testing.T) {
	app := &internal.App{
		Config: internal.Config{
			TrustedOrigins: []string{"http:localhost:8080"},
		},
	}

	nextHandler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	})

	middlewareHandler := Headers(app, nextHandler)

	testCases := []struct {
		name            string
		originHeader    string
		expectedOrigin  string
		expectedVary    []string
		expectedReferer string
		expectedCreds   string
		expectedServer  string
	}{
		{
			name:            "trusted origin",
			originHeader:    "http://localhost:8080",
			expectedOrigin:  "http://localhost:8080",
			expectedVary:    []string{"origin", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
			expectedReferer: "origin-when-cross-origin",
			expectedCreds:   "true",
			expectedServer:  "Go",
		},
		{
			name:            "untrusted origin",
			originHeader:    "http://untrusted.com",
			expectedOrigin:  "",
			expectedVary:    []string{"origin", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
			expectedReferer: "origin-when-cross-origin",
			expectedCreds:   "true",
			expectedServer:  "Go",
		},
		{
			name:            "no origin header",
			originHeader:    "",
			expectedOrigin:  "",
			expectedVary:    []string{"origin", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
			expectedReferer: "origin-when-cross-origin",
			expectedCreds:   "true",
			expectedServer:  "Go",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Origin", tc.originHeader)

			res := httptest.NewRecorder()

			middlewareHandler.ServeHTTP(res, req)

			if res.Header().Get("Access-Control-Allow-Origin") != tc.expectedOrigin {
				t.Errorf("expected origin %q, got %q", tc.expectedOrigin, res.Header().Get("Access-Control-Allow-Origin"))
			}

			for _, v := range tc.expectedVary {
				if !contains(res.Header()["Vary"], v) {
					t.Errorf("expected Vary header to contain %q, but did not", v)
				}
			}

			if res.Header().Get("referrer-policy") != tc.expectedReferer {
				t.Errorf("expected referer-policy %q, got %q", tc.expectedReferer, res.Header().Get("referrer-policy"))
			}

			if res.Header().Get("Access-Control-Allow-Credentials") != tc.expectedCreds {
				t.Errorf("expected Access-Control-Allow-Credentials %q, got %q", tc.expectedCreds, res.Header().Get("Access-Control-Allow-Credentials"))
			}

			if res.Header().Get("Server") != tc.expectedServer {
				t.Errorf("expected Server %q, got %q", tc.expectedServer, res.Header().Get("Server"))
			}
		})
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
