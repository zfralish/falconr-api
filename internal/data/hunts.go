package data

import (
	"github.com/google/uuid"
	"time"
)

type Hunt struct {
	ID        uuid.UUID `json:"id"`
	BirdID    uuid.UUID `json:"bird_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	PreyType  string    `json:"prey_type"`
	PreyCount int       `json:"prey_count"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
