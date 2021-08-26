package usecase

import (
	"fmt"

	"github.com/vkstack/expense-tracker/entities"
	"github.com/vkstack/expense-tracker/interfaces/iodevice"
	"github.com/vkstack/expense-tracker/interfaces/irepository"
	"github.com/vkstack/expense-tracker/interfaces/iusecase"
)

type ExpenseManagerUseCase struct {
	userRepo    irepository.IUserRepo
	groupRepo   irepository.IGroupRepo
	expenseRepo irepository.IExpenseRepo

	outputDevice iodevice.IOut
}

func NewExpenseManagerUseCase(userRepo irepository.IUserRepo, groupRepo irepository.IGroupRepo, expenseRepo irepository.IExpenseRepo, outputDevice iodevice.IOut) iusecase.IExpensifyUseCase {
	return &ExpenseManagerUseCase{
		userRepo:    userRepo,
		groupRepo:   groupRepo,
		expenseRepo: expenseRepo,

		outputDevice: outputDevice,
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

func (expensifyUC *ExpenseManagerUseCase) AddExpense(userID, groupID string, amount float64, expType entities.ExpenseType, distribution map[string]float64) {
	user := expensifyUC.userRepo.GetUser(userID)
	expense, err := entities.NewExpense(groupID, amount, user, expType, distribution)
	if err != nil {
		expensifyUC.outputDevice.Write("Error: " + err.Error())
		return
	}
	expensifyUC.expenseRepo.SaveExpense(expense)
	expensifyUC.expenseRepo.PutBalances(expense)
	expensifyUC.outputDevice.Write(fmt.Sprint("Successfully Saved expense:\n", expense))

	// group := expensifyUC.groupRepo.GetGroup(groupID)
}

func (expensifyUC *ExpenseManagerUseCase) GetMyExpenses(userID, groupID string) {
}

func (expensifyUC *ExpenseManagerUseCase) GetMyBalance(userID, groupID string) {

}

func (expensifyUC *ExpenseManagerUseCase) GetAllBalance(groupID string) {

}
