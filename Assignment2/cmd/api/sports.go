package main

import (
	"Assignment2/internal/validator"
	"fmt"
	"net/http"
	"time"

	"Assignment2/internal/data"
)

func (app *application) createSportHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title                 string                     `json:"title"`
		Description           string                     `json:"description"`
		Type                  string                     `json:"type"`
		Brand                 string                     `json:"brand"`
		Sex                   string                     `json:"sex"`
		SportsEquipmentNumber data.SportsEquipmentNumber `json:"sports_equipment_number"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sport := &data.Sports{
		Title:                 input.Title,
		Description:           input.Description,
		Type:                  input.Type,
		Brand:                 input.Brand,
		Sex:                   input.Sex,
		SportsEquipmentNumber: input.SportsEquipmentNumber,
	}

	v := validator.New()
	if data.ValidateSport(v, sport); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showSportHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	sport := data.Sports{
		ID:                    id,
		CreatedAt:             time.Now(),
		Title:                 "T-shirt",
		Description:           "T-shirt made 100% cotton",
		Type:                  "sportswear",
		Brand:                 "Puma",
		Sex:                   "Male",
		SportsEquipmentNumber: 200,
		Version:               1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sport": sport}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
