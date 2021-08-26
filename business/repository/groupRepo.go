package repository

import (
	"errors"

	"github.com/vkstack/expense-tracker/entities"
	"github.com/vkstack/expense-tracker/interfaces/irepository"
)

var (
	errGroupNotFound = errors.New("group not found")
)

type groupRepository struct {
	groups map[string]*entities.Group
}

func NewGroupRepository() irepository.IGroupRepo {
	return &groupRepository{
		groups: make(map[string]*entities.Group),
	}
}

func (groupRepo *groupRepository) SaveGroup(group *entities.Group) (bool, error) {
	groupRepo.groups[group.GetID()] = group
	return true, nil
}

func (groupRepo *groupRepository) AddMembers(groupID string, members ...*entities.User) (bool, error) {
	if group, ok := groupRepo.groups[groupID]; ok {
		group.AddMembers(members...)
		return true, nil
	}
	return false, errGroupNotFound
}

func (groupRepo *groupRepository) RemoveMember(groupID string, userID string) (bool, error) {
	if group, ok := groupRepo.groups[groupID]; ok {
		group.RemoveMember(userID)
		return true, nil
	}
	return false, errGroupNotFound
}

func (groupRepo *groupRepository) GetGroup(groupID string) *entities.Group {
	return groupRepo.groups[groupID]
}

func (groupRepo *groupRepository) GetMembers(groupID string) []*entities.User {
	if group, ok := groupRepo.groups[groupID]; ok {
		return group.GetMembers()
	}
	return nil
}
