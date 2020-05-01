package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/tevity/go-getting-started/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) != 2 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	encodeResponseAsJSON(users, w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	user, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(user, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	updatedUser, err := models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(updatedUser, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	user.ID = id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	updatedUser, err := models.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(updatedUser, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusNoContent)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	return user, err
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
