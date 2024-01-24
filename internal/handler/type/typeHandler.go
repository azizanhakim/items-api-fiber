package typeHandler

import (
	"fmt"

	"github.com/azizanhakim/items-api-fiber/database"
	"github.com/azizanhakim/items-api-fiber/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateTypes(c *fiber.Ctx) error {
	db := database.DB
	tipe := new(model.Type)

	// Store the body in the tipe and return error if encountered
	err := c.BodyParser(tipe)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Please check again your input.", "data": err})
	}

	if tipe.Name == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Please input type name.", "data": nil})
	}

	// Create the tipe and return error if encountered
	err = db.Create(&tipe).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create type", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Type created", "data": tipe})
}

func GetTypes(c *fiber.Ctx) error {
	db := database.DB
	var types []model.Type

	// Find all types in the database
	db.Find(&types)

	if len(types) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No type present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Types found", "data": types})
}

func GetType(c *fiber.Ctx) error {
	db := database.DB
	var tipe model.Type

	// Read the param typeId
	id := c.Params("typeId")

	// Find all types in the database
	db.Find(&tipe, "id = ?", id)
	fmt.Println("your ID:", tipe.ID)
	if tipe.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No type present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Type found", "data": tipe})
}

func UpdateType(c *fiber.Ctx) error {
	type UpdateType struct {
		Name    string `json:"name"`
		IsHeavy bool   `json:"isHeavy"`
	}

	db := database.DB
	var tipe model.Type

	// Read the param
	id := c.Params("typeId")

	db.Find(&tipe, "id = ?", id)

	if tipe.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "type not found", "data": nil})
	}

	var updateType UpdateType
	err := c.BodyParser(&updateType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Please check again your input", "data": err})
	}

	tipe.Name = updateType.Name
	tipe.IsHeavy = updateType.IsHeavy

	err = db.Save(&tipe).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update type", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Type successfuly updated", "data": tipe})

}

func DeleteType(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("typeId")

	var tipe model.Type

	db.Find(&tipe, "id = ?", id)
	if tipe.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Type is not present", "data": nil})
	}

	err := db.Delete(&tipe).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete type", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Type successfully deleted", "data": tipe})
}
