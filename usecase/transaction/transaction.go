package transaction

import (
	"mywallet/apperror"
	"mywallet/dto/request"
	"mywallet/dto/response"
	"mywallet/model"
	"mywallet/pkg/pagination"
	"mywallet/shared/constant"
	"mywallet/shared/utils"
	"time"

	"gorm.io/gorm"
)

func (uc *TransactionUsecase) Transfer(senderUserID uint, req request.TransferRequest) (*response.TransferResponse, error) {
	// Get receiver user by email
	receiverUser, err := uc.u.FindByEmail(req.ReceiverEmail)
	if err != nil {
		return nil, err
	}

	var newBalance float64
	var txID uint
	var senderWalletID, receiverWalletID uint
	var createdAt time.Time

	// Execute transfer in a database transaction (ACID)
	err = uc.db.Transaction(func(tx *gorm.DB) error {
		// Fetch both wallets with locks (prevents race conditions)
		senderWallet, err := uc.w.FindByUserIDWithLock(tx, senderUserID)
		if err != nil {
			return apperror.ErrWalletNotFound
		}
		senderWalletID = senderWallet.ID

		receiverWallet, err := uc.w.FindByUserIDWithLock(tx, receiverUser.ID)
		if err != nil {
			return apperror.ErrWalletNotFound
		}
		receiverWalletID = receiverWallet.ID

		// Validate transfer - business logic
		if req.Amount <= 0 {
			return apperror.ErrInvalidAmount
		}
		if senderWallet.ID == receiverWallet.ID {
			return apperror.ErrSelfTransfer
		}
		if senderWallet.Balance < req.Amount {
			return apperror.ErrInsufficientBalance
		}

		// Create transaction record
		txRecord := &model.Transaction{
			TransactionType:  string(constant.TransactionTypeTransfer),
			SenderWalletID:   &senderWallet.ID,
			ReceiverWalletID: receiverWallet.ID,
			Amount:           req.Amount,
			Status:           string(constant.TransactionStatusPending),
			Description:      req.Description,
		}

		// Save transaction
		if err := uc.t.CreateTx(tx, txRecord); err != nil {
			return err
		}
		txID = txRecord.ID
		createdAt = txRecord.CreatedAt

		// Execute transfer - update balances
		senderWallet.Balance -= req.Amount
		receiverWallet.Balance += req.Amount
		newBalance = senderWallet.Balance

		// Save both wallets
		if err := uc.w.UpdateTx(tx, senderWallet); err != nil {
			return err
		}
		if err := uc.w.UpdateTx(tx, receiverWallet); err != nil {
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

	return &response.TransferResponse{
		TransactionID:    txID,
		SenderWalletID:   senderWalletID,
		ReceiverWalletID: receiverWalletID,
		Amount:           req.Amount,
		NewBalance:       newBalance,
		CreatedAt:        createdAt,
		Status:           string(constant.TransactionStatusSuccess),
	}, nil
}

func (uc *TransactionUsecase) GetHistory(userID uint, page, limit int) ([]response.TransactionResponse, *response.PaginationMeta, error) {
	// Get user's wallet
	wallet, err := uc.w.GetWalletByUserID(userID)
	if err != nil {
		return nil, nil, apperror.ErrWalletNotFound
	}

	// Create pagination params
	paginationParams := pagination.NewPaginationParams(page, limit)

	// Get transactions
	transactions, total, err := uc.t.FindByWalletID(
		wallet.ID,
		paginationParams.Limit,
		paginationParams.Offset(),
	)
	if err != nil {
		return nil, nil, err
	}

	// Convert to response
	txResponses := utils.ModelTransactionsToResponse(transactions)

	// Build pagination metadata
	paginationMeta := &response.PaginationMeta{
		Page:       paginationParams.Page,
		Limit:      paginationParams.Limit,
		Total:      total,
		TotalPages: pagination.CalculateTotalPages(total, paginationParams.Limit),
	}

	return txResponses, paginationMeta, nil
}
