package auth

import (
	"ai-chat/internal/app/auth/authController"
	"ai-chat/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB) {
	auth := router.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		return authController.Login(c, db)
	})
	auth.Post("/register", func(c *fiber.Ctx) error {
		return authController.Register(c, db)
	})

	protected := auth.Group("/me", middleware.Protected())
	protected.Get("/", func(c *fiber.Ctx) error {
    	return authController.GetMe(c, db)
	})
}
