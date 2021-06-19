package database

import (
	"admin/src/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic("Could not connect to the database")
	} else {
		fmt.Println("Connecting..... wait!!")
	}
	DB = database
	database.AutoMigrate(models.User{})
}
