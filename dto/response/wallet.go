package response

import "time"

type WalletResponse struct {
	ID      uint    `json:"wallet_id"`
	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

type TopUpResponse struct {
	WalletID      uint      `json:"wallet_id"`
	NewBalance    float64   `json:"new_balance"`
	TransactionID uint      `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
}
