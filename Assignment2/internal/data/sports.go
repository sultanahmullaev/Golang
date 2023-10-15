package data

import (
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
	v.Check(sport.Type != "", "ty", "must be provided")
	v.Check(len(sport.Type) <= 300, "type", "must not be more than 300 bytes long")
	v.Check(sport.Brand != "", "brand", "must be provided")
	v.Check(len(sport.Brand) <= 300, "brand", "must not be more than 300 bytes long")
	v.Check(sport.Sex != "", "sex", "must be provided")
	v.Check(len(sport.Sex) <= 10, "sex", "must not be more than 10 bytes long")
}
