package repository

import (
	"sync"

	"github.com/vkstack/expense-tracker/entities"
	"github.com/vkstack/expense-tracker/interfaces/repository"
)

type UserRepo struct {
	mu    sync.Mutex
	users map[string]*entities.User
}

func (userRepo *UserRepo) SaveUser(user *entities.User) (bool, error) {
	userRepo.mu.Lock()
	defer userRepo.mu.Unlock()
	userRepo.users[user.GetID()] = user
	return true, nil
}

func (userRepo *UserRepo) GetUser(userID string) *entities.User {
	return userRepo.users[userID]
}
func (userRepo *UserRepo) GetUsers(userIDs ...string) (users []*entities.User) {
	for _, userID := range userIDs {
		users = append(users, userRepo.users[userID])
	}
	return users
}

func NewUserRepo() repository.IUserRepo {
	return &UserRepo{
		users: make(map[string]*entities.User),
	}
}
