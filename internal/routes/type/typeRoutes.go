package typeRoutes

import (
	typeHandler "github.com/azizanhakim/items-api-fiber/internal/handler/type"
	"github.com/gofiber/fiber/v2"
)

func SetupTypeRoutes(router fiber.Router) {
	tipe := router.Group("/type")

	// Create an Item
	tipe.Post("/", typeHandler.CreateTypes)

	// Read all Item
	tipe.Get("/", typeHandler.GetTypes)

	// Read one Item
	tipe.Get("/:typeId", typeHandler.GetType)

	// Update one Item
	tipe.Put("/:typeId", typeHandler.UpdateType)

	// Delete one Item
	tipe.Delete("/:typeId", typeHandler.DeleteType)

}
