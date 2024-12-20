package model

type Chest struct { // chests
	ID      int64 `gorm:"primaryKey"`
	Gain    int64 `gorm:"not null"`
	PlaceID int64 `json:"place_id" gorm:"not null"`
}
