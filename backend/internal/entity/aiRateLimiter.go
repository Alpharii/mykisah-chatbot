package entity

import "time"

type AiRateLimiter struct {
	Type      string    `gorm:"size:10;primaryKey"` // "day" | "minute"
	Count     uint
	Limit     uint
	LastReset time.Time
}
