package domain

import (
	"RESTful/dto"
	"RESTful/errs"
)

type Account struct {
	Id          string
	CustomerID  string
	OpeningDate string
	Type        string
	Amount      float64
	Status      string
}

type AccountRepo interface {
	Save(a Account) (*Account, *errs.AppError)
	FindBy(id string) (*Account, *errs.AppError)
	SaveTransaction(ts Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.Id}
}

func (a Account) CanWithdrawal(amount float64) bool {
	return a.Amount >= amount
}
