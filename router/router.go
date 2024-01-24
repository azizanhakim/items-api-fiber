package router

import (
	dashboardRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/dashboard"
	itemRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/item"
	loginRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/login"
	typeRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/type"
	userRoutes "github.com/azizanhakim/items-api-fiber/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupUnauthenticatedRouter(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	itemRoutes.SetupItemRoutes(api)
	typeRoutes.SetupTypeRoutes(api)
	loginRoutes.SetupLoginRoutes(api)
	userRoutes.SetupUserRoutes(api)
}

func SetupAuthenticatedRouter(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	dashboardRoutes.SetupDashboardRoutes(api)
}
