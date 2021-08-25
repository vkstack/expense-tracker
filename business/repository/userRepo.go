package repository

import (
	"github.com/vkstack/expense-tracker/entities"
	"github.com/vkstack/expense-tracker/interfaces/repository"
)

type UserRepo struct {
	users map[string]*entities.User
}

func (userRepo *UserRepo) SaveUser(user *entities.User) (bool, error) {
	userRepo.users[user.GetID()] = user
	return true, nil
}

func (userRepo *UserRepo) GetUser(userID string) *entities.User {
	return userRepo.users[userID]
}

func NewUserRepo() repository.IUserRepo {
	return &UserRepo{
		users: make(map[string]*entities.User),
	}
}
