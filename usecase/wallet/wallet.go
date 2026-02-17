package wallet

import (
	"mywallet/apperror"
	"mywallet/dto/request"
	"mywallet/dto/response"
	"mywallet/model"
	"mywallet/shared/constant"
	"mywallet/shared/utils"
	"time"

	"gorm.io/gorm"
)

func (uc *WalletUsecase) GetBalance(userID uint) (*response.WalletResponse, error) {
	wallet, err := uc.w.GetWalletByUserID(userID)
	if err != nil {
		return nil, err
	}

	walletResp := utils.ModelWalletToResponse(wallet)
	return &walletResp, nil
}

func (uc *WalletUsecase) TopUp(userID uint, req request.TopUpRequest) (*response.TopUpResponse, error) {
	if err := uc.w.ValidateTopUp(req.Amount); err != nil {
		return nil, err
	}

	var (
		newBalance float64
		txID       uint
		walletID   uint
		createdAt  time.Time
	)
	// Execute all operations in a single database transaction
	// Auto-commits on success, auto-rollbacks on error
	err := uc.db.Transaction(func(tx *gorm.DB) error {
		// Lock wallet row for update (prevents race conditions)
		wallet, err := uc.w.FindByUserIDWithLock(tx, userID)
		if err != nil {
			return apperror.ErrWalletNotFound
		}
		walletID = wallet.ID

		// Create transaction record
		txRecord := &model.Transaction{
			TransactionType:  string(constant.TransactionTypeTopUp),
			ReceiverWalletID: wallet.ID,
			Amount:           req.Amount,
			Status:           string(constant.TransactionStatusPending),
			Description:      "Top up",
		}
		if err := uc.t.CreateTx(tx, txRecord); err != nil {
			return err
		}
		txID = txRecord.ID
		createdAt = txRecord.CreatedAt

		// Update wallet balance
		wallet.Balance += req.Amount
		newBalance = wallet.Balance

		if err := uc.w.UpdateTx(tx, wallet); err != nil {
			return err
		}

		// Mark transaction as success
		txRecord.Status = string(constant.TransactionStatusSuccess)
		if err := uc.t.UpdateTx(tx, txRecord); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &response.TopUpResponse{
		WalletID:      walletID,
		NewBalance:    newBalance,
		TransactionID: txID,
		CreatedAt:     createdAt,
	}, nil
}
