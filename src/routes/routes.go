package routes

import (
	"admin/src/controllers"
	"admin/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)
	// admin.Get("user", controllers.User)

	// Middleware

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Get("logout", controllers.Logout)
	adminAuthenticated.Put("users/info", controllers.ProfileUpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)
	adminAuthenticated.Get("vendors", controllers.Vendor)
	adminAuthenticated.Get("products", controllers.Products)
}
