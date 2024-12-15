package main

import (
	"net/http"
	"github.com/duixe/social_app/internal/env"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
		"env": env.Envs.CurrentEnv,
		"version": env.Envs.Version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}