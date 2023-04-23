package dto

import (
	"RESTful/errs"
	"strings"
)

type TransactionRequest struct {
	AccountID       string  `json:"account_id"`
	CustomerID      string  `json:"customer_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (t TransactionRequest) IsWithdrawal() bool {
	return t.TransactionType == "withdrawal"
}

func (t TransactionRequest) Validate() *errs.AppError {
	lowerReq := strings.ToLower(t.TransactionType)
	if lowerReq != "deposit" && lowerReq != "withdrawal" {
		return errs.NewValidateError("transaction type must be deposit or withdrawal")
	}

	if t.Amount < 0 {
		return errs.NewValidateError("transaction amount cannot be less than 0")
	}

	return nil
}
