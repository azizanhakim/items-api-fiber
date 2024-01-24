package itemHandler

import (
	"github.com/azizanhakim/items-api-fiber/database"
	"github.com/azizanhakim/items-api-fiber/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateItems(c *fiber.Ctx) error {
	db := database.DB
	item := new(model.Item)

	// Store the body in the item and return error if encountered
	err := c.BodyParser(item)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Please check again your input.", "data": err})
	}

	if item.Name == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Please input item name.", "data": nil})
	} else if item.Color == nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Please input item color.", "data": nil})
	}

	// Search for item type
	err = db.First(&item.Type, item.TypeID).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not find item type", "data": err})
	}
	// Create the item and return error if encountered
	err = db.Create(&item).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create item", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Item created", "data": item})
}

func GetItems(c *fiber.Ctx) error {
	db := database.DB
	var items []model.Item

	// Find all types in the database
	db.Preload("Type").Find(&items)

	if len(items) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No item present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Items found", "data": items})
}

func GetItem(c *fiber.Ctx) error {
	db := database.DB
	var item model.Item

	id := c.Params("itemId")

	db.Find(&item, "id = ?", id)

	if item.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No item present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Item found", "data": item})
}

func UpdateItem(c *fiber.Ctx) error {
	type UpdateItem struct {
		Name   string  `json:"name"`
		TypeID uint    `json:"typeID"`
		Price  int     `json:"price"`
		Color  *string `json:"color"`
		Qty    int     `json:"qty"`
	}

	db := database.DB
	var item model.Item

	id := c.Params("itemId")

	db.Preload("Type").Find(&item, "id = ?", id)

	if item.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No item present", "data": nil})
	}

	var updateItem UpdateItem
	err := c.BodyParser(&updateItem)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Please check again your input", "data": err})
	}

	var tipe model.Type
	db.Find(&tipe, "id = ?", updateItem.TypeID)

	if tipe.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No type present", "data": nil})
	}

	item.Name = updateItem.Name
	item.Price = updateItem.Price
	item.Color = updateItem.Color
	item.Qty = updateItem.Qty
	item.Type = tipe

	err = db.Save(&item).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update item", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Item successfully updated", "data": item})

}

func DeleteItem(c *fiber.Ctx) error {
	var item model.Item

	db := database.DB
	id := c.Params("itemId")

	db.Find(&item, "id = ?", id)

	if item.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Item not found", "data": nil})
	}

	err := db.Delete(&item).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete item", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Item successfully deleted", "data": item})
}
