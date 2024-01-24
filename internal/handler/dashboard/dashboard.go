package dashboardHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Dashboard(c *fiber.Ctx) error {
	// Access the token from the context
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok || userToken == nil {
		// Handle the case when the user token is nil
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
	}

	if userToken == nil {
		// Handle the case when the token is nil
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized access",
			"data":    nil,
		})
	}

	claims := userToken.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.SendString("Welcome to the dashboard, " + name)
}
