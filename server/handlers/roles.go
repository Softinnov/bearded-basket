package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
)

func indexRoles(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	roles, err := models.GetRoles(c)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	utils.WriteJSON(w, roles)
	return http.StatusOK, nil
}
