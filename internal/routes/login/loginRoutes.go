package loginRoutes

import (
	loginHandler "github.com/azizanhakim/items-api-fiber/internal/handler/login"
	"github.com/gofiber/fiber/v2"
)

func SetupLoginRoutes(router fiber.Router) {
	login := router.Group("/login")

	login.Get("/", loginHandler.Login)
}
