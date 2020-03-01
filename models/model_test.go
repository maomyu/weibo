package models

import "testing"

func TestAddUser(t *testing.T) {
	user := User{
		Username: "test1",
		Password: "123",
		Sex:      1,
		Age:      23,
	}
	AddUser(user)
}
