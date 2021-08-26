package entities

type UserExpenseShare struct {
	userID       string
	contribution float64
}

func NewUserExpenseShare(userID string, contribution float64) *UserExpenseShare {
	return &UserExpenseShare{userID: userID, contribution: contribution}
}

func (usertExpShare *UserExpenseShare) AddCredit(amount float64) {
	usertExpShare.contribution += amount
}

func (userExpShare *UserExpenseShare) GetShare() (string, float64) {
	return userExpShare.userID, userExpShare.contribution
}
