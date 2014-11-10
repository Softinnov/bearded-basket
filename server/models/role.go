package models

import (
	"fmt"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type Role struct {
	Id      int    `json:"r_id"`
	Libelle string `json:"r_libelle"`
}

func GetRole(c *utils.Context, id int) (*Role, *utils.SError) {
	r := Role{}

	e := c.DB.QueryRow("SELECT r_id, r_libelle FROM role WHERE r_id=?", id).
		Scan(&r.Id, &r.Libelle)
	if e != nil {
		return nil, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("GetRole: %s", e),
		}
	}
	return &r, nil
}

func GetRoles(c *utils.Context) ([]*Role, *utils.SError) {
	rs := make([]*Role, 0)

	rows, e := c.DB.Query("SELECT r_id, r_libelle FROM role")
	if e != nil {
		return nil, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("GetRoles: %s", e),
		}
	}
	defer rows.Close()
	for rows.Next() {
		var r Role

		e := rows.Scan(&r.Id, &r.Libelle)
		if e != nil {
			return nil, &utils.SError{StatusInternalServerError,
				nil,
				fmt.Errorf("GetRoles: %s", e),
			}
		}
		rs = append(rs, &r)
	}
	return rs, nil
}
