package controllers

import (
	"admin/src/database"
	"admin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var link []models.Link

	database.DB.Where("user_id=?", id).Find(&link)

	return c.JSON(link)
}
