package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type User struct {
	Id       int    `json:"id,omitempty"`
	Pdv      int    `json:"pdv,omitempty"`
	Nom      string `json:"nom,omitempty"`
	Prenom   string `json:"prenom,omitempty"`
	Role     int8   `json:"role,omitempty"`
	Password string `json:"password,omitempty"`
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

func buildSqlSets(b []byte) (string, error) {
	var buf bytes.Buffer
	var data map[string]interface{}

	json.Unmarshal(b, &data)
	flag := false
	for key, val := range data {
		if flag {
			buf.WriteString(", ")
		}
		buf.WriteString(key + "=" + fmt.Sprintf("%v", val))
		flag = true
	}
	return buf.String(), nil
}

func UpdateUser(c *utils.Context, user *User) error {
	m, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := buildSqlSets(m)
	if err != nil {
		return err
	}
	stmt, err := c.DB.Prepare(fmt.Sprintf("UPDATE utilisateur SET %s WHERE u_id=?", req))
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	_, err = stmt.Exec(user.Id)
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	return nil
}
