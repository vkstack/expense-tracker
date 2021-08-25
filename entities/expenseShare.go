package entities

import "errors"

type ExpenseType int

const (
	ExpenseTypeEqual ExpenseType = iota + 1
	ExpenseTypeExact
	ExpenseTypePercentage
)

var errInvalidExpenseType = errors.New("invalid expense type")

type UserExpenseShare struct {
	userID       string
	contribution float64
}

type ExpenseShare struct {
	etype        ExpenseType
	amount       float64
	count        int
	distribution []UserExpenseShare

	contribution []UserExpenseShare
}

func (share *ExpenseShare) eval() error {
	if share.etype != ExpenseTypeEqual {
		vals := 0.0
		for _, val := range share.contribution {
			vals += val.contribution
		}
		if (share.etype == ExpenseTypeExact && vals != share.amount) ||
			(share.etype == ExpenseTypePercentage && vals != 100) {
			return errors.New("distributions is incorrect, Please check")
		}
	}
	if share.etype == ExpenseTypeExact {
		share.contribution = share.distribution
		return nil
	}
	var amount float64 = share.amount
	for _, userShare := range share.distribution {
		var uexshare = UserExpenseShare{userID: userShare.userID}
		if share.etype == ExpenseTypeEqual {
			uexshare.contribution = share.amount / float64(share.count)
		} else {
			uexshare.contribution = (share.amount * userShare.contribution) / 100.0
		}
		amount -= uexshare.contribution
		share.contribution = append(share.contribution, uexshare)
	}
	share.contribution[0].contribution += amount
	return nil
}

func NewExpenseShare(etype ExpenseType, distribution []UserExpenseShare) (*ExpenseShare, error) {
	if etype >= ExpenseTypePercentage || distribution == nil {
		return nil, errInvalidExpenseType
	}
	expShare := ExpenseShare{
		etype:        etype,
		distribution: distribution,
		count:        len(distribution),
	}
	if err := expShare.eval(); err != nil {
		return nil, err
	}
	return &expShare, nil
}
