package iusecase

import "github.com/vkstack/expense-tracker/entities"

type IExpensifyUseCase interface {
	CreateUser(name, email, phone string) (bool, error)

	CreateGroup(name, descriptions string, memberIDs ...string)

	AddExpense(userID, groupID string, amount float64, expType entities.ExpenseType, distribution map[string]float64)
	GetMyExpenses(userID, groupID string)
	GetMyBalance(userID, groupID string)
	GetAllBalance(groupID string)
}
