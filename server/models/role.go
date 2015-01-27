package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type Role struct {
	Id      int    `json:"r_id"`
	Libelle string `json:"r_libelle"`
}

func (r *Role) Scan(v []*string, c []string) {
	for k, val := range v {
		switch c[k] {
		case "r_id":
			i, e := strconv.Atoi(*val)
			if e != nil {
				log.Printf("%s\n", e)
				continue
			}
			r.Id = i
		case "r_libelle":
			if val != nil {
				r.Libelle = *val
			}
		}
	}
}

func GetRoles(c *utils.Context) ([]*Role, *utils.SError) {
	rs := make([]*Role, 0)

	res, e := c.HTTPdb.Query("SELECT r_id, r_libelle FROM role")
	if e != nil {
		return nil, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("GetRoles: %s", e),
		}
	}
	for _, re := range res.Data {
		var r Role

		r.Scan(re, res.Columns)

		rs = append(rs, &r)
	}
	return rs, nil
}
