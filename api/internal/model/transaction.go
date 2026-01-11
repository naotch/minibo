package model

import "gorm.io/gorm"

type TransactionCategory int

const (
	EXP TransactionCategory = iota
	INC
)

type Transaction struct {
	gorm.Model
	UserID   uint                `gorm:"not null;index"`
	Title    string              `gorm:"not null"`
	Category TransactionCategory `gorm:"not null"`
	Amount   int                 `gorm:"not null"`
}
