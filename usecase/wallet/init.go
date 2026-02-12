package wallet

import (
	"mywallet/config"
	"mywallet/repository/transaction"
	"mywallet/repository/wallet"

	"gorm.io/gorm"
)

type WalletUsecase struct {
	cfg config.Config
	db  *gorm.DB
	w   wallet.WalletRepositoryItf
	t   transaction.TransactionRepositoryItf
}

func InitWalletUsecase(
	cfg config.Config,
	db *gorm.DB,
	walletRepository wallet.WalletRepository,
	transactionRepository transaction.TransactionRepository,
) *WalletUsecase {
	return &WalletUsecase{
		cfg: cfg,
		db:  db,
		w:   walletRepository,
		t:   transactionRepository,
	}
}
