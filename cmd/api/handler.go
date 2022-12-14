package main

import (
	"net/http"
)


type RequestPayload struct{
	Action string `json:"action"`
	Auth AuthPayload `json:"auth, omitempty"`
}

type AuthPayload struct{
	Email string `json:"email`
	Password string `json:"password`
}

func (app *Config) BrokerService(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit broker service",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request){

	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

// 	switch requestPayload.Action
}