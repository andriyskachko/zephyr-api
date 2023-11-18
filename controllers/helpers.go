package controllers

import "net/http"

func WriteJsonResponse(w http.ResponseWriter, status int, response []byte) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
