package controllers

import (
	"admin/src/database"
	"admin/src/models"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var products models.Product

	if err := c.BodyParser(&products); err != nil {
		return err
	}

	database.DB.Create(&products)

	return c.JSON(products)
}
