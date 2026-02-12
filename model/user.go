package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Email        string         `gorm:"unique;not null;index"`
	Name         string         `gorm:"not null"`
	PasswordHash string         `gorm:"not null"`

	// Relations (use pointer to break circular dependency)
	Wallet *Wallet `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
