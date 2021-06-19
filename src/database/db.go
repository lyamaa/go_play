package database

import (
	"admin/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	DB, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/root"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(models.User{})
}
