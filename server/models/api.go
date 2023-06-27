package models

import "github.com/golang-jwt/jwt"

type R[T any] struct {
	Status  int16  `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type JWTClaims struct {
	OpenId string `json:"token"`
	jwt.StandardClaims
}
