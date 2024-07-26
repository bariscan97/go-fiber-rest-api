package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type MyCustomClaims struct {
	Id       uuid.UUID `json:omitempty`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	jwt.StandardClaims
}
