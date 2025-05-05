package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1 * 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}