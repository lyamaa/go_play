package controllers

import (
	"admin/src/database"
	"admin/src/models"

	"github.com/gofiber/fiber/v2"
)

func Order(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Find(&orders)

	return c.JSON(orders)

}
