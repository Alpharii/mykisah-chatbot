package chatController

import (
	"ai-chat/config"
	"ai-chat/internal/entity"
	ratelimiter "ai-chat/utils/rateLimiter"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/genai"
	"gorm.io/gorm"
)

func StreamAiChat(c *fiber.Ctx, db *gorm.DB) error {
	fmt.Println("hitted")
	sessionId, _ := strconv.Atoi(c.Query("session_id"))
	messageId := c.Query("message_id")
	
	if sessionId == 0 {
    return c.Status(400).JSON(fiber.Map{
        "error": "session_id is required",
    })
	}

	// pastikan limiter ada
	if err := ratelimiter.EnsureAiRateLimiter(db); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to init ai limiter",
		})
	}

	allowed, msg, err := ratelimiter.CheckAiRateLimit(db)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !allowed {
		// return sebagai pesan AI
		return c.JSON(fiber.Map{
			"message": fiber.Map{
				"role":    "assistant",
				"content": msg,
			},
		})
	}

	var userMessage *entity.ChatMessage
	db.Where("id = ?", messageId).First(&userMessage)

	var msgs []entity.ChatMessage
	db.Where("session_id = ?", sessionId).Order("id ASC").Find(&msgs)

	history := []*genai.Content{}
	for _, m := range msgs {
		if m.Role == "user" {
			history = append(history, genai.NewContentFromText(m.Content, genai.RoleUser))
		} else {
			history = append(history, genai.NewContentFromText(m.Content, genai.RoleModel))
		}
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	ctx := context.Background()

	chat, _ := config.GeminiClient.Chats.Create(
		ctx,
		"gemini-2.5-flash",
		nil,
		history,
	)

	resp, err := chat.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(userMessage.Content),
		config.GeminiConfig,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var fullResponse string
	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}
		for _, p := range cand.Content.Parts {
			if p.Text != "" {
				fullResponse += p.Text
			}
		}
	}

	aiMsg := entity.ChatMessage{
		SessionID: uint(sessionId),
		Role:      "assistant",
		Content:   fullResponse,
		CreatedAt: time.Now(),
	}
	db.Create(&aiMsg)

	return c.JSON(fiber.Map{
		"message": aiMsg,
	})
}