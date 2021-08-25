package entities

import (
	"time"

	"github.com/rs/xid"
)

type Expense struct {
	id      string
	groupID string

	amount     float64
	expensedAt time.Time

	paidBy *User
}

func (expense *Expense) GetID() string {
	return expense.id
}

func (expense *Expense) GetGroupID() string {
	return expense.groupID
}

func (expense *Expense) GetSpender() *User {
	return expense.paidBy
}

func NewExpense(groupid string, amount float64, paidBy *User) *Expense {
	return &Expense{id: xid.New().String(), groupID: groupid, amount: amount, paidBy: paidBy, expensedAt: time.Now()}
}
