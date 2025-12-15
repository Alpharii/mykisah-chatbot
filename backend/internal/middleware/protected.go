package middleware

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {

		tokenStr := c.Cookies("token")

		if tokenStr == "" {
			auth := c.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "unauthorized",
				})
			}
			tokenStr = strings.TrimPrefix(auth, "Bearer ")
		} else {
			decoded, err := base64.StdEncoding.DecodeString(tokenStr)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "invalid token encoding",
				})
			}
			tokenStr = string(decoded)
			tokenStr = strings.Trim(tokenStr, `"`)
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(
					fiber.StatusUnauthorized,
					"invalid signing method",
				)
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