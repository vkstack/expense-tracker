package entities

type Group struct {
	groupID      string
	name         string
	descriptions string
	members      map[string]*User
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

func NewGroup(id, name, email, phone string) *User {
	return &User{userID: id, name: name, email: email, phone: phone}
}
