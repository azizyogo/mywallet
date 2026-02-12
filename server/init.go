package server

import (
	"log"
	"mywallet/config"
	transactionRepo "mywallet/repository/transaction"
	userRepo "mywallet/repository/user"
	walletRepo "mywallet/repository/wallet"
	transactionUsecase "mywallet/usecase/transaction"
	userUsecase "mywallet/usecase/user"
	walletUsecase "mywallet/usecase/wallet"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	Cfg config.Config

	// Domain services
	userRepository        userRepo.UserRepository
	walletRepository      walletRepo.WalletRepository
	transactionRepository transactionRepo.TransactionRepository

	// Usecases
	UserUsecase        *userUsecase.UserUsecase
	WalletUsecase      *walletUsecase.WalletUsecase
	TransactionUsecase *transactionUsecase.TransactionUsecase
)

func Init(c config.Config) error {
	Cfg = c

	var err error
	db, err = initMySQL(Cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return err
	}

	initLayers(db, Cfg)

	return nil
}

func Close() {
	if db != nil {
		sqlDB, _ := db.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}
}

func initLayers(db *gorm.DB, cfg config.Config) {
	// initialize repositories
	userRepository = userRepo.InitRepository(&userRepo.UserResource{DB: db})
	walletRepository = walletRepo.InitRepository(&walletRepo.WalletResource{DB: db})
	transactionRepository = transactionRepo.InitRepository(&transactionRepo.TransactionResource{DB: db})

	// initialize usecases
	UserUsecase = userUsecase.InitUserUsecase(
		cfg,
		userRepository,
		walletRepository,
	)
	WalletUsecase = walletUsecase.InitWalletUsecase(
		cfg,
		db,
		walletRepository,
		transactionRepository,
	)
	TransactionUsecase = transactionUsecase.InitTransactionUsecase(
		db,
		userRepository,
		walletRepository,
		transactionRepository,
	)
}

func initMySQL(cfg config.Config) (*gorm.DB, error) {
	logMode := logger.Info
	if cfg.GinMode == "release" {
		logMode = logger.Silent
	}

	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
