package irepository

import "github.com/vkstack/expense-tracker/entities"

type IExpenseRepo interface {
	SaveExpense(expense *entities.Expense) (bool, error)
	GetExpense(expenseID string) *entities.Expense
	PutBalances(expense *entities.Expense)

	GetUserBalance(userID string) map[string]*entities.UserExpenseShare
	GetAllBalance() map[string]map[string]*entities.UserExpenseShare
}
