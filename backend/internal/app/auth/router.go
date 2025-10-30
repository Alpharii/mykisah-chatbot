package auth

import (
	"ai-chat/internal/app/auth/controller"
	"ai-chat/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB) {
	auth := router.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		return controller.Login(c, db)
	})
	auth.Post("/register", func(c *fiber.Ctx) error {
		return controller.Register(c, db)
	})

	protected := auth.Group("/me", middleware.Protected())
	protected.Get("/", func(c *fiber.Ctx) error {
    	return controller.GetMe(c, db)
	})
}
