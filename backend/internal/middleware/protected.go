package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected()fiber.Handler{
	return func (c*fiber.Ctx) error {
		auth:= c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer "){
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unaothorized",
			})
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return  nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid Signin Method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
}