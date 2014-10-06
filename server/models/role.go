package models

import (
	"log"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type Role struct {
	Id      int    `json:"r_id"`
	Libelle string `json:"r_libelle"`
}

func GetRole(c *utils.Context, id int) (*Role, error) {
	role := Role{}

	err := c.DB.QueryRow("SELECT r_id, r_libelle FROM role WHERE r_id=?", id).
		Scan(&role.Id, &role.Libelle)
	if err != nil {
		log.Println("GetRole:", err)
		return nil, err
	}
	return &role, nil
}

func GetRoles(c *utils.Context) ([]*Role, error) {
	var roles []*Role

	rows, err := c.DB.Query("SELECT r_id, r_libelle FROM role")
	if err != nil {
		log.Println("GetRoles:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var role Role

		err := rows.Scan(&role.Id, &role.Libelle)
		if err != nil {
			log.Println("GetRoles:", err)
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}
