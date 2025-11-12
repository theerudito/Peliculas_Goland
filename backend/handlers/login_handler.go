package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
	"github.com/theerudito/peliculas/models"
)

func PostLogin(c *fiber.Ctx) error {

	var loginRequest models.Login

	conn := db.GetDB()

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cuerpo de solicitud inválido"})
	}

	var user models.Login

	err := conn.QueryRow(`
		SELECT username, password 
		FROM login 
		WHERE username = $1
	`, strings.ToUpper(loginRequest.UserName)).Scan(&user.UserName, &user.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Usuario o contraseña incorrectos"})
	}

	password, err := helpers.DesencriptarDato(user.Password)

	if password != loginRequest.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Usuario o contraseña incorrectos"})
	}

	token, err := helpers.GenerateToken(user.UserName + user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al generar token"})
	}

	return c.Status(fiber.StatusOK).JSON(models.LoginDTO{
		UserName: user.UserName,
		Token:    token,
	})

}
