package handler

import (
	"encoding/json"
	"net/http"
)
// respondJSON makes the response in payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}){
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Set specific headers
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.Header().Set("Accept", "application/vnd.api+json")

	// Set response
	w.Write([]byte(response))
}

// respondError makes the error response in payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
