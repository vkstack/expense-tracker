package repository

import (
	"github.com/vkstack/expense-tracker/entities"
	"github.com/vkstack/expense-tracker/interfaces/repository"
)

type expenseRepository struct {
	expenses map[string]*entities.Expense
}

func (expenseRepo *expenseRepository) SaveExpense(expense *entities.Expense) (bool, error) {
	expenseRepo.expenses[expense.GetID()] = expense
	return true, nil
}

func (expenseRepo *expenseRepository) GetExpense(expenseID string) *entities.Expense {
	return expenseRepo.expenses[expenseID]
}

func NewExpenseRepository() repository.IExpenseRepo {
	return &expenseRepository{
		expenses: make(map[string]*entities.Expense),
	}
}
