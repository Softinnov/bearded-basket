package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/utils"
)

type User struct {
	Id       int64     `json:"u_id,omitempty"`
	Pdv      int       `json:"u_pdv,omitempty"`
	Nom      string    `json:"u_nom,omitempty"`
	Prenom   string    `json:"u_prenom,omitempty"`
	Role     int8      `json:"u_role,omitempty"`
	Password string    `json:"u_pass,omitempty"`
	Login    string    `json:"u_login,omitempty"`
	Supprime int8      `json:"u_supprime,omitempty"`
	Cree     time.Time `json:"u_cree,omitempty"`
	Modifie  time.Time `json:"u_modifie,omitempty"`
	FaitPar  int64     `json:"u_faitpar,omitempty"`
}

func (u *User) Scan(v []*string, c []string) {
	const longForm = "2006-01-02 15:04:05"
	for k, val := range v {
		switch c[k] {
		case "u_id":
			i, e := strconv.ParseInt(*val, 10, 64)
			if e != nil {
				log.Printf("%s\n", e)
				continue
			}
			u.Id = i
		case "u_pdv":
			i, e := strconv.ParseInt(*val, 10, 32)
			if e != nil {
				log.Printf("%s\n", e)
				continue
			}
			u.Pdv = int(i)
		case "u_nom":
			if val != nil {
				u.Nom = *val
			}
		case "u_prenom":
			if val != nil {
				u.Prenom = *val
			}
		case "u_role":
			i, e := strconv.ParseInt(*val, 10, 8)
			if e != nil {
				log.Printf("%s\n", e)
				continue
			}
			u.Role = int8(i)
		case "u_pass":
			if val != nil {
				u.Prenom = *val
			}
		case "u_login":
			if val != nil {
				u.Login = *val
			}
		case "u_supprime":
			i, e := strconv.ParseInt(*val, 10, 8)
			if e != nil {
				log.Printf("%s\n", e)
				continue
			}
			u.Supprime = int8(i)
		case "u_cree":
			if val != nil {
				t, e := time.Parse(longForm, *val)
				if e != nil {
					fmt.Println(e)
				}
				u.Cree = t
			}
			fmt.Printf("u_cree: %v %v\n", u.Cree, *val)
		case "u_modifie":
			if val != nil {
				t, e := time.Parse(longForm, *val)
				if e != nil {
					fmt.Println(e)
				}
				u.Modifie = t
			}
			fmt.Printf("u_modifie: %v %v\n", u.Cree, *val)
		case "u_faitpar":
			i, e := strconv.ParseInt(*val, 10, 64)
			if e != nil {
				log.Printf("%s\n", e)
				continue
			}
			u.FaitPar = i
		}
	}
}

func GetUser(c *utils.Context, id int64) (*User, *utils.SError) {
	u := User{}
	req := fmt.Sprintf("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login, u_supprime, u_faitpar FROM utilisateur WHERE u_id=%d", id)
	res, e := c.HTTPdb.Query(req)
	if e != nil {
		return nil, &utils.SError{StatusBadRequest,
			fmt.Errorf("champs utilisateur incorrect"),
			fmt.Errorf("GetUser: %s\n", e),
		}
	}
	log.Printf("%v\n", res.Data)
	if len(res.Data) == 1 {
		u.Scan(res.Data[0], res.Columns)
	} else {
		return nil, &utils.SError{StatusBadRequest,
			fmt.Errorf("utilisateur incorrect"),
			fmt.Errorf("GetUser: User not found\n", e),
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

	req := fmt.Sprintf("SELECT u_id, u_pdv, u_nom, u_prenom, u_role, u_login FROM utilisateur WHERE u_pdv=%d AND u_supprime=0", c.Session.PdvId)
	res, e := c.HTTPdb.Query(req)
	if e != nil {
		return nil, &utils.SError{StatusBadRequest,
			fmt.Errorf("session incorrecte"),
			fmt.Errorf("GetUsersFromSession: %s", e),
		}
	}

	for _, re := range res.Data {
		var u User

		u.Scan(re, res.Columns)

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
	req := fmt.Sprintf("SELECT COUNT(*) FROM utilisateur WHERE u_login=%q AND u_supprime=0", u.Login)
	res, e := c.HTTPdb.Query(req)
	if e != nil {
		return 0, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("CreateUser: %s", e),
		}
	}

	var count int

	if len(res.Data) == 1 && len(res.Data[0]) == 1 {
		i, e := strconv.ParseInt(*res.Data[0][0], 10, 32)
		if e != nil {
			return 0, &utils.SError{StatusInternalServerError,
				nil,
				fmt.Errorf("CreateUser: %s", e),
			}
		}
		count = int(i)
	}

	if count > 0 {
		return 0, &utils.SError{StatusBadRequest,
			errors.New("login déjà existant"),
			fmt.Errorf("CreateUser: login already exists (%d)", count),
		}
	}

	now := time.Now()
	m, e := json.Marshal(struct {
		*User
		Pdv      int       `json:"u_pdv"`
		Password string    `json:"u_pass"`
		Supprime int8      `json:"u_supprime"`
		FaitPar  int64     `json:"u_faitpar"`
		Cree     time.Time `json:"u_cree"`
		Modifie  time.Time `json:"u_modifie"`
	}{
		User: u,

		Pdv:      c.Session.PdvId,
		Password: hashPassword(u.Password),
		Supprime: 0,
		FaitPar:  c.Session.Id,
		Cree:     now,
		Modifie:  now,
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
	req = fmt.Sprintf("INSERT INTO utilisateur SET %s", r)
	res, e = c.HTTPdb.Exec(req)
	if e != nil {
		return 0, &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("CreateUser: %s", e),
		}
	}

	rid := res.Infos["lastInsertId"]
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
		Prenom   string    `json:"u_prenom,omitempty"`
		Nom      string    `json:"u_nom,omitempty"`
		Role     int8      `json:"u_role,omitempty"`
		Password string    `json:"u_pass,omitempty"`
		Modifie  time.Time `json:"u_modifie,omitempty"`
	}{
		Prenom:   up.Prenom,
		Nom:      up.Nom,
		Password: hashPassword(up.Password),
		Role:     up.Role,
		Modifie:  time.Now(),
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
	req := fmt.Sprintf("UPDATE utilisateur SET %s WHERE u_id=%d", r, u.Id)
	_, e = c.HTTPdb.Exec(req)
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
	req := fmt.Sprintf("UPDATE utilisateur SET u_supprime=1 WHERE u_id=%d", u.Id)
	_, e := c.HTTPdb.Exec(req)
	if e != nil {
		return &utils.SError{StatusInternalServerError,
			nil,
			fmt.Errorf("RemoveUser: %s", e),
		}
	}
	return nil
}
