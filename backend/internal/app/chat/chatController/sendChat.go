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

	var session entity.ChatSession
	if err := db.First(&session, payload.SessionId).Error; err != nil{
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
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

	if (session.Title == ""){
		title := payload.Message
		if len(title) > 50 {
			title = title[:50] + "..."
		}

		db.Model(&session).Updates(map[string]interface{}{
			"title":      title,
			"updated_at": time.Now(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message_id": chat.ID,
		"stream_url": fmt.Sprintf("/chat/stream?session_id=%d&message_id=%d", chat.SessionID, chat.ID),
	})
}