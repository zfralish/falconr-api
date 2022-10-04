package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Birds BirdModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Birds: BirdModel{DB: db},
	}
}
