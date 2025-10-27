package entitiy

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID			uint
	Username	string
	Email		string
	Password	string
	IsActive	bool
	ActivatedAt	sql.NullTime
	CreatedAt	time.Time
	UpdatedAt	time.Time
}