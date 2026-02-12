package transaction

import (
	"mywallet/model"

	"gorm.io/gorm"
)

type (
	TransactionRepositoryItf interface {
		CreateTx(tx *gorm.DB, transaction *model.Transaction) error
		UpdateTx(tx *gorm.DB, transaction *model.Transaction) error
		FindByWalletID(walletID uint, limit, offset int) ([]model.Transaction, int64, error)
	}

	TransactionRepository struct {
		resource TransactionResourceItf
	}

	TransactionResourceItf interface {
		createTx(tx *gorm.DB, transaction *model.Transaction) error
		updateTx(tx *gorm.DB, transaction *model.Transaction) error
		findByWalletID(walletID uint, limit, offset int) ([]model.Transaction, int64, error)
	}

	TransactionResource struct {
		DB *gorm.DB
	}
)

func InitRepository(rsc TransactionResourceItf) TransactionRepository {
	return TransactionRepository{
		resource: rsc,
	}
}

func (d TransactionRepository) CreateTx(tx *gorm.DB, transaction *model.Transaction) error {
	return d.resource.createTx(tx, transaction)
}

func (d TransactionRepository) UpdateTx(tx *gorm.DB, transaction *model.Transaction) error {
	return d.resource.updateTx(tx, transaction)
}

func (d TransactionRepository) FindByWalletID(walletID uint, limit, offset int) ([]model.Transaction, int64, error) {
	return d.resource.findByWalletID(walletID, limit, offset)
}
