package ratelimiter

import (
	"ai-chat/internal/entity"
	"time"

	"gorm.io/gorm"
)

func EnsureAiRateLimiter(db *gorm.DB) error {
	var count int64
	db.Model(&entity.AiRateLimiter{}).Count(&count)

	if count > 0 {
		return nil
	}

	now := time.Now()

	limiters := []entity.AiRateLimiter{
		{
			Type:      "day",
			Limit:     17,
			Count:     0,
			LastReset: now,
		},
		{
			Type:      "minute",
			Limit:     4,
			Count:     0,
			LastReset: now,
		},
	}

	return db.Create(&limiters).Error
}
