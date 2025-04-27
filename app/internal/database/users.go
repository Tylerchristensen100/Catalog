package database

import (
	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func GetAllUsers(app *internal.App) []models.User {
	var users []models.User

	app.DB.Find(&users)

	return users
}

func GetUserById(app *internal.App, id int32) models.User {
	var user models.User

	app.DB.First(&user, id)

	return user
}

func GetUserByClientId(app *internal.App, clientId int32) models.User {
	var user models.User

	err := app.DB.Where(models.User{ClientID: clientId}).First(&user).Error
	if err != nil {
		app.Log.Error("Error fetching user by client ID", "error", err.Error())
	}

	return user
}

func GetUserByUsername(app *internal.App, username string) models.User {
	var user models.User

	app.DB.Where(models.User{Username: username}).First(&user)

	return user
}

func CreateUser(app *internal.App, user models.User) error {
	result := app.DB.Create(&user)

	return result.Error
}

func UpdateUser(app *internal.App, user models.User) error {
	err := app.DB.Where(&models.User{ClientID: user.ClientID}).Save(&user).Error
	if err != nil {
		app.Log.Error("Error updating user", "error", err.Error())
	}
	return err
}
