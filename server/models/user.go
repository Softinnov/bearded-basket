package models

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type User struct {
	Id       int    `json:"u_id,omitempty"`
	Pdv      int    `json:"u_pdv,omitempty"`
	Nom      string `json:"u_nom,omitempty"`
	Prenom   string `json:"u_prenom,omitempty"`
	Role     int8   `json:"u_role,omitempty"`
	Password string `json:"u_pass,omitempty"`
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

func CreateUser(c *utils.Context) (*User, error) {
	return nil, nil
}

func GetUsersFromSession(c *utils.Context, s *Session) ([]*User, error) {
	users := make([]*User, 0)

	rows, err := c.DB.
		Query("SELECT u_id, u_pdv, u_nom, u_prenom, u_role FROM utilisateur WHERE u_pdv=? AND u_supprime=0", s.PdvId)
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

	err := json.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}
	flag := false
	for key, val := range data {
		if flag {
			buf.WriteString(", ")
		}
		switch val.(type) {
		case string:
			buf.WriteString(key + "=" + fmt.Sprintf("%q", val))
		default:
			buf.WriteString(key + "=" + fmt.Sprintf("%v", val))
		}
		flag = true
	}
	return buf.String(), nil
}

func hashPassword(p string) string {
	hasher := md5.New()
	hasher.Write([]byte(p))

	return hex.EncodeToString(hasher.Sum(nil))
}

func UpdateUser(c *utils.Context, id int, user *User) error {
	if user.Password != "" {
		user.Password = hashPassword(user.Password)
	}
	m, err := json.Marshal(user)
	if err != nil {
		return err
	}
	r, err := buildSqlSets(m)
	if err != nil {
		return err
	}
	req := fmt.Sprintf("UPDATE utilisateur SET %s WHERE u_id=%v", r, id)
	stmt, err := c.DB.Prepare(req)
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	return nil
}

func RemoveUser(c *utils.Context, id int) error {
	req := fmt.Sprintf("UPDATE utilisateur SET u_supprime=1 WHERE u_id=%v", id)
	stmt, err := c.DB.Prepare(req)
	if err != nil {
		log.Println("RemoveUser:", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("RemoveUser:", err)
		return err
	}
	return nil
}
