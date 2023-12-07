package main

import (
	"testing"
)

func Test_Firstname_Required(t *testing.T) {
	_, error := createUser("", "lastname", "msisdn")
	if error.Error() == "Firstname cannot be empty" {
		t.Log("Firstname field validation success")
	} else {
		t.Error("Firstname field validation failed")
	}
}

func Test_Lastname_Nil_Required(t *testing.T) {
	_, error := createUser("firstname", "", "msisdn")
	if error.Error() == "Lastname cannot be empty" {
		t.Log("Lastname field validation success")
	} else {
		t.Error("Lastname field validation failed")
	}
}

func Test_Msisdn_Nil_Required(t *testing.T) {
	_, error := createUser("firstname", "lastname", "")
	if error.Error() == "msisdin cannot be empty" {
		t.Log("Msisdn field validation success")
	} else {
		t.Error("Msisdn field validation failed")
	}
}

func Test_Create_User(t *testing.T) {
	_, error := createUser("firstname", "lastname", "12345678")
	if error.Error() == "" {
		t.Log("User created")
	} else {
		t.Error("User created")
	}
}
