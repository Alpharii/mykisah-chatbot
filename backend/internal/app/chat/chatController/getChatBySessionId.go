package chatController

import (
	"ai-chat/internal/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetChatBySessionId(c *fiber.Ctx, db *gorm.DB) error {
	userId := c.Locals("user_id")
	if(userId == nil){
		return c.Status(401).JSON(fiber.Map{"error":"unathorized"})
	}

	sessiondId := c.Params("id")

	var chatSession entity.ChatSession
	if err := db.Where("id = ? AND user_id = ?", sessiondId, userId).First(&chatSession).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"error": "session is not yours"})
		}

		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	
	var chatMsgs []entity.ChatMessage
	if err := db.Where("session_id = ?", chatSession.ID).Order("id ASC").Find(&chatMsgs).Error; err != nil{
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.Status(200).JSON(chatMsgs)
}