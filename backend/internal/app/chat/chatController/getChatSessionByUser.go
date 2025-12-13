package chatController

import (
	"ai-chat/internal/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetChatSessionByUser(c *fiber.Ctx, db *gorm.DB) error {
	userId := c.Locals("user_id")
	if(userId == nil){
		return c.Status(401).JSON(fiber.Map{"error": "unaothorized"})
	}

	var user *entity.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil{
		return c.Status(401).JSON(fiber.Map{"error": "user not found"})
	}

	var chatSession []entity.ChatSession
	if err := db.Where("user_id = ?", userId).Find(&chatSession).Error; err != nil{
		return c.Status(404).JSON(fiber.Map{"error" : "not found"})
	}

	return c.Status(200).JSON(chatSession)
}