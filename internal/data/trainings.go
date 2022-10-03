package data

import (
	"github.com/google/uuid"
	"time"
)

type Training struct {
	ID           uuid.UUID `json:"id"`
	BirdID       uuid.UUID `json:"bird_id"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	TrainingType string    `json:"training_type"`
	Notes        string    `json:"notes"`
	Performance  int64     `json:"performance"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
