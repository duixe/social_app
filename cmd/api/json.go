package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

// create a singleton of validator to be used in payload struct
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxByte := 1_048_578 //1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type ErrorFormat struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &ErrorFormat{message})
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type jsonFormat struct {
		Data any `json:"data"`
	}

	return writeJSON(w, status, &jsonFormat{Data: data})
}
