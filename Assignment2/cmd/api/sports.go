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
		Title                 *string                     `json:"title"`
		Description           *string                     `json:"description"`
		Type                  *string                     `json:"type"`
		Brand                 *string                     `json:"brand"`
		Sex                   *string                     `json:"sex"`
		SportsEquipmentNumber *data.SportsEquipmentNumber `json:"sports_equipment_number"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		sport.Title = *input.Title
	}
	if input.Description != nil {
		sport.Description = *input.Description
	}
	if input.Type != nil {
		sport.Type = *input.Type
	}
	if input.Brand != nil {
		sport.Brand = *input.Brand
	}
	if input.Sex != nil {
		sport.Sex = *input.Sex
	}
	if input.SportsEquipmentNumber != nil {
		sport.SportsEquipmentNumber = *input.SportsEquipmentNumber
	}

	v := validator.New()

	if data.ValidateSport(v, sport); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Sports.Update(sport)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
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

func (app *application) listSportHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		Type  string
		Brand string
		Sex   string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Type = app.readString(qs, "type", "")
	input.Brand = app.readString(qs, "brand", "")
	input.Sex = app.readString(qs, "sex", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "type", "brand", "sex", "sports_equipment_number", "-id", "-title", "-type", "-brand", "-sex", "-sports_equipment_number"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	sports, metadata, err := app.models.Sports.GetAll(input.Title, input.Type, input.Brand, input.Sex, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sports": sports, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
