package entity

type ExpenseType int

const (
	ExpenseTypeExact ExpenseType = iota + 1
	ExpenseTypePercent
	ExpenseTypeEqual
)
