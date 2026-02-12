package transaction

import (
	"mywallet/repository/transaction"
	"mywallet/repository/user"
	"mywallet/repository/wallet"

	"gorm.io/gorm"
)

type TransactionUsecase struct {
	db *gorm.DB
	u  user.UserRepositoryItf
	w  wallet.WalletRepositoryItf
	t  transaction.TransactionRepositoryItf
}

func InitTransactionUsecase(
	db *gorm.DB,
	userRepo user.UserRepositoryItf,
	walletRepository wallet.WalletRepositoryItf,
	transactionRepository transaction.TransactionRepositoryItf,
) *TransactionUsecase {
	return &TransactionUsecase{
		db: db,
		u:  userRepo,
		w:  walletRepository,
		t:  transactionRepository,
	}
}
