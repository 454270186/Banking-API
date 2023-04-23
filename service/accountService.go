package service

import (
	"RESTful/domain"
	"RESTful/dto"
	"RESTful/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepo
}

func NewAccountService(repository domain.AccountRepo) AccountService {
	return DefaultAccountService{repository}
}

func (as DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	a := domain.Account{
		CustomerID: req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Type: req.AccountType,
		Amount: req.Amount,
		Status: "True",
	}

	newAccount, err := as.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToAccountResponseDTO()

	return &response, nil
}