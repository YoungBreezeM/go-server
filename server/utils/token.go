package utils

import (
	"math/rand"
	"server/config"
	"server/log"
	"server/models"

	"github.com/golang-jwt/jwt"
)

func GeneratorToken(JWTClaims *models.JWTClaims) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims)

	token, err := claims.SignedString(config.JWTSECRET)
	if err != nil {
		log.Log.Errorln(err)
	}

	return token
}

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func ParseToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSECRET, nil
	})
}
