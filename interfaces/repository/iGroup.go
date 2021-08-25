package repository

import (
	"github.com/vkstack/expense-tracker/entities"
)

type IGroupRepo interface {
	SaveGroup(group *entities.Group) (bool, error)
	AddMembers(groupID string, members ...*entities.User) (bool, error)
	RemoveMember(groupID string, userID string) (bool, error)

	GetGroup(groupID string) *entities.Group
	GetMembers(groupID string) []*entities.User
}
