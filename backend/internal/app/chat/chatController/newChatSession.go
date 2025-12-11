package chatController

import (
	"ai-chat/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewChatSession(c *fiber.Ctx, db *gorm.DB) error {
	userId := c.Locals("user_id")
	if (userId == nil){
		return c.Status(401).JSON(fiber.Map{"error": "unaothorized"})
	}

	var user entity.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil{
		return c.Status(401).JSON(fiber.Map{"error": "user not found"})
	}

	session := entity.ChatSession{
		UserID:	user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&session).Error; err != nil{
		return c.Status(500).JSON(fiber.Map{"error": "failded to create user"})
	}


	return c.Status(201).JSON(session)
}
