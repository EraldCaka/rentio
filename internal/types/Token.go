package types

import "github.com/golang-jwt/jwt/v5"

type JWTToken struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	jwt.RegisteredClaims
}
