package ratelimiter

import (
	"ai-chat/internal/entity"
	"time"

	"gorm.io/gorm"
)

func CheckAiRateLimit(db *gorm.DB) (bool, string, error) {
	var limiters []entity.AiRateLimiter
	if err := db.Find(&limiters).Error; err != nil {
		return false, "", err
	}

	for i := range limiters {
		l := &limiters[i]

		if ShouldReset(l) {
			l.Count = 0
			l.LastReset = time.Now()
			db.Save(l)
		}

		if l.Count >= l.Limit {
			if l.Type == "day" {
				return false, "Limit harian AI sudah tercapai", nil
			}
			if l.Type == "minute" {
				return false, "Limit AI per menit sudah tercapai", nil
			}
		}
	}

	// increment semua limiter
	for _, l := range limiters {
		db.Model(&entity.AiRateLimiter{}).
			Where("type = ?", l.Type).
			Update("count", gorm.Expr("count + 1"))
	}

	return true, "", nil
}
