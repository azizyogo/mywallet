package response

import "time"

type TransactionResponse struct {
	ID               uint      `json:"id"`
	Type             string    `json:"type"`
	Amount           float64   `json:"amount"`
	Description      string    `json:"description,omitempty"`
	SenderWalletID   *uint     `json:"sender_wallet_id,omitempty"`
	ReceiverWalletID uint      `json:"receiver_wallet_id"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

type TransferResponse struct {
	TransactionID    uint      `json:"transaction_id"`
	SenderWalletID   uint      `json:"sender_wallet_id"`
	ReceiverWalletID uint      `json:"receiver_wallet_id"`
	Amount           float64   `json:"amount"`
	NewBalance       float64   `json:"new_balance"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

// PaginationMeta contains pagination metadata
type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}
