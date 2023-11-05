package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Sports SportsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Sports: SportsModel{DB: db},
	}
}
