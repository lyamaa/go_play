package database

import (
	"admin/src/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:root@/root"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	} else {
		fmt.Println("Connecting..... wait!!")
	}
	DB = database
	database.AutoMigrate(models.User{})
}
