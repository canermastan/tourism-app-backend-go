package model

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID         int64  `gorm:"primaryKey"`
	UserID     int64  `json:"user_id" gorm:"not null" validate:"required"`
	PlaceID    int    `json:"place_id" gorm:"not null" validate:"required"`
	Rating     byte   `json:"rating" validate:"gte=0,lte=5,required"`
	Comment    string `json:"comment" validate:"max=255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	LikesCount int            `json:"likes_count"`
	Visibility bool           `json:"visibility"`
}
