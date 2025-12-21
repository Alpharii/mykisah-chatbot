package ratelimiter

import (
	"ai-chat/internal/entity"
	"time"
)

func ShouldReset(limiter *entity.AiRateLimiter) bool {
	now := time.Now()

	switch limiter.Type {
	case "day":
		return now.Format("2006-01-02") != limiter.LastReset.Format("2006-01-02")
	case "minute":
		return now.Sub(limiter.LastReset) >= time.Minute
	default:
		return false
	}
}
