package irepository

import "github.com/vkstack/expense-tracker/entities"

type IUserRepo interface {
	SaveUser(user *entities.User) (bool, error)
	GetUser(userID string) *entities.User
	GetUsers(userIDs ...string) []*entities.User
}
