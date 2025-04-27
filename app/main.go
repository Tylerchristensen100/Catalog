package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"strings"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/constants"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/helpers"
	"catalog.tylerChristensen/internal/router"
	"github.com/coreos/go-oidc"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		debug := os.Getenv(constants.DEBUG_KEY)
		if debug == "true" {
			log.Fatalf("Error loading .env file: %v", err)
		} else {
			log.Printf("No .env file found in the parent directory")
		}
	}

	username := os.Getenv(constants.DB_USERNAME_KEY)
	password := os.Getenv(constants.DB_PASSWORD_KEY)
	address := os.Getenv(constants.DB_ADDRESS_KEY)
	if username == "" || password == "" || address == "" {
		log.Fatal("One or more environment variables are missing. " +
			"Please set the following environment variables: " +
			constants.DB_USERNAME_KEY + ", " +
			constants.DB_PASSWORD_KEY + ", " +
			constants.DB_ADDRESS_KEY)
	}

	flag.Parse()

	db := database.InitDB(username, password, address)

	tmpl, err := helpers.LoadTemplates("./templates/components", "./templates")
	if err != nil {
		slog.Error("Error loading templates", "error", err)
		log.Fatal(err)
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	config := config(log)
	app := internal.App{
		Log:     log,
		DB:      db,
		Context: context.Background(),
		Templ:   tmpl,
		Config:  config,
	}

	app.Router = router.Router(&app)

	app.Serve()
}

func config(log *slog.Logger) internal.Config {
	clientID, exists := os.LookupEnv(constants.OauthClientIDKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.OauthClientIDKey, "value", clientID)
		os.Exit(1)
	}
	clientSecret, exists := os.LookupEnv(constants.OauthClientSecretKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.OauthClientSecretKey, "value", clientSecret)
		os.Exit(1)
	}
	redirectURI, exists := os.LookupEnv(constants.OauthRedirectURIKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.OauthRedirectURIKey, "value", redirectURI)
		os.Exit(1)
	}
	scopesStr, exists := os.LookupEnv(constants.OauthScopesKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.OauthScopesKey, "value", scopesStr)
		os.Exit(1)
	}
	provider, exists := os.LookupEnv(constants.OauthProviderKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.OauthProviderKey, "value", provider)
		os.Exit(1)
	}
	authKeyPath, exists := os.LookupEnv(constants.AuthKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.AuthKey, "value", authKeyPath)
		os.Exit(1)
	}
	domain, exists := os.LookupEnv(constants.DomainKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.DomainKey, "value", domain)
		os.Exit(1)
	}
	trustedOrigins, exists := os.LookupEnv(constants.TrustedOriginsKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.TrustedOriginsKey, "value", trustedOrigins)
		os.Exit(1)
	}
	isDev, exists := os.LookupEnv(constants.DevelopmentKey)
	if !exists {
		log.Error("Missing required environment variable", "key", constants.DevelopmentKey, "value", isDev)
		os.Exit(1)
	}

	scopes := sliceFromString(scopesStr)
	scopes = append(scopes, oidc.ScopeOpenID)

	return internal.Config{
		TrustedOrigins: sliceFromString(trustedOrigins),
		Domain:         domain,
		Development:    isDev == "true",
		AuthDomain:     provider,
		AuthKey:        authKeyPath,
		Oauth: oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURI,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://" + provider + "/oauth/v2/authorize",
				TokenURL: "https://" + provider + "/oauth/v2/token",
			},
			Scopes: scopes,
		},
	}
}

func sliceFromString(str string) []string {
	return strings.Split(str, ",")
}
