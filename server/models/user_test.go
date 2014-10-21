package models

import (
	"testing"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/utils"
)

func TestGetUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	tes := []struct {
		name  string
		works bool
		user  *User
	}{
		{
			name:  "Working user",
			works: true,
			user: &User{
				Id:      1,
				Pdv:     0,
				Prenom:  "(super)",
				Nom:     "administrateur",
				Role:    9,
				Login:   "admin",
				FaitPar: 0,
			},
		},
		{
			name:  "No corresponding user",
			works: false,
			user:  &User{Id: 0},
		},
	}
	for _, te := range tes {

		u, err := GetUser(c, te.user.Id)

		if te.works {
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
			if u == nil {
				t.Errorf("Expected a user, got nil")
			} else if *te.user != *u {
				t.Errorf("Expected %#v, got %#v", te.user, u)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if u != nil {
				t.Errorf("Unexpected user, got %#v", u)
			}
		}
	}
}

func TestGetCurrentUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	tes := []struct {
		name    string
		works   bool
		user    *User
		session *utils.Session
	}{
		{
			name:  "Working user",
			works: true,
			user: &User{
				Id:      1,
				Pdv:     0,
				Prenom:  "(super)",
				Nom:     "administrateur",
				Role:    9,
				Login:   "admin",
				FaitPar: 0,
			},
			session: &utils.Session{Id: 1, PdvId: 0},
		},
		{
			name:    "Working user",
			works:   false,
			user:    nil,
			session: &utils.Session{Id: 0},
		},
	}
	for _, te := range tes {

		c.Session = te.session
		u, err := GetCurrentUser(c)

		if te.works {
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
			if u == nil {
				t.Errorf("Expected a user, got nil")
			} else if *te.user != *u {
				t.Errorf("Expected %#v, got %#v", te.user, u)
			}
		} else {
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if u != nil {
				t.Errorf("Unexpected user, got %#v", u)
			}
		}
	}
}

func TestGetUsersFromSession(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	tes := []struct {
		name    string
		works   bool
		users   []*User
		session *utils.Session
	}{
		{
			name:    "Working session",
			works:   true,
			users:   []*User{{}, {}, {}},
			session: &utils.Session{PdvId: 0},
		},
	}
	for _, te := range tes {

		c.Session = te.session
		us, err := GetUsersFromSession(c)

		if te.works {
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
			if len(te.users) != len(us) {
				t.Errorf("Expected %d users, got %d", len(te.users), len(us))
			}
		} else {
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if 0 != len(us) {
				t.Errorf("Unexpected %d user, got %d", 0, len(us))
			}
		}
	}
}

func TestCreateUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	tes := []struct {
		name    string
		works   bool
		user    *User
		session *utils.Session
	}{
		{
			name:  "Working User",
			works: true,
			user: &User{
				Pdv:      0,
				Nom:      "NomTest",
				Prenom:   "PrenomTest",
				Role:     3,
				Password: "coucou",
				Login:    "loginTest",
				FaitPar:  1,
			},
			session: &utils.Session{Id: 1, PdvId: 0, Role: 5},
		},
		{
			name:  "Empty Role User",
			works: false,
			user: &User{
				Pdv:      0,
				Nom:      "NomTest",
				Prenom:   "PrenomTest",
				Password: "coucou",
				Login:    "loginTest",
				FaitPar:  1,
			},
			session: &utils.Session{Id: 1, PdvId: 0, Role: 5},
		},
		{
			name:  "Same Role",
			works: false,
			user: &User{
				Pdv:      0,
				Nom:      "NomTest",
				Prenom:   "PrenomTest",
				Role:     6,
				Password: "coucou",
				Login:    "loginTest",
				FaitPar:  1,
			},
			session: &utils.Session{Id: 1, PdvId: 0, Role: 5},
		},
		{
			name:  "Empty password User",
			works: false,
			user: &User{
				Pdv:      0,
				Nom:      "NomTest",
				Prenom:   "PrenomTest",
				Role:     4,
				Password: "",
				Login:    "loginTest",
				FaitPar:  1,
			},
			session: &utils.Session{Id: 1, PdvId: 0, Role: 5},
		},
		{
			name:  "Empty field (Name) User",
			works: false,
			user: &User{
				Pdv:      0,
				Nom:      "",
				Prenom:   "PrenomTest",
				Role:     4,
				Password: "coucou",
				Login:    "loginTest",
				FaitPar:  1,
			},
			session: &utils.Session{Id: 1, PdvId: 0, Role: 5},
		},
	}

	for _, te := range tes {

		c.Session = te.session
		cid, err := CreateUser(c, te.user)

		if te.works {
			if err != nil {
				t.Fatalf("%s: Unexpected error: %s", te.name, err)
			}

			u, err := GetUser(c, cid)
			if err != nil {
				t.Fatalf("%s: Unexpected error: %s", te.name, err)
			}
			if u == nil {
				t.Fatalf("%s: Expected a user, got nil", te.name)
			}

			te.user.Id = cid
			te.user.Password = ""
			if *te.user != *u {
				t.Errorf("%s: Expected user %#v, got %#v", te.name, te.user, u)
			}
		} else {
			if err == nil {
				t.Fatalf("%s: Expected error got nil", te.name)
			}
		}
	}
}

func TestUpdateUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	tes := []struct {
		name    string
		works   bool
		changes *User
		wanted  *User
		session *utils.Session
	}{
		{
			name:  "Working User",
			works: true,
			changes: &User{
				Nom: "NiniTest",
			},
			wanted: &User{
				Id:      15,
				Pdv:     1,
				Nom:     "NiniTest",
				Prenom:  "Delphine",
				Role:    3,
				Login:   "Delphine",
				FaitPar: 2,
			},
			session: &utils.Session{PdvId: 1, Role: 3},
		},
		//{
		//	name:  "Higher Role",
		//	works: false,
		//	user: &User{
		//		Pdv:      0,
		//		Nom:      "NomTest",
		//		Prenom:   "PrenomTest",
		//		Role:     6,
		//		Password: "coucou",
		//		Login:    "loginTest",
		//		FaitPar:  1,
		//	},
		//	session: &utils.Session{
		//		Id:    1,
		//		PdvId: 0,
		//		Role:  5,
		//	},
		//},
	}

	for _, te := range tes {

		c.Session = te.session
		tu, err := GetUser(c, te.wanted.Id)
		err = tu.UpdateUser(c, te.wanted)

		if te.works {
			if err != nil {
				t.Fatalf("%s: Unexpected error: %s", te.name, err)
			}

			u, err := GetUser(c, te.wanted.Id)
			if err != nil {
				t.Fatalf("%s: Unexpected error: %s", te.name, err)
			}
			if u == nil {
				t.Fatalf("%s: Expected a user, got nil", te.name)
			}

			if *te.wanted != *u {
				t.Errorf("%s: Expected user %#v, got %#v", te.name, te.wanted, u)
			}
		} else {
			if err == nil {
				t.Fatalf("%s: Expected error got nil", te.name)
			}
		}
	}
}
