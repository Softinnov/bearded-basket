package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Softinnov/bearded-basket/server/database"
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

func GetUser(c *utils.Context, id int64) (*User, *utils.SError) {
	u := User{}

	e := c.DB.
		QueryRow("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login, u_supprime, u_faitpar FROM utilisateur WHERE u_id=?", id).
		Scan(&u.Id, &u.Pdv, &u.Nom, &u.Prenom, &u.Role, &u.Login, &u.Supprime, &u.FaitPar)
	if e != nil {
		return nil, &utils.SError{StatusBadRequest,
			fmt.Errorf("champs utilisateur incorrect"),
			fmt.Errorf("GetUser: %s\n", e),
		}
	}
	return &u, nil
}

func GetCurrentUser(c *utils.Context) (*User, *utils.SError) {
	u, e := GetUser(c, c.Session.Id)
	if e != nil {
		return nil, e
	}
	return u, nil
}

func GetUsersFromSession(c *utils.Context) ([]*User, *utils.SError) {
	us := make([]*User, 0)

	r, e := c.DB.
		Query("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login FROM utilisateur WHERE u_pdv=? AND u_supprime=0", c.Session.PdvId)
	if e != nil {
		return nil, &utils.SError{StatusBadRequest,
			fmt.Errorf("session incorrecte"),
			fmt.Errorf("GetUsersFromSession:", e),
		}
	}
	defer r.Close()

	for r.Next() {
		var u User
		e := r.Scan(&u.Id, &u.Pdv, &u.Nom, &u.Prenom, &u.Role, &u.Login)
		if e != nil {
			return nil, &utils.SError{StatusInternalServerError,
				nil,
				fmt.Errorf("GetUsersFromSession:", e),
			}
		}
		us = append(us, &u)
	}
	return us, nil
}

func hashPassword(p string) string {
	if p == "" {
		return ""
	}
	hasher := md5.New()
	hasher.Write([]byte(p))

	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateUser(c *utils.Context, u *User) (int64, *utils.SError) {
	if u.Password == "" || u.Prenom == "" || u.Nom == "" || u.Login == "" || u.Role >= c.Session.Role || u.Role < 1 {
		return 0, &utils.SError{StatusBadRequest,
			errors.New("champs incorrects"),
			errors.New("CreateUser: incorrect fields"),
		}
	}
	us, err := GetUsersFromSession(c)
	if err != nil {
		return 0, err
	}
	if len(us) >= 10 {
		return 0, &utils.SError{StatusBadRequest,
			errors.New("Nombre limite de 10 utilisateurs atteint"),
			fmt.Errorf("CreateUser: maximum users reached %d", len(us)),
		}
	}
	m, e := json.Marshal(struct {
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
	if e != nil {
		return 0, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("CreateUser: %s", e),
		}
	}
	r, e := database.BuildSqlSets(m)
	if e != nil {
		return 0, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("CreateUser: %s", e),
		}
	}
	req := fmt.Sprintf("INSERT INTO utilisateur SET %s", r)
	log.Println(req)
	stmt, e := c.DB.Prepare(req)
	if e != nil {
		return 0, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("CreateUser: %s", e),
		}
	}
	res, e := stmt.Exec()
	if e != nil {
		return 0, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("CreateUser: %s", e),
		}
	}
	rid, e := res.LastInsertId()
	if e != nil {
		return 0, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("CreateUser: %s", e),
		}
	}
	return rid, nil
}

func (u *User) UpdateUser(c *utils.Context, up *User) *utils.SError {
	if c.Session.PdvId != u.Pdv || c.Session.Role <= up.Role ||
		(c.Session.Id == u.Id && c.Session.Role != up.Role) {
		return &utils.SError{StatusBadRequest,
			fmt.Errorf("incorrect fields"),
			fmt.Errorf("UpdateUser: incorrect fields"),
		}
	}
	if up.Password != "" {
		u.Password = hashPassword(up.Password)
	}
	m, e := json.Marshal(struct {
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
	if e != nil {
		return &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("UpdateUser: %s", e),
		}
	}
	r, e := database.BuildSqlSets(m)
	if e != nil {
		return &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("UpdateUser: %s", e),
		}
	}
	req := fmt.Sprintf("UPDATE utilisateur SET %s WHERE u_id=%v", r, u.Id)
	stmt, e := c.DB.Prepare(req)
	if e != nil {
		return &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("UpdateUser: %s", e),
		}
	}
	_, e = stmt.Exec()
	if e != nil {
		return &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("UpdateUser: %s", e),
		}
	}
	return nil
}

func (u *User) RemoveUser(c *utils.Context) *utils.SError {
	if c.Session.PdvId != u.Pdv || c.Session.Role <= u.Role ||
		c.Session.Id == u.Id {
		return &utils.SError{StatusUnauthorized,
			fmt.Errorf("utilisateur incorrect"),
			fmt.Errorf("RemoveUser: PDV: %d != %d", c.Session.PdvId, u.Pdv),
		}
	}
	req := fmt.Sprintf("UPDATE utilisateur SET u_supprime=1 WHERE u_id=%v", u.Id)
	stmt, e := c.DB.Prepare(req)
	if e != nil {
		return &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("RemoveUser: %s", e),
		}
	}
	_, e = stmt.Exec()
	if e != nil {
		return &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("RemoveUser: %s", e),
		}
	}
	return nil
}
