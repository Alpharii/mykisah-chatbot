package chatController

import (
	"ai-chat/config"
	"ai-chat/internal/entity"
	"bufio"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/genai"
	"gorm.io/gorm"
)

func StreamAiChat(c *fiber.Ctx, db *gorm.DB) error {
	sessionId, _ := strconv.Atoi(c.Query("session_id"))
	userMessage := c.Query("message")
	

	if sessionId == 0 {
    return c.Status(400).JSON(fiber.Map{
        "error": "session_id is required",
    })
	}


	var msgs []entity.ChatMessage
	db.Where("session_id = ?", sessionId).Order("id ASC").Find(&msgs)

	fmt.Println(sessionId)
	fmt.Println(msgs)

	history := []*genai.Content{}
	for _, m := range msgs {
		if m.Role == "user" {
			history = append(history, genai.NewContentFromText(m.Content, genai.RoleUser))
		} else {
			history = append(history, genai.NewContentFromText(m.Content, genai.RoleModel))
		}
	}

	fmt.Println(history)

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	ctx := context.Background()

	chat, err := config.GeminiClient.Chats.Create(
		ctx,
		"gemini-2.5-flash",
		nil,
		history,
	)
	if err != nil {
		return c.Status(500).SendString("Error creating chat session: " + err.Error())
	}

	stream := chat.SendMessageStream(ctx, genai.Part{Text: userMessage})

	fullResponse := ""

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {

		for chunk := range stream {

			if chunk == nil {
				continue
			}

			if len(chunk.Candidates) == 0 {
				continue
			}

			cand := chunk.Candidates[0]

			if cand.Content == nil {
				continue
			}

			if len(cand.Content.Parts) == 0 {
				continue
			}

			for _, p := range cand.Content.Parts {
				if p == nil || p.Text == "" {
					continue
				}

				fullResponse += p.Text

				fmt.Fprintf(w, "data: %s\n\n", p.Text)
				w.Flush()
			}
		}


		aiMsg := entity.ChatMessage{
			SessionID: uint(sessionId),
			Role:      "assistant",
			Content:   fullResponse,
			CreatedAt: time.Now(),
		}
		db.Create(&aiMsg)

		fmt.Fprintf(w, "event: done\ndata: end\n\n")
		w.Flush()
	})

	return nil
}
