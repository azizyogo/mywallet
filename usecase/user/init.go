package user

import (
	"mywallet/config"
	"mywallet/repository/user"
	"mywallet/repository/wallet"
)

type UserUsecase struct {
	cfg config.Config
	u   user.UserRepositoryItf
	w   wallet.WalletRepositoryItf
}

func InitUserUsecase(
	cfg config.Config,
	userRepository user.UserRepository,
	walletRepository wallet.WalletRepository,
) *UserUsecase {
	return &UserUsecase{
		cfg: cfg,
		u:   userRepository,
		w:   walletRepository,
	}
}
