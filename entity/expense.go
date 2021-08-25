package entity

import "time"

type Expense struct {
	expenseID string
	groupID   string
	userID    string

	title       string
	description string
	time        time.Time

	amount float64
	etype  ExpenseType
	// group
}
