package structs

import (
	"github.com/dgrijalva/jwt-go"
)

type DefaultResponse struct {
	Status  int64       `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Token struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}
