package entities

import (
	"fmt"

	"github.com/rs/xid"
)

type User struct {
	userID string
	name   string
	email  string
	phone  string
}

func (user *User) GetID() string {
	return user.userID
}

func NewUser(name, email, phone string) *User {
	return &User{userID: xid.New().String(), name: name, email: email, phone: phone}
}

func (user *User) UserCreationMessage() string {
	return fmt.Sprintf("UserID: %s\n Name: %s\n Email:%s\nPhone:%s\n", user.userID, user.name, user.email, user.phone)
}

func (user *User) String() string {
	return `----------------------------------------------------------------------------------------------\n` + fmt.Sprintf("User \t%s(%s)\n created\n", user.name, user.userID) + `----------------------------------------------------------------------------------------------\n`
}
