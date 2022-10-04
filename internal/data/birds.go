package data

import (
	"database/sql"
	"falconr-api/internal/validator"
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

type BirdModel struct {
	DB *sql.DB
}

func ValidateBird(v *validator.Validator, bird *Bird) {
	v.Check(validator.IsValidUUID(bird.ID.String()), "id", "must be provided")
	v.Check(bird.FalconerID != "", "falconerId", "must not be empty")
	v.Check(bird.Name != "", "name", "must be provided")
	v.Check(bird.Species != "", "species", "must be provided")
}

func (b BirdModel) Insert(bird *Bird) error {
	query := `
        INSERT INTO birds (id, falconer_id, name, species, trap_date)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at`

	args := []any{bird.ID, bird.FalconerID, bird.Name, bird.Species, bird.TrapDate}

	return b.DB.QueryRow(query, args...).Scan(&bird.ID, &bird.CreatedAt, &bird.UpdatedAt)
}

func (b BirdModel) Get(id uuid.UUID) (*Bird, error) {
	return nil, nil
}

func (b BirdModel) Update(bird *Bird) error {
	return nil
}

func (b BirdModel) Delete(id uuid.UUID) error {
	return nil
}
