package itemRoutes

import (
	itemHandler "github.com/azizanhakim/items-api-fiber/internal/handler/item"
	"github.com/gofiber/fiber/v2"
)

func SetupItemRoutes(router fiber.Router) {
	item := router.Group("/item")

	// Create an Item
	item.Post("/", itemHandler.CreateItems)

	// Read all Item
	item.Get("/", itemHandler.GetItems)

	// Read one Item
	item.Get("/:itemId", itemHandler.GetItem)

	// Update one Item
	item.Put("/:itemId", itemHandler.UpdateItem)

	// Delete one Item
	item.Delete("/:itemId", itemHandler.DeleteItem)

}
