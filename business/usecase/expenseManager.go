package usecase

import (
	"github.com/vkstack/expense-tracker/entities"
	"github.com/vkstack/expense-tracker/interfaces/iodevice"
	"github.com/vkstack/expense-tracker/interfaces/repository"
	"github.com/vkstack/expense-tracker/interfaces/usecase"
)

type ExpenseManagerUseCase struct {
	userRepo    repository.IUserRepo
	groupRepo   repository.IGroupRepo
	expenseRepo repository.IExpenseRepo

	outputDevice iodevice.IOut
}

func NewExpenseManagerUseCase(userRepo repository.IUserRepo, groupRepo repository.IGroupRepo, expenseRepo repository.IExpenseRepo) usecase.IExpensifyUseCase {
	return &ExpenseManagerUseCase{
		userRepo:    userRepo,
		groupRepo:   groupRepo,
		expenseRepo: expenseRepo,
	}
}

func (expensifyUC *ExpenseManagerUseCase) CreateUser(name, email, phone string) (bool, error) {
	user := entities.NewUser(name, email, phone)
	expensifyUC.userRepo.SaveUser(user)
	expensifyUC.outputDevice.Write(user.UserCreationMessage())
	return true, nil
}

func (expensifyUC *ExpenseManagerUseCase) CreateGroup(name, descriptions string, memberIDs ...string) {
	members := expensifyUC.userRepo.GetUsers(memberIDs...)
	group := entities.NewGroup(name, descriptions, members...)
	expensifyUC.groupRepo.SaveGroup(group)
}

func (expensifyUC *ExpenseManagerUseCase) AddExpense(userID, groupID string, expenseShare entities.ExpenseShare) {

}

func (expensifyUC *ExpenseManagerUseCase) GetMyExpenses(userID string) {

}

func (expensifyUC *ExpenseManagerUseCase) GetMyBalance(userID string) {

}
