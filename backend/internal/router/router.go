package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()

	app.Get("/", func (c*fiber.Ctx) error {
		return c.JSON("hi from backend")
	})

	return app
}