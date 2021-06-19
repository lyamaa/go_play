package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/root"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
