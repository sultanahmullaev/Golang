package main

import (
	"Assignment2/internal/data"
	"Assignment2/internal/validator"
	"errors"
	"fmt"
	"net/http"
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

	err = app.models.Sports.Insert(sport)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sports/%d", sport.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"sport": sport}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showSportHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	sport, err := app.models.Sports.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sport": sport}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateSportHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	sport, err := app.models.Sports.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title                 string                     `json:"title"`
		Description           string                     `json:"description"`
		Type                  string                     `json:"type"`
		Brand                 string                     `json:"brand"`
		Sex                   string                     `json:"sex"`
		SportsEquipmentNumber data.SportsEquipmentNumber `json:"sports_equipment_number"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sport.Title = input.Title
	sport.Description = input.Description
	sport.Type = input.Type
	sport.Brand = input.Brand
	sport.Sex = input.Sex
	sport.SportsEquipmentNumber = input.SportsEquipmentNumber

	v := validator.New()
	if data.ValidateSport(v, sport); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Sports.Update(sport)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sport": sport}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSportHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Sports.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "sport successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
