package auth

import (
	"ai-chat/internal/entitiy"
	"ai-chat/utils"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)


func Login(c*fiber.Ctx, db*gorm.DB) error {
	payload := struct {
		Email		string	`json:"email"`
		Password	string	`json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return  err
	}

	user := entitiy.User{
		Email: payload.Email,
	}

	if err := db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user not found"})
	}

	isAuthorized := utils.CheckHashedPassword(user.Password, payload.Password)
	if !isAuthorized {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong password"})
	}

	user.Password = ""

	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(int(user.ID)),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create token"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "login success",
		"token": token,
		"user":  user,
	})
}

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