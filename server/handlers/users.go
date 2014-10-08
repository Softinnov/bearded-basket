package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/mux"
)

func checkPdvId(c *utils.Context, id int, r *http.Request) (int, error) {
	user, err := models.GetUser(c, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	session, err := models.GetSessionFromCookies(c.Store, r)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if session.PdvId != user.Pdv {
		return http.StatusUnauthorized, errors.New("")
	}
	return http.StatusOK, nil
}

func editUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if s, err := checkPdvId(c, id, r); err != nil {
		return s, err
	}

	var user struct {
		Prenom   string `json:"u_prenom"`
		Nom      string `json:"u_nom"`
		Role     int8   `json:"u_role"`
		Password string `json:"u_pass"`
	}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&user); err != nil {
		return http.StatusBadRequest, err
	}
	defer r.Body.Close()

	u := &models.User{
		Prenom:   user.Prenom,
		Nom:      user.Nom,
		Role:     user.Role,
		Password: user.Password,
	}
	if err = models.UpdateUser(c, id, u); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusAccepted, nil
}

func indexUsers(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	session, err := models.GetSessionFromCookies(c.Store, r)
	if err != nil {
		return http.StatusUnauthorized, err
	}
	users, err := models.GetUsersFromSession(c, session)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, users)
	return http.StatusOK, nil
}

func newUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	var user struct {
		Prenom   string `json:"u_prenom"`
		Nom      string `json:"u_nom"`
		Role     int8   `json:"u_role"`
		Password string `json:"u_pass"`
		Login    string `json:"u_login"`
	}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&user); err != nil {
		return http.StatusBadRequest, err
	}
	defer r.Body.Close()

	s, err := models.GetSessionFromCookies(c.Store, r)
	if err != nil {
		return http.StatusUnauthorized, err
	}
	u := &models.User{
		Prenom:   user.Prenom,
		Nom:      user.Nom,
		Role:     user.Role,
		Password: user.Password,
		Login:    user.Login,
	}
	if err := models.CreateUser(c, u, s); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusCreated, nil
}

func deleteUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if s, err := checkPdvId(c, id, r); err != nil {
		return s, err
	}
	if err = models.RemoveUser(c, id); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusOK, nil
}

func getCurrentUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	s, err := models.GetSessionFromCookies(c.Store, r)
	if err != nil {
		return http.StatusUnauthorized, err
	}
	user, err := models.GetCurrentUser(c, s)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, user)
	return http.StatusOK, nil
}
