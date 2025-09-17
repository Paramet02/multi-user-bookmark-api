package model


import (
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID int    `json:"ToUserModel"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}