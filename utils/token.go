package utils

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func parseToken(token string) (*jwt.Token, error) {
	tokenString := strings.Split(token, " ")[1]
	
	result, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	
	if err != nil {
		return nil, ErrInvalidToken
	}

	return result, nil
}

func GetIDFromToken(token string) (uint, error) {
	result, err := parseToken(token)
	if err != nil {
		return 0, err 
	}
	claims, ok := result.Claims.(jwt.MapClaims)
	if ok && result.Valid {
		return uint(claims["id"].(float64)), nil
	}

	return 0, ErrInvalidToken
}