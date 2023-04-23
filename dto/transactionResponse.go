package dto

type TransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	NewBalance      float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
