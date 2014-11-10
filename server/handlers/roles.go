package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
)

func indexRoles(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError {
	roles, e := models.GetRoles(c)
	if e != nil {
		return convHSt(e)
	}
	if e := WriteJSON(w, http.StatusOK, roles); e != nil {
		return &utils.SError{http.StatusInternalServerError, nil, e}
	}
	return nil
}
