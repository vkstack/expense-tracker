package repository

import (
	"bytes"
	"encoding/gob"
	"sync"

	"github.com/vkstack/expense-tracker/entities"
	"github.com/vkstack/expense-tracker/interfaces/irepository"
)

type expenseRepository struct {
	mu sync.Mutex

	expenses map[string]*entities.Expense

	balances map[string]map[string]*entities.UserExpenseShare

	// userExpenses map[string][]*entities.Expense

	buff *bytes.Buffer
	enc  *gob.Encoder
	dec  *gob.Decoder
	// groupExpenses map[string][]*entities.Expense
}

func NewExpenseRepository() irepository.IExpenseRepo {
	buff := new(bytes.Buffer)
	return &expenseRepository{
		buff:     buff,
		enc:      gob.NewEncoder(buff),
		dec:      gob.NewDecoder(buff),
		expenses: make(map[string]*entities.Expense),
		balances: make(map[string]map[string]*entities.UserExpenseShare),
	}
}

func (expenseRepo *expenseRepository) SaveExpense(expense *entities.Expense) (bool, error) {
	expenseRepo.expenses[expense.GetID()] = expense
	return true, nil
}

func (expenseRepo *expenseRepository) GetExpense(expenseID string) *entities.Expense {
	return expenseRepo.expenses[expenseID]
}

func (expenseRepo *expenseRepository) PutBalances(expense *entities.Expense) {
	expenseRepo.mu.Lock()
	defer expenseRepo.mu.Unlock()
	userID := expense.GetSpender().GetID()
	if _, ok := expenseRepo.balances[userID]; !ok {
		expenseRepo.balances[userID] = make(map[string]*entities.UserExpenseShare)
	}
	for _, expShare := range expense.GetContribution() {
		benifitter, credit := expShare.GetShare()
		if _, ok := expenseRepo.balances[userID][benifitter]; !ok {
			expenseRepo.balances[userID][benifitter] = entities.NewUserExpenseShare(benifitter, credit)
		} else {
			expenseRepo.balances[userID][benifitter].AddCredit(credit)
		}
		if _, ok := expenseRepo.balances[benifitter]; !ok {
			expenseRepo.balances[benifitter] = make(map[string]*entities.UserExpenseShare)
		}
		if _, ok := expenseRepo.balances[benifitter][userID]; !ok {
			expenseRepo.balances[benifitter][userID] = entities.NewUserExpenseShare(userID, -credit)
		} else {
			expenseRepo.balances[benifitter][userID].AddCredit(-credit)
		}
	}
}

func (expenseRepo *expenseRepository) GetUserBalance(userID string) map[string]*entities.UserExpenseShare {
	expenseRepo.mu.Lock()
	defer expenseRepo.mu.Unlock()
	var resp map[string]*entities.UserExpenseShare
	expenseRepo.enc.Encode(expenseRepo.balances[userID])
	expenseRepo.dec.Decode(&resp)
	return resp
}

func (expenseRepo *expenseRepository) GetAllBalance() map[string]map[string]*entities.UserExpenseShare {
	expenseRepo.mu.Lock()
	defer expenseRepo.mu.Unlock()
	var resp map[string]map[string]*entities.UserExpenseShare
	expenseRepo.enc.Encode(expenseRepo.balances)
	expenseRepo.dec.Decode(&resp)
	return resp
}
