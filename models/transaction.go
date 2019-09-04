package models

import "time"

// Transaction represent model of transaction data
type Transaction struct {
	UserID        int       `json:"user_id"`
	TransactionID string    `json:"transaction_id"`
	Type          string    `json:"transaction_type"`
	Amount        int       `json:"amount"`
	Desc          string    `json:"desc"`
	CreatedAt     time.Time `json:"created_at"`
}
