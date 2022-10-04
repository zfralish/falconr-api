package main

import (
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

	// Note that the movie variable contains a *pointer* to a Movie struct.
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

	// Call the Insert() method on our movies model, passing in a pointer to the
	// validated movie struct. This will create a record in the database and update the
	// movie struct with the system-generated information.
	err = app.models.Birds.Insert(bird)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// When sending a HTTP response, we want to include a Location header to let the
	// client know which URL they can find the newly-created resource at. We make an
	// empty http.Header map and then use the Set() method to add a new Location header,
	// interpolating the system-generated ID for our new movie in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/birds/%d", bird.ID))

	// Write a JSON response with a 201 Created status code, the movie data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"bird": bird}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
