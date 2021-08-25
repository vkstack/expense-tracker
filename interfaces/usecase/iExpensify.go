package usecase

import "github.com/vkstack/expense-tracker/entities"

type IExpensifyUseCase interface {
	CreateUser(name, email, phone string) (bool, error)
}

type IUserUseCase interface {
	CreateGroup(name, descriptions string, memberIDs ...string)
	AddExpense(userID, groupID string, expenseShare entities.ExpenseShare, benificiaries ...string)

	GetMyExpenses(userID string)
	GetMyBalance(userID string)
}
