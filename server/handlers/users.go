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
	session, err := models.GetFromCookies(c.Store, r)
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
		Prenom   string `json:"prenom"`
		Nom      string `json:"nom"`
		Role     int8   `json:"role"`
		Password string `json:"password"`
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
	if err = models.UpdateUser(c, u); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusAccepted, nil
}

func indexUsers(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	session, err := models.GetFromCookies(c.Store, r)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	users, err := models.GetUsersFromSession(c, session)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, users)
	return http.StatusOK, nil
}
