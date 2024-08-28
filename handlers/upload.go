package handlers

import (
	"fmt"
	"net/http"
	"time"

	"repo/config"
	"repo/database"
	"repo/models"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)



func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to get file")
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to open file")
	}
	defer src.Close()

	// Initialize MinIO client
	minioClient := config.InitMinioClient()

	// Generate a unique file name or use the original name
	fileName := file.Filename
	bucketName := "dev-bucket"

	// Upload the file to MinIO
	_, err = minioClient.PutObject(c.Context(), bucketName, fileName, src, file.Size, minio.PutObjectOptions{ContentType: file.Header["Content-Type"][0]})
	if err != nil {
		fmt.Println("MinIO Upload Error:", err)
		return c.Status(http.StatusInternalServerError).SendString("Failed to upload file to MinIO")
	}

	// Save metadata to DB
	fileMeta := models.FileMetadata{FileName: fileName, UploadTime: time.Now()}
	if err := database.DB.Create(&fileMeta).Error; err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to save metadata to database")
	}
	fmt.Printf("Uploading file: %s, size: %d, content type: %s\n", fileName, file.Size, file.Header["Content-Type"][0])

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("File %s uploaded and metadata saved", fileName),
	})
}
