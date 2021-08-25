package usecase

import "github.com/vkstack/expense-tracker/entities"

type IExpensifyUseCase interface {
	CreateUser(name, email, phone string) (bool, error)

	CreateGroup(name, descriptions string, memberIDs ...string)
	AddExpense(userID, groupID string, expenseShare entities.ExpenseShare)

	GetMyExpenses(userID string)
	GetMyBalance(userID string)
}
