package controllers

import (
	"net/http"
)

// RegisterControllers afds
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
