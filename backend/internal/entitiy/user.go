package entitiy

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID			uint			`gorm:"primarykey"`
	Username  string         	`gorm:"size:100;not null"`
	Email     string         	`gorm:"size:100;uniqueIndex;not null"`
	Password  string         	`gorm:"size:255;not null"`
	IsActive	bool
	ActivatedAt	sql.NullTime
	CreatedAt	time.Time
	UpdatedAt	time.Time
}