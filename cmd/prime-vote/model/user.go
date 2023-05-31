package model

import "time"

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Username  string `gorm:"uniqueIndex"`
	Password  string // Hashed password
	IsVoted   bool
}
