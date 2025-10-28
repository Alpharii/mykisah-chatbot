package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB) {
	auth := router.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		return Login(c, db)
	})
	auth.Post("/register", func(c *fiber.Ctx) error {
		return Register(c, db)
	})
}
