package wallet

import (
	"mywallet/apperror"
	"mywallet/model"

	"gorm.io/gorm"
)

type (
	WalletRepositoryItf interface {
		CreateWallet(userID uint) (*model.Wallet, error)
		GetWalletByUserID(userID uint) (*model.Wallet, error)
		ValidateTopUp(amount float64) error
		FindByUserIDWithLock(tx *gorm.DB, userID uint) (*model.Wallet, error)
		UpdateTx(tx *gorm.DB, wallet *model.Wallet) error
		// GetWalletByID(id uint) (*model.Wallet, error)
		// ValidateTransfer(senderWallet, receiverWallet *model.Wallet, amount float64) error
	}

	WalletRepository struct {
		resource WalletResourceItf
	}

	WalletResourceItf interface {
		create(wallet *model.Wallet) error
		findByUserID(userID uint) (*model.Wallet, error)
		findByID(id uint) (*model.Wallet, error)
		findByUserIDWithLock(tx *gorm.DB, userID uint) (*model.Wallet, error)
		updateTx(tx *gorm.DB, wallet *model.Wallet) error
		// Update(wallet *model.Wallet) error
		// UpdateWithOptimisticLock(wallet *model.Wallet, oldVersion int) (bool, error)
	}

	WalletResource struct {
		DB *gorm.DB
	}
)

func InitRepository(rsc WalletResourceItf) WalletRepository {
	return WalletRepository{
		resource: rsc,
	}
}

func (d WalletRepository) CreateWallet(userID uint) (*model.Wallet, error) {
	wallet := &model.Wallet{
		UserID:  userID,
		Balance: 0.0,
	}
	if err := d.resource.create(wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (d WalletRepository) GetWalletByUserID(userID uint) (*model.Wallet, error) {
	wallet, err := d.resource.findByUserID(userID)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (d WalletRepository) ValidateTopUp(amount float64) error {
	if amount <= 0 {
		return apperror.ErrInvalidAmount
	}
	return nil
}

func (d WalletRepository) FindByUserIDWithLock(tx *gorm.DB, userID uint) (*model.Wallet, error) {
	return d.resource.findByUserIDWithLock(tx, userID)
}

func (d WalletRepository) UpdateTx(tx *gorm.DB, wallet *model.Wallet) error {
	return d.resource.updateTx(tx, wallet)
}
