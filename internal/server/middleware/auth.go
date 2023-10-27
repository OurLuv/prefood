package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId uint `json:"user_id"`
	*jwt.StandardClaims
}

func CreateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userId,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (uint, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	return claims.UserId, nil
}
