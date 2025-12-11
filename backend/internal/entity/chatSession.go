package entity

import (
	"time"
)

type ChatSession struct {
	ID        uint           `gorm:"primarykey"`
	UserID    uint           `gorm:"not null"`
	Title     string         `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Messages []ChatMessage `gorm:"foreignKey:SessionID"`
}
