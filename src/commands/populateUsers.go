package main

import (
	"admin/src/database"
	"admin/src/models"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		vendor := models.User{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			IsVendor:  true,
		}
		vendor.SetPassword("12345678")

		database.DB.Create(&vendor)

	}
}
