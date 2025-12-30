package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ImagesController(c *fiber.Ctx) error {

	fileName := c.Params("id")

	if fileName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "nombre del archivo requerido"})
	}

	if strings.Contains(fileName, "..") {
		return c.Status(400).JSON(fiber.Map{"error": "nombre de archivo no válido"})
	}

	root := os.Getenv("Directory_Path")

	folder := os.Getenv("Images")

	if root == "" || folder == "" {
		return c.Status(500).JSON(fiber.Map{"error": "variables de entorno no configuradas"})
	}

	filePath := filepath.Join(root, folder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "imagen no encontrada"})
	}

	return c.SendFile(filePath)

}

func VideoController(c *fiber.Ctx) error {

	fileName := c.Params("id")

	if fileName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "nombre del archivo requerido"})
	}

	if strings.Contains(fileName, "..") {
		return c.Status(400).JSON(fiber.Map{"error": "nombre de archivo no válido"})
	}

	root := os.Getenv("Directory_Path")
	folder := os.Getenv("Videos")
	if root == "" || folder == "" {
		return c.Status(500).JSON(fiber.Map{"error": "variables de entorno no configuradas"})
	}

	filePath := filepath.Join(root, folder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "imagen no encontrada"})
	}

	return c.SendFile(filePath)

}
