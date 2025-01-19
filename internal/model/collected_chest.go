package model

import "time"

type CollectedChest struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64 `json:"user_id" gorm:"not null"`
	ChestID   int64 `json:"chest_id" gorm:"not null"`
	CreatedAt time.Time
}
