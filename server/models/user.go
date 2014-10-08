package models

import (
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
	Login    string `json:"u_login,omitempty"`
	Supprime *int8  `json:"u_supprime,omitempty"`
	FaitPar  int    `json:"u_faitpar,omitempty"`
}

func GetUser(c *utils.Context, id int) (*User, error) {
	user := User{}

	err := c.DB.
		QueryRow("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login FROM utilisateur WHERE u_id=?", id).
		Scan(&user.Id, &user.Pdv, &user.Nom, &user.Prenom, &user.Role, &user.Login)
	if err != nil {
		log.Println("GetUser:", err)
		return nil, err
	}
	return &user, nil
}

func GetCurrentUser(c *utils.Context, s *Session) (*User, error) {
	user := User{}

	err := c.DB.
		QueryRow("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login FROM utilisateur WHERE u_id=?", s.Id).
		Scan(&user.Id, &user.Pdv, &user.Nom, &user.Prenom, &user.Role, &user.Login)
	if err != nil {
		log.Println("GetCurrentUser:", err)
		return nil, err
	}
	return &user, nil
}

func CreateUser(c *utils.Context, user *User, s *Session) error {
	if user.Password != "" {
		user.Password = hashPassword(user.Password)
	}
	user.Pdv = s.PdvId
	user.Supprime = new(int8)
	user.FaitPar = s.Id
	fmt.Printf("%#v\n", user)
	m, err := json.Marshal(user)
	if err != nil {
		log.Println("CreateUser:", err)
		return err
	}
	fmt.Printf("%s\n", m)
	r, err := utils.BuildSqlSets(m)
	if err != nil {
		log.Println("CreateUser:", err)
		return err
	}
	fmt.Println(r)
	req := fmt.Sprintf("INSERT INTO utilisateur SET %s", r)
	stmt, err := c.DB.Prepare(req)
	if err != nil {
		log.Println("CreateUser:", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("CreateUser:", err)
		return err
	}
	return nil
}

func GetUsersFromSession(c *utils.Context, s *Session) ([]*User, error) {
	users := make([]*User, 0)

	rows, err := c.DB.
		Query("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login FROM utilisateur WHERE u_pdv=? AND u_supprime=0", s.PdvId)
	if err != nil {
		log.Println("GetUsersFromSession:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Pdv, &user.Nom, &user.Prenom, &user.Role, &user.Login)
		if err != nil {
			log.Println("GetUsersFromSession:", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func hashPassword(p string) string {
	hasher := md5.New()
	hasher.Write([]byte(p))

	return hex.EncodeToString(hasher.Sum(nil))
}

func UpdateUser(c *utils.Context, id int, user *User, session *Session) error {
	if user.Password != "" {
		user.Password = hashPassword(user.Password)
	}
	m, err := json.Marshal(user)
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	r, err := utils.BuildSqlSets(m)
	if err != nil {
		log.Println("UpdateUser:", err)
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
