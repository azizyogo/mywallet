package model

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint           `gorm:"unique;not null;index"`
	Balance   float64        `gorm:"type:decimal(19,2);default:0.00"`

	// Relations (use pointers to break circular dependencies)
	User                 *User          `gorm:"foreignKey:UserID"`
	SentTransactions     []*Transaction `gorm:"foreignKey:SenderWalletID"`
	ReceivedTransactions []*Transaction `gorm:"foreignKey:ReceiverWalletID"`
}

func (Wallet) TableName() string {
	return "wallets"
}
