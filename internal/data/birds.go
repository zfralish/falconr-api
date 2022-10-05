package data

import (
	"database/sql"
	"errors"
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
	query := `
        SELECT *
        FROM birds
        WHERE id = $1`

	var bird Bird

	err := b.DB.QueryRow(query, id).Scan(
		&bird.ID,
		&bird.FalconerID,
		&bird.Name,
		&bird.Species,
		&bird.TrapDate,
		&bird.CreatedAt,
		&bird.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &bird, nil
}

func (b BirdModel) Update(bird *Bird) error {
	query := `
        UPDATE birds 
        SET name = $1, falconer_id = $2, species = $3, trap_date = $4, updated_at = $5
        WHERE id = $6
        returning updated_at`

	args := []any{
		bird.Name,
		bird.FalconerID,
		bird.Species,
		bird.TrapDate,
		time.Now(),
		bird.ID,
	}

	return b.DB.QueryRow(query, args...).Scan(&bird.UpdatedAt)
}

func (b BirdModel) Delete(id uuid.UUID) error {
	return nil
}
