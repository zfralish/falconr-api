package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/birds", app.createBirdHandler)
	router.HandlerFunc(http.MethodGet, "/v1/birds/:id", app.getBirdHandler)
	router.HandlerFunc(http.MethodPut, "/v1/birds/:id", app.updateBirdHandler)

	return router
}
