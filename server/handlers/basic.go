package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type basicHandler struct {
	c  *utils.Context
	fn func(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError
}

func (ah basicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := ah.fn(ah.c, w, r)
	LogHTTP(e, w, r)
}
