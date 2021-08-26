package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/xid"
)

type ExpenseType int

const (
	ExpenseTypeEqual ExpenseType = iota + 1
	ExpenseTypeExact
	ExpenseTypePercentage
)

var errInvalidExpenseType = errors.New("invalid expense type")

type Expense struct {
	id      string
	groupID string

	expensedAt time.Time
	paidBy     *User

	etype        ExpenseType
	amount       float64
	count        int
	distribution []*UserExpenseShare

	contribution []*UserExpenseShare
}

func (expense *Expense) GetContribution() []*UserExpenseShare {
	return expense.contribution
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

func (expense *Expense) eval() error {
	if expense.etype != ExpenseTypeEqual {
		vals := 0.0
		for _, val := range expense.contribution {
			vals += val.contribution
		}
		if (expense.etype == ExpenseTypeExact && vals != expense.amount) ||
			(expense.etype == ExpenseTypePercentage && vals != 100) {
			return errors.New("distributions is incorrect, Please check")
		}
	}
	if expense.etype == ExpenseTypeExact {
		expense.contribution = expense.distribution
		return nil
	}
	var amount float64 = expense.amount
	for _, userShare := range expense.distribution {
		var uexshare = &UserExpenseShare{userID: userShare.userID}
		if expense.etype == ExpenseTypeEqual {
			uexshare.contribution = expense.amount / float64(expense.count)
		} else {
			uexshare.contribution = (expense.amount * userShare.contribution) / 100.0
		}
		amount -= uexshare.contribution
		expense.contribution = append(expense.contribution, uexshare)
	}
	expense.contribution[0].contribution += amount
	return nil
}

func NewExpense(groupid string, amount float64, paidBy *User, etype ExpenseType, distribution map[string]float64) (*Expense, error) {
	expense := &Expense{
		id: xid.New().String(), groupID: groupid, amount: amount,
		paidBy: paidBy, expensedAt: time.Now(),
		etype: etype,
		count: len(distribution),
	}
	if etype >= ExpenseTypePercentage || distribution == nil {
		return nil, errInvalidExpenseType
	}
	for userID, share := range distribution {
		expense.distribution = append(expense.distribution, &UserExpenseShare{userID, share})
	}
	if err := expense.eval(); err != nil {
		return nil, err
	}
	return expense, nil
}

func (expense *Expense) String() string {
	return `----------------------------------------------------------------------------------------------\n` + fmt.Sprintf(`expenseID:\t%s\nspendBy:\t%s\namount:\t%f\ngroup:\t%s\n`, expense.id, expense.paidBy.userID, expense.amount, expense.groupID) + `----------------------------------------------------------------------------------------------\n`
}
