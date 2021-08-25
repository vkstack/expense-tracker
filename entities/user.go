package entities

import "github.com/rs/xid"

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
