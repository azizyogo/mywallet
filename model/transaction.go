package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID               uint      `gorm:"primaryKey"`
	CreatedAt        time.Time `gorm:"index"`
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	TransactionType  string         `gorm:"type:enum('TOPUP','TRANSFER');not null"`
	SenderWalletID   *uint          `gorm:"index"`
	ReceiverWalletID uint           `gorm:"not null;index"`
	Amount           float64        `gorm:"type:decimal(19,2);not null"`
	Status           string         `gorm:"type:enum('PENDING','SUCCESS','FAILED');default:'PENDING';index"`
	Description      string         `gorm:"type:varchar(500)"`

	// Relations (use pointers to avoid circular dependencies)
	SenderWallet   *Wallet `gorm:"foreignKey:SenderWalletID"`
	ReceiverWallet *Wallet `gorm:"foreignKey:ReceiverWalletID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
