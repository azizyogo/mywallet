package utils

import (
	"mywallet/dto/response"
	"mywallet/model"
	"strings"
)

// MaskEmail masks email address for privacy (e.g., j***@example.com)
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}
	local := parts[0]
	if len(local) <= 1 {
		return email
	}
	return string(local[0]) + "***@" + parts[1]
}

func ModelUserToResponse(user *model.User) response.UserResponse {
	return response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		// Email:     MaskEmail(user.Email), // use this if email masking is desired
		CreatedAt: user.CreatedAt,
	}
}

func ModelWalletToResponse(wallet *model.Wallet) response.WalletResponse {
	return response.WalletResponse{
		ID:      wallet.ID,
		UserID:  wallet.UserID,
		Balance: wallet.Balance,
	}
}

func ModelTransactionToResponse(tx *model.Transaction) response.TransactionResponse {
	return response.TransactionResponse{
		ID:               tx.ID,
		Type:             tx.TransactionType,
		Amount:           tx.Amount,
		Description:      tx.Description,
		SenderWalletID:   tx.SenderWalletID,
		ReceiverWalletID: tx.ReceiverWalletID,
		Status:           tx.Status,
		CreatedAt:        tx.CreatedAt,
	}
}

func ModelTransactionsToResponse(txs []model.Transaction) []response.TransactionResponse {
	result := make([]response.TransactionResponse, len(txs))
	for i, tx := range txs {
		result[i] = ModelTransactionToResponse(&tx)
	}
	return result
}
