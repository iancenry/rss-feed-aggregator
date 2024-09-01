package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// marshal payload to JSON string and return as bytes
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// set headers, status code, and write payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("Responding with 5XX level error %s", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	
	}
	RespondWithJSON(w, code, errorResponse{
		Error: message,
	})
}