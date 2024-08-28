package main

import (
	"net/http"
)

// http handler for readiness probe
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}