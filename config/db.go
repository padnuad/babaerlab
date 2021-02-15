package config

import (
	"baberlab/domain"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialDatabase() error {
	var connectErr error

	// connectionString := fmt.Sprintf("host=localhost port=5433 user=postgres dbname=fillgoods-lab password=1234 sslmode=disable")
	// DB, connectErr = gorm.Open("postgres", connectionString)
	if connectErr != nil {
		fmt.Println(connectErr)
		return connectErr
	}

	autoMigrateDatabase()

	fmt.Println("Database connection successfully...")
	return nil
}

func autoMigrateDatabase() {
	DB.AutoMigrate(&domain.User{}) // &Owner{},
	// &Shop{},
	// &Booking{}
}
