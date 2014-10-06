package handlers

import (
	"errors"
	"fmt"
	"log"
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

	role, err := strconv.Atoi(r.PostFormValue("u_role"))
	if err != nil {
		return http.StatusBadRequest, err
	}
	user := &models.User{
		Id:     id,
		Nom:    r.PostFormValue("u_nom"),
		Prenom: r.PostFormValue("u_prenom"),
		Role:   int8(role),
	}

	if s, err := checkPdvId(c, user.Id, r); err != nil {
		return s, err
	}
	if err = models.UpdateUser(c, user); err != nil {
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
		log.Println("UsersHandler:", err)
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, users)
	return http.StatusOK, nil
}
