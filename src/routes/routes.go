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

	// Middlewares
	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)

	// ADMIN AUTH ROUTES
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Put("users/info", controllers.ProfileUpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)

	// VENDOR ROUTES
	adminAuthenticated.Get("vendors", controllers.Vendor)

	// PRODUCT ROUTES
	adminAuthenticated.Get("products", controllers.Products)
	adminAuthenticated.Post("products", controllers.CreateProduct)
	adminAuthenticated.Get("products/:id", controllers.GetProduct)
	adminAuthenticated.Put("products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("products/:id", controllers.DeletePRoduct)

	// Link ROUTE
	adminAuthenticated.Get("users/:id/links", controllers.Link)

	// Order Routes
	adminAuthenticated.Get("orders", controllers.Orders)

	// VENDOR
	vendor := api.Group("vendor")
	// Vendor auth
	vendor.Post("register", controllers.Register)
	vendor.Post("login", controllers.Login)

	vendorAuthenticated := vendor.Use(middlewares.IsAuthenticated)

	// VENDOR AUTH ROUTES
	vendorAuthenticated.Get("user", controllers.User)
	vendorAuthenticated.Post("logout", controllers.Logout)
	vendorAuthenticated.Put("users/info", controllers.ProfileUpdateInfo)
	vendorAuthenticated.Put("users/password", controllers.UpdatePassword)

}
