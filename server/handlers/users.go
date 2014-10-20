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

func editUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 64)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	uEdit, err := models.GetUser(c, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	session, err := models.GetSessionFromCookies(c.Store, r)
	if err != nil {
		return http.StatusUnauthorized, err
	}
	if session.PdvId != uEdit.Pdv {
		return http.StatusUnauthorized, errors.New("")
	}
	if session.Role < uEdit.Role {
		return http.StatusUnauthorized, errors.New("")
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
	if _, err = models.UpdateUser(c, id, u, session); err != nil {
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
	if _, err := models.CreateUser(c, u, s); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusCreated, nil
}

func deleteUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 64)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	u, err := models.GetUser(c, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	session, err := models.GetSessionFromCookies(c.Store, r)
	if err != nil {
		return http.StatusUnauthorized, err
	}
	if session.PdvId != u.Pdv {
		return http.StatusUnauthorized, errors.New("")
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
