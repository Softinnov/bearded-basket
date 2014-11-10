package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
)

var MStatus = map[int]int{
	models.StatusMovedPermanently:    http.StatusMovedPermanently,
	models.StatusBadRequest:          http.StatusBadRequest,
	models.StatusUnauthorized:        http.StatusUnauthorized,
	models.StatusNotFound:            http.StatusNotFound,
	models.StatusInternalServerError: http.StatusInternalServerError,
}

func convHSt(e *utils.SError) *utils.SError {
	return &utils.SError{
		Status: MStatus[e.Status],
		Front:  e.Front,
		Back:   e.Back,
	}
}
