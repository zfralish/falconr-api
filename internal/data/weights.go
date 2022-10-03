package data

import (
	"github.com/google/uuid"
	"time"
)

type Weight struct {
	ID        uuid.UUID `json:"id"`
	BirdID    uuid.UUID `json:"bird_id"`
	Weight    float32   `json:"weight"`
	WTime     time.Time `json:"w_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
