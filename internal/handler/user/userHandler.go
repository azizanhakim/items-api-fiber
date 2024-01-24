package userHandler

import (
	"fmt"

	"github.com/azizanhakim/items-api-fiber/database"
	"github.com/azizanhakim/items-api-fiber/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Please check again your input.", "data": err})
	}

	if user.Username == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Please input username.", "data": nil})
	}

	if user.Password == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Please input password.", "data": nil})
	}

	// Check if a user with the same username already exists
	var existingUser model.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		// A user with the same username already exists
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Username already exists.", "data": nil})
	}

	err = user.SetPassword(user.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not hashed the password", "data": err})
	}

	// Create the user and return error if encountered
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User created", "data": user})
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	// Find all users in the database
	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	// Read the param userId
	id := c.Params("userId")

	// Find user by user ID in the database
	db.Find(&user, "id = ?", id)
	fmt.Println("your ID:", user.ID)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Fullname string `json:"fullname"`
	}

	db := database.DB
	var user model.User

	// Read the param
	id := c.Params("userId")

	// Find the user by Id
	db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}

	var updateUser UpdateUser

	err := c.BodyParser(&updateUser)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Please check again your input", "data": err})
	}

	user.Username = updateUser.Username
	user.Password = updateUser.Password
	user.Fullname = updateUser.Fullname

	// hashed the new password
	err = user.SetPassword(user.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not hashed the new password", "data": err})
	}

	err = db.Save(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User successfuly updated", "data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("userId")

	var user model.User

	db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User is not present", "data": nil})
	}

	err := db.Delete(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": user})
}
