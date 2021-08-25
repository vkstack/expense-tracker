package repository

import "github.com/vkstack/expense-tracker/entities"

type IExpenseRepo interface {
	SaveExpense(expense *entities.Expense) (bool, error)
	GetExpense(expenseID string) *entities.Expense

	// GetExpensesFromUser(userID string) []*entities.Expense
	// GetExpensesFromGroup(groupID string) []*entities.Expense
}
