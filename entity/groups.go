package entity

type Group struct {
	GroupID int
	Name,
	Description string
	Members map[int]*User
}
