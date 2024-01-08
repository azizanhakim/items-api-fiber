package router

import (
	itemRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/item"
	loginRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/login"
	typeRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/type"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	itemRoutes.SetupItemRoutes(api)
	typeRoutes.SetupTypeRoutes(api)
	loginRoutes.SetupLoginRoutes(api)
}
