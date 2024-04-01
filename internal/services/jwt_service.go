package services

import (
	"github.com/EraldCaka/rentio/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func ValidateJWTToken(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func CreateJWTToken(tokenData types.JWTToken, secretKey string, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":       tokenData.ID,
		"username": tokenData.Username,
		"exp":      time.Now().Add(expirationTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/*
		TODO: IMPLEMENT A LOGIC TO STORE THE JWT TOKEN INSIDE THE DATABASE
	*/
	return token.SignedString([]byte(secretKey))
}
