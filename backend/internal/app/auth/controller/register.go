package controller

import (
	"ai-chat/internal/entitiy"
	"ai-chat/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func Register(c*fiber.Ctx, db*gorm.DB) error {
	payload := struct {
		Username	string	`json:"username"`
		Email		string	`json:"email"`
		Password	string	`json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	
	hashedPass, err := utils.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	
	user := entitiy.User{
		Username: payload.Username,
		Password: hashedPass,
		Email: payload.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}