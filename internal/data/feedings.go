package data

import (
	"github.com/google/uuid"
	"time"
)

type Feeding struct {
	ID        uuid.UUID `json:"id"`
	BirdID    uuid.UUID `json:"bird_id"`
	FTime     time.Time `json:"f_time"`
	FoodType  string    `json:"food_type"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
