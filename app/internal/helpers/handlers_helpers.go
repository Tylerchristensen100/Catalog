package helpers

import (
	"testing"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/models"
)

func GetUserFromUsername(app *internal.App, auth *auth.AuthInfo) (models.User, error) {
	var user models.User

	if (testing.Testing() && auth == nil) || app.Config.Development {
		return models.User{
			ID:       999,
			Username: "10802612@uvu.edu",
			Roles:    "admin,faculty,student",
			ClientID: 10802612,
		}, nil
	}
	result := app.DB.First(&user, models.User{Username: auth.Username})
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
