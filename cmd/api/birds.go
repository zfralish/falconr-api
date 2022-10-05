package main

import (
	"errors"
	"falconr-api/internal/data"
	"falconr-api/internal/validator"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (app *application) createBirdHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string    `json:"name"`
		Species    string    `json:"species"`
		FalconerId string    `json:"falconer_id"`
		TrapDate   time.Time `json:"trap_date"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	bird := &data.Bird{
		ID:         uuid.New(),
		Name:       input.Name,
		Species:    input.Species,
		FalconerID: input.FalconerId,
		TrapDate:   input.TrapDate,
	}

	v := validator.New()

	if data.ValidateBird(v, bird); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Birds.Insert(bird)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/birds/%d", bird.ID))

	err = app.writeJSON(w, http.StatusCreated, bird, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBirdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	bird, err := app.models.Birds.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, bird, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateBirdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	bird, err := app.models.Birds.Get(id)
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
		Name       string    `json:"name"`
		Species    string    `json:"species"`
		TrapDate   time.Time `json:"trap_date"`
		FalconerId string    `json:"falconer_id"`
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	bird.Name = input.Name
	bird.Species = input.Species
	bird.TrapDate = input.TrapDate
	bird.FalconerID = input.FalconerId

	v := validator.New()

	if data.ValidateBird(v, bird); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Birds.Update(bird)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, bird, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
