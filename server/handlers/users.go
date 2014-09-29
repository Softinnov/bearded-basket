package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/mux"
)

func editUser(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return http.StatusInternalServerError, err
	}
	role, err := strconv.Atoi(r.PostFormValue("u_role"))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	user := &models.User{
		Id:     id,
		Nom:    r.PostFormValue("u_nom"),
		Prenom: r.PostFormValue("u_prenom"),
		Role:   int8(role),
	}
	fmt.Printf("%#v\n", user)
	if err = models.UpdateUser(c, user); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Fprint(w, "Success")
	u, _ := models.GetUser(c, id)
	fmt.Printf("%#v\n", u)
	return 0, nil
}

func indexUsers(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	session, err := models.GetFromCookies(c.Store, r)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	//check if pdvid is same in session and user context session
	users, err := models.GetUsersFromSession(c, session)
	if err != nil {
		log.Println("UsersHandler:", err)
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, users)
	return 0, nil
}
