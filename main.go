package main

import (
	"github.com/azizanhakim/items-api-fiber/database"
	loginHandler "github.com/azizanhakim/items-api-fiber/internal/login"
	"github.com/azizanhakim/items-api-fiber/router"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Start the fiber app
	app := fiber.New()

	database.ConnectDB()

	router.SetupRouter(app)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    loginHandler.PrivateKey.Public(),
		},
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	// Listen on PORT 3000
	app.Listen(":3000")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
