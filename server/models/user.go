package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Softinnov/bearded-basket/server/utils"
)

type User struct {
	Id       int64  `json:"u_id,omitempty"`
	Pdv      int    `json:"u_pdv,omitempty"`
	Nom      string `json:"u_nom,omitempty"`
	Prenom   string `json:"u_prenom,omitempty"`
	Role     int8   `json:"u_role,omitempty"`
	Password string `json:"u_pass,omitempty"`
	Login    string `json:"u_login,omitempty"`
	Supprime int8   `json:"u_supprime,omitempty"`
	FaitPar  int64  `json:"u_faitpar,omitempty"`
}

func GetUser(c *utils.Context, id int64) (*User, error) {
	u := User{}

	err := c.DB.
		QueryRow("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login, u_supprime, u_faitpar FROM utilisateur WHERE u_id=?", id).
		Scan(&u.Id, &u.Pdv, &u.Nom, &u.Prenom, &u.Role, &u.Login, &u.Supprime, &u.FaitPar)
	if err != nil {
		log.Println("GetUser:", err)
		return nil, err
	}
	return &u, nil
}

func GetCurrentUser(c *utils.Context) (*User, error) {
	u, err := GetUser(c, c.Session.Id)
	if err != nil {
		log.Println("GetCurrentUser: GetUser call failed")
		return nil, err
	}
	return u, nil
}

func GetUsersFromSession(c *utils.Context) ([]*User, error) {
	users := make([]*User, 0)

	rows, err := c.DB.
		Query("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login FROM utilisateur WHERE u_pdv=? AND u_supprime=0", c.Session.PdvId)
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
	if p == "" {
		return ""
	}
	hasher := md5.New()
	hasher.Write([]byte(p))

	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateUser(c *utils.Context, u *User) (int64, error) {
	if u.Password == "" || u.Prenom == "" || u.Nom == "" || u.Login == "" || u.Role >= c.Session.Role || u.Role < 1 {
		log.Println("CreateUser: incorrect fields")
		return 0, errors.New("incorrect fields")
	}
	m, err := json.Marshal(struct {
		*User
		Pdv      int    `json:"u_pdv"`
		Password string `json:"u_pass"`
		Supprime int8   `json:"u_supprime"`
		FaitPar  int64  `json:"u_faitpar"`
	}{
		User: u,

		Pdv:      c.Session.PdvId,
		Password: hashPassword(u.Password),
		Supprime: 0,
		FaitPar:  c.Session.Id,
	})
	if err != nil {
		log.Println("CreateUser:", err)
		return 0, err
	}
	r, err := utils.BuildSqlSets(m)
	if err != nil {
		log.Println("CreateUser:", err)
		return 0, err
	}
	req := fmt.Sprintf("INSERT INTO utilisateur SET %s", r)
	log.Println(req)
	stmt, err := c.DB.Prepare(req)
	if err != nil {
		log.Println("CreateUser:", err)
		return 0, err
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Println("CreateUser:", err)
		return 0, err
	}
	rid, err := res.LastInsertId()
	if err != nil {
		log.Println("CreateUser:", err)
		return 0, err
	}
	return rid, nil
}

func (u *User) UpdateUser(c *utils.Context, up *User) error {
	if c.Session.PdvId != u.Pdv || c.Session.Role < up.Role ||
		(c.Session.Id == u.Id && c.Session.Role != up.Role) {
		return errors.New("UpdateUser: incorrect fields")
	}
	if up.Password != "" {
		u.Password = hashPassword(up.Password)
	}
	m, err := json.Marshal(struct {
		Prenom   string `json:"u_prenom,omitempty"`
		Nom      string `json:"u_nom,omitempty"`
		Role     int8   `json:"u_role,omitempty"`
		Password string `json:"u_pass,omitempty"`
	}{
		Prenom:   up.Prenom,
		Nom:      up.Nom,
		Password: hashPassword(up.Password),
		Role:     up.Role,
	})
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	r, err := utils.BuildSqlSets(m)
	if err != nil {
		log.Println("UpdateUser:", err)
		return err
	}
	req := fmt.Sprintf("UPDATE utilisateur SET %s WHERE u_id=%v", r, u.Id)
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

func (u *User) RemoveUser(c *utils.Context) error {
	if c.Session.PdvId != u.Pdv {
		return utils.NewMError(utils.Unauthorised, "RemoveUser: PDV: %d != %d", c.Session.PdvId, u.Pdv)
	}
	req := fmt.Sprintf("UPDATE utilisateur SET u_supprime=1 WHERE u_id=%v", u.Id)
	stmt, err := c.DB.Prepare(req)
	if err != nil {
		return utils.NewMError(utils.Internal, "RemoveUser: %s", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return utils.NewMError(utils.Internal, "RemoveUser: %s", err)
	}
	return nil
}
