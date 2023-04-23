package service

import (
	"RESTful/domain"
	"RESTful/dto"
	"RESTful/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
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

func (as DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsWithdrawal() {
		account, err := as.repo.FindBy(req.AccountID)
		if err != nil {
			return nil, err;
		}

		if !account.CanWithdrawal(req.Amount) {
			return nil, errs.NewValidateError("Insufficient balance in account")
		}
	}

	ts := domain.Transaction{
		AccountID: req.AccountID,
		Amount: req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, appErr := as.repo.SaveTransaction(ts)
	if appErr != nil {
		return nil, appErr
	}

	response := transaction.ToTransResponseDTO()
	return &response, nil
}