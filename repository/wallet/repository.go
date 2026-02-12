package wallet

import (
	"mywallet/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (rsc WalletResource) create(wallet *model.Wallet) error {
	return rsc.DB.Create(wallet).Error
}

func (rsc WalletResource) findByUserID(userID uint) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := rsc.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (rsc WalletResource) findByID(id uint) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := rsc.DB.Where("id = ?", id).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (rsc WalletResource) findByUserIDWithLock(tx *gorm.DB, userID uint) (*model.Wallet, error) {
	var wallet model.Wallet
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("user_id = ?", userID).
		First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (rsc WalletResource) updateTx(tx *gorm.DB, wallet *model.Wallet) error {
	return tx.Save(wallet).Error
}
