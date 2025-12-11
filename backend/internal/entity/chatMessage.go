package entity

import (
	"time"
)

type ChatMessage struct {
	ID        uint           `gorm:"primarykey"`
	SessionID uint           `gorm:"not null"`
	Role      string         `gorm:"size:10;not null"` // "user" atau "assistant"
	Content   string         `gorm:"type:text;not null"`
	CreatedAt time.Time
}
