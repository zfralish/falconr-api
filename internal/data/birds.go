package data

import (
	"github.com/google/uuid"
	"time"
)

type Bird struct {
	ID         uuid.UUID `json:"id"`
	FalconerID string    `json:"falconer_id"`
	Name       string    `json:"name"`
	Species    string    `json:"species"`
	TrapDate   time.Time `json:"trap_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
