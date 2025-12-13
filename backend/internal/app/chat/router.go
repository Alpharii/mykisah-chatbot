package chat

import (
	"ai-chat/internal/app/chat/chatController"
	"ai-chat/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChatRouter(router fiber.Router, db *gorm.DB){
	routes := router.Group("/chat", middleware.Protected())

	routes.Get("/session", func(c *fiber.Ctx) error {
		return chatController.GetChatSessionByUser(c, db)
	})

	routes.Get("/session/:id", func(c *fiber.Ctx) error {
		return chatController.GetChatBySessionId(c, db)
	})

	routes.Post("/session/new", func(c *fiber.Ctx) error {
		return chatController.NewChatSession(c, db)
	})

	routes.Post("/send", func(c *fiber.Ctx) error {
		return chatController.SendChat(c, db)
	})

	routes.Get("/stream", func(c *fiber.Ctx) error {
		return chatController.StreamAiChat(c, db)
	})
}