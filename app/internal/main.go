package internal

import (
	"context"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type App struct {
	Router      http.Handler
	DB          *gorm.DB
	Log         *slog.Logger
	Context     context.Context
	Config      Config
	Templ       *template.Template
	Auth        *middleware.Interceptor[*oauth.IntrospectionContext]
	OauthConfig *oauth2.Config
}

type Config struct {
	TrustedOrigins []string
	Development    bool
	Domain         string
	AuthDomain     string
	AuthKey        string
	Oauth          oauth2.Config
}

func (app *App) Serve() {
	server := &http.Server{
		Addr:     ":3000",
		Handler:  app.Router,
		ErrorLog: slog.NewLogLogger(app.Log.Handler(), slog.LevelError),
	}

	app.Log.Info("Starting Server on", slog.String("address", server.Addr))
	err := server.ListenAndServe()
	if err != nil {
		app.Log.Error("Error starting server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
