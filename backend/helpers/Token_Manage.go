package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generate_Token(value string) (string, error) {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error cargando el archivo .env")
	// }

	// secret := os.Getenv("secret.env")
	// if secret == "" {
	// 	log.Fatal("no encontrado en el entorno")
	// }

	// log.Println("Key_Token:", secret)

	var jwtSecret = []byte("2025")

	claims := jwt.MapClaims{
		"user": value,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
