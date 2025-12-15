package router

import (
	"ai-chat/internal/app/auth"
	"ai-chat/internal/app/chat"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()

	frontendUrl := os.Getenv("FRONTEND_URL")

	app.Use(cors.New(cors.Config{
		AllowOrigins:     frontendUrl,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	router := app.Group("/api")
	
	router.Get("/", func (c*fiber.Ctx) error {
		return c.JSON("hi from backend")
	})
	
	auth.AuthRouter(router, db)
	chat.ChatRouter(router, db)

	return app
}