package userRoutes

import (
	userHandler "github.com/azizanhakim/items-api-fiber/internal/handler/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	// Create an User
	user.Post("/", userHandler.CreateUser)

	// Read all Users
	user.Get("/", userHandler.GetUsers)

	// Read one User
	user.Get("/:userId", userHandler.GetUser)

	// Update one User
	user.Put("/:userId", userHandler.UpdateUser)

	// Delete one User
	user.Delete("/:userId", userHandler.DeleteUser)

}
