package main

import (
	"log"

	"github.com/azizanhakim/items-api-fiber/database"
	loginHandler "github.com/azizanhakim/items-api-fiber/internal/handler/login"
	"github.com/azizanhakim/items-api-fiber/router"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start the fiber app
	app := fiber.New()

	database.ConnectDB()

	// Unauthenticated route
	router.SetupUnauthenticatedRouter(app)

	// JWT Middleware
	app.Use(func(c *fiber.Ctx) error {
		// Configure JWT middleware
		config := jwtware.Config{
			SigningKey: jwtware.SigningKey{
				JWTAlg: jwtware.RS256,
				Key:    loginHandler.PublicKey,
			},
			SuccessHandler: func(c *fiber.Ctx) error {
				return c.Next()
			},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				log.Printf("JWT Middleware Error: %v", err)
				return c.SendStatus(fiber.StatusUnauthorized)
			},
		}

		return jwtware.New(config)(c)
	})

	// Authenticated route
	router.SetupAuthenticatedRouter(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
