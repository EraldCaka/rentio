package middleware

import (
	"github.com/EraldCaka/rentio/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, route := range Routes {
			if c.Path() == route {
				return c.Next()
			}
		}

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return fiber.ErrUnauthorized
		}

		token, err := services.ValidateJWTToken(tokenString, secretKey)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return fiber.ErrUnauthorized
		}

		c.Locals("user", claims)
		return c.Next()
	}
}
