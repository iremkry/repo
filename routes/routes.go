package routes

import (
    "log"
    "repo/database"
    "repo/models"

    "github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetFileMetadata(c *fiber.Ctx) error {
	var items []models.FileMetadata

	if err := database.DB.Find(&items).Error; err != nil {
		log.Println("Error retrieving file metadata:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
	}

	return c.JSON(items)
}

