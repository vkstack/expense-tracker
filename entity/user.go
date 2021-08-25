package entity

import (
	"github.com/vkstack/expense-tracker/utility"
)

type User struct {
	userID int    `binding:"required"`
	name   string `binding:"required"`
	email  string `binding:"required"`
	phone  string `binding:"required"`
}

func NewUser(id int, name, email, phone string) (user *User, err error) {
	user = &User{
		userID: id,
		name:   name,
		email:  email,
		phone:  phone,
	}
	if err := utility.Validator.Validate(user); err != nil {
		return nil, err
	}
	return user, nil
}
