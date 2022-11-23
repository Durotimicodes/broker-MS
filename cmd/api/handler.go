package main

import (
	"net/http"
)

func (app *Config) BrokerService(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit broker service",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
