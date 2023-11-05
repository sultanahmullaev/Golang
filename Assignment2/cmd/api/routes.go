package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/sports", app.listSportsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sports", app.createSportHandler)
	router.HandlerFunc(http.MethodGet, "/v1/sports/:id", app.showSportHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/sports/:id", app.updateSportHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/sports/:id", app.deleteSportHandler)

	return app.recoverPanic(router)
}
