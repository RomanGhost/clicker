package jwtservice

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	UserID    uint
	UserLogin string
	TimeLimit int64
}

func (t *JWTToken) ToString() (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  t.UserID,
		"logU": t.UserLogin,
		"exp":  t.TimeLimit,
	})

	tokenString, err := token.SignedString([]byte("testPhrase"))
	if err != nil {
		return nil, err
	}

	return &tokenString, err
}

func GetFromJWT(tokenString string) (*JWTToken, error) {
	secret := []byte("testPhrase")
	parsedToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Ошибка при парсинге токена: %v", err)
	}

	// Проверяем, что токен валиден и приводим claims к jwt.MapClaims:
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// Извлекаем информацию:
		subVal := claims["sub"].(float64)

		userID := uint(subVal)

		return &JWTToken{
			UserID:    userID,
			UserLogin: claims["logU"].(string),
			TimeLimit: claims["exp"].(int64),
		}, nil
	} else {
		return nil, fmt.Errorf("token is invalid")
	}
}
