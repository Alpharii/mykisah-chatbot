package router

import (
	"ai-chat/internal/app/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()
	router := app.Group("/api")
	
	router.Get("/", func (c*fiber.Ctx) error {
		return c.JSON("hi from backend")
	})
	
	auth.AuthRouter(router, db)

	return app
}