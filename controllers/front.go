package controllers

import (
	"encoding/json"
	"net/http"
)

// RegisterControllers afds
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
