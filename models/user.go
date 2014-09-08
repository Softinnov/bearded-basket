package models

import "github.com/Softinnov/bearded-basket/database"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func GetUser(id int) (*User, error) {
	var user User

	err := database.DB.QueryRow("SELECT id, name, gender FROM users WHERE id=?", id).Scan(&user.Id, &user.Name, &user.Gender)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers() ([]*User, error) {
	var users []*User

	rows, err := database.DB.Query("SELECT id, name, gender FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Gender)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
