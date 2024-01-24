package dashboardRoutes

import (
	dashboardHandler "github.com/azizanhakim/items-api-fiber/internal/handler/dashboard"
	"github.com/gofiber/fiber/v2"
)

func SetupDashboardRoutes(router fiber.Router) {
	dashboard := router.Group("/dashboard")

	dashboard.Get("/", dashboardHandler.Dashboard)

}
