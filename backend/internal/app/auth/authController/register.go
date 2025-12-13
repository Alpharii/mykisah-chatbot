package authController

import (
	"ai-chat/internal/entity"
	"ai-chat/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx, db *gorm.DB) error {
	payload := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid payload"})
	}

	var user entity.User
	err := db.Where("email = ?", payload.Email).First(&user).Error

	switch err {
	case nil:
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "email sudah terdaftar"})

	case gorm.ErrRecordNotFound:
		hashedPass, err := utils.HashPassword(payload.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "failed to hash password"})
		}

		user = entity.User{
			Username: payload.Username,
			Email:    payload.Email,
			Password: hashedPass,
		}

		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "failed to create user"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		})

	default:
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
}
