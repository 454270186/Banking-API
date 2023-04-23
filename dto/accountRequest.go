package dto

import (
	"RESTful/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

// Validate checks the validation of new account request
func (req NewAccountRequest) Validate() *errs.AppError {
	if req.Amount < 5000 {
		return errs.NewValidateError("need atleast 5000 to create a new account")
	}

	lowerAcType := strings.ToLower(req.AccountType)
	if lowerAcType != "checking" && lowerAcType != "saving" {
		return errs.NewValidateError("account type must be checking or saving")
	}

	return nil
}