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

	router.HandlerFunc(http.MethodGet, "/v1/sports", app.requirePermission("sports:read", app.listSportHandler))
	router.HandlerFunc(http.MethodPost, "/v1/sports", app.requirePermission("sports:write", app.createSportHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sports/:id", app.requirePermission("sports:read", app.showSportHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/sports/:id", app.requirePermission("sports:write", app.updateSportHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/sports/:id", app.requirePermission("sports:write", app.deleteSportHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
