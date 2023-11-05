package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"Assignment2/internal/validator"
)

type Sports struct {
	ID                    int64                 `json:"id"`
	CreatedAt             time.Time             `json:"-"`
	Title                 string                `json:"title"`
	Description           string                `json:"description,omitempty"`
	Type                  string                `json:"type,omitempty"`
	Brand                 string                `json:"brand,omitempty"`
	Sex                   string                `json:"sex,omitempty"`
	SportsEquipmentNumber SportsEquipmentNumber `json:"sports_equipment_number,omitempty"`
	Version               int32                 `json:"version"`
}

func ValidateSport(v *validator.Validator, sport *Sports) {
	v.Check(sport.Title != "", "title", "must be provided")
	v.Check(len(sport.Title) <= 100, "title", "must not be more than 100 bytes long")
	v.Check(sport.Description != "", "description", "must be provided")
	v.Check(len(sport.Description) <= 3000, "description", "must not be more than 3000 bytes long")
	v.Check(sport.Type != "", "type", "must be provided")
	v.Check(len(sport.Type) <= 300, "type", "must not be more than 300 bytes long")
	v.Check(sport.Brand != "", "brand", "must be provided")
	v.Check(len(sport.Brand) <= 300, "brand", "must not be more than 300 bytes long")
	v.Check(sport.Sex != "", "sex", "must be provided")
	v.Check(len(sport.Sex) <= 10, "sex", "must not be more than 10 bytes long")
}

type SportsModel struct {
	DB *sql.DB
}

func (s SportsModel) Insert(sport *Sports) error {
	query := `
			  INSERT INTO sports (title, description, type, brand, sex, sports_equipment_number)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, created_at, version`

	args := []interface{}{sport.Title, sport.Description, sport.Type, sport.Brand, sport.Sex, sport.SportsEquipmentNumber}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.DB.QueryRowContext(ctx, query, args...).Scan(&sport.ID, &sport.CreatedAt, &sport.Version)
}

func (s SportsModel) Get(id int64) (*Sports, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
			  SELECT id, created_at, title, description, type, brand, sex, sports_equipment_number, version
			  FROM sports
			  WHERE id = $1`

	var sport Sports

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := s.DB.QueryRowContext(ctx, query, id).Scan(
		&sport.ID,
		&sport.CreatedAt,
		&sport.Title,
		&sport.Description,
		&sport.Type,
		&sport.Brand,
		&sport.Sex,
		&sport.SportsEquipmentNumber,
		&sport.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &sport, nil
}

func (s SportsModel) Update(sport *Sports) error {
	query := `
			  UPDATE sports
			  SET title = $1, description = $2, type = $3, brand = $4, sex = $5, sports_equipment_number = $6, version = version + 1
			  WHERE id = $7 AND version = $8
			  RETURNING version`

	args := []interface{}{
		sport.Title,
		sport.Description,
		sport.Type,
		sport.Brand,
		sport.Sex,
		sport.SportsEquipmentNumber,
		sport.Version,
		sport.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.DB.QueryRowContext(ctx, query, args...).Scan(&sport.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (s SportsModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
			  DELETE FROM sports
			  WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := s.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
