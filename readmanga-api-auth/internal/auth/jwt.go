package auth

import (
	"fmt"
	"readmanga-api-auth/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("KEY_JWT"))

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(string(jwtKey))
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Algoritmo inesperado: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Não foi possível converter os claims")
	}
	return claims, nil
}
