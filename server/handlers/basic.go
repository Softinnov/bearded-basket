package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type basicHandler struct {
	c  *utils.Context
	fn func(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error)
}

func (ah basicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := ah.fn(ah.c, w, r)
	utils.LogHTTP(w, err, status, r)
}
