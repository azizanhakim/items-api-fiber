package typeRoutes

import (
	typeHandler "github.com/azizanhakim/items-api-fiber/internal/handler/type"
	"github.com/gofiber/fiber/v2"
)

func SetupTypeRoutes(router fiber.Router) {
	tipe := router.Group("/type")

	// Create an Type
	tipe.Post("/", typeHandler.CreateTypes)

	// Read all Types
	tipe.Get("/", typeHandler.GetTypes)

	// Read one Type
	tipe.Get("/:typeId", typeHandler.GetType)

	// Update one Type
	tipe.Put("/:typeId", typeHandler.UpdateType)

	// Delete one Type
	tipe.Delete("/:typeId", typeHandler.DeleteType)

}
