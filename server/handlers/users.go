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

func getCurrentUser(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError {
	u, e := models.GetCurrentUser(c)
	if e != nil {
		return convHSt(e)
	}
	if e := WriteJSON(w, http.StatusOK, u); e != nil {
		return &utils.SError{http.StatusInternalServerError, nil, e}
	}
	return nil
}

func indexUsers(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError {
	us, e := models.GetUsersFromSession(c)
	if e != nil {
		return convHSt(e)
	}
	if e := WriteJSON(w, http.StatusOK, us); e != nil {
		return &utils.SError{http.StatusInternalServerError, nil, e}
	}
	return nil
}

func newUser(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError {
	u := models.User{}

	dec := json.NewDecoder(r.Body)
	if e := dec.Decode(&u); e != nil {
		return &utils.SError{http.StatusBadRequest, fmt.Errorf("champs utilisateur incorrects"), e}
	}
	defer r.Body.Close()

	id, e := models.CreateUser(c, &u)
	if e != nil {
		return convHSt(e)
	}
	u.Id = id
	u.Password = ""
	if e := WriteJSON(w, http.StatusCreated, u); e != nil {
		return &utils.SError{http.StatusInternalServerError, nil, e}
	}
	return nil
}

func editUser(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError {
	id, e := strconv.ParseInt(mux.Vars(r)["id"], 0, 64)
	if e != nil {
		return &utils.SError{http.StatusInternalServerError, nil, e}
	}
	u, err := models.GetUser(c, id)
	if err != nil {
		return convHSt(err)
	}
	up := models.User{}
	dec := json.NewDecoder(r.Body)
	if e = dec.Decode(&up); e != nil {
		return &utils.SError{http.StatusBadRequest,
			fmt.Errorf("champs utilisateur incorrects"),
			e,
		}
	}
	defer r.Body.Close()

	if err = u.UpdateUser(c, &up); err != nil {
		return convHSt(err)
	}
	return nil
}

func deleteUser(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError {
	id, e := strconv.ParseInt(mux.Vars(r)["id"], 0, 64)
	if e != nil {
		return &utils.SError{http.StatusInternalServerError, nil, e}
	}
	u, err := models.GetUser(c, id)
	if err != nil {
		return convHSt(err)
	}
	if err = u.RemoveUser(c); err != nil {
		return convHSt(err)
	}
	return nil
}
