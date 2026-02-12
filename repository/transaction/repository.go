package transaction

import (
	"mywallet/model"

	"gorm.io/gorm"
)

func (rsc TransactionResource) createTx(tx *gorm.DB, transaction *model.Transaction) error {
	return tx.Create(transaction).Error
}

func (rsc TransactionResource) updateTx(tx *gorm.DB, transaction *model.Transaction) error {
	return tx.Save(transaction).Error
}

func (rsc TransactionResource) findByWalletID(walletID uint, limit, offset int) ([]model.Transaction, int64, error) {
	var transactions []model.Transaction
	var total int64

	// Count total transactions
	countQuery := rsc.DB.Model(&model.Transaction{}).
		Where("(sender_wallet_id = ? OR receiver_wallet_id = ?)", walletID, walletID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated transactions
	err := rsc.DB.Where("(sender_wallet_id = ? OR receiver_wallet_id = ?)", walletID, walletID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error

	if err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}
