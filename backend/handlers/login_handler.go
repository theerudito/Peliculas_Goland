package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
	"github.com/theerudito/peliculas/models"
)

func POST_Login(c *fiber.Ctx) error {

	var loginRequest models.Login

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	var user models.Login

	err := db.DB.QueryRow(`
		SELECT username, password 
		FROM login 
		WHERE username = ?
	`, strings.ToUpper(loginRequest.UserName)).Scan(&user.UserName, &user.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Usuario o contraseña incorrectos",
		})
	}

	isMatch, err := helpers.ComparePassword(strings.ToLower(loginRequest.Password), user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al comparar contraseña",
		})
	}

	if !isMatch {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Usuario o contraseña incorrectos",
		})
	}

	token, err := helpers.Generate_Token(user.UserName + user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al generar token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.LoginDTO{
		UserName: user.UserName,
		Token:    token,
	})

}
