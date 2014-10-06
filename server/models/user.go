package models

import (
	"log"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type User struct {
	Id     int    `json:"u_id,omitempty"`
	Pdv    int    `json:"u_pdv,omitempty"`
	Nom    string `json:"u_nom,omitempty"`
	Prenom string `json:"u_prenom,omitempty"`
	Role   int8   `json:"u_role,omitempty"`
}

func GetUser(c *utils.Context, id int) (*User, error) {
	user := User{}

	err := c.DB.
		QueryRow("SELECT u_id, u_pdv, u_nom, u_prenom, u_role FROM utilisateur WHERE u_id=?", id).
		Scan(&user.Id, &user.Pdv, &user.Nom, &user.Prenom, &user.Role)
	if err != nil {
		log.Println("GetUser:", err)
		return nil, err
	}
	return &user, nil
}

func GetUsersFromSession(c *utils.Context, s *Session) ([]*User, error) {
	var users []*User

	rows, err := c.DB.
		Query("SELECT u_id, u_pdv, u_nom, u_prenom, u_role FROM utilisateur WHERE u_pdv=?", s.PdvId)
	if err != nil {
		log.Println("GetUsers:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Pdv, &user.Nom, &user.Prenom, &user.Role)
		if err != nil {
			log.Println("GetUsers:", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func UpdateUser(c *utils.Context, user *User) error {
	stmt, err := c.DB.Prepare("UPDATE utilisateur SET u_nom=?, u_prenom=?, u_role=? WHERE u_id=?")
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	_, err = stmt.Exec(user.Nom, user.Prenom, user.Role, user.Id)
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	return nil
}
