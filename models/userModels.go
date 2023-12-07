package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Msisdn    string `json:"msisdn"`
}

func (user *User) String() string {
	return fmt.Sprintf("id: %s, firstname: %s, lastname: %s, msisdn: %s", user.Id, user.Firstname, user.Lastname, user.Msisdn)
}

func createUser(id string, firstname string, lastname string, msisdn string) (User, error) {
	if firstname == "" {
		return User{}, errors.New("Firstname cannot be empty")
	}
	if lastname == "" {
		return User{}, errors.New("Lastname cannot be empty")
	}
	if msisdn == "" {
		return User{}, errors.New("msisdin cannot be empty")
	}
	user := User{id, firstname, lastname, msisdn}
	return user, errors.New("")
}
