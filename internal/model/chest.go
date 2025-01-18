package model

type Chest struct { // chests
	ID      int64 `gorm:"primaryKey"`
	Gain    int64 `json:"gain" gorm:"not null"`
	PlaceID int64 `json:"place_id" gorm:"not null"`
}

type ValidationError struct {
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
