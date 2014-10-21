package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/mux"
)

func getCurrentUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	user, err := models.GetCurrentUser(c)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, user)
	return http.StatusOK, nil
}

func indexUsers(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	us, err := models.GetUsersFromSession(c)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, us)
	return http.StatusOK, nil
}

func newUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	u := models.User{}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&u); err != nil {
		return http.StatusBadRequest, err
	}
	defer r.Body.Close()

	if _, err := models.CreateUser(c, &u); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusCreated, nil
}

func editUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 64)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	u, err := models.GetUser(c, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	up := models.User{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&up); err != nil {
		return http.StatusBadRequest, err
	}
	defer r.Body.Close()

	if err = u.UpdateUser(c, &up); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusAccepted, nil
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
	if err = u.RemoveUser(c); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	return http.StatusOK, nil
}
