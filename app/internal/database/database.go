package database

import (
	"fmt"
	"log"
	"time"

	"catalog.tylerChristensen/internal/models"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var retries = 0

func InitDB(username string, password string, address string) *gorm.DB {

	dsn := username + ":" + password + "@tcp(" + address + ")/catalog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		if retries < 15 {
			fmt.Println("Failed to connect to database, retrying in 5 seconds")
			fmt.Println("Login info being used is:  Username: ", username, " Password: ", password, "Address: ", address)

			time.Sleep(5 * time.Second)
			retries++

			return InitDB(username, password, address)
		}
		log.Fatal("failed to connect database")
	}
	AutomigrateDB(db)
	return db
}

func AutomigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Cip{})
	if err != nil {
		fmt.Println("Error migrating Cip:", err)
	}
	err = db.AutoMigrate(&models.Course{})
	if err != nil {
		fmt.Println("Error migrating Course:", err)
	}
	err = db.AutoMigrate(&models.School{})
	if err != nil {
		fmt.Println("Error migrating School:", err)
	}
	err = db.AutoMigrate(&models.GradLevel{})
	if err != nil {
		fmt.Println("Error migrating GradLevel:", err)
	}
	err = db.AutoMigrate(&models.Program{})
	if err != nil {
		fmt.Println("Error migrating Program:", err)
	}
	err = db.AutoMigrate(&models.Major{})
	if err != nil {
		fmt.Println("Error migrating Major:", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error migrating User:", err)
	}

}
