package handler

import (
	"net/http"
)

// http handler for readiness probe
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	// RespondWithJSON(w, http.StatusOK, struct{}{})
	RespondWithJSON(w, http.StatusOK, struct{
		Message string
	}{
		Message: "API up and running",
	})
}