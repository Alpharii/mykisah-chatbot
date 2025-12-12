package chatController

import (
	"ai-chat/internal/entity"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SendChat(c *fiber.Ctx, db *gorm.DB)error{
	payload := struct {
		SessionId	uint	`json:"sessionId"`
		Message		string	`json:"message"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	chat := entity.ChatMessage{
		SessionID: payload.SessionId,
		Content: payload.Message,
		Role: "user",
		CreatedAt: time.Now(),
	}

	if err := db.Create(&chat).Error; err != nil{
		return c.Status(500).JSON(fiber.Map{"error": "failded to create chat message"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message_id": chat.ID,
		"stream_url": fmt.Sprintf("/chat/stream?session_id=%d&message=%s", chat.SessionID, chat.Content),
	})
}