package controller

import (
	"ai-chat/internal/entitiy"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMe (c*fiber.Ctx, db*gorm.DB) error {
	fmt.Println("hitted")
	userId := c.Locals("user_id")
	if userId == nil {
		return c.Status(401).JSON(fiber.Map{"error": "unaothorized"})
	}

	var user entitiy.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user not found"})
	}
	
	return c.JSON(user)
}