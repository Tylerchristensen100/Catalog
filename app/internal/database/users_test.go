package database

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/models"
)

func TestGetAllUsers(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testUsers := generateTestUsers()

	users := GetAllUsers(&app)
	if len(users) != len(testUsers) {
		t.Errorf("Expected %d users, but got %d", len(testUsers), len(users))
	}
	for i, user := range users {
		if user.Username != testUsers[i].Username {
			t.Errorf("Expected user %s, but got %s", testUsers[i].Username, user.Username)
		}
	}
	t.Logf("Found %d users", len(users))

}

func TestGetUserByUsername(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testUsers := generateTestUsers()

	users := GetUserByUsername(&app, testUsers[0].Username)
	if users.Username != testUsers[0].Username {
		t.Errorf("Expected user %s, but got %s", testUsers[0].Username, users.Username)
	}
	t.Logf("Found user %s", users.Username)
}

func TestGetUserById(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testUsers := generateTestUsers()

	users := GetUserById(&app, testUsers[0].ID)
	if users.ID != testUsers[0].ID {
		t.Errorf("Expected user %d, but got %d", testUsers[0].ID, users.ID)
	}
	t.Logf("Found user %d", users.ID)
}

func TestGetUserByClientId(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}
	testUsers := generateTestUsers()

	user := GetUserByClientId(&app, testUsers[0].ClientID)
	if user.ClientID != testUsers[0].ClientID {
		t.Errorf("Expected user %d, but got %d", testUsers[0].ClientID, user.ClientID)
	}
	t.Logf("Found user %d", user.ClientID)
}

func TestCreateUser(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	testUser := models.User{
		Username: fmt.Sprintf("testuser %d", rand.Int31()),
		Roles:    "user",
		ClientID: rand.Int31(),
	}

	err := CreateUser(&app, testUser)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	newUser := GetUserByUsername(&app, testUser.Username)
	if newUser.Username != testUser.Username {
		t.Errorf("Expected user %s, but got %s", testUser.Username, newUser.Username)
	}

	t.Logf("Created user %s", testUser.Username)
}

func TestUpdateUser(t *testing.T) {
	app := internal.App{
		DB: SetupTestDB(t),
	}

	testUsers := generateTestUsers()
	testRole := "testRole"

	testUsers[0].Roles = testRole
	err := UpdateUser(&app, testUsers[0])
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	updatedUser := GetUserByClientId(&app, testUsers[0].ClientID)
	if updatedUser.Roles != testRole {
		t.Errorf("Expected user %s to be updated to %s, but got %s", testUsers[0].Username, testRole, updatedUser.Roles)
	}

	t.Logf("Updated user %s to %s", testUsers[0].Username, testRole)
}
