package entities

import (
	"fmt"

	"github.com/rs/xid"
)

type Group struct {
	groupID     string
	name        string
	description string
	members     map[string]*User
}

func (group *Group) GetID() string {
	return group.groupID
}

func (group *Group) AddMembers(members ...*User) {
	for _, member := range members {
		group.members[member.GetID()] = member
	}
}

func (group *Group) GetMembers() (members []*User) {
	for _, member := range group.members {
		members = append(members, member)
	}
	return members
}

func (group *Group) RemoveMember(userID string) {
	delete(group.members, userID)
}

func NewGroup(name, description string, members ...*User) *Group {
	group := &Group{
		groupID:     xid.New().String(),
		name:        name,
		description: description,
	}
	for _, member := range members {
		group.members[member.userID] = member
	}
	return group
}

func (group *Group) CreationMessage() string {
	message := fmt.Sprintf("Created A group,\nGroupID: %s\nTitle:%s\nDescription:%s\n The memebers are:\n", group.groupID, group.name, group.description)
	idx := 1
	for _, user := range group.members {
		message += fmt.Sprintf("%d:\t%s (%s)\n", idx, user.name, user.userID)
		idx++
	}
	return message
}
